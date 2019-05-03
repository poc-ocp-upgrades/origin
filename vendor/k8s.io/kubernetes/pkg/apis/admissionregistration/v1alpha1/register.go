package v1alpha1

import (
 admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "admissionregistration.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &admissionregistrationv1alpha1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs)
}
