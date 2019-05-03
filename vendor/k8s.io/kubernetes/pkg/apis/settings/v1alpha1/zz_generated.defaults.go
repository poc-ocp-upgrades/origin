package v1alpha1

import (
 v1alpha1 "k8s.io/api/settings/v1alpha1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1alpha1.PodPreset{}, func(obj interface{}) {
  SetObjectDefaults_PodPreset(obj.(*v1alpha1.PodPreset))
 })
 scheme.AddTypeDefaultingFunc(&v1alpha1.PodPresetList{}, func(obj interface{}) {
  SetObjectDefaults_PodPresetList(obj.(*v1alpha1.PodPresetList))
 })
 return nil
}
func SetObjectDefaults_PodPreset(in *v1alpha1.PodPreset) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Spec.Env {
  a := &in.Spec.Env[i]
  if a.ValueFrom != nil {
   if a.ValueFrom.FieldRef != nil {
    v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
   }
  }
 }
 for i := range in.Spec.Volumes {
  a := &in.Spec.Volumes[i]
  v1.SetDefaults_Volume(a)
  if a.VolumeSource.HostPath != nil {
   v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
  }
  if a.VolumeSource.Secret != nil {
   v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
  }
  if a.VolumeSource.ISCSI != nil {
   v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
  }
  if a.VolumeSource.RBD != nil {
   v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
  }
  if a.VolumeSource.DownwardAPI != nil {
   v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
   for j := range a.VolumeSource.DownwardAPI.Items {
    b := &a.VolumeSource.DownwardAPI.Items[j]
    if b.FieldRef != nil {
     v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
    }
   }
  }
  if a.VolumeSource.ConfigMap != nil {
   v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
  }
  if a.VolumeSource.AzureDisk != nil {
   v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
  }
  if a.VolumeSource.Projected != nil {
   v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
   for j := range a.VolumeSource.Projected.Sources {
    b := &a.VolumeSource.Projected.Sources[j]
    if b.DownwardAPI != nil {
     for k := range b.DownwardAPI.Items {
      c := &b.DownwardAPI.Items[k]
      if c.FieldRef != nil {
       v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
      }
     }
    }
    if b.ServiceAccountToken != nil {
     v1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
    }
   }
  }
  if a.VolumeSource.ScaleIO != nil {
   v1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
  }
 }
}
func SetObjectDefaults_PodPresetList(in *v1alpha1.PodPresetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_PodPreset(a)
 }
}
