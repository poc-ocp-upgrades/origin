package scheduling

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *PriorityClass) DeepCopyInto(out *PriorityClass) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 return
}
func (in *PriorityClass) DeepCopy() *PriorityClass {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PriorityClass)
 in.DeepCopyInto(out)
 return out
}
func (in *PriorityClass) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PriorityClassList) DeepCopyInto(out *PriorityClassList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PriorityClass, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PriorityClassList) DeepCopy() *PriorityClassList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PriorityClassList)
 in.DeepCopyInto(out)
 return out
}
func (in *PriorityClassList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
