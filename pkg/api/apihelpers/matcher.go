package apihelpers

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func InternalListOptionsToSelectors(options *metainternal.ListOptions) (labels.Selector, fields.Selector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	label := labels.Everything()
	if options != nil && options.LabelSelector != nil {
		label = options.LabelSelector
	}
	field := fields.Everything()
	if options != nil && options.FieldSelector != nil {
		field = options.FieldSelector
	}
	return label, field
}
