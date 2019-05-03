package gce

import (
 "context"
 "fmt"
 "net/http"
 "path"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/klog"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newRoutesMetricContext(request string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("routes", request, unusedMetricLabel, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) ListRoutes(ctx context.Context, clusterName string) ([]*cloudprovider.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newRoutesMetricContext("list")
 prefix := truncateClusterName(clusterName)
 f := filter.Regexp("name", prefix+"-.*").AndRegexp("network", g.NetworkURL()).AndRegexp("description", k8sNodeRouteTag)
 routes, err := g.c.Routes().List(ctx, f)
 if err != nil {
  return nil, mc.Observe(err)
 }
 var croutes []*cloudprovider.Route
 for _, r := range routes {
  target := path.Base(r.NextHopInstance)
  targetNodeName := types.NodeName(target)
  croutes = append(croutes, &cloudprovider.Route{Name: r.Name, TargetNode: targetNodeName, DestinationCIDR: r.DestRange})
 }
 return croutes, mc.Observe(nil)
}
func (g *Cloud) CreateRoute(ctx context.Context, clusterName string, nameHint string, route *cloudprovider.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newRoutesMetricContext("create")
 targetInstance, err := g.getInstanceByName(mapNodeNameToInstanceName(route.TargetNode))
 if err != nil {
  return mc.Observe(err)
 }
 cr := &compute.Route{Name: truncateClusterName(clusterName) + "-" + nameHint, DestRange: route.DestinationCIDR, NextHopInstance: fmt.Sprintf("zones/%s/instances/%s", targetInstance.Zone, targetInstance.Name), Network: g.NetworkURL(), Priority: 1000, Description: k8sNodeRouteTag}
 err = g.c.Routes().Insert(ctx, meta.GlobalKey(cr.Name), cr)
 if isHTTPErrorCode(err, http.StatusConflict) {
  klog.Infof("Route %q already exists.", cr.Name)
  err = nil
 }
 return mc.Observe(err)
}
func (g *Cloud) DeleteRoute(ctx context.Context, clusterName string, route *cloudprovider.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newRoutesMetricContext("delete")
 return mc.Observe(g.c.Routes().Delete(ctx, meta.GlobalKey(route.Name)))
}
func truncateClusterName(clusterName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(clusterName) > 26 {
  return clusterName[:26]
 }
 return clusterName
}
