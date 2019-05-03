package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 coordination "k8s.io/kubernetes/pkg/apis/coordination"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type LeasesGetter interface {
 Leases(namespace string) LeaseInterface
}
type LeaseInterface interface {
 Create(*coordination.Lease) (*coordination.Lease, error)
 Update(*coordination.Lease) (*coordination.Lease, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*coordination.Lease, error)
 List(opts v1.ListOptions) (*coordination.LeaseList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coordination.Lease, err error)
 LeaseExpansion
}
type leases struct {
 client rest.Interface
 ns     string
}

func newLeases(c *CoordinationClient, namespace string) *leases {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &leases{client: c.RESTClient(), ns: namespace}
}
func (c *leases) Get(name string, options v1.GetOptions) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &coordination.Lease{}
 err = c.client.Get().Namespace(c.ns).Resource("leases").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *leases) List(opts v1.ListOptions) (result *coordination.LeaseList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &coordination.LeaseList{}
 err = c.client.Get().Namespace(c.ns).Resource("leases").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *leases) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("leases").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *leases) Create(lease *coordination.Lease) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &coordination.Lease{}
 err = c.client.Post().Namespace(c.ns).Resource("leases").Body(lease).Do().Into(result)
 return
}
func (c *leases) Update(lease *coordination.Lease) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &coordination.Lease{}
 err = c.client.Put().Namespace(c.ns).Resource("leases").Name(lease.Name).Body(lease).Do().Into(result)
 return
}
func (c *leases) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("leases").Name(name).Body(options).Do().Error()
}
func (c *leases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("leases").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *leases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &coordination.Lease{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("leases").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
