package proc

import (
	"os"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"os/signal"
	"syscall"
	"k8s.io/klog"
)

func StartReaper() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if os.Getpid() == 1 {
		klog.V(4).Infof("Launching reaper")
		go func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGCHLD)
			for {
				sig := <-sigs
				klog.V(4).Infof("Signal received: %v", sig)
				for {
					cpid, _ := syscall.Wait4(-1, nil, syscall.WNOHANG, nil)
					if cpid < 1 {
						break
					}
					klog.V(4).Infof("Reaped process with pid %d", cpid)
				}
			}
		}()
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
