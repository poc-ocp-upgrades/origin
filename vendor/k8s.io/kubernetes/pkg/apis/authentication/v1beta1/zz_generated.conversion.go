package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/authentication/v1beta1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 authentication "k8s.io/kubernetes/pkg/apis/authentication"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReview)(nil), (*authentication.TokenReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_TokenReview_To_authentication_TokenReview(a.(*v1beta1.TokenReview), b.(*authentication.TokenReview), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*authentication.TokenReview)(nil), (*v1beta1.TokenReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_authentication_TokenReview_To_v1beta1_TokenReview(a.(*authentication.TokenReview), b.(*v1beta1.TokenReview), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReviewSpec)(nil), (*authentication.TokenReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(a.(*v1beta1.TokenReviewSpec), b.(*authentication.TokenReviewSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*authentication.TokenReviewSpec)(nil), (*v1beta1.TokenReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(a.(*authentication.TokenReviewSpec), b.(*v1beta1.TokenReviewSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReviewStatus)(nil), (*authentication.TokenReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(a.(*v1beta1.TokenReviewStatus), b.(*authentication.TokenReviewStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*authentication.TokenReviewStatus)(nil), (*v1beta1.TokenReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(a.(*authentication.TokenReviewStatus), b.(*v1beta1.TokenReviewStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.UserInfo)(nil), (*authentication.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_UserInfo_To_authentication_UserInfo(a.(*v1beta1.UserInfo), b.(*authentication.UserInfo), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*authentication.UserInfo)(nil), (*v1beta1.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_authentication_UserInfo_To_v1beta1_UserInfo(a.(*authentication.UserInfo), b.(*v1beta1.UserInfo), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in *v1beta1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_TokenReview_To_authentication_TokenReview(in *v1beta1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in, out, s)
}
func autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *v1beta1.TokenReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *v1beta1.TokenReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in, out, s)
}
func autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1beta1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Token = in.Token
 out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
 return nil
}
func Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1beta1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in, out, s)
}
func autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1beta1.TokenReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Token = in.Token
 out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
 return nil
}
func Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1beta1.TokenReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in, out, s)
}
func autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1beta1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Authenticated = in.Authenticated
 if err := Convert_v1beta1_UserInfo_To_authentication_UserInfo(&in.User, &out.User, s); err != nil {
  return err
 }
 out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
 out.Error = in.Error
 return nil
}
func Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1beta1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in, out, s)
}
func autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1beta1.TokenReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Authenticated = in.Authenticated
 if err := Convert_authentication_UserInfo_To_v1beta1_UserInfo(&in.User, &out.User, s); err != nil {
  return err
 }
 out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
 out.Error = in.Error
 return nil
}
func Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1beta1.TokenReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in, out, s)
}
func autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in *v1beta1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Username = in.Username
 out.UID = in.UID
 out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
 out.Extra = *(*map[string]authentication.ExtraValue)(unsafe.Pointer(&in.Extra))
 return nil
}
func Convert_v1beta1_UserInfo_To_authentication_UserInfo(in *v1beta1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in, out, s)
}
func autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *v1beta1.UserInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Username = in.Username
 out.UID = in.UID
 out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
 out.Extra = *(*map[string]v1beta1.ExtraValue)(unsafe.Pointer(&in.Extra))
 return nil
}
func Convert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *v1beta1.UserInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
