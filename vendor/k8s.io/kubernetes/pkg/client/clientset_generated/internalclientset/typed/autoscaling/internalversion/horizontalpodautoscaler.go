package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type HorizontalPodAutoscalersGetter interface {
 HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerInterface
}
type HorizontalPodAutoscalerInterface interface {
 Create(*autoscaling.HorizontalPodAutoscaler) (*autoscaling.HorizontalPodAutoscaler, error)
 Update(*autoscaling.HorizontalPodAutoscaler) (*autoscaling.HorizontalPodAutoscaler, error)
 UpdateStatus(*autoscaling.HorizontalPodAutoscaler) (*autoscaling.HorizontalPodAutoscaler, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*autoscaling.HorizontalPodAutoscaler, error)
 List(opts v1.ListOptions) (*autoscaling.HorizontalPodAutoscalerList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *autoscaling.HorizontalPodAutoscaler, err error)
 HorizontalPodAutoscalerExpansion
}
type horizontalPodAutoscalers struct {
 client rest.Interface
 ns     string
}

func newHorizontalPodAutoscalers(c *AutoscalingClient, namespace string) *horizontalPodAutoscalers {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &horizontalPodAutoscalers{client: c.RESTClient(), ns: namespace}
}
func (c *horizontalPodAutoscalers) Get(name string, options v1.GetOptions) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.HorizontalPodAutoscaler{}
 err = c.client.Get().Namespace(c.ns).Resource("horizontalpodautoscalers").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *horizontalPodAutoscalers) List(opts v1.ListOptions) (result *autoscaling.HorizontalPodAutoscalerList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &autoscaling.HorizontalPodAutoscalerList{}
 err = c.client.Get().Namespace(c.ns).Resource("horizontalpodautoscalers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *horizontalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("horizontalpodautoscalers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *horizontalPodAutoscalers) Create(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.HorizontalPodAutoscaler{}
 err = c.client.Post().Namespace(c.ns).Resource("horizontalpodautoscalers").Body(horizontalPodAutoscaler).Do().Into(result)
 return
}
func (c *horizontalPodAutoscalers) Update(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.HorizontalPodAutoscaler{}
 err = c.client.Put().Namespace(c.ns).Resource("horizontalpodautoscalers").Name(horizontalPodAutoscaler.Name).Body(horizontalPodAutoscaler).Do().Into(result)
 return
}
func (c *horizontalPodAutoscalers) UpdateStatus(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.HorizontalPodAutoscaler{}
 err = c.client.Put().Namespace(c.ns).Resource("horizontalpodautoscalers").Name(horizontalPodAutoscaler.Name).SubResource("status").Body(horizontalPodAutoscaler).Do().Into(result)
 return
}
func (c *horizontalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("horizontalpodautoscalers").Name(name).Body(options).Do().Error()
}
func (c *horizontalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("horizontalpodautoscalers").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *horizontalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.HorizontalPodAutoscaler{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("horizontalpodautoscalers").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
