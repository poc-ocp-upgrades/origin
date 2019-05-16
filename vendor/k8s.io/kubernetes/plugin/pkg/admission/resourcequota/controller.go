package resourcequota

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
	_ "k8s.io/kubernetes/pkg/util/reflector/prometheus"
	_ "k8s.io/kubernetes/pkg/util/workqueue/prometheus"
	resourcequotaapi "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
	"sort"
	"strings"
	"sync"
	"time"
)

type Evaluator interface {
	Evaluate(a admission.Attributes) error
}
type quotaEvaluator struct {
	quotaAccessor       QuotaAccessor
	lockAcquisitionFunc func([]corev1.ResourceQuota) func()
	ignoredResources    map[schema.GroupResource]struct{}
	registry            quota.Registry
	queue               *workqueue.Type
	workLock            sync.Mutex
	work                map[string][]*admissionWaiter
	dirtyWork           map[string][]*admissionWaiter
	inProgress          sets.String
	workers             int
	stopCh              <-chan struct{}
	init                sync.Once
	config              *resourcequotaapi.Configuration
}
type admissionWaiter struct {
	attributes admission.Attributes
	finished   chan struct{}
	result     error
}
type defaultDeny struct{}

func (defaultDeny) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "DEFAULT DENY"
}
func IsDefaultDeny(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return false
	}
	_, ok := err.(defaultDeny)
	return ok
}
func newAdmissionWaiter(a admission.Attributes) *admissionWaiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &admissionWaiter{attributes: a, finished: make(chan struct{}), result: defaultDeny{}}
}
func NewQuotaEvaluator(quotaAccessor QuotaAccessor, ignoredResources map[schema.GroupResource]struct{}, quotaRegistry quota.Registry, lockAcquisitionFunc func([]corev1.ResourceQuota) func(), config *resourcequotaapi.Configuration, workers int, stopCh <-chan struct{}) Evaluator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if config == nil {
		config = &resourcequotaapi.Configuration{}
	}
	return &quotaEvaluator{quotaAccessor: quotaAccessor, lockAcquisitionFunc: lockAcquisitionFunc, ignoredResources: ignoredResources, registry: quotaRegistry, queue: workqueue.NewNamed("admission_quota_controller"), work: map[string][]*admissionWaiter{}, dirtyWork: map[string][]*admissionWaiter{}, inProgress: sets.String{}, workers: workers, stopCh: stopCh, config: config}
}
func (e *quotaEvaluator) run() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	for i := 0; i < e.workers; i++ {
		go wait.Until(e.doWork, time.Second, e.stopCh)
	}
	<-e.stopCh
	klog.Infof("Shutting down quota evaluator")
	e.queue.ShutDown()
}
func (e *quotaEvaluator) doWork() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	workFunc := func() bool {
		ns, admissionAttributes, quit := e.getWork()
		if quit {
			return true
		}
		defer e.completeWork(ns)
		if len(admissionAttributes) == 0 {
			return false
		}
		e.checkAttributes(ns, admissionAttributes)
		return false
	}
	for {
		if quit := workFunc(); quit {
			klog.Infof("quota evaluator worker shutdown")
			return
		}
	}
}
func (e *quotaEvaluator) checkAttributes(ns string, admissionAttributes []*admissionWaiter) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer func() {
		for _, admissionAttribute := range admissionAttributes {
			close(admissionAttribute.finished)
		}
	}()
	quotas, err := e.quotaAccessor.GetQuotas(ns)
	if err != nil {
		for _, admissionAttribute := range admissionAttributes {
			admissionAttribute.result = err
		}
		return
	}
	limitedResourcesDisabled := len(e.config.LimitedResources) == 0
	if len(quotas) == 0 && limitedResourcesDisabled {
		for _, admissionAttribute := range admissionAttributes {
			admissionAttribute.result = nil
		}
		return
	}
	if e.lockAcquisitionFunc != nil {
		releaseLocks := e.lockAcquisitionFunc(quotas)
		defer releaseLocks()
	}
	e.checkQuotas(quotas, admissionAttributes, 3)
}
func (e *quotaEvaluator) checkQuotas(quotas []corev1.ResourceQuota, admissionAttributes []*admissionWaiter, remainingRetries int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	originalQuotas, err := copyQuotas(quotas)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	atLeastOneChanged := false
	for i := range admissionAttributes {
		admissionAttribute := admissionAttributes[i]
		newQuotas, err := e.checkRequest(quotas, admissionAttribute.attributes)
		if err != nil {
			admissionAttribute.result = err
			continue
		}
		if admissionAttribute.attributes.IsDryRun() {
			admissionAttribute.result = nil
			continue
		}
		atLeastOneChangeForThisWaiter := false
		for j := range newQuotas {
			if !quota.Equals(quotas[j].Status.Used, newQuotas[j].Status.Used) {
				atLeastOneChanged = true
				atLeastOneChangeForThisWaiter = true
				break
			}
		}
		if !atLeastOneChangeForThisWaiter {
			admissionAttribute.result = nil
		}
		quotas = newQuotas
	}
	if !atLeastOneChanged {
		return
	}
	var updatedFailedQuotas []corev1.ResourceQuota
	var lastErr error
	for i := range quotas {
		newQuota := quotas[i]
		if quota.Equals(originalQuotas[i].Status.Used, newQuota.Status.Used) {
			continue
		}
		if err := e.quotaAccessor.UpdateQuotaStatus(&newQuota); err != nil {
			updatedFailedQuotas = append(updatedFailedQuotas, newQuota)
			lastErr = err
		}
	}
	if len(updatedFailedQuotas) == 0 {
		for _, admissionAttribute := range admissionAttributes {
			if IsDefaultDeny(admissionAttribute.result) {
				admissionAttribute.result = nil
			}
		}
		return
	}
	if remainingRetries <= 0 {
		for _, admissionAttribute := range admissionAttributes {
			if IsDefaultDeny(admissionAttribute.result) {
				admissionAttribute.result = lastErr
			}
		}
		return
	}
	newQuotas, err := e.quotaAccessor.GetQuotas(quotas[0].Namespace)
	if err != nil {
		for _, admissionAttribute := range admissionAttributes {
			if IsDefaultDeny(admissionAttribute.result) {
				admissionAttribute.result = lastErr
			}
		}
		return
	}
	quotasToCheck := []corev1.ResourceQuota{}
	for _, newQuota := range newQuotas {
		for _, oldQuota := range updatedFailedQuotas {
			if newQuota.Name == oldQuota.Name {
				quotasToCheck = append(quotasToCheck, newQuota)
				break
			}
		}
	}
	e.checkQuotas(quotasToCheck, admissionAttributes, remainingRetries-1)
}
func copyQuotas(in []corev1.ResourceQuota) ([]corev1.ResourceQuota, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out := make([]corev1.ResourceQuota, 0, len(in))
	for _, quota := range in {
		out = append(out, *quota.DeepCopy())
	}
	return out, nil
}
func filterLimitedResourcesByGroupResource(input []resourcequotaapi.LimitedResource, groupResource schema.GroupResource) []resourcequotaapi.LimitedResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := []resourcequotaapi.LimitedResource{}
	for i := range input {
		limitedResource := input[i]
		limitedGroupResource := schema.GroupResource{Group: limitedResource.APIGroup, Resource: limitedResource.Resource}
		if limitedGroupResource == groupResource {
			result = append(result, limitedResource)
		}
	}
	return result
}
func limitedByDefault(usage corev1.ResourceList, limitedResources []resourcequotaapi.LimitedResource) []corev1.ResourceName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := []corev1.ResourceName{}
	for _, limitedResource := range limitedResources {
		for k, v := range usage {
			if v.Sign() == 1 {
				for _, matchContain := range limitedResource.MatchContains {
					if strings.Contains(string(k), matchContain) {
						result = append(result, k)
						break
					}
				}
			}
		}
	}
	return result
}
func getMatchedLimitedScopes(evaluator quota.Evaluator, inputObject runtime.Object, limitedResources []resourcequotaapi.LimitedResource) ([]corev1.ScopedResourceSelectorRequirement, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scopes := []corev1.ScopedResourceSelectorRequirement{}
	for _, limitedResource := range limitedResources {
		matched, err := evaluator.MatchingScopes(inputObject, limitedResource.MatchScopes)
		if err != nil {
			klog.Errorf("Error while matching limited Scopes: %v", err)
			return []corev1.ScopedResourceSelectorRequirement{}, err
		}
		for _, scope := range matched {
			scopes = append(scopes, scope)
		}
	}
	return scopes, nil
}
func (e *quotaEvaluator) checkRequest(quotas []corev1.ResourceQuota, a admission.Attributes) ([]corev1.ResourceQuota, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	evaluator := e.registry.Get(a.GetResource().GroupResource())
	if evaluator == nil {
		return quotas, nil
	}
	return CheckRequest(quotas, a, evaluator, e.config.LimitedResources)
}
func CheckRequest(quotas []corev1.ResourceQuota, a admission.Attributes, evaluator quota.Evaluator, limited []resourcequotaapi.LimitedResource) ([]corev1.ResourceQuota, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !evaluator.Handles(a) {
		return quotas, nil
	}
	inputObject := a.GetObject()
	limitedScopes, err := getMatchedLimitedScopes(evaluator, inputObject, limited)
	if err != nil {
		return quotas, nil
	}
	limitedResourceNames := []corev1.ResourceName{}
	limitedResources := filterLimitedResourcesByGroupResource(limited, a.GetResource().GroupResource())
	if len(limitedResources) > 0 {
		deltaUsage, err := evaluator.Usage(inputObject)
		if err != nil {
			return quotas, err
		}
		limitedResourceNames = limitedByDefault(deltaUsage, limitedResources)
	}
	limitedResourceNamesSet := quota.ToSet(limitedResourceNames)
	interestingQuotaIndexes := []int{}
	restrictedResourcesSet := sets.String{}
	restrictedScopes := []corev1.ScopedResourceSelectorRequirement{}
	for i := range quotas {
		resourceQuota := quotas[i]
		scopeSelectors := getScopeSelectorsFromQuota(resourceQuota)
		localRestrictedScopes, err := evaluator.MatchingScopes(inputObject, scopeSelectors)
		if err != nil {
			return nil, fmt.Errorf("error matching scopes of quota %s, err: %v", resourceQuota.Name, err)
		}
		for _, scope := range localRestrictedScopes {
			restrictedScopes = append(restrictedScopes, scope)
		}
		match, err := evaluator.Matches(&resourceQuota, inputObject)
		if err != nil {
			klog.Errorf("Error occurred while matching resource quota, %v, against input object. Err: %v", resourceQuota, err)
			return quotas, err
		}
		if !match {
			continue
		}
		hardResources := quota.ResourceNames(resourceQuota.Status.Hard)
		restrictedResources := evaluator.MatchingResources(hardResources)
		if err := evaluator.Constraints(restrictedResources, inputObject); err != nil {
			return nil, admission.NewForbidden(a, fmt.Errorf("failed quota: %s: %v", resourceQuota.Name, err))
		}
		if !hasUsageStats(&resourceQuota, restrictedResources) {
			return nil, admission.NewForbidden(a, fmt.Errorf("status unknown for quota: %s, resources: %s", resourceQuota.Name, prettyPrintResourceNames(restrictedResources)))
		}
		interestingQuotaIndexes = append(interestingQuotaIndexes, i)
		localRestrictedResourcesSet := quota.ToSet(restrictedResources)
		restrictedResourcesSet.Insert(localRestrictedResourcesSet.List()...)
	}
	hasNoCoveringQuota := limitedResourceNamesSet.Difference(restrictedResourcesSet)
	if len(hasNoCoveringQuota) > 0 {
		return quotas, admission.NewForbidden(a, fmt.Errorf("insufficient quota to consume: %v", strings.Join(hasNoCoveringQuota.List(), ",")))
	}
	scopesHasNoCoveringQuota, err := evaluator.UncoveredQuotaScopes(limitedScopes, restrictedScopes)
	if err != nil {
		return quotas, err
	}
	if len(scopesHasNoCoveringQuota) > 0 {
		return quotas, fmt.Errorf("insufficient quota to match these scopes: %v", scopesHasNoCoveringQuota)
	}
	if len(interestingQuotaIndexes) == 0 {
		return quotas, nil
	}
	namespace := a.GetNamespace()
	if accessor, err := meta.Accessor(inputObject); namespace != "" && err == nil {
		if accessor.GetNamespace() == "" {
			accessor.SetNamespace(namespace)
		}
	}
	deltaUsage, err := evaluator.Usage(inputObject)
	if err != nil {
		return quotas, err
	}
	if negativeUsage := quota.IsNegative(deltaUsage); len(negativeUsage) > 0 {
		return nil, admission.NewForbidden(a, fmt.Errorf("quota usage is negative for resource(s): %s", prettyPrintResourceNames(negativeUsage)))
	}
	if admission.Update == a.GetOperation() {
		prevItem := a.GetOldObject()
		if prevItem == nil {
			return nil, admission.NewForbidden(a, fmt.Errorf("unable to get previous usage since prior version of object was not found"))
		}
		metadata, err := meta.Accessor(prevItem)
		if err == nil && len(metadata.GetResourceVersion()) > 0 {
			prevUsage, innerErr := evaluator.Usage(prevItem)
			if innerErr != nil {
				return quotas, innerErr
			}
			deltaUsage = quota.SubtractWithNonNegativeResult(deltaUsage, prevUsage)
		}
	}
	if quota.IsZero(deltaUsage) {
		return quotas, nil
	}
	outQuotas, err := copyQuotas(quotas)
	if err != nil {
		return nil, err
	}
	for _, index := range interestingQuotaIndexes {
		resourceQuota := outQuotas[index]
		hardResources := quota.ResourceNames(resourceQuota.Status.Hard)
		requestedUsage := quota.Mask(deltaUsage, hardResources)
		newUsage := quota.Add(resourceQuota.Status.Used, requestedUsage)
		maskedNewUsage := quota.Mask(newUsage, quota.ResourceNames(requestedUsage))
		if allowed, exceeded := quota.LessThanOrEqual(maskedNewUsage, resourceQuota.Status.Hard); !allowed {
			failedRequestedUsage := quota.Mask(requestedUsage, exceeded)
			failedUsed := quota.Mask(resourceQuota.Status.Used, exceeded)
			failedHard := quota.Mask(resourceQuota.Status.Hard, exceeded)
			return nil, admission.NewForbidden(a, fmt.Errorf("exceeded quota: %s, requested: %s, used: %s, limited: %s", resourceQuota.Name, prettyPrint(failedRequestedUsage), prettyPrint(failedUsed), prettyPrint(failedHard)))
		}
		outQuotas[index].Status.Used = newUsage
	}
	return outQuotas, nil
}
func getScopeSelectorsFromQuota(quota corev1.ResourceQuota) []corev1.ScopedResourceSelectorRequirement {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func (e *quotaEvaluator) Evaluate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.init.Do(func() {
		go e.run()
	})
	gvr := a.GetResource()
	gr := gvr.GroupResource()
	if _, ok := e.ignoredResources[gr]; ok {
		return nil
	}
	evaluator := e.registry.Get(gr)
	if evaluator == nil {
		evaluator = generic.NewObjectCountEvaluator(gr, nil, "")
		e.registry.Add(evaluator)
		klog.Infof("quota admission added evaluator for: %s", gr)
	}
	if !evaluator.Handles(a) {
		return nil
	}
	waiter := newAdmissionWaiter(a)
	e.addWork(waiter)
	select {
	case <-waiter.finished:
	case <-time.After(10 * time.Second):
		return apierrors.NewInternalError(fmt.Errorf("resource quota evaluates timeout"))
	}
	return waiter.result
}
func (e *quotaEvaluator) addWork(a *admissionWaiter) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.workLock.Lock()
	defer e.workLock.Unlock()
	ns := a.attributes.GetNamespace()
	e.queue.Add(ns)
	if e.inProgress.Has(ns) {
		e.dirtyWork[ns] = append(e.dirtyWork[ns], a)
		return
	}
	e.work[ns] = append(e.work[ns], a)
}
func (e *quotaEvaluator) completeWork(ns string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.workLock.Lock()
	defer e.workLock.Unlock()
	e.queue.Done(ns)
	e.work[ns] = e.dirtyWork[ns]
	delete(e.dirtyWork, ns)
	e.inProgress.Delete(ns)
}
func (e *quotaEvaluator) getWork() (string, []*admissionWaiter, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	uncastNS, shutdown := e.queue.Get()
	if shutdown {
		return "", []*admissionWaiter{}, shutdown
	}
	ns := uncastNS.(string)
	e.workLock.Lock()
	defer e.workLock.Unlock()
	work := e.work[ns]
	delete(e.work, ns)
	delete(e.dirtyWork, ns)
	e.inProgress.Insert(ns)
	return ns, work, false
}
func prettyPrint(item corev1.ResourceList) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := []string{}
	keys := []string{}
	for key := range item {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := item[corev1.ResourceName(key)]
		constraint := key + "=" + value.String()
		parts = append(parts, constraint)
	}
	return strings.Join(parts, ",")
}
func prettyPrintResourceNames(a []corev1.ResourceName) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	values := []string{}
	for _, value := range a {
		values = append(values, string(value))
	}
	sort.Strings(values)
	return strings.Join(values, ",")
}
func hasUsageStats(resourceQuota *corev1.ResourceQuota, interestingResources []corev1.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	interestingSet := quota.ToSet(interestingResources)
	for resourceName := range resourceQuota.Status.Hard {
		if !interestingSet.Has(string(resourceName)) {
			continue
		}
		if _, found := resourceQuota.Status.Used[resourceName]; !found {
			return false
		}
	}
	return true
}
