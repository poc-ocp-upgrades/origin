package winuserspace

import (
 "fmt"
 "net"
 "strconv"
 "strings"
 "sync"
 "sync/atomic"
 "time"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 utilnet "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/proxy"
 "k8s.io/kubernetes/pkg/util/netsh"
)

const allAvailableInterfaces string = ""

type portal struct {
 ip         string
 port       int
 isExternal bool
}
type serviceInfo struct {
 isAliveAtomic       int32
 portal              portal
 protocol            v1.Protocol
 socket              proxySocket
 timeout             time.Duration
 activeClients       *clientCache
 dnsClients          *dnsClientCache
 sessionAffinityType v1.ServiceAffinity
}

func (info *serviceInfo) setAlive(b bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var i int32
 if b {
  i = 1
 }
 atomic.StoreInt32(&info.isAliveAtomic, i)
}
func (info *serviceInfo) isAlive() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return atomic.LoadInt32(&info.isAliveAtomic) != 0
}
func logTimeout(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if e, ok := err.(net.Error); ok {
  if e.Timeout() {
   klog.V(3).Infof("connection to endpoint closed due to inactivity")
   return true
  }
 }
 return false
}

type Proxier struct {
 loadBalancer   LoadBalancer
 mu             sync.Mutex
 serviceMap     map[ServicePortPortalName]*serviceInfo
 syncPeriod     time.Duration
 udpIdleTimeout time.Duration
 portMapMutex   sync.Mutex
 portMap        map[portMapKey]*portMapValue
 numProxyLoops  int32
 netsh          netsh.Interface
 hostIP         net.IP
}

var _ proxy.ProxyProvider = &Proxier{}

type portMapKey struct {
 ip       string
 port     int
 protocol v1.Protocol
}

func (k *portMapKey) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%s/%s", net.JoinHostPort(k.ip, strconv.Itoa(k.port)), k.protocol)
}

type portMapValue struct {
 owner  ServicePortPortalName
 socket interface{ Close() error }
}

var (
 ErrProxyOnLocalhost = fmt.Errorf("cannot proxy on localhost")
)
var localhostIPv4 = net.ParseIP("127.0.0.1")
var localhostIPv6 = net.ParseIP("::1")

