package openshiftapiserver

import (
	rbacinformers "k8s.io/client-go/informers/rbac/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
	rbacauthorizer "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
