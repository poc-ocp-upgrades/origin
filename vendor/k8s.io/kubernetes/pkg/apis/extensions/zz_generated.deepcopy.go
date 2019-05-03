package extensions

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *HTTPIngressPath) DeepCopyInto(out *HTTPIngressPath) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.Backend = in.Backend
 return
}
func (in *HTTPIngressPath) DeepCopy() *HTTPIngressPath {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HTTPIngressPath)
 in.DeepCopyInto(out)
 return out
}
func (in *HTTPIngressRuleValue) DeepCopyInto(out *HTTPIngressRuleValue) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Paths != nil {
  in, out := &in.Paths, &out.Paths
  *out = make([]HTTPIngressPath, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *HTTPIngressRuleValue) DeepCopy() *HTTPIngressRuleValue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HTTPIngressRuleValue)
 in.DeepCopyInto(out)
 return out
}
func (in *Ingress) DeepCopyInto(out *Ingress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Ingress) DeepCopy() *Ingress {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Ingress)
 in.DeepCopyInto(out)
 return out
}
func (in *Ingress) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *IngressBackend) DeepCopyInto(out *IngressBackend) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.ServicePort = in.ServicePort
 return
}
func (in *IngressBackend) DeepCopy() *IngressBackend {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressBackend)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressList) DeepCopyInto(out *IngressList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Ingress, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *IngressList) DeepCopy() *IngressList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressList)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.IngressRuleValue.DeepCopyInto(&out.IngressRuleValue)
 return
}
func (in *IngressRule) DeepCopy() *IngressRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressRule)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressRuleValue) DeepCopyInto(out *IngressRuleValue) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.HTTP != nil {
  in, out := &in.HTTP, &out.HTTP
  *out = new(HTTPIngressRuleValue)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *IngressRuleValue) DeepCopy() *IngressRuleValue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressRuleValue)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Backend != nil {
  in, out := &in.Backend, &out.Backend
  *out = new(IngressBackend)
  **out = **in
 }
 if in.TLS != nil {
  in, out := &in.TLS, &out.TLS
  *out = make([]IngressTLS, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Rules != nil {
  in, out := &in.Rules, &out.Rules
  *out = make([]IngressRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *IngressSpec) DeepCopy() *IngressSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
 return
}
func (in *IngressStatus) DeepCopy() *IngressStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *IngressTLS) DeepCopyInto(out *IngressTLS) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Hosts != nil {
  in, out := &in.Hosts, &out.Hosts
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *IngressTLS) DeepCopy() *IngressTLS {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IngressTLS)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerDummy) DeepCopyInto(out *ReplicationControllerDummy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 return
}
func (in *ReplicationControllerDummy) DeepCopy() *ReplicationControllerDummy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerDummy)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerDummy) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
