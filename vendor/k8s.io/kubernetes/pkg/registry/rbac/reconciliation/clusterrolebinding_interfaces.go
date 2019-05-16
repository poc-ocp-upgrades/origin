package reconciliation

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type ClusterRoleBindingAdapter struct{ ClusterRoleBinding *rbacv1.ClusterRoleBinding }

func (o ClusterRoleBindingAdapter) GetObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding
}
func (o ClusterRoleBindingAdapter) GetNamespace() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.Namespace
}
func (o ClusterRoleBindingAdapter) GetName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.Name
}
func (o ClusterRoleBindingAdapter) GetUID() types.UID {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.UID
}
func (o ClusterRoleBindingAdapter) GetLabels() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.Labels
}
func (o ClusterRoleBindingAdapter) SetLabels(in map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRoleBinding.Labels = in
}
func (o ClusterRoleBindingAdapter) GetAnnotations() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.Annotations
}
func (o ClusterRoleBindingAdapter) SetAnnotations(in map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRoleBinding.Annotations = in
}
func (o ClusterRoleBindingAdapter) GetRoleRef() rbacv1.RoleRef {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.RoleRef
}
func (o ClusterRoleBindingAdapter) GetSubjects() []rbacv1.Subject {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRoleBinding.Subjects
}
func (o ClusterRoleBindingAdapter) SetSubjects(in []rbacv1.Subject) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRoleBinding.Subjects = in
}

type ClusterRoleBindingClientAdapter struct {
	Client rbacv1client.ClusterRoleBindingInterface
}

func (c ClusterRoleBindingClientAdapter) Get(namespace, name string) (RoleBinding, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ClusterRoleBindingAdapter{ClusterRoleBinding: ret}, err
}
func (c ClusterRoleBindingClientAdapter) Create(in RoleBinding) (RoleBinding, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Create(in.(ClusterRoleBindingAdapter).ClusterRoleBinding)
	if err != nil {
		return nil, err
	}
	return ClusterRoleBindingAdapter{ClusterRoleBinding: ret}, err
}
func (c ClusterRoleBindingClientAdapter) Update(in RoleBinding) (RoleBinding, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Update(in.(ClusterRoleBindingAdapter).ClusterRoleBinding)
	if err != nil {
		return nil, err
	}
	return ClusterRoleBindingAdapter{ClusterRoleBinding: ret}, err
}
func (c ClusterRoleBindingClientAdapter) Delete(namespace, name string, uid types.UID) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.Client.Delete(name, &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &uid}})
}
