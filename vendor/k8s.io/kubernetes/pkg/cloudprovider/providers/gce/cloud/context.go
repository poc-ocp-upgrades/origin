package cloud

import (
 "context"
 "time"
)

const (
 defaultCallTimeout = 1 * time.Hour
)

func ContextWithCallTimeout() (context.Context, context.CancelFunc) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return context.WithTimeout(context.Background(), defaultCallTimeout)
}
