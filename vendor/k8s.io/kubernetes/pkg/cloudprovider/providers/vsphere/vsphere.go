package vsphere

import (
 "context"
 "errors"
 "fmt"
 "io"
 "net"
 "net/url"
 "os"
 "path"
 "path/filepath"
 "runtime"
 "strings"
 "sync"
 "time"
 "gopkg.in/gcfg.v1"
 "github.com/vmware/govmomi/vapi/rest"
 "github.com/vmware/govmomi/vapi/tags"
 "github.com/vmware/govmomi/vim25/mo"
 "k8s.io/api/core/v1"
 k8stypes "k8s.io/apimachinery/pkg/types"
 "k8s.io/client-go/informers"
 "k8s.io/client-go/tools/cache"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/klog"
 v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib/diskmanagers"
)

const (
 ProviderName                  = "vsphere"
 VolDir                        = "kubevols"
 RoundTripperDefaultCount      = 3
 DummyVMPrefixName             = "vsphere-k8s"
 CleanUpDummyVMRoutineInterval = 5
)

var cleanUpRoutineInitialized = false
var datastoreFolderIDMap = make(map[string]map[string]string)
var cleanUpRoutineInitLock sync.Mutex
var cleanUpDummyVMLock sync.RWMutex

const (
 MissingUsernameErrMsg = "Username is missing"
 MissingPasswordErrMsg = "Password is missing"
)

var (
 ErrUsernameMissing = errors.New(MissingUsernameErrMsg)
 ErrPasswordMissing = errors.New(MissingPasswordErrMsg)
)

type VSphere struct {
 cfg                  *VSphereConfig
 hostName             string
 vsphereInstanceMap   map[string]*VSphereInstance
 nodeManager          *NodeManager
 vmUUID               string
 isSecretInfoProvided bool
}
type VSphereInstance struct {
 conn *vclib.VSphereConnection
 cfg  *VirtualCenterConfig
}
type VirtualCenterConfig struct {
 User              string `gcfg:"user"`
 Password          string `gcfg:"password"`
 VCenterPort       string `gcfg:"port"`
 Datacenters       string `gcfg:"datacenters"`
 RoundTripperCount uint   `gcfg:"soap-roundtrip-count"`
 Thumbprint        string `gcfg:"thumbprint"`
}
type VSphereConfig struct {
 Global struct {
  User              string `gcfg:"user"`
  Password          string `gcfg:"password"`
  VCenterIP         string `gcfg:"server"`
  VCenterPort       string `gcfg:"port"`
  InsecureFlag      bool   `gcfg:"insecure-flag"`
  CAFile            string `gcfg:"ca-file"`
  Thumbprint        string `gcfg:"thumbprint"`
  Datacenter        string `gcfg:"datacenter"`
  Datacenters       string `gcfg:"datacenters"`
  DefaultDatastore  string `gcfg:"datastore"`
  WorkingDir        string `gcfg:"working-dir"`
  RoundTripperCount uint   `gcfg:"soap-roundtrip-count"`
  VMUUID            string `gcfg:"vm-uuid"`
  VMName            string `gcfg:"vm-name"`
  SecretName        string `gcfg:"secret-name"`
  SecretNamespace   string `gcfg:"secret-namespace"`
 }
 VirtualCenter map[string]*VirtualCenterConfig
 Network       struct {
  PublicNetwork string `gcfg:"public-network"`
 }
 Disk struct {
  SCSIControllerType string `dcfg:"scsicontrollertype"`
 }
 Workspace struct {
  VCenterIP        string `gcfg:"server"`
  Datacenter       string `gcfg:"datacenter"`
  Folder           string `gcfg:"folder"`
  DefaultDatastore string `gcfg:"default-datastore"`
  ResourcePoolPath string `gcfg:"resourcepool-path"`
 }
 Labels struct {
  Zone   string `gcfg:"zone"`
  Region string `gcfg:"region"`
 }
}
type Volumes interface {
 AttachDisk(vmDiskPath string, storagePolicyName string, nodeName k8stypes.NodeName) (diskUUID string, err error)
 DetachDisk(volPath string, nodeName k8stypes.NodeName) error
 DiskIsAttached(volPath string, nodeName k8stypes.NodeName) (bool, error)
 DisksAreAttached(nodeVolumes map[k8stypes.NodeName][]string) (map[k8stypes.NodeName]map[string]bool, error)
 CreateVolume(volumeOptions *vclib.VolumeOptions) (volumePath string, err error)
 DeleteVolume(vmDiskPath string) error
}

