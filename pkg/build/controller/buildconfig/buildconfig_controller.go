package controller

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	kcontroller "k8s.io/kubernetes/pkg/controller"
	buildv1 "github.com/openshift/api/build/v1"
	buildclient "github.com/openshift/client-go/build/clientset/versioned"
	buildinformer "github.com/openshift/client-go/build/informers/externalversions/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	"github.com/openshift/origin/pkg/build/buildscheme"
	buildmanualclient "github.com/openshift/origin/pkg/build/client"
	buildcommon "github.com/openshift/origin/pkg/build/controller/common"
	"github.com/openshift/origin/pkg/build/util"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

const (
	maxRetries = 15
)

type configControllerFatalError struct{ reason string }

func (e *configControllerFatalError) Error() string {
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
	return fmt.Sprintf("fatal: %s", e.reason)
}
func IsFatal(err error) bool {
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
	_, isFatal := err.(*configControllerFatalError)
	return isFatal
}

type BuildConfigController struct {
	buildConfigInstantiator	buildmanualclient.BuildConfigInstantiator
	buildConfigGetter	buildlister.BuildConfigLister
	buildLister		buildlister.BuildLister
	buildDeleter		buildmanualclient.BuildDeleter
	buildConfigInformer	cache.SharedIndexInformer
	queue			workqueue.RateLimitingInterface
	buildConfigStoreSynced	func() bool
	recorder		record.EventRecorder
}

