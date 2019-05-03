package v1beta1

import (
 authorizationv1beta1 "k8s.io/api/authorization/v1beta1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "authorization.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &authorizationv1beta1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs, addConversionFuncs)
}
