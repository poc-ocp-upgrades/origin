package gce

import (
 "fmt"
 "net/http"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
)

type addressManager struct {
 logPrefix   string
 svc         CloudAddressService
 name        string
 serviceName string
 targetIP    string
 addressType cloud.LbScheme
 region      string
 subnetURL   string
 tryRelease  bool
}

func newAddressManager(svc CloudAddressService, serviceName, region, subnetURL, name, targetIP string, addressType cloud.LbScheme) *addressManager {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &addressManager{svc: svc, logPrefix: fmt.Sprintf("AddressManager(%q)", name), region: region, serviceName: serviceName, name: name, targetIP: targetIP, addressType: addressType, tryRelease: true, subnetURL: subnetURL}
}
func (am *addressManager) HoldAddress() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("%v: attempting hold of IP %q Type %q", am.logPrefix, am.targetIP, am.addressType)
 addr, err := am.svc.GetRegionAddress(am.name, am.region)
 if err != nil && !isNotFound(err) {
  return "", err
 }
 if addr != nil {
  validationError := am.validateAddress(addr)
  if validationError == nil {
   klog.V(4).Infof("%v: address %q already reserves IP %q Type %q. No further action required.", am.logPrefix, addr.Name, addr.Address, addr.AddressType)
   return addr.Address, nil
  }
  klog.V(2).Infof("%v: deleting existing address because %v", am.logPrefix, validationError)
  err := am.svc.DeleteRegionAddress(addr.Name, am.region)
  if err != nil {
   if isNotFound(err) {
    klog.V(4).Infof("%v: address %q was not found. Ignoring.", am.logPrefix, addr.Name)
   } else {
    return "", err
   }
  } else {
   klog.V(4).Infof("%v: successfully deleted previous address %q", am.logPrefix, addr.Name)
  }
 }
 return am.ensureAddressReservation()
}
func (am *addressManager) ReleaseAddress() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !am.tryRelease {
  klog.V(4).Infof("%v: not attempting release of address %q.", am.logPrefix, am.targetIP)
  return nil
 }
 klog.V(4).Infof("%v: releasing address %q named %q", am.logPrefix, am.targetIP, am.name)
 err := am.svc.DeleteRegionAddress(am.name, am.region)
 if err != nil {
  if isNotFound(err) {
   klog.Warningf("%v: address %q was not found. Ignoring.", am.logPrefix, am.name)
   return nil
  }
  return err
 }
 klog.V(4).Infof("%v: successfully released IP %q named %q", am.logPrefix, am.targetIP, am.name)
 return nil
}
func (am *addressManager) ensureAddressReservation() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newAddr := &compute.Address{Name: am.name, Description: fmt.Sprintf(`{"kubernetes.io/service-name":"%s"}`, am.serviceName), Address: am.targetIP, AddressType: string(am.addressType), Subnetwork: am.subnetURL}
 reserveErr := am.svc.ReserveRegionAddress(newAddr, am.region)
 if reserveErr == nil {
  if newAddr.Address != "" {
   klog.V(4).Infof("%v: successfully reserved IP %q with name %q", am.logPrefix, newAddr.Address, newAddr.Name)
   return newAddr.Address, nil
  }
  addr, err := am.svc.GetRegionAddress(newAddr.Name, am.region)
  if err != nil {
   return "", err
  }
  klog.V(4).Infof("%v: successfully created address %q which reserved IP %q", am.logPrefix, addr.Name, addr.Address)
  return addr.Address, nil
 } else if !isHTTPErrorCode(reserveErr, http.StatusConflict) && !isHTTPErrorCode(reserveErr, http.StatusBadRequest) {
  return "", reserveErr
 }
 if am.targetIP == "" {
  return "", fmt.Errorf("failed to reserve address %q with no specific IP, err: %v", am.name, reserveErr)
 }
 addr, err := am.svc.GetRegionAddressByIP(am.region, am.targetIP)
 if err != nil {
  return "", fmt.Errorf("failed to get address by IP %q after reservation attempt, err: %q, reservation err: %q", am.targetIP, err, reserveErr)
 }
 if err := am.validateAddress(addr); err != nil {
  return "", err
 }
 if am.isManagedAddress(addr) {
  klog.Warningf("%v: address %q unexpectedly existed with IP %q.", am.logPrefix, addr.Name, am.targetIP)
 } else {
  klog.V(4).Infof("%v: address %q was already reserved with name: %q, description: %q", am.logPrefix, am.targetIP, addr.Name, addr.Description)
  am.tryRelease = false
 }
 return addr.Address, nil
}
func (am *addressManager) validateAddress(addr *compute.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if am.targetIP != "" && am.targetIP != addr.Address {
  return fmt.Errorf("address %q does not have the expected IP %q, actual: %q", addr.Name, am.targetIP, addr.Address)
 }
 if addr.AddressType != string(am.addressType) {
  return fmt.Errorf("address %q does not have the expected address type %q, actual: %q", addr.Name, am.addressType, addr.AddressType)
 }
 return nil
}
func (am *addressManager) isManagedAddress(addr *compute.Address) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return addr.Name == am.name
}
func ensureAddressDeleted(svc CloudAddressService, name, region string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ignoreNotFound(svc.DeleteRegionAddress(name, region))
}
