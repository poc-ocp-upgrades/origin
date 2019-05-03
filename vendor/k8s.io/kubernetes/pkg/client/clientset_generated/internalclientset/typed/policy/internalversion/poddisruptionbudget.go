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

type PodDisruptionBudgetsGetter interface {
 PodDisruptionBudgets(namespace string) PodDisruptionBudgetInterface
}
type PodDisruptionBudgetInterface interface {
 Create(*policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error)
 Update(*policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error)
 UpdateStatus(*policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*policy.PodDisruptionBudget, error)
 List(opts v1.ListOptions) (*policy.PodDisruptionBudgetList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodDisruptionBudget, err error)
 PodDisruptionBudgetExpansion
}
type podDisruptionBudgets struct {
 client rest.Interface
 ns     string
}

func newPodDisruptionBudgets(c *PolicyClient, namespace string) *podDisruptionBudgets {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &podDisruptionBudgets{client: c.RESTClient(), ns: namespace}
}
func (c *podDisruptionBudgets) Get(name string, options v1.GetOptions) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodDisruptionBudget{}
 err = c.client.Get().Namespace(c.ns).Resource("poddisruptionbudgets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *podDisruptionBudgets) List(opts v1.ListOptions) (result *policy.PodDisruptionBudgetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &policy.PodDisruptionBudgetList{}
 err = c.client.Get().Namespace(c.ns).Resource("poddisruptionbudgets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *podDisruptionBudgets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("poddisruptionbudgets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *podDisruptionBudgets) Create(podDisruptionBudget *policy.PodDisruptionBudget) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodDisruptionBudget{}
 err = c.client.Post().Namespace(c.ns).Resource("poddisruptionbudgets").Body(podDisruptionBudget).Do().Into(result)
 return
}
func (c *podDisruptionBudgets) Update(podDisruptionBudget *policy.PodDisruptionBudget) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodDisruptionBudget{}
 err = c.client.Put().Namespace(c.ns).Resource("poddisruptionbudgets").Name(podDisruptionBudget.Name).Body(podDisruptionBudget).Do().Into(result)
 return
}
func (c *podDisruptionBudgets) UpdateStatus(podDisruptionBudget *policy.PodDisruptionBudget) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodDisruptionBudget{}
 err = c.client.Put().Namespace(c.ns).Resource("poddisruptionbudgets").Name(podDisruptionBudget.Name).SubResource("status").Body(podDisruptionBudget).Do().Into(result)
 return
}
func (c *podDisruptionBudgets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("poddisruptionbudgets").Name(name).Body(options).Do().Error()
}
func (c *podDisruptionBudgets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("poddisruptionbudgets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *podDisruptionBudgets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &policy.PodDisruptionBudget{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("poddisruptionbudgets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
