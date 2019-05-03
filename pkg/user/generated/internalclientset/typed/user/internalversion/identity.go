package internalversion

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	scheme "github.com/openshift/origin/pkg/user/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

type IdentitiesGetter interface{ Identities() IdentityInterface }
type IdentityInterface interface {
	Create(*user.Identity) (*user.Identity, error)
	Update(*user.Identity) (*user.Identity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*user.Identity, error)
	List(opts v1.ListOptions) (*user.IdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Identity, err error)
	IdentityExpansion
}
type identities struct{ client rest.Interface }

func newIdentities(c *UserClient) *identities {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &identities{client: c.RESTClient()}
}
func (c *identities) Get(name string, options v1.GetOptions) (result *user.Identity, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Identity{}
	err = c.client.Get().Resource("identities").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *identities) List(opts v1.ListOptions) (result *user.IdentityList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &user.IdentityList{}
	err = c.client.Get().Resource("identities").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *identities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("identities").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *identities) Create(identity *user.Identity) (result *user.Identity, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Identity{}
	err = c.client.Post().Resource("identities").Body(identity).Do().Into(result)
	return
}
func (c *identities) Update(identity *user.Identity) (result *user.Identity, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Identity{}
	err = c.client.Put().Resource("identities").Name(identity.Name).Body(identity).Do().Into(result)
	return
}
func (c *identities) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("identities").Name(name).Body(options).Do().Error()
}
func (c *identities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("identities").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *identities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Identity, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &user.Identity{}
	err = c.client.Patch(pt).Resource("identities").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
