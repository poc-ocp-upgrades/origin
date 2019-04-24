package internalversion

import (
	"time"
	security "github.com/openshift/origin/pkg/security/apis/security"
	scheme "github.com/openshift/origin/pkg/security/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type SecurityContextConstraintsGetter interface {
	SecurityContextConstraints() SecurityContextConstraintsInterface
}
type SecurityContextConstraintsInterface interface {
	Create(*security.SecurityContextConstraints) (*security.SecurityContextConstraints, error)
	Update(*security.SecurityContextConstraints) (*security.SecurityContextConstraints, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*security.SecurityContextConstraints, error)
	List(opts v1.ListOptions) (*security.SecurityContextConstraintsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *security.SecurityContextConstraints, err error)
	SecurityContextConstraintsExpansion
}
type securityContextConstraints struct{ client rest.Interface }

func newSecurityContextConstraints(c *SecurityClient) *securityContextConstraints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &securityContextConstraints{client: c.RESTClient()}
}
func (c *securityContextConstraints) Get(name string, options v1.GetOptions) (result *security.SecurityContextConstraints, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &security.SecurityContextConstraints{}
	err = c.client.Get().Resource("securitycontextconstraints").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *securityContextConstraints) List(opts v1.ListOptions) (result *security.SecurityContextConstraintsList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &security.SecurityContextConstraintsList{}
	err = c.client.Get().Resource("securitycontextconstraints").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *securityContextConstraints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("securitycontextconstraints").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *securityContextConstraints) Create(securityContextConstraints *security.SecurityContextConstraints) (result *security.SecurityContextConstraints, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &security.SecurityContextConstraints{}
	err = c.client.Post().Resource("securitycontextconstraints").Body(securityContextConstraints).Do().Into(result)
	return
}
func (c *securityContextConstraints) Update(securityContextConstraints *security.SecurityContextConstraints) (result *security.SecurityContextConstraints, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &security.SecurityContextConstraints{}
	err = c.client.Put().Resource("securitycontextconstraints").Name(securityContextConstraints.Name).Body(securityContextConstraints).Do().Into(result)
	return
}
func (c *securityContextConstraints) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("securitycontextconstraints").Name(name).Body(options).Do().Error()
}
func (c *securityContextConstraints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("securitycontextconstraints").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *securityContextConstraints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *security.SecurityContextConstraints, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &security.SecurityContextConstraints{}
	err = c.client.Patch(pt).Resource("securitycontextconstraints").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
