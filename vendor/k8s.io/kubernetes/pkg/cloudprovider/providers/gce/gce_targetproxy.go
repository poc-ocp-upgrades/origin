package gce

import (
 compute "google.golang.org/api/compute/v1"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newTargetProxyMetricContext(request string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("targetproxy", request, unusedMetricLabel, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetTargetHTTPProxy(name string) (*compute.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("get")
 v, err := g.c.TargetHttpProxies().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) CreateTargetHTTPProxy(proxy *compute.TargetHttpProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("create")
 return mc.Observe(g.c.TargetHttpProxies().Insert(ctx, meta.GlobalKey(proxy.Name), proxy))
}
func (g *Cloud) SetURLMapForTargetHTTPProxy(proxy *compute.TargetHttpProxy, urlMapLink string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 ref := &compute.UrlMapReference{UrlMap: urlMapLink}
 mc := newTargetProxyMetricContext("set_url_map")
 return mc.Observe(g.c.TargetHttpProxies().SetUrlMap(ctx, meta.GlobalKey(proxy.Name), ref))
}
func (g *Cloud) DeleteTargetHTTPProxy(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("delete")
 return mc.Observe(g.c.TargetHttpProxies().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) ListTargetHTTPProxies() ([]*compute.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("list")
 v, err := g.c.TargetHttpProxies().List(ctx, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) GetTargetHTTPSProxy(name string) (*compute.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("get")
 v, err := g.c.TargetHttpsProxies().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) CreateTargetHTTPSProxy(proxy *compute.TargetHttpsProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("create")
 return mc.Observe(g.c.TargetHttpsProxies().Insert(ctx, meta.GlobalKey(proxy.Name), proxy))
}
func (g *Cloud) SetURLMapForTargetHTTPSProxy(proxy *compute.TargetHttpsProxy, urlMapLink string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("set_url_map")
 ref := &compute.UrlMapReference{UrlMap: urlMapLink}
 return mc.Observe(g.c.TargetHttpsProxies().SetUrlMap(ctx, meta.GlobalKey(proxy.Name), ref))
}
func (g *Cloud) SetSslCertificateForTargetHTTPSProxy(proxy *compute.TargetHttpsProxy, sslCertURLs []string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("set_ssl_cert")
 req := &compute.TargetHttpsProxiesSetSslCertificatesRequest{SslCertificates: sslCertURLs}
 return mc.Observe(g.c.TargetHttpsProxies().SetSslCertificates(ctx, meta.GlobalKey(proxy.Name), req))
}
func (g *Cloud) DeleteTargetHTTPSProxy(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("delete")
 return mc.Observe(g.c.TargetHttpsProxies().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) ListTargetHTTPSProxies() ([]*compute.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newTargetProxyMetricContext("list")
 v, err := g.c.TargetHttpsProxies().List(ctx, filter.None)
 return v, mc.Observe(err)
}
