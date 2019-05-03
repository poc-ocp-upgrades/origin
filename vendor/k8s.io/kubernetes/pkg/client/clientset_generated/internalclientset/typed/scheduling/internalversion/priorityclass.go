package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type PriorityClassesGetter interface{ PriorityClasses() PriorityClassInterface }
type PriorityClassInterface interface {
 Create(*scheduling.PriorityClass) (*scheduling.PriorityClass, error)
 Update(*scheduling.PriorityClass) (*scheduling.PriorityClass, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*scheduling.PriorityClass, error)
 List(opts v1.ListOptions) (*scheduling.PriorityClassList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *scheduling.PriorityClass, err error)
 PriorityClassExpansion
}
type priorityClasses struct{ client rest.Interface }

func newPriorityClasses(c *SchedulingClient) *priorityClasses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &priorityClasses{client: c.RESTClient()}
}
func (c *priorityClasses) Get(name string, options v1.GetOptions) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &scheduling.PriorityClass{}
 err = c.client.Get().Resource("priorityclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *priorityClasses) List(opts v1.ListOptions) (result *scheduling.PriorityClassList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &scheduling.PriorityClassList{}
 err = c.client.Get().Resource("priorityclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *priorityClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("priorityclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *priorityClasses) Create(priorityClass *scheduling.PriorityClass) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &scheduling.PriorityClass{}
 err = c.client.Post().Resource("priorityclasses").Body(priorityClass).Do().Into(result)
 return
}
func (c *priorityClasses) Update(priorityClass *scheduling.PriorityClass) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &scheduling.PriorityClass{}
 err = c.client.Put().Resource("priorityclasses").Name(priorityClass.Name).Body(priorityClass).Do().Into(result)
 return
}
func (c *priorityClasses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("priorityclasses").Name(name).Body(options).Do().Error()
}
func (c *priorityClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("priorityclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *priorityClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &scheduling.PriorityClass{}
 err = c.client.Patch(pt).Resource("priorityclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
