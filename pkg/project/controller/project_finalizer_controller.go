package controller

import (
	"fmt"
	goformat "fmt"
	projectapiv1 "github.com/openshift/api/project/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ProjectFinalizerController struct {
	client      kubernetes.Interface
	queue       workqueue.RateLimitingInterface
	cacheSynced cache.InformerSynced
	nsLister    corev1listers.NamespaceLister
	syncHandler func(key string) error
}

func NewProjectFinalizerController(namespaces corev1informers.NamespaceInformer, client kubernetes.Interface) *ProjectFinalizerController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c := &ProjectFinalizerController{client: client, cacheSynced: namespaces.Informer().HasSynced, queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), nsLister: namespaces.Lister()}
	namespaces.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		c.enqueueNamespace(obj)
	}, UpdateFunc: func(oldObj, newObj interface{}) {
		c.enqueueNamespace(newObj)
	}})
	c.syncHandler = c.syncNamespace
	return c
}
func (c *ProjectFinalizerController) Run(stopCh <-chan struct{}, workers int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, c.cacheSynced) {
		return
	}
	klog.V(5).Infof("Starting workers")
	for i := 0; i < workers; i++ {
		go c.worker()
	}
	<-stopCh
	klog.V(1).Infof("Shutting down")
}
func (c *ProjectFinalizerController) enqueueNamespace(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.queue.Add(key)
}
func (c *ProjectFinalizerController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for {
		if !c.work() {
			return
		}
	}
}
func (c *ProjectFinalizerController) work() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	if err := c.syncHandler(key.(string)); err == nil {
		c.queue.Forget(key)
	} else {
		runtime.HandleError(fmt.Errorf("error syncing namespace, it will be retried: %v", err))
		c.queue.AddRateLimited(key)
	}
	return true
}
func (c *ProjectFinalizerController) syncNamespace(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns, err := c.nsLister.Get(key)
	if errors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	found := false
	for _, finalizerName := range ns.Spec.Finalizers {
		if projectapiv1.FinalizerOrigin == finalizerName {
			found = true
		}
	}
	if !found {
		return nil
	}
	return c.finalize(ns.DeepCopy())
}
func (c *ProjectFinalizerController) finalize(namespace *v1.Namespace) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finalizerSet := sets.NewString()
	for i := range namespace.Spec.Finalizers {
		finalizerSet.Insert(string(namespace.Spec.Finalizers[i]))
	}
	finalizerSet.Delete(string(projectapiv1.FinalizerOrigin))
	namespace.Spec.Finalizers = make([]v1.FinalizerName, 0, len(finalizerSet))
	for _, value := range finalizerSet.List() {
		namespace.Spec.Finalizers = append(namespace.Spec.Finalizers, v1.FinalizerName(value))
	}
	_, err := c.client.CoreV1().Namespaces().Finalize(namespace)
	return err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
