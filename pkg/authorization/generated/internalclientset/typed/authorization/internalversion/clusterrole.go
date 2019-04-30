package internalversion

import (
	"time"
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	scheme "github.com/openshift/origin/pkg/authorization/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ClusterRolesGetter interface{ ClusterRoles() ClusterRoleInterface }
type ClusterRoleInterface interface {
	Create(*authorization.ClusterRole) (*authorization.ClusterRole, error)
	Update(*authorization.ClusterRole) (*authorization.ClusterRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*authorization.ClusterRole, error)
	List(opts v1.ListOptions) (*authorization.ClusterRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.ClusterRole, err error)
	ClusterRoleExpansion
}
type clusterRoles struct{ client rest.Interface }

func newClusterRoles(c *AuthorizationClient) *clusterRoles {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterRoles{client: c.RESTClient()}
}
func (c *clusterRoles) Get(name string, options v1.GetOptions) (result *authorization.ClusterRole, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ClusterRole{}
	err = c.client.Get().Resource("clusterroles").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *clusterRoles) List(opts v1.ListOptions) (result *authorization.ClusterRoleList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authorization.ClusterRoleList{}
	err = c.client.Get().Resource("clusterroles").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *clusterRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("clusterroles").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterRoles) Create(clusterRole *authorization.ClusterRole) (result *authorization.ClusterRole, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ClusterRole{}
	err = c.client.Post().Resource("clusterroles").Body(clusterRole).Do().Into(result)
	return
}
func (c *clusterRoles) Update(clusterRole *authorization.ClusterRole) (result *authorization.ClusterRole, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ClusterRole{}
	err = c.client.Put().Resource("clusterroles").Name(clusterRole.Name).Body(clusterRole).Do().Into(result)
	return
}
func (c *clusterRoles) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("clusterroles").Name(name).Body(options).Do().Error()
}
func (c *clusterRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("clusterroles").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.ClusterRole, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ClusterRole{}
	err = c.client.Patch(pt).Resource("clusterroles").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
