package admissionregistration

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Initializer) DeepCopyInto(out *Initializer) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Rules != nil {
  in, out := &in.Rules, &out.Rules
  *out = make([]Rule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *Initializer) DeepCopy() *Initializer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Initializer)
 in.DeepCopyInto(out)
 return out
}
func (in *InitializerConfiguration) DeepCopyInto(out *InitializerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Initializers != nil {
  in, out := &in.Initializers, &out.Initializers
  *out = make([]Initializer, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *InitializerConfiguration) DeepCopy() *InitializerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(InitializerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *InitializerConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *InitializerConfigurationList) DeepCopyInto(out *InitializerConfigurationList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]InitializerConfiguration, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *InitializerConfigurationList) DeepCopy() *InitializerConfigurationList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(InitializerConfigurationList)
 in.DeepCopyInto(out)
 return out
}
func (in *InitializerConfigurationList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *MutatingWebhookConfiguration) DeepCopyInto(out *MutatingWebhookConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Webhooks != nil {
  in, out := &in.Webhooks, &out.Webhooks
  *out = make([]Webhook, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *MutatingWebhookConfiguration) DeepCopy() *MutatingWebhookConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MutatingWebhookConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *MutatingWebhookConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *MutatingWebhookConfigurationList) DeepCopyInto(out *MutatingWebhookConfigurationList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]MutatingWebhookConfiguration, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *MutatingWebhookConfigurationList) DeepCopy() *MutatingWebhookConfigurationList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MutatingWebhookConfigurationList)
 in.DeepCopyInto(out)
 return out
}
func (in *MutatingWebhookConfigurationList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *Rule) DeepCopyInto(out *Rule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.APIGroups != nil {
  in, out := &in.APIGroups, &out.APIGroups
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.APIVersions != nil {
  in, out := &in.APIVersions, &out.APIVersions
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Resources != nil {
  in, out := &in.Resources, &out.Resources
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *Rule) DeepCopy() *Rule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Rule)
 in.DeepCopyInto(out)
 return out
}
func (in *RuleWithOperations) DeepCopyInto(out *RuleWithOperations) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Operations != nil {
  in, out := &in.Operations, &out.Operations
  *out = make([]OperationType, len(*in))
  copy(*out, *in)
 }
 in.Rule.DeepCopyInto(&out.Rule)
 return
}
func (in *RuleWithOperations) DeepCopy() *RuleWithOperations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RuleWithOperations)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Path != nil {
  in, out := &in.Path, &out.Path
  *out = new(string)
  **out = **in
 }
 return
}
func (in *ServiceReference) DeepCopy() *ServiceReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceReference)
 in.DeepCopyInto(out)
 return out
}
func (in *ValidatingWebhookConfiguration) DeepCopyInto(out *ValidatingWebhookConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Webhooks != nil {
  in, out := &in.Webhooks, &out.Webhooks
  *out = make([]Webhook, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ValidatingWebhookConfiguration) DeepCopy() *ValidatingWebhookConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ValidatingWebhookConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *ValidatingWebhookConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ValidatingWebhookConfigurationList) DeepCopyInto(out *ValidatingWebhookConfigurationList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ValidatingWebhookConfiguration, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ValidatingWebhookConfigurationList) DeepCopy() *ValidatingWebhookConfigurationList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ValidatingWebhookConfigurationList)
 in.DeepCopyInto(out)
 return out
}
func (in *ValidatingWebhookConfigurationList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *Webhook) DeepCopyInto(out *Webhook) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.ClientConfig.DeepCopyInto(&out.ClientConfig)
 if in.Rules != nil {
  in, out := &in.Rules, &out.Rules
  *out = make([]RuleWithOperations, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.FailurePolicy != nil {
  in, out := &in.FailurePolicy, &out.FailurePolicy
  *out = new(FailurePolicyType)
  **out = **in
 }
 if in.NamespaceSelector != nil {
  in, out := &in.NamespaceSelector, &out.NamespaceSelector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.SideEffects != nil {
  in, out := &in.SideEffects, &out.SideEffects
  *out = new(SideEffectClass)
  **out = **in
 }
 return
}
func (in *Webhook) DeepCopy() *Webhook {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Webhook)
 in.DeepCopyInto(out)
 return out
}
func (in *WebhookClientConfig) DeepCopyInto(out *WebhookClientConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.URL != nil {
  in, out := &in.URL, &out.URL
  *out = new(string)
  **out = **in
 }
 if in.Service != nil {
  in, out := &in.Service, &out.Service
  *out = new(ServiceReference)
  (*in).DeepCopyInto(*out)
 }
 if in.CABundle != nil {
  in, out := &in.CABundle, &out.CABundle
  *out = make([]byte, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *WebhookClientConfig) DeepCopy() *WebhookClientConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(WebhookClientConfig)
 in.DeepCopyInto(out)
 return out
}
