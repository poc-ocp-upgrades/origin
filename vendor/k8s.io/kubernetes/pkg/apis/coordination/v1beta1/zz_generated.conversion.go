package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/coordination/v1beta1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 coordination "k8s.io/kubernetes/pkg/apis/coordination"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.Lease)(nil), (*coordination.Lease)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Lease_To_coordination_Lease(a.(*v1beta1.Lease), b.(*coordination.Lease), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*coordination.Lease)(nil), (*v1beta1.Lease)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_coordination_Lease_To_v1beta1_Lease(a.(*coordination.Lease), b.(*v1beta1.Lease), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.LeaseList)(nil), (*coordination.LeaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_LeaseList_To_coordination_LeaseList(a.(*v1beta1.LeaseList), b.(*coordination.LeaseList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*coordination.LeaseList)(nil), (*v1beta1.LeaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_coordination_LeaseList_To_v1beta1_LeaseList(a.(*coordination.LeaseList), b.(*v1beta1.LeaseList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.LeaseSpec)(nil), (*coordination.LeaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_LeaseSpec_To_coordination_LeaseSpec(a.(*v1beta1.LeaseSpec), b.(*coordination.LeaseSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*coordination.LeaseSpec)(nil), (*v1beta1.LeaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_coordination_LeaseSpec_To_v1beta1_LeaseSpec(a.(*coordination.LeaseSpec), b.(*v1beta1.LeaseSpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_Lease_To_coordination_Lease(in *v1beta1.Lease, out *coordination.Lease, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_LeaseSpec_To_coordination_LeaseSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_Lease_To_coordination_Lease(in *v1beta1.Lease, out *coordination.Lease, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_Lease_To_coordination_Lease(in, out, s)
}
func autoConvert_coordination_Lease_To_v1beta1_Lease(in *coordination.Lease, out *v1beta1.Lease, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_coordination_LeaseSpec_To_v1beta1_LeaseSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_coordination_Lease_To_v1beta1_Lease(in *coordination.Lease, out *v1beta1.Lease, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_coordination_Lease_To_v1beta1_Lease(in, out, s)
}
func autoConvert_v1beta1_LeaseList_To_coordination_LeaseList(in *v1beta1.LeaseList, out *coordination.LeaseList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]coordination.Lease)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_LeaseList_To_coordination_LeaseList(in *v1beta1.LeaseList, out *coordination.LeaseList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_LeaseList_To_coordination_LeaseList(in, out, s)
}
func autoConvert_coordination_LeaseList_To_v1beta1_LeaseList(in *coordination.LeaseList, out *v1beta1.LeaseList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.Lease)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_coordination_LeaseList_To_v1beta1_LeaseList(in *coordination.LeaseList, out *v1beta1.LeaseList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_coordination_LeaseList_To_v1beta1_LeaseList(in, out, s)
}
func autoConvert_v1beta1_LeaseSpec_To_coordination_LeaseSpec(in *v1beta1.LeaseSpec, out *coordination.LeaseSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HolderIdentity = (*string)(unsafe.Pointer(in.HolderIdentity))
 out.LeaseDurationSeconds = (*int32)(unsafe.Pointer(in.LeaseDurationSeconds))
 out.AcquireTime = (*v1.MicroTime)(unsafe.Pointer(in.AcquireTime))
 out.RenewTime = (*v1.MicroTime)(unsafe.Pointer(in.RenewTime))
 out.LeaseTransitions = (*int32)(unsafe.Pointer(in.LeaseTransitions))
 return nil
}
func Convert_v1beta1_LeaseSpec_To_coordination_LeaseSpec(in *v1beta1.LeaseSpec, out *coordination.LeaseSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_LeaseSpec_To_coordination_LeaseSpec(in, out, s)
}
func autoConvert_coordination_LeaseSpec_To_v1beta1_LeaseSpec(in *coordination.LeaseSpec, out *v1beta1.LeaseSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HolderIdentity = (*string)(unsafe.Pointer(in.HolderIdentity))
 out.LeaseDurationSeconds = (*int32)(unsafe.Pointer(in.LeaseDurationSeconds))
 out.AcquireTime = (*v1.MicroTime)(unsafe.Pointer(in.AcquireTime))
 out.RenewTime = (*v1.MicroTime)(unsafe.Pointer(in.RenewTime))
 out.LeaseTransitions = (*int32)(unsafe.Pointer(in.LeaseTransitions))
 return nil
}
func Convert_coordination_LeaseSpec_To_v1beta1_LeaseSpec(in *coordination.LeaseSpec, out *v1beta1.LeaseSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_coordination_LeaseSpec_To_v1beta1_LeaseSpec(in, out, s)
}
