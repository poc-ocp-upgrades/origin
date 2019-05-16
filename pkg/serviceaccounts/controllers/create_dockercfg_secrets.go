package controllers

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	informers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/credentialprovider"
	"k8s.io/kubernetes/pkg/registry/core/secret"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"sync"
	"time"
	gotime "time"
)

const (
	ServiceAccountTokenSecretNameKey   = "openshift.io/token-secret.name"
	MaxRetriesBeforeResync             = 5
	ServiceAccountTokenValueAnnotation = "openshift.io/token-secret.value"
	CreateDockercfgSecretsController   = "openshift.io/create-dockercfg-secrets"
	PendingTokenAnnotation             = "openshift.io/create-dockercfg-secrets.pending-token"
	DeprecatedKubeCreatedByAnnotation  = "kubernetes.io/created-by"
	maxNameLength                      = 63
	randomLength                       = 5
	maxSecretPrefixNameLength          = maxNameLength - randomLength
)

type DockercfgControllerOptions struct {
	Resync                time.Duration
	DockerURLsInitialized chan struct{}
}

func NewDockercfgController(serviceAccounts informers.ServiceAccountInformer, secrets informers.SecretInformer, cl kclientset.Interface, options DockercfgControllerOptions) *DockercfgController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e := &DockercfgController{client: cl, queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), dockerURLsInitialized: options.DockerURLsInitialized}
	serviceAccountCache := serviceAccounts.Informer().GetStore()
	e.serviceAccountController = serviceAccounts.Informer().GetController()
	serviceAccounts.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		serviceAccount := obj.(*v1.ServiceAccount)
		klog.V(5).Infof("Adding service account %s", serviceAccount.Name)
		e.enqueueServiceAccount(serviceAccount)
	}, UpdateFunc: func(old, cur interface{}) {
		serviceAccount := cur.(*v1.ServiceAccount)
		klog.V(5).Infof("Updating service account %s", serviceAccount.Name)
		e.enqueueServiceAccount(serviceAccount)
	}}, options.Resync)
	e.serviceAccountCache = NewEtcdMutationCache(serviceAccountCache)
	e.secretCache = secrets.Informer().GetIndexer()
	e.secretController = secrets.Informer().GetController()
	secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Secret:
			return t.Type == v1.SecretTypeServiceAccountToken
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: func(cur interface{}) {
		e.handleTokenSecretUpdate(nil, cur)
	}, UpdateFunc: func(old, cur interface{}) {
		e.handleTokenSecretUpdate(old, cur)
	}, DeleteFunc: e.handleTokenSecretDelete}}, options.Resync)
	e.syncHandler = e.syncServiceAccount
	return e
}

type DockercfgController struct {
	client                   kclientset.Interface
	dockerURLLock            sync.Mutex
	dockerURLs               []string
	dockerURLsInitialized    chan struct{}
	serviceAccountCache      MutationCache
	serviceAccountController cache.Controller
	secretCache              cache.Store
	secretController         cache.Controller
	queue                    workqueue.RateLimitingInterface
	syncHandler              func(serviceKey string) error
}

