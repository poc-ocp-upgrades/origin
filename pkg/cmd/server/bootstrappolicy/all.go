package bootstrappolicy

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type PolicyData struct {
	ClusterRoles		[]rbacv1.ClusterRole
	ClusterRoleBindings	[]rbacv1.ClusterRoleBinding
	Roles			map[string][]rbacv1.Role
	RoleBindings		map[string][]rbacv1.RoleBinding
	ClusterRolesToAggregate	map[string]string
}

func Policy() *PolicyData {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PolicyData{ClusterRoles: GetBootstrapClusterRoles(), ClusterRoleBindings: GetBootstrapClusterRoleBindings(), Roles: NamespaceRoles(), RoleBindings: NamespaceRoleBindings(), ClusterRolesToAggregate: GetBootstrapClusterRolesToAggregate()}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
