package v0

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Policy) DeepCopyInto(out *Policy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
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
