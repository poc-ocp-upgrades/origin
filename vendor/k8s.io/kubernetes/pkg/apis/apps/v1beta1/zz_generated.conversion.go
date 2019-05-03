package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/apps/v1beta1"
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.ControllerRevision)(nil), (*apps.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ControllerRevision_To_apps_ControllerRevision(a.(*v1beta1.ControllerRevision), b.(*apps.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevision)(nil), (*v1beta1.ControllerRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevision_To_v1beta1_ControllerRevision(a.(*apps.ControllerRevision), b.(*v1beta1.ControllerRevision), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ControllerRevisionList)(nil), (*apps.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ControllerRevisionList_To_apps_ControllerRevisionList(a.(*v1beta1.ControllerRevisionList), b.(*apps.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ControllerRevisionList)(nil), (*v1beta1.ControllerRevisionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ControllerRevisionList_To_v1beta1_ControllerRevisionList(a.(*apps.ControllerRevisionList), b.(*v1beta1.ControllerRevisionList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.Deployment)(nil), (*apps.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Deployment_To_apps_Deployment(a.(*v1beta1.Deployment), b.(*apps.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.Deployment)(nil), (*v1beta1.Deployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_Deployment_To_v1beta1_Deployment(a.(*apps.Deployment), b.(*v1beta1.Deployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentCondition)(nil), (*apps.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentCondition_To_apps_DeploymentCondition(a.(*v1beta1.DeploymentCondition), b.(*apps.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentCondition)(nil), (*v1beta1.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentCondition_To_v1beta1_DeploymentCondition(a.(*apps.DeploymentCondition), b.(*v1beta1.DeploymentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentList)(nil), (*apps.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentList_To_apps_DeploymentList(a.(*v1beta1.DeploymentList), b.(*apps.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentList)(nil), (*v1beta1.DeploymentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentList_To_v1beta1_DeploymentList(a.(*apps.DeploymentList), b.(*v1beta1.DeploymentList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentRollback)(nil), (*apps.DeploymentRollback)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentRollback_To_apps_DeploymentRollback(a.(*v1beta1.DeploymentRollback), b.(*apps.DeploymentRollback), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentRollback)(nil), (*v1beta1.DeploymentRollback)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentRollback_To_v1beta1_DeploymentRollback(a.(*apps.DeploymentRollback), b.(*v1beta1.DeploymentRollback), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1beta1.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentSpec)(nil), (*v1beta1.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1beta1_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1beta1.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentStatus)(nil), (*apps.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentStatus_To_apps_DeploymentStatus(a.(*v1beta1.DeploymentStatus), b.(*apps.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStatus)(nil), (*v1beta1.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStatus_To_v1beta1_DeploymentStatus(a.(*apps.DeploymentStatus), b.(*v1beta1.DeploymentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1beta1.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DeploymentStrategy)(nil), (*v1beta1.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1beta1_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1beta1.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RollbackConfig)(nil), (*apps.RollbackConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollbackConfig_To_apps_RollbackConfig(a.(*v1beta1.RollbackConfig), b.(*apps.RollbackConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollbackConfig)(nil), (*v1beta1.RollbackConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollbackConfig_To_v1beta1_RollbackConfig(a.(*apps.RollbackConfig), b.(*v1beta1.RollbackConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1beta1.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1beta1.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1beta1_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1beta1.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RollingUpdateStatefulSetStrategy)(nil), (*apps.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(a.(*v1beta1.RollingUpdateStatefulSetStrategy), b.(*apps.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateStatefulSetStrategy)(nil), (*v1beta1.RollingUpdateStatefulSetStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta1_RollingUpdateStatefulSetStrategy(a.(*apps.RollingUpdateStatefulSetStrategy), b.(*v1beta1.RollingUpdateStatefulSetStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.Scale)(nil), (*autoscaling.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Scale_To_autoscaling_Scale(a.(*v1beta1.Scale), b.(*autoscaling.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.Scale)(nil), (*v1beta1.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_Scale_To_v1beta1_Scale(a.(*autoscaling.Scale), b.(*v1beta1.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ScaleSpec)(nil), (*autoscaling.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ScaleSpec_To_autoscaling_ScaleSpec(a.(*v1beta1.ScaleSpec), b.(*autoscaling.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleSpec)(nil), (*v1beta1.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleSpec_To_v1beta1_ScaleSpec(a.(*autoscaling.ScaleSpec), b.(*v1beta1.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ScaleStatus)(nil), (*autoscaling.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ScaleStatus_To_autoscaling_ScaleStatus(a.(*v1beta1.ScaleStatus), b.(*autoscaling.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1beta1.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1beta1_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1beta1.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSet)(nil), (*apps.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSet_To_apps_StatefulSet(a.(*v1beta1.StatefulSet), b.(*apps.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSet)(nil), (*v1beta1.StatefulSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSet_To_v1beta1_StatefulSet(a.(*apps.StatefulSet), b.(*v1beta1.StatefulSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSetCondition)(nil), (*apps.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetCondition_To_apps_StatefulSetCondition(a.(*v1beta1.StatefulSetCondition), b.(*apps.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetCondition)(nil), (*v1beta1.StatefulSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetCondition_To_v1beta1_StatefulSetCondition(a.(*apps.StatefulSetCondition), b.(*v1beta1.StatefulSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSetList)(nil), (*apps.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetList_To_apps_StatefulSetList(a.(*v1beta1.StatefulSetList), b.(*apps.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetList)(nil), (*v1beta1.StatefulSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetList_To_v1beta1_StatefulSetList(a.(*apps.StatefulSetList), b.(*v1beta1.StatefulSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1beta1.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetSpec)(nil), (*v1beta1.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1beta1.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSetStatus)(nil), (*apps.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(a.(*v1beta1.StatefulSetStatus), b.(*apps.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetStatus)(nil), (*v1beta1.StatefulSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(a.(*apps.StatefulSetStatus), b.(*v1beta1.StatefulSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1beta1.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1beta1.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1beta1_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1beta1.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentSpec)(nil), (*v1beta1.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentSpec_To_v1beta1_DeploymentSpec(a.(*apps.DeploymentSpec), b.(*v1beta1.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.DeploymentStrategy)(nil), (*v1beta1.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DeploymentStrategy_To_v1beta1_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1beta1.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1beta1.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1beta1_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1beta1.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetSpec)(nil), (*v1beta1.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(a.(*apps.StatefulSetSpec), b.(*v1beta1.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.StatefulSetUpdateStrategy)(nil), (*v1beta1.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_StatefulSetUpdateStrategy_To_v1beta1_StatefulSetUpdateStrategy(a.(*apps.StatefulSetUpdateStrategy), b.(*v1beta1.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1beta1.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1beta1_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1beta1.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.DeploymentSpec)(nil), (*apps.DeploymentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentSpec_To_apps_DeploymentSpec(a.(*v1beta1.DeploymentSpec), b.(*apps.DeploymentSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1beta1.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.RollingUpdateDeployment)(nil), (*apps.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(a.(*v1beta1.RollingUpdateDeployment), b.(*apps.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.ScaleStatus)(nil), (*autoscaling.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ScaleStatus_To_autoscaling_ScaleStatus(a.(*v1beta1.ScaleStatus), b.(*autoscaling.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.StatefulSetSpec)(nil), (*apps.StatefulSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(a.(*v1beta1.StatefulSetSpec), b.(*apps.StatefulSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.StatefulSetUpdateStrategy)(nil), (*apps.StatefulSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(a.(*v1beta1.StatefulSetUpdateStrategy), b.(*apps.StatefulSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_ControllerRevision_To_apps_ControllerRevision(in *v1beta1.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_v1beta1_ControllerRevision_To_apps_ControllerRevision(in *v1beta1.ControllerRevision, out *apps.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ControllerRevision_To_apps_ControllerRevision(in, out, s)
}
func autoConvert_apps_ControllerRevision_To_v1beta1_ControllerRevision(in *apps.ControllerRevision, out *v1beta1.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Data, &out.Data, s); err != nil {
  return err
 }
 out.Revision = in.Revision
 return nil
}
func Convert_apps_ControllerRevision_To_v1beta1_ControllerRevision(in *apps.ControllerRevision, out *v1beta1.ControllerRevision, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevision_To_v1beta1_ControllerRevision(in, out, s)
}
func autoConvert_v1beta1_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1beta1.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_ControllerRevision_To_apps_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_ControllerRevisionList_To_apps_ControllerRevisionList(in *v1beta1.ControllerRevisionList, out *apps.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ControllerRevisionList_To_apps_ControllerRevisionList(in, out, s)
}
func autoConvert_apps_ControllerRevisionList_To_v1beta1_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1beta1.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.ControllerRevision, len(*in))
  for i := range *in {
   if err := Convert_apps_ControllerRevision_To_v1beta1_ControllerRevision(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ControllerRevisionList_To_v1beta1_ControllerRevisionList(in *apps.ControllerRevisionList, out *v1beta1.ControllerRevisionList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ControllerRevisionList_To_v1beta1_ControllerRevisionList(in, out, s)
}
func autoConvert_v1beta1_Deployment_To_apps_Deployment(in *v1beta1.Deployment, out *apps.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_DeploymentSpec_To_apps_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_DeploymentStatus_To_apps_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_Deployment_To_apps_Deployment(in *v1beta1.Deployment, out *apps.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Deployment_To_apps_Deployment(in, out, s)
}
func autoConvert_apps_Deployment_To_v1beta1_Deployment(in *apps.Deployment, out *v1beta1.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DeploymentSpec_To_v1beta1_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStatus_To_v1beta1_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_Deployment_To_v1beta1_Deployment(in *apps.Deployment, out *v1beta1.Deployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_Deployment_To_v1beta1_Deployment(in, out, s)
}
func autoConvert_v1beta1_DeploymentCondition_To_apps_DeploymentCondition(in *v1beta1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
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
func Convert_v1beta1_DeploymentCondition_To_apps_DeploymentCondition(in *v1beta1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DeploymentCondition_To_apps_DeploymentCondition(in, out, s)
}
func autoConvert_apps_DeploymentCondition_To_v1beta1_DeploymentCondition(in *apps.DeploymentCondition, out *v1beta1.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.DeploymentConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastUpdateTime = in.LastUpdateTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DeploymentCondition_To_v1beta1_DeploymentCondition(in *apps.DeploymentCondition, out *v1beta1.DeploymentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentCondition_To_v1beta1_DeploymentCondition(in, out, s)
}
func autoConvert_v1beta1_DeploymentList_To_apps_DeploymentList(in *v1beta1.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.Deployment, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_Deployment_To_apps_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_DeploymentList_To_apps_DeploymentList(in *v1beta1.DeploymentList, out *apps.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DeploymentList_To_apps_DeploymentList(in, out, s)
}
func autoConvert_apps_DeploymentList_To_v1beta1_DeploymentList(in *apps.DeploymentList, out *v1beta1.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.Deployment, len(*in))
  for i := range *in {
   if err := Convert_apps_Deployment_To_v1beta1_Deployment(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DeploymentList_To_v1beta1_DeploymentList(in *apps.DeploymentList, out *v1beta1.DeploymentList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentList_To_v1beta1_DeploymentList(in, out, s)
}
func autoConvert_v1beta1_DeploymentRollback_To_apps_DeploymentRollback(in *v1beta1.DeploymentRollback, out *apps.DeploymentRollback, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.UpdatedAnnotations = *(*map[string]string)(unsafe.Pointer(&in.UpdatedAnnotations))
 if err := Convert_v1beta1_RollbackConfig_To_apps_RollbackConfig(&in.RollbackTo, &out.RollbackTo, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_DeploymentRollback_To_apps_DeploymentRollback(in *v1beta1.DeploymentRollback, out *apps.DeploymentRollback, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DeploymentRollback_To_apps_DeploymentRollback(in, out, s)
}
func autoConvert_apps_DeploymentRollback_To_v1beta1_DeploymentRollback(in *apps.DeploymentRollback, out *v1beta1.DeploymentRollback, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.UpdatedAnnotations = *(*map[string]string)(unsafe.Pointer(&in.UpdatedAnnotations))
 if err := Convert_apps_RollbackConfig_To_v1beta1_RollbackConfig(&in.RollbackTo, &out.RollbackTo, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_DeploymentRollback_To_v1beta1_DeploymentRollback(in *apps.DeploymentRollback, out *v1beta1.DeploymentRollback, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentRollback_To_v1beta1_DeploymentRollback(in, out, s)
}
func autoConvert_v1beta1_DeploymentSpec_To_apps_DeploymentSpec(in *v1beta1.DeploymentSpec, out *apps.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_DeploymentStrategy_To_apps_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.RollbackTo = (*apps.RollbackConfig)(unsafe.Pointer(in.RollbackTo))
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_apps_DeploymentSpec_To_v1beta1_DeploymentSpec(in *apps.DeploymentSpec, out *v1beta1.DeploymentSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DeploymentStrategy_To_v1beta1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 out.Paused = in.Paused
 out.RollbackTo = (*v1beta1.RollbackConfig)(unsafe.Pointer(in.RollbackTo))
 out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
 return nil
}
func autoConvert_v1beta1_DeploymentStatus_To_apps_DeploymentStatus(in *v1beta1.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
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
func Convert_v1beta1_DeploymentStatus_To_apps_DeploymentStatus(in *v1beta1.DeploymentStatus, out *apps.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DeploymentStatus_To_apps_DeploymentStatus(in, out, s)
}
func autoConvert_apps_DeploymentStatus_To_v1beta1_DeploymentStatus(in *apps.DeploymentStatus, out *v1beta1.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.Replicas = in.Replicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.UnavailableReplicas = in.UnavailableReplicas
 out.Conditions = *(*[]v1beta1.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 return nil
}
func Convert_apps_DeploymentStatus_To_v1beta1_DeploymentStatus(in *apps.DeploymentStatus, out *v1beta1.DeploymentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DeploymentStatus_To_v1beta1_DeploymentStatus(in, out, s)
}
func autoConvert_v1beta1_DeploymentStrategy_To_apps_DeploymentStrategy(in *v1beta1.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDeployment)
  if err := Convert_v1beta1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_DeploymentStrategy_To_v1beta1_DeploymentStrategy(in *apps.DeploymentStrategy, out *v1beta1.DeploymentStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.DeploymentStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta1.RollingUpdateDeployment)
  if err := Convert_apps_RollingUpdateDeployment_To_v1beta1_RollingUpdateDeployment(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_v1beta1_RollbackConfig_To_apps_RollbackConfig(in *v1beta1.RollbackConfig, out *apps.RollbackConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Revision = in.Revision
 return nil
}
func Convert_v1beta1_RollbackConfig_To_apps_RollbackConfig(in *v1beta1.RollbackConfig, out *apps.RollbackConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_RollbackConfig_To_apps_RollbackConfig(in, out, s)
}
func autoConvert_apps_RollbackConfig_To_v1beta1_RollbackConfig(in *apps.RollbackConfig, out *v1beta1.RollbackConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Revision = in.Revision
 return nil
}
func Convert_apps_RollbackConfig_To_v1beta1_RollbackConfig(in *apps.RollbackConfig, out *v1beta1.RollbackConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_RollbackConfig_To_v1beta1_RollbackConfig(in, out, s)
}
func autoConvert_v1beta1_RollingUpdateDeployment_To_apps_RollingUpdateDeployment(in *v1beta1.RollingUpdateDeployment, out *apps.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDeployment_To_v1beta1_RollingUpdateDeployment(in *apps.RollingUpdateDeployment, out *v1beta1.RollingUpdateDeployment, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1beta1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta1.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta1.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta1_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta1.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Partition, &out.Partition, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta1_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta1.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta1_RollingUpdateStatefulSetStrategy(in, out, s)
}
func autoConvert_v1beta1_Scale_To_autoscaling_Scale(in *v1beta1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_Scale_To_autoscaling_Scale(in *v1beta1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Scale_To_autoscaling_Scale(in, out, s)
}
func autoConvert_autoscaling_Scale_To_v1beta1_Scale(in *autoscaling.Scale, out *v1beta1.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_autoscaling_ScaleSpec_To_v1beta1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_autoscaling_ScaleStatus_To_v1beta1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_autoscaling_Scale_To_v1beta1_Scale(in *autoscaling.Scale, out *v1beta1.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_Scale_To_v1beta1_Scale(in, out, s)
}
func autoConvert_v1beta1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1beta1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_v1beta1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1beta1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}
func autoConvert_autoscaling_ScaleSpec_To_v1beta1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1beta1.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_autoscaling_ScaleSpec_To_v1beta1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1beta1.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_ScaleSpec_To_v1beta1_ScaleSpec(in, out, s)
}
func autoConvert_v1beta1_ScaleStatus_To_autoscaling_ScaleStatus(in *v1beta1.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func autoConvert_autoscaling_ScaleStatus_To_v1beta1_ScaleStatus(in *autoscaling.ScaleStatus, out *v1beta1.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func autoConvert_v1beta1_StatefulSet_To_apps_StatefulSet(in *v1beta1.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_StatefulSet_To_apps_StatefulSet(in *v1beta1.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StatefulSet_To_apps_StatefulSet(in, out, s)
}
func autoConvert_apps_StatefulSet_To_v1beta1_StatefulSet(in *apps.StatefulSet, out *v1beta1.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_StatefulSet_To_v1beta1_StatefulSet(in *apps.StatefulSet, out *v1beta1.StatefulSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSet_To_v1beta1_StatefulSet(in, out, s)
}
func autoConvert_v1beta1_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1beta1.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta1_StatefulSetCondition_To_apps_StatefulSetCondition(in *v1beta1.StatefulSetCondition, out *apps.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StatefulSetCondition_To_apps_StatefulSetCondition(in, out, s)
}
func autoConvert_apps_StatefulSetCondition_To_v1beta1_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1beta1.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.StatefulSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_StatefulSetCondition_To_v1beta1_StatefulSetCondition(in *apps.StatefulSetCondition, out *v1beta1.StatefulSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetCondition_To_v1beta1_StatefulSetCondition(in, out, s)
}
func autoConvert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in *v1beta1.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_StatefulSet_To_apps_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in *v1beta1.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in, out, s)
}
func autoConvert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in *apps.StatefulSetList, out *v1beta1.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.StatefulSet, len(*in))
  for i := range *in {
   if err := Convert_apps_StatefulSet_To_v1beta1_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in *apps.StatefulSetList, out *v1beta1.StatefulSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in, out, s)
}
func autoConvert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(in *v1beta1.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
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
 if err := Convert_v1beta1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(in *apps.StatefulSetSpec, out *v1beta1.StatefulSetSpec, s conversion.Scope) error {
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
 out.PodManagementPolicy = v1beta1.PodManagementPolicyType(in.PodManagementPolicy)
 if err := Convert_apps_StatefulSetUpdateStrategy_To_v1beta1_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func autoConvert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1beta1.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
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
func Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1beta1.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in, out, s)
}
func autoConvert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1beta1.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
 out.Replicas = in.Replicas
 out.ReadyReplicas = in.ReadyReplicas
 out.CurrentReplicas = in.CurrentReplicas
 out.UpdatedReplicas = in.UpdatedReplicas
 out.CurrentRevision = in.CurrentRevision
 out.UpdateRevision = in.UpdateRevision
 out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
 out.Conditions = *(*[]v1beta1.StatefulSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1beta1.StatefulSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in, out, s)
}
func autoConvert_v1beta1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *v1beta1.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateStatefulSetStrategy)
  if err := Convert_v1beta1_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func autoConvert_apps_StatefulSetUpdateStrategy_To_v1beta1_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *v1beta1.StatefulSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.StatefulSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta1.RollingUpdateStatefulSetStrategy)
  if err := Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta1_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
