package namespaceconditions

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1lister "k8s.io/client-go/listers/core/v1"
)

const runLevelLabel = "openshift.io/run-level"

var (
	skipRunLevelZeroSelector	labels.Selector
	skipRunLevelOneSelector		labels.Selector
)

func init() {
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
	var err error
	skipRunLevelZeroSelector, err = labels.Parse(runLevelLabel + " notin ( 0 )")
	if err != nil {
		panic(err)
	}
	skipRunLevelOneSelector, err = labels.Parse(runLevelLabel + " notin ( 0,1 )")
	if err != nil {
		panic(err)
	}
}

type pluginHandlerWithNamespaceLabelConditions struct {
	admissionPlugin		admission.Interface
	namespaceClient		corev1client.NamespacesGetter
	namespaceLister		corev1lister.NamespaceLister
	namespaceSelector	labels.Selector
}

var _ admission.ValidationInterface = &pluginHandlerWithNamespaceLabelConditions{}
var _ admission.MutationInterface = &pluginHandlerWithNamespaceLabelConditions{}

func (p pluginHandlerWithNamespaceLabelConditions) Handles(operation admission.Operation) bool {
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
	return p.admissionPlugin.Handles(operation)
}
func (p pluginHandlerWithNamespaceLabelConditions) Admit(a admission.Attributes) error {
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
	if !p.shouldRunAdmission(a) {
		return nil
	}
	mutatingHandler, ok := p.admissionPlugin.(admission.MutationInterface)
	if !ok {
		return nil
	}
	return mutatingHandler.Admit(a)
}
func (p pluginHandlerWithNamespaceLabelConditions) Validate(a admission.Attributes) error {
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
	if !p.shouldRunAdmission(a) {
		return nil
	}
	validatingHandler, ok := p.admissionPlugin.(admission.ValidationInterface)
	if !ok {
		return nil
	}
	return validatingHandler.Validate(a)
}
func (p pluginHandlerWithNamespaceLabelConditions) shouldRunAdmission(attr admission.Attributes) bool {
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
	namespaceName := attr.GetNamespace()
	if len(namespaceName) == 0 && attr.GetResource().Resource != "namespaces" {
		return true
	}
	namespaceLabels, err := p.getNamespaceLabels(attr)
	if err != nil {
		return true
	}
	return p.namespaceSelector.Matches(labels.Set(namespaceLabels))
}
func (p pluginHandlerWithNamespaceLabelConditions) getNamespaceLabels(attr admission.Attributes) (map[string]string, error) {
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
	if attr.GetResource().Resource == "namespaces" && len(attr.GetSubresource()) == 0 && (attr.GetOperation() == admission.Create || attr.GetOperation() == admission.Update) {
		accessor, err := meta.Accessor(attr.GetObject())
		if err != nil {
			return nil, err
		}
		return accessor.GetLabels(), nil
	}
	namespaceName := attr.GetNamespace()
	namespace, err := p.namespaceLister.Get(namespaceName)
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, err
	}
	if apierrors.IsNotFound(err) {
		namespace, err = p.namespaceClient.Namespaces().Get(namespaceName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	}
	return namespace.Labels, nil
}
