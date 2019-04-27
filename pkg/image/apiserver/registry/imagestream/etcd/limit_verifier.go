package etcd

import (
	"fmt"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1informers "k8s.io/client-go/informers/core/v1"
	imageadmission "github.com/openshift/origin/pkg/image/apiserver/admission/limitrange"
)

func ImageLimitVerifier(limitRangeInformer corev1informers.LimitRangeInformer) imageadmission.LimitVerifier {
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
	limitRangeInformer.Informer()
	return imageadmission.NewLimitVerifier(imageadmission.LimitRangesForNamespaceFunc(func(ns string) ([]*corev1.LimitRange, error) {
		list, err := limitRangeInformer.Lister().LimitRanges(ns).List(labels.Everything())
		if err != nil {
			return nil, err
		}
		if len(list) == 0 && len(limitRangeInformer.Informer().LastSyncResourceVersion()) == 0 {
			klog.V(4).Infof("LimitVerifier still waiting for ranges to load: %#v", limitRangeInformer.Informer())
			forbiddenErr := apierrors.NewForbidden(schema.GroupResource{Resource: "limitranges"}, "", fmt.Errorf("the server is still loading limit information"))
			forbiddenErr.ErrStatus.Details.RetryAfterSeconds = 1
			return nil, forbiddenErr
		}
		return list, nil
	}))
}
