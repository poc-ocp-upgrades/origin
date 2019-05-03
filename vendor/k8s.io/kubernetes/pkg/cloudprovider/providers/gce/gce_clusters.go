package gce

import (
 "context"
 "fmt"
 "google.golang.org/api/container/v1"
 "k8s.io/klog"
)

func newClustersMetricContext(request, zone string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("clusters", request, unusedMetricLabel, zone, computeV1Version)
}
func (g *Cloud) ListClusters(ctx context.Context) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allClusters := []string{}
 for _, zone := range g.managedZones {
  clusters, err := g.listClustersInZone(zone)
  if err != nil {
   return nil, err
  }
  allClusters = append(allClusters, clusters...)
 }
 return allClusters, nil
}
func (g *Cloud) GetManagedClusters(ctx context.Context) ([]*container.Cluster, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 managedClusters := []*container.Cluster{}
 if g.regional {
  var err error
  managedClusters, err = g.getClustersInLocation(g.region)
  if err != nil {
   return nil, err
  }
 } else if len(g.managedZones) >= 1 {
  for _, zone := range g.managedZones {
   clusters, err := g.getClustersInLocation(zone)
   if err != nil {
    return nil, err
   }
   managedClusters = append(managedClusters, clusters...)
  }
 } else {
  return nil, fmt.Errorf("no zones associated with this cluster(%s)", g.ProjectID())
 }
 return managedClusters, nil
}
func (g *Cloud) Master(ctx context.Context, clusterName string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "k8s-" + clusterName + "-master.internal", nil
}
func (g *Cloud) listClustersInZone(zone string) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clusters, err := g.getClustersInLocation(zone)
 if err != nil {
  return nil, err
 }
 result := []string{}
 for _, cluster := range clusters {
  result = append(result, cluster.Name)
 }
 return result, nil
}
func (g *Cloud) getClustersInLocation(zoneOrRegion string) ([]*container.Cluster, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mc := newClustersMetricContext("list_zone", zoneOrRegion)
 location := getLocationName(g.projectID, zoneOrRegion)
 list, err := g.containerService.Projects.Locations.Clusters.List(location).Do()
 if err != nil {
  return nil, mc.Observe(err)
 }
 if list.Header.Get("nextPageToken") != "" {
  klog.Errorf("Failed to get all clusters for request, received next page token %s", list.Header.Get("nextPageToken"))
 }
 return list.Clusters, mc.Observe(nil)
}
