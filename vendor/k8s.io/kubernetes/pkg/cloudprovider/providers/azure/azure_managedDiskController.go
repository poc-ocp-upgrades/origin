package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	kwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	"path"
	"strconv"
	"strings"
)

const (
	defaultDiskIOPSReadWrite = 500
	defaultDiskMBpsReadWrite = 100
)

type ManagedDiskController struct{ common *controllerCommon }
type ManagedDiskOptions struct {
	DiskName           string
	SizeGB             int
	PVCName            string
	ResourceGroup      string
	AvailabilityZone   string
	Tags               map[string]string
	StorageAccountType compute.DiskStorageAccountTypes
	DiskIOPSReadWrite  string
	DiskMBpsReadWrite  string
}

func (c *ManagedDiskController) CreateManagedDisk(options *ManagedDiskOptions) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	klog.V(4).Infof("azureDisk - creating new managed Name:%s StorageAccountType:%s Size:%v", options.DiskName, options.StorageAccountType, options.SizeGB)
	var createZones *[]string
	if len(options.AvailabilityZone) > 0 {
		zoneList := []string{c.common.cloud.GetZoneID(options.AvailabilityZone)}
		createZones = &zoneList
	}
	newTags := make(map[string]*string)
	azureDDTag := "kubernetes-azure-dd"
	newTags["created-by"] = &azureDDTag
	if options.Tags != nil {
		for k, v := range options.Tags {
			newKey := strings.Replace(k, "/", "-", -1)
			newValue := strings.Replace(v, "/", "-", -1)
			newTags[newKey] = &newValue
		}
	}
	diskSizeGB := int32(options.SizeGB)
	diskSku := compute.DiskStorageAccountTypes(options.StorageAccountType)
	diskProperties := compute.DiskProperties{DiskSizeGB: &diskSizeGB, CreationData: &compute.CreationData{CreateOption: compute.Empty}}
	if diskSku == compute.UltraSSDLRS {
		diskIOPSReadWrite := int64(defaultDiskIOPSReadWrite)
		if options.DiskIOPSReadWrite != "" {
			v, err := strconv.Atoi(options.DiskIOPSReadWrite)
			if err != nil {
				return "", fmt.Errorf("AzureDisk - failed to parse DiskIOPSReadWrite: %v", err)
			}
			diskIOPSReadWrite = int64(v)
		}
		diskProperties.DiskIOPSReadWrite = to.Int64Ptr(diskIOPSReadWrite)
		diskMBpsReadWrite := int32(defaultDiskMBpsReadWrite)
		if options.DiskMBpsReadWrite != "" {
			v, err := strconv.Atoi(options.DiskMBpsReadWrite)
			if err != nil {
				return "", fmt.Errorf("AzureDisk - failed to parse DiskMBpsReadWrite: %v", err)
			}
			diskMBpsReadWrite = int32(v)
		}
		diskProperties.DiskMBpsReadWrite = to.Int32Ptr(diskMBpsReadWrite)
	} else {
		if options.DiskIOPSReadWrite != "" {
			return "", fmt.Errorf("AzureDisk - DiskIOPSReadWrite parameter is only applicable in UltraSSD_LRS disk type")
		}
		if options.DiskMBpsReadWrite != "" {
			return "", fmt.Errorf("AzureDisk - DiskMBpsReadWrite parameter is only applicable in UltraSSD_LRS disk type")
		}
	}
	model := compute.Disk{Location: &c.common.location, Tags: newTags, Zones: createZones, Sku: &compute.DiskSku{Name: diskSku}, DiskProperties: &diskProperties}
	if options.ResourceGroup == "" {
		options.ResourceGroup = c.common.resourceGroup
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	_, err = c.common.cloud.DisksClient.CreateOrUpdate(ctx, options.ResourceGroup, options.DiskName, model)
	if err != nil {
		return "", err
	}
	diskID := ""
	err = kwait.ExponentialBackoff(defaultBackOff, func() (bool, error) {
		provisionState, id, err := c.getDisk(options.ResourceGroup, options.DiskName)
		diskID = id
		if err != nil {
			return false, err
		}
		if strings.ToLower(provisionState) == "succeeded" {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		klog.V(2).Infof("azureDisk - created new MD Name:%s StorageAccountType:%s Size:%v but was unable to confirm provisioningState in poll process", options.DiskName, options.StorageAccountType, options.SizeGB)
	} else {
		klog.V(2).Infof("azureDisk - created new MD Name:%s StorageAccountType:%s Size:%v", options.DiskName, options.StorageAccountType, options.SizeGB)
	}
	return diskID, nil
}
func (c *ManagedDiskController) DeleteManagedDisk(diskURI string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	diskName := path.Base(diskURI)
	resourceGroup, err := getResourceGroupFromDiskURI(diskURI)
	if err != nil {
		return err
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	_, err = c.common.cloud.DisksClient.Delete(ctx, resourceGroup, diskName)
	if err != nil {
		return err
	}
	klog.V(2).Infof("azureDisk - deleted a managed disk: %s", diskURI)
	return nil
}
func (c *ManagedDiskController) getDisk(resourceGroup, diskName string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := getContextWithCancel()
	defer cancel()
	result, err := c.common.cloud.DisksClient.Get(ctx, resourceGroup, diskName)
	if err != nil {
		return "", "", err
	}
	if result.DiskProperties != nil && (*result.DiskProperties).ProvisioningState != nil {
		return *(*result.DiskProperties).ProvisioningState, *result.ID, nil
	}
	return "", "", err
}
func (c *ManagedDiskController) ResizeDisk(diskURI string, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := getContextWithCancel()
	defer cancel()
	diskName := path.Base(diskURI)
	resourceGroup, err := getResourceGroupFromDiskURI(diskURI)
	if err != nil {
		return oldSize, err
	}
	result, err := c.common.cloud.DisksClient.Get(ctx, resourceGroup, diskName)
	if err != nil {
		return oldSize, err
	}
	if result.DiskProperties == nil || result.DiskProperties.DiskSizeGB == nil {
		return oldSize, fmt.Errorf("DiskProperties of disk(%s) is nil", diskName)
	}
	requestBytes := newSize.Value()
	requestGiB := int32(util.RoundUpSize(requestBytes, 1024*1024*1024))
	newSizeQuant := resource.MustParse(fmt.Sprintf("%dGi", requestGiB))
	klog.V(2).Infof("azureDisk - begin to resize disk(%s) with new size(%d), old size(%v)", diskName, requestGiB, oldSize)
	if *result.DiskProperties.DiskSizeGB >= requestGiB {
		return newSizeQuant, nil
	}
	result.DiskProperties.DiskSizeGB = &requestGiB
	ctx, cancel = getContextWithCancel()
	defer cancel()
	if _, err := c.common.cloud.DisksClient.CreateOrUpdate(ctx, resourceGroup, diskName, result); err != nil {
		return oldSize, err
	}
	klog.V(2).Infof("azureDisk - resize disk(%s) with new size(%d) completed", diskName, requestGiB)
	return newSizeQuant, nil
}
func getResourceGroupFromDiskURI(diskURI string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fields := strings.Split(diskURI, "/")
	if len(fields) != 9 || fields[3] != "resourceGroups" {
		return "", fmt.Errorf("invalid disk URI: %s", diskURI)
	}
	return fields[4], nil
}
func (c *Cloud) GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.AzureDisk == nil {
		return nil, nil
	}
	if pv.Spec.AzureDisk.DiskName == volume.ProvisionedVolumeName {
		return nil, nil
	}
	return c.GetAzureDiskLabels(pv.Spec.AzureDisk.DataDiskURI)
}
func (c *Cloud) GetAzureDiskLabels(diskURI string) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	diskName := path.Base(diskURI)
	resourceGroup, err := getResourceGroupFromDiskURI(diskURI)
	if err != nil {
		klog.Errorf("Failed to get resource group for AzureDisk %q: %v", diskName, err)
		return nil, err
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	disk, err := c.DisksClient.Get(ctx, resourceGroup, diskName)
	if err != nil {
		klog.Errorf("Failed to get information for AzureDisk %q: %v", diskName, err)
		return nil, err
	}
	if disk.Zones == nil || len(*disk.Zones) == 0 {
		klog.V(4).Infof("Azure disk %q is not zoned", diskName)
		return nil, nil
	}
	zones := *disk.Zones
	zoneID, err := strconv.Atoi(zones[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse zone %v for AzureDisk %v: %v", zones, diskName, err)
	}
	zone := c.makeZone(zoneID)
	klog.V(4).Infof("Got zone %q for Azure disk %q", zone, diskName)
	labels := map[string]string{kubeletapis.LabelZoneRegion: c.Location, kubeletapis.LabelZoneFailureDomain: zone}
	return labels, nil
}
