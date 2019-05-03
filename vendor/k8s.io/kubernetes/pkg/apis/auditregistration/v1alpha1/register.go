package v1alpha1

import (
 auditregistrationv1alpha1 "k8s.io/api/auditregistration/v1alpha1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "auditregistration.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &auditregistrationv1alpha1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs)
}
