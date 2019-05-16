package v1beta1

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/apis/abac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const allAuthenticated = "system:authenticated"

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return scheme.AddConversionFuncs(func(in *Policy, out *api.Policy, s conversion.Scope) error {
		if err := autoConvert_v1beta1_Policy_To_abac_Policy(in, out, s); err != nil {
			return err
		}
		if in.Spec.User == "*" || in.Spec.Group == "*" {
			out.Spec.Group = allAuthenticated
			out.Spec.User = ""
		}
		return nil
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
