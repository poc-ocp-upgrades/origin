package azure

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"hash/crc32"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/uuid"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	loadBalancerMinimumPriority    = 500
	loadBalancerMaximumPriority    = 4096
	machineIDTemplate              = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s"
	availabilitySetIDTemplate      = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/availabilitySets/%s"
	frontendIPConfigIDTemplate     = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/loadBalancers/%s/frontendIPConfigurations/%s"
	backendPoolIDTemplate          = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/loadBalancers/%s/backendAddressPools/%s"
	loadBalancerProbeIDTemplate    = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/loadBalancers/%s/probes/%s"
	InternalLoadBalancerNameSuffix = "-internal"
	nodeLabelRole                  = "kubernetes.io/role"
	nicFailedState                 = "Failed"
	storageAccountNameMaxLength    = 24
)

var errNotInVMSet = errors.New("vm is not in the vmset")
var providerIDRE = regexp.MustCompile(`^` + CloudProviderName + `://(?:.*)/Microsoft.Compute/virtualMachines/(.+)$`)
var backendPoolIDRE = regexp.MustCompile(`^/subscriptions/(?:.*)/resourceGroups/(?:.*)/providers/Microsoft.Network/loadBalancers/(.+)/backendAddressPools/(?:.*)`)
var nicResourceGroupRE = regexp.MustCompile(`.*/subscriptions/(?:.*)/resourceGroups/(.+)/providers/Microsoft.Network/networkInterfaces/(?:.*)`)
var publicIPResourceGroupRE = regexp.MustCompile(`.*/subscriptions/(?:.*)/resourceGroups/(.+)/providers/Microsoft.Network/publicIPAddresses/(?:.*)`)

