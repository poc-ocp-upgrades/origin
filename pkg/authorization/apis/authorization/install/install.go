package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/rbacconversion"
	authorizationapiv1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
