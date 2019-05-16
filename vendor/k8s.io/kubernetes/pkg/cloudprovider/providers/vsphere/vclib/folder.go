package vclib

import (
	"context"
	"github.com/vmware/govmomi/object"
	"k8s.io/klog"
)

type Folder struct {
	*object.Folder
	Datacenter *Datacenter
}

func (folder *Folder) GetVirtualMachines(ctx context.Context) ([]*VirtualMachine, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmFolders, err := folder.Children(ctx)
	if err != nil {
		klog.Errorf("Failed to get children from Folder: %s. err: %+v", folder.InventoryPath, err)
		return nil, err
	}
	var vmObjList []*VirtualMachine
	for _, vmFolder := range vmFolders {
		if vmFolder.Reference().Type == VirtualMachineType {
			vmObj := VirtualMachine{object.NewVirtualMachine(folder.Client(), vmFolder.Reference()), folder.Datacenter}
			vmObjList = append(vmObjList, &vmObj)
		}
	}
	return vmObjList, nil
}
