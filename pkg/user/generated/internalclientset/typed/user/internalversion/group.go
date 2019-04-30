package internalversion

import (
	"time"
	user "github.com/openshift/origin/pkg/user/apis/user"
	scheme "github.com/openshift/origin/pkg/user/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type GroupsGetter interface{ Groups() GroupInterface }
type GroupInterface interface {
	Create(*user.Group) (*user.Group, error)
	Update(*user.Group) (*user.Group, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*user.Group, error)
	List(opts v1.ListOptions) (*user.GroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Group, err error)
	GroupExpansion
}
type groups struct{ client rest.Interface }

func newGroups(c *UserClient) *groups {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &groups{client: c.RESTClient()}
}
func (c *groups) Get(name string, options v1.GetOptions) (result *user.Group, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Group{}
	err = c.client.Get().Resource("groups").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *groups) List(opts v1.ListOptions) (result *user.GroupList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &user.GroupList{}
	err = c.client.Get().Resource("groups").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *groups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("groups").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *groups) Create(group *user.Group) (result *user.Group, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Group{}
	err = c.client.Post().Resource("groups").Body(group).Do().Into(result)
	return
}
func (c *groups) Update(group *user.Group) (result *user.Group, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Group{}
	err = c.client.Put().Resource("groups").Name(group.Name).Body(group).Do().Into(result)
	return
}
func (c *groups) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("groups").Name(name).Body(options).Do().Error()
}
func (c *groups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("groups").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *groups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Group, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Group{}
	err = c.client.Patch(pt).Resource("groups").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
