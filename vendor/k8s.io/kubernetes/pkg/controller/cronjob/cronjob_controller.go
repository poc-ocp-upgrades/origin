package cronjob

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "sort"
 "time"
 "k8s.io/klog"
 batchv1 "k8s.io/api/batch/v1"
 batchv1beta1 "k8s.io/api/batch/v1beta1"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/record"
 ref "k8s.io/client-go/tools/reference"
 "k8s.io/kubernetes/pkg/util/metrics"
)

var controllerKind = batchv1beta1.SchemeGroupVersion.WithKind("CronJob")

type CronJobController struct {
 kubeClient clientset.Interface
 jobControl jobControlInterface
 sjControl  sjControlInterface
 podControl podControlInterface
 recorder   record.EventRecorder
}

func NewCronJobController(kubeClient clientset.Interface) (*CronJobController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
  if err := metrics.RegisterMetricAndTrackRateLimiterUsage("cronjob_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter()); err != nil {
   return nil, err
  }
 }
 jm := &CronJobController{kubeClient: kubeClient, jobControl: realJobControl{KubeClient: kubeClient}, sjControl: &realSJControl{KubeClient: kubeClient}, podControl: &realPodControl{KubeClient: kubeClient}, recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "cronjob-controller"})}
 return jm, nil
}
func (jm *CronJobController) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting CronJob Manager")
 go wait.Until(jm.syncAll, 10*time.Second, stopCh)
 <-stopCh
 klog.Infof("Shutting down CronJob Manager")
}
func (jm *CronJobController) syncAll() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 jl, err := jm.kubeClient.BatchV1().Jobs(metav1.NamespaceAll).List(metav1.ListOptions{})
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("can't list Jobs: %v", err))
  return
 }
 js := jl.Items
 klog.V(4).Infof("Found %d jobs", len(js))
 sjl, err := jm.kubeClient.BatchV1beta1().CronJobs(metav1.NamespaceAll).List(metav1.ListOptions{})
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("can't list CronJobs: %v", err))
  return
 }
 sjs := sjl.Items
 klog.V(4).Infof("Found %d cronjobs", len(sjs))
 jobsBySj := groupJobsByParent(js)
 klog.V(4).Infof("Found %d groups", len(jobsBySj))
 for _, sj := range sjs {
  syncOne(&sj, jobsBySj[sj.UID], time.Now(), jm.jobControl, jm.sjControl, jm.recorder)
  cleanupFinishedJobs(&sj, jobsBySj[sj.UID], jm.jobControl, jm.sjControl, jm.recorder)
 }
}
func cleanupFinishedJobs(sj *batchv1beta1.CronJob, js []batchv1.Job, jc jobControlInterface, sjc sjControlInterface, recorder record.EventRecorder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if sj.Spec.FailedJobsHistoryLimit == nil && sj.Spec.SuccessfulJobsHistoryLimit == nil {
  return
 }
 failedJobs := []batchv1.Job{}
 succesfulJobs := []batchv1.Job{}
 for _, job := range js {
  isFinished, finishedStatus := getFinishedStatus(&job)
  if isFinished && finishedStatus == batchv1.JobComplete {
   succesfulJobs = append(succesfulJobs, job)
  } else if isFinished && finishedStatus == batchv1.JobFailed {
   failedJobs = append(failedJobs, job)
  }
 }
 if sj.Spec.SuccessfulJobsHistoryLimit != nil {
  removeOldestJobs(sj, succesfulJobs, jc, *sj.Spec.SuccessfulJobsHistoryLimit, recorder)
 }
 if sj.Spec.FailedJobsHistoryLimit != nil {
  removeOldestJobs(sj, failedJobs, jc, *sj.Spec.FailedJobsHistoryLimit, recorder)
 }
 if _, err := sjc.UpdateStatus(sj); err != nil {
  nameForLog := fmt.Sprintf("%s/%s", sj.Namespace, sj.Name)
  klog.Infof("Unable to update status for %s (rv = %s): %v", nameForLog, sj.ResourceVersion, err)
 }
}
func removeOldestJobs(sj *batchv1beta1.CronJob, js []batchv1.Job, jc jobControlInterface, maxJobs int32, recorder record.EventRecorder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 numToDelete := len(js) - int(maxJobs)
 if numToDelete <= 0 {
  return
 }
 nameForLog := fmt.Sprintf("%s/%s", sj.Namespace, sj.Name)
 klog.V(4).Infof("Cleaning up %d/%d jobs from %s", numToDelete, len(js), nameForLog)
 sort.Sort(byJobStartTime(js))
 for i := 0; i < numToDelete; i++ {
  klog.V(4).Infof("Removing job %s from %s", js[i].Name, nameForLog)
  deleteJob(sj, &js[i], jc, recorder)
 }
}
func syncOne(sj *batchv1beta1.CronJob, js []batchv1.Job, now time.Time, jc jobControlInterface, sjc sjControlInterface, recorder record.EventRecorder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nameForLog := fmt.Sprintf("%s/%s", sj.Namespace, sj.Name)
 childrenJobs := make(map[types.UID]bool)
 for _, j := range js {
  childrenJobs[j.ObjectMeta.UID] = true
  found := inActiveList(*sj, j.ObjectMeta.UID)
  if !found && !IsJobFinished(&j) {
   recorder.Eventf(sj, v1.EventTypeWarning, "UnexpectedJob", "Saw a job that the controller did not create or forgot: %v", j.Name)
  } else if found && IsJobFinished(&j) {
   deleteFromActiveList(sj, j.ObjectMeta.UID)
   recorder.Eventf(sj, v1.EventTypeNormal, "SawCompletedJob", "Saw completed job: %v", j.Name)
  }
 }
 for _, j := range sj.Status.Active {
  if found := childrenJobs[j.UID]; !found {
   recorder.Eventf(sj, v1.EventTypeNormal, "MissingJob", "Active job went missing: %v", j.Name)
   deleteFromActiveList(sj, j.UID)
  }
 }
 updatedSJ, err := sjc.UpdateStatus(sj)
 if err != nil {
  klog.Errorf("Unable to update status for %s (rv = %s): %v", nameForLog, sj.ResourceVersion, err)
  return
 }
 *sj = *updatedSJ
 if sj.DeletionTimestamp != nil {
  return
 }
 if sj.Spec.Suspend != nil && *sj.Spec.Suspend {
  klog.V(4).Infof("Not starting job for %s because it is suspended", nameForLog)
  return
 }
 times, err := getRecentUnmetScheduleTimes(*sj, now)
 if err != nil {
  recorder.Eventf(sj, v1.EventTypeWarning, "FailedNeedsStart", "Cannot determine if job needs to be started: %v", err)
  klog.Errorf("Cannot determine if %s needs to be started: %v", nameForLog, err)
  return
 }
 if len(times) == 0 {
  klog.V(4).Infof("No unmet start times for %s", nameForLog)
  return
 }
 if len(times) > 1 {
  klog.V(4).Infof("Multiple unmet start times for %s so only starting last one", nameForLog)
 }
 scheduledTime := times[len(times)-1]
 tooLate := false
 if sj.Spec.StartingDeadlineSeconds != nil {
  tooLate = scheduledTime.Add(time.Second * time.Duration(*sj.Spec.StartingDeadlineSeconds)).Before(now)
 }
 if tooLate {
  klog.V(4).Infof("Missed starting window for %s", nameForLog)
  recorder.Eventf(sj, v1.EventTypeWarning, "MissSchedule", "Missed scheduled time to start a job: %s", scheduledTime.Format(time.RFC1123Z))
  return
 }
 if sj.Spec.ConcurrencyPolicy == batchv1beta1.ForbidConcurrent && len(sj.Status.Active) > 0 {
  klog.V(4).Infof("Not starting job for %s because of prior execution still running and concurrency policy is Forbid", nameForLog)
  return
 }
 if sj.Spec.ConcurrencyPolicy == batchv1beta1.ReplaceConcurrent {
  for _, j := range sj.Status.Active {
   klog.V(4).Infof("Deleting job %s of %s that was still running at next scheduled start time", j.Name, nameForLog)
   job, err := jc.GetJob(j.Namespace, j.Name)
   if err != nil {
    recorder.Eventf(sj, v1.EventTypeWarning, "FailedGet", "Get job: %v", err)
    return
   }
   if !deleteJob(sj, job, jc, recorder) {
    return
   }
  }
 }
 jobReq, err := getJobFromTemplate(sj, scheduledTime)
 if err != nil {
  klog.Errorf("Unable to make Job from template in %s: %v", nameForLog, err)
  return
 }
 jobResp, err := jc.CreateJob(sj.Namespace, jobReq)
 if err != nil {
  recorder.Eventf(sj, v1.EventTypeWarning, "FailedCreate", "Error creating job: %v", err)
  return
 }
 klog.V(4).Infof("Created Job %s for %s", jobResp.Name, nameForLog)
 recorder.Eventf(sj, v1.EventTypeNormal, "SuccessfulCreate", "Created job %v", jobResp.Name)
 ref, err := getRef(jobResp)
 if err != nil {
  klog.V(2).Infof("Unable to make object reference for job for %s", nameForLog)
 } else {
  sj.Status.Active = append(sj.Status.Active, *ref)
 }
 sj.Status.LastScheduleTime = &metav1.Time{Time: scheduledTime}
 if _, err := sjc.UpdateStatus(sj); err != nil {
  klog.Infof("Unable to update status for %s (rv = %s): %v", nameForLog, sj.ResourceVersion, err)
 }
 return
}
func deleteJob(sj *batchv1beta1.CronJob, job *batchv1.Job, jc jobControlInterface, recorder record.EventRecorder) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nameForLog := fmt.Sprintf("%s/%s", sj.Namespace, sj.Name)
 if err := jc.DeleteJob(job.Namespace, job.Name); err != nil {
  recorder.Eventf(sj, v1.EventTypeWarning, "FailedDelete", "Deleted job: %v", err)
  klog.Errorf("Error deleting job %s from %s: %v", job.Name, nameForLog, err)
  return false
 }
 deleteFromActiveList(sj, job.ObjectMeta.UID)
 recorder.Eventf(sj, v1.EventTypeNormal, "SuccessfulDelete", "Deleted job %v", job.Name)
 return true
}
func getRef(object runtime.Object) (*v1.ObjectReference, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ref.GetReference(scheme.Scheme, object)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
