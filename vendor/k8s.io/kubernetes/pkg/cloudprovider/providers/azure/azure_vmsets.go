package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type VMSet interface {
	GetInstanceIDByNodeName(name string) (string, error)
	GetInstanceTypeByNodeName(name string) (string, error)
	GetIPByNodeName(name string) (string, string, error)
	GetPrimaryInterface(nodeName string) (network.Interface, error)
	GetNodeNameByProviderID(providerID string) (types.NodeName, error)
	GetZoneByNodeName(name string) (cloudprovider.Zone, error)
	GetPrimaryVMSetName() string
	GetVMSetNames(service *v1.Service, nodes []*v1.Node) (availabilitySetNames *[]string, err error)
	EnsureHostsInPool(service *v1.Service, nodes []*v1.Node, backendPoolID string, vmSetName string, isInternal bool) error
	EnsureBackendPoolDeleted(service *v1.Service, poolID, vmSetName string, backendAddressPools *[]network.BackendAddressPool) error
	AttachDisk(isManagedDisk bool, diskName, diskURI string, nodeName types.NodeName, lun int32, cachingMode compute.CachingTypes) error
	DetachDiskByName(diskName, diskURI string, nodeName types.NodeName) error
	GetDataDisks(nodeName types.NodeName) ([]compute.DataDisk, error)
	GetPowerStatusByNodeName(name string) (string, error)
}
