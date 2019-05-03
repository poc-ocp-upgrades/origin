package v1beta1

import (
 coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "coordination.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &coordinationv1beta1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterDefaults)
}
