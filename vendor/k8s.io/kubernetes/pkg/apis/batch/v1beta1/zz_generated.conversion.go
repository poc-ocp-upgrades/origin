package v1beta1

import (
 unsafe "unsafe"
 v1beta1 "k8s.io/api/batch/v1beta1"
 corev1 "k8s.io/api/core/v1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 batch "k8s.io/kubernetes/pkg/apis/batch"
 batchv1 "k8s.io/kubernetes/pkg/apis/batch/v1"
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
 if err := s.AddGeneratedConversionFunc((*v1beta1.CronJob)(nil), (*batch.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CronJob_To_batch_CronJob(a.(*v1beta1.CronJob), b.(*batch.CronJob), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.CronJob)(nil), (*v1beta1.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_CronJob_To_v1beta1_CronJob(a.(*batch.CronJob), b.(*v1beta1.CronJob), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobList)(nil), (*batch.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CronJobList_To_batch_CronJobList(a.(*v1beta1.CronJobList), b.(*batch.CronJobList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.CronJobList)(nil), (*v1beta1.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_CronJobList_To_v1beta1_CronJobList(a.(*batch.CronJobList), b.(*v1beta1.CronJobList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobSpec)(nil), (*batch.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CronJobSpec_To_batch_CronJobSpec(a.(*v1beta1.CronJobSpec), b.(*batch.CronJobSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.CronJobSpec)(nil), (*v1beta1.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_CronJobSpec_To_v1beta1_CronJobSpec(a.(*batch.CronJobSpec), b.(*v1beta1.CronJobSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobStatus)(nil), (*batch.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_CronJobStatus_To_batch_CronJobStatus(a.(*v1beta1.CronJobStatus), b.(*batch.CronJobStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.CronJobStatus)(nil), (*v1beta1.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_CronJobStatus_To_v1beta1_CronJobStatus(a.(*batch.CronJobStatus), b.(*v1beta1.CronJobStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.JobTemplate)(nil), (*batch.JobTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_JobTemplate_To_batch_JobTemplate(a.(*v1beta1.JobTemplate), b.(*batch.JobTemplate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobTemplate)(nil), (*v1beta1.JobTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobTemplate_To_v1beta1_JobTemplate(a.(*batch.JobTemplate), b.(*v1beta1.JobTemplate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1beta1.JobTemplateSpec)(nil), (*batch.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(a.(*v1beta1.JobTemplateSpec), b.(*batch.JobTemplateSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*batch.JobTemplateSpec)(nil), (*v1beta1.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(a.(*batch.JobTemplateSpec), b.(*v1beta1.JobTemplateSpec), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1beta1_CronJob_To_batch_CronJob(in *v1beta1.CronJob, out *batch.CronJob, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_CronJobSpec_To_batch_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1beta1_CronJobStatus_To_batch_CronJobStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_CronJob_To_batch_CronJob(in *v1beta1.CronJob, out *batch.CronJob, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CronJob_To_batch_CronJob(in, out, s)
}
func autoConvert_batch_CronJob_To_v1beta1_CronJob(in *batch.CronJob, out *v1beta1.CronJob, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_batch_CronJobSpec_To_v1beta1_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_batch_CronJobStatus_To_v1beta1_CronJobStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_batch_CronJob_To_v1beta1_CronJob(in *batch.CronJob, out *v1beta1.CronJob, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_CronJob_To_v1beta1_CronJob(in, out, s)
}
func autoConvert_v1beta1_CronJobList_To_batch_CronJobList(in *v1beta1.CronJobList, out *batch.CronJobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]batch.CronJob, len(*in))
  for i := range *in {
   if err := Convert_v1beta1_CronJob_To_batch_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1beta1_CronJobList_To_batch_CronJobList(in *v1beta1.CronJobList, out *batch.CronJobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CronJobList_To_batch_CronJobList(in, out, s)
}
func autoConvert_batch_CronJobList_To_v1beta1_CronJobList(in *batch.CronJobList, out *v1beta1.CronJobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1beta1.CronJob, len(*in))
  for i := range *in {
   if err := Convert_batch_CronJob_To_v1beta1_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_batch_CronJobList_To_v1beta1_CronJobList(in *batch.CronJobList, out *v1beta1.CronJobList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_CronJobList_To_v1beta1_CronJobList(in, out, s)
}
func autoConvert_v1beta1_CronJobSpec_To_batch_CronJobSpec(in *v1beta1.CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Schedule = in.Schedule
 out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
 out.ConcurrencyPolicy = batch.ConcurrencyPolicy(in.ConcurrencyPolicy)
 out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
 if err := Convert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
  return err
 }
 out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
 out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
 return nil
}
func Convert_v1beta1_CronJobSpec_To_batch_CronJobSpec(in *v1beta1.CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CronJobSpec_To_batch_CronJobSpec(in, out, s)
}
func autoConvert_batch_CronJobSpec_To_v1beta1_CronJobSpec(in *batch.CronJobSpec, out *v1beta1.CronJobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Schedule = in.Schedule
 out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
 out.ConcurrencyPolicy = v1beta1.ConcurrencyPolicy(in.ConcurrencyPolicy)
 out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
 if err := Convert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
  return err
 }
 out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
 out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
 return nil
}
func Convert_batch_CronJobSpec_To_v1beta1_CronJobSpec(in *batch.CronJobSpec, out *v1beta1.CronJobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_CronJobSpec_To_v1beta1_CronJobSpec(in, out, s)
}
func autoConvert_v1beta1_CronJobStatus_To_batch_CronJobStatus(in *v1beta1.CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Active = *(*[]core.ObjectReference)(unsafe.Pointer(&in.Active))
 out.LastScheduleTime = (*v1.Time)(unsafe.Pointer(in.LastScheduleTime))
 return nil
}
func Convert_v1beta1_CronJobStatus_To_batch_CronJobStatus(in *v1beta1.CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_CronJobStatus_To_batch_CronJobStatus(in, out, s)
}
func autoConvert_batch_CronJobStatus_To_v1beta1_CronJobStatus(in *batch.CronJobStatus, out *v1beta1.CronJobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Active = *(*[]corev1.ObjectReference)(unsafe.Pointer(&in.Active))
 out.LastScheduleTime = (*v1.Time)(unsafe.Pointer(in.LastScheduleTime))
 return nil
}
func Convert_batch_CronJobStatus_To_v1beta1_CronJobStatus(in *batch.CronJobStatus, out *v1beta1.CronJobStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_CronJobStatus_To_v1beta1_CronJobStatus(in, out, s)
}
func autoConvert_v1beta1_JobTemplate_To_batch_JobTemplate(in *v1beta1.JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_JobTemplate_To_batch_JobTemplate(in *v1beta1.JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_JobTemplate_To_batch_JobTemplate(in, out, s)
}
func autoConvert_batch_JobTemplate_To_v1beta1_JobTemplate(in *batch.JobTemplate, out *v1beta1.JobTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_batch_JobTemplate_To_v1beta1_JobTemplate(in *batch.JobTemplate, out *v1beta1.JobTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_JobTemplate_To_v1beta1_JobTemplate(in, out, s)
}
func autoConvert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(in *v1beta1.JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := batchv1.Convert_v1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(in *v1beta1.JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1beta1_JobTemplateSpec_To_batch_JobTemplateSpec(in, out, s)
}
func autoConvert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in *batch.JobTemplateSpec, out *v1beta1.JobTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := batchv1.Convert_batch_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in *batch.JobTemplateSpec, out *v1beta1.JobTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_batch_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in, out, s)
}
