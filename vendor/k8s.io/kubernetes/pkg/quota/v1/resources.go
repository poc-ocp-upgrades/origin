package quota

import (
 "sort"
 "strings"
 corev1 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/resource"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 "k8s.io/apimachinery/pkg/util/sets"
)

func Equals(a corev1.ResourceList, b corev1.ResourceList) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(a) != len(b) {
  return false
 }
 for key, value1 := range a {
  value2, found := b[key]
  if !found {
   return false
  }
  if value1.Cmp(value2) != 0 {
   return false
  }
 }
 return true
}
func V1Equals(a corev1.ResourceList, b corev1.ResourceList) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(a) != len(b) {
  return false
 }
 for key, value1 := range a {
  value2, found := b[key]
  if !found {
   return false
  }
  if value1.Cmp(value2) != 0 {
   return false
  }
 }
 return true
}
func LessThanOrEqual(a corev1.ResourceList, b corev1.ResourceList) (bool, []corev1.ResourceName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := true
 resourceNames := []corev1.ResourceName{}
 for key, value := range b {
  if other, found := a[key]; found {
   if other.Cmp(value) > 0 {
    result = false
    resourceNames = append(resourceNames, key)
   }
  }
 }
 return result, resourceNames
}
func Max(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := corev1.ResourceList{}
 for key, value := range a {
  if other, found := b[key]; found {
   if value.Cmp(other) <= 0 {
    result[key] = *other.Copy()
    continue
   }
  }
  result[key] = *value.Copy()
 }
 for key, value := range b {
  if _, found := result[key]; !found {
   result[key] = *value.Copy()
  }
 }
 return result
}
func Add(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := corev1.ResourceList{}
 for key, value := range a {
  quantity := *value.Copy()
  if other, found := b[key]; found {
   quantity.Add(other)
  }
  result[key] = quantity
 }
 for key, value := range b {
  if _, found := result[key]; !found {
   quantity := *value.Copy()
   result[key] = quantity
  }
 }
 return result
}
func SubtractWithNonNegativeResult(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zero := resource.MustParse("0")
 result := corev1.ResourceList{}
 for key, value := range a {
  quantity := *value.Copy()
  if other, found := b[key]; found {
   quantity.Sub(other)
  }
  if quantity.Cmp(zero) > 0 {
   result[key] = quantity
  } else {
   result[key] = zero
  }
 }
 for key := range b {
  if _, found := result[key]; !found {
   result[key] = zero
  }
 }
 return result
}
func Subtract(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := corev1.ResourceList{}
 for key, value := range a {
  quantity := *value.Copy()
  if other, found := b[key]; found {
   quantity.Sub(other)
  }
  result[key] = quantity
 }
 for key, value := range b {
  if _, found := result[key]; !found {
   quantity := *value.Copy()
   quantity.Neg()
   result[key] = quantity
  }
 }
 return result
}
func Mask(resources corev1.ResourceList, names []corev1.ResourceName) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nameSet := ToSet(names)
 result := corev1.ResourceList{}
 for key, value := range resources {
  if nameSet.Has(string(key)) {
   result[key] = *value.Copy()
  }
 }
 return result
}
func ResourceNames(resources corev1.ResourceList) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := []corev1.ResourceName{}
 for resourceName := range resources {
  result = append(result, resourceName)
 }
 return result
}
func Contains(items []corev1.ResourceName, item corev1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, i := range items {
  if i == item {
   return true
  }
 }
 return false
}
func ContainsPrefix(prefixSet []string, item corev1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, prefix := range prefixSet {
  if strings.HasPrefix(string(item), prefix) {
   return true
  }
 }
 return false
}
func Intersection(a []corev1.ResourceName, b []corev1.ResourceName) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := make([]corev1.ResourceName, 0, len(a))
 for _, item := range a {
  if Contains(result, item) {
   continue
  }
  if !Contains(b, item) {
   continue
  }
  result = append(result, item)
 }
 sort.Slice(result, func(i, j int) bool {
  return result[i] < result[j]
 })
 return result
}
func Difference(a []corev1.ResourceName, b []corev1.ResourceName) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := make([]corev1.ResourceName, 0, len(a))
 for _, item := range a {
  if Contains(b, item) || Contains(result, item) {
   continue
  }
  result = append(result, item)
 }
 sort.Slice(result, func(i, j int) bool {
  return result[i] < result[j]
 })
 return result
}
func IsZero(a corev1.ResourceList) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zero := resource.MustParse("0")
 for _, v := range a {
  if v.Cmp(zero) != 0 {
   return false
  }
 }
 return true
}
func IsNegative(a corev1.ResourceList) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 results := []corev1.ResourceName{}
 zero := resource.MustParse("0")
 for k, v := range a {
  if v.Cmp(zero) < 0 {
   results = append(results, k)
  }
 }
 return results
}
func ToSet(resourceNames []corev1.ResourceName) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := sets.NewString()
 for _, resourceName := range resourceNames {
  result.Insert(string(resourceName))
 }
 return result
}
func CalculateUsage(namespaceName string, scopes []corev1.ResourceQuotaScope, hardLimits corev1.ResourceList, registry Registry, scopeSelector *corev1.ScopeSelector) (corev1.ResourceList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hardResources := ResourceNames(hardLimits)
 potentialResources := []corev1.ResourceName{}
 evaluators := registry.List()
 for _, evaluator := range evaluators {
  potentialResources = append(potentialResources, evaluator.MatchingResources(hardResources)...)
 }
 matchedResources := Intersection(hardResources, potentialResources)
 errors := []error{}
 newUsage := corev1.ResourceList{}
 for _, evaluator := range evaluators {
  intersection := evaluator.MatchingResources(matchedResources)
  if len(intersection) == 0 {
   continue
  }
  usageStatsOptions := UsageStatsOptions{Namespace: namespaceName, Scopes: scopes, Resources: intersection, ScopeSelector: scopeSelector}
  stats, err := evaluator.UsageStats(usageStatsOptions)
  if err != nil {
   errors = append(errors, err)
   matchedResources = Difference(matchedResources, intersection)
   continue
  }
  newUsage = Add(newUsage, stats.Used)
 }
 newUsage = Mask(newUsage, matchedResources)
 return newUsage, utilerrors.NewAggregate(errors)
}
