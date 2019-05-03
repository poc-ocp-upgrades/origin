package v2alpha1

import (
 batchv2alpha1 "k8s.io/api/batch/v2alpha1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "batch"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v2alpha1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &batchv2alpha1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs, addConversionFuncs)
}
