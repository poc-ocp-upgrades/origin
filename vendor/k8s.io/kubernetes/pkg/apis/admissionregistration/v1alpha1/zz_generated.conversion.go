package v1alpha1

import (
 unsafe "unsafe"
 v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.Initializer)(nil), (*admissionregistration.Initializer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_Initializer_To_admissionregistration_Initializer(a.(*v1alpha1.Initializer), b.(*admissionregistration.Initializer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.Initializer)(nil), (*v1alpha1.Initializer)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_Initializer_To_v1alpha1_Initializer(a.(*admissionregistration.Initializer), b.(*v1alpha1.Initializer), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.InitializerConfiguration)(nil), (*admissionregistration.InitializerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(a.(*v1alpha1.InitializerConfiguration), b.(*admissionregistration.InitializerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.InitializerConfiguration)(nil), (*v1alpha1.InitializerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(a.(*admissionregistration.InitializerConfiguration), b.(*v1alpha1.InitializerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.InitializerConfigurationList)(nil), (*admissionregistration.InitializerConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(a.(*v1alpha1.InitializerConfigurationList), b.(*admissionregistration.InitializerConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.InitializerConfigurationList)(nil), (*v1alpha1.InitializerConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(a.(*admissionregistration.InitializerConfigurationList), b.(*v1alpha1.InitializerConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.Rule)(nil), (*admissionregistration.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_Rule_To_admissionregistration_Rule(a.(*v1alpha1.Rule), b.(*admissionregistration.Rule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.Rule)(nil), (*v1alpha1.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_Rule_To_v1alpha1_Rule(a.(*admissionregistration.Rule), b.(*v1alpha1.Rule), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *v1alpha1.Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Rules = *(*[]admissionregistration.Rule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *v1alpha1.Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in, out, s)
}
func autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *v1alpha1.Initializer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Rules = *(*[]v1alpha1.Rule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *v1alpha1.Initializer, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in, out, s)
}
func autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *v1alpha1.InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Initializers = *(*[]admissionregistration.Initializer)(unsafe.Pointer(&in.Initializers))
 return nil
}
func Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *v1alpha1.InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in, out, s)
}
func autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *v1alpha1.InitializerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Initializers = *(*[]v1alpha1.Initializer)(unsafe.Pointer(&in.Initializers))
 return nil
}
func Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *v1alpha1.InitializerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *v1alpha1.InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]admissionregistration.InitializerConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *v1alpha1.InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in, out, s)
}
func autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *v1alpha1.InitializerConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1alpha1.InitializerConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *v1alpha1.InitializerConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in, out, s)
}
func autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in *v1alpha1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 return nil
}
func Convert_v1alpha1_Rule_To_admissionregistration_Rule(in *v1alpha1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in, out, s)
}
func autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *v1alpha1.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 return nil
}
func Convert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *v1alpha1.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in, out, s)
}
