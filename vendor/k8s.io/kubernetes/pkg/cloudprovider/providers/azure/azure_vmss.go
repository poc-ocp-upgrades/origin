package azure

import (
 "errors"
 "fmt"
 "regexp"
 "sort"
 "strconv"
 "strings"
 "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
 "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
 "github.com/Azure/go-autorest/autorest/to"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 cloudprovider "k8s.io/cloud-provider"
)

var (
 ErrorNotVmssInstance   = errors.New("not a vmss instance")
 scaleSetNameRE         = regexp.MustCompile(`.*/subscriptions/(?:.*)/Microsoft.Compute/virtualMachineScaleSets/(.+)/virtualMachines(?:.*)`)
 resourceGroupRE        = regexp.MustCompile(`.*/subscriptions/(?:.*)/resourceGroups/(.+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?:.*)/virtualMachines(?:.*)`)
 vmssNicResourceGroupRE = regexp.MustCompile(`.*/subscriptions/(?:.*)/resourceGroups/(.+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?:.*)/virtualMachines/(?:.*)/networkInterfaces/(?:.*)`)
 vmssMachineIDTemplate  = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines/%s"
)

type scaleSet struct {
 *Cloud
 availabilitySet                VMSet
 vmssCache                      *timedCache
 vmssVMCache                    *timedCache
 nodeNameToScaleSetMappingCache *timedCache
 availabilitySetNodesCache      *timedCache
}

