package v1

import (
 appsv1 "k8s.io/api/apps/v1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "apps"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &appsv1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs, addConversionFuncs)
}
