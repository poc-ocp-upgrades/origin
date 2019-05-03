package batch

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *CronJob) DeepCopyInto(out *CronJob) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *CronJob) DeepCopy() *CronJob {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CronJob)
 in.DeepCopyInto(out)
 return out
}
func (in *CronJob) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *CronJobList) DeepCopyInto(out *CronJobList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]CronJob, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *CronJobList) DeepCopy() *CronJobList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CronJobList)
 in.DeepCopyInto(out)
 return out
}
func (in *CronJobList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *CronJobSpec) DeepCopyInto(out *CronJobSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.StartingDeadlineSeconds != nil {
  in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
  *out = new(int64)
  **out = **in
 }
 if in.Suspend != nil {
  in, out := &in.Suspend, &out.Suspend
  *out = new(bool)
  **out = **in
 }
 in.JobTemplate.DeepCopyInto(&out.JobTemplate)
 if in.SuccessfulJobsHistoryLimit != nil {
  in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
  *out = new(int32)
  **out = **in
 }
 if in.FailedJobsHistoryLimit != nil {
  in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *CronJobSpec) DeepCopy() *CronJobSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CronJobSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *CronJobStatus) DeepCopyInto(out *CronJobStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Active != nil {
  in, out := &in.Active, &out.Active
  *out = make([]core.ObjectReference, len(*in))
  copy(*out, *in)
 }
 if in.LastScheduleTime != nil {
  in, out := &in.LastScheduleTime, &out.LastScheduleTime
  *out = (*in).DeepCopy()
 }
 return
}
func (in *CronJobStatus) DeepCopy() *CronJobStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CronJobStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *Job) DeepCopyInto(out *Job) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Job) DeepCopy() *Job {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Job)
 in.DeepCopyInto(out)
 return out
}
func (in *Job) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *JobCondition) DeepCopyInto(out *JobCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *JobCondition) DeepCopy() *JobCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *JobList) DeepCopyInto(out *JobList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Job, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *JobList) DeepCopy() *JobList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobList)
 in.DeepCopyInto(out)
 return out
}
func (in *JobList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Parallelism != nil {
  in, out := &in.Parallelism, &out.Parallelism
  *out = new(int32)
  **out = **in
 }
 if in.Completions != nil {
  in, out := &in.Completions, &out.Completions
  *out = new(int32)
  **out = **in
 }
 if in.ActiveDeadlineSeconds != nil {
  in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
  *out = new(int64)
  **out = **in
 }
 if in.BackoffLimit != nil {
  in, out := &in.BackoffLimit, &out.BackoffLimit
  *out = new(int32)
  **out = **in
 }
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.ManualSelector != nil {
  in, out := &in.ManualSelector, &out.ManualSelector
  *out = new(bool)
  **out = **in
 }
 in.Template.DeepCopyInto(&out.Template)
 if in.TTLSecondsAfterFinished != nil {
  in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *JobSpec) DeepCopy() *JobSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]JobCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.StartTime != nil {
  in, out := &in.StartTime, &out.StartTime
  *out = (*in).DeepCopy()
 }
 if in.CompletionTime != nil {
  in, out := &in.CompletionTime, &out.CompletionTime
  *out = (*in).DeepCopy()
 }
 return
}
func (in *JobStatus) DeepCopy() *JobStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *JobTemplate) DeepCopyInto(out *JobTemplate) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Template.DeepCopyInto(&out.Template)
 return
}
func (in *JobTemplate) DeepCopy() *JobTemplate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobTemplate)
 in.DeepCopyInto(out)
 return out
}
func (in *JobTemplate) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *JobTemplateSpec) DeepCopyInto(out *JobTemplateSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *JobTemplateSpec) DeepCopy() *JobTemplateSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobTemplateSpec)
 in.DeepCopyInto(out)
 return out
}
