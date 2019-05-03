package v1beta1

import (
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 abac "k8s.io/kubernetes/pkg/apis/abac"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*Policy)(nil), (*abac.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Policy_To_abac_Policy(a.(*Policy), b.(*abac.Policy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*abac.Policy)(nil), (*Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_abac_Policy_To_v1beta1_Policy(a.(*abac.Policy), b.(*Policy), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*PolicySpec)(nil), (*abac.PolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PolicySpec_To_abac_PolicySpec(a.(*PolicySpec), b.(*abac.PolicySpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*abac.PolicySpec)(nil), (*PolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_abac_PolicySpec_To_v1beta1_PolicySpec(a.(*abac.PolicySpec), b.(*PolicySpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_Policy_To_abac_Policy(in *Policy, out *abac.Policy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1beta1_PolicySpec_To_abac_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_Policy_To_abac_Policy(in *Policy, out *abac.Policy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Policy_To_abac_Policy(in, out, s)
}
func autoConvert_abac_Policy_To_v1beta1_Policy(in *abac.Policy, out *Policy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_abac_PolicySpec_To_v1beta1_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_abac_Policy_To_v1beta1_Policy(in *abac.Policy, out *Policy, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_abac_Policy_To_v1beta1_Policy(in, out, s)
}
func autoConvert_v1beta1_PolicySpec_To_abac_PolicySpec(in *PolicySpec, out *abac.PolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.User = in.User
 out.Group = in.Group
 out.Readonly = in.Readonly
 out.APIGroup = in.APIGroup
 out.Resource = in.Resource
 out.Namespace = in.Namespace
 out.NonResourcePath = in.NonResourcePath
 return nil
}
func Convert_v1beta1_PolicySpec_To_abac_PolicySpec(in *PolicySpec, out *abac.PolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PolicySpec_To_abac_PolicySpec(in, out, s)
}
func autoConvert_abac_PolicySpec_To_v1beta1_PolicySpec(in *abac.PolicySpec, out *PolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.User = in.User
 out.Group = in.Group
 out.Readonly = in.Readonly
 out.APIGroup = in.APIGroup
 out.Resource = in.Resource
 out.Namespace = in.Namespace
 out.NonResourcePath = in.NonResourcePath
 return nil
}
func Convert_abac_PolicySpec_To_v1beta1_PolicySpec(in *abac.PolicySpec, out *PolicySpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_abac_PolicySpec_To_v1beta1_PolicySpec(in, out, s)
}
