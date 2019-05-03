package azure

import (
 "fmt"
 "time"
 "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/types"
 kwait "k8s.io/apimachinery/pkg/util/wait"
 cloudprovider "k8s.io/cloud-provider"
)

const (
 storageAccountNameTemplate             = "pvc%s"
 maxStorageAccounts                     = 100
 maxDisksPerStorageAccounts             = 60
 storageAccountUtilizationBeforeGrowing = 0.5
 maxLUN                                 = 64
 errLeaseFailed                         = "AcquireDiskLeaseFailed"
 errLeaseIDMissing                      = "LeaseIdMissing"
 errContainerNotFound                   = "ContainerNotFound"
 errDiskBlobNotFound                    = "DiskBlobNotFound"
)

var defaultBackOff = kwait.Backoff{Steps: 20, Duration: 2 * time.Second, Factor: 1.5, Jitter: 0.0}

type controllerCommon struct {
 subscriptionID        string
 location              string
 storageEndpointSuffix string
 resourceGroup         string
 cloud                 *Cloud
}

func (c *controllerCommon) getNodeVMSet(nodeName types.NodeName) (VMSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c.cloud.VMType == vmTypeStandard {
  return c.cloud.vmSet, nil
 }
 ss, ok := c.cloud.vmSet.(*scaleSet)
 if !ok {
  return nil, fmt.Errorf("error of converting vmSet (%q) to scaleSet with vmType %q", c.cloud.vmSet, c.cloud.VMType)
 }
 managedByAS, err := ss.isNodeManagedByAvailabilitySet(mapNodeNameToVMName(nodeName))
 if err != nil {
  return nil, err
 }
 if managedByAS {
  return ss.availabilitySet, nil
 }
 return ss, nil
}
func (c *controllerCommon) AttachDisk(isManagedDisk bool, diskName, diskURI string, nodeName types.NodeName, lun int32, cachingMode compute.CachingTypes) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmset, err := c.getNodeVMSet(nodeName)
 if err != nil {
  return err
 }
 return vmset.AttachDisk(isManagedDisk, diskName, diskURI, nodeName, lun, cachingMode)
}
func (c *controllerCommon) DetachDiskByName(diskName, diskURI string, nodeName types.NodeName) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmset, err := c.getNodeVMSet(nodeName)
 if err != nil {
  return err
 }
 return vmset.DetachDiskByName(diskName, diskURI, nodeName)
}
func (c *controllerCommon) getNodeDataDisks(nodeName types.NodeName) ([]compute.DataDisk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmset, err := c.getNodeVMSet(nodeName)
 if err != nil {
  return nil, err
 }
 return vmset.GetDataDisks(nodeName)
}
func (c *controllerCommon) GetDiskLun(diskName, diskURI string, nodeName types.NodeName) (int32, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disks, err := c.getNodeDataDisks(nodeName)
 if err != nil {
  klog.Errorf("error of getting data disks for node %q: %v", nodeName, err)
  return -1, err
 }
 for _, disk := range disks {
  if disk.Lun != nil && (disk.Name != nil && diskName != "" && *disk.Name == diskName) || (disk.Vhd != nil && disk.Vhd.URI != nil && diskURI != "" && *disk.Vhd.URI == diskURI) || (disk.ManagedDisk != nil && *disk.ManagedDisk.ID == diskURI) {
   klog.V(2).Infof("azureDisk - find disk: lun %d name %q uri %q", *disk.Lun, diskName, diskURI)
   return *disk.Lun, nil
  }
 }
 return -1, fmt.Errorf("Cannot find Lun for disk %s", diskName)
}
func (c *controllerCommon) GetNextDiskLun(nodeName types.NodeName) (int32, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disks, err := c.getNodeDataDisks(nodeName)
 if err != nil {
  klog.Errorf("error of getting data disks for node %q: %v", nodeName, err)
  return -1, err
 }
 used := make([]bool, maxLUN)
 for _, disk := range disks {
  if disk.Lun != nil {
   used[*disk.Lun] = true
  }
 }
 for k, v := range used {
  if !v {
   return int32(k), nil
  }
 }
 return -1, fmt.Errorf("all luns are used")
}
func (c *controllerCommon) DisksAreAttached(diskNames []string, nodeName types.NodeName) (map[string]bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 attached := make(map[string]bool)
 for _, diskName := range diskNames {
  attached[diskName] = false
 }
 disks, err := c.getNodeDataDisks(nodeName)
 if err != nil {
  if err == cloudprovider.InstanceNotFound {
   klog.Warningf("azureDisk - Cannot find node %q, DisksAreAttached will assume disks %v are not attached to it.", nodeName, diskNames)
   return attached, nil
  }
  return attached, err
 }
 for _, disk := range disks {
  for _, diskName := range diskNames {
   if disk.Name != nil && diskName != "" && *disk.Name == diskName {
    attached[diskName] = true
   }
  }
 }
 return attached, nil
}
