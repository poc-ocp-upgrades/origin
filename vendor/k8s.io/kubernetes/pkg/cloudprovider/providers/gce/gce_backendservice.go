package gce

import (
 computealpha "google.golang.org/api/compute/v0.alpha"
 computebeta "google.golang.org/api/compute/v0.beta"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newBackendServiceMetricContext(request, region string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newBackendServiceMetricContextWithVersion(request, region, computeV1Version)
}
func newBackendServiceMetricContextWithVersion(request, region, version string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("backendservice", request, region, unusedMetricLabel, version)
}
func (g *Cloud) GetGlobalBackendService(name string) (*compute.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("get", "")
 v, err := g.c.BackendServices().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) GetBetaGlobalBackendService(name string) (*computebeta.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("get", "", computeBetaVersion)
 v, err := g.c.BetaBackendServices().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) GetAlphaGlobalBackendService(name string) (*computealpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("get", "", computeAlphaVersion)
 v, err := g.c.AlphaBackendServices().Get(ctx, meta.GlobalKey(name))
 return v, mc.Observe(err)
}
func (g *Cloud) UpdateGlobalBackendService(bg *compute.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("update", "")
 return mc.Observe(g.c.BackendServices().Update(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) UpdateBetaGlobalBackendService(bg *computebeta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("update", "", computeBetaVersion)
 return mc.Observe(g.c.BetaBackendServices().Update(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) UpdateAlphaGlobalBackendService(bg *computealpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("update", "", computeAlphaVersion)
 return mc.Observe(g.c.AlphaBackendServices().Update(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) DeleteGlobalBackendService(name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("delete", "")
 return mc.Observe(g.c.BackendServices().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) CreateGlobalBackendService(bg *compute.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("create", "")
 return mc.Observe(g.c.BackendServices().Insert(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) CreateBetaGlobalBackendService(bg *computebeta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("create", "", computeBetaVersion)
 return mc.Observe(g.c.BetaBackendServices().Insert(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) CreateAlphaGlobalBackendService(bg *computealpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("create", "", computeAlphaVersion)
 return mc.Observe(g.c.AlphaBackendServices().Insert(ctx, meta.GlobalKey(bg.Name), bg))
}
func (g *Cloud) ListGlobalBackendServices() ([]*compute.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("list", "")
 v, err := g.c.BackendServices().List(ctx, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) GetGlobalBackendServiceHealth(name string, instanceGroupLink string) (*compute.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("get_health", "")
 groupRef := &compute.ResourceGroupReference{Group: instanceGroupLink}
 v, err := g.c.BackendServices().GetHealth(ctx, meta.GlobalKey(name), groupRef)
 return v, mc.Observe(err)
}
func (g *Cloud) GetRegionBackendService(name, region string) (*compute.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("get", region)
 v, err := g.c.RegionBackendServices().Get(ctx, meta.RegionalKey(name, region))
 return v, mc.Observe(err)
}
func (g *Cloud) UpdateRegionBackendService(bg *compute.BackendService, region string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("update", region)
 return mc.Observe(g.c.RegionBackendServices().Update(ctx, meta.RegionalKey(bg.Name, region), bg))
}
func (g *Cloud) DeleteRegionBackendService(name, region string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("delete", region)
 return mc.Observe(g.c.RegionBackendServices().Delete(ctx, meta.RegionalKey(name, region)))
}
func (g *Cloud) CreateRegionBackendService(bg *compute.BackendService, region string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("create", region)
 return mc.Observe(g.c.RegionBackendServices().Insert(ctx, meta.RegionalKey(bg.Name, region), bg))
}
func (g *Cloud) ListRegionBackendServices(region string) ([]*compute.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("list", region)
 v, err := g.c.RegionBackendServices().List(ctx, region, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) GetRegionalBackendServiceHealth(name, region string, instanceGroupLink string) (*compute.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContext("get_health", region)
 ref := &compute.ResourceGroupReference{Group: instanceGroupLink}
 v, err := g.c.RegionBackendServices().GetHealth(ctx, meta.RegionalKey(name, region), ref)
 return v, mc.Observe(err)
}
func (g *Cloud) SetSecurityPolicyForBetaGlobalBackendService(backendServiceName string, securityPolicyReference *computebeta.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("set_security_policy", "", computeBetaVersion)
 return mc.Observe(g.c.BetaBackendServices().SetSecurityPolicy(ctx, meta.GlobalKey(backendServiceName), securityPolicyReference))
}
func (g *Cloud) SetSecurityPolicyForAlphaGlobalBackendService(backendServiceName string, securityPolicyReference *computealpha.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newBackendServiceMetricContextWithVersion("set_security_policy", "", computeAlphaVersion)
 return mc.Observe(g.c.AlphaBackendServices().SetSecurityPolicy(ctx, meta.GlobalKey(backendServiceName), securityPolicyReference))
}
