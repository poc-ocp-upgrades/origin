package authorization

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

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
func (in *LocalSubjectAccessReview) DeepCopyInto(out *LocalSubjectAccessReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 out.Status = in.Status
 return
}
func (in *LocalSubjectAccessReview) DeepCopy() *LocalSubjectAccessReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LocalSubjectAccessReview)
 in.DeepCopyInto(out)
 return out
}
func (in *LocalSubjectAccessReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NonResourceAttributes) DeepCopyInto(out *NonResourceAttributes) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NonResourceAttributes) DeepCopy() *NonResourceAttributes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NonResourceAttributes)
 in.DeepCopyInto(out)
 return out
}
func (in *NonResourceRule) DeepCopyInto(out *NonResourceRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Verbs != nil {
  in, out := &in.Verbs, &out.Verbs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.NonResourceURLs != nil {
  in, out := &in.NonResourceURLs, &out.NonResourceURLs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *NonResourceRule) DeepCopy() *NonResourceRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NonResourceRule)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceAttributes) DeepCopyInto(out *ResourceAttributes) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ResourceAttributes) DeepCopy() *ResourceAttributes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceAttributes)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceRule) DeepCopyInto(out *ResourceRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Verbs != nil {
  in, out := &in.Verbs, &out.Verbs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.APIGroups != nil {
  in, out := &in.APIGroups, &out.APIGroups
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Resources != nil {
  in, out := &in.Resources, &out.Resources
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.ResourceNames != nil {
  in, out := &in.ResourceNames, &out.ResourceNames
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ResourceRule) DeepCopy() *ResourceRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceRule)
 in.DeepCopyInto(out)
 return out
}
func (in *SelfSubjectAccessReview) DeepCopyInto(out *SelfSubjectAccessReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 out.Status = in.Status
 return
}
func (in *SelfSubjectAccessReview) DeepCopy() *SelfSubjectAccessReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SelfSubjectAccessReview)
 in.DeepCopyInto(out)
 return out
}
func (in *SelfSubjectAccessReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *SelfSubjectAccessReviewSpec) DeepCopyInto(out *SelfSubjectAccessReviewSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ResourceAttributes != nil {
  in, out := &in.ResourceAttributes, &out.ResourceAttributes
  *out = new(ResourceAttributes)
  **out = **in
 }
 if in.NonResourceAttributes != nil {
  in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
  *out = new(NonResourceAttributes)
  **out = **in
 }
 return
}
func (in *SelfSubjectAccessReviewSpec) DeepCopy() *SelfSubjectAccessReviewSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SelfSubjectAccessReviewSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *SelfSubjectRulesReview) DeepCopyInto(out *SelfSubjectRulesReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 out.Spec = in.Spec
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *SelfSubjectRulesReview) DeepCopy() *SelfSubjectRulesReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SelfSubjectRulesReview)
 in.DeepCopyInto(out)
 return out
}
func (in *SelfSubjectRulesReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *SelfSubjectRulesReviewSpec) DeepCopyInto(out *SelfSubjectRulesReviewSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SelfSubjectRulesReviewSpec) DeepCopy() *SelfSubjectRulesReviewSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SelfSubjectRulesReviewSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *SubjectAccessReview) DeepCopyInto(out *SubjectAccessReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 out.Status = in.Status
 return
}
func (in *SubjectAccessReview) DeepCopy() *SubjectAccessReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SubjectAccessReview)
 in.DeepCopyInto(out)
 return out
}
func (in *SubjectAccessReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *SubjectAccessReviewSpec) DeepCopyInto(out *SubjectAccessReviewSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ResourceAttributes != nil {
  in, out := &in.ResourceAttributes, &out.ResourceAttributes
  *out = new(ResourceAttributes)
  **out = **in
 }
 if in.NonResourceAttributes != nil {
  in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
  *out = new(NonResourceAttributes)
  **out = **in
 }
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
func (in *SubjectAccessReviewSpec) DeepCopy() *SubjectAccessReviewSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SubjectAccessReviewSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *SubjectAccessReviewStatus) DeepCopyInto(out *SubjectAccessReviewStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SubjectAccessReviewStatus) DeepCopy() *SubjectAccessReviewStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SubjectAccessReviewStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *SubjectRulesReviewStatus) DeepCopyInto(out *SubjectRulesReviewStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ResourceRules != nil {
  in, out := &in.ResourceRules, &out.ResourceRules
  *out = make([]ResourceRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.NonResourceRules != nil {
  in, out := &in.NonResourceRules, &out.NonResourceRules
  *out = make([]NonResourceRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *SubjectRulesReviewStatus) DeepCopy() *SubjectRulesReviewStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SubjectRulesReviewStatus)
 in.DeepCopyInto(out)
 return out
}