func readConfig(config io.Reader) (VSphereConfig, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if config == nil {
  err := fmt.Errorf("no vSphere cloud provider config file given")
  return VSphereConfig{}, err
 }
 var cfg VSphereConfig
 err := gcfg.ReadInto(&cfg, config)
 return cfg, err
}
func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vclib.RegisterMetrics()
 cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
  if config == nil {
   return newWorkerNode()
  }
  cfg, err := readConfig(config)
  if err != nil {
   return nil, err
  }
  return newControllerNode(cfg)
 })
}
func (vs *VSphere) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (vs *VSphere) SetInformers(informerFactory informers.SharedInformerFactory) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if vs.cfg == nil {
  return
 }
 if vs.isSecretInfoProvided {
  secretCredentialManager := &SecretCredentialManager{SecretName: vs.cfg.Global.SecretName, SecretNamespace: vs.cfg.Global.SecretNamespace, SecretLister: informerFactory.Core().V1().Secrets().Lister(), Cache: &SecretCache{VirtualCenter: make(map[string]*Credential)}}
  vs.nodeManager.UpdateCredentialManager(secretCredentialManager)
 }
 klog.V(4).Infof("Setting up node informers for vSphere Cloud Provider")
 nodeInformer := informerFactory.Core().V1().Nodes().Informer()
 nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: vs.NodeAdded, DeleteFunc: vs.NodeDeleted})
 klog.V(4).Infof("Node informers in vSphere cloud provider initialized")
}
func newWorkerNode() (*VSphere, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 vs := VSphere{}
 vs.hostName, err = os.Hostname()
 if err != nil {
  klog.Errorf("Failed to get hostname. err: %+v", err)
  return nil, err
 }
 vs.vmUUID, err = GetVMUUID()
 if err != nil {
  klog.Errorf("Failed to get uuid. err: %+v", err)
  return nil, err
 }
 return &vs, nil
}
func populateVsphereInstanceMap(cfg *VSphereConfig) (map[string]*VSphereInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vsphereInstanceMap := make(map[string]*VSphereInstance)
 isSecretInfoProvided := true
 if cfg.Global.SecretName == "" || cfg.Global.SecretNamespace == "" {
  klog.Warningf("SecretName and/or SecretNamespace is not provided. " + "VCP will use username and password from config file")
  isSecretInfoProvided = false
 }
 if isSecretInfoProvided {
  if cfg.Global.User != "" {
   klog.Warning("Global.User and Secret info provided. VCP will use secret to get credentials")
   cfg.Global.User = ""
  }
  if cfg.Global.Password != "" {
   klog.Warning("Global.Password and Secret info provided. VCP will use secret to get credentials")
   cfg.Global.Password = ""
  }
 }
 if cfg.VirtualCenter == nil || len(cfg.VirtualCenter) == 0 {
  klog.V(4).Infof("Config is not per virtual center and is in old format.")
  if !isSecretInfoProvided {
   if cfg.Global.User == "" {
    klog.Error("Global.User is empty!")
    return nil, ErrUsernameMissing
   }
   if cfg.Global.Password == "" {
    klog.Error("Global.Password is empty!")
    return nil, ErrPasswordMissing
   }
  }
  if cfg.Global.WorkingDir == "" {
   klog.Error("Global.WorkingDir is empty!")
   return nil, errors.New("Global.WorkingDir is empty!")
  }
  if cfg.Global.VCenterIP == "" {
   klog.Error("Global.VCenterIP is empty!")
   return nil, errors.New("Global.VCenterIP is empty!")
  }
  if cfg.Global.Datacenter == "" {
   klog.Error("Global.Datacenter is empty!")
   return nil, errors.New("Global.Datacenter is empty!")
  }
  cfg.Workspace.VCenterIP = cfg.Global.VCenterIP
  cfg.Workspace.Datacenter = cfg.Global.Datacenter
  cfg.Workspace.Folder = cfg.Global.WorkingDir
  cfg.Workspace.DefaultDatastore = cfg.Global.DefaultDatastore
  vcConfig := VirtualCenterConfig{User: cfg.Global.User, Password: cfg.Global.Password, VCenterPort: cfg.Global.VCenterPort, Datacenters: cfg.Global.Datacenter, RoundTripperCount: cfg.Global.RoundTripperCount, Thumbprint: cfg.Global.Thumbprint}
  vSphereConn := vclib.VSphereConnection{Username: vcConfig.User, Password: vcConfig.Password, Hostname: cfg.Global.VCenterIP, Insecure: cfg.Global.InsecureFlag, RoundTripperCount: vcConfig.RoundTripperCount, Port: vcConfig.VCenterPort, CACert: cfg.Global.CAFile, Thumbprint: cfg.Global.Thumbprint}
  vsphereIns := VSphereInstance{conn: &vSphereConn, cfg: &vcConfig}
  vsphereInstanceMap[cfg.Global.VCenterIP] = &vsphereIns
 } else {
  if cfg.Workspace.VCenterIP == "" || cfg.Workspace.Folder == "" || cfg.Workspace.Datacenter == "" {
   msg := fmt.Sprintf("All fields in workspace are mandatory."+" vsphere.conf does not have the workspace specified correctly. cfg.Workspace: %+v", cfg.Workspace)
   klog.Error(msg)
   return nil, errors.New(msg)
  }
  for vcServer, vcConfig := range cfg.VirtualCenter {
   klog.V(4).Infof("Initializing vc server %s", vcServer)
   if vcServer == "" {
    klog.Error("vsphere.conf does not have the VirtualCenter IP address specified")
    return nil, errors.New("vsphere.conf does not have the VirtualCenter IP address specified")
   }
   if !isSecretInfoProvided {
    if vcConfig.User == "" {
     vcConfig.User = cfg.Global.User
     if vcConfig.User == "" {
      klog.Errorf("vcConfig.User is empty for vc %s!", vcServer)
      return nil, ErrUsernameMissing
     }
    }
    if vcConfig.Password == "" {
     vcConfig.Password = cfg.Global.Password
     if vcConfig.Password == "" {
      klog.Errorf("vcConfig.Password is empty for vc %s!", vcServer)
      return nil, ErrPasswordMissing
     }
    }
   } else {
    if vcConfig.User != "" {
     klog.Warningf("vcConfig.User for server %s and Secret info provided. VCP will use secret to get credentials", vcServer)
     vcConfig.User = ""
    }
    if vcConfig.Password != "" {
     klog.Warningf("vcConfig.Password for server %s and Secret info provided. VCP will use secret to get credentials", vcServer)
     vcConfig.Password = ""
    }
   }
   if vcConfig.VCenterPort == "" {
    vcConfig.VCenterPort = cfg.Global.VCenterPort
   }
   if vcConfig.Datacenters == "" {
    if cfg.Global.Datacenters != "" {
     vcConfig.Datacenters = cfg.Global.Datacenters
    } else {
     vcConfig.Datacenters = cfg.Global.Datacenter
    }
   }
   if vcConfig.RoundTripperCount == 0 {
    vcConfig.RoundTripperCount = cfg.Global.RoundTripperCount
   }
   vSphereConn := vclib.VSphereConnection{Username: vcConfig.User, Password: vcConfig.Password, Hostname: vcServer, Insecure: cfg.Global.InsecureFlag, RoundTripperCount: vcConfig.RoundTripperCount, Port: vcConfig.VCenterPort, CACert: cfg.Global.CAFile, Thumbprint: vcConfig.Thumbprint}
   vsphereIns := VSphereInstance{conn: &vSphereConn, cfg: vcConfig}
   vsphereInstanceMap[vcServer] = &vsphereIns
  }
 }
 return vsphereInstanceMap, nil
}

