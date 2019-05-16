package v1

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return scheme.AddConversionFuncs()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
