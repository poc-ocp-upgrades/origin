package v1beta1

import (
 certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "certificates.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

func Kind(kind string) schema.GroupKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &certificatesv1beta1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs)
}
