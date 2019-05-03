package generic

import (
 "fmt"
 "sync/atomic"
 corev1 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/resource"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apiserver/pkg/admission"
 "k8s.io/client-go/informers"
 "k8s.io/client-go/tools/cache"
 quota "k8s.io/kubernetes/pkg/quota/v1"
)

type InformerForResourceFunc func(schema.GroupVersionResource) (informers.GenericInformer, error)

func ListerFuncForResourceFunc(f InformerForResourceFunc) quota.ListerForResourceFunc {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(gvr schema.GroupVersionResource) (cache.GenericLister, error) {
  informer, err := f(gvr)
  if err != nil {
   return nil, err
  }
  return &protectedLister{hasSynced: cachedHasSynced(informer.Informer().HasSynced), notReadyErr: fmt.Errorf("%v not yet synced", gvr), delegate: informer.Lister()}, nil
 }
}
func cachedHasSynced(hasSynced func() bool) func() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cache := &atomic.Value{}
 cache.Store(false)
 return func() bool {
  if cache.Load().(bool) {
   return true
  }
  if hasSynced() {
   cache.Store(true)
   return true
  }
  return false
 }
}

type protectedLister struct {
 hasSynced   func() bool
 notReadyErr error
 delegate    cache.GenericLister
}

func (p *protectedLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !p.hasSynced() {
  return nil, p.notReadyErr
 }
 return p.delegate.List(selector)
}
func (p *protectedLister) Get(name string) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !p.hasSynced() {
  return nil, p.notReadyErr
 }
 return p.delegate.Get(name)
}
func (p *protectedLister) ByNamespace(namespace string) cache.GenericNamespaceLister {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &protectedNamespaceLister{p.hasSynced, p.notReadyErr, p.delegate.ByNamespace(namespace)}
}

type protectedNamespaceLister struct {
 hasSynced   func() bool
 notReadyErr error
 delegate    cache.GenericNamespaceLister
}

func (p *protectedNamespaceLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !p.hasSynced() {
  return nil, p.notReadyErr
 }
 return p.delegate.List(selector)
}
func (p *protectedNamespaceLister) Get(name string) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !p.hasSynced() {
  return nil, p.notReadyErr
 }
 return p.delegate.Get(name)
}
func ListResourceUsingListerFunc(l quota.ListerForResourceFunc, resource schema.GroupVersionResource) ListFuncByNamespace {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(namespace string) ([]runtime.Object, error) {
  lister, err := l(resource)
  if err != nil {
   return nil, err
  }
  return lister.ByNamespace(namespace).List(labels.Everything())
 }
}
func ObjectCountQuotaResourceNameFor(groupResource schema.GroupResource) corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(groupResource.Group) == 0 {
  return corev1.ResourceName("count/" + groupResource.Resource)
 }
 return corev1.ResourceName("count/" + groupResource.Resource + "." + groupResource.Group)
}

type ListFuncByNamespace func(namespace string) ([]runtime.Object, error)
type MatchesScopeFunc func(scope corev1.ScopedResourceSelectorRequirement, object runtime.Object) (bool, error)
type UsageFunc func(object runtime.Object) (corev1.ResourceList, error)
type MatchingResourceNamesFunc func(input []corev1.ResourceName) []corev1.ResourceName

