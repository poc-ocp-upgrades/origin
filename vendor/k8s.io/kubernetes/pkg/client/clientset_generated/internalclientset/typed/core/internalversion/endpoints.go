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

type EndpointsGetter interface {
 Endpoints(namespace string) EndpointsInterface
}
type EndpointsInterface interface {
 Create(*core.Endpoints) (*core.Endpoints, error)
 Update(*core.Endpoints) (*core.Endpoints, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Endpoints, error)
 List(opts v1.ListOptions) (*core.EndpointsList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Endpoints, err error)
 EndpointsExpansion
}
type endpoints struct {
 client rest.Interface
 ns     string
}

func newEndpoints(c *CoreClient, namespace string) *endpoints {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &endpoints{client: c.RESTClient(), ns: namespace}
}
func (c *endpoints) Get(name string, options v1.GetOptions) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Endpoints{}
 err = c.client.Get().Namespace(c.ns).Resource("endpoints").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *endpoints) List(opts v1.ListOptions) (result *core.EndpointsList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.EndpointsList{}
 err = c.client.Get().Namespace(c.ns).Resource("endpoints").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *endpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("endpoints").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *endpoints) Create(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Endpoints{}
 err = c.client.Post().Namespace(c.ns).Resource("endpoints").Body(endpoints).Do().Into(result)
 return
}
func (c *endpoints) Update(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Endpoints{}
 err = c.client.Put().Namespace(c.ns).Resource("endpoints").Name(endpoints.Name).Body(endpoints).Do().Into(result)
 return
}
func (c *endpoints) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("endpoints").Name(name).Body(options).Do().Error()
}
func (c *endpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("endpoints").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *endpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Endpoints{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("endpoints").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
