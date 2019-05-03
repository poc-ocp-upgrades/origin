package podautoscaler

import (
 "time"
 "k8s.io/client-go/util/workqueue"
)

type FixedItemIntervalRateLimiter struct{ interval time.Duration }

var _ workqueue.RateLimiter = &FixedItemIntervalRateLimiter{}

func NewFixedItemIntervalRateLimiter(interval time.Duration) workqueue.RateLimiter {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FixedItemIntervalRateLimiter{interval: interval}
}
func (r *FixedItemIntervalRateLimiter) When(item interface{}) time.Duration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.interval
}
func (r *FixedItemIntervalRateLimiter) NumRequeues(item interface{}) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return 1
}
func (r *FixedItemIntervalRateLimiter) Forget(item interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func NewDefaultHPARateLimiter(interval time.Duration) workqueue.RateLimiter {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return NewFixedItemIntervalRateLimiter(interval)
}