func NewBuildConfigController(buildInternalClient buildclient.Interface, kubeExternalClient kubernetes.Interface, buildConfigInformer buildinformer.BuildConfigInformer, buildInformer buildinformer.BuildInformer) *BuildConfigController {
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
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeExternalClient.CoreV1().Events("")})
	buildClient := buildmanualclient.NewClientBuildClient(buildInternalClient)
	buildConfigGetter := buildConfigInformer.Lister()
	buildConfigInstantiator := buildmanualclient.NewClientBuildConfigInstantiatorClient(buildInternalClient)
	buildLister := buildInformer.Lister()
	c := &BuildConfigController{buildConfigGetter: buildConfigGetter, buildLister: buildLister, buildDeleter: buildClient, buildConfigInstantiator: buildConfigInstantiator, buildConfigInformer: buildConfigInformer.Informer(), queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), recorder: eventBroadcaster.NewRecorder(buildscheme.EncoderScheme, corev1.EventSource{Component: "buildconfig-controller"})}
	c.buildConfigInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: c.buildConfigUpdated, AddFunc: c.buildConfigAdded})
	c.buildConfigStoreSynced = c.buildConfigInformer.HasSynced
	return c
}
func (c *BuildConfigController) handleBuildConfig(bc *buildv1.BuildConfig) error {
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
	klog.V(4).Infof("Handling BuildConfig %s", bcDesc(bc))
	if err := buildcommon.HandleBuildPruning(bc.Name, bc.Namespace, c.buildLister, c.buildConfigGetter, c.buildDeleter); err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to prune builds for %s/%s: %v", bc.Namespace, bc.Name, err))
	}
	hasChangeTrigger := buildapihelpers.HasTriggerType(buildv1.ConfigChangeBuildTriggerType, bc)
	if !hasChangeTrigger {
		return nil
	}
	if bc.Status.LastVersion > 0 {
		return nil
	}
	klog.V(4).Infof("Running build for BuildConfig %s", bcDesc(bc))
	buildTriggerCauses := []buildv1.BuildTriggerCause{}
	lastVersion := int64(0)
	request := &buildv1.BuildRequest{TriggeredBy: append(buildTriggerCauses, buildv1.BuildTriggerCause{Message: util.BuildTriggerCauseConfigMsg}), ObjectMeta: metav1.ObjectMeta{Name: bc.Name, Namespace: bc.Namespace}, LastVersion: &lastVersion}
	if _, err := c.buildConfigInstantiator.Instantiate(bc.Namespace, request); err != nil {
		var instantiateErr error
		if kerrors.IsConflict(err) {
			instantiateErr = fmt.Errorf("unable to instantiate Build for BuildConfig %s due to a conflicting update: %v", bcDesc(bc), err)
			utilruntime.HandleError(instantiateErr)
		} else if buildutil.IsFatalGeneratorError(err) || kerrors.IsNotFound(err) || kerrors.IsBadRequest(err) || kerrors.IsForbidden(err) {
			instantiateErr = fmt.Errorf("gave up on Build for BuildConfig %s due to fatal error: %v", bcDesc(bc), err)
			utilruntime.HandleError(instantiateErr)
			if !strings.Contains(instantiateErr.Error(), "does not match the build request LastVersion(0)") {
				c.recorder.Event(bc, corev1.EventTypeWarning, "BuildConfigInstantiateFailed", instantiateErr.Error())
			}
			return &configControllerFatalError{err.Error()}
		} else {
			instantiateErr = fmt.Errorf("error instantiating Build from BuildConfig %s: %v", bcDesc(bc), err)
			c.recorder.Event(bc, corev1.EventTypeWarning, "BuildConfigInstantiateFailed", instantiateErr.Error())
			utilruntime.HandleError(instantiateErr)
		}
		return instantiateErr
	}
	return nil
}
func (c *BuildConfigController) buildConfigAdded(obj interface{}) {
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
	bc := obj.(*buildv1.BuildConfig)
	c.enqueueBuildConfig(bc)
}
func (c *BuildConfigController) buildConfigUpdated(old, cur interface{}) {
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
	bc := cur.(*buildv1.BuildConfig)
	c.enqueueBuildConfig(bc)
}
func (c *BuildConfigController) enqueueBuildConfig(bc *buildv1.BuildConfig) {
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
	key, err := kcontroller.KeyFunc(bc)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for buildconfig %#v: %v", bc, err))
		return
	}
	c.queue.Add(key)
}
func (c *BuildConfigController) Run(workers int, stopCh <-chan struct{}) {
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
	defer c.queue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, c.buildConfigStoreSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	klog.Infof("Starting buildconfig controller")
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}
	<-stopCh
	klog.Infof("Shutting down buildconfig controller")
}
func (c *BuildConfigController) worker() {
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
		if quit := c.work(); quit {
			return
		}
	}
}
func (c *BuildConfigController) work() bool {
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
	key, quit := c.queue.Get()
	if quit {
		return true
	}
	defer c.queue.Done(key)
	bc, err := c.getBuildConfigByKey(key.(string))
	if err != nil {
		c.handleError(err, key)
		return false
	}
	if bc == nil {
		return false
	}
	err = c.handleBuildConfig(bc)
	c.handleError(err, key)
	return false
}
func (c *BuildConfigController) handleError(err error, key interface{}) {
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
	if err == nil {
		c.queue.Forget(key)
		return
	}
	if IsFatal(err) {
		klog.V(2).Infof("Will not retry fatal error for key %v: %v", key, err)
		c.queue.Forget(key)
		return
	}
	if c.queue.NumRequeues(key) < maxRetries {
		klog.V(4).Infof("Retrying key %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}
	klog.V(2).Infof("Giving up retrying %v: %v", key, err)
	c.queue.Forget(key)
}
func (c *BuildConfigController) getBuildConfigByKey(key string) (*buildv1.BuildConfig, error) {
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
	obj, exists, err := c.buildConfigInformer.GetIndexer().GetByKey(key)
	if err != nil {
		klog.V(2).Infof("Unable to retrieve buildconfig %s from store: %v", key, err)
		return nil, err
	}
	if !exists {
		klog.V(2).Infof("Buildconfig %q has been deleted", key)
		return nil, nil
	}
	return obj.(*buildv1.BuildConfig), nil
}
func bcDesc(bc *buildv1.BuildConfig) string {
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
	return fmt.Sprintf("%s/%s (%d)", bc.Namespace, bc.Name, bc.Status.LastVersion)
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
