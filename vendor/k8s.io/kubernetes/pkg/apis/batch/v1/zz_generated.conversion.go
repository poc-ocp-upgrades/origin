package v1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/batch/v1"
 corev1 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 batch "k8s.io/kubernetes/pkg/apis/batch"
 core "k8s.io/kubernetes/pkg/apis/core"
 apiscorev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1.Job)(nil), (*batch.Job)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Job_To_batch_Job(a.(*v1.Job), b.(*batch.Job), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.Job)(nil), (*v1.Job)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_Job_To_v1_Job(a.(*batch.Job), b.(*v1.Job), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.JobCondition)(nil), (*batch.JobCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_JobCondition_To_batch_JobCondition(a.(*v1.JobCondition), b.(*batch.JobCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobCondition)(nil), (*v1.JobCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobCondition_To_v1_JobCondition(a.(*batch.JobCondition), b.(*v1.JobCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.JobList)(nil), (*batch.JobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_JobList_To_batch_JobList(a.(*v1.JobList), b.(*batch.JobList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobList)(nil), (*v1.JobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobList_To_v1_JobList(a.(*batch.JobList), b.(*v1.JobList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.JobSpec)(nil), (*batch.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_JobSpec_To_batch_JobSpec(a.(*v1.JobSpec), b.(*batch.JobSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobSpec)(nil), (*v1.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobSpec_To_v1_JobSpec(a.(*batch.JobSpec), b.(*v1.JobSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.JobStatus)(nil), (*batch.JobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_JobStatus_To_batch_JobStatus(a.(*v1.JobStatus), b.(*batch.JobStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobStatus)(nil), (*v1.JobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobStatus_To_v1_JobStatus(a.(*batch.JobStatus), b.(*v1.JobStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*batch.JobSpec)(nil), (*v1.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobSpec_To_v1_JobSpec(a.(*batch.JobSpec), b.(*v1.JobSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.JobSpec)(nil), (*batch.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_JobSpec_To_batch_JobSpec(a.(*v1.JobSpec), b.(*batch.JobSpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_JobStatus_To_batch_JobStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Job_To_batch_Job(in, out, s)
}
func autoConvert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_batch_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_batch_JobStatus_To_v1_JobStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_Job_To_v1_Job(in, out, s)
}
func autoConvert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = batch.JobConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_JobCondition_To_batch_JobCondition(in, out, s)
}
func autoConvert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.JobConditionType(in.Type)
 out.Status = corev1.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_JobCondition_To_v1_JobCondition(in, out, s)
}
func autoConvert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]batch.Job, len(*in))
  for i := range *in {
   if err := Convert_v1_Job_To_batch_Job(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_JobList_To_batch_JobList(in, out, s)
}
func autoConvert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.Job, len(*in))
  for i := range *in {
   if err := Convert_batch_Job_To_v1_Job(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_JobList_To_v1_JobList(in, out, s)
}
func autoConvert_v1_JobSpec_To_batch_JobSpec(in *v1.JobSpec, out *batch.JobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
 out.Completions = (*int32)(unsafe.Pointer(in.Completions))
 out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
 out.BackoffLimit = (*int32)(unsafe.Pointer(in.BackoffLimit))
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
 if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.TTLSecondsAfterFinished = (*int32)(unsafe.Pointer(in.TTLSecondsAfterFinished))
 return nil
}
func autoConvert_batch_JobSpec_To_v1_JobSpec(in *batch.JobSpec, out *v1.JobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
 out.Completions = (*int32)(unsafe.Pointer(in.Completions))
 out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
 out.BackoffLimit = (*int32)(unsafe.Pointer(in.BackoffLimit))
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
 if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 out.TTLSecondsAfterFinished = (*int32)(unsafe.Pointer(in.TTLSecondsAfterFinished))
 return nil
}
func autoConvert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Conditions = *(*[]batch.JobCondition)(unsafe.Pointer(&in.Conditions))
 out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
 out.CompletionTime = (*metav1.Time)(unsafe.Pointer(in.CompletionTime))
 out.Active = in.Active
 out.Succeeded = in.Succeeded
 out.Failed = in.Failed
 return nil
}
func Convert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_JobStatus_To_batch_JobStatus(in, out, s)
}
func autoConvert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Conditions = *(*[]v1.JobCondition)(unsafe.Pointer(&in.Conditions))
 out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
 out.CompletionTime = (*metav1.Time)(unsafe.Pointer(in.CompletionTime))
 out.Active = in.Active
 out.Succeeded = in.Succeeded
 out.Failed = in.Failed
 return nil
}
func Convert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_JobStatus_To_v1_JobStatus(in, out, s)
}
