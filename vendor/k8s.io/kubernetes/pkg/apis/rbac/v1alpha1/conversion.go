package v1alpha1

import (
	goformat "fmt"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	api "k8s.io/kubernetes/pkg/apis/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const allAuthenticated = "system:authenticated"

func Convert_v1alpha1_Subject_To_rbac_Subject(in *rbacv1alpha1.Subject, out *api.Subject, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := autoConvert_v1alpha1_Subject_To_rbac_Subject(in, out, s); err != nil {
		return err
	}
	switch {
	case in.Kind == rbacv1alpha1.ServiceAccountKind:
		out.APIGroup = ""
	case in.Kind == rbacv1alpha1.UserKind:
		out.APIGroup = GroupName
	case in.Kind == rbacv1alpha1.GroupKind:
		out.APIGroup = GroupName
	default:
		if gv, err := schema.ParseGroupVersion(in.APIVersion); err == nil {
			out.APIGroup = gv.Group
		}
	}
	if out.Kind == rbacv1alpha1.UserKind && out.APIGroup == GroupName && out.Name == "*" {
		out.Kind = rbacv1alpha1.GroupKind
		out.Name = allAuthenticated
	}
	return nil
}
func Convert_rbac_Subject_To_v1alpha1_Subject(in *api.Subject, out *rbacv1alpha1.Subject, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := autoConvert_rbac_Subject_To_v1alpha1_Subject(in, out, s); err != nil {
		return err
	}
	switch {
	case in.Kind == rbacv1alpha1.ServiceAccountKind && in.APIGroup == "":
		out.APIVersion = "v1"
	case in.Kind == rbacv1alpha1.UserKind && in.APIGroup == GroupName:
		out.APIVersion = SchemeGroupVersion.String()
	case in.Kind == rbacv1alpha1.GroupKind && in.APIGroup == GroupName:
		out.APIVersion = SchemeGroupVersion.String()
	default:
		out.APIVersion = schema.GroupVersion{Group: in.APIGroup}.String()
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
