package gce

import (
	computealpha "google.golang.org/api/compute/v0.alpha"
	compute "google.golang.org/api/compute/v1"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newForwardingRuleMetricContext(request, region string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newForwardingRuleMetricContextWithVersion(request, region, computeV1Version)
}
func newForwardingRuleMetricContextWithVersion(request, region, version string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("forwardingrule", request, region, unusedMetricLabel, version)
}
func (g *Cloud) CreateGlobalForwardingRule(rule *compute.ForwardingRule) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("create", "")
	return mc.Observe(g.c.GlobalForwardingRules().Insert(ctx, meta.GlobalKey(rule.Name), rule))
}
func (g *Cloud) SetProxyForGlobalForwardingRule(forwardingRuleName, targetProxyLink string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("set_proxy", "")
	target := &compute.TargetReference{Target: targetProxyLink}
	return mc.Observe(g.c.GlobalForwardingRules().SetTarget(ctx, meta.GlobalKey(forwardingRuleName), target))
}
func (g *Cloud) DeleteGlobalForwardingRule(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("delete", "")
	return mc.Observe(g.c.GlobalForwardingRules().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) GetGlobalForwardingRule(name string) (*compute.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("get", "")
	v, err := g.c.GlobalForwardingRules().Get(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) ListGlobalForwardingRules() ([]*compute.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("list", "")
	v, err := g.c.GlobalForwardingRules().List(ctx, filter.None)
	return v, mc.Observe(err)
}
func (g *Cloud) GetRegionForwardingRule(name, region string) (*compute.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("get", region)
	v, err := g.c.ForwardingRules().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) GetAlphaRegionForwardingRule(name, region string) (*computealpha.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContextWithVersion("get", region, computeAlphaVersion)
	v, err := g.c.AlphaForwardingRules().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) ListRegionForwardingRules(region string) ([]*compute.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("list", region)
	v, err := g.c.ForwardingRules().List(ctx, region, filter.None)
	return v, mc.Observe(err)
}
func (g *Cloud) ListAlphaRegionForwardingRules(region string) ([]*computealpha.ForwardingRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContextWithVersion("list", region, computeAlphaVersion)
	v, err := g.c.AlphaForwardingRules().List(ctx, region, filter.None)
	return v, mc.Observe(err)
}
func (g *Cloud) CreateRegionForwardingRule(rule *compute.ForwardingRule, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("create", region)
	return mc.Observe(g.c.ForwardingRules().Insert(ctx, meta.RegionalKey(rule.Name, region), rule))
}
func (g *Cloud) CreateAlphaRegionForwardingRule(rule *computealpha.ForwardingRule, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContextWithVersion("create", region, computeAlphaVersion)
	return mc.Observe(g.c.AlphaForwardingRules().Insert(ctx, meta.RegionalKey(rule.Name, region), rule))
}
func (g *Cloud) DeleteRegionForwardingRule(name, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newForwardingRuleMetricContext("delete", region)
	return mc.Observe(g.c.ForwardingRules().Delete(ctx, meta.RegionalKey(name, region)))
}
func (g *Cloud) getNetworkTierFromForwardingRule(name, region string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !g.AlphaFeatureGate.Enabled(AlphaFeatureNetworkTiers) {
		return cloud.NetworkTierDefault.ToGCEValue(), nil
	}
	fwdRule, err := g.GetAlphaRegionForwardingRule(name, region)
	if err != nil {
		return handleAlphaNetworkTierGetError(err)
	}
	return fwdRule.NetworkTier, nil
}
