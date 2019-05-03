package v1beta2

import (
 appsv1beta2 "k8s.io/api/apps/v1beta2"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "apps"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta2"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &appsv1beta2.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs, addConversionFuncs)
}
