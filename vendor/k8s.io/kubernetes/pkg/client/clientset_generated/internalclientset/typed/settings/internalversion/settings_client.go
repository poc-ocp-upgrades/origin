package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type SettingsInterface interface {
 RESTClient() rest.Interface
 PodPresetsGetter
}
type SettingsClient struct{ restClient rest.Interface }

func (c *SettingsClient) PodPresets(namespace string) PodPresetInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newPodPresets(c, namespace)
}
func NewForConfig(c *rest.Config) (*SettingsClient, error) {
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
 return &SettingsClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *SettingsClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *SettingsClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &SettingsClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("settings.k8s.io")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("settings.k8s.io")[0]
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
func (c *SettingsClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
