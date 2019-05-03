package gce

import (
 compute "google.golang.org/api/compute/v1"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newInstanceGroupMetricContext(request string, zone string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("instancegroup", request, unusedMetricLabel, zone, computeV1Version)
}
func (g *Cloud) CreateInstanceGroup(ig *compute.InstanceGroup, zone string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("create", zone)
 return mc.Observe(g.c.InstanceGroups().Insert(ctx, meta.ZonalKey(ig.Name, zone), ig))
}
func (g *Cloud) DeleteInstanceGroup(name string, zone string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("delete", zone)
 return mc.Observe(g.c.InstanceGroups().Delete(ctx, meta.ZonalKey(name, zone)))
}
func (g *Cloud) ListInstanceGroups(zone string) ([]*compute.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("list", zone)
 v, err := g.c.InstanceGroups().List(ctx, zone, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) ListInstancesInInstanceGroup(name string, zone string, state string) ([]*compute.InstanceWithNamedPorts, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("list_instances", zone)
 req := &compute.InstanceGroupsListInstancesRequest{InstanceState: state}
 v, err := g.c.InstanceGroups().ListInstances(ctx, meta.ZonalKey(name, zone), req, filter.None)
 return v, mc.Observe(err)
}
func (g *Cloud) AddInstancesToInstanceGroup(name string, zone string, instanceRefs []*compute.InstanceReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("add_instances", zone)
 if len(instanceRefs) == 0 {
  return nil
 }
 req := &compute.InstanceGroupsAddInstancesRequest{Instances: instanceRefs}
 return mc.Observe(g.c.InstanceGroups().AddInstances(ctx, meta.ZonalKey(name, zone), req))
}
func (g *Cloud) RemoveInstancesFromInstanceGroup(name string, zone string, instanceRefs []*compute.InstanceReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("remove_instances", zone)
 if len(instanceRefs) == 0 {
  return nil
 }
 req := &compute.InstanceGroupsRemoveInstancesRequest{Instances: instanceRefs}
 return mc.Observe(g.c.InstanceGroups().RemoveInstances(ctx, meta.ZonalKey(name, zone), req))
}
func (g *Cloud) SetNamedPortsOfInstanceGroup(igName, zone string, namedPorts []*compute.NamedPort) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("set_namedports", zone)
 req := &compute.InstanceGroupsSetNamedPortsRequest{NamedPorts: namedPorts}
 return mc.Observe(g.c.InstanceGroups().SetNamedPorts(ctx, meta.ZonalKey(igName, zone), req))
}
func (g *Cloud) GetInstanceGroup(name string, zone string) (*compute.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstanceGroupMetricContext("get", zone)
 v, err := g.c.InstanceGroups().Get(ctx, meta.ZonalKey(name, zone))
 return v, mc.Observe(err)
}
