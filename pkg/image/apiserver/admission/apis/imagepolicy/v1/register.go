package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (obj *ImagePolicyConfig) GetObjectKind() schema.ObjectKind {
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
	return &obj.TypeMeta
}

var GroupVersion = schema.GroupVersion{Group: "image.openshift.io", Version: "v1"}
var (
	schemeBuilder	= runtime.NewSchemeBuilder(addKnownTypes, addDefaultingFuncs)
	Install		= schemeBuilder.AddToScheme
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
	scheme.AddKnownTypes(GroupVersion, &ImagePolicyConfig{})
	return nil
}
