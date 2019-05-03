package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
)

type FakePolicy struct{ *testing.Fake }

func (c *FakePolicy) Evictions(namespace string) internalversion.EvictionInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeEvictions{c, namespace}
}
func (c *FakePolicy) PodDisruptionBudgets(namespace string) internalversion.PodDisruptionBudgetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePodDisruptionBudgets{c, namespace}
}
func (c *FakePolicy) PodSecurityPolicies() internalversion.PodSecurityPolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePodSecurityPolicies{c}
}
func (c *FakePolicy) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
