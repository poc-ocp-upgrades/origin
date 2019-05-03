package v0

import (
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/kubernetes/pkg/apis/abac"
)

const GroupName = "abac.authorization.kubernetes.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v0"}

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := addKnownTypes(abac.Scheme); err != nil {
  panic(err)
 }
 if err := addConversionFuncs(abac.Scheme); err != nil {
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
 localSchemeBuilder.Register(addKnownTypes, addConversionFuncs)
}
func addKnownTypes(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddKnownTypes(SchemeGroupVersion, &Policy{})
 return nil
}
