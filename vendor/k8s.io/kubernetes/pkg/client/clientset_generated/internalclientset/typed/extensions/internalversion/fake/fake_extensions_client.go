package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/extensions/internalversion"
)

type FakeExtensions struct{ *testing.Fake }

func (c *FakeExtensions) Ingresses(namespace string) internalversion.IngressInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeIngresses{c, namespace}
}
func (c *FakeExtensions) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
