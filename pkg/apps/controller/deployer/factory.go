package deployment

import (
	"fmt"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	kcoreinformers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	kv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kcontroller "k8s.io/kubernetes/pkg/controller"
	"time"
)

func NewDeployerController(rcInformer kcoreinformers.ReplicationControllerInformer, podInformer kcoreinformers.PodInformer, kubeClientset kclientset.Interface, sa, image string, env []v1.EnvVar) *DeploymentController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&kv1core.EventSinkImpl{Interface: kubeClientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: "deployer-controller"})
	c := &DeploymentController{rn: kubeClientset.CoreV1(), pn: kubeClientset.CoreV1(), queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), rcLister: rcInformer.Lister(), rcListerSynced: rcInformer.Informer().HasSynced, podLister: podInformer.Lister(), podListerSynced: podInformer.Informer().HasSynced, serviceAccount: sa, deployerImage: image, environment: env, recorder: recorder}
	rcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addReplicationController, UpdateFunc: c.updateReplicationController})
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: c.updatePod, DeleteFunc: c.deletePod})
	return c
}
func (c *DeploymentController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	klog.Infof("Starting deployer controller")
	if !cache.WaitForCacheSync(stopCh, c.rcListerSynced, c.podListerSynced) {
		return
	}
	klog.Infof("Deployer controller caches are synced. Starting workers.")
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}
	<-stopCh
	klog.Infof("Shutting down deployer controller")
}
func (c *DeploymentController) addReplicationController(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc := obj.(*v1.ReplicationController)
	if !appsutil.IsOwnedByConfig(rc) {
		return
	}
	c.enqueueReplicationController(rc)
}
func (c *DeploymentController) updateReplicationController(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	curRC := cur.(*v1.ReplicationController)
	if !appsutil.IsOwnedByConfig(curRC) {
		return
	}
	c.enqueueReplicationController(curRC)
}
func (c *DeploymentController) updatePod(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	curPod := cur.(*v1.Pod)
	oldPod := old.(*v1.Pod)
	if curPod.ResourceVersion == oldPod.ResourceVersion {
		return
	}
	if rc, err := c.rcForDeployerPod(curPod); err == nil && rc != nil {
		c.enqueueReplicationController(rc)
	}
}
func (c *DeploymentController) deletePod(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, ok := obj.(*v1.Pod)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone: %+v", obj))
			return
		}
		pod, ok = tombstone.Obj.(*v1.Pod)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod: %+v", obj))
			return
		}
	}
	if rc, err := c.rcForDeployerPod(pod); err == nil && rc != nil {
		c.enqueueReplicationController(rc)
	}
}
func (c *DeploymentController) enqueueReplicationController(rc *v1.ReplicationController) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := kcontroller.KeyFunc(rc)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", rc, err))
		return
	}
	c.queue.Add(key)
}
func (c *DeploymentController) rcForDeployerPod(pod *v1.Pod) (*v1.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rcName := appsutil.DeploymentNameFor(pod)
	if len(rcName) == 0 {
		return nil, nil
	}
	key := pod.Namespace + "/" + rcName
	return c.getByKey(key)
}
func (c *DeploymentController) worker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if quit := c.work(); quit {
			return
		}
	}
}
func (c *DeploymentController) work() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := c.queue.Get()
	if quit {
		return true
	}
	defer c.queue.Done(key)
	rc, err := c.getByKey(key.(string))
	if err != nil {
		utilruntime.HandleError(err)
	}
	if rc == nil {
		return false
	}
	willBeDropped := c.queue.NumRequeues(key) >= maxRetryCount-2
	err = c.handle(rc, willBeDropped)
	c.handleErr(err, key, rc)
	return false
}
func (c *DeploymentController) getByKey(key string) (*v1.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return nil, err
	}
	rc, err := c.rcLister.ReplicationControllers(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.V(4).Infof("Replication controller %q has been deleted", key)
		return nil, nil
	}
	if err != nil {
		klog.Infof("Unable to retrieve replication controller %q from store: %v", key, err)
		c.queue.AddRateLimited(key)
		return nil, err
	}
	return rc, nil
}