func (az *Cloud) getStandardMachineID(resourceGroup, machineName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(machineIDTemplate, az.SubscriptionID, resourceGroup, machineName)
}
func (az *Cloud) getAvailabilitySetID(resourceGroup, availabilitySetName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(availabilitySetIDTemplate, az.SubscriptionID, resourceGroup, availabilitySetName)
}
func (az *Cloud) getFrontendIPConfigID(lbName, backendPoolName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(frontendIPConfigIDTemplate, az.SubscriptionID, az.ResourceGroup, lbName, backendPoolName)
}
func (az *Cloud) getBackendPoolID(lbName, backendPoolName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(backendPoolIDTemplate, az.SubscriptionID, az.ResourceGroup, lbName, backendPoolName)
}
func (az *Cloud) getLoadBalancerProbeID(lbName, lbRuleName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(loadBalancerProbeIDTemplate, az.SubscriptionID, az.ResourceGroup, lbName, lbRuleName)
}
func (az *Cloud) mapLoadBalancerNameToVMSet(lbName string, clusterName string) (vmSetName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmSetName = strings.TrimSuffix(lbName, InternalLoadBalancerNameSuffix)
	if strings.EqualFold(clusterName, vmSetName) {
		vmSetName = az.vmSet.GetPrimaryVMSetName()
	}
	return vmSetName
}
func (az *Cloud) getAzureLoadBalancerName(clusterName string, vmSetName string, isInternal bool) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lbNamePrefix := vmSetName
	if strings.EqualFold(vmSetName, az.vmSet.GetPrimaryVMSetName()) || az.useStandardLoadBalancer() {
		lbNamePrefix = clusterName
	}
	if isInternal {
		return fmt.Sprintf("%s%s", lbNamePrefix, InternalLoadBalancerNameSuffix)
	}
	return lbNamePrefix
}
func isMasterNode(node *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if val, ok := node.Labels[nodeLabelRole]; ok && val == "master" {
		return true
	}
	return false
}
func getLastSegment(ID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(ID, "/")
	name := parts[len(parts)-1]
	if len(name) == 0 {
		return "", fmt.Errorf("resource name was missing from identifier")
	}
	return name, nil
}
func getProtocolsFromKubernetesProtocol(protocol v1.Protocol) (*network.TransportProtocol, *network.SecurityRuleProtocol, *network.ProbeProtocol, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var transportProto network.TransportProtocol
	var securityProto network.SecurityRuleProtocol
	var probeProto network.ProbeProtocol
	switch protocol {
	case v1.ProtocolTCP:
		transportProto = network.TransportProtocolTCP
		securityProto = network.SecurityRuleProtocolTCP
		probeProto = network.ProbeProtocolTCP
		return &transportProto, &securityProto, &probeProto, nil
	case v1.ProtocolUDP:
		transportProto = network.TransportProtocolUDP
		securityProto = network.SecurityRuleProtocolUDP
		return &transportProto, &securityProto, nil, nil
	default:
		return &transportProto, &securityProto, &probeProto, fmt.Errorf("only TCP and UDP are supported for Azure LoadBalancers")
	}
}
func getPrimaryInterfaceID(machine compute.VirtualMachine) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func getPrimaryIPConfig(nic network.Interface) (*network.InterfaceIPConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if nic.IPConfigurations == nil {
		return nil, fmt.Errorf("nic.IPConfigurations for nic (nicname=%q) is nil", *nic.Name)
	}
	if len(*nic.IPConfigurations) == 1 {
		return &((*nic.IPConfigurations)[0]), nil
	}
	for _, ref := range *nic.IPConfigurations {
		if *ref.Primary {
			return &ref, nil
		}
	}
	return nil, fmt.Errorf("failed to determine the primary ipconfig. nicname=%q", *nic.Name)
}
func isInternalLoadBalancer(lb *network.LoadBalancer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasSuffix(*lb.Name, InternalLoadBalancerNameSuffix)
}
func getBackendPoolName(clusterName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return clusterName
}
func (az *Cloud) getLoadBalancerRuleName(service *v1.Service, protocol v1.Protocol, port int32, subnetName *string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prefix := az.getRulePrefix(service)
	if subnetName == nil {
		return fmt.Sprintf("%s-%s-%d", prefix, protocol, port)
	}
	return fmt.Sprintf("%s-%s-%s-%d", prefix, *subnetName, protocol, port)
}
func (az *Cloud) getSecurityRuleName(service *v1.Service, port v1.ServicePort, sourceAddrPrefix string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if useSharedSecurityRule(service) {
		safePrefix := strings.Replace(sourceAddrPrefix, "/", "_", -1)
		return fmt.Sprintf("shared-%s-%d-%s", port.Protocol, port.Port, safePrefix)
	}
	safePrefix := strings.Replace(sourceAddrPrefix, "/", "_", -1)
	rulePrefix := az.getRulePrefix(service)
	return fmt.Sprintf("%s-%s-%d-%s", rulePrefix, port.Protocol, port.Port, safePrefix)
}
func getServiceName(service *v1.Service) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s", service.Namespace, service.Name)
}
func (az *Cloud) getRulePrefix(service *v1.Service) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az.GetLoadBalancerName(context.TODO(), "", service)
}
func (az *Cloud) getPublicIPName(clusterName string, service *v1.Service) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s-%s", clusterName, az.GetLoadBalancerName(context.TODO(), clusterName, service))
}
func (az *Cloud) serviceOwnsRule(service *v1.Service, rule string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prefix := az.getRulePrefix(service)
	return strings.HasPrefix(strings.ToUpper(rule), strings.ToUpper(prefix))
}
func (az *Cloud) serviceOwnsFrontendIP(fip network.FrontendIPConfiguration, service *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	baseName := az.GetLoadBalancerName(context.TODO(), "", service)
	return strings.HasPrefix(*fip.Name, baseName)
}
func (az *Cloud) getFrontendIPConfigName(service *v1.Service, subnetName *string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	baseName := az.GetLoadBalancerName(context.TODO(), "", service)
	if subnetName != nil {
		return fmt.Sprintf("%s-%s", baseName, *subnetName)
	}
	return baseName
}
func getNextAvailablePriority(rules []network.SecurityRule) (int32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var smallest int32 = loadBalancerMinimumPriority
	var spread int32 = 1
outer:
	for smallest < loadBalancerMaximumPriority {
		for _, rule := range rules {
			if *rule.Priority == smallest {
				smallest += spread
				continue outer
			}
		}
		return smallest, nil
	}
	return -1, fmt.Errorf("securityGroup priorities are exhausted")
}
func (az *Cloud) getIPForMachine(nodeName types.NodeName) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az.vmSet.GetIPByNodeName(string(nodeName))
}

var polyTable = crc32.MakeTable(crc32.Koopman)

func MakeCRC32(str string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	crc := crc32.New(polyTable)
	crc.Write([]byte(str))
	hash := crc.Sum32()
	return strconv.FormatUint(uint64(hash), 10)
}

type availabilitySet struct{ *Cloud }

