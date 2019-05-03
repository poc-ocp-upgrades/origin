package coordination

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Lease) DeepCopyInto(out *Lease) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *Lease) DeepCopy() *Lease {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Lease)
 in.DeepCopyInto(out)
 return out
}
func (in *Lease) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *LeaseList) DeepCopyInto(out *LeaseList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Lease, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *LeaseList) DeepCopy() *LeaseList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LeaseList)
 in.DeepCopyInto(out)
 return out
}
func (in *LeaseList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *LeaseSpec) DeepCopyInto(out *LeaseSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.HolderIdentity != nil {
  in, out := &in.HolderIdentity, &out.HolderIdentity
  *out = new(string)
  **out = **in
 }
 if in.LeaseDurationSeconds != nil {
  in, out := &in.LeaseDurationSeconds, &out.LeaseDurationSeconds
  *out = new(int32)
  **out = **in
 }
 if in.AcquireTime != nil {
  in, out := &in.AcquireTime, &out.AcquireTime
  *out = (*in).DeepCopy()
 }
 if in.RenewTime != nil {
  in, out := &in.RenewTime, &out.RenewTime
  *out = (*in).DeepCopy()
 }
 if in.LeaseTransitions != nil {
  in, out := &in.LeaseTransitions, &out.LeaseTransitions
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *LeaseSpec) DeepCopy() *LeaseSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LeaseSpec)
 in.DeepCopyInto(out)
 return out
}
