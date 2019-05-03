package rbac

import (
 "fmt"
 "strings"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/sets"
)

func ResourceMatches(rule *PolicyRule, combinedRequestedResource, requestedSubresource string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ruleResource := range rule.Resources {
  if ruleResource == ResourceAll {
   return true
  }
  if ruleResource == combinedRequestedResource {
   return true
  }
  if len(requestedSubresource) == 0 {
   continue
  }
  if len(ruleResource) == len(requestedSubresource)+2 && strings.HasPrefix(ruleResource, "*/") && strings.HasSuffix(ruleResource, requestedSubresource) {
   return true
  }
 }
 return false
}
func SubjectsStrings(subjects []Subject) ([]string, []string, []string, []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 users := []string{}
 groups := []string{}
 sas := []string{}
 others := []string{}
 for _, subject := range subjects {
  switch subject.Kind {
  case ServiceAccountKind:
   sas = append(sas, fmt.Sprintf("%s/%s", subject.Namespace, subject.Name))
  case UserKind:
   users = append(users, subject.Name)
  case GroupKind:
   groups = append(groups, subject.Name)
  default:
   others = append(others, fmt.Sprintf("%s/%s/%s", subject.Kind, subject.Namespace, subject.Name))
  }
 }
 return users, groups, sas, others
}
func (r PolicyRule) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "PolicyRule" + r.CompactString()
}
func (r PolicyRule) CompactString() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 formatStringParts := []string{}
 formatArgs := []interface{}{}
 if len(r.APIGroups) > 0 {
  formatStringParts = append(formatStringParts, "APIGroups:%q")
  formatArgs = append(formatArgs, r.APIGroups)
 }
 if len(r.Resources) > 0 {
  formatStringParts = append(formatStringParts, "Resources:%q")
  formatArgs = append(formatArgs, r.Resources)
 }
 if len(r.NonResourceURLs) > 0 {
  formatStringParts = append(formatStringParts, "NonResourceURLs:%q")
  formatArgs = append(formatArgs, r.NonResourceURLs)
 }
 if len(r.ResourceNames) > 0 {
  formatStringParts = append(formatStringParts, "ResourceNames:%q")
  formatArgs = append(formatArgs, r.ResourceNames)
 }
 if len(r.Verbs) > 0 {
  formatStringParts = append(formatStringParts, "Verbs:%q")
  formatArgs = append(formatArgs, r.Verbs)
 }
 formatString := "{" + strings.Join(formatStringParts, ", ") + "}"
 return fmt.Sprintf(formatString, formatArgs...)
}

type PolicyRuleBuilder struct{ PolicyRule PolicyRule }

func NewRule(verbs ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &PolicyRuleBuilder{PolicyRule: PolicyRule{Verbs: sets.NewString(verbs...).List()}}
}
func (r *PolicyRuleBuilder) Groups(groups ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.APIGroups = combine(r.PolicyRule.APIGroups, groups)
 return r
}
func (r *PolicyRuleBuilder) Resources(resources ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.Resources = combine(r.PolicyRule.Resources, resources)
 return r
}
func (r *PolicyRuleBuilder) Names(names ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.ResourceNames = combine(r.PolicyRule.ResourceNames, names)
 return r
}
func (r *PolicyRuleBuilder) URLs(urls ...string) *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.PolicyRule.NonResourceURLs = combine(r.PolicyRule.NonResourceURLs, urls)
 return r
}
func (r *PolicyRuleBuilder) RuleOrDie() PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := r.Rule()
 if err != nil {
  panic(err)
 }
 return ret
}
func combine(s1, s2 []string) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s := sets.NewString(s1...)
 s.Insert(s2...)
 return s.List()
}
func (r *PolicyRuleBuilder) Rule() (PolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(r.PolicyRule.Verbs) == 0 {
  return PolicyRule{}, fmt.Errorf("verbs are required: %#v", r.PolicyRule)
 }
 switch {
 case len(r.PolicyRule.NonResourceURLs) > 0:
  if len(r.PolicyRule.APIGroups) != 0 || len(r.PolicyRule.Resources) != 0 || len(r.PolicyRule.ResourceNames) != 0 {
   return PolicyRule{}, fmt.Errorf("non-resource rule may not have apiGroups, resources, or resourceNames: %#v", r.PolicyRule)
  }
 case len(r.PolicyRule.Resources) > 0:
  if len(r.PolicyRule.APIGroups) == 0 {
   return PolicyRule{}, fmt.Errorf("resource rule must have apiGroups: %#v", r.PolicyRule)
  }
  if len(r.PolicyRule.ResourceNames) != 0 {
   illegalVerbs := []string{}
   for _, verb := range r.PolicyRule.Verbs {
    switch verb {
    case "list", "watch", "create", "deletecollection":
     illegalVerbs = append(illegalVerbs, verb)
    }
   }
   if len(illegalVerbs) > 0 {
    return PolicyRule{}, fmt.Errorf("verbs %v do not have names available: %#v", illegalVerbs, r.PolicyRule)
   }
  }
 default:
  return PolicyRule{}, fmt.Errorf("a rule must have either nonResourceURLs or resources: %#v", r.PolicyRule)
 }
 return r.PolicyRule, nil
}

