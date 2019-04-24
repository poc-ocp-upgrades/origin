package rsync

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

func NewDefaultCopyStrategy(o *RsyncOptions) CopyStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strategies := copyStrategies{}
	if hasLocalRsync() {
		if isWindows() {
			strategies = append(strategies, NewRsyncDaemonStrategy(o))
		} else {
			strategies = append(strategies, NewRsyncStrategy(o))
		}
	} else {
		warnNoRsync(o.ErrOut)
	}
	return append(strategies, NewTarStrategy(o))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
