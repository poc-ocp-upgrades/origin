package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/coordination/internalversion"
)

type FakeCoordination struct{ *testing.Fake }

func (c *FakeCoordination) Leases(namespace string) internalversion.LeaseInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeLeases{c, namespace}
}
func (c *FakeCoordination) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
