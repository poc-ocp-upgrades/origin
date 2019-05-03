package v1alpha1

import (
 settingsv1alpha1 "k8s.io/api/settings/v1alpha1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "settings.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &settingsv1alpha1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterDefaults)
}
