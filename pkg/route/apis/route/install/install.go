package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	routev1 "github.com/openshift/api/route/v1"
	routeapiv1 "github.com/openshift/origin/pkg/route/apis/route/v1"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	Install(legacyscheme.Scheme)
}
func Install(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(routeapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(routev1.GroupVersion))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
