package image

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	imagev1informer "github.com/openshift/client-go/image/informers/externalversions/image/v1"
	"github.com/openshift/origin/pkg/api/legacy"
)

var legacyObjectCountAliases = map[schema.GroupVersionResource]corev1.ResourceName{imagev1.GroupVersion.WithResource("imagestreams"): imagev1.ResourceImageStreams}

func NewReplenishmentEvaluators(f quota.ListerForResourceFunc, isInformer imagev1informer.ImageStreamInformer, imageClient imagev1typedclient.ImageStreamTagsGetter) []quota.Evaluator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := []quota.Evaluator{NewImageStreamTagEvaluator(isInformer.Lister(), imageClient), NewImageStreamImportEvaluator(isInformer.Lister())}
	for gvr, alias := range legacyObjectCountAliases {
		result = append(result, generic.NewObjectCountEvaluator(gvr.GroupResource(), generic.ListResourceUsingListerFunc(f, gvr), alias))
	}
	return result
}
func NewReplenishmentEvaluatorsForAdmission(isInformer imagev1informer.ImageStreamInformer, imageClient imagev1typedclient.ImageStreamTagsGetter) []quota.Evaluator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := []quota.Evaluator{NewImageStreamTagEvaluator(isInformer.Lister(), imageClient), NewImageStreamImportEvaluator(isInformer.Lister()), &evaluatorForLegacyResource{Evaluator: NewImageStreamTagEvaluator(isInformer.Lister(), imageClient), LegacyGroupResource: legacy.Resource("imagestreamtags")}, &evaluatorForLegacyResource{Evaluator: NewImageStreamImportEvaluator(isInformer.Lister()), LegacyGroupResource: legacy.Resource("imagestreamimports")}}
	for gvr, alias := range legacyObjectCountAliases {
		result = append(result, generic.NewObjectCountEvaluator(gvr.GroupResource(), generic.ListResourceUsingListerFunc(nil, gvr), alias))
	}
	result = append(result, generic.NewObjectCountEvaluator(legacy.Resource("imagestreams"), generic.ListResourceUsingListerFunc(nil, imagev1.GroupVersion.WithResource("imagestreams")), imagev1.ResourceImageStreams))
	return result
}

type evaluatorForLegacyResource struct {
	quota.Evaluator
	LegacyGroupResource	schema.GroupResource
}

func (e *evaluatorForLegacyResource) GroupResource() schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.LegacyGroupResource
}
