package bootstrappolicy

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/klog"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	"strings"
)

var (
	namespaceRoles        = map[string][]rbacv1.Role{}
	namespaceRoleBindings = map[string][]rbacv1.RoleBinding{}
)

func addNamespaceRole(namespace string, role rbacv1.Role) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(namespace, "kube-") {
		klog.Fatalf(`roles can only be bootstrapped into reserved namespaces starting with "kube-", not %q`, namespace)
	}
	existingRoles := namespaceRoles[namespace]
	for _, existingRole := range existingRoles {
		if role.Name == existingRole.Name {
			klog.Fatalf("role %q was already registered in %q", role.Name, namespace)
		}
	}
	role.Namespace = namespace
	addDefaultMetadata(&role)
	existingRoles = append(existingRoles, role)
	namespaceRoles[namespace] = existingRoles
}
func addNamespaceRoleBinding(namespace string, roleBinding rbacv1.RoleBinding) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(namespace, "kube-") {
		klog.Fatalf(`rolebindings can only be bootstrapped into reserved namespaces starting with "kube-", not %q`, namespace)
	}
	existingRoleBindings := namespaceRoleBindings[namespace]
	for _, existingRoleBinding := range existingRoleBindings {
		if roleBinding.Name == existingRoleBinding.Name {
			klog.Fatalf("rolebinding %q was already registered in %q", roleBinding.Name, namespace)
		}
	}
	roleBinding.Namespace = namespace
	addDefaultMetadata(&roleBinding)
	existingRoleBindings = append(existingRoleBindings, roleBinding)
	namespaceRoleBindings[namespace] = existingRoleBindings
}
func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "extension-apiserver-authentication-reader"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("configmaps").Names("extension-apiserver-authentication").RuleOrDie()}})
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + "bootstrap-signer"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("secrets").RuleOrDie()}})
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + "cloud-provider"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create", "get", "list", "watch").Groups(legacyGroup).Resources("configmaps").RuleOrDie()}})
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + "token-cleaner"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "delete").Groups(legacyGroup).Resources("secrets").RuleOrDie(), eventsRule()}})
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "system::leader-locking-kube-controller-manager"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("watch").Groups(legacyGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(legacyGroup).Resources("configmaps").Names("kube-controller-manager").RuleOrDie()}})
	addNamespaceRole(metav1.NamespaceSystem, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "system::leader-locking-kube-scheduler"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("watch").Groups(legacyGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(legacyGroup).Resources("configmaps").Names("kube-scheduler").RuleOrDie()}})
	delegatedAuthBinding := rbacv1helpers.NewRoleBinding("extension-apiserver-authentication-reader", metav1.NamespaceSystem).Users(user.KubeControllerManager, user.KubeScheduler).BindingOrDie()
	delegatedAuthBinding.Name = "system::extension-apiserver-authentication-reader"
	addNamespaceRoleBinding(metav1.NamespaceSystem, delegatedAuthBinding)
	addNamespaceRoleBinding(metav1.NamespaceSystem, rbacv1helpers.NewRoleBinding("system::leader-locking-kube-controller-manager", metav1.NamespaceSystem).Users(user.KubeControllerManager).SAs(metav1.NamespaceSystem, "kube-controller-manager").BindingOrDie())
	addNamespaceRoleBinding(metav1.NamespaceSystem, rbacv1helpers.NewRoleBinding("system::leader-locking-kube-scheduler", metav1.NamespaceSystem).Users(user.KubeScheduler).SAs(metav1.NamespaceSystem, "kube-scheduler").BindingOrDie())
	addNamespaceRoleBinding(metav1.NamespaceSystem, rbacv1helpers.NewRoleBinding(saRolePrefix+"bootstrap-signer", metav1.NamespaceSystem).SAs(metav1.NamespaceSystem, "bootstrap-signer").BindingOrDie())
	addNamespaceRoleBinding(metav1.NamespaceSystem, rbacv1helpers.NewRoleBinding(saRolePrefix+"cloud-provider", metav1.NamespaceSystem).SAs(metav1.NamespaceSystem, "cloud-provider").BindingOrDie())
	addNamespaceRoleBinding(metav1.NamespaceSystem, rbacv1helpers.NewRoleBinding(saRolePrefix+"token-cleaner", metav1.NamespaceSystem).SAs(metav1.NamespaceSystem, "token-cleaner").BindingOrDie())
	addNamespaceRole(metav1.NamespacePublic, rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + "bootstrap-signer"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(legacyGroup).Resources("configmaps").Names("cluster-info").RuleOrDie(), eventsRule()}})
	addNamespaceRoleBinding(metav1.NamespacePublic, rbacv1helpers.NewRoleBinding(saRolePrefix+"bootstrap-signer", metav1.NamespacePublic).SAs(metav1.NamespaceSystem, "bootstrap-signer").BindingOrDie())
}
func NamespaceRoles() map[string][]rbacv1.Role {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return namespaceRoles
}
func NamespaceRoleBindings() map[string][]rbacv1.RoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return namespaceRoleBindings
}
