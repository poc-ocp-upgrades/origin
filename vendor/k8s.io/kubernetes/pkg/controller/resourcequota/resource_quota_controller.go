package resourcequota

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

type NamespacedResourcesFunc func() ([]*metav1.APIResourceList, error)
type ReplenishmentFunc func(groupResource schema.GroupResource, namespace string)
type InformerFactory interface {
	ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error)
	Start(stopCh <-chan struct{})
}
type ResourceQuotaControllerOptions struct {
	QuotaClient               corev1client.ResourceQuotasGetter
	ResourceQuotaInformer     coreinformers.ResourceQuotaInformer
	ResyncPeriod              controller.ResyncPeriodFunc
	Registry                  quota.Registry
	DiscoveryFunc             NamespacedResourcesFunc
	IgnoredResourcesFunc      func() map[schema.GroupResource]struct{}
	InformersStarted          <-chan struct{}
	InformerFactory           InformerFactory
	ReplenishmentResyncPeriod controller.ResyncPeriodFunc
}
type ResourceQuotaController struct {
	rqClient            corev1client.ResourceQuotasGetter
	rqLister            corelisters.ResourceQuotaLister
	informerSyncedFuncs []cache.InformerSynced
	queue               workqueue.RateLimitingInterface
	missingUsageQueue   workqueue.RateLimitingInterface
	syncHandler         func(key string) error
	resyncPeriod        controller.ResyncPeriodFunc
	registry            quota.Registry
	quotaMonitor        *QuotaMonitor
	workerLock          sync.RWMutex
}

