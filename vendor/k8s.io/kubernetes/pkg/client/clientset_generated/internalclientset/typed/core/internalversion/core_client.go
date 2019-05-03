package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type CoreInterface interface {
 RESTClient() rest.Interface
 ComponentStatusesGetter
 ConfigMapsGetter
 EndpointsGetter
 EventsGetter
 LimitRangesGetter
 NamespacesGetter
 NodesGetter
 PersistentVolumesGetter
 PersistentVolumeClaimsGetter
 PodsGetter
 PodTemplatesGetter
 ReplicationControllersGetter
 ResourceQuotasGetter
 SecretsGetter
 ServicesGetter
 ServiceAccountsGetter
}
type CoreClient struct{ restClient rest.Interface }

func (c *CoreClient) ComponentStatuses() ComponentStatusInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newComponentStatuses(c)
}
func (c *CoreClient) ConfigMaps(namespace string) ConfigMapInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newConfigMaps(c, namespace)
}
func (c *CoreClient) Endpoints(namespace string) EndpointsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newEndpoints(c, namespace)
}
func (c *CoreClient) Events(namespace string) EventInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newEvents(c, namespace)
}
func (c *CoreClient) LimitRanges(namespace string) LimitRangeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newLimitRanges(c, namespace)
}
func (c *CoreClient) Namespaces() NamespaceInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newNamespaces(c)
}
func (c *CoreClient) Nodes() NodeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newNodes(c)
}
func (c *CoreClient) PersistentVolumes() PersistentVolumeInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPersistentVolumes(c)
}
func (c *CoreClient) PersistentVolumeClaims(namespace string) PersistentVolumeClaimInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPersistentVolumeClaims(c, namespace)
}
func (c *CoreClient) Pods(namespace string) PodInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPods(c, namespace)
}
func (c *CoreClient) PodTemplates(namespace string) PodTemplateInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPodTemplates(c, namespace)
}
func (c *CoreClient) ReplicationControllers(namespace string) ReplicationControllerInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newReplicationControllers(c, namespace)
}
func (c *CoreClient) ResourceQuotas(namespace string) ResourceQuotaInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newResourceQuotas(c, namespace)
}
func (c *CoreClient) Secrets(namespace string) SecretInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newSecrets(c, namespace)
}
func (c *CoreClient) Services(namespace string) ServiceInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newServices(c, namespace)
}
func (c *CoreClient) ServiceAccounts(namespace string) ServiceAccountInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newServiceAccounts(c, namespace)
}
func NewForConfig(c *rest.Config) (*CoreClient, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config := *c
 if err := setConfigDefaults(&config); err != nil {
  return nil, err
 }
 client, err := rest.RESTClientFor(&config)
 if err != nil {
  return nil, err
 }
 return &CoreClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *CoreClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *CoreClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &CoreClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/api"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("")[0]
  config.GroupVersion = &gv
 }
 config.NegotiatedSerializer = scheme.Codecs
 if config.QPS == 0 {
  config.QPS = 5
 }
 if config.Burst == 0 {
  config.Burst = 10
 }
 return nil
}
func (c *CoreClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
