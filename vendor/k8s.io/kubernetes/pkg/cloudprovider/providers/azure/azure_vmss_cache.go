package azure

import (
 "fmt"
 "strings"
 "time"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/util/sets"
)

var (
 vmssNameSeparator                 = "_"
 vmssCacheSeparator                = "#"
 nodeNameToScaleSetMappingKey      = "k8sNodeNameToScaleSetMappingKey"
 availabilitySetNodesKey           = "k8sAvailabilitySetNodesKey"
 vmssCacheTTL                      = time.Minute
 vmssVMCacheTTL                    = time.Minute
 availabilitySetNodesCacheTTL      = 5 * time.Minute
 nodeNameToScaleSetMappingCacheTTL = 5 * time.Minute
)

type nodeNameToScaleSetMapping map[string]string

func (ss *scaleSet) makeVmssVMName(scaleSetName, instanceID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%s%s%s", scaleSetName, vmssNameSeparator, instanceID)
}
func extractVmssVMName(name string) (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 split := strings.SplitAfter(name, vmssNameSeparator)
 if len(split) < 2 {
  klog.V(3).Infof("Failed to extract vmssVMName %q", name)
  return "", "", ErrorNotVmssInstance
 }
 ssName := strings.Join(split[0:len(split)-1], "")
 ssName = ssName[:len(ssName)-1]
 instanceID := split[len(split)-1]
 return ssName, instanceID, nil
}
func (ss *scaleSet) newVmssCache() (*timedCache, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getter := func(key string) (interface{}, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  result, err := ss.VirtualMachineScaleSetsClient.Get(ctx, ss.ResourceGroup, key)
  exists, message, realErr := checkResourceExistsFromError(err)
  if realErr != nil {
   return nil, realErr
  }
  if !exists {
   klog.V(2).Infof("Virtual machine scale set %q not found with message: %q", key, message)
   return nil, nil
  }
  return &result, nil
 }
 return newTimedcache(vmssCacheTTL, getter)
}
func (ss *scaleSet) newNodeNameToScaleSetMappingCache() (*timedCache, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getter := func(key string) (interface{}, error) {
  localCache := make(nodeNameToScaleSetMapping)
  allResourceGroups, err := ss.GetResourceGroups()
  if err != nil {
   return nil, err
  }
  for _, resourceGroup := range allResourceGroups.List() {
   scaleSetNames, err := ss.listScaleSets(resourceGroup)
   if err != nil {
    return nil, err
   }
   for _, ssName := range scaleSetNames {
    vms, err := ss.listScaleSetVMs(ssName, resourceGroup)
    if err != nil {
     return nil, err
    }
    for _, vm := range vms {
     if vm.OsProfile == nil || vm.OsProfile.ComputerName == nil {
      klog.Warningf("failed to get computerName for vmssVM (%q)", ssName)
      continue
     }
     computerName := strings.ToLower(*vm.OsProfile.ComputerName)
     localCache[computerName] = ssName
    }
   }
  }
  return localCache, nil
 }
 return newTimedcache(nodeNameToScaleSetMappingCacheTTL, getter)
}
func (ss *scaleSet) newAvailabilitySetNodesCache() (*timedCache, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getter := func(key string) (interface{}, error) {
  localCache := sets.NewString()
  resourceGroups, err := ss.GetResourceGroups()
  if err != nil {
   return nil, err
  }
  for _, resourceGroup := range resourceGroups.List() {
   vmList, err := ss.Cloud.VirtualMachineClientListWithRetry(resourceGroup)
   if err != nil {
    return nil, err
   }
   for _, vm := range vmList {
    if vm.Name != nil {
     localCache.Insert(*vm.Name)
    }
   }
  }
  return localCache, nil
 }
 return newTimedcache(availabilitySetNodesCacheTTL, getter)
}
func buildVmssCacheKey(resourceGroup, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%s%s%s", resourceGroup, vmssCacheSeparator, name)
}
func extractVmssCacheKey(key string) (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 keyItems := strings.Split(key, vmssCacheSeparator)
 if len(keyItems) != 2 {
  return "", "", fmt.Errorf("key %q is not in format '<resouceGroup>#<vmName>'", key)
 }
 resourceGroup := keyItems[0]
 vmName := keyItems[1]
 return resourceGroup, vmName, nil
}
func (ss *scaleSet) newVmssVMCache() (*timedCache, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getter := func(key string) (interface{}, error) {
  resourceGroup, vmName, err := extractVmssCacheKey(key)
  if err != nil {
   return nil, err
  }
  ssName, instanceID, err := extractVmssVMName(vmName)
  if err != nil {
   return nil, err
  }
  if ssName == "" {
   return nil, nil
  }
  ctx, cancel := getContextWithCancel()
  defer cancel()
  result, err := ss.VirtualMachineScaleSetVMsClient.Get(ctx, resourceGroup, ssName, instanceID)
  exists, message, realErr := checkResourceExistsFromError(err)
  if realErr != nil {
   return nil, realErr
  }
  if !exists {
   klog.V(2).Infof("Virtual machine scale set VM %q not found with message: %q", key, message)
   return nil, nil
  }
  if result.InstanceView == nil {
   viewCtx, viewCancel := getContextWithCancel()
   defer viewCancel()
   view, err := ss.VirtualMachineScaleSetVMsClient.GetInstanceView(viewCtx, resourceGroup, ssName, instanceID)
   exists, message, realErr = checkResourceExistsFromError(err)
   if realErr != nil {
    return nil, realErr
   }
   if !exists {
    klog.V(2).Infof("Virtual machine scale set VM %q not found with message: %q", key, message)
    return nil, nil
   }
   result.InstanceView = &view
  }
  return &result, nil
 }
 return newTimedcache(vmssVMCacheTTL, getter)
}
func (ss *scaleSet) getScaleSetNameByNodeName(nodeName string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 getScaleSetName := func(nodeName string) (string, error) {
  nodeNameMapping, err := ss.nodeNameToScaleSetMappingCache.Get(nodeNameToScaleSetMappingKey)
  if err != nil {
   return "", err
  }
  realMapping := nodeNameMapping.(nodeNameToScaleSetMapping)
  if ssName, ok := realMapping[nodeName]; ok {
   return ssName, nil
  }
  return "", nil
 }
 ssName, err := getScaleSetName(nodeName)
 if err != nil {
  return "", err
 }
 if ssName != "" {
  return ssName, nil
 }
 ss.nodeNameToScaleSetMappingCache.Delete(nodeNameToScaleSetMappingKey)
 return getScaleSetName(nodeName)
}
func (ss *scaleSet) isNodeManagedByAvailabilitySet(nodeName string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cached, err := ss.availabilitySetNodesCache.Get(availabilitySetNodesKey)
 if err != nil {
  return false, err
 }
 availabilitySetNodes := cached.(sets.String)
 return availabilitySetNodes.Has(nodeName), nil
}
