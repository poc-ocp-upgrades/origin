package controller

import (
	"errors"
	"fmt"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	kcontroller "k8s.io/kubernetes/pkg/controller"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	imagev1lister "github.com/openshift/client-go/image/listers/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metrics "github.com/openshift/origin/pkg/image/metrics/prometheus"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

var ErrNotImportable = errors.New("requested image cannot be imported")

type Notifier interface {
	Importing(stream *imagev1.ImageStream)
}
type ImageStreamController struct {
	client		imagev1typedclient.ImageV1Interface
	queue		workqueue.RateLimitingInterface
	syncHandler	func(isKey string) error
	lister		imagev1lister.ImageStreamLister
	listerSynced	cache.InformerSynced
	notifier	Notifier
	importCounter	*ImportMetricCounter
}

func (c *ImageStreamController) SetNotifier(n Notifier) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.notifier = n
}
func (c *ImageStreamController) Run(workers int, stopCh <-chan struct{}) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	klog.Infof("Starting image stream controller")
	if !cache.WaitForCacheSync(stopCh, c.listerSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}
	metrics.InitializeImportCollector(false, c.importCounter.Collect)
	<-stopCh
	klog.Infof("Shutting down image stream controller")
}
func (c *ImageStreamController) addImageStream(obj interface{}) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stream, ok := obj.(*imagev1.ImageStream); ok {
		c.enqueueImageStream(stream)
	}
}
func (c *ImageStreamController) updateImageStream(old, cur interface{}) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	curStream, ok := cur.(*imagev1.ImageStream)
	if !ok {
		return
	}
	oldStream, ok := old.(*imagev1.ImageStream)
	if !ok {
		return
	}
	if curStream.ResourceVersion == oldStream.ResourceVersion {
		return
	}
	c.enqueueImageStream(curStream)
}
func (c *ImageStreamController) enqueueImageStream(stream *imagev1.ImageStream) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := kcontroller.KeyFunc(stream)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for image stream %#v: %v", stream, err))
		return
	}
	c.queue.Add(key)
}
func (c *ImageStreamController) worker() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for c.processNextWorkItem() {
	}
}
func (c *ImageStreamController) processNextWorkItem() bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	utilruntime.HandleError(fmt.Errorf("Error syncing image stream %q: %v", key, err))
	c.queue.AddRateLimited(key)
	return true
}
func (c *ImageStreamController) syncImageStream(key string) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing image stream %q (%v)", key, time.Since(startTime))
	}()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	stream, err := c.lister.ImageStreams(namespace).Get(name)
	if apierrs.IsNotFound(err) {
		klog.V(4).Infof("ImageStream has been deleted: %v", key)
		return nil
	}
	if err != nil {
		return err
	}
	klog.V(3).Infof("Queued import of stream %s/%s...", stream.Namespace, stream.Name)
	result, err := handleImageStream(stream, c.client.RESTClient(), c.notifier)
	c.importCounter.Increment(result, err)
	return err
}
func tagImportable(tagRef imagev1.TagReference) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return !(tagRef.From == nil || tagRef.From.Kind != "DockerImage" || tagRef.Reference)
}
func tagNeedsImport(stream *imagev1.ImageStream, tagRef imagev1.TagReference, importWhenGenerationNil bool) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !tagImportable(tagRef) {
		return false
	}
	if tagRef.Generation == nil {
		return importWhenGenerationNil
	}
	return *tagRef.Generation > imageutil.LatestObservedTagGeneration(stream, tagRef.Name)
}
func needsImport(stream *imagev1.ImageStream) (ok bool, partial bool) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stream.Annotations == nil || len(stream.Annotations[imageapi.DockerImageRepositoryCheckAnnotation]) == 0 {
		if len(stream.Spec.DockerImageRepository) > 0 {
			return true, false
		}
		for _, tagRef := range stream.Spec.Tags {
			if tagImportable(tagRef) {
				return true, true
			}
		}
	}
	for _, tagRef := range stream.Spec.Tags {
		if tagNeedsImport(stream, tagRef, false) {
			return true, true
		}
	}
	return false, false
}
func handleImageStream(stream *imagev1.ImageStream, client rest.Interface, notifier Notifier) (*imagev1.ImageStreamImport, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	ok, partial := needsImport(stream)
	if !ok {
		return nil, nil
	}
	klog.V(3).Infof("Importing stream %s/%s partial=%t...", stream.Namespace, stream.Name, partial)
	if notifier != nil {
		notifier.Importing(stream)
	}
	isi := &imagev1.ImageStreamImport{ObjectMeta: metav1.ObjectMeta{Name: stream.Name, Namespace: stream.Namespace, ResourceVersion: stream.ResourceVersion, UID: stream.UID}, Spec: imagev1.ImageStreamImportSpec{Import: true}}
	for _, tagRef := range stream.Spec.Tags {
		if tagImportable(tagRef) && (tagNeedsImport(stream, tagRef, true) || !partial) {
			isi.Spec.Images = append(isi.Spec.Images, imagev1.ImageImportSpec{From: corev1.ObjectReference{Kind: "DockerImage", Name: tagRef.From.Name}, To: &corev1.LocalObjectReference{Name: tagRef.Name}, ImportPolicy: tagRef.ImportPolicy, ReferencePolicy: tagRef.ReferencePolicy})
		}
	}
	if repo := stream.Spec.DockerImageRepository; !partial && len(repo) > 0 {
		insecure := stream.Annotations[imageapi.InsecureRepositoryAnnotation] == "true"
		isi.Spec.Repository = &imagev1.RepositoryImportSpec{From: corev1.ObjectReference{Kind: "DockerImage", Name: repo}, ImportPolicy: imagev1.TagImportPolicy{Insecure: insecure}}
	}
	if isi.Spec.Repository == nil && len(isi.Spec.Images) == 0 {
		klog.V(4).Infof("Did not find any tags or repository needing import")
		return nil, nil
	}
	result := &imagev1.ImageStreamImport{}
	err := client.Post().Namespace(stream.Namespace).Resource(imagev1.Resource("imagestreamimports").Resource).Body(isi).Timeout(time.Hour).Do().Into(result)
	if err != nil {
		if apierrs.IsNotFound(err) && isStatusErrorKind(err, "imageStream") {
			return result, ErrNotImportable
		}
		klog.V(4).Infof("Import stream %s/%s partial=%t error: %v", stream.Namespace, stream.Name, partial, err)
		return result, err
	}
	klog.V(5).Infof("Import stream %s/%s partial=%t import: %#v", stream.Namespace, stream.Name, partial, result.Status.Import)
	return result, nil
}
func isStatusErrorKind(err error, kind string) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s, ok := err.(apierrs.APIStatus); ok {
		if details := s.Status().Details; details != nil {
			return kind == details.Kind
		}
	}
	return false
}
