package v1alpha1

import (
 storagev1alpha1 "k8s.io/api/storage/v1alpha1"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "storage.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
 localSchemeBuilder = &storagev1alpha1.SchemeBuilder
 AddToScheme        = localSchemeBuilder.AddToScheme
)
