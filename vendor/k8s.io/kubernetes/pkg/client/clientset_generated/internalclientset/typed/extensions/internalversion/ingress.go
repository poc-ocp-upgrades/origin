package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 extensions "k8s.io/kubernetes/pkg/apis/extensions"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type IngressesGetter interface {
 Ingresses(namespace string) IngressInterface
}
type IngressInterface interface {
 Create(*extensions.Ingress) (*extensions.Ingress, error)
 Update(*extensions.Ingress) (*extensions.Ingress, error)
 UpdateStatus(*extensions.Ingress) (*extensions.Ingress, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*extensions.Ingress, error)
 List(opts v1.ListOptions) (*extensions.IngressList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.Ingress, err error)
 IngressExpansion
}
type ingresses struct {
 client rest.Interface
 ns     string
}

func newIngresses(c *ExtensionsClient, namespace string) *ingresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &ingresses{client: c.RESTClient(), ns: namespace}
}
func (c *ingresses) Get(name string, options v1.GetOptions) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &extensions.Ingress{}
 err = c.client.Get().Namespace(c.ns).Resource("ingresses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *ingresses) List(opts v1.ListOptions) (result *extensions.IngressList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &extensions.IngressList{}
 err = c.client.Get().Namespace(c.ns).Resource("ingresses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *ingresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("ingresses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *ingresses) Create(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &extensions.Ingress{}
 err = c.client.Post().Namespace(c.ns).Resource("ingresses").Body(ingress).Do().Into(result)
 return
}
func (c *ingresses) Update(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &extensions.Ingress{}
 err = c.client.Put().Namespace(c.ns).Resource("ingresses").Name(ingress.Name).Body(ingress).Do().Into(result)
 return
}
func (c *ingresses) UpdateStatus(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &extensions.Ingress{}
 err = c.client.Put().Namespace(c.ns).Resource("ingresses").Name(ingress.Name).SubResource("status").Body(ingress).Do().Into(result)
 return
}
func (c *ingresses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("ingresses").Name(name).Body(options).Do().Error()
}
func (c *ingresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("ingresses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *ingresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &extensions.Ingress{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("ingresses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