func NewProxier(loadBalancer LoadBalancer, listenIP net.IP, netsh netsh.Interface, pr utilnet.PortRange, syncPeriod, udpIdleTimeout time.Duration) (*Proxier, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if listenIP.Equal(localhostIPv4) || listenIP.Equal(localhostIPv6) {
  return nil, ErrProxyOnLocalhost
 }
 hostIP, err := utilnet.ChooseHostInterface()
 if err != nil {
  return nil, fmt.Errorf("failed to select a host interface: %v", err)
 }
 klog.V(2).Infof("Setting proxy IP to %v", hostIP)
 return createProxier(loadBalancer, listenIP, netsh, hostIP, syncPeriod, udpIdleTimeout)
}
func createProxier(loadBalancer LoadBalancer, listenIP net.IP, netsh netsh.Interface, hostIP net.IP, syncPeriod, udpIdleTimeout time.Duration) (*Proxier, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &Proxier{loadBalancer: loadBalancer, serviceMap: make(map[ServicePortPortalName]*serviceInfo), portMap: make(map[portMapKey]*portMapValue), syncPeriod: syncPeriod, udpIdleTimeout: udpIdleTimeout, netsh: netsh, hostIP: hostIP}, nil
}
func (proxier *Proxier) Sync() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.cleanupStaleStickySessions()
}
func (proxier *Proxier) SyncLoop() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 t := time.NewTicker(proxier.syncPeriod)
 defer t.Stop()
 for {
  <-t.C
  klog.V(6).Infof("Periodic sync")
  proxier.Sync()
 }
}
func (proxier *Proxier) cleanupStaleStickySessions() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 defer proxier.mu.Unlock()
 servicePortNameMap := make(map[proxy.ServicePortName]bool)
 for name := range proxier.serviceMap {
  servicePortName := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: name.Namespace, Name: name.Name}, Port: name.Port}
  if servicePortNameMap[servicePortName] == false {
   servicePortNameMap[servicePortName] = true
   proxier.loadBalancer.CleanupStaleStickySessions(servicePortName)
  }
 }
}
func (proxier *Proxier) stopProxy(service ServicePortPortalName, info *serviceInfo) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 defer proxier.mu.Unlock()
 return proxier.stopProxyInternal(service, info)
}
func (proxier *Proxier) stopProxyInternal(service ServicePortPortalName, info *serviceInfo) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 delete(proxier.serviceMap, service)
 info.setAlive(false)
 err := info.socket.Close()
 return err
}
func (proxier *Proxier) getServiceInfo(service ServicePortPortalName) (*serviceInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 defer proxier.mu.Unlock()
 info, ok := proxier.serviceMap[service]
 return info, ok
}
func (proxier *Proxier) setServiceInfo(service ServicePortPortalName, info *serviceInfo) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 defer proxier.mu.Unlock()
 proxier.serviceMap[service] = info
}
func (proxier *Proxier) addServicePortPortal(servicePortPortalName ServicePortPortalName, protocol v1.Protocol, listenIP string, port int, timeout time.Duration) (*serviceInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var serviceIP net.IP
 if listenIP != allAvailableInterfaces {
  if serviceIP = net.ParseIP(listenIP); serviceIP == nil {
   return nil, fmt.Errorf("could not parse ip '%q'", listenIP)
  }
  args := proxier.netshIpv4AddressAddArgs(serviceIP)
  if existed, err := proxier.netsh.EnsureIPAddress(args, serviceIP); err != nil {
   return nil, err
  } else if !existed {
   klog.V(3).Infof("Added ip address to fowarder interface for service %q at %s/%s", servicePortPortalName, net.JoinHostPort(listenIP, strconv.Itoa(port)), protocol)
  }
 }
 sock, err := newProxySocket(protocol, serviceIP, port)
 if err != nil {
  return nil, err
 }
 si := &serviceInfo{isAliveAtomic: 1, portal: portal{ip: listenIP, port: port, isExternal: false}, protocol: protocol, socket: sock, timeout: timeout, activeClients: newClientCache(), dnsClients: newDNSClientCache(), sessionAffinityType: v1.ServiceAffinityNone}
 proxier.setServiceInfo(servicePortPortalName, si)
 klog.V(2).Infof("Proxying for service %q at %s/%s", servicePortPortalName, net.JoinHostPort(listenIP, strconv.Itoa(port)), protocol)
 go func(service ServicePortPortalName, proxier *Proxier) {
  defer runtime.HandleCrash()
  atomic.AddInt32(&proxier.numProxyLoops, 1)
  sock.ProxyLoop(service, si, proxier)
  atomic.AddInt32(&proxier.numProxyLoops, -1)
 }(servicePortPortalName, proxier)
 return si, nil
}
func (proxier *Proxier) closeServicePortPortal(servicePortPortalName ServicePortPortalName, info *serviceInfo) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := proxier.stopProxy(servicePortPortalName, info); err != nil {
  return err
 }
 if info.portal.ip != allAvailableInterfaces {
  serviceIP := net.ParseIP(info.portal.ip)
  args := proxier.netshIpv4AddressDeleteArgs(serviceIP)
  if err := proxier.netsh.DeleteIPAddress(args); err != nil {
   return err
  }
 }
 return nil
}
func getListenIPPortMap(service *v1.Service, listenPort int, nodePort int) map[string]int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 listenIPPortMap := make(map[string]int)
 listenIPPortMap[service.Spec.ClusterIP] = listenPort
 for _, ip := range service.Spec.ExternalIPs {
  listenIPPortMap[ip] = listenPort
 }
 for _, ingress := range service.Status.LoadBalancer.Ingress {
  listenIPPortMap[ingress.IP] = listenPort
 }
 if nodePort != 0 {
  listenIPPortMap[allAvailableInterfaces] = nodePort
 }
 return listenIPPortMap
}
func (proxier *Proxier) mergeService(service *v1.Service) map[ServicePortPortalName]bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if service == nil {
  return nil
 }
 svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
 if !helper.IsServiceIPSet(service) {
  klog.V(3).Infof("Skipping service %s due to clusterIP = %q", svcName, service.Spec.ClusterIP)
  return nil
 }
 existingPortPortals := make(map[ServicePortPortalName]bool)
 for i := range service.Spec.Ports {
  servicePort := &service.Spec.Ports[i]
  listenIPPortMap := getListenIPPortMap(service, int(servicePort.Port), int(servicePort.NodePort))
  protocol := servicePort.Protocol
  for listenIP, listenPort := range listenIPPortMap {
   servicePortPortalName := ServicePortPortalName{NamespacedName: svcName, Port: servicePort.Name, PortalIPName: listenIP}
   existingPortPortals[servicePortPortalName] = true
   info, exists := proxier.getServiceInfo(servicePortPortalName)
   if exists && sameConfig(info, service, protocol, listenPort) {
    continue
   }
   if exists {
    klog.V(4).Infof("Something changed for service %q: stopping it", servicePortPortalName)
    if err := proxier.closeServicePortPortal(servicePortPortalName, info); err != nil {
     klog.Errorf("Failed to close service port portal %q: %v", servicePortPortalName, err)
    }
   }
   klog.V(1).Infof("Adding new service %q at %s/%s", servicePortPortalName, net.JoinHostPort(listenIP, strconv.Itoa(listenPort)), protocol)
   info, err := proxier.addServicePortPortal(servicePortPortalName, protocol, listenIP, listenPort, proxier.udpIdleTimeout)
   if err != nil {
    klog.Errorf("Failed to start proxy for %q: %v", servicePortPortalName, err)
    continue
   }
   info.sessionAffinityType = service.Spec.SessionAffinity
   klog.V(10).Infof("info: %#v", info)
  }
  if len(listenIPPortMap) > 0 {
   servicePortName := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: service.Namespace, Name: service.Name}, Port: servicePort.Name}
   timeoutSeconds := 0
   if service.Spec.SessionAffinity == v1.ServiceAffinityClientIP {
    timeoutSeconds = int(*service.Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds)
   }
   proxier.loadBalancer.NewService(servicePortName, service.Spec.SessionAffinity, timeoutSeconds)
  }
 }
 return existingPortPortals
}
func (proxier *Proxier) unmergeService(service *v1.Service, existingPortPortals map[ServicePortPortalName]bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if service == nil {
  return
 }
 svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
 if !helper.IsServiceIPSet(service) {
  klog.V(3).Infof("Skipping service %s due to clusterIP = %q", svcName, service.Spec.ClusterIP)
  return
 }
 servicePortNameMap := make(map[proxy.ServicePortName]bool)
 for name := range existingPortPortals {
  servicePortName := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: name.Namespace, Name: name.Name}, Port: name.Port}
  servicePortNameMap[servicePortName] = true
 }
 for i := range service.Spec.Ports {
  servicePort := &service.Spec.Ports[i]
  serviceName := proxy.ServicePortName{NamespacedName: svcName, Port: servicePort.Name}
  listenIPPortMap := getListenIPPortMap(service, int(servicePort.Port), int(servicePort.NodePort))
  for listenIP := range listenIPPortMap {
   servicePortPortalName := ServicePortPortalName{NamespacedName: svcName, Port: servicePort.Name, PortalIPName: listenIP}
   if existingPortPortals[servicePortPortalName] {
    continue
   }
   klog.V(1).Infof("Stopping service %q", servicePortPortalName)
   info, exists := proxier.getServiceInfo(servicePortPortalName)
   if !exists {
    klog.Errorf("Service %q is being removed but doesn't exist", servicePortPortalName)
    continue
   }
   if err := proxier.closeServicePortPortal(servicePortPortalName, info); err != nil {
    klog.Errorf("Failed to close service port portal %q: %v", servicePortPortalName, err)
   }
  }
  if !servicePortNameMap[serviceName] {
   proxier.loadBalancer.DeleteService(serviceName)
  }
 }
}
func (proxier *Proxier) OnServiceAdd(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = proxier.mergeService(service)
}
func (proxier *Proxier) OnServiceUpdate(oldService, service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 existingPortPortals := proxier.mergeService(service)
 proxier.unmergeService(oldService, existingPortPortals)
}
func (proxier *Proxier) OnServiceDelete(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.unmergeService(service, map[ServicePortPortalName]bool{})
}
func (proxier *Proxier) OnServiceSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func sameConfig(info *serviceInfo, service *v1.Service, protocol v1.Protocol, listenPort int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return info.protocol == protocol && info.portal.port == listenPort && info.sessionAffinityType == service.Spec.SessionAffinity
}
func isTooManyFDsError(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.Contains(err.Error(), "too many open files")
}
func isClosedError(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.HasSuffix(err.Error(), "use of closed network connection")
}
func (proxier *Proxier) netshIpv4AddressAddArgs(destIP net.IP) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 intName := proxier.netsh.GetInterfaceToAddIP()
 args := []string{"interface", "ipv4", "add", "address", "name=" + intName, "address=" + destIP.String()}
 return args
}
func (proxier *Proxier) netshIpv4AddressDeleteArgs(destIP net.IP) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 intName := proxier.netsh.GetInterfaceToAddIP()
 args := []string{"interface", "ipv4", "delete", "address", "name=" + intName, "address=" + destIP.String()}
 return args
}
