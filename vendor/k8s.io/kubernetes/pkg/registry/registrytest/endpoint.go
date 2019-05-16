package registrytest

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	api "k8s.io/kubernetes/pkg/apis/core"
	"sync"
)

type EndpointRegistry struct {
	Endpoints *api.EndpointsList
	Updates   []api.Endpoints
	Err       error
	lock      sync.Mutex
}

func (e *EndpointRegistry) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.Endpoints, e.Err
}
func (e *EndpointRegistry) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &api.Endpoints{}
}
func (e *EndpointRegistry) NewList() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &api.EndpointsList{}
}
func (e *EndpointRegistry) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.Err != nil {
		return nil, e.Err
	}
	if e.Endpoints != nil {
		for _, endpoint := range e.Endpoints.Items {
			if endpoint.Name == name {
				return &endpoint, nil
			}
		}
	}
	return nil, errors.NewNotFound(api.Resource("endpoints"), name)
}
func (e *EndpointRegistry) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("unimplemented!")
}
func (e *EndpointRegistry) Create(ctx context.Context, endpoints runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("unimplemented!")
}
func (e *EndpointRegistry) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreateOnUpdate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := objInfo.UpdatedObject(ctx, nil)
	if err != nil {
		return nil, false, err
	}
	endpoints := obj.(*api.Endpoints)
	e.lock.Lock()
	defer e.lock.Unlock()
	e.Updates = append(e.Updates, *endpoints)
	if e.Err != nil {
		return nil, false, e.Err
	}
	if e.Endpoints == nil {
		e.Endpoints = &api.EndpointsList{Items: []api.Endpoints{*endpoints}}
		return endpoints, false, nil
	}
	for ix := range e.Endpoints.Items {
		if e.Endpoints.Items[ix].Name == endpoints.Name {
			e.Endpoints.Items[ix] = *endpoints
		}
	}
	e.Endpoints.Items = append(e.Endpoints.Items, *endpoints)
	return endpoints, false, nil
}
func (e *EndpointRegistry) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.Err != nil {
		return nil, false, e.Err
	}
	if e.Endpoints != nil {
		var newList []api.Endpoints
		for _, endpoint := range e.Endpoints.Items {
			if endpoint.Name != name {
				newList = append(newList, endpoint)
			}
		}
		e.Endpoints.Items = newList
	}
	return nil, true, nil
}
func (e *EndpointRegistry) DeleteCollection(ctx context.Context, _ *metav1.DeleteOptions, _ *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("unimplemented!")
}
