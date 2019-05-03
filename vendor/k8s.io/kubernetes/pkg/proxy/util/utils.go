package util

import (
 "context"
 "errors"
 "fmt"
 "net"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/client-go/tools/record"
 helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 utilnet "k8s.io/kubernetes/pkg/util/net"
 "k8s.io/klog"
)

const (
 IPv4ZeroCIDR = "0.0.0.0/0"
 IPv6ZeroCIDR = "::/0"
)

var (
 ErrAddressNotAllowed = errors.New("address not allowed")
 ErrNoAddresses       = errors.New("No addresses for hostname")
)

func IsZeroCIDR(cidr string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if cidr == IPv4ZeroCIDR || cidr == IPv6ZeroCIDR {
  return true
 }
 return false
}
func IsProxyableIP(ip string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 netIP := net.ParseIP(ip)
 if netIP == nil {
  return ErrAddressNotAllowed
 }
 return isProxyableIP(netIP)
}
func isProxyableIP(ip net.IP) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsInterfaceLocalMulticast() {
  return ErrAddressNotAllowed
 }
 return nil
}

type Resolver interface {
 LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error)
}

func IsProxyableHostname(ctx context.Context, resolv Resolver, hostname string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resp, err := resolv.LookupIPAddr(ctx, hostname)
 if err != nil {
  return err
 }
 if len(resp) == 0 {
  return ErrNoAddresses
 }
 for _, host := range resp {
  if err := isProxyableIP(host.IP); err != nil {
   return err
  }
 }
 return nil
}
func IsLocalIP(ip string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 addrs, err := net.InterfaceAddrs()
 if err != nil {
  return false, err
 }
 for i := range addrs {
  intf, _, err := net.ParseCIDR(addrs[i].String())
  if err != nil {
   return false, err
  }
  if net.ParseIP(ip).Equal(intf) {
   return true, nil
  }
 }
 return false, nil
}
func ShouldSkipService(svcName types.NamespacedName, service *v1.Service) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !helper.IsServiceIPSet(service) {
  klog.V(3).Infof("Skipping service %s due to clusterIP = %q", svcName, service.Spec.ClusterIP)
  return true
 }
 if service.Spec.Type == v1.ServiceTypeExternalName {
  klog.V(3).Infof("Skipping service %s due to Type=ExternalName", svcName)
  return true
 }
 return false
}
func GetNodeAddresses(cidrs []string, nw NetworkInterfacer) (sets.String, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 uniqueAddressList := sets.NewString()
 if len(cidrs) == 0 {
  uniqueAddressList.Insert(IPv4ZeroCIDR)
  uniqueAddressList.Insert(IPv6ZeroCIDR)
  return uniqueAddressList, nil
 }
 for _, cidr := range cidrs {
  if IsZeroCIDR(cidr) {
   uniqueAddressList.Insert(cidr)
  }
 }
 for _, cidr := range cidrs {
  if IsZeroCIDR(cidr) {
   continue
  }
  _, ipNet, _ := net.ParseCIDR(cidr)
  itfs, err := nw.Interfaces()
  if err != nil {
   return nil, fmt.Errorf("error listing all interfaces from host, error: %v", err)
  }
  for _, itf := range itfs {
   addrs, err := nw.Addrs(&itf)
   if err != nil {
    return nil, fmt.Errorf("error getting address from interface %s, error: %v", itf.Name, err)
   }
   for _, addr := range addrs {
    if addr == nil {
     continue
    }
    ip, _, err := net.ParseCIDR(addr.String())
    if err != nil {
     return nil, fmt.Errorf("error parsing CIDR for interface %s, error: %v", itf.Name, err)
    }
    if ipNet.Contains(ip) {
     if utilnet.IsIPv6(ip) && !uniqueAddressList.Has(IPv6ZeroCIDR) {
      uniqueAddressList.Insert(ip.String())
     }
     if !utilnet.IsIPv6(ip) && !uniqueAddressList.Has(IPv4ZeroCIDR) {
      uniqueAddressList.Insert(ip.String())
     }
    }
   }
  }
 }
 return uniqueAddressList, nil
}
func LogAndEmitIncorrectIPVersionEvent(recorder record.EventRecorder, fieldName, fieldValue, svcNamespace, svcName string, svcUID types.UID) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errMsg := fmt.Sprintf("%s in %s has incorrect IP version", fieldValue, fieldName)
 klog.Errorf("%s (service %s/%s).", errMsg, svcNamespace, svcName)
 if recorder != nil {
  recorder.Eventf(&v1.ObjectReference{Kind: "Service", Name: svcName, Namespace: svcNamespace, UID: svcUID}, v1.EventTypeWarning, "KubeProxyIncorrectIPVersion", errMsg)
 }
}
