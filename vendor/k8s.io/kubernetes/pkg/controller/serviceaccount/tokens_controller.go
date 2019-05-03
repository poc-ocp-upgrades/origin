package serviceaccount

import (
 "bytes"
 "fmt"
 "time"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/types"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 informers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 listersv1 "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 clientretry "k8s.io/client-go/util/retry"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/registry/core/secret"
 "k8s.io/kubernetes/pkg/serviceaccount"
 "k8s.io/kubernetes/pkg/util/metrics"
)

const ServiceServingCASecretKey = "service-ca.crt"

var RemoveTokenBackoff = wait.Backoff{Steps: 10, Duration: 100 * time.Millisecond, Jitter: 1.0}

type TokensControllerOptions struct {
 TokenGenerator       serviceaccount.TokenGenerator
 ServiceAccountResync time.Duration
 SecretResync         time.Duration
 RootCA               []byte
 MaxRetries           int
 ServiceServingCA     []byte
}

func NewTokensController(serviceAccounts informers.ServiceAccountInformer, secrets informers.SecretInformer, cl clientset.Interface, options TokensControllerOptions) (*TokensController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 maxRetries := options.MaxRetries
 if maxRetries == 0 {
  maxRetries = 10
 }
 e := &TokensController{client: cl, token: options.TokenGenerator, rootCA: options.RootCA, serviceServingCA: options.ServiceServingCA, syncServiceAccountQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "serviceaccount_tokens_service"), syncSecretQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "serviceaccount_tokens_secret"), maxRetries: maxRetries}
 if cl != nil && cl.CoreV1().RESTClient().GetRateLimiter() != nil {
  if err := metrics.RegisterMetricAndTrackRateLimiterUsage("serviceaccount_tokens_controller", cl.CoreV1().RESTClient().GetRateLimiter()); err != nil {
   return nil, err
  }
 }
 e.serviceAccounts = serviceAccounts.Lister()
 e.serviceAccountSynced = serviceAccounts.Informer().HasSynced
 serviceAccounts.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: e.queueServiceAccountSync, UpdateFunc: e.queueServiceAccountUpdateSync, DeleteFunc: e.queueServiceAccountSync}, options.ServiceAccountResync)
 secretCache := secrets.Informer().GetIndexer()
 e.updatedSecrets = cache.NewIntegerResourceVersionMutationCache(secretCache, secretCache, 60*time.Second, true)
 e.secretSynced = secrets.Informer().HasSynced
 secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
  switch t := obj.(type) {
  case *v1.Secret:
   return t.Type == v1.SecretTypeServiceAccountToken
  default:
   utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
   return false
  }
 }, Handler: cache.ResourceEventHandlerFuncs{AddFunc: e.queueSecretSync, UpdateFunc: e.queueSecretUpdateSync, DeleteFunc: e.queueSecretSync}}, options.SecretResync)
 return e, nil
}

type TokensController struct {
 client                  clientset.Interface
 token                   serviceaccount.TokenGenerator
 rootCA                  []byte
 serviceServingCA        []byte
 serviceAccounts         listersv1.ServiceAccountLister
 updatedSecrets          cache.MutationCache
 serviceAccountSynced    cache.InformerSynced
 secretSynced            cache.InformerSynced
 syncServiceAccountQueue workqueue.RateLimitingInterface
 syncSecretQueue         workqueue.RateLimitingInterface
 maxRetries              int
}

