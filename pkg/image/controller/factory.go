package controller

import (
	goformat "fmt"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned"
	imagev1informer "github.com/openshift/client-go/image/informers/externalversions/image/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/client-go/util/workqueue"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type ScheduledImageStreamControllerOptions struct {
	Resync                   time.Duration
	Enabled                  bool
	DefaultBucketSize        int
	MaxImageImportsPerMinute int
}

func (opts ScheduledImageStreamControllerOptions) Buckets() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buckets := opts.DefaultBucketSize
	switch {
	case opts.Resync > time.Hour:
		return buckets * 2
	case opts.Resync < 10*time.Minute:
		return buckets / 2
	}
	return buckets
}
func (opts ScheduledImageStreamControllerOptions) BucketsToQPS() float32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	seconds := float32(opts.Resync / time.Second)
	return 1.0 / seconds * float32(opts.Buckets())
}
func (opts ScheduledImageStreamControllerOptions) GetRateLimiter() flowcontrol.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if opts.MaxImageImportsPerMinute <= 0 {
		return flowcontrol.NewFakeAlwaysRateLimiter()
	}
	importRate := float32(opts.MaxImageImportsPerMinute) / float32(time.Minute/time.Second)
	importBurst := opts.MaxImageImportsPerMinute * 2
	return flowcontrol.NewTokenBucketRateLimiter(importRate, importBurst)
}
func NewImageStreamController(client imagev1client.Interface, informer imagev1informer.ImageStreamInformer) *ImageStreamController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controller := &ImageStreamController{queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ImageStreamController"), client: client.ImageV1(), lister: informer.Lister(), listerSynced: informer.Informer().HasSynced, importCounter: NewImportMetricCounter()}
	controller.syncHandler = controller.syncImageStream
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.addImageStream, UpdateFunc: controller.updateImageStream})
	return controller
}
func NewScheduledImageStreamController(client imagev1client.Interface, informer imagev1informer.ImageStreamInformer, opts ScheduledImageStreamControllerOptions) *ScheduledImageStreamController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bucketLimiter := flowcontrol.NewTokenBucketRateLimiter(opts.BucketsToQPS(), 1)
	controller := &ScheduledImageStreamController{enabled: opts.Enabled, rateLimiter: opts.GetRateLimiter(), client: client.ImageV1().RESTClient(), lister: informer.Lister(), listerSynced: informer.Informer().HasSynced, importCounter: NewImportMetricCounter()}
	controller.scheduler = newScheduler(opts.Buckets(), bucketLimiter, controller.syncTimed)
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.addImageStream, UpdateFunc: controller.updateImageStream, DeleteFunc: controller.deleteImageStream})
	return controller
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
