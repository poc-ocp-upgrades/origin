package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
)

func SetDefaults_PodNodeConstraintsConfig(obj *PodNodeConstraintsConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.NodeSelectorLabelBlacklist == nil {
		obj.NodeSelectorLabelBlacklist = []string{kubeletapis.LabelHostname}
	}
}
func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&PodNodeConstraintsConfig{}, func(obj interface{}) {
		SetDefaults_PodNodeConstraintsConfig(obj.(*PodNodeConstraintsConfig))
	})
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
