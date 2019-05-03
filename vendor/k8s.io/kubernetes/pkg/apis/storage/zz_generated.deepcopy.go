package storage

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *StorageClass) DeepCopyInto(out *StorageClass) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Parameters != nil {
  in, out := &in.Parameters, &out.Parameters
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.ReclaimPolicy != nil {
  in, out := &in.ReclaimPolicy, &out.ReclaimPolicy
  *out = new(core.PersistentVolumeReclaimPolicy)
  **out = **in
 }
 if in.MountOptions != nil {
  in, out := &in.MountOptions, &out.MountOptions
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.AllowVolumeExpansion != nil {
  in, out := &in.AllowVolumeExpansion, &out.AllowVolumeExpansion
  *out = new(bool)
  **out = **in
 }
 if in.VolumeBindingMode != nil {
  in, out := &in.VolumeBindingMode, &out.VolumeBindingMode
  *out = new(VolumeBindingMode)
  **out = **in
 }
 if in.AllowedTopologies != nil {
  in, out := &in.AllowedTopologies, &out.AllowedTopologies
  *out = make([]core.TopologySelectorTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *StorageClass) DeepCopy() *StorageClass {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StorageClass)
 in.DeepCopyInto(out)
 return out
}
func (in *StorageClass) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *StorageClassList) DeepCopyInto(out *StorageClassList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]StorageClass, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *StorageClassList) DeepCopy() *StorageClassList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StorageClassList)
 in.DeepCopyInto(out)
 return out
}
func (in *StorageClassList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *VolumeAttachment) DeepCopyInto(out *VolumeAttachment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *VolumeAttachment) DeepCopy() *VolumeAttachment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeAttachment)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeAttachment) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *VolumeAttachmentList) DeepCopyInto(out *VolumeAttachmentList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]VolumeAttachment, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *VolumeAttachmentList) DeepCopy() *VolumeAttachmentList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeAttachmentList)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeAttachmentList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *VolumeAttachmentSource) DeepCopyInto(out *VolumeAttachmentSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.PersistentVolumeName != nil {
  in, out := &in.PersistentVolumeName, &out.PersistentVolumeName
  *out = new(string)
  **out = **in
 }
 return
}
func (in *VolumeAttachmentSource) DeepCopy() *VolumeAttachmentSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeAttachmentSource)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeAttachmentSpec) DeepCopyInto(out *VolumeAttachmentSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Source.DeepCopyInto(&out.Source)
 return
}
func (in *VolumeAttachmentSpec) DeepCopy() *VolumeAttachmentSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeAttachmentSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeAttachmentStatus) DeepCopyInto(out *VolumeAttachmentStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.AttachmentMetadata != nil {
  in, out := &in.AttachmentMetadata, &out.AttachmentMetadata
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.AttachError != nil {
  in, out := &in.AttachError, &out.AttachError
  *out = new(VolumeError)
  (*in).DeepCopyInto(*out)
 }
 if in.DetachError != nil {
  in, out := &in.DetachError, &out.DetachError
  *out = new(VolumeError)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *VolumeAttachmentStatus) DeepCopy() *VolumeAttachmentStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeAttachmentStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeError) DeepCopyInto(out *VolumeError) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Time.DeepCopyInto(&out.Time)
 return
}
func (in *VolumeError) DeepCopy() *VolumeError {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeError)
 in.DeepCopyInto(out)
 return out
}
