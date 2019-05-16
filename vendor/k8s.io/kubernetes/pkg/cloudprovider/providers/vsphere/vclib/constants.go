package vclib

const (
	ThinDiskType             = "thin"
	PreallocatedDiskType     = "preallocated"
	EagerZeroedThickDiskType = "eagerZeroedThick"
	ZeroedThickDiskType      = "zeroedThick"
)
const (
	SCSIControllerLimit       = 4
	SCSIControllerDeviceLimit = 15
	SCSIDeviceSlots           = 16
	SCSIReservedSlot          = 7
	SCSIControllerType        = "scsi"
	LSILogicControllerType    = "lsiLogic"
	BusLogicControllerType    = "busLogic"
	LSILogicSASControllerType = "lsiLogic-sas"
	PVSCSIControllerType      = "pvscsi"
)
const (
	LogLevel                 = 4
	DatastoreProperty        = "datastore"
	ResourcePoolProperty     = "resourcePool"
	DatastoreInfoProperty    = "info"
	VirtualMachineType       = "VirtualMachine"
	RoundTripperDefaultCount = 3
	VSANDatastoreType        = "vsan"
	DummyVMPrefixName        = "vsphere-k8s"
	ActivePowerState         = "poweredOn"
)
const (
	TestDefaultDatacenter = "DC0"
	TestDefaultDatastore  = "LocalDS_0"
	TestDefaultNetwork    = "VM Network"
	testNameNotFound      = "enoent"
)
