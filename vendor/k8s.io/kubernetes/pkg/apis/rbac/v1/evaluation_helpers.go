package v1

import (
 "fmt"
 "strings"
 rbacv1 "k8s.io/api/rbac/v1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

func RoleRefGroupKind(roleRef rbacv1.RoleRef) schema.GroupKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return schema.GroupKind{Group: roleRef.APIGroup, Kind: roleRef.Kind}
}
func VerbMatches(rule *rbacv1.PolicyRule, requestedVerb string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ruleVerb := range rule.Verbs {
  if ruleVerb == rbacv1.VerbAll {
   return true
  }
  if ruleVerb == requestedVerb {
   return true
  }
 }
 return false
}
func APIGroupMatches(rule *rbacv1.PolicyRule, requestedGroup string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ruleGroup := range rule.APIGroups {
  if ruleGroup == rbacv1.APIGroupAll {
   return true
  }
  if ruleGroup == requestedGroup {
   return true
  }
 }
 return false
}
func ResourceMatches(rule *rbacv1.PolicyRule, combinedRequestedResource, requestedSubresource string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ruleResource := range rule.Resources {
  if ruleResource == rbacv1.ResourceAll {
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
func ResourceNameMatches(rule *rbacv1.PolicyRule, requestedName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(rule.ResourceNames) == 0 {
  return true
 }
 for _, ruleName := range rule.ResourceNames {
  if ruleName == requestedName {
   return true
  }
 }
 return false
}
func NonResourceURLMatches(rule *rbacv1.PolicyRule, requestedURL string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ruleURL := range rule.NonResourceURLs {
  if ruleURL == rbacv1.NonResourceAll {
   return true
  }
  if ruleURL == requestedURL {
   return true
  }
  if strings.HasSuffix(ruleURL, "*") && strings.HasPrefix(requestedURL, strings.TrimRight(ruleURL, "*")) {
   return true
  }
 }
 return false
}
func SubjectsStrings(subjects []rbacv1.Subject) ([]string, []string, []string, []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 users := []string{}
 groups := []string{}
 sas := []string{}
 others := []string{}
 for _, subject := range subjects {
  switch subject.Kind {
  case rbacv1.ServiceAccountKind:
   sas = append(sas, fmt.Sprintf("%s/%s", subject.Namespace, subject.Name))
  case rbacv1.UserKind:
   users = append(users, subject.Name)
  case rbacv1.GroupKind:
   groups = append(groups, subject.Name)
  default:
   others = append(others, fmt.Sprintf("%s/%s/%s", subject.Kind, subject.Namespace, subject.Name))
  }
 }
 return users, groups, sas, others
}
func String(r rbacv1.PolicyRule) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "PolicyRule" + CompactString(r)
}
func CompactString(r rbacv1.PolicyRule) string {
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

type SortableRuleSlice []rbacv1.PolicyRule

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