func (e *TokensController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer e.syncServiceAccountQueue.ShutDown()
 defer e.syncSecretQueue.ShutDown()
 if !controller.WaitForCacheSync("tokens", stopCh, e.serviceAccountSynced, e.secretSynced) {
  return
 }
 klog.V(5).Infof("Starting workers")
 for i := 0; i < workers; i++ {
  go wait.Until(e.syncServiceAccount, 0, stopCh)
  go wait.Until(e.syncSecret, 0, stopCh)
 }
 <-stopCh
 klog.V(1).Infof("Shutting down")
}
func (e *TokensController) queueServiceAccountSync(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if serviceAccount, ok := obj.(*v1.ServiceAccount); ok {
  e.syncServiceAccountQueue.Add(makeServiceAccountKey(serviceAccount))
 }
}
func (e *TokensController) queueServiceAccountUpdateSync(oldObj interface{}, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if serviceAccount, ok := newObj.(*v1.ServiceAccount); ok {
  e.syncServiceAccountQueue.Add(makeServiceAccountKey(serviceAccount))
 }
}
func (e *TokensController) retryOrForget(queue workqueue.RateLimitingInterface, key interface{}, requeue bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !requeue {
  queue.Forget(key)
  return
 }
 requeueCount := queue.NumRequeues(key)
 if requeueCount < e.maxRetries {
  queue.AddRateLimited(key)
  return
 }
 klog.V(4).Infof("retried %d times: %#v", requeueCount, key)
 queue.Forget(key)
}
func (e *TokensController) queueSecretSync(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if secret, ok := obj.(*v1.Secret); ok {
  e.syncSecretQueue.Add(makeSecretQueueKey(secret))
 }
}
func (e *TokensController) queueSecretUpdateSync(oldObj interface{}, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if secret, ok := newObj.(*v1.Secret); ok {
  e.syncSecretQueue.Add(makeSecretQueueKey(secret))
 }
}
func (e *TokensController) syncServiceAccount() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := e.syncServiceAccountQueue.Get()
 if quit {
  return
 }
 defer e.syncServiceAccountQueue.Done(key)
 retry := false
 defer func() {
  e.retryOrForget(e.syncServiceAccountQueue, key, retry)
 }()
 saInfo, err := parseServiceAccountKey(key)
 if err != nil {
  klog.Error(err)
  return
 }
 sa, err := e.getServiceAccount(saInfo.namespace, saInfo.name, saInfo.uid, false)
 switch {
 case err != nil:
  klog.Error(err)
  retry = true
 case sa == nil:
  klog.V(4).Infof("syncServiceAccount(%s/%s), service account deleted, removing tokens", saInfo.namespace, saInfo.name)
  sa = &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: saInfo.namespace, Name: saInfo.name, UID: saInfo.uid}}
  retry, err = e.deleteTokens(sa)
  if err != nil {
   klog.Errorf("error deleting serviceaccount tokens for %s/%s: %v", saInfo.namespace, saInfo.name, err)
  }
 default:
  retry, err = e.ensureReferencedToken(sa)
  if err != nil {
   klog.Errorf("error synchronizing serviceaccount %s/%s: %v", saInfo.namespace, saInfo.name, err)
  }
 }
}
func (e *TokensController) syncSecret() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := e.syncSecretQueue.Get()
 if quit {
  return
 }
 defer e.syncSecretQueue.Done(key)
 retry := false
 defer func() {
  e.retryOrForget(e.syncSecretQueue, key, retry)
 }()
 secretInfo, err := parseSecretQueueKey(key)
 if err != nil {
  klog.Error(err)
  return
 }
 secret, err := e.getSecret(secretInfo.namespace, secretInfo.name, secretInfo.uid, false)
 switch {
 case err != nil:
  klog.Error(err)
  retry = true
 case secret == nil:
  if sa, saErr := e.getServiceAccount(secretInfo.namespace, secretInfo.saName, secretInfo.saUID, false); saErr == nil && sa != nil {
   if err := clientretry.RetryOnConflict(RemoveTokenBackoff, func() error {
    return e.removeSecretReference(secretInfo.namespace, secretInfo.saName, secretInfo.saUID, secretInfo.name)
   }); err != nil {
    klog.Error(err)
   }
  }
 default:
  sa, saErr := e.getServiceAccount(secretInfo.namespace, secretInfo.saName, secretInfo.saUID, true)
  switch {
  case saErr != nil:
   klog.Error(saErr)
   retry = true
  case sa == nil:
   klog.V(4).Infof("syncSecret(%s/%s), service account does not exist, deleting token", secretInfo.namespace, secretInfo.name)
   if retriable, err := e.deleteToken(secretInfo.namespace, secretInfo.name, secretInfo.uid); err != nil {
    klog.Errorf("error deleting serviceaccount token %s/%s for service account %s: %v", secretInfo.namespace, secretInfo.name, secretInfo.saName, err)
    retry = retriable
   }
  default:
   if retriable, err := e.generateTokenIfNeeded(sa, secret); err != nil {
    klog.Errorf("error populating serviceaccount token %s/%s for service account %s: %v", secretInfo.namespace, secretInfo.name, secretInfo.saName, err)
    retry = retriable
   }
  }
 }
}
func (e *TokensController) deleteTokens(serviceAccount *v1.ServiceAccount) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tokens, err := e.listTokenSecrets(serviceAccount)
 if err != nil {
  return false, err
 }
 retry := false
 errs := []error{}
 for _, token := range tokens {
  r, err := e.deleteToken(token.Namespace, token.Name, token.UID)
  if err != nil {
   errs = append(errs, err)
  }
  if r {
   retry = true
  }
 }
 return retry, utilerrors.NewAggregate(errs)
}
func (e *TokensController) deleteToken(ns, name string, uid types.UID) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var opts *metav1.DeleteOptions
 if len(uid) > 0 {
  opts = &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &uid}}
 }
 err := e.client.CoreV1().Secrets(ns).Delete(name, opts)
 if err == nil || apierrors.IsNotFound(err) || apierrors.IsConflict(err) {
  return false, nil
 }
 return true, err
}
func (e *TokensController) ensureReferencedToken(serviceAccount *v1.ServiceAccount) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if hasToken, err := e.hasReferencedToken(serviceAccount); err != nil {
  return false, err
 } else if hasToken {
  return false, nil
 }
 serviceAccounts := e.client.CoreV1().ServiceAccounts(serviceAccount.Namespace)
 liveServiceAccount, err := serviceAccounts.Get(serviceAccount.Name, metav1.GetOptions{})
 if err != nil {
  return true, err
 }
 if liveServiceAccount.ResourceVersion != serviceAccount.ResourceVersion {
  klog.V(4).Infof("liveServiceAccount.ResourceVersion (%s) does not match cache (%s), retrying", liveServiceAccount.ResourceVersion, serviceAccount.ResourceVersion)
  return true, nil
 }
 secret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: secret.Strategy.GenerateName(fmt.Sprintf("%s-token-", serviceAccount.Name)), Namespace: serviceAccount.Namespace, Annotations: map[string]string{v1.ServiceAccountNameKey: serviceAccount.Name, v1.ServiceAccountUIDKey: string(serviceAccount.UID)}}, Type: v1.SecretTypeServiceAccountToken, Data: map[string][]byte{}}
 token, err := e.token.GenerateToken(serviceaccount.LegacyClaims(*serviceAccount, *secret))
 if err != nil {
  return true, err
 }
 secret.Data[v1.ServiceAccountTokenKey] = []byte(token)
 secret.Data[v1.ServiceAccountNamespaceKey] = []byte(serviceAccount.Namespace)
 if e.rootCA != nil && len(e.rootCA) > 0 {
  secret.Data[v1.ServiceAccountRootCAKey] = e.rootCA
 }
 if e.serviceServingCA != nil && len(e.serviceServingCA) > 0 {
  secret.Data[ServiceServingCASecretKey] = e.serviceServingCA
 }
 createdToken, err := e.client.CoreV1().Secrets(serviceAccount.Namespace).Create(secret)
 if err != nil {
  return true, err
 }
 e.updatedSecrets.Mutation(createdToken)
 addedReference := false
 err = clientretry.RetryOnConflict(clientretry.DefaultRetry, func() error {
  defer func() {
   liveServiceAccount = nil
  }()
  if liveServiceAccount == nil {
   liveServiceAccount, err = serviceAccounts.Get(serviceAccount.Name, metav1.GetOptions{})
   if err != nil {
    return err
   }
   if liveServiceAccount.UID != serviceAccount.UID {
    return nil
   }
   if hasToken, err := e.hasReferencedToken(liveServiceAccount); err != nil {
    return nil
   } else if hasToken {
    return nil
   }
  }
  liveServiceAccount.Secrets = append(liveServiceAccount.Secrets, v1.ObjectReference{Name: secret.Name})
  if _, err := serviceAccounts.Update(liveServiceAccount); err != nil {
   return err
  }
  addedReference = true
  return nil
 })
 if !addedReference {
  klog.V(2).Infof("deleting secret %s/%s because reference couldn't be added (%v)", secret.Namespace, secret.Name, err)
  deleteOpts := &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &createdToken.UID}}
  if deleteErr := e.client.CoreV1().Secrets(createdToken.Namespace).Delete(createdToken.Name, deleteOpts); deleteErr != nil {
   klog.Error(deleteErr)
  }
 }
 if err != nil {
  if apierrors.IsConflict(err) || apierrors.IsNotFound(err) {
   return false, nil
  }
  return true, err
 }
 return false, nil
}
func (e *TokensController) hasReferencedToken(serviceAccount *v1.ServiceAccount) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(serviceAccount.Secrets) == 0 {
  return false, nil
 }
 allSecrets, err := e.listTokenSecrets(serviceAccount)
 if err != nil {
  return false, err
 }
 referencedSecrets := getSecretReferences(serviceAccount)
 for _, secret := range allSecrets {
  if referencedSecrets.Has(secret.Name) {
   return true, nil
  }
 }
 return false, nil
}
func (e *TokensController) secretUpdateNeeded(secret *v1.Secret) (bool, bool, bool, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 caData := secret.Data[v1.ServiceAccountRootCAKey]
 needsCA := len(e.rootCA) > 0 && bytes.Compare(caData, e.rootCA) != 0
 needsServiceServingCA := len(e.serviceServingCA) > 0 && bytes.Compare(secret.Data[ServiceServingCASecretKey], e.serviceServingCA) != 0
 needsNamespace := len(secret.Data[v1.ServiceAccountNamespaceKey]) == 0
 tokenData := secret.Data[v1.ServiceAccountTokenKey]
 needsToken := len(tokenData) == 0
 return needsCA, needsServiceServingCA, needsNamespace, needsToken
}
func (e *TokensController) generateTokenIfNeeded(serviceAccount *v1.ServiceAccount, cachedSecret *v1.Secret) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if needsCA, needsServiceServingCA, needsNamespace, needsToken := e.secretUpdateNeeded(cachedSecret); !needsCA && !needsServiceServingCA && !needsToken && !needsNamespace {
  return false, nil
 }
 secrets := e.client.CoreV1().Secrets(cachedSecret.Namespace)
 liveSecret, err := secrets.Get(cachedSecret.Name, metav1.GetOptions{})
 if err != nil {
  return !apierrors.IsNotFound(err), err
 }
 if liveSecret.ResourceVersion != cachedSecret.ResourceVersion {
  klog.V(2).Infof("secret %s/%s is not up to date, skipping token population", liveSecret.Namespace, liveSecret.Name)
  return false, nil
 }
 needsCA, needsServiceServingCA, needsNamespace, needsToken := e.secretUpdateNeeded(liveSecret)
 if !needsCA && !needsServiceServingCA && !needsToken && !needsNamespace {
  return false, nil
 }
 if liveSecret.Annotations == nil {
  liveSecret.Annotations = map[string]string{}
 }
 if liveSecret.Data == nil {
  liveSecret.Data = map[string][]byte{}
 }
 if needsCA {
  liveSecret.Data[v1.ServiceAccountRootCAKey] = e.rootCA
 }
 if needsServiceServingCA {
  liveSecret.Data[ServiceServingCASecretKey] = e.serviceServingCA
 }
 if needsNamespace {
  liveSecret.Data[v1.ServiceAccountNamespaceKey] = []byte(liveSecret.Namespace)
 }
 if needsToken {
  token, err := e.token.GenerateToken(serviceaccount.LegacyClaims(*serviceAccount, *liveSecret))
  if err != nil {
   return false, err
  }
  liveSecret.Data[v1.ServiceAccountTokenKey] = []byte(token)
 }
 liveSecret.Annotations[v1.ServiceAccountNameKey] = serviceAccount.Name
 liveSecret.Annotations[v1.ServiceAccountUIDKey] = string(serviceAccount.UID)
 _, err = secrets.Update(liveSecret)
 if apierrors.IsConflict(err) || apierrors.IsNotFound(err) {
  return false, nil
 }
 if err != nil {
  return true, err
 }
 return false, nil
}
func (e *TokensController) removeSecretReference(saNamespace string, saName string, saUID types.UID, secretName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 serviceAccounts := e.client.CoreV1().ServiceAccounts(saNamespace)
 serviceAccount, err := serviceAccounts.Get(saName, metav1.GetOptions{})
 if apierrors.IsNotFound(err) {
  return nil
 }
 if err != nil {
  return err
 }
 if len(saUID) > 0 && saUID != serviceAccount.UID {
  return nil
 }
 if !getSecretReferences(serviceAccount).Has(secretName) {
  return nil
 }
 secrets := []v1.ObjectReference{}
 for _, s := range serviceAccount.Secrets {
  if s.Name != secretName {
   secrets = append(secrets, s)
  }
 }
 serviceAccount.Secrets = secrets
 _, err = serviceAccounts.Update(serviceAccount)
 if apierrors.IsNotFound(err) {
  return nil
 }
 return err
}
func (e *TokensController) getServiceAccount(ns string, name string, uid types.UID, fetchOnCacheMiss bool) (*v1.ServiceAccount, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sa, err := e.serviceAccounts.ServiceAccounts(ns).Get(name)
 if err != nil && !apierrors.IsNotFound(err) {
  return nil, err
 }
 if sa != nil {
  if len(uid) == 0 || uid == sa.UID {
   return sa, nil
  }
 }
 if !fetchOnCacheMiss {
  return nil, nil
 }
 sa, err = e.client.CoreV1().ServiceAccounts(ns).Get(name, metav1.GetOptions{})
 if apierrors.IsNotFound(err) {
  return nil, nil
 }
 if err != nil {
  return nil, err
 }
 if len(uid) == 0 || uid == sa.UID {
  return sa, nil
 }
 return nil, nil
}
func (e *TokensController) getSecret(ns string, name string, uid types.UID, fetchOnCacheMiss bool) (*v1.Secret, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, exists, err := e.updatedSecrets.GetByKey(makeCacheKey(ns, name))
 if err != nil {
  return nil, err
 }
 if exists {
  secret, ok := obj.(*v1.Secret)
  if !ok {
   return nil, fmt.Errorf("expected *v1.Secret, got %#v", secret)
  }
  if len(uid) == 0 || uid == secret.UID {
   return secret, nil
  }
 }
 if !fetchOnCacheMiss {
  return nil, nil
 }
 secret, err := e.client.CoreV1().Secrets(ns).Get(name, metav1.GetOptions{})
 if apierrors.IsNotFound(err) {
  return nil, nil
 }
 if err != nil {
  return nil, err
 }
 if len(uid) == 0 || uid == secret.UID {
  return secret, nil
 }
 return nil, nil
}
func (e *TokensController) listTokenSecrets(serviceAccount *v1.ServiceAccount) ([]*v1.Secret, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 namespaceSecrets, err := e.updatedSecrets.ByIndex("namespace", serviceAccount.Namespace)
 if err != nil {
  return nil, err
 }
 items := []*v1.Secret{}
 for _, obj := range namespaceSecrets {
  secret := obj.(*v1.Secret)
  if serviceaccount.IsServiceAccountToken(secret, serviceAccount) {
   items = append(items, secret)
  }
 }
 return items, nil
}
func serviceAccountNameAndUID(secret *v1.Secret) (string, string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if secret.Type != v1.SecretTypeServiceAccountToken {
  return "", ""
 }
 return secret.Annotations[v1.ServiceAccountNameKey], secret.Annotations[v1.ServiceAccountUIDKey]
}
func getSecretReferences(serviceAccount *v1.ServiceAccount) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 references := sets.NewString()
 for _, secret := range serviceAccount.Secrets {
  references.Insert(secret.Name)
 }
 return references
}

