package cloud

import (
	"context"
	"time"
)

const (
	defaultCallTimeout = 1 * time.Hour
)

func ContextWithCallTimeout() (context.Context, context.CancelFunc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return context.WithTimeout(context.Background(), defaultCallTimeout)
}
