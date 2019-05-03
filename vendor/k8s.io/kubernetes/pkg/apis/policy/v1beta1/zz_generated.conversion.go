package v1beta1

import (
 unsafe "unsafe"
 corev1 "k8s.io/api/core/v1"
 v1beta1 "k8s.io/api/policy/v1beta1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 intstr "k8s.io/apimachinery/pkg/util/intstr"
 core "k8s.io/kubernetes/pkg/apis/core"
 policy "k8s.io/kubernetes/pkg/apis/policy"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.AllowedFlexVolume)(nil), (*policy.AllowedFlexVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(a.(*v1beta1.AllowedFlexVolume), b.(*policy.AllowedFlexVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.AllowedFlexVolume)(nil), (*v1beta1.AllowedFlexVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(a.(*policy.AllowedFlexVolume), b.(*v1beta1.AllowedFlexVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.AllowedHostPath)(nil), (*policy.AllowedHostPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(a.(*v1beta1.AllowedHostPath), b.(*policy.AllowedHostPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.AllowedHostPath)(nil), (*v1beta1.AllowedHostPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(a.(*policy.AllowedHostPath), b.(*v1beta1.AllowedHostPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.Eviction)(nil), (*policy.Eviction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Eviction_To_policy_Eviction(a.(*v1beta1.Eviction), b.(*policy.Eviction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.Eviction)(nil), (*v1beta1.Eviction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_Eviction_To_v1beta1_Eviction(a.(*policy.Eviction), b.(*v1beta1.Eviction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.FSGroupStrategyOptions)(nil), (*policy.FSGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(a.(*v1beta1.FSGroupStrategyOptions), b.(*policy.FSGroupStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.FSGroupStrategyOptions)(nil), (*v1beta1.FSGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(a.(*policy.FSGroupStrategyOptions), b.(*v1beta1.FSGroupStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.HostPortRange)(nil), (*policy.HostPortRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_HostPortRange_To_policy_HostPortRange(a.(*v1beta1.HostPortRange), b.(*policy.HostPortRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.HostPortRange)(nil), (*v1beta1.HostPortRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_HostPortRange_To_v1beta1_HostPortRange(a.(*policy.HostPortRange), b.(*v1beta1.HostPortRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IDRange)(nil), (*policy.IDRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IDRange_To_policy_IDRange(a.(*v1beta1.IDRange), b.(*policy.IDRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.IDRange)(nil), (*v1beta1.IDRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_IDRange_To_v1beta1_IDRange(a.(*policy.IDRange), b.(*v1beta1.IDRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodDisruptionBudget)(nil), (*policy.PodDisruptionBudget)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(a.(*v1beta1.PodDisruptionBudget), b.(*policy.PodDisruptionBudget), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudget)(nil), (*v1beta1.PodDisruptionBudget)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(a.(*policy.PodDisruptionBudget), b.(*v1beta1.PodDisruptionBudget), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodDisruptionBudgetList)(nil), (*policy.PodDisruptionBudgetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(a.(*v1beta1.PodDisruptionBudgetList), b.(*policy.PodDisruptionBudgetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetList)(nil), (*v1beta1.PodDisruptionBudgetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(a.(*policy.PodDisruptionBudgetList), b.(*v1beta1.PodDisruptionBudgetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodDisruptionBudgetSpec)(nil), (*policy.PodDisruptionBudgetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(a.(*v1beta1.PodDisruptionBudgetSpec), b.(*policy.PodDisruptionBudgetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetSpec)(nil), (*v1beta1.PodDisruptionBudgetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(a.(*policy.PodDisruptionBudgetSpec), b.(*v1beta1.PodDisruptionBudgetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodDisruptionBudgetStatus)(nil), (*policy.PodDisruptionBudgetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(a.(*v1beta1.PodDisruptionBudgetStatus), b.(*policy.PodDisruptionBudgetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetStatus)(nil), (*v1beta1.PodDisruptionBudgetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(a.(*policy.PodDisruptionBudgetStatus), b.(*v1beta1.PodDisruptionBudgetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodSecurityPolicy)(nil), (*policy.PodSecurityPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(a.(*v1beta1.PodSecurityPolicy), b.(*policy.PodSecurityPolicy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodSecurityPolicy)(nil), (*v1beta1.PodSecurityPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(a.(*policy.PodSecurityPolicy), b.(*v1beta1.PodSecurityPolicy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodSecurityPolicyList)(nil), (*policy.PodSecurityPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(a.(*v1beta1.PodSecurityPolicyList), b.(*policy.PodSecurityPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodSecurityPolicyList)(nil), (*v1beta1.PodSecurityPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(a.(*policy.PodSecurityPolicyList), b.(*v1beta1.PodSecurityPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PodSecurityPolicySpec)(nil), (*policy.PodSecurityPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(a.(*v1beta1.PodSecurityPolicySpec), b.(*policy.PodSecurityPolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.PodSecurityPolicySpec)(nil), (*v1beta1.PodSecurityPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(a.(*policy.PodSecurityPolicySpec), b.(*v1beta1.PodSecurityPolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RunAsGroupStrategyOptions)(nil), (*policy.RunAsGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RunAsGroupStrategyOptions_To_policy_RunAsGroupStrategyOptions(a.(*v1beta1.RunAsGroupStrategyOptions), b.(*policy.RunAsGroupStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.RunAsGroupStrategyOptions)(nil), (*v1beta1.RunAsGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_RunAsGroupStrategyOptions_To_v1beta1_RunAsGroupStrategyOptions(a.(*policy.RunAsGroupStrategyOptions), b.(*v1beta1.RunAsGroupStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RunAsUserStrategyOptions)(nil), (*policy.RunAsUserStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(a.(*v1beta1.RunAsUserStrategyOptions), b.(*policy.RunAsUserStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.RunAsUserStrategyOptions)(nil), (*v1beta1.RunAsUserStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(a.(*policy.RunAsUserStrategyOptions), b.(*v1beta1.RunAsUserStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.SELinuxStrategyOptions)(nil), (*policy.SELinuxStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(a.(*v1beta1.SELinuxStrategyOptions), b.(*policy.SELinuxStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.SELinuxStrategyOptions)(nil), (*v1beta1.SELinuxStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(a.(*policy.SELinuxStrategyOptions), b.(*v1beta1.SELinuxStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.SupplementalGroupsStrategyOptions)(nil), (*policy.SupplementalGroupsStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(a.(*v1beta1.SupplementalGroupsStrategyOptions), b.(*policy.SupplementalGroupsStrategyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*policy.SupplementalGroupsStrategyOptions)(nil), (*v1beta1.SupplementalGroupsStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(a.(*policy.SupplementalGroupsStrategyOptions), b.(*v1beta1.SupplementalGroupsStrategyOptions), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in *v1beta1.AllowedFlexVolume, out *policy.AllowedFlexVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 return nil
}
func Convert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in *v1beta1.AllowedFlexVolume, out *policy.AllowedFlexVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in, out, s)
}
func autoConvert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in *policy.AllowedFlexVolume, out *v1beta1.AllowedFlexVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 return nil
}
func Convert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in *policy.AllowedFlexVolume, out *v1beta1.AllowedFlexVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in, out, s)
}
func autoConvert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in *v1beta1.AllowedHostPath, out *policy.AllowedHostPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PathPrefix = in.PathPrefix
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in *v1beta1.AllowedHostPath, out *policy.AllowedHostPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in, out, s)
}
func autoConvert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in *policy.AllowedHostPath, out *v1beta1.AllowedHostPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PathPrefix = in.PathPrefix
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in *policy.AllowedHostPath, out *v1beta1.AllowedHostPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in, out, s)
}
func autoConvert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
 return nil
}
func Convert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Eviction_To_policy_Eviction(in, out, s)
}
func autoConvert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
 return nil
}
func Convert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_Eviction_To_v1beta1_Eviction(in, out, s)
}
func autoConvert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in *v1beta1.FSGroupStrategyOptions, out *policy.FSGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = policy.FSGroupStrategyType(in.Rule)
 out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in *v1beta1.FSGroupStrategyOptions, out *policy.FSGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in, out, s)
}
func autoConvert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in *policy.FSGroupStrategyOptions, out *v1beta1.FSGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = v1beta1.FSGroupStrategyType(in.Rule)
 out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in *policy.FSGroupStrategyOptions, out *v1beta1.FSGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in, out, s)
}
func autoConvert_v1beta1_HostPortRange_To_policy_HostPortRange(in *v1beta1.HostPortRange, out *policy.HostPortRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Min = in.Min
 out.Max = in.Max
 return nil
}
func Convert_v1beta1_HostPortRange_To_policy_HostPortRange(in *v1beta1.HostPortRange, out *policy.HostPortRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_HostPortRange_To_policy_HostPortRange(in, out, s)
}
func autoConvert_policy_HostPortRange_To_v1beta1_HostPortRange(in *policy.HostPortRange, out *v1beta1.HostPortRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Min = in.Min
 out.Max = in.Max
 return nil
}
func Convert_policy_HostPortRange_To_v1beta1_HostPortRange(in *policy.HostPortRange, out *v1beta1.HostPortRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_HostPortRange_To_v1beta1_HostPortRange(in, out, s)
}
func autoConvert_v1beta1_IDRange_To_policy_IDRange(in *v1beta1.IDRange, out *policy.IDRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Min = in.Min
 out.Max = in.Max
 return nil
}
func Convert_v1beta1_IDRange_To_policy_IDRange(in *v1beta1.IDRange, out *policy.IDRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IDRange_To_policy_IDRange(in, out, s)
}
func autoConvert_policy_IDRange_To_v1beta1_IDRange(in *policy.IDRange, out *v1beta1.IDRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Min = in.Min
 out.Max = in.Max
 return nil
}
func Convert_policy_IDRange_To_v1beta1_IDRange(in *policy.IDRange, out *v1beta1.IDRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_IDRange_To_v1beta1_IDRange(in, out, s)
}
func autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in, out, s)
}
func autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in, out, s)
}
func autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]policy.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in, out, s)
}
func autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in, out, s)
}
func autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
 out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
 out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
 return nil
}
func Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in, out, s)
}
func autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
 out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
 out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
 return nil
}
func Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in, out, s)
}
func autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
 out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
 out.CurrentHealthy = in.CurrentHealthy
 out.DesiredHealthy = in.DesiredHealthy
 out.ExpectedPods = in.ExpectedPods
 return nil
}
func Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in, out, s)
}
func autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObservedGeneration = in.ObservedGeneration
 out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
 out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
 out.CurrentHealthy = in.CurrentHealthy
 out.DesiredHealthy = in.DesiredHealthy
 out.ExpectedPods = in.ExpectedPods
 return nil
}
func Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in, out, s)
}
func autoConvert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in *v1beta1.PodSecurityPolicy, out *policy.PodSecurityPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in *v1beta1.PodSecurityPolicy, out *policy.PodSecurityPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in, out, s)
}
func autoConvert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in *policy.PodSecurityPolicy, out *v1beta1.PodSecurityPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in *policy.PodSecurityPolicy, out *v1beta1.PodSecurityPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in, out, s)
}
func autoConvert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in *v1beta1.PodSecurityPolicyList, out *policy.PodSecurityPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]policy.PodSecurityPolicy, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in *v1beta1.PodSecurityPolicyList, out *policy.PodSecurityPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in, out, s)
}
func autoConvert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in *policy.PodSecurityPolicyList, out *v1beta1.PodSecurityPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.PodSecurityPolicy, len(*in))
  for i := range *in {
   if err := Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in *policy.PodSecurityPolicyList, out *v1beta1.PodSecurityPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in, out, s)
}
func autoConvert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in *v1beta1.PodSecurityPolicySpec, out *policy.PodSecurityPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Privileged = in.Privileged
 out.DefaultAddCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
 out.RequiredDropCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
 out.AllowedCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
 out.Volumes = *(*[]policy.FSType)(unsafe.Pointer(&in.Volumes))
 out.HostNetwork = in.HostNetwork
 out.HostPorts = *(*[]policy.HostPortRange)(unsafe.Pointer(&in.HostPorts))
 out.HostPID = in.HostPID
 out.HostIPC = in.HostIPC
 if err := Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
  return err
 }
 out.RunAsGroup = (*policy.RunAsGroupStrategyOptions)(unsafe.Pointer(in.RunAsGroup))
 if err := Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
  return err
 }
 out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
 out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
 if err := v1.Convert_Pointer_bool_To_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
  return err
 }
 out.AllowedHostPaths = *(*[]policy.AllowedHostPath)(unsafe.Pointer(&in.AllowedHostPaths))
 out.AllowedFlexVolumes = *(*[]policy.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
 out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
 out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
 out.AllowedProcMountTypes = *(*[]core.ProcMountType)(unsafe.Pointer(&in.AllowedProcMountTypes))
 return nil
}
func Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in *v1beta1.PodSecurityPolicySpec, out *policy.PodSecurityPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in, out, s)
}
func autoConvert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in *policy.PodSecurityPolicySpec, out *v1beta1.PodSecurityPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Privileged = in.Privileged
 out.DefaultAddCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
 out.RequiredDropCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
 out.AllowedCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
 out.Volumes = *(*[]v1beta1.FSType)(unsafe.Pointer(&in.Volumes))
 out.HostNetwork = in.HostNetwork
 out.HostPorts = *(*[]v1beta1.HostPortRange)(unsafe.Pointer(&in.HostPorts))
 out.HostPID = in.HostPID
 out.HostIPC = in.HostIPC
 if err := Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, s); err != nil {
  return err
 }
 if err := Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
  return err
 }
 out.RunAsGroup = (*v1beta1.RunAsGroupStrategyOptions)(unsafe.Pointer(in.RunAsGroup))
 if err := Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
  return err
 }
 if err := Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
  return err
 }
 out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
 out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
 if err := v1.Convert_bool_To_Pointer_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
  return err
 }
 out.AllowedHostPaths = *(*[]v1beta1.AllowedHostPath)(unsafe.Pointer(&in.AllowedHostPaths))
 out.AllowedFlexVolumes = *(*[]v1beta1.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
 out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
 out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
 out.AllowedProcMountTypes = *(*[]corev1.ProcMountType)(unsafe.Pointer(&in.AllowedProcMountTypes))
 return nil
}
func Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in *policy.PodSecurityPolicySpec, out *v1beta1.PodSecurityPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in, out, s)
}
func autoConvert_v1beta1_RunAsGroupStrategyOptions_To_policy_RunAsGroupStrategyOptions(in *v1beta1.RunAsGroupStrategyOptions, out *policy.RunAsGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = policy.RunAsGroupStrategy(in.Rule)
 out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_v1beta1_RunAsGroupStrategyOptions_To_policy_RunAsGroupStrategyOptions(in *v1beta1.RunAsGroupStrategyOptions, out *policy.RunAsGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_RunAsGroupStrategyOptions_To_policy_RunAsGroupStrategyOptions(in, out, s)
}
func autoConvert_policy_RunAsGroupStrategyOptions_To_v1beta1_RunAsGroupStrategyOptions(in *policy.RunAsGroupStrategyOptions, out *v1beta1.RunAsGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = v1beta1.RunAsGroupStrategy(in.Rule)
 out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_policy_RunAsGroupStrategyOptions_To_v1beta1_RunAsGroupStrategyOptions(in *policy.RunAsGroupStrategyOptions, out *v1beta1.RunAsGroupStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_RunAsGroupStrategyOptions_To_v1beta1_RunAsGroupStrategyOptions(in, out, s)
}
func autoConvert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in *v1beta1.RunAsUserStrategyOptions, out *policy.RunAsUserStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = policy.RunAsUserStrategy(in.Rule)
 out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in *v1beta1.RunAsUserStrategyOptions, out *policy.RunAsUserStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in, out, s)
}
func autoConvert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in *policy.RunAsUserStrategyOptions, out *v1beta1.RunAsUserStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = v1beta1.RunAsUserStrategy(in.Rule)
 out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in *policy.RunAsUserStrategyOptions, out *v1beta1.RunAsUserStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in, out, s)
}
func autoConvert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in *v1beta1.SELinuxStrategyOptions, out *policy.SELinuxStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = policy.SELinuxStrategy(in.Rule)
 out.SELinuxOptions = (*core.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 return nil
}
func Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in *v1beta1.SELinuxStrategyOptions, out *policy.SELinuxStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in, out, s)
}
func autoConvert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in *policy.SELinuxStrategyOptions, out *v1beta1.SELinuxStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = v1beta1.SELinuxStrategy(in.Rule)
 out.SELinuxOptions = (*corev1.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 return nil
}
func Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in *policy.SELinuxStrategyOptions, out *v1beta1.SELinuxStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in, out, s)
}
func autoConvert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in *v1beta1.SupplementalGroupsStrategyOptions, out *policy.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = policy.SupplementalGroupsStrategyType(in.Rule)
 out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in *v1beta1.SupplementalGroupsStrategyOptions, out *policy.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in, out, s)
}
func autoConvert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in *policy.SupplementalGroupsStrategyOptions, out *v1beta1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Rule = v1beta1.SupplementalGroupsStrategyType(in.Rule)
 out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
 return nil
}
func Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in *policy.SupplementalGroupsStrategyOptions, out *v1beta1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in, out, s)
}
