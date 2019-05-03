package v1beta1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/core/v1"
 v1beta1 "k8s.io/api/storage/v1beta1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
 storage "k8s.io/kubernetes/pkg/apis/storage"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.StorageClass)(nil), (*storage.StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StorageClass_To_storage_StorageClass(a.(*v1beta1.StorageClass), b.(*storage.StorageClass), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.StorageClass)(nil), (*v1beta1.StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_StorageClass_To_v1beta1_StorageClass(a.(*storage.StorageClass), b.(*v1beta1.StorageClass), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StorageClassList)(nil), (*storage.StorageClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StorageClassList_To_storage_StorageClassList(a.(*v1beta1.StorageClassList), b.(*storage.StorageClassList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.StorageClassList)(nil), (*v1beta1.StorageClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_StorageClassList_To_v1beta1_StorageClassList(a.(*storage.StorageClassList), b.(*v1beta1.StorageClassList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachment)(nil), (*storage.VolumeAttachment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(a.(*v1beta1.VolumeAttachment), b.(*storage.VolumeAttachment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachment)(nil), (*v1beta1.VolumeAttachment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(a.(*storage.VolumeAttachment), b.(*v1beta1.VolumeAttachment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentList)(nil), (*storage.VolumeAttachmentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(a.(*v1beta1.VolumeAttachmentList), b.(*storage.VolumeAttachmentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentList)(nil), (*v1beta1.VolumeAttachmentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(a.(*storage.VolumeAttachmentList), b.(*v1beta1.VolumeAttachmentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentSource)(nil), (*storage.VolumeAttachmentSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(a.(*v1beta1.VolumeAttachmentSource), b.(*storage.VolumeAttachmentSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentSource)(nil), (*v1beta1.VolumeAttachmentSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(a.(*storage.VolumeAttachmentSource), b.(*v1beta1.VolumeAttachmentSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentSpec)(nil), (*storage.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(a.(*v1beta1.VolumeAttachmentSpec), b.(*storage.VolumeAttachmentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentSpec)(nil), (*v1beta1.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(a.(*storage.VolumeAttachmentSpec), b.(*v1beta1.VolumeAttachmentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentStatus)(nil), (*storage.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(a.(*v1beta1.VolumeAttachmentStatus), b.(*storage.VolumeAttachmentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentStatus)(nil), (*v1beta1.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(a.(*storage.VolumeAttachmentStatus), b.(*v1beta1.VolumeAttachmentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeError)(nil), (*storage.VolumeError)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_VolumeError_To_storage_VolumeError(a.(*v1beta1.VolumeError), b.(*storage.VolumeError), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*storage.VolumeError)(nil), (*v1beta1.VolumeError)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_storage_VolumeError_To_v1beta1_VolumeError(a.(*storage.VolumeError), b.(*v1beta1.VolumeError), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in *v1beta1.StorageClass, out *storage.StorageClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Provisioner = in.Provisioner
 out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
 out.ReclaimPolicy = (*core.PersistentVolumeReclaimPolicy)(unsafe.Pointer(in.ReclaimPolicy))
 out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
 out.AllowVolumeExpansion = (*bool)(unsafe.Pointer(in.AllowVolumeExpansion))
 out.VolumeBindingMode = (*storage.VolumeBindingMode)(unsafe.Pointer(in.VolumeBindingMode))
 out.AllowedTopologies = *(*[]core.TopologySelectorTerm)(unsafe.Pointer(&in.AllowedTopologies))
 return nil
}
func Convert_v1beta1_StorageClass_To_storage_StorageClass(in *v1beta1.StorageClass, out *storage.StorageClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in, out, s)
}
func autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *v1beta1.StorageClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Provisioner = in.Provisioner
 out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
 out.ReclaimPolicy = (*v1.PersistentVolumeReclaimPolicy)(unsafe.Pointer(in.ReclaimPolicy))
 out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
 out.AllowVolumeExpansion = (*bool)(unsafe.Pointer(in.AllowVolumeExpansion))
 out.VolumeBindingMode = (*v1beta1.VolumeBindingMode)(unsafe.Pointer(in.VolumeBindingMode))
 out.AllowedTopologies = *(*[]v1.TopologySelectorTerm)(unsafe.Pointer(&in.AllowedTopologies))
 return nil
}
func Convert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *v1beta1.StorageClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in, out, s)
}
func autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in *v1beta1.StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]storage.StorageClass)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_StorageClassList_To_storage_StorageClassList(in *v1beta1.StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in, out, s)
}
func autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *v1beta1.StorageClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.StorageClass)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *v1beta1.StorageClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in, out, s)
}
func autoConvert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in *v1beta1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in *v1beta1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in, out, s)
}
func autoConvert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in *storage.VolumeAttachment, out *v1beta1.VolumeAttachment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in *storage.VolumeAttachment, out *v1beta1.VolumeAttachment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in, out, s)
}
func autoConvert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1beta1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]storage.VolumeAttachment)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1beta1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in, out, s)
}
func autoConvert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1beta1.VolumeAttachmentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.VolumeAttachment)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1beta1.VolumeAttachmentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in, out, s)
}
func autoConvert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1beta1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
 return nil
}
func Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1beta1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in, out, s)
}
func autoConvert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1beta1.VolumeAttachmentSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
 return nil
}
func Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1beta1.VolumeAttachmentSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in, out, s)
}
func autoConvert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1beta1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Attacher = in.Attacher
 if err := Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
  return err
 }
 out.NodeName = in.NodeName
 return nil
}
func Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1beta1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in, out, s)
}
func autoConvert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1beta1.VolumeAttachmentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Attacher = in.Attacher
 if err := Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
  return err
 }
 out.NodeName = in.NodeName
 return nil
}
func Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1beta1.VolumeAttachmentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in, out, s)
}
func autoConvert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1beta1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Attached = in.Attached
 out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
 out.AttachError = (*storage.VolumeError)(unsafe.Pointer(in.AttachError))
 out.DetachError = (*storage.VolumeError)(unsafe.Pointer(in.DetachError))
 return nil
}
func Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1beta1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in, out, s)
}
func autoConvert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1beta1.VolumeAttachmentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Attached = in.Attached
 out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
 out.AttachError = (*v1beta1.VolumeError)(unsafe.Pointer(in.AttachError))
 out.DetachError = (*v1beta1.VolumeError)(unsafe.Pointer(in.DetachError))
 return nil
}
func Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1beta1.VolumeAttachmentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in, out, s)
}
func autoConvert_v1beta1_VolumeError_To_storage_VolumeError(in *v1beta1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Time = in.Time
 out.Message = in.Message
 return nil
}
func Convert_v1beta1_VolumeError_To_storage_VolumeError(in *v1beta1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_VolumeError_To_storage_VolumeError(in, out, s)
}
func autoConvert_storage_VolumeError_To_v1beta1_VolumeError(in *storage.VolumeError, out *v1beta1.VolumeError, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Time = in.Time
 out.Message = in.Message
 return nil
}
func Convert_storage_VolumeError_To_v1beta1_VolumeError(in *storage.VolumeError, out *v1beta1.VolumeError, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_storage_VolumeError_To_v1beta1_VolumeError(in, out, s)
}
