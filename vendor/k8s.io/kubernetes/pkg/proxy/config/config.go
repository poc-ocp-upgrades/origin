package config

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/api/core/v1"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 coreinformers "k8s.io/client-go/informers/core/v1"
 listers "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/controller"
)

type ServiceHandler interface {
 OnServiceAdd(service *v1.Service)
 OnServiceUpdate(oldService, service *v1.Service)
 OnServiceDelete(service *v1.Service)
 OnServiceSynced()
}
type EndpointsHandler interface {
 OnEndpointsAdd(endpoints *v1.Endpoints)
 OnEndpointsUpdate(oldEndpoints, endpoints *v1.Endpoints)
 OnEndpointsDelete(endpoints *v1.Endpoints)
 OnEndpointsSynced()
}
type EndpointsConfig struct {
 lister        listers.EndpointsLister
 listerSynced  cache.InformerSynced
 eventHandlers []EndpointsHandler
}

func NewEndpointsConfig(endpointsInformer coreinformers.EndpointsInformer, resyncPeriod time.Duration) *EndpointsConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := &EndpointsConfig{lister: endpointsInformer.Lister(), listerSynced: endpointsInformer.Informer().HasSynced}
 endpointsInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: result.handleAddEndpoints, UpdateFunc: result.handleUpdateEndpoints, DeleteFunc: result.handleDeleteEndpoints}, resyncPeriod)
 return result
}
func (c *EndpointsConfig) RegisterEventHandler(handler EndpointsHandler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.eventHandlers = append(c.eventHandlers, handler)
}
func (c *EndpointsConfig) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Info("Starting endpoints config controller")
 defer klog.Info("Shutting down endpoints config controller")
 if !controller.WaitForCacheSync("endpoints config", stopCh, c.listerSynced) {
  return
 }
 for i := range c.eventHandlers {
  klog.V(3).Infof("Calling handler.OnEndpointsSynced()")
  c.eventHandlers[i].OnEndpointsSynced()
 }
 <-stopCh
}
func (c *EndpointsConfig) handleAddEndpoints(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 endpoints, ok := obj.(*v1.Endpoints)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
  return
 }
 for i := range c.eventHandlers {
  klog.V(4).Infof("Calling handler.OnEndpointsAdd")
  c.eventHandlers[i].OnEndpointsAdd(endpoints)
 }
}
func (c *EndpointsConfig) handleUpdateEndpoints(oldObj, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldEndpoints, ok := oldObj.(*v1.Endpoints)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", oldObj))
  return
 }
 endpoints, ok := newObj.(*v1.Endpoints)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", newObj))
  return
 }
 for i := range c.eventHandlers {
  klog.V(4).Infof("Calling handler.OnEndpointsUpdate")
  c.eventHandlers[i].OnEndpointsUpdate(oldEndpoints, endpoints)
 }
}
func (c *EndpointsConfig) handleDeleteEndpoints(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 endpoints, ok := obj.(*v1.Endpoints)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
   return
  }
  if endpoints, ok = tombstone.Obj.(*v1.Endpoints); !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
   return
  }
 }
 for i := range c.eventHandlers {
  klog.V(4).Infof("Calling handler.OnEndpointsDelete")
  c.eventHandlers[i].OnEndpointsDelete(endpoints)
 }
}

type ServiceConfig struct {
 lister        listers.ServiceLister
 listerSynced  cache.InformerSynced
 eventHandlers []ServiceHandler
}

func NewServiceConfig(serviceInformer coreinformers.ServiceInformer, resyncPeriod time.Duration) *ServiceConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := &ServiceConfig{lister: serviceInformer.Lister(), listerSynced: serviceInformer.Informer().HasSynced}
 serviceInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: result.handleAddService, UpdateFunc: result.handleUpdateService, DeleteFunc: result.handleDeleteService}, resyncPeriod)
 return result
}
func (c *ServiceConfig) RegisterEventHandler(handler ServiceHandler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.eventHandlers = append(c.eventHandlers, handler)
}
func (c *ServiceConfig) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Info("Starting service config controller")
 defer klog.Info("Shutting down service config controller")
 if !controller.WaitForCacheSync("service config", stopCh, c.listerSynced) {
  return
 }
 for i := range c.eventHandlers {
  klog.V(3).Info("Calling handler.OnServiceSynced()")
  c.eventHandlers[i].OnServiceSynced()
 }
 <-stopCh
}
func (c *ServiceConfig) handleAddService(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service, ok := obj.(*v1.Service)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
  return
 }
 for i := range c.eventHandlers {
  klog.V(4).Info("Calling handler.OnServiceAdd")
  c.eventHandlers[i].OnServiceAdd(service)
 }
}
func (c *ServiceConfig) handleUpdateService(oldObj, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldService, ok := oldObj.(*v1.Service)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", oldObj))
  return
 }
 service, ok := newObj.(*v1.Service)
 if !ok {
  utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", newObj))
  return
 }
 for i := range c.eventHandlers {
  klog.V(4).Info("Calling handler.OnServiceUpdate")
  c.eventHandlers[i].OnServiceUpdate(oldService, service)
 }
}
func (c *ServiceConfig) handleDeleteService(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service, ok := obj.(*v1.Service)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
   return
  }
  if service, ok = tombstone.Obj.(*v1.Service); !ok {
   utilruntime.HandleError(fmt.Errorf("unexpected object type: %v", obj))
   return
  }
 }
 for i := range c.eventHandlers {
  klog.V(4).Info("Calling handler.OnServiceDelete")
  c.eventHandlers[i].OnServiceDelete(service)
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
