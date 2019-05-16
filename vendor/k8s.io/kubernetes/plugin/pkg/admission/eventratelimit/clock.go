package eventratelimit

import (
	"time"
)

type realClock struct{}

func (realClock) Now() time.Time {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return time.Now()
}
func (realClock) Sleep(d time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	time.Sleep(d)
}
