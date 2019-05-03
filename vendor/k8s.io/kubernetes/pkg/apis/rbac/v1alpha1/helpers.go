package v1alpha1

import (
 "fmt"
 rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PolicyRuleBuilder struct {
 PolicyRule rbacv1alpha1.PolicyRule `protobuf:"bytes,1,opt,name=policyRule"`
}

func NewRule(verbs ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &PolicyRuleBuilder{PolicyRule: rbacv1alpha1.PolicyRule{Verbs: verbs}}
}
func (r *PolicyRuleBuilder) Groups(groups ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.APIGroups = append(r.PolicyRule.APIGroups, groups...)
 return r
}
func (r *PolicyRuleBuilder) Resources(resources ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.Resources = append(r.PolicyRule.Resources, resources...)
 return r
}
func (r *PolicyRuleBuilder) Names(names ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.ResourceNames = append(r.PolicyRule.ResourceNames, names...)
 return r
}
func (r *PolicyRuleBuilder) URLs(urls ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.NonResourceURLs = append(r.PolicyRule.NonResourceURLs, urls...)
 return r
}
func (r *PolicyRuleBuilder) RuleOrDie() rbacv1alpha1.PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := r.Rule()
 if err != nil {
  panic(err)
 }
 return ret
}
func (r *PolicyRuleBuilder) Rule() (rbacv1alpha1.PolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(r.PolicyRule.Verbs) == 0 {
  return rbacv1alpha1.PolicyRule{}, fmt.Errorf("verbs are required: %#v", r.PolicyRule)
 }
 switch {
 case len(r.PolicyRule.NonResourceURLs) > 0:
  if len(r.PolicyRule.APIGroups) != 0 || len(r.PolicyRule.Resources) != 0 || len(r.PolicyRule.ResourceNames) != 0 {
   return rbacv1alpha1.PolicyRule{}, fmt.Errorf("non-resource rule may not have apiGroups, resources, or resourceNames: %#v", r.PolicyRule)
  }
 case len(r.PolicyRule.Resources) > 0:
  if len(r.PolicyRule.NonResourceURLs) != 0 {
   return rbacv1alpha1.PolicyRule{}, fmt.Errorf("resource rule may not have nonResourceURLs: %#v", r.PolicyRule)
  }
  if len(r.PolicyRule.APIGroups) == 0 {
   return rbacv1alpha1.PolicyRule{}, fmt.Errorf("resource rule must have apiGroups: %#v", r.PolicyRule)
  }
 default:
  return rbacv1alpha1.PolicyRule{}, fmt.Errorf("a rule must have either nonResourceURLs or resources: %#v", r.PolicyRule)
 }
 return r.PolicyRule, nil
}

type ClusterRoleBindingBuilder struct {
 ClusterRoleBinding rbacv1alpha1.ClusterRoleBinding `protobuf:"bytes,1,opt,name=clusterRoleBinding"`
}

func NewClusterBinding(clusterRoleName string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &ClusterRoleBindingBuilder{ClusterRoleBinding: rbacv1alpha1.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: clusterRoleName}, RoleRef: rbacv1alpha1.RoleRef{APIGroup: GroupName, Kind: "ClusterRole", Name: clusterRoleName}}}
}
func (r *ClusterRoleBindingBuilder) Groups(groups ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, group := range groups {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1alpha1.Subject{Kind: rbacv1alpha1.GroupKind, Name: group})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) Users(users ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, user := range users {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1alpha1.Subject{Kind: rbacv1alpha1.UserKind, Name: user})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) SAs(namespace string, serviceAccountNames ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, saName := range serviceAccountNames {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1alpha1.Subject{Kind: rbacv1alpha1.ServiceAccountKind, Namespace: namespace, Name: saName})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) BindingOrDie() rbacv1alpha1.ClusterRoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := r.Binding()
 if err != nil {
  panic(err)
 }
 return ret
}
func (r *ClusterRoleBindingBuilder) Binding() (rbacv1alpha1.ClusterRoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(r.ClusterRoleBinding.Subjects) == 0 {
  return rbacv1alpha1.ClusterRoleBinding{}, fmt.Errorf("subjects are required: %#v", r.ClusterRoleBinding)
 }
 return r.ClusterRoleBinding, nil
}
