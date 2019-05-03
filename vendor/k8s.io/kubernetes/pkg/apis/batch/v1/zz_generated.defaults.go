package v1

import (
 v1 "k8s.io/api/batch/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1.Job{}, func(obj interface{}) {
  SetObjectDefaults_Job(obj.(*v1.Job))
 })
 scheme.AddTypeDefaultingFunc(&v1.JobList{}, func(obj interface{}) {
  SetObjectDefaults_JobList(obj.(*v1.JobList))
 })
 return nil
}
func SetObjectDefaults_Job(in *v1.Job) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_Job(in)
 corev1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
 for i := range in.Spec.Template.Spec.Volumes {
  a := &in.Spec.Template.Spec.Volumes[i]
  corev1.SetDefaults_Volume(a)
  if a.VolumeSource.HostPath != nil {
   corev1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
  }
  if a.VolumeSource.Secret != nil {
   corev1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
  }
  if a.VolumeSource.ISCSI != nil {
   corev1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
  }
  if a.VolumeSource.RBD != nil {
   corev1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
  }
  if a.VolumeSource.DownwardAPI != nil {
   corev1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
   for j := range a.VolumeSource.DownwardAPI.Items {
    b := &a.VolumeSource.DownwardAPI.Items[j]
    if b.FieldRef != nil {
     corev1.SetDefaults_ObjectFieldSelector(b.FieldRef)
    }
   }
  }
  if a.VolumeSource.ConfigMap != nil {
   corev1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
  }
  if a.VolumeSource.AzureDisk != nil {
   corev1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
  }
  if a.VolumeSource.Projected != nil {
   corev1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
   for j := range a.VolumeSource.Projected.Sources {
    b := &a.VolumeSource.Projected.Sources[j]
    if b.DownwardAPI != nil {
     for k := range b.DownwardAPI.Items {
      c := &b.DownwardAPI.Items[k]
      if c.FieldRef != nil {
       corev1.SetDefaults_ObjectFieldSelector(c.FieldRef)
      }
     }
    }
    if b.ServiceAccountToken != nil {
     corev1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
    }
   }
  }
  if a.VolumeSource.ScaleIO != nil {
   corev1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
  }
 }
 for i := range in.Spec.Template.Spec.InitContainers {
  a := &in.Spec.Template.Spec.InitContainers[i]
  corev1.SetDefaults_Container(a)
  for j := range a.Ports {
   b := &a.Ports[j]
   corev1.SetDefaults_ContainerPort(b)
  }
  for j := range a.Env {
   b := &a.Env[j]
   if b.ValueFrom != nil {
    if b.ValueFrom.FieldRef != nil {
     corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
    }
   }
  }
  corev1.SetDefaults_ResourceList(&a.Resources.Limits)
  corev1.SetDefaults_ResourceList(&a.Resources.Requests)
  if a.LivenessProbe != nil {
   corev1.SetDefaults_Probe(a.LivenessProbe)
   if a.LivenessProbe.Handler.HTTPGet != nil {
    corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
   }
  }
  if a.ReadinessProbe != nil {
   corev1.SetDefaults_Probe(a.ReadinessProbe)
   if a.ReadinessProbe.Handler.HTTPGet != nil {
    corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
   }
  }
  if a.Lifecycle != nil {
   if a.Lifecycle.PostStart != nil {
    if a.Lifecycle.PostStart.HTTPGet != nil {
     corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
    }
   }
   if a.Lifecycle.PreStop != nil {
    if a.Lifecycle.PreStop.HTTPGet != nil {
     corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
    }
   }
  }
 }
 for i := range in.Spec.Template.Spec.Containers {
  a := &in.Spec.Template.Spec.Containers[i]
  corev1.SetDefaults_Container(a)
  for j := range a.Ports {
   b := &a.Ports[j]
   corev1.SetDefaults_ContainerPort(b)
  }
  for j := range a.Env {
   b := &a.Env[j]
   if b.ValueFrom != nil {
    if b.ValueFrom.FieldRef != nil {
     corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
    }
   }
  }
  corev1.SetDefaults_ResourceList(&a.Resources.Limits)
  corev1.SetDefaults_ResourceList(&a.Resources.Requests)
  if a.LivenessProbe != nil {
   corev1.SetDefaults_Probe(a.LivenessProbe)
   if a.LivenessProbe.Handler.HTTPGet != nil {
    corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
   }
  }
  if a.ReadinessProbe != nil {
   corev1.SetDefaults_Probe(a.ReadinessProbe)
   if a.ReadinessProbe.Handler.HTTPGet != nil {
    corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
   }
  }
  if a.Lifecycle != nil {
   if a.Lifecycle.PostStart != nil {
    if a.Lifecycle.PostStart.HTTPGet != nil {
     corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
    }
   }
   if a.Lifecycle.PreStop != nil {
    if a.Lifecycle.PreStop.HTTPGet != nil {
     corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
    }
   }
  }
 }
}
func SetObjectDefaults_JobList(in *v1.JobList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_Job(a)
 }
}
