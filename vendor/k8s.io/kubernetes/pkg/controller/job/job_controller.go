package job

import (
	"fmt"
	goformat "fmt"
	batch "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	batchinformers "k8s.io/client-go/informers/batch/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	batchv1listers "k8s.io/client-go/listers/batch/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/integer"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
	"math"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"sort"
	"sync"
	"time"
	gotime "time"
)

const statusUpdateRetries = 3

var controllerKind = batch.SchemeGroupVersion.WithKind("Job")
var (
	DefaultJobBackOff = 10 * time.Second
	MaxJobBackOff     = 360 * time.Second
)

type JobController struct {
	kubeClient     clientset.Interface
	podControl     controller.PodControlInterface
	updateHandler  func(job *batch.Job) error
	syncHandler    func(jobKey string) (bool, error)
	podStoreSynced cache.InformerSynced
	jobStoreSynced cache.InformerSynced
	expectations   controller.ControllerExpectationsInterface
	jobLister      batchv1listers.JobLister
	podStore       corelisters.PodLister
	queue          workqueue.RateLimitingInterface
	recorder       record.EventRecorder
}

func NewJobController(podInformer coreinformers.PodInformer, jobInformer batchinformers.JobInformer, kubeClient clientset.Interface) *JobController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("job_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
	}
	jm := &JobController{kubeClient: kubeClient, podControl: controller.RealPodControl{KubeClient: kubeClient, Recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "job-controller"})}, expectations: controller.NewControllerExpectations(), queue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(DefaultJobBackOff, MaxJobBackOff), "job"), recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "job-controller"})}
	jobInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		jm.enqueueController(obj, true)
	}, UpdateFunc: jm.updateJob, DeleteFunc: func(obj interface{}) {
		jm.enqueueController(obj, true)
	}})
	jm.jobLister = jobInformer.Lister()
	jm.jobStoreSynced = jobInformer.Informer().HasSynced
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: jm.addPod, UpdateFunc: jm.updatePod, DeleteFunc: jm.deletePod})
	jm.podStore = podInformer.Lister()
	jm.podStoreSynced = podInformer.Informer().HasSynced
	jm.updateHandler = jm.updateJobStatus
	jm.syncHandler = jm.syncJob
	return jm
}
func (jm *JobController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer jm.queue.ShutDown()
	klog.Infof("Starting job controller")
	defer klog.Infof("Shutting down job controller")
	if !controller.WaitForCacheSync("job", stopCh, jm.podStoreSynced, jm.jobStoreSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(jm.worker, time.Second, stopCh)
	}
	<-stopCh
}
func (jm *JobController) getPodJobs(pod *v1.Pod) []*batch.Job {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jobs, err := jm.jobLister.GetPodJobs(pod)
	if err != nil {
		return nil
	}
	if len(jobs) > 1 {
		utilruntime.HandleError(fmt.Errorf("user error! more than one job is selecting pods with labels: %+v", pod.Labels))
	}
	ret := make([]*batch.Job, 0, len(jobs))
	for i := range jobs {
		ret = append(ret, &jobs[i])
	}
	return ret
}
func (jm *JobController) resolveControllerRef(namespace string, controllerRef *metav1.OwnerReference) *batch.Job {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if controllerRef.Kind != controllerKind.Kind {
		return nil
	}
	job, err := jm.jobLister.Jobs(namespace).Get(controllerRef.Name)
	if err != nil {
		return nil
	}
	if job.UID != controllerRef.UID {
		return nil
	}
	return job
}
func (jm *JobController) addPod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := obj.(*v1.Pod)
	if pod.DeletionTimestamp != nil {
		jm.deletePod(pod)
		return
	}
	if controllerRef := metav1.GetControllerOf(pod); controllerRef != nil {
		job := jm.resolveControllerRef(pod.Namespace, controllerRef)
		if job == nil {
			return
		}
		jobKey, err := controller.KeyFunc(job)
		if err != nil {
			return
		}
		jm.expectations.CreationObserved(jobKey)
		jm.enqueueController(job, true)
		return
	}
	for _, job := range jm.getPodJobs(pod) {
		jm.enqueueController(job, true)
	}
}
func (jm *JobController) updatePod(old, cur interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	curPod := cur.(*v1.Pod)
	oldPod := old.(*v1.Pod)
	if curPod.ResourceVersion == oldPod.ResourceVersion {
		return
	}
	if curPod.DeletionTimestamp != nil {
		jm.deletePod(curPod)
		return
	}
	immediate := curPod.Status.Phase != v1.PodFailed
	curControllerRef := metav1.GetControllerOf(curPod)
	oldControllerRef := metav1.GetControllerOf(oldPod)
	controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
	if controllerRefChanged && oldControllerRef != nil {
		if job := jm.resolveControllerRef(oldPod.Namespace, oldControllerRef); job != nil {
			jm.enqueueController(job, immediate)
		}
	}
	if curControllerRef != nil {
		job := jm.resolveControllerRef(curPod.Namespace, curControllerRef)
		if job == nil {
			return
		}
		jm.enqueueController(job, immediate)
		return
	}
	labelChanged := !reflect.DeepEqual(curPod.Labels, oldPod.Labels)
	if labelChanged || controllerRefChanged {
		for _, job := range jm.getPodJobs(curPod) {
			jm.enqueueController(job, immediate)
		}
	}
}
func (jm *JobController) deletePod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := obj.(*v1.Pod)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %+v", obj))
			return
		}
		pod, ok = tombstone.Obj.(*v1.Pod)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod %+v", obj))
			return
		}
	}
	controllerRef := metav1.GetControllerOf(pod)
	if controllerRef == nil {
		return
	}
	job := jm.resolveControllerRef(pod.Namespace, controllerRef)
	if job == nil {
		return
	}
	jobKey, err := controller.KeyFunc(job)
	if err != nil {
		return
	}
	jm.expectations.DeletionObserved(jobKey)
	jm.enqueueController(job, true)
}
func (jm *JobController) updateJob(old, cur interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldJob := old.(*batch.Job)
	curJob := cur.(*batch.Job)
	key, err := controller.KeyFunc(curJob)
	if err != nil {
		return
	}
	jm.enqueueController(curJob, true)
	if curJob.Status.StartTime != nil {
		curADS := curJob.Spec.ActiveDeadlineSeconds
		if curADS == nil {
			return
		}
		oldADS := oldJob.Spec.ActiveDeadlineSeconds
		if oldADS == nil || *oldADS != *curADS {
			now := metav1.Now()
			start := curJob.Status.StartTime.Time
			passed := now.Time.Sub(start)
			total := time.Duration(*curADS) * time.Second
			jm.queue.AddAfter(key, total-passed)
			klog.V(4).Infof("job ActiveDeadlineSeconds updated, will rsync after %d seconds", total-passed)
		}
	}
}
func (jm *JobController) enqueueController(obj interface{}, immediate bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}
	backoff := time.Duration(0)
	if !immediate {
		backoff = getBackoff(jm.queue, key)
	}
	jm.queue.AddAfter(key, backoff)
}
func (jm *JobController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for jm.processNextWorkItem() {
	}
}
func (jm *JobController) processNextWorkItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := jm.queue.Get()
	if quit {
		return false
	}
	defer jm.queue.Done(key)
	forget, err := jm.syncHandler(key.(string))
	if err == nil {
		if forget {
			jm.queue.Forget(key)
		}
		return true
	}
	utilruntime.HandleError(fmt.Errorf("Error syncing job: %v", err))
	jm.queue.AddRateLimited(key)
	return true
}
func (jm *JobController) getPodsForJob(j *batch.Job) ([]*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	selector, err := metav1.LabelSelectorAsSelector(j.Spec.Selector)
	if err != nil {
		return nil, fmt.Errorf("couldn't convert Job selector: %v", err)
	}
	pods, err := jm.podStore.Pods(j.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	canAdoptFunc := controller.RecheckDeletionTimestamp(func() (metav1.Object, error) {
		fresh, err := jm.kubeClient.BatchV1().Jobs(j.Namespace).Get(j.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if fresh.UID != j.UID {
			return nil, fmt.Errorf("original Job %v/%v is gone: got uid %v, wanted %v", j.Namespace, j.Name, fresh.UID, j.UID)
		}
		return fresh, nil
	})
	cm := controller.NewPodControllerRefManager(jm.podControl, j, selector, controllerKind, canAdoptFunc)
	return cm.ClaimPods(pods)
}
func (jm *JobController) syncJob(key string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing job %q (%v)", key, time.Since(startTime))
	}()
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return false, err
	}
	if len(ns) == 0 || len(name) == 0 {
		return false, fmt.Errorf("invalid job key %q: either namespace or name is missing", key)
	}
	sharedJob, err := jm.jobLister.Jobs(ns).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			klog.V(4).Infof("Job has been deleted: %v", key)
			jm.expectations.DeleteExpectations(key)
			return true, nil
		}
		return false, err
	}
	job := *sharedJob
	if IsJobFinished(&job) {
		return true, nil
	}
	previousRetry := jm.queue.NumRequeues(key)
	jobNeedsSync := jm.expectations.SatisfiedExpectations(key)
	pods, err := jm.getPodsForJob(&job)
	if err != nil {
		return false, err
	}
	activePods := controller.FilterActivePods(pods)
	active := int32(len(activePods))
	succeeded, failed := getStatus(pods)
	conditions := len(job.Status.Conditions)
	if job.Status.StartTime == nil {
		now := metav1.Now()
		job.Status.StartTime = &now
		if job.Spec.ActiveDeadlineSeconds != nil {
			klog.V(4).Infof("Job %s have ActiveDeadlineSeconds will sync after %d seconds", key, *job.Spec.ActiveDeadlineSeconds)
			jm.queue.AddAfter(key, time.Duration(*job.Spec.ActiveDeadlineSeconds)*time.Second)
		}
	}
	var manageJobErr error
	jobFailed := false
	var failureReason string
	var failureMessage string
	jobHaveNewFailure := failed > job.Status.Failed
	exceedsBackoffLimit := jobHaveNewFailure && (active != *job.Spec.Parallelism) && (int32(previousRetry)+1 > *job.Spec.BackoffLimit)
	if exceedsBackoffLimit || pastBackoffLimitOnFailure(&job, pods) {
		jobFailed = true
		failureReason = "BackoffLimitExceeded"
		failureMessage = "Job has reached the specified backoff limit"
	} else if pastActiveDeadline(&job) {
		jobFailed = true
		failureReason = "DeadlineExceeded"
		failureMessage = "Job was active longer than specified deadline"
	}
	if jobFailed {
		errCh := make(chan error, active)
		jm.deleteJobPods(&job, activePods, errCh)
		select {
		case manageJobErr = <-errCh:
			if manageJobErr != nil {
				break
			}
		default:
		}
		failed += active
		active = 0
		job.Status.Conditions = append(job.Status.Conditions, newCondition(batch.JobFailed, failureReason, failureMessage))
		jm.recorder.Event(&job, v1.EventTypeWarning, failureReason, failureMessage)
	} else {
		if jobNeedsSync && job.DeletionTimestamp == nil {
			active, manageJobErr = jm.manageJob(activePods, succeeded, &job)
		}
		completions := succeeded
		complete := false
		if job.Spec.Completions == nil {
			if succeeded > 0 && active == 0 {
				complete = true
			}
		} else {
			if completions >= *job.Spec.Completions {
				complete = true
				if active > 0 {
					jm.recorder.Event(&job, v1.EventTypeWarning, "TooManyActivePods", "Too many active pods running after completion count reached")
				}
				if completions > *job.Spec.Completions {
					jm.recorder.Event(&job, v1.EventTypeWarning, "TooManySucceededPods", "Too many succeeded pods running after completion count reached")
				}
			}
		}
		if complete {
			job.Status.Conditions = append(job.Status.Conditions, newCondition(batch.JobComplete, "", ""))
			now := metav1.Now()
			job.Status.CompletionTime = &now
		}
	}
	forget := false
	if job.Status.Succeeded < succeeded {
		forget = true
	}
	if job.Status.Active != active || job.Status.Succeeded != succeeded || job.Status.Failed != failed || len(job.Status.Conditions) != conditions {
		job.Status.Active = active
		job.Status.Succeeded = succeeded
		job.Status.Failed = failed
		if err := jm.updateHandler(&job); err != nil {
			return forget, err
		}
		if jobHaveNewFailure && !IsJobFinished(&job) {
			return forget, fmt.Errorf("failed pod(s) detected for job key %q", key)
		}
		forget = true
	}
	return forget, manageJobErr
}
func (jm *JobController) deleteJobPods(job *batch.Job, pods []*v1.Pod, errCh chan<- error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait := sync.WaitGroup{}
	nbPods := len(pods)
	wait.Add(nbPods)
	for i := int32(0); i < int32(nbPods); i++ {
		go func(ix int32) {
			defer wait.Done()
			if err := jm.podControl.DeletePod(job.Namespace, pods[ix].Name, job); err != nil {
				defer utilruntime.HandleError(err)
				klog.V(2).Infof("Failed to delete %v, job %q/%q deadline exceeded", pods[ix].Name, job.Namespace, job.Name)
				errCh <- err
			}
		}(i)
	}
	wait.Wait()
}
func pastBackoffLimitOnFailure(job *batch.Job, pods []*v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if job.Spec.Template.Spec.RestartPolicy != v1.RestartPolicyOnFailure {
		return false
	}
	result := int32(0)
	for i := range pods {
		po := pods[i]
		if po.Status.Phase != v1.PodRunning {
			continue
		}
		for j := range po.Status.InitContainerStatuses {
			stat := po.Status.InitContainerStatuses[j]
			result += stat.RestartCount
		}
		for j := range po.Status.ContainerStatuses {
			stat := po.Status.ContainerStatuses[j]
			result += stat.RestartCount
		}
	}
	if *job.Spec.BackoffLimit == 0 {
		return result > 0
	}
	return result >= *job.Spec.BackoffLimit
}
func pastActiveDeadline(job *batch.Job) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if job.Spec.ActiveDeadlineSeconds == nil || job.Status.StartTime == nil {
		return false
	}
	now := metav1.Now()
	start := job.Status.StartTime.Time
	duration := now.Time.Sub(start)
	allowedDuration := time.Duration(*job.Spec.ActiveDeadlineSeconds) * time.Second
	return duration >= allowedDuration
}
func newCondition(conditionType batch.JobConditionType, reason, message string) batch.JobCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return batch.JobCondition{Type: conditionType, Status: v1.ConditionTrue, LastProbeTime: metav1.Now(), LastTransitionTime: metav1.Now(), Reason: reason, Message: message}
}
func getStatus(pods []*v1.Pod) (succeeded, failed int32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	succeeded = int32(filterPods(pods, v1.PodSucceeded))
	failed = int32(filterPods(pods, v1.PodFailed))
	return
}
func (jm *JobController) manageJob(activePods []*v1.Pod, succeeded int32, job *batch.Job) (int32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var activeLock sync.Mutex
	active := int32(len(activePods))
	parallelism := *job.Spec.Parallelism
	jobKey, err := controller.KeyFunc(job)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for job %#v: %v", job, err))
		return 0, nil
	}
	var errCh chan error
	if active > parallelism {
		diff := active - parallelism
		errCh = make(chan error, diff)
		jm.expectations.ExpectDeletions(jobKey, int(diff))
		klog.V(4).Infof("Too many pods running job %q, need %d, deleting %d", jobKey, parallelism, diff)
		sort.Sort(controller.ActivePods(activePods))
		active -= diff
		wait := sync.WaitGroup{}
		wait.Add(int(diff))
		for i := int32(0); i < diff; i++ {
			go func(ix int32) {
				defer wait.Done()
				if err := jm.podControl.DeletePod(job.Namespace, activePods[ix].Name, job); err != nil {
					defer utilruntime.HandleError(err)
					klog.V(2).Infof("Failed to delete %v, decrementing expectations for job %q/%q", activePods[ix].Name, job.Namespace, job.Name)
					jm.expectations.DeletionObserved(jobKey)
					activeLock.Lock()
					active++
					activeLock.Unlock()
					errCh <- err
				}
			}(i)
		}
		wait.Wait()
	} else if active < parallelism {
		wantActive := int32(0)
		if job.Spec.Completions == nil {
			if succeeded > 0 {
				wantActive = active
			} else {
				wantActive = parallelism
			}
		} else {
			wantActive = *job.Spec.Completions - succeeded
			if wantActive > parallelism {
				wantActive = parallelism
			}
		}
		diff := wantActive - active
		if diff < 0 {
			utilruntime.HandleError(fmt.Errorf("More active than wanted: job %q, want %d, have %d", jobKey, wantActive, active))
			diff = 0
		}
		jm.expectations.ExpectCreations(jobKey, int(diff))
		errCh = make(chan error, diff)
		klog.V(4).Infof("Too few pods running job %q, need %d, creating %d", jobKey, wantActive, diff)
		active += diff
		wait := sync.WaitGroup{}
		for batchSize := int32(integer.IntMin(int(diff), controller.SlowStartInitialBatchSize)); diff > 0; batchSize = integer.Int32Min(2*batchSize, diff) {
			errorCount := len(errCh)
			wait.Add(int(batchSize))
			for i := int32(0); i < batchSize; i++ {
				go func() {
					defer wait.Done()
					err := jm.podControl.CreatePodsWithControllerRef(job.Namespace, &job.Spec.Template, job, metav1.NewControllerRef(job, controllerKind))
					if err != nil && errors.IsTimeout(err) {
						return
					}
					if err != nil {
						defer utilruntime.HandleError(err)
						klog.V(2).Infof("Failed creation, decrementing expectations for job %q/%q", job.Namespace, job.Name)
						jm.expectations.CreationObserved(jobKey)
						activeLock.Lock()
						active--
						activeLock.Unlock()
						errCh <- err
					}
				}()
			}
			wait.Wait()
			skippedPods := diff - batchSize
			if errorCount < len(errCh) && skippedPods > 0 {
				klog.V(2).Infof("Slow-start failure. Skipping creation of %d pods, decrementing expectations for job %q/%q", skippedPods, job.Namespace, job.Name)
				active -= skippedPods
				for i := int32(0); i < skippedPods; i++ {
					jm.expectations.CreationObserved(jobKey)
				}
				break
			}
			diff -= batchSize
		}
	}
	select {
	case err := <-errCh:
		if err != nil {
			return active, err
		}
	default:
	}
	return active, nil
}
func (jm *JobController) updateJobStatus(job *batch.Job) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jobClient := jm.kubeClient.BatchV1().Jobs(job.Namespace)
	var err error
	for i := 0; i <= statusUpdateRetries; i = i + 1 {
		var newJob *batch.Job
		newJob, err = jobClient.Get(job.Name, metav1.GetOptions{})
		if err != nil {
			break
		}
		newJob.Status = job.Status
		if _, err = jobClient.UpdateStatus(newJob); err == nil {
			break
		}
	}
	return err
}
func getBackoff(queue workqueue.RateLimitingInterface, key interface{}) time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	exp := queue.NumRequeues(key)
	if exp <= 0 {
		return time.Duration(0)
	}
	backoff := float64(DefaultJobBackOff.Nanoseconds()) * math.Pow(2, float64(exp-1))
	if backoff > math.MaxInt64 {
		return MaxJobBackOff
	}
	calculated := time.Duration(backoff)
	if calculated > MaxJobBackOff {
		return MaxJobBackOff
	}
	return calculated
}
func filterPods(pods []*v1.Pod, phase v1.PodPhase) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := 0
	for i := range pods {
		if phase == pods[i].Status.Phase {
			result++
		}
	}
	return result
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
