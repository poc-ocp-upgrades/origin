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

type OAuthAuthorizeTokensGetter interface {
	OAuthAuthorizeTokens() OAuthAuthorizeTokenInterface
}
type OAuthAuthorizeTokenInterface interface {
	Create(*oauth.OAuthAuthorizeToken) (*oauth.OAuthAuthorizeToken, error)
	Update(*oauth.OAuthAuthorizeToken) (*oauth.OAuthAuthorizeToken, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*oauth.OAuthAuthorizeToken, error)
	List(opts v1.ListOptions) (*oauth.OAuthAuthorizeTokenList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAuthorizeToken, err error)
	OAuthAuthorizeTokenExpansion
}
type oAuthAuthorizeTokens struct{ client rest.Interface }

func newOAuthAuthorizeTokens(c *OauthClient) *oAuthAuthorizeTokens {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &oAuthAuthorizeTokens{client: c.RESTClient()}
}
func (c *oAuthAuthorizeTokens) Get(name string, options v1.GetOptions) (result *oauth.OAuthAuthorizeToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAuthorizeToken{}
	err = c.client.Get().Resource("oauthauthorizetokens").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *oAuthAuthorizeTokens) List(opts v1.ListOptions) (result *oauth.OAuthAuthorizeTokenList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &oauth.OAuthAuthorizeTokenList{}
	err = c.client.Get().Resource("oauthauthorizetokens").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *oAuthAuthorizeTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("oauthauthorizetokens").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *oAuthAuthorizeTokens) Create(oAuthAuthorizeToken *oauth.OAuthAuthorizeToken) (result *oauth.OAuthAuthorizeToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAuthorizeToken{}
	err = c.client.Post().Resource("oauthauthorizetokens").Body(oAuthAuthorizeToken).Do().Into(result)
	return
}
func (c *oAuthAuthorizeTokens) Update(oAuthAuthorizeToken *oauth.OAuthAuthorizeToken) (result *oauth.OAuthAuthorizeToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAuthorizeToken{}
	err = c.client.Put().Resource("oauthauthorizetokens").Name(oAuthAuthorizeToken.Name).Body(oAuthAuthorizeToken).Do().Into(result)
	return
}
func (c *oAuthAuthorizeTokens) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("oauthauthorizetokens").Name(name).Body(options).Do().Error()
}
func (c *oAuthAuthorizeTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("oauthauthorizetokens").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *oAuthAuthorizeTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAuthorizeToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAuthorizeToken{}
	err = c.client.Patch(pt).Resource("oauthauthorizetokens").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
