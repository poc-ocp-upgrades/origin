package podautoscaler

import (
	"k8s.io/client-go/util/workqueue"
	"time"
)

type FixedItemIntervalRateLimiter struct{ interval time.Duration }

var _ workqueue.RateLimiter = &FixedItemIntervalRateLimiter{}

func NewFixedItemIntervalRateLimiter(interval time.Duration) workqueue.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FixedItemIntervalRateLimiter{interval: interval}
}
func (r *FixedItemIntervalRateLimiter) When(item interface{}) time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.interval
}
func (r *FixedItemIntervalRateLimiter) NumRequeues(item interface{}) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return 1
}
func (r *FixedItemIntervalRateLimiter) Forget(item interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func NewDefaultHPARateLimiter(interval time.Duration) workqueue.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewFixedItemIntervalRateLimiter(interval)
}
