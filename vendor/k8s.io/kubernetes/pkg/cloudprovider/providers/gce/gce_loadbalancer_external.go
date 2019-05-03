package gce

import (
 "context"
 "fmt"
 "net/http"
 "strconv"
 "strings"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 "k8s.io/apimachinery/pkg/util/sets"
 apiservice "k8s.io/kubernetes/pkg/api/v1/service"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 netsets "k8s.io/kubernetes/pkg/util/net/sets"
 computealpha "google.golang.org/api/compute/v0.alpha"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/klog"
)

func (g *Cloud) ensureExternalLoadBalancer(clusterName string, clusterID string, apiService *v1.Service, existingFwdRule *compute.ForwardingRule, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(nodes) == 0 {
  return nil, fmt.Errorf("Cannot EnsureLoadBalancer() with no hosts")
 }
 hostNames := nodeNames(nodes)
 supportsNodesHealthCheck := supportsNodesHealthCheck(nodes)
 hosts, err := g.getInstancesByNames(hostNames)
 if err != nil {
  return nil, err
 }
 loadBalancerName := g.GetLoadBalancerName(context.TODO(), clusterName, apiService)
 requestedIP := apiService.Spec.LoadBalancerIP
 ports := apiService.Spec.Ports
 portStr := []string{}
 for _, p := range apiService.Spec.Ports {
  portStr = append(portStr, fmt.Sprintf("%s/%d", p.Protocol, p.Port))
 }
 serviceName := types.NamespacedName{Namespace: apiService.Namespace, Name: apiService.Name}
 lbRefStr := fmt.Sprintf("%v(%v)", loadBalancerName, serviceName)
 klog.V(2).Infof("ensureExternalLoadBalancer(%s, %v, %v, %v, %v, %v)", lbRefStr, g.region, requestedIP, portStr, hostNames, apiService.Annotations)
 netTier, err := g.getServiceNetworkTier(apiService)
 if err != nil {
  klog.Errorf("ensureExternalLoadBalancer(%s): Failed to get the desired network tier: %v.", lbRefStr, err)
  return nil, err
 }
 klog.V(4).Infof("ensureExternalLoadBalancer(%s): Desired network tier %q.", lbRefStr, netTier)
 if g.AlphaFeatureGate.Enabled(AlphaFeatureNetworkTiers) {
  g.deleteWrongNetworkTieredResources(loadBalancerName, lbRefStr, netTier)
 }
 fwdRuleExists, fwdRuleNeedsUpdate, fwdRuleIP, err := g.forwardingRuleNeedsUpdate(loadBalancerName, g.region, requestedIP, ports)
 if err != nil {
  return nil, err
 }
 if !fwdRuleExists {
  klog.V(2).Infof("ensureExternalLoadBalancer(%s): Forwarding rule %v doesn't exist.", lbRefStr, loadBalancerName)
 }
 ipAddressToUse := ""
 isUserOwnedIP := false
 isSafeToReleaseIP := false
 defer func() {
  if isUserOwnedIP {
   return
  }
  if isSafeToReleaseIP {
   if err := g.DeleteRegionAddress(loadBalancerName, g.region); err != nil && !isNotFound(err) {
    klog.Errorf("ensureExternalLoadBalancer(%s): Failed to release static IP %s in region %v: %v.", lbRefStr, ipAddressToUse, g.region, err)
   } else if isNotFound(err) {
    klog.V(2).Infof("ensureExternalLoadBalancer(%s): IP address %s is not reserved.", lbRefStr, ipAddressToUse)
   } else {
    klog.Infof("ensureExternalLoadBalancer(%s): Released static IP %s.", lbRefStr, ipAddressToUse)
   }
  } else {
   klog.Warningf("ensureExternalLoadBalancer(%s): Orphaning static IP %s in region %v: %v.", lbRefStr, ipAddressToUse, g.region, err)
  }
 }()
 if requestedIP != "" {
  isUserOwnedIP, err = verifyUserRequestedIP(g, g.region, requestedIP, fwdRuleIP, lbRefStr, netTier)
  if err != nil {
   return nil, err
  }
  ipAddressToUse = requestedIP
 }
 if !isUserOwnedIP {
  ipAddr, existed, err := ensureStaticIP(g, loadBalancerName, serviceName.String(), g.region, fwdRuleIP, netTier)
  if err != nil {
   return nil, fmt.Errorf("failed to ensure a static IP for load balancer (%s): %v", lbRefStr, err)
  }
  klog.Infof("ensureExternalLoadBalancer(%s): Ensured IP address %s (tier: %s).", lbRefStr, ipAddr, netTier)
  isSafeToReleaseIP = !existed
  ipAddressToUse = ipAddr
 }
 sourceRanges, err := apiservice.GetLoadBalancerSourceRanges(apiService)
 if err != nil {
  return nil, err
 }
 firewallExists, firewallNeedsUpdate, err := g.firewallNeedsUpdate(loadBalancerName, serviceName.String(), g.region, ipAddressToUse, ports, sourceRanges)
 if err != nil {
  return nil, err
 }
 if firewallNeedsUpdate {
  desc := makeFirewallDescription(serviceName.String(), ipAddressToUse)
  if firewallExists {
   klog.Infof("ensureExternalLoadBalancer(%s): Updating firewall.", lbRefStr)
   if err := g.updateFirewall(apiService, MakeFirewallName(loadBalancerName), g.region, desc, sourceRanges, ports, hosts); err != nil {
    return nil, err
   }
   klog.Infof("ensureExternalLoadBalancer(%s): Updated firewall.", lbRefStr)
  } else {
   klog.Infof("ensureExternalLoadBalancer(%s): Creating firewall.", lbRefStr)
   if err := g.createFirewall(apiService, MakeFirewallName(loadBalancerName), g.region, desc, sourceRanges, ports, hosts); err != nil {
    return nil, err
   }
   klog.Infof("ensureExternalLoadBalancer(%s): Created firewall.", lbRefStr)
  }
 }
 tpExists, tpNeedsRecreation, err := g.targetPoolNeedsRecreation(loadBalancerName, g.region, apiService.Spec.SessionAffinity)
 if err != nil {
  return nil, err
 }
 if !tpExists {
  klog.Infof("ensureExternalLoadBalancer(%s): Target pool for service doesn't exist.", lbRefStr)
 }
 var hcToCreate, hcToDelete *compute.HttpHealthCheck
 hcLocalTrafficExisting, err := g.GetHTTPHealthCheck(loadBalancerName)
 if err != nil && !isHTTPErrorCode(err, http.StatusNotFound) {
  return nil, fmt.Errorf("error checking HTTP health check for load balancer (%s): %v", lbRefStr, err)
 }
 if path, healthCheckNodePort := apiservice.GetServiceHealthCheckPathPort(apiService); path != "" {
  klog.V(4).Infof("ensureExternalLoadBalancer(%s): Service needs local traffic health checks on: %d%s.", lbRefStr, healthCheckNodePort, path)
  if hcLocalTrafficExisting == nil {
   klog.V(2).Infof("ensureExternalLoadBalancer(%s): Updating from nodes health checks to local traffic health checks.", lbRefStr)
   if supportsNodesHealthCheck {
    hcToDelete = makeHTTPHealthCheck(MakeNodesHealthCheckName(clusterID), GetNodesHealthCheckPath(), GetNodesHealthCheckPort())
   }
   tpNeedsRecreation = true
  }
  hcToCreate = makeHTTPHealthCheck(loadBalancerName, path, healthCheckNodePort)
 } else {
  klog.V(4).Infof("ensureExternalLoadBalancer(%s): Service needs nodes health checks.", lbRefStr)
  if hcLocalTrafficExisting != nil {
   klog.V(2).Infof("ensureExternalLoadBalancer(%s): Updating from local traffic health checks to nodes health checks.", lbRefStr)
   hcToDelete = hcLocalTrafficExisting
   tpNeedsRecreation = true
  }
  if supportsNodesHealthCheck {
   hcToCreate = makeHTTPHealthCheck(MakeNodesHealthCheckName(clusterID), GetNodesHealthCheckPath(), GetNodesHealthCheckPort())
  }
 }
 if fwdRuleExists && (fwdRuleNeedsUpdate || tpNeedsRecreation) {
  isSafeToReleaseIP = false
  if err := g.DeleteRegionForwardingRule(loadBalancerName, g.region); err != nil && !isNotFound(err) {
   return nil, fmt.Errorf("failed to delete existing forwarding rule for load balancer (%s) update: %v", lbRefStr, err)
  }
  klog.Infof("ensureExternalLoadBalancer(%s): Deleted forwarding rule.", lbRefStr)
 }
 if err := g.ensureTargetPoolAndHealthCheck(tpExists, tpNeedsRecreation, apiService, loadBalancerName, clusterID, ipAddressToUse, hosts, hcToCreate, hcToDelete); err != nil {
  return nil, err
 }
 if tpNeedsRecreation || fwdRuleNeedsUpdate {
  klog.Infof("ensureExternalLoadBalancer(%s): Creating forwarding rule, IP %s (tier: %s).", lbRefStr, ipAddressToUse, netTier)
  if err := createForwardingRule(g, loadBalancerName, serviceName.String(), g.region, ipAddressToUse, g.targetPoolURL(loadBalancerName), ports, netTier); err != nil {
   return nil, fmt.Errorf("failed to create forwarding rule for load balancer (%s): %v", lbRefStr, err)
  }
  isSafeToReleaseIP = true
  klog.Infof("ensureExternalLoadBalancer(%s): Created forwarding rule, IP %s.", lbRefStr, ipAddressToUse)
 }
 status := &v1.LoadBalancerStatus{}
 status.Ingress = []v1.LoadBalancerIngress{{IP: ipAddressToUse}}
 return status, nil
}
func (g *Cloud) updateExternalLoadBalancer(clusterName string, service *v1.Service, nodes []*v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hosts, err := g.getInstancesByNames(nodeNames(nodes))
 if err != nil {
  return err
 }
 loadBalancerName := g.GetLoadBalancerName(context.TODO(), clusterName, service)
 return g.updateTargetPool(loadBalancerName, hosts)
}
func (g *Cloud) ensureExternalLoadBalancerDeleted(clusterName, clusterID string, service *v1.Service) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := g.GetLoadBalancerName(context.TODO(), clusterName, service)
 serviceName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
 lbRefStr := fmt.Sprintf("%v(%v)", loadBalancerName, serviceName)
 var hcNames []string
 if path, _ := apiservice.GetServiceHealthCheckPathPort(service); path != "" {
  hcToDelete, err := g.GetHTTPHealthCheck(loadBalancerName)
  if err != nil && !isHTTPErrorCode(err, http.StatusNotFound) {
   klog.Infof("ensureExternalLoadBalancerDeleted(%s): Failed to retrieve health check:%v.", lbRefStr, err)
   return err
  }
  if err == nil {
   hcNames = append(hcNames, hcToDelete.Name)
  }
 } else {
  hcNames = append(hcNames, loadBalancerName)
  hcNames = append(hcNames, MakeNodesHealthCheckName(clusterID))
 }
 errs := utilerrors.AggregateGoroutines(func() error {
  klog.Infof("ensureExternalLoadBalancerDeleted(%s): Deleting firewall rule.", lbRefStr)
  fwName := MakeFirewallName(loadBalancerName)
  err := ignoreNotFound(g.DeleteFirewall(fwName))
  if isForbidden(err) && g.OnXPN() {
   klog.V(4).Infof("ensureExternalLoadBalancerDeleted(%s): Do not have permission to delete firewall rule %v (on XPN). Raising event.", lbRefStr, fwName)
   g.raiseFirewallChangeNeededEvent(service, FirewallToGCloudDeleteCmd(fwName, g.NetworkProjectID()))
   return nil
  }
  return err
 }, func() error {
  klog.Infof("ensureExternalLoadBalancerDeleted(%s): Deleting IP address.", lbRefStr)
  return ignoreNotFound(g.DeleteRegionAddress(loadBalancerName, g.region))
 }, func() error {
  klog.Infof("ensureExternalLoadBalancerDeleted(%s): Deleting forwarding rule.", lbRefStr)
  if err := ignoreNotFound(g.DeleteRegionForwardingRule(loadBalancerName, g.region)); err != nil {
   return err
  }
  klog.Infof("ensureExternalLoadBalancerDeleted(%s): Deleting target pool.", lbRefStr)
  if err := g.DeleteExternalTargetPoolAndChecks(service, loadBalancerName, g.region, clusterID, hcNames...); err != nil {
   return err
  }
  return nil
 })
 if errs != nil {
  return utilerrors.Flatten(errs)
 }
 return nil
}
func (g *Cloud) DeleteExternalTargetPoolAndChecks(service *v1.Service, name, region, clusterID string, hcNames ...string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 serviceName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
 lbRefStr := fmt.Sprintf("%v(%v)", name, serviceName)
 if err := g.DeleteTargetPool(name, region); err != nil && isHTTPErrorCode(err, http.StatusNotFound) {
  klog.Infof("DeleteExternalTargetPoolAndChecks(%v): Target pool already deleted. Continuing to delete other resources.", lbRefStr)
 } else if err != nil {
  klog.Warningf("DeleteExternalTargetPoolAndChecks(%v): Failed to delete target pool, got error %s.", lbRefStr, err.Error())
  return err
 }
 for _, hcName := range hcNames {
  if err := func() error {
   isNodesHealthCheck := hcName != name
   if isNodesHealthCheck {
    g.sharedResourceLock.Lock()
    defer g.sharedResourceLock.Unlock()
   }
   klog.Infof("DeleteExternalTargetPoolAndChecks(%v): Deleting health check %v.", lbRefStr, hcName)
   if err := g.DeleteHTTPHealthCheck(hcName); err != nil {
    if isInUsedByError(err) {
     klog.V(4).Infof("DeleteExternalTargetPoolAndChecks(%v): Health check %v is in used: %v.", lbRefStr, hcName, err)
     return nil
    } else if !isHTTPErrorCode(err, http.StatusNotFound) {
     klog.Warningf("DeleteExternalTargetPoolAndChecks(%v): Failed to delete health check %v: %v.", lbRefStr, hcName, err)
     return err
    }
    klog.V(4).Infof("DeleteExternalTargetPoolAndChecks(%v): Health check %v is already deleted.", lbRefStr, hcName)
   }
   fwName := MakeHealthCheckFirewallName(clusterID, hcName, isNodesHealthCheck)
   klog.Infof("DeleteExternalTargetPoolAndChecks(%v): Deleting health check firewall %v.", lbRefStr, fwName)
   if err := ignoreNotFound(g.DeleteFirewall(fwName)); err != nil {
    if isForbidden(err) && g.OnXPN() {
     klog.V(4).Infof("DeleteExternalTargetPoolAndChecks(%v): Do not have permission to delete firewall rule %v (on XPN). Raising event.", lbRefStr, fwName)
     g.raiseFirewallChangeNeededEvent(service, FirewallToGCloudDeleteCmd(fwName, g.NetworkProjectID()))
     return nil
    }
    return err
   }
   return nil
  }(); err != nil {
   return err
  }
 }
 return nil
}
func verifyUserRequestedIP(s CloudAddressService, region, requestedIP, fwdRuleIP, lbRef string, desiredNetTier cloud.NetworkTier) (isUserOwnedIP bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if requestedIP == "" {
  return false, nil
 }
 existingAddress, err := s.GetRegionAddressByIP(region, requestedIP)
 if err != nil && !isNotFound(err) {
  klog.Errorf("verifyUserRequestedIP: failed to check whether the requested IP %q for LB %s exists: %v", requestedIP, lbRef, err)
  return false, err
 }
 if err == nil {
  netTierStr, err := s.getNetworkTierFromAddress(existingAddress.Name, region)
  if err != nil {
   return false, fmt.Errorf("failed to check the network tier of the IP %q: %v", requestedIP, err)
  }
  netTier := cloud.NetworkTierGCEValueToType(netTierStr)
  if netTier != desiredNetTier {
   klog.Errorf("verifyUserRequestedIP: requested static IP %q (name: %s) for LB %s has network tier %s, need %s.", requestedIP, existingAddress.Name, lbRef, netTier, desiredNetTier)
   return false, fmt.Errorf("requrested IP %q belongs to the %s network tier; expected %s", requestedIP, netTier, desiredNetTier)
  }
  klog.V(4).Infof("verifyUserRequestedIP: the requested static IP %q (name: %s, tier: %s) for LB %s exists.", requestedIP, existingAddress.Name, netTier, lbRef)
  return true, nil
 }
 if requestedIP == fwdRuleIP {
  klog.V(4).Infof("verifyUserRequestedIP: the requested IP %q is not static, but is currently in use by for LB %s", requestedIP, lbRef)
  return false, nil
 }
 klog.Errorf("verifyUserRequestedIP: requested IP %q for LB %s is neither static nor assigned to the LB", requestedIP, lbRef)
 return false, fmt.Errorf("requested ip %q is neither static nor assigned to the LB", requestedIP)
}
func (g *Cloud) ensureTargetPoolAndHealthCheck(tpExists, tpNeedsRecreation bool, svc *v1.Service, loadBalancerName, clusterID, ipAddressToUse string, hosts []*gceInstance, hcToCreate, hcToDelete *compute.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 serviceName := types.NamespacedName{Namespace: svc.Namespace, Name: svc.Name}
 lbRefStr := fmt.Sprintf("%v(%v)", loadBalancerName, serviceName)
 if tpExists && tpNeedsRecreation {
  var hcNames []string
  if hcToDelete != nil {
   hcNames = append(hcNames, hcToDelete.Name)
  }
  if err := g.DeleteExternalTargetPoolAndChecks(svc, loadBalancerName, g.region, clusterID, hcNames...); err != nil {
   return fmt.Errorf("failed to delete existing target pool for load balancer (%s) update: %v", lbRefStr, err)
  }
  klog.Infof("ensureTargetPoolAndHealthCheck(%s): Deleted target pool.", lbRefStr)
 }
 if tpNeedsRecreation {
  createInstances := hosts
  if len(hosts) > maxTargetPoolCreateInstances {
   createInstances = createInstances[:maxTargetPoolCreateInstances]
  }
  if err := g.createTargetPoolAndHealthCheck(svc, loadBalancerName, serviceName.String(), ipAddressToUse, g.region, clusterID, createInstances, hcToCreate); err != nil {
   return fmt.Errorf("failed to create target pool for load balancer (%s): %v", lbRefStr, err)
  }
  if hcToCreate != nil {
   klog.Infof("ensureTargetPoolAndHealthCheck(%s): Created health checks %v.", lbRefStr, hcToCreate.Name)
  }
  if len(hosts) <= maxTargetPoolCreateInstances {
   klog.Infof("ensureTargetPoolAndHealthCheck(%s): Created target pool.", lbRefStr)
  } else {
   klog.Infof("ensureTargetPoolAndHealthCheck(%s): Created initial target pool (now updating the remaining %d hosts).", lbRefStr, len(hosts)-maxTargetPoolCreateInstances)
   if err := g.updateTargetPool(loadBalancerName, hosts); err != nil {
    return fmt.Errorf("failed to update target pool for load balancer (%s): %v", lbRefStr, err)
   }
   klog.Infof("ensureTargetPoolAndHealthCheck(%s): Updated target pool (with %d hosts).", lbRefStr, len(hosts)-maxTargetPoolCreateInstances)
  }
 } else if tpExists {
  if err := g.updateTargetPool(loadBalancerName, hosts); err != nil {
   return fmt.Errorf("failed to update target pool for load balancer (%s): %v", lbRefStr, err)
  }
  klog.Infof("ensureTargetPoolAndHealthCheck(%s): Updated target pool (with %d hosts).", lbRefStr, len(hosts))
  if hcToCreate != nil {
   if hc, err := g.ensureHTTPHealthCheck(hcToCreate.Name, hcToCreate.RequestPath, int32(hcToCreate.Port)); err != nil || hc == nil {
    return fmt.Errorf("Failed to ensure health check for %v port %d path %v: %v", loadBalancerName, hcToCreate.Port, hcToCreate.RequestPath, err)
   }
  }
 } else {
  klog.Errorf("ensureTargetPoolAndHealthCheck(%s): target pool not exists and doesn't need to be created.", lbRefStr)
 }
 return nil
}
func (g *Cloud) createTargetPoolAndHealthCheck(svc *v1.Service, name, serviceName, ipAddress, region, clusterID string, hosts []*gceInstance, hc *compute.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hcLinks := []string{}
 if hc != nil {
  isNodesHealthCheck := hc.Name != name
  if isNodesHealthCheck {
   g.sharedResourceLock.Lock()
   defer g.sharedResourceLock.Unlock()
  }
  if err := g.ensureHTTPHealthCheckFirewall(svc, serviceName, ipAddress, region, clusterID, hosts, hc.Name, int32(hc.Port), isNodesHealthCheck); err != nil {
   return err
  }
  var err error
  hcRequestPath, hcPort := hc.RequestPath, hc.Port
  if hc, err = g.ensureHTTPHealthCheck(hc.Name, hc.RequestPath, int32(hc.Port)); err != nil || hc == nil {
   return fmt.Errorf("Failed to ensure health check for %v port %d path %v: %v", name, hcPort, hcRequestPath, err)
  }
  hcLinks = append(hcLinks, hc.SelfLink)
 }
 var instances []string
 for _, host := range hosts {
  instances = append(instances, host.makeComparableHostPath())
 }
 klog.Infof("Creating targetpool %v with %d healthchecks", name, len(hcLinks))
 pool := &compute.TargetPool{Name: name, Description: fmt.Sprintf(`{"kubernetes.io/service-name":"%s"}`, serviceName), Instances: instances, SessionAffinity: translateAffinityType(svc.Spec.SessionAffinity), HealthChecks: hcLinks}
 if err := g.CreateTargetPool(pool, region); err != nil && !isHTTPErrorCode(err, http.StatusConflict) {
  return err
 }
 return nil
}
func (g *Cloud) updateTargetPool(loadBalancerName string, hosts []*gceInstance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pool, err := g.GetTargetPool(loadBalancerName, g.region)
 if err != nil {
  return err
 }
 existing := sets.NewString()
 for _, instance := range pool.Instances {
  existing.Insert(hostURLToComparablePath(instance))
 }
 var toAdd []*compute.InstanceReference
 var toRemove []*compute.InstanceReference
 for _, host := range hosts {
  link := host.makeComparableHostPath()
  if !existing.Has(link) {
   toAdd = append(toAdd, &compute.InstanceReference{Instance: link})
  }
  existing.Delete(link)
 }
 for link := range existing {
  toRemove = append(toRemove, &compute.InstanceReference{Instance: link})
 }
 if len(toAdd) > 0 {
  if err := g.AddInstancesToTargetPool(loadBalancerName, g.region, toAdd); err != nil {
   return err
  }
 }
 if len(toRemove) > 0 {
  if err := g.RemoveInstancesFromTargetPool(loadBalancerName, g.region, toRemove); err != nil {
   return err
  }
 }
 updatedPool, err := g.GetTargetPool(loadBalancerName, g.region)
 if err != nil {
  return err
 }
 if len(updatedPool.Instances) != len(hosts) {
  klog.Errorf("Unexpected number of instances (%d) in target pool %s after updating (expected %d). Instances in updated pool: %s", len(updatedPool.Instances), loadBalancerName, len(hosts), strings.Join(updatedPool.Instances, ","))
  return fmt.Errorf("Unexpected number of instances (%d) in target pool %s after update (expected %d)", len(updatedPool.Instances), loadBalancerName, len(hosts))
 }
 return nil
}
func (g *Cloud) targetPoolURL(name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return g.service.BasePath + strings.Join([]string{g.projectID, "regions", g.region, "targetPools", name}, "/")
}
func makeHTTPHealthCheck(name, path string, port int32) *compute.HttpHealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &compute.HttpHealthCheck{Name: name, Port: int64(port), RequestPath: path, Host: "", Description: makeHealthCheckDescription(name), CheckIntervalSec: gceHcCheckIntervalSeconds, TimeoutSec: gceHcTimeoutSeconds, HealthyThreshold: gceHcHealthyThreshold, UnhealthyThreshold: gceHcUnhealthyThreshold}
}
func mergeHTTPHealthChecks(hc, newHC *compute.HttpHealthCheck) *compute.HttpHealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if hc.CheckIntervalSec > newHC.CheckIntervalSec {
  newHC.CheckIntervalSec = hc.CheckIntervalSec
 }
 if hc.TimeoutSec > newHC.TimeoutSec {
  newHC.TimeoutSec = hc.TimeoutSec
 }
 if hc.UnhealthyThreshold > newHC.UnhealthyThreshold {
  newHC.UnhealthyThreshold = hc.UnhealthyThreshold
 }
 if hc.HealthyThreshold > newHC.HealthyThreshold {
  newHC.HealthyThreshold = hc.HealthyThreshold
 }
 return newHC
}
func needToUpdateHTTPHealthChecks(hc, newHC *compute.HttpHealthCheck) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 changed := hc.Port != newHC.Port || hc.RequestPath != newHC.RequestPath || hc.Description != newHC.Description
 changed = changed || hc.CheckIntervalSec < newHC.CheckIntervalSec || hc.TimeoutSec < newHC.TimeoutSec
 changed = changed || hc.UnhealthyThreshold < newHC.UnhealthyThreshold || hc.HealthyThreshold < newHC.HealthyThreshold
 return changed
}
func (g *Cloud) ensureHTTPHealthCheck(name, path string, port int32) (hc *compute.HttpHealthCheck, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newHC := makeHTTPHealthCheck(name, path, port)
 hc, err = g.GetHTTPHealthCheck(name)
 if hc == nil || err != nil && isHTTPErrorCode(err, http.StatusNotFound) {
  klog.Infof("Did not find health check %v, creating port %v path %v", name, port, path)
  if err = g.CreateHTTPHealthCheck(newHC); err != nil {
   return nil, err
  }
  hc, err = g.GetHTTPHealthCheck(name)
  if err != nil {
   klog.Errorf("Failed to get http health check %v", err)
   return nil, err
  }
  klog.Infof("Created HTTP health check %v healthCheckNodePort: %d", name, port)
  return hc, nil
 }
 klog.V(4).Infof("Checking http health check params %s", name)
 if needToUpdateHTTPHealthChecks(hc, newHC) {
  klog.Warningf("Health check %v exists but parameters have drifted - updating...", name)
  newHC = mergeHTTPHealthChecks(hc, newHC)
  if err := g.UpdateHTTPHealthCheck(newHC); err != nil {
   klog.Warningf("Failed to reconcile http health check %v parameters", name)
   return nil, err
  }
  klog.V(4).Infof("Corrected health check %v parameters successful", name)
  hc, err = g.GetHTTPHealthCheck(name)
  if err != nil {
   return nil, err
  }
 }
 return hc, nil
}
func (g *Cloud) forwardingRuleNeedsUpdate(name, region string, loadBalancerIP string, ports []v1.ServicePort) (exists bool, needsUpdate bool, ipAddress string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fwd, err := g.GetRegionForwardingRule(name, region)
 if err != nil {
  if isHTTPErrorCode(err, http.StatusNotFound) {
   return false, true, "", nil
  }
  return true, false, "", fmt.Errorf("error getting load balancer's forwarding rule: %v", err)
 }
 if loadBalancerIP != "" && loadBalancerIP != fwd.IPAddress {
  klog.Infof("LoadBalancer ip for forwarding rule %v was expected to be %v, but was actually %v", fwd.Name, fwd.IPAddress, loadBalancerIP)
  return true, true, fwd.IPAddress, nil
 }
 portRange, err := loadBalancerPortRange(ports)
 if err != nil {
  return true, false, "", err
 }
 if portRange != fwd.PortRange {
  klog.Infof("LoadBalancer port range for forwarding rule %v was expected to be %v, but was actually %v", fwd.Name, fwd.PortRange, portRange)
  return true, true, fwd.IPAddress, nil
 }
 if string(ports[0].Protocol) != fwd.IPProtocol {
  klog.Infof("LoadBalancer protocol for forwarding rule %v was expected to be %v, but was actually %v", fwd.Name, fwd.IPProtocol, string(ports[0].Protocol))
  return true, true, fwd.IPAddress, nil
 }
 return true, false, fwd.IPAddress, nil
}
func (g *Cloud) targetPoolNeedsRecreation(name, region string, affinityType v1.ServiceAffinity) (exists bool, needsRecreation bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tp, err := g.GetTargetPool(name, region)
 if err != nil {
  if isHTTPErrorCode(err, http.StatusNotFound) {
   return false, true, nil
  }
  return true, false, fmt.Errorf("error getting load balancer's target pool: %v", err)
 }
 if tp.SessionAffinity != "" && translateAffinityType(affinityType) != tp.SessionAffinity {
  klog.Infof("LoadBalancer target pool %v changed affinity from %v to %v", name, tp.SessionAffinity, affinityType)
  return true, true, nil
 }
 return true, false, nil
}
func (h *gceInstance) makeComparableHostPath() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("/zones/%s/instances/%s", h.Zone, h.Name)
}
func nodeNames(nodes []*v1.Node) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := make([]string, len(nodes))
 for i, node := range nodes {
  ret[i] = node.Name
 }
 return ret
}
func hostURLToComparablePath(hostURL string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 idx := strings.Index(hostURL, "/zones/")
 if idx < 0 {
  return ""
 }
 return hostURL[idx:]
}
func loadBalancerPortRange(ports []v1.ServicePort) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(ports) == 0 {
  return "", fmt.Errorf("no ports specified for GCE load balancer")
 }
 if ports[0].Protocol != v1.ProtocolTCP && ports[0].Protocol != v1.ProtocolUDP {
  return "", fmt.Errorf("Invalid protocol %s, only TCP and UDP are supported", string(ports[0].Protocol))
 }
 minPort := int32(65536)
 maxPort := int32(0)
 for i := range ports {
  if ports[i].Port < minPort {
   minPort = ports[i].Port
  }
  if ports[i].Port > maxPort {
   maxPort = ports[i].Port
  }
 }
 return fmt.Sprintf("%d-%d", minPort, maxPort), nil
}
func translateAffinityType(affinityType v1.ServiceAffinity) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch affinityType {
 case v1.ServiceAffinityClientIP:
  return gceAffinityTypeClientIP
 case v1.ServiceAffinityNone:
  return gceAffinityTypeNone
 default:
  klog.Errorf("Unexpected affinity type: %v", affinityType)
  return gceAffinityTypeNone
 }
}
func (g *Cloud) firewallNeedsUpdate(name, serviceName, region, ipAddress string, ports []v1.ServicePort, sourceRanges netsets.IPNet) (exists bool, needsUpdate bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fw, err := g.GetFirewall(MakeFirewallName(name))
 if err != nil {
  if isHTTPErrorCode(err, http.StatusNotFound) {
   return false, true, nil
  }
  return false, false, fmt.Errorf("error getting load balancer's firewall: %v", err)
 }
 if fw.Description != makeFirewallDescription(serviceName, ipAddress) {
  return true, true, nil
 }
 if len(fw.Allowed) != 1 || (fw.Allowed[0].IPProtocol != "tcp" && fw.Allowed[0].IPProtocol != "udp") {
  return true, true, nil
 }
 allowedPorts := make([]string, len(ports))
 for ix := range ports {
  allowedPorts[ix] = strconv.Itoa(int(ports[ix].Port))
 }
 if !equalStringSets(allowedPorts, fw.Allowed[0].Ports) {
  return true, true, nil
 }
 actualSourceRanges, err := netsets.ParseIPNets(fw.SourceRanges...)
 if err != nil {
  klog.Warningf("Error parsing firewall SourceRanges: %v", fw.SourceRanges)
  return true, true, nil
 }
 if !sourceRanges.Equal(actualSourceRanges) {
  return true, true, nil
 }
 return true, false, nil
}
func (g *Cloud) ensureHTTPHealthCheckFirewall(svc *v1.Service, serviceName, ipAddress, region, clusterID string, hosts []*gceInstance, hcName string, hcPort int32, isNodesHealthCheck bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 desc := fmt.Sprintf(`{"kubernetes.io/cluster-id":"%s"}`, clusterID)
 if !isNodesHealthCheck {
  desc = makeFirewallDescription(serviceName, ipAddress)
 }
 sourceRanges := lbSrcRngsFlag.ipn
 ports := []v1.ServicePort{{Protocol: "tcp", Port: hcPort}}
 fwName := MakeHealthCheckFirewallName(clusterID, hcName, isNodesHealthCheck)
 fw, err := g.GetFirewall(fwName)
 if err != nil {
  if !isHTTPErrorCode(err, http.StatusNotFound) {
   return fmt.Errorf("error getting firewall for health checks: %v", err)
  }
  klog.Infof("Creating firewall %v for health checks.", fwName)
  if err := g.createFirewall(svc, fwName, region, desc, sourceRanges, ports, hosts); err != nil {
   return err
  }
  klog.Infof("Created firewall %v for health checks.", fwName)
  return nil
 }
 if fw.Description != desc || len(fw.Allowed) != 1 || fw.Allowed[0].IPProtocol != string(ports[0].Protocol) || !equalStringSets(fw.Allowed[0].Ports, []string{strconv.Itoa(int(ports[0].Port))}) || !equalStringSets(fw.SourceRanges, sourceRanges.StringSlice()) {
  klog.Warningf("Firewall %v exists but parameters have drifted - updating...", fwName)
  if err := g.updateFirewall(svc, fwName, region, desc, sourceRanges, ports, hosts); err != nil {
   klog.Warningf("Failed to reconcile firewall %v parameters.", fwName)
   return err
  }
  klog.V(4).Infof("Corrected firewall %v parameters successful", fwName)
 }
 return nil
}
func createForwardingRule(s CloudForwardingRuleService, name, serviceName, region, ipAddress, target string, ports []v1.ServicePort, netTier cloud.NetworkTier) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portRange, err := loadBalancerPortRange(ports)
 if err != nil {
  return err
 }
 desc := makeServiceDescription(serviceName)
 ipProtocol := string(ports[0].Protocol)
 switch netTier {
 case cloud.NetworkTierPremium:
  rule := &compute.ForwardingRule{Name: name, Description: desc, IPAddress: ipAddress, IPProtocol: ipProtocol, PortRange: portRange, Target: target}
  err = s.CreateRegionForwardingRule(rule, region)
 default:
  rule := &computealpha.ForwardingRule{Name: name, Description: desc, IPAddress: ipAddress, IPProtocol: ipProtocol, PortRange: portRange, Target: target, NetworkTier: netTier.ToGCEValue()}
  err = s.CreateAlphaRegionForwardingRule(rule, region)
 }
 if err != nil && !isHTTPErrorCode(err, http.StatusConflict) {
  return err
 }
 return nil
}
func (g *Cloud) createFirewall(svc *v1.Service, name, region, desc string, sourceRanges netsets.IPNet, ports []v1.ServicePort, hosts []*gceInstance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 firewall, err := g.firewallObject(name, region, desc, sourceRanges, ports, hosts)
 if err != nil {
  return err
 }
 if err = g.CreateFirewall(firewall); err != nil {
  if isHTTPErrorCode(err, http.StatusConflict) {
   return nil
  } else if isForbidden(err) && g.OnXPN() {
   klog.V(4).Infof("createFirewall(%v): do not have permission to create firewall rule (on XPN). Raising event.", firewall.Name)
   g.raiseFirewallChangeNeededEvent(svc, FirewallToGCloudCreateCmd(firewall, g.NetworkProjectID()))
   return nil
  }
  return err
 }
 return nil
}
func (g *Cloud) updateFirewall(svc *v1.Service, name, region, desc string, sourceRanges netsets.IPNet, ports []v1.ServicePort, hosts []*gceInstance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 firewall, err := g.firewallObject(name, region, desc, sourceRanges, ports, hosts)
 if err != nil {
  return err
 }
 if err = g.UpdateFirewall(firewall); err != nil {
  if isHTTPErrorCode(err, http.StatusConflict) {
   return nil
  } else if isForbidden(err) && g.OnXPN() {
   klog.V(4).Infof("updateFirewall(%v): do not have permission to update firewall rule (on XPN). Raising event.", firewall.Name)
   g.raiseFirewallChangeNeededEvent(svc, FirewallToGCloudUpdateCmd(firewall, g.NetworkProjectID()))
   return nil
  }
  return err
 }
 return nil
}
func (g *Cloud) firewallObject(name, region, desc string, sourceRanges netsets.IPNet, ports []v1.ServicePort, hosts []*gceInstance) (*compute.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allowedPorts := make([]string, len(ports))
 for ix := range ports {
  allowedPorts[ix] = strconv.Itoa(int(ports[ix].Port))
 }
 hostTags := g.nodeTags
 if len(hostTags) == 0 {
  var err error
  if hostTags, err = g.computeHostTags(hosts); err != nil {
   return nil, fmt.Errorf("no node tags supplied and also failed to parse the given lists of hosts for tags. Abort creating firewall rule")
  }
 }
 firewall := &compute.Firewall{Name: name, Description: desc, Network: g.networkURL, SourceRanges: sourceRanges.StringSlice(), TargetTags: hostTags, Allowed: []*compute.FirewallAllowed{{IPProtocol: strings.ToLower(string(ports[0].Protocol)), Ports: allowedPorts}}}
 return firewall, nil
}
func ensureStaticIP(s CloudAddressService, name, serviceName, region, existingIP string, netTier cloud.NetworkTier) (ipAddress string, existing bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 existed := false
 desc := makeServiceDescription(serviceName)
 var creationErr error
 switch netTier {
 case cloud.NetworkTierPremium:
  addressObj := &compute.Address{Name: name, Description: desc}
  if existingIP != "" {
   addressObj.Address = existingIP
  }
  creationErr = s.ReserveRegionAddress(addressObj, region)
 default:
  addressObj := &computealpha.Address{Name: name, Description: desc, NetworkTier: netTier.ToGCEValue()}
  if existingIP != "" {
   addressObj.Address = existingIP
  }
  creationErr = s.ReserveAlphaRegionAddress(addressObj, region)
 }
 if creationErr != nil {
  if !isHTTPErrorCode(creationErr, http.StatusConflict) && !isHTTPErrorCode(creationErr, http.StatusBadRequest) {
   return "", false, fmt.Errorf("error creating gce static IP address: %v", creationErr)
  }
  existed = true
 }
 addr, err := s.GetRegionAddress(name, region)
 if err != nil {
  return "", false, fmt.Errorf("error getting static IP address: %v", err)
 }
 return addr.Address, existed, nil
}
func (g *Cloud) getServiceNetworkTier(svc *v1.Service) (cloud.NetworkTier, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !g.AlphaFeatureGate.Enabled(AlphaFeatureNetworkTiers) {
  return cloud.NetworkTierDefault, nil
 }
 tier, err := GetServiceNetworkTier(svc)
 if err != nil {
  return cloud.NetworkTier(""), err
 }
 return tier, nil
}
func (g *Cloud) deleteWrongNetworkTieredResources(lbName, lbRef string, desiredNetTier cloud.NetworkTier) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 logPrefix := fmt.Sprintf("deleteWrongNetworkTieredResources:(%s)", lbRef)
 if err := deleteFWDRuleWithWrongTier(g, g.region, lbName, logPrefix, desiredNetTier); err != nil {
  return err
 }
 if err := deleteAddressWithWrongTier(g, g.region, lbName, logPrefix, desiredNetTier); err != nil {
  return err
 }
 return nil
}
func deleteFWDRuleWithWrongTier(s CloudForwardingRuleService, region, name, logPrefix string, desiredNetTier cloud.NetworkTier) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tierStr, err := s.getNetworkTierFromForwardingRule(name, region)
 if isNotFound(err) {
  return nil
 } else if err != nil {
  return err
 }
 existingTier := cloud.NetworkTierGCEValueToType(tierStr)
 if existingTier == desiredNetTier {
  return nil
 }
 klog.V(2).Infof("%s: Network tiers do not match; existing forwarding rule: %q, desired: %q. Deleting the forwarding rule", logPrefix, existingTier, desiredNetTier)
 err = s.DeleteRegionForwardingRule(name, region)
 return ignoreNotFound(err)
}
func deleteAddressWithWrongTier(s CloudAddressService, region, name, logPrefix string, desiredNetTier cloud.NetworkTier) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tierStr, err := s.getNetworkTierFromAddress(name, region)
 if isNotFound(err) {
  return nil
 } else if err != nil {
  return err
 }
 existingTier := cloud.NetworkTierGCEValueToType(tierStr)
 if existingTier == desiredNetTier {
  return nil
 }
 klog.V(2).Infof("%s: Network tiers do not match; existing address: %q, desired: %q. Deleting the address", logPrefix, existingTier, desiredNetTier)
 err = s.DeleteRegionAddress(name, region)
 return ignoreNotFound(err)
}
