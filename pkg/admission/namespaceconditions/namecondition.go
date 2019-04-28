package namespaceconditions

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
)

type pluginHandlerWithNamespaceNameConditions struct {
	admissionPlugin		admission.Interface
	namespacesToExclude	sets.String
}

var _ admission.ValidationInterface = &pluginHandlerWithNamespaceNameConditions{}
var _ admission.MutationInterface = &pluginHandlerWithNamespaceNameConditions{}

func (p pluginHandlerWithNamespaceNameConditions) Handles(operation admission.Operation) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.admissionPlugin.Handles(operation)
}
func (p pluginHandlerWithNamespaceNameConditions) Admit(a admission.Attributes) error {
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
func (p pluginHandlerWithNamespaceNameConditions) Validate(a admission.Attributes) error {
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
func (p pluginHandlerWithNamespaceNameConditions) shouldRunAdmission(attr admission.Attributes) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaceName := attr.GetNamespace()
	if p.namespacesToExclude.Has(namespaceName) {
		return false
	}
	if (attr.GetResource().GroupResource() == schema.GroupResource{Resource: "namespaces"}) && p.namespacesToExclude.Has(attr.GetName()) {
		return false
	}
	return true
}