func newAvailabilitySet(az *Cloud) VMSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &availabilitySet{Cloud: az}
}
func (as *availabilitySet) GetInstanceIDByNodeName(name string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var machine compute.VirtualMachine
	var err error
	machine, err = as.getVirtualMachine(types.NodeName(name))
	if err == cloudprovider.InstanceNotFound {
		return "", cloudprovider.InstanceNotFound
	}
	if err != nil {
		if as.CloudProviderBackoff {
			klog.V(2).Infof("GetInstanceIDByNodeName(%s) backing off", name)
			machine, err = as.GetVirtualMachineWithRetry(types.NodeName(name))
			if err != nil {
				klog.V(2).Infof("GetInstanceIDByNodeName(%s) abort backoff", name)
				return "", err
			}
		} else {
			return "", err
		}
	}
	return *machine.ID, nil
}
func (as *availabilitySet) GetPowerStatusByNodeName(name string) (powerState string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vm, err := as.getVirtualMachine(types.NodeName(name))
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
func (as *availabilitySet) GetNodeNameByProviderID(providerID string) (types.NodeName, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matches := providerIDRE.FindStringSubmatch(providerID)
	if len(matches) != 2 {
		return "", errors.New("error splitting providerID")
	}
	return types.NodeName(matches[1]), nil
}
func (as *availabilitySet) GetInstanceTypeByNodeName(name string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	machine, err := as.getVirtualMachine(types.NodeName(name))
	if err != nil {
		klog.Errorf("as.GetInstanceTypeByNodeName(%s) failed: as.getVirtualMachine(%s) err=%v", name, name, err)
		return "", err
	}
	return string(machine.HardwareProfile.VMSize), nil
}
func (as *availabilitySet) GetZoneByNodeName(name string) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vm, err := as.getVirtualMachine(types.NodeName(name))
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
		failureDomain = as.makeZone(zoneID)
	} else {
		failureDomain = strconv.Itoa(int(*vm.VirtualMachineProperties.InstanceView.PlatformFaultDomain))
	}
	zone := cloudprovider.Zone{FailureDomain: failureDomain, Region: *(vm.Location)}
	return zone, nil
}
func (as *availabilitySet) GetPrimaryVMSetName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return as.Config.PrimaryAvailabilitySetName
}
func (as *availabilitySet) GetIPByNodeName(name string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nic, err := as.GetPrimaryInterface(name)
	if err != nil {
		return "", "", err
	}
	ipConfig, err := getPrimaryIPConfig(nic)
	if err != nil {
		klog.Errorf("as.GetIPByNodeName(%s) failed: getPrimaryIPConfig(%v), err=%v", name, nic, err)
		return "", "", err
	}
	privateIP := *ipConfig.PrivateIPAddress
	publicIP := ""
	if ipConfig.PublicIPAddress != nil && ipConfig.PublicIPAddress.ID != nil {
		pipID := *ipConfig.PublicIPAddress.ID
		pipName, err := getLastSegment(pipID)
		if err != nil {
			return "", "", fmt.Errorf("failed to publicIP name for node %q with pipID %q", name, pipID)
		}
		pip, existsPip, err := as.getPublicIPAddress(as.ResourceGroup, pipName)
		if err != nil {
			return "", "", err
		}
		if existsPip {
			publicIP = *pip.IPAddress
		}
	}
	return privateIP, publicIP, nil
}
func (as *availabilitySet) getAgentPoolAvailabiliySets(nodes []*v1.Node) (agentPoolAvailabilitySets *[]string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vms, err := as.VirtualMachineClientListWithRetry(as.ResourceGroup)
	if err != nil {
		klog.Errorf("as.getNodeAvailabilitySet - VirtualMachineClientListWithRetry failed, err=%v", err)
		return nil, err
	}
	vmNameToAvailabilitySetID := make(map[string]string, len(vms))
	for vmx := range vms {
		vm := vms[vmx]
		if vm.AvailabilitySet != nil {
			vmNameToAvailabilitySetID[*vm.Name] = *vm.AvailabilitySet.ID
		}
	}
	availabilitySetIDs := sets.NewString()
	agentPoolAvailabilitySets = &[]string{}
	for nx := range nodes {
		nodeName := (*nodes[nx]).Name
		if isMasterNode(nodes[nx]) {
			continue
		}
		asID, ok := vmNameToAvailabilitySetID[nodeName]
		if !ok {
			klog.Errorf("as.getNodeAvailabilitySet - Node(%s) has no availability sets", nodeName)
			return nil, fmt.Errorf("Node (%s) - has no availability sets", nodeName)
		}
		if availabilitySetIDs.Has(asID) {
			continue
		}
		asName, err := getLastSegment(asID)
		if err != nil {
			klog.Errorf("as.getNodeAvailabilitySet - Node (%s)- getLastSegment(%s), err=%v", nodeName, asID, err)
			return nil, err
		}
		asName = strings.ToLower(asName)
		*agentPoolAvailabilitySets = append(*agentPoolAvailabilitySets, asName)
	}
	return agentPoolAvailabilitySets, nil
}
func (as *availabilitySet) GetVMSetNames(service *v1.Service, nodes []*v1.Node) (availabilitySetNames *[]string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hasMode, isAuto, serviceAvailabilitySetNames := getServiceLoadBalancerMode(service)
	if !hasMode {
		availabilitySetNames = &[]string{as.Config.PrimaryAvailabilitySetName}
		return availabilitySetNames, nil
	}
	availabilitySetNames, err = as.getAgentPoolAvailabiliySets(nodes)
	if err != nil {
		klog.Errorf("as.GetVMSetNames - getAgentPoolAvailabiliySets failed err=(%v)", err)
		return nil, err
	}
	if len(*availabilitySetNames) == 0 {
		klog.Errorf("as.GetVMSetNames - No availability sets found for nodes in the cluster, node count(%d)", len(nodes))
		return nil, fmt.Errorf("No availability sets found for nodes, node count(%d)", len(nodes))
	}
	sort.Strings(*availabilitySetNames)
	if !isAuto {
		if serviceAvailabilitySetNames == nil || len(serviceAvailabilitySetNames) == 0 {
			return nil, fmt.Errorf("service annotation for LoadBalancerMode is empty, it should have __auto__ or availability sets value")
		}
		var found bool
		for sasx := range serviceAvailabilitySetNames {
			for asx := range *availabilitySetNames {
				if strings.EqualFold((*availabilitySetNames)[asx], serviceAvailabilitySetNames[sasx]) {
					found = true
					serviceAvailabilitySetNames[sasx] = (*availabilitySetNames)[asx]
					break
				}
			}
			if !found {
				klog.Errorf("as.GetVMSetNames - Availability set (%s) in service annotation not found", serviceAvailabilitySetNames[sasx])
				return nil, fmt.Errorf("availability set (%s) - not found", serviceAvailabilitySetNames[sasx])
			}
		}
		availabilitySetNames = &serviceAvailabilitySetNames
	}
	return availabilitySetNames, nil
}
func (as *availabilitySet) GetPrimaryInterface(nodeName string) (network.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return as.getPrimaryInterfaceWithVMSet(nodeName, "")
}
func extractResourceGroupByNicID(nicID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matches := nicResourceGroupRE.FindStringSubmatch(nicID)
	if len(matches) != 2 {
		return "", fmt.Errorf("error of extracting resourceGroup from nicID %q", nicID)
	}
	return matches[1], nil
}
func extractResourceGroupByPipID(pipID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matches := publicIPResourceGroupRE.FindStringSubmatch(pipID)
	if len(matches) != 2 {
		return "", fmt.Errorf("error of extracting resourceGroup from pipID %q", pipID)
	}
	return matches[1], nil
}
func (as *availabilitySet) getPrimaryInterfaceWithVMSet(nodeName, vmSetName string) (network.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var machine compute.VirtualMachine
	machine, err := as.GetVirtualMachineWithRetry(types.NodeName(nodeName))
	if err != nil {
		klog.V(2).Infof("GetPrimaryInterface(%s, %s) abort backoff", nodeName, vmSetName)
		return network.Interface{}, err
	}
	primaryNicID, err := getPrimaryInterfaceID(machine)
	if err != nil {
		return network.Interface{}, err
	}
	nicName, err := getLastSegment(primaryNicID)
	if err != nil {
		return network.Interface{}, err
	}
	nodeResourceGroup, err := as.GetNodeResourceGroup(nodeName)
	if err != nil {
		return network.Interface{}, err
	}
	if vmSetName != "" && !as.useStandardLoadBalancer() {
		expectedAvailabilitySetName := as.getAvailabilitySetID(nodeResourceGroup, vmSetName)
		if machine.AvailabilitySet == nil || !strings.EqualFold(*machine.AvailabilitySet.ID, expectedAvailabilitySetName) {
			klog.V(3).Infof("GetPrimaryInterface: nic (%s) is not in the availabilitySet(%s)", nicName, vmSetName)
			return network.Interface{}, errNotInVMSet
		}
	}
	nicResourceGroup, err := extractResourceGroupByNicID(primaryNicID)
	if err != nil {
		return network.Interface{}, err
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	nic, err := as.InterfacesClient.Get(ctx, nicResourceGroup, nicName, "")
	if err != nil {
		return network.Interface{}, err
	}
	return nic, nil
}
func (as *availabilitySet) ensureHostInPool(service *v1.Service, nodeName types.NodeName, backendPoolID string, vmSetName string, isInternal bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmName := mapNodeNameToVMName(nodeName)
	serviceName := getServiceName(service)
	nic, err := as.getPrimaryInterfaceWithVMSet(vmName, vmSetName)
	if err != nil {
		if err == errNotInVMSet {
			klog.V(3).Infof("ensureHostInPool skips node %s because it is not in the vmSet %s", nodeName, vmSetName)
			return nil
		}
		klog.Errorf("error: az.ensureHostInPool(%s), az.vmSet.GetPrimaryInterface.Get(%s, %s), err=%v", nodeName, vmName, vmSetName, err)
		return err
	}
	if nic.ProvisioningState != nil && *nic.ProvisioningState == nicFailedState {
		klog.V(3).Infof("ensureHostInPool skips node %s because its primary nic %s is in Failed state", nodeName, *nic.Name)
		return nil
	}
	var primaryIPConfig *network.InterfaceIPConfiguration
	primaryIPConfig, err = getPrimaryIPConfig(nic)
	if err != nil {
		return err
	}
	foundPool := false
	newBackendPools := []network.BackendAddressPool{}
	if primaryIPConfig.LoadBalancerBackendAddressPools != nil {
		newBackendPools = *primaryIPConfig.LoadBalancerBackendAddressPools
	}
	for _, existingPool := range newBackendPools {
		if strings.EqualFold(backendPoolID, *existingPool.ID) {
			foundPool = true
			break
		}
	}
	if !foundPool {
		if as.useStandardLoadBalancer() && len(newBackendPools) > 0 {
			for _, pool := range newBackendPools {
				backendPool := *pool.ID
				matches := backendPoolIDRE.FindStringSubmatch(backendPool)
				if len(matches) == 2 {
					lbName := matches[1]
					if strings.HasSuffix(lbName, InternalLoadBalancerNameSuffix) == isInternal {
						klog.V(4).Infof("Node %q has already been added to LB %q, omit adding it to a new one", nodeName, lbName)
						return nil
					}
				}
			}
		}
		newBackendPools = append(newBackendPools, network.BackendAddressPool{ID: to.StringPtr(backendPoolID)})
		primaryIPConfig.LoadBalancerBackendAddressPools = &newBackendPools
		nicName := *nic.Name
		klog.V(3).Infof("nicupdate(%s): nic(%s) - updating", serviceName, nicName)
		ctx, cancel := getContextWithCancel()
		defer cancel()
		resp, err := as.InterfacesClient.CreateOrUpdate(ctx, as.ResourceGroup, *nic.Name, nic)
		klog.V(10).Infof("InterfacesClient.CreateOrUpdate(%q): end", *nic.Name)
		if as.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
			klog.V(2).Infof("nicupdate(%s) backing off: nic(%s) - updating, err=%v", serviceName, nicName, err)
			retryErr := as.CreateOrUpdateInterfaceWithRetry(service, nic)
			if retryErr != nil {
				err = retryErr
				klog.V(2).Infof("nicupdate(%s) abort backoff: nic(%s) - updating", serviceName, nicName)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (as *availabilitySet) EnsureHostsInPool(service *v1.Service, nodes []*v1.Node, backendPoolID string, vmSetName string, isInternal bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hostUpdates := make([]func() error, 0, len(nodes))
	for _, node := range nodes {
		localNodeName := node.Name
		if as.useStandardLoadBalancer() && as.excludeMasterNodesFromStandardLB() && isMasterNode(node) {
			klog.V(4).Infof("Excluding master node %q from load balancer backendpool %q", localNodeName, backendPoolID)
			continue
		}
		if as.ShouldNodeExcludedFromLoadBalancer(node) {
			klog.V(4).Infof("Excluding unmanaged/external-resource-group node %q", localNodeName)
			continue
		}
		f := func() error {
			err := as.ensureHostInPool(service, types.NodeName(localNodeName), backendPoolID, vmSetName, isInternal)
			if err != nil {
				return fmt.Errorf("ensure(%s): backendPoolID(%s) - failed to ensure host in pool: %q", getServiceName(service), backendPoolID, err)
			}
			return nil
		}
		hostUpdates = append(hostUpdates, f)
	}
	errs := utilerrors.AggregateGoroutines(hostUpdates...)
	if errs != nil {
		return utilerrors.Flatten(errs)
	}
	return nil
}
func (as *availabilitySet) EnsureBackendPoolDeleted(service *v1.Service, poolID, vmSetName string, backendAddressPools *[]network.BackendAddressPool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func generateStorageAccountName(accountNamePrefix string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	uniqueID := strings.Replace(string(uuid.NewUUID()), "-", "", -1)
	accountName := strings.ToLower(accountNamePrefix + uniqueID)
	if len(accountName) > storageAccountNameMaxLength {
		return accountName[:storageAccountNameMaxLength-1]
	}
	return accountName
}
