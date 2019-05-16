package v1

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func SetDefaults_PodNodeConstraintsConfig(obj *PodNodeConstraintsConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.NodeSelectorLabelBlacklist == nil {
		obj.NodeSelectorLabelBlacklist = []string{kubeletapis.LabelHostname}
	}
}
func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddTypeDefaultingFunc(&PodNodeConstraintsConfig{}, func(obj interface{}) {
		SetDefaults_PodNodeConstraintsConfig(obj.(*PodNodeConstraintsConfig))
	})
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
