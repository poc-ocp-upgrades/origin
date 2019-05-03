package certificates

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *CertificateSigningRequest) DeepCopyInto(out *CertificateSigningRequest) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *CertificateSigningRequest) DeepCopy() *CertificateSigningRequest {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CertificateSigningRequest)
 in.DeepCopyInto(out)
 return out
}
func (in *CertificateSigningRequest) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *CertificateSigningRequestCondition) DeepCopyInto(out *CertificateSigningRequestCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
 return
}
func (in *CertificateSigningRequestCondition) DeepCopy() *CertificateSigningRequestCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CertificateSigningRequestCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *CertificateSigningRequestList) DeepCopyInto(out *CertificateSigningRequestList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]CertificateSigningRequest, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *CertificateSigningRequestList) DeepCopy() *CertificateSigningRequestList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CertificateSigningRequestList)
 in.DeepCopyInto(out)
 return out
}
func (in *CertificateSigningRequestList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *CertificateSigningRequestSpec) DeepCopyInto(out *CertificateSigningRequestSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Request != nil {
  in, out := &in.Request, &out.Request
  *out = make([]byte, len(*in))
  copy(*out, *in)
 }
 if in.Usages != nil {
  in, out := &in.Usages, &out.Usages
  *out = make([]KeyUsage, len(*in))
  copy(*out, *in)
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
func (in *CertificateSigningRequestSpec) DeepCopy() *CertificateSigningRequestSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CertificateSigningRequestSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *CertificateSigningRequestStatus) DeepCopyInto(out *CertificateSigningRequestStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]CertificateSigningRequestCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Certificate != nil {
  in, out := &in.Certificate, &out.Certificate
  *out = make([]byte, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *CertificateSigningRequestStatus) DeepCopy() *CertificateSigningRequestStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CertificateSigningRequestStatus)
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
