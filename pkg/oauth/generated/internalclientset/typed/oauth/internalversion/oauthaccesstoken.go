package internalversion

import (
	"time"
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	scheme "github.com/openshift/origin/pkg/oauth/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type OAuthAccessTokensGetter interface {
	OAuthAccessTokens() OAuthAccessTokenInterface
}
type OAuthAccessTokenInterface interface {
	Create(*oauth.OAuthAccessToken) (*oauth.OAuthAccessToken, error)
	Update(*oauth.OAuthAccessToken) (*oauth.OAuthAccessToken, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*oauth.OAuthAccessToken, error)
	List(opts v1.ListOptions) (*oauth.OAuthAccessTokenList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAccessToken, err error)
	OAuthAccessTokenExpansion
}
type oAuthAccessTokens struct{ client rest.Interface }

func newOAuthAccessTokens(c *OauthClient) *oAuthAccessTokens {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &oAuthAccessTokens{client: c.RESTClient()}
}
func (c *oAuthAccessTokens) Get(name string, options v1.GetOptions) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAccessToken{}
	err = c.client.Get().Resource("oauthaccesstokens").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *oAuthAccessTokens) List(opts v1.ListOptions) (result *oauth.OAuthAccessTokenList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &oauth.OAuthAccessTokenList{}
	err = c.client.Get().Resource("oauthaccesstokens").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *oAuthAccessTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("oauthaccesstokens").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *oAuthAccessTokens) Create(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAccessToken{}
	err = c.client.Post().Resource("oauthaccesstokens").Body(oAuthAccessToken).Do().Into(result)
	return
}
func (c *oAuthAccessTokens) Update(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAccessToken{}
	err = c.client.Put().Resource("oauthaccesstokens").Name(oAuthAccessToken.Name).Body(oAuthAccessToken).Do().Into(result)
	return
}
func (c *oAuthAccessTokens) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("oauthaccesstokens").Name(name).Body(options).Do().Error()
}
func (c *oAuthAccessTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("oauthaccesstokens").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *oAuthAccessTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &oauth.OAuthAccessToken{}
	err = c.client.Patch(pt).Resource("oauthaccesstokens").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
