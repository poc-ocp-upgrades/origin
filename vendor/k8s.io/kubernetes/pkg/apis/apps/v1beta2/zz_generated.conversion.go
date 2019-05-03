package v1beta2

import (
 unsafe "unsafe"
 v1beta2 "k8s.io/api/apps/v1beta2"
 v1 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 apps "k8s.io/kubernetes/pkg/apis/apps"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 core "k8s.io/kubernetes/pkg/apis/core"
 corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta2.ControllerRevision)(nil), (*apps.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ControllerRevision_To_apps_ControllerRevision(a.(*v1beta2.ControllerRevision), b.(*apps.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevision)(nil), (*v1beta2.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevision_To_v1beta2_ControllerRevision(a.(*apps.ControllerRevision), b.(*v1beta2.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ControllerRevisionList)(nil), (*apps.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ControllerRevisionList_To_apps_ControllerRevisionList(a.(*v1beta2.ControllerRevisionList), b.(*apps.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevisionList)(nil), (*v1beta2.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevisionList_To_v1beta2_ControllerRevisionList(a.(*apps.ControllerRevisionList), b.(*v1beta2.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSet)(nil), (*apps.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSet_To_apps_DaemonSet(a.(*v1beta2.DaemonSet), b.(*apps.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSet)(nil), (*v1beta2.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSet_To_v1beta2_DaemonSet(a.(*apps.DaemonSet), b.(*v1beta2.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSetCondition)(nil), (*apps.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetCondition_To_apps_DaemonSetCondition(a.(*v1beta2.DaemonSetCondition), b.(*apps.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetCondition)(nil), (*v1beta2.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetCondition_To_v1beta2_DaemonSetCondition(a.(*apps.DaemonSetCondition), b.(*v1beta2.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSetList)(nil), (*apps.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetList_To_apps_DaemonSetList(a.(*v1beta2.DaemonSetList), b.(*apps.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetList)(nil), (*v1beta2.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetList_To_v1beta2_DaemonSetList(a.(*apps.DaemonSetList), b.(*v1beta2.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSetSpec)(nil), (*apps.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetSpec_To_apps_DaemonSetSpec(a.(*v1beta2.DaemonSetSpec), b.(*apps.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetSpec)(nil), (*v1beta2.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetSpec_To_v1beta2_DaemonSetSpec(a.(*apps.DaemonSetSpec), b.(*v1beta2.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSetStatus)(nil), (*apps.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetStatus_To_apps_DaemonSetStatus(a.(*v1beta2.DaemonSetStatus), b.(*apps.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetStatus)(nil), (*v1beta2.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetStatus_To_v1beta2_DaemonSetStatus(a.(*apps.DaemonSetStatus), b.(*v1beta2.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DaemonSetUpdateStrategy)(nil), (*apps.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(a.(*v1beta2.DaemonSetUpdateStrategy), b.(*apps.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetUpdateStrategy)(nil), (*v1beta2.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetUpdateStrategy_To_v1beta2_DaemonSetUpdateStrategy(a.(*apps.DaemonSetUpdateStrategy), b.(*v1beta2.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.Deployment)(nil), (*apps.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_Deployment_To_apps_Deployment(a.(*v1beta2.Deployment), b.(*apps.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.Deployment)(nil), (*v1beta2.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_Deployment_To_v1beta2_Deployment(a.(*apps.Deployment), b.(*v1beta2.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DeploymentCondition)(nil), (*apps.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentCondition_To_apps_DeploymentCondition(a.(*v1beta2.DeploymentCondition), b.(*apps.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentCondition)(nil), (*v1beta2.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentCondition_To_v1beta2_DeploymentCondition(a.(*apps.DeploymentCondition), b.(*v1beta2.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DeploymentList)(nil), (*apps.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentList_To_apps_DeploymentList(a.(*v1beta2.DeploymentList), b.(*apps.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentList)(nil), (*v1beta2.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentList_To_v1beta2_DeploymentList(a.(*apps.DeploymentList), b.(*v1beta2.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1beta2.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentSpec)(nil), (*v1beta2.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1beta2_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1beta2.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DeploymentStatus)(nil), (*apps.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentStatus_To_apps_DeploymentStatus(a.(*v1beta2.DeploymentStatus), b.(*apps.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStatus)(nil), (*v1beta2.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStatus_To_v1beta2_DeploymentStatus(a.(*apps.DeploymentStatus), b.(*v1beta2.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1beta2.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStrategy)(nil), (*v1beta2.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1beta2_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1beta2.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ReplicaSet)(nil), (*apps.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSet_To_apps_ReplicaSet(a.(*v1beta2.ReplicaSet), b.(*apps.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSet)(nil), (*v1beta2.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSet_To_v1beta2_ReplicaSet(a.(*apps.ReplicaSet), b.(*v1beta2.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ReplicaSetCondition)(nil), (*apps.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSetCondition_To_apps_ReplicaSetCondition(a.(*v1beta2.ReplicaSetCondition), b.(*apps.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetCondition)(nil), (*v1beta2.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetCondition_To_v1beta2_ReplicaSetCondition(a.(*apps.ReplicaSetCondition), b.(*v1beta2.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ReplicaSetList)(nil), (*apps.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSetList_To_apps_ReplicaSetList(a.(*v1beta2.ReplicaSetList), b.(*apps.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetList)(nil), (*v1beta2.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetList_To_v1beta2_ReplicaSetList(a.(*apps.ReplicaSetList), b.(*v1beta2.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1beta2.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1beta2.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1beta2_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1beta2.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ReplicaSetStatus)(nil), (*apps.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSetStatus_To_apps_ReplicaSetStatus(a.(*v1beta2.ReplicaSetStatus), b.(*apps.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetStatus)(nil), (*v1beta2.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetStatus_To_v1beta2_ReplicaSetStatus(a.(*apps.ReplicaSetStatus), b.(*v1beta2.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1beta2.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1beta2.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1beta2_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1beta2.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1beta2.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1beta2.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1beta2_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1beta2.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.RollingUpdateStatefulSetStrategy)(nil), (*apps.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(a.(*v1beta2.RollingUpdateStatefulSetStrategy), b.(*apps.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateStatefulSetStrategy)(nil), (*v1beta2.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(a.(*apps.RollingUpdateStatefulSetStrategy), b.(*v1beta2.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.Scale)(nil), (*autoscaling.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_Scale_To_autoscaling_Scale(a.(*v1beta2.Scale), b.(*autoscaling.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.Scale)(nil), (*v1beta2.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_Scale_To_v1beta2_Scale(a.(*autoscaling.Scale), b.(*v1beta2.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ScaleSpec)(nil), (*autoscaling.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ScaleSpec_To_autoscaling_ScaleSpec(a.(*v1beta2.ScaleSpec), b.(*autoscaling.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleSpec)(nil), (*v1beta2.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleSpec_To_v1beta2_ScaleSpec(a.(*autoscaling.ScaleSpec), b.(*v1beta2.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.ScaleStatus)(nil), (*autoscaling.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ScaleStatus_To_autoscaling_ScaleStatus(a.(*v1beta2.ScaleStatus), b.(*autoscaling.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1beta2.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1beta2_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1beta2.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSet)(nil), (*apps.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSet_To_apps_StatefulSet(a.(*v1beta2.StatefulSet), b.(*apps.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSet)(nil), (*v1beta2.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSet_To_v1beta2_StatefulSet(a.(*apps.StatefulSet), b.(*v1beta2.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSetCondition)(nil), (*apps.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetCondition_To_apps_StatefulSetCondition(a.(*v1beta2.StatefulSetCondition), b.(*apps.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetCondition)(nil), (*v1beta2.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetCondition_To_v1beta2_StatefulSetCondition(a.(*apps.StatefulSetCondition), b.(*v1beta2.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSetList)(nil), (*apps.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetList_To_apps_StatefulSetList(a.(*v1beta2.StatefulSetList), b.(*apps.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetList)(nil), (*v1beta2.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetList_To_v1beta2_StatefulSetList(a.(*apps.StatefulSetList), b.(*v1beta2.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1beta2.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetSpec)(nil), (*v1beta2.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1beta2.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSetStatus)(nil), (*apps.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(a.(*v1beta2.StatefulSetStatus), b.(*apps.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetStatus)(nil), (*v1beta2.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(a.(*apps.StatefulSetStatus), b.(*v1beta2.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta2.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1beta2.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1beta2.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1beta2.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSetSpec)(nil), (*v1beta2.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetSpec_To_v1beta2_DaemonSetSpec(a.(*apps.DaemonSetSpec), b.(*v1beta2.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSetUpdateStrategy)(nil), (*v1beta2.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetUpdateStrategy_To_v1beta2_DaemonSetUpdateStrategy(a.(*apps.DaemonSetUpdateStrategy), b.(*v1beta2.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSet)(nil), (*v1beta2.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSet_To_v1beta2_DaemonSet(a.(*apps.DaemonSet), b.(*v1beta2.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentSpec)(nil), (*v1beta2.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1beta2_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1beta2.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentStrategy)(nil), (*v1beta2.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1beta2_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1beta2.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.Deployment)(nil), (*v1beta2.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_Deployment_To_v1beta2_Deployment(a.(*apps.Deployment), b.(*v1beta2.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1beta2.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1beta2_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1beta2.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1beta2.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1beta2_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1beta2.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1beta2.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1beta2_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1beta2.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetSpec)(nil), (*v1beta2.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1beta2.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetStatus)(nil), (*v1beta2.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(a.(*apps.StatefulSetStatus), b.(*v1beta2.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1beta2.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1beta2.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1beta2.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1beta2_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1beta2.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.DaemonSetSpec)(nil), (*apps.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetSpec_To_apps_DaemonSetSpec(a.(*v1beta2.DaemonSetSpec), b.(*apps.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.DaemonSetUpdateStrategy)(nil), (*apps.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(a.(*v1beta2.DaemonSetUpdateStrategy), b.(*apps.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.DaemonSet)(nil), (*apps.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DaemonSet_To_apps_DaemonSet(a.(*v1beta2.DaemonSet), b.(*apps.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1beta2.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1beta2.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.Deployment)(nil), (*apps.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_Deployment_To_apps_Deployment(a.(*v1beta2.Deployment), b.(*apps.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1beta2.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1beta2.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1beta2.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.ScaleStatus)(nil), (*autoscaling.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_ScaleStatus_To_autoscaling_ScaleStatus(a.(*v1beta2.ScaleStatus), b.(*autoscaling.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1beta2.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.StatefulSetStatus)(nil), (*apps.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(a.(*v1beta2.StatefulSetStatus), b.(*apps.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta2.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1beta2.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta2_ControllerRevision_To_apps_ControllerRevision(in *v1beta2.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_v1beta2_ControllerRevision_To_apps_ControllerRevision(in *v1beta2.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ControllerRevision_To_apps_ControllerRevision(in, out, s)
}
func autoConvert_apps_ControllerRevision_To_v1beta2_ControllerRevision(in *apps.ControllerRevision, out *v1beta2.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_apps_ControllerRevision_To_v1beta2_ControllerRevision(in *apps.ControllerRevision, out *v1beta2.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevision_To_v1beta2_ControllerRevision(in, out, s)
}
func autoConvert_v1beta2_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1beta2.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_v1beta2_ControllerRevision_To_apps_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta2_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1beta2.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ControllerRevisionList_To_apps_ControllerRevisionList(in, out, s)
}
func autoConvert_apps_ControllerRevisionList_To_v1beta2_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1beta2.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta2.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_apps_ControllerRevision_To_v1beta2_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ControllerRevisionList_To_v1beta2_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1beta2.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevisionList_To_v1beta2_ControllerRevisionList(in, out, s)
}
func autoConvert_v1beta2_DaemonSet_To_apps_DaemonSet(in *v1beta2.DaemonSet, out *apps.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta2_DaemonSetSpec_To_apps_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_DaemonSetStatus_To_apps_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_DaemonSet_To_v1beta2_DaemonSet(in *apps.DaemonSet, out *v1beta2.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DaemonSetSpec_To_v1beta2_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetStatus_To_v1beta2_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta2_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1beta2.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta2_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1beta2.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DaemonSetCondition_To_apps_DaemonSetCondition(in, out, s)
}
func autoConvert_apps_DaemonSetCondition_To_v1beta2_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1beta2.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.DaemonSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DaemonSetCondition_To_v1beta2_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1beta2.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetCondition_To_v1beta2_DaemonSetCondition(in, out, s)
}
func autoConvert_v1beta2_DaemonSetList_To_apps_DaemonSetList(in *v1beta2.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta2_DaemonSet_To_apps_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta2_DaemonSetList_To_apps_DaemonSetList(in *v1beta2.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DaemonSetList_To_apps_DaemonSetList(in, out, s)
}
func autoConvert_apps_DaemonSetList_To_v1beta2_DaemonSetList(in *apps.DaemonSetList, out *v1beta2.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta2.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_apps_DaemonSet_To_v1beta2_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DaemonSetList_To_v1beta2_DaemonSetList(in *apps.DaemonSetList, out *v1beta2.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetList_To_v1beta2_DaemonSetList(in, out, s)
}
func autoConvert_v1beta2_DaemonSetSpec_To_apps_DaemonSetSpec(in *v1beta2.DaemonSetSpec, out *apps.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_apps_DaemonSetSpec_To_v1beta2_DaemonSetSpec(in *apps.DaemonSetSpec, out *v1beta2.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetUpdateStrategy_To_v1beta2_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_v1beta2_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1beta2.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CurrentNumberScheduled = in.CurrentNumberScheduled
 out.NumberMisscheduled = in.NumberMisscheduled
 out.DesiredNumberScheduled = in.DesiredNumberScheduled
 out.NumberReady = in.NumberReady
 out.ObservedGeneration = in.ObservedGeneration
 out.UpdatedNumberScheduled = in.UpdatedNumberScheduled
 out.NumberAvailable = in.NumberAvailable
 out.NumberUnavailable = in.NumberUnavailable
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 out.Conditions = *(*[]apps.DaemonSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_v1beta2_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1beta2.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DaemonSetStatus_To_apps_DaemonSetStatus(in, out, s)
}
func autoConvert_apps_DaemonSetStatus_To_v1beta2_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1beta2.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CurrentNumberScheduled = in.CurrentNumberScheduled
 out.NumberMisscheduled = in.NumberMisscheduled
 out.DesiredNumberScheduled = in.DesiredNumberScheduled
 out.NumberReady = in.NumberReady
 out.ObservedGeneration = in.ObservedGeneration
 out.UpdatedNumberScheduled = in.UpdatedNumberScheduled
 out.NumberAvailable = in.NumberAvailable
 out.NumberUnavailable = in.NumberUnavailable
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 out.Conditions = *(*[]v1beta2.DaemonSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_DaemonSetStatus_To_v1beta2_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1beta2.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetStatus_To_v1beta2_DaemonSetStatus(in, out, s)
}
func autoConvert_v1beta2_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in *v1beta2.DaemonSetUpdateStrategy, out *apps.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDaemonSet)
  if err := Convert_v1beta2_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_DaemonSetUpdateStrategy_To_v1beta2_DaemonSetUpdateStrategy(in *apps.DaemonSetUpdateStrategy, out *v1beta2.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta2.RollingUpdateDaemonSet)
  if err := Convert_apps_RollingUpdateDaemonSet_To_v1beta2_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_v1beta2_Deployment_To_apps_Deployment(in *v1beta2.Deployment, out *apps.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta2_DeploymentSpec_To_apps_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_DeploymentStatus_To_apps_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_Deployment_To_v1beta2_Deployment(in *apps.Deployment, out *v1beta2.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DeploymentSpec_To_v1beta2_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStatus_To_v1beta2_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta2_DeploymentCondition_To_apps_DeploymentCondition(in *v1beta2.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DeploymentConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastUpdateTime = in.LastUpdateTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta2_DeploymentCondition_To_apps_DeploymentCondition(in *v1beta2.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DeploymentCondition_To_apps_DeploymentCondition(in, out, s)
}
func autoConvert_apps_DeploymentCondition_To_v1beta2_DeploymentCondition(in *apps.DeploymentCondition, out *v1beta2.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.DeploymentConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastUpdateTime = in.LastUpdateTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DeploymentCondition_To_v1beta2_DeploymentCondition(in *apps.DeploymentCondition, out *v1beta2.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentCondition_To_v1beta2_DeploymentCondition(in, out, s)
}
func autoConvert_v1beta2_DeploymentList_To_apps_DeploymentList(in *v1beta2.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.Deployment, len(*in))
  for i := range *in {
   if err := Convert_v1beta2_Deployment_To_apps_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta2_DeploymentList_To_apps_DeploymentList(in *v1beta2.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DeploymentList_To_apps_DeploymentList(in, out, s)
}
func autoConvert_apps_DeploymentList_To_v1beta2_DeploymentList(in *apps.DeploymentList, out *v1beta2.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta2.Deployment, len(*in))
  for i := range *in {
   if err := Convert_apps_Deployment_To_v1beta2_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DeploymentList_To_v1beta2_DeploymentList(in *apps.DeploymentList, out *v1beta2.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentList_To_v1beta2_DeploymentList(in, out, s)
}
func autoConvert_v1beta2_DeploymentSpec_To_apps_DeploymentSpec(in *v1beta2.DeploymentSpec, out *apps.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_DeploymentStrategy_To_apps_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_apps_DeploymentSpec_To_v1beta2_DeploymentSpec(in *apps.DeploymentSpec, out *v1beta2.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStrategy_To_v1beta2_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_v1beta2_DeploymentStatus_To_apps_DeploymentStatus(in *v1beta2.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.Replicas = in.Replicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.UnavailableReplicas = in.UnavailableReplicas
 out.Conditions = *(*[]apps.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 return nil
}
func Convert_v1beta2_DeploymentStatus_To_apps_DeploymentStatus(in *v1beta2.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_DeploymentStatus_To_apps_DeploymentStatus(in, out, s)
}
func autoConvert_apps_DeploymentStatus_To_v1beta2_DeploymentStatus(in *apps.DeploymentStatus, out *v1beta2.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.Replicas = in.Replicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.UnavailableReplicas = in.UnavailableReplicas
 out.Conditions = *(*[]v1beta2.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 return nil
}
func Convert_apps_DeploymentStatus_To_v1beta2_DeploymentStatus(in *apps.DeploymentStatus, out *v1beta2.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentStatus_To_v1beta2_DeploymentStatus(in, out, s)
}
func autoConvert_v1beta2_DeploymentStrategy_To_apps_DeploymentStrategy(in *v1beta2.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDeployment)
  if err := Convert_v1beta2_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_DeploymentStrategy_To_v1beta2_DeploymentStrategy(in *apps.DeploymentStrategy, out *v1beta2.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta2.RollingUpdateDeployment)
  if err := Convert_apps_RollingUpdateDeployment_To_v1beta2_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_v1beta2_ReplicaSet_To_apps_ReplicaSet(in *v1beta2.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta2_ReplicaSetSpec_To_apps_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_ReplicaSetStatus_To_apps_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta2_ReplicaSet_To_apps_ReplicaSet(in *v1beta2.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ReplicaSet_To_apps_ReplicaSet(in, out, s)
}
func autoConvert_apps_ReplicaSet_To_v1beta2_ReplicaSet(in *apps.ReplicaSet, out *v1beta2.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_ReplicaSetSpec_To_v1beta2_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_ReplicaSetStatus_To_v1beta2_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_ReplicaSet_To_v1beta2_ReplicaSet(in *apps.ReplicaSet, out *v1beta2.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSet_To_v1beta2_ReplicaSet(in, out, s)
}
func autoConvert_v1beta2_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1beta2.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.ReplicaSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta2_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1beta2.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ReplicaSetCondition_To_apps_ReplicaSetCondition(in, out, s)
}
func autoConvert_apps_ReplicaSetCondition_To_v1beta2_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1beta2.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.ReplicaSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_ReplicaSetCondition_To_v1beta2_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1beta2.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetCondition_To_v1beta2_ReplicaSetCondition(in, out, s)
}
func autoConvert_v1beta2_ReplicaSetList_To_apps_ReplicaSetList(in *v1beta2.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta2_ReplicaSet_To_apps_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta2_ReplicaSetList_To_apps_ReplicaSetList(in *v1beta2.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ReplicaSetList_To_apps_ReplicaSetList(in, out, s)
}
func autoConvert_apps_ReplicaSetList_To_v1beta2_ReplicaSetList(in *apps.ReplicaSetList, out *v1beta2.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta2.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_apps_ReplicaSet_To_v1beta2_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ReplicaSetList_To_v1beta2_ReplicaSetList(in *apps.ReplicaSetList, out *v1beta2.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetList_To_v1beta2_ReplicaSetList(in, out, s)
}
func autoConvert_v1beta2_ReplicaSetSpec_To_apps_ReplicaSetSpec(in *v1beta2.ReplicaSetSpec, out *apps.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_ReplicaSetSpec_To_v1beta2_ReplicaSetSpec(in *apps.ReplicaSetSpec, out *v1beta2.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta2_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1beta2.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]apps.ReplicaSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_v1beta2_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1beta2.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ReplicaSetStatus_To_apps_ReplicaSetStatus(in, out, s)
}
func autoConvert_apps_ReplicaSetStatus_To_v1beta2_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1beta2.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]v1beta2.ReplicaSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_ReplicaSetStatus_To_v1beta2_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1beta2.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetStatus_To_v1beta2_ReplicaSetStatus(in, out, s)
}
func autoConvert_v1beta2_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(in *v1beta2.RollingUpdateDaemonSet, out *apps.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDaemonSet_To_v1beta2_RollingUpdateDaemonSet(in *apps.RollingUpdateDaemonSet, out *v1beta2.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1beta2_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(in *v1beta2.RollingUpdateDeployment, out *apps.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDeployment_To_v1beta2_RollingUpdateDeployment(in *apps.RollingUpdateDeployment, out *v1beta2.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta2.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta2.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta2.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta2.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_v1beta2_Scale_To_autoscaling_Scale(in *v1beta2.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta2_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta2_Scale_To_autoscaling_Scale(in *v1beta2.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_Scale_To_autoscaling_Scale(in, out, s)
}
func autoConvert_autoscaling_Scale_To_v1beta2_Scale(in *autoscaling.Scale, out *v1beta2.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_autoscaling_ScaleSpec_To_v1beta2_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_autoscaling_ScaleStatus_To_v1beta2_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_autoscaling_Scale_To_v1beta2_Scale(in *autoscaling.Scale, out *v1beta2.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_Scale_To_v1beta2_Scale(in, out, s)
}
func autoConvert_v1beta2_ScaleSpec_To_autoscaling_ScaleSpec(in *v1beta2.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_v1beta2_ScaleSpec_To_autoscaling_ScaleSpec(in *v1beta2.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}
func autoConvert_autoscaling_ScaleSpec_To_v1beta2_ScaleSpec(in *autoscaling.ScaleSpec, out *v1beta2.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_autoscaling_ScaleSpec_To_v1beta2_ScaleSpec(in *autoscaling.ScaleSpec, out *v1beta2.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_ScaleSpec_To_v1beta2_ScaleSpec(in, out, s)
}
func autoConvert_v1beta2_ScaleStatus_To_autoscaling_ScaleStatus(in *v1beta2.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func autoConvert_autoscaling_ScaleStatus_To_v1beta2_ScaleStatus(in *autoscaling.ScaleStatus, out *v1beta2.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func autoConvert_v1beta2_StatefulSet_To_apps_StatefulSet(in *v1beta2.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta2_StatefulSet_To_apps_StatefulSet(in *v1beta2.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_StatefulSet_To_apps_StatefulSet(in, out, s)
}
func autoConvert_apps_StatefulSet_To_v1beta2_StatefulSet(in *apps.StatefulSet, out *v1beta2.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_StatefulSet_To_v1beta2_StatefulSet(in *apps.StatefulSet, out *v1beta2.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSet_To_v1beta2_StatefulSet(in, out, s)
}
func autoConvert_v1beta2_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1beta2.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta2_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1beta2.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_StatefulSetCondition_To_apps_StatefulSetCondition(in, out, s)
}
func autoConvert_apps_StatefulSetCondition_To_v1beta2_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1beta2.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.StatefulSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_StatefulSetCondition_To_v1beta2_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1beta2.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetCondition_To_v1beta2_StatefulSetCondition(in, out, s)
}
func autoConvert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in *v1beta2.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta2_StatefulSet_To_apps_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in *v1beta2.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in, out, s)
}
func autoConvert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in *apps.StatefulSetList, out *v1beta2.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta2.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_apps_StatefulSet_To_v1beta2_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in *apps.StatefulSetList, out *v1beta2.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in, out, s)
}
func autoConvert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(in *v1beta2.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.VolumeClaimTemplates = *(*[]core.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = apps.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(in *apps.StatefulSetSpec, out *v1beta2.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.VolumeClaimTemplates = *(*[]v1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = v1beta2.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1beta2.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int64_To_Pointer_int64(&in.ObservedGeneration, &out.ObservedGeneration, s); err != nil {
  return err
 }
 out.Replicas = in.Replicas
 out.ReadyReplicas = in.ReadyReplicas
 out.CurrentReplicas = in.CurrentReplicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.CurrentRevision = in.CurrentRevision
 out.UpdateRevision = in.UpdateRevision
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 out.Conditions = *(*[]apps.StatefulSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func autoConvert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1beta2.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int64_To_int64(&in.ObservedGeneration, &out.ObservedGeneration, s); err != nil {
  return err
 }
 out.Replicas = in.Replicas
 out.ReadyReplicas = in.ReadyReplicas
 out.CurrentReplicas = in.CurrentReplicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.CurrentRevision = in.CurrentRevision
 out.UpdateRevision = in.UpdateRevision
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 out.Conditions = *(*[]v1beta2.StatefulSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func autoConvert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *v1beta2.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateStatefulSetStrategy)
  if err := Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *v1beta2.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta2.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta2.RollingUpdateStatefulSetStrategy)
  if err := Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
