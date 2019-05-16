package namespace

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/namespace/deletion"
	"k8s.io/kubernetes/pkg/util/metrics"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	namespaceDeletionGracePeriod = 5 * time.Second
)

type NamespaceController struct {
	lister                     corelisters.NamespaceLister
	listerSynced               cache.InformerSynced
	queue                      workqueue.RateLimitingInterface
	namespacedResourcesDeleter deletion.NamespacedResourcesDeleterInterface
}

func NewNamespaceController(kubeClient clientset.Interface, dynamicClient dynamic.Interface, discoverResourcesFn func() ([]*metav1.APIResourceList, error), namespaceInformer coreinformers.NamespaceInformer, resyncPeriod time.Duration, finalizerToken v1.FinalizerName) *NamespaceController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespaceController := &NamespaceController{queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "namespace"), namespacedResourcesDeleter: deletion.NewNamespacedResourcesDeleter(kubeClient.CoreV1().Namespaces(), dynamicClient, kubeClient.CoreV1(), discoverResourcesFn, finalizerToken, true)}
	if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("namespace_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
	}
	namespaceInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		namespace := obj.(*v1.Namespace)
		namespaceController.enqueueNamespace(namespace)
	}, UpdateFunc: func(oldObj, newObj interface{}) {
		namespace := newObj.(*v1.Namespace)
		namespaceController.enqueueNamespace(namespace)
	}}, resyncPeriod)
	namespaceController.lister = namespaceInformer.Lister()
	namespaceController.listerSynced = namespaceInformer.Informer().HasSynced
	return namespaceController
}
func (nm *NamespaceController) enqueueNamespace(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}
	namespace := obj.(*v1.Namespace)
	if namespace.DeletionTimestamp == nil || namespace.DeletionTimestamp.IsZero() {
		return
	}
	nm.queue.AddAfter(key, namespaceDeletionGracePeriod)
}
func (nm *NamespaceController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	workFunc := func() bool {
		key, quit := nm.queue.Get()
		if quit {
			return true
		}
		defer nm.queue.Done(key)
		err := nm.syncNamespaceFromKey(key.(string))
		if err == nil {
			nm.queue.Forget(key)
			return false
		}
		if estimate, ok := err.(*deletion.ResourcesRemainingError); ok {
			t := estimate.Estimate/2 + 1
			klog.V(4).Infof("Content remaining in namespace %s, waiting %d seconds", key, t)
			nm.queue.AddAfter(key, time.Duration(t)*time.Second)
		} else {
			nm.queue.AddRateLimited(key)
			utilruntime.HandleError(err)
		}
		return false
	}
	for {
		quit := workFunc()
		if quit {
			return
		}
	}
}
func (nm *NamespaceController) syncNamespaceFromKey(key string) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing namespace %q (%v)", key, time.Since(startTime))
	}()
	namespace, err := nm.lister.Get(key)
	if errors.IsNotFound(err) {
		klog.Infof("Namespace has been deleted %v", key)
		return nil
	}
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to retrieve namespace %v from store: %v", key, err))
		return err
	}
	return nm.namespacedResourcesDeleter.Delete(namespace.Name)
}
func (nm *NamespaceController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer nm.queue.ShutDown()
	klog.Infof("Starting namespace controller")
	defer klog.Infof("Shutting down namespace controller")
	if !controller.WaitForCacheSync("namespace", stopCh, nm.listerSynced) {
		return
	}
	klog.V(5).Info("Starting workers of namespace controller")
	for i := 0; i < workers; i++ {
		go wait.Until(nm.worker, time.Second, stopCh)
	}
	<-stopCh
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
