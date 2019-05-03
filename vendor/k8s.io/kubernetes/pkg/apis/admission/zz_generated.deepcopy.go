package admission

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AdmissionRequest) DeepCopyInto(out *AdmissionRequest) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.Kind = in.Kind
 out.Resource = in.Resource
 in.UserInfo.DeepCopyInto(&out.UserInfo)
 if in.Object != nil {
  out.Object = in.Object.DeepCopyObject()
 }
 if in.OldObject != nil {
  out.OldObject = in.OldObject.DeepCopyObject()
 }
 if in.DryRun != nil {
  in, out := &in.DryRun, &out.DryRun
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *AdmissionRequest) DeepCopy() *AdmissionRequest {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AdmissionRequest)
 in.DeepCopyInto(out)
 return out
}
func (in *AdmissionResponse) DeepCopyInto(out *AdmissionResponse) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Result != nil {
  in, out := &in.Result, &out.Result
  *out = new(v1.Status)
  (*in).DeepCopyInto(*out)
 }
 if in.Patch != nil {
  in, out := &in.Patch, &out.Patch
  *out = make([]byte, len(*in))
  copy(*out, *in)
 }
 if in.PatchType != nil {
  in, out := &in.PatchType, &out.PatchType
  *out = new(PatchType)
  **out = **in
 }
 if in.AuditAnnotations != nil {
  in, out := &in.AuditAnnotations, &out.AuditAnnotations
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *AdmissionResponse) DeepCopy() *AdmissionResponse {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AdmissionResponse)
 in.DeepCopyInto(out)
 return out
}
func (in *AdmissionReview) DeepCopyInto(out *AdmissionReview) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.Request != nil {
  in, out := &in.Request, &out.Request
  *out = new(AdmissionRequest)
  (*in).DeepCopyInto(*out)
 }
 if in.Response != nil {
  in, out := &in.Response, &out.Response
  *out = new(AdmissionResponse)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *AdmissionReview) DeepCopy() *AdmissionReview {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AdmissionReview)
 in.DeepCopyInto(out)
 return out
}
func (in *AdmissionReview) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
