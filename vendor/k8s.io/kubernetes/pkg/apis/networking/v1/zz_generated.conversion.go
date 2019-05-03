package v1

import (
 unsafe "unsafe"
 corev1 "k8s.io/api/core/v1"
 v1 "k8s.io/api/networking/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 intstr "k8s.io/apimachinery/pkg/util/intstr"
 core "k8s.io/kubernetes/pkg/apis/core"
 networking "k8s.io/kubernetes/pkg/apis/networking"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1.IPBlock)(nil), (*networking.IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_IPBlock_To_networking_IPBlock(a.(*v1.IPBlock), b.(*networking.IPBlock), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.IPBlock)(nil), (*v1.IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_IPBlock_To_v1_IPBlock(a.(*networking.IPBlock), b.(*v1.IPBlock), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicy)(nil), (*networking.NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicy_To_networking_NetworkPolicy(a.(*v1.NetworkPolicy), b.(*networking.NetworkPolicy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicy)(nil), (*v1.NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicy_To_v1_NetworkPolicy(a.(*networking.NetworkPolicy), b.(*v1.NetworkPolicy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicyEgressRule)(nil), (*networking.NetworkPolicyEgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(a.(*v1.NetworkPolicyEgressRule), b.(*networking.NetworkPolicyEgressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyEgressRule)(nil), (*v1.NetworkPolicyEgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(a.(*networking.NetworkPolicyEgressRule), b.(*v1.NetworkPolicyEgressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicyIngressRule)(nil), (*networking.NetworkPolicyIngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(a.(*v1.NetworkPolicyIngressRule), b.(*networking.NetworkPolicyIngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyIngressRule)(nil), (*v1.NetworkPolicyIngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(a.(*networking.NetworkPolicyIngressRule), b.(*v1.NetworkPolicyIngressRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicyList)(nil), (*networking.NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(a.(*v1.NetworkPolicyList), b.(*networking.NetworkPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyList)(nil), (*v1.NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(a.(*networking.NetworkPolicyList), b.(*v1.NetworkPolicyList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicyPeer)(nil), (*networking.NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(a.(*v1.NetworkPolicyPeer), b.(*networking.NetworkPolicyPeer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyPeer)(nil), (*v1.NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(a.(*networking.NetworkPolicyPeer), b.(*v1.NetworkPolicyPeer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicyPort)(nil), (*networking.NetworkPolicyPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(a.(*v1.NetworkPolicyPort), b.(*networking.NetworkPolicyPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyPort)(nil), (*v1.NetworkPolicyPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(a.(*networking.NetworkPolicyPort), b.(*v1.NetworkPolicyPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NetworkPolicySpec)(nil), (*networking.NetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(a.(*v1.NetworkPolicySpec), b.(*networking.NetworkPolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicySpec)(nil), (*v1.NetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(a.(*networking.NetworkPolicySpec), b.(*v1.NetworkPolicySpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_IPBlock_To_networking_IPBlock(in *v1.IPBlock, out *networking.IPBlock, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CIDR = in.CIDR
 out.Except = *(*[]string)(unsafe.Pointer(&in.Except))
 return nil
}
func Convert_v1_IPBlock_To_networking_IPBlock(in *v1.IPBlock, out *networking.IPBlock, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_IPBlock_To_networking_IPBlock(in, out, s)
}
func autoConvert_networking_IPBlock_To_v1_IPBlock(in *networking.IPBlock, out *v1.IPBlock, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CIDR = in.CIDR
 out.Except = *(*[]string)(unsafe.Pointer(&in.Except))
 return nil
}
func Convert_networking_IPBlock_To_v1_IPBlock(in *networking.IPBlock, out *v1.IPBlock, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_IPBlock_To_v1_IPBlock(in, out, s)
}
func autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *v1.NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *v1.NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in, out, s)
}
func autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *v1.NetworkPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *v1.NetworkPolicy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in, out, s)
}
func autoConvert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in *v1.NetworkPolicyEgressRule, out *networking.NetworkPolicyEgressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]networking.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
 out.To = *(*[]networking.NetworkPolicyPeer)(unsafe.Pointer(&in.To))
 return nil
}
func Convert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in *v1.NetworkPolicyEgressRule, out *networking.NetworkPolicyEgressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in, out, s)
}
func autoConvert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in *networking.NetworkPolicyEgressRule, out *v1.NetworkPolicyEgressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]v1.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
 out.To = *(*[]v1.NetworkPolicyPeer)(unsafe.Pointer(&in.To))
 return nil
}
func Convert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in *networking.NetworkPolicyEgressRule, out *v1.NetworkPolicyEgressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in, out, s)
}
func autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *v1.NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]networking.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
 out.From = *(*[]networking.NetworkPolicyPeer)(unsafe.Pointer(&in.From))
 return nil
}
func Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *v1.NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in, out, s)
}
func autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *v1.NetworkPolicyIngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]v1.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
 out.From = *(*[]v1.NetworkPolicyPeer)(unsafe.Pointer(&in.From))
 return nil
}
func Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *v1.NetworkPolicyIngressRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in, out, s)
}
func autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *v1.NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]networking.NetworkPolicy)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *v1.NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in, out, s)
}
func autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *v1.NetworkPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.NetworkPolicy)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *v1.NetworkPolicyList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in, out, s)
}
func autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *v1.NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.PodSelector))
 out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
 out.IPBlock = (*networking.IPBlock)(unsafe.Pointer(in.IPBlock))
 return nil
}
func Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *v1.NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in, out, s)
}
func autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *v1.NetworkPolicyPeer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.PodSelector))
 out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
 out.IPBlock = (*v1.IPBlock)(unsafe.Pointer(in.IPBlock))
 return nil
}
func Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *v1.NetworkPolicyPeer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in, out, s)
}
func autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *v1.NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Protocol = (*core.Protocol)(unsafe.Pointer(in.Protocol))
 out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
 return nil
}
func Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *v1.NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in, out, s)
}
func autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *v1.NetworkPolicyPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Protocol = (*corev1.Protocol)(unsafe.Pointer(in.Protocol))
 out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
 return nil
}
func Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *v1.NetworkPolicyPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in, out, s)
}
func autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *v1.NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodSelector = in.PodSelector
 out.Ingress = *(*[]networking.NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
 out.Egress = *(*[]networking.NetworkPolicyEgressRule)(unsafe.Pointer(&in.Egress))
 out.PolicyTypes = *(*[]networking.PolicyType)(unsafe.Pointer(&in.PolicyTypes))
 return nil
}
func Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *v1.NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in, out, s)
}
func autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *v1.NetworkPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodSelector = in.PodSelector
 out.Ingress = *(*[]v1.NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
 out.Egress = *(*[]v1.NetworkPolicyEgressRule)(unsafe.Pointer(&in.Egress))
 out.PolicyTypes = *(*[]v1.PolicyType)(unsafe.Pointer(&in.PolicyTypes))
 return nil
}
func Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *v1.NetworkPolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in, out, s)
}
