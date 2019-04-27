package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "config.templateservicebroker.openshift.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}
var (
	SchemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme		= SchemeBuilder.AddToScheme
	localSchemeBuilder	= &SchemeBuilder
)

func addKnownTypes(scheme *runtime.Scheme) error {
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
	scheme.AddKnownTypes(SchemeGroupVersion, &TemplateServiceBrokerConfig{})
	return nil
}
