package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type PolicyInterface interface {
 RESTClient() rest.Interface
 EvictionsGetter
 PodDisruptionBudgetsGetter
 PodSecurityPoliciesGetter
}
type PolicyClient struct{ restClient rest.Interface }

func (c *PolicyClient) Evictions(namespace string) EvictionInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newEvictions(c, namespace)
}
func (c *PolicyClient) PodDisruptionBudgets(namespace string) PodDisruptionBudgetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPodDisruptionBudgets(c, namespace)
}
func (c *PolicyClient) PodSecurityPolicies() PodSecurityPolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPodSecurityPolicies(c)
}
func NewForConfig(c *rest.Config) (*PolicyClient, error) {
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
 return &PolicyClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *PolicyClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *PolicyClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &PolicyClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("policy")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("policy")[0]
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
func (c *PolicyClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
