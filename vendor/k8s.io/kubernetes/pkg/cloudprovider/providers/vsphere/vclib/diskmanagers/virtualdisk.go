package diskmanagers

import (
 "context"
 "fmt"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
)

type VirtualDisk struct {
 DiskPath      string
 VolumeOptions *vclib.VolumeOptions
 VMOptions     *vclib.VMOptions
}

const (
 VirtualDiskCreateOperation = "Create"
 VirtualDiskDeleteOperation = "Delete"
)

type VirtualDiskProvider interface {
 Create(ctx context.Context, datastore *vclib.Datastore) (string, error)
 Delete(ctx context.Context, datacenter *vclib.Datacenter) error
}

func getDiskManager(disk *VirtualDisk, diskOperation string) VirtualDiskProvider {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var diskProvider VirtualDiskProvider
 switch diskOperation {
 case VirtualDiskDeleteOperation:
  diskProvider = virtualDiskManager{disk.DiskPath, disk.VolumeOptions}
 case VirtualDiskCreateOperation:
  if disk.VolumeOptions.StoragePolicyName != "" || disk.VolumeOptions.VSANStorageProfileData != "" || disk.VolumeOptions.StoragePolicyID != "" {
   diskProvider = vmDiskManager{disk.DiskPath, disk.VolumeOptions, disk.VMOptions}
  } else {
   diskProvider = virtualDiskManager{disk.DiskPath, disk.VolumeOptions}
  }
 }
 return diskProvider
}
func (virtualDisk *VirtualDisk) Create(ctx context.Context, datastore *vclib.Datastore) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if virtualDisk.VolumeOptions.DiskFormat == "" {
  virtualDisk.VolumeOptions.DiskFormat = vclib.ThinDiskType
 }
 if !virtualDisk.VolumeOptions.VerifyVolumeOptions() {
  klog.Error("VolumeOptions verification failed. volumeOptions: ", virtualDisk.VolumeOptions)
  return "", vclib.ErrInvalidVolumeOptions
 }
 if virtualDisk.VolumeOptions.StoragePolicyID != "" && virtualDisk.VolumeOptions.StoragePolicyName != "" {
  return "", fmt.Errorf("Storage Policy ID and Storage Policy Name both set, Please set only one parameter")
 }
 return getDiskManager(virtualDisk, VirtualDiskCreateOperation).Create(ctx, datastore)
}
func (virtualDisk *VirtualDisk) Delete(ctx context.Context, datacenter *vclib.Datacenter) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return getDiskManager(virtualDisk, VirtualDiskDeleteOperation).Delete(ctx, datacenter)
}
