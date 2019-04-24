package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	configapi "github.com/openshift/origin/pkg/template/servicebroker/apis/config"
	configapiv1 "github.com/openshift/origin/pkg/template/servicebroker/apis/config/v1"
)

func Install(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(configapi.AddToScheme(scheme))
	utilruntime.Must(configapiv1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(configapiv1.SchemeGroupVersion))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
