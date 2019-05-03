package openstack

import (
 "context"
 "fmt"
 "net"
 "reflect"
 "strings"
 "time"
 "github.com/gophercloud/gophercloud"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/external"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
 v2monitors "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
 v2pools "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
 "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
 neutronports "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
 "github.com/gophercloud/gophercloud/pagination"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/api/v1/service"
)

const (
 loadbalancerActiveInitDelay                    = 1 * time.Second
 loadbalancerActiveFactor                       = 1.2
 loadbalancerActiveSteps                        = 19
 loadbalancerDeleteInitDelay                    = 1 * time.Second
 loadbalancerDeleteFactor                       = 1.2
 loadbalancerDeleteSteps                        = 13
 activeStatus                                   = "ACTIVE"
 errorStatus                                    = "ERROR"
 ServiceAnnotationLoadBalancerFloatingNetworkID = "loadbalancer.openstack.org/floating-network-id"
 ServiceAnnotationLoadBalancerSubnetID          = "loadbalancer.openstack.org/subnet-id"
 ServiceAnnotationLoadBalancerInternal          = "service.beta.kubernetes.io/openstack-internal-load-balancer"
)

type LbaasV2 struct{ LoadBalancer }
type empty struct{}

func networkExtensions(client *gophercloud.ServiceClient) (map[string]bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 seen := make(map[string]bool)
 pager := extensions.List(client)
 err := pager.EachPage(func(page pagination.Page) (bool, error) {
  exts, err := extensions.ExtractExtensions(page)
  if err != nil {
   return false, err
  }
  for _, ext := range exts {
   seen[ext.Alias] = true
  }
  return true, nil
 })
 return seen, err
}
func getFloatingIPByPortID(client *gophercloud.ServiceClient, portID string) (*floatingips.FloatingIP, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 opts := floatingips.ListOpts{PortID: portID}
 pager := floatingips.List(client, opts)
 floatingIPList := make([]floatingips.FloatingIP, 0, 1)
 err := pager.EachPage(func(page pagination.Page) (bool, error) {
  f, err := floatingips.ExtractFloatingIPs(page)
  if err != nil {
   return false, err
  }
  floatingIPList = append(floatingIPList, f...)
  if len(floatingIPList) > 1 {
   return false, ErrMultipleResults
  }
  return true, nil
 })
 if err != nil {
  if isNotFound(err) {
   return nil, ErrNotFound
  }
  return nil, err
 }
 if len(floatingIPList) == 0 {
  return nil, ErrNotFound
 } else if len(floatingIPList) > 1 {
  return nil, ErrMultipleResults
 }
 return &floatingIPList[0], nil
}
func getLoadbalancerByName(client *gophercloud.ServiceClient, name string) (*loadbalancers.LoadBalancer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 opts := loadbalancers.ListOpts{Name: name}
 pager := loadbalancers.List(client, opts)
 loadbalancerList := make([]loadbalancers.LoadBalancer, 0, 1)
 err := pager.EachPage(func(page pagination.Page) (bool, error) {
  v, err := loadbalancers.ExtractLoadBalancers(page)
  if err != nil {
   return false, err
  }
  loadbalancerList = append(loadbalancerList, v...)
  if len(loadbalancerList) > 1 {
   return false, ErrMultipleResults
  }
  return true, nil
 })
 if err != nil {
  if isNotFound(err) {
   return nil, ErrNotFound
  }
  return nil, err
 }
 if len(loadbalancerList) == 0 {
  return nil, ErrNotFound
 } else if len(loadbalancerList) > 1 {
  return nil, ErrMultipleResults
 }
 return &loadbalancerList[0], nil
}
func getListenersByLoadBalancerID(client *gophercloud.ServiceClient, id string) ([]listeners.Listener, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var existingListeners []listeners.Listener
 err := listeners.List(client, listeners.ListOpts{LoadbalancerID: id}).EachPage(func(page pagination.Page) (bool, error) {
  listenerList, err := listeners.ExtractListeners(page)
  if err != nil {
   return false, err
  }
  for _, l := range listenerList {
   for _, lb := range l.Loadbalancers {
    if lb.ID == id {
     existingListeners = append(existingListeners, l)
     break
    }
   }
  }
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return existingListeners, nil
}
func getListenerForPort(existingListeners []listeners.Listener, port v1.ServicePort) *listeners.Listener {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, l := range existingListeners {
  if listeners.Protocol(l.Protocol) == toListenersProtocol(port.Protocol) && l.ProtocolPort == int(port.Port) {
   return &l
  }
 }
 return nil
}
func getPoolByListenerID(client *gophercloud.ServiceClient, loadbalancerID string, listenerID string) (*v2pools.Pool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 listenerPools := make([]v2pools.Pool, 0, 1)
 err := v2pools.List(client, v2pools.ListOpts{LoadbalancerID: loadbalancerID}).EachPage(func(page pagination.Page) (bool, error) {
  poolsList, err := v2pools.ExtractPools(page)
  if err != nil {
   return false, err
  }
  for _, p := range poolsList {
   for _, l := range p.Listeners {
    if l.ID == listenerID {
     listenerPools = append(listenerPools, p)
    }
   }
  }
  if len(listenerPools) > 1 {
   return false, ErrMultipleResults
  }
  return true, nil
 })
 if err != nil {
  if isNotFound(err) {
   return nil, ErrNotFound
  }
  return nil, err
 }
 if len(listenerPools) == 0 {
  return nil, ErrNotFound
 } else if len(listenerPools) > 1 {
  return nil, ErrMultipleResults
 }
 return &listenerPools[0], nil
}
func getMembersByPoolID(client *gophercloud.ServiceClient, id string) ([]v2pools.Member, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var members []v2pools.Member
 err := v2pools.ListMembers(client, id, v2pools.ListMembersOpts{}).EachPage(func(page pagination.Page) (bool, error) {
  membersList, err := v2pools.ExtractMembers(page)
  if err != nil {
   return false, err
  }
  members = append(members, membersList...)
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return members, nil
}
func memberExists(members []v2pools.Member, addr string, port int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, member := range members {
  if member.Address == addr && member.ProtocolPort == port {
   return true
  }
 }
 return false
}
func popListener(existingListeners []listeners.Listener, id string) []listeners.Listener {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i, existingListener := range existingListeners {
  if existingListener.ID == id {
   existingListeners[i] = existingListeners[len(existingListeners)-1]
   existingListeners = existingListeners[:len(existingListeners)-1]
   break
  }
 }
 return existingListeners
}
func popMember(members []v2pools.Member, addr string, port int) []v2pools.Member {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i, member := range members {
  if member.Address == addr && member.ProtocolPort == port {
   members[i] = members[len(members)-1]
   members = members[:len(members)-1]
  }
 }
 return members
}
func getSecurityGroupName(service *v1.Service) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 securityGroupName := fmt.Sprintf("lb-sg-%s-%s-%s", service.UID, service.Namespace, service.Name)
 if len(securityGroupName) > 255 {
  securityGroupName = securityGroupName[:255]
 }
 return securityGroupName
}
func getSecurityGroupRules(client *gophercloud.ServiceClient, opts rules.ListOpts) ([]rules.SecGroupRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pager := rules.List(client, opts)
 var securityRules []rules.SecGroupRule
 err := pager.EachPage(func(page pagination.Page) (bool, error) {
  ruleList, err := rules.ExtractRules(page)
  if err != nil {
   return false, err
  }
  securityRules = append(securityRules, ruleList...)
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return securityRules, nil
}
func waitLoadbalancerActiveProvisioningStatus(client *gophercloud.ServiceClient, loadbalancerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 backoff := wait.Backoff{Duration: loadbalancerActiveInitDelay, Factor: loadbalancerActiveFactor, Steps: loadbalancerActiveSteps}
 var provisioningStatus string
 err := wait.ExponentialBackoff(backoff, func() (bool, error) {
  loadbalancer, err := loadbalancers.Get(client, loadbalancerID).Extract()
  if err != nil {
   return false, err
  }
  provisioningStatus = loadbalancer.ProvisioningStatus
  if loadbalancer.ProvisioningStatus == activeStatus {
   return true, nil
  } else if loadbalancer.ProvisioningStatus == errorStatus {
   return true, fmt.Errorf("loadbalancer has gone into ERROR state")
  } else {
   return false, nil
  }
 })
 if err == wait.ErrWaitTimeout {
  err = fmt.Errorf("loadbalancer failed to go into ACTIVE provisioning status within alloted time")
 }
 return provisioningStatus, err
}
func waitLoadbalancerDeleted(client *gophercloud.ServiceClient, loadbalancerID string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 backoff := wait.Backoff{Duration: loadbalancerDeleteInitDelay, Factor: loadbalancerDeleteFactor, Steps: loadbalancerDeleteSteps}
 err := wait.ExponentialBackoff(backoff, func() (bool, error) {
  _, err := loadbalancers.Get(client, loadbalancerID).Extract()
  if err != nil {
   if isNotFound(err) {
    return true, nil
   }
   return false, err
  }
  return false, nil
 })
 if err == wait.ErrWaitTimeout {
  err = fmt.Errorf("loadbalancer failed to delete within the alloted time")
 }
 return err
}
func toRuleProtocol(protocol v1.Protocol) rules.RuleProtocol {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch protocol {
 case v1.ProtocolTCP:
  return rules.ProtocolTCP
 case v1.ProtocolUDP:
  return rules.ProtocolUDP
 default:
  return rules.RuleProtocol(strings.ToLower(string(protocol)))
 }
}
func toListenersProtocol(protocol v1.Protocol) listeners.Protocol {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch protocol {
 case v1.ProtocolTCP:
  return listeners.ProtocolTCP
 default:
  return listeners.Protocol(string(protocol))
 }
}
func createNodeSecurityGroup(client *gophercloud.ServiceClient, nodeSecurityGroupID string, port int, protocol v1.Protocol, lbSecGroup string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 v4NodeSecGroupRuleCreateOpts := rules.CreateOpts{Direction: rules.DirIngress, PortRangeMax: port, PortRangeMin: port, Protocol: toRuleProtocol(protocol), RemoteGroupID: lbSecGroup, SecGroupID: nodeSecurityGroupID, EtherType: rules.EtherType4}
 v6NodeSecGroupRuleCreateOpts := rules.CreateOpts{Direction: rules.DirIngress, PortRangeMax: port, PortRangeMin: port, Protocol: toRuleProtocol(protocol), RemoteGroupID: lbSecGroup, SecGroupID: nodeSecurityGroupID, EtherType: rules.EtherType6}
 _, err := rules.Create(client, v4NodeSecGroupRuleCreateOpts).Extract()
 if err != nil {
  return err
 }
 _, err = rules.Create(client, v6NodeSecGroupRuleCreateOpts).Extract()
 if err != nil {
  return err
 }
 return nil
}
func (lbaas *LbaasV2) createLoadBalancer(service *v1.Service, name string, internalAnnotation bool) (*loadbalancers.LoadBalancer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 createOpts := loadbalancers.CreateOpts{Name: name, Description: fmt.Sprintf("Kubernetes external service %s", name), VipSubnetID: lbaas.opts.SubnetID, Provider: lbaas.opts.LBProvider}
 loadBalancerIP := service.Spec.LoadBalancerIP
 if loadBalancerIP != "" && internalAnnotation {
  createOpts.VipAddress = loadBalancerIP
 }
 loadbalancer, err := loadbalancers.Create(lbaas.lb, createOpts).Extract()
 if err != nil {
  return nil, fmt.Errorf("error creating loadbalancer %v: %v", createOpts, err)
 }
 return loadbalancer, nil
}
func (lbaas *LbaasV2) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := lbaas.GetLoadBalancerName(ctx, clusterName, service)
 loadbalancer, err := getLoadbalancerByName(lbaas.lb, loadBalancerName)
 if err == ErrNotFound {
  return nil, false, nil
 }
 if loadbalancer == nil {
  return nil, false, err
 }
 status := &v1.LoadBalancerStatus{}
 portID := loadbalancer.VipPortID
 if portID != "" {
  floatIP, err := getFloatingIPByPortID(lbaas.network, portID)
  if err != nil && err != ErrNotFound {
   return nil, false, fmt.Errorf("error getting floating ip for port %s: %v", portID, err)
  }
  if floatIP != nil {
   status.Ingress = []v1.LoadBalancerIngress{{IP: floatIP.FloatingIP}}
  }
 } else {
  status.Ingress = []v1.LoadBalancerIngress{{IP: loadbalancer.VipAddress}}
 }
 return status, true, err
}
func (lbaas *LbaasV2) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.DefaultLoadBalancerName(service)
}
func nodeAddressForLB(node *v1.Node) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 addrs := node.Status.Addresses
 if len(addrs) == 0 {
  return "", ErrNoAddressFound
 }
 allowedAddrTypes := []v1.NodeAddressType{v1.NodeInternalIP, v1.NodeExternalIP}
 for _, allowedAddrType := range allowedAddrTypes {
  for _, addr := range addrs {
   if addr.Type == allowedAddrType {
    return addr.Address, nil
   }
  }
 }
 return "", ErrNoAddressFound
}
func getStringFromServiceAnnotation(service *v1.Service, annotationKey string, defaultSetting string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("getStringFromServiceAnnotation(%v, %v, %v)", service, annotationKey, defaultSetting)
 if annotationValue, ok := service.Annotations[annotationKey]; ok {
  klog.V(4).Infof("Found a Service Annotation: %v = %v", annotationKey, annotationValue)
  return annotationValue
 }
 klog.V(4).Infof("Could not find a Service Annotation; falling back on cloud-config setting: %v = %v", annotationKey, defaultSetting)
 return defaultSetting
}
func getSubnetIDForLB(compute *gophercloud.ServiceClient, node v1.Node) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ipAddress, err := nodeAddressForLB(&node)
 if err != nil {
  return "", err
 }
 instanceID := node.Spec.ProviderID
 if ind := strings.LastIndex(instanceID, "/"); ind >= 0 {
  instanceID = instanceID[(ind + 1):]
 }
 interfaces, err := getAttachedInterfacesByID(compute, instanceID)
 if err != nil {
  return "", err
 }
 for _, intf := range interfaces {
  for _, fixedIP := range intf.FixedIPs {
   if fixedIP.IPAddress == ipAddress {
    return fixedIP.SubnetID, nil
   }
  }
 }
 return "", ErrNotFound
}
func getNodeSecurityGroupIDForLB(compute *gophercloud.ServiceClient, network *gophercloud.ServiceClient, nodes []*v1.Node) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 secGroupNames := sets.NewString()
 for _, node := range nodes {
  nodeName := types.NodeName(node.Name)
  srv, err := getServerByName(compute, nodeName)
  if err != nil {
   return []string{}, err
  }
  secGroupNames.Insert(srv.SecurityGroups[0]["name"].(string))
 }
 secGroupIDs := make([]string, secGroupNames.Len())
 for i, name := range secGroupNames.List() {
  secGroupID, err := groups.IDFromName(network, name)
  if err != nil {
   return []string{}, err
  }
  secGroupIDs[i] = secGroupID
 }
 return secGroupIDs, nil
}
func isSecurityGroupNotFound(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errType := reflect.TypeOf(err).String()
 errTypeSlice := strings.Split(errType, ".")
 errTypeValue := ""
 if len(errTypeSlice) != 0 {
  errTypeValue = errTypeSlice[len(errTypeSlice)-1]
 }
 if errTypeValue == "ErrResourceNotFound" {
  return true
 }
 return false
}
func getFloatingNetworkIDForLB(client *gophercloud.ServiceClient) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var floatingNetworkIds []string
 type NetworkWithExternalExt struct {
  networks.Network
  external.NetworkExternalExt
 }
 err := networks.List(client, networks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
  var externalNetwork []NetworkWithExternalExt
  err := networks.ExtractNetworksInto(page, &externalNetwork)
  if err != nil {
   return false, err
  }
  for _, externalNet := range externalNetwork {
   if externalNet.External {
    floatingNetworkIds = append(floatingNetworkIds, externalNet.ID)
   }
  }
  if len(floatingNetworkIds) > 1 {
   return false, ErrMultipleResults
  }
  return true, nil
 })
 if err != nil {
  if isNotFound(err) {
   return "", ErrNotFound
  }
  if err == ErrMultipleResults {
   klog.V(4).Infof("find multiple external networks, pick the first one when there are no explicit configuration.")
   return floatingNetworkIds[0], nil
  }
  return "", err
 }
 if len(floatingNetworkIds) == 0 {
  return "", ErrNotFound
 }
 return floatingNetworkIds[0], nil
}
func (lbaas *LbaasV2) EnsureLoadBalancer(ctx context.Context, clusterName string, apiService *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v, %v, %v)", clusterName, apiService.Namespace, apiService.Name, apiService.Spec.LoadBalancerIP, apiService.Spec.Ports, nodes, apiService.Annotations)
 if len(nodes) == 0 {
  return nil, fmt.Errorf("there are no available nodes for LoadBalancer service %s/%s", apiService.Namespace, apiService.Name)
 }
 lbaas.opts.SubnetID = getStringFromServiceAnnotation(apiService, ServiceAnnotationLoadBalancerSubnetID, lbaas.opts.SubnetID)
 if len(lbaas.opts.SubnetID) == 0 {
  subnetID, err := getSubnetIDForLB(lbaas.compute, *nodes[0])
  if err != nil {
   klog.Warningf("Failed to find subnet-id for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
   return nil, fmt.Errorf("no subnet-id for service %s/%s : subnet-id not set in cloud provider config, "+"and failed to find subnet-id from OpenStack: %v", apiService.Namespace, apiService.Name, err)
  }
  lbaas.opts.SubnetID = subnetID
 }
 ports := apiService.Spec.Ports
 if len(ports) == 0 {
  return nil, fmt.Errorf("no ports provided to openstack load balancer")
 }
 floatingPool := getStringFromServiceAnnotation(apiService, ServiceAnnotationLoadBalancerFloatingNetworkID, lbaas.opts.FloatingNetworkID)
 if len(floatingPool) == 0 {
  var err error
  floatingPool, err = getFloatingNetworkIDForLB(lbaas.network)
  if err != nil {
   klog.Warningf("Failed to find floating-network-id for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
  }
 }
 var internalAnnotation bool
 internal := getStringFromServiceAnnotation(apiService, ServiceAnnotationLoadBalancerInternal, "false")
 switch internal {
 case "true":
  klog.V(4).Infof("Ensure an internal loadbalancer service.")
  internalAnnotation = true
 case "false":
  if len(floatingPool) != 0 {
   klog.V(4).Infof("Ensure an external loadbalancer service, using floatingPool: %v", floatingPool)
   internalAnnotation = false
  } else {
   return nil, fmt.Errorf("floating-network-id or loadbalancer.openstack.org/floating-network-id should be specified when ensuring an external loadbalancer service")
  }
 default:
  return nil, fmt.Errorf("unknown service.beta.kubernetes.io/openstack-internal-load-balancer annotation: %v, specify \"true\" or \"false\" ", internal)
 }
 for _, port := range ports {
  if port.Protocol != v1.ProtocolTCP {
   return nil, fmt.Errorf("only TCP LoadBalancer is supported for openstack load balancers")
  }
 }
 sourceRanges, err := service.GetLoadBalancerSourceRanges(apiService)
 if err != nil {
  return nil, fmt.Errorf("failed to get source ranges for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
 }
 if !service.IsAllowAll(sourceRanges) && !lbaas.opts.ManageSecurityGroups {
  return nil, fmt.Errorf("source range restrictions are not supported for openstack load balancers without managing security groups")
 }
 affinity := apiService.Spec.SessionAffinity
 var persistence *v2pools.SessionPersistence
 switch affinity {
 case v1.ServiceAffinityNone:
  persistence = nil
 case v1.ServiceAffinityClientIP:
  persistence = &v2pools.SessionPersistence{Type: "SOURCE_IP"}
 default:
  return nil, fmt.Errorf("unsupported load balancer affinity: %v", affinity)
 }
 name := lbaas.GetLoadBalancerName(ctx, clusterName, apiService)
 loadbalancer, err := getLoadbalancerByName(lbaas.lb, name)
 if err != nil {
  if err != ErrNotFound {
   return nil, fmt.Errorf("error getting loadbalancer %s: %v", name, err)
  }
  klog.V(2).Infof("Creating loadbalancer %s", name)
  loadbalancer, err = lbaas.createLoadBalancer(apiService, name, internalAnnotation)
  if err != nil {
   return nil, fmt.Errorf("error creating loadbalancer %s: %v", name, err)
  }
 } else {
  klog.V(2).Infof("LoadBalancer %s already exists", name)
 }
 provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
 if err != nil {
  return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
 }
 lbmethod := v2pools.LBMethod(lbaas.opts.LBMethod)
 if lbmethod == "" {
  lbmethod = v2pools.LBMethodRoundRobin
 }
 oldListeners, err := getListenersByLoadBalancerID(lbaas.lb, loadbalancer.ID)
 if err != nil {
  return nil, fmt.Errorf("error getting LB %s listeners: %v", name, err)
 }
 for portIndex, port := range ports {
  listener := getListenerForPort(oldListeners, port)
  if listener == nil {
   klog.V(4).Infof("Creating listener for port %d", int(port.Port))
   listener, err = listeners.Create(lbaas.lb, listeners.CreateOpts{Name: fmt.Sprintf("listener_%s_%d", name, portIndex), Protocol: listeners.Protocol(port.Protocol), ProtocolPort: int(port.Port), LoadbalancerID: loadbalancer.ID}).Extract()
   if err != nil {
    return nil, fmt.Errorf("error creating LB listener: %v", err)
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  klog.V(4).Infof("Listener for %s port %d: %s", string(port.Protocol), int(port.Port), listener.ID)
  oldListeners = popListener(oldListeners, listener.ID)
  pool, err := getPoolByListenerID(lbaas.lb, loadbalancer.ID, listener.ID)
  if err != nil && err != ErrNotFound {
   return nil, fmt.Errorf("error getting pool for listener %s: %v", listener.ID, err)
  }
  if pool == nil {
   klog.V(4).Infof("Creating pool for listener %s", listener.ID)
   pool, err = v2pools.Create(lbaas.lb, v2pools.CreateOpts{Name: fmt.Sprintf("pool_%s_%d", name, portIndex), Protocol: v2pools.Protocol(port.Protocol), LBMethod: lbmethod, ListenerID: listener.ID, Persistence: persistence}).Extract()
   if err != nil {
    return nil, fmt.Errorf("error creating pool for listener %s: %v", listener.ID, err)
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  klog.V(4).Infof("Pool for listener %s: %s", listener.ID, pool.ID)
  members, err := getMembersByPoolID(lbaas.lb, pool.ID)
  if err != nil && !isNotFound(err) {
   return nil, fmt.Errorf("error getting pool members %s: %v", pool.ID, err)
  }
  for _, node := range nodes {
   addr, err := nodeAddressForLB(node)
   if err != nil {
    if err == ErrNotFound {
     klog.Warningf("Failed to create LB pool member for node %s: %v", node.Name, err)
     continue
    } else {
     return nil, fmt.Errorf("error getting address for node %s: %v", node.Name, err)
    }
   }
   if !memberExists(members, addr, int(port.NodePort)) {
    klog.V(4).Infof("Creating member for pool %s", pool.ID)
    _, err := v2pools.CreateMember(lbaas.lb, pool.ID, v2pools.CreateMemberOpts{Name: fmt.Sprintf("member_%s_%d_%s", name, portIndex, node.Name), ProtocolPort: int(port.NodePort), Address: addr, SubnetID: lbaas.opts.SubnetID}).Extract()
    if err != nil {
     return nil, fmt.Errorf("error creating LB pool member for node: %s, %v", node.Name, err)
    }
    provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
    if err != nil {
     return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
    }
   } else {
    members = popMember(members, addr, int(port.NodePort))
   }
   klog.V(4).Infof("Ensured pool %s has member for %s at %s", pool.ID, node.Name, addr)
  }
  for _, member := range members {
   klog.V(4).Infof("Deleting obsolete member %s for pool %s address %s", member.ID, pool.ID, member.Address)
   err := v2pools.DeleteMember(lbaas.lb, pool.ID, member.ID).ExtractErr()
   if err != nil && !isNotFound(err) {
    return nil, fmt.Errorf("error deleting obsolete member %s for pool %s address %s: %v", member.ID, pool.ID, member.Address, err)
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  monitorID := pool.MonitorID
  if monitorID == "" && lbaas.opts.CreateMonitor {
   klog.V(4).Infof("Creating monitor for pool %s", pool.ID)
   monitor, err := v2monitors.Create(lbaas.lb, v2monitors.CreateOpts{Name: fmt.Sprintf("monitor_%s_%d", name, portIndex), PoolID: pool.ID, Type: string(port.Protocol), Delay: int(lbaas.opts.MonitorDelay.Duration.Seconds()), Timeout: int(lbaas.opts.MonitorTimeout.Duration.Seconds()), MaxRetries: int(lbaas.opts.MonitorMaxRetries)}).Extract()
   if err != nil {
    return nil, fmt.Errorf("error creating LB pool healthmonitor: %v", err)
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
   monitorID = monitor.ID
  } else if lbaas.opts.CreateMonitor == false {
   klog.V(4).Infof("Do not create monitor for pool %s when create-monitor is false", pool.ID)
  }
  if monitorID != "" {
   klog.V(4).Infof("Monitor for pool %s: %s", pool.ID, monitorID)
  }
 }
 for _, listener := range oldListeners {
  klog.V(4).Infof("Deleting obsolete listener %s:", listener.ID)
  pool, err := getPoolByListenerID(lbaas.lb, loadbalancer.ID, listener.ID)
  if err != nil && err != ErrNotFound {
   return nil, fmt.Errorf("error getting pool for obsolete listener %s: %v", listener.ID, err)
  }
  if pool != nil {
   monitorID := pool.MonitorID
   if monitorID != "" {
    klog.V(4).Infof("Deleting obsolete monitor %s for pool %s", monitorID, pool.ID)
    err = v2monitors.Delete(lbaas.lb, monitorID).ExtractErr()
    if err != nil && !isNotFound(err) {
     return nil, fmt.Errorf("error deleting obsolete monitor %s for pool %s: %v", monitorID, pool.ID, err)
    }
    provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
    if err != nil {
     return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
    }
   }
   members, err := getMembersByPoolID(lbaas.lb, pool.ID)
   if err != nil && !isNotFound(err) {
    return nil, fmt.Errorf("error getting members for pool %s: %v", pool.ID, err)
   }
   if members != nil {
    for _, member := range members {
     klog.V(4).Infof("Deleting obsolete member %s for pool %s address %s", member.ID, pool.ID, member.Address)
     err := v2pools.DeleteMember(lbaas.lb, pool.ID, member.ID).ExtractErr()
     if err != nil && !isNotFound(err) {
      return nil, fmt.Errorf("error deleting obsolete member %s for pool %s address %s: %v", member.ID, pool.ID, member.Address, err)
     }
     provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
     if err != nil {
      return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
     }
    }
   }
   klog.V(4).Infof("Deleting obsolete pool %s for listener %s", pool.ID, listener.ID)
   err = v2pools.Delete(lbaas.lb, pool.ID).ExtractErr()
   if err != nil && !isNotFound(err) {
    return nil, fmt.Errorf("error deleting obsolete pool %s for listener %s: %v", pool.ID, listener.ID, err)
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  err = listeners.Delete(lbaas.lb, listener.ID).ExtractErr()
  if err != nil && !isNotFound(err) {
   return nil, fmt.Errorf("error deleteting obsolete listener: %v", err)
  }
  provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
  if err != nil {
   return nil, fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
  }
  klog.V(2).Infof("Deleted obsolete listener: %s", listener.ID)
 }
 portID := loadbalancer.VipPortID
 floatIP, err := getFloatingIPByPortID(lbaas.network, portID)
 if err != nil && err != ErrNotFound {
  return nil, fmt.Errorf("error getting floating ip for port %s: %v", portID, err)
 }
 if floatIP == nil && floatingPool != "" && !internalAnnotation {
  klog.V(4).Infof("Creating floating ip for loadbalancer %s port %s", loadbalancer.ID, portID)
  floatIPOpts := floatingips.CreateOpts{FloatingNetworkID: floatingPool, PortID: portID}
  loadBalancerIP := apiService.Spec.LoadBalancerIP
  if loadBalancerIP != "" {
   floatIPOpts.FloatingIP = loadBalancerIP
  }
  floatIP, err = floatingips.Create(lbaas.network, floatIPOpts).Extract()
  if err != nil {
   return nil, fmt.Errorf("error creating LB floatingip %+v: %v", floatIPOpts, err)
  }
 }
 status := &v1.LoadBalancerStatus{}
 if floatIP != nil {
  status.Ingress = []v1.LoadBalancerIngress{{IP: floatIP.FloatingIP}}
 } else {
  status.Ingress = []v1.LoadBalancerIngress{{IP: loadbalancer.VipAddress}}
 }
 if lbaas.opts.ManageSecurityGroups {
  err := lbaas.ensureSecurityGroup(clusterName, apiService, nodes, loadbalancer)
  if err != nil {
   _ = lbaas.EnsureLoadBalancerDeleted(ctx, clusterName, apiService)
   return status, err
  }
 }
 return status, nil
}
func (lbaas *LbaasV2) ensureSecurityGroup(clusterName string, apiService *v1.Service, nodes []*v1.Node, loadbalancer *loadbalancers.LoadBalancer) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 if len(lbaas.opts.NodeSecurityGroupIDs) == 0 {
  lbaas.opts.NodeSecurityGroupIDs, err = getNodeSecurityGroupIDForLB(lbaas.compute, lbaas.network, nodes)
  if err != nil {
   return fmt.Errorf("failed to find node-security-group for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
  }
 }
 klog.V(4).Infof("find node-security-group %v for loadbalancer service %s/%s", lbaas.opts.NodeSecurityGroupIDs, apiService.Namespace, apiService.Name)
 ports := apiService.Spec.Ports
 if len(ports) == 0 {
  return fmt.Errorf("no ports provided to openstack load balancer")
 }
 sourceRanges, err := service.GetLoadBalancerSourceRanges(apiService)
 if err != nil {
  return fmt.Errorf("failed to get source ranges for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
 }
 lbSecGroupName := getSecurityGroupName(apiService)
 lbSecGroupID, err := groups.IDFromName(lbaas.network, lbSecGroupName)
 if err != nil {
  if isSecurityGroupNotFound(err) {
   lbSecGroupID = ""
  } else {
   return fmt.Errorf("error occurred finding security group: %s: %v", lbSecGroupName, err)
  }
 }
 if len(lbSecGroupID) == 0 {
  lbSecGroupCreateOpts := groups.CreateOpts{Name: getSecurityGroupName(apiService), Description: fmt.Sprintf("Security Group for %s/%s Service LoadBalancer in cluster %s", apiService.Namespace, apiService.Name, clusterName)}
  lbSecGroup, err := groups.Create(lbaas.network, lbSecGroupCreateOpts).Extract()
  if err != nil {
   return fmt.Errorf("failed to create Security Group for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
  }
  lbSecGroupID = lbSecGroup.ID
  for _, port := range ports {
   for _, sourceRange := range sourceRanges.StringSlice() {
    ethertype := rules.EtherType4
    network, _, err := net.ParseCIDR(sourceRange)
    if err != nil {
     return fmt.Errorf("error parsing source range %s as a CIDR: %v", sourceRange, err)
    }
    if network.To4() == nil {
     ethertype = rules.EtherType6
    }
    lbSecGroupRuleCreateOpts := rules.CreateOpts{Direction: rules.DirIngress, PortRangeMax: int(port.Port), PortRangeMin: int(port.Port), Protocol: toRuleProtocol(port.Protocol), RemoteIPPrefix: sourceRange, SecGroupID: lbSecGroup.ID, EtherType: ethertype}
    _, err = rules.Create(lbaas.network, lbSecGroupRuleCreateOpts).Extract()
    if err != nil {
     return fmt.Errorf("error occurred creating rule for SecGroup %s: %v", lbSecGroup.ID, err)
    }
   }
  }
  lbSecGroupRuleCreateOpts := rules.CreateOpts{Direction: rules.DirIngress, PortRangeMax: 4, PortRangeMin: 3, Protocol: rules.ProtocolICMP, RemoteIPPrefix: "0.0.0.0/0", SecGroupID: lbSecGroup.ID, EtherType: rules.EtherType4}
  _, err = rules.Create(lbaas.network, lbSecGroupRuleCreateOpts).Extract()
  if err != nil {
   return fmt.Errorf("error occurred creating rule for SecGroup %s: %v", lbSecGroup.ID, err)
  }
  lbSecGroupRuleCreateOpts = rules.CreateOpts{Direction: rules.DirIngress, PortRangeMax: 0, PortRangeMin: 2, Protocol: rules.ProtocolICMP, RemoteIPPrefix: "::/0", SecGroupID: lbSecGroup.ID, EtherType: rules.EtherType6}
  _, err = rules.Create(lbaas.network, lbSecGroupRuleCreateOpts).Extract()
  if err != nil {
   return fmt.Errorf("error occurred creating rule for SecGroup %s: %v", lbSecGroup.ID, err)
  }
  portID := loadbalancer.VipPortID
  port, err := getPortByID(lbaas.network, portID)
  if err != nil {
   return err
  }
  found := false
  for _, portSecurityGroups := range port.SecurityGroups {
   if portSecurityGroups == lbSecGroup.ID {
    found = true
    break
   }
  }
  if !found {
   port.SecurityGroups = append(port.SecurityGroups, lbSecGroup.ID)
   updateOpts := neutronports.UpdateOpts{SecurityGroups: &port.SecurityGroups}
   res := neutronports.Update(lbaas.network, portID, updateOpts)
   if res.Err != nil {
    msg := fmt.Sprintf("Error occurred updating port %s for loadbalancer service %s/%s: %v", portID, apiService.Namespace, apiService.Name, res.Err)
    return fmt.Errorf(msg)
   }
  }
 }
 for _, port := range ports {
  for _, nodeSecurityGroupID := range lbaas.opts.NodeSecurityGroupIDs {
   opts := rules.ListOpts{Direction: string(rules.DirIngress), SecGroupID: nodeSecurityGroupID, RemoteGroupID: lbSecGroupID, PortRangeMax: int(port.NodePort), PortRangeMin: int(port.NodePort), Protocol: string(port.Protocol)}
   secGroupRules, err := getSecurityGroupRules(lbaas.network, opts)
   if err != nil && !isNotFound(err) {
    msg := fmt.Sprintf("Error finding rules for remote group id %s in security group id %s: %v", lbSecGroupID, nodeSecurityGroupID, err)
    return fmt.Errorf(msg)
   }
   if len(secGroupRules) != 0 {
    continue
   }
   err = createNodeSecurityGroup(lbaas.network, nodeSecurityGroupID, int(port.NodePort), port.Protocol, lbSecGroupID)
   if err != nil {
    return fmt.Errorf("error occurred creating security group for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
   }
  }
 }
 return nil
}
func (lbaas *LbaasV2) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := lbaas.GetLoadBalancerName(ctx, clusterName, service)
 klog.V(4).Infof("UpdateLoadBalancer(%v, %v, %v)", clusterName, loadBalancerName, nodes)
 lbaas.opts.SubnetID = getStringFromServiceAnnotation(service, ServiceAnnotationLoadBalancerSubnetID, lbaas.opts.SubnetID)
 if len(lbaas.opts.SubnetID) == 0 && len(nodes) > 0 {
  subnetID, err := getSubnetIDForLB(lbaas.compute, *nodes[0])
  if err != nil {
   klog.Warningf("Failed to find subnet-id for loadbalancer service %s/%s: %v", service.Namespace, service.Name, err)
   return fmt.Errorf("no subnet-id for service %s/%s : subnet-id not set in cloud provider config, "+"and failed to find subnet-id from OpenStack: %v", service.Namespace, service.Name, err)
  }
  lbaas.opts.SubnetID = subnetID
 }
 ports := service.Spec.Ports
 if len(ports) == 0 {
  return fmt.Errorf("no ports provided to openstack load balancer")
 }
 loadbalancer, err := getLoadbalancerByName(lbaas.lb, loadBalancerName)
 if err != nil {
  return err
 }
 if loadbalancer == nil {
  return fmt.Errorf("loadbalancer %s does not exist", loadBalancerName)
 }
 type portKey struct {
  Protocol listeners.Protocol
  Port     int
 }
 var listenerIDs []string
 lbListeners := make(map[portKey]listeners.Listener)
 allListeners, err := getListenersByLoadBalancerID(lbaas.lb, loadbalancer.ID)
 if err != nil {
  return fmt.Errorf("error getting listeners for LB %s: %v", loadBalancerName, err)
 }
 for _, l := range allListeners {
  key := portKey{Protocol: listeners.Protocol(l.Protocol), Port: l.ProtocolPort}
  lbListeners[key] = l
  listenerIDs = append(listenerIDs, l.ID)
 }
 lbPools := make(map[string]v2pools.Pool)
 for _, listenerID := range listenerIDs {
  pool, err := getPoolByListenerID(lbaas.lb, loadbalancer.ID, listenerID)
  if err != nil {
   return fmt.Errorf("error getting pool for listener %s: %v", listenerID, err)
  }
  lbPools[listenerID] = *pool
 }
 addrs := make(map[string]*v1.Node)
 for _, node := range nodes {
  addr, err := nodeAddressForLB(node)
  if err != nil {
   return err
  }
  addrs[addr] = node
 }
 for portIndex, port := range ports {
  listener, ok := lbListeners[portKey{Protocol: toListenersProtocol(port.Protocol), Port: int(port.Port)}]
  if !ok {
   return fmt.Errorf("loadbalancer %s does not contain required listener for port %d and protocol %s", loadBalancerName, port.Port, port.Protocol)
  }
  pool, ok := lbPools[listener.ID]
  if !ok {
   return fmt.Errorf("loadbalancer %s does not contain required pool for listener %s", loadBalancerName, listener.ID)
  }
  getMembers, err := getMembersByPoolID(lbaas.lb, pool.ID)
  if err != nil {
   return fmt.Errorf("error getting pool members %s: %v", pool.ID, err)
  }
  members := make(map[string]v2pools.Member)
  for _, member := range getMembers {
   members[member.Address] = member
  }
  for addr, node := range addrs {
   if _, ok := members[addr]; ok && members[addr].ProtocolPort == int(port.NodePort) {
    continue
   }
   _, err := v2pools.CreateMember(lbaas.lb, pool.ID, v2pools.CreateMemberOpts{Name: fmt.Sprintf("member_%s_%d_%s", loadbalancer.Name, portIndex, node.Name), Address: addr, ProtocolPort: int(port.NodePort), SubnetID: lbaas.opts.SubnetID}).Extract()
   if err != nil {
    return err
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  for _, member := range members {
   if _, ok := addrs[member.Address]; ok && member.ProtocolPort == int(port.NodePort) {
    continue
   }
   err = v2pools.DeleteMember(lbaas.lb, pool.ID, member.ID).ExtractErr()
   if err != nil && !isNotFound(err) {
    return err
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
 }
 if lbaas.opts.ManageSecurityGroups {
  err := lbaas.updateSecurityGroup(clusterName, service, nodes, loadbalancer)
  if err != nil {
   return fmt.Errorf("failed to update Security Group for loadbalancer service %s/%s: %v", service.Namespace, service.Name, err)
  }
 }
 return nil
}
func (lbaas *LbaasV2) updateSecurityGroup(clusterName string, apiService *v1.Service, nodes []*v1.Node, loadbalancer *loadbalancers.LoadBalancer) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 originalNodeSecurityGroupIDs := lbaas.opts.NodeSecurityGroupIDs
 var err error
 lbaas.opts.NodeSecurityGroupIDs, err = getNodeSecurityGroupIDForLB(lbaas.compute, lbaas.network, nodes)
 if err != nil {
  return fmt.Errorf("failed to find node-security-group for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
 }
 klog.V(4).Infof("find node-security-group %v for loadbalancer service %s/%s", lbaas.opts.NodeSecurityGroupIDs, apiService.Namespace, apiService.Name)
 original := sets.NewString(originalNodeSecurityGroupIDs...)
 current := sets.NewString(lbaas.opts.NodeSecurityGroupIDs...)
 removals := original.Difference(current)
 lbSecGroupName := getSecurityGroupName(apiService)
 lbSecGroupID, err := groups.IDFromName(lbaas.network, lbSecGroupName)
 if err != nil {
  return fmt.Errorf("error occurred finding security group: %s: %v", lbSecGroupName, err)
 }
 ports := apiService.Spec.Ports
 if len(ports) == 0 {
  return fmt.Errorf("no ports provided to openstack load balancer")
 }
 for _, port := range ports {
  for removal := range removals {
   opts := rules.ListOpts{Direction: string(rules.DirIngress), SecGroupID: removal, RemoteGroupID: lbSecGroupID, PortRangeMax: int(port.NodePort), PortRangeMin: int(port.NodePort), Protocol: string(port.Protocol)}
   secGroupRules, err := getSecurityGroupRules(lbaas.network, opts)
   if err != nil && !isNotFound(err) {
    return fmt.Errorf("error finding rules for remote group id %s in security group id %s: %v", lbSecGroupID, removal, err)
   }
   for _, rule := range secGroupRules {
    res := rules.Delete(lbaas.network, rule.ID)
    if res.Err != nil && !isNotFound(res.Err) {
     return fmt.Errorf("error occurred deleting security group rule: %s: %v", rule.ID, res.Err)
    }
   }
  }
  for _, nodeSecurityGroupID := range lbaas.opts.NodeSecurityGroupIDs {
   opts := rules.ListOpts{Direction: string(rules.DirIngress), SecGroupID: nodeSecurityGroupID, RemoteGroupID: lbSecGroupID, PortRangeMax: int(port.NodePort), PortRangeMin: int(port.NodePort), Protocol: string(port.Protocol)}
   secGroupRules, err := getSecurityGroupRules(lbaas.network, opts)
   if err != nil && !isNotFound(err) {
    return fmt.Errorf("error finding rules for remote group id %s in security group id %s: %v", lbSecGroupID, nodeSecurityGroupID, err)
   }
   if len(secGroupRules) != 0 {
    continue
   }
   err = createNodeSecurityGroup(lbaas.network, nodeSecurityGroupID, int(port.NodePort), port.Protocol, lbSecGroupID)
   if err != nil {
    return fmt.Errorf("error occurred creating security group for loadbalancer service %s/%s: %v", apiService.Namespace, apiService.Name, err)
   }
  }
 }
 return nil
}
func (lbaas *LbaasV2) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 loadBalancerName := lbaas.GetLoadBalancerName(ctx, clusterName, service)
 klog.V(4).Infof("EnsureLoadBalancerDeleted(%v, %v)", clusterName, loadBalancerName)
 loadbalancer, err := getLoadbalancerByName(lbaas.lb, loadBalancerName)
 if err != nil && err != ErrNotFound {
  return err
 }
 if loadbalancer == nil {
  return nil
 }
 if loadbalancer.VipPortID != "" {
  portID := loadbalancer.VipPortID
  floatingIP, err := getFloatingIPByPortID(lbaas.network, portID)
  if err != nil && err != ErrNotFound {
   return err
  }
  if floatingIP != nil {
   err = floatingips.Delete(lbaas.network, floatingIP.ID).ExtractErr()
   if err != nil && !isNotFound(err) {
    return err
   }
  }
 }
 listenerList, err := getListenersByLoadBalancerID(lbaas.lb, loadbalancer.ID)
 if err != nil {
  return fmt.Errorf("error getting LB %s listeners: %v", loadbalancer.ID, err)
 }
 var poolIDs []string
 var monitorIDs []string
 for _, listener := range listenerList {
  pool, err := getPoolByListenerID(lbaas.lb, loadbalancer.ID, listener.ID)
  if err != nil && err != ErrNotFound {
   return fmt.Errorf("error getting pool for listener %s: %v", listener.ID, err)
  }
  if pool != nil {
   poolIDs = append(poolIDs, pool.ID)
   if pool.MonitorID != "" {
    monitorIDs = append(monitorIDs, pool.MonitorID)
   }
  }
 }
 for _, monitorID := range monitorIDs {
  err := v2monitors.Delete(lbaas.lb, monitorID).ExtractErr()
  if err != nil && !isNotFound(err) {
   return err
  }
  provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
  if err != nil {
   return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
  }
 }
 for _, poolID := range poolIDs {
  membersList, err := getMembersByPoolID(lbaas.lb, poolID)
  if err != nil && !isNotFound(err) {
   return fmt.Errorf("error getting pool members %s: %v", poolID, err)
  }
  for _, member := range membersList {
   err := v2pools.DeleteMember(lbaas.lb, poolID, member.ID).ExtractErr()
   if err != nil && !isNotFound(err) {
    return err
   }
   provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
   if err != nil {
    return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
   }
  }
  err = v2pools.Delete(lbaas.lb, poolID).ExtractErr()
  if err != nil && !isNotFound(err) {
   return err
  }
  provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
  if err != nil {
   return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
  }
 }
 for _, listener := range listenerList {
  err := listeners.Delete(lbaas.lb, listener.ID).ExtractErr()
  if err != nil && !isNotFound(err) {
   return err
  }
  provisioningStatus, err := waitLoadbalancerActiveProvisioningStatus(lbaas.lb, loadbalancer.ID)
  if err != nil {
   return fmt.Errorf("failed to loadbalance ACTIVE provisioning status %v: %v", provisioningStatus, err)
  }
 }
 err = loadbalancers.Delete(lbaas.lb, loadbalancer.ID).ExtractErr()
 if err != nil && !isNotFound(err) {
  return err
 }
 err = waitLoadbalancerDeleted(lbaas.lb, loadbalancer.ID)
 if err != nil {
  return fmt.Errorf("failed to delete loadbalancer: %v", err)
 }
 if lbaas.opts.ManageSecurityGroups {
  err := lbaas.EnsureSecurityGroupDeleted(clusterName, service)
  if err != nil {
   return fmt.Errorf("Failed to delete Security Group for loadbalancer service %s/%s: %v", service.Namespace, service.Name, err)
  }
 }
 return nil
}
func (lbaas *LbaasV2) EnsureSecurityGroupDeleted(clusterName string, service *v1.Service) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 lbSecGroupName := getSecurityGroupName(service)
 lbSecGroupID, err := groups.IDFromName(lbaas.network, lbSecGroupName)
 if err != nil {
  if isSecurityGroupNotFound(err) {
   return nil
  }
  return fmt.Errorf("Error occurred finding security group: %s: %v", lbSecGroupName, err)
 }
 lbSecGroup := groups.Delete(lbaas.network, lbSecGroupID)
 if lbSecGroup.Err != nil && !isNotFound(lbSecGroup.Err) {
  return lbSecGroup.Err
 }
 if len(lbaas.opts.NodeSecurityGroupIDs) == 0 {
  klog.Warningf("Can not find node-security-group from all the nodes of this cluster when delete loadbalancer service %s/%s", service.Namespace, service.Name)
 } else {
  for _, nodeSecurityGroupID := range lbaas.opts.NodeSecurityGroupIDs {
   opts := rules.ListOpts{SecGroupID: nodeSecurityGroupID, RemoteGroupID: lbSecGroupID}
   secGroupRules, err := getSecurityGroupRules(lbaas.network, opts)
   if err != nil && !isNotFound(err) {
    msg := fmt.Sprintf("Error finding rules for remote group id %s in security group id %s: %v", lbSecGroupID, nodeSecurityGroupID, err)
    return fmt.Errorf(msg)
   }
   for _, rule := range secGroupRules {
    res := rules.Delete(lbaas.network, rule.ID)
    if res.Err != nil && !isNotFound(res.Err) {
     return fmt.Errorf("Error occurred deleting security group rule: %s: %v", rule.ID, res.Err)
    }
   }
  }
 }
 return nil
}
