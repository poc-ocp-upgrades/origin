package ipam

import (
 "fmt"
 "math/rand"
 "net"
 "sync"
 "time"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 informers "k8s.io/client-go/informers/core/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 cloudprovider "k8s.io/cloud-provider"
 v1node "k8s.io/kubernetes/pkg/api/v1/node"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
 "k8s.io/kubernetes/pkg/controller"
 nodeutil "k8s.io/kubernetes/pkg/controller/util/node"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 utilnode "k8s.io/kubernetes/pkg/util/node"
 utiltaints "k8s.io/kubernetes/pkg/util/taints"
)

type nodeProcessingInfo struct{ retries int }
type cloudCIDRAllocator struct {
 client            clientset.Interface
 cloud             *gce.Cloud
 nodeLister        corelisters.NodeLister
 nodesSynced       cache.InformerSynced
 nodeUpdateChannel chan string
 recorder          record.EventRecorder
 lock              sync.Mutex
 nodesInProcessing map[string]*nodeProcessingInfo
}

var _ CIDRAllocator = (*cloudCIDRAllocator)(nil)

func NewCloudCIDRAllocator(client clientset.Interface, cloud cloudprovider.Interface, nodeInformer informers.NodeInformer) (CIDRAllocator, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if client == nil {
  klog.Fatalf("kubeClient is nil when starting NodeController")
 }
 eventBroadcaster := record.NewBroadcaster()
 recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "cidrAllocator"})
 eventBroadcaster.StartLogging(klog.Infof)
 klog.V(0).Infof("Sending events to api server.")
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: client.CoreV1().Events("")})
 gceCloud, ok := cloud.(*gce.Cloud)
 if !ok {
  err := fmt.Errorf("cloudCIDRAllocator does not support %v provider", cloud.ProviderName())
  return nil, err
 }
 ca := &cloudCIDRAllocator{client: client, cloud: gceCloud, nodeLister: nodeInformer.Lister(), nodesSynced: nodeInformer.Informer().HasSynced, nodeUpdateChannel: make(chan string, cidrUpdateQueueSize), recorder: recorder, nodesInProcessing: map[string]*nodeProcessingInfo{}}
 nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(ca.AllocateOrOccupyCIDR), UpdateFunc: nodeutil.CreateUpdateNodeHandler(func(_, newNode *v1.Node) error {
  if newNode.Spec.PodCIDR == "" {
   return ca.AllocateOrOccupyCIDR(newNode)
  }
  networkUnavailableTaint := &v1.Taint{Key: schedulerapi.TaintNodeNetworkUnavailable, Effect: v1.TaintEffectNoSchedule}
  _, cond := v1node.GetNodeCondition(&newNode.Status, v1.NodeNetworkUnavailable)
  if cond == nil || cond.Status != v1.ConditionFalse || utiltaints.TaintExists(newNode.Spec.Taints, networkUnavailableTaint) {
   return ca.AllocateOrOccupyCIDR(newNode)
  }
  return nil
 }), DeleteFunc: nodeutil.CreateDeleteNodeHandler(ca.ReleaseCIDR)})
 klog.V(0).Infof("Using cloud CIDR allocator (provider: %v)", cloud.ProviderName())
 return ca, nil
}
func (ca *cloudCIDRAllocator) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting cloud CIDR allocator")
 defer klog.Infof("Shutting down cloud CIDR allocator")
 if !controller.WaitForCacheSync("cidrallocator", stopCh, ca.nodesSynced) {
  return
 }
 for i := 0; i < cidrUpdateWorkers; i++ {
  go ca.worker(stopCh)
 }
 <-stopCh
}
func (ca *cloudCIDRAllocator) worker(stopChan <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for {
  select {
  case workItem, ok := <-ca.nodeUpdateChannel:
   if !ok {
    klog.Warning("Channel nodeCIDRUpdateChannel was unexpectedly closed")
    return
   }
   if err := ca.updateCIDRAllocation(workItem); err == nil {
    klog.V(3).Infof("Updated CIDR for %q", workItem)
   } else {
    klog.Errorf("Error updating CIDR for %q: %v", workItem, err)
    if canRetry, timeout := ca.retryParams(workItem); canRetry {
     klog.V(2).Infof("Retrying update for %q after %v", workItem, timeout)
     time.AfterFunc(timeout, func() {
      ca.nodeUpdateChannel <- workItem
     })
     continue
    }
    klog.Errorf("Exceeded retry count for %q, dropping from queue", workItem)
   }
   ca.removeNodeFromProcessing(workItem)
  case <-stopChan:
   return
  }
 }
}
func (ca *cloudCIDRAllocator) insertNodeToProcessing(nodeName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ca.lock.Lock()
 defer ca.lock.Unlock()
 if _, found := ca.nodesInProcessing[nodeName]; found {
  return false
 }
 ca.nodesInProcessing[nodeName] = &nodeProcessingInfo{}
 return true
}
func (ca *cloudCIDRAllocator) retryParams(nodeName string) (bool, time.Duration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ca.lock.Lock()
 defer ca.lock.Unlock()
 entry, ok := ca.nodesInProcessing[nodeName]
 if !ok {
  klog.Errorf("Cannot get retryParams for %q as entry does not exist", nodeName)
  return false, 0
 }
 count := entry.retries + 1
 if count > updateMaxRetries {
  return false, 0
 }
 ca.nodesInProcessing[nodeName].retries = count
 return true, nodeUpdateRetryTimeout(count)
}
func nodeUpdateRetryTimeout(count int) time.Duration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 timeout := updateRetryTimeout
 for i := 0; i < count && timeout < maxUpdateRetryTimeout; i++ {
  timeout *= 2
 }
 if timeout > maxUpdateRetryTimeout {
  timeout = maxUpdateRetryTimeout
 }
 return time.Duration(timeout.Nanoseconds()/2 + rand.Int63n(timeout.Nanoseconds()))
}
func (ca *cloudCIDRAllocator) removeNodeFromProcessing(nodeName string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ca.lock.Lock()
 defer ca.lock.Unlock()
 delete(ca.nodesInProcessing, nodeName)
}
func (ca *cloudCIDRAllocator) AllocateOrOccupyCIDR(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if node == nil {
  return nil
 }
 if !ca.insertNodeToProcessing(node.Name) {
  klog.V(2).Infof("Node %v is already in a process of CIDR assignment.", node.Name)
  return nil
 }
 klog.V(4).Infof("Putting node %s into the work queue", node.Name)
 ca.nodeUpdateChannel <- node.Name
 return nil
}
func (ca *cloudCIDRAllocator) updateCIDRAllocation(nodeName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, err := ca.nodeLister.Get(nodeName)
 if err != nil {
  if errors.IsNotFound(err) {
   return nil
  }
  klog.Errorf("Failed while getting node %v for updating Node.Spec.PodCIDR: %v", nodeName, err)
  return err
 }
 cidrs, err := ca.cloud.AliasRanges(types.NodeName(nodeName))
 if err != nil {
  nodeutil.RecordNodeStatusChange(ca.recorder, node, "CIDRNotAvailable")
  return fmt.Errorf("failed to allocate cidr: %v", err)
 }
 if len(cidrs) == 0 {
  nodeutil.RecordNodeStatusChange(ca.recorder, node, "CIDRNotAvailable")
  return fmt.Errorf("failed to allocate cidr: Node %v has no CIDRs", node.Name)
 }
 _, cidr, err := net.ParseCIDR(cidrs[0])
 if err != nil {
  return fmt.Errorf("failed to parse string '%s' as a CIDR: %v", cidrs[0], err)
 }
 podCIDR := cidr.String()
 if node.Spec.PodCIDR == podCIDR {
  klog.V(4).Infof("Node %v already has allocated CIDR %v. It matches the proposed one.", node.Name, podCIDR)
 } else {
  if node.Spec.PodCIDR != "" {
   klog.Errorf("PodCIDR being reassigned! Node %v spec has %v, but cloud provider has assigned %v", node.Name, node.Spec.PodCIDR, podCIDR)
  }
  for i := 0; i < cidrUpdateRetries; i++ {
   if err = utilnode.PatchNodeCIDR(ca.client, types.NodeName(node.Name), podCIDR); err == nil {
    klog.Infof("Set node %v PodCIDR to %v", node.Name, podCIDR)
    break
   }
  }
 }
 if err != nil {
  nodeutil.RecordNodeStatusChange(ca.recorder, node, "CIDRAssignmentFailed")
  klog.Errorf("Failed to update node %v PodCIDR to %v after multiple attempts: %v", node.Name, podCIDR, err)
  return err
 }
 err = utilnode.SetNodeCondition(ca.client, types.NodeName(node.Name), v1.NodeCondition{Type: v1.NodeNetworkUnavailable, Status: v1.ConditionFalse, Reason: "RouteCreated", Message: "NodeController create implicit route", LastTransitionTime: metav1.Now()})
 if err != nil {
  klog.Errorf("Error setting route status for node %v: %v", node.Name, err)
 }
 return err
}
func (ca *cloudCIDRAllocator) ReleaseCIDR(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(2).Infof("Node %v PodCIDR (%v) will be released by external cloud provider (not managed by controller)", node.Name, node.Spec.PodCIDR)
 return nil
}
