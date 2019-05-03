package v1

import (
	"github.com/openshift/origin/pkg/route/apiserver/admission/apis/ingressadmission"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (obj *IngressAdmissionConfig) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &obj.TypeMeta
}

var GroupVersion = schema.GroupVersion{Group: "route.openshift.io", Version: "v1"}
var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, ingressadmission.Install)
	Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(GroupVersion, &IngressAdmissionConfig{})
	return nil
}
