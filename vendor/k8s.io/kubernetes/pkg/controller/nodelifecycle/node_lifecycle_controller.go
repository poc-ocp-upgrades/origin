package nodelifecycle

import (
 "context"
 "fmt"
 "hash/fnv"
 "io"
 "sync"
 "time"
 "k8s.io/klog"
 coordv1beta1 "k8s.io/api/coordination/v1beta1"
 "k8s.io/api/core/v1"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 coordinformers "k8s.io/client-go/informers/coordination/v1beta1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 extensionsinformers "k8s.io/client-go/informers/extensions/v1beta1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 coordlisters "k8s.io/client-go/listers/coordination/v1beta1"
 corelisters "k8s.io/client-go/listers/core/v1"
 extensionslisters "k8s.io/client-go/listers/extensions/v1beta1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/flowcontrol"
 "k8s.io/client-go/util/workqueue"
 cloudprovider "k8s.io/cloud-provider"
 v1node "k8s.io/kubernetes/pkg/api/v1/node"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/nodelifecycle/scheduler"
 nodeutil "k8s.io/kubernetes/pkg/controller/util/node"
 "k8s.io/kubernetes/pkg/features"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 "k8s.io/kubernetes/pkg/util/metrics"
 utilnode "k8s.io/kubernetes/pkg/util/node"
 "k8s.io/kubernetes/pkg/util/system"
 taintutils "k8s.io/kubernetes/pkg/util/taints"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 Register()
}

var (
 UnreachableTaintTemplate         = &v1.Taint{Key: schedulerapi.TaintNodeUnreachable, Effect: v1.TaintEffectNoExecute}
 NotReadyTaintTemplate            = &v1.Taint{Key: schedulerapi.TaintNodeNotReady, Effect: v1.TaintEffectNoExecute}
 nodeConditionToTaintKeyStatusMap = map[v1.NodeConditionType]map[v1.ConditionStatus]string{v1.NodeReady: {v1.ConditionFalse: schedulerapi.TaintNodeNotReady, v1.ConditionUnknown: schedulerapi.TaintNodeUnreachable}, v1.NodeMemoryPressure: {v1.ConditionTrue: schedulerapi.TaintNodeMemoryPressure}, v1.NodeOutOfDisk: {v1.ConditionTrue: schedulerapi.TaintNodeOutOfDisk}, v1.NodeDiskPressure: {v1.ConditionTrue: schedulerapi.TaintNodeDiskPressure}, v1.NodeNetworkUnavailable: {v1.ConditionTrue: schedulerapi.TaintNodeNetworkUnavailable}, v1.NodePIDPressure: {v1.ConditionTrue: schedulerapi.TaintNodePIDPressure}}
 taintKeyToNodeConditionMap       = map[string]v1.NodeConditionType{schedulerapi.TaintNodeNotReady: v1.NodeReady, schedulerapi.TaintNodeUnreachable: v1.NodeReady, schedulerapi.TaintNodeNetworkUnavailable: v1.NodeNetworkUnavailable, schedulerapi.TaintNodeMemoryPressure: v1.NodeMemoryPressure, schedulerapi.TaintNodeOutOfDisk: v1.NodeOutOfDisk, schedulerapi.TaintNodeDiskPressure: v1.NodeDiskPressure, schedulerapi.TaintNodePIDPressure: v1.NodePIDPressure}
)

type ZoneState string

const (
 stateInitial           = ZoneState("Initial")
 stateNormal            = ZoneState("Normal")
 stateFullDisruption    = ZoneState("FullDisruption")
 statePartialDisruption = ZoneState("PartialDisruption")
)
const (
 retrySleepTime = 20 * time.Millisecond
)

type nodeHealthData struct {
 probeTimestamp           metav1.Time
 readyTransitionTimestamp metav1.Time
 status                   *v1.NodeStatus
 lease                    *coordv1beta1.Lease
}
type Controller struct {
 taintManager                *scheduler.NoExecuteTaintManager
 podInformerSynced           cache.InformerSynced
 cloud                       cloudprovider.Interface
 kubeClient                  clientset.Interface
 now                         func() metav1.Time
 enterPartialDisruptionFunc  func(nodeNum int) float32
 enterFullDisruptionFunc     func(nodeNum int) float32
 computeZoneStateFunc        func(nodeConditions []*v1.NodeCondition) (int, ZoneState)
 knownNodeSet                map[string]*v1.Node
 nodeHealthMap               map[string]*nodeHealthData
 evictorLock                 sync.Mutex
 zonePodEvictor              map[string]*scheduler.RateLimitedTimedQueue
 zoneNoExecuteTainter        map[string]*scheduler.RateLimitedTimedQueue
 zoneStates                  map[string]ZoneState
 daemonSetStore              extensionslisters.DaemonSetLister
 daemonSetInformerSynced     cache.InformerSynced
 leaseLister                 coordlisters.LeaseLister
 leaseInformerSynced         cache.InformerSynced
 nodeLister                  corelisters.NodeLister
 nodeInformerSynced          cache.InformerSynced
 nodeExistsInCloudProvider   func(types.NodeName) (bool, error)
 nodeShutdownInCloudProvider func(context.Context, *v1.Node) (bool, error)
 recorder                    record.EventRecorder
 nodeMonitorPeriod           time.Duration
 nodeStartupGracePeriod      time.Duration
 nodeMonitorGracePeriod      time.Duration
 podEvictionTimeout          time.Duration
 evictionLimiterQPS          float32
 secondaryEvictionLimiterQPS float32
 largeClusterThreshold       int32
 unhealthyZoneThreshold      float32
 runTaintManager             bool
 useTaintBasedEvictions      bool
 taintNodeByCondition        bool
 nodeUpdateQueue             workqueue.Interface
}

