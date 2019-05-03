package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 networking "k8s.io/kubernetes/pkg/apis/networking"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type NetworkPoliciesGetter interface {
 NetworkPolicies(namespace string) NetworkPolicyInterface
}
type NetworkPolicyInterface interface {
 Create(*networking.NetworkPolicy) (*networking.NetworkPolicy, error)
 Update(*networking.NetworkPolicy) (*networking.NetworkPolicy, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*networking.NetworkPolicy, error)
 List(opts v1.ListOptions) (*networking.NetworkPolicyList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *networking.NetworkPolicy, err error)
 NetworkPolicyExpansion
}
type networkPolicies struct {
 client rest.Interface
 ns     string
}

func newNetworkPolicies(c *NetworkingClient, namespace string) *networkPolicies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &networkPolicies{client: c.RESTClient(), ns: namespace}
}
func (c *networkPolicies) Get(name string, options v1.GetOptions) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &networking.NetworkPolicy{}
 err = c.client.Get().Namespace(c.ns).Resource("networkpolicies").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *networkPolicies) List(opts v1.ListOptions) (result *networking.NetworkPolicyList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &networking.NetworkPolicyList{}
 err = c.client.Get().Namespace(c.ns).Resource("networkpolicies").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *networkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("networkpolicies").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *networkPolicies) Create(networkPolicy *networking.NetworkPolicy) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &networking.NetworkPolicy{}
 err = c.client.Post().Namespace(c.ns).Resource("networkpolicies").Body(networkPolicy).Do().Into(result)
 return
}
func (c *networkPolicies) Update(networkPolicy *networking.NetworkPolicy) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &networking.NetworkPolicy{}
 err = c.client.Put().Namespace(c.ns).Resource("networkpolicies").Name(networkPolicy.Name).Body(networkPolicy).Do().Into(result)
 return
}
func (c *networkPolicies) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("networkpolicies").Name(name).Body(options).Do().Error()
}
func (c *networkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("networkpolicies").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *networkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &networking.NetworkPolicy{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("networkpolicies").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
