package v1beta1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/core/v1"
 v1beta1 "k8s.io/api/events/v1beta1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1beta1.Event)(nil), (*core.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Event_To_core_Event(a.(*v1beta1.Event), b.(*core.Event), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Event)(nil), (*v1beta1.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Event_To_v1beta1_Event(a.(*core.Event), b.(*v1beta1.Event), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.EventList)(nil), (*core.EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_EventList_To_core_EventList(a.(*v1beta1.EventList), b.(*core.EventList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EventList)(nil), (*v1beta1.EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EventList_To_v1beta1_EventList(a.(*core.EventList), b.(*v1beta1.EventList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.EventSeries)(nil), (*core.EventSeries)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_EventSeries_To_core_EventSeries(a.(*v1beta1.EventSeries), b.(*core.EventSeries), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EventSeries)(nil), (*v1beta1.EventSeries)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EventSeries_To_v1beta1_EventSeries(a.(*core.EventSeries), b.(*v1beta1.EventSeries), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.Event)(nil), (*v1beta1.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Event_To_v1beta1_Event(a.(*core.Event), b.(*v1beta1.Event), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1beta1.Event)(nil), (*core.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_Event_To_core_Event(a.(*v1beta1.Event), b.(*core.Event), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_Event_To_core_Event(in *v1beta1.Event, out *core.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.EventTime = in.EventTime
 out.Series = (*core.EventSeries)(unsafe.Pointer(in.Series))
 out.ReportingController = in.ReportingController
 out.ReportingInstance = in.ReportingInstance
 out.Action = in.Action
 out.Reason = in.Reason
 out.Related = (*core.ObjectReference)(unsafe.Pointer(in.Related))
 out.Type = in.Type
 return nil
}
func autoConvert_core_Event_To_v1beta1_Event(in *core.Event, out *v1beta1.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Reason = in.Reason
 out.Type = in.Type
 out.EventTime = in.EventTime
 out.Series = (*v1beta1.EventSeries)(unsafe.Pointer(in.Series))
 out.Action = in.Action
 out.Related = (*v1.ObjectReference)(unsafe.Pointer(in.Related))
 out.ReportingController = in.ReportingController
 out.ReportingInstance = in.ReportingInstance
 return nil
}
func autoConvert_v1beta1_EventList_To_core_EventList(in *v1beta1.EventList, out *core.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.Event, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_Event_To_core_Event(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_EventList_To_core_EventList(in *v1beta1.EventList, out *core.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_EventList_To_core_EventList(in, out, s)
}
func autoConvert_core_EventList_To_v1beta1_EventList(in *core.EventList, out *v1beta1.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.Event, len(*in))
  for i := range *in {
   if err := Convert_core_Event_To_v1beta1_Event(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_EventList_To_v1beta1_EventList(in *core.EventList, out *v1beta1.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EventList_To_v1beta1_EventList(in, out, s)
}
func autoConvert_v1beta1_EventSeries_To_core_EventSeries(in *v1beta1.EventSeries, out *core.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Count = in.Count
 out.LastObservedTime = in.LastObservedTime
 out.State = core.EventSeriesState(in.State)
 return nil
}
func Convert_v1beta1_EventSeries_To_core_EventSeries(in *v1beta1.EventSeries, out *core.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_EventSeries_To_core_EventSeries(in, out, s)
}
func autoConvert_core_EventSeries_To_v1beta1_EventSeries(in *core.EventSeries, out *v1beta1.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Count = in.Count
 out.LastObservedTime = in.LastObservedTime
 out.State = v1beta1.EventSeriesState(in.State)
 return nil
}
func Convert_core_EventSeries_To_v1beta1_EventSeries(in *core.EventSeries, out *v1beta1.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EventSeries_To_v1beta1_EventSeries(in, out, s)
}