func NewNodeLifecycleController(leaseInformer coordinformers.LeaseInformer, podInformer coreinformers.PodInformer, nodeInformer coreinformers.NodeInformer, daemonSetInformer extensionsinformers.DaemonSetInformer, cloud cloudprovider.Interface, kubeClient clientset.Interface, nodeMonitorPeriod time.Duration, nodeStartupGracePeriod time.Duration, nodeMonitorGracePeriod time.Duration, podEvictionTimeout time.Duration, evictionLimiterQPS float32, secondaryEvictionLimiterQPS float32, largeClusterThreshold int32, unhealthyZoneThreshold float32, runTaintManager bool, useTaintBasedEvictions bool, taintNodeByCondition bool) (*Controller, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if kubeClient == nil {
  klog.Fatalf("kubeClient is nil when starting Controller")
 }
 eventBroadcaster := record.NewBroadcaster()
 recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "node-controller"})
 eventBroadcaster.StartLogging(klog.Infof)
 klog.Infof("Sending events to api server.")
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(kubeClient.CoreV1().RESTClient()).Events("")})
 if kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
  metrics.RegisterMetricAndTrackRateLimiterUsage("node_lifecycle_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
 }
 nc := &Controller{cloud: cloud, kubeClient: kubeClient, now: metav1.Now, knownNodeSet: make(map[string]*v1.Node), nodeHealthMap: make(map[string]*nodeHealthData), nodeExistsInCloudProvider: func(nodeName types.NodeName) (bool, error) {
  return nodeutil.ExistsInCloudProvider(cloud, nodeName)
 }, nodeShutdownInCloudProvider: func(ctx context.Context, node *v1.Node) (bool, error) {
  return nodeutil.ShutdownInCloudProvider(ctx, cloud, node)
 }, recorder: recorder, nodeMonitorPeriod: nodeMonitorPeriod, nodeStartupGracePeriod: nodeStartupGracePeriod, nodeMonitorGracePeriod: nodeMonitorGracePeriod, zonePodEvictor: make(map[string]*scheduler.RateLimitedTimedQueue), zoneNoExecuteTainter: make(map[string]*scheduler.RateLimitedTimedQueue), zoneStates: make(map[string]ZoneState), podEvictionTimeout: podEvictionTimeout, evictionLimiterQPS: evictionLimiterQPS, secondaryEvictionLimiterQPS: secondaryEvictionLimiterQPS, largeClusterThreshold: largeClusterThreshold, unhealthyZoneThreshold: unhealthyZoneThreshold, runTaintManager: runTaintManager, useTaintBasedEvictions: useTaintBasedEvictions && runTaintManager, taintNodeByCondition: taintNodeByCondition, nodeUpdateQueue: workqueue.New()}
 if useTaintBasedEvictions {
  klog.Infof("Controller is using taint based evictions.")
 }
 nc.enterPartialDisruptionFunc = nc.ReducedQPSFunc
 nc.enterFullDisruptionFunc = nc.HealthyQPSFunc
 nc.computeZoneStateFunc = nc.ComputeZoneState
 podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  pod := obj.(*v1.Pod)
  if nc.taintManager != nil {
   nc.taintManager.PodUpdated(nil, pod)
  }
 }, UpdateFunc: func(prev, obj interface{}) {
  prevPod := prev.(*v1.Pod)
  newPod := obj.(*v1.Pod)
  if nc.taintManager != nil {
   nc.taintManager.PodUpdated(prevPod, newPod)
  }
 }, DeleteFunc: func(obj interface{}) {
  pod, isPod := obj.(*v1.Pod)
  if !isPod {
   deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
   if !ok {
    klog.Errorf("Received unexpected object: %v", obj)
    return
   }
   pod, ok = deletedState.Obj.(*v1.Pod)
   if !ok {
    klog.Errorf("DeletedFinalStateUnknown contained non-Pod object: %v", deletedState.Obj)
    return
   }
  }
  if nc.taintManager != nil {
   nc.taintManager.PodUpdated(pod, nil)
  }
 }})
 nc.podInformerSynced = podInformer.Informer().HasSynced
 if nc.runTaintManager {
  podLister := podInformer.Lister()
  podGetter := func(name, namespace string) (*v1.Pod, error) {
   return podLister.Pods(namespace).Get(name)
  }
  nodeLister := nodeInformer.Lister()
  nodeGetter := func(name string) (*v1.Node, error) {
   return nodeLister.Get(name)
  }
  nc.taintManager = scheduler.NewNoExecuteTaintManager(kubeClient, podGetter, nodeGetter)
  nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(func(node *v1.Node) error {
   nc.taintManager.NodeUpdated(nil, node)
   return nil
  }), UpdateFunc: nodeutil.CreateUpdateNodeHandler(func(oldNode, newNode *v1.Node) error {
   nc.taintManager.NodeUpdated(oldNode, newNode)
   return nil
  }), DeleteFunc: nodeutil.CreateDeleteNodeHandler(func(node *v1.Node) error {
   nc.taintManager.NodeUpdated(node, nil)
   return nil
  })})
 }
 if nc.taintNodeByCondition {
  klog.Infof("Controller will taint node by condition.")
  nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(func(node *v1.Node) error {
   nc.nodeUpdateQueue.Add(node.Name)
   return nil
  }), UpdateFunc: nodeutil.CreateUpdateNodeHandler(func(_, newNode *v1.Node) error {
   nc.nodeUpdateQueue.Add(newNode.Name)
   return nil
  })})
 }
 nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(func(node *v1.Node) error {
  return nc.doFixDeprecatedTaintKeyPass(node)
 }), UpdateFunc: nodeutil.CreateUpdateNodeHandler(func(_, newNode *v1.Node) error {
  return nc.doFixDeprecatedTaintKeyPass(newNode)
 })})
 nc.leaseLister = leaseInformer.Lister()
 if utilfeature.DefaultFeatureGate.Enabled(features.NodeLease) {
  nc.leaseInformerSynced = leaseInformer.Informer().HasSynced
 } else {
  nc.leaseInformerSynced = func() bool {
   return true
  }
 }
 nc.nodeLister = nodeInformer.Lister()
 nc.nodeInformerSynced = nodeInformer.Informer().HasSynced
 nc.daemonSetStore = daemonSetInformer.Lister()
 nc.daemonSetInformerSynced = daemonSetInformer.Informer().HasSynced
 return nc, nil
}
func (nc *Controller) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting node controller")
 defer klog.Infof("Shutting down node controller")
 if !controller.WaitForCacheSync("taint", stopCh, nc.leaseInformerSynced, nc.nodeInformerSynced, nc.podInformerSynced, nc.daemonSetInformerSynced) {
  return
 }
 if nc.runTaintManager {
  go nc.taintManager.Run(stopCh)
 }
 if nc.taintNodeByCondition {
  defer nc.nodeUpdateQueue.ShutDown()
  for i := 0; i < scheduler.UpdateWorkerSize; i++ {
   go wait.Until(nc.doNoScheduleTaintingPassWorker, time.Second, stopCh)
  }
 }
 if nc.useTaintBasedEvictions {
  go wait.Until(nc.doNoExecuteTaintingPass, scheduler.NodeEvictionPeriod, stopCh)
 } else {
  go wait.Until(nc.doEvictionPass, scheduler.NodeEvictionPeriod, stopCh)
 }
 go wait.Until(func() {
  if err := nc.monitorNodeHealth(); err != nil {
   klog.Errorf("Error monitoring node health: %v", err)
  }
 }, nc.nodeMonitorPeriod, stopCh)
 <-stopCh
}
func (nc *Controller) doFixDeprecatedTaintKeyPass(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 taintsToAdd := []*v1.Taint{}
 taintsToDel := []*v1.Taint{}
 for _, taint := range node.Spec.Taints {
  if taint.Key == schedulerapi.DeprecatedTaintNodeNotReady {
   tDel := taint
   taintsToDel = append(taintsToDel, &tDel)
   tAdd := taint
   tAdd.Key = schedulerapi.TaintNodeNotReady
   taintsToAdd = append(taintsToAdd, &tAdd)
  }
  if taint.Key == schedulerapi.DeprecatedTaintNodeUnreachable {
   tDel := taint
   taintsToDel = append(taintsToDel, &tDel)
   tAdd := taint
   tAdd.Key = schedulerapi.TaintNodeUnreachable
   taintsToAdd = append(taintsToAdd, &tAdd)
  }
 }
 if len(taintsToAdd) == 0 && len(taintsToDel) == 0 {
  return nil
 }
 klog.Warningf("Detected deprecated taint keys: %v on node: %v, will substitute them with %v", taintsToDel, node.GetName(), taintsToAdd)
 if !nodeutil.SwapNodeControllerTaint(nc.kubeClient, taintsToAdd, taintsToDel, node) {
  return fmt.Errorf("failed to swap taints of node %+v", node)
 }
 return nil
}
func (nc *Controller) doNoScheduleTaintingPassWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for {
  obj, shutdown := nc.nodeUpdateQueue.Get()
  if shutdown {
   return
  }
  nodeName := obj.(string)
  if err := nc.doNoScheduleTaintingPass(nodeName); err != nil {
   klog.Errorf("Failed to taint NoSchedule on node <%s>, requeue it: %v", nodeName, err)
  }
  nc.nodeUpdateQueue.Done(nodeName)
 }
}
func (nc *Controller) doNoScheduleTaintingPass(nodeName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, err := nc.nodeLister.Get(nodeName)
 if err != nil {
  if apierrors.IsNotFound(err) {
   return nil
  }
  return err
 }
 var taints []v1.Taint
 for _, condition := range node.Status.Conditions {
  if taintMap, found := nodeConditionToTaintKeyStatusMap[condition.Type]; found {
   if taintKey, found := taintMap[condition.Status]; found {
    taints = append(taints, v1.Taint{Key: taintKey, Effect: v1.TaintEffectNoSchedule})
   }
  }
 }
 if node.Spec.Unschedulable {
  taints = append(taints, v1.Taint{Key: schedulerapi.TaintNodeUnschedulable, Effect: v1.TaintEffectNoSchedule})
 }
 nodeTaints := taintutils.TaintSetFilter(node.Spec.Taints, func(t *v1.Taint) bool {
  if t.Effect != v1.TaintEffectNoSchedule {
   return false
  }
  if t.Key == schedulerapi.TaintNodeUnschedulable {
   return true
  }
  _, found := taintKeyToNodeConditionMap[t.Key]
  return found
 })
 taintsToAdd, taintsToDel := taintutils.TaintSetDiff(taints, nodeTaints)
 if len(taintsToAdd) == 0 && len(taintsToDel) == 0 {
  return nil
 }
 if !nodeutil.SwapNodeControllerTaint(nc.kubeClient, taintsToAdd, taintsToDel, node) {
  return fmt.Errorf("failed to swap taints of node %+v", node)
 }
 return nil
}
func (nc *Controller) doNoExecuteTaintingPass() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 for k := range nc.zoneNoExecuteTainter {
  nc.zoneNoExecuteTainter[k].Try(func(value scheduler.TimedValue) (bool, time.Duration) {
   node, err := nc.nodeLister.Get(value.Value)
   if apierrors.IsNotFound(err) {
    klog.Warningf("Node %v no longer present in nodeLister!", value.Value)
    return true, 0
   } else if err != nil {
    klog.Warningf("Failed to get Node %v from the nodeLister: %v", value.Value, err)
    return false, 50 * time.Millisecond
   } else {
    zone := utilnode.GetZoneKey(node)
    evictionsNumber.WithLabelValues(zone).Inc()
   }
   _, condition := v1node.GetNodeCondition(&node.Status, v1.NodeReady)
   taintToAdd := v1.Taint{}
   oppositeTaint := v1.Taint{}
   if condition.Status == v1.ConditionFalse {
    taintToAdd = *NotReadyTaintTemplate
    oppositeTaint = *UnreachableTaintTemplate
   } else if condition.Status == v1.ConditionUnknown {
    taintToAdd = *UnreachableTaintTemplate
    oppositeTaint = *NotReadyTaintTemplate
   } else {
    klog.V(4).Infof("Node %v was in a taint queue, but it's ready now. Ignoring taint request.", value.Value)
    return true, 0
   }
   return nodeutil.SwapNodeControllerTaint(nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{&oppositeTaint}, node), 0
  })
 }
}
func (nc *Controller) doEvictionPass() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 for k := range nc.zonePodEvictor {
  nc.zonePodEvictor[k].Try(func(value scheduler.TimedValue) (bool, time.Duration) {
   node, err := nc.nodeLister.Get(value.Value)
   if apierrors.IsNotFound(err) {
    klog.Warningf("Node %v no longer present in nodeLister!", value.Value)
   } else if err != nil {
    klog.Warningf("Failed to get Node %v from the nodeLister: %v", value.Value, err)
   } else {
    zone := utilnode.GetZoneKey(node)
    evictionsNumber.WithLabelValues(zone).Inc()
   }
   nodeUID, _ := value.UID.(string)
   remaining, err := nodeutil.DeletePods(nc.kubeClient, nc.recorder, value.Value, nodeUID, nc.daemonSetStore)
   if err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to evict node %q: %v", value.Value, err))
    return false, 0
   }
   if remaining {
    klog.Infof("Pods awaiting deletion due to Controller eviction")
   }
   return true, 0
  })
 }
}
func (nc *Controller) monitorNodeHealth() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodes, err := nc.nodeLister.List(labels.Everything())
 if err != nil {
  return err
 }
 added, deleted, newZoneRepresentatives := nc.classifyNodes(nodes)
 for i := range newZoneRepresentatives {
  nc.addPodEvictorForNewZone(newZoneRepresentatives[i])
 }
 for i := range added {
  klog.V(1).Infof("Controller observed a new Node: %#v", added[i].Name)
  nodeutil.RecordNodeEvent(nc.recorder, added[i].Name, string(added[i].UID), v1.EventTypeNormal, "RegisteredNode", fmt.Sprintf("Registered Node %v in Controller", added[i].Name))
  nc.knownNodeSet[added[i].Name] = added[i]
  nc.addPodEvictorForNewZone(added[i])
  if nc.useTaintBasedEvictions {
   nc.markNodeAsReachable(added[i])
  } else {
   nc.cancelPodEviction(added[i])
  }
 }
 for i := range deleted {
  klog.V(1).Infof("Controller observed a Node deletion: %v", deleted[i].Name)
  nodeutil.RecordNodeEvent(nc.recorder, deleted[i].Name, string(deleted[i].UID), v1.EventTypeNormal, "RemovingNode", fmt.Sprintf("Removing Node %v from Controller", deleted[i].Name))
  delete(nc.knownNodeSet, deleted[i].Name)
 }
 zoneToNodeConditions := map[string][]*v1.NodeCondition{}
 for i := range nodes {
  var gracePeriod time.Duration
  var observedReadyCondition v1.NodeCondition
  var currentReadyCondition *v1.NodeCondition
  node := nodes[i].DeepCopy()
  if err := wait.PollImmediate(retrySleepTime, retrySleepTime*scheduler.NodeHealthUpdateRetry, func() (bool, error) {
   gracePeriod, observedReadyCondition, currentReadyCondition, err = nc.tryUpdateNodeHealth(node)
   if err == nil {
    return true, nil
   }
   name := node.Name
   node, err = nc.kubeClient.CoreV1().Nodes().Get(name, metav1.GetOptions{})
   if err != nil {
    klog.Errorf("Failed while getting a Node to retry updating node health. Probably Node %s was deleted.", name)
    return false, err
   }
   return false, nil
  }); err != nil {
   klog.Errorf("Update health of Node '%v' from Controller error: %v. "+"Skipping - no pods will be evicted.", node.Name, err)
   continue
  }
  if !system.IsMasterNode(node.Name) {
   zoneToNodeConditions[utilnode.GetZoneKey(node)] = append(zoneToNodeConditions[utilnode.GetZoneKey(node)], currentReadyCondition)
  }
  decisionTimestamp := nc.now()
  if currentReadyCondition != nil {
   if observedReadyCondition.Status == v1.ConditionFalse {
    if nc.useTaintBasedEvictions {
     if taintutils.TaintExists(node.Spec.Taints, UnreachableTaintTemplate) {
      taintToAdd := *NotReadyTaintTemplate
      if !nodeutil.SwapNodeControllerTaint(nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{UnreachableTaintTemplate}, node) {
       klog.Errorf("Failed to instantly swap UnreachableTaint to NotReadyTaint. Will try again in the next cycle.")
      }
     } else if nc.markNodeForTainting(node) {
      klog.V(2).Infof("Node %v is NotReady as of %v. Adding it to the Taint queue.", node.Name, decisionTimestamp)
     }
    } else {
     if decisionTimestamp.After(nc.nodeHealthMap[node.Name].readyTransitionTimestamp.Add(nc.podEvictionTimeout)) {
      if nc.evictPods(node) {
       klog.V(2).Infof("Node is NotReady. Adding Pods on Node %s to eviction queue: %v is later than %v + %v", node.Name, decisionTimestamp, nc.nodeHealthMap[node.Name].readyTransitionTimestamp, nc.podEvictionTimeout)
      }
     }
    }
   }
   if observedReadyCondition.Status == v1.ConditionUnknown {
    if nc.useTaintBasedEvictions {
     if taintutils.TaintExists(node.Spec.Taints, NotReadyTaintTemplate) {
      taintToAdd := *UnreachableTaintTemplate
      if !nodeutil.SwapNodeControllerTaint(nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{NotReadyTaintTemplate}, node) {
       klog.Errorf("Failed to instantly swap UnreachableTaint to NotReadyTaint. Will try again in the next cycle.")
      }
     } else if nc.markNodeForTainting(node) {
      klog.V(2).Infof("Node %v is unresponsive as of %v. Adding it to the Taint queue.", node.Name, decisionTimestamp)
     }
    } else {
     if decisionTimestamp.After(nc.nodeHealthMap[node.Name].probeTimestamp.Add(nc.podEvictionTimeout)) {
      if nc.evictPods(node) {
       klog.V(2).Infof("Node is unresponsive. Adding Pods on Node %s to eviction queues: %v is later than %v + %v", node.Name, decisionTimestamp, nc.nodeHealthMap[node.Name].readyTransitionTimestamp, nc.podEvictionTimeout-gracePeriod)
      }
     }
    }
   }
   if observedReadyCondition.Status == v1.ConditionTrue {
    if nc.useTaintBasedEvictions {
     removed, err := nc.markNodeAsReachable(node)
     if err != nil {
      klog.Errorf("Failed to remove taints from node %v. Will retry in next iteration.", node.Name)
     }
     if removed {
      klog.V(2).Infof("Node %s is healthy again, removing all taints", node.Name)
     }
    } else {
     if nc.cancelPodEviction(node) {
      klog.V(2).Infof("Node %s is ready again, cancelled pod eviction", node.Name)
     }
    }
    err := nc.markNodeAsNotShutdown(node)
    if err != nil {
     klog.Errorf("Failed to remove taints from node %v. Will retry in next iteration.", node.Name)
    }
   }
   if currentReadyCondition.Status != v1.ConditionTrue && observedReadyCondition.Status == v1.ConditionTrue {
    nodeutil.RecordNodeStatusChange(nc.recorder, node, "NodeNotReady")
    if err = nodeutil.MarkAllPodsNotReady(nc.kubeClient, node); err != nil {
     utilruntime.HandleError(fmt.Errorf("Unable to mark all pods NotReady on node %v: %v", node.Name, err))
    }
   }
   if currentReadyCondition.Status != v1.ConditionTrue && nc.cloud != nil {
    shutdown, err := nc.nodeShutdownInCloudProvider(context.TODO(), node)
    if err != nil {
     klog.Errorf("Error determining if node %v shutdown in cloud: %v", node.Name, err)
    }
    if shutdown && err == nil {
     err = controller.AddOrUpdateTaintOnNode(nc.kubeClient, node.Name, controller.ShutdownTaint)
     if err != nil {
      klog.Errorf("Error patching node taints: %v", err)
     }
     continue
    }
    exists, err := nc.nodeExistsInCloudProvider(types.NodeName(node.Name))
    if err != nil {
     klog.Errorf("Error determining if node %v exists in cloud: %v", node.Name, err)
     continue
    }
    if !exists {
     klog.V(2).Infof("Deleting node (no longer present in cloud provider): %s", node.Name)
     nodeutil.RecordNodeEvent(nc.recorder, node.Name, string(node.UID), v1.EventTypeNormal, "DeletingNode", fmt.Sprintf("Deleting Node %v because it's not present according to cloud provider", node.Name))
     go func(nodeName string) {
      defer utilruntime.HandleCrash()
      if err := nodeutil.ForcefullyDeleteNode(nc.kubeClient, nodeName); err != nil {
       klog.Errorf("Unable to forcefully delete node %q: %v", nodeName, err)
      }
     }(node.Name)
    }
   }
  }
 }
 nc.handleDisruption(zoneToNodeConditions, nodes)
 return nil
}
func (nc *Controller) tryUpdateNodeHealth(node *v1.Node) (time.Duration, v1.NodeCondition, *v1.NodeCondition, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 var gracePeriod time.Duration
 var observedReadyCondition v1.NodeCondition
 _, currentReadyCondition := v1node.GetNodeCondition(&node.Status, v1.NodeReady)
 if currentReadyCondition == nil {
  observedReadyCondition = v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionUnknown, LastHeartbeatTime: node.CreationTimestamp, LastTransitionTime: node.CreationTimestamp}
  gracePeriod = nc.nodeStartupGracePeriod
  if _, found := nc.nodeHealthMap[node.Name]; found {
   nc.nodeHealthMap[node.Name].status = &node.Status
  } else {
   nc.nodeHealthMap[node.Name] = &nodeHealthData{status: &node.Status, probeTimestamp: node.CreationTimestamp, readyTransitionTimestamp: node.CreationTimestamp}
  }
 } else {
  observedReadyCondition = *currentReadyCondition
  gracePeriod = nc.nodeMonitorGracePeriod
 }
 savedNodeHealth, found := nc.nodeHealthMap[node.Name]
 var savedCondition *v1.NodeCondition
 var savedLease *coordv1beta1.Lease
 if found {
  _, savedCondition = v1node.GetNodeCondition(savedNodeHealth.status, v1.NodeReady)
  savedLease = savedNodeHealth.lease
 }
 _, observedCondition := v1node.GetNodeCondition(&node.Status, v1.NodeReady)
 if !found {
  klog.Warningf("Missing timestamp for Node %s. Assuming now as a timestamp.", node.Name)
  savedNodeHealth = &nodeHealthData{status: &node.Status, probeTimestamp: nc.now(), readyTransitionTimestamp: nc.now()}
 } else if savedCondition == nil && observedCondition != nil {
  klog.V(1).Infof("Creating timestamp entry for newly observed Node %s", node.Name)
  savedNodeHealth = &nodeHealthData{status: &node.Status, probeTimestamp: nc.now(), readyTransitionTimestamp: nc.now()}
 } else if savedCondition != nil && observedCondition == nil {
  klog.Errorf("ReadyCondition was removed from Status of Node %s", node.Name)
  savedNodeHealth = &nodeHealthData{status: &node.Status, probeTimestamp: nc.now(), readyTransitionTimestamp: nc.now()}
 } else if savedCondition != nil && observedCondition != nil && savedCondition.LastHeartbeatTime != observedCondition.LastHeartbeatTime {
  var transitionTime metav1.Time
  if savedCondition.LastTransitionTime != observedCondition.LastTransitionTime {
   klog.V(3).Infof("ReadyCondition for Node %s transitioned from %v to %v", node.Name, savedCondition, observedCondition)
   transitionTime = nc.now()
  } else {
   transitionTime = savedNodeHealth.readyTransitionTimestamp
  }
  if klog.V(5) {
   klog.V(5).Infof("Node %s ReadyCondition updated. Updating timestamp: %+v vs %+v.", node.Name, savedNodeHealth.status, node.Status)
  } else {
   klog.V(3).Infof("Node %s ReadyCondition updated. Updating timestamp.", node.Name)
  }
  savedNodeHealth = &nodeHealthData{status: &node.Status, probeTimestamp: nc.now(), readyTransitionTimestamp: transitionTime}
 }
 var observedLease *coordv1beta1.Lease
 if utilfeature.DefaultFeatureGate.Enabled(features.NodeLease) {
  observedLease, _ = nc.leaseLister.Leases(v1.NamespaceNodeLease).Get(node.Name)
  if observedLease != nil && (savedLease == nil || savedLease.Spec.RenewTime.Before(observedLease.Spec.RenewTime)) {
   savedNodeHealth.lease = observedLease
   savedNodeHealth.probeTimestamp = nc.now()
  }
 }
 nc.nodeHealthMap[node.Name] = savedNodeHealth
 if nc.now().After(savedNodeHealth.probeTimestamp.Add(gracePeriod)) {
  if currentReadyCondition == nil {
   klog.V(2).Infof("node %v is never updated by kubelet", node.Name)
   node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionUnknown, Reason: "NodeStatusNeverUpdated", Message: fmt.Sprintf("Kubelet never posted node status."), LastHeartbeatTime: node.CreationTimestamp, LastTransitionTime: nc.now()})
  } else {
   klog.V(4).Infof("node %v hasn't been updated for %+v. Last ready condition is: %+v", node.Name, nc.now().Time.Sub(savedNodeHealth.probeTimestamp.Time), observedReadyCondition)
   if observedReadyCondition.Status != v1.ConditionUnknown {
    currentReadyCondition.Status = v1.ConditionUnknown
    currentReadyCondition.Reason = "NodeStatusUnknown"
    currentReadyCondition.Message = "Kubelet stopped posting node status."
    currentReadyCondition.LastHeartbeatTime = observedReadyCondition.LastHeartbeatTime
    currentReadyCondition.LastTransitionTime = nc.now()
   }
  }
  remainingNodeConditionTypes := []v1.NodeConditionType{v1.NodeOutOfDisk, v1.NodeMemoryPressure, v1.NodeDiskPressure, v1.NodePIDPressure}
  nowTimestamp := nc.now()
  for _, nodeConditionType := range remainingNodeConditionTypes {
   _, currentCondition := v1node.GetNodeCondition(&node.Status, nodeConditionType)
   if currentCondition == nil {
    klog.V(2).Infof("Condition %v of node %v was never updated by kubelet", nodeConditionType, node.Name)
    node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: nodeConditionType, Status: v1.ConditionUnknown, Reason: "NodeStatusNeverUpdated", Message: "Kubelet never posted node status.", LastHeartbeatTime: node.CreationTimestamp, LastTransitionTime: nowTimestamp})
   } else {
    klog.V(4).Infof("node %v hasn't been updated for %+v. Last %v is: %+v", node.Name, nc.now().Time.Sub(savedNodeHealth.probeTimestamp.Time), nodeConditionType, currentCondition)
    if currentCondition.Status != v1.ConditionUnknown {
     currentCondition.Status = v1.ConditionUnknown
     currentCondition.Reason = "NodeStatusUnknown"
     currentCondition.Message = "Kubelet stopped posting node status."
     currentCondition.LastTransitionTime = nowTimestamp
    }
   }
  }
  _, currentCondition := v1node.GetNodeCondition(&node.Status, v1.NodeReady)
  if !apiequality.Semantic.DeepEqual(currentCondition, &observedReadyCondition) {
   if _, err = nc.kubeClient.CoreV1().Nodes().UpdateStatus(node); err != nil {
    klog.Errorf("Error updating node %s: %v", node.Name, err)
    return gracePeriod, observedReadyCondition, currentReadyCondition, err
   }
   nc.nodeHealthMap[node.Name] = &nodeHealthData{status: &node.Status, probeTimestamp: nc.nodeHealthMap[node.Name].probeTimestamp, readyTransitionTimestamp: nc.now(), lease: observedLease}
   return gracePeriod, observedReadyCondition, currentReadyCondition, nil
  }
 }
 return gracePeriod, observedReadyCondition, currentReadyCondition, err
}
func (nc *Controller) handleDisruption(zoneToNodeConditions map[string][]*v1.NodeCondition, nodes []*v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newZoneStates := map[string]ZoneState{}
 allAreFullyDisrupted := true
 for k, v := range zoneToNodeConditions {
  zoneSize.WithLabelValues(k).Set(float64(len(v)))
  unhealthy, newState := nc.computeZoneStateFunc(v)
  zoneHealth.WithLabelValues(k).Set(float64(100*(len(v)-unhealthy)) / float64(len(v)))
  unhealthyNodes.WithLabelValues(k).Set(float64(unhealthy))
  if newState != stateFullDisruption {
   allAreFullyDisrupted = false
  }
  newZoneStates[k] = newState
  if _, had := nc.zoneStates[k]; !had {
   klog.Errorf("Setting initial state for unseen zone: %v", k)
   nc.zoneStates[k] = stateInitial
  }
 }
 allWasFullyDisrupted := true
 for k, v := range nc.zoneStates {
  if _, have := zoneToNodeConditions[k]; !have {
   zoneSize.WithLabelValues(k).Set(0)
   zoneHealth.WithLabelValues(k).Set(100)
   unhealthyNodes.WithLabelValues(k).Set(0)
   delete(nc.zoneStates, k)
   continue
  }
  if v != stateFullDisruption {
   allWasFullyDisrupted = false
   break
  }
 }
 if !allAreFullyDisrupted || !allWasFullyDisrupted {
  if allAreFullyDisrupted {
   klog.V(0).Info("Controller detected that all Nodes are not-Ready. Entering master disruption mode.")
   for i := range nodes {
    if nc.useTaintBasedEvictions {
     _, err := nc.markNodeAsReachable(nodes[i])
     if err != nil {
      klog.Errorf("Failed to remove taints from Node %v", nodes[i].Name)
     }
    } else {
     nc.cancelPodEviction(nodes[i])
    }
   }
   for k := range nc.zoneStates {
    if nc.useTaintBasedEvictions {
     nc.zoneNoExecuteTainter[k].SwapLimiter(0)
    } else {
     nc.zonePodEvictor[k].SwapLimiter(0)
    }
   }
   for k := range nc.zoneStates {
    nc.zoneStates[k] = stateFullDisruption
   }
   return
  }
  if allWasFullyDisrupted {
   klog.V(0).Info("Controller detected that some Nodes are Ready. Exiting master disruption mode.")
   now := nc.now()
   for i := range nodes {
    v := nc.nodeHealthMap[nodes[i].Name]
    v.probeTimestamp = now
    v.readyTransitionTimestamp = now
    nc.nodeHealthMap[nodes[i].Name] = v
   }
   for k := range nc.zoneStates {
    nc.setLimiterInZone(k, len(zoneToNodeConditions[k]), newZoneStates[k])
    nc.zoneStates[k] = newZoneStates[k]
   }
   return
  }
  for k, v := range nc.zoneStates {
   newState := newZoneStates[k]
   if v == newState {
    continue
   }
   klog.V(0).Infof("Controller detected that zone %v is now in state %v.", k, newState)
   nc.setLimiterInZone(k, len(zoneToNodeConditions[k]), newState)
   nc.zoneStates[k] = newState
  }
 }
}
func (nc *Controller) setLimiterInZone(zone string, zoneSize int, state ZoneState) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch state {
 case stateNormal:
  if nc.useTaintBasedEvictions {
   nc.zoneNoExecuteTainter[zone].SwapLimiter(nc.evictionLimiterQPS)
  } else {
   nc.zonePodEvictor[zone].SwapLimiter(nc.evictionLimiterQPS)
  }
 case statePartialDisruption:
  if nc.useTaintBasedEvictions {
   nc.zoneNoExecuteTainter[zone].SwapLimiter(nc.enterPartialDisruptionFunc(zoneSize))
  } else {
   nc.zonePodEvictor[zone].SwapLimiter(nc.enterPartialDisruptionFunc(zoneSize))
  }
 case stateFullDisruption:
  if nc.useTaintBasedEvictions {
   nc.zoneNoExecuteTainter[zone].SwapLimiter(nc.enterFullDisruptionFunc(zoneSize))
  } else {
   nc.zonePodEvictor[zone].SwapLimiter(nc.enterFullDisruptionFunc(zoneSize))
  }
 }
}
func (nc *Controller) classifyNodes(allNodes []*v1.Node) (added, deleted, newZoneRepresentatives []*v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range allNodes {
  if _, has := nc.knownNodeSet[allNodes[i].Name]; !has {
   added = append(added, allNodes[i])
  } else {
   zone := utilnode.GetZoneKey(allNodes[i])
   if _, found := nc.zoneStates[zone]; !found {
    newZoneRepresentatives = append(newZoneRepresentatives, allNodes[i])
   }
  }
 }
 if len(nc.knownNodeSet)+len(added) != len(allNodes) {
  knowSetCopy := map[string]*v1.Node{}
  for k, v := range nc.knownNodeSet {
   knowSetCopy[k] = v
  }
  for i := range allNodes {
   delete(knowSetCopy, allNodes[i].Name)
  }
  for i := range knowSetCopy {
   deleted = append(deleted, knowSetCopy[i])
  }
 }
 return
}
func (nc *Controller) HealthyQPSFunc(nodeNum int) float32 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nc.evictionLimiterQPS
}
func (nc *Controller) ReducedQPSFunc(nodeNum int) float32 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if int32(nodeNum) > nc.largeClusterThreshold {
  return nc.secondaryEvictionLimiterQPS
 }
 return 0
}
func (nc *Controller) addPodEvictorForNewZone(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 zone := utilnode.GetZoneKey(node)
 if _, found := nc.zoneStates[zone]; !found {
  nc.zoneStates[zone] = stateInitial
  if !nc.useTaintBasedEvictions {
   nc.zonePodEvictor[zone] = scheduler.NewRateLimitedTimedQueue(flowcontrol.NewTokenBucketRateLimiter(nc.evictionLimiterQPS, scheduler.EvictionRateLimiterBurst))
  } else {
   nc.zoneNoExecuteTainter[zone] = scheduler.NewRateLimitedTimedQueue(flowcontrol.NewTokenBucketRateLimiter(nc.evictionLimiterQPS, scheduler.EvictionRateLimiterBurst))
  }
  klog.Infof("Initializing eviction metric for zone: %v", zone)
  evictionsNumber.WithLabelValues(zone).Add(0)
 }
}
func (nc *Controller) cancelPodEviction(node *v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zone := utilnode.GetZoneKey(node)
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 wasDeleting := nc.zonePodEvictor[zone].Remove(node.Name)
 if wasDeleting {
  klog.V(2).Infof("Cancelling pod Eviction on Node: %v", node.Name)
  return true
 }
 return false
}
func (nc *Controller) evictPods(node *v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 return nc.zonePodEvictor[utilnode.GetZoneKey(node)].Add(node.Name, string(node.UID))
}
func (nc *Controller) markNodeForTainting(node *v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 return nc.zoneNoExecuteTainter[utilnode.GetZoneKey(node)].Add(node.Name, string(node.UID))
}
func (nc *Controller) markNodeAsReachable(node *v1.Node) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 err := controller.RemoveTaintOffNode(nc.kubeClient, node.Name, node, UnreachableTaintTemplate)
 if err != nil {
  klog.Errorf("Failed to remove taint from node %v: %v", node.Name, err)
  return false, err
 }
 err = controller.RemoveTaintOffNode(nc.kubeClient, node.Name, node, NotReadyTaintTemplate)
 if err != nil {
  klog.Errorf("Failed to remove taint from node %v: %v", node.Name, err)
  return false, err
 }
 return nc.zoneNoExecuteTainter[utilnode.GetZoneKey(node)].Remove(node.Name), nil
}
func (nc *Controller) markNodeAsNotShutdown(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nc.evictorLock.Lock()
 defer nc.evictorLock.Unlock()
 err := controller.RemoveTaintOffNode(nc.kubeClient, node.Name, node, controller.ShutdownTaint)
 if err != nil {
  klog.Errorf("Failed to remove taint from node %v: %v", node.Name, err)
  return err
 }
 return nil
}
func (nc *Controller) ComputeZoneState(nodeReadyConditions []*v1.NodeCondition) (int, ZoneState) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 readyNodes := 0
 notReadyNodes := 0
 for i := range nodeReadyConditions {
  if nodeReadyConditions[i] != nil && nodeReadyConditions[i].Status == v1.ConditionTrue {
   readyNodes++
  } else {
   notReadyNodes++
  }
 }
 switch {
 case readyNodes == 0 && notReadyNodes > 0:
  return notReadyNodes, stateFullDisruption
 case notReadyNodes > 2 && float32(notReadyNodes)/float32(notReadyNodes+readyNodes) >= nc.unhealthyZoneThreshold:
  return notReadyNodes, statePartialDisruption
 default:
  return notReadyNodes, stateNormal
 }
}
func hash(val string, max int) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hasher := fnv.New32a()
 io.WriteString(hasher, val)
 return int(hasher.Sum32()) % max
}
