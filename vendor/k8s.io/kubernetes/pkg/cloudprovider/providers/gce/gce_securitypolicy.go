package gce

import (
	computebeta "google.golang.org/api/compute/v0.beta"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newSecurityPolicyMetricContextWithVersion(request, version string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("securitypolicy", request, "", unusedMetricLabel, version)
}
func (g *Cloud) GetBetaSecurityPolicy(name string) (*computebeta.SecurityPolicy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("get", computeBetaVersion)
	v, err := g.c.BetaSecurityPolicies().Get(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) ListBetaSecurityPolicy() ([]*computebeta.SecurityPolicy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("list", computeBetaVersion)
	v, err := g.c.BetaSecurityPolicies().List(ctx, filter.None)
	return v, mc.Observe(err)
}
func (g *Cloud) CreateBetaSecurityPolicy(sp *computebeta.SecurityPolicy) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("create", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().Insert(ctx, meta.GlobalKey(sp.Name), sp))
}
func (g *Cloud) DeleteBetaSecurityPolicy(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("delete", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) PatchBetaSecurityPolicy(sp *computebeta.SecurityPolicy) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("patch", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().Patch(ctx, meta.GlobalKey(sp.Name), sp))
}
func (g *Cloud) GetRuleForBetaSecurityPolicy(name string) (*computebeta.SecurityPolicyRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("get_rule", computeBetaVersion)
	v, err := g.c.BetaSecurityPolicies().GetRule(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) AddRuletoBetaSecurityPolicy(name string, spr *computebeta.SecurityPolicyRule) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("add_rule", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().AddRule(ctx, meta.GlobalKey(name), spr))
}
func (g *Cloud) PatchRuleForBetaSecurityPolicy(name string, spr *computebeta.SecurityPolicyRule) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("patch_rule", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().PatchRule(ctx, meta.GlobalKey(name), spr))
}
func (g *Cloud) RemoveRuleFromBetaSecurityPolicy(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newSecurityPolicyMetricContextWithVersion("remove_rule", computeBetaVersion)
	return mc.Observe(g.c.BetaSecurityPolicies().RemoveRule(ctx, meta.GlobalKey(name)))
}
