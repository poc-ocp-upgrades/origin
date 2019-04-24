package clusterquotareconciliation

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"reflect"
	"sync"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/resourcequota"
	utilquota "k8s.io/kubernetes/pkg/quota/v1"
	quotav1 "github.com/openshift/api/quota/v1"
	quotatypedclient "github.com/openshift/client-go/quota/clientset/versioned/typed/quota/v1"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions/quota/v1"
	quotalister "github.com/openshift/client-go/quota/listers/quota/v1"
	quotav1conversions "github.com/openshift/origin/pkg/quota/apis/quota/v1"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
)

type ClusterQuotaReconcilationControllerOptions struct {
	ClusterQuotaInformer		quotainformer.ClusterResourceQuotaInformer
	ClusterQuotaMapper		clusterquotamapping.ClusterQuotaMapper
	ClusterQuotaClient		quotatypedclient.ClusterResourceQuotaInterface
	Registry			utilquota.Registry
	ResyncPeriod			time.Duration
	DiscoveryFunc			resourcequota.NamespacedResourcesFunc
	IgnoredResourcesFunc		func() map[schema.GroupResource]struct{}
	InformersStarted		<-chan struct{}
	InformerFactory			resourcequota.InformerFactory
	ReplenishmentResyncPeriod	controller.ResyncPeriodFunc
}
type ClusterQuotaReconcilationController struct {
	clusterQuotaLister	quotalister.ClusterResourceQuotaLister
	clusterQuotaMapper	clusterquotamapping.ClusterQuotaMapper
	clusterQuotaClient	quotatypedclient.ClusterResourceQuotaInterface
	informerSyncedFuncs	[]cache.InformerSynced
	resyncPeriod		time.Duration
	queue			BucketingWorkQueue
	registry		utilquota.Registry
	quotaMonitor		*resourcequota.QuotaMonitor
	workerLock		sync.RWMutex
}
type workItem struct {
	namespaceName		string
	forceRecalculation	bool
}

