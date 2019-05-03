package config

import (
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "kubecontrollermanager.config.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
var (
 SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
 AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddKnownTypes(SchemeGroupVersion, &KubeControllerManagerConfiguration{})
 return nil
}
