package image

import (
	"fmt"
	"github.com/openshift/api/image"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	imagev1lister "github.com/openshift/client-go/image/listers/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imagev1conversions "github.com/openshift/origin/pkg/image/apis/image/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kadmission "k8s.io/apiserver/pkg/admission"
	kquota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
)

var imageStreamTagResources = []corev1.ResourceName{imagev1.ResourceImageStreams}

type imageStreamTagEvaluator struct {
	store     imagev1lister.ImageStreamLister
	istGetter imagev1typedclient.ImageStreamTagsGetter
}

func NewImageStreamTagEvaluator(store imagev1lister.ImageStreamLister, istGetter imagev1typedclient.ImageStreamTagsGetter) kquota.Evaluator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageStreamTagEvaluator{store: store, istGetter: istGetter}
}
func (i *imageStreamTagEvaluator) Constraints(required []corev1.ResourceName, object runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, okInt := object.(*imageapi.ImageStreamTag)
	_, okExt := object.(*imagev1.ImageStreamTag)
	if !okInt && !okExt {
		return fmt.Errorf("unexpected input object %v", object)
	}
	return nil
}
func (i *imageStreamTagEvaluator) GroupResource() schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return image.Resource("imagestreamtags")
}
func (i *imageStreamTagEvaluator) Handles(a kadmission.Attributes) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	operation := a.GetOperation()
	return operation == kadmission.Create || operation == kadmission.Update
}
func (i *imageStreamTagEvaluator) Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	matchesScopeFunc := func(corev1.ScopedResourceSelectorRequirement, runtime.Object) (bool, error) {
		return true, nil
	}
	return generic.Matches(resourceQuota, item, i.MatchingResources, matchesScopeFunc)
}
func (p *imageStreamTagEvaluator) MatchingScopes(item runtime.Object, scopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (p *imageStreamTagEvaluator) UncoveredQuotaScopes(limitedScopes []corev1.ScopedResourceSelectorRequirement, matchedQuotaScopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []corev1.ScopedResourceSelectorRequirement{}, nil
}
func (i *imageStreamTagEvaluator) MatchingResources(input []corev1.ResourceName) []corev1.ResourceName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kquota.Intersection(input, imageStreamTagResources)
}
func (i *imageStreamTagEvaluator) Usage(item runtime.Object) (corev1.ResourceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if istInternal, ok := item.(*imageapi.ImageStreamTag); ok {
		out := &imagev1.ImageStreamTag{}
		if err := imagev1conversions.Convert_image_ImageStreamTag_To_v1_ImageStreamTag(istInternal, out, nil); err != nil {
			return corev1.ResourceList{}, fmt.Errorf("error converting ImageStreamImport: %v", err)
		}
		item = out
	}
	ist, ok := item.(*imagev1.ImageStreamTag)
	if !ok {
		return corev1.ResourceList{}, nil
	}
	res := map[corev1.ResourceName]resource.Quantity{imagev1.ResourceImageStreams: *resource.NewQuantity(0, resource.BinarySI)}
	isName, _, err := imageapi.ParseImageStreamTagName(ist.Name)
	if err != nil {
		return corev1.ResourceList{}, err
	}
	is, err := i.store.ImageStreams(ist.Namespace).Get(isName)
	if err != nil && !kerrors.IsNotFound(err) {
		utilruntime.HandleError(fmt.Errorf("failed to get image stream %s/%s: %v", ist.Namespace, isName, err))
	}
	if is == nil || kerrors.IsNotFound(err) {
		res[imagev1.ResourceImageStreams] = *resource.NewQuantity(1, resource.BinarySI)
	}
	return res, nil
}
func (i *imageStreamTagEvaluator) UsageStats(options kquota.UsageStatsOptions) (kquota.UsageStats, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kquota.UsageStats{}, nil
}
