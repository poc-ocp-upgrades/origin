package fake

import (
	internalversion "github.com/openshift/origin/pkg/route/generated/internalclientset/typed/route/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRoute struct{ *testing.Fake }

func (c *FakeRoute) Routes(namespace string) internalversion.RouteResourceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeRoutes{c, namespace}
}
func (c *FakeRoute) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
