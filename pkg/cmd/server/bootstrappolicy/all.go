package bootstrappolicy

import (
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type PolicyData struct {
	ClusterRoles            []rbacv1.ClusterRole
	ClusterRoleBindings     []rbacv1.ClusterRoleBinding
	Roles                   map[string][]rbacv1.Role
	RoleBindings            map[string][]rbacv1.RoleBinding
	ClusterRolesToAggregate map[string]string
}

func Policy() *PolicyData {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PolicyData{ClusterRoles: GetBootstrapClusterRoles(), ClusterRoleBindings: GetBootstrapClusterRoleBindings(), Roles: NamespaceRoles(), RoleBindings: NamespaceRoleBindings(), ClusterRolesToAggregate: GetBootstrapClusterRolesToAggregate()}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
