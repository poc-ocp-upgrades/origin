package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/scheduling/internalversion"
)

type FakeScheduling struct{ *testing.Fake }

func (c *FakeScheduling) PriorityClasses() internalversion.PriorityClassInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePriorityClasses{c}
}
func (c *FakeScheduling) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
