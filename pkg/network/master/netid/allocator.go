package netid

import (
	"errors"
	goformat "fmt"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Interface interface {
	Allocate(uint32) error
	AllocateNext() (uint32, error)
	Release(uint32) error
	Has(uint32) bool
}

var (
	ErrFull       = errors.New("range is full")
	ErrNotInRange = errors.New("provided netid is not in the valid range")
	ErrAllocated  = errors.New("provided netid is already allocated")
)

type Allocator struct {
	netIDRange *NetIDRange
	alloc      allocator.Interface
}

var _ Interface = &Allocator{}

func New(r *NetIDRange, allocatorFactory allocator.AllocatorFactory) *Allocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Allocator{netIDRange: r, alloc: allocatorFactory(int(r.Size), r.String())}
}
func NewInMemory(r *NetIDRange) *Allocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return New(r, func(max int, rangeSpec string) allocator.Interface {
		return allocator.NewAllocationMap(max, rangeSpec)
	})
}
func (r *Allocator) Free() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.alloc.Free()
}
func (r *Allocator) Allocate(id uint32) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.netIDRange.Contains(id)
	if !ok {
		return ErrNotInRange
	}
	allocated, err := r.alloc.Allocate(int(offset))
	if err != nil {
		return err
	}
	if !allocated {
		return ErrAllocated
	}
	return nil
}
func (r *Allocator) AllocateNext() (uint32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	offset, ok, err := r.alloc.AllocateNext()
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, ErrFull
	}
	return r.netIDRange.Base + uint32(offset), nil
}
func (r *Allocator) Release(id uint32) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.netIDRange.Contains(id)
	if !ok {
		return nil
	}
	return r.alloc.Release(int(offset))
}
func (r *Allocator) Has(id uint32) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.netIDRange.Contains(id)
	if !ok {
		return false
	}
	return r.alloc.Has(int(offset))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
