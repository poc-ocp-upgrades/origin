package ipallocator

import (
	"errors"
	"fmt"
	goformat "fmt"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	"math/big"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Interface interface {
	Allocate(net.IP) error
	AllocateNext() (net.IP, error)
	Release(net.IP) error
	ForEach(func(net.IP))
	Has(ip net.IP) bool
}

var (
	ErrFull              = errors.New("range is full")
	ErrAllocated         = errors.New("provided IP is already allocated")
	ErrMismatchedNetwork = errors.New("the provided network does not match the current range")
)

type ErrNotInRange struct{ ValidRange string }

func (e *ErrNotInRange) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("provided IP is not in the valid range. The range of valid IPs is %s", e.ValidRange)
}

type Range struct {
	net   *net.IPNet
	base  *big.Int
	max   int
	alloc allocator.Interface
}

func NewAllocatorCIDRRange(cidr *net.IPNet, allocatorFactory allocator.AllocatorFactory) *Range {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	max := RangeSize(cidr)
	base := bigForIP(cidr.IP)
	rangeSpec := cidr.String()
	r := Range{net: cidr, base: base.Add(base, big.NewInt(1)), max: maximum(0, int(max-2))}
	r.alloc = allocatorFactory(r.max, rangeSpec)
	return &r
}
func NewCIDRRange(cidr *net.IPNet) *Range {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewAllocatorCIDRRange(cidr, func(max int, rangeSpec string) allocator.Interface {
		return allocator.NewAllocationMap(max, rangeSpec)
	})
}
func NewFromSnapshot(snap *api.RangeAllocation) (*Range, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, ipnet, err := net.ParseCIDR(snap.Range)
	if err != nil {
		return nil, err
	}
	r := NewCIDRRange(ipnet)
	if err := r.Restore(ipnet, snap.Data); err != nil {
		return nil, err
	}
	return r, nil
}
func maximum(a, b int) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a > b {
		return a
	}
	return b
}
func (r *Range) Free() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.alloc.Free()
}
func (r *Range) Used() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.max - r.alloc.Free()
}
func (r *Range) CIDR() net.IPNet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return *r.net
}
func (r *Range) Allocate(ip net.IP) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(ip)
	if !ok {
		return &ErrNotInRange{r.net.String()}
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
func (r *Range) AllocateNext() (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	offset, ok, err := r.alloc.AllocateNext()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFull
	}
	return addIPOffset(r.base, offset), nil
}
func (r *Range) Release(ip net.IP) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(ip)
	if !ok {
		return nil
	}
	return r.alloc.Release(offset)
}
func (r *Range) ForEach(fn func(net.IP)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.alloc.ForEach(func(offset int) {
		ip, _ := GetIndexedIP(r.net, offset+1)
		fn(ip)
	})
}
func (r *Range) Has(ip net.IP) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ok, offset := r.contains(ip)
	if !ok {
		return false
	}
	return r.alloc.Has(offset)
}
func (r *Range) Snapshot(dst *api.RangeAllocation) error {
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
func (r *Range) Restore(net *net.IPNet, data []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !net.IP.Equal(r.net.IP) || net.Mask.String() != r.net.Mask.String() {
		return ErrMismatchedNetwork
	}
	snapshottable, ok := r.alloc.(allocator.Snapshottable)
	if !ok {
		return fmt.Errorf("not a snapshottable allocator")
	}
	snapshottable.Restore(net.String(), data)
	return nil
}
func (r *Range) contains(ip net.IP) (bool, int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !r.net.Contains(ip) {
		return false, 0
	}
	offset := calculateIPOffset(r.base, ip)
	if offset < 0 || offset >= r.max {
		return false, 0
	}
	return true, offset
}
func bigForIP(ip net.IP) *big.Int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b := ip.To4()
	if b == nil {
		b = ip.To16()
	}
	return big.NewInt(0).SetBytes(b)
}
func addIPOffset(base *big.Int, offset int) net.IP {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return net.IP(big.NewInt(0).Add(base, big.NewInt(int64(offset))).Bytes())
}
func calculateIPOffset(base *big.Int, ip net.IP) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return int(big.NewInt(0).Sub(bigForIP(ip), base).Int64())
}
func RangeSize(subnet *net.IPNet) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ones, bits := subnet.Mask.Size()
	if bits == 32 && (bits-ones) >= 31 || bits == 128 && (bits-ones) >= 127 {
		return 0
	}
	if bits == 128 && (bits-ones) >= 16 {
		return int64(1) << uint(16)
	} else {
		return int64(1) << uint(bits-ones)
	}
}
func GetIndexedIP(subnet *net.IPNet, index int) (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ip := addIPOffset(bigForIP(subnet.IP), index)
	if !subnet.Contains(ip) {
		return nil, fmt.Errorf("can't generate IP with index %d from subnet. subnet too small. subnet: %q", index, subnet)
	}
	return ip, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