func MatchesNoScopeFunc(scope corev1.ScopedResourceSelectorRequirement, object runtime.Object) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false, nil
}
func Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object, matchFunc MatchingResourceNamesFunc, scopeFunc MatchesScopeFunc) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if resourceQuota == nil {
  return false, fmt.Errorf("expected non-nil quota")
 }
 matchResource := len(matchFunc(quota.ResourceNames(resourceQuota.Status.Hard))) > 0
 matchScope := true
 for _, scope := range getScopeSelectorsFromQuota(resourceQuota) {
  innerMatch, err := scopeFunc(scope, item)
  if err != nil {
   return false, err
  }
  matchScope = matchScope && innerMatch
 }
 return matchResource && matchScope, nil
}
func getScopeSelectorsFromQuota(quota *corev1.ResourceQuota) []corev1.ScopedResourceSelectorRequirement {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selectors := []corev1.ScopedResourceSelectorRequirement{}
 for _, scope := range quota.Spec.Scopes {
  selectors = append(selectors, corev1.ScopedResourceSelectorRequirement{ScopeName: scope, Operator: corev1.ScopeSelectorOpExists})
 }
 if quota.Spec.ScopeSelector != nil {
  for _, scopeSelector := range quota.Spec.ScopeSelector.MatchExpressions {
   selectors = append(selectors, scopeSelector)
  }
 }
 return selectors
}
func CalculateUsageStats(options quota.UsageStatsOptions, listFunc ListFuncByNamespace, scopeFunc MatchesScopeFunc, usageFunc UsageFunc) (quota.UsageStats, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := quota.UsageStats{Used: corev1.ResourceList{}}
 for _, resourceName := range options.Resources {
  result.Used[resourceName] = resource.Quantity{Format: resource.DecimalSI}
 }
 items, err := listFunc(options.Namespace)
 if err != nil {
  return result, fmt.Errorf("failed to list content: %v", err)
 }
 for _, item := range items {
  matchesScopes := true
  for _, scope := range options.Scopes {
   innerMatch, err := scopeFunc(corev1.ScopedResourceSelectorRequirement{ScopeName: scope}, item)
   if err != nil {
    return result, nil
   }
   if !innerMatch {
    matchesScopes = false
   }
  }
  if options.ScopeSelector != nil {
   for _, selector := range options.ScopeSelector.MatchExpressions {
    innerMatch, err := scopeFunc(selector, item)
    if err != nil {
     return result, nil
    }
    matchesScopes = matchesScopes && innerMatch
   }
  }
  if matchesScopes {
   usage, err := usageFunc(item)
   if err != nil {
    return result, err
   }
   result.Used = quota.Add(result.Used, usage)
  }
 }
 return result, nil
}

type objectCountEvaluator struct {
 groupResource       schema.GroupResource
 listFuncByNamespace ListFuncByNamespace
 resourceNames       []corev1.ResourceName
}

func (o *objectCountEvaluator) Constraints(required []corev1.ResourceName, item runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (o *objectCountEvaluator) Handles(a admission.Attributes) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 operation := a.GetOperation()
 return operation == admission.Create
}
func (o *objectCountEvaluator) Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return Matches(resourceQuota, item, o.MatchingResources, MatchesNoScopeFunc)
}
func (o *objectCountEvaluator) MatchingResources(input []corev1.ResourceName) []corev1.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return quota.Intersection(input, o.resourceNames)
}
func (o *objectCountEvaluator) MatchingScopes(item runtime.Object, scopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (o *objectCountEvaluator) UncoveredQuotaScopes(limitedScopes []corev1.ScopedResourceSelectorRequirement, matchedQuotaScopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (o *objectCountEvaluator) Usage(object runtime.Object) (corev1.ResourceList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 quantity := resource.NewQuantity(1, resource.DecimalSI)
 resourceList := corev1.ResourceList{}
 for _, resourceName := range o.resourceNames {
  resourceList[resourceName] = *quantity
 }
 return resourceList, nil
}
func (o *objectCountEvaluator) GroupResource() schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.groupResource
}
func (o *objectCountEvaluator) UsageStats(options quota.UsageStatsOptions) (quota.UsageStats, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return CalculateUsageStats(options, o.listFuncByNamespace, MatchesNoScopeFunc, o.Usage)
}

var _ quota.Evaluator = &objectCountEvaluator{}

func NewObjectCountEvaluator(groupResource schema.GroupResource, listFuncByNamespace ListFuncByNamespace, alias corev1.ResourceName) quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resourceNames := []corev1.ResourceName{ObjectCountQuotaResourceNameFor(groupResource)}
 if len(alias) > 0 {
  resourceNames = append(resourceNames, alias)
 }
 return &objectCountEvaluator{groupResource: groupResource, listFuncByNamespace: listFuncByNamespace, resourceNames: resourceNames}
}
