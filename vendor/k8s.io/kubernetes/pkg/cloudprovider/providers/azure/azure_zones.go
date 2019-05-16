package azure

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"strconv"
	"strings"
)

func (az *Cloud) makeZone(zoneID int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s-%d", strings.ToLower(az.Location), zoneID)
}
func (az *Cloud) isAvailabilityZone(zone string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasPrefix(zone, fmt.Sprintf("%s-", az.Location))
}
func (az *Cloud) GetZoneID(zoneLabel string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !az.isAvailabilityZone(zoneLabel) {
		return ""
	}
	return strings.TrimPrefix(zoneLabel, fmt.Sprintf("%s-", az.Location))
}
func (az *Cloud) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	metadata, err := az.metadata.GetMetadata()
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	if metadata.Compute == nil {
		return cloudprovider.Zone{}, fmt.Errorf("failure of getting compute information from instance metadata")
	}
	zone := ""
	if metadata.Compute.Zone != "" {
		zoneID, err := strconv.Atoi(metadata.Compute.Zone)
		if err != nil {
			return cloudprovider.Zone{}, fmt.Errorf("failed to parse zone ID %q: %v", metadata.Compute.Zone, err)
		}
		zone = az.makeZone(zoneID)
	} else {
		klog.V(3).Infof("Availability zone is not enabled for the node, falling back to fault domain")
		zone = metadata.Compute.FaultDomain
	}
	return cloudprovider.Zone{FailureDomain: zone, Region: az.Location}, nil
}
func (az *Cloud) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.IsNodeUnmanagedByProviderID(providerID) {
		klog.V(2).Infof("GetZoneByProviderID: omitting unmanaged node %q", providerID)
		return cloudprovider.Zone{}, nil
	}
	nodeName, err := az.vmSet.GetNodeNameByProviderID(providerID)
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	return az.GetZoneByNodeName(ctx, nodeName)
}
func (az *Cloud) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	unmanaged, err := az.IsNodeUnmanaged(string(nodeName))
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	if unmanaged {
		klog.V(2).Infof("GetZoneByNodeName: omitting unmanaged node %q", nodeName)
		return cloudprovider.Zone{}, nil
	}
	return az.vmSet.GetZoneByNodeName(string(nodeName))
}
