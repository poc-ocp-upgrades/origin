package scheduler

import (
 "fmt"
 "hash/fnv"
 "io"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/apis/core/helper"
 v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/klog"
)

const (
 NodeUpdateChannelSize = 10
 UpdateWorkerSize      = 8
 podUpdateChannelSize  = 1
 retries               = 5
)

type nodeUpdateItem struct{ nodeName string }
type podUpdateItem struct {
 podName      string
 podNamespace string
 nodeName     string
}

func hash(val string, max int) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hasher := fnv.New32a()
 io.WriteString(hasher, val)
 return int(hasher.Sum32() % uint32(max))
}

type GetPodFunc func(name, namespace string) (*v1.Pod, error)
type GetNodeFunc func(name string) (*v1.Node, error)
type NoExecuteTaintManager struct {
 client             clientset.Interface
 recorder           record.EventRecorder
 getPod             GetPodFunc
 getNode            GetNodeFunc
 taintEvictionQueue *TimedWorkerQueue
 taintedNodesLock   sync.Mutex
 taintedNodes       map[string][]v1.Taint
 nodeUpdateChannels []chan nodeUpdateItem
 podUpdateChannels  []chan podUpdateItem
 nodeUpdateQueue    workqueue.Interface
 podUpdateQueue     workqueue.Interface
}

