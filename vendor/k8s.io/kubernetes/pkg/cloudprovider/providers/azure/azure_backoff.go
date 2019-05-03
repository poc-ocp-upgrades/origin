package azure

import (
 "context"
 "fmt"
 "net/http"
 "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
 "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/wait"
 cloudprovider "k8s.io/cloud-provider"
)

func (az *Cloud) requestBackoff() (resourceRequestBackoff wait.Backoff) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if az.CloudProviderBackoff {
  return az.resourceRequestBackoff
 }
 resourceRequestBackoff = wait.Backoff{Steps: 1}
 return resourceRequestBackoff
}
func (az *Cloud) Event(obj runtime.Object, eventtype, reason, message string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj != nil && reason != "" {
  az.eventRecorder.Event(obj, eventtype, reason, message)
 }
}
func (az *Cloud) GetVirtualMachineWithRetry(name types.NodeName) (compute.VirtualMachine, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var machine compute.VirtualMachine
 var retryErr error
 err := wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  machine, retryErr = az.getVirtualMachine(name)
  if retryErr == cloudprovider.InstanceNotFound {
   return true, cloudprovider.InstanceNotFound
  }
  if retryErr != nil {
   klog.Errorf("GetVirtualMachineWithRetry(%s): backoff failure, will retry, err=%v", name, retryErr)
   return false, nil
  }
  klog.V(2).Infof("GetVirtualMachineWithRetry(%s): backoff success", name)
  return true, nil
 })
 if err == wait.ErrWaitTimeout {
  err = retryErr
 }
 return machine, err
}
func (az *Cloud) VirtualMachineClientListWithRetry(resourceGroup string) ([]compute.VirtualMachine, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allNodes := []compute.VirtualMachine{}
 err := wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  var retryErr error
  ctx, cancel := getContextWithCancel()
  defer cancel()
  allNodes, retryErr = az.VirtualMachinesClient.List(ctx, resourceGroup)
  if retryErr != nil {
   klog.Errorf("VirtualMachinesClient.List(%v) - backoff: failure, will retry,err=%v", resourceGroup, retryErr)
   return false, retryErr
  }
  klog.V(2).Infof("VirtualMachinesClient.List(%v) - backoff: success", resourceGroup)
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return allNodes, err
}
func (az *Cloud) GetIPForMachineWithRetry(name types.NodeName) (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ip, publicIP string
 err := wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  var retryErr error
  ip, publicIP, retryErr = az.getIPForMachine(name)
  if retryErr != nil {
   klog.Errorf("GetIPForMachineWithRetry(%s): backoff failure, will retry,err=%v", name, retryErr)
   return false, nil
  }
  klog.V(2).Infof("GetIPForMachineWithRetry(%s): backoff success", name)
  return true, nil
 })
 return ip, publicIP, err
}
func (az *Cloud) CreateOrUpdateSGWithRetry(service *v1.Service, sg network.SecurityGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.SecurityGroupsClient.CreateOrUpdate(ctx, az.ResourceGroup, *sg.Name, sg)
  klog.V(10).Infof("SecurityGroupsClient.CreateOrUpdate(%s): end", *sg.Name)
  done, err := az.processHTTPRetryResponse(service, "CreateOrUpdateSecurityGroup", resp, err)
  if done && err == nil {
   az.nsgCache.Delete(*sg.Name)
  }
  return done, err
 })
}
func (az *Cloud) CreateOrUpdateLBWithRetry(service *v1.Service, lb network.LoadBalancer) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.LoadBalancerClient.CreateOrUpdate(ctx, az.ResourceGroup, *lb.Name, lb)
  klog.V(10).Infof("LoadBalancerClient.CreateOrUpdate(%s): end", *lb.Name)
  done, err := az.processHTTPRetryResponse(service, "CreateOrUpdateLoadBalancer", resp, err)
  if done && err == nil {
   az.lbCache.Delete(*lb.Name)
  }
  return done, err
 })
}
func (az *Cloud) ListLBWithRetry(service *v1.Service) ([]network.LoadBalancer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var allLBs []network.LoadBalancer
 err := wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  var retryErr error
  ctx, cancel := getContextWithCancel()
  defer cancel()
  allLBs, retryErr = az.LoadBalancerClient.List(ctx, az.ResourceGroup)
  if retryErr != nil {
   az.Event(service, v1.EventTypeWarning, "ListLoadBalancers", retryErr.Error())
   klog.Errorf("LoadBalancerClient.List(%v) - backoff: failure, will retry,err=%v", az.ResourceGroup, retryErr)
   return false, retryErr
  }
  klog.V(2).Infof("LoadBalancerClient.List(%v) - backoff: success", az.ResourceGroup)
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return allLBs, nil
}
func (az *Cloud) ListPIPWithRetry(service *v1.Service, pipResourceGroup string) ([]network.PublicIPAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var allPIPs []network.PublicIPAddress
 err := wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  var retryErr error
  ctx, cancel := getContextWithCancel()
  defer cancel()
  allPIPs, retryErr = az.PublicIPAddressesClient.List(ctx, pipResourceGroup)
  if retryErr != nil {
   az.Event(service, v1.EventTypeWarning, "ListPublicIPs", retryErr.Error())
   klog.Errorf("PublicIPAddressesClient.List(%v) - backoff: failure, will retry,err=%v", pipResourceGroup, retryErr)
   return false, retryErr
  }
  klog.V(2).Infof("PublicIPAddressesClient.List(%v) - backoff: success", pipResourceGroup)
  return true, nil
 })
 if err != nil {
  return nil, err
 }
 return allPIPs, nil
}
func (az *Cloud) CreateOrUpdatePIPWithRetry(service *v1.Service, pipResourceGroup string, pip network.PublicIPAddress) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.PublicIPAddressesClient.CreateOrUpdate(ctx, pipResourceGroup, *pip.Name, pip)
  klog.V(10).Infof("PublicIPAddressesClient.CreateOrUpdate(%s, %s): end", pipResourceGroup, *pip.Name)
  return az.processHTTPRetryResponse(service, "CreateOrUpdatePublicIPAddress", resp, err)
 })
}
func (az *Cloud) CreateOrUpdateInterfaceWithRetry(service *v1.Service, nic network.Interface) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.InterfacesClient.CreateOrUpdate(ctx, az.ResourceGroup, *nic.Name, nic)
  klog.V(10).Infof("InterfacesClient.CreateOrUpdate(%s): end", *nic.Name)
  return az.processHTTPRetryResponse(service, "CreateOrUpdateInterface", resp, err)
 })
}
func (az *Cloud) DeletePublicIPWithRetry(service *v1.Service, pipResourceGroup string, pipName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.PublicIPAddressesClient.Delete(ctx, pipResourceGroup, pipName)
  return az.processHTTPRetryResponse(service, "DeletePublicIPAddress", resp, err)
 })
}
func (az *Cloud) DeleteLBWithRetry(service *v1.Service, lbName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.LoadBalancerClient.Delete(ctx, az.ResourceGroup, lbName)
  done, err := az.processHTTPRetryResponse(service, "DeleteLoadBalancer", resp, err)
  if done && err == nil {
   az.lbCache.Delete(lbName)
  }
  return done, err
 })
}
func (az *Cloud) CreateOrUpdateRouteTableWithRetry(routeTable network.RouteTable) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.RouteTablesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, routeTable)
  return az.processHTTPRetryResponse(nil, "", resp, err)
 })
}
func (az *Cloud) CreateOrUpdateRouteWithRetry(route network.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.RoutesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, *route.Name, route)
  klog.V(10).Infof("RoutesClient.CreateOrUpdate(%s): end", *route.Name)
  return az.processHTTPRetryResponse(nil, "", resp, err)
 })
}
func (az *Cloud) DeleteRouteWithRetry(routeName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.RoutesClient.Delete(ctx, az.ResourceGroup, az.RouteTableName, routeName)
  klog.V(10).Infof("RoutesClient.Delete(%s): end", az.RouteTableName)
  return az.processHTTPRetryResponse(nil, "", resp, err)
 })
}
func (az *Cloud) CreateOrUpdateVMWithRetry(resourceGroup, vmName string, newVM compute.VirtualMachine) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  ctx, cancel := getContextWithCancel()
  defer cancel()
  resp, err := az.VirtualMachinesClient.CreateOrUpdate(ctx, resourceGroup, vmName, newVM)
  klog.V(10).Infof("VirtualMachinesClient.CreateOrUpdate(%s): end", vmName)
  return az.processHTTPRetryResponse(nil, "", resp, err)
 })
}
func (az *Cloud) UpdateVmssVMWithRetry(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.VirtualMachineScaleSetVM) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return wait.ExponentialBackoff(az.requestBackoff(), func() (bool, error) {
  resp, err := az.VirtualMachineScaleSetVMsClient.Update(ctx, resourceGroupName, VMScaleSetName, instanceID, parameters)
  klog.V(10).Infof("VirtualMachinesClient.CreateOrUpdate(%s,%s): end", VMScaleSetName, instanceID)
  return az.processHTTPRetryResponse(nil, "", resp, err)
 })
}
func isSuccessHTTPResponse(resp http.Response) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if 199 < resp.StatusCode && resp.StatusCode < 300 {
  return true
 }
 return false
}
func shouldRetryHTTPRequest(resp *http.Response, err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err != nil {
  return true
 }
 if resp != nil {
  if 399 < resp.StatusCode && resp.StatusCode < 600 {
   return true
  }
 }
 return false
}
func (az *Cloud) processHTTPRetryResponse(service *v1.Service, reason string, resp *http.Response, err error) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if resp != nil && isSuccessHTTPResponse(*resp) {
  return true, nil
 }
 if shouldRetryHTTPRequest(resp, err) {
  if err != nil {
   az.Event(service, v1.EventTypeWarning, reason, err.Error())
   klog.Errorf("processHTTPRetryResponse: backoff failure, will retry, err=%v", err)
  } else {
   az.Event(service, v1.EventTypeWarning, reason, fmt.Sprintf("Azure HTTP response %d", resp.StatusCode))
   klog.Errorf("processHTTPRetryResponse: backoff failure, will retry, HTTP response=%d", resp.StatusCode)
  }
  return false, nil
 }
 return true, nil
}