type ClusterRoleBindingBuilder struct{ ClusterRoleBinding ClusterRoleBinding }

func NewClusterBinding(clusterRoleName string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &ClusterRoleBindingBuilder{ClusterRoleBinding: ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: clusterRoleName}, RoleRef: RoleRef{APIGroup: GroupName, Kind: "ClusterRole", Name: clusterRoleName}}}
}
func (r *ClusterRoleBindingBuilder) Groups(groups ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, group := range groups {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, Subject{Kind: GroupKind, APIGroup: GroupName, Name: group})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) Users(users ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, user := range users {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, Subject{Kind: UserKind, APIGroup: GroupName, Name: user})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) SAs(namespace string, serviceAccountNames ...string) *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, saName := range serviceAccountNames {
  r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, Subject{Kind: ServiceAccountKind, Namespace: namespace, Name: saName})
 }
 return r
}
func (r *ClusterRoleBindingBuilder) BindingOrDie() ClusterRoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := r.Binding()
 if err != nil {
  panic(err)
 }
 return ret
}
func (r *ClusterRoleBindingBuilder) Binding() (ClusterRoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(r.ClusterRoleBinding.Subjects) == 0 {
  return ClusterRoleBinding{}, fmt.Errorf("subjects are required: %#v", r.ClusterRoleBinding)
 }
 return r.ClusterRoleBinding, nil
}

type RoleBindingBuilder struct{ RoleBinding RoleBinding }

func NewRoleBinding(roleName, namespace string) *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RoleBindingBuilder{RoleBinding: RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: roleName, Namespace: namespace}, RoleRef: RoleRef{APIGroup: GroupName, Kind: "Role", Name: roleName}}}
}
func NewRoleBindingForClusterRole(roleName, namespace string) *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RoleBindingBuilder{RoleBinding: RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: roleName, Namespace: namespace}, RoleRef: RoleRef{APIGroup: GroupName, Kind: "ClusterRole", Name: roleName}}}
}
func (r *RoleBindingBuilder) Groups(groups ...string) *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, group := range groups {
  r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, Subject{Kind: GroupKind, APIGroup: GroupName, Name: group})
 }
 return r
}
func (r *RoleBindingBuilder) Users(users ...string) *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, user := range users {
  r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, Subject{Kind: UserKind, APIGroup: GroupName, Name: user})
 }
 return r
}
func (r *RoleBindingBuilder) SAs(namespace string, serviceAccountNames ...string) *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, saName := range serviceAccountNames {
  r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, Subject{Kind: ServiceAccountKind, Namespace: namespace, Name: saName})
 }
 return r
}
func (r *RoleBindingBuilder) BindingOrDie() RoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := r.Binding()
 if err != nil {
  panic(err)
 }
 return ret
}
func (r *RoleBindingBuilder) Binding() (RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(r.RoleBinding.Subjects) == 0 {
  return RoleBinding{}, fmt.Errorf("subjects are required: %#v", r.RoleBinding)
 }
 return r.RoleBinding, nil
}

type SortableRuleSlice []PolicyRule

func (s SortableRuleSlice) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(s)
}
func (s SortableRuleSlice) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s[i], s[j] = s[j], s[i]
}
func (s SortableRuleSlice) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.Compare(s[i].String(), s[j].String()) < 0
}
