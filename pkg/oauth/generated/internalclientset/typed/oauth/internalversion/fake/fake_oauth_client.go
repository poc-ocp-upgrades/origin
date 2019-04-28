package fake

import (
	internalversion "github.com/openshift/origin/pkg/oauth/generated/internalclientset/typed/oauth/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOauth struct{ *testing.Fake }

func (c *FakeOauth) OAuthAccessTokens() internalversion.OAuthAccessTokenInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeOAuthAccessTokens{c}
}
func (c *FakeOauth) OAuthAuthorizeTokens() internalversion.OAuthAuthorizeTokenInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeOAuthAuthorizeTokens{c}
}
func (c *FakeOauth) OAuthClients() internalversion.OAuthClientInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeOAuthClients{c}
}
func (c *FakeOauth) OAuthClientAuthorizations() internalversion.OAuthClientAuthorizationInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeOAuthClientAuthorizations{c}
}
func (c *FakeOauth) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
