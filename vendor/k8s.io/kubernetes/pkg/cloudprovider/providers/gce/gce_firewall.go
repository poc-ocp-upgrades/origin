package gce

import (
	compute "google.golang.org/api/compute/v1"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newFirewallMetricContext(request string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("firewall", request, unusedMetricLabel, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetFirewall(name string) (*compute.Firewall, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newFirewallMetricContext("get")
	v, err := g.c.Firewalls().Get(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) CreateFirewall(f *compute.Firewall) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newFirewallMetricContext("create")
	return mc.Observe(g.c.Firewalls().Insert(ctx, meta.GlobalKey(f.Name), f))
}
func (g *Cloud) DeleteFirewall(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newFirewallMetricContext("delete")
	return mc.Observe(g.c.Firewalls().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) UpdateFirewall(f *compute.Firewall) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newFirewallMetricContext("update")
	return mc.Observe(g.c.Firewalls().Update(ctx, meta.GlobalKey(f.Name), f))
}
