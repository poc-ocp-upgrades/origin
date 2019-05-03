package controllers

import (
	"fmt"
	"k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"time"
)

const NumServiceAccountUpdateRetries = 10

type DockercfgDeletedControllerOptions struct{ Resync time.Duration }

func NewDockercfgDeletedController(secrets informers.SecretInformer, cl kclientset.Interface, options DockercfgDeletedControllerOptions) *DockercfgDeletedController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := &DockercfgDeletedController{client: cl}
	e.secretController = secrets.Informer().GetController()
	secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Secret:
			return t.Type == v1.SecretTypeDockercfg
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{DeleteFunc: e.secretDeleted}}, options.Resync)
	return e
}

type DockercfgDeletedController struct {
	client           kclientset.Interface
	secretController cache.Controller
}

func (e *DockercfgDeletedController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	klog.Infof("Starting DockercfgDeletedController controller")
	defer klog.Infof("Shutting down DockercfgDeletedController controller")
	if !cache.WaitForCacheSync(stopCh, e.secretController.HasSynced) {
		return
	}
	klog.V(1).Infof("caches synced")
	<-stopCh
}
func (e *DockercfgDeletedController) secretDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dockercfgSecret, ok := obj.(*v1.Secret)
	if !ok {
		return
	}
	if _, exists := dockercfgSecret.Annotations[ServiceAccountTokenSecretNameKey]; !exists {
		return
	}
	for i := 1; i <= NumServiceAccountUpdateRetries; i++ {
		if err := e.removeDockercfgSecretReference(dockercfgSecret); err != nil {
			if kapierrors.IsConflict(err) && i < NumServiceAccountUpdateRetries {
				time.Sleep(wait.Jitter(100*time.Millisecond, 0.0))
				continue
			}
			klog.Error(err)
			break
		}
		break
	}
	if err := e.client.CoreV1().Secrets(dockercfgSecret.Namespace).Delete(dockercfgSecret.Annotations[ServiceAccountTokenSecretNameKey], nil); (err != nil) && !kapierrors.IsNotFound(err) {
		utilruntime.HandleError(err)
	}
}
func (e *DockercfgDeletedController) removeDockercfgSecretReference(dockercfgSecret *v1.Secret) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceAccount, err := e.getServiceAccount(dockercfgSecret)
	if kapierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	changed := false
	secrets := []v1.ObjectReference{}
	for _, s := range serviceAccount.Secrets {
		if s.Name == dockercfgSecret.Name {
			changed = true
			continue
		}
		secrets = append(secrets, s)
	}
	serviceAccount.Secrets = secrets
	imagePullSecrets := []v1.LocalObjectReference{}
	for _, s := range serviceAccount.ImagePullSecrets {
		if s.Name == dockercfgSecret.Name {
			changed = true
			continue
		}
		imagePullSecrets = append(imagePullSecrets, s)
	}
	serviceAccount.ImagePullSecrets = imagePullSecrets
	if changed {
		_, err = e.client.CoreV1().ServiceAccounts(dockercfgSecret.Namespace).Update(serviceAccount)
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *DockercfgDeletedController) getServiceAccount(secret *v1.Secret) (*v1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	saName, saUID := secret.Annotations[v1.ServiceAccountNameKey], secret.Annotations[v1.ServiceAccountUIDKey]
	if len(saName) == 0 || len(saUID) == 0 {
		return nil, nil
	}
	serviceAccount, err := e.client.CoreV1().ServiceAccounts(secret.Namespace).Get(saName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if saUID != string(serviceAccount.UID) {
		return nil, fmt.Errorf("secret (%v) service account UID (%v) does not match service account (%v) UID (%v)", secret.Name, saUID, serviceAccount.Name, serviceAccount.UID)
	}
	return serviceAccount, nil
}
