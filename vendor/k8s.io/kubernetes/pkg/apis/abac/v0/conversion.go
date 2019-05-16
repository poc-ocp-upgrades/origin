package v0

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
		out.Spec.User = in.User
		out.Spec.Group = in.Group
		out.Spec.Namespace = in.Namespace
		out.Spec.Resource = in.Resource
		out.Spec.Readonly = in.Readonly
		if len(in.User) == 0 && len(in.Group) == 0 {
			out.Spec.Group = allAuthenticated
		}
		if in.User == "*" || in.Group == "*" {
			out.Spec.Group = allAuthenticated
			out.Spec.User = ""
		}
		if len(in.Namespace) == 0 {
			out.Spec.Namespace = "*"
		}
		if len(in.Resource) == 0 {
			out.Spec.Resource = "*"
		}
		out.Spec.APIGroup = "*"
		if len(in.Namespace) == 0 && len(in.Resource) == 0 {
			out.Spec.NonResourcePath = "*"
		}
		return nil
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