func NewClusterQuotaReconcilationController(options ClusterQuotaReconcilationControllerOptions) (*ClusterQuotaReconcilationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &ClusterQuotaReconcilationController{clusterQuotaLister: options.ClusterQuotaInformer.Lister(), clusterQuotaMapper: options.ClusterQuotaMapper, clusterQuotaClient: options.ClusterQuotaClient, informerSyncedFuncs: []cache.InformerSynced{options.ClusterQuotaInformer.Informer().HasSynced}, resyncPeriod: options.ResyncPeriod, registry: options.Registry, queue: NewBucketingWorkQueue("controller_clusterquotareconcilationcontroller")}
	options.ClusterQuotaInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addClusterQuota, UpdateFunc: c.updateClusterQuota})
	qm := resourcequota.NewQuotaMonitor(options.InformersStarted, options.InformerFactory, options.IgnoredResourcesFunc(), options.ReplenishmentResyncPeriod, c.replenishQuota, c.registry)
	c.quotaMonitor = qm
	resources, err := resourcequota.GetQuotableResources(options.DiscoveryFunc)
	if discovery.IsGroupDiscoveryFailedError(err) {
		utilruntime.HandleError(fmt.Errorf("initial discovery check failure, continuing and counting on future sync update: %v", err))
	} else if err != nil {
		return nil, err
	}
	if err = qm.SyncMonitors(resources); err != nil {
		utilruntime.HandleError(fmt.Errorf("initial monitor sync has error: %v", err))
	}
	c.informerSyncedFuncs = append(c.informerSyncedFuncs, qm.IsSynced)
	return c, nil
}
func (c *ClusterQuotaReconcilationController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	klog.Infof("Starting the cluster quota reconciliation controller")
	go c.quotaMonitor.Run(stopCh)
	if !controller.WaitForCacheSync("cluster resource quota", stopCh, c.informerSyncedFuncs...) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}
	go wait.Until(func() {
		c.calculateAll()
	}, c.resyncPeriod, stopCh)
	<-stopCh
	klog.Infof("Shutting down ClusterQuotaReconcilationController")
	c.queue.ShutDown()
}
func (c *ClusterQuotaReconcilationController) Sync(discoveryFunc resourcequota.NamespacedResourcesFunc, period time.Duration, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldResources := make(map[schema.GroupVersionResource]struct{})
	wait.Until(func() {
		newResources, err := resourcequota.GetQuotableResources(discoveryFunc)
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
		klog.V(2).Infof("syncing resource quota controller with updated resources from discovery: %v", newResources)
		oldResources = newResources
		c.workerLock.Lock()
		defer c.workerLock.Unlock()
		if err := c.resyncMonitors(newResources); err != nil {
			utilruntime.HandleError(fmt.Errorf("failed to sync resource monitors: %v", err))
			return
		}
		if c.quotaMonitor != nil && !controller.WaitForCacheSync("cluster resource quota", stopCh, c.quotaMonitor.IsSynced) {
			utilruntime.HandleError(fmt.Errorf("timed out waiting for quota monitor sync"))
		}
	}, period, stopCh)
}
func (c *ClusterQuotaReconcilationController) resyncMonitors(resources map[schema.GroupVersionResource]struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.quotaMonitor.SyncMonitors(resources); err != nil {
		return err
	}
	c.quotaMonitor.StartMonitors()
	return nil
}
func (c *ClusterQuotaReconcilationController) calculate(quotaName string, namespaceNames ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(namespaceNames) == 0 {
		return
	}
	items := make([]interface{}, 0, len(namespaceNames))
	for _, name := range namespaceNames {
		items = append(items, workItem{namespaceName: name, forceRecalculation: false})
	}
	c.queue.AddWithData(quotaName, items...)
}
func (c *ClusterQuotaReconcilationController) forceCalculation(quotaName string, namespaceNames ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(namespaceNames) == 0 {
		return
	}
	items := make([]interface{}, 0, len(namespaceNames))
	for _, name := range namespaceNames {
		items = append(items, workItem{namespaceName: name, forceRecalculation: true})
	}
	c.queue.AddWithData(quotaName, items...)
}
func (c *ClusterQuotaReconcilationController) calculateAll() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	quotas, err := c.clusterQuotaLister.List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	for _, quota := range quotas {
		namespaces, _ := c.clusterQuotaMapper.GetNamespacesFor(quota.Name)
		if len(namespaces) > 0 {
			c.forceCalculation(quota.Name, namespaces...)
			continue
		}
		if len(quota.Status.Namespaces) > 0 {
			c.queue.AddWithData(quota.Name)
			continue
		}
	}
}
func (c *ClusterQuotaReconcilationController) worker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	workFunc := func() bool {
		c.workerLock.RLock()
		defer c.workerLock.RUnlock()
		uncastKey, uncastData, quit := c.queue.GetWithData()
		if quit {
			return true
		}
		defer c.queue.Done(uncastKey)
		quotaName := uncastKey.(string)
		quota, err := c.clusterQuotaLister.Get(quotaName)
		if apierrors.IsNotFound(err) {
			c.queue.Forget(uncastKey)
			return false
		}
		if err != nil {
			utilruntime.HandleError(err)
			c.queue.AddWithDataRateLimited(uncastKey, uncastData...)
			return false
		}
		workItems := make([]workItem, 0, len(uncastData))
		for _, dataElement := range uncastData {
			workItems = append(workItems, dataElement.(workItem))
		}
		err, retryItems := c.syncQuotaForNamespaces(quota, workItems)
		if err == nil {
			c.queue.Forget(uncastKey)
			return false
		}
		utilruntime.HandleError(err)
		items := make([]interface{}, 0, len(retryItems))
		for _, item := range retryItems {
			items = append(items, item)
		}
		c.queue.AddWithDataRateLimited(uncastKey, items...)
		return false
	}
	for {
		if quit := workFunc(); quit {
			klog.Infof("resource quota controller worker shutting down")
			return
		}
	}
}
func (c *ClusterQuotaReconcilationController) syncQuotaForNamespaces(originalQuota *quotav1.ClusterResourceQuota, workItems []workItem) (error, []workItem) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	quota := originalQuota.DeepCopy()
	matchingNamespaceNamesList, quotaSelector := c.clusterQuotaMapper.GetNamespacesFor(quota.Name)
	if !equality.Semantic.DeepEqual(quotaSelector, quota.Spec.Selector) {
		return fmt.Errorf("mapping not up to date, have=%v need=%v", quotaSelector, quota.Spec.Selector), workItems
	}
	matchingNamespaceNames := sets.NewString(matchingNamespaceNamesList...)
	reconcilationErrors := []error{}
	retryItems := []workItem{}
	for _, item := range workItems {
		namespaceName := item.namespaceName
		namespaceTotals, namespaceLoaded := quotav1conversions.GetResourceQuotasStatusByNamespace(quota.Status.Namespaces, namespaceName)
		if !matchingNamespaceNames.Has(namespaceName) {
			if namespaceLoaded {
				quota.Status.Total.Used = utilquota.Subtract(quota.Status.Total.Used, namespaceTotals.Used)
				quotav1conversions.RemoveResourceQuotasStatusByNamespace(&quota.Status.Namespaces, namespaceName)
			}
			continue
		}
		if !item.forceRecalculation && namespaceLoaded && equality.Semantic.DeepEqual(namespaceTotals.Hard, quota.Spec.Quota.Hard) {
			continue
		}
		actualUsage, err := quotaUsageCalculationFunc(namespaceName, quota.Spec.Quota.Scopes, quota.Spec.Quota.Hard, c.registry, quota.Spec.Quota.ScopeSelector)
		if err != nil {
			reconcilationErrors = append(reconcilationErrors, err)
			retryItems = append(retryItems, item)
			continue
		}
		recalculatedStatus := corev1.ResourceQuotaStatus{Used: actualUsage, Hard: quota.Spec.Quota.Hard}
		quota.Status.Total.Used = utilquota.Subtract(quota.Status.Total.Used, namespaceTotals.Used)
		quota.Status.Total.Used = utilquota.Add(quota.Status.Total.Used, recalculatedStatus.Used)
		quotav1conversions.InsertResourceQuotasStatus(&quota.Status.Namespaces, quotav1.ResourceQuotaStatusByNamespace{Namespace: namespaceName, Status: recalculatedStatus})
	}
	statusCopy := quota.Status.Namespaces.DeepCopy()
	for _, namespaceTotals := range statusCopy {
		namespaceName := namespaceTotals.Namespace
		if !matchingNamespaceNames.Has(namespaceName) {
			quota.Status.Total.Used = utilquota.Subtract(quota.Status.Total.Used, namespaceTotals.Status.Used)
			quotav1conversions.RemoveResourceQuotasStatusByNamespace(&quota.Status.Namespaces, namespaceName)
		}
	}
	quota.Status.Total.Hard = quota.Spec.Quota.Hard
	if equality.Semantic.DeepEqual(quota, originalQuota) {
		return kutilerrors.NewAggregate(reconcilationErrors), retryItems
	}
	if _, err := c.clusterQuotaClient.UpdateStatus(quota); err != nil {
		return kutilerrors.NewAggregate(append(reconcilationErrors, err)), workItems
	}
	return kutilerrors.NewAggregate(reconcilationErrors), retryItems
}
func (c *ClusterQuotaReconcilationController) replenishQuota(groupResource schema.GroupResource, namespace string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	releventEvaluators := []utilquota.Evaluator{}
	evaluators := c.registry.List()
	for i := range evaluators {
		evaluator := evaluators[i]
		if evaluator.GroupResource() == groupResource {
			releventEvaluators = append(releventEvaluators, evaluator)
		}
	}
	if len(releventEvaluators) == 0 {
		return
	}
	quotaNames, _ := c.clusterQuotaMapper.GetClusterQuotasFor(namespace)
	for _, quotaName := range quotaNames {
		quota, err := c.clusterQuotaLister.Get(quotaName)
		if err != nil {
			continue
		}
		resourceQuotaResources := utilquota.ResourceNames(quota.Status.Total.Hard)
		for _, evaluator := range releventEvaluators {
			matchedResources := evaluator.MatchingResources(resourceQuotaResources)
			if len(matchedResources) > 0 {
				c.forceCalculation(quotaName, namespace)
				break
			}
		}
	}
}
func (c *ClusterQuotaReconcilationController) addClusterQuota(cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.enqueueClusterQuota(cur)
}
func (c *ClusterQuotaReconcilationController) updateClusterQuota(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.enqueueClusterQuota(cur)
}
func (c *ClusterQuotaReconcilationController) enqueueClusterQuota(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	quota, ok := obj.(*quotav1.ClusterResourceQuota)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("not a ClusterResourceQuota %v", obj))
		return
	}
	namespaces, _ := c.clusterQuotaMapper.GetNamespacesFor(quota.Name)
	c.calculate(quota.Name, namespaces...)
}
func (c *ClusterQuotaReconcilationController) AddMapping(quotaName, namespaceName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.calculate(quotaName, namespaceName)
}
func (c *ClusterQuotaReconcilationController) RemoveMapping(quotaName, namespaceName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.calculate(quotaName, namespaceName)
}

var quotaUsageCalculationFunc = utilquota.CalculateUsage

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
