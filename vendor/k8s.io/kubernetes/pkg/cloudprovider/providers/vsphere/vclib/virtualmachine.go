package vclib

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"k8s.io/klog"
	"strings"
	"time"
)

type VirtualMachine struct {
	*object.VirtualMachine
	Datacenter *Datacenter
}

func (vm *VirtualMachine) IsDiskAttached(ctx context.Context, diskPath string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	device, err := vm.getVirtualDeviceByPath(ctx, diskPath)
	if err != nil {
		return false, err
	}
	if device != nil {
		return true, nil
	}
	return false, nil
}
func (vm *VirtualMachine) DeleteVM(ctx context.Context) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	destroyTask, err := vm.Destroy(ctx)
	if err != nil {
		klog.Errorf("Failed to delete the VM: %q. err: %+v", vm.InventoryPath, err)
		return err
	}
	return destroyTask.Wait(ctx)
}
func (vm *VirtualMachine) AttachDisk(ctx context.Context, vmDiskPath string, volumeOptions *VolumeOptions) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !CheckControllerSupported(volumeOptions.SCSIControllerType) {
		return "", fmt.Errorf("Not a valid SCSI Controller Type. Valid options are %q", SCSIControllerTypeValidOptions())
	}
	vmDiskPathCopy := vmDiskPath
	vmDiskPath = RemoveStorageClusterORFolderNameFromVDiskPath(vmDiskPath)
	attached, err := vm.IsDiskAttached(ctx, vmDiskPath)
	if err != nil {
		klog.Errorf("Error occurred while checking if disk is attached on VM: %q. vmDiskPath: %q, err: %+v", vm.InventoryPath, vmDiskPath, err)
		return "", err
	}
	if attached {
		diskUUID, _ := vm.Datacenter.GetVirtualDiskPage83Data(ctx, vmDiskPath)
		return diskUUID, nil
	}
	if volumeOptions.StoragePolicyName != "" {
		pbmClient, err := NewPbmClient(ctx, vm.Client())
		if err != nil {
			klog.Errorf("Error occurred while creating new pbmClient. err: %+v", err)
			return "", err
		}
		volumeOptions.StoragePolicyID, err = pbmClient.ProfileIDByName(ctx, volumeOptions.StoragePolicyName)
		if err != nil {
			klog.Errorf("Failed to get Profile ID by name: %s. err: %+v", volumeOptions.StoragePolicyName, err)
			return "", err
		}
	}
	dsObj, err := vm.Datacenter.GetDatastoreByPath(ctx, vmDiskPathCopy)
	if err != nil {
		klog.Errorf("Failed to get datastore from vmDiskPath: %q. err: %+v", vmDiskPath, err)
		return "", err
	}
	disk, newSCSIController, err := vm.CreateDiskSpec(ctx, vmDiskPath, dsObj, volumeOptions)
	if err != nil {
		klog.Errorf("Error occurred while creating disk spec. err: %+v", err)
		return "", err
	}
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		klog.Errorf("Failed to retrieve VM devices for VM: %q. err: %+v", vm.InventoryPath, err)
		return "", err
	}
	virtualMachineConfigSpec := types.VirtualMachineConfigSpec{}
	deviceConfigSpec := &types.VirtualDeviceConfigSpec{Device: disk, Operation: types.VirtualDeviceConfigSpecOperationAdd}
	if volumeOptions.StoragePolicyID != "" {
		profileSpec := &types.VirtualMachineDefinedProfileSpec{ProfileId: volumeOptions.StoragePolicyID}
		deviceConfigSpec.Profile = append(deviceConfigSpec.Profile, profileSpec)
	}
	virtualMachineConfigSpec.DeviceChange = append(virtualMachineConfigSpec.DeviceChange, deviceConfigSpec)
	requestTime := time.Now()
	task, err := vm.Reconfigure(ctx, virtualMachineConfigSpec)
	if err != nil {
		RecordvSphereMetric(APIAttachVolume, requestTime, err)
		klog.Errorf("Failed to attach the disk with storagePolicy: %q on VM: %q. err - %+v", volumeOptions.StoragePolicyID, vm.InventoryPath, err)
		if newSCSIController != nil {
			vm.deleteController(ctx, newSCSIController, vmDevices)
		}
		return "", err
	}
	err = task.Wait(ctx)
	RecordvSphereMetric(APIAttachVolume, requestTime, err)
	if err != nil {
		klog.Errorf("Failed to attach the disk with storagePolicy: %+q on VM: %q. err - %+v", volumeOptions.StoragePolicyID, vm.InventoryPath, err)
		if newSCSIController != nil {
			vm.deleteController(ctx, newSCSIController, vmDevices)
		}
		return "", err
	}
	diskUUID, err := vm.Datacenter.GetVirtualDiskPage83Data(ctx, vmDiskPath)
	if err != nil {
		klog.Errorf("Error occurred while getting Disk Info from VM: %q. err: %v", vm.InventoryPath, err)
		vm.DetachDisk(ctx, vmDiskPath)
		if newSCSIController != nil {
			vm.deleteController(ctx, newSCSIController, vmDevices)
		}
		return "", err
	}
	return diskUUID, nil
}
func (vm *VirtualMachine) DetachDisk(ctx context.Context, vmDiskPath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmDiskPath = RemoveStorageClusterORFolderNameFromVDiskPath(vmDiskPath)
	device, err := vm.getVirtualDeviceByPath(ctx, vmDiskPath)
	if err != nil {
		klog.Errorf("Disk ID not found for VM: %q with diskPath: %q", vm.InventoryPath, vmDiskPath)
		return err
	}
	if device == nil {
		klog.Errorf("No virtual device found with diskPath: %q on VM: %q", vmDiskPath, vm.InventoryPath)
		return fmt.Errorf("No virtual device found with diskPath: %q on VM: %q", vmDiskPath, vm.InventoryPath)
	}
	requestTime := time.Now()
	err = vm.RemoveDevice(ctx, true, device)
	RecordvSphereMetric(APIDetachVolume, requestTime, err)
	if err != nil {
		klog.Errorf("Error occurred while removing disk device for VM: %q. err: %v", vm.InventoryPath, err)
		return err
	}
	return nil
}
func (vm *VirtualMachine) GetResourcePool(ctx context.Context) (*object.ResourcePool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmMoList, err := vm.Datacenter.GetVMMoList(ctx, []*VirtualMachine{vm}, []string{"resourcePool"})
	if err != nil {
		klog.Errorf("Failed to get resource pool from VM: %q. err: %+v", vm.InventoryPath, err)
		return nil, err
	}
	return object.NewResourcePool(vm.Client(), vmMoList[0].ResourcePool.Reference()), nil
}
func (vm *VirtualMachine) IsActive(ctx context.Context) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmMoList, err := vm.Datacenter.GetVMMoList(ctx, []*VirtualMachine{vm}, []string{"summary"})
	if err != nil {
		klog.Errorf("Failed to get VM Managed object with property summary. err: +%v", err)
		return false, err
	}
	if vmMoList[0].Summary.Runtime.PowerState == ActivePowerState {
		return true, nil
	}
	return false, nil
}
func (vm *VirtualMachine) GetAllAccessibleDatastores(ctx context.Context) ([]*DatastoreInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	host, err := vm.HostSystem(ctx)
	if err != nil {
		klog.Errorf("Failed to get host system for VM: %q. err: %+v", vm.InventoryPath, err)
		return nil, err
	}
	var hostSystemMo mo.HostSystem
	s := object.NewSearchIndex(vm.Client())
	err = s.Properties(ctx, host.Reference(), []string{DatastoreProperty}, &hostSystemMo)
	if err != nil {
		klog.Errorf("Failed to retrieve datastores for host: %+v. err: %+v", host, err)
		return nil, err
	}
	var dsRefList []types.ManagedObjectReference
	for _, dsRef := range hostSystemMo.Datastore {
		dsRefList = append(dsRefList, dsRef)
	}
	var dsMoList []mo.Datastore
	pc := property.DefaultCollector(vm.Client())
	properties := []string{DatastoreInfoProperty}
	err = pc.Retrieve(ctx, dsRefList, properties, &dsMoList)
	if err != nil {
		klog.Errorf("Failed to get Datastore managed objects from datastore objects."+" dsObjList: %+v, properties: %+v, err: %v", dsRefList, properties, err)
		return nil, err
	}
	klog.V(9).Infof("Result dsMoList: %+v", dsMoList)
	var dsObjList []*DatastoreInfo
	for _, dsMo := range dsMoList {
		dsObjList = append(dsObjList, &DatastoreInfo{&Datastore{object.NewDatastore(vm.Client(), dsMo.Reference()), vm.Datacenter}, dsMo.Info.GetDatastoreInfo()})
	}
	return dsObjList, nil
}
func (vm *VirtualMachine) CreateDiskSpec(ctx context.Context, diskPath string, dsObj *Datastore, volumeOptions *VolumeOptions) (*types.VirtualDisk, types.BaseVirtualDevice, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var newSCSIController types.BaseVirtualDevice
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		klog.Errorf("Failed to retrieve VM devices. err: %+v", err)
		return nil, nil, err
	}
	scsiControllersOfRequiredType := getSCSIControllersOfType(vmDevices, volumeOptions.SCSIControllerType)
	scsiController := getAvailableSCSIController(scsiControllersOfRequiredType)
	if scsiController == nil {
		newSCSIController, err = vm.createAndAttachSCSIController(ctx, volumeOptions.SCSIControllerType)
		if err != nil {
			klog.Errorf("Failed to create SCSI controller for VM :%q with err: %+v", vm.InventoryPath, err)
			return nil, nil, err
		}
		vmDevices, err := vm.Device(ctx)
		if err != nil {
			klog.Errorf("Failed to retrieve VM devices. err: %v", err)
			return nil, nil, err
		}
		scsiControllersOfRequiredType := getSCSIControllersOfType(vmDevices, volumeOptions.SCSIControllerType)
		scsiController = getAvailableSCSIController(scsiControllersOfRequiredType)
		if scsiController == nil {
			klog.Errorf("Cannot find SCSI controller of type: %q in VM", volumeOptions.SCSIControllerType)
			vm.deleteController(ctx, newSCSIController, vmDevices)
			return nil, nil, fmt.Errorf("Cannot find SCSI controller of type: %q in VM", volumeOptions.SCSIControllerType)
		}
	}
	disk := vmDevices.CreateDisk(scsiController, dsObj.Reference(), diskPath)
	unitNumber, err := getNextUnitNumber(vmDevices, scsiController)
	if err != nil {
		klog.Errorf("Cannot attach disk to VM, unitNumber limit reached - %+v.", err)
		return nil, nil, err
	}
	*disk.UnitNumber = unitNumber
	backing := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
	backing.DiskMode = string(types.VirtualDiskModeIndependent_persistent)
	if volumeOptions.CapacityKB != 0 {
		disk.CapacityInKB = int64(volumeOptions.CapacityKB)
	}
	if volumeOptions.DiskFormat != "" {
		var diskFormat string
		diskFormat = DiskFormatValidType[volumeOptions.DiskFormat]
		switch diskFormat {
		case ThinDiskType:
			backing.ThinProvisioned = types.NewBool(true)
		case EagerZeroedThickDiskType:
			backing.EagerlyScrub = types.NewBool(true)
		default:
			backing.ThinProvisioned = types.NewBool(false)
		}
	}
	return disk, newSCSIController, nil
}
func (vm *VirtualMachine) GetVirtualDiskPath(ctx context.Context) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		klog.Errorf("Failed to get the devices for VM: %q. err: %+v", vm.InventoryPath, err)
		return "", err
	}
	for _, device := range vmDevices {
		if vmDevices.TypeName(device) == "VirtualDisk" {
			virtualDevice := device.GetVirtualDevice()
			if backing, ok := virtualDevice.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
				return backing.FileName, nil
			}
		}
	}
	return "", nil
}
func (vm *VirtualMachine) createAndAttachSCSIController(ctx context.Context, diskControllerType string) (types.BaseVirtualDevice, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		klog.Errorf("Failed to retrieve VM devices for VM: %q. err: %+v", vm.InventoryPath, err)
		return nil, err
	}
	allSCSIControllers := getSCSIControllers(vmDevices)
	if len(allSCSIControllers) >= SCSIControllerLimit {
		klog.Errorf("SCSI Controller Limit of %d has been reached, cannot create another SCSI controller", SCSIControllerLimit)
		return nil, fmt.Errorf("SCSI Controller Limit of %d has been reached, cannot create another SCSI controller", SCSIControllerLimit)
	}
	newSCSIController, err := vmDevices.CreateSCSIController(diskControllerType)
	if err != nil {
		klog.Errorf("Failed to create new SCSI controller on VM: %q. err: %+v", vm.InventoryPath, err)
		return nil, err
	}
	configNewSCSIController := newSCSIController.(types.BaseVirtualSCSIController).GetVirtualSCSIController()
	hotAndRemove := true
	configNewSCSIController.HotAddRemove = &hotAndRemove
	configNewSCSIController.SharedBus = types.VirtualSCSISharing(types.VirtualSCSISharingNoSharing)
	err = vm.AddDevice(context.TODO(), newSCSIController)
	if err != nil {
		klog.V(LogLevel).Infof("Cannot add SCSI controller to VM: %q. err: %+v", vm.InventoryPath, err)
		vm.deleteController(ctx, newSCSIController, vmDevices)
		return nil, err
	}
	return newSCSIController, nil
}
func (vm *VirtualMachine) getVirtualDeviceByPath(ctx context.Context, diskPath string) (types.BaseVirtualDevice, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		klog.Errorf("Failed to get the devices for VM: %q. err: %+v", vm.InventoryPath, err)
		return nil, err
	}
	for _, device := range vmDevices {
		if vmDevices.TypeName(device) == "VirtualDisk" {
			virtualDevice := device.GetVirtualDevice()
			if backing, ok := virtualDevice.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
				if matchVirtualDiskAndVolPath(backing.FileName, diskPath) {
					klog.V(LogLevel).Infof("Found VirtualDisk backing with filename %q for diskPath %q", backing.FileName, diskPath)
					return device, nil
				}
			}
		}
	}
	return nil, nil
}
func matchVirtualDiskAndVolPath(diskPath, volPath string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fileExt := ".vmdk"
	diskPath = strings.TrimSuffix(diskPath, fileExt)
	volPath = strings.TrimSuffix(volPath, fileExt)
	return diskPath == volPath
}
func (vm *VirtualMachine) deleteController(ctx context.Context, controllerDevice types.BaseVirtualDevice, vmDevices object.VirtualDeviceList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerDeviceList := vmDevices.SelectByType(controllerDevice)
	if len(controllerDeviceList) < 1 {
		return ErrNoDevicesFound
	}
	device := controllerDeviceList[len(controllerDeviceList)-1]
	err := vm.RemoveDevice(ctx, true, device)
	if err != nil {
		klog.Errorf("Error occurred while removing device on VM: %q. err: %+v", vm.InventoryPath, err)
		return err
	}
	return nil
}
func (vm *VirtualMachine) RenewVM(client *vim25.Client) VirtualMachine {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dc := Datacenter{Datacenter: object.NewDatacenter(client, vm.Datacenter.Reference())}
	newVM := object.NewVirtualMachine(client, vm.VirtualMachine.Reference())
	return VirtualMachine{VirtualMachine: newVM, Datacenter: &dc}
}
