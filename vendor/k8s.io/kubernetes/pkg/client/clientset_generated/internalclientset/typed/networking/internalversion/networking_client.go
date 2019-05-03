package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type NetworkingInterface interface {
 RESTClient() rest.Interface
 NetworkPoliciesGetter
}
type NetworkingClient struct{ restClient rest.Interface }

func (c *NetworkingClient) NetworkPolicies(namespace string) NetworkPolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newNetworkPolicies(c, namespace)
}
func NewForConfig(c *rest.Config) (*NetworkingClient, error) {
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
 return &NetworkingClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *NetworkingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *NetworkingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &NetworkingClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("networking.k8s.io")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("networking.k8s.io")[0]
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
func (c *NetworkingClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
