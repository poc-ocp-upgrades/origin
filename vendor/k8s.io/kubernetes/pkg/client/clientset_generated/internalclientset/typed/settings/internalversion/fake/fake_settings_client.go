package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion"
)

type FakeSettings struct{ *testing.Fake }

func (c *FakeSettings) PodPresets(namespace string) internalversion.PodPresetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePodPresets{c, namespace}
}
func (c *FakeSettings) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
