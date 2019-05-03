package validation

import (
 "reflect"
 rbacv1 "k8s.io/api/rbac/v1"
)

type simpleResource struct {
 Group             string
 Resource          string
 ResourceNameExist bool
 ResourceName      string
}

func CompactRules(rules []rbacv1.PolicyRule) ([]rbacv1.PolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 compacted := make([]rbacv1.PolicyRule, 0, len(rules))
 simpleRules := map[simpleResource]*rbacv1.PolicyRule{}
 for _, rule := range rules {
  if resource, isSimple := isSimpleResourceRule(&rule); isSimple {
   if existingRule, ok := simpleRules[resource]; ok {
    if existingRule.Verbs == nil {
     existingRule.Verbs = []string{}
    }
    existingRule.Verbs = append(existingRule.Verbs, rule.Verbs...)
   } else {
    simpleRules[resource] = rule.DeepCopy()
   }
  } else {
   compacted = append(compacted, rule)
  }
 }
 for _, simpleRule := range simpleRules {
  compacted = append(compacted, *simpleRule)
 }
 return compacted, nil
}
func isSimpleResourceRule(rule *rbacv1.PolicyRule) (simpleResource, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource := simpleResource{}
 if len(rule.ResourceNames) > 1 || len(rule.NonResourceURLs) > 0 {
  return resource, false
 }
 if len(rule.APIGroups) != 1 || len(rule.Resources) != 1 {
  return resource, false
 }
 simpleRule := &rbacv1.PolicyRule{APIGroups: rule.APIGroups, Resources: rule.Resources, Verbs: rule.Verbs, ResourceNames: rule.ResourceNames}
 if !reflect.DeepEqual(simpleRule, rule) {
  return resource, false
 }
 if len(rule.ResourceNames) == 0 {
  resource = simpleResource{Group: rule.APIGroups[0], Resource: rule.Resources[0], ResourceNameExist: false}
 } else {
  resource = simpleResource{Group: rule.APIGroups[0], Resource: rule.Resources[0], ResourceNameExist: true, ResourceName: rule.ResourceNames[0]}
 }
 return resource, true
}
