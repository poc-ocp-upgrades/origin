package project

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/core"
)

const (
	GroupName = "project.openshift.io"
)

var (
	schemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme)
	Install			= schemeBuilder.AddToScheme
	SchemeGroupVersion	= schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	AddToScheme		= schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
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
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
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
	scheme.AddKnownTypes(SchemeGroupVersion, &Project{}, &ProjectList{}, &ProjectRequest{})
	return nil
}
