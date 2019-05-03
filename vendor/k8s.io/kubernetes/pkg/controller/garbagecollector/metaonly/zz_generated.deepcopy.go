package metaonly

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *MetadataOnlyObject) DeepCopyInto(out *MetadataOnlyObject) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 return
}
func (in *MetadataOnlyObject) DeepCopy() *MetadataOnlyObject {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetadataOnlyObject)
 in.DeepCopyInto(out)
 return out
}
func (in *MetadataOnlyObject) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *MetadataOnlyObjectList) DeepCopyInto(out *MetadataOnlyObjectList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]MetadataOnlyObject, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *MetadataOnlyObjectList) DeepCopy() *MetadataOnlyObjectList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetadataOnlyObjectList)
 in.DeepCopyInto(out)
 return out
}
func (in *MetadataOnlyObjectList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
