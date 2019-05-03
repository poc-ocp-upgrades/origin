package core

import (
 "fmt"
 "strings"
 "time"
 corev1 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/resource"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/util/clock"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apiserver/pkg/admission"
 api "k8s.io/kubernetes/pkg/apis/core"
 k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
 "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
 "k8s.io/kubernetes/pkg/kubeapiserver/admission/util"
 quota "k8s.io/kubernetes/pkg/quota/v1"
 "k8s.io/kubernetes/pkg/quota/v1/generic"
)

var podObjectCountName = generic.ObjectCountQuotaResourceNameFor(corev1.SchemeGroupVersion.WithResource("pods").GroupResource())
var podResources = []corev1.ResourceName{podObjectCountName, corev1.ResourceCPU, corev1.ResourceMemory, corev1.ResourceEphemeralStorage, corev1.ResourceRequestsCPU, corev1.ResourceRequestsMemory, corev1.ResourceRequestsEphemeralStorage, corev1.ResourceLimitsCPU, corev1.ResourceLimitsMemory, corev1.ResourceLimitsEphemeralStorage, corev1.ResourcePods}
var podResourcePrefixes = []string{corev1.ResourceHugePagesPrefix, corev1.ResourceRequestsHugePagesPrefix}
var requestedResourcePrefixes = []string{corev1.ResourceHugePagesPrefix}

func maskResourceWithPrefix(resource corev1.ResourceName, prefix string) corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return corev1.ResourceName(fmt.Sprintf("%s%s", prefix, string(resource)))
}
func isExtendedResourceNameForQuota(name corev1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return !helper.IsNativeResource(name) && strings.HasPrefix(string(name), corev1.DefaultResourceRequestsPrefix)
}

var validationSet = sets.NewString(string(corev1.ResourceCPU), string(corev1.ResourceMemory), string(corev1.ResourceRequestsCPU), string(corev1.ResourceRequestsMemory), string(corev1.ResourceLimitsCPU), string(corev1.ResourceLimitsMemory))

func NewPodEvaluator(f quota.ListerForResourceFunc, clock clock.Clock) quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 listFuncByNamespace := generic.ListResourceUsingListerFunc(f, corev1.SchemeGroupVersion.WithResource("pods"))
 podEvaluator := &podEvaluator{listFuncByNamespace: listFuncByNamespace, clock: clock}
 return podEvaluator
}

type podEvaluator struct {
 listFuncByNamespace generic.ListFuncByNamespace
 clock               clock.Clock
}

