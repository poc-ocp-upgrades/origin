package settings

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *PodPreset) DeepCopyInto(out *PodPreset) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *PodPreset) DeepCopy() *PodPreset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodPreset)
 in.DeepCopyInto(out)
 return out
}
func (in *PodPreset) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodPresetList) DeepCopyInto(out *PodPresetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PodPreset, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodPresetList) DeepCopy() *PodPresetList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodPresetList)
 in.DeepCopyInto(out)
 return out
}
func (in *PodPresetList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodPresetSpec) DeepCopyInto(out *PodPresetSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Selector.DeepCopyInto(&out.Selector)
 if in.Env != nil {
  in, out := &in.Env, &out.Env
  *out = make([]core.EnvVar, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.EnvFrom != nil {
  in, out := &in.EnvFrom, &out.EnvFrom
  *out = make([]core.EnvFromSource, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]core.Volume, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.VolumeMounts != nil {
  in, out := &in.VolumeMounts, &out.VolumeMounts
  *out = make([]core.VolumeMount, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodPresetSpec) DeepCopy() *PodPresetSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodPresetSpec)
 in.DeepCopyInto(out)
 return out
}
