package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/apps/internalversion"
)

type FakeApps struct{ *testing.Fake }

func (c *FakeApps) ControllerRevisions(namespace string) internalversion.ControllerRevisionInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeControllerRevisions{c, namespace}
}
func (c *FakeApps) DaemonSets(namespace string) internalversion.DaemonSetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeDaemonSets{c, namespace}
}
func (c *FakeApps) Deployments(namespace string) internalversion.DeploymentInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeDeployments{c, namespace}
}
func (c *FakeApps) ReplicaSets(namespace string) internalversion.ReplicaSetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeReplicaSets{c, namespace}
}
func (c *FakeApps) StatefulSets(namespace string) internalversion.StatefulSetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeStatefulSets{c, namespace}
}
func (c *FakeApps) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
