package genericinformers

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type GenericResourceInformer interface {
	ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error)
	Start(stopCh <-chan struct{})
}
type GenericInternalResourceInformerFunc func(resource schema.GroupVersionResource) (informers.GenericInformer, error)

func (fn GenericInternalResourceInformerFunc) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resource.Version = runtime.APIVersionInternal
	return fn(resource)
}
func (fn GenericInternalResourceInformerFunc) Start(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}

type GenericResourceInformerFunc func(resource schema.GroupVersionResource) (informers.GenericInformer, error)

func (fn GenericResourceInformerFunc) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fn(resource)
}
func (fn GenericResourceInformerFunc) Start(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}

type genericInformers struct {
	startFn func(stopCh <-chan struct{})
	generic []GenericResourceInformer
	bias    map[schema.GroupVersionResource]schema.GroupVersionResource
}

func NewGenericInformers(startFn func(stopCh <-chan struct{}), informers ...GenericResourceInformer) genericInformers {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return genericInformers{startFn: startFn, generic: informers}
}
func (i genericInformers) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if try, ok := i.bias[resource]; ok {
		if res, err := i.ForResource(try); err == nil {
			return res, nil
		}
	}
	var firstErr error
	for _, generic := range i.generic {
		informer, err := generic.ForResource(resource)
		if err == nil {
			return informer, nil
		}
		if firstErr == nil {
			firstErr = err
		}
	}
	klog.V(4).Infof("Couldn't find informer for %v", resource)
	return nil, firstErr
}
func (i genericInformers) Start(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	i.startFn(stopCh)
	for _, generic := range i.generic {
		generic.Start(stopCh)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
