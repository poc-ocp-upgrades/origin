package gce

import (
 "context"
 "strings"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/apimachinery/pkg/types"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
)

func newZonesMetricContext(request, region string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("zones", request, region, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.Zone{FailureDomain: g.localZone, Region: g.region}, nil
}
func (g *Cloud) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, zone, _, err := splitProviderID(providerID)
 if err != nil {
  return cloudprovider.Zone{}, err
 }
 region, err := GetGCERegion(zone)
 if err != nil {
  return cloudprovider.Zone{}, err
 }
 return cloudprovider.Zone{FailureDomain: zone, Region: region}, nil
}
func (g *Cloud) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceName := mapNodeNameToInstanceName(nodeName)
 instance, err := g.getInstanceByName(instanceName)
 if err != nil {
  return cloudprovider.Zone{}, err
 }
 region, err := GetGCERegion(instance.Zone)
 if err != nil {
  return cloudprovider.Zone{}, err
 }
 return cloudprovider.Zone{FailureDomain: instance.Zone, Region: region}, nil
}
func (g *Cloud) ListZonesInRegion(region string) ([]*compute.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newZonesMetricContext("list", region)
 list, err := g.c.Zones().List(ctx, filter.Regexp("region", g.getRegionLink(region)))
 if err != nil {
  return nil, mc.Observe(err)
 }
 return list, mc.Observe(err)
}
func (g *Cloud) getRegionLink(region string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return g.service.BasePath + strings.Join([]string{g.projectID, "regions", region}, "/")
}
