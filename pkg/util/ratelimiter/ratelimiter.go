package ratelimiter

import (
	goformat "fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	kcache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/flowcontrol"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type HandlerFunc func() error
type RateLimitedFunction struct {
	Handler HandlerFunc
	queue   kcache.Queue
	flowcontrol.RateLimiter
}

func NewRateLimitedFunction(keyFunc kcache.KeyFunc, interval int, handlerFunc HandlerFunc) *RateLimitedFunction {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fifo := kcache.NewFIFO(keyFunc)
	qps := float32(1000.0)
	if interval > 0 {
		qps = float32(1.0 / float32(interval))
	}
	limiter := flowcontrol.NewTokenBucketRateLimiter(qps, 1)
	return &RateLimitedFunction{handlerFunc, fifo, limiter}
}
func (rlf *RateLimitedFunction) RunUntil(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	go utilwait.Until(rlf.pop, 0, stopCh)
}
func (rlf *RateLimitedFunction) pop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rlf.RateLimiter.Accept()
	if _, err := rlf.queue.Pop(func(_ interface{}) error {
		return rlf.Handler()
	}); err != nil {
		utilruntime.HandleError(err)
	}
}
func (rlf *RateLimitedFunction) Invoke(resource interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rlf.queue.AddIfNotPresent(resource)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