func (p *podEvaluator) Constraints(required []corev1.ResourceName, item runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, err := toExternalPodOrError(item)
 if err != nil {
  return err
 }
 requiredSet := quota.ToSet(required).Intersection(validationSet)
 missingSet := sets.NewString()
 for i := range pod.Spec.Containers {
  enforcePodContainerConstraints(&pod.Spec.Containers[i], requiredSet, missingSet)
 }
 for i := range pod.Spec.InitContainers {
  enforcePodContainerConstraints(&pod.Spec.InitContainers[i], requiredSet, missingSet)
 }
 if len(missingSet) == 0 {
  return nil
 }
 return fmt.Errorf("must specify %s", strings.Join(missingSet.List(), ","))
}
func (p *podEvaluator) GroupResource() schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return corev1.SchemeGroupVersion.WithResource("pods").GroupResource()
}
func (p *podEvaluator) Handles(a admission.Attributes) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 op := a.GetOperation()
 if op == admission.Create {
  return true
 }
 initializationCompletion, err := util.IsInitializationCompletion(a)
 if err != nil {
  utilruntime.HandleError(err)
  return true
 }
 return initializationCompletion
}
func (p *podEvaluator) Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return generic.Matches(resourceQuota, item, p.MatchingResources, podMatchesScopeFunc)
}
func (p *podEvaluator) MatchingResources(input []corev1.ResourceName) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := quota.Intersection(input, podResources)
 for _, resource := range input {
  if quota.ContainsPrefix(podResourcePrefixes, resource) {
   result = append(result, resource)
  }
  if isExtendedResourceNameForQuota(resource) {
   result = append(result, resource)
  }
 }
 return result
}
func (p *podEvaluator) MatchingScopes(item runtime.Object, scopeSelectors []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 matchedScopes := []corev1.ScopedResourceSelectorRequirement{}
 for _, selector := range scopeSelectors {
  match, err := podMatchesScopeFunc(selector, item)
  if err != nil {
   return []corev1.ScopedResourceSelectorRequirement{}, fmt.Errorf("error on matching scope %v: %v", selector, err)
  }
  if match {
   matchedScopes = append(matchedScopes, selector)
  }
 }
 return matchedScopes, nil
}
func (p *podEvaluator) UncoveredQuotaScopes(limitedScopes []corev1.ScopedResourceSelectorRequirement, matchedQuotaScopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 uncoveredScopes := []corev1.ScopedResourceSelectorRequirement{}
 for _, selector := range limitedScopes {
  isCovered := false
  for _, matchedScopeSelector := range matchedQuotaScopes {
   if matchedScopeSelector.ScopeName == selector.ScopeName {
    isCovered = true
    break
   }
  }
  if !isCovered {
   uncoveredScopes = append(uncoveredScopes, selector)
  }
 }
 return uncoveredScopes, nil
}
func (p *podEvaluator) Usage(item runtime.Object) (corev1.ResourceList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return PodUsageFunc(item, p.clock)
}
func (p *podEvaluator) UsageStats(options quota.UsageStatsOptions) (quota.UsageStats, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return generic.CalculateUsageStats(options, p.listFuncByNamespace, podMatchesScopeFunc, p.Usage)
}

var _ quota.Evaluator = &podEvaluator{}

