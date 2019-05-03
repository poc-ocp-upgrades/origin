package reconcilers

import (
 "fmt"
 "net"
 "path"
 "sync"
 "time"
 "k8s.io/klog"
 corev1 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 kruntime "k8s.io/apimachinery/pkg/runtime"
 apirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/storage"
 corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
 endpointsv1 "k8s.io/kubernetes/pkg/api/v1/endpoints"
)

type Leases interface {
 ListLeases() ([]string, error)
 UpdateLease(ip string) error
 RemoveLease(ip string) error
}
type storageLeases struct {
 storage   storage.Interface
 baseKey   string
 leaseTime time.Duration
}

var _ Leases = &storageLeases{}

func (s *storageLeases) ListLeases() ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ipInfoList := &corev1.EndpointsList{}
 if err := s.storage.List(apirequest.NewDefaultContext(), s.baseKey, "0", storage.Everything, ipInfoList); err != nil {
  return nil, err
 }
 ipList := make([]string, len(ipInfoList.Items))
 for i, ip := range ipInfoList.Items {
  ipList[i] = ip.Subsets[0].Addresses[0].IP
 }
 klog.V(6).Infof("Current master IPs listed in storage are %v", ipList)
 return ipList, nil
}
func (s *storageLeases) UpdateLease(ip string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := path.Join(s.baseKey, ip)
 return s.storage.GuaranteedUpdate(apirequest.NewDefaultContext(), key, &corev1.Endpoints{}, true, nil, func(input kruntime.Object, respMeta storage.ResponseMeta) (kruntime.Object, *uint64, error) {
  existing := input.(*corev1.Endpoints)
  existing.Subsets = []corev1.EndpointSubset{{Addresses: []corev1.EndpointAddress{{IP: ip}}}}
  leaseTime := uint64(s.leaseTime / time.Second)
  existing.Generation++
  klog.V(6).Infof("Resetting TTL on master IP %q listed in storage to %v", ip, leaseTime)
  return existing, &leaseTime, nil
 })
}
func (s *storageLeases) RemoveLease(ip string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return s.storage.Delete(apirequest.NewDefaultContext(), s.baseKey+"/"+ip, &corev1.Endpoints{}, nil)
}
func NewLeases(storage storage.Interface, baseKey string, leaseTime time.Duration) Leases {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storageLeases{storage: storage, baseKey: baseKey, leaseTime: leaseTime}
}

type leaseEndpointReconciler struct {
 endpointClient        corev1client.EndpointsGetter
 masterLeases          Leases
 stopReconcilingCalled bool
 reconcilingLock       sync.Mutex
}

func NewLeaseEndpointReconciler(endpointClient corev1client.EndpointsGetter, masterLeases Leases) EndpointReconciler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &leaseEndpointReconciler{endpointClient: endpointClient, masterLeases: masterLeases, stopReconcilingCalled: false}
}
func (r *leaseEndpointReconciler) ReconcileEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.reconcilingLock.Lock()
 defer r.reconcilingLock.Unlock()
 if r.stopReconcilingCalled {
  return nil
 }
 if err := r.masterLeases.UpdateLease(ip.String()); err != nil {
  return err
 }
 return r.doReconcile(serviceName, endpointPorts, reconcilePorts)
}
func (r *leaseEndpointReconciler) doReconcile(serviceName string, endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 e, err := r.endpointClient.Endpoints(corev1.NamespaceDefault).Get(serviceName, metav1.GetOptions{})
 shouldCreate := false
 if err != nil {
  if !errors.IsNotFound(err) {
   return err
  }
  shouldCreate = true
  e = &corev1.Endpoints{ObjectMeta: metav1.ObjectMeta{Name: serviceName, Namespace: corev1.NamespaceDefault}}
 }
 masterIPs, err := r.masterLeases.ListLeases()
 if err != nil {
  return err
 }
 if len(masterIPs) == 0 {
  return fmt.Errorf("no master IPs were listed in storage, refusing to erase all endpoints for the kubernetes service")
 }
 formatCorrect, ipCorrect, portsCorrect := checkEndpointSubsetFormatWithLease(e, masterIPs, endpointPorts, reconcilePorts)
 if formatCorrect && ipCorrect && portsCorrect {
  return nil
 }
 if !formatCorrect {
  e.Subsets = []corev1.EndpointSubset{{Addresses: []corev1.EndpointAddress{}, Ports: endpointPorts}}
 }
 if !formatCorrect || !ipCorrect {
  e.Subsets[0].Addresses = make([]corev1.EndpointAddress, len(masterIPs))
  for ind, ip := range masterIPs {
   e.Subsets[0].Addresses[ind] = corev1.EndpointAddress{IP: ip}
  }
  e.Subsets = endpointsv1.RepackSubsets(e.Subsets)
 }
 if !portsCorrect {
  e.Subsets[0].Ports = endpointPorts
 }
 klog.Warningf("Resetting endpoints for master service %q to %v", serviceName, masterIPs)
 if shouldCreate {
  if _, err = r.endpointClient.Endpoints(corev1.NamespaceDefault).Create(e); errors.IsAlreadyExists(err) {
   err = nil
  }
 } else {
  _, err = r.endpointClient.Endpoints(corev1.NamespaceDefault).Update(e)
 }
 return err
}
func checkEndpointSubsetFormatWithLease(e *corev1.Endpoints, expectedIPs []string, ports []corev1.EndpointPort, reconcilePorts bool) (formatCorrect bool, ipsCorrect bool, portsCorrect bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(e.Subsets) != 1 {
  return false, false, false
 }
 sub := &e.Subsets[0]
 portsCorrect = true
 if reconcilePorts {
  if len(sub.Ports) != len(ports) {
   portsCorrect = false
  } else {
   for i, port := range ports {
    if port != sub.Ports[i] {
     portsCorrect = false
     break
    }
   }
  }
 }
 ipsCorrect = true
 if len(sub.Addresses) != len(expectedIPs) {
  ipsCorrect = false
 } else {
  presentAddrs := make(map[string]bool, len(expectedIPs))
  for _, ip := range expectedIPs {
   presentAddrs[ip] = false
  }
  for _, addr := range sub.Addresses {
   if alreadySeen, ok := presentAddrs[addr.IP]; alreadySeen || !ok {
    ipsCorrect = false
    break
   }
   presentAddrs[addr.IP] = true
  }
 }
 return true, ipsCorrect, portsCorrect
}
func (r *leaseEndpointReconciler) RemoveEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := r.masterLeases.RemoveLease(ip.String()); err != nil {
  return err
 }
 return r.doReconcile(serviceName, endpointPorts, true)
}
func (r *leaseEndpointReconciler) StopReconciling() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.reconcilingLock.Lock()
 defer r.reconcilingLock.Unlock()
 r.stopReconcilingCalled = true
}
