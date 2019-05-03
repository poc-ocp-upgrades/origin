package gce

import (
 "fmt"
 "strings"
 computebeta "google.golang.org/api/compute/v0.beta"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newNetworkEndpointGroupMetricContext(request string, zone string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("networkendpointgroup_", request, unusedMetricLabel, zone, computeBetaVersion)
}
func (g *Cloud) GetNetworkEndpointGroup(name string, zone string) (*computebeta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("get", zone)
 v, err := g.c.BetaNetworkEndpointGroups().Get(ctx, meta.ZonalKey(name, zone))
 return v, mc.Observe(err)
}
func (g *Cloud) ListNetworkEndpointGroup(zone string) ([]*computebeta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("list", zone)
 negs, err := g.c.BetaNetworkEndpointGroups().List(ctx, zone, filter.None)
 return negs, mc.Observe(err)
}
func (g *Cloud) AggregatedListNetworkEndpointGroup() (map[string][]*computebeta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("aggregated_list", "")
 all, err := g.c.BetaNetworkEndpointGroups().AggregatedList(ctx, filter.None)
 if err != nil {
  return nil, mc.Observe(err)
 }
 ret := map[string][]*computebeta.NetworkEndpointGroup{}
 for key, byZone := range all {
  parts := strings.Split(key, "/")
  if len(parts) != 2 {
   return nil, mc.Observe(fmt.Errorf("invalid key for AggregatedListNetworkEndpointGroup: %q", key))
  }
  zone := parts[1]
  ret[zone] = append(ret[zone], byZone...)
 }
 return ret, mc.Observe(nil)
}
func (g *Cloud) CreateNetworkEndpointGroup(neg *computebeta.NetworkEndpointGroup, zone string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("create", zone)
 return mc.Observe(g.c.BetaNetworkEndpointGroups().Insert(ctx, meta.ZonalKey(neg.Name, zone), neg))
}
func (g *Cloud) DeleteNetworkEndpointGroup(name string, zone string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("delete", zone)
 return mc.Observe(g.c.BetaNetworkEndpointGroups().Delete(ctx, meta.ZonalKey(name, zone)))
}
func (g *Cloud) AttachNetworkEndpoints(name, zone string, endpoints []*computebeta.NetworkEndpoint) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("attach", zone)
 req := &computebeta.NetworkEndpointGroupsAttachEndpointsRequest{NetworkEndpoints: endpoints}
 return mc.Observe(g.c.BetaNetworkEndpointGroups().AttachNetworkEndpoints(ctx, meta.ZonalKey(name, zone), req))
}
func (g *Cloud) DetachNetworkEndpoints(name, zone string, endpoints []*computebeta.NetworkEndpoint) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("detach", zone)
 req := &computebeta.NetworkEndpointGroupsDetachEndpointsRequest{NetworkEndpoints: endpoints}
 return mc.Observe(g.c.BetaNetworkEndpointGroups().DetachNetworkEndpoints(ctx, meta.ZonalKey(name, zone), req))
}
func (g *Cloud) ListNetworkEndpoints(name, zone string, showHealthStatus bool) ([]*computebeta.NetworkEndpointWithHealthStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newNetworkEndpointGroupMetricContext("list_networkendpoints", zone)
 healthStatus := "SKIP"
 if showHealthStatus {
  healthStatus = "SHOW"
 }
 req := &computebeta.NetworkEndpointGroupsListEndpointsRequest{HealthStatus: healthStatus}
 l, err := g.c.BetaNetworkEndpointGroups().ListNetworkEndpoints(ctx, meta.ZonalKey(name, zone), req, filter.None)
 return l, mc.Observe(err)
}
