package controllers

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
	"k8s.io/klog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	informers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/credentialprovider"
)

type DockerRegistryServiceControllerOptions struct {
	Resync			time.Duration
	ClusterDNSSuffix	string
	DockercfgController	*DockercfgController
	AdditionalRegistryURLs	[]string
	DockerURLsInitialized	chan struct{}
}
type serviceLocation struct {
	namespace	string
	name		string
}

var serviceLocations = []serviceLocation{{namespace: "default", name: "docker-registry"}, {namespace: "openshift-image-registry", name: "image-registry"}}

func NewDockerRegistryServiceController(secrets informers.SecretInformer, serviceInformer informers.ServiceInformer, cl kclientset.Interface, options DockerRegistryServiceControllerOptions) *DockerRegistryServiceController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := &DockerRegistryServiceController{client: cl, additionalRegistryURLs: options.AdditionalRegistryURLs, clusterDNSSuffix: options.ClusterDNSSuffix, dockercfgController: options.DockercfgController, registryLocationQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), secretsToUpdate: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), dockerURLsInitialized: options.DockerURLsInitialized}
	serviceInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Service:
			for _, location := range serviceLocations {
				if t.Namespace == location.namespace && t.Name == location.name {
					return true
				}
			}
		}
		return false
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		e.enqueueRegistryLocationQueue()
	}, UpdateFunc: func(old, cur interface{}) {
		e.enqueueRegistryLocationQueue()
	}, DeleteFunc: func(obj interface{}) {
		e.enqueueRegistryLocationQueue()
	}}})
	e.servicesSynced = serviceInformer.Informer().HasSynced
	e.serviceLister = serviceInformer.Lister()
	e.syncRegistryLocationHandler = e.syncRegistryLocationChange
	e.secretCache = secrets.Informer().GetIndexer()
	e.secretsSynced = secrets.Informer().GetController().HasSynced
	e.syncSecretHandler = e.syncSecretUpdate
	return e
}

type DockerRegistryServiceController struct {
	client				kclientset.Interface
	clusterDNSSuffix		string
	additionalRegistryURLs		[]string
	dockercfgController		*DockercfgController
	serviceLister			listers.ServiceLister
	servicesSynced			func() bool
	syncRegistryLocationHandler	func() error
	secretCache			cache.Store
	secretsSynced			func() bool
	syncSecretHandler		func(key string) error
	registryURLs			sets.String
	registryURLLock			sync.RWMutex
	registryLocationQueue		workqueue.RateLimitingInterface
	secretsToUpdate			workqueue.RateLimitingInterface
	dockerURLsInitialized		chan struct{}
	initialSecretsCheckDone		bool
}

