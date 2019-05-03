package core

import (
 "k8s.io/apimachinery/pkg/runtime/schema"
)

func (obj *ObjectReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}
func (obj *ObjectReference) GroupVersionKind() schema.GroupVersionKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return schema.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}
func (obj *ObjectReference) GetObjectKind() schema.ObjectKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return obj
}
