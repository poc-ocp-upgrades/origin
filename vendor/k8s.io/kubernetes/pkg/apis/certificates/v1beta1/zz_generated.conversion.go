package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/certificates/v1beta1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 certificates "k8s.io/kubernetes/pkg/apis/certificates"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.CertificateSigningRequest)(nil), (*certificates.CertificateSigningRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(a.(*v1beta1.CertificateSigningRequest), b.(*certificates.CertificateSigningRequest), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*certificates.CertificateSigningRequest)(nil), (*v1beta1.CertificateSigningRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(a.(*certificates.CertificateSigningRequest), b.(*v1beta1.CertificateSigningRequest), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CertificateSigningRequestCondition)(nil), (*certificates.CertificateSigningRequestCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(a.(*v1beta1.CertificateSigningRequestCondition), b.(*certificates.CertificateSigningRequestCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*certificates.CertificateSigningRequestCondition)(nil), (*v1beta1.CertificateSigningRequestCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(a.(*certificates.CertificateSigningRequestCondition), b.(*v1beta1.CertificateSigningRequestCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CertificateSigningRequestList)(nil), (*certificates.CertificateSigningRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(a.(*v1beta1.CertificateSigningRequestList), b.(*certificates.CertificateSigningRequestList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*certificates.CertificateSigningRequestList)(nil), (*v1beta1.CertificateSigningRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(a.(*certificates.CertificateSigningRequestList), b.(*v1beta1.CertificateSigningRequestList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CertificateSigningRequestSpec)(nil), (*certificates.CertificateSigningRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(a.(*v1beta1.CertificateSigningRequestSpec), b.(*certificates.CertificateSigningRequestSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*certificates.CertificateSigningRequestSpec)(nil), (*v1beta1.CertificateSigningRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(a.(*certificates.CertificateSigningRequestSpec), b.(*v1beta1.CertificateSigningRequestSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CertificateSigningRequestStatus)(nil), (*certificates.CertificateSigningRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(a.(*v1beta1.CertificateSigningRequestStatus), b.(*certificates.CertificateSigningRequestStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*certificates.CertificateSigningRequestStatus)(nil), (*v1beta1.CertificateSigningRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(a.(*certificates.CertificateSigningRequestStatus), b.(*v1beta1.CertificateSigningRequestStatus), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *v1beta1.CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *v1beta1.CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in, out, s)
}
func autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *v1beta1.CertificateSigningRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *v1beta1.CertificateSigningRequest, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in, out, s)
}
func autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *v1beta1.CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = certificates.RequestConditionType(in.Type)
 out.Reason = in.Reason
 out.Message = in.Message
 out.LastUpdateTime = in.LastUpdateTime
 return nil
}
func Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *v1beta1.CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in, out, s)
}
func autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *v1beta1.CertificateSigningRequestCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1beta1.RequestConditionType(in.Type)
 out.Reason = in.Reason
 out.Message = in.Message
 out.LastUpdateTime = in.LastUpdateTime
 return nil
}
func Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *v1beta1.CertificateSigningRequestCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in, out, s)
}
func autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *v1beta1.CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]certificates.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *v1beta1.CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in, out, s)
}
func autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *v1beta1.CertificateSigningRequestList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *v1beta1.CertificateSigningRequestList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in, out, s)
}
func autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *v1beta1.CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
 out.Usages = *(*[]certificates.KeyUsage)(unsafe.Pointer(&in.Usages))
 out.Username = in.Username
 out.UID = in.UID
 out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
 out.Extra = *(*map[string]certificates.ExtraValue)(unsafe.Pointer(&in.Extra))
 return nil
}
func Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *v1beta1.CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in, out, s)
}
func autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *v1beta1.CertificateSigningRequestSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
 out.Usages = *(*[]v1beta1.KeyUsage)(unsafe.Pointer(&in.Usages))
 out.Username = in.Username
 out.UID = in.UID
 out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
 out.Extra = *(*map[string]v1beta1.ExtraValue)(unsafe.Pointer(&in.Extra))
 return nil
}
func Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *v1beta1.CertificateSigningRequestSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in, out, s)
}
func autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *v1beta1.CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Conditions = *(*[]certificates.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
 out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
 return nil
}
func Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *v1beta1.CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in, out, s)
}
func autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *v1beta1.CertificateSigningRequestStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Conditions = *(*[]v1beta1.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
 out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
 return nil
}
func Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *v1beta1.CertificateSigningRequestStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in, out, s)
}
