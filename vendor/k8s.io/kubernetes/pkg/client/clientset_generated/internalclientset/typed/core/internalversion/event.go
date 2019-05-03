package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 core "k8s.io/kubernetes/pkg/apis/core"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type EventsGetter interface {
 Events(namespace string) EventInterface
}
type EventInterface interface {
 Create(*core.Event) (*core.Event, error)
 Update(*core.Event) (*core.Event, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Event, error)
 List(opts v1.ListOptions) (*core.EventList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Event, err error)
 EventExpansion
}
type events struct {
 client rest.Interface
 ns     string
}

func newEvents(c *CoreClient, namespace string) *events {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &events{client: c.RESTClient(), ns: namespace}
}
func (c *events) Get(name string, options v1.GetOptions) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Event{}
 err = c.client.Get().Namespace(c.ns).Resource("events").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *events) List(opts v1.ListOptions) (result *core.EventList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.EventList{}
 err = c.client.Get().Namespace(c.ns).Resource("events").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *events) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("events").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *events) Create(event *core.Event) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Event{}
 err = c.client.Post().Namespace(c.ns).Resource("events").Body(event).Do().Into(result)
 return
}
func (c *events) Update(event *core.Event) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Event{}
 err = c.client.Put().Namespace(c.ns).Resource("events").Name(event.Name).Body(event).Do().Into(result)
 return
}
func (c *events) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("events").Name(name).Body(options).Do().Error()
}
func (c *events) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("events").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *events) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Event{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("events").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
