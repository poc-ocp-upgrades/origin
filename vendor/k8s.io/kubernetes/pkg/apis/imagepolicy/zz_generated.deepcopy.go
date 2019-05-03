package imagepolicy

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ImageReview) DeepCopyInto(out *ImageReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *ImageReview) DeepCopy() *ImageReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ImageReview)
 in.DeepCopyInto(out)
 return out
}
func (in *ImageReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ImageReviewContainerSpec) DeepCopyInto(out *ImageReviewContainerSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ImageReviewContainerSpec) DeepCopy() *ImageReviewContainerSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ImageReviewContainerSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ImageReviewSpec) DeepCopyInto(out *ImageReviewSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Containers != nil {
  in, out := &in.Containers, &out.Containers
  *out = make([]ImageReviewContainerSpec, len(*in))
  copy(*out, *in)
 }
 if in.Annotations != nil {
  in, out := &in.Annotations, &out.Annotations
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *ImageReviewSpec) DeepCopy() *ImageReviewSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ImageReviewSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ImageReviewStatus) DeepCopyInto(out *ImageReviewStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.AuditAnnotations != nil {
  in, out := &in.AuditAnnotations, &out.AuditAnnotations
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *ImageReviewStatus) DeepCopy() *ImageReviewStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ImageReviewStatus)
 in.DeepCopyInto(out)
 return out
}
