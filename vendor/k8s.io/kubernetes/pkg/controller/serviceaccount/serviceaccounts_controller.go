package serviceaccount

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type ServiceAccountsControllerOptions struct {
	ServiceAccounts      []v1.ServiceAccount
	ServiceAccountResync time.Duration
	NamespaceResync      time.Duration
}

func DefaultServiceAccountsControllerOptions() ServiceAccountsControllerOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ServiceAccountsControllerOptions{ServiceAccounts: []v1.ServiceAccount{{ObjectMeta: metav1.ObjectMeta{Name: "default"}}}}
}
func NewServiceAccountsController(saInformer coreinformers.ServiceAccountInformer, nsInformer coreinformers.NamespaceInformer, cl clientset.Interface, options ServiceAccountsControllerOptions) (*ServiceAccountsController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e := &ServiceAccountsController{client: cl, serviceAccountsToEnsure: options.ServiceAccounts, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "serviceaccount")}
	if cl != nil && cl.CoreV1().RESTClient().GetRateLimiter() != nil {
		if err := metrics.RegisterMetricAndTrackRateLimiterUsage("serviceaccount_controller", cl.CoreV1().RESTClient().GetRateLimiter()); err != nil {
			return nil, err
		}
	}
	saInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{DeleteFunc: e.serviceAccountDeleted}, options.ServiceAccountResync)
	e.saLister = saInformer.Lister()
	e.saListerSynced = saInformer.Informer().HasSynced
	nsInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: e.namespaceAdded, UpdateFunc: e.namespaceUpdated}, options.NamespaceResync)
	e.nsLister = nsInformer.Lister()
	e.nsListerSynced = nsInformer.Informer().HasSynced
	e.syncHandler = e.syncNamespace
	return e, nil
}

type ServiceAccountsController struct {
	client                  clientset.Interface
	serviceAccountsToEnsure []v1.ServiceAccount
	syncHandler             func(key string) error
	saLister                corelisters.ServiceAccountLister
	saListerSynced          cache.InformerSynced
	nsLister                corelisters.NamespaceLister
	nsListerSynced          cache.InformerSynced
	queue                   workqueue.RateLimitingInterface
}

func (c *ServiceAccountsController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	klog.Infof("Starting service account controller")
	defer klog.Infof("Shutting down service account controller")
	if !controller.WaitForCacheSync("service account", stopCh, c.saListerSynced, c.nsListerSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}
	<-stopCh
}
func (c *ServiceAccountsController) serviceAccountDeleted(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sa, ok := obj.(*v1.ServiceAccount)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		sa, ok = tombstone.Obj.(*v1.ServiceAccount)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a ServiceAccount %#v", obj))
			return
		}
	}
	c.queue.Add(sa.Namespace)
}
func (c *ServiceAccountsController) namespaceAdded(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace := obj.(*v1.Namespace)
	c.queue.Add(namespace.Name)
}
func (c *ServiceAccountsController) namespaceUpdated(oldObj interface{}, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newNamespace := newObj.(*v1.Namespace)
	c.queue.Add(newNamespace.Name)
}
func (c *ServiceAccountsController) runWorker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for c.processNextWorkItem() {
	}
}
func (c *ServiceAccountsController) processNextWorkItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	err := c.syncHandler(key.(string))
	if err == nil {
		c.queue.Forget(key)
		return true
	}
	utilruntime.HandleError(fmt.Errorf("%v failed with : %v", key, err))
	c.queue.AddRateLimited(key)
	return true
}
func (c *ServiceAccountsController) syncNamespace(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing namespace %q (%v)", key, time.Since(startTime))
	}()
	ns, err := c.nsLister.Get(key)
	if apierrs.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if ns.Status.Phase != v1.NamespaceActive {
		return nil
	}
	createFailures := []error{}
	for i := range c.serviceAccountsToEnsure {
		sa := c.serviceAccountsToEnsure[i]
		switch _, err := c.saLister.ServiceAccounts(ns.Name).Get(sa.Name); {
		case err == nil:
			continue
		case apierrs.IsNotFound(err):
		case err != nil:
			return err
		}
		sa.Namespace = ns.Name
		if _, err := c.client.CoreV1().ServiceAccounts(ns.Name).Create(&sa); err != nil && !apierrs.IsAlreadyExists(err) {
			createFailures = append(createFailures, err)
		}
	}
	return utilerrors.Flatten(utilerrors.NewAggregate(createFailures))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
