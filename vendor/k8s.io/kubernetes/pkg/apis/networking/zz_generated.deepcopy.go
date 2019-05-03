package networking

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 intstr "k8s.io/apimachinery/pkg/util/intstr"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *IPBlock) DeepCopyInto(out *IPBlock) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Except != nil {
  in, out := &in.Except, &out.Except
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *IPBlock) DeepCopy() *IPBlock {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IPBlock)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicy)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicy) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NetworkPolicyEgressRule) DeepCopyInto(out *NetworkPolicyEgressRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]NetworkPolicyPort, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.To != nil {
  in, out := &in.To, &out.To
  *out = make([]NetworkPolicyPeer, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NetworkPolicyEgressRule) DeepCopy() *NetworkPolicyEgressRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicyEgressRule)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicyIngressRule) DeepCopyInto(out *NetworkPolicyIngressRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]NetworkPolicyPort, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.From != nil {
  in, out := &in.From, &out.From
  *out = make([]NetworkPolicyPeer, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NetworkPolicyIngressRule) DeepCopy() *NetworkPolicyIngressRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicyIngressRule)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicyList) DeepCopyInto(out *NetworkPolicyList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]NetworkPolicy, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NetworkPolicyList) DeepCopy() *NetworkPolicyList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicyList)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicyList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NetworkPolicyPeer) DeepCopyInto(out *NetworkPolicyPeer) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.PodSelector != nil {
  in, out := &in.PodSelector, &out.PodSelector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.NamespaceSelector != nil {
  in, out := &in.NamespaceSelector, &out.NamespaceSelector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.IPBlock != nil {
  in, out := &in.IPBlock, &out.IPBlock
  *out = new(IPBlock)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *NetworkPolicyPeer) DeepCopy() *NetworkPolicyPeer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicyPeer)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicyPort) DeepCopyInto(out *NetworkPolicyPort) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Protocol != nil {
  in, out := &in.Protocol, &out.Protocol
  *out = new(core.Protocol)
  **out = **in
 }
 if in.Port != nil {
  in, out := &in.Port, &out.Port
  *out = new(intstr.IntOrString)
  **out = **in
 }
 return
}
func (in *NetworkPolicyPort) DeepCopy() *NetworkPolicyPort {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicyPort)
 in.DeepCopyInto(out)
 return out
}
func (in *NetworkPolicySpec) DeepCopyInto(out *NetworkPolicySpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.PodSelector.DeepCopyInto(&out.PodSelector)
 if in.Ingress != nil {
  in, out := &in.Ingress, &out.Ingress
  *out = make([]NetworkPolicyIngressRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Egress != nil {
  in, out := &in.Egress, &out.Egress
  *out = make([]NetworkPolicyEgressRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.PolicyTypes != nil {
  in, out := &in.PolicyTypes, &out.PolicyTypes
  *out = make([]PolicyType, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *NetworkPolicySpec) DeepCopy() *NetworkPolicySpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NetworkPolicySpec)
 in.DeepCopyInto(out)
 return out
}
