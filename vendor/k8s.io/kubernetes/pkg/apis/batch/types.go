package batch

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type Job struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   JobSpec
 Status JobStatus
}
type JobList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []Job
}
type JobTemplate struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Template JobTemplateSpec
}
type JobTemplateSpec struct {
 metav1.ObjectMeta
 Spec JobSpec
}
type JobSpec struct {
 Parallelism             *int32
 Completions             *int32
 ActiveDeadlineSeconds   *int64
 BackoffLimit            *int32
 Selector                *metav1.LabelSelector
 ManualSelector          *bool
 Template                api.PodTemplateSpec
 TTLSecondsAfterFinished *int32
}
type JobStatus struct {
 Conditions     []JobCondition
 StartTime      *metav1.Time
 CompletionTime *metav1.Time
 Active         int32
 Succeeded      int32
 Failed         int32
}
type JobConditionType string

const (
 JobComplete JobConditionType = "Complete"
 JobFailed   JobConditionType = "Failed"
)

type JobCondition struct {
 Type               JobConditionType
 Status             api.ConditionStatus
 LastProbeTime      metav1.Time
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
type CronJob struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   CronJobSpec
 Status CronJobStatus
}
type CronJobList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []CronJob
}
type CronJobSpec struct {
 Schedule                   string
 StartingDeadlineSeconds    *int64
 ConcurrencyPolicy          ConcurrencyPolicy
 Suspend                    *bool
 JobTemplate                JobTemplateSpec
 SuccessfulJobsHistoryLimit *int32
 FailedJobsHistoryLimit     *int32
}
type ConcurrencyPolicy string

const (
 AllowConcurrent   ConcurrencyPolicy = "Allow"
 ForbidConcurrent  ConcurrencyPolicy = "Forbid"
 ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

type CronJobStatus struct {
 Active           []api.ObjectReference
 LastScheduleTime *metav1.Time
}
