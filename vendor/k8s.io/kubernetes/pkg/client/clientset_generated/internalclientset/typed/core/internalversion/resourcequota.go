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

type ResourceQuotasGetter interface {
 ResourceQuotas(namespace string) ResourceQuotaInterface
}
type ResourceQuotaInterface interface {
 Create(*core.ResourceQuota) (*core.ResourceQuota, error)
 Update(*core.ResourceQuota) (*core.ResourceQuota, error)
 UpdateStatus(*core.ResourceQuota) (*core.ResourceQuota, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.ResourceQuota, error)
 List(opts v1.ListOptions) (*core.ResourceQuotaList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ResourceQuota, err error)
 ResourceQuotaExpansion
}
type resourceQuotas struct {
 client rest.Interface
 ns     string
}

func newResourceQuotas(c *CoreClient, namespace string) *resourceQuotas {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &resourceQuotas{client: c.RESTClient(), ns: namespace}
}
func (c *resourceQuotas) Get(name string, options v1.GetOptions) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ResourceQuota{}
 err = c.client.Get().Namespace(c.ns).Resource("resourcequotas").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *resourceQuotas) List(opts v1.ListOptions) (result *core.ResourceQuotaList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.ResourceQuotaList{}
 err = c.client.Get().Namespace(c.ns).Resource("resourcequotas").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *resourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("resourcequotas").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *resourceQuotas) Create(resourceQuota *core.ResourceQuota) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ResourceQuota{}
 err = c.client.Post().Namespace(c.ns).Resource("resourcequotas").Body(resourceQuota).Do().Into(result)
 return
}
func (c *resourceQuotas) Update(resourceQuota *core.ResourceQuota) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ResourceQuota{}
 err = c.client.Put().Namespace(c.ns).Resource("resourcequotas").Name(resourceQuota.Name).Body(resourceQuota).Do().Into(result)
 return
}
func (c *resourceQuotas) UpdateStatus(resourceQuota *core.ResourceQuota) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ResourceQuota{}
 err = c.client.Put().Namespace(c.ns).Resource("resourcequotas").Name(resourceQuota.Name).SubResource("status").Body(resourceQuota).Do().Into(result)
 return
}
func (c *resourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("resourcequotas").Name(name).Body(options).Do().Error()
}
func (c *resourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("resourcequotas").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *resourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ResourceQuota{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("resourcequotas").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
