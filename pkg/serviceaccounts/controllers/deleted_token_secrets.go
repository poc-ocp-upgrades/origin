package controllers

import (
	"fmt"
	"time"
	"k8s.io/klog"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	informers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type DockercfgTokenDeletedControllerOptions struct{ Resync time.Duration }

func NewDockercfgTokenDeletedController(secrets informers.SecretInformer, cl kclientset.Interface, options DockercfgTokenDeletedControllerOptions) *DockercfgTokenDeletedController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := &DockercfgTokenDeletedController{client: cl}
	e.secretController = secrets.Informer().GetController()
	secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Secret:
			return t.Type == v1.SecretTypeServiceAccountToken
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{DeleteFunc: e.secretDeleted}}, options.Resync)
	return e
}

type DockercfgTokenDeletedController struct {
	client			kclientset.Interface
	secretController	cache.Controller
}

func (e *DockercfgTokenDeletedController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	klog.Infof("Starting DockercfgTokenDeletedController controller")
	defer klog.Infof("Shutting down DockercfgTokenDeletedController controller")
	if !cache.WaitForCacheSync(stopCh, e.secretController.HasSynced) {
		return
	}
	klog.V(1).Infof("caches synced")
	<-stopCh
}
func (e *DockercfgTokenDeletedController) secretDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tokenSecret, ok := obj.(*v1.Secret)
	if !ok {
		return
	}
	dockercfgSecrets, err := e.findDockercfgSecrets(tokenSecret)
	if err != nil {
		klog.Error(err)
		return
	}
	if len(dockercfgSecrets) == 0 {
		return
	}
	for _, dockercfgSecret := range dockercfgSecrets {
		if err := e.client.CoreV1().Secrets(dockercfgSecret.Namespace).Delete(dockercfgSecret.Name, nil); (err != nil) && !apierrors.IsNotFound(err) {
			utilruntime.HandleError(err)
		}
	}
}
func (e *DockercfgTokenDeletedController) findDockercfgSecrets(tokenSecret *v1.Secret) ([]*v1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dockercfgSecrets := []*v1.Secret{}
	options := metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(api.SecretTypeField, string(v1.SecretTypeDockercfg)).String()}
	potentialSecrets, err := e.client.CoreV1().Secrets(tokenSecret.Namespace).List(options)
	if err != nil {
		return nil, err
	}
	for i, currSecret := range potentialSecrets.Items {
		if currSecret.Annotations[ServiceAccountTokenSecretNameKey] == tokenSecret.Name {
			dockercfgSecrets = append(dockercfgSecrets, &potentialSecrets.Items[i])
		}
	}
	return dockercfgSecrets, nil
}
