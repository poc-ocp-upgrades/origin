package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/admissionregistration/internalversion"
)

type FakeAdmissionregistration struct{ *testing.Fake }

func (c *FakeAdmissionregistration) InitializerConfigurations() internalversion.InitializerConfigurationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeInitializerConfigurations{c}
}
func (c *FakeAdmissionregistration) MutatingWebhookConfigurations() internalversion.MutatingWebhookConfigurationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeMutatingWebhookConfigurations{c}
}
func (c *FakeAdmissionregistration) ValidatingWebhookConfigurations() internalversion.ValidatingWebhookConfigurationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeValidatingWebhookConfigurations{c}
}
func (c *FakeAdmissionregistration) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
