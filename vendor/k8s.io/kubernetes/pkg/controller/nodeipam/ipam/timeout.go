package ipam

import (
	"time"
)

type Timeout struct {
	Resync       time.Duration
	MaxBackoff   time.Duration
	InitialRetry time.Duration
	errs         int
	current      time.Duration
}

func (b *Timeout) Update(ok bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ok {
		b.errs = 0
		b.current = b.Resync
		return
	}
	b.errs++
	if b.errs == 1 {
		b.current = b.InitialRetry
		return
	}
	b.current *= 2
	if b.current >= b.MaxBackoff {
		b.current = b.MaxBackoff
	}
}
func (b *Timeout) Next() time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if b.errs == 0 {
		return b.Resync
	}
	return b.current
}
