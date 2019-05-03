package vsphere

import (
 "context"
 "fmt"
 "strings"
 "sync"
 "k8s.io/api/core/v1"
 k8stypes "k8s.io/apimachinery/pkg/types"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
)

type NodeInfo struct {
 dataCenter *vclib.Datacenter
 vm         *vclib.VirtualMachine
 vcServer   string
 vmUUID     string
}
type NodeManager struct {
 vsphereInstanceMap    map[string]*VSphereInstance
 nodeInfoMap           map[string]*NodeInfo
 registeredNodes       map[string]*v1.Node
 credentialManager     *SecretCredentialManager
 registeredNodesLock   sync.RWMutex
 nodeInfoLock          sync.RWMutex
 credentialManagerLock sync.Mutex
}
type NodeDetails struct {
 NodeName string
 vm       *vclib.VirtualMachine
 VMUUID   string
}

const (
 POOL_SIZE  = 8
 QUEUE_SIZE = POOL_SIZE * 10
)

func (nm *NodeManager) DiscoverNode(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 type VmSearch struct {
  vc         string
  datacenter *vclib.Datacenter
 }
 var mutex = &sync.Mutex{}
 var globalErrMutex = &sync.Mutex{}
 var queueChannel chan *VmSearch
 var wg sync.WaitGroup
 var globalErr *error
 queueChannel = make(chan *VmSearch, QUEUE_SIZE)
 nodeUUID, err := GetNodeUUID(node)
 if err != nil {
  klog.Errorf("Node Discovery failed to get node uuid for node %s with error: %v", node.Name, err)
  return err
 }
 klog.V(4).Infof("Discovering node %s with uuid %s", node.ObjectMeta.Name, nodeUUID)
 vmFound := false
 globalErr = nil
 setGlobalErr := func(err error) {
  globalErrMutex.Lock()
  globalErr = &err
  globalErrMutex.Unlock()
 }
 setVMFound := func(found bool) {
  mutex.Lock()
  vmFound = found
  mutex.Unlock()
 }
 getVMFound := func() bool {
  mutex.Lock()
  found := vmFound
  mutex.Unlock()
  return found
 }
 go func() {
  var datacenterObjs []*vclib.Datacenter
  for vc, vsi := range nm.vsphereInstanceMap {
   found := getVMFound()
   if found == true {
    break
   }
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   err := nm.vcConnect(ctx, vsi)
   if err != nil {
    klog.V(4).Info("Discovering node error vc:", err)
    setGlobalErr(err)
    continue
   }
   if vsi.cfg.Datacenters == "" {
    datacenterObjs, err = vclib.GetAllDatacenter(ctx, vsi.conn)
    if err != nil {
     klog.V(4).Info("Discovering node error dc:", err)
     setGlobalErr(err)
     continue
    }
   } else {
    datacenters := strings.Split(vsi.cfg.Datacenters, ",")
    for _, dc := range datacenters {
     dc = strings.TrimSpace(dc)
     if dc == "" {
      continue
     }
     datacenterObj, err := vclib.GetDatacenter(ctx, vsi.conn, dc)
     if err != nil {
      klog.V(4).Info("Discovering node error dc:", err)
      setGlobalErr(err)
      continue
     }
     datacenterObjs = append(datacenterObjs, datacenterObj)
    }
   }
   for _, datacenterObj := range datacenterObjs {
    found := getVMFound()
    if found == true {
     break
    }
    klog.V(4).Infof("Finding node %s in vc=%s and datacenter=%s", node.Name, vc, datacenterObj.Name())
    queueChannel <- &VmSearch{vc: vc, datacenter: datacenterObj}
   }
  }
  close(queueChannel)
 }()
 for i := 0; i < POOL_SIZE; i++ {
  go func() {
   for res := range queueChannel {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    vm, err := res.datacenter.GetVMByUUID(ctx, nodeUUID)
    if err != nil {
     klog.V(4).Infof("Error while looking for vm=%+v in vc=%s and datacenter=%s: %v", vm, res.vc, res.datacenter.Name(), err)
     if err != vclib.ErrNoVMFound {
      setGlobalErr(err)
     } else {
      klog.V(4).Infof("Did not find node %s in vc=%s and datacenter=%s", node.Name, res.vc, res.datacenter.Name())
     }
     continue
    }
    if vm != nil {
     klog.V(4).Infof("Found node %s as vm=%+v in vc=%s and datacenter=%s", node.Name, vm, res.vc, res.datacenter.Name())
     nodeInfo := &NodeInfo{dataCenter: res.datacenter, vm: vm, vcServer: res.vc, vmUUID: nodeUUID}
     nm.addNodeInfo(node.ObjectMeta.Name, nodeInfo)
     for range queueChannel {
     }
     setVMFound(true)
     break
    }
   }
   wg.Done()
  }()
  wg.Add(1)
 }
 wg.Wait()
 if vmFound {
  return nil
 }
 if globalErr != nil {
  return *globalErr
 }
 klog.V(4).Infof("Discovery Node: %q vm not found", node.Name)
 return vclib.ErrNoVMFound
}
func (nm *NodeManager) RegisterNode(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.addNode(node)
 nm.DiscoverNode(node)
 return nil
}
func (nm *NodeManager) UnRegisterNode(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.removeNode(node)
 return nil
}
func (nm *NodeManager) RediscoverNode(nodeName k8stypes.NodeName) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node, err := nm.GetNode(nodeName)
 if err != nil {
  return err
 }
 return nm.DiscoverNode(&node)
}
func (nm *NodeManager) GetNode(nodeName k8stypes.NodeName) (v1.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.registeredNodesLock.RLock()
 node := nm.registeredNodes[convertToString(nodeName)]
 nm.registeredNodesLock.RUnlock()
 if node == nil {
  return v1.Node{}, vclib.ErrNoVMFound
 }
 return *node, nil
}
func (nm *NodeManager) addNode(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.registeredNodesLock.Lock()
 nm.registeredNodes[node.ObjectMeta.Name] = node
 nm.registeredNodesLock.Unlock()
}
func (nm *NodeManager) removeNode(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.registeredNodesLock.Lock()
 delete(nm.registeredNodes, node.ObjectMeta.Name)
 nm.registeredNodesLock.Unlock()
 nm.nodeInfoLock.Lock()
 delete(nm.nodeInfoMap, node.ObjectMeta.Name)
 nm.nodeInfoLock.Unlock()
}
func (nm *NodeManager) GetNodeInfo(nodeName k8stypes.NodeName) (NodeInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getNodeInfo := func(nodeName k8stypes.NodeName) *NodeInfo {
  nm.nodeInfoLock.RLock()
  nodeInfo := nm.nodeInfoMap[convertToString(nodeName)]
  nm.nodeInfoLock.RUnlock()
  return nodeInfo
 }
 nodeInfo := getNodeInfo(nodeName)
 var err error
 if nodeInfo == nil {
  klog.V(4).Infof("No VM found for node %q. Initiating rediscovery.", convertToString(nodeName))
  err = nm.RediscoverNode(nodeName)
  if err != nil {
   klog.Errorf("Error %q node info for node %q not found", err, convertToString(nodeName))
   return NodeInfo{}, err
  }
  nodeInfo = getNodeInfo(nodeName)
 } else {
  klog.V(4).Infof("Renewing NodeInfo %+v for node %q", nodeInfo, convertToString(nodeName))
  nodeInfo, err = nm.renewNodeInfo(nodeInfo, true)
  if err != nil {
   klog.Errorf("Error %q occurred while renewing NodeInfo for %q", err, convertToString(nodeName))
   return NodeInfo{}, err
  }
  nm.addNodeInfo(convertToString(nodeName), nodeInfo)
 }
 return *nodeInfo, nil
}
func (nm *NodeManager) GetNodeDetails() ([]NodeDetails, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.registeredNodesLock.Lock()
 defer nm.registeredNodesLock.Unlock()
 var nodeDetails []NodeDetails
 for nodeName, nodeObj := range nm.registeredNodes {
  nodeInfo, err := nm.GetNodeInfoWithNodeObject(nodeObj)
  if err != nil {
   return nil, err
  }
  klog.V(4).Infof("Updated NodeInfo %v for node %q.", nodeInfo, nodeName)
  nodeDetails = append(nodeDetails, NodeDetails{nodeName, nodeInfo.vm, nodeInfo.vmUUID})
 }
 return nodeDetails, nil
}
func (nm *NodeManager) addNodeInfo(nodeName string, nodeInfo *NodeInfo) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.nodeInfoLock.Lock()
 nm.nodeInfoMap[nodeName] = nodeInfo
 nm.nodeInfoLock.Unlock()
}
func (nm *NodeManager) GetVSphereInstance(nodeName k8stypes.NodeName) (VSphereInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeInfo, err := nm.GetNodeInfo(nodeName)
 if err != nil {
  klog.V(4).Infof("node info for node %q not found", convertToString(nodeName))
  return VSphereInstance{}, err
 }
 vsphereInstance := nm.vsphereInstanceMap[nodeInfo.vcServer]
 if vsphereInstance == nil {
  return VSphereInstance{}, fmt.Errorf("vSphereInstance for vc server %q not found while looking for node %q", nodeInfo.vcServer, convertToString(nodeName))
 }
 return *vsphereInstance, nil
}
func (nm *NodeManager) renewNodeInfo(nodeInfo *NodeInfo, reconnect bool) (*NodeInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := context.WithCancel(context.Background())
 defer cancel()
 vsphereInstance := nm.vsphereInstanceMap[nodeInfo.vcServer]
 if vsphereInstance == nil {
  err := fmt.Errorf("vSphereInstance for vSphere %q not found while refershing NodeInfo for VM %q", nodeInfo.vcServer, nodeInfo.vm)
  return nil, err
 }
 if reconnect {
  err := nm.vcConnect(ctx, vsphereInstance)
  if err != nil {
   return nil, err
  }
 }
 vm := nodeInfo.vm.RenewVM(vsphereInstance.conn.Client)
 return &NodeInfo{vm: &vm, dataCenter: vm.Datacenter, vcServer: nodeInfo.vcServer, vmUUID: nodeInfo.vmUUID}, nil
}
func (nodeInfo *NodeInfo) VM() *vclib.VirtualMachine {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if nodeInfo == nil {
  return nil
 }
 return nodeInfo.vm
}
func (nm *NodeManager) vcConnect(ctx context.Context, vsphereInstance *VSphereInstance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := vsphereInstance.conn.Connect(ctx)
 if err == nil {
  return nil
 }
 credentialManager := nm.CredentialManager()
 if !vclib.IsInvalidCredentialsError(err) || credentialManager == nil {
  klog.Errorf("Cannot connect to vCenter with err: %v", err)
  return err
 }
 klog.V(4).Infof("Invalid credentials. Cannot connect to server %q. Fetching credentials from secrets.", vsphereInstance.conn.Hostname)
 credentials, err := credentialManager.GetCredential(vsphereInstance.conn.Hostname)
 if err != nil {
  klog.Errorf("Failed to get credentials from Secret Credential Manager with err: %v", err)
  return err
 }
 vsphereInstance.conn.UpdateCredentials(credentials.User, credentials.Password)
 return vsphereInstance.conn.Connect(ctx)
}
func (nm *NodeManager) GetNodeInfoWithNodeObject(node *v1.Node) (NodeInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName := node.Name
 getNodeInfo := func(nodeName string) *NodeInfo {
  nm.nodeInfoLock.RLock()
  nodeInfo := nm.nodeInfoMap[nodeName]
  nm.nodeInfoLock.RUnlock()
  return nodeInfo
 }
 nodeInfo := getNodeInfo(nodeName)
 var err error
 if nodeInfo == nil {
  klog.V(4).Infof("No VM found for node %q. Initiating rediscovery.", nodeName)
  err = nm.DiscoverNode(node)
  if err != nil {
   klog.Errorf("Error %q node info for node %q not found", err, nodeName)
   return NodeInfo{}, err
  }
  nodeInfo = getNodeInfo(nodeName)
 } else {
  klog.V(4).Infof("Renewing NodeInfo %+v for node %q", nodeInfo, nodeName)
  nodeInfo, err = nm.renewNodeInfo(nodeInfo, true)
  if err != nil {
   klog.Errorf("Error %q occurred while renewing NodeInfo for %q", err, nodeName)
   return NodeInfo{}, err
  }
  nm.addNodeInfo(nodeName, nodeInfo)
 }
 return *nodeInfo, nil
}
func (nm *NodeManager) CredentialManager() *SecretCredentialManager {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.credentialManagerLock.Lock()
 defer nm.credentialManagerLock.Unlock()
 return nm.credentialManager
}
func (nm *NodeManager) UpdateCredentialManager(credentialManager *SecretCredentialManager) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nm.credentialManagerLock.Lock()
 defer nm.credentialManagerLock.Unlock()
 nm.credentialManager = credentialManager
}
