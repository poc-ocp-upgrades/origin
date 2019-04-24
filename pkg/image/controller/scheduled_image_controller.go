package controller

import (
	"fmt"
	"k8s.io/klog"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/flowcontrol"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1lister "github.com/openshift/client-go/image/listers/image/v1"
	metrics "github.com/openshift/origin/pkg/image/metrics/prometheus"
)

type uniqueItem struct {
	uid		string
	resourceVersion	string
}
type ScheduledImageStreamController struct {
	enabled		bool
	client		rest.Interface
	lister		imagev1lister.ImageStreamLister
	listerSynced	cache.InformerSynced
	rateLimiter	flowcontrol.RateLimiter
	scheduler	*scheduler
	importCounter	*ImportMetricCounter
}

func (s *ScheduledImageStreamController) Importing(stream *imagev1.ImageStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !s.enabled {
		return
	}
	klog.V(5).Infof("DEBUG: stream %s was just imported", stream.Name)
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(stream)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to get the key for stream %s: %v", stream.Name, err))
		return
	}
	s.scheduler.Delay(key)
}
func (s *ScheduledImageStreamController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	klog.Infof("Starting scheduled import controller")
	if !cache.WaitForCacheSync(stopCh, s.listerSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	go s.scheduler.RunUntil(stopCh)
	metrics.InitializeImportCollector(true, s.importCounter.Collect)
	<-stopCh
	klog.Infof("Shutting down image stream controller")
}
func (s *ScheduledImageStreamController) addImageStream(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream := obj.(*imagev1.ImageStream)
	s.enqueueImageStream(stream)
}
func (s *ScheduledImageStreamController) updateImageStream(old, cur interface{}) {
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
	s.enqueueImageStream(curStream)
}
func (s *ScheduledImageStreamController) deleteImageStream(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to get namespace key for %#v", obj))
		return
	}
	s.scheduler.Remove(key, nil)
}
func (s *ScheduledImageStreamController) enqueueImageStream(stream *imagev1.ImageStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !s.enabled {
		return
	}
	if needsScheduling(stream) {
		key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(stream)
		if err != nil {
			klog.V(2).Infof("unable to get namespace key function for stream %s/%s: %v", stream.Namespace, stream.Name, err)
			return
		}
		s.scheduler.Add(key, uniqueItem{uid: string(stream.UID), resourceVersion: stream.ResourceVersion})
	}
}
func (s *ScheduledImageStreamController) syncTimed(key, value interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !s.enabled {
		s.scheduler.Remove(key, value)
		return
	}
	if s.rateLimiter != nil && !s.rateLimiter.TryAccept() {
		klog.V(5).Infof("DEBUG: check of %s exceeded rate limit, will retry later", key)
		return
	}
	namespace, name, err := cache.SplitMetaNamespaceKey(key.(string))
	if err != nil {
		klog.V(2).Infof("unable to split namespace key for key %q: %v", key, err)
		return
	}
	if err := s.syncTimedByName(namespace, name); err != nil {
		if err == ErrNotImportable {
			s.scheduler.Remove(key, value)
			return
		}
		utilruntime.HandleError(err)
		return
	}
}
func (s *ScheduledImageStreamController) syncTimedByName(namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sharedStream, err := s.lister.ImageStreams(namespace).Get(name)
	if err != nil {
		if apierrs.IsNotFound(err) {
			return ErrNotImportable
		}
		return err
	}
	if !needsScheduling(sharedStream) {
		return ErrNotImportable
	}
	stream := sharedStream.DeepCopy()
	resetScheduledTags(stream)
	klog.V(3).Infof("Scheduled import of stream %s/%s...", stream.Namespace, stream.Name)
	result, err := handleImageStream(stream, s.client, nil)
	s.importCounter.Increment(result, err)
	return err
}
func resetScheduledTags(stream *imagev1.ImageStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	next := stream.Generation + 1
	for tag, tagRef := range stream.Spec.Tags {
		if tagImportable(tagRef) && tagRef.ImportPolicy.Scheduled {
			tagRef.Generation = &next
			stream.Spec.Tags[tag] = tagRef
		}
	}
}
func needsScheduling(stream *imagev1.ImageStream) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, tagRef := range stream.Spec.Tags {
		if tagImportable(tagRef) && tagRef.ImportPolicy.Scheduled {
			return true
		}
	}
	return false
}