func NewResourceQuotaController(options *ResourceQuotaControllerOptions) (*ResourceQuotaController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rq := &ResourceQuotaController{rqClient: options.QuotaClient, rqLister: options.ResourceQuotaInformer.Lister(), informerSyncedFuncs: []cache.InformerSynced{options.ResourceQuotaInformer.Informer().HasSynced}, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "resourcequota_primary"), missingUsageQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "resourcequota_priority"), resyncPeriod: options.ResyncPeriod, registry: options.Registry}
	rq.syncHandler = rq.syncResourceQuotaFromKey
	options.ResourceQuotaInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: rq.addQuota, UpdateFunc: func(old, cur interface{}) {
		oldResourceQuota := old.(*v1.ResourceQuota)
		curResourceQuota := cur.(*v1.ResourceQuota)
		if quota.V1Equals(oldResourceQuota.Spec.Hard, curResourceQuota.Spec.Hard) {
			return
		}
		rq.addQuota(curResourceQuota)
	}, DeleteFunc: rq.enqueueResourceQuota}, rq.resyncPeriod())
	if options.DiscoveryFunc != nil {
		qm := &QuotaMonitor{informersStarted: options.InformersStarted, informerFactory: options.InformerFactory, ignoredResources: options.IgnoredResourcesFunc(), resourceChanges: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "resource_quota_controller_resource_changes"), resyncPeriod: options.ReplenishmentResyncPeriod, replenishmentFunc: rq.replenishQuota, registry: rq.registry}
		rq.quotaMonitor = qm
		resources, err := GetQuotableResources(options.DiscoveryFunc)
		if discovery.IsGroupDiscoveryFailedError(err) {
			utilruntime.HandleError(fmt.Errorf("initial discovery check failure, continuing and counting on future sync update: %v", err))
		} else if err != nil {
			return nil, err
		}
		if err = qm.SyncMonitors(resources); err != nil {
			utilruntime.HandleError(fmt.Errorf("initial monitor sync has error: %v", err))
		}
		rq.informerSyncedFuncs = append(rq.informerSyncedFuncs, qm.IsSynced)
	}
	return rq, nil
}
func (rq *ResourceQuotaController) enqueueAll() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer klog.V(4).Infof("Resource quota controller queued all resource quota for full calculation of usage")
	rqs, err := rq.rqLister.List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to enqueue all - error listing resource quotas: %v", err))
		return
	}
	for i := range rqs {
		key, err := controller.KeyFunc(rqs[i])
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", rqs[i], err))
			continue
		}
		rq.queue.Add(key)
	}
}
func (rq *ResourceQuotaController) enqueueResourceQuota(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	rq.queue.Add(key)
}
func (rq *ResourceQuotaController) addQuota(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	resourceQuota := obj.(*v1.ResourceQuota)
	if !apiequality.Semantic.DeepEqual(resourceQuota.Spec.Hard, resourceQuota.Status.Hard) {
		rq.missingUsageQueue.Add(key)
		return
	}
	for constraint := range resourceQuota.Status.Hard {
		if _, usageFound := resourceQuota.Status.Used[constraint]; !usageFound {
			matchedResources := []v1.ResourceName{v1.ResourceName(constraint)}
			for _, evaluator := range rq.registry.List() {
				if intersection := evaluator.MatchingResources(matchedResources); len(intersection) > 0 {
					rq.missingUsageQueue.Add(key)
					return
				}
			}
		}
	}
	rq.queue.Add(key)
}
func (rq *ResourceQuotaController) worker(queue workqueue.RateLimitingInterface) func() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	workFunc := func() bool {
		key, quit := queue.Get()
		if quit {
			return true
		}
		defer queue.Done(key)
		rq.workerLock.RLock()
		defer rq.workerLock.RUnlock()
		err := rq.syncHandler(key.(string))
		if err == nil {
			queue.Forget(key)
			return false
		}
		utilruntime.HandleError(err)
		queue.AddRateLimited(key)
		return false
	}
	return func() {
		for {
			if quit := workFunc(); quit {
				klog.Infof("resource quota controller worker shutting down")
				return
			}
		}
	}
}
func (rq *ResourceQuotaController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer rq.queue.ShutDown()
	klog.Infof("Starting resource quota controller")
	defer klog.Infof("Shutting down resource quota controller")
	if rq.quotaMonitor != nil {
		go rq.quotaMonitor.Run(stopCh)
	}
	if !controller.WaitForCacheSync("resource quota", stopCh, rq.informerSyncedFuncs...) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(rq.worker(rq.queue), time.Second, stopCh)
		go wait.Until(rq.worker(rq.missingUsageQueue), time.Second, stopCh)
	}
	go wait.Until(func() {
		rq.enqueueAll()
	}, rq.resyncPeriod(), stopCh)
	<-stopCh
}
func (rq *ResourceQuotaController) syncResourceQuotaFromKey(key string) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing resource quota %q (%v)", key, time.Since(startTime))
	}()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	quota, err := rq.rqLister.ResourceQuotas(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Infof("Resource quota has been deleted %v", key)
		return nil
	}
	if err != nil {
		klog.Infof("Unable to retrieve resource quota %v from store: %v", key, err)
		return err
	}
	return rq.syncResourceQuota(quota)
}
func (rq *ResourceQuotaController) syncResourceQuota(resourceQuota *v1.ResourceQuota) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	statusLimitsDirty := !apiequality.Semantic.DeepEqual(resourceQuota.Spec.Hard, resourceQuota.Status.Hard)
	dirty := statusLimitsDirty || resourceQuota.Status.Hard == nil || resourceQuota.Status.Used == nil
	used := v1.ResourceList{}
	if resourceQuota.Status.Used != nil {
		used = quota.Add(v1.ResourceList{}, resourceQuota.Status.Used)
	}
	hardLimits := quota.Add(v1.ResourceList{}, resourceQuota.Spec.Hard)
	errors := []error{}
	newUsage, err := quota.CalculateUsage(resourceQuota.Namespace, resourceQuota.Spec.Scopes, hardLimits, rq.registry, resourceQuota.Spec.ScopeSelector)
	if err != nil {
		errors = append(errors, err)
	}
	for key, value := range newUsage {
		used[key] = value
	}
	hardResources := quota.ResourceNames(hardLimits)
	used = quota.Mask(used, hardResources)
	usage := resourceQuota.DeepCopy()
	usage.Status = v1.ResourceQuotaStatus{Hard: hardLimits, Used: used}
	dirty = dirty || !quota.Equals(usage.Status.Used, resourceQuota.Status.Used)
	if dirty {
		_, err = rq.rqClient.ResourceQuotas(usage.Namespace).UpdateStatus(usage)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return utilerrors.NewAggregate(errors)
}
func (rq *ResourceQuotaController) replenishQuota(groupResource schema.GroupResource, namespace string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	evaluator := rq.registry.Get(groupResource)
	if evaluator == nil {
		return
	}
	resourceQuotas, err := rq.rqLister.ResourceQuotas(namespace).List(labels.Everything())
	if errors.IsNotFound(err) {
		utilruntime.HandleError(fmt.Errorf("quota controller could not find ResourceQuota associated with namespace: %s, could take up to %v before a quota replenishes", namespace, rq.resyncPeriod()))
		return
	}
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error checking to see if namespace %s has any ResourceQuota associated with it: %v", namespace, err))
		return
	}
	if len(resourceQuotas) == 0 {
		return
	}
	for i := range resourceQuotas {
		resourceQuota := resourceQuotas[i]
		resourceQuotaResources := quota.ResourceNames(resourceQuota.Status.Hard)
		if intersection := evaluator.MatchingResources(resourceQuotaResources); len(intersection) > 0 {
			rq.enqueueResourceQuota(resourceQuota)
		}
	}
}
func (rq *ResourceQuotaController) Sync(discoveryFunc NamespacedResourcesFunc, period time.Duration, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldResources := make(map[schema.GroupVersionResource]struct{})
	wait.Until(func() {
		newResources, err := GetQuotableResources(discoveryFunc)
		if err != nil {
			utilruntime.HandleError(err)
			if discovery.IsGroupDiscoveryFailedError(err) && len(newResources) > 0 {
				for k, v := range oldResources {
					newResources[k] = v
				}
			} else {
				return
			}
		}
		if reflect.DeepEqual(oldResources, newResources) {
			klog.V(4).Infof("no resource updates from discovery, skipping resource quota sync")
			return
		}
		rq.workerLock.Lock()
		defer rq.workerLock.Unlock()
		if klog.V(2) {
			klog.Infof("syncing resource quota controller with updated resources from discovery: %s", printDiff(oldResources, newResources))
		}
		if err := rq.resyncMonitors(newResources); err != nil {
			utilruntime.HandleError(fmt.Errorf("failed to sync resource monitors: %v", err))
			return
		}
		if rq.quotaMonitor != nil && !controller.WaitForCacheSync("resource quota", waitForStopOrTimeout(stopCh, period), rq.quotaMonitor.IsSynced) {
			utilruntime.HandleError(fmt.Errorf("timed out waiting for quota monitor sync"))
			return
		}
		oldResources = newResources
		klog.V(2).Infof("synced quota controller")
	}, period, stopCh)
}
func printDiff(oldResources, newResources map[schema.GroupVersionResource]struct{}) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	removed := sets.NewString()
	for oldResource := range oldResources {
		if _, ok := newResources[oldResource]; !ok {
			removed.Insert(fmt.Sprintf("%+v", oldResource))
		}
	}
	added := sets.NewString()
	for newResource := range newResources {
		if _, ok := oldResources[newResource]; !ok {
			added.Insert(fmt.Sprintf("%+v", newResource))
		}
	}
	return fmt.Sprintf("added: %v, removed: %v", added.List(), removed.List())
}
func waitForStopOrTimeout(stopCh <-chan struct{}, timeout time.Duration) <-chan struct{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stopChWithTimeout := make(chan struct{})
	go func() {
		defer close(stopChWithTimeout)
		select {
		case <-stopCh:
		case <-time.After(timeout):
		}
	}()
	return stopChWithTimeout
}
func (rq *ResourceQuotaController) resyncMonitors(resources map[schema.GroupVersionResource]struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rq.quotaMonitor == nil {
		return nil
	}
	if err := rq.quotaMonitor.SyncMonitors(resources); err != nil {
		return err
	}
	rq.quotaMonitor.StartMonitors()
	return nil
}
func GetQuotableResources(discoveryFunc NamespacedResourcesFunc) (map[schema.GroupVersionResource]struct{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	possibleResources, discoveryErr := discoveryFunc()
	if discoveryErr != nil && len(possibleResources) == 0 {
		return nil, fmt.Errorf("failed to discover resources: %v", discoveryErr)
	}
	quotableResources := discovery.FilteredBy(discovery.SupportsAllVerbs{Verbs: []string{"create", "list", "watch", "delete"}}, possibleResources)
	quotableGroupVersionResources, err := discovery.GroupVersionResources(quotableResources)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse resources: %v", err)
	}
	return quotableGroupVersionResources, discoveryErr
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
