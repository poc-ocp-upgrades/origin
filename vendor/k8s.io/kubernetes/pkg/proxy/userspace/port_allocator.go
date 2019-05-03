package userspace

import (
 "errors"
 "math/big"
 "math/rand"
 "sync"
 "time"
 "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/wait"
)

var (
 errPortRangeNoPortsRemaining = errors.New("port allocation failed; there are no remaining ports left to allocate in the accepted range")
)

type PortAllocator interface {
 AllocateNext() (int, error)
 Release(int)
}
type randomAllocator struct{}

func (r *randomAllocator) AllocateNext() (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return 0, nil
}
func (r *randomAllocator) Release(_ int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func newPortAllocator(r net.PortRange) PortAllocator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if r.Base == 0 {
  return &randomAllocator{}
 }
 return newPortRangeAllocator(r, true)
}

const (
 portsBufSize         = 16
 nextFreePortCooldown = 500 * time.Millisecond
 allocateNextTimeout  = 1 * time.Second
)

type rangeAllocator struct {
 net.PortRange
 ports chan int
 used  big.Int
 lock  sync.Mutex
 rand  *rand.Rand
}

func newPortRangeAllocator(r net.PortRange, autoFill bool) PortAllocator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if r.Base == 0 || r.Size == 0 {
  panic("illegal argument: may not specify an empty port range")
 }
 ra := &rangeAllocator{PortRange: r, ports: make(chan int, portsBufSize), rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
 if autoFill {
  go wait.Forever(func() {
   ra.fillPorts()
  }, nextFreePortCooldown)
 }
 return ra
}
func (r *rangeAllocator) fillPorts() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for {
  if !r.fillPortsOnce() {
   return
  }
 }
}
func (r *rangeAllocator) fillPortsOnce() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 port := r.nextFreePort()
 if port == -1 {
  return false
 }
 r.ports <- port
 return true
}
func (r *rangeAllocator) nextFreePort() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 j := r.rand.Intn(r.Size)
 if b := r.used.Bit(j); b == 0 {
  r.used.SetBit(&r.used, j, 1)
  return j + r.Base
 }
 for i := j + 1; i < r.Size; i++ {
  if b := r.used.Bit(i); b == 0 {
   r.used.SetBit(&r.used, i, 1)
   return i + r.Base
  }
 }
 for i := 0; i < j; i++ {
  if b := r.used.Bit(i); b == 0 {
   r.used.SetBit(&r.used, i, 1)
   return i + r.Base
  }
 }
 return -1
}
func (r *rangeAllocator) AllocateNext() (port int, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 select {
 case port = <-r.ports:
 case <-time.After(allocateNextTimeout):
  err = errPortRangeNoPortsRemaining
 }
 return
}
func (r *rangeAllocator) Release(port int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 port -= r.Base
 if port < 0 || port >= r.Size {
  return
 }
 r.lock.Lock()
 defer r.lock.Unlock()
 r.used.SetBit(&r.used, port, 0)
}
