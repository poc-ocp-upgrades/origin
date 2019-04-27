package trigger

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"k8s.io/klog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	kv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/controller"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1informer "github.com/openshift/client-go/image/informers/externalversions/image/v1"
	imagev1lister "github.com/openshift/client-go/image/listers/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/trigger"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

const (
	maxRetries		= 5
	maxResourceInterval	= 30 * time.Second
)

var ErrUnresolvedTag = fmt.Errorf("one or more triggers on this object cannot be resolved")

type TriggerSource struct {
	Resource	schema.GroupResource
	Informer	cache.SharedInformer
	Store		cache.Store
	TriggerFn	func(prefix string) trigger.Indexer
	Reactor		trigger.ImageReactor
}
type tagRetriever struct {
	lister imagev1lister.ImageStreamLister
}

var _ trigger.TagRetriever = tagRetriever{}

func NewTagRetriever(lister imagev1lister.ImageStreamLister) trigger.TagRetriever {
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
	return tagRetriever{lister}
}
func (r tagRetriever) ImageStreamTag(namespace, name string) (ref string, rv int64, ok bool) {
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
	streamName, tag, ok := imageapi.SplitImageStreamTag(name)
	if !ok {
		return "", 0, false
	}
	is, err := r.lister.ImageStreams(namespace).Get(streamName)
	if err != nil {
		return "", 0, false
	}
	rv, err = strconv.ParseInt(is.ResourceVersion, 10, 64)
	if err != nil {
		return "", 0, false
	}
	ref, ok = imageutil.ResolveLatestTaggedImage(is, tag)
	return ref, rv, ok
}
func defaultResourceFailureDelay(requeue int) (time.Duration, bool) {
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
	if requeue > 5 {
		return maxResourceInterval, true
	}
	t := time.Duration(math.Pow(2.0, float64(requeue)) * float64(time.Second))
	if t > maxResourceInterval {
		t = maxResourceInterval
	}
	return t, true
}

type TriggerController struct {
	eventRecorder		record.EventRecorder
	triggerCache		cache.ThreadSafeStore
	triggerSources		map[string]TriggerSource
	syncImageStreamFn	func(key string) error
	syncResourceFn		func(key string) error
	enqueueImageStreamFn	func(is *imagev1.ImageStream)
	resourceFailureDelayFn	func(requeue int) (time.Duration, bool)
	lister			imagev1lister.ImageStreamLister
	tagRetriever		trigger.TagRetriever
	queue			workqueue.RateLimitingInterface
	imageChangeQueue	workqueue.RateLimitingInterface
	syncs			[]cache.InformerSynced
}

