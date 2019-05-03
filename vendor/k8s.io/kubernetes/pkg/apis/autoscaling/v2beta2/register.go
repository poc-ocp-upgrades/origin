package v2beta2

import (
 autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "autoscaling"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v2beta2"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &autoscalingv2beta2.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(addDefaultingFuncs)
}
