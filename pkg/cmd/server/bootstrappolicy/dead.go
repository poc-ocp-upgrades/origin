package bootstrappolicy

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

var (
	deadClusterRoles        = []rbacv1.ClusterRole{}
	deadClusterRoleBindings = []rbacv1.ClusterRoleBinding{}
)

func addDeadClusterRole(name string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, existingRole := range deadClusterRoles {
		if name == existingRole.Name {
			klog.Fatalf("role %q was already registered", name)
		}
	}
	deadClusterRole := rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: name}}
	addDefaultMetadata(&deadClusterRole)
	deadClusterRoles = append(deadClusterRoles, deadClusterRole)
}
func addDeadClusterRoleBinding(name, roleName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, existing := range deadClusterRoleBindings {
		if name == existing.Name {
			klog.Fatalf("%q was already registered", name)
		}
	}
	deadClusterRoleBinding := rbacv1.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: name}, RoleRef: rbacv1.RoleRef{APIGroup: rbacv1.GroupName, Kind: "ClusterRole", Name: roleName}}
	addDefaultMetadata(&deadClusterRoleBinding)
	deadClusterRoleBindings = append(deadClusterRoleBindings, deadClusterRoleBinding)
}
func GetDeadClusterRoles() []rbacv1.ClusterRole {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return deadClusterRoles
}
func GetDeadClusterRoleBindings() []rbacv1.ClusterRoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return deadClusterRoleBindings
}
func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addDeadClusterRole("system:replication-controller")
	addDeadClusterRole("system:endpoint-controller")
	addDeadClusterRole("system:replicaset-controller")
	addDeadClusterRole("system:garbage-collector-controller")
	addDeadClusterRole("system:job-controller")
	addDeadClusterRole("system:hpa-controller")
	addDeadClusterRole("system:daemonset-controller")
	addDeadClusterRole("system:disruption-controller")
	addDeadClusterRole("system:namespace-controller")
	addDeadClusterRole("system:gc-controller")
	addDeadClusterRole("system:certificate-signing-controller")
	addDeadClusterRole("system:statefulset-controller")
	addDeadClusterRole("system:build-controller")
	addDeadClusterRole("system:deploymentconfig-controller")
	addDeadClusterRole("system:deployment-controller")
	addDeadClusterRoleBinding("system:nodes", "system:node")
	addDeadClusterRoleBinding("system:discovery-binding", "system:discovery")
}
