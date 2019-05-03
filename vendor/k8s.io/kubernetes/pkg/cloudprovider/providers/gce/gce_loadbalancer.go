package gce

import (
 "context"
 "flag"
 "fmt"
 "net"
 "sort"
 "strings"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 netsets "k8s.io/kubernetes/pkg/util/net/sets"
)

type cidrs struct {
 ipn   netsets.IPNet
 isSet bool
}

var (
 lbSrcRngsFlag cidrs
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 lbSrcRngsFlag.ipn, err = netsets.ParseIPNets([]string{"130.211.0.0/22", "35.191.0.0/16", "209.85.152.0/22", "209.85.204.0/22"}...)
 if err != nil {
  panic("Incorrect default GCE L7 source ranges")
 }
 flag.Var(&lbSrcRngsFlag, "cloud-provider-gce-lb-src-cidrs", "CIDRs opened in GCE firewall for LB traffic proxy & health checks")
}
func (c *cidrs) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s := c.ipn.StringSlice()
 sort.Strings(s)
 return strings.Join(s, ",")
}
func (c *cidrs) Set(value string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !c.isSet {
  c.isSet = true
  c.ipn = make(netsets.IPNet)
 } else {
  return fmt.Errorf("GCE LB CIDRs have already been set")
 }
 for _, cidr := range strings.Split(value, ",") {
  _, ipnet, err := net.ParseCIDR(cidr)
  if err != nil {
   return err
  }
  c.ipn.Insert(ipnet)
 }
 return nil
}
func LoadBalancerSrcRanges() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return lbSrcRngsFlag.ipn.StringSlice()
}
func (g *Cloud) GetLoadBalancer(ctx context.Context, clusterName string, svc *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := g.GetLoadBalancerName(ctx, clusterName, svc)
 fwd, err := g.GetRegionForwardingRule(loadBalancerName, g.region)
 if err == nil {
  status := &v1.LoadBalancerStatus{}
  status.Ingress = []v1.LoadBalancerIngress{{IP: fwd.IPAddress}}
  return status, true, nil
 }
 return nil, false, ignoreNotFound(err)
}
func (g *Cloud) GetLoadBalancerName(ctx context.Context, clusterName string, svc *v1.Service) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.DefaultLoadBalancerName(svc)
}
func (g *Cloud) EnsureLoadBalancer(ctx context.Context, clusterName string, svc *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := g.GetLoadBalancerName(ctx, clusterName, svc)
 desiredScheme := getSvcScheme(svc)
 clusterID, err := g.ClusterID.GetID()
 if err != nil {
  return nil, err
 }
 klog.V(4).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v): ensure %v loadbalancer", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, desiredScheme)
 existingFwdRule, err := g.GetRegionForwardingRule(loadBalancerName, g.region)
 if err != nil && !isNotFound(err) {
  return nil, err
 }
 if existingFwdRule != nil {
  existingScheme := cloud.LbScheme(strings.ToUpper(existingFwdRule.LoadBalancingScheme))
  if existingScheme != desiredScheme {
   klog.V(4).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v): deleting existing %v loadbalancer", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, existingScheme)
   switch existingScheme {
   case cloud.SchemeInternal:
    err = g.ensureInternalLoadBalancerDeleted(clusterName, clusterID, svc)
   default:
    err = g.ensureExternalLoadBalancerDeleted(clusterName, clusterID, svc)
   }
   klog.V(4).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v): done deleting existing %v loadbalancer. err: %v", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, existingScheme, err)
   if err != nil {
    return nil, err
   }
   existingFwdRule = nil
  }
 }
 var status *v1.LoadBalancerStatus
 switch desiredScheme {
 case cloud.SchemeInternal:
  status, err = g.ensureInternalLoadBalancer(clusterName, clusterID, svc, existingFwdRule, nodes)
 default:
  status, err = g.ensureExternalLoadBalancer(clusterName, clusterID, svc, existingFwdRule, nodes)
 }
 klog.V(4).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v): done ensuring loadbalancer. err: %v", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, err)
 return status, err
}
func (g *Cloud) UpdateLoadBalancer(ctx context.Context, clusterName string, svc *v1.Service, nodes []*v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := g.GetLoadBalancerName(ctx, clusterName, svc)
 scheme := getSvcScheme(svc)
 clusterID, err := g.ClusterID.GetID()
 if err != nil {
  return err
 }
 klog.V(4).Infof("UpdateLoadBalancer(%v, %v, %v, %v, %v): updating with %d nodes", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, len(nodes))
 switch scheme {
 case cloud.SchemeInternal:
  err = g.updateInternalLoadBalancer(clusterName, clusterID, svc, nodes)
 default:
  err = g.updateExternalLoadBalancer(clusterName, svc, nodes)
 }
 klog.V(4).Infof("UpdateLoadBalancer(%v, %v, %v, %v, %v): done updating. err: %v", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, err)
 return err
}
func (g *Cloud) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, svc *v1.Service) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := g.GetLoadBalancerName(ctx, clusterName, svc)
 scheme := getSvcScheme(svc)
 clusterID, err := g.ClusterID.GetID()
 if err != nil {
  return err
 }
 klog.V(4).Infof("EnsureLoadBalancerDeleted(%v, %v, %v, %v, %v): deleting loadbalancer", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region)
 switch scheme {
 case cloud.SchemeInternal:
  err = g.ensureInternalLoadBalancerDeleted(clusterName, clusterID, svc)
 default:
  err = g.ensureExternalLoadBalancerDeleted(clusterName, clusterID, svc)
 }
 klog.V(4).Infof("EnsureLoadBalancerDeleted(%v, %v, %v, %v, %v): done deleting loadbalancer. err: %v", clusterName, svc.Namespace, svc.Name, loadBalancerName, g.region, err)
 return err
}
func getSvcScheme(svc *v1.Service) cloud.LbScheme {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if typ, ok := GetLoadBalancerAnnotationType(svc); ok && typ == LBTypeInternal {
  return cloud.SchemeInternal
 }
 return cloud.SchemeExternal
}
