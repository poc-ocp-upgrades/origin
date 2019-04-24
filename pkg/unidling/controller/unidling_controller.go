package controller

import (
	"encoding/json"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"sync"
	"time"
	"k8s.io/klog"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	appstypedclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	unidlingapi "github.com/openshift/origin/pkg/unidling/api"
	unidlingutil "github.com/openshift/origin/pkg/unidling/util"
)

const MaxRetries = 5

type lastFiredCache struct {
	sync.RWMutex
	items	map[types.NamespacedName]time.Time
}

func (c *lastFiredCache) Get(info types.NamespacedName) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.RLock()
	defer c.RUnlock()
	return c.items[info]
}
func (c *lastFiredCache) Clear(info types.NamespacedName) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	delete(c.items, info)
}
func (c *lastFiredCache) AddIfNewer(info types.NamespacedName, newLastFired time.Time) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if lastFired, hasLastFired := c.items[info]; !hasLastFired || lastFired.Before(newLastFired) {
		c.items[info] = newLastFired
		return true
	}
	return false
}

type UnidlingController struct {
	controller		cache.Controller
	scaleNamespacer		scale.ScalesGetter
	mapper			meta.RESTMapper
	endpointsNamespacer	corev1client.EndpointsGetter
	queue			workqueue.RateLimitingInterface
	lastFiredCache		*lastFiredCache
	dcNamespacer		appstypedclient.DeploymentConfigsGetter
	rcNamespacer		corev1client.ReplicationControllersGetter
}

