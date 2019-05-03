package azure

import (
 "context"
 "fmt"
 "os"
 "strings"
 "k8s.io/api/core/v1"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/klog"
)

const (
 vmPowerStatePrefix      = "PowerState/"
 vmPowerStateStopped     = "stopped"
 vmPowerStateDeallocated = "deallocated"
)

func (az *Cloud) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 unmanaged, err := az.IsNodeUnmanaged(string(name))
 if err != nil {
  return nil, err
 }
 if unmanaged {
  klog.V(4).Infof("NodeAddresses: omitting unmanaged node %q", name)
  return nil, nil
 }
 addressGetter := func(nodeName types.NodeName) ([]v1.NodeAddress, error) {
  ip, publicIP, err := az.GetIPForMachineWithRetry(nodeName)
  if err != nil {
   klog.V(2).Infof("NodeAddresses(%s) abort backoff: %v", nodeName, err)
   return nil, err
  }
  addresses := []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: ip}, {Type: v1.NodeHostName, Address: string(name)}}
  if len(publicIP) > 0 {
   addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: publicIP})
  }
  return addresses, nil
 }
 if az.UseInstanceMetadata {
  metadata, err := az.metadata.GetMetadata()
  if err != nil {
   return nil, err
  }
  if metadata.Compute == nil || metadata.Network == nil {
   return nil, fmt.Errorf("failure of getting instance metadata")
  }
  isLocalInstance, err := az.isCurrentInstance(name, metadata.Compute.Name)
  if err != nil {
   return nil, err
  }
  if !isLocalInstance {
   return addressGetter(name)
  }
  if len(metadata.Network.Interface) == 0 {
   return nil, fmt.Errorf("no interface is found for the instance")
  }
  ipAddress := metadata.Network.Interface[0]
  addresses := []v1.NodeAddress{{Type: v1.NodeHostName, Address: string(name)}}
  if len(ipAddress.IPV4.IPAddress) > 0 && len(ipAddress.IPV4.IPAddress[0].PrivateIP) > 0 {
   address := ipAddress.IPV4.IPAddress[0]
   addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalIP, Address: address.PrivateIP})
   if len(address.PublicIP) > 0 {
    addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: address.PublicIP})
   }
  }
  if len(ipAddress.IPV6.IPAddress) > 0 && len(ipAddress.IPV6.IPAddress[0].PrivateIP) > 0 {
   address := ipAddress.IPV6.IPAddress[0]
   addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalIP, Address: address.PrivateIP})
   if len(address.PublicIP) > 0 {
    addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: address.PublicIP})
   }
  }
  if len(addresses) == 1 {
   az.metadata.imsCache.Delete(metadataCacheKey)
   return nil, fmt.Errorf("get empty IP addresses from instance metadata service")
  }
  return addresses, nil
 }
 return addressGetter(name)
}
func (az *Cloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if az.IsNodeUnmanagedByProviderID(providerID) {
  klog.V(4).Infof("NodeAddressesByProviderID: omitting unmanaged node %q", providerID)
  return nil, nil
 }
 name, err := az.vmSet.GetNodeNameByProviderID(providerID)
 if err != nil {
  return nil, err
 }
 return az.NodeAddresses(ctx, name)
}
func (az *Cloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if az.IsNodeUnmanagedByProviderID(providerID) {
  klog.V(4).Infof("InstanceExistsByProviderID: assuming unmanaged node %q exists", providerID)
  return true, nil
 }
 name, err := az.vmSet.GetNodeNameByProviderID(providerID)
 if err != nil {
  return false, err
 }
 _, err = az.InstanceID(ctx, name)
 if err != nil {
  if err == cloudprovider.InstanceNotFound {
   return false, nil
  }
  return false, err
 }
 return true, nil
}
func (az *Cloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName, err := az.vmSet.GetNodeNameByProviderID(providerID)
 if err != nil {
  return false, err
 }
 powerStatus, err := az.vmSet.GetPowerStatusByNodeName(string(nodeName))
 if err != nil {
  return false, err
 }
 klog.V(5).Infof("InstanceShutdownByProviderID gets power status %q for node %q", powerStatus, nodeName)
 return strings.ToLower(powerStatus) == vmPowerStateStopped || strings.ToLower(powerStatus) == vmPowerStateDeallocated, nil
}
func (az *Cloud) isCurrentInstance(name types.NodeName, metadataVMName string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 nodeName := mapNodeNameToVMName(name)
 if az.VMType == vmTypeVMSS {
  metadataVMName, err = os.Hostname()
  if err != nil {
   return false, err
  }
 }
 metadataVMName = strings.ToLower(metadataVMName)
 return (metadataVMName == nodeName), err
}
func (az *Cloud) InstanceID(ctx context.Context, name types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeName := mapNodeNameToVMName(name)
 unmanaged, err := az.IsNodeUnmanaged(nodeName)
 if err != nil {
  return "", err
 }
 if unmanaged {
  klog.V(4).Infof("InstanceID: getting ID %q for unmanaged node %q", name, name)
  return nodeName, nil
 }
 if az.UseInstanceMetadata {
  metadata, err := az.metadata.GetMetadata()
  if err != nil {
   return "", err
  }
  if metadata.Compute == nil {
   return "", fmt.Errorf("failure of getting instance metadata")
  }
  isLocalInstance, err := az.isCurrentInstance(name, metadata.Compute.Name)
  if err != nil {
   return "", err
  }
  if !isLocalInstance {
   return az.vmSet.GetInstanceIDByNodeName(nodeName)
  }
  resourceGroup := metadata.Compute.ResourceGroup
  if az.VMType == vmTypeStandard {
   return az.getStandardMachineID(resourceGroup, nodeName), nil
  }
  ssName, instanceID, err := extractVmssVMName(metadata.Compute.Name)
  if err != nil {
   if err == ErrorNotVmssInstance {
    return az.getStandardMachineID(resourceGroup, nodeName), nil
   }
   return "", err
  }
  return az.getVmssMachineID(resourceGroup, ssName, instanceID), nil
 }
 return az.vmSet.GetInstanceIDByNodeName(nodeName)
}
func (az *Cloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if az.IsNodeUnmanagedByProviderID(providerID) {
  klog.V(4).Infof("InstanceTypeByProviderID: omitting unmanaged node %q", providerID)
  return "", nil
 }
 name, err := az.vmSet.GetNodeNameByProviderID(providerID)
 if err != nil {
  return "", err
 }
 return az.InstanceType(ctx, name)
}
func (az *Cloud) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 unmanaged, err := az.IsNodeUnmanaged(string(name))
 if err != nil {
  return "", err
 }
 if unmanaged {
  klog.V(4).Infof("InstanceType: omitting unmanaged node %q", name)
  return "", nil
 }
 if az.UseInstanceMetadata {
  metadata, err := az.metadata.GetMetadata()
  if err != nil {
   return "", err
  }
  if metadata.Compute == nil {
   return "", fmt.Errorf("failure of getting instance metadata")
  }
  isLocalInstance, err := az.isCurrentInstance(name, metadata.Compute.Name)
  if err != nil {
   return "", err
  }
  if isLocalInstance {
   if metadata.Compute.VMSize != "" {
    return metadata.Compute.VMSize, nil
   }
  }
 }
 return az.vmSet.GetInstanceTypeByNodeName(string(name))
}
func (az *Cloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.NotImplemented
}
func (az *Cloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return types.NodeName(hostname), nil
}
func mapNodeNameToVMName(nodeName types.NodeName) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return string(nodeName)
}
