package util

import (
	"time"
)

type Clock interface{ Now() time.Time }
type RealClock struct{}

func (RealClock) Now() time.Time {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return time.Now()
}
