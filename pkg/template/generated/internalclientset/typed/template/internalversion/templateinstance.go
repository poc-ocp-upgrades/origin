package internalversion

import (
	"time"
	template "github.com/openshift/origin/pkg/template/apis/template"
	scheme "github.com/openshift/origin/pkg/template/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type TemplateInstancesGetter interface {
	TemplateInstances(namespace string) TemplateInstanceInterface
}
type TemplateInstanceInterface interface {
	Create(*template.TemplateInstance) (*template.TemplateInstance, error)
	Update(*template.TemplateInstance) (*template.TemplateInstance, error)
	UpdateStatus(*template.TemplateInstance) (*template.TemplateInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*template.TemplateInstance, error)
	List(opts v1.ListOptions) (*template.TemplateInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *template.TemplateInstance, err error)
	TemplateInstanceExpansion
}
type templateInstances struct {
	client	rest.Interface
	ns	string
}

func newTemplateInstances(c *TemplateClient, namespace string) *templateInstances {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &templateInstances{client: c.RESTClient(), ns: namespace}
}
func (c *templateInstances) Get(name string, options v1.GetOptions) (result *template.TemplateInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &template.TemplateInstance{}
	err = c.client.Get().Namespace(c.ns).Resource("templateinstances").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *templateInstances) List(opts v1.ListOptions) (result *template.TemplateInstanceList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &template.TemplateInstanceList{}
	err = c.client.Get().Namespace(c.ns).Resource("templateinstances").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *templateInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("templateinstances").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *templateInstances) Create(templateInstance *template.TemplateInstance) (result *template.TemplateInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &template.TemplateInstance{}
	err = c.client.Post().Namespace(c.ns).Resource("templateinstances").Body(templateInstance).Do().Into(result)
	return
}
func (c *templateInstances) Update(templateInstance *template.TemplateInstance) (result *template.TemplateInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &template.TemplateInstance{}
	err = c.client.Put().Namespace(c.ns).Resource("templateinstances").Name(templateInstance.Name).Body(templateInstance).Do().Into(result)
	return
}
func (c *templateInstances) UpdateStatus(templateInstance *template.TemplateInstance) (result *template.TemplateInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &template.TemplateInstance{}
	err = c.client.Put().Namespace(c.ns).Resource("templateinstances").Name(templateInstance.Name).SubResource("status").Body(templateInstance).Do().Into(result)
	return
}
func (c *templateInstances) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("templateinstances").Name(name).Body(options).Do().Error()
}
func (c *templateInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("templateinstances").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *templateInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *template.TemplateInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &template.TemplateInstance{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("templateinstances").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
