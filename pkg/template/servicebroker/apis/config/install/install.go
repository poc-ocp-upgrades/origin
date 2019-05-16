package install

import (
	goformat "fmt"
	configapi "github.com/openshift/origin/pkg/template/servicebroker/apis/config"
	configapiv1 "github.com/openshift/origin/pkg/template/servicebroker/apis/config/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Install(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(configapi.AddToScheme(scheme))
	utilruntime.Must(configapiv1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(configapiv1.SchemeGroupVersion))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
