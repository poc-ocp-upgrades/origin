package v1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/apps/v1"
 corev1 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 apps "k8s.io/kubernetes/pkg/apis/apps"
 core "k8s.io/kubernetes/pkg/apis/core"
 apiscorev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1.ControllerRevision)(nil), (*apps.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ControllerRevision_To_apps_ControllerRevision(a.(*v1.ControllerRevision), b.(*apps.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevision)(nil), (*v1.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevision_To_v1_ControllerRevision(a.(*apps.ControllerRevision), b.(*v1.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ControllerRevisionList)(nil), (*apps.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ControllerRevisionList_To_apps_ControllerRevisionList(a.(*v1.ControllerRevisionList), b.(*apps.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevisionList)(nil), (*v1.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevisionList_To_v1_ControllerRevisionList(a.(*apps.ControllerRevisionList), b.(*v1.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSet)(nil), (*apps.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSet_To_apps_DaemonSet(a.(*v1.DaemonSet), b.(*apps.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSet)(nil), (*v1.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSet_To_v1_DaemonSet(a.(*apps.DaemonSet), b.(*v1.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSetCondition)(nil), (*apps.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetCondition_To_apps_DaemonSetCondition(a.(*v1.DaemonSetCondition), b.(*apps.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetCondition)(nil), (*v1.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetCondition_To_v1_DaemonSetCondition(a.(*apps.DaemonSetCondition), b.(*v1.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSetList)(nil), (*apps.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetList_To_apps_DaemonSetList(a.(*v1.DaemonSetList), b.(*apps.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetList)(nil), (*v1.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetList_To_v1_DaemonSetList(a.(*apps.DaemonSetList), b.(*v1.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSetSpec)(nil), (*apps.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(a.(*v1.DaemonSetSpec), b.(*apps.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetSpec)(nil), (*v1.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(a.(*apps.DaemonSetSpec), b.(*v1.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSetStatus)(nil), (*apps.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetStatus_To_apps_DaemonSetStatus(a.(*v1.DaemonSetStatus), b.(*apps.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetStatus)(nil), (*v1.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetStatus_To_v1_DaemonSetStatus(a.(*apps.DaemonSetStatus), b.(*v1.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonSetUpdateStrategy)(nil), (*apps.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(a.(*v1.DaemonSetUpdateStrategy), b.(*apps.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetUpdateStrategy)(nil), (*v1.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(a.(*apps.DaemonSetUpdateStrategy), b.(*v1.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Deployment)(nil), (*apps.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Deployment_To_apps_Deployment(a.(*v1.Deployment), b.(*apps.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.Deployment)(nil), (*v1.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_Deployment_To_v1_Deployment(a.(*apps.Deployment), b.(*v1.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DeploymentCondition)(nil), (*apps.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentCondition_To_apps_DeploymentCondition(a.(*v1.DeploymentCondition), b.(*apps.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentCondition)(nil), (*v1.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentCondition_To_v1_DeploymentCondition(a.(*apps.DeploymentCondition), b.(*v1.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DeploymentList)(nil), (*apps.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentList_To_apps_DeploymentList(a.(*v1.DeploymentList), b.(*apps.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentList)(nil), (*v1.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentList_To_v1_DeploymentList(a.(*apps.DeploymentList), b.(*v1.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentSpec)(nil), (*v1.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DeploymentStatus)(nil), (*apps.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentStatus_To_apps_DeploymentStatus(a.(*v1.DeploymentStatus), b.(*apps.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStatus)(nil), (*v1.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStatus_To_v1_DeploymentStatus(a.(*apps.DeploymentStatus), b.(*v1.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStrategy)(nil), (*v1.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicaSet)(nil), (*apps.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSet_To_apps_ReplicaSet(a.(*v1.ReplicaSet), b.(*apps.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSet)(nil), (*v1.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSet_To_v1_ReplicaSet(a.(*apps.ReplicaSet), b.(*v1.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicaSetCondition)(nil), (*apps.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSetCondition_To_apps_ReplicaSetCondition(a.(*v1.ReplicaSetCondition), b.(*apps.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetCondition)(nil), (*v1.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetCondition_To_v1_ReplicaSetCondition(a.(*apps.ReplicaSetCondition), b.(*v1.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicaSetList)(nil), (*apps.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSetList_To_apps_ReplicaSetList(a.(*v1.ReplicaSetList), b.(*apps.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetList)(nil), (*v1.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetList_To_v1_ReplicaSetList(a.(*apps.ReplicaSetList), b.(*v1.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicaSetStatus)(nil), (*apps.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSetStatus_To_apps_ReplicaSetStatus(a.(*v1.ReplicaSetStatus), b.(*apps.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetStatus)(nil), (*v1.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetStatus_To_v1_ReplicaSetStatus(a.(*apps.ReplicaSetStatus), b.(*v1.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RollingUpdateStatefulSetStrategy)(nil), (*apps.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(a.(*v1.RollingUpdateStatefulSetStrategy), b.(*apps.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateStatefulSetStrategy)(nil), (*v1.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateStatefulSetStrategy_To_v1_RollingUpdateStatefulSetStrategy(a.(*apps.RollingUpdateStatefulSetStrategy), b.(*v1.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSet)(nil), (*apps.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSet_To_apps_StatefulSet(a.(*v1.StatefulSet), b.(*apps.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSet)(nil), (*v1.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSet_To_v1_StatefulSet(a.(*apps.StatefulSet), b.(*v1.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSetCondition)(nil), (*apps.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetCondition_To_apps_StatefulSetCondition(a.(*v1.StatefulSetCondition), b.(*apps.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetCondition)(nil), (*v1.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetCondition_To_v1_StatefulSetCondition(a.(*apps.StatefulSetCondition), b.(*v1.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSetList)(nil), (*apps.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetList_To_apps_StatefulSetList(a.(*v1.StatefulSetList), b.(*apps.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetList)(nil), (*v1.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetList_To_v1_StatefulSetList(a.(*apps.StatefulSetList), b.(*v1.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetSpec)(nil), (*v1.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSetStatus)(nil), (*apps.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(a.(*v1.StatefulSetStatus), b.(*apps.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetStatus)(nil), (*v1.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(a.(*apps.StatefulSetStatus), b.(*v1.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSetSpec)(nil), (*v1.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(a.(*apps.DaemonSetSpec), b.(*v1.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSetUpdateStrategy)(nil), (*v1.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(a.(*apps.DaemonSetUpdateStrategy), b.(*v1.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DaemonSet)(nil), (*v1.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSet_To_v1_DaemonSet(a.(*apps.DaemonSet), b.(*v1.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentSpec)(nil), (*v1.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentStrategy)(nil), (*v1.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.Deployment)(nil), (*v1.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_Deployment_To_v1_Deployment(a.(*apps.Deployment), b.(*v1.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetSpec)(nil), (*v1.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetStatus)(nil), (*v1.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(a.(*apps.StatefulSetStatus), b.(*v1.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.DaemonSetSpec)(nil), (*apps.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(a.(*v1.DaemonSetSpec), b.(*apps.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.DaemonSetUpdateStrategy)(nil), (*apps.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(a.(*v1.DaemonSetUpdateStrategy), b.(*apps.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.DaemonSet)(nil), (*apps.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonSet_To_apps_DaemonSet(a.(*v1.DaemonSet), b.(*apps.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.Deployment)(nil), (*apps.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Deployment_To_apps_Deployment(a.(*v1.Deployment), b.(*apps.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.StatefulSetStatus)(nil), (*apps.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(a.(*v1.StatefulSetStatus), b.(*apps.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_ControllerRevision_To_apps_ControllerRevision(in *v1.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_v1_ControllerRevision_To_apps_ControllerRevision(in *v1.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ControllerRevision_To_apps_ControllerRevision(in, out, s)
}
func autoConvert_apps_ControllerRevision_To_v1_ControllerRevision(in *apps.ControllerRevision, out *v1.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_apps_ControllerRevision_To_v1_ControllerRevision(in *apps.ControllerRevision, out *v1.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevision_To_v1_ControllerRevision(in, out, s)
}
func autoConvert_v1_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_v1_ControllerRevision_To_apps_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ControllerRevisionList_To_apps_ControllerRevisionList(in, out, s)
}
func autoConvert_apps_ControllerRevisionList_To_v1_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_apps_ControllerRevision_To_v1_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ControllerRevisionList_To_v1_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevisionList_To_v1_ControllerRevisionList(in, out, s)
}
func autoConvert_v1_DaemonSet_To_apps_DaemonSet(in *v1.DaemonSet, out *apps.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_DaemonSetStatus_To_apps_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_DaemonSet_To_v1_DaemonSet(in *apps.DaemonSet, out *v1.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetStatus_To_v1_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DaemonSetCondition_To_apps_DaemonSetCondition(in, out, s)
}
func autoConvert_apps_DaemonSetCondition_To_v1_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.DaemonSetConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DaemonSetCondition_To_v1_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetCondition_To_v1_DaemonSetCondition(in, out, s)
}
func autoConvert_v1_DaemonSetList_To_apps_DaemonSetList(in *v1.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_v1_DaemonSet_To_apps_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_DaemonSetList_To_apps_DaemonSetList(in *v1.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DaemonSetList_To_apps_DaemonSetList(in, out, s)
}
func autoConvert_apps_DaemonSetList_To_v1_DaemonSetList(in *apps.DaemonSetList, out *v1.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_apps_DaemonSet_To_v1_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DaemonSetList_To_v1_DaemonSetList(in *apps.DaemonSetList, out *v1.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetList_To_v1_DaemonSetList(in, out, s)
}
func autoConvert_v1_DaemonSetSpec_To_apps_DaemonSetSpec(in *v1.DaemonSetSpec, out *apps.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_apps_DaemonSetSpec_To_v1_DaemonSetSpec(in *apps.DaemonSetSpec, out *v1.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_v1_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
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
func Convert_v1_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DaemonSetStatus_To_apps_DaemonSetStatus(in, out, s)
}
func autoConvert_apps_DaemonSetStatus_To_v1_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1.DaemonSetStatus, s conversion.Scope) error {
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
 out.Conditions = *(*[]v1.DaemonSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_DaemonSetStatus_To_v1_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetStatus_To_v1_DaemonSetStatus(in, out, s)
}
func autoConvert_v1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in *v1.DaemonSetUpdateStrategy, out *apps.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDaemonSet)
  if err := Convert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(in *apps.DaemonSetUpdateStrategy, out *v1.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1.RollingUpdateDaemonSet)
  if err := Convert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_v1_Deployment_To_apps_Deployment(in *v1.Deployment, out *apps.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_DeploymentSpec_To_apps_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_DeploymentStatus_To_apps_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_Deployment_To_v1_Deployment(in *apps.Deployment, out *v1.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DeploymentSpec_To_v1_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStatus_To_v1_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_DeploymentCondition_To_apps_DeploymentCondition(in *v1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
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
func Convert_v1_DeploymentCondition_To_apps_DeploymentCondition(in *v1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DeploymentCondition_To_apps_DeploymentCondition(in, out, s)
}
func autoConvert_apps_DeploymentCondition_To_v1_DeploymentCondition(in *apps.DeploymentCondition, out *v1.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.DeploymentConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastUpdateTime = in.LastUpdateTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DeploymentCondition_To_v1_DeploymentCondition(in *apps.DeploymentCondition, out *v1.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentCondition_To_v1_DeploymentCondition(in, out, s)
}
func autoConvert_v1_DeploymentList_To_apps_DeploymentList(in *v1.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.Deployment, len(*in))
  for i := range *in {
   if err := Convert_v1_Deployment_To_apps_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_DeploymentList_To_apps_DeploymentList(in *v1.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DeploymentList_To_apps_DeploymentList(in, out, s)
}
func autoConvert_apps_DeploymentList_To_v1_DeploymentList(in *apps.DeploymentList, out *v1.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.Deployment, len(*in))
  for i := range *in {
   if err := Convert_apps_Deployment_To_v1_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DeploymentList_To_v1_DeploymentList(in *apps.DeploymentList, out *v1.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentList_To_v1_DeploymentList(in, out, s)
}
func autoConvert_v1_DeploymentSpec_To_apps_DeploymentSpec(in *v1.DeploymentSpec, out *apps.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_apps_DeploymentSpec_To_v1_DeploymentSpec(in *apps.DeploymentSpec, out *v1.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_v1_DeploymentStatus_To_apps_DeploymentStatus(in *v1.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
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
func Convert_v1_DeploymentStatus_To_apps_DeploymentStatus(in *v1.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DeploymentStatus_To_apps_DeploymentStatus(in, out, s)
}
func autoConvert_apps_DeploymentStatus_To_v1_DeploymentStatus(in *apps.DeploymentStatus, out *v1.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.Replicas = in.Replicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.UnavailableReplicas = in.UnavailableReplicas
 out.Conditions = *(*[]v1.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 return nil
}
func Convert_apps_DeploymentStatus_To_v1_DeploymentStatus(in *apps.DeploymentStatus, out *v1.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentStatus_To_v1_DeploymentStatus(in, out, s)
}
func autoConvert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(in *v1.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDeployment)
  if err := Convert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(in *apps.DeploymentStrategy, out *v1.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1.RollingUpdateDeployment)
  if err := Convert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_v1_ReplicaSet_To_apps_ReplicaSet(in *v1.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ReplicaSetStatus_To_apps_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ReplicaSet_To_apps_ReplicaSet(in *v1.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicaSet_To_apps_ReplicaSet(in, out, s)
}
func autoConvert_apps_ReplicaSet_To_v1_ReplicaSet(in *apps.ReplicaSet, out *v1.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_ReplicaSetStatus_To_v1_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_ReplicaSet_To_v1_ReplicaSet(in *apps.ReplicaSet, out *v1.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSet_To_v1_ReplicaSet(in, out, s)
}
func autoConvert_v1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.ReplicaSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in, out, s)
}
func autoConvert_apps_ReplicaSetCondition_To_v1_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.ReplicaSetConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_ReplicaSetCondition_To_v1_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetCondition_To_v1_ReplicaSetCondition(in, out, s)
}
func autoConvert_v1_ReplicaSetList_To_apps_ReplicaSetList(in *v1.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_v1_ReplicaSet_To_apps_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_ReplicaSetList_To_apps_ReplicaSetList(in *v1.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicaSetList_To_apps_ReplicaSetList(in, out, s)
}
func autoConvert_apps_ReplicaSetList_To_v1_ReplicaSetList(in *apps.ReplicaSetList, out *v1.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_apps_ReplicaSet_To_v1_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ReplicaSetList_To_v1_ReplicaSetList(in *apps.ReplicaSetList, out *v1.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetList_To_v1_ReplicaSetList(in, out, s)
}
func autoConvert_v1_ReplicaSetSpec_To_apps_ReplicaSetSpec(in *v1.ReplicaSetSpec, out *apps.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_apps_ReplicaSetSpec_To_v1_ReplicaSetSpec(in *apps.ReplicaSetSpec, out *v1.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
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
func Convert_v1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in, out, s)
}
func autoConvert_apps_ReplicaSetStatus_To_v1_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]v1.ReplicaSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_ReplicaSetStatus_To_v1_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetStatus_To_v1_ReplicaSetStatus(in, out, s)
}
func autoConvert_v1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(in *v1.RollingUpdateDaemonSet, out *apps.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(in *apps.RollingUpdateDaemonSet, out *v1.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(in *v1.RollingUpdateDeployment, out *apps.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(in *apps.RollingUpdateDeployment, out *v1.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_RollingUpdateStatefulSetStrategy_To_v1_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_v1_StatefulSet_To_apps_StatefulSet(in *v1.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_StatefulSet_To_apps_StatefulSet(in *v1.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_StatefulSet_To_apps_StatefulSet(in, out, s)
}
func autoConvert_apps_StatefulSet_To_v1_StatefulSet(in *apps.StatefulSet, out *v1.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_StatefulSet_To_v1_StatefulSet(in *apps.StatefulSet, out *v1.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSet_To_v1_StatefulSet(in, out, s)
}
func autoConvert_v1_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_StatefulSetCondition_To_apps_StatefulSetCondition(in, out, s)
}
func autoConvert_apps_StatefulSetCondition_To_v1_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.StatefulSetConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_StatefulSetCondition_To_v1_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetCondition_To_v1_StatefulSetCondition(in, out, s)
}
func autoConvert_v1_StatefulSetList_To_apps_StatefulSetList(in *v1.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_v1_StatefulSet_To_apps_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_StatefulSetList_To_apps_StatefulSetList(in *v1.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_StatefulSetList_To_apps_StatefulSetList(in, out, s)
}
func autoConvert_apps_StatefulSetList_To_v1_StatefulSetList(in *apps.StatefulSetList, out *v1.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_apps_StatefulSet_To_v1_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_StatefulSetList_To_v1_StatefulSetList(in *apps.StatefulSetList, out *v1.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetList_To_v1_StatefulSetList(in, out, s)
}
func autoConvert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(in *v1.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.VolumeClaimTemplates = *(*[]core.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = apps.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(in *apps.StatefulSetSpec, out *v1.StatefulSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.VolumeClaimTemplates = *(*[]corev1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
 out.ServiceName = in.ServiceName
 out.PodManagementPolicy = v1.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
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
func autoConvert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1.StatefulSetStatus, s conversion.Scope) error {
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
 out.Conditions = *(*[]v1.StatefulSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func autoConvert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *v1.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateStatefulSetStrategy)
  if err := Convert_v1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *v1.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1.RollingUpdateStatefulSetStrategy)
  if err := Convert_apps_RollingUpdateStatefulSetStrategy_To_v1_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
