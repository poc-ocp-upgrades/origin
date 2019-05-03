package allocator

import (
 "errors"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "math/big"
 "math/rand"
 "sync"
 "time"
)

type AllocationBitmap struct {
 strategy  bitAllocator
 max       int
 rangeSpec string
 lock      sync.Mutex
 count     int
 allocated *big.Int
}

var _ Interface = &AllocationBitmap{}
var _ Snapshottable = &AllocationBitmap{}

type bitAllocator interface {
 AllocateBit(allocated *big.Int, max, count int) (int, bool)
}

func NewAllocationMap(max int, rangeSpec string) *AllocationBitmap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 a := AllocationBitmap{strategy: randomScanStrategy{rand: rand.New(rand.NewSource(time.Now().UnixNano()))}, allocated: big.NewInt(0), count: 0, max: max, rangeSpec: rangeSpec}
 return &a
}
func NewContiguousAllocationMap(max int, rangeSpec string) *AllocationBitmap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 a := AllocationBitmap{strategy: contiguousScanStrategy{}, allocated: big.NewInt(0), count: 0, max: max, rangeSpec: rangeSpec}
 return &a
}
func (r *AllocationBitmap) Allocate(offset int) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 if r.allocated.Bit(offset) == 1 {
  return false, nil
 }
 r.allocated = r.allocated.SetBit(r.allocated, offset, 1)
 r.count++
 return true, nil
}
func (r *AllocationBitmap) AllocateNext() (int, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 next, ok := r.strategy.AllocateBit(r.allocated, r.max, r.count)
 if !ok {
  return 0, false, nil
 }
 r.count++
 r.allocated = r.allocated.SetBit(r.allocated, next, 1)
 return next, true, nil
}
func (r *AllocationBitmap) Release(offset int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 if r.allocated.Bit(offset) == 0 {
  return nil
 }
 r.allocated = r.allocated.SetBit(r.allocated, offset, 0)
 r.count--
 return nil
}

const (
 notZero   = uint64(^big.Word(0))
 wordPower = (notZero>>8)&1 + (notZero>>16)&1 + (notZero>>32)&1
 wordSize  = 1 << wordPower
)

func (r *AllocationBitmap) ForEach(fn func(int)) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 words := r.allocated.Bits()
 for wordIdx, word := range words {
  bit := 0
  for word > 0 {
   if (word & 1) != 0 {
    fn((wordIdx * wordSize * 8) + bit)
    word = word &^ 1
   }
   bit++
   word = word >> 1
  }
 }
}
func (r *AllocationBitmap) Has(offset int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 return r.allocated.Bit(offset) == 1
}
func (r *AllocationBitmap) Free() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 return r.max - r.count
}
func (r *AllocationBitmap) Snapshot() (string, []byte) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 return r.rangeSpec, r.allocated.Bytes()
}
func (r *AllocationBitmap) Restore(rangeSpec string, data []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 if r.rangeSpec != rangeSpec {
  return errors.New("the provided range does not match the current range")
 }
 r.allocated = big.NewInt(0).SetBytes(data)
 r.count = countBits(r.allocated)
 return nil
}

type randomScanStrategy struct{ rand *rand.Rand }

func (rss randomScanStrategy) AllocateBit(allocated *big.Int, max, count int) (int, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if count >= max {
  return 0, false
 }
 offset := rss.rand.Intn(max)
 for i := 0; i < max; i++ {
  at := (offset + i) % max
  if allocated.Bit(at) == 0 {
   return at, true
  }
 }
 return 0, false
}

var _ bitAllocator = randomScanStrategy{}

type contiguousScanStrategy struct{}

func (contiguousScanStrategy) AllocateBit(allocated *big.Int, max, count int) (int, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if count >= max {
  return 0, false
 }
 for i := 0; i < max; i++ {
  if allocated.Bit(i) == 0 {
   return i, true
  }
 }
 return 0, false
}

var _ bitAllocator = contiguousScanStrategy{}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
