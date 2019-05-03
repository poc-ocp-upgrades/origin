package v1

import (
	quotav1 "github.com/openshift/api/quota/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetResourceQuotasStatusByNamespace(namespaceStatuses quotav1.ResourceQuotasStatusByNamespace, namespace string) (corev1.ResourceQuotaStatus, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range namespaceStatuses {
		curr := namespaceStatuses[i]
		if curr.Namespace == namespace {
			return curr.Status, true
		}
	}
	return corev1.ResourceQuotaStatus{}, false
}
func RemoveResourceQuotasStatusByNamespace(namespaceStatuses *quotav1.ResourceQuotasStatusByNamespace, namespace string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newNamespaceStatuses := quotav1.ResourceQuotasStatusByNamespace{}
	for i := range *namespaceStatuses {
		curr := (*namespaceStatuses)[i]
		if curr.Namespace == namespace {
			continue
		}
		newNamespaceStatuses = append(newNamespaceStatuses, curr)
	}
	*namespaceStatuses = newNamespaceStatuses
}
func InsertResourceQuotasStatus(namespaceStatuses *quotav1.ResourceQuotasStatusByNamespace, newStatus quotav1.ResourceQuotaStatusByNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newNamespaceStatuses := quotav1.ResourceQuotasStatusByNamespace{}
	found := false
	for i := range *namespaceStatuses {
		curr := (*namespaceStatuses)[i]
		if curr.Namespace == newStatus.Namespace {
			newNamespaceStatuses = append(newNamespaceStatuses, newStatus)
			found = true
			continue
		}
		newNamespaceStatuses = append(newNamespaceStatuses, curr)
	}
	if !found {
		newNamespaceStatuses = append(newNamespaceStatuses, newStatus)
	}
	*namespaceStatuses = newNamespaceStatuses
}

var accessor = meta.NewAccessor()

func GetMatcher(selector quotav1.ClusterResourceQuotaSelector) (func(obj runtime.Object) (bool, error), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var labelSelector labels.Selector
	if selector.LabelSelector != nil {
		var err error
		labelSelector, err = metav1.LabelSelectorAsSelector(selector.LabelSelector)
		if err != nil {
			return nil, err
		}
	}
	var annotationSelector map[string]string
	if len(selector.AnnotationSelector) > 0 {
		annotationSelector = make(map[string]string, len(selector.AnnotationSelector))
		for k, v := range selector.AnnotationSelector {
			annotationSelector[k] = v
		}
	}
	return func(obj runtime.Object) (bool, error) {
		if labelSelector != nil {
			objLabels, err := accessor.Labels(obj)
			if err != nil {
				return false, err
			}
			if !labelSelector.Matches(labels.Set(objLabels)) {
				return false, nil
			}
		}
		if annotationSelector != nil {
			objAnnotations, err := accessor.Annotations(obj)
			if err != nil {
				return false, err
			}
			for k, v := range annotationSelector {
				if objValue, exists := objAnnotations[k]; !exists || objValue != v {
					return false, nil
				}
			}
		}
		return true, nil
	}, nil
}
func GetObjectMatcher(selector quotav1.ClusterResourceQuotaSelector) (func(obj metav1.Object) (bool, error), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var labelSelector labels.Selector
	if selector.LabelSelector != nil {
		var err error
		labelSelector, err = metav1.LabelSelectorAsSelector(selector.LabelSelector)
		if err != nil {
			return nil, err
		}
	}
	var annotationSelector map[string]string
	if len(selector.AnnotationSelector) > 0 {
		annotationSelector = make(map[string]string, len(selector.AnnotationSelector))
		for k, v := range selector.AnnotationSelector {
			annotationSelector[k] = v
		}
	}
	return func(obj metav1.Object) (bool, error) {
		if labelSelector != nil {
			if !labelSelector.Matches(labels.Set(obj.GetLabels())) {
				return false, nil
			}
		}
		if annotationSelector != nil {
			objAnnotations := obj.GetAnnotations()
			for k, v := range annotationSelector {
				if objValue, exists := objAnnotations[k]; !exists || objValue != v {
					return false, nil
				}
			}
		}
		return true, nil
	}, nil
}
