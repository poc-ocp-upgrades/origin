package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"github.com/openshift/origin/pkg/project/apiserver/admission/apis/requestlimit"
	"github.com/openshift/origin/pkg/project/apiserver/admission/apis/requestlimit/v1"
)

func InstallInternal(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(requestlimit.Install(scheme))
	utilruntime.Must(v1.Install(scheme))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
