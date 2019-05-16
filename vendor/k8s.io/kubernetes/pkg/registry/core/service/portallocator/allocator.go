package portallocator

import (
	"errors"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/net"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Interface interface {
	Allocate(int) error
	AllocateNext() (int, error)
	Release(int) error
	ForEach(func(int))
	Has(int) bool
}

var (
	ErrFull              = errors.New("range is full")
	ErrAllocated         = errors.New("provided port is already allocated")
	ErrMismatchedNetwork = errors.New("the provided port range does not match the current port range")
)

type ErrNotInRange struct{ ValidPorts string }

func (e *ErrNotInRange) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("provided port is not in the valid range. The range of valid ports is %s", e.ValidPorts)
}

type PortAllocator struct {
	portRange net.PortRange
	alloc     allocator.Interface
}

var _ Interface = &PortAllocator{}

func NewPortAllocatorCustom(pr net.PortRange, allocatorFactory allocator.AllocatorFactory) *PortAllocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	max := pr.Size
	rangeSpec := pr.String()
	a := &PortAllocator{portRange: pr}
	a.alloc = allocatorFactory(max, rangeSpec)
	return a
}
func NewPortAllocator(pr net.PortRange) *PortAllocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewPortAllocatorCustom(pr, func(max int, rangeSpec string) allocator.Interface {
		return allocator.NewAllocationMap(max, rangeSpec)
	})
}
func NewFromSnapshot(snap *api.RangeAllocation) (*PortAllocator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pr, err := net.ParsePortRange(snap.Range)
	if err != nil {
		return nil, err
	}
	r := NewPortAllocator(*pr)
	if err := r.Restore(*pr, snap.Data); err != nil {
		return nil, err
	}
	return r, nil
}
func (r *PortAllocator) Free() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.alloc.Free()
}
func (r *PortAllocator) Used() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.portRange.Size - r.alloc.Free()
}
func (r *PortAllocator) Allocate(port int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(port)
	if !ok {
		validPorts := r.portRange.String()
		return &ErrNotInRange{validPorts}
	}
	allocated, err := r.alloc.Allocate(offset)
	if err != nil {
		return err
	}
	if !allocated {
		return ErrAllocated
	}
	return nil
}
func (r *PortAllocator) AllocateNext() (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	offset, ok, err := r.alloc.AllocateNext()
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, ErrFull
	}
	return r.portRange.Base + offset, nil
}
func (r *PortAllocator) ForEach(fn func(int)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.alloc.ForEach(func(offset int) {
		fn(r.portRange.Base + offset)
	})
}
func (r *PortAllocator) Release(port int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(port)
	if !ok {
		klog.Warningf("port is not in the range when release it. port: %v", port)
		return nil
	}
	return r.alloc.Release(offset)
}
func (r *PortAllocator) Has(port int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(port)
	if !ok {
		return false
	}
	return r.alloc.Has(offset)
}
func (r *PortAllocator) Snapshot(dst *api.RangeAllocation) error {
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
func (r *PortAllocator) Restore(pr net.PortRange, data []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pr.String() != r.portRange.String() {
		return ErrMismatchedNetwork
	}
	snapshottable, ok := r.alloc.(allocator.Snapshottable)
	if !ok {
		return fmt.Errorf("not a snapshottable allocator")
	}
	return snapshottable.Restore(pr.String(), data)
}
func (r *PortAllocator) contains(port int) (bool, int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !r.portRange.Contains(port) {
		return false, 0
	}
	offset := port - r.portRange.Base
	return true, offset
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
