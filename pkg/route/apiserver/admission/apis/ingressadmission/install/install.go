package install

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/route/apiserver/admission/apis/ingressadmission"
	"github.com/openshift/origin/pkg/route/apiserver/admission/apis/ingressadmission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func InstallInternal(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(ingressadmission.Install(scheme))
	utilruntime.Must(v1.Install(scheme))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
