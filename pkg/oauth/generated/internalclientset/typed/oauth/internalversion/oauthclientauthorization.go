package internalversion

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	scheme "github.com/openshift/origin/pkg/oauth/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

type OAuthClientAuthorizationsGetter interface {
	OAuthClientAuthorizations() OAuthClientAuthorizationInterface
}
type OAuthClientAuthorizationInterface interface {
	Create(*oauth.OAuthClientAuthorization) (*oauth.OAuthClientAuthorization, error)
	Update(*oauth.OAuthClientAuthorization) (*oauth.OAuthClientAuthorization, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*oauth.OAuthClientAuthorization, error)
	List(opts v1.ListOptions) (*oauth.OAuthClientAuthorizationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthClientAuthorization, err error)
	OAuthClientAuthorizationExpansion
}
type oAuthClientAuthorizations struct{ client rest.Interface }

func newOAuthClientAuthorizations(c *OauthClient) *oAuthClientAuthorizations {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &oAuthClientAuthorizations{client: c.RESTClient()}
}
func (c *oAuthClientAuthorizations) Get(name string, options v1.GetOptions) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthClientAuthorization{}
	err = c.client.Get().Resource("oauthclientauthorizations").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *oAuthClientAuthorizations) List(opts v1.ListOptions) (result *oauth.OAuthClientAuthorizationList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &oauth.OAuthClientAuthorizationList{}
	err = c.client.Get().Resource("oauthclientauthorizations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *oAuthClientAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("oauthclientauthorizations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *oAuthClientAuthorizations) Create(oAuthClientAuthorization *oauth.OAuthClientAuthorization) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthClientAuthorization{}
	err = c.client.Post().Resource("oauthclientauthorizations").Body(oAuthClientAuthorization).Do().Into(result)
	return
}
func (c *oAuthClientAuthorizations) Update(oAuthClientAuthorization *oauth.OAuthClientAuthorization) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthClientAuthorization{}
	err = c.client.Put().Resource("oauthclientauthorizations").Name(oAuthClientAuthorization.Name).Body(oAuthClientAuthorization).Do().Into(result)
	return
}
func (c *oAuthClientAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("oauthclientauthorizations").Name(name).Body(options).Do().Error()
}
func (c *oAuthClientAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("oauthclientauthorizations").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *oAuthClientAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthClientAuthorization{}
	err = c.client.Patch(pt).Resource("oauthclientauthorizations").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
