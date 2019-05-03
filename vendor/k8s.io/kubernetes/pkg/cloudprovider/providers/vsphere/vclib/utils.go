package vclib

import (
 "fmt"
 "path/filepath"
 "regexp"
 "strings"
 "github.com/vmware/govmomi/find"
 "github.com/vmware/govmomi/object"
 "github.com/vmware/govmomi/vim25/mo"
 "github.com/vmware/govmomi/vim25/soap"
 "github.com/vmware/govmomi/vim25/types"
 "k8s.io/klog"
)

func IsNotFound(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, ok := err.(*find.NotFoundError)
 if ok {
  return true
 }
 _, ok = err.(*find.DefaultNotFoundError)
 if ok {
  return true
 }
 return false
}
func getFinder(dc *Datacenter) *find.Finder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 finder := find.NewFinder(dc.Client(), false)
 finder.SetDatacenter(dc.Datacenter)
 return finder
}
func formatVirtualDiskUUID(uuid string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 uuidwithNoSpace := strings.Replace(uuid, " ", "", -1)
 uuidWithNoHypens := strings.Replace(uuidwithNoSpace, "-", "", -1)
 return strings.ToLower(uuidWithNoHypens)
}
func getSCSIControllersOfType(vmDevices object.VirtualDeviceList, scsiType string) []*types.VirtualController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var scsiControllers []*types.VirtualController
 for _, device := range vmDevices {
  devType := vmDevices.Type(device)
  if devType == scsiType {
   if c, ok := device.(types.BaseVirtualController); ok {
    scsiControllers = append(scsiControllers, c.GetVirtualController())
   }
  }
 }
 return scsiControllers
}
func getAvailableSCSIController(scsiControllers []*types.VirtualController) *types.VirtualController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, controller := range scsiControllers {
  if len(controller.Device) < SCSIControllerDeviceLimit {
   return controller
  }
 }
 return nil
}
func getNextUnitNumber(devices object.VirtualDeviceList, c types.BaseVirtualController) (int32, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var takenUnitNumbers [SCSIDeviceSlots]bool
 takenUnitNumbers[SCSIReservedSlot] = true
 key := c.GetVirtualController().Key
 for _, device := range devices {
  d := device.GetVirtualDevice()
  if d.ControllerKey == key {
   if d.UnitNumber != nil {
    takenUnitNumbers[*d.UnitNumber] = true
   }
  }
 }
 for unitNumber, takenUnitNumber := range takenUnitNumbers {
  if !takenUnitNumber {
   return int32(unitNumber), nil
  }
 }
 return -1, fmt.Errorf("SCSI Controller with key=%d does not have any available slots", key)
}
func getSCSIControllers(vmDevices object.VirtualDeviceList) []*types.VirtualController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var scsiControllers []*types.VirtualController
 for _, device := range vmDevices {
  devType := vmDevices.Type(device)
  switch devType {
  case SCSIControllerType, strings.ToLower(LSILogicControllerType), strings.ToLower(BusLogicControllerType), PVSCSIControllerType, strings.ToLower(LSILogicSASControllerType):
   if c, ok := device.(types.BaseVirtualController); ok {
    scsiControllers = append(scsiControllers, c.GetVirtualController())
   }
  }
 }
 return scsiControllers
}
func RemoveStorageClusterORFolderNameFromVDiskPath(vDiskPath string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 datastore := regexp.MustCompile("\\[(.*?)\\]").FindStringSubmatch(vDiskPath)[1]
 if filepath.Base(datastore) != datastore {
  vDiskPath = strings.Replace(vDiskPath, datastore, filepath.Base(datastore), 1)
 }
 return vDiskPath
}
func GetPathFromVMDiskPath(vmDiskPath string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 datastorePathObj := new(object.DatastorePath)
 isSuccess := datastorePathObj.FromString(vmDiskPath)
 if !isSuccess {
  klog.Errorf("Failed to parse vmDiskPath: %s", vmDiskPath)
  return ""
 }
 return datastorePathObj.Path
}
func GetDatastorePathObjFromVMDiskPath(vmDiskPath string) (*object.DatastorePath, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 datastorePathObj := new(object.DatastorePath)
 isSuccess := datastorePathObj.FromString(vmDiskPath)
 if !isSuccess {
  klog.Errorf("Failed to parse volPath: %s", vmDiskPath)
  return nil, fmt.Errorf("Failed to parse volPath: %s", vmDiskPath)
 }
 return datastorePathObj, nil
}
func IsValidUUID(uuid string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
 return r.MatchString(uuid)
}
func IsManagedObjectNotFoundError(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isManagedObjectNotFoundError := false
 if soap.IsSoapFault(err) {
  _, isManagedObjectNotFoundError = soap.ToSoapFault(err).VimFault().(types.ManagedObjectNotFound)
 }
 return isManagedObjectNotFoundError
}
func IsInvalidCredentialsError(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isInvalidCredentialsError := false
 if soap.IsSoapFault(err) {
  _, isInvalidCredentialsError = soap.ToSoapFault(err).VimFault().(types.InvalidLogin)
 }
 return isInvalidCredentialsError
}
func VerifyVolumePathsForVM(vmMo mo.VirtualMachine, volPaths []string, nodeName string, nodeVolumeMap map[string]map[string]bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmDevices := object.VirtualDeviceList(vmMo.Config.Hardware.Device)
 VerifyVolumePathsForVMDevices(vmDevices, volPaths, nodeName, nodeVolumeMap)
}
func VerifyVolumePathsForVMDevices(vmDevices object.VirtualDeviceList, volPaths []string, nodeName string, nodeVolumeMap map[string]map[string]bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 volPathsMap := make(map[string]bool)
 for _, volPath := range volPaths {
  volPathsMap[volPath] = true
 }
 for _, device := range vmDevices {
  if vmDevices.TypeName(device) == "VirtualDisk" {
   virtualDevice := device.GetVirtualDevice()
   if backing, ok := virtualDevice.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
    if volPathsMap[backing.FileName] {
     setNodeVolumeMap(nodeVolumeMap, backing.FileName, nodeName, true)
    }
   }
  }
 }
}
