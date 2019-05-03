package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 policy "k8s.io/kubernetes/pkg/apis/policy"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type PodSecurityPoliciesGetter interface {
 PodSecurityPolicies() PodSecurityPolicyInterface
}
type PodSecurityPolicyInterface interface {
 Create(*policy.PodSecurityPolicy) (*policy.PodSecurityPolicy, error)
 Update(*policy.PodSecurityPolicy) (*policy.PodSecurityPolicy, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*policy.PodSecurityPolicy, error)
 List(opts v1.ListOptions) (*policy.PodSecurityPolicyList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodSecurityPolicy, err error)
 PodSecurityPolicyExpansion
}
type podSecurityPolicies struct{ client rest.Interface }

func newPodSecurityPolicies(c *PolicyClient) *podSecurityPolicies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &podSecurityPolicies{client: c.RESTClient()}
}
func (c *podSecurityPolicies) Get(name string, options v1.GetOptions) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodSecurityPolicy{}
 err = c.client.Get().Resource("podsecuritypolicies").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *podSecurityPolicies) List(opts v1.ListOptions) (result *policy.PodSecurityPolicyList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &policy.PodSecurityPolicyList{}
 err = c.client.Get().Resource("podsecuritypolicies").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *podSecurityPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("podsecuritypolicies").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *podSecurityPolicies) Create(podSecurityPolicy *policy.PodSecurityPolicy) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodSecurityPolicy{}
 err = c.client.Post().Resource("podsecuritypolicies").Body(podSecurityPolicy).Do().Into(result)
 return
}
func (c *podSecurityPolicies) Update(podSecurityPolicy *policy.PodSecurityPolicy) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodSecurityPolicy{}
 err = c.client.Put().Resource("podsecuritypolicies").Name(podSecurityPolicy.Name).Body(podSecurityPolicy).Do().Into(result)
 return
}
func (c *podSecurityPolicies) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("podsecuritypolicies").Name(name).Body(options).Do().Error()
}
func (c *podSecurityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("podsecuritypolicies").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *podSecurityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodSecurityPolicy{}
 err = c.client.Patch(pt).Resource("podsecuritypolicies").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
