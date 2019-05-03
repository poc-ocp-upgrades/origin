package v1beta1

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Policy) DeepCopyInto(out *Policy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.Spec = in.Spec
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
func (in *Policy) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PolicySpec) DeepCopy() *PolicySpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PolicySpec)
 in.DeepCopyInto(out)
 return out
}
