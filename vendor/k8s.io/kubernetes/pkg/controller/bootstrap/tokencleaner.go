package bootstrap

import (
 "fmt"
 "time"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 bootstrapapi "k8s.io/cluster-bootstrap/token/api"
 "k8s.io/klog"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/util/metrics"
)

type TokenCleanerOptions struct {
 TokenSecretNamespace string
 SecretResync         time.Duration
}

func DefaultTokenCleanerOptions() TokenCleanerOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return TokenCleanerOptions{TokenSecretNamespace: api.NamespaceSystem}
}

type TokenCleaner struct {
 tokenSecretNamespace string
 client               clientset.Interface
 secretLister         corelisters.SecretLister
 secretSynced         cache.InformerSynced
 queue                workqueue.RateLimitingInterface
}

func NewTokenCleaner(cl clientset.Interface, secrets coreinformers.SecretInformer, options TokenCleanerOptions) (*TokenCleaner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 e := &TokenCleaner{client: cl, secretLister: secrets.Lister(), secretSynced: secrets.Informer().HasSynced, tokenSecretNamespace: options.TokenSecretNamespace, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "token_cleaner")}
 if cl.CoreV1().RESTClient().GetRateLimiter() != nil {
  if err := metrics.RegisterMetricAndTrackRateLimiterUsage("token_cleaner", cl.CoreV1().RESTClient().GetRateLimiter()); err != nil {
   return nil, err
  }
 }
 secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
  switch t := obj.(type) {
  case *v1.Secret:
   return t.Type == bootstrapapi.SecretTypeBootstrapToken && t.Namespace == e.tokenSecretNamespace
  default:
   utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
   return false
  }
 }, Handler: cache.ResourceEventHandlerFuncs{AddFunc: e.enqueueSecrets, UpdateFunc: func(oldSecret, newSecret interface{}) {
  e.enqueueSecrets(newSecret)
 }}}, options.SecretResync)
 return e, nil
}
func (tc *TokenCleaner) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer tc.queue.ShutDown()
 klog.Infof("Starting token cleaner controller")
 defer klog.Infof("Shutting down token cleaner controller")
 if !controller.WaitForCacheSync("token_cleaner", stopCh, tc.secretSynced) {
  return
 }
 go wait.Until(tc.worker, 10*time.Second, stopCh)
 <-stopCh
}
func (tc *TokenCleaner) enqueueSecrets(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(obj)
 if err != nil {
  utilruntime.HandleError(err)
  return
 }
 tc.queue.Add(key)
}
func (tc *TokenCleaner) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for tc.processNextWorkItem() {
 }
}
func (tc *TokenCleaner) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := tc.queue.Get()
 if quit {
  return false
 }
 defer tc.queue.Done(key)
 if err := tc.syncFunc(key.(string)); err != nil {
  tc.queue.AddRateLimited(key)
  utilruntime.HandleError(fmt.Errorf("Sync %v failed with : %v", key, err))
  return true
 }
 tc.queue.Forget(key)
 return true
}
func (tc *TokenCleaner) syncFunc(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished syncing secret %q (%v)", key, time.Since(startTime))
 }()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 ret, err := tc.secretLister.Secrets(namespace).Get(name)
 if apierrors.IsNotFound(err) {
  klog.V(3).Infof("secret has been deleted: %v", key)
  return nil
 }
 if err != nil {
  return err
 }
 if ret.Type == bootstrapapi.SecretTypeBootstrapToken {
  tc.evalSecret(ret)
 }
 return nil
}
func (tc *TokenCleaner) evalSecret(o interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 secret := o.(*v1.Secret)
 if isSecretExpired(secret) {
  klog.V(3).Infof("Deleting expired secret %s/%s", secret.Namespace, secret.Name)
  var options *metav1.DeleteOptions
  if len(secret.UID) > 0 {
   options = &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &secret.UID}}
  }
  err := tc.client.CoreV1().Secrets(secret.Namespace).Delete(secret.Name, options)
  if err != nil && !apierrors.IsConflict(err) && !apierrors.IsNotFound(err) {
   klog.V(3).Infof("Error deleting Secret: %v", err)
  }
 }
}
