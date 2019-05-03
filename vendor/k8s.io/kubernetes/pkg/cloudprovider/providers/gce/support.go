package gce

import (
 "context"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

type gceProjectRouter struct{ gce *Cloud }

func (r *gceProjectRouter) ProjectID(ctx context.Context, version meta.Version, service string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch service {
 case "Firewalls", "Routes":
  return r.gce.NetworkProjectID()
 default:
  return r.gce.projectID
 }
}

type gceRateLimiter struct{ gce *Cloud }

func (l *gceRateLimiter) Accept(ctx context.Context, key *cloud.RateLimitKey) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if key.Operation == "Get" && key.Service == "Operations" {
  rl := &cloud.MinimumRateLimiter{RateLimiter: &cloud.AcceptRateLimiter{Acceptor: l.gce.operationPollRateLimiter}, Minimum: operationPollInterval}
  return rl.Accept(ctx, key)
 }
 return nil
}
func CreateGCECloudWithCloud(config *CloudConfig, c cloud.Cloud) (*Cloud, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 gceCloud, err := CreateGCECloud(config)
 if err == nil {
  gceCloud.c = c
 }
 return gceCloud, err
}
