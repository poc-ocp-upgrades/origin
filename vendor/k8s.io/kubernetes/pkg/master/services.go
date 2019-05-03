package master

import (
 "fmt"
 "net"
 "k8s.io/klog"
 kubeoptions "k8s.io/kubernetes/pkg/kubeapiserver/options"
 "k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
)

func DefaultServiceIPRange(passedServiceClusterIPRange net.IPNet) (net.IPNet, net.IP, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 serviceClusterIPRange := passedServiceClusterIPRange
 if passedServiceClusterIPRange.IP == nil {
  klog.Infof("Network range for service cluster IPs is unspecified. Defaulting to %v.", kubeoptions.DefaultServiceIPCIDR)
  serviceClusterIPRange = kubeoptions.DefaultServiceIPCIDR
 }
 if size := ipallocator.RangeSize(&serviceClusterIPRange); size < 8 {
  return net.IPNet{}, net.IP{}, fmt.Errorf("The service cluster IP range must be at least %d IP addresses", 8)
 }
 apiServerServiceIP, err := ipallocator.GetIndexedIP(&serviceClusterIPRange, 1)
 if err != nil {
  return net.IPNet{}, net.IP{}, err
 }
 klog.V(4).Infof("Setting service IP to %q (read-write).", apiServerServiceIP)
 return serviceClusterIPRange, apiServerServiceIP, nil
}
