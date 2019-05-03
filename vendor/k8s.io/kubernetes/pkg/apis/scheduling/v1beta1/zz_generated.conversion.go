package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/scheduling/v1beta1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.PriorityClass)(nil), (*scheduling.PriorityClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PriorityClass_To_scheduling_PriorityClass(a.(*v1beta1.PriorityClass), b.(*scheduling.PriorityClass), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*scheduling.PriorityClass)(nil), (*v1beta1.PriorityClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_scheduling_PriorityClass_To_v1beta1_PriorityClass(a.(*scheduling.PriorityClass), b.(*v1beta1.PriorityClass), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.PriorityClassList)(nil), (*scheduling.PriorityClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_PriorityClassList_To_scheduling_PriorityClassList(a.(*v1beta1.PriorityClassList), b.(*scheduling.PriorityClassList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*scheduling.PriorityClassList)(nil), (*v1beta1.PriorityClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_scheduling_PriorityClassList_To_v1beta1_PriorityClassList(a.(*scheduling.PriorityClassList), b.(*v1beta1.PriorityClassList), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_PriorityClass_To_scheduling_PriorityClass(in *v1beta1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Value = in.Value
 out.GlobalDefault = in.GlobalDefault
 out.Description = in.Description
 return nil
}
func Convert_v1beta1_PriorityClass_To_scheduling_PriorityClass(in *v1beta1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PriorityClass_To_scheduling_PriorityClass(in, out, s)
}
func autoConvert_scheduling_PriorityClass_To_v1beta1_PriorityClass(in *scheduling.PriorityClass, out *v1beta1.PriorityClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Value = in.Value
 out.GlobalDefault = in.GlobalDefault
 out.Description = in.Description
 return nil
}
func Convert_scheduling_PriorityClass_To_v1beta1_PriorityClass(in *scheduling.PriorityClass, out *v1beta1.PriorityClass, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_scheduling_PriorityClass_To_v1beta1_PriorityClass(in, out, s)
}
func autoConvert_v1beta1_PriorityClassList_To_scheduling_PriorityClassList(in *v1beta1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]scheduling.PriorityClass)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1beta1_PriorityClassList_To_scheduling_PriorityClassList(in *v1beta1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_PriorityClassList_To_scheduling_PriorityClassList(in, out, s)
}
func autoConvert_scheduling_PriorityClassList_To_v1beta1_PriorityClassList(in *scheduling.PriorityClassList, out *v1beta1.PriorityClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1beta1.PriorityClass)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_scheduling_PriorityClassList_To_v1beta1_PriorityClassList(in *scheduling.PriorityClassList, out *v1beta1.PriorityClassList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_scheduling_PriorityClassList_To_v1beta1_PriorityClassList(in, out, s)
}