func (e *DockerRegistryServiceController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer e.registryLocationQueue.ShutDown()
	klog.Infof("Starting DockerRegistryServiceController controller")
	defer klog.Infof("Shutting down DockerRegistryServiceController controller")
	ready := make(chan struct{})
	go e.waitForDockerURLs(ready, stopCh)
	select {
	case <-ready:
	case <-stopCh:
		return
	}
	klog.V(1).Infof("caches synced")
	go wait.Until(e.watchForDockerURLChanges, time.Second, stopCh)
	for i := 0; i < workers; i++ {
		go wait.Until(e.watchForDockercfgSecretUpdates, time.Second, stopCh)
	}
	<-stopCh
}
func (e *DockerRegistryServiceController) enqueueRegistryLocationQueue() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.registryLocationQueue.Add("check")
}
func (e *DockerRegistryServiceController) waitForDockerURLs(ready chan<- struct{}, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	if !cache.WaitForCacheSync(stopCh, e.servicesSynced, e.secretsSynced) {
		return
	}
	urls := e.getDockerRegistryLocations()
	e.setRegistryURLs(urls...)
	e.dockercfgController.SetDockerURLs(urls...)
	close(e.dockerURLsInitialized)
	close(ready)
	return
}
func (e *DockerRegistryServiceController) setRegistryURLs(registryURLs ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.registryURLLock.Lock()
	defer e.registryURLLock.Unlock()
	e.registryURLs = sets.NewString(registryURLs...)
}
func (e *DockerRegistryServiceController) getRegistryURLs() sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.registryURLLock.RLock()
	defer e.registryURLLock.RUnlock()
	return sets.NewString(e.registryURLs.List()...)
}
func (e *DockerRegistryServiceController) watchForDockerURLChanges() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	workFn := func() bool {
		key, quit := e.registryLocationQueue.Get()
		if quit {
			return true
		}
		defer e.registryLocationQueue.Done(key)
		if err := e.syncRegistryLocationHandler(); err == nil {
			e.registryLocationQueue.Forget(key)
		} else {
			utilruntime.HandleError(fmt.Errorf("error syncing service, it will be retried: %v", err))
			e.registryLocationQueue.AddRateLimited(key)
		}
		return false
	}
	for {
		if workFn() {
			return
		}
	}
}
func (e *DockerRegistryServiceController) getDockerRegistryLocations() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := append([]string{}, e.additionalRegistryURLs...)
	for _, location := range serviceLocations {
		ret = append(ret, getDockerRegistryLocations(e.serviceLister, location, e.clusterDNSSuffix)...)
	}
	klog.V(4).Infof("found docker registry urls: %v", ret)
	return ret
}
func getDockerRegistryLocations(lister listers.ServiceLister, location serviceLocation, clusterDNSSuffix string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	service, err := lister.Services(location.namespace).Get(location.name)
	if err != nil {
		return []string{}
	}
	hasClusterIP := (len(service.Spec.ClusterIP) > 0) && (net.ParseIP(service.Spec.ClusterIP) != nil)
	if hasClusterIP && len(service.Spec.Ports) > 0 {
		ret := []string{net.JoinHostPort(service.Spec.ClusterIP, fmt.Sprintf("%d", service.Spec.Ports[0].Port)), net.JoinHostPort(fmt.Sprintf("%s.%s.svc", service.Name, service.Namespace), fmt.Sprintf("%d", service.Spec.Ports[0].Port))}
		if len(clusterDNSSuffix) > 0 {
			ret = append(ret, net.JoinHostPort(fmt.Sprintf("%s.%s.svc."+clusterDNSSuffix, service.Name, service.Namespace), fmt.Sprintf("%d", service.Spec.Ports[0].Port)))
		}
		return ret
	}
	return []string{}
}
func (e *DockerRegistryServiceController) syncRegistryLocationChange() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newLocations := e.getDockerRegistryLocations()
	newDockerRegistryLocations := sets.NewString(newLocations...)
	existingURLs := e.getRegistryURLs()
	if existingURLs.Equal(newDockerRegistryLocations) && e.initialSecretsCheckDone {
		klog.V(3).Infof("No effective update: %v", newDockerRegistryLocations)
		return nil
	}
	klog.V(1).Infof("Updating registry URLs from %v to %v", existingURLs, newDockerRegistryLocations)
	e.dockercfgController.SetDockerURLs(newDockerRegistryLocations.List()...)
	e.setRegistryURLs(newDockerRegistryLocations.List()...)
	e.initialSecretsCheckDone = true
	for _, obj := range e.secretCache.List() {
		switch t := obj.(type) {
		case *v1.Secret:
			if t.Type != v1.SecretTypeDockercfg {
				continue
			}
			if t.Annotations == nil {
				continue
			}
			if _, hasTokenSecret := t.Annotations[ServiceAccountTokenSecretNameKey]; !hasTokenSecret {
				continue
			}
		default:
			utilruntime.HandleError(fmt.Errorf("object passed to %T that is not expected: %T", e, obj))
			continue
		}
		key, err := controller.KeyFunc(obj)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", obj, err))
			continue
		}
		e.secretsToUpdate.Add(key)
	}
	return nil
}
func (e *DockerRegistryServiceController) watchForDockercfgSecretUpdates() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	workFn := func() bool {
		key, quit := e.secretsToUpdate.Get()
		if quit {
			return true
		}
		defer e.secretsToUpdate.Done(key)
		if err := e.syncSecretHandler(key.(string)); err == nil {
			e.secretsToUpdate.Forget(key)
		} else {
			utilruntime.HandleError(fmt.Errorf("error syncing service, it will be retried: %v", err))
			e.secretsToUpdate.AddRateLimited(key)
		}
		return false
	}
	for {
		if workFn() {
			return
		}
	}
}
func (e *DockerRegistryServiceController) syncSecretUpdate(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := e.secretCache.GetByKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to retrieve secret %v from store: %v", key, err))
		return err
	}
	if !exists {
		return nil
	}
	dockerRegistryURLs := e.getRegistryURLs()
	sharedDockercfgSecret := obj.(*v1.Secret)
	dockercfg := &credentialprovider.DockerConfig{}
	json.Unmarshal(sharedDockercfgSecret.Data[v1.DockerConfigKey], dockercfg)
	dockercfgMap := map[string]credentialprovider.DockerConfigEntry(*dockercfg)
	existingDockercfgSecretLocations := sets.StringKeySet(dockercfgMap)
	if existingDockercfgSecretLocations.Equal(dockerRegistryURLs) {
		return nil
	}
	dockercfgSecret := obj.(runtime.Object).DeepCopyObject().(*v1.Secret)
	dockerCredentials := dockercfgSecret.Annotations[ServiceAccountTokenValueAnnotation]
	if len(dockerCredentials) == 0 && len(existingDockercfgSecretLocations) > 0 {
		dockerCredentials = dockercfgMap[existingDockercfgSecretLocations.List()[0]].Password
	}
	if len(dockerCredentials) == 0 {
		tokenSecretKey := dockercfgSecret.Namespace + "/" + dockercfgSecret.Annotations[ServiceAccountTokenSecretNameKey]
		tokenSecret, exists, err := e.secretCache.GetByKey(tokenSecretKey)
		if !exists {
			utilruntime.HandleError(fmt.Errorf("cannot determine SA token due to missing secret: %v", tokenSecretKey))
			return nil
		}
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("cannot determine SA token: %v", err))
			return nil
		}
		dockerCredentials = string(tokenSecret.(*v1.Secret).Data[v1.ServiceAccountTokenKey])
	}
	newDockercfgMap := credentialprovider.DockerConfig{}
	for key := range dockerRegistryURLs {
		newDockercfgMap[key] = credentialprovider.DockerConfigEntry{Username: "serviceaccount", Password: dockerCredentials, Email: "serviceaccount@example.org"}
	}
	dockercfgContent, err := json.Marshal(&newDockercfgMap)
	if err != nil {
		utilruntime.HandleError(err)
		return nil
	}
	dockercfgSecret.Data[v1.DockerConfigKey] = dockercfgContent
	if _, err := e.client.CoreV1().Secrets(dockercfgSecret.Namespace).Update(dockercfgSecret); err != nil {
		return err
	}
	return nil
}
