package internalversion

import (
	"time"
	project "github.com/openshift/origin/pkg/project/apis/project"
	scheme "github.com/openshift/origin/pkg/project/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ProjectsGetter interface {
	Projects() ProjectResourceInterface
}
type ProjectResourceInterface interface {
	Create(*project.Project) (*project.Project, error)
	Update(*project.Project) (*project.Project, error)
	UpdateStatus(*project.Project) (*project.Project, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*project.Project, error)
	List(opts v1.ListOptions) (*project.ProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *project.Project, err error)
	ProjectResourceExpansion
}
type projects struct{ client rest.Interface }

func newProjects(c *ProjectClient) *projects {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &projects{client: c.RESTClient()}
}
func (c *projects) Get(name string, options v1.GetOptions) (result *project.Project, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &project.Project{}
	err = c.client.Get().Resource("projects").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *projects) List(opts v1.ListOptions) (result *project.ProjectList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &project.ProjectList{}
	err = c.client.Get().Resource("projects").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *projects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("projects").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *projects) Create(projectResource *project.Project) (result *project.Project, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &project.Project{}
	err = c.client.Post().Resource("projects").Body(projectResource).Do().Into(result)
	return
}
func (c *projects) Update(projectResource *project.Project) (result *project.Project, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &project.Project{}
	err = c.client.Put().Resource("projects").Name(projectResource.Name).Body(projectResource).Do().Into(result)
	return
}
func (c *projects) UpdateStatus(projectResource *project.Project) (result *project.Project, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &project.Project{}
	err = c.client.Put().Resource("projects").Name(projectResource.Name).SubResource("status").Body(projectResource).Do().Into(result)
	return
}
func (c *projects) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("projects").Name(name).Body(options).Do().Error()
}
func (c *projects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("projects").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *projects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *project.Project, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &project.Project{}
	err = c.client.Patch(pt).Resource("projects").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
