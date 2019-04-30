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

type FakeOAuthClientAuthorizations struct{ Fake *FakeOauth }

var oauthclientauthorizationsResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "", Resource: "oauthclientauthorizations"}
var oauthclientauthorizationsKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "", Kind: "OAuthClientAuthorization"}

func (c *FakeOAuthClientAuthorizations) Get(name string, options v1.GetOptions) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(oauthclientauthorizationsResource, name), &oauth.OAuthClientAuthorization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthClientAuthorization), err
}
func (c *FakeOAuthClientAuthorizations) List(opts v1.ListOptions) (result *oauth.OAuthClientAuthorizationList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(oauthclientauthorizationsResource, oauthclientauthorizationsKind, opts), &oauth.OAuthClientAuthorizationList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauth.OAuthClientAuthorizationList{ListMeta: obj.(*oauth.OAuthClientAuthorizationList).ListMeta}
	for _, item := range obj.(*oauth.OAuthClientAuthorizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeOAuthClientAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(oauthclientauthorizationsResource, opts))
}
func (c *FakeOAuthClientAuthorizations) Create(oAuthClientAuthorization *oauth.OAuthClientAuthorization) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(oauthclientauthorizationsResource, oAuthClientAuthorization), &oauth.OAuthClientAuthorization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthClientAuthorization), err
}
func (c *FakeOAuthClientAuthorizations) Update(oAuthClientAuthorization *oauth.OAuthClientAuthorization) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(oauthclientauthorizationsResource, oAuthClientAuthorization), &oauth.OAuthClientAuthorization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthClientAuthorization), err
}
func (c *FakeOAuthClientAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(oauthclientauthorizationsResource, name), &oauth.OAuthClientAuthorization{})
	return err
}
func (c *FakeOAuthClientAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(oauthclientauthorizationsResource, listOptions)
	_, err := c.Fake.Invokes(action, &oauth.OAuthClientAuthorizationList{})
	return err
}
func (c *FakeOAuthClientAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthClientAuthorization, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(oauthclientauthorizationsResource, name, pt, data, subresources...), &oauth.OAuthClientAuthorization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauth.OAuthClientAuthorization), err
}
