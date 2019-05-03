package vclib

import (
 "strings"
 "k8s.io/klog"
)

type VolumeOptions struct {
 CapacityKB             int
 Tags                   map[string]string
 Name                   string
 DiskFormat             string
 Datastore              string
 VSANStorageProfileData string
 StoragePolicyName      string
 StoragePolicyID        string
 SCSIControllerType     string
}

var (
 DiskFormatValidType     = map[string]string{ThinDiskType: ThinDiskType, strings.ToLower(EagerZeroedThickDiskType): EagerZeroedThickDiskType, strings.ToLower(ZeroedThickDiskType): PreallocatedDiskType}
 SCSIControllerValidType = []string{LSILogicControllerType, LSILogicSASControllerType, PVSCSIControllerType}
)

func DiskformatValidOptions() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validopts := ""
 for diskformat := range DiskFormatValidType {
  validopts += diskformat + ", "
 }
 validopts = strings.TrimSuffix(validopts, ", ")
 return validopts
}
func CheckDiskFormatSupported(diskFormat string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if DiskFormatValidType[diskFormat] == "" {
  klog.Errorf("Not a valid Disk Format. Valid options are %+q", DiskformatValidOptions())
  return false
 }
 return true
}
func SCSIControllerTypeValidOptions() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validopts := ""
 for _, controllerType := range SCSIControllerValidType {
  validopts += (controllerType + ", ")
 }
 validopts = strings.TrimSuffix(validopts, ", ")
 return validopts
}
func CheckControllerSupported(ctrlType string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range SCSIControllerValidType {
  if ctrlType == c {
   return true
  }
 }
 klog.Errorf("Not a valid SCSI Controller Type. Valid options are %q", SCSIControllerTypeValidOptions())
 return false
}
func (volumeOptions VolumeOptions) VerifyVolumeOptions() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if volumeOptions.SCSIControllerType != "" {
  isValid := CheckControllerSupported(volumeOptions.SCSIControllerType)
  if !isValid {
   return false
  }
 }
 if volumeOptions.DiskFormat != ThinDiskType {
  isValid := CheckDiskFormatSupported(volumeOptions.DiskFormat)
  if !isValid {
   return false
  }
 }
 return true
}
