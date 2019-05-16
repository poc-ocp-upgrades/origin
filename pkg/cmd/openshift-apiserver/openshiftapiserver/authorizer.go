package openshiftapiserver

import (
	goformat "fmt"
	rbacinformers "k8s.io/client-go/informers/rbac/v1"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
	rbacauthorizer "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewRuleResolver(informers rbacinformers.Interface) rbacregistryvalidation.AuthorizationRuleResolver {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rbacregistryvalidation.NewDefaultRuleResolver(&rbacauthorizer.RoleGetter{Lister: informers.Roles().Lister()}, &rbacauthorizer.RoleBindingLister{Lister: informers.RoleBindings().Lister()}, &rbacauthorizer.ClusterRoleGetter{Lister: informers.ClusterRoles().Lister()}, &rbacauthorizer.ClusterRoleBindingLister{Lister: informers.ClusterRoleBindings().Lister()})
}
func NewSubjectLocator(informers rbacinformers.Interface) rbacauthorizer.SubjectLocator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rbacauthorizer.NewSubjectAccessEvaluator(&rbacauthorizer.RoleGetter{Lister: informers.Roles().Lister()}, &rbacauthorizer.RoleBindingLister{Lister: informers.RoleBindings().Lister()}, &rbacauthorizer.ClusterRoleGetter{Lister: informers.ClusterRoles().Lister()}, &rbacauthorizer.ClusterRoleBindingLister{Lister: informers.ClusterRoleBindings().Lister()}, "")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
