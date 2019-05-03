package v1alpha1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/core/v1"
 v1alpha1 "k8s.io/api/settings/v1alpha1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
 settings "k8s.io/kubernetes/pkg/apis/settings"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PodPreset)(nil), (*settings.PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PodPreset_To_settings_PodPreset(a.(*v1alpha1.PodPreset), b.(*settings.PodPreset), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*settings.PodPreset)(nil), (*v1alpha1.PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_settings_PodPreset_To_v1alpha1_PodPreset(a.(*settings.PodPreset), b.(*v1alpha1.PodPreset), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PodPresetList)(nil), (*settings.PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(a.(*v1alpha1.PodPresetList), b.(*settings.PodPresetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*settings.PodPresetList)(nil), (*v1alpha1.PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(a.(*settings.PodPresetList), b.(*v1alpha1.PodPresetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PodPresetSpec)(nil), (*settings.PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(a.(*v1alpha1.PodPresetSpec), b.(*settings.PodPresetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*settings.PodPresetSpec)(nil), (*v1alpha1.PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(a.(*settings.PodPresetSpec), b.(*v1alpha1.PodPresetSpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in *v1alpha1.PodPreset, out *settings.PodPreset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_PodPreset_To_settings_PodPreset(in *v1alpha1.PodPreset, out *settings.PodPreset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in, out, s)
}
func autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *v1alpha1.PodPreset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *v1alpha1.PodPreset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in, out, s)
}
func autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *v1alpha1.PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]settings.PodPreset, len(*in))
  for i := range *in {
   if err := Convert_v1alpha1_PodPreset_To_settings_PodPreset(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *v1alpha1.PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in, out, s)
}
func autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *v1alpha1.PodPresetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1alpha1.PodPreset, len(*in))
  for i := range *in {
   if err := Convert_settings_PodPreset_To_v1alpha1_PodPreset(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *v1alpha1.PodPresetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in, out, s)
}
func autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *v1alpha1.PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = in.Selector
 out.Env = *(*[]core.EnvVar)(unsafe.Pointer(&in.Env))
 out.EnvFrom = *(*[]core.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]core.Volume, len(*in))
  for i := range *in {
   if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
    return err
   }
  }
 } else {
  out.Volumes = nil
 }
 out.VolumeMounts = *(*[]core.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
 return nil
}
func Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *v1alpha1.PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in, out, s)
}
func autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *v1alpha1.PodPresetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = in.Selector
 out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
 out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]v1.Volume, len(*in))
  for i := range *in {
   if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
    return err
   }
  }
 } else {
  out.Volumes = nil
 }
 out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
 return nil
}
func Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *v1alpha1.PodPresetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in, out, s)
}
