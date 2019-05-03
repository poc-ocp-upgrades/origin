package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type StorageInterface interface {
 RESTClient() rest.Interface
 StorageClassesGetter
 VolumeAttachmentsGetter
}
type StorageClient struct{ restClient rest.Interface }

func (c *StorageClient) StorageClasses() StorageClassInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newStorageClasses(c)
}
func (c *StorageClient) VolumeAttachments() VolumeAttachmentInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newVolumeAttachments(c)
}
func NewForConfig(c *rest.Config) (*StorageClient, error) {
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
 return &StorageClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *StorageClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *StorageClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &StorageClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("storage.k8s.io")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("storage.k8s.io")[0]
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
func (c *StorageClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
