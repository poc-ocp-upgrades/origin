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

type RoleBindingRestrictionsGetter interface {
	RoleBindingRestrictions(namespace string) RoleBindingRestrictionInterface
}
type RoleBindingRestrictionInterface interface {
	Create(*authorization.RoleBindingRestriction) (*authorization.RoleBindingRestriction, error)
	Update(*authorization.RoleBindingRestriction) (*authorization.RoleBindingRestriction, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*authorization.RoleBindingRestriction, error)
	List(opts v1.ListOptions) (*authorization.RoleBindingRestrictionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.RoleBindingRestriction, err error)
	RoleBindingRestrictionExpansion
}
type roleBindingRestrictions struct {
	client	rest.Interface
	ns	string
}

func newRoleBindingRestrictions(c *AuthorizationClient, namespace string) *roleBindingRestrictions {
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
	return &roleBindingRestrictions{client: c.RESTClient(), ns: namespace}
}
func (c *roleBindingRestrictions) Get(name string, options v1.GetOptions) (result *authorization.RoleBindingRestriction, err error) {
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
	result = &authorization.RoleBindingRestriction{}
	err = c.client.Get().Namespace(c.ns).Resource("rolebindingrestrictions").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *roleBindingRestrictions) List(opts v1.ListOptions) (result *authorization.RoleBindingRestrictionList, err error) {
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
	result = &authorization.RoleBindingRestrictionList{}
	err = c.client.Get().Namespace(c.ns).Resource("rolebindingrestrictions").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *roleBindingRestrictions) Watch(opts v1.ListOptions) (watch.Interface, error) {
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
	return c.client.Get().Namespace(c.ns).Resource("rolebindingrestrictions").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *roleBindingRestrictions) Create(roleBindingRestriction *authorization.RoleBindingRestriction) (result *authorization.RoleBindingRestriction, err error) {
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
	result = &authorization.RoleBindingRestriction{}
	err = c.client.Post().Namespace(c.ns).Resource("rolebindingrestrictions").Body(roleBindingRestriction).Do().Into(result)
	return
}
func (c *roleBindingRestrictions) Update(roleBindingRestriction *authorization.RoleBindingRestriction) (result *authorization.RoleBindingRestriction, err error) {
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
	result = &authorization.RoleBindingRestriction{}
	err = c.client.Put().Namespace(c.ns).Resource("rolebindingrestrictions").Name(roleBindingRestriction.Name).Body(roleBindingRestriction).Do().Into(result)
	return
}
func (c *roleBindingRestrictions) Delete(name string, options *v1.DeleteOptions) error {
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
	return c.client.Delete().Namespace(c.ns).Resource("rolebindingrestrictions").Name(name).Body(options).Do().Error()
}
func (c *roleBindingRestrictions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
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
	return c.client.Delete().Namespace(c.ns).Resource("rolebindingrestrictions").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *roleBindingRestrictions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.RoleBindingRestriction, err error) {
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
	result = &authorization.RoleBindingRestriction{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("rolebindingrestrictions").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
