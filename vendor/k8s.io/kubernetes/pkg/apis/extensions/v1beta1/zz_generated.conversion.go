package v1beta1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/core/v1"
 v1beta1 "k8s.io/api/extensions/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 apps "k8s.io/kubernetes/pkg/apis/apps"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 core "k8s.io/kubernetes/pkg/apis/core"
 corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
 extensions "k8s.io/kubernetes/pkg/apis/extensions"
 networking "k8s.io/kubernetes/pkg/apis/networking"
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSet)(nil), (*apps.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSet_To_apps_DaemonSet(a.(*v1beta1.DaemonSet), b.(*apps.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSet)(nil), (*v1beta1.DaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSet_To_v1beta1_DaemonSet(a.(*apps.DaemonSet), b.(*v1beta1.DaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSetCondition)(nil), (*apps.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSetCondition_To_apps_DaemonSetCondition(a.(*v1beta1.DaemonSetCondition), b.(*apps.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetCondition)(nil), (*v1beta1.DaemonSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetCondition_To_v1beta1_DaemonSetCondition(a.(*apps.DaemonSetCondition), b.(*v1beta1.DaemonSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSetList)(nil), (*apps.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSetList_To_apps_DaemonSetList(a.(*v1beta1.DaemonSetList), b.(*apps.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetList)(nil), (*v1beta1.DaemonSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetList_To_v1beta1_DaemonSetList(a.(*apps.DaemonSetList), b.(*v1beta1.DaemonSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSetSpec)(nil), (*apps.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSetSpec_To_apps_DaemonSetSpec(a.(*v1beta1.DaemonSetSpec), b.(*apps.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetSpec)(nil), (*v1beta1.DaemonSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetSpec_To_v1beta1_DaemonSetSpec(a.(*apps.DaemonSetSpec), b.(*v1beta1.DaemonSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSetStatus)(nil), (*apps.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSetStatus_To_apps_DaemonSetStatus(a.(*v1beta1.DaemonSetStatus), b.(*apps.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetStatus)(nil), (*v1beta1.DaemonSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetStatus_To_v1beta1_DaemonSetStatus(a.(*apps.DaemonSetStatus), b.(*v1beta1.DaemonSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.DaemonSetUpdateStrategy)(nil), (*apps.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(a.(*v1beta1.DaemonSetUpdateStrategy), b.(*apps.DaemonSetUpdateStrategy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.DaemonSetUpdateStrategy)(nil), (*v1beta1.DaemonSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_DaemonSetUpdateStrategy_To_v1beta1_DaemonSetUpdateStrategy(a.(*apps.DaemonSetUpdateStrategy), b.(*v1beta1.DaemonSetUpdateStrategy), scope)
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.HTTPIngressPath)(nil), (*extensions.HTTPIngressPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_HTTPIngressPath_To_extensions_HTTPIngressPath(a.(*v1beta1.HTTPIngressPath), b.(*extensions.HTTPIngressPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.HTTPIngressPath)(nil), (*v1beta1.HTTPIngressPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_HTTPIngressPath_To_v1beta1_HTTPIngressPath(a.(*extensions.HTTPIngressPath), b.(*v1beta1.HTTPIngressPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.HTTPIngressRuleValue)(nil), (*extensions.HTTPIngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_HTTPIngressRuleValue_To_extensions_HTTPIngressRuleValue(a.(*v1beta1.HTTPIngressRuleValue), b.(*extensions.HTTPIngressRuleValue), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.HTTPIngressRuleValue)(nil), (*v1beta1.HTTPIngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(a.(*extensions.HTTPIngressRuleValue), b.(*v1beta1.HTTPIngressRuleValue), scope)
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.Ingress)(nil), (*extensions.Ingress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Ingress_To_extensions_Ingress(a.(*v1beta1.Ingress), b.(*extensions.Ingress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.Ingress)(nil), (*v1beta1.Ingress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_Ingress_To_v1beta1_Ingress(a.(*extensions.Ingress), b.(*v1beta1.Ingress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressBackend)(nil), (*extensions.IngressBackend)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressBackend_To_extensions_IngressBackend(a.(*v1beta1.IngressBackend), b.(*extensions.IngressBackend), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressBackend)(nil), (*v1beta1.IngressBackend)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressBackend_To_v1beta1_IngressBackend(a.(*extensions.IngressBackend), b.(*v1beta1.IngressBackend), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressList)(nil), (*extensions.IngressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressList_To_extensions_IngressList(a.(*v1beta1.IngressList), b.(*extensions.IngressList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressList)(nil), (*v1beta1.IngressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressList_To_v1beta1_IngressList(a.(*extensions.IngressList), b.(*v1beta1.IngressList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressRule)(nil), (*extensions.IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressRule_To_extensions_IngressRule(a.(*v1beta1.IngressRule), b.(*extensions.IngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressRule)(nil), (*v1beta1.IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressRule_To_v1beta1_IngressRule(a.(*extensions.IngressRule), b.(*v1beta1.IngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressRuleValue)(nil), (*extensions.IngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressRuleValue_To_extensions_IngressRuleValue(a.(*v1beta1.IngressRuleValue), b.(*extensions.IngressRuleValue), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressRuleValue)(nil), (*v1beta1.IngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressRuleValue_To_v1beta1_IngressRuleValue(a.(*extensions.IngressRuleValue), b.(*v1beta1.IngressRuleValue), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressSpec)(nil), (*extensions.IngressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressSpec_To_extensions_IngressSpec(a.(*v1beta1.IngressSpec), b.(*extensions.IngressSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressSpec)(nil), (*v1beta1.IngressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressSpec_To_v1beta1_IngressSpec(a.(*extensions.IngressSpec), b.(*v1beta1.IngressSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressStatus)(nil), (*extensions.IngressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressStatus_To_extensions_IngressStatus(a.(*v1beta1.IngressStatus), b.(*extensions.IngressStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressStatus)(nil), (*v1beta1.IngressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressStatus_To_v1beta1_IngressStatus(a.(*extensions.IngressStatus), b.(*v1beta1.IngressStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.IngressTLS)(nil), (*extensions.IngressTLS)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IngressTLS_To_extensions_IngressTLS(a.(*v1beta1.IngressTLS), b.(*extensions.IngressTLS), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.IngressTLS)(nil), (*v1beta1.IngressTLS)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_IngressTLS_To_v1beta1_IngressTLS(a.(*extensions.IngressTLS), b.(*v1beta1.IngressTLS), scope)
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicaSet)(nil), (*apps.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSet_To_apps_ReplicaSet(a.(*v1beta1.ReplicaSet), b.(*apps.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSet)(nil), (*v1beta1.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSet_To_v1beta1_ReplicaSet(a.(*apps.ReplicaSet), b.(*v1beta1.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicaSetCondition)(nil), (*apps.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSetCondition_To_apps_ReplicaSetCondition(a.(*v1beta1.ReplicaSetCondition), b.(*apps.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetCondition)(nil), (*v1beta1.ReplicaSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetCondition_To_v1beta1_ReplicaSetCondition(a.(*apps.ReplicaSetCondition), b.(*v1beta1.ReplicaSetCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicaSetList)(nil), (*apps.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSetList_To_apps_ReplicaSetList(a.(*v1beta1.ReplicaSetList), b.(*apps.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetList)(nil), (*v1beta1.ReplicaSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetList_To_v1beta1_ReplicaSetList(a.(*apps.ReplicaSetList), b.(*v1beta1.ReplicaSetList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1beta1.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1beta1.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1beta1_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1beta1.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicaSetStatus)(nil), (*apps.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSetStatus_To_apps_ReplicaSetStatus(a.(*v1beta1.ReplicaSetStatus), b.(*apps.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.ReplicaSetStatus)(nil), (*v1beta1.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetStatus_To_v1beta1_ReplicaSetStatus(a.(*apps.ReplicaSetStatus), b.(*v1beta1.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ReplicationControllerDummy)(nil), (*extensions.ReplicationControllerDummy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicationControllerDummy_To_extensions_ReplicationControllerDummy(a.(*v1beta1.ReplicationControllerDummy), b.(*extensions.ReplicationControllerDummy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*extensions.ReplicationControllerDummy)(nil), (*v1beta1.ReplicationControllerDummy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_extensions_ReplicationControllerDummy_To_v1beta1_ReplicationControllerDummy(a.(*extensions.ReplicationControllerDummy), b.(*v1beta1.ReplicationControllerDummy), scope)
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1beta1.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1beta1.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1beta1_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1beta1.RollingUpdateDaemonSet), scope)
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
 if err := s.AddConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1beta1.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1beta1_ReplicaSetSpec(a.(*apps.ReplicaSetSpec), b.(*v1beta1.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDaemonSet)(nil), (*v1beta1.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDaemonSet_To_v1beta1_RollingUpdateDaemonSet(a.(*apps.RollingUpdateDaemonSet), b.(*v1beta1.RollingUpdateDaemonSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.RollingUpdateDeployment)(nil), (*v1beta1.RollingUpdateDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_RollingUpdateDeployment_To_v1beta1_RollingUpdateDeployment(a.(*apps.RollingUpdateDeployment), b.(*v1beta1.RollingUpdateDeployment), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*autoscaling.ScaleStatus)(nil), (*v1beta1.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_autoscaling_ScaleStatus_To_v1beta1_ScaleStatus(a.(*autoscaling.ScaleStatus), b.(*v1beta1.ScaleStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.IPBlock)(nil), (*v1beta1.IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_IPBlock_To_v1beta1_IPBlock(a.(*networking.IPBlock), b.(*v1beta1.IPBlock), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicyEgressRule)(nil), (*v1beta1.NetworkPolicyEgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyEgressRule_To_v1beta1_NetworkPolicyEgressRule(a.(*networking.NetworkPolicyEgressRule), b.(*v1beta1.NetworkPolicyEgressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicyIngressRule)(nil), (*v1beta1.NetworkPolicyIngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyIngressRule_To_v1beta1_NetworkPolicyIngressRule(a.(*networking.NetworkPolicyIngressRule), b.(*v1beta1.NetworkPolicyIngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicyList)(nil), (*v1beta1.NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList(a.(*networking.NetworkPolicyList), b.(*v1beta1.NetworkPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicyPeer)(nil), (*v1beta1.NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(a.(*networking.NetworkPolicyPeer), b.(*v1beta1.NetworkPolicyPeer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicyPort)(nil), (*v1beta1.NetworkPolicyPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyPort_To_v1beta1_NetworkPolicyPort(a.(*networking.NetworkPolicyPort), b.(*v1beta1.NetworkPolicyPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicySpec)(nil), (*v1beta1.NetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicySpec_To_v1beta1_NetworkPolicySpec(a.(*networking.NetworkPolicySpec), b.(*v1beta1.NetworkPolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*networking.NetworkPolicy)(nil), (*v1beta1.NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy(a.(*networking.NetworkPolicy), b.(*v1beta1.NetworkPolicy), scope)
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
 if err := s.AddConversionFunc((*v1beta1.IPBlock)(nil), (*networking.IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_IPBlock_To_networking_IPBlock(a.(*v1beta1.IPBlock), b.(*networking.IPBlock), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicyEgressRule)(nil), (*networking.NetworkPolicyEgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(a.(*v1beta1.NetworkPolicyEgressRule), b.(*networking.NetworkPolicyEgressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicyIngressRule)(nil), (*networking.NetworkPolicyIngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(a.(*v1beta1.NetworkPolicyIngressRule), b.(*networking.NetworkPolicyIngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicyList)(nil), (*networking.NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList(a.(*v1beta1.NetworkPolicyList), b.(*networking.NetworkPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicyPeer)(nil), (*networking.NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(a.(*v1beta1.NetworkPolicyPeer), b.(*networking.NetworkPolicyPeer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicyPort)(nil), (*networking.NetworkPolicyPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicyPort_To_networking_NetworkPolicyPort(a.(*v1beta1.NetworkPolicyPort), b.(*networking.NetworkPolicyPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicySpec)(nil), (*networking.NetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicySpec_To_networking_NetworkPolicySpec(a.(*v1beta1.NetworkPolicySpec), b.(*networking.NetworkPolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.NetworkPolicy)(nil), (*networking.NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy(a.(*v1beta1.NetworkPolicy), b.(*networking.NetworkPolicy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.ReplicaSetSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ReplicaSetSpec_To_apps_ReplicaSetSpec(a.(*v1beta1.ReplicaSetSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.RollingUpdateDaemonSet)(nil), (*apps.RollingUpdateDaemonSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(a.(*v1beta1.RollingUpdateDaemonSet), b.(*apps.RollingUpdateDaemonSet), scope)
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
func autoConvert_v1beta1_DaemonSet_To_apps_DaemonSet(in *v1beta1.DaemonSet, out *apps.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_DaemonSetSpec_To_apps_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_DaemonSetStatus_To_apps_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_DaemonSet_To_apps_DaemonSet(in *v1beta1.DaemonSet, out *apps.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSet_To_apps_DaemonSet(in, out, s)
}
func autoConvert_apps_DaemonSet_To_v1beta1_DaemonSet(in *apps.DaemonSet, out *v1beta1.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_DaemonSetSpec_To_v1beta1_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetStatus_To_v1beta1_DaemonSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_DaemonSet_To_v1beta1_DaemonSet(in *apps.DaemonSet, out *v1beta1.DaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSet_To_v1beta1_DaemonSet(in, out, s)
}
func autoConvert_v1beta1_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1beta1.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta1_DaemonSetCondition_To_apps_DaemonSetCondition(in *v1beta1.DaemonSetCondition, out *apps.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSetCondition_To_apps_DaemonSetCondition(in, out, s)
}
func autoConvert_apps_DaemonSetCondition_To_v1beta1_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1beta1.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.DaemonSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_DaemonSetCondition_To_v1beta1_DaemonSetCondition(in *apps.DaemonSetCondition, out *v1beta1.DaemonSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetCondition_To_v1beta1_DaemonSetCondition(in, out, s)
}
func autoConvert_v1beta1_DaemonSetList_To_apps_DaemonSetList(in *v1beta1.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_DaemonSet_To_apps_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_DaemonSetList_To_apps_DaemonSetList(in *v1beta1.DaemonSetList, out *apps.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSetList_To_apps_DaemonSetList(in, out, s)
}
func autoConvert_apps_DaemonSetList_To_v1beta1_DaemonSetList(in *apps.DaemonSetList, out *v1beta1.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.DaemonSet, len(*in))
  for i := range *in {
   if err := Convert_apps_DaemonSet_To_v1beta1_DaemonSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_DaemonSetList_To_v1beta1_DaemonSetList(in *apps.DaemonSetList, out *v1beta1.DaemonSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetList_To_v1beta1_DaemonSetList(in, out, s)
}
func autoConvert_v1beta1_DaemonSetSpec_To_apps_DaemonSetSpec(in *v1beta1.DaemonSetSpec, out *apps.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.TemplateGeneration = in.TemplateGeneration
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func Convert_v1beta1_DaemonSetSpec_To_apps_DaemonSetSpec(in *v1beta1.DaemonSetSpec, out *apps.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSetSpec_To_apps_DaemonSetSpec(in, out, s)
}
func autoConvert_apps_DaemonSetSpec_To_v1beta1_DaemonSetSpec(in *apps.DaemonSetSpec, out *v1beta1.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 if err := Convert_apps_DaemonSetUpdateStrategy_To_v1beta1_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.TemplateGeneration = in.TemplateGeneration
 out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
 return nil
}
func Convert_apps_DaemonSetSpec_To_v1beta1_DaemonSetSpec(in *apps.DaemonSetSpec, out *v1beta1.DaemonSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetSpec_To_v1beta1_DaemonSetSpec(in, out, s)
}
func autoConvert_v1beta1_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1beta1.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
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
func Convert_v1beta1_DaemonSetStatus_To_apps_DaemonSetStatus(in *v1beta1.DaemonSetStatus, out *apps.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSetStatus_To_apps_DaemonSetStatus(in, out, s)
}
func autoConvert_apps_DaemonSetStatus_To_v1beta1_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1beta1.DaemonSetStatus, s conversion.Scope) error {
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
 out.Conditions = *(*[]v1beta1.DaemonSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_DaemonSetStatus_To_v1beta1_DaemonSetStatus(in *apps.DaemonSetStatus, out *v1beta1.DaemonSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetStatus_To_v1beta1_DaemonSetStatus(in, out, s)
}
func autoConvert_v1beta1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in *v1beta1.DaemonSetUpdateStrategy, out *apps.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(apps.RollingUpdateDaemonSet)
  if err := Convert_v1beta1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_v1beta1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in *v1beta1.DaemonSetUpdateStrategy, out *apps.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_DaemonSetUpdateStrategy_To_apps_DaemonSetUpdateStrategy(in, out, s)
}
func autoConvert_apps_DaemonSetUpdateStrategy_To_v1beta1_DaemonSetUpdateStrategy(in *apps.DaemonSetUpdateStrategy, out *v1beta1.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.DaemonSetUpdateStrategyType(in.Type)
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(v1beta1.RollingUpdateDaemonSet)
  if err := Convert_apps_RollingUpdateDaemonSet_To_v1beta1_RollingUpdateDaemonSet(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.RollingUpdate = nil
 }
 return nil
}
func Convert_apps_DaemonSetUpdateStrategy_To_v1beta1_DaemonSetUpdateStrategy(in *apps.DaemonSetUpdateStrategy, out *v1beta1.DaemonSetUpdateStrategy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_DaemonSetUpdateStrategy_To_v1beta1_DaemonSetUpdateStrategy(in, out, s)
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
func autoConvert_v1beta1_HTTPIngressPath_To_extensions_HTTPIngressPath(in *v1beta1.HTTPIngressPath, out *extensions.HTTPIngressPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 if err := Convert_v1beta1_IngressBackend_To_extensions_IngressBackend(&in.Backend, &out.Backend, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_HTTPIngressPath_To_extensions_HTTPIngressPath(in *v1beta1.HTTPIngressPath, out *extensions.HTTPIngressPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_HTTPIngressPath_To_extensions_HTTPIngressPath(in, out, s)
}
func autoConvert_extensions_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in *extensions.HTTPIngressPath, out *v1beta1.HTTPIngressPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 if err := Convert_extensions_IngressBackend_To_v1beta1_IngressBackend(&in.Backend, &out.Backend, s); err != nil {
  return err
 }
 return nil
}
func Convert_extensions_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in *extensions.HTTPIngressPath, out *v1beta1.HTTPIngressPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in, out, s)
}
func autoConvert_v1beta1_HTTPIngressRuleValue_To_extensions_HTTPIngressRuleValue(in *v1beta1.HTTPIngressRuleValue, out *extensions.HTTPIngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Paths = *(*[]extensions.HTTPIngressPath)(unsafe.Pointer(&in.Paths))
 return nil
}
func Convert_v1beta1_HTTPIngressRuleValue_To_extensions_HTTPIngressRuleValue(in *v1beta1.HTTPIngressRuleValue, out *extensions.HTTPIngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_HTTPIngressRuleValue_To_extensions_HTTPIngressRuleValue(in, out, s)
}
func autoConvert_extensions_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in *extensions.HTTPIngressRuleValue, out *v1beta1.HTTPIngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Paths = *(*[]v1beta1.HTTPIngressPath)(unsafe.Pointer(&in.Paths))
 return nil
}
func Convert_extensions_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in *extensions.HTTPIngressRuleValue, out *v1beta1.HTTPIngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in, out, s)
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
func autoConvert_v1beta1_Ingress_To_extensions_Ingress(in *v1beta1.Ingress, out *extensions.Ingress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_IngressSpec_To_extensions_IngressSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_IngressStatus_To_extensions_IngressStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_Ingress_To_extensions_Ingress(in *v1beta1.Ingress, out *extensions.Ingress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Ingress_To_extensions_Ingress(in, out, s)
}
func autoConvert_extensions_Ingress_To_v1beta1_Ingress(in *extensions.Ingress, out *v1beta1.Ingress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_extensions_IngressSpec_To_v1beta1_IngressSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_extensions_IngressStatus_To_v1beta1_IngressStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_extensions_Ingress_To_v1beta1_Ingress(in *extensions.Ingress, out *v1beta1.Ingress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_Ingress_To_v1beta1_Ingress(in, out, s)
}
func autoConvert_v1beta1_IngressBackend_To_extensions_IngressBackend(in *v1beta1.IngressBackend, out *extensions.IngressBackend, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceName = in.ServiceName
 out.ServicePort = in.ServicePort
 return nil
}
func Convert_v1beta1_IngressBackend_To_extensions_IngressBackend(in *v1beta1.IngressBackend, out *extensions.IngressBackend, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressBackend_To_extensions_IngressBackend(in, out, s)
}
func autoConvert_extensions_IngressBackend_To_v1beta1_IngressBackend(in *extensions.IngressBackend, out *v1beta1.IngressBackend, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceName = in.ServiceName
 out.ServicePort = in.ServicePort
 return nil
}
func Convert_extensions_IngressBackend_To_v1beta1_IngressBackend(in *extensions.IngressBackend, out *v1beta1.IngressBackend, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressBackend_To_v1beta1_IngressBackend(in, out, s)
}
func autoConvert_v1beta1_IngressList_To_extensions_IngressList(in *v1beta1.IngressList, out *extensions.IngressList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]extensions.Ingress)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_IngressList_To_extensions_IngressList(in *v1beta1.IngressList, out *extensions.IngressList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressList_To_extensions_IngressList(in, out, s)
}
func autoConvert_extensions_IngressList_To_v1beta1_IngressList(in *extensions.IngressList, out *v1beta1.IngressList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.Ingress)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_extensions_IngressList_To_v1beta1_IngressList(in *extensions.IngressList, out *v1beta1.IngressList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressList_To_v1beta1_IngressList(in, out, s)
}
func autoConvert_v1beta1_IngressRule_To_extensions_IngressRule(in *v1beta1.IngressRule, out *extensions.IngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Host = in.Host
 if err := Convert_v1beta1_IngressRuleValue_To_extensions_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_IngressRule_To_extensions_IngressRule(in *v1beta1.IngressRule, out *extensions.IngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressRule_To_extensions_IngressRule(in, out, s)
}
func autoConvert_extensions_IngressRule_To_v1beta1_IngressRule(in *extensions.IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Host = in.Host
 if err := Convert_extensions_IngressRuleValue_To_v1beta1_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, s); err != nil {
  return err
 }
 return nil
}
func Convert_extensions_IngressRule_To_v1beta1_IngressRule(in *extensions.IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressRule_To_v1beta1_IngressRule(in, out, s)
}
func autoConvert_v1beta1_IngressRuleValue_To_extensions_IngressRuleValue(in *v1beta1.IngressRuleValue, out *extensions.IngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HTTP = (*extensions.HTTPIngressRuleValue)(unsafe.Pointer(in.HTTP))
 return nil
}
func Convert_v1beta1_IngressRuleValue_To_extensions_IngressRuleValue(in *v1beta1.IngressRuleValue, out *extensions.IngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressRuleValue_To_extensions_IngressRuleValue(in, out, s)
}
func autoConvert_extensions_IngressRuleValue_To_v1beta1_IngressRuleValue(in *extensions.IngressRuleValue, out *v1beta1.IngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HTTP = (*v1beta1.HTTPIngressRuleValue)(unsafe.Pointer(in.HTTP))
 return nil
}
func Convert_extensions_IngressRuleValue_To_v1beta1_IngressRuleValue(in *extensions.IngressRuleValue, out *v1beta1.IngressRuleValue, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressRuleValue_To_v1beta1_IngressRuleValue(in, out, s)
}
func autoConvert_v1beta1_IngressSpec_To_extensions_IngressSpec(in *v1beta1.IngressSpec, out *extensions.IngressSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Backend = (*extensions.IngressBackend)(unsafe.Pointer(in.Backend))
 out.TLS = *(*[]extensions.IngressTLS)(unsafe.Pointer(&in.TLS))
 out.Rules = *(*[]extensions.IngressRule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_v1beta1_IngressSpec_To_extensions_IngressSpec(in *v1beta1.IngressSpec, out *extensions.IngressSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressSpec_To_extensions_IngressSpec(in, out, s)
}
func autoConvert_extensions_IngressSpec_To_v1beta1_IngressSpec(in *extensions.IngressSpec, out *v1beta1.IngressSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Backend = (*v1beta1.IngressBackend)(unsafe.Pointer(in.Backend))
 out.TLS = *(*[]v1beta1.IngressTLS)(unsafe.Pointer(&in.TLS))
 out.Rules = *(*[]v1beta1.IngressRule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_extensions_IngressSpec_To_v1beta1_IngressSpec(in *extensions.IngressSpec, out *v1beta1.IngressSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressSpec_To_v1beta1_IngressSpec(in, out, s)
}
func autoConvert_v1beta1_IngressStatus_To_extensions_IngressStatus(in *v1beta1.IngressStatus, out *extensions.IngressStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.Convert(&in.LoadBalancer, &out.LoadBalancer, 0); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_IngressStatus_To_extensions_IngressStatus(in *v1beta1.IngressStatus, out *extensions.IngressStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressStatus_To_extensions_IngressStatus(in, out, s)
}
func autoConvert_extensions_IngressStatus_To_v1beta1_IngressStatus(in *extensions.IngressStatus, out *v1beta1.IngressStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.Convert(&in.LoadBalancer, &out.LoadBalancer, 0); err != nil {
  return err
 }
 return nil
}
func Convert_extensions_IngressStatus_To_v1beta1_IngressStatus(in *extensions.IngressStatus, out *v1beta1.IngressStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressStatus_To_v1beta1_IngressStatus(in, out, s)
}
func autoConvert_v1beta1_IngressTLS_To_extensions_IngressTLS(in *v1beta1.IngressTLS, out *extensions.IngressTLS, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hosts = *(*[]string)(unsafe.Pointer(&in.Hosts))
 out.SecretName = in.SecretName
 return nil
}
func Convert_v1beta1_IngressTLS_To_extensions_IngressTLS(in *v1beta1.IngressTLS, out *extensions.IngressTLS, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_IngressTLS_To_extensions_IngressTLS(in, out, s)
}
func autoConvert_extensions_IngressTLS_To_v1beta1_IngressTLS(in *extensions.IngressTLS, out *v1beta1.IngressTLS, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hosts = *(*[]string)(unsafe.Pointer(&in.Hosts))
 out.SecretName = in.SecretName
 return nil
}
func Convert_extensions_IngressTLS_To_v1beta1_IngressTLS(in *extensions.IngressTLS, out *v1beta1.IngressTLS, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_IngressTLS_To_v1beta1_IngressTLS(in, out, s)
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
 if err := metav1.Convert_Pointer_bool_To_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
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
 out.DefaultAddCapabilities = *(*[]v1.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
 out.RequiredDropCapabilities = *(*[]v1.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
 out.AllowedCapabilities = *(*[]v1.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
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
 if err := metav1.Convert_bool_To_Pointer_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
  return err
 }
 out.AllowedHostPaths = *(*[]v1beta1.AllowedHostPath)(unsafe.Pointer(&in.AllowedHostPaths))
 out.AllowedFlexVolumes = *(*[]v1beta1.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
 out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
 out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
 out.AllowedProcMountTypes = *(*[]v1.ProcMountType)(unsafe.Pointer(&in.AllowedProcMountTypes))
 return nil
}
func Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in *policy.PodSecurityPolicySpec, out *v1beta1.PodSecurityPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in, out, s)
}
func autoConvert_v1beta1_ReplicaSet_To_apps_ReplicaSet(in *v1beta1.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_ReplicaSetSpec_To_apps_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_ReplicaSetStatus_To_apps_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_ReplicaSet_To_apps_ReplicaSet(in *v1beta1.ReplicaSet, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ReplicaSet_To_apps_ReplicaSet(in, out, s)
}
func autoConvert_apps_ReplicaSet_To_v1beta1_ReplicaSet(in *apps.ReplicaSet, out *v1beta1.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_ReplicaSetSpec_To_v1beta1_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_apps_ReplicaSetStatus_To_v1beta1_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_ReplicaSet_To_v1beta1_ReplicaSet(in *apps.ReplicaSet, out *v1beta1.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSet_To_v1beta1_ReplicaSet(in, out, s)
}
func autoConvert_v1beta1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1beta1.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = apps.ReplicaSetConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1beta1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in *v1beta1.ReplicaSetCondition, out *apps.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ReplicaSetCondition_To_apps_ReplicaSetCondition(in, out, s)
}
func autoConvert_apps_ReplicaSetCondition_To_v1beta1_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1beta1.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.ReplicaSetConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_apps_ReplicaSetCondition_To_v1beta1_ReplicaSetCondition(in *apps.ReplicaSetCondition, out *v1beta1.ReplicaSetCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetCondition_To_v1beta1_ReplicaSetCondition(in, out, s)
}
func autoConvert_v1beta1_ReplicaSetList_To_apps_ReplicaSetList(in *v1beta1.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]apps.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_ReplicaSet_To_apps_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_ReplicaSetList_To_apps_ReplicaSetList(in *v1beta1.ReplicaSetList, out *apps.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ReplicaSetList_To_apps_ReplicaSetList(in, out, s)
}
func autoConvert_apps_ReplicaSetList_To_v1beta1_ReplicaSetList(in *apps.ReplicaSetList, out *v1beta1.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.ReplicaSet, len(*in))
  for i := range *in {
   if err := Convert_apps_ReplicaSet_To_v1beta1_ReplicaSet(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_apps_ReplicaSetList_To_v1beta1_ReplicaSetList(in *apps.ReplicaSetList, out *v1beta1.ReplicaSetList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetList_To_v1beta1_ReplicaSetList(in, out, s)
}
func autoConvert_v1beta1_ReplicaSetSpec_To_apps_ReplicaSetSpec(in *v1beta1.ReplicaSetSpec, out *apps.ReplicaSetSpec, s conversion.Scope) error {
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
func autoConvert_apps_ReplicaSetSpec_To_v1beta1_ReplicaSetSpec(in *apps.ReplicaSetSpec, out *v1beta1.ReplicaSetSpec, s conversion.Scope) error {
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
func autoConvert_v1beta1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1beta1.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
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
func Convert_v1beta1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in *v1beta1.ReplicaSetStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ReplicaSetStatus_To_apps_ReplicaSetStatus(in, out, s)
}
func autoConvert_apps_ReplicaSetStatus_To_v1beta1_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1beta1.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]v1beta1.ReplicaSetCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_apps_ReplicaSetStatus_To_v1beta1_ReplicaSetStatus(in *apps.ReplicaSetStatus, out *v1beta1.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_apps_ReplicaSetStatus_To_v1beta1_ReplicaSetStatus(in, out, s)
}
func autoConvert_v1beta1_ReplicationControllerDummy_To_extensions_ReplicationControllerDummy(in *v1beta1.ReplicationControllerDummy, out *extensions.ReplicationControllerDummy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func Convert_v1beta1_ReplicationControllerDummy_To_extensions_ReplicationControllerDummy(in *v1beta1.ReplicationControllerDummy, out *extensions.ReplicationControllerDummy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ReplicationControllerDummy_To_extensions_ReplicationControllerDummy(in, out, s)
}
func autoConvert_extensions_ReplicationControllerDummy_To_v1beta1_ReplicationControllerDummy(in *extensions.ReplicationControllerDummy, out *v1beta1.ReplicationControllerDummy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func Convert_extensions_ReplicationControllerDummy_To_v1beta1_ReplicationControllerDummy(in *extensions.ReplicationControllerDummy, out *v1beta1.ReplicationControllerDummy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_extensions_ReplicationControllerDummy_To_v1beta1_ReplicationControllerDummy(in, out, s)
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
func autoConvert_v1beta1_RollingUpdateDaemonSet_To_apps_RollingUpdateDaemonSet(in *v1beta1.RollingUpdateDaemonSet, out *apps.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func autoConvert_apps_RollingUpdateDaemonSet_To_v1beta1_RollingUpdateDaemonSet(in *apps.RollingUpdateDaemonSet, out *v1beta1.RollingUpdateDaemonSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
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
 out.SELinuxOptions = (*v1.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 return nil
}
func Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in *policy.SELinuxStrategyOptions, out *v1beta1.SELinuxStrategyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in, out, s)
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
