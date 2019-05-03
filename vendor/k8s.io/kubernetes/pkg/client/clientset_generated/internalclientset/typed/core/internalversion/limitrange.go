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

type LimitRangesGetter interface {
 LimitRanges(namespace string) LimitRangeInterface
}
type LimitRangeInterface interface {
 Create(*core.LimitRange) (*core.LimitRange, error)
 Update(*core.LimitRange) (*core.LimitRange, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.LimitRange, error)
 List(opts v1.ListOptions) (*core.LimitRangeList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.LimitRange, err error)
 LimitRangeExpansion
}
type limitRanges struct {
 client rest.Interface
 ns     string
}

func newLimitRanges(c *CoreClient, namespace string) *limitRanges {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &limitRanges{client: c.RESTClient(), ns: namespace}
}
func (c *limitRanges) Get(name string, options v1.GetOptions) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.LimitRange{}
 err = c.client.Get().Namespace(c.ns).Resource("limitranges").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *limitRanges) List(opts v1.ListOptions) (result *core.LimitRangeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.LimitRangeList{}
 err = c.client.Get().Namespace(c.ns).Resource("limitranges").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *limitRanges) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("limitranges").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *limitRanges) Create(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.LimitRange{}
 err = c.client.Post().Namespace(c.ns).Resource("limitranges").Body(limitRange).Do().Into(result)
 return
}
func (c *limitRanges) Update(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.LimitRange{}
 err = c.client.Put().Namespace(c.ns).Resource("limitranges").Name(limitRange.Name).Body(limitRange).Do().Into(result)
 return
}
func (c *limitRanges) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("limitranges").Name(name).Body(options).Do().Error()
}
func (c *limitRanges) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("limitranges").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *limitRanges) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.LimitRange{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("limitranges").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
