package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/networking/internalversion"
)

type FakeNetworking struct{ *testing.Fake }

func (c *FakeNetworking) NetworkPolicies(namespace string) internalversion.NetworkPolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeNetworkPolicies{c, namespace}
}
func (c *FakeNetworking) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
