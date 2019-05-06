package util

import (
	"time"
)

type Clock interface{ Now() time.Time }
type RealClock struct{}

func (RealClock) Now() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Now()
}
