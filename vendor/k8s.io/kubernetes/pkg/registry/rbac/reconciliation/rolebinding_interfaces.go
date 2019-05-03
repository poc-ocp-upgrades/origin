package reconciliation

import (
 rbacv1 "k8s.io/api/rbac/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
 rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type RoleBindingAdapter struct{ RoleBinding *rbacv1.RoleBinding }

func (o RoleBindingAdapter) GetObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding
}
func (o RoleBindingAdapter) GetNamespace() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.Namespace
}
func (o RoleBindingAdapter) GetName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.Name
}
func (o RoleBindingAdapter) GetUID() types.UID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.UID
}
func (o RoleBindingAdapter) GetLabels() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.Labels
}
func (o RoleBindingAdapter) SetLabels(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.RoleBinding.Labels = in
}
func (o RoleBindingAdapter) GetAnnotations() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.Annotations
}
func (o RoleBindingAdapter) SetAnnotations(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.RoleBinding.Annotations = in
}
func (o RoleBindingAdapter) GetRoleRef() rbacv1.RoleRef {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.RoleRef
}
func (o RoleBindingAdapter) GetSubjects() []rbacv1.Subject {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.RoleBinding.Subjects
}
func (o RoleBindingAdapter) SetSubjects(in []rbacv1.Subject) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.RoleBinding.Subjects = in
}

type RoleBindingClientAdapter struct {
 Client          rbacv1client.RoleBindingsGetter
 NamespaceClient corev1client.NamespaceInterface
}

func (c RoleBindingClientAdapter) Get(namespace, name string) (RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.RoleBindings(namespace).Get(name, metav1.GetOptions{})
 if err != nil {
  return nil, err
 }
 return RoleBindingAdapter{RoleBinding: ret}, err
}
func (c RoleBindingClientAdapter) Create(in RoleBinding) (RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := tryEnsureNamespace(c.NamespaceClient, in.GetNamespace()); err != nil {
  return nil, err
 }
 ret, err := c.Client.RoleBindings(in.GetNamespace()).Create(in.(RoleBindingAdapter).RoleBinding)
 if err != nil {
  return nil, err
 }
 return RoleBindingAdapter{RoleBinding: ret}, err
}
func (c RoleBindingClientAdapter) Update(in RoleBinding) (RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.RoleBindings(in.GetNamespace()).Update(in.(RoleBindingAdapter).RoleBinding)
 if err != nil {
  return nil, err
 }
 return RoleBindingAdapter{RoleBinding: ret}, err
}
func (c RoleBindingClientAdapter) Delete(namespace, name string, uid types.UID) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Client.RoleBindings(namespace).Delete(name, &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &uid}})
}