type serviceAccountQueueKey struct {
 namespace string
 name      string
 uid       types.UID
}

func makeServiceAccountKey(sa *v1.ServiceAccount) interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return serviceAccountQueueKey{namespace: sa.Namespace, name: sa.Name, uid: sa.UID}
}
func parseServiceAccountKey(key interface{}) (serviceAccountQueueKey, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 queueKey, ok := key.(serviceAccountQueueKey)
 if !ok || len(queueKey.namespace) == 0 || len(queueKey.name) == 0 || len(queueKey.uid) == 0 {
  return serviceAccountQueueKey{}, fmt.Errorf("invalid serviceaccount key: %#v", key)
 }
 return queueKey, nil
}

type secretQueueKey struct {
 namespace string
 name      string
 uid       types.UID
 saName    string
 saUID     types.UID
}

func makeSecretQueueKey(secret *v1.Secret) interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return secretQueueKey{namespace: secret.Namespace, name: secret.Name, uid: secret.UID, saName: secret.Annotations[v1.ServiceAccountNameKey], saUID: types.UID(secret.Annotations[v1.ServiceAccountUIDKey])}
}
func parseSecretQueueKey(key interface{}) (secretQueueKey, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 queueKey, ok := key.(secretQueueKey)
 if !ok || len(queueKey.namespace) == 0 || len(queueKey.name) == 0 || len(queueKey.uid) == 0 || len(queueKey.saName) == 0 {
  return secretQueueKey{}, fmt.Errorf("invalid secret key: %#v", key)
 }
 return queueKey, nil
}
func makeCacheKey(namespace, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return namespace + "/" + name
}