var getVMUUID = GetVMUUID

func newControllerNode(cfg VSphereConfig) (*VSphere, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vs, err := buildVSphereFromConfig(cfg)
 if err != nil {
  return nil, err
 }
 vs.hostName, err = os.Hostname()
 if err != nil {
  klog.Errorf("Failed to get hostname. err: %+v", err)
  return nil, err
 }
 if cfg.Global.VMUUID != "" {
  vs.vmUUID = cfg.Global.VMUUID
 } else {
  vs.vmUUID, err = getVMUUID()
  if err != nil {
   klog.Errorf("Failed to get uuid. err: %+v", err)
   return nil, err
  }
 }
 runtime.SetFinalizer(vs, logout)
 return vs, nil
}
func buildVSphereFromConfig(cfg VSphereConfig) (*VSphere, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isSecretInfoProvided := false
 if cfg.Global.SecretName != "" && cfg.Global.SecretNamespace != "" {
  isSecretInfoProvided = true
 }
 if cfg.Disk.SCSIControllerType == "" {
  cfg.Disk.SCSIControllerType = vclib.PVSCSIControllerType
 } else if !vclib.CheckControllerSupported(cfg.Disk.SCSIControllerType) {
  klog.Errorf("%v is not a supported SCSI Controller type. Please configure 'lsilogic-sas' OR 'pvscsi'", cfg.Disk.SCSIControllerType)
  return nil, errors.New("Controller type not supported. Please configure 'lsilogic-sas' OR 'pvscsi'")
 }
 if cfg.Global.WorkingDir != "" {
  cfg.Global.WorkingDir = path.Clean(cfg.Global.WorkingDir)
 }
 if cfg.Global.RoundTripperCount == 0 {
  cfg.Global.RoundTripperCount = RoundTripperDefaultCount
 }
 if cfg.Global.VCenterPort == "" {
  cfg.Global.VCenterPort = "443"
 }
 vsphereInstanceMap, err := populateVsphereInstanceMap(&cfg)
 if err != nil {
  return nil, err
 }
 vs := VSphere{vsphereInstanceMap: vsphereInstanceMap, nodeManager: &NodeManager{vsphereInstanceMap: vsphereInstanceMap, nodeInfoMap: make(map[string]*NodeInfo), registeredNodes: make(map[string]*v1.Node)}, isSecretInfoProvided: isSecretInfoProvided, cfg: &cfg}
 return &vs, nil
}
func logout(vs *VSphere) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, vsphereIns := range vs.vsphereInstanceMap {
  vsphereIns.conn.Logout(context.TODO())
 }
}
func (vs *VSphere) Instances() (cloudprovider.Instances, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return vs, true
}
func getLocalIP() ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vmwareOUI := map[string]bool{"00:05:69": true, "00:0c:29": true, "00:1c:14": true, "00:50:56": true}
 addrs := []v1.NodeAddress{}
 ifaces, err := net.Interfaces()
 if err != nil {
  klog.Errorf("net.Interfaces() failed for NodeAddresses - %v", err)
  return nil, err
 }
 for _, i := range ifaces {
  if i.Flags&net.FlagLoopback != 0 {
   continue
  }
  localAddrs, err := i.Addrs()
  if err != nil {
   klog.Warningf("Failed to extract addresses for NodeAddresses - %v", err)
  } else {
   for _, addr := range localAddrs {
    if ipnet, ok := addr.(*net.IPNet); ok {
     if ipnet.IP.To4() != nil {
      vmMACAddr := strings.ToLower(i.HardwareAddr.String())
      if len(vmMACAddr) < 17 {
       klog.V(4).Infof("Skipping invalid MAC address: %q", vmMACAddr)
       continue
      }
      if vmwareOUI[vmMACAddr[:8]] {
       v1helper.AddToNodeAddresses(&addrs, v1.NodeAddress{Type: v1.NodeExternalIP, Address: ipnet.IP.String()}, v1.NodeAddress{Type: v1.NodeInternalIP, Address: ipnet.IP.String()})
       klog.V(4).Infof("Detected local IP address as %q", ipnet.IP.String())
      } else {
       klog.Warningf("Failed to patch IP as MAC address %q does not belong to a VMware platform", vmMACAddr)
      }
     }
    }
   }
  }
 }
 return addrs, nil
}
func (vs *VSphere) getVSphereInstance(nodeName k8stypes.NodeName) (*VSphereInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vsphereIns, err := vs.nodeManager.GetVSphereInstance(nodeName)
 if err != nil {
  klog.Errorf("Cannot find node %q in cache. Node not found!!!", nodeName)
  return nil, err
 }
 return &vsphereIns, nil
}
func (vs *VSphere) getVSphereInstanceForServer(vcServer string, ctx context.Context) (*VSphereInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 vsphereIns, ok := vs.vsphereInstanceMap[vcServer]
 if !ok {
  klog.Errorf("cannot find vcServer %q in cache. VC not found!!!", vcServer)
  return nil, errors.New(fmt.Sprintf("Cannot find node %q in vsphere configuration map", vcServer))
 }
 err := vs.nodeManager.vcConnect(ctx, vsphereIns)
 if err != nil {
  klog.Errorf("failed connecting to vcServer %q with error %+v", vcServer, err)
  return nil, err
 }
 return vsphereIns, nil
}
func (vs *VSphere) getVMFromNodeName(ctx context.Context, nodeName k8stypes.NodeName) (*vclib.VirtualMachine, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeInfo, err := vs.nodeManager.GetNodeInfo(nodeName)
 if err != nil {
  return nil, err
 }
 return nodeInfo.vm, nil
}
func (vs *VSphere) NodeAddresses(ctx context.Context, nodeName k8stypes.NodeName) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if vs.hostName == convertToString(nodeName) {
  addrs, err := getLocalIP()
  if err != nil {
   return nil, err
  }
  v1helper.AddToNodeAddresses(&addrs, v1.NodeAddress{Type: v1.NodeHostName, Address: vs.hostName})
  return addrs, nil
 }
 if vs.cfg == nil {
  return nil, cloudprovider.InstanceNotFound
 }
 addrs := []v1.NodeAddress{}
 ctx, cancel := context.WithCancel(context.Background())
 defer cancel()
 vsi, err := vs.getVSphereInstance(nodeName)
 if err != nil {
  return nil, err
 }
 err = vs.nodeManager.vcConnect(ctx, vsi)
 if err != nil {
  return nil, err
 }
 vm, err := vs.getVMFromNodeName(ctx, nodeName)
 if err != nil {
  klog.Errorf("Failed to get VM object for node: %q. err: +%v", convertToString(nodeName), err)
  return nil, err
 }
 vmMoList, err := vm.Datacenter.GetVMMoList(ctx, []*vclib.VirtualMachine{vm}, []string{"guest.net"})
 if err != nil {
  klog.Errorf("Failed to get VM Managed object with property guest.net for node: %q. err: +%v", convertToString(nodeName), err)
  return nil, err
 }
 for _, v := range vmMoList[0].Guest.Net {
  if vs.cfg.Network.PublicNetwork == v.Network {
   for _, ip := range v.IpAddress {
    if net.ParseIP(ip).To4() != nil {
     v1helper.AddToNodeAddresses(&addrs, v1.NodeAddress{Type: v1.NodeExternalIP, Address: ip}, v1.NodeAddress{Type: v1.NodeInternalIP, Address: ip})
    }
   }
  }
 }
 return addrs, nil
}
func (vs *VSphere) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return vs.NodeAddresses(ctx, convertToK8sType(providerID))
}
func (vs *VSphere) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.NotImplemented
}
func (vs *VSphere) CurrentNodeName(ctx context.Context, hostname string) (k8stypes.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return convertToK8sType(vs.hostName), nil
}
func convertToString(nodeName k8stypes.NodeName) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return string(nodeName)
}
func convertToK8sType(vmName string) k8stypes.NodeName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return k8stypes.NodeName(vmName)
}
func (vs *VSphere) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName, err := vs.GetNodeNameFromProviderID(providerID)
 if err != nil {
  klog.Errorf("Error while getting nodename for providerID %s", providerID)
  return false, err
 }
 _, err = vs.InstanceID(ctx, convertToK8sType(nodeName))
 if err == nil {
  return true, nil
 }
 return false, err
}
func (vs *VSphere) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName, err := vs.GetNodeNameFromProviderID(providerID)
 if err != nil {
  klog.Errorf("Error while getting nodename for providerID %s", providerID)
  return false, err
 }
 vsi, err := vs.getVSphereInstance(convertToK8sType(nodeName))
 if err != nil {
  return false, err
 }
 if err := vs.nodeManager.vcConnect(ctx, vsi); err != nil {
  return false, err
 }
 vm, err := vs.getVMFromNodeName(ctx, convertToK8sType(nodeName))
 if err != nil {
  klog.Errorf("Failed to get VM object for node: %q. err: +%v", nodeName, err)
  return false, err
 }
 isActive, err := vm.IsActive(ctx)
 if err != nil {
  klog.Errorf("Failed to check whether node %q is active. err: %+v.", nodeName, err)
  return false, err
 }
 return !isActive, nil
}
func (vs *VSphere) InstanceID(ctx context.Context, nodeName k8stypes.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceIDInternal := func() (string, error) {
  if vs.hostName == convertToString(nodeName) {
   return vs.vmUUID, nil
  }
  if vs.cfg == nil {
   return "", fmt.Errorf("The current node can't detremine InstanceID for %q", convertToString(nodeName))
  }
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstance(nodeName)
  if err != nil {
   return "", err
  }
  err = vs.nodeManager.vcConnect(ctx, vsi)
  if err != nil {
   return "", err
  }
  vm, err := vs.getVMFromNodeName(ctx, nodeName)
  if err != nil {
   klog.Errorf("Failed to get VM object for node: %q. err: +%v", convertToString(nodeName), err)
   return "", err
  }
  isActive, err := vm.IsActive(ctx)
  if err != nil {
   klog.Errorf("Failed to check whether node %q is active. err: %+v.", convertToString(nodeName), err)
   return "", err
  }
  if isActive {
   return vs.vmUUID, nil
  }
  klog.Warningf("The VM: %s is not in %s state", convertToString(nodeName), vclib.ActivePowerState)
  return "", cloudprovider.InstanceNotFound
 }
 instanceID, err := instanceIDInternal()
 if err != nil {
  if vclib.IsManagedObjectNotFoundError(err) {
   err = vs.nodeManager.RediscoverNode(nodeName)
   if err == nil {
    klog.V(4).Infof("InstanceID: Found node %q", convertToString(nodeName))
    instanceID, err = instanceIDInternal()
   } else if err == vclib.ErrNoVMFound {
    return "", cloudprovider.InstanceNotFound
   }
  }
 }
 return instanceID, err
}
func (vs *VSphere) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "", nil
}
func (vs *VSphere) InstanceType(ctx context.Context, name k8stypes.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "", nil
}
func (vs *VSphere) Clusters() (cloudprovider.Clusters, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, true
}
func (vs *VSphere) ProviderName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ProviderName
}
func (vs *VSphere) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, false
}
func (vs *VSphere) Zones() (cloudprovider.Zones, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if vs.cfg == nil {
  klog.V(1).Info("The vSphere cloud provider does not support zones")
  return nil, false
 }
 return vs, true
}
func (vs *VSphere) Routes() (cloudprovider.Routes, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, false
}
func (vs *VSphere) AttachDisk(vmDiskPath string, storagePolicyName string, nodeName k8stypes.NodeName) (diskUUID string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 attachDiskInternal := func(vmDiskPath string, storagePolicyName string, nodeName k8stypes.NodeName) (diskUUID string, err error) {
  if nodeName == "" {
   nodeName = convertToK8sType(vs.hostName)
  }
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstance(nodeName)
  if err != nil {
   return "", err
  }
  err = vs.nodeManager.vcConnect(ctx, vsi)
  if err != nil {
   return "", err
  }
  vm, err := vs.getVMFromNodeName(ctx, nodeName)
  if err != nil {
   klog.Errorf("Failed to get VM object for node: %q. err: +%v", convertToString(nodeName), err)
   return "", err
  }
  diskUUID, err = vm.AttachDisk(ctx, vmDiskPath, &vclib.VolumeOptions{SCSIControllerType: vclib.PVSCSIControllerType, StoragePolicyName: storagePolicyName})
  if err != nil {
   klog.Errorf("Failed to attach disk: %s for node: %s. err: +%v", vmDiskPath, convertToString(nodeName), err)
   return "", err
  }
  return diskUUID, nil
 }
 requestTime := time.Now()
 diskUUID, err = attachDiskInternal(vmDiskPath, storagePolicyName, nodeName)
 if err != nil {
  if vclib.IsManagedObjectNotFoundError(err) {
   err = vs.nodeManager.RediscoverNode(nodeName)
   if err == nil {
    klog.V(4).Infof("AttachDisk: Found node %q", convertToString(nodeName))
    diskUUID, err = attachDiskInternal(vmDiskPath, storagePolicyName, nodeName)
    klog.V(4).Infof("AttachDisk: Retry: diskUUID %s, err +%v", diskUUID, err)
   }
  }
 }
 klog.V(4).Infof("AttachDisk executed for node %s and volume %s with diskUUID %s. Err: %s", convertToString(nodeName), vmDiskPath, diskUUID, err)
 vclib.RecordvSphereMetric(vclib.OperationAttachVolume, requestTime, err)
 return diskUUID, err
}
func (vs *VSphere) DetachDisk(volPath string, nodeName k8stypes.NodeName) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 detachDiskInternal := func(volPath string, nodeName k8stypes.NodeName) error {
  if nodeName == "" {
   nodeName = convertToK8sType(vs.hostName)
  }
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstance(nodeName)
  if err != nil {
   if err == vclib.ErrNoVMFound {
    klog.Infof("Node %q does not exist, disk %s is already detached from node.", convertToString(nodeName), volPath)
    return nil
   }
   return err
  }
  err = vs.nodeManager.vcConnect(ctx, vsi)
  if err != nil {
   return err
  }
  vm, err := vs.getVMFromNodeName(ctx, nodeName)
  if err != nil {
   if err == vclib.ErrNoVMFound {
    klog.Infof("Node %q does not exist, disk %s is already detached from node.", convertToString(nodeName), volPath)
    return nil
   }
   klog.Errorf("Failed to get VM object for node: %q. err: +%v", convertToString(nodeName), err)
   return err
  }
  err = vm.DetachDisk(ctx, volPath)
  if err != nil {
   klog.Errorf("Failed to detach disk: %s for node: %s. err: +%v", volPath, convertToString(nodeName), err)
   return err
  }
  return nil
 }
 requestTime := time.Now()
 err := detachDiskInternal(volPath, nodeName)
 if err != nil {
  if vclib.IsManagedObjectNotFoundError(err) {
   err = vs.nodeManager.RediscoverNode(nodeName)
   if err == nil {
    err = detachDiskInternal(volPath, nodeName)
   }
  }
 }
 vclib.RecordvSphereMetric(vclib.OperationDetachVolume, requestTime, err)
 return err
}
func (vs *VSphere) DiskIsAttached(volPath string, nodeName k8stypes.NodeName) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 diskIsAttachedInternal := func(volPath string, nodeName k8stypes.NodeName) (bool, error) {
  var vSphereInstance string
  if nodeName == "" {
   vSphereInstance = vs.hostName
   nodeName = convertToK8sType(vSphereInstance)
  } else {
   vSphereInstance = convertToString(nodeName)
  }
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstance(nodeName)
  if err != nil {
   return false, err
  }
  err = vs.nodeManager.vcConnect(ctx, vsi)
  if err != nil {
   return false, err
  }
  vm, err := vs.getVMFromNodeName(ctx, nodeName)
  if err != nil {
   if err == vclib.ErrNoVMFound {
    klog.Warningf("Node %q does not exist, vsphere CP will assume disk %v is not attached to it.", nodeName, volPath)
    return false, nil
   }
   klog.Errorf("Failed to get VM object for node: %q. err: +%v", vSphereInstance, err)
   return false, err
  }
  volPath = vclib.RemoveStorageClusterORFolderNameFromVDiskPath(volPath)
  attached, err := vm.IsDiskAttached(ctx, volPath)
  if err != nil {
   klog.Errorf("DiskIsAttached failed to determine whether disk %q is still attached on node %q", volPath, vSphereInstance)
  }
  klog.V(4).Infof("DiskIsAttached result: %v and error: %q, for volume: %q", attached, err, volPath)
  return attached, err
 }
 requestTime := time.Now()
 isAttached, err := diskIsAttachedInternal(volPath, nodeName)
 if err != nil {
  if vclib.IsManagedObjectNotFoundError(err) {
   err = vs.nodeManager.RediscoverNode(nodeName)
   if err == vclib.ErrNoVMFound {
    isAttached, err = false, nil
   } else if err == nil {
    isAttached, err = diskIsAttachedInternal(volPath, nodeName)
   }
  }
 }
 vclib.RecordvSphereMetric(vclib.OperationDiskIsAttached, requestTime, err)
 return isAttached, err
}
func (vs *VSphere) DisksAreAttached(nodeVolumes map[k8stypes.NodeName][]string) (map[k8stypes.NodeName]map[string]bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disksAreAttachedInternal := func(nodeVolumes map[k8stypes.NodeName][]string) (map[k8stypes.NodeName]map[string]bool, error) {
  disksAreAttach := func(ctx context.Context, nodeVolumes map[k8stypes.NodeName][]string, attached map[string]map[string]bool, retry bool) ([]k8stypes.NodeName, error) {
   var wg sync.WaitGroup
   var localAttachedMaps []map[string]map[string]bool
   var nodesToRetry []k8stypes.NodeName
   var globalErr error
   globalErr = nil
   globalErrMutex := &sync.Mutex{}
   nodesToRetryMutex := &sync.Mutex{}
   dcNodes := make(map[string][]k8stypes.NodeName)
   for nodeName := range nodeVolumes {
    nodeInfo, err := vs.nodeManager.GetNodeInfo(nodeName)
    if err != nil {
     klog.Errorf("Failed to get node info: %+v. err: %+v", nodeInfo.vm, err)
     return nodesToRetry, err
    }
    VC_DC := nodeInfo.vcServer + nodeInfo.dataCenter.String()
    dcNodes[VC_DC] = append(dcNodes[VC_DC], nodeName)
   }
   for _, nodes := range dcNodes {
    localAttachedMap := make(map[string]map[string]bool)
    localAttachedMaps = append(localAttachedMaps, localAttachedMap)
    go func() {
     nodesToRetryLocal, err := vs.checkDiskAttached(ctx, nodes, nodeVolumes, localAttachedMap, retry)
     if err != nil {
      if !vclib.IsManagedObjectNotFoundError(err) {
       globalErrMutex.Lock()
       globalErr = err
       globalErrMutex.Unlock()
       klog.Errorf("Failed to check disk attached for nodes: %+v. err: %+v", nodes, err)
      }
     }
     nodesToRetryMutex.Lock()
     nodesToRetry = append(nodesToRetry, nodesToRetryLocal...)
     nodesToRetryMutex.Unlock()
     wg.Done()
    }()
    wg.Add(1)
   }
   wg.Wait()
   if globalErr != nil {
    return nodesToRetry, globalErr
   }
   for _, localAttachedMap := range localAttachedMaps {
    for key, value := range localAttachedMap {
     attached[key] = value
    }
   }
   return nodesToRetry, nil
  }
  klog.V(4).Infof("Starting DisksAreAttached API for vSphere with nodeVolumes: %+v", nodeVolumes)
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  disksAttached := make(map[k8stypes.NodeName]map[string]bool)
  if len(nodeVolumes) == 0 {
   return disksAttached, nil
  }
  vmVolumes, err := vs.convertVolPathsToDevicePaths(ctx, nodeVolumes)
  if err != nil {
   klog.Errorf("Failed to convert volPaths to devicePaths: %+v. err: %+v", nodeVolumes, err)
   return nil, err
  }
  attached := make(map[string]map[string]bool)
  nodesToRetry, err := disksAreAttach(ctx, vmVolumes, attached, false)
  if err != nil {
   return nil, err
  }
  if len(nodesToRetry) != 0 {
   remainingNodesVolumes := make(map[k8stypes.NodeName][]string)
   for _, nodeName := range nodesToRetry {
    err = vs.nodeManager.RediscoverNode(nodeName)
    if err != nil {
     if err == vclib.ErrNoVMFound {
      klog.V(4).Infof("node %s not found. err: %+v", nodeName, err)
      continue
     }
     klog.Errorf("Failed to rediscover node %s. err: %+v", nodeName, err)
     return nil, err
    }
    remainingNodesVolumes[nodeName] = nodeVolumes[nodeName]
   }
   if len(remainingNodesVolumes) != 0 {
    nodesToRetry, err = disksAreAttach(ctx, remainingNodesVolumes, attached, true)
    if err != nil || len(nodesToRetry) != 0 {
     klog.Errorf("Failed to retry disksAreAttach  for nodes %+v. err: %+v", remainingNodesVolumes, err)
     return nil, err
    }
   }
   for nodeName, volPaths := range attached {
    disksAttached[convertToK8sType(nodeName)] = volPaths
   }
  }
  klog.V(4).Infof("DisksAreAttach successfully executed. result: %+v", attached)
  return disksAttached, nil
 }
 requestTime := time.Now()
 attached, err := disksAreAttachedInternal(nodeVolumes)
 vclib.RecordvSphereMetric(vclib.OperationDisksAreAttached, requestTime, err)
 return attached, err
}
func (vs *VSphere) CreateVolume(volumeOptions *vclib.VolumeOptions) (canonicalVolumePath string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(1).Infof("Starting to create a vSphere volume with volumeOptions: %+v", volumeOptions)
 createVolumeInternal := func(volumeOptions *vclib.VolumeOptions) (canonicalVolumePath string, err error) {
  var datastore string
  if volumeOptions.Datastore == "" {
   datastore = vs.cfg.Workspace.DefaultDatastore
  } else {
   datastore = volumeOptions.Datastore
  }
  datastore = strings.TrimSpace(datastore)
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstanceForServer(vs.cfg.Workspace.VCenterIP, ctx)
  if err != nil {
   return "", err
  }
  dc, err := vclib.GetDatacenter(ctx, vsi.conn, vs.cfg.Workspace.Datacenter)
  if err != nil {
   return "", err
  }
  var vmOptions *vclib.VMOptions
  if volumeOptions.VSANStorageProfileData != "" || volumeOptions.StoragePolicyName != "" {
   cleanUpDummyVMLock.RLock()
   defer cleanUpDummyVMLock.RUnlock()
   cleanUpRoutineInitLock.Lock()
   if !cleanUpRoutineInitialized {
    klog.V(1).Infof("Starting a clean up routine to remove stale dummy VM's")
    go vs.cleanUpDummyVMs(DummyVMPrefixName)
    cleanUpRoutineInitialized = true
   }
   cleanUpRoutineInitLock.Unlock()
   vmOptions, err = vs.setVMOptions(ctx, dc, vs.cfg.Workspace.ResourcePoolPath)
   if err != nil {
    klog.Errorf("Failed to set VM options requires to create a vsphere volume. err: %+v", err)
    return "", err
   }
  }
  if volumeOptions.StoragePolicyName != "" && volumeOptions.Datastore == "" {
   datastore, err = getPbmCompatibleDatastore(ctx, dc, volumeOptions.StoragePolicyName, vs.nodeManager)
   if err != nil {
    klog.Errorf("Failed to get pbm compatible datastore with storagePolicy: %s. err: %+v", volumeOptions.StoragePolicyName, err)
    return "", err
   }
  } else {
   sharedDsList, err := getSharedDatastoresInK8SCluster(ctx, dc, vs.nodeManager)
   if err != nil {
    klog.Errorf("Failed to get shared datastore: %+v", err)
    return "", err
   }
   found := false
   for _, sharedDs := range sharedDsList {
    if datastore == sharedDs.Info.Name {
     found = true
     break
    }
   }
   if !found {
    msg := fmt.Sprintf("The specified datastore %s is not a shared datastore across node VMs", datastore)
    return "", errors.New(msg)
   }
  }
  ds, err := dc.GetDatastoreByName(ctx, datastore)
  if err != nil {
   return "", err
  }
  volumeOptions.Datastore = datastore
  kubeVolsPath := filepath.Clean(ds.Path(VolDir)) + "/"
  err = ds.CreateDirectory(ctx, kubeVolsPath, false)
  if err != nil && err != vclib.ErrFileAlreadyExist {
   klog.Errorf("Cannot create dir %#v. err %s", kubeVolsPath, err)
   return "", err
  }
  volumePath := kubeVolsPath + volumeOptions.Name + ".vmdk"
  disk := diskmanagers.VirtualDisk{DiskPath: volumePath, VolumeOptions: volumeOptions, VMOptions: vmOptions}
  volumePath, err = disk.Create(ctx, ds)
  if err != nil {
   klog.Errorf("Failed to create a vsphere volume with volumeOptions: %+v on datastore: %s. err: %+v", volumeOptions, datastore, err)
   return "", err
  }
  canonicalVolumePath, err = getcanonicalVolumePath(ctx, dc, volumePath)
  if err != nil {
   klog.Errorf("Failed to get canonical vsphere volume path for volume: %s with volumeOptions: %+v on datastore: %s. err: %+v", volumePath, volumeOptions, datastore, err)
   return "", err
  }
  if filepath.Base(datastore) != datastore {
   canonicalVolumePath = strings.Replace(canonicalVolumePath, filepath.Base(datastore), datastore, 1)
  }
  return canonicalVolumePath, nil
 }
 requestTime := time.Now()
 canonicalVolumePath, err = createVolumeInternal(volumeOptions)
 vclib.RecordCreateVolumeMetric(volumeOptions, requestTime, err)
 klog.V(4).Infof("The canonical volume path for the newly created vSphere volume is %q", canonicalVolumePath)
 return canonicalVolumePath, err
}
func (vs *VSphere) DeleteVolume(vmDiskPath string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(1).Infof("Starting to delete vSphere volume with vmDiskPath: %s", vmDiskPath)
 deleteVolumeInternal := func(vmDiskPath string) error {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  vsi, err := vs.getVSphereInstanceForServer(vs.cfg.Workspace.VCenterIP, ctx)
  if err != nil {
   return err
  }
  dc, err := vclib.GetDatacenter(ctx, vsi.conn, vs.cfg.Workspace.Datacenter)
  if err != nil {
   return err
  }
  disk := diskmanagers.VirtualDisk{DiskPath: vmDiskPath, VolumeOptions: &vclib.VolumeOptions{}, VMOptions: &vclib.VMOptions{}}
  err = disk.Delete(ctx, dc)
  if err != nil {
   klog.Errorf("Failed to delete vsphere volume with vmDiskPath: %s. err: %+v", vmDiskPath, err)
  }
  return err
 }
 requestTime := time.Now()
 err := deleteVolumeInternal(vmDiskPath)
 vclib.RecordvSphereMetric(vclib.OperationDeleteVolume, requestTime, err)
 return err
}
func (vs *VSphere) HasClusterID() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (vs *VSphere) NodeAdded(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, ok := obj.(*v1.Node)
 if node == nil || !ok {
  klog.Warningf("NodeAdded: unrecognized object %+v", obj)
  return
 }
 klog.V(4).Infof("Node added: %+v", node)
 vs.nodeManager.RegisterNode(node)
}
func (vs *VSphere) NodeDeleted(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, ok := obj.(*v1.Node)
 if node == nil || !ok {
  klog.Warningf("NodeDeleted: unrecognized object %+v", obj)
  return
 }
 klog.V(4).Infof("Node deleted: %+v", node)
 vs.nodeManager.UnRegisterNode(node)
}
func (vs *VSphere) NodeManager() (nodeManager *NodeManager) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if vs == nil {
  return nil
 }
 return vs.nodeManager
}
func withTagsClient(ctx context.Context, connection *vclib.VSphereConnection, f func(c *rest.Client) error) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := rest.NewClient(connection.Client)
 user := url.UserPassword(connection.Username, connection.Password)
 if err := c.Login(ctx, user); err != nil {
  return err
 }
 defer c.Logout(ctx)
 return f(c)
}
func (vs *VSphere) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName, err := vs.CurrentNodeName(ctx, vs.hostName)
 if err != nil {
  klog.Errorf("Cannot get node name.")
  return cloudprovider.Zone{}, err
 }
 zone := cloudprovider.Zone{}
 vsi, err := vs.getVSphereInstanceForServer(vs.cfg.Workspace.VCenterIP, ctx)
 if err != nil {
  klog.Errorf("Cannot connent to vsphere. Get zone for node %s error", nodeName)
  return cloudprovider.Zone{}, err
 }
 dc, err := vclib.GetDatacenter(ctx, vsi.conn, vs.cfg.Workspace.Datacenter)
 if err != nil {
  klog.Errorf("Cannot connent to datacenter. Get zone for node %s error", nodeName)
  return cloudprovider.Zone{}, err
 }
 vmHost, err := dc.GetHostByVMUUID(ctx, vs.vmUUID)
 if err != nil {
  klog.Errorf("Cannot find VM runtime host. Get zone for node %s error", nodeName)
  return cloudprovider.Zone{}, err
 }
 pc := vsi.conn.Client.ServiceContent.PropertyCollector
 err = withTagsClient(ctx, vsi.conn, func(c *rest.Client) error {
  client := tags.NewManager(c)
  objects, err := mo.Ancestors(ctx, vsi.conn.Client, pc, *vmHost)
  if err != nil {
   return err
  }
  for i := range objects {
   obj := objects[len(objects)-1-i]
   tags, err := client.ListAttachedTags(ctx, obj)
   if err != nil {
    klog.Errorf("Cannot list attached tags. Get zone for node %s: %s", nodeName, err)
    return err
   }
   for _, value := range tags {
    tag, err := client.GetTag(ctx, value)
    if err != nil {
     klog.Errorf("Get tag %s: %s", value, err)
     return err
    }
    category, err := client.GetCategory(ctx, tag.CategoryID)
    if err != nil {
     klog.Errorf("Get category %s error", value)
     return err
    }
    found := func() {
     klog.Errorf("Found %q tag (%s) for %s attached to %s", category.Name, tag.Name, vs.vmUUID, obj.Reference())
    }
    switch {
    case category.Name == vs.cfg.Labels.Zone:
     zone.FailureDomain = tag.Name
     found()
    case category.Name == vs.cfg.Labels.Region:
     zone.Region = tag.Name
     found()
    }
    if zone.FailureDomain != "" && zone.Region != "" {
     return nil
    }
   }
  }
  if zone.Region == "" {
   if vs.cfg.Labels.Region != "" {
    return fmt.Errorf("vSphere region category %q does not match any tags for node %s [%s]", vs.cfg.Labels.Region, nodeName, vs.vmUUID)
   }
  }
  if zone.FailureDomain == "" {
   if vs.cfg.Labels.Zone != "" {
    return fmt.Errorf("vSphere zone category %q does not match any tags for node %s [%s]", vs.cfg.Labels.Zone, nodeName, vs.vmUUID)
   }
  }
  return nil
 })
 if err != nil {
  klog.Errorf("Get zone for node %s: %s", nodeName, err)
  return cloudprovider.Zone{}, err
 }
 return zone, nil
}
func (vs *VSphere) GetZoneByNodeName(ctx context.Context, nodeName k8stypes.NodeName) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.Zone{}, cloudprovider.NotImplemented
}
func (vs *VSphere) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.Zone{}, cloudprovider.NotImplemented
}
