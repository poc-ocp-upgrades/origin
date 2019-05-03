package auditregistration

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AuditSink) DeepCopyInto(out *AuditSink) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *AuditSink) DeepCopy() *AuditSink {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AuditSink)
 in.DeepCopyInto(out)
 return out
}
func (in *AuditSink) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *AuditSinkList) DeepCopyInto(out *AuditSinkList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]AuditSink, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *AuditSinkList) DeepCopy() *AuditSinkList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AuditSinkList)
 in.DeepCopyInto(out)
 return out
}
func (in *AuditSinkList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *AuditSinkSpec) DeepCopyInto(out *AuditSinkSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Policy.DeepCopyInto(&out.Policy)
 in.Webhook.DeepCopyInto(&out.Webhook)
 return
}
func (in *AuditSinkSpec) DeepCopy() *AuditSinkSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AuditSinkSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *Policy) DeepCopyInto(out *Policy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Stages != nil {
  in, out := &in.Stages, &out.Stages
  *out = make([]Stage, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *Policy) DeepCopy() *Policy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Policy)
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
func (in *Webhook) DeepCopyInto(out *Webhook) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Throttle != nil {
  in, out := &in.Throttle, &out.Throttle
  *out = new(WebhookThrottleConfig)
  (*in).DeepCopyInto(*out)
 }
 in.ClientConfig.DeepCopyInto(&out.ClientConfig)
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
func (in *WebhookThrottleConfig) DeepCopyInto(out *WebhookThrottleConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.QPS != nil {
  in, out := &in.QPS, &out.QPS
  *out = new(int64)
  **out = **in
 }
 if in.Burst != nil {
  in, out := &in.Burst, &out.Burst
  *out = new(int64)
  **out = **in
 }
 return
}
func (in *WebhookThrottleConfig) DeepCopy() *WebhookThrottleConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(WebhookThrottleConfig)
 in.DeepCopyInto(out)
 return out
}
