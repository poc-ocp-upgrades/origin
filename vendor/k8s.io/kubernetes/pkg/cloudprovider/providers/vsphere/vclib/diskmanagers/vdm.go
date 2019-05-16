package diskmanagers

import (
	"context"
	goformat "fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type virtualDiskManager struct {
	diskPath      string
	volumeOptions *vclib.VolumeOptions
}

func (diskManager virtualDiskManager) Create(ctx context.Context, datastore *vclib.Datastore) (canonicalDiskPath string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if diskManager.volumeOptions.SCSIControllerType == "" {
		diskManager.volumeOptions.SCSIControllerType = vclib.LSILogicControllerType
	}
	diskFormat := vclib.DiskFormatValidType[diskManager.volumeOptions.DiskFormat]
	vdm := object.NewVirtualDiskManager(datastore.Client())
	vmDiskSpec := &types.FileBackedVirtualDiskSpec{VirtualDiskSpec: types.VirtualDiskSpec{AdapterType: diskManager.volumeOptions.SCSIControllerType, DiskType: diskFormat}, CapacityKb: int64(diskManager.volumeOptions.CapacityKB)}
	requestTime := time.Now()
	task, err := vdm.CreateVirtualDisk(ctx, diskManager.diskPath, datastore.Datacenter.Datacenter, vmDiskSpec)
	if err != nil {
		vclib.RecordvSphereMetric(vclib.APICreateVolume, requestTime, err)
		klog.Errorf("Failed to create virtual disk: %s. err: %+v", diskManager.diskPath, err)
		return "", err
	}
	taskInfo, err := task.WaitForResult(ctx, nil)
	vclib.RecordvSphereMetric(vclib.APICreateVolume, requestTime, err)
	if err != nil {
		klog.Errorf("Failed to complete virtual disk creation: %s. err: %+v", diskManager.diskPath, err)
		return "", err
	}
	canonicalDiskPath = taskInfo.Result.(string)
	return canonicalDiskPath, nil
}
func (diskManager virtualDiskManager) Delete(ctx context.Context, datacenter *vclib.Datacenter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	virtualDiskManager := object.NewVirtualDiskManager(datacenter.Client())
	diskPath := vclib.RemoveStorageClusterORFolderNameFromVDiskPath(diskManager.diskPath)
	requestTime := time.Now()
	task, err := virtualDiskManager.DeleteVirtualDisk(ctx, diskPath, datacenter.Datacenter)
	if err != nil {
		klog.Errorf("Failed to delete virtual disk. err: %v", err)
		vclib.RecordvSphereMetric(vclib.APIDeleteVolume, requestTime, err)
		return err
	}
	err = task.Wait(ctx)
	vclib.RecordvSphereMetric(vclib.APIDeleteVolume, requestTime, err)
	if err != nil {
		klog.Errorf("Failed to delete virtual disk. err: %v", err)
		return err
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
