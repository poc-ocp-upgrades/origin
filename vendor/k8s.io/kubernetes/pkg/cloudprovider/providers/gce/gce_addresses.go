package gce

import (
	"fmt"
	computealpha "google.golang.org/api/compute/v0.alpha"
	computebeta "google.golang.org/api/compute/v0.beta"
	compute "google.golang.org/api/compute/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newAddressMetricContext(request, region string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newAddressMetricContextWithVersion(request, region, computeV1Version)
}
func newAddressMetricContextWithVersion(request, region, version string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("address", request, region, unusedMetricLabel, version)
}
func (g *Cloud) ReserveGlobalAddress(addr *compute.Address) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("reserve", "")
	return mc.Observe(g.c.GlobalAddresses().Insert(ctx, meta.GlobalKey(addr.Name), addr))
}
func (g *Cloud) DeleteGlobalAddress(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("delete", "")
	return mc.Observe(g.c.GlobalAddresses().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) GetGlobalAddress(name string) (*compute.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("get", "")
	v, err := g.c.GlobalAddresses().Get(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) ReserveRegionAddress(addr *compute.Address, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("reserve", region)
	return mc.Observe(g.c.Addresses().Insert(ctx, meta.RegionalKey(addr.Name, region), addr))
}
func (g *Cloud) ReserveAlphaRegionAddress(addr *computealpha.Address, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("reserve", region)
	return mc.Observe(g.c.AlphaAddresses().Insert(ctx, meta.RegionalKey(addr.Name, region), addr))
}
func (g *Cloud) ReserveBetaRegionAddress(addr *computebeta.Address, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("reserve", region)
	return mc.Observe(g.c.BetaAddresses().Insert(ctx, meta.RegionalKey(addr.Name, region), addr))
}
func (g *Cloud) DeleteRegionAddress(name, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("delete", region)
	return mc.Observe(g.c.Addresses().Delete(ctx, meta.RegionalKey(name, region)))
}
func (g *Cloud) GetRegionAddress(name, region string) (*compute.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("get", region)
	v, err := g.c.Addresses().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) GetAlphaRegionAddress(name, region string) (*computealpha.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("get", region)
	v, err := g.c.AlphaAddresses().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) GetBetaRegionAddress(name, region string) (*computebeta.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("get", region)
	v, err := g.c.BetaAddresses().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) GetRegionAddressByIP(region, ipAddress string) (*compute.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("list", region)
	addrs, err := g.c.Addresses().List(ctx, region, filter.Regexp("address", ipAddress))
	mc.Observe(err)
	if err != nil {
		return nil, err
	}
	if len(addrs) > 1 {
		klog.Warningf("More than one addresses matching the IP %q: %v", ipAddress, addrNames(addrs))
	}
	for _, addr := range addrs {
		if addr.Address == ipAddress {
			return addr, nil
		}
	}
	return nil, makeGoogleAPINotFoundError(fmt.Sprintf("Address with IP %q was not found in region %q", ipAddress, region))
}
func (g *Cloud) GetBetaRegionAddressByIP(region, ipAddress string) (*computebeta.Address, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newAddressMetricContext("list", region)
	addrs, err := g.c.BetaAddresses().List(ctx, region, filter.Regexp("address", ipAddress))
	mc.Observe(err)
	if err != nil {
		return nil, err
	}
	if len(addrs) > 1 {
		klog.Warningf("More than one addresses matching the IP %q: %v", ipAddress, addrNames(addrs))
	}
	for _, addr := range addrs {
		if addr.Address == ipAddress {
			return addr, nil
		}
	}
	return nil, makeGoogleAPINotFoundError(fmt.Sprintf("Address with IP %q was not found in region %q", ipAddress, region))
}
func (g *Cloud) getNetworkTierFromAddress(name, region string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !g.AlphaFeatureGate.Enabled(AlphaFeatureNetworkTiers) {
		return cloud.NetworkTierDefault.ToGCEValue(), nil
	}
	addr, err := g.GetAlphaRegionAddress(name, region)
	if err != nil {
		return handleAlphaNetworkTierGetError(err)
	}
	return addr.NetworkTier, nil
}
func addrNames(items interface{}) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var ret []string
	switch items := items.(type) {
	case []compute.Address:
		for _, a := range items {
			ret = append(ret, a.Name)
		}
	case []computebeta.Address:
		for _, a := range items {
			ret = append(ret, a.Name)
		}
	}
	return ret
}