func NewTriggerEventBroadcaster(client kv1core.CoreV1Interface) record.EventBroadcaster {
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
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&kv1core.EventSinkImpl{Interface: client.Events("")})
	return eventBroadcaster
}
func NewTriggerController(eventBroadcaster record.EventBroadcaster, isInformer imagev1informer.ImageStreamInformer, sources ...TriggerSource) *TriggerController {
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
	lister := isInformer.Lister()
	c := &TriggerController{eventRecorder: eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: "image-trigger-controller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "image-trigger"), imageChangeQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "image-trigger-reactions"), lister: lister, tagRetriever: NewTagRetriever(lister), triggerCache: NewTriggerCache(), resourceFailureDelayFn: defaultResourceFailureDelay}
	c.syncImageStreamFn = c.syncImageStream
	c.syncResourceFn = c.syncResource
	c.enqueueImageStreamFn = c.enqueueImageStream
	isInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addImageStreamNotification, UpdateFunc: c.updateImageStreamNotification})
	c.syncs = []cache.InformerSynced{isInformer.Informer().HasSynced}
	triggers, syncs, err := setupTriggerSources(c.triggerCache, c.tagRetriever, sources, c.imageChangeQueue)
	if err != nil {
		panic(err)
	}
	c.triggerSources = triggers
	c.syncs = append(c.syncs, syncs...)
	return c
}
func setupTriggerSources(triggerCache cache.ThreadSafeStore, tagRetriever trigger.TagRetriever, sources []TriggerSource, imageChangeQueue workqueue.RateLimitingInterface) (map[string]TriggerSource, []cache.InformerSynced, error) {
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
	var syncs []cache.InformerSynced
	triggerSources := make(map[string]TriggerSource)
	for _, source := range sources {
		if source.Store == nil {
			source.Store = source.Informer.GetStore()
		}
		prefix := source.Resource.String() + "/"
		if _, ok := triggerSources[source.Resource.String()]; ok {
			return nil, nil, fmt.Errorf("duplicate resource names registered in %#v", sources)
		}
		triggerSources[source.Resource.String()] = source
		handler := ProcessEvents(triggerCache, source.TriggerFn(prefix), imageChangeQueue, tagRetriever)
		source.Informer.AddEventHandler(handler)
		syncs = append(syncs, source.Informer.HasSynced)
	}
	return triggerSources, syncs, nil
}
func (c *TriggerController) Run(workers int, stopCh <-chan struct{}) {
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
	klog.Infof("Starting trigger controller")
	if !cache.WaitForCacheSync(stopCh, c.syncs...) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(c.imageStreamWorker, time.Second, stopCh)
	}
	for i := 0; i < workers; i++ {
		go wait.Until(c.resourceWorker, time.Second, stopCh)
	}
	<-stopCh
	klog.Infof("Shutting down trigger controller")
}
func (c *TriggerController) addImageStreamNotification(obj interface{}) {
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
	is := obj.(*imagev1.ImageStream)
	c.enqueueImageStreamFn(is)
}
func (c *TriggerController) updateImageStreamNotification(old, cur interface{}) {
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
	c.enqueueImageStreamFn(cur.(*imagev1.ImageStream))
}
func (c *TriggerController) enqueueImageStream(is *imagev1.ImageStream) {
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
	key, err := controller.KeyFunc(is)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", is, err))
		return
	}
	c.queue.Add(key)
}
func (c *TriggerController) imageStreamWorker() {
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
	for c.processNextImageStream() {
	}
	klog.V(4).Infof("Image stream worker stopped")
}
func (c *TriggerController) processNextImageStream() bool {
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
		return false
	}
	defer c.queue.Done(key)
	err := c.syncImageStreamFn(key.(string))
	c.handleImageStreamErr(err, key)
	return true
}
func (c *TriggerController) handleImageStreamErr(err error, key interface{}) {
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
	if c.queue.NumRequeues(key) < maxRetries {
		klog.V(4).Infof("Error syncing image stream %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	klog.V(4).Infof("Dropping image stream %q out of the queue: %v", key, err)
	c.queue.Forget(key)
}
func (c *TriggerController) resourceWorker() {
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
	for c.processNextResource() {
	}
	klog.V(4).Infof("Resource worker stopped")
}
func (c *TriggerController) processNextResource() bool {
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
	key, quit := c.imageChangeQueue.Get()
	if quit {
		return false
	}
	defer c.imageChangeQueue.Done(key.(string))
	err := c.syncResourceFn(key.(string))
	c.handleResourceErr(err, key.(string))
	return true
}
func (c *TriggerController) handleResourceErr(err error, key string) {
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
		c.imageChangeQueue.Forget(key)
		return
	}
	if delay, ok := c.resourceFailureDelayFn(c.imageChangeQueue.NumRequeues(key)); ok {
		klog.V(4).Infof("Error syncing resource %s: %v", key, err)
		c.imageChangeQueue.AddAfter(key, delay)
		return
	}
	utilruntime.HandleError(err)
	klog.V(4).Infof("Dropping resource %q out of the queue: %v", key, err)
	c.imageChangeQueue.Forget(key)
}
func (c *TriggerController) syncImageStream(key string) error {
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
	if klog.V(4) {
		startTime := time.Now()
		klog.Infof("Started syncing image stream %q", key)
		defer func() {
			klog.Infof("Finished syncing image stream %q (%v)", key, time.Since(startTime))
		}()
	}
	triggered, err := c.triggerCache.ByIndex("images", key)
	if err != nil {
		return err
	}
	if len(triggered) == 0 {
		return nil
	}
	for _, t := range triggered {
		entry := t.(*trigger.CacheEntry)
		c.imageChangeQueue.Add(entry.Key)
	}
	return nil
}
func (c *TriggerController) syncResource(key string) error {
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
	if klog.V(4) {
		startTime := time.Now()
		klog.Infof("Started syncing resource %q", key)
		defer func() {
			klog.Infof("Finished syncing resource %q (%v)", key, time.Since(startTime))
		}()
	}
	parts := strings.SplitN(key, "/", 2)
	source := c.triggerSources[parts[0]]
	obj, exists, err := source.Store.GetByKey(parts[1])
	if err != nil {
		return fmt.Errorf("unable to retrieve %s %s from store: %v", parts[0], parts[1], err)
	}
	if !exists {
		return nil
	}
	return source.Reactor.ImageChanged(obj.(runtime.Object), c.tagRetriever)
}
