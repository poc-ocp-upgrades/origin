package diskmanagers

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"hash/fnv"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
	"strings"
)

type vmDiskManager struct {
	diskPath      string
	volumeOptions *vclib.VolumeOptions
	vmOptions     *vclib.VMOptions
}

func (vmdisk vmDiskManager) Create(ctx context.Context, datastore *vclib.Datastore) (canonicalDiskPath string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if vmdisk.volumeOptions.SCSIControllerType == "" {
		vmdisk.volumeOptions.SCSIControllerType = vclib.PVSCSIControllerType
	}
	pbmClient, err := vclib.NewPbmClient(ctx, datastore.Client())
	if err != nil {
		klog.Errorf("Error occurred while creating new pbmClient, err: %+v", err)
		return "", err
	}
	if vmdisk.volumeOptions.StoragePolicyID == "" && vmdisk.volumeOptions.StoragePolicyName != "" {
		vmdisk.volumeOptions.StoragePolicyID, err = pbmClient.ProfileIDByName(ctx, vmdisk.volumeOptions.StoragePolicyName)
		if err != nil {
			klog.Errorf("Error occurred while getting Profile Id from Profile Name: %s, err: %+v", vmdisk.volumeOptions.StoragePolicyName, err)
			return "", err
		}
	}
	if vmdisk.volumeOptions.StoragePolicyID != "" {
		compatible, faultMessage, err := datastore.IsCompatibleWithStoragePolicy(ctx, vmdisk.volumeOptions.StoragePolicyID)
		if err != nil {
			klog.Errorf("Error occurred while checking datastore compatibility with storage policy id: %s, err: %+v", vmdisk.volumeOptions.StoragePolicyID, err)
			return "", err
		}
		if !compatible {
			klog.Errorf("Datastore: %s is not compatible with Policy: %s", datastore.Name(), vmdisk.volumeOptions.StoragePolicyName)
			return "", fmt.Errorf("User specified datastore is not compatible with the storagePolicy: %q. Failed with faults: %+q", vmdisk.volumeOptions.StoragePolicyName, faultMessage)
		}
	}
	storageProfileSpec := &types.VirtualMachineDefinedProfileSpec{}
	if vmdisk.volumeOptions.StoragePolicyID != "" {
		storageProfileSpec.ProfileId = vmdisk.volumeOptions.StoragePolicyID
	} else if vmdisk.volumeOptions.VSANStorageProfileData != "" {
		dsType, err := datastore.GetType(ctx)
		if err != nil {
			return "", err
		}
		if dsType != vclib.VSANDatastoreType {
			klog.Errorf("The specified datastore: %q is not a VSAN datastore", datastore.Name())
			return "", fmt.Errorf("The specified datastore: %q is not a VSAN datastore."+" The policy parameters will work only with VSAN Datastore."+" So, please specify a valid VSAN datastore in Storage class definition.", datastore.Name())
		}
		storageProfileSpec.ProfileId = ""
		storageProfileSpec.ProfileData = &types.VirtualMachineProfileRawData{ExtensionKey: "com.vmware.vim.sps", ObjectData: vmdisk.volumeOptions.VSANStorageProfileData}
	} else {
		klog.Errorf("Both volumeOptions.StoragePolicyID and volumeOptions.VSANStorageProfileData are not set. One of them should be set")
		return "", fmt.Errorf("Both volumeOptions.StoragePolicyID and volumeOptions.VSANStorageProfileData are not set. One of them should be set")
	}
	var dummyVM *vclib.VirtualMachine
	fnvHash := fnv.New32a()
	fnvHash.Write([]byte(vmdisk.volumeOptions.Name))
	dummyVMFullName := vclib.DummyVMPrefixName + "-" + fmt.Sprint(fnvHash.Sum32())
	dummyVM, err = datastore.Datacenter.GetVMByPath(ctx, vmdisk.vmOptions.VMFolder.InventoryPath+"/"+dummyVMFullName)
	if err != nil {
		klog.V(1).Infof("Creating Dummy VM: %q", dummyVMFullName)
		dummyVM, err = vmdisk.createDummyVM(ctx, datastore.Datacenter, dummyVMFullName)
		if err != nil {
			klog.Errorf("Failed to create Dummy VM. err: %v", err)
			return "", err
		}
	}
	virtualMachineConfigSpec := types.VirtualMachineConfigSpec{}
	disk, _, err := dummyVM.CreateDiskSpec(ctx, vmdisk.diskPath, datastore, vmdisk.volumeOptions)
	if err != nil {
		klog.Errorf("Failed to create Disk Spec. err: %v", err)
		return "", err
	}
	deviceConfigSpec := &types.VirtualDeviceConfigSpec{Device: disk, Operation: types.VirtualDeviceConfigSpecOperationAdd, FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate}
	deviceConfigSpec.Profile = append(deviceConfigSpec.Profile, storageProfileSpec)
	virtualMachineConfigSpec.DeviceChange = append(virtualMachineConfigSpec.DeviceChange, deviceConfigSpec)
	fileAlreadyExist := false
	task, err := dummyVM.Reconfigure(ctx, virtualMachineConfigSpec)
	if err != nil {
		klog.Errorf("Failed to reconfig. err: %v", err)
		return "", err
	}
	err = task.Wait(ctx)
	if err != nil {
		fileAlreadyExist = isAlreadyExists(vmdisk.diskPath, err)
		if fileAlreadyExist {
			klog.V(vclib.LogLevel).Infof("File: %v already exists", vmdisk.diskPath)
		} else {
			klog.Errorf("Failed to attach the disk to VM: %q with err: %+v", dummyVMFullName, err)
			return "", err
		}
	}
	err = dummyVM.DetachDisk(ctx, vmdisk.diskPath)
	if err != nil {
		if vclib.DiskNotFoundErrMsg == err.Error() && fileAlreadyExist {
			klog.V(vclib.LogLevel).Infof("File: %v is already detached", vmdisk.diskPath)
		} else {
			klog.Errorf("Failed to detach the disk: %q from VM: %q with err: %+v", vmdisk.diskPath, dummyVMFullName, err)
			return "", err
		}
	}
	err = dummyVM.DeleteVM(ctx)
	if err != nil {
		klog.Errorf("Failed to destroy the vm: %q with err: %+v", dummyVMFullName, err)
	}
	return vmdisk.diskPath, nil
}
func (vmdisk vmDiskManager) Delete(ctx context.Context, datacenter *vclib.Datacenter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Errorf("vmDiskManager.Delete is not supported")
}
func (vmdisk vmDiskManager) createDummyVM(ctx context.Context, datacenter *vclib.Datacenter, vmName string) (*vclib.VirtualMachine, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	virtualMachineConfigSpec := types.VirtualMachineConfigSpec{Name: vmName, Files: &types.VirtualMachineFileInfo{VmPathName: "[" + vmdisk.volumeOptions.Datastore + "]"}, NumCPUs: 1, MemoryMB: 4, DeviceChange: []types.BaseVirtualDeviceConfigSpec{&types.VirtualDeviceConfigSpec{Operation: types.VirtualDeviceConfigSpecOperationAdd, Device: &types.ParaVirtualSCSIController{VirtualSCSIController: types.VirtualSCSIController{SharedBus: types.VirtualSCSISharingNoSharing, VirtualController: types.VirtualController{BusNumber: 0, VirtualDevice: types.VirtualDevice{Key: 1000}}}}}}}
	task, err := vmdisk.vmOptions.VMFolder.CreateVM(ctx, virtualMachineConfigSpec, vmdisk.vmOptions.VMResourcePool, nil)
	if err != nil {
		klog.Errorf("Failed to create VM. err: %+v", err)
		return nil, err
	}
	dummyVMTaskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		klog.Errorf("Error occurred while waiting for create VM task result. err: %+v", err)
		return nil, err
	}
	vmRef := dummyVMTaskInfo.Result.(object.Reference)
	dummyVM := object.NewVirtualMachine(datacenter.Client(), vmRef.Reference())
	return &vclib.VirtualMachine{VirtualMachine: dummyVM, Datacenter: datacenter}, nil
}
func CleanUpDummyVMs(ctx context.Context, folder *vclib.Folder, dc *vclib.Datacenter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmList, err := folder.GetVirtualMachines(ctx)
	if err != nil {
		klog.V(4).Infof("Failed to get virtual machines in the kubernetes cluster: %s, err: %+v", folder.InventoryPath, err)
		return err
	}
	if vmList == nil || len(vmList) == 0 {
		klog.Errorf("No virtual machines found in the kubernetes cluster: %s", folder.InventoryPath)
		return fmt.Errorf("No virtual machines found in the kubernetes cluster: %s", folder.InventoryPath)
	}
	var dummyVMList []*vclib.VirtualMachine
	for _, vm := range vmList {
		vmName, err := vm.ObjectName(ctx)
		if err != nil {
			klog.V(4).Infof("Unable to get name from VM with err: %+v", err)
			continue
		}
		if strings.HasPrefix(vmName, vclib.DummyVMPrefixName) {
			vmObj := vclib.VirtualMachine{VirtualMachine: object.NewVirtualMachine(dc.Client(), vm.Reference()), Datacenter: dc}
			dummyVMList = append(dummyVMList, &vmObj)
		}
	}
	for _, vm := range dummyVMList {
		err = vm.DeleteVM(ctx)
		if err != nil {
			klog.V(4).Infof("Unable to delete dummy VM with err: %+v", err)
			continue
		}
	}
	return nil
}
func isAlreadyExists(path string, err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errorMessage := fmt.Sprintf("Cannot complete the operation because the file or folder %s already exists", path)
	if errorMessage == err.Error() {
		return true
	}
	return false
}
