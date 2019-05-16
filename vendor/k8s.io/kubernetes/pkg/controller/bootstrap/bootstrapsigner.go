package bootstrap

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

type BootstrapSignerOptions struct {
	ConfigMapNamespace   string
	ConfigMapName        string
	TokenSecretNamespace string
	ConfigMapResync      time.Duration
	SecretResync         time.Duration
}

func DefaultBootstrapSignerOptions() BootstrapSignerOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return BootstrapSignerOptions{ConfigMapNamespace: api.NamespacePublic, ConfigMapName: bootstrapapi.ConfigMapClusterInfo, TokenSecretNamespace: api.NamespaceSystem}
}

type BootstrapSigner struct {
	client             clientset.Interface
	configMapKey       string
	configMapName      string
	configMapNamespace string
	secretNamespace    string
	syncQueue          workqueue.RateLimitingInterface
	secretLister       corelisters.SecretLister
	secretSynced       cache.InformerSynced
	configMapLister    corelisters.ConfigMapLister
	configMapSynced    cache.InformerSynced
}

func NewBootstrapSigner(cl clientset.Interface, secrets informers.SecretInformer, configMaps informers.ConfigMapInformer, options BootstrapSignerOptions) (*BootstrapSigner, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e := &BootstrapSigner{client: cl, configMapKey: options.ConfigMapNamespace + "/" + options.ConfigMapName, configMapName: options.ConfigMapName, configMapNamespace: options.ConfigMapNamespace, secretNamespace: options.TokenSecretNamespace, secretLister: secrets.Lister(), secretSynced: secrets.Informer().HasSynced, configMapLister: configMaps.Lister(), configMapSynced: configMaps.Informer().HasSynced, syncQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "bootstrap_signer_queue")}
	if cl.CoreV1().RESTClient().GetRateLimiter() != nil {
		if err := metrics.RegisterMetricAndTrackRateLimiterUsage("bootstrap_signer", cl.CoreV1().RESTClient().GetRateLimiter()); err != nil {
			return nil, err
		}
	}
	configMaps.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.ConfigMap:
			return t.Name == options.ConfigMapName && t.Namespace == options.ConfigMapNamespace
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: func(_ interface{}) {
		e.pokeConfigMapSync()
	}, UpdateFunc: func(_, _ interface{}) {
		e.pokeConfigMapSync()
	}}}, options.ConfigMapResync)
	secrets.Informer().AddEventHandlerWithResyncPeriod(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Secret:
			return t.Type == bootstrapapi.SecretTypeBootstrapToken && t.Namespace == e.secretNamespace
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: func(_ interface{}) {
		e.pokeConfigMapSync()
	}, UpdateFunc: func(_, _ interface{}) {
		e.pokeConfigMapSync()
	}, DeleteFunc: func(_ interface{}) {
		e.pokeConfigMapSync()
	}}}, options.SecretResync)
	return e, nil
}
func (e *BootstrapSigner) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer e.syncQueue.ShutDown()
	if !controller.WaitForCacheSync("bootstrap_signer", stopCh, e.configMapSynced, e.secretSynced) {
		return
	}
	klog.V(5).Infof("Starting workers")
	go wait.Until(e.serviceConfigMapQueue, 0, stopCh)
	<-stopCh
	klog.V(1).Infof("Shutting down")
}
func (e *BootstrapSigner) pokeConfigMapSync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.syncQueue.Add(e.configMapKey)
}
func (e *BootstrapSigner) serviceConfigMapQueue() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := e.syncQueue.Get()
	if quit {
		return
	}
	defer e.syncQueue.Done(key)
	e.signConfigMap()
}
func (e *BootstrapSigner) signConfigMap() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	origCM := e.getConfigMap()
	if origCM == nil {
		return
	}
	var needUpdate = false
	newCM := origCM.DeepCopy()
	content, ok := newCM.Data[bootstrapapi.KubeConfigKey]
	if !ok {
		klog.V(3).Infof("No %s key in %s/%s ConfigMap", bootstrapapi.KubeConfigKey, origCM.Namespace, origCM.Name)
		return
	}
	sigs := map[string]string{}
	for key, value := range newCM.Data {
		if strings.HasPrefix(key, bootstrapapi.JWSSignatureKeyPrefix) {
			tokenID := strings.TrimPrefix(key, bootstrapapi.JWSSignatureKeyPrefix)
			sigs[tokenID] = value
			delete(newCM.Data, key)
		}
	}
	tokens := e.getTokens()
	for tokenID, tokenValue := range tokens {
		sig, err := computeDetachedSig(content, tokenID, tokenValue)
		if err != nil {
			utilruntime.HandleError(err)
		}
		oldSig, _ := sigs[tokenID]
		if sig != oldSig {
			needUpdate = true
		}
		delete(sigs, tokenID)
		newCM.Data[bootstrapapi.JWSSignatureKeyPrefix+tokenID] = sig
	}
	if len(sigs) != 0 {
		needUpdate = true
	}
	if needUpdate {
		e.updateConfigMap(newCM)
	}
}
func (e *BootstrapSigner) updateConfigMap(cm *v1.ConfigMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := e.client.CoreV1().ConfigMaps(cm.Namespace).Update(cm)
	if err != nil && !apierrors.IsConflict(err) && !apierrors.IsNotFound(err) {
		klog.V(3).Infof("Error updating ConfigMap: %v", err)
	}
}
func (e *BootstrapSigner) getConfigMap() *v1.ConfigMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMap, err := e.configMapLister.ConfigMaps(e.configMapNamespace).Get(e.configMapName)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			utilruntime.HandleError(err)
		}
		return nil
	}
	return configMap
}
func (e *BootstrapSigner) listSecrets() []*v1.Secret {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secrets, err := e.secretLister.Secrets(e.secretNamespace).List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(err)
		return nil
	}
	items := []*v1.Secret{}
	for _, secret := range secrets {
		if secret.Type == bootstrapapi.SecretTypeBootstrapToken {
			items = append(items, secret)
		}
	}
	return items
}
func (e *BootstrapSigner) getTokens() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := map[string]string{}
	secretObjs := e.listSecrets()
	for _, secret := range secretObjs {
		tokenID, tokenSecret, ok := validateSecretForSigning(secret)
		if !ok {
			continue
		}
		if _, ok := ret[tokenID]; ok {
			klog.V(1).Infof("Duplicate bootstrap tokens found for id %s, ignoring on in %s/%s", tokenID, secret.Namespace, secret.Name)
			continue
		}
		ret[tokenID] = tokenSecret
	}
	return ret
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
