package ports

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	ProxyStatusPort                    = 10249
	KubeletPort                        = 10250
	InsecureSchedulerPort              = 10251
	InsecureKubeControllerManagerPort  = 10252
	InsecureCloudControllerManagerPort = 10253
	KubeletReadOnlyPort                = 10255
	ProxyHealthzPort                   = 10256
	KubeControllerManagerPort          = 10257
	CloudControllerManagerPort         = 10258
	KubeSchedulerPort                  = 10259
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