func deletePodHandler(c clientset.Interface, emitEventFunc func(types.NamespacedName)) func(args *WorkArgs) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(args *WorkArgs) error {
  ns := args.NamespacedName.Namespace
  name := args.NamespacedName.Name
  klog.V(0).Infof("NoExecuteTaintManager is deleting Pod: %v", args.NamespacedName.String())
  if emitEventFunc != nil {
   emitEventFunc(args.NamespacedName)
  }
  var err error
  for i := 0; i < retries; i++ {
   err = c.CoreV1().Pods(ns).Delete(name, &metav1.DeleteOptions{})
   if err == nil {
    break
   }
   time.Sleep(10 * time.Millisecond)
  }
  return err
 }
}
func getNoExecuteTaints(taints []v1.Taint) []v1.Taint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := []v1.Taint{}
 for i := range taints {
  if taints[i].Effect == v1.TaintEffectNoExecute {
   result = append(result, taints[i])
  }
 }
 return result
}
func getPodsAssignedToNode(c clientset.Interface, nodeName string) ([]v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector := fields.SelectorFromSet(fields.Set{"spec.nodeName": nodeName})
 pods, err := c.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{FieldSelector: selector.String(), LabelSelector: labels.Everything().String()})
 for i := 0; i < retries && err != nil; i++ {
  pods, err = c.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{FieldSelector: selector.String(), LabelSelector: labels.Everything().String()})
  time.Sleep(100 * time.Millisecond)
 }
 if err != nil {
  return []v1.Pod{}, fmt.Errorf("failed to get Pods assigned to node %v", nodeName)
 }
 return pods.Items, nil
}
func getMinTolerationTime(tolerations []v1.Toleration) time.Duration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 minTolerationTime := int64(-1)
 if len(tolerations) == 0 {
  return 0
 }
 for i := range tolerations {
  if tolerations[i].TolerationSeconds != nil {
   tolerationSeconds := *(tolerations[i].TolerationSeconds)
   if tolerationSeconds <= 0 {
    return 0
   } else if tolerationSeconds < minTolerationTime || minTolerationTime == -1 {
    minTolerationTime = tolerationSeconds
   }
  }
 }
 return time.Duration(minTolerationTime) * time.Second
}
func NewNoExecuteTaintManager(c clientset.Interface, getPod GetPodFunc, getNode GetNodeFunc) *NoExecuteTaintManager {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "taint-controller"})
 eventBroadcaster.StartLogging(klog.Infof)
 if c != nil {
  klog.V(0).Infof("Sending events to api server.")
  eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: c.CoreV1().Events("")})
 } else {
  klog.Fatalf("kubeClient is nil when starting NodeController")
 }
 tm := &NoExecuteTaintManager{client: c, recorder: recorder, getPod: getPod, getNode: getNode, taintedNodes: make(map[string][]v1.Taint), nodeUpdateQueue: workqueue.New(), podUpdateQueue: workqueue.New()}
 tm.taintEvictionQueue = CreateWorkerQueue(deletePodHandler(c, tm.emitPodDeletionEvent))
 return tm
}
func (tc *NoExecuteTaintManager) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(0).Infof("Starting NoExecuteTaintManager")
 for i := 0; i < UpdateWorkerSize; i++ {
  tc.nodeUpdateChannels = append(tc.nodeUpdateChannels, make(chan nodeUpdateItem, NodeUpdateChannelSize))
  tc.podUpdateChannels = append(tc.podUpdateChannels, make(chan podUpdateItem, podUpdateChannelSize))
 }
 go func(stopCh <-chan struct{}) {
  for {
   item, shutdown := tc.nodeUpdateQueue.Get()
   if shutdown {
    break
   }
   nodeUpdate := item.(nodeUpdateItem)
   hash := hash(nodeUpdate.nodeName, UpdateWorkerSize)
   select {
   case <-stopCh:
    tc.nodeUpdateQueue.Done(item)
    return
   case tc.nodeUpdateChannels[hash] <- nodeUpdate:
   }
  }
 }(stopCh)
 go func(stopCh <-chan struct{}) {
  for {
   item, shutdown := tc.podUpdateQueue.Get()
   if shutdown {
    break
   }
   podUpdate := item.(podUpdateItem)
   hash := hash(podUpdate.nodeName, UpdateWorkerSize)
   select {
   case <-stopCh:
    tc.podUpdateQueue.Done(item)
    return
   case tc.podUpdateChannels[hash] <- podUpdate:
   }
  }
 }(stopCh)
 wg := sync.WaitGroup{}
 wg.Add(UpdateWorkerSize)
 for i := 0; i < UpdateWorkerSize; i++ {
  go tc.worker(i, wg.Done, stopCh)
 }
 wg.Wait()
}
func (tc *NoExecuteTaintManager) worker(worker int, done func(), stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer done()
 for {
  select {
  case <-stopCh:
   return
  case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
   tc.handleNodeUpdate(nodeUpdate)
   tc.nodeUpdateQueue.Done(nodeUpdate)
  case podUpdate := <-tc.podUpdateChannels[worker]:
  priority:
   for {
    select {
    case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
     tc.handleNodeUpdate(nodeUpdate)
     tc.nodeUpdateQueue.Done(nodeUpdate)
    default:
     break priority
    }
   }
   tc.handlePodUpdate(podUpdate)
   tc.podUpdateQueue.Done(podUpdate)
  }
 }
}
func (tc *NoExecuteTaintManager) PodUpdated(oldPod *v1.Pod, newPod *v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podName := ""
 podNamespace := ""
 nodeName := ""
 oldTolerations := []v1.Toleration{}
 if oldPod != nil {
  podName = oldPod.Name
  podNamespace = oldPod.Namespace
  nodeName = oldPod.Spec.NodeName
  oldTolerations = oldPod.Spec.Tolerations
 }
 newTolerations := []v1.Toleration{}
 if newPod != nil {
  podName = newPod.Name
  podNamespace = newPod.Namespace
  nodeName = newPod.Spec.NodeName
  newTolerations = newPod.Spec.Tolerations
 }
 if oldPod != nil && newPod != nil && helper.Semantic.DeepEqual(oldTolerations, newTolerations) && oldPod.Spec.NodeName == newPod.Spec.NodeName {
  return
 }
 updateItem := podUpdateItem{podName: podName, podNamespace: podNamespace, nodeName: nodeName}
 tc.podUpdateQueue.Add(updateItem)
}
func (tc *NoExecuteTaintManager) NodeUpdated(oldNode *v1.Node, newNode *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName := ""
 oldTaints := []v1.Taint{}
 if oldNode != nil {
  nodeName = oldNode.Name
  oldTaints = getNoExecuteTaints(oldNode.Spec.Taints)
 }
 newTaints := []v1.Taint{}
 if newNode != nil {
  nodeName = newNode.Name
  newTaints = getNoExecuteTaints(newNode.Spec.Taints)
 }
 if oldNode != nil && newNode != nil && helper.Semantic.DeepEqual(oldTaints, newTaints) {
  return
 }
 updateItem := nodeUpdateItem{nodeName: nodeName}
 tc.nodeUpdateQueue.Add(updateItem)
}
func (tc *NoExecuteTaintManager) cancelWorkWithEvent(nsName types.NamespacedName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if tc.taintEvictionQueue.CancelWork(nsName.String()) {
  tc.emitCancelPodDeletionEvent(nsName)
 }
}
func (tc *NoExecuteTaintManager) processPodOnNode(podNamespacedName types.NamespacedName, nodeName string, tolerations []v1.Toleration, taints []v1.Taint, now time.Time) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(taints) == 0 {
  tc.cancelWorkWithEvent(podNamespacedName)
 }
 allTolerated, usedTolerations := v1helper.GetMatchingTolerations(taints, tolerations)
 if !allTolerated {
  klog.V(2).Infof("Not all taints are tolerated after update for Pod %v on %v", podNamespacedName.String(), nodeName)
  tc.cancelWorkWithEvent(podNamespacedName)
  tc.taintEvictionQueue.AddWork(NewWorkArgs(podNamespacedName.Name, podNamespacedName.Namespace), time.Now(), time.Now())
  return
 }
 minTolerationTime := getMinTolerationTime(usedTolerations)
 if minTolerationTime < 0 {
  klog.V(4).Infof("New tolerations for %v tolerate forever. Scheduled deletion won't be cancelled if already scheduled.", podNamespacedName.String())
  return
 }
 startTime := now
 triggerTime := startTime.Add(minTolerationTime)
 scheduledEviction := tc.taintEvictionQueue.GetWorkerUnsafe(podNamespacedName.String())
 if scheduledEviction != nil {
  startTime = scheduledEviction.CreatedAt
  if startTime.Add(minTolerationTime).Before(triggerTime) {
   return
  }
  tc.cancelWorkWithEvent(podNamespacedName)
 }
 tc.taintEvictionQueue.AddWork(NewWorkArgs(podNamespacedName.Name, podNamespacedName.Namespace), startTime, triggerTime)
}
func (tc *NoExecuteTaintManager) handlePodUpdate(podUpdate podUpdateItem) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, err := tc.getPod(podUpdate.podName, podUpdate.podNamespace)
 if err != nil {
  if apierrors.IsNotFound(err) {
   podNamespacedName := types.NamespacedName{Namespace: podUpdate.podNamespace, Name: podUpdate.podName}
   klog.V(4).Infof("Noticed pod deletion: %#v", podNamespacedName)
   tc.cancelWorkWithEvent(podNamespacedName)
   return
  }
  utilruntime.HandleError(fmt.Errorf("could not get pod %s/%s: %v", podUpdate.podName, podUpdate.podNamespace, err))
  return
 }
 if pod.Spec.NodeName != podUpdate.nodeName {
  return
 }
 podNamespacedName := types.NamespacedName{Namespace: pod.Namespace, Name: pod.Name}
 klog.V(4).Infof("Noticed pod update: %#v", podNamespacedName)
 nodeName := pod.Spec.NodeName
 if nodeName == "" {
  return
 }
 taints, ok := func() ([]v1.Taint, bool) {
  tc.taintedNodesLock.Lock()
  defer tc.taintedNodesLock.Unlock()
  taints, ok := tc.taintedNodes[nodeName]
  return taints, ok
 }()
 if !ok {
  return
 }
 tc.processPodOnNode(podNamespacedName, nodeName, pod.Spec.Tolerations, taints, time.Now())
}
func (tc *NoExecuteTaintManager) handleNodeUpdate(nodeUpdate nodeUpdateItem) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, err := tc.getNode(nodeUpdate.nodeName)
 if err != nil {
  if apierrors.IsNotFound(err) {
   klog.V(4).Infof("Noticed node deletion: %#v", nodeUpdate.nodeName)
   tc.taintedNodesLock.Lock()
   defer tc.taintedNodesLock.Unlock()
   delete(tc.taintedNodes, nodeUpdate.nodeName)
   return
  }
  utilruntime.HandleError(fmt.Errorf("cannot get node %s: %v", nodeUpdate.nodeName, err))
  return
 }
 klog.V(4).Infof("Noticed node update: %#v", nodeUpdate)
 taints := getNoExecuteTaints(node.Spec.Taints)
 func() {
  tc.taintedNodesLock.Lock()
  defer tc.taintedNodesLock.Unlock()
  klog.V(4).Infof("Updating known taints on node %v: %v", node.Name, taints)
  if len(taints) == 0 {
   delete(tc.taintedNodes, node.Name)
  } else {
   tc.taintedNodes[node.Name] = taints
  }
 }()
 pods, err := getPodsAssignedToNode(tc.client, node.Name)
 if err != nil {
  klog.Errorf(err.Error())
  return
 }
 if len(pods) == 0 {
  return
 }
 if len(taints) == 0 {
  klog.V(4).Infof("All taints were removed from the Node %v. Cancelling all evictions...", node.Name)
  for i := range pods {
   tc.cancelWorkWithEvent(types.NamespacedName{Namespace: pods[i].Namespace, Name: pods[i].Name})
  }
  return
 }
 now := time.Now()
 for i := range pods {
  pod := &pods[i]
  podNamespacedName := types.NamespacedName{Namespace: pod.Namespace, Name: pod.Name}
  tc.processPodOnNode(podNamespacedName, node.Name, pod.Spec.Tolerations, taints, now)
 }
}
func (tc *NoExecuteTaintManager) emitPodDeletionEvent(nsName types.NamespacedName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if tc.recorder == nil {
  return
 }
 ref := &v1.ObjectReference{Kind: "Pod", Name: nsName.Name, Namespace: nsName.Namespace}
 tc.recorder.Eventf(ref, v1.EventTypeNormal, "TaintManagerEviction", "Marking for deletion Pod %s", nsName.String())
}
func (tc *NoExecuteTaintManager) emitCancelPodDeletionEvent(nsName types.NamespacedName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if tc.recorder == nil {
  return
 }
 ref := &v1.ObjectReference{Kind: "Pod", Name: nsName.Name, Namespace: nsName.Namespace}
 tc.recorder.Eventf(ref, v1.EventTypeNormal, "TaintManagerEviction", "Cancelling deletion of Pod %s", nsName.String())
}
