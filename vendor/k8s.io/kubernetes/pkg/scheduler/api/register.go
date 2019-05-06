package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var Scheme = runtime.NewScheme()
var SchemeGroupVersion = schema.GroupVersion{Group: "", Version: runtime.APIVersionInternal}
var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := addKnownTypes(Scheme); err != nil {
		panic(err)
	}
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddIgnoredConversionType(&metav1.TypeMeta{}, &metav1.TypeMeta{}); err != nil {
		return err
	}
	scheme.AddKnownTypes(SchemeGroupVersion, &Policy{})
	return nil
}
