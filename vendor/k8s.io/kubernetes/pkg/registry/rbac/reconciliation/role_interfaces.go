package reconciliation

import (
 rbacv1 "k8s.io/api/rbac/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
 rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type RoleRuleOwner struct{ Role *rbacv1.Role }

func (o RoleRuleOwner) GetObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role
}
func (o RoleRuleOwner) GetNamespace() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role.Namespace
}
func (o RoleRuleOwner) GetName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role.Name
}
func (o RoleRuleOwner) GetLabels() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role.Labels
}
func (o RoleRuleOwner) SetLabels(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.Role.Labels = in
}
func (o RoleRuleOwner) GetAnnotations() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role.Annotations
}
func (o RoleRuleOwner) SetAnnotations(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.Role.Annotations = in
}
func (o RoleRuleOwner) GetRules() []rbacv1.PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.Role.Rules
}
func (o RoleRuleOwner) SetRules(in []rbacv1.PolicyRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.Role.Rules = in
}
func (o RoleRuleOwner) GetAggregationRule() *rbacv1.AggregationRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (o RoleRuleOwner) SetAggregationRule(in *rbacv1.AggregationRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}

type RoleModifier struct {
 Client          rbacv1client.RolesGetter
 NamespaceClient corev1client.NamespaceInterface
}

func (c RoleModifier) Get(namespace, name string) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.Roles(namespace).Get(name, metav1.GetOptions{})
 if err != nil {
  return nil, err
 }
 return RoleRuleOwner{Role: ret}, err
}
func (c RoleModifier) Create(in RuleOwner) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := tryEnsureNamespace(c.NamespaceClient, in.GetNamespace()); err != nil {
  return nil, err
 }
 ret, err := c.Client.Roles(in.GetNamespace()).Create(in.(RoleRuleOwner).Role)
 if err != nil {
  return nil, err
 }
 return RoleRuleOwner{Role: ret}, err
}
func (c RoleModifier) Update(in RuleOwner) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.Roles(in.GetNamespace()).Update(in.(RoleRuleOwner).Role)
 if err != nil {
  return nil, err
 }
 return RoleRuleOwner{Role: ret}, err
}
