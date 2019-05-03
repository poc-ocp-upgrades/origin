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

type ServicesGetter interface {
 Services(namespace string) ServiceInterface
}
type ServiceInterface interface {
 Create(*core.Service) (*core.Service, error)
 Update(*core.Service) (*core.Service, error)
 UpdateStatus(*core.Service) (*core.Service, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Service, error)
 List(opts v1.ListOptions) (*core.ServiceList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Service, err error)
 ServiceExpansion
}
type services struct {
 client rest.Interface
 ns     string
}

func newServices(c *CoreClient, namespace string) *services {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &services{client: c.RESTClient(), ns: namespace}
}
func (c *services) Get(name string, options v1.GetOptions) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Service{}
 err = c.client.Get().Namespace(c.ns).Resource("services").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *services) List(opts v1.ListOptions) (result *core.ServiceList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.ServiceList{}
 err = c.client.Get().Namespace(c.ns).Resource("services").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *services) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("services").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *services) Create(service *core.Service) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Service{}
 err = c.client.Post().Namespace(c.ns).Resource("services").Body(service).Do().Into(result)
 return
}
func (c *services) Update(service *core.Service) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Service{}
 err = c.client.Put().Namespace(c.ns).Resource("services").Name(service.Name).Body(service).Do().Into(result)
 return
}
func (c *services) UpdateStatus(service *core.Service) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Service{}
 err = c.client.Put().Namespace(c.ns).Resource("services").Name(service.Name).SubResource("status").Body(service).Do().Into(result)
 return
}
func (c *services) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("services").Name(name).Body(options).Do().Error()
}
func (c *services) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("services").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *services) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Service{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("services").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
