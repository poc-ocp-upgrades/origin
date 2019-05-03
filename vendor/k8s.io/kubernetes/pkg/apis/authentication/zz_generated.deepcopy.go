package authentication

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *BoundObjectReference) DeepCopyInto(out *BoundObjectReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *BoundObjectReference) DeepCopy() *BoundObjectReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(BoundObjectReference)
 in.DeepCopyInto(out)
 return out
}
func (in ExtraValue) DeepCopyInto(out *ExtraValue) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(ExtraValue, len(*in))
  copy(*out, *in)
  return
 }
}
func (in ExtraValue) DeepCopy() ExtraValue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtraValue)
 in.DeepCopyInto(out)
 return *out
}
func (in *TokenRequest) DeepCopyInto(out *TokenRequest) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *TokenRequest) DeepCopy() *TokenRequest {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenRequest)
 in.DeepCopyInto(out)
 return out
}
func (in *TokenRequest) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *TokenRequestSpec) DeepCopyInto(out *TokenRequestSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Audiences != nil {
  in, out := &in.Audiences, &out.Audiences
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.BoundObjectRef != nil {
  in, out := &in.BoundObjectRef, &out.BoundObjectRef
  *out = new(BoundObjectReference)
  **out = **in
 }
 return
}
func (in *TokenRequestSpec) DeepCopy() *TokenRequestSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenRequestSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *TokenRequestStatus) DeepCopyInto(out *TokenRequestStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.ExpirationTimestamp.DeepCopyInto(&out.ExpirationTimestamp)
 return
}
func (in *TokenRequestStatus) DeepCopy() *TokenRequestStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenRequestStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *TokenReview) DeepCopyInto(out *TokenReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *TokenReview) DeepCopy() *TokenReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenReview)
 in.DeepCopyInto(out)
 return out
}
func (in *TokenReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *TokenReviewSpec) DeepCopyInto(out *TokenReviewSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Audiences != nil {
  in, out := &in.Audiences, &out.Audiences
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *TokenReviewSpec) DeepCopy() *TokenReviewSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenReviewSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *TokenReviewStatus) DeepCopyInto(out *TokenReviewStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.User.DeepCopyInto(&out.User)
 if in.Audiences != nil {
  in, out := &in.Audiences, &out.Audiences
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *TokenReviewStatus) DeepCopy() *TokenReviewStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TokenReviewStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Groups != nil {
  in, out := &in.Groups, &out.Groups
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Extra != nil {
  in, out := &in.Extra, &out.Extra
  *out = make(map[string]ExtraValue, len(*in))
  for key, val := range *in {
   var outVal []string
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = make(ExtraValue, len(*in))
    copy(*out, *in)
   }
   (*out)[key] = outVal
  }
 }
 return
}
func (in *UserInfo) DeepCopy() *UserInfo {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(UserInfo)
 in.DeepCopyInto(out)
 return out
}
