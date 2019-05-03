package validation

import (
 "strings"
 rbacv1 "k8s.io/api/rbac/v1"
)

func Covers(ownerRules, servantRules []rbacv1.PolicyRule) (bool, []rbacv1.PolicyRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 subrules := []rbacv1.PolicyRule{}
 for _, servantRule := range servantRules {
  subrules = append(subrules, BreakdownRule(servantRule)...)
 }
 uncoveredRules := []rbacv1.PolicyRule{}
 for _, subrule := range subrules {
  covered := false
  for _, ownerRule := range ownerRules {
   if ruleCovers(ownerRule, subrule) {
    covered = true
    break
   }
  }
  if !covered {
   uncoveredRules = append(uncoveredRules, subrule)
  }
 }
 return (len(uncoveredRules) == 0), uncoveredRules
}
func BreakdownRule(rule rbacv1.PolicyRule) []rbacv1.PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 subrules := []rbacv1.PolicyRule{}
 for _, group := range rule.APIGroups {
  for _, resource := range rule.Resources {
   for _, verb := range rule.Verbs {
    if len(rule.ResourceNames) > 0 {
     for _, resourceName := range rule.ResourceNames {
      subrules = append(subrules, rbacv1.PolicyRule{APIGroups: []string{group}, Resources: []string{resource}, Verbs: []string{verb}, ResourceNames: []string{resourceName}})
     }
    } else {
     subrules = append(subrules, rbacv1.PolicyRule{APIGroups: []string{group}, Resources: []string{resource}, Verbs: []string{verb}})
    }
   }
  }
 }
 for _, nonResourceURL := range rule.NonResourceURLs {
  for _, verb := range rule.Verbs {
   subrules = append(subrules, rbacv1.PolicyRule{NonResourceURLs: []string{nonResourceURL}, Verbs: []string{verb}})
  }
 }
 return subrules
}
func has(set []string, ele string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, s := range set {
  if s == ele {
   return true
  }
 }
 return false
}
func hasAll(set, contains []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 owning := make(map[string]struct{}, len(set))
 for _, ele := range set {
  owning[ele] = struct{}{}
 }
 for _, ele := range contains {
  if _, ok := owning[ele]; !ok {
   return false
  }
 }
 return true
}
func resourceCoversAll(setResources, coversResources []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if has(setResources, rbacv1.ResourceAll) || hasAll(setResources, coversResources) {
  return true
 }
 for _, path := range coversResources {
  if has(setResources, path) {
   continue
  }
  if !strings.Contains(path, "/") {
   return false
  }
  tokens := strings.SplitN(path, "/", 2)
  resourceToCheck := "*/" + tokens[1]
  if !has(setResources, resourceToCheck) {
   return false
  }
 }
 return true
}
func nonResourceURLsCoversAll(set, covers []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, path := range covers {
  covered := false
  for _, owner := range set {
   if nonResourceURLCovers(owner, path) {
    covered = true
    break
   }
  }
  if !covered {
   return false
  }
 }
 return true
}
func nonResourceURLCovers(ownerPath, subPath string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ownerPath == subPath {
  return true
 }
 return strings.HasSuffix(ownerPath, "*") && strings.HasPrefix(subPath, strings.TrimRight(ownerPath, "*"))
}
func ruleCovers(ownerRule, subRule rbacv1.PolicyRule) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 verbMatches := has(ownerRule.Verbs, rbacv1.VerbAll) || hasAll(ownerRule.Verbs, subRule.Verbs)
 groupMatches := has(ownerRule.APIGroups, rbacv1.APIGroupAll) || hasAll(ownerRule.APIGroups, subRule.APIGroups)
 resourceMatches := resourceCoversAll(ownerRule.Resources, subRule.Resources)
 nonResourceURLMatches := nonResourceURLsCoversAll(ownerRule.NonResourceURLs, subRule.NonResourceURLs)
 resourceNameMatches := false
 if len(subRule.ResourceNames) == 0 {
  resourceNameMatches = (len(ownerRule.ResourceNames) == 0)
 } else {
  resourceNameMatches = (len(ownerRule.ResourceNames) == 0) || hasAll(ownerRule.ResourceNames, subRule.ResourceNames)
 }
 return verbMatches && groupMatches && resourceMatches && resourceNameMatches && nonResourceURLMatches
}
