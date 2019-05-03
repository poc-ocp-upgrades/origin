package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
)

type FakeCore struct{ *testing.Fake }

func (c *FakeCore) ComponentStatuses() internalversion.ComponentStatusInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeComponentStatuses{c}
}
func (c *FakeCore) ConfigMaps(namespace string) internalversion.ConfigMapInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeConfigMaps{c, namespace}
}
func (c *FakeCore) Endpoints(namespace string) internalversion.EndpointsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeEndpoints{c, namespace}
}
func (c *FakeCore) Events(namespace string) internalversion.EventInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeEvents{c, namespace}
}
func (c *FakeCore) LimitRanges(namespace string) internalversion.LimitRangeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeLimitRanges{c, namespace}
}
func (c *FakeCore) Namespaces() internalversion.NamespaceInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeNamespaces{c}
}
func (c *FakeCore) Nodes() internalversion.NodeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeNodes{c}
}
func (c *FakeCore) PersistentVolumes() internalversion.PersistentVolumeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePersistentVolumes{c}
}
func (c *FakeCore) PersistentVolumeClaims(namespace string) internalversion.PersistentVolumeClaimInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePersistentVolumeClaims{c, namespace}
}
func (c *FakeCore) Pods(namespace string) internalversion.PodInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePods{c, namespace}
}
func (c *FakeCore) PodTemplates(namespace string) internalversion.PodTemplateInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakePodTemplates{c, namespace}
}
func (c *FakeCore) ReplicationControllers(namespace string) internalversion.ReplicationControllerInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeReplicationControllers{c, namespace}
}
func (c *FakeCore) ResourceQuotas(namespace string) internalversion.ResourceQuotaInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeResourceQuotas{c, namespace}
}
func (c *FakeCore) Secrets(namespace string) internalversion.SecretInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeSecrets{c, namespace}
}
func (c *FakeCore) Services(namespace string) internalversion.ServiceInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeServices{c, namespace}
}
func (c *FakeCore) ServiceAccounts(namespace string) internalversion.ServiceAccountInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeServiceAccounts{c, namespace}
}
func (c *FakeCore) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
