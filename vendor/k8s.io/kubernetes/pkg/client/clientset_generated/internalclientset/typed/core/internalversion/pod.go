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

type PodsGetter interface {
 Pods(namespace string) PodInterface
}
type PodInterface interface {
 Create(*core.Pod) (*core.Pod, error)
 Update(*core.Pod) (*core.Pod, error)
 UpdateStatus(*core.Pod) (*core.Pod, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Pod, error)
 List(opts v1.ListOptions) (*core.PodList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Pod, err error)
 PodExpansion
}
type pods struct {
 client rest.Interface
 ns     string
}

func newPods(c *CoreClient, namespace string) *pods {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &pods{client: c.RESTClient(), ns: namespace}
}
func (c *pods) Get(name string, options v1.GetOptions) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Pod{}
 err = c.client.Get().Namespace(c.ns).Resource("pods").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *pods) List(opts v1.ListOptions) (result *core.PodList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.PodList{}
 err = c.client.Get().Namespace(c.ns).Resource("pods").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *pods) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("pods").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *pods) Create(pod *core.Pod) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Pod{}
 err = c.client.Post().Namespace(c.ns).Resource("pods").Body(pod).Do().Into(result)
 return
}
func (c *pods) Update(pod *core.Pod) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Pod{}
 err = c.client.Put().Namespace(c.ns).Resource("pods").Name(pod.Name).Body(pod).Do().Into(result)
 return
}
func (c *pods) UpdateStatus(pod *core.Pod) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Pod{}
 err = c.client.Put().Namespace(c.ns).Resource("pods").Name(pod.Name).SubResource("status").Body(pod).Do().Into(result)
 return
}
func (c *pods) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("pods").Name(name).Body(options).Do().Error()
}
func (c *pods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("pods").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *pods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Pod{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("pods").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
