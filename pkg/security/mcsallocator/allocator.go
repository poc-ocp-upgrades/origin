package mcsallocator

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/security/mcs"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Interface interface {
	Allocate(*mcs.Label) error
	AllocateNext() (*mcs.Label, error)
	Release(*mcs.Label) error
}

var (
	ErrFull            = errors.New("range is full")
	ErrNotInRange      = errors.New("provided label is not in the valid range")
	ErrAllocated       = errors.New("provided label is already allocated")
	ErrMismatchedRange = errors.New("the provided label does not match the current label range")
)

type Allocator struct {
	r     *mcs.Range
	alloc allocator.Interface
}

var _ Interface = &Allocator{}

func New(r *mcs.Range, factory allocator.AllocatorFactory) *Allocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Allocator{r: r, alloc: factory(int(r.Size()), r.String())}
}
func NewInMemory(r *mcs.Range) *Allocator {
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
func (r *Allocator) Allocate(label *mcs.Label) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(label)
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
func (r *Allocator) AllocateNext() (*mcs.Label, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	offset, ok, err := r.alloc.AllocateNext()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFull
	}
	label, ok := r.r.LabelAt(uint64(offset))
	if !ok {
		return nil, ErrNotInRange
	}
	return label, nil
}
func (r *Allocator) Release(label *mcs.Label) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(label)
	if !ok {
		return nil
	}
	return r.alloc.Release(int(offset))
}
func (r *Allocator) Has(label *mcs.Label) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(label)
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
func (r *Allocator) Restore(into *mcs.Range, data []byte) error {
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
func (r *Allocator) contains(label *mcs.Label) (bool, uint64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.r.Offset(label)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
