package cronjob

import (
	"fmt"
	"github.com/robfig/cron"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"time"
)

func inActiveList(sj batchv1beta1.CronJob, uid types.UID) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, j := range sj.Status.Active {
		if j.UID == uid {
			return true
		}
	}
	return false
}
func deleteFromActiveList(sj *batchv1beta1.CronJob, uid types.UID) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if sj == nil {
		return
	}
	newActive := []v1.ObjectReference{}
	for _, j := range sj.Status.Active {
		if j.UID != uid {
			newActive = append(newActive, j)
		}
	}
	sj.Status.Active = newActive
}
func getParentUIDFromJob(j batchv1.Job) (types.UID, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerRef := metav1.GetControllerOf(&j)
	if controllerRef == nil {
		return types.UID(""), false
	}
	if controllerRef.Kind != "CronJob" {
		klog.V(4).Infof("Job with non-CronJob parent, name %s namespace %s", j.Name, j.Namespace)
		return types.UID(""), false
	}
	return controllerRef.UID, true
}
func groupJobsByParent(js []batchv1.Job) map[types.UID][]batchv1.Job {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jobsBySj := make(map[types.UID][]batchv1.Job)
	for _, job := range js {
		parentUID, found := getParentUIDFromJob(job)
		if !found {
			klog.V(4).Infof("Unable to get parent uid from job %s in namespace %s", job.Name, job.Namespace)
			continue
		}
		jobsBySj[parentUID] = append(jobsBySj[parentUID], job)
	}
	return jobsBySj
}
func getRecentUnmetScheduleTimes(sj batchv1beta1.CronJob, now time.Time) ([]time.Time, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	starts := []time.Time{}
	sched, err := cron.ParseStandard(sj.Spec.Schedule)
	if err != nil {
		return starts, fmt.Errorf("Unparseable schedule: %s : %s", sj.Spec.Schedule, err)
	}
	var earliestTime time.Time
	if sj.Status.LastScheduleTime != nil {
		earliestTime = sj.Status.LastScheduleTime.Time
	} else {
		earliestTime = sj.ObjectMeta.CreationTimestamp.Time
	}
	if sj.Spec.StartingDeadlineSeconds != nil {
		schedulingDeadline := now.Add(-time.Second * time.Duration(*sj.Spec.StartingDeadlineSeconds))
		if schedulingDeadline.After(earliestTime) {
			earliestTime = schedulingDeadline
		}
	}
	if earliestTime.After(now) {
		return []time.Time{}, nil
	}
	for t := sched.Next(earliestTime); !t.After(now); t = sched.Next(t) {
		starts = append(starts, t)
		if len(starts) > 100 {
			return []time.Time{}, fmt.Errorf("Too many missed start time (> 100). Set or decrease .spec.startingDeadlineSeconds or check clock skew.")
		}
	}
	return starts, nil
}
func getJobFromTemplate(sj *batchv1beta1.CronJob, scheduledTime time.Time) (*batchv1.Job, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	labels := copyLabels(&sj.Spec.JobTemplate)
	annotations := copyAnnotations(&sj.Spec.JobTemplate)
	name := fmt.Sprintf("%s-%d", sj.Name, getTimeHash(scheduledTime))
	job := &batchv1.Job{ObjectMeta: metav1.ObjectMeta{Labels: labels, Annotations: annotations, Name: name, OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(sj, controllerKind)}}}
	if err := legacyscheme.Scheme.Convert(&sj.Spec.JobTemplate.Spec, &job.Spec, nil); err != nil {
		return nil, fmt.Errorf("unable to convert job template: %v", err)
	}
	return job, nil
}
func getTimeHash(scheduledTime time.Time) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return scheduledTime.Unix()
}
func getFinishedStatus(j *batchv1.Job) (bool, batchv1.JobConditionType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, c := range j.Status.Conditions {
		if (c.Type == batchv1.JobComplete || c.Type == batchv1.JobFailed) && c.Status == v1.ConditionTrue {
			return true, c.Type
		}
	}
	return false, ""
}
func IsJobFinished(j *batchv1.Job) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	isFinished, _ := getFinishedStatus(j)
	return isFinished
}

type byJobStartTime []batchv1.Job

func (o byJobStartTime) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(o)
}
func (o byJobStartTime) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o[i], o[j] = o[j], o[i]
}
func (o byJobStartTime) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o[i].Status.StartTime == nil && o[j].Status.StartTime != nil {
		return false
	}
	if o[i].Status.StartTime != nil && o[j].Status.StartTime == nil {
		return true
	}
	if o[i].Status.StartTime.Equal(o[j].Status.StartTime) {
		return o[i].Name < o[j].Name
	}
	return o[i].Status.StartTime.Before(o[j].Status.StartTime)
}
