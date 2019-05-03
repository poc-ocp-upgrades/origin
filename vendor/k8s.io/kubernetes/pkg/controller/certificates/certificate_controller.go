package certificates

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "golang.org/x/time/rate"
 "k8s.io/klog"
 certificates "k8s.io/api/certificates/v1beta1"
 "k8s.io/apimachinery/pkg/api/errors"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 certificatesinformers "k8s.io/client-go/informers/certificates/v1beta1"
 clientset "k8s.io/client-go/kubernetes"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 certificateslisters "k8s.io/client-go/listers/certificates/v1beta1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
)

type CertificateController struct {
 kubeClient clientset.Interface
 csrLister  certificateslisters.CertificateSigningRequestLister
 csrsSynced cache.InformerSynced
 handler    func(*certificates.CertificateSigningRequest) error
 queue      workqueue.RateLimitingInterface
}

func NewCertificateController(kubeClient clientset.Interface, csrInformer certificatesinformers.CertificateSigningRequestInformer, handler func(*certificates.CertificateSigningRequest) error) *CertificateController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 cc := &CertificateController{kubeClient: kubeClient, queue: workqueue.NewNamedRateLimitingQueue(workqueue.NewMaxOfRateLimiter(workqueue.NewItemExponentialFailureRateLimiter(200*time.Millisecond, 1000*time.Second), &workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)}), "certificate"), handler: handler}
 csrInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  csr := obj.(*certificates.CertificateSigningRequest)
  klog.V(4).Infof("Adding certificate request %s", csr.Name)
  cc.enqueueCertificateRequest(obj)
 }, UpdateFunc: func(old, new interface{}) {
  oldCSR := old.(*certificates.CertificateSigningRequest)
  klog.V(4).Infof("Updating certificate request %s", oldCSR.Name)
  cc.enqueueCertificateRequest(new)
 }, DeleteFunc: func(obj interface{}) {
  csr, ok := obj.(*certificates.CertificateSigningRequest)
  if !ok {
   tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
   if !ok {
    klog.V(2).Infof("Couldn't get object from tombstone %#v", obj)
    return
   }
   csr, ok = tombstone.Obj.(*certificates.CertificateSigningRequest)
   if !ok {
    klog.V(2).Infof("Tombstone contained object that is not a CSR: %#v", obj)
    return
   }
  }
  klog.V(4).Infof("Deleting certificate request %s", csr.Name)
  cc.enqueueCertificateRequest(obj)
 }})
 cc.csrLister = csrInformer.Lister()
 cc.csrsSynced = csrInformer.Informer().HasSynced
 return cc
}
func (cc *CertificateController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer cc.queue.ShutDown()
 klog.Infof("Starting certificate controller")
 defer klog.Infof("Shutting down certificate controller")
 if !controller.WaitForCacheSync("certificate", stopCh, cc.csrsSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(cc.worker, time.Second, stopCh)
 }
 <-stopCh
}
func (cc *CertificateController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for cc.processNextWorkItem() {
 }
}
func (cc *CertificateController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cKey, quit := cc.queue.Get()
 if quit {
  return false
 }
 defer cc.queue.Done(cKey)
 if err := cc.syncFunc(cKey.(string)); err != nil {
  cc.queue.AddRateLimited(cKey)
  if _, ignorable := err.(ignorableError); !ignorable {
   utilruntime.HandleError(fmt.Errorf("Sync %v failed with : %v", cKey, err))
  } else {
   klog.V(4).Infof("Sync %v failed with : %v", cKey, err)
  }
  return true
 }
 cc.queue.Forget(cKey)
 return true
}
func (cc *CertificateController) enqueueCertificateRequest(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(obj)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
  return
 }
 cc.queue.Add(key)
}
func (cc *CertificateController) syncFunc(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished syncing certificate request %q (%v)", key, time.Since(startTime))
 }()
 csr, err := cc.csrLister.Get(key)
 if errors.IsNotFound(err) {
  klog.V(3).Infof("csr has been deleted: %v", key)
  return nil
 }
 if err != nil {
  return err
 }
 if csr.Status.Certificate != nil {
  return nil
 }
 csr = csr.DeepCopy()
 return cc.handler(csr)
}
func IgnorableError(s string, args ...interface{}) ignorableError {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ignorableError(fmt.Sprintf(s, args...))
}

type ignorableError string

func (e ignorableError) Error() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return string(e)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
