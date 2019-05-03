package gce

import (
 "k8s.io/klog"
 computealpha "google.golang.org/api/compute/v0.alpha"
 computebeta "google.golang.org/api/compute/v0.beta"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/api/core/v1"
 utilversion "k8s.io/apimachinery/pkg/util/version"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
 "k8s.io/kubernetes/pkg/master/ports"
)

const (
 nodesHealthCheckPath   = "/healthz"
 lbNodesHealthCheckPort = ports.ProxyHealthzPort
)

var (
 minNodesHealthCheckVersion *utilversion.Version
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if v, err := utilversion.ParseGeneric("1.7.2"); err != nil {
  klog.Fatalf("Failed to parse version for minNodesHealthCheckVersion: %v", err)
 } else {
  minNodesHealthCheckVersion = v
 }
}
func newHealthcheckMetricContext(request string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newHealthcheckMetricContextWithVersion(request, computeV1Version)
}
func newHealthcheckMetricContextWithVersion(request, version string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("healthcheck", request, unusedMetricLabel, unusedMetricLabel, version)
}
func (g *Cloud) GetHTTPHealthCheck(name string) (*compute.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("get_legacy")
 v, err := g.c.HttpHealthChecks().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) UpdateHTTPHealthCheck(hc *compute.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("update_legacy")
 return mc.Observe(g.c.HttpHealthChecks().Update(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) DeleteHTTPHealthCheck(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("delete_legacy")
 return mc.Observe(g.c.HttpHealthChecks().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) CreateHTTPHealthCheck(hc *compute.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("create_legacy")
 return mc.Observe(g.c.HttpHealthChecks().Insert(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) ListHTTPHealthChecks() ([]*compute.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("list_legacy")
 v, err := g.c.HttpHealthChecks().List(ctx, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) GetHTTPSHealthCheck(name string) (*compute.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("get_legacy")
 v, err := g.c.HttpsHealthChecks().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) UpdateHTTPSHealthCheck(hc *compute.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("update_legacy")
 return mc.Observe(g.c.HttpsHealthChecks().Update(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) DeleteHTTPSHealthCheck(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("delete_legacy")
 return mc.Observe(g.c.HttpsHealthChecks().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) CreateHTTPSHealthCheck(hc *compute.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("create_legacy")
 return mc.Observe(g.c.HttpsHealthChecks().Insert(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) ListHTTPSHealthChecks() ([]*compute.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("list_legacy")
 v, err := g.c.HttpsHealthChecks().List(ctx, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) GetHealthCheck(name string) (*compute.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("get")
 v, err := g.c.HealthChecks().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) GetAlphaHealthCheck(name string) (*computealpha.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("get", computeAlphaVersion)
 v, err := g.c.AlphaHealthChecks().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) GetBetaHealthCheck(name string) (*computebeta.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("get", computeBetaVersion)
 v, err := g.c.BetaHealthChecks().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) UpdateHealthCheck(hc *compute.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("update")
 return mc.Observe(g.c.HealthChecks().Update(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) UpdateAlphaHealthCheck(hc *computealpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("update", computeAlphaVersion)
 return mc.Observe(g.c.AlphaHealthChecks().Update(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) UpdateBetaHealthCheck(hc *computebeta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("update", computeBetaVersion)
 return mc.Observe(g.c.BetaHealthChecks().Update(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) DeleteHealthCheck(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("delete")
 return mc.Observe(g.c.HealthChecks().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) CreateHealthCheck(hc *compute.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("create")
 return mc.Observe(g.c.HealthChecks().Insert(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) CreateAlphaHealthCheck(hc *computealpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("create", computeAlphaVersion)
 return mc.Observe(g.c.AlphaHealthChecks().Insert(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) CreateBetaHealthCheck(hc *computebeta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContextWithVersion("create", computeBetaVersion)
 return mc.Observe(g.c.BetaHealthChecks().Insert(ctx, meta.GlobalKey(hc.Name), hc))
}
func (g *Cloud) ListHealthChecks() ([]*compute.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newHealthcheckMetricContext("list")
 v, err := g.c.HealthChecks().List(ctx, filter.None)
 return v, mc.Observe(err)
}
func GetNodesHealthCheckPort() int32 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return lbNodesHealthCheckPort
}
func GetNodesHealthCheckPath() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nodesHealthCheckPath
}
func isAtLeastMinNodesHealthCheckVersion(vstring string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 version, err := utilversion.ParseGeneric(vstring)
 if err != nil {
  klog.Errorf("vstring (%s) is not a valid version string: %v", vstring, err)
  return false
 }
 return version.AtLeast(minNodesHealthCheckVersion)
}
func supportsNodesHealthCheck(nodes []*v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, node := range nodes {
  if !isAtLeastMinNodesHealthCheckVersion(node.Status.NodeInfo.KubeProxyVersion) {
   return false
  }
 }
 return true
}
