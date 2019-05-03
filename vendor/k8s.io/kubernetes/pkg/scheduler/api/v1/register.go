package v1

import (
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "", Version: "v1"}

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := addKnownTypes(schedulerapi.Scheme); err != nil {
  panic(err)
 }
}

var (
 SchemeBuilder      runtime.SchemeBuilder
 localSchemeBuilder = &SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addKnownTypes)
}
func addKnownTypes(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddKnownTypes(SchemeGroupVersion, &Policy{})
 return nil
}
