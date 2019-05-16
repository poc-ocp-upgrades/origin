package allocation

import (
	"github.com/openshift/origin/pkg/route"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	"k8s.io/klog"
)

type RouteAllocationController struct{ Plugin route.AllocationPlugin }

func (c *RouteAllocationController) AllocateRouterShard(route *routeapi.Route) (*routeapi.RouterShard, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Allocating router shard for Route: %s [alias=%s]", route.Name, route.Spec.Host)
	shard, err := c.Plugin.Allocate(route)
	if err != nil {
		klog.Errorf("unable to allocate router shard: %v", err)
		return shard, err
	}
	klog.V(4).Infof("Route %s allocated to shard %s [suffix=%s]", route.Name, shard.ShardName, shard.DNSSuffix)
	return shard, err
}
func (c *RouteAllocationController) GenerateHostname(route *routeapi.Route, shard *routeapi.RouterShard) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Generating host name for Route: %s", route.Name)
	s := c.Plugin.GenerateHostname(route, shard)
	klog.V(4).Infof("Route: %s, generated host name/alias=%s", route.Name, s)
	return s
}