func enforcePodContainerConstraints(container *corev1.Container, requiredSet, missingSet sets.String) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 requests := container.Resources.Requests
 limits := container.Resources.Limits
 containerUsage := podComputeUsageHelper(requests, limits)
 containerSet := quota.ToSet(quota.ResourceNames(containerUsage))
 if !containerSet.Equal(requiredSet) {
  difference := requiredSet.Difference(containerSet)
  missingSet.Insert(difference.List()...)
 }
}
func podComputeUsageHelper(requests corev1.ResourceList, limits corev1.ResourceList) corev1.ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := corev1.ResourceList{}
 result[corev1.ResourcePods] = resource.MustParse("1")
 if request, found := requests[corev1.ResourceCPU]; found {
  result[corev1.ResourceCPU] = request
  result[corev1.ResourceRequestsCPU] = request
 }
 if limit, found := limits[corev1.ResourceCPU]; found {
  result[corev1.ResourceLimitsCPU] = limit
 }
 if request, found := requests[corev1.ResourceMemory]; found {
  result[corev1.ResourceMemory] = request
  result[corev1.ResourceRequestsMemory] = request
 }
 if limit, found := limits[corev1.ResourceMemory]; found {
  result[corev1.ResourceLimitsMemory] = limit
 }
 if request, found := requests[corev1.ResourceEphemeralStorage]; found {
  result[corev1.ResourceEphemeralStorage] = request
  result[corev1.ResourceRequestsEphemeralStorage] = request
 }
 if limit, found := limits[corev1.ResourceEphemeralStorage]; found {
  result[corev1.ResourceLimitsEphemeralStorage] = limit
 }
 for resource, request := range requests {
  if quota.ContainsPrefix(requestedResourcePrefixes, resource) {
   result[resource] = request
   result[maskResourceWithPrefix(resource, corev1.DefaultResourceRequestsPrefix)] = request
  }
  if helper.IsExtendedResourceName(resource) {
   result[maskResourceWithPrefix(resource, corev1.DefaultResourceRequestsPrefix)] = request
  }
 }
 return result
}
func toExternalPodOrError(obj runtime.Object) (*corev1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := &corev1.Pod{}
 switch t := obj.(type) {
 case *corev1.Pod:
  pod = t
 case *api.Pod:
  if err := k8s_api_v1.Convert_core_Pod_To_v1_Pod(t, pod, nil); err != nil {
   return nil, err
  }
 default:
  return nil, fmt.Errorf("expect *api.Pod or *v1.Pod, got %v", t)
 }
 return pod, nil
}
func podMatchesScopeFunc(selector corev1.ScopedResourceSelectorRequirement, object runtime.Object) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, err := toExternalPodOrError(object)
 if err != nil {
  return false, err
 }
 switch selector.ScopeName {
 case corev1.ResourceQuotaScopeTerminating:
  return isTerminating(pod), nil
 case corev1.ResourceQuotaScopeNotTerminating:
  return !isTerminating(pod), nil
 case corev1.ResourceQuotaScopeBestEffort:
  return isBestEffort(pod), nil
 case corev1.ResourceQuotaScopeNotBestEffort:
  return !isBestEffort(pod), nil
 case corev1.ResourceQuotaScopePriorityClass:
  return podMatchesSelector(pod, selector)
 }
 return false, nil
}
func PodUsageFunc(obj runtime.Object, clock clock.Clock) (corev1.ResourceList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, err := toExternalPodOrError(obj)
 if err != nil {
  return corev1.ResourceList{}, err
 }
 result := corev1.ResourceList{podObjectCountName: *(resource.NewQuantity(1, resource.DecimalSI))}
 if !QuotaV1Pod(pod, clock) {
  return result, nil
 }
 requests := corev1.ResourceList{}
 limits := corev1.ResourceList{}
 for i := range pod.Spec.Containers {
  requests = quota.Add(requests, pod.Spec.Containers[i].Resources.Requests)
  limits = quota.Add(limits, pod.Spec.Containers[i].Resources.Limits)
 }
 for i := range pod.Spec.InitContainers {
  requests = quota.Max(requests, pod.Spec.InitContainers[i].Resources.Requests)
  limits = quota.Max(limits, pod.Spec.InitContainers[i].Resources.Limits)
 }
 result = quota.Add(result, podComputeUsageHelper(requests, limits))
 return result, nil
}
func isBestEffort(pod *corev1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return qos.GetPodQOS(pod) == corev1.PodQOSBestEffort
}
func isTerminating(pod *corev1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pod.Spec.ActiveDeadlineSeconds != nil && *pod.Spec.ActiveDeadlineSeconds >= int64(0) {
  return true
 }
 return false
}
func podMatchesSelector(pod *corev1.Pod, selector corev1.ScopedResourceSelectorRequirement) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 labelSelector, err := helper.ScopedResourceSelectorRequirementsAsSelector(selector)
 if err != nil {
  return false, fmt.Errorf("failed to parse and convert selector: %v", err)
 }
 var m map[string]string
 if len(pod.Spec.PriorityClassName) != 0 {
  m = map[string]string{string(corev1.ResourceQuotaScopePriorityClass): pod.Spec.PriorityClassName}
 }
 if labelSelector.Matches(labels.Set(m)) {
  return true, nil
 }
 return false, nil
}
func QuotaV1Pod(pod *corev1.Pod, clock clock.Clock) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if corev1.PodFailed == pod.Status.Phase || corev1.PodSucceeded == pod.Status.Phase {
  return false
 }
 if pod.DeletionTimestamp != nil && pod.DeletionGracePeriodSeconds != nil {
  now := clock.Now()
  deletionTime := pod.DeletionTimestamp.Time
  gracePeriod := time.Duration(*pod.DeletionGracePeriodSeconds) * time.Second
  if now.After(deletionTime.Add(gracePeriod)) {
   return false
  }
 }
 return true
}
