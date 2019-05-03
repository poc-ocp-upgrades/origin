package v1

import (
 autoscalingv1 "k8s.io/api/autoscaling/v1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "autoscaling"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &autoscalingv1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs, addConversionFuncs)
}
