package service

import (
 "context"
 "fmt"
 "net/http"
 "net/url"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/proxy"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/capabilities"
)

type ProxyREST struct {
 Redirector     rest.Redirector
 ProxyTransport http.RoundTripper
}

var _ = rest.Connecter(&ProxyREST{})
var proxyMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

func (r *ProxyREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.ServiceProxyOptions{}
}
func (r *ProxyREST) ConnectMethods() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return proxyMethods
}
func (r *ProxyREST) NewConnectOptions() (runtime.Object, bool, string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.ServiceProxyOptions{}, true, "path"
}
func (r *ProxyREST) Connect(ctx context.Context, id string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxyOpts, ok := opts.(*api.ServiceProxyOptions)
 if !ok {
  return nil, fmt.Errorf("Invalid options object: %#v", opts)
 }
 location, transport, err := r.Redirector.ResourceLocation(ctx, id)
 if err != nil {
  return nil, err
 }
 location.Path = net.JoinPreservingTrailingSlash(location.Path, proxyOpts.Path)
 return newThrottledUpgradeAwareProxyHandler(location, transport, true, false, responder), nil
}
func newThrottledUpgradeAwareProxyHandler(location *url.URL, transport http.RoundTripper, wrapTransport, upgradeRequired bool, responder rest.Responder) *proxy.UpgradeAwareHandler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 handler := proxy.NewUpgradeAwareHandler(location, transport, wrapTransport, upgradeRequired, proxy.NewErrorResponder(responder))
 handler.MaxBytesPerSec = capabilities.Get().PerConnectionBandwidthLimitBytesPerSec
 return handler
}
