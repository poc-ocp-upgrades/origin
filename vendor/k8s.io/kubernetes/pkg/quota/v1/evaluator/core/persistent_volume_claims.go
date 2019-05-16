package core

import (
	"fmt"
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/initialization"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/features"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	api "k8s.io/kubernetes/pkg/apis/core"
	k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/apis/core/v1/helper"
	k8sfeatures "k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/kubeapiserver/admission/util"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var pvcObjectCountName = generic.ObjectCountQuotaResourceNameFor(corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims").GroupResource())
var pvcResources = []corev1.ResourceName{corev1.ResourcePersistentVolumeClaims, corev1.ResourceRequestsStorage}

const storageClassSuffix string = ".storageclass.storage.k8s.io/"

func V1ResourceByStorageClass(storageClass string, resourceName corev1.ResourceName) corev1.ResourceName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return corev1.ResourceName(string(storageClass + storageClassSuffix + string(resourceName)))
}
func NewPersistentVolumeClaimEvaluator(f quota.ListerForResourceFunc) quota.Evaluator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	listFuncByNamespace := generic.ListResourceUsingListerFunc(f, corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims"))
	pvcEvaluator := &pvcEvaluator{listFuncByNamespace: listFuncByNamespace}
	return pvcEvaluator
}

type pvcEvaluator struct{ listFuncByNamespace generic.ListFuncByNamespace }

func (p *pvcEvaluator) Constraints(required []corev1.ResourceName, item runtime.Object) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (p *pvcEvaluator) GroupResource() schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims").GroupResource()
}
func (p *pvcEvaluator) Handles(a admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	op := a.GetOperation()
	if op == admission.Create {
		return true
	}
	if op == admission.Update && utilfeature.DefaultFeatureGate.Enabled(k8sfeatures.ExpandPersistentVolumes) {
		initialized, err := initialization.IsObjectInitialized(a.GetObject())
		if err != nil {
			utilruntime.HandleError(err)
			return true
		}
		return initialized
	}
	initializationCompletion, err := util.IsInitializationCompletion(a)
	if err != nil {
		utilruntime.HandleError(err)
		return true
	}
	return initializationCompletion
}
func (p *pvcEvaluator) Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return generic.Matches(resourceQuota, item, p.MatchingResources, generic.MatchesNoScopeFunc)
}
func (p *pvcEvaluator) MatchingScopes(item runtime.Object, scopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (p *pvcEvaluator) UncoveredQuotaScopes(limitedScopes []corev1.ScopedResourceSelectorRequirement, matchedQuotaScopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (p *pvcEvaluator) MatchingResources(items []corev1.ResourceName) []corev1.ResourceName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := []corev1.ResourceName{}
	for _, item := range items {
		if quota.Contains([]corev1.ResourceName{pvcObjectCountName}, item) {
			result = append(result, item)
			continue
		}
		if quota.Contains(pvcResources, item) {
			result = append(result, item)
			continue
		}
		for _, resource := range pvcResources {
			byStorageClass := storageClassSuffix + string(resource)
			if strings.HasSuffix(string(item), byStorageClass) {
				result = append(result, item)
				break
			}
		}
	}
	return result
}
func (p *pvcEvaluator) Usage(item runtime.Object) (corev1.ResourceList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := corev1.ResourceList{}
	pvc, err := toExternalPersistentVolumeClaimOrError(item)
	if err != nil {
		return result, err
	}
	result[corev1.ResourcePersistentVolumeClaims] = *(resource.NewQuantity(1, resource.DecimalSI))
	result[pvcObjectCountName] = *(resource.NewQuantity(1, resource.DecimalSI))
	if utilfeature.DefaultFeatureGate.Enabled(features.Initializers) {
		if !initialization.IsInitialized(pvc.Initializers) {
			return result, nil
		}
	}
	storageClassRef := helper.GetPersistentVolumeClaimClass(pvc)
	if len(storageClassRef) > 0 {
		storageClassClaim := corev1.ResourceName(storageClassRef + storageClassSuffix + string(corev1.ResourcePersistentVolumeClaims))
		result[storageClassClaim] = *(resource.NewQuantity(1, resource.DecimalSI))
	}
	if request, found := pvc.Spec.Resources.Requests[corev1.ResourceStorage]; found {
		result[corev1.ResourceRequestsStorage] = request
		if len(storageClassRef) > 0 {
			storageClassStorage := corev1.ResourceName(storageClassRef + storageClassSuffix + string(corev1.ResourceRequestsStorage))
			result[storageClassStorage] = request
		}
	}
	return result, nil
}
func (p *pvcEvaluator) UsageStats(options quota.UsageStatsOptions) (quota.UsageStats, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return generic.CalculateUsageStats(options, p.listFuncByNamespace, generic.MatchesNoScopeFunc, p.Usage)
}

var _ quota.Evaluator = &pvcEvaluator{}

func toExternalPersistentVolumeClaimOrError(obj runtime.Object) (*corev1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvc := &corev1.PersistentVolumeClaim{}
	switch t := obj.(type) {
	case *corev1.PersistentVolumeClaim:
		pvc = t
	case *api.PersistentVolumeClaim:
		if err := k8s_api_v1.Convert_core_PersistentVolumeClaim_To_v1_PersistentVolumeClaim(t, pvc, nil); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("expect *api.PersistentVolumeClaim or *v1.PersistentVolumeClaim, got %v", t)
	}
	return pvc, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
