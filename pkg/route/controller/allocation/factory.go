package allocation

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/route"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RouteAllocationControllerFactory struct{}

func (factory *RouteAllocationControllerFactory) Create(plugin route.AllocationPlugin) *RouteAllocationController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &RouteAllocationController{Plugin: plugin}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
