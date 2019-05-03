package v1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/autoscaling/v1"
 corev1 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1.CrossVersionObjectReference)(nil), (*autoscaling.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(a.(*v1.CrossVersionObjectReference), b.(*autoscaling.CrossVersionObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.CrossVersionObjectReference)(nil), (*v1.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(a.(*autoscaling.CrossVersionObjectReference), b.(*v1.CrossVersionObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ExternalMetricSource)(nil), (*autoscaling.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(a.(*v1.ExternalMetricSource), b.(*autoscaling.ExternalMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ExternalMetricSource)(nil), (*v1.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ExternalMetricSource_To_v1_ExternalMetricSource(a.(*autoscaling.ExternalMetricSource), b.(*v1.ExternalMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ExternalMetricStatus)(nil), (*autoscaling.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(a.(*v1.ExternalMetricStatus), b.(*autoscaling.ExternalMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ExternalMetricStatus)(nil), (*v1.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ExternalMetricStatus_To_v1_ExternalMetricStatus(a.(*autoscaling.ExternalMetricStatus), b.(*v1.ExternalMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HorizontalPodAutoscaler)(nil), (*autoscaling.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(a.(*v1.HorizontalPodAutoscaler), b.(*autoscaling.HorizontalPodAutoscaler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscaler)(nil), (*v1.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(a.(*autoscaling.HorizontalPodAutoscaler), b.(*v1.HorizontalPodAutoscaler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HorizontalPodAutoscalerCondition)(nil), (*autoscaling.HorizontalPodAutoscalerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(a.(*v1.HorizontalPodAutoscalerCondition), b.(*autoscaling.HorizontalPodAutoscalerCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerCondition)(nil), (*v1.HorizontalPodAutoscalerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(a.(*autoscaling.HorizontalPodAutoscalerCondition), b.(*v1.HorizontalPodAutoscalerCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HorizontalPodAutoscalerList)(nil), (*autoscaling.HorizontalPodAutoscalerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(a.(*v1.HorizontalPodAutoscalerList), b.(*autoscaling.HorizontalPodAutoscalerList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerList)(nil), (*v1.HorizontalPodAutoscalerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(a.(*autoscaling.HorizontalPodAutoscalerList), b.(*v1.HorizontalPodAutoscalerList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HorizontalPodAutoscalerSpec)(nil), (*autoscaling.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(a.(*v1.HorizontalPodAutoscalerSpec), b.(*autoscaling.HorizontalPodAutoscalerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerSpec)(nil), (*v1.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(a.(*autoscaling.HorizontalPodAutoscalerSpec), b.(*v1.HorizontalPodAutoscalerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HorizontalPodAutoscalerStatus)(nil), (*autoscaling.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(a.(*v1.HorizontalPodAutoscalerStatus), b.(*autoscaling.HorizontalPodAutoscalerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerStatus)(nil), (*v1.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(a.(*autoscaling.HorizontalPodAutoscalerStatus), b.(*v1.HorizontalPodAutoscalerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.MetricSpec)(nil), (*autoscaling.MetricSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_MetricSpec_To_autoscaling_MetricSpec(a.(*v1.MetricSpec), b.(*autoscaling.MetricSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.MetricSpec)(nil), (*v1.MetricSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_MetricSpec_To_v1_MetricSpec(a.(*autoscaling.MetricSpec), b.(*v1.MetricSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.MetricStatus)(nil), (*autoscaling.MetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_MetricStatus_To_autoscaling_MetricStatus(a.(*v1.MetricStatus), b.(*autoscaling.MetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.MetricStatus)(nil), (*v1.MetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_MetricStatus_To_v1_MetricStatus(a.(*autoscaling.MetricStatus), b.(*v1.MetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ObjectMetricSource)(nil), (*autoscaling.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(a.(*v1.ObjectMetricSource), b.(*autoscaling.ObjectMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ObjectMetricSource)(nil), (*v1.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(a.(*autoscaling.ObjectMetricSource), b.(*v1.ObjectMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ObjectMetricStatus)(nil), (*autoscaling.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(a.(*v1.ObjectMetricStatus), b.(*autoscaling.ObjectMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ObjectMetricStatus)(nil), (*v1.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(a.(*autoscaling.ObjectMetricStatus), b.(*v1.ObjectMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodsMetricSource)(nil), (*autoscaling.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(a.(*v1.PodsMetricSource), b.(*autoscaling.PodsMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.PodsMetricSource)(nil), (*v1.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(a.(*autoscaling.PodsMetricSource), b.(*v1.PodsMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodsMetricStatus)(nil), (*autoscaling.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(a.(*v1.PodsMetricStatus), b.(*autoscaling.PodsMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.PodsMetricStatus)(nil), (*v1.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(a.(*autoscaling.PodsMetricStatus), b.(*v1.PodsMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceMetricSource)(nil), (*autoscaling.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(a.(*v1.ResourceMetricSource), b.(*autoscaling.ResourceMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ResourceMetricSource)(nil), (*v1.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(a.(*autoscaling.ResourceMetricSource), b.(*v1.ResourceMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceMetricStatus)(nil), (*autoscaling.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(a.(*v1.ResourceMetricStatus), b.(*autoscaling.ResourceMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ResourceMetricStatus)(nil), (*v1.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(a.(*autoscaling.ResourceMetricStatus), b.(*v1.ResourceMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Scale)(nil), (*autoscaling.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Scale_To_autoscaling_Scale(a.(*v1.Scale), b.(*autoscaling.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.Scale)(nil), (*v1.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_Scale_To_v1_Scale(a.(*autoscaling.Scale), b.(*v1.Scale), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScaleSpec)(nil), (*autoscaling.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(a.(*v1.ScaleSpec), b.(*autoscaling.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleSpec)(nil), (*v1.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(a.(*autoscaling.ScaleSpec), b.(*v1.ScaleSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScaleStatus)(nil), (*autoscaling.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(a.(*v1.ScaleStatus), b.(*autoscaling.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ExternalMetricSource)(nil), (*v1.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ExternalMetricSource_To_v1_ExternalMetricSource(a.(*autoscaling.ExternalMetricSource), b.(*v1.ExternalMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ExternalMetricStatus)(nil), (*v1.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ExternalMetricStatus_To_v1_ExternalMetricStatus(a.(*autoscaling.ExternalMetricStatus), b.(*v1.ExternalMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.HorizontalPodAutoscalerSpec)(nil), (*v1.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(a.(*autoscaling.HorizontalPodAutoscalerSpec), b.(*v1.HorizontalPodAutoscalerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.HorizontalPodAutoscalerStatus)(nil), (*v1.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(a.(*autoscaling.HorizontalPodAutoscalerStatus), b.(*v1.HorizontalPodAutoscalerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.HorizontalPodAutoscaler)(nil), (*v1.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(a.(*autoscaling.HorizontalPodAutoscaler), b.(*v1.HorizontalPodAutoscaler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.MetricTarget)(nil), (*v1.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_MetricTarget_To_v1_CrossVersionObjectReference(a.(*autoscaling.MetricTarget), b.(*v1.CrossVersionObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ObjectMetricSource)(nil), (*v1.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(a.(*autoscaling.ObjectMetricSource), b.(*v1.ObjectMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ObjectMetricStatus)(nil), (*v1.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(a.(*autoscaling.ObjectMetricStatus), b.(*v1.ObjectMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.PodsMetricSource)(nil), (*v1.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(a.(*autoscaling.PodsMetricSource), b.(*v1.PodsMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.PodsMetricStatus)(nil), (*v1.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(a.(*autoscaling.PodsMetricStatus), b.(*v1.PodsMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ResourceMetricSource)(nil), (*v1.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(a.(*autoscaling.ResourceMetricSource), b.(*v1.ResourceMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ResourceMetricStatus)(nil), (*v1.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(a.(*autoscaling.ResourceMetricStatus), b.(*v1.ResourceMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.CrossVersionObjectReference)(nil), (*autoscaling.MetricTarget)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CrossVersionObjectReference_To_autoscaling_MetricTarget(a.(*v1.CrossVersionObjectReference), b.(*autoscaling.MetricTarget), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ExternalMetricSource)(nil), (*autoscaling.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(a.(*v1.ExternalMetricSource), b.(*autoscaling.ExternalMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ExternalMetricStatus)(nil), (*autoscaling.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(a.(*v1.ExternalMetricStatus), b.(*autoscaling.ExternalMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.HorizontalPodAutoscalerSpec)(nil), (*autoscaling.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(a.(*v1.HorizontalPodAutoscalerSpec), b.(*autoscaling.HorizontalPodAutoscalerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.HorizontalPodAutoscalerStatus)(nil), (*autoscaling.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(a.(*v1.HorizontalPodAutoscalerStatus), b.(*autoscaling.HorizontalPodAutoscalerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.HorizontalPodAutoscaler)(nil), (*autoscaling.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(a.(*v1.HorizontalPodAutoscaler), b.(*autoscaling.HorizontalPodAutoscaler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ObjectMetricSource)(nil), (*autoscaling.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(a.(*v1.ObjectMetricSource), b.(*autoscaling.ObjectMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ObjectMetricStatus)(nil), (*autoscaling.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(a.(*v1.ObjectMetricStatus), b.(*autoscaling.ObjectMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.PodsMetricSource)(nil), (*autoscaling.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(a.(*v1.PodsMetricSource), b.(*autoscaling.PodsMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.PodsMetricStatus)(nil), (*autoscaling.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(a.(*v1.PodsMetricStatus), b.(*autoscaling.PodsMetricStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ResourceMetricSource)(nil), (*autoscaling.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(a.(*v1.ResourceMetricSource), b.(*autoscaling.ResourceMetricSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ResourceMetricStatus)(nil), (*autoscaling.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(a.(*v1.ResourceMetricStatus), b.(*autoscaling.ResourceMetricStatus), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Name = in.Name
 out.APIVersion = in.APIVersion
 return nil
}
func Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}
func autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v1.CrossVersionObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Name = in.Name
 out.APIVersion = in.APIVersion
 return nil
}
func Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v1.CrossVersionObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in, out, s)
}
func autoConvert_v1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(in *v1.ExternalMetricSource, out *autoscaling.ExternalMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_autoscaling_ExternalMetricSource_To_v1_ExternalMetricSource(in *autoscaling.ExternalMetricSource, out *v1.ExternalMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(in *v1.ExternalMetricStatus, out *autoscaling.ExternalMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_autoscaling_ExternalMetricStatus_To_v1_ExternalMetricStatus(in *autoscaling.ExternalMetricStatus, out *v1.ExternalMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *v1.HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *v1.HorizontalPodAutoscaler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = autoscaling.HorizontalPodAutoscalerConditionType(in.Type)
 out.Status = autoscaling.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in, out, s)
}
func autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.HorizontalPodAutoscalerConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in, out, s)
}
func autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]autoscaling.HorizontalPodAutoscaler, len(*in))
  for i := range *in {
   if err := Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}
func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v1.HorizontalPodAutoscalerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.HorizontalPodAutoscaler, len(*in))
  for i := range *in {
   if err := Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v1.HorizontalPodAutoscalerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in, out, s)
}
func autoConvert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
  return err
 }
 out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
 out.MaxReplicas = in.MaxReplicas
 return nil
}
func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *v1.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
  return err
 }
 out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
 out.MaxReplicas = in.MaxReplicas
 return nil
}
func autoConvert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
 out.LastScaleTime = (*metav1.Time)(unsafe.Pointer(in.LastScaleTime))
 out.CurrentReplicas = in.CurrentReplicas
 out.DesiredReplicas = in.DesiredReplicas
 return nil
}
func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
 out.LastScaleTime = (*metav1.Time)(unsafe.Pointer(in.LastScaleTime))
 out.CurrentReplicas = in.CurrentReplicas
 out.DesiredReplicas = in.DesiredReplicas
 return nil
}
func autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in *v1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = autoscaling.MetricSourceType(in.Type)
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(autoscaling.ObjectMetricSource)
  if err := Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Object = nil
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(autoscaling.PodsMetricSource)
  if err := Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Pods = nil
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(autoscaling.ResourceMetricSource)
  if err := Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Resource = nil
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(autoscaling.ExternalMetricSource)
  if err := Convert_v1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.External = nil
 }
 return nil
}
func Convert_v1_MetricSpec_To_autoscaling_MetricSpec(in *v1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in, out, s)
}
func autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *v1.MetricSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.MetricSourceType(in.Type)
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(v1.ObjectMetricSource)
  if err := Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Object = nil
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(v1.PodsMetricSource)
  if err := Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Pods = nil
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(v1.ResourceMetricSource)
  if err := Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Resource = nil
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(v1.ExternalMetricSource)
  if err := Convert_autoscaling_ExternalMetricSource_To_v1_ExternalMetricSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.External = nil
 }
 return nil
}
func Convert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *v1.MetricSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in, out, s)
}
func autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in *v1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = autoscaling.MetricSourceType(in.Type)
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(autoscaling.ObjectMetricStatus)
  if err := Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Object = nil
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(autoscaling.PodsMetricStatus)
  if err := Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Pods = nil
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(autoscaling.ResourceMetricStatus)
  if err := Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Resource = nil
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(autoscaling.ExternalMetricStatus)
  if err := Convert_v1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.External = nil
 }
 return nil
}
func Convert_v1_MetricStatus_To_autoscaling_MetricStatus(in *v1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in, out, s)
}
func autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *v1.MetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.MetricSourceType(in.Type)
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(v1.ObjectMetricStatus)
  if err := Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Object = nil
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(v1.PodsMetricStatus)
  if err := Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Pods = nil
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(v1.ResourceMetricStatus)
  if err := Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Resource = nil
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(v1.ExternalMetricStatus)
  if err := Convert_autoscaling_ExternalMetricStatus_To_v1_ExternalMetricStatus(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.External = nil
 }
 return nil
}
func Convert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *v1.MetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in, out, s)
}
func autoConvert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_MetricTarget(&in.Target, &out.Target, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v1.ObjectMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_autoscaling_MetricTarget_To_v1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v1.ObjectMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v1.PodsMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v1.PodsMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = core.ResourceName(in.Name)
 return nil
}
func autoConvert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v1.ResourceMetricSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = corev1.ResourceName(in.Name)
 return nil
}
func autoConvert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = core.ResourceName(in.Name)
 return nil
}
func autoConvert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v1.ResourceMetricStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = corev1.ResourceName(in.Name)
 return nil
}
func autoConvert_v1_Scale_To_autoscaling_Scale(in *v1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Scale_To_autoscaling_Scale(in *v1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Scale_To_autoscaling_Scale(in, out, s)
}
func autoConvert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *v1.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *v1.Scale, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_Scale_To_v1_Scale(in, out, s)
}
func autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}
func autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 return nil
}
func Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1.ScaleSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in, out, s)
}
func autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *v1.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.Selector = in.Selector
 return nil
}
func Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *v1.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in, out, s)
}
func autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *v1.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.Selector = in.Selector
 return nil
}
func Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *v1.ScaleStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in, out, s)
}
