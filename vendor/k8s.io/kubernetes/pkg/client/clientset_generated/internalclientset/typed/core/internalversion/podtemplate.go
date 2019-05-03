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

type PodTemplatesGetter interface {
 PodTemplates(namespace string) PodTemplateInterface
}
type PodTemplateInterface interface {
 Create(*core.PodTemplate) (*core.PodTemplate, error)
 Update(*core.PodTemplate) (*core.PodTemplate, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.PodTemplate, error)
 List(opts v1.ListOptions) (*core.PodTemplateList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PodTemplate, err error)
 PodTemplateExpansion
}
type podTemplates struct {
 client rest.Interface
 ns     string
}

func newPodTemplates(c *CoreClient, namespace string) *podTemplates {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &podTemplates{client: c.RESTClient(), ns: namespace}
}
func (c *podTemplates) Get(name string, options v1.GetOptions) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PodTemplate{}
 err = c.client.Get().Namespace(c.ns).Resource("podtemplates").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *podTemplates) List(opts v1.ListOptions) (result *core.PodTemplateList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.PodTemplateList{}
 err = c.client.Get().Namespace(c.ns).Resource("podtemplates").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *podTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("podtemplates").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *podTemplates) Create(podTemplate *core.PodTemplate) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PodTemplate{}
 err = c.client.Post().Namespace(c.ns).Resource("podtemplates").Body(podTemplate).Do().Into(result)
 return
}
func (c *podTemplates) Update(podTemplate *core.PodTemplate) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PodTemplate{}
 err = c.client.Put().Namespace(c.ns).Resource("podtemplates").Name(podTemplate.Name).Body(podTemplate).Do().Into(result)
 return
}
func (c *podTemplates) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("podtemplates").Name(name).Body(options).Do().Error()
}
func (c *podTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("podtemplates").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *podTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PodTemplate{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("podtemplates").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
