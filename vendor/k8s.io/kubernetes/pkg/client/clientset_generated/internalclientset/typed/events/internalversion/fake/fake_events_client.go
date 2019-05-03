package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
)

type FakeEvents struct{ *testing.Fake }

func (c *FakeEvents) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
