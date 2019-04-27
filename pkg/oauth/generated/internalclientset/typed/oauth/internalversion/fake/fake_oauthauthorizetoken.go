package fake

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeOAuthAuthorizeTokens struct{ Fake *FakeOauth }

var oauthauthorizetokensResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "", Resource: "oauthauthorizetokens"}
var oauthauthorizetokensKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "", Kind: "OAuthAuthorizeToken"}

func (c *FakeOAuthAuthorizeTokens) Get(name string, options v1.GetOptions) (result *oauth.OAuthAuthorizeToken, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(oauthauthorizetokensResource, name), &oauth.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAuthorizeToken), err
}
func (c *FakeOAuthAuthorizeTokens) List(opts v1.ListOptions) (result *oauth.OAuthAuthorizeTokenList, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(oauthauthorizetokensResource, oauthauthorizetokensKind, opts), &oauth.OAuthAuthorizeTokenList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauth.OAuthAuthorizeTokenList{ListMeta: obj.(*oauth.OAuthAuthorizeTokenList).ListMeta}
	for _, item := range obj.(*oauth.OAuthAuthorizeTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeOAuthAuthorizeTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(oauthauthorizetokensResource, opts))
}
func (c *FakeOAuthAuthorizeTokens) Create(oAuthAuthorizeToken *oauth.OAuthAuthorizeToken) (result *oauth.OAuthAuthorizeToken, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(oauthauthorizetokensResource, oAuthAuthorizeToken), &oauth.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAuthorizeToken), err
}
func (c *FakeOAuthAuthorizeTokens) Update(oAuthAuthorizeToken *oauth.OAuthAuthorizeToken) (result *oauth.OAuthAuthorizeToken, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(oauthauthorizetokensResource, oAuthAuthorizeToken), &oauth.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAuthorizeToken), err
}
func (c *FakeOAuthAuthorizeTokens) Delete(name string, options *v1.DeleteOptions) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(oauthauthorizetokensResource, name), &oauth.OAuthAuthorizeToken{})
	return err
}
func (c *FakeOAuthAuthorizeTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(oauthauthorizetokensResource, listOptions)
	_, err := c.Fake.Invokes(action, &oauth.OAuthAuthorizeTokenList{})
	return err
}
func (c *FakeOAuthAuthorizeTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAuthorizeToken, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(oauthauthorizetokensResource, name, pt, data, subresources...), &oauth.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAuthorizeToken), err
}
