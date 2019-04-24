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

type FakeOAuthAccessTokens struct{ Fake *FakeOauth }

var oauthaccesstokensResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "", Resource: "oauthaccesstokens"}
var oauthaccesstokensKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "", Kind: "OAuthAccessToken"}

func (c *FakeOAuthAccessTokens) Get(name string, options v1.GetOptions) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(oauthaccesstokensResource, name), &oauth.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAccessToken), err
}
func (c *FakeOAuthAccessTokens) List(opts v1.ListOptions) (result *oauth.OAuthAccessTokenList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(oauthaccesstokensResource, oauthaccesstokensKind, opts), &oauth.OAuthAccessTokenList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauth.OAuthAccessTokenList{ListMeta: obj.(*oauth.OAuthAccessTokenList).ListMeta}
	for _, item := range obj.(*oauth.OAuthAccessTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeOAuthAccessTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(oauthaccesstokensResource, opts))
}
func (c *FakeOAuthAccessTokens) Create(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(oauthaccesstokensResource, oAuthAccessToken), &oauth.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAccessToken), err
}
func (c *FakeOAuthAccessTokens) Update(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(oauthaccesstokensResource, oAuthAccessToken), &oauth.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAccessToken), err
}
func (c *FakeOAuthAccessTokens) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(oauthaccesstokensResource, name), &oauth.OAuthAccessToken{})
	return err
}
func (c *FakeOAuthAccessTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(oauthaccesstokensResource, listOptions)
	_, err := c.Fake.Invokes(action, &oauth.OAuthAccessTokenList{})
	return err
}
func (c *FakeOAuthAccessTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAccessToken, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(oauthaccesstokensResource, name, pt, data, subresources...), &oauth.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthAccessToken), err
}
