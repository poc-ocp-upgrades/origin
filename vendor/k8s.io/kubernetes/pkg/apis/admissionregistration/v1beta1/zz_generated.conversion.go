package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/admissionregistration/v1beta1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.MutatingWebhookConfiguration)(nil), (*admissionregistration.MutatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(a.(*v1beta1.MutatingWebhookConfiguration), b.(*admissionregistration.MutatingWebhookConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.MutatingWebhookConfiguration)(nil), (*v1beta1.MutatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(a.(*admissionregistration.MutatingWebhookConfiguration), b.(*v1beta1.MutatingWebhookConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.MutatingWebhookConfigurationList)(nil), (*admissionregistration.MutatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(a.(*v1beta1.MutatingWebhookConfigurationList), b.(*admissionregistration.MutatingWebhookConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.MutatingWebhookConfigurationList)(nil), (*v1beta1.MutatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(a.(*admissionregistration.MutatingWebhookConfigurationList), b.(*v1beta1.MutatingWebhookConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.Rule)(nil), (*admissionregistration.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Rule_To_admissionregistration_Rule(a.(*v1beta1.Rule), b.(*admissionregistration.Rule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.Rule)(nil), (*v1beta1.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_Rule_To_v1beta1_Rule(a.(*admissionregistration.Rule), b.(*v1beta1.Rule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.RuleWithOperations)(nil), (*admissionregistration.RuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(a.(*v1beta1.RuleWithOperations), b.(*admissionregistration.RuleWithOperations), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.RuleWithOperations)(nil), (*v1beta1.RuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(a.(*admissionregistration.RuleWithOperations), b.(*v1beta1.RuleWithOperations), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ServiceReference)(nil), (*admissionregistration.ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(a.(*v1beta1.ServiceReference), b.(*admissionregistration.ServiceReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.ServiceReference)(nil), (*v1beta1.ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(a.(*admissionregistration.ServiceReference), b.(*v1beta1.ServiceReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ValidatingWebhookConfiguration)(nil), (*admissionregistration.ValidatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(a.(*v1beta1.ValidatingWebhookConfiguration), b.(*admissionregistration.ValidatingWebhookConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingWebhookConfiguration)(nil), (*v1beta1.ValidatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(a.(*admissionregistration.ValidatingWebhookConfiguration), b.(*v1beta1.ValidatingWebhookConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.ValidatingWebhookConfigurationList)(nil), (*admissionregistration.ValidatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(a.(*v1beta1.ValidatingWebhookConfigurationList), b.(*admissionregistration.ValidatingWebhookConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingWebhookConfigurationList)(nil), (*v1beta1.ValidatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(a.(*admissionregistration.ValidatingWebhookConfigurationList), b.(*v1beta1.ValidatingWebhookConfigurationList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.Webhook)(nil), (*admissionregistration.Webhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Webhook_To_admissionregistration_Webhook(a.(*v1beta1.Webhook), b.(*admissionregistration.Webhook), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.Webhook)(nil), (*v1beta1.Webhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_Webhook_To_v1beta1_Webhook(a.(*admissionregistration.Webhook), b.(*v1beta1.Webhook), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.WebhookClientConfig)(nil), (*admissionregistration.WebhookClientConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(a.(*v1beta1.WebhookClientConfig), b.(*admissionregistration.WebhookClientConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admissionregistration.WebhookClientConfig)(nil), (*v1beta1.WebhookClientConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(a.(*admissionregistration.WebhookClientConfig), b.(*v1beta1.WebhookClientConfig), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1beta1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Webhooks = *(*[]admissionregistration.Webhook)(unsafe.Pointer(&in.Webhooks))
 return nil
}
func Convert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1beta1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in, out, s)
}
func autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1beta1.MutatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Webhooks = *(*[]v1beta1.Webhook)(unsafe.Pointer(&in.Webhooks))
 return nil
}
func Convert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1beta1.MutatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in, out, s)
}
func autoConvert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1beta1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]admissionregistration.MutatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1beta1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in, out, s)
}
func autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1beta1.MutatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.MutatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1beta1.MutatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in, out, s)
}
func autoConvert_v1beta1_Rule_To_admissionregistration_Rule(in *v1beta1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 return nil
}
func Convert_v1beta1_Rule_To_admissionregistration_Rule(in *v1beta1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Rule_To_admissionregistration_Rule(in, out, s)
}
func autoConvert_admissionregistration_Rule_To_v1beta1_Rule(in *admissionregistration.Rule, out *v1beta1.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 return nil
}
func Convert_admissionregistration_Rule_To_v1beta1_Rule(in *admissionregistration.Rule, out *v1beta1.Rule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_Rule_To_v1beta1_Rule(in, out, s)
}
func autoConvert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1beta1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Operations = *(*[]admissionregistration.OperationType)(unsafe.Pointer(&in.Operations))
 if err := Convert_v1beta1_Rule_To_admissionregistration_Rule(&in.Rule, &out.Rule, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1beta1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in, out, s)
}
func autoConvert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1beta1.RuleWithOperations, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Operations = *(*[]v1beta1.OperationType)(unsafe.Pointer(&in.Operations))
 if err := Convert_admissionregistration_Rule_To_v1beta1_Rule(&in.Rule, &out.Rule, s); err != nil {
  return err
 }
 return nil
}
func Convert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1beta1.RuleWithOperations, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in, out, s)
}
func autoConvert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in *v1beta1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.Path = (*string)(unsafe.Pointer(in.Path))
 return nil
}
func Convert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in *v1beta1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in, out, s)
}
func autoConvert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in *admissionregistration.ServiceReference, out *v1beta1.ServiceReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.Path = (*string)(unsafe.Pointer(in.Path))
 return nil
}
func Convert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in *admissionregistration.ServiceReference, out *v1beta1.ServiceReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in, out, s)
}
func autoConvert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1beta1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Webhooks = *(*[]admissionregistration.Webhook)(unsafe.Pointer(&in.Webhooks))
 return nil
}
func Convert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1beta1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in, out, s)
}
func autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1beta1.ValidatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Webhooks = *(*[]v1beta1.Webhook)(unsafe.Pointer(&in.Webhooks))
 return nil
}
func Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1beta1.ValidatingWebhookConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in, out, s)
}
func autoConvert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1beta1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]admissionregistration.ValidatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1beta1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in, out, s)
}
func autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1beta1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.ValidatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1beta1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in, out, s)
}
func autoConvert_v1beta1_Webhook_To_admissionregistration_Webhook(in *v1beta1.Webhook, out *admissionregistration.Webhook, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
  return err
 }
 out.Rules = *(*[]admissionregistration.RuleWithOperations)(unsafe.Pointer(&in.Rules))
 out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
 out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
 out.SideEffects = (*admissionregistration.SideEffectClass)(unsafe.Pointer(in.SideEffects))
 return nil
}
func Convert_v1beta1_Webhook_To_admissionregistration_Webhook(in *v1beta1.Webhook, out *admissionregistration.Webhook, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Webhook_To_admissionregistration_Webhook(in, out, s)
}
func autoConvert_admissionregistration_Webhook_To_v1beta1_Webhook(in *admissionregistration.Webhook, out *v1beta1.Webhook, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
  return err
 }
 out.Rules = *(*[]v1beta1.RuleWithOperations)(unsafe.Pointer(&in.Rules))
 out.FailurePolicy = (*v1beta1.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
 out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
 out.SideEffects = (*v1beta1.SideEffectClass)(unsafe.Pointer(in.SideEffects))
 return nil
}
func Convert_admissionregistration_Webhook_To_v1beta1_Webhook(in *admissionregistration.Webhook, out *v1beta1.Webhook, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_Webhook_To_v1beta1_Webhook(in, out, s)
}
func autoConvert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1beta1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.URL = (*string)(unsafe.Pointer(in.URL))
 out.Service = (*admissionregistration.ServiceReference)(unsafe.Pointer(in.Service))
 out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
 return nil
}
func Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1beta1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in, out, s)
}
func autoConvert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1beta1.WebhookClientConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.URL = (*string)(unsafe.Pointer(in.URL))
 out.Service = (*v1beta1.ServiceReference)(unsafe.Pointer(in.Service))
 out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
 return nil
}
func Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1beta1.WebhookClientConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in, out, s)
}
