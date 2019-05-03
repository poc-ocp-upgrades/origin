package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type SchedulingInterface interface {
 RESTClient() rest.Interface
 PriorityClassesGetter
}
type SchedulingClient struct{ restClient rest.Interface }

func (c *SchedulingClient) PriorityClasses() PriorityClassInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPriorityClasses(c)
}
func NewForConfig(c *rest.Config) (*SchedulingClient, error) {
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
 return &SchedulingClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *SchedulingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *SchedulingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &SchedulingClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("scheduling.k8s.io")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("scheduling.k8s.io")[0]
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
func (c *SchedulingClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