func (e *DockercfgController) handleTokenSecretUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secret := newObj.(*v1.Secret)
	if secret.Annotations[DeprecatedKubeCreatedByAnnotation] != CreateDockercfgSecretsController {
		return
	}
	isPopulated := len(secret.Data[v1.ServiceAccountTokenKey]) > 0
	wasPopulated := false
	if oldObj != nil {
		oldSecret := oldObj.(*v1.Secret)
		wasPopulated = len(oldSecret.Data[v1.ServiceAccountTokenKey]) > 0
		klog.V(5).Infof("Updating token secret %s/%s", secret.Namespace, secret.Name)
	} else {
		klog.V(5).Infof("Adding token secret %s/%s", secret.Namespace, secret.Name)
	}
	if !wasPopulated && isPopulated {
		e.enqueueServiceAccountForToken(secret)
	}
}
func (e *DockercfgController) handleTokenSecretDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secret, isSecret := obj.(*v1.Secret)
	if !isSecret {
		tombstone, objIsTombstone := obj.(cache.DeletedFinalStateUnknown)
		if !objIsTombstone {
			klog.V(2).Infof("Expected tombstone object when deleting token, got %v", obj)
			return
		}
		secret, isSecret = tombstone.Obj.(*v1.Secret)
		if !isSecret {
			klog.V(2).Infof("Expected tombstone object to contain secret, got: %v", obj)
			return
		}
	}
	if secret.Annotations[DeprecatedKubeCreatedByAnnotation] != CreateDockercfgSecretsController {
		return
	}
	if len(secret.Data[v1.ServiceAccountTokenKey]) > 0 {
		return
	}
	e.enqueueServiceAccountForToken(secret)
}
func (e *DockercfgController) enqueueServiceAccountForToken(tokenSecret *v1.Secret) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serviceAccount := &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: tokenSecret.Annotations[v1.ServiceAccountNameKey], Namespace: tokenSecret.Namespace, UID: types.UID(tokenSecret.Annotations[v1.ServiceAccountUIDKey])}}
	key, err := controller.KeyFunc(serviceAccount)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error syncing token secret %s/%s: %v", tokenSecret.Namespace, tokenSecret.Name, err))
		return
	}
	e.queue.Add(key)
}
func (e *DockercfgController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer e.queue.ShutDown()
	klog.Infof("Starting DockercfgController controller")
	defer klog.Infof("Shutting down DockercfgController controller")
	ready := make(chan struct{})
	go e.waitForDockerURLs(ready, stopCh)
	select {
	case <-ready:
	case <-stopCh:
		return
	}
	klog.V(1).Infof("urls found")
	if !cache.WaitForCacheSync(stopCh, e.serviceAccountController.HasSynced, e.secretController.HasSynced) {
		return
	}
	klog.V(1).Infof("caches synced")
	for i := 0; i < workers; i++ {
		go wait.Until(e.worker, time.Second, stopCh)
	}
	<-stopCh
}
func (c *DockercfgController) waitForDockerURLs(ready chan<- struct{}, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	select {
	case <-c.dockerURLsInitialized:
	case <-stopCh:
		return
	}
	close(ready)
}
func (e *DockercfgController) enqueueServiceAccount(serviceAccount *v1.ServiceAccount) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !needsDockercfgSecret(serviceAccount) {
		return
	}
	key, err := controller.KeyFunc(serviceAccount)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", serviceAccount, err)
		return
	}
	e.queue.Add(key)
}
func (e *DockercfgController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for {
		if !e.work() {
			return
		}
	}
}
func (e *DockercfgController) work() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := e.queue.Get()
	if quit {
		return false
	}
	defer e.queue.Done(key)
	if err := e.syncHandler(key.(string)); err == nil {
		e.queue.Forget(key)
	} else {
		if e.queue.NumRequeues(key) > MaxRetriesBeforeResync {
			utilruntime.HandleError(fmt.Errorf("error syncing service, it will be tried again on a resync %v: %v", key, err))
			e.queue.Forget(key)
		} else {
			klog.V(4).Infof("error syncing service, it will be retried %v: %v", key, err)
			e.queue.AddRateLimited(key)
		}
	}
	return true
}
func (e *DockercfgController) SetDockerURLs(newDockerURLs ...string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.dockerURLLock.Lock()
	defer e.dockerURLLock.Unlock()
	e.dockerURLs = newDockerURLs
}
func needsDockercfgSecret(serviceAccount *v1.ServiceAccount) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mountableDockercfgSecrets, imageDockercfgPullSecrets := getGeneratedDockercfgSecretNames(serviceAccount)
	if len(imageDockercfgPullSecrets) > 0 && len(mountableDockercfgSecrets) > 0 {
		return false
	}
	return true
}
func (e *DockercfgController) syncServiceAccount(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, exists, err := e.serviceAccountCache.GetByKey(key)
	if err != nil {
		klog.V(4).Infof("Unable to retrieve service account %v from store: %v", key, err)
		return err
	}
	if !exists {
		klog.V(4).Infof("Service account has been deleted %v", key)
		return nil
	}
	if !needsDockercfgSecret(obj.(*v1.ServiceAccount)) {
		return nil
	}
	serviceAccount := obj.(*v1.ServiceAccount).DeepCopyObject().(*v1.ServiceAccount)
	mountableDockercfgSecrets, imageDockercfgPullSecrets := getGeneratedDockercfgSecretNames(serviceAccount)
	foundPullSecret := len(imageDockercfgPullSecrets) > 0
	foundMountableSecret := len(mountableDockercfgSecrets) > 0
	if foundPullSecret || foundMountableSecret {
		switch {
		case foundPullSecret:
			serviceAccount.Secrets = append(serviceAccount.Secrets, v1.ObjectReference{Name: imageDockercfgPullSecrets.List()[0]})
		case foundMountableSecret:
			serviceAccount.ImagePullSecrets = append(serviceAccount.ImagePullSecrets, v1.LocalObjectReference{Name: mountableDockercfgSecrets.List()[0]})
		}
		delete(serviceAccount.Annotations, PendingTokenAnnotation)
		updatedSA, err := e.client.CoreV1().ServiceAccounts(serviceAccount.Namespace).Update(serviceAccount)
		if err == nil {
			e.serviceAccountCache.Mutation(updatedSA)
		}
		return err
	}
	dockercfgSecret, created, err := e.createDockerPullSecret(serviceAccount)
	if err != nil {
		return err
	}
	if !created {
		klog.V(5).Infof("The dockercfg secret was not created for service account %s/%s, will retry", serviceAccount.Namespace, serviceAccount.Name)
		return nil
	}
	first := true
	err = retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		if !first {
			obj, exists, err := e.serviceAccountCache.GetByKey(key)
			if err != nil {
				return err
			}
			if !exists || !needsDockercfgSecret(obj.(*v1.ServiceAccount)) || serviceAccount.UID != obj.(*v1.ServiceAccount).UID {
				klog.V(2).Infof("Deleting secret because the work is already done %s/%s", dockercfgSecret.Namespace, dockercfgSecret.Name)
				e.client.CoreV1().Secrets(dockercfgSecret.Namespace).Delete(dockercfgSecret.Name, nil)
				return nil
			}
			serviceAccount = obj.(*v1.ServiceAccount).DeepCopyObject().(*v1.ServiceAccount)
		}
		first = false
		serviceAccount.Secrets = append(serviceAccount.Secrets, v1.ObjectReference{Name: dockercfgSecret.Name})
		serviceAccount.ImagePullSecrets = append(serviceAccount.ImagePullSecrets, v1.LocalObjectReference{Name: dockercfgSecret.Name})
		delete(serviceAccount.Annotations, PendingTokenAnnotation)
		updatedSA, err := e.client.CoreV1().ServiceAccounts(serviceAccount.Namespace).Update(serviceAccount)
		if err == nil {
			e.serviceAccountCache.Mutation(updatedSA)
		}
		return err
	})
	if err != nil {
		klog.V(2).Infof("Deleting secret %s/%s (err=%v)", dockercfgSecret.Namespace, dockercfgSecret.Name, err)
		e.client.CoreV1().Secrets(dockercfgSecret.Namespace).Delete(dockercfgSecret.Name, nil)
	}
	return err
}
func (e *DockercfgController) createTokenSecret(serviceAccount *v1.ServiceAccount) (*v1.Secret, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pendingTokenName := serviceAccount.Annotations[PendingTokenAnnotation]
	if len(pendingTokenName) == 0 {
		pendingTokenName = secret.Strategy.GenerateName(getTokenSecretNamePrefix(serviceAccount.Name))
		if serviceAccount.Annotations == nil {
			serviceAccount.Annotations = map[string]string{}
		}
		serviceAccount.Annotations[PendingTokenAnnotation] = pendingTokenName
		updatedServiceAccount, err := e.client.CoreV1().ServiceAccounts(serviceAccount.Namespace).Update(serviceAccount)
		if kapierrors.IsConflict(err) {
			return nil, false, nil
		}
		if err != nil {
			return nil, false, err
		}
		serviceAccount = updatedServiceAccount
	}
	existingTokenSecretObj, exists, err := e.secretCache.GetByKey(serviceAccount.Namespace + "/" + pendingTokenName)
	if err != nil {
		return nil, false, err
	}
	if exists {
		existingTokenSecret := existingTokenSecretObj.(*v1.Secret)
		return existingTokenSecret, len(existingTokenSecret.Data[v1.ServiceAccountTokenKey]) > 0, nil
	}
	tokenSecret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: pendingTokenName, Namespace: serviceAccount.Namespace, Annotations: map[string]string{v1.ServiceAccountNameKey: serviceAccount.Name, v1.ServiceAccountUIDKey: string(serviceAccount.UID), DeprecatedKubeCreatedByAnnotation: CreateDockercfgSecretsController}}, Type: v1.SecretTypeServiceAccountToken, Data: map[string][]byte{}}
	klog.V(4).Infof("Creating token secret %q for service account %s/%s", tokenSecret.Name, serviceAccount.Namespace, serviceAccount.Name)
	token, err := e.client.CoreV1().Secrets(tokenSecret.Namespace).Create(tokenSecret)
	if kapierrors.IsAlreadyExists(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return token, len(token.Data[v1.ServiceAccountTokenKey]) > 0, nil
}
func (e *DockercfgController) createDockerPullSecret(serviceAccount *v1.ServiceAccount) (*v1.Secret, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokenSecret, isPopulated, err := e.createTokenSecret(serviceAccount)
	if err != nil {
		return nil, false, err
	}
	if !isPopulated {
		klog.V(5).Infof("Token secret for service account %s/%s is not populated yet", serviceAccount.Namespace, serviceAccount.Name)
		return nil, false, nil
	}
	dockercfgSecret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: secret.Strategy.GenerateName(getDockercfgSecretNamePrefix(serviceAccount.Name)), Namespace: tokenSecret.Namespace, Annotations: map[string]string{v1.ServiceAccountNameKey: serviceAccount.Name, v1.ServiceAccountUIDKey: string(serviceAccount.UID), ServiceAccountTokenSecretNameKey: string(tokenSecret.Name), ServiceAccountTokenValueAnnotation: string(tokenSecret.Data[v1.ServiceAccountTokenKey])}}, Type: v1.SecretTypeDockercfg, Data: map[string][]byte{}}
	klog.V(4).Infof("Creating dockercfg secret %q for service account %s/%s", dockercfgSecret.Name, serviceAccount.Namespace, serviceAccount.Name)
	e.dockerURLLock.Lock()
	defer e.dockerURLLock.Unlock()
	dockercfg := credentialprovider.DockerConfig{}
	for _, dockerURL := range e.dockerURLs {
		dockercfg[dockerURL] = credentialprovider.DockerConfigEntry{Username: "serviceaccount", Password: string(tokenSecret.Data[v1.ServiceAccountTokenKey]), Email: "serviceaccount@example.org"}
	}
	dockercfgContent, err := json.Marshal(&dockercfg)
	if err != nil {
		return nil, false, err
	}
	dockercfgSecret.Data[v1.DockerConfigKey] = dockercfgContent
	createdSecret, err := e.client.CoreV1().Secrets(tokenSecret.Namespace).Create(dockercfgSecret)
	return createdSecret, err == nil, err
}
func getGeneratedDockercfgSecretNames(serviceAccount *v1.ServiceAccount) (sets.String, sets.String) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mountableDockercfgSecrets := sets.String{}
	imageDockercfgPullSecrets := sets.String{}
	secretNamePrefix := getDockercfgSecretNamePrefix(serviceAccount.Name)
	for _, s := range serviceAccount.Secrets {
		if strings.HasPrefix(s.Name, secretNamePrefix) {
			mountableDockercfgSecrets.Insert(s.Name)
		}
	}
	for _, s := range serviceAccount.ImagePullSecrets {
		if strings.HasPrefix(s.Name, secretNamePrefix) {
			imageDockercfgPullSecrets.Insert(s.Name)
		}
	}
	return mountableDockercfgSecrets, imageDockercfgPullSecrets
}
func getDockercfgSecretNamePrefix(serviceAccountName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apihelpers.GetName(serviceAccountName, "dockercfg-", maxSecretPrefixNameLength)
}
func getTokenSecretNamePrefix(serviceAccountName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apihelpers.GetName(serviceAccountName, "token-", maxSecretPrefixNameLength)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
