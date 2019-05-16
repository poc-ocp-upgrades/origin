package algorithmprovider

import (
	goformat "fmt"
	"k8s.io/kubernetes/pkg/scheduler/algorithmprovider/defaults"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ApplyFeatureGates() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaults.ApplyFeatureGates()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
