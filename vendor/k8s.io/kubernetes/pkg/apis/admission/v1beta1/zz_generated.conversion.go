package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/admission/v1beta1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 types "k8s.io/apimachinery/pkg/types"
 admission "k8s.io/kubernetes/pkg/apis/admission"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.AdmissionRequest)(nil), (*admission.AdmissionRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_AdmissionRequest_To_admission_AdmissionRequest(a.(*v1beta1.AdmissionRequest), b.(*admission.AdmissionRequest), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admission.AdmissionRequest)(nil), (*v1beta1.AdmissionRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admission_AdmissionRequest_To_v1beta1_AdmissionRequest(a.(*admission.AdmissionRequest), b.(*v1beta1.AdmissionRequest), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.AdmissionResponse)(nil), (*admission.AdmissionResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_AdmissionResponse_To_admission_AdmissionResponse(a.(*v1beta1.AdmissionResponse), b.(*admission.AdmissionResponse), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admission.AdmissionResponse)(nil), (*v1beta1.AdmissionResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admission_AdmissionResponse_To_v1beta1_AdmissionResponse(a.(*admission.AdmissionResponse), b.(*v1beta1.AdmissionResponse), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.AdmissionReview)(nil), (*admission.AdmissionReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_AdmissionReview_To_admission_AdmissionReview(a.(*v1beta1.AdmissionReview), b.(*admission.AdmissionReview), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*admission.AdmissionReview)(nil), (*v1beta1.AdmissionReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_admission_AdmissionReview_To_v1beta1_AdmissionReview(a.(*admission.AdmissionReview), b.(*v1beta1.AdmissionReview), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_AdmissionRequest_To_admission_AdmissionRequest(in *v1beta1.AdmissionRequest, out *admission.AdmissionRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = types.UID(in.UID)
 out.Kind = in.Kind
 out.Resource = in.Resource
 out.SubResource = in.SubResource
 out.Name = in.Name
 out.Namespace = in.Namespace
 out.Operation = admission.Operation(in.Operation)
 if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
  return err
 }
 if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Object, &out.Object, s); err != nil {
  return err
 }
 if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.OldObject, &out.OldObject, s); err != nil {
  return err
 }
 out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
 return nil
}
func Convert_v1beta1_AdmissionRequest_To_admission_AdmissionRequest(in *v1beta1.AdmissionRequest, out *admission.AdmissionRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_AdmissionRequest_To_admission_AdmissionRequest(in, out, s)
}
func autoConvert_admission_AdmissionRequest_To_v1beta1_AdmissionRequest(in *admission.AdmissionRequest, out *v1beta1.AdmissionRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = types.UID(in.UID)
 out.Kind = in.Kind
 out.Resource = in.Resource
 out.SubResource = in.SubResource
 out.Name = in.Name
 out.Namespace = in.Namespace
 out.Operation = v1beta1.Operation(in.Operation)
 if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
  return err
 }
 if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Object, &out.Object, s); err != nil {
  return err
 }
 if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.OldObject, &out.OldObject, s); err != nil {
  return err
 }
 out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
 return nil
}
func Convert_admission_AdmissionRequest_To_v1beta1_AdmissionRequest(in *admission.AdmissionRequest, out *v1beta1.AdmissionRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admission_AdmissionRequest_To_v1beta1_AdmissionRequest(in, out, s)
}
func autoConvert_v1beta1_AdmissionResponse_To_admission_AdmissionResponse(in *v1beta1.AdmissionResponse, out *admission.AdmissionResponse, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = types.UID(in.UID)
 out.Allowed = in.Allowed
 out.Result = (*v1.Status)(unsafe.Pointer(in.Result))
 out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
 out.PatchType = (*admission.PatchType)(unsafe.Pointer(in.PatchType))
 out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
 return nil
}
func Convert_v1beta1_AdmissionResponse_To_admission_AdmissionResponse(in *v1beta1.AdmissionResponse, out *admission.AdmissionResponse, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_AdmissionResponse_To_admission_AdmissionResponse(in, out, s)
}
func autoConvert_admission_AdmissionResponse_To_v1beta1_AdmissionResponse(in *admission.AdmissionResponse, out *v1beta1.AdmissionResponse, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = types.UID(in.UID)
 out.Allowed = in.Allowed
 out.Result = (*v1.Status)(unsafe.Pointer(in.Result))
 out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
 out.PatchType = (*v1beta1.PatchType)(unsafe.Pointer(in.PatchType))
 out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
 return nil
}
func Convert_admission_AdmissionResponse_To_v1beta1_AdmissionResponse(in *admission.AdmissionResponse, out *v1beta1.AdmissionResponse, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admission_AdmissionResponse_To_v1beta1_AdmissionResponse(in, out, s)
}
func autoConvert_v1beta1_AdmissionReview_To_admission_AdmissionReview(in *v1beta1.AdmissionReview, out *admission.AdmissionReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Request != nil {
  in, out := &in.Request, &out.Request
  *out = new(admission.AdmissionRequest)
  if err := Convert_v1beta1_AdmissionRequest_To_admission_AdmissionRequest(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Request = nil
 }
 out.Response = (*admission.AdmissionResponse)(unsafe.Pointer(in.Response))
 return nil
}
func Convert_v1beta1_AdmissionReview_To_admission_AdmissionReview(in *v1beta1.AdmissionReview, out *admission.AdmissionReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_AdmissionReview_To_admission_AdmissionReview(in, out, s)
}
func autoConvert_admission_AdmissionReview_To_v1beta1_AdmissionReview(in *admission.AdmissionReview, out *v1beta1.AdmissionReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Request != nil {
  in, out := &in.Request, &out.Request
  *out = new(v1beta1.AdmissionRequest)
  if err := Convert_admission_AdmissionRequest_To_v1beta1_AdmissionRequest(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Request = nil
 }
 out.Response = (*v1beta1.AdmissionResponse)(unsafe.Pointer(in.Response))
 return nil
}
func Convert_admission_AdmissionReview_To_v1beta1_AdmissionReview(in *admission.AdmissionReview, out *v1beta1.AdmissionReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_admission_AdmissionReview_To_v1beta1_AdmissionReview(in, out, s)
}
