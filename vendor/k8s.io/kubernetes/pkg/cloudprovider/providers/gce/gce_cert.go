package gce

import (
 compute "google.golang.org/api/compute/v1"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newCertMetricContext(request string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("cert", request, unusedMetricLabel, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetSslCertificate(name string) (*compute.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newCertMetricContext("get")
 v, err := g.c.SslCertificates().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) CreateSslCertificate(sslCerts *compute.SslCertificate) (*compute.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newCertMetricContext("create")
 err := g.c.SslCertificates().Insert(ctx, meta.GlobalKey(sslCerts.Name), sslCerts)
 if err != nil {
  return nil, mc.Observe(err)
 }
 return g.GetSslCertificate(sslCerts.Name)
}
func (g *Cloud) DeleteSslCertificate(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newCertMetricContext("delete")
 return mc.Observe(g.c.SslCertificates().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) ListSslCertificates() ([]*compute.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newCertMetricContext("list")
 v, err := g.c.SslCertificates().List(ctx, filter.None)
 return v, mc.Observe(err)
}
