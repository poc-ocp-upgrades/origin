package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/auditregistration/internalversion"
)

type FakeAuditregistration struct{ *testing.Fake }

func (c *FakeAuditregistration) AuditSinks() internalversion.AuditSinkInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeAuditSinks{c}
}
func (c *FakeAuditregistration) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
