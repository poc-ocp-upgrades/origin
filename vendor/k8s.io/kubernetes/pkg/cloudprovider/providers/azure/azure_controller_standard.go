package azure

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"strings"
)

func (as *availabilitySet) AttachDisk(isManagedDisk bool, diskName, diskURI string, nodeName types.NodeName, lun int32, cachingMode compute.CachingTypes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vm, err := as.getVirtualMachine(nodeName)
	if err != nil {
		return err
	}
	vmName := mapNodeNameToVMName(nodeName)
	nodeResourceGroup, err := as.GetNodeResourceGroup(vmName)
	if err != nil {
		return err
	}
	disks := *vm.StorageProfile.DataDisks
	if isManagedDisk {
		disks = append(disks, compute.DataDisk{Name: &diskName, Lun: &lun, Caching: cachingMode, CreateOption: "attach", ManagedDisk: &compute.ManagedDiskParameters{ID: &diskURI}})
	} else {
		disks = append(disks, compute.DataDisk{Name: &diskName, Vhd: &compute.VirtualHardDisk{URI: &diskURI}, Lun: &lun, Caching: cachingMode, CreateOption: "attach"})
	}
	newVM := compute.VirtualMachine{Location: vm.Location, VirtualMachineProperties: &compute.VirtualMachineProperties{HardwareProfile: vm.HardwareProfile, StorageProfile: &compute.StorageProfile{DataDisks: &disks}}}
	klog.V(2).Infof("azureDisk - update(%s): vm(%s) - attach disk(%s)", nodeResourceGroup, vmName, diskName)
	ctx, cancel := getContextWithCancel()
	defer cancel()
	defer as.cloud.vmCache.Delete(vmName)
	_, err = as.VirtualMachinesClient.CreateOrUpdate(ctx, nodeResourceGroup, vmName, newVM)
	if err != nil {
		klog.Errorf("azureDisk - attach disk(%s) failed, err: %v", diskName, err)
		detail := err.Error()
		if strings.Contains(detail, errLeaseFailed) || strings.Contains(detail, errDiskBlobNotFound) {
			klog.V(2).Infof("azureDisk - err %v, try detach disk(%s)", err, diskName)
			as.DetachDiskByName(diskName, diskURI, nodeName)
		}
	} else {
		klog.V(2).Infof("azureDisk - attach disk(%s) succeeded", diskName)
	}
	return err
}
func (as *availabilitySet) DetachDiskByName(diskName, diskURI string, nodeName types.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vm, err := as.getVirtualMachine(nodeName)
	if err != nil {
		klog.Warningf("azureDisk - cannot find node %s, skip detaching disk %s", nodeName, diskName)
		return nil
	}
	vmName := mapNodeNameToVMName(nodeName)
	nodeResourceGroup, err := as.GetNodeResourceGroup(vmName)
	if err != nil {
		return err
	}
	disks := *vm.StorageProfile.DataDisks
	bFoundDisk := false
	for i, disk := range disks {
		if disk.Lun != nil && (disk.Name != nil && diskName != "" && *disk.Name == diskName) || (disk.Vhd != nil && disk.Vhd.URI != nil && diskURI != "" && *disk.Vhd.URI == diskURI) || (disk.ManagedDisk != nil && diskURI != "" && *disk.ManagedDisk.ID == diskURI) {
			klog.V(2).Infof("azureDisk - detach disk: name %q uri %q", diskName, diskURI)
			disks = append(disks[:i], disks[i+1:]...)
			bFoundDisk = true
			break
		}
	}
	if !bFoundDisk {
		return fmt.Errorf("detach azure disk failure, disk %s not found, diskURI: %s", diskName, diskURI)
	}
	newVM := compute.VirtualMachine{Location: vm.Location, VirtualMachineProperties: &compute.VirtualMachineProperties{HardwareProfile: vm.HardwareProfile, StorageProfile: &compute.StorageProfile{DataDisks: &disks}}}
	klog.V(2).Infof("azureDisk - update(%s): vm(%s) - detach disk(%s)", nodeResourceGroup, vmName, diskName)
	ctx, cancel := getContextWithCancel()
	defer cancel()
	defer as.cloud.vmCache.Delete(vmName)
	resp, err := as.VirtualMachinesClient.CreateOrUpdate(ctx, nodeResourceGroup, vmName, newVM)
	if as.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
		klog.V(2).Infof("azureDisk - update(%s) backing off: vm(%s) detach disk(%s, %s), err: %v", nodeResourceGroup, vmName, diskName, diskURI, err)
		retryErr := as.CreateOrUpdateVMWithRetry(nodeResourceGroup, vmName, newVM)
		if retryErr != nil {
			err = retryErr
			klog.V(2).Infof("azureDisk - update(%s) abort backoff: vm(%s) detach disk(%s, %s), err: %v", nodeResourceGroup, vmName, diskName, diskURI, err)
		}
	}
	if err != nil {
		klog.Errorf("azureDisk - detach disk(%s, %s)) failed, err: %v", diskName, diskURI, err)
	} else {
		klog.V(2).Infof("azureDisk - detach disk(%s, %s) succeeded", diskName, diskURI)
	}
	return err
}
func (as *availabilitySet) GetDataDisks(nodeName types.NodeName) ([]compute.DataDisk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vm, err := as.getVirtualMachine(nodeName)
	if err != nil {
		return nil, err
	}
	if vm.StorageProfile.DataDisks == nil {
		return nil, nil
	}
	return *vm.StorageProfile.DataDisks, nil
}
