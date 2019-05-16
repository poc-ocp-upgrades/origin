package route

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/fields"
	runtime "k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func RouteFieldSelector(obj runtime.Object, fieldSet fields.Set) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	route, ok := obj.(*Route)
	if !ok {
		return fmt.Errorf("%T not a Route", obj)
	}
	fieldSet["spec.path"] = route.Spec.Path
	fieldSet["spec.host"] = route.Spec.Host
	fieldSet["spec.to.name"] = route.Spec.To.Name
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
