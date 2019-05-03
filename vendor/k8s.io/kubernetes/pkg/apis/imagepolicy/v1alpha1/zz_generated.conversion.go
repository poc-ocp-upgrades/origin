package v1alpha1

import (
 unsafe "unsafe"
 v1alpha1 "k8s.io/api/imagepolicy/v1alpha1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 imagepolicy "k8s.io/kubernetes/pkg/apis/imagepolicy"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ImageReview)(nil), (*imagepolicy.ImageReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ImageReview_To_imagepolicy_ImageReview(a.(*v1alpha1.ImageReview), b.(*imagepolicy.ImageReview), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*imagepolicy.ImageReview)(nil), (*v1alpha1.ImageReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_imagepolicy_ImageReview_To_v1alpha1_ImageReview(a.(*imagepolicy.ImageReview), b.(*v1alpha1.ImageReview), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ImageReviewContainerSpec)(nil), (*imagepolicy.ImageReviewContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ImageReviewContainerSpec_To_imagepolicy_ImageReviewContainerSpec(a.(*v1alpha1.ImageReviewContainerSpec), b.(*imagepolicy.ImageReviewContainerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*imagepolicy.ImageReviewContainerSpec)(nil), (*v1alpha1.ImageReviewContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_imagepolicy_ImageReviewContainerSpec_To_v1alpha1_ImageReviewContainerSpec(a.(*imagepolicy.ImageReviewContainerSpec), b.(*v1alpha1.ImageReviewContainerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ImageReviewSpec)(nil), (*imagepolicy.ImageReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ImageReviewSpec_To_imagepolicy_ImageReviewSpec(a.(*v1alpha1.ImageReviewSpec), b.(*imagepolicy.ImageReviewSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*imagepolicy.ImageReviewSpec)(nil), (*v1alpha1.ImageReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_imagepolicy_ImageReviewSpec_To_v1alpha1_ImageReviewSpec(a.(*imagepolicy.ImageReviewSpec), b.(*v1alpha1.ImageReviewSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ImageReviewStatus)(nil), (*imagepolicy.ImageReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ImageReviewStatus_To_imagepolicy_ImageReviewStatus(a.(*v1alpha1.ImageReviewStatus), b.(*imagepolicy.ImageReviewStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*imagepolicy.ImageReviewStatus)(nil), (*v1alpha1.ImageReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_imagepolicy_ImageReviewStatus_To_v1alpha1_ImageReviewStatus(a.(*imagepolicy.ImageReviewStatus), b.(*v1alpha1.ImageReviewStatus), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_ImageReview_To_imagepolicy_ImageReview(in *v1alpha1.ImageReview, out *imagepolicy.ImageReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1alpha1_ImageReviewSpec_To_imagepolicy_ImageReviewSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_ImageReviewStatus_To_imagepolicy_ImageReviewStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_ImageReview_To_imagepolicy_ImageReview(in *v1alpha1.ImageReview, out *imagepolicy.ImageReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ImageReview_To_imagepolicy_ImageReview(in, out, s)
}
func autoConvert_imagepolicy_ImageReview_To_v1alpha1_ImageReview(in *imagepolicy.ImageReview, out *v1alpha1.ImageReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_imagepolicy_ImageReviewSpec_To_v1alpha1_ImageReviewSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_imagepolicy_ImageReviewStatus_To_v1alpha1_ImageReviewStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_imagepolicy_ImageReview_To_v1alpha1_ImageReview(in *imagepolicy.ImageReview, out *v1alpha1.ImageReview, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_imagepolicy_ImageReview_To_v1alpha1_ImageReview(in, out, s)
}
func autoConvert_v1alpha1_ImageReviewContainerSpec_To_imagepolicy_ImageReviewContainerSpec(in *v1alpha1.ImageReviewContainerSpec, out *imagepolicy.ImageReviewContainerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Image = in.Image
 return nil
}
func Convert_v1alpha1_ImageReviewContainerSpec_To_imagepolicy_ImageReviewContainerSpec(in *v1alpha1.ImageReviewContainerSpec, out *imagepolicy.ImageReviewContainerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ImageReviewContainerSpec_To_imagepolicy_ImageReviewContainerSpec(in, out, s)
}
func autoConvert_imagepolicy_ImageReviewContainerSpec_To_v1alpha1_ImageReviewContainerSpec(in *imagepolicy.ImageReviewContainerSpec, out *v1alpha1.ImageReviewContainerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Image = in.Image
 return nil
}
func Convert_imagepolicy_ImageReviewContainerSpec_To_v1alpha1_ImageReviewContainerSpec(in *imagepolicy.ImageReviewContainerSpec, out *v1alpha1.ImageReviewContainerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_imagepolicy_ImageReviewContainerSpec_To_v1alpha1_ImageReviewContainerSpec(in, out, s)
}
func autoConvert_v1alpha1_ImageReviewSpec_To_imagepolicy_ImageReviewSpec(in *v1alpha1.ImageReviewSpec, out *imagepolicy.ImageReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Containers = *(*[]imagepolicy.ImageReviewContainerSpec)(unsafe.Pointer(&in.Containers))
 out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
 out.Namespace = in.Namespace
 return nil
}
func Convert_v1alpha1_ImageReviewSpec_To_imagepolicy_ImageReviewSpec(in *v1alpha1.ImageReviewSpec, out *imagepolicy.ImageReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ImageReviewSpec_To_imagepolicy_ImageReviewSpec(in, out, s)
}
func autoConvert_imagepolicy_ImageReviewSpec_To_v1alpha1_ImageReviewSpec(in *imagepolicy.ImageReviewSpec, out *v1alpha1.ImageReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Containers = *(*[]v1alpha1.ImageReviewContainerSpec)(unsafe.Pointer(&in.Containers))
 out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
 out.Namespace = in.Namespace
 return nil
}
func Convert_imagepolicy_ImageReviewSpec_To_v1alpha1_ImageReviewSpec(in *imagepolicy.ImageReviewSpec, out *v1alpha1.ImageReviewSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_imagepolicy_ImageReviewSpec_To_v1alpha1_ImageReviewSpec(in, out, s)
}
func autoConvert_v1alpha1_ImageReviewStatus_To_imagepolicy_ImageReviewStatus(in *v1alpha1.ImageReviewStatus, out *imagepolicy.ImageReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Allowed = in.Allowed
 out.Reason = in.Reason
 out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
 return nil
}
func Convert_v1alpha1_ImageReviewStatus_To_imagepolicy_ImageReviewStatus(in *v1alpha1.ImageReviewStatus, out *imagepolicy.ImageReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ImageReviewStatus_To_imagepolicy_ImageReviewStatus(in, out, s)
}
func autoConvert_imagepolicy_ImageReviewStatus_To_v1alpha1_ImageReviewStatus(in *imagepolicy.ImageReviewStatus, out *v1alpha1.ImageReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Allowed = in.Allowed
 out.Reason = in.Reason
 out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
 return nil
}
func Convert_imagepolicy_ImageReviewStatus_To_v1alpha1_ImageReviewStatus(in *imagepolicy.ImageReviewStatus, out *v1alpha1.ImageReviewStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_imagepolicy_ImageReviewStatus_To_v1alpha1_ImageReviewStatus(in, out, s)
}
