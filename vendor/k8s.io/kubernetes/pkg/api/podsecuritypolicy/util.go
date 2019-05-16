package podsecuritypolicy

import (
	goformat "fmt"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/features"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func DropDisabledAlphaFields(pspSpec *policy.PodSecurityPolicySpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.ProcMountType) {
		pspSpec.AllowedProcMountTypes = nil
	}
	if !utilfeature.DefaultFeatureGate.Enabled(features.RunAsGroup) {
		pspSpec.RunAsGroup = nil
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
