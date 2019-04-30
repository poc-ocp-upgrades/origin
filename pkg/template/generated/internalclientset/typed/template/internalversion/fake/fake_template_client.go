package fake

import (
	internalversion "github.com/openshift/origin/pkg/template/generated/internalclientset/typed/template/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTemplate struct{ *testing.Fake }

func (c *FakeTemplate) BrokerTemplateInstances() internalversion.BrokerTemplateInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeBrokerTemplateInstances{c}
}
func (c *FakeTemplate) Templates(namespace string) internalversion.TemplateResourceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeTemplates{c, namespace}
}
func (c *FakeTemplate) TemplateInstances(namespace string) internalversion.TemplateInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeTemplateInstances{c, namespace}
}
func (c *FakeTemplate) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
