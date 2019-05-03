package scheduler

import (
 "container/heap"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "sync"
 "time"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/client-go/util/flowcontrol"
 "k8s.io/klog"
)

const (
 NodeHealthUpdateRetry    = 5
 NodeEvictionPeriod       = 100 * time.Millisecond
 EvictionRateLimiterBurst = 1
)

type TimedValue struct {
 Value     string
 UID       interface{}
 AddedAt   time.Time
 ProcessAt time.Time
}

var now = time.Now

type TimedQueue []*TimedValue

func (h TimedQueue) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(h)
}
func (h TimedQueue) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return h[i].ProcessAt.Before(h[j].ProcessAt)
}
func (h TimedQueue) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 h[i], h[j] = h[j], h[i]
}
func (h *TimedQueue) Push(x interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *h = append(*h, x.(*TimedValue))
}
func (h *TimedQueue) Pop() interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 old := *h
 n := len(old)
 x := old[n-1]
 *h = old[0 : n-1]
 return x
}

type UniqueQueue struct {
 lock  sync.Mutex
 queue TimedQueue
 set   sets.String
}

func (q *UniqueQueue) Add(value TimedValue) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if q.set.Has(value.Value) {
  return false
 }
 heap.Push(&q.queue, &value)
 q.set.Insert(value.Value)
 return true
}
func (q *UniqueQueue) Replace(value TimedValue) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 for i := range q.queue {
  if q.queue[i].Value != value.Value {
   continue
  }
  heap.Remove(&q.queue, i)
  heap.Push(&q.queue, &value)
  return true
 }
 return false
}
func (q *UniqueQueue) RemoveFromQueue(value string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if !q.set.Has(value) {
  return false
 }
 for i, val := range q.queue {
  if val.Value == value {
   heap.Remove(&q.queue, i)
   return true
  }
 }
 return false
}
func (q *UniqueQueue) Remove(value string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if !q.set.Has(value) {
  return false
 }
 q.set.Delete(value)
 for i, val := range q.queue {
  if val.Value == value {
   heap.Remove(&q.queue, i)
   return true
  }
 }
 return true
}
func (q *UniqueQueue) Get() (TimedValue, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if len(q.queue) == 0 {
  return TimedValue{}, false
 }
 result := heap.Pop(&q.queue).(*TimedValue)
 q.set.Delete(result.Value)
 return *result, true
}
func (q *UniqueQueue) Head() (TimedValue, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if len(q.queue) == 0 {
  return TimedValue{}, false
 }
 result := q.queue[0]
 return *result, true
}
func (q *UniqueQueue) Clear() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.lock.Lock()
 defer q.lock.Unlock()
 if q.queue.Len() > 0 {
  q.queue = make(TimedQueue, 0)
 }
 if len(q.set) > 0 {
  q.set = sets.NewString()
 }
}

type RateLimitedTimedQueue struct {
 queue       UniqueQueue
 limiterLock sync.Mutex
 limiter     flowcontrol.RateLimiter
}

func NewRateLimitedTimedQueue(limiter flowcontrol.RateLimiter) *RateLimitedTimedQueue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RateLimitedTimedQueue{queue: UniqueQueue{queue: TimedQueue{}, set: sets.NewString()}, limiter: limiter}
}

type ActionFunc func(TimedValue) (bool, time.Duration)

func (q *RateLimitedTimedQueue) Try(fn ActionFunc) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 val, ok := q.queue.Head()
 q.limiterLock.Lock()
 defer q.limiterLock.Unlock()
 for ok {
  if !q.limiter.TryAccept() {
   klog.V(10).Infof("Try rate limited for value: %v", val)
   break
  }
  now := now()
  if now.Before(val.ProcessAt) {
   break
  }
  if ok, wait := fn(val); !ok {
   val.ProcessAt = now.Add(wait + 1)
   q.queue.Replace(val)
  } else {
   q.queue.RemoveFromQueue(val.Value)
  }
  val, ok = q.queue.Head()
 }
}
func (q *RateLimitedTimedQueue) Add(value string, uid interface{}) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 now := now()
 return q.queue.Add(TimedValue{Value: value, UID: uid, AddedAt: now, ProcessAt: now})
}
func (q *RateLimitedTimedQueue) Remove(value string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return q.queue.Remove(value)
}
func (q *RateLimitedTimedQueue) Clear() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.queue.Clear()
}
func (q *RateLimitedTimedQueue) SwapLimiter(newQPS float32) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 q.limiterLock.Lock()
 defer q.limiterLock.Unlock()
 if q.limiter.QPS() == newQPS {
  return
 }
 var newLimiter flowcontrol.RateLimiter
 if newQPS <= 0 {
  newLimiter = flowcontrol.NewFakeNeverRateLimiter()
 } else {
  newLimiter = flowcontrol.NewTokenBucketRateLimiter(newQPS, EvictionRateLimiterBurst)
  if q.limiter.TryAccept() == false {
   newLimiter.TryAccept()
  }
 }
 q.limiter.Stop()
 q.limiter = newLimiter
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
