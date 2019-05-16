package vclib

import (
	"context"
	"errors"
	"fmt"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"k8s.io/klog"
	"path/filepath"
	"strings"
)

type Datacenter struct{ *object.Datacenter }

func GetDatacenter(ctx context.Context, connection *VSphereConnection, datacenterPath string) (*Datacenter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := find.NewFinder(connection.Client, false)
	datacenter, err := finder.Datacenter(ctx, datacenterPath)
	if err != nil {
		klog.Errorf("Failed to find the datacenter: %s. err: %+v", datacenterPath, err)
		return nil, err
	}
	dc := Datacenter{datacenter}
	return &dc, nil
}
func GetAllDatacenter(ctx context.Context, connection *VSphereConnection) ([]*Datacenter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var dc []*Datacenter
	finder := find.NewFinder(connection.Client, false)
	datacenters, err := finder.DatacenterList(ctx, "*")
	if err != nil {
		klog.Errorf("Failed to find the datacenter. err: %+v", err)
		return nil, err
	}
	for _, datacenter := range datacenters {
		dc = append(dc, &(Datacenter{datacenter}))
	}
	return dc, nil
}
func (dc *Datacenter) GetVMByUUID(ctx context.Context, vmUUID string) (*VirtualMachine, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := object.NewSearchIndex(dc.Client())
	vmUUID = strings.ToLower(strings.TrimSpace(vmUUID))
	svm, err := s.FindByUuid(ctx, dc.Datacenter, vmUUID, true, nil)
	if err != nil {
		klog.Errorf("Failed to find VM by UUID. VM UUID: %s, err: %+v", vmUUID, err)
		return nil, err
	}
	if svm == nil {
		klog.Errorf("Unable to find VM by UUID. VM UUID: %s", vmUUID)
		return nil, ErrNoVMFound
	}
	virtualMachine := VirtualMachine{object.NewVirtualMachine(dc.Client(), svm.Reference()), dc}
	return &virtualMachine, nil
}
func (dc *Datacenter) GetHostByVMUUID(ctx context.Context, vmUUID string) (*types.ManagedObjectReference, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	virtualMachine, err := dc.GetVMByUUID(ctx, vmUUID)
	var vmMo mo.VirtualMachine
	pc := property.DefaultCollector(virtualMachine.Client())
	err = pc.RetrieveOne(ctx, virtualMachine.Reference(), []string{"summary.runtime.host"}, &vmMo)
	if err != nil {
		klog.Errorf("Failed to retrive VM runtime host, err: %v", err)
		return nil, err
	}
	host := vmMo.Summary.Runtime.Host
	klog.Infof("%s host is %s", virtualMachine.Reference(), host)
	return host, nil
}
func (dc *Datacenter) GetVMByPath(ctx context.Context, vmPath string) (*VirtualMachine, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := getFinder(dc)
	vm, err := finder.VirtualMachine(ctx, vmPath)
	if err != nil {
		klog.Errorf("Failed to find VM by Path. VM Path: %s, err: %+v", vmPath, err)
		return nil, err
	}
	virtualMachine := VirtualMachine{vm, dc}
	return &virtualMachine, nil
}
func (dc *Datacenter) GetAllDatastores(ctx context.Context) (map[string]*DatastoreInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := getFinder(dc)
	datastores, err := finder.DatastoreList(ctx, "*")
	if err != nil {
		klog.Errorf("Failed to get all the datastores. err: %+v", err)
		return nil, err
	}
	var dsList []types.ManagedObjectReference
	for _, ds := range datastores {
		dsList = append(dsList, ds.Reference())
	}
	var dsMoList []mo.Datastore
	pc := property.DefaultCollector(dc.Client())
	properties := []string{DatastoreInfoProperty}
	err = pc.Retrieve(ctx, dsList, properties, &dsMoList)
	if err != nil {
		klog.Errorf("Failed to get Datastore managed objects from datastore objects."+" dsObjList: %+v, properties: %+v, err: %v", dsList, properties, err)
		return nil, err
	}
	dsURLInfoMap := make(map[string]*DatastoreInfo)
	for _, dsMo := range dsMoList {
		dsURLInfoMap[dsMo.Info.GetDatastoreInfo().Url] = &DatastoreInfo{&Datastore{object.NewDatastore(dc.Client(), dsMo.Reference()), dc}, dsMo.Info.GetDatastoreInfo()}
	}
	klog.V(9).Infof("dsURLInfoMap : %+v", dsURLInfoMap)
	return dsURLInfoMap, nil
}
func (dc *Datacenter) GetDatastoreByPath(ctx context.Context, vmDiskPath string) (*Datastore, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	datastorePathObj := new(object.DatastorePath)
	isSuccess := datastorePathObj.FromString(vmDiskPath)
	if !isSuccess {
		klog.Errorf("Failed to parse vmDiskPath: %s", vmDiskPath)
		return nil, errors.New("Failed to parse vmDiskPath")
	}
	return dc.GetDatastoreByName(ctx, datastorePathObj.Datastore)
}
func (dc *Datacenter) GetDatastoreByName(ctx context.Context, name string) (*Datastore, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := getFinder(dc)
	ds, err := finder.Datastore(ctx, name)
	if err != nil {
		klog.Errorf("Failed while searching for datastore: %s. err: %+v", name, err)
		return nil, err
	}
	datastore := Datastore{ds, dc}
	return &datastore, nil
}
func (dc *Datacenter) GetResourcePool(ctx context.Context, resourcePoolPath string) (*object.ResourcePool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := getFinder(dc)
	var resourcePool *object.ResourcePool
	var err error
	resourcePool, err = finder.ResourcePoolOrDefault(ctx, resourcePoolPath)
	if err != nil {
		klog.Errorf("Failed to get the ResourcePool for path '%s'. err: %+v", resourcePoolPath, err)
		return nil, err
	}
	return resourcePool, nil
}
func (dc *Datacenter) GetFolderByPath(ctx context.Context, folderPath string) (*Folder, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finder := getFinder(dc)
	vmFolder, err := finder.Folder(ctx, folderPath)
	if err != nil {
		klog.Errorf("Failed to get the folder reference for %s. err: %+v", folderPath, err)
		return nil, err
	}
	folder := Folder{vmFolder, dc}
	return &folder, nil
}
func (dc *Datacenter) GetVMMoList(ctx context.Context, vmObjList []*VirtualMachine, properties []string) ([]mo.VirtualMachine, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var vmMoList []mo.VirtualMachine
	var vmRefs []types.ManagedObjectReference
	if len(vmObjList) < 1 {
		klog.Errorf("VirtualMachine Object list is empty")
		return nil, fmt.Errorf("VirtualMachine Object list is empty")
	}
	for _, vmObj := range vmObjList {
		vmRefs = append(vmRefs, vmObj.Reference())
	}
	pc := property.DefaultCollector(dc.Client())
	err := pc.Retrieve(ctx, vmRefs, properties, &vmMoList)
	if err != nil {
		klog.Errorf("Failed to get VM managed objects from VM objects. vmObjList: %+v, properties: %+v, err: %v", vmObjList, properties, err)
		return nil, err
	}
	return vmMoList, nil
}
func (dc *Datacenter) GetVirtualDiskPage83Data(ctx context.Context, diskPath string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(diskPath) > 0 && filepath.Ext(diskPath) != ".vmdk" {
		diskPath += ".vmdk"
	}
	vdm := object.NewVirtualDiskManager(dc.Client())
	diskUUID, err := vdm.QueryVirtualDiskUuid(ctx, diskPath, dc.Datacenter)
	if err != nil {
		klog.Warningf("QueryVirtualDiskUuid failed for diskPath: %q. err: %+v", diskPath, err)
		return "", err
	}
	diskUUID = formatVirtualDiskUUID(diskUUID)
	return diskUUID, nil
}
func (dc *Datacenter) GetDatastoreMoList(ctx context.Context, dsObjList []*Datastore, properties []string) ([]mo.Datastore, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var dsMoList []mo.Datastore
	var dsRefs []types.ManagedObjectReference
	if len(dsObjList) < 1 {
		klog.Errorf("Datastore Object list is empty")
		return nil, fmt.Errorf("Datastore Object list is empty")
	}
	for _, dsObj := range dsObjList {
		dsRefs = append(dsRefs, dsObj.Reference())
	}
	pc := property.DefaultCollector(dc.Client())
	err := pc.Retrieve(ctx, dsRefs, properties, &dsMoList)
	if err != nil {
		klog.Errorf("Failed to get Datastore managed objects from datastore objects. dsObjList: %+v, properties: %+v, err: %v", dsObjList, properties, err)
		return nil, err
	}
	return dsMoList, nil
}
func (dc *Datacenter) CheckDisksAttached(ctx context.Context, nodeVolumes map[string][]string) (map[string]map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attached := make(map[string]map[string]bool)
	var vmList []*VirtualMachine
	for nodeName, volPaths := range nodeVolumes {
		for _, volPath := range volPaths {
			setNodeVolumeMap(attached, volPath, nodeName, false)
		}
		vm, err := dc.GetVMByPath(ctx, nodeName)
		if err != nil {
			if IsNotFound(err) {
				klog.Warningf("Node %q does not exist, vSphere CP will assume disks %v are not attached to it.", nodeName, volPaths)
			}
			continue
		}
		vmList = append(vmList, vm)
	}
	if len(vmList) == 0 {
		klog.V(2).Infof("vSphere CP will assume no disks are attached to any node.")
		return attached, nil
	}
	vmMoList, err := dc.GetVMMoList(ctx, vmList, []string{"config.hardware.device", "name"})
	if err != nil {
		klog.Errorf("Failed to get VM Managed object for nodes: %+v. err: +%v", vmList, err)
		return nil, err
	}
	for _, vmMo := range vmMoList {
		if vmMo.Config == nil {
			klog.Errorf("Config is not available for VM: %q", vmMo.Name)
			continue
		}
		for nodeName, volPaths := range nodeVolumes {
			if nodeName == vmMo.Name {
				verifyVolumePathsForVM(vmMo, volPaths, attached)
			}
		}
	}
	return attached, nil
}
func verifyVolumePathsForVM(vmMo mo.VirtualMachine, volPaths []string, nodeVolumeMap map[string]map[string]bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, volPath := range volPaths {
		vmDevices := object.VirtualDeviceList(vmMo.Config.Hardware.Device)
		for _, device := range vmDevices {
			if vmDevices.TypeName(device) == "VirtualDisk" {
				virtualDevice := device.GetVirtualDevice()
				if backing, ok := virtualDevice.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
					if backing.FileName == volPath {
						setNodeVolumeMap(nodeVolumeMap, volPath, vmMo.Name, true)
					}
				}
			}
		}
	}
}
func setNodeVolumeMap(nodeVolumeMap map[string]map[string]bool, volumePath string, nodeName string, check bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeMap := nodeVolumeMap[nodeName]
	if volumeMap == nil {
		volumeMap = make(map[string]bool)
		nodeVolumeMap[nodeName] = volumeMap
	}
	volumeMap[volumePath] = check
}
