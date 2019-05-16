package uidallocator

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/security/uid"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Interface interface {
	Allocate(uid.Block) error
	AllocateNext() (uid.Block, error)
	Release(uid.Block) error
}

var (
	ErrFull            = errors.New("range is full")
	ErrNotInRange      = errors.New("provided UID range is not in the valid range")
	ErrAllocated       = errors.New("provided UID range is already allocated")
	ErrMismatchedRange = errors.New("the provided UID range does not match the current UID range")
)

type Allocator struct {
	r     *uid.Range
	alloc allocator.Interface
}

var _ Interface = &Allocator{}

func New(r *uid.Range, factory allocator.AllocatorFactory) *Allocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Allocator{r: r, alloc: factory(int(r.Size()), r.String())}
}
func NewInMemory(r *uid.Range) *Allocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	factory := func(max int, rangeSpec string) allocator.Interface {
		return allocator.NewContiguousAllocationMap(max, rangeSpec)
	}
	return New(r, factory)
}
func (r *Allocator) Free() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.alloc.Free()
}
func (r *Allocator) Allocate(block uid.Block) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(block)
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
func (r *Allocator) AllocateNext() (uid.Block, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	offset, ok, err := r.alloc.AllocateNext()
	if err != nil {
		return uid.Block{}, err
	}
	if !ok {
		return uid.Block{}, ErrFull
	}
	block, ok := r.r.BlockAt(uint32(offset))
	if !ok {
		return uid.Block{}, ErrNotInRange
	}
	return block, nil
}
func (r *Allocator) Release(block uid.Block) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(block)
	if !ok {
		return nil
	}
	return r.alloc.Release(int(offset))
}
func (r *Allocator) Has(block uid.Block) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(block)
	if !ok {
		return false
	}
	return r.alloc.Has(int(offset))
}
func (r *Allocator) Snapshot(dst *api.RangeAllocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	snapshottable, ok := r.alloc.(allocator.Snapshottable)
	if !ok {
		return fmt.Errorf("not a snapshottable allocator")
	}
	rangeString, data := snapshottable.Snapshot()
	dst.Range = rangeString
	dst.Data = data
	return nil
}
func (r *Allocator) Restore(into *uid.Range, data []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if into.String() != r.r.String() {
		return ErrMismatchedRange
	}
	snapshottable, ok := r.alloc.(allocator.Snapshottable)
	if !ok {
		return fmt.Errorf("not a snapshottable allocator")
	}
	return snapshottable.Restore(into.String(), data)
}
func (r *Allocator) contains(block uid.Block) (bool, uint32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.r.Offset(block)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