func newScaleSet(az *Cloud) (VMSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 ss := &scaleSet{Cloud: az, availabilitySet: newAvailabilitySet(az)}
 ss.nodeNameToScaleSetMappingCache, err = ss.newNodeNameToScaleSetMappingCache()
 if err != nil {
  return nil, err
 }
 ss.availabilitySetNodesCache, err = ss.newAvailabilitySetNodesCache()
 if err != nil {
  return nil, err
 }
 ss.vmssCache, err = ss.newVmssCache()
 if err != nil {
  return nil, err
 }
 ss.vmssVMCache, err = ss.newVmssVMCache()
 if err != nil {
  return nil, err
 }
 return ss, nil
}
func (ss *scaleSet) getVmssVM(nodeName string) (ssName, instanceID string, vm compute.VirtualMachineScaleSetVM, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceID, err = getScaleSetVMInstanceID(nodeName)
 if err != nil {
  return ssName, instanceID, vm, err
 }
 ssName, err = ss.getScaleSetNameByNodeName(nodeName)
 if err != nil {
  return ssName, instanceID, vm, err
 }
 if ssName == "" {
  return "", "", vm, cloudprovider.InstanceNotFound
 }
 resourceGroup, err := ss.GetNodeResourceGroup(nodeName)
 if err != nil {
  return "", "", vm, err
 }
 klog.V(4).Infof("getVmssVM gets scaleSetName (%q) and instanceID (%q) for node %q", ssName, instanceID, nodeName)
 key := buildVmssCacheKey(resourceGroup, ss.makeVmssVMName(ssName, instanceID))
 cachedVM, err := ss.vmssVMCache.Get(key)
 if err != nil {
  return ssName, instanceID, vm, err
 }
 if cachedVM == nil {
  klog.Errorf("Can't find node (%q) in any scale sets", nodeName)
  return ssName, instanceID, vm, cloudprovider.InstanceNotFound
 }
 return ssName, instanceID, *(cachedVM.(*compute.VirtualMachineScaleSetVM)), nil
}
func (ss *scaleSet) GetPowerStatusByNodeName(name string) (powerState string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, _, vm, err := ss.getVmssVM(name)
 if err != nil {
  return powerState, err
 }
 if vm.InstanceView != nil && vm.InstanceView.Statuses != nil {
  statuses := *vm.InstanceView.Statuses
  for _, status := range statuses {
   state := to.String(status.Code)
   if strings.HasPrefix(state, vmPowerStatePrefix) {
    return strings.TrimPrefix(state, vmPowerStatePrefix), nil
   }
  }
 }
 return "", fmt.Errorf("failed to get power status for node %q", name)
}
func (ss *scaleSet) getVmssVMByInstanceID(resourceGroup, scaleSetName, instanceID string) (vm compute.VirtualMachineScaleSetVM, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmName := ss.makeVmssVMName(scaleSetName, instanceID)
 key := buildVmssCacheKey(resourceGroup, vmName)
 cachedVM, err := ss.vmssVMCache.Get(key)
 if err != nil {
  return vm, err
 }
 if cachedVM == nil {
  klog.Errorf("couldn't find vmss virtual machine by scaleSetName (%s) and instanceID (%s)", scaleSetName, instanceID)
  return vm, cloudprovider.InstanceNotFound
 }
 return *(cachedVM.(*compute.VirtualMachineScaleSetVM)), nil
}
func (ss *scaleSet) GetInstanceIDByNodeName(name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 managedByAS, err := ss.isNodeManagedByAvailabilitySet(name)
 if err != nil {
  klog.Errorf("Failed to check isNodeManagedByAvailabilitySet: %v", err)
  return "", err
 }
 if managedByAS {
  return ss.availabilitySet.GetInstanceIDByNodeName(name)
 }
 _, _, vm, err := ss.getVmssVM(name)
 if err != nil {
  return "", err
 }
 return *vm.ID, nil
}
func (ss *scaleSet) GetNodeNameByProviderID(providerID string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scaleSetName, err := extractScaleSetNameByProviderID(providerID)
 if err != nil {
  klog.V(4).Infof("Can not extract scale set name from providerID (%s), assuming it is mananaged by availability set: %v", providerID, err)
  return ss.availabilitySet.GetNodeNameByProviderID(providerID)
 }
 resourceGroup, err := extractResourceGroupByProviderID(providerID)
 if err != nil {
  return "", fmt.Errorf("error of extracting resource group for node %q", providerID)
 }
 instanceID, err := getLastSegment(providerID)
 if err != nil {
  klog.V(4).Infof("Can not extract instanceID from providerID (%s), assuming it is mananaged by availability set: %v", providerID, err)
  return ss.availabilitySet.GetNodeNameByProviderID(providerID)
 }
 vm, err := ss.getVmssVMByInstanceID(resourceGroup, scaleSetName, instanceID)
 if err != nil {
  return "", err
 }
 if vm.OsProfile != nil && vm.OsProfile.ComputerName != nil {
  nodeName := strings.ToLower(*vm.OsProfile.ComputerName)
  return types.NodeName(nodeName), nil
 }
 return "", nil
}
func (ss *scaleSet) GetInstanceTypeByNodeName(name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 managedByAS, err := ss.isNodeManagedByAvailabilitySet(name)
 if err != nil {
  klog.Errorf("Failed to check isNodeManagedByAvailabilitySet: %v", err)
  return "", err
 }
 if managedByAS {
  return ss.availabilitySet.GetInstanceTypeByNodeName(name)
 }
 _, _, vm, err := ss.getVmssVM(name)
 if err != nil {
  return "", err
 }
 if vm.Sku != nil && vm.Sku.Name != nil {
  return *vm.Sku.Name, nil
 }
 return "", nil
}
func (ss *scaleSet) GetZoneByNodeName(name string) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 managedByAS, err := ss.isNodeManagedByAvailabilitySet(name)
 if err != nil {
  klog.Errorf("Failed to check isNodeManagedByAvailabilitySet: %v", err)
  return cloudprovider.Zone{}, err
 }
 if managedByAS {
  return ss.availabilitySet.GetZoneByNodeName(name)
 }
 _, _, vm, err := ss.getVmssVM(name)
 if err != nil {
  return cloudprovider.Zone{}, err
 }
 var failureDomain string
 if vm.Zones != nil && len(*vm.Zones) > 0 {
  zones := *vm.Zones
  zoneID, err := strconv.Atoi(zones[0])
  if err != nil {
   return cloudprovider.Zone{}, fmt.Errorf("failed to parse zone %q: %v", zones, err)
  }
  failureDomain = ss.makeZone(zoneID)
 } else if vm.InstanceView != nil && vm.InstanceView.PlatformFaultDomain != nil {
  failureDomain = strconv.Itoa(int(*vm.InstanceView.PlatformFaultDomain))
 }
 return cloudprovider.Zone{FailureDomain: failureDomain, Region: *vm.Location}, nil
}
func (ss *scaleSet) GetPrimaryVMSetName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ss.Config.PrimaryScaleSetName
}
func (ss *scaleSet) GetIPByNodeName(nodeName string) (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nic, err := ss.GetPrimaryInterface(nodeName)
 if err != nil {
  klog.Errorf("error: ss.GetIPByNodeName(%s), GetPrimaryInterface(%q), err=%v", nodeName, nodeName, err)
  return "", "", err
 }
 ipConfig, err := getPrimaryIPConfig(nic)
 if err != nil {
  klog.Errorf("error: ss.GetIPByNodeName(%s), getPrimaryIPConfig(%v), err=%v", nodeName, nic, err)
  return "", "", err
 }
 internalIP := *ipConfig.PrivateIPAddress
 publicIP := ""
 if ipConfig.PublicIPAddress != nil && ipConfig.PublicIPAddress.ID != nil {
  pipID := *ipConfig.PublicIPAddress.ID
  pipName, err := getLastSegment(pipID)
  if err != nil {
   return "", "", fmt.Errorf("failed to get publicIP name for node %q with pipID %q", nodeName, pipID)
  }
  resourceGroup, err := ss.GetNodeResourceGroup(nodeName)
  if err != nil {
   return "", "", err
  }
  pip, existsPip, err := ss.getPublicIPAddress(resourceGroup, pipName)
  if err != nil {
   return "", "", err
  }
  if existsPip {
   publicIP = *pip.IPAddress
  }
 }
 return internalIP, publicIP, nil
}
func (ss *scaleSet) getPrimaryInterfaceID(machine compute.VirtualMachineScaleSetVM) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(*machine.NetworkProfile.NetworkInterfaces) == 1 {
  return *(*machine.NetworkProfile.NetworkInterfaces)[0].ID, nil
 }
 for _, ref := range *machine.NetworkProfile.NetworkInterfaces {
  if *ref.Primary {
   return *ref.ID, nil
  }
 }
 return "", fmt.Errorf("failed to find a primary nic for the vm. vmname=%q", *machine.Name)
}
func getScaleSetVMInstanceID(machineName string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nameLength := len(machineName)
 if nameLength < 6 {
  return "", ErrorNotVmssInstance
 }
 instanceID, err := strconv.ParseUint(machineName[nameLength-6:], 36, 64)
 if err != nil {
  return "", ErrorNotVmssInstance
 }
 return fmt.Sprintf("%d", instanceID), nil
}
func extractScaleSetNameByProviderID(providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 matches := scaleSetNameRE.FindStringSubmatch(providerID)
 if len(matches) != 2 {
  return "", ErrorNotVmssInstance
 }
 return matches[1], nil
}
func extractResourceGroupByProviderID(providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 matches := resourceGroupRE.FindStringSubmatch(providerID)
 if len(matches) != 2 {
  return "", ErrorNotVmssInstance
 }
 return matches[1], nil
}
func (ss *scaleSet) listScaleSets(resourceGroup string) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 ctx, cancel := getContextWithCancel()
 defer cancel()
 allScaleSets, err := ss.VirtualMachineScaleSetsClient.List(ctx, resourceGroup)
 if err != nil {
  klog.Errorf("VirtualMachineScaleSetsClient.List failed: %v", err)
  return nil, err
 }
 ssNames := make([]string, len(allScaleSets))
 for i := range allScaleSets {
  ssNames[i] = *(allScaleSets[i].Name)
 }
 return ssNames, nil
}
func (ss *scaleSet) listScaleSetVMs(scaleSetName, resourceGroup string) ([]compute.VirtualMachineScaleSetVM, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 ctx, cancel := getContextWithCancel()
 defer cancel()
 allVMs, err := ss.VirtualMachineScaleSetVMsClient.List(ctx, resourceGroup, scaleSetName, "", "", string(compute.InstanceView))
 if err != nil {
  klog.Errorf("VirtualMachineScaleSetVMsClient.List failed: %v", err)
  return nil, err
 }
 return allVMs, nil
}
func (ss *scaleSet) getAgentPoolScaleSets(nodes []*v1.Node) (*[]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 agentPoolScaleSets := &[]string{}
 for nx := range nodes {
  if isMasterNode(nodes[nx]) {
   continue
  }
  if ss.ShouldNodeExcludedFromLoadBalancer(nodes[nx]) {
   continue
  }
  nodeName := nodes[nx].Name
  ssName, err := ss.getScaleSetNameByNodeName(nodeName)
  if err != nil {
   return nil, err
  }
  if ssName == "" {
   klog.V(3).Infof("Node %q is not belonging to any known scale sets", nodeName)
   continue
  }
  *agentPoolScaleSets = append(*agentPoolScaleSets, ssName)
 }
 return agentPoolScaleSets, nil
}
func (ss *scaleSet) GetVMSetNames(service *v1.Service, nodes []*v1.Node) (vmSetNames *[]string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hasMode, isAuto, serviceVMSetNames := getServiceLoadBalancerMode(service)
 if !hasMode {
  scaleSetNames := &[]string{ss.Config.PrimaryScaleSetName}
  return scaleSetNames, nil
 }
 scaleSetNames, err := ss.getAgentPoolScaleSets(nodes)
 if err != nil {
  klog.Errorf("ss.GetVMSetNames - getAgentPoolScaleSets failed err=(%v)", err)
  return nil, err
 }
 if len(*scaleSetNames) == 0 {
  klog.Errorf("ss.GetVMSetNames - No scale sets found for nodes in the cluster, node count(%d)", len(nodes))
  return nil, fmt.Errorf("No scale sets found for nodes, node count(%d)", len(nodes))
 }
 sort.Strings(*scaleSetNames)
 if !isAuto {
  if serviceVMSetNames == nil || len(serviceVMSetNames) == 0 {
   return nil, fmt.Errorf("service annotation for LoadBalancerMode is empty, it should have __auto__ or availability sets value")
  }
  var found bool
  for sasx := range serviceVMSetNames {
   for asx := range *scaleSetNames {
    if strings.EqualFold((*scaleSetNames)[asx], serviceVMSetNames[sasx]) {
     found = true
     serviceVMSetNames[sasx] = (*scaleSetNames)[asx]
     break
    }
   }
   if !found {
    klog.Errorf("ss.GetVMSetNames - scale set (%s) in service annotation not found", serviceVMSetNames[sasx])
    return nil, fmt.Errorf("scale set (%s) - not found", serviceVMSetNames[sasx])
   }
  }
  vmSetNames = &serviceVMSetNames
 }
 return vmSetNames, nil
}
func extractResourceGroupByVMSSNicID(nicID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 matches := vmssNicResourceGroupRE.FindStringSubmatch(nicID)
 if len(matches) != 2 {
  return "", fmt.Errorf("error of extracting resourceGroup from nicID %q", nicID)
 }
 return matches[1], nil
}
func (ss *scaleSet) GetPrimaryInterface(nodeName string) (network.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 managedByAS, err := ss.isNodeManagedByAvailabilitySet(nodeName)
 if err != nil {
  klog.Errorf("Failed to check isNodeManagedByAvailabilitySet: %v", err)
  return network.Interface{}, err
 }
 if managedByAS {
  return ss.availabilitySet.GetPrimaryInterface(nodeName)
 }
 ssName, instanceID, vm, err := ss.getVmssVM(nodeName)
 if err != nil {
  if err == ErrorNotVmssInstance {
   return ss.availabilitySet.GetPrimaryInterface(nodeName)
  }
  klog.Errorf("error: ss.GetPrimaryInterface(%s), ss.getVmssVM(%s), err=%v", nodeName, nodeName, err)
  return network.Interface{}, err
 }
 primaryInterfaceID, err := ss.getPrimaryInterfaceID(vm)
 if err != nil {
  klog.Errorf("error: ss.GetPrimaryInterface(%s), ss.getPrimaryInterfaceID(), err=%v", nodeName, err)
  return network.Interface{}, err
 }
 nicName, err := getLastSegment(primaryInterfaceID)
 if err != nil {
  klog.Errorf("error: ss.GetPrimaryInterface(%s), getLastSegment(%s), err=%v", nodeName, primaryInterfaceID, err)
  return network.Interface{}, err
 }
 resourceGroup, err := extractResourceGroupByVMSSNicID(primaryInterfaceID)
 if err != nil {
  return network.Interface{}, err
 }
 ctx, cancel := getContextWithCancel()
 defer cancel()
 nic, err := ss.InterfacesClient.GetVirtualMachineScaleSetNetworkInterface(ctx, resourceGroup, ssName, instanceID, nicName, "")
 if err != nil {
  klog.Errorf("error: ss.GetPrimaryInterface(%s), ss.GetVirtualMachineScaleSetNetworkInterface.Get(%s, %s, %s), err=%v", nodeName, resourceGroup, ssName, nicName, err)
  return network.Interface{}, err
 }
 if nic.Location == nil || *nic.Location == "" {
  nic.Location = vm.Location
 }
 return nic, nil
}
func (ss *scaleSet) getScaleSetWithRetry(service *v1.Service, name string) (compute.VirtualMachineScaleSet, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var result compute.VirtualMachineScaleSet
 var exists bool
 err := wait.ExponentialBackoff(ss.requestBackoff(), func() (bool, error) {
  cached, retryErr := ss.vmssCache.Get(name)
  if retryErr != nil {
   ss.Event(service, v1.EventTypeWarning, "GetVirtualMachineScaleSet", retryErr.Error())
   klog.Errorf("backoff: failure for scale set %q, will retry,err=%v", name, retryErr)
   return false, nil
  }
  klog.V(4).Infof("backoff: success for scale set %q", name)
  if cached != nil {
   exists = true
   result = *(cached.(*compute.VirtualMachineScaleSet))
  }
  return true, nil
 })
 return result, exists, err
}
func (ss *scaleSet) getPrimaryNetworkConfiguration(networkConfigurationList *[]compute.VirtualMachineScaleSetNetworkConfiguration, scaleSetName string) (*compute.VirtualMachineScaleSetNetworkConfiguration, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 networkConfigurations := *networkConfigurationList
 if len(networkConfigurations) == 1 {
  return &networkConfigurations[0], nil
 }
 for idx := range networkConfigurations {
  networkConfig := &networkConfigurations[idx]
  if networkConfig.Primary != nil && *networkConfig.Primary == true {
   return networkConfig, nil
  }
 }
 return nil, fmt.Errorf("failed to find a primary network configuration for the scale set %q", scaleSetName)
}
func (ss *scaleSet) getPrimaryIPConfigForScaleSet(config *compute.VirtualMachineScaleSetNetworkConfiguration, scaleSetName string) (*compute.VirtualMachineScaleSetIPConfiguration, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ipConfigurations := *config.IPConfigurations
 if len(ipConfigurations) == 1 {
  return &ipConfigurations[0], nil
 }
 for idx := range ipConfigurations {
  ipConfig := &ipConfigurations[idx]
  if ipConfig.Primary != nil && *ipConfig.Primary == true {
   return ipConfig, nil
  }
 }
 return nil, fmt.Errorf("failed to find a primary IP configuration for the scale set %q", scaleSetName)
}
func (ss *scaleSet) createOrUpdateVMSSWithRetry(service *v1.Service, virtualMachineScaleSet compute.VirtualMachineScaleSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(ss.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := ss.VirtualMachineScaleSetsClient.CreateOrUpdate(ctx, ss.ResourceGroup, *virtualMachineScaleSet.Name, virtualMachineScaleSet)
  klog.V(10).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate(%s): end", *virtualMachineScaleSet.Name)
  return ss.processHTTPRetryResponse(service, "CreateOrUpdateVMSS", resp, err)
 })
}
func (ss *scaleSet) updateVMSSInstancesWithRetry(service *v1.Service, scaleSetName string, vmInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(ss.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := ss.VirtualMachineScaleSetsClient.UpdateInstances(ctx, ss.ResourceGroup, scaleSetName, vmInstanceIDs)
  klog.V(10).Infof("VirtualMachineScaleSetsClient.UpdateInstances(%s): end", scaleSetName)
  return ss.processHTTPRetryResponse(service, "CreateOrUpdateVMSSInstance", resp, err)
 })
}
func (ss *scaleSet) getNodesScaleSets(nodes []*v1.Node) (map[string]sets.String, []*v1.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scalesets := make(map[string]sets.String)
 standardNodes := []*v1.Node{}
 for _, curNode := range nodes {
  if ss.useStandardLoadBalancer() && ss.excludeMasterNodesFromStandardLB() && isMasterNode(curNode) {
   klog.V(4).Infof("Excluding master node %q from load balancer backendpool", curNode.Name)
   continue
  }
  if ss.ShouldNodeExcludedFromLoadBalancer(curNode) {
   klog.V(4).Infof("Excluding unmanaged/external-resource-group node %q", curNode.Name)
   continue
  }
  curScaleSetName, err := extractScaleSetNameByProviderID(curNode.Spec.ProviderID)
  if err != nil {
   klog.V(4).Infof("Node %q is not belonging to any scale sets, assuming it is belong to availability sets", curNode.Name)
   standardNodes = append(standardNodes, curNode)
   continue
  }
  if _, ok := scalesets[curScaleSetName]; !ok {
   scalesets[curScaleSetName] = sets.NewString()
  }
  instanceID, err := getLastSegment(curNode.Spec.ProviderID)
  if err != nil {
   klog.Errorf("Failed to get instance ID for node %q: %v", curNode.Spec.ProviderID, err)
   return nil, nil, err
  }
  scalesets[curScaleSetName].Insert(instanceID)
 }
 return scalesets, standardNodes, nil
}
func (ss *scaleSet) ensureHostsInVMSetPool(service *v1.Service, backendPoolID string, vmSetName string, instanceIDs []string, isInternal bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(3).Infof("ensuring hosts %q of scaleset %q in LB backendpool %q", instanceIDs, vmSetName, backendPoolID)
 serviceName := getServiceName(service)
 virtualMachineScaleSet, exists, err := ss.getScaleSetWithRetry(service, vmSetName)
 if err != nil {
  klog.Errorf("ss.getScaleSetWithRetry(%s) for service %q failed: %v", vmSetName, serviceName, err)
  return err
 }
 if !exists {
  errorMessage := fmt.Errorf("Scale set %q not found", vmSetName)
  klog.Errorf("%v", errorMessage)
  return errorMessage
 }
 networkConfigureList := virtualMachineScaleSet.VirtualMachineProfile.NetworkProfile.NetworkInterfaceConfigurations
 primaryNetworkConfiguration, err := ss.getPrimaryNetworkConfiguration(networkConfigureList, vmSetName)
 if err != nil {
  return err
 }
 primaryIPConfiguration, err := ss.getPrimaryIPConfigForScaleSet(primaryNetworkConfiguration, vmSetName)
 if err != nil {
  return err
 }
 foundPool := false
 newBackendPools := []compute.SubResource{}
 if primaryIPConfiguration.LoadBalancerBackendAddressPools != nil {
  newBackendPools = *primaryIPConfiguration.LoadBalancerBackendAddressPools
 }
 for _, existingPool := range newBackendPools {
  if strings.EqualFold(backendPoolID, *existingPool.ID) {
   foundPool = true
   break
  }
 }
 if !foundPool {
  if ss.useStandardLoadBalancer() && len(newBackendPools) > 0 {
   for _, pool := range newBackendPools {
    backendPool := *pool.ID
    matches := backendPoolIDRE.FindStringSubmatch(backendPool)
    if len(matches) == 2 {
     lbName := matches[1]
     if strings.HasSuffix(lbName, InternalLoadBalancerNameSuffix) == isInternal {
      klog.V(4).Infof("vmss %q has already been added to LB %q, omit adding it to a new one", vmSetName, lbName)
      return nil
     }
    }
   }
  }
  newBackendPools = append(newBackendPools, compute.SubResource{ID: to.StringPtr(backendPoolID)})
  primaryIPConfiguration.LoadBalancerBackendAddressPools = &newBackendPools
  ctx, cancel := getContextWithCancel()
  defer cancel()
  klog.V(3).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate for service (%s): scale set (%s) - updating", serviceName, vmSetName)
  resp, err := ss.VirtualMachineScaleSetsClient.CreateOrUpdate(ctx, ss.ResourceGroup, vmSetName, virtualMachineScaleSet)
  klog.V(10).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate(%q): end", vmSetName)
  if ss.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
   klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate for service (%s): scale set (%s) - updating, err=%v", serviceName, vmSetName, err)
   retryErr := ss.createOrUpdateVMSSWithRetry(service, virtualMachineScaleSet)
   if retryErr != nil {
    err = retryErr
    klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate for service (%s) abort backoff: scale set (%s) - updating", serviceName, vmSetName)
   }
  }
  if err != nil {
   return err
  }
 }
 vmInstanceIDs := compute.VirtualMachineScaleSetVMInstanceRequiredIDs{InstanceIds: &instanceIDs}
 ctx, cancel := getContextWithCancel()
 defer cancel()
 instanceResp, err := ss.VirtualMachineScaleSetsClient.UpdateInstances(ctx, ss.ResourceGroup, vmSetName, vmInstanceIDs)
 klog.V(10).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate(%q): end", vmSetName)
 if ss.CloudProviderBackoff && shouldRetryHTTPRequest(instanceResp, err) {
  klog.V(2).Infof("VirtualMachineScaleSetsClient.UpdateInstances for service (%s): scale set (%s) - updating, err=%v", serviceName, vmSetName, err)
  retryErr := ss.updateVMSSInstancesWithRetry(service, vmSetName, vmInstanceIDs)
  if retryErr != nil {
   err = retryErr
   klog.V(2).Infof("VirtualMachineScaleSetsClient.UpdateInstances for service (%s) abort backoff: scale set (%s) - updating", serviceName, vmSetName)
  }
 }
 if err != nil {
  return err
 }
 return nil
}
func (ss *scaleSet) EnsureHostsInPool(service *v1.Service, nodes []*v1.Node, backendPoolID string, vmSetName string, isInternal bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 serviceName := getServiceName(service)
 scalesets, standardNodes, err := ss.getNodesScaleSets(nodes)
 if err != nil {
  klog.Errorf("getNodesScaleSets() for service %q failed: %v", serviceName, err)
  return err
 }
 for ssName, instanceIDs := range scalesets {
  if !ss.useStandardLoadBalancer() && !strings.EqualFold(ssName, vmSetName) {
   continue
  }
  if instanceIDs.Len() == 0 {
   klog.V(3).Infof("scale set %q has 0 nodes, adding it to load balancer anyway", ssName)
   instanceIDs.Insert("*")
  }
  err := ss.ensureHostsInVMSetPool(service, backendPoolID, ssName, instanceIDs.List(), isInternal)
  if err != nil {
   klog.Errorf("ensureHostsInVMSetPool() with scaleSet %q for service %q failed: %v", ssName, serviceName, err)
   return err
  }
 }
 if ss.useStandardLoadBalancer() && len(standardNodes) > 0 {
  err := ss.availabilitySet.EnsureHostsInPool(service, standardNodes, backendPoolID, "", isInternal)
  if err != nil {
   klog.Errorf("availabilitySet.EnsureHostsInPool() for service %q failed: %v", serviceName, err)
   return err
  }
 }
 return nil
}
func (ss *scaleSet) ensureScaleSetBackendPoolDeleted(service *v1.Service, poolID, ssName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(3).Infof("ensuring backend pool %q deleted from scaleset %q", poolID, ssName)
 virtualMachineScaleSet, exists, err := ss.getScaleSetWithRetry(service, ssName)
 if err != nil {
  klog.Errorf("ss.ensureScaleSetBackendPoolDeleted(%s, %s) getScaleSetWithRetry(%s) failed: %v", poolID, ssName, ssName, err)
  return err
 }
 if !exists {
  klog.V(2).Infof("ss.ensureScaleSetBackendPoolDeleted(%s, %s), scale set %s has already been non-exist", poolID, ssName, ssName)
  return nil
 }
 networkConfigureList := virtualMachineScaleSet.VirtualMachineProfile.NetworkProfile.NetworkInterfaceConfigurations
 primaryNetworkConfiguration, err := ss.getPrimaryNetworkConfiguration(networkConfigureList, ssName)
 if err != nil {
  return err
 }
 primaryIPConfiguration, err := ss.getPrimaryIPConfigForScaleSet(primaryNetworkConfiguration, ssName)
 if err != nil {
  return err
 }
 if primaryIPConfiguration.LoadBalancerBackendAddressPools == nil || len(*primaryIPConfiguration.LoadBalancerBackendAddressPools) == 0 {
  return nil
 }
 existingBackendPools := *primaryIPConfiguration.LoadBalancerBackendAddressPools
 newBackendPools := []compute.SubResource{}
 foundPool := false
 for i := len(existingBackendPools) - 1; i >= 0; i-- {
  curPool := existingBackendPools[i]
  if strings.EqualFold(poolID, *curPool.ID) {
   klog.V(10).Infof("ensureScaleSetBackendPoolDeleted gets unwanted backend pool %q for scale set %q", poolID, ssName)
   foundPool = true
   newBackendPools = append(existingBackendPools[:i], existingBackendPools[i+1:]...)
  }
 }
 if !foundPool {
  return nil
 }
 primaryIPConfiguration.LoadBalancerBackendAddressPools = &newBackendPools
 klog.V(3).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate: scale set (%s) - updating", ssName)
 ctx, cancel := getContextWithCancel()
 defer cancel()
 resp, err := ss.VirtualMachineScaleSetsClient.CreateOrUpdate(ctx, ss.ResourceGroup, ssName, virtualMachineScaleSet)
 klog.V(10).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate(%q): end", ssName)
 if ss.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
  klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate: scale set (%s) - updating, err=%v", ssName, err)
  retryErr := ss.createOrUpdateVMSSWithRetry(service, virtualMachineScaleSet)
  if retryErr != nil {
   err = retryErr
   klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate abort backoff: scale set (%s) - updating", ssName)
  }
 }
 if err != nil {
  return err
 }
 instanceIDs := []string{"*"}
 vmInstanceIDs := compute.VirtualMachineScaleSetVMInstanceRequiredIDs{InstanceIds: &instanceIDs}
 instanceCtx, instanceCancel := getContextWithCancel()
 defer instanceCancel()
 instanceResp, err := ss.VirtualMachineScaleSetsClient.UpdateInstances(instanceCtx, ss.ResourceGroup, ssName, vmInstanceIDs)
 klog.V(10).Infof("VirtualMachineScaleSetsClient.UpdateInstances(%q): end", ssName)
 if ss.CloudProviderBackoff && shouldRetryHTTPRequest(instanceResp, err) {
  klog.V(2).Infof("VirtualMachineScaleSetsClient.UpdateInstances scale set (%s) - updating, err=%v", ssName, err)
  retryErr := ss.updateVMSSInstancesWithRetry(service, ssName, vmInstanceIDs)
  if retryErr != nil {
   err = retryErr
   klog.V(2).Infof("VirtualMachineScaleSetsClient.UpdateInstances abort backoff: scale set (%s) - updating", ssName)
  }
 }
 if err != nil {
  return err
 }
 if len(newBackendPools) == 0 {
  updateCtx, updateCancel := getContextWithCancel()
  defer updateCancel()
  klog.V(3).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate: scale set (%s) - updating second time", ssName)
  resp, err = ss.VirtualMachineScaleSetsClient.CreateOrUpdate(updateCtx, ss.ResourceGroup, ssName, virtualMachineScaleSet)
  klog.V(10).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate(%q): end", ssName)
  if ss.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
   klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate: scale set (%s) - updating, err=%v", ssName, err)
   retryErr := ss.createOrUpdateVMSSWithRetry(service, virtualMachineScaleSet)
   if retryErr != nil {
    klog.V(2).Infof("VirtualMachineScaleSetsClient.CreateOrUpdate abort backoff: scale set (%s) - updating", ssName)
   }
  }
 }
 return nil
}
func (ss *scaleSet) EnsureBackendPoolDeleted(service *v1.Service, poolID, vmSetName string, backendAddressPools *[]network.BackendAddressPool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if backendAddressPools == nil {
  return nil
 }
 scalesets := sets.NewString()
 for _, backendPool := range *backendAddressPools {
  if strings.EqualFold(*backendPool.ID, poolID) && backendPool.BackendIPConfigurations != nil {
   for _, ipConfigurations := range *backendPool.BackendIPConfigurations {
    if ipConfigurations.ID == nil {
     continue
    }
    ssName, err := extractScaleSetNameByProviderID(*ipConfigurations.ID)
    if err != nil {
     klog.V(4).Infof("backend IP configuration %q is not belonging to any vmss, omit it", *ipConfigurations.ID)
     continue
    }
    scalesets.Insert(ssName)
   }
   break
  }
 }
 for ssName := range scalesets {
  if !ss.useStandardLoadBalancer() && !strings.EqualFold(ssName, vmSetName) {
   continue
  }
  err := ss.ensureScaleSetBackendPoolDeleted(service, poolID, ssName)
  if err != nil {
   klog.Errorf("ensureScaleSetBackendPoolDeleted() with scaleSet %q failed: %v", ssName, err)
   return err
  }
 }
 return nil
}
func (az *Cloud) getVmssMachineID(resourceGroup, scaleSetName, instanceID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf(vmssMachineIDTemplate, az.SubscriptionID, resourceGroup, scaleSetName, instanceID)
}
