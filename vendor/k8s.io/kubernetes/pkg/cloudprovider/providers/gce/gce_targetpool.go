package gce

import (
	compute "google.golang.org/api/compute/v1"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newTargetPoolMetricContext(request, region string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("targetpool", request, region, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetTargetPool(name, region string) (*compute.TargetPool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newTargetPoolMetricContext("get", region)
	v, err := g.c.TargetPools().Get(ctx, meta.RegionalKey(name, region))
	return v, mc.Observe(err)
}
func (g *Cloud) CreateTargetPool(tp *compute.TargetPool, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newTargetPoolMetricContext("create", region)
	return mc.Observe(g.c.TargetPools().Insert(ctx, meta.RegionalKey(tp.Name, region), tp))
}
func (g *Cloud) DeleteTargetPool(name, region string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newTargetPoolMetricContext("delete", region)
	return mc.Observe(g.c.TargetPools().Delete(ctx, meta.RegionalKey(name, region)))
}
func (g *Cloud) AddInstancesToTargetPool(name, region string, instanceRefs []*compute.InstanceReference) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	req := &compute.TargetPoolsAddInstanceRequest{Instances: instanceRefs}
	mc := newTargetPoolMetricContext("add_instances", region)
	return mc.Observe(g.c.TargetPools().AddInstance(ctx, meta.RegionalKey(name, region), req))
}
func (g *Cloud) RemoveInstancesFromTargetPool(name, region string, instanceRefs []*compute.InstanceReference) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	req := &compute.TargetPoolsRemoveInstanceRequest{Instances: instanceRefs}
	mc := newTargetPoolMetricContext("remove_instances", region)
	return mc.Observe(g.c.TargetPools().RemoveInstance(ctx, meta.RegionalKey(name, region), req))
}