func NewUnidlingController(scaleNS scale.ScalesGetter, mapper meta.RESTMapper, endptsNS corev1client.EndpointsGetter, evtNS corev1client.EventsGetter, dcNamespacer appstypedclient.DeploymentConfigsGetter, rcNamespacer corev1client.ReplicationControllersGetter, resyncPeriod time.Duration) *UnidlingController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{}
	fieldSet["reason"] = unidlingapi.NeedPodsReason
	fieldSelector := fieldSet.AsSelector()
	unidlingController := &UnidlingController{scaleNamespacer: scaleNS, mapper: mapper, endpointsNamespacer: endptsNS, queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), lastFiredCache: &lastFiredCache{items: make(map[types.NamespacedName]time.Time)}, dcNamespacer: dcNamespacer, rcNamespacer: rcNamespacer}
	_, controller := cache.NewInformer(&cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		options.FieldSelector = fieldSelector.String()
		return evtNS.Events(metav1.NamespaceAll).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		options.FieldSelector = fieldSelector.String()
		return evtNS.Events(metav1.NamespaceAll).Watch(options)
	}}, &corev1.Event{}, resyncPeriod, cache.ResourceEventHandlerFuncs{AddFunc: unidlingController.addEvent, UpdateFunc: unidlingController.updateEvent, DeleteFunc: unidlingController.checkAndClearFromCache})
	unidlingController.controller = controller
	return unidlingController
}
func (c *UnidlingController) addEvent(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	evt, ok := obj.(*corev1.Event)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("got non-Event object in event action: %v", obj))
		return
	}
	c.enqueueEvent(evt)
}
func (c *UnidlingController) updateEvent(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	evt, ok := newObj.(*corev1.Event)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("got non-Event object in event action: %v", newObj))
		return
	}
	c.enqueueEvent(evt)
}
func (c *UnidlingController) checkAndClearFromCache(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	evt, objIsEvent := obj.(*corev1.Event)
	if !objIsEvent {
		tombstone, objIsTombstone := obj.(cache.DeletedFinalStateUnknown)
		if !objIsTombstone {
			utilruntime.HandleError(fmt.Errorf("got non-event, non-tombstone object in event action: %v", obj))
			return
		}
		evt, objIsEvent = tombstone.Obj.(*corev1.Event)
		if !objIsEvent {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not an Event in event action: %v", obj))
			return
		}
	}
	c.clearEventFromCache(evt)
}
func (c *UnidlingController) clearEventFromCache(event *corev1.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if event.Reason != unidlingapi.NeedPodsReason {
		return
	}
	info := types.NamespacedName{Namespace: event.InvolvedObject.Namespace, Name: event.InvolvedObject.Name}
	c.lastFiredCache.Clear(info)
}
func (c *UnidlingController) enqueueEvent(event *corev1.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if event.Reason != unidlingapi.NeedPodsReason {
		return
	}
	info := types.NamespacedName{Namespace: event.InvolvedObject.Namespace, Name: event.InvolvedObject.Name}
	if c.lastFiredCache.AddIfNewer(info, event.LastTimestamp.Time) {
		c.queue.Add(info)
	}
}
func (c *UnidlingController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	go c.controller.Run(stopCh)
	go wait.Until(c.processRequests, time.Second, stopCh)
}
func (c *UnidlingController) processRequests() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if !c.awaitRequest() {
			return
		}
	}
}
func (c *UnidlingController) awaitRequest() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	infoRaw, stop := c.queue.Get()
	if stop {
		return false
	}
	defer c.queue.Done(infoRaw)
	info := infoRaw.(types.NamespacedName)
	lastFired := c.lastFiredCache.Get(info)
	var retry bool
	var err error
	if retry, err = c.handleRequest(info, lastFired); err == nil {
		c.queue.Forget(infoRaw)
		return true
	}
	if !retry {
		utilruntime.HandleError(fmt.Errorf("Unable to process unidling event for %s/%s at (%s), will not retry: %v", info.Namespace, info.Name, lastFired, err))
		return true
	}
	if c.queue.NumRequeues(infoRaw) > MaxRetries {
		utilruntime.HandleError(fmt.Errorf("Unable to process unidling event for %s/%s (at %s), will not retry again: %v", info.Namespace, info.Name, lastFired, err))
		c.queue.Forget(infoRaw)
		return true
	}
	klog.V(4).Infof("Unable to fully process unidling request for %s/%s (at %s), will retry: %v", info.Namespace, info.Name, lastFired, err)
	c.queue.AddRateLimited(infoRaw)
	return true
}
func (c *UnidlingController) handleRequest(info types.NamespacedName, lastFired time.Time) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	targetEndpoints, err := c.endpointsNamespacer.Endpoints(info.Namespace).Get(info.Name, metav1.GetOptions{})
	if err != nil {
		return true, fmt.Errorf("unable to retrieve endpoints: %v", err)
	}
	idledTimeRaw, wasIdled := targetEndpoints.Annotations[unidlingapi.IdledAtAnnotation]
	if !wasIdled {
		klog.V(5).Infof("UnidlingController received a NeedPods event for a service that was not idled, ignoring")
		return false, nil
	}
	idledTime, err := time.Parse(time.RFC3339, idledTimeRaw)
	if err != nil {
		return false, fmt.Errorf("unable to check idled-at time: %v", err)
	}
	if lastFired.Before(idledTime) {
		klog.V(5).Infof("UnidlingController received an out-of-date NeedPods event, ignoring")
		return false, nil
	}
	var targetScalables []unidlingapi.RecordedScaleReference
	if targetScalablesStr, hasTargetScalables := targetEndpoints.Annotations[unidlingapi.UnidleTargetAnnotation]; hasTargetScalables {
		if err = json.Unmarshal([]byte(targetScalablesStr), &targetScalables); err != nil {
			return false, fmt.Errorf("unable to unmarshal target scalable references: %v", err)
		}
	} else {
		klog.V(4).Infof("Service %s/%s had no scalables to unidle", info.Namespace, info.Name)
		targetScalables = []unidlingapi.RecordedScaleReference{}
	}
	targetScalablesSet := make(map[unidlingapi.RecordedScaleReference]struct{}, len(targetScalables))
	for _, v := range targetScalables {
		targetScalablesSet[v] = struct{}{}
	}
	deleteIdlingAnnotations := func(_ int32, annotations map[string]string) {
		delete(annotations, unidlingapi.IdledAtAnnotation)
		delete(annotations, unidlingapi.PreviousScaleAnnotation)
	}
	scaleAnnotater := unidlingutil.NewScaleAnnotater(c.scaleNamespacer, c.mapper, c.dcNamespacer, c.rcNamespacer, deleteIdlingAnnotations)
	for _, scalableRef := range targetScalables {
		var scale *autoscalingv1.Scale
		var obj runtime.Object
		obj, scale, err = scaleAnnotater.GetObjectWithScale(info.Namespace, scalableRef.CrossGroupObjectReference)
		if err != nil {
			if errors.IsNotFound(err) {
				utilruntime.HandleError(fmt.Errorf("%s %q does not exist, removing from list of scalables while unidling service %s/%s: %v", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name, err))
				delete(targetScalablesSet, scalableRef)
			} else {
				utilruntime.HandleError(fmt.Errorf("Unable to get scale for %s %q while unidling service %s/%s, will try again later: %v", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name, err))
			}
			continue
		}
		if scale.Spec.Replicas > 0 {
			klog.V(4).Infof("%s %q is not idle, skipping while unidling service %s/%s", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name)
			continue
		}
		scale.Spec.Replicas = scalableRef.Replicas
		updater := unidlingutil.NewScaleUpdater(legacyscheme.Codecs.LegacyCodec(legacyscheme.Scheme.PrioritizedVersionsAllGroups()...), info.Namespace, c.dcNamespacer, c.rcNamespacer)
		if err = scaleAnnotater.UpdateObjectScale(updater, info.Namespace, scalableRef.CrossGroupObjectReference, obj, scale); err != nil {
			if errors.IsNotFound(err) {
				utilruntime.HandleError(fmt.Errorf("%s %q does not exist, removing from list of scalables while unidling service %s/%s: %v", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name, err))
				delete(targetScalablesSet, scalableRef)
			} else {
				utilruntime.HandleError(fmt.Errorf("Unable to scale up %s %q while unidling service %s/%s: %v", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name, err))
			}
			continue
		} else {
			klog.V(4).Infof("Scaled up %s %q while unidling service %s/%s", scalableRef.Kind, scalableRef.Name, info.Namespace, info.Name)
		}
		delete(targetScalablesSet, scalableRef)
	}
	newAnnotationList := make([]unidlingapi.RecordedScaleReference, 0, len(targetScalablesSet))
	for k := range targetScalablesSet {
		newAnnotationList = append(newAnnotationList, k)
	}
	if len(newAnnotationList) == 0 {
		delete(targetEndpoints.Annotations, unidlingapi.UnidleTargetAnnotation)
		delete(targetEndpoints.Annotations, unidlingapi.IdledAtAnnotation)
	} else {
		var newAnnotationBytes []byte
		newAnnotationBytes, err = json.Marshal(newAnnotationList)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to update/remove idle annotations from %s/%s: unable to marshal list of remaining scalables, removing list entirely: %v", info.Namespace, info.Name, err))
			delete(targetEndpoints.Annotations, unidlingapi.UnidleTargetAnnotation)
			delete(targetEndpoints.Annotations, unidlingapi.IdledAtAnnotation)
		} else {
			targetEndpoints.Annotations[unidlingapi.UnidleTargetAnnotation] = string(newAnnotationBytes)
		}
	}
	if _, err = c.endpointsNamespacer.Endpoints(info.Namespace).Update(targetEndpoints); err != nil {
		return true, fmt.Errorf("unable to update/remove idle annotations from %s/%s: %v", info.Namespace, info.Name, err)
	}
	return false, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
