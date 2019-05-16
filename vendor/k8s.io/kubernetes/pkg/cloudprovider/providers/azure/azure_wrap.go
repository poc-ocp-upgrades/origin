package azure

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	vmCacheTTL            = time.Minute
	lbCacheTTL            = 2 * time.Minute
	nsgCacheTTL           = 2 * time.Minute
	rtCacheTTL            = 2 * time.Minute
	azureNodeProviderIDRE = regexp.MustCompile(`^azure:///subscriptions/(?:.*)/resourceGroups/(?:.*)/providers/Microsoft.Compute/(?:.*)`)
)

func checkResourceExistsFromError(err error) (bool, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return true, "", nil
	}
	v, ok := err.(autorest.DetailedError)
	if !ok {
		return false, "", err
	}
	if v.StatusCode == http.StatusNotFound {
		return false, err.Error(), nil
	}
	return false, "", v
}
func ignoreStatusNotFoundFromError(err error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return nil
	}
	v, ok := err.(autorest.DetailedError)
	if ok && v.StatusCode == http.StatusNotFound {
		return nil
	}
	return err
}
func (az *Cloud) getVirtualMachine(nodeName types.NodeName) (vm compute.VirtualMachine, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmName := string(nodeName)
	cachedVM, err := az.vmCache.Get(vmName)
	if err != nil {
		return vm, err
	}
	if cachedVM == nil {
		return vm, cloudprovider.InstanceNotFound
	}
	return *(cachedVM.(*compute.VirtualMachine)), nil
}
func (az *Cloud) getRouteTable() (routeTable network.RouteTable, exists bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cachedRt, err := az.rtCache.Get(az.RouteTableName)
	if err != nil {
		return routeTable, false, err
	}
	if cachedRt == nil {
		return routeTable, false, nil
	}
	return *(cachedRt.(*network.RouteTable)), true, nil
}
func (az *Cloud) getPublicIPAddress(pipResourceGroup string, pipName string) (pip network.PublicIPAddress, exists bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resourceGroup := az.ResourceGroup
	if pipResourceGroup != "" {
		resourceGroup = pipResourceGroup
	}
	var realErr error
	var message string
	ctx, cancel := getContextWithCancel()
	defer cancel()
	pip, err = az.PublicIPAddressesClient.Get(ctx, resourceGroup, pipName, "")
	exists, message, realErr = checkResourceExistsFromError(err)
	if realErr != nil {
		return pip, false, realErr
	}
	if !exists {
		klog.V(2).Infof("Public IP %q not found with message: %q", pipName, message)
		return pip, false, nil
	}
	return pip, exists, err
}
func (az *Cloud) getSubnet(virtualNetworkName string, subnetName string) (subnet network.Subnet, exists bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var realErr error
	var message string
	var rg string
	if len(az.VnetResourceGroup) > 0 {
		rg = az.VnetResourceGroup
	} else {
		rg = az.ResourceGroup
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	subnet, err = az.SubnetsClient.Get(ctx, rg, virtualNetworkName, subnetName, "")
	exists, message, realErr = checkResourceExistsFromError(err)
	if realErr != nil {
		return subnet, false, realErr
	}
	if !exists {
		klog.V(2).Infof("Subnet %q not found with message: %q", subnetName, message)
		return subnet, false, nil
	}
	return subnet, exists, err
}
func (az *Cloud) getAzureLoadBalancer(name string) (lb network.LoadBalancer, exists bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cachedLB, err := az.lbCache.Get(name)
	if err != nil {
		return lb, false, err
	}
	if cachedLB == nil {
		return lb, false, nil
	}
	return *(cachedLB.(*network.LoadBalancer)), true, nil
}
func (az *Cloud) getSecurityGroup() (nsg network.SecurityGroup, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.SecurityGroupName == "" {
		return nsg, fmt.Errorf("securityGroupName is not configured")
	}
	securityGroup, err := az.nsgCache.Get(az.SecurityGroupName)
	if err != nil {
		return nsg, err
	}
	if securityGroup == nil {
		return nsg, fmt.Errorf("nsg %q not found", az.SecurityGroupName)
	}
	return *(securityGroup.(*network.SecurityGroup)), nil
}
func (az *Cloud) newVMCache() (*timedCache, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getter := func(key string) (interface{}, error) {
		ctx, cancel := getContextWithCancel()
		defer cancel()
		resourceGroup, err := az.GetNodeResourceGroup(key)
		if err != nil {
			return nil, err
		}
		vm, err := az.VirtualMachinesClient.Get(ctx, resourceGroup, key, compute.InstanceView)
		exists, message, realErr := checkResourceExistsFromError(err)
		if realErr != nil {
			return nil, realErr
		}
		if !exists {
			klog.V(2).Infof("Virtual machine %q not found with message: %q", key, message)
			return nil, nil
		}
		return &vm, nil
	}
	return newTimedcache(vmCacheTTL, getter)
}
func (az *Cloud) newLBCache() (*timedCache, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getter := func(key string) (interface{}, error) {
		ctx, cancel := getContextWithCancel()
		defer cancel()
		lb, err := az.LoadBalancerClient.Get(ctx, az.ResourceGroup, key, "")
		exists, message, realErr := checkResourceExistsFromError(err)
		if realErr != nil {
			return nil, realErr
		}
		if !exists {
			klog.V(2).Infof("Load balancer %q not found with message: %q", key, message)
			return nil, nil
		}
		return &lb, nil
	}
	return newTimedcache(lbCacheTTL, getter)
}
func (az *Cloud) newNSGCache() (*timedCache, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getter := func(key string) (interface{}, error) {
		ctx, cancel := getContextWithCancel()
		defer cancel()
		nsg, err := az.SecurityGroupsClient.Get(ctx, az.ResourceGroup, key, "")
		exists, message, realErr := checkResourceExistsFromError(err)
		if realErr != nil {
			return nil, realErr
		}
		if !exists {
			klog.V(2).Infof("Security group %q not found with message: %q", key, message)
			return nil, nil
		}
		return &nsg, nil
	}
	return newTimedcache(nsgCacheTTL, getter)
}
func (az *Cloud) newRouteTableCache() (*timedCache, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getter := func(key string) (interface{}, error) {
		ctx, cancel := getContextWithCancel()
		defer cancel()
		rt, err := az.RouteTablesClient.Get(ctx, az.ResourceGroup, key, "")
		exists, message, realErr := checkResourceExistsFromError(err)
		if realErr != nil {
			return nil, realErr
		}
		if !exists {
			klog.V(2).Infof("Route table %q not found with message: %q", key, message)
			return nil, nil
		}
		return &rt, nil
	}
	return newTimedcache(rtCacheTTL, getter)
}
func (az *Cloud) useStandardLoadBalancer() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.EqualFold(az.LoadBalancerSku, loadBalancerSkuStandard)
}
func (az *Cloud) excludeMasterNodesFromStandardLB() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az.ExcludeMasterFromStandardLB != nil && *az.ExcludeMasterFromStandardLB
}
func (az *Cloud) IsNodeUnmanaged(nodeName string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	unmanagedNodes, err := az.GetUnmanagedNodes()
	if err != nil {
		return false, err
	}
	return unmanagedNodes.Has(nodeName), nil
}
func (az *Cloud) IsNodeUnmanagedByProviderID(providerID string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !azureNodeProviderIDRE.Match([]byte(providerID))
}
