package userspace

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/kubernetes/pkg/proxy"
	utilproxy "k8s.io/kubernetes/pkg/proxy/util"
	"k8s.io/kubernetes/pkg/util/conntrack"
	"k8s.io/kubernetes/pkg/util/iptables"
	utilexec "k8s.io/utils/exec"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type portal struct {
	ip         net.IP
	port       int
	isExternal bool
}
type ServiceInfo struct {
	Timeout             time.Duration
	ActiveClients       *ClientCache
	isAliveAtomic       int32
	portal              portal
	protocol            v1.Protocol
	proxyPort           int
	socket              ProxySocket
	nodePort            int
	loadBalancerStatus  v1.LoadBalancerStatus
	sessionAffinityType v1.ServiceAffinity
	stickyMaxAgeSeconds int
	externalIPs         []string
}

func (info *ServiceInfo) setAlive(b bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var i int32
	if b {
		i = 1
	}
	atomic.StoreInt32(&info.isAliveAtomic, i)
}
func (info *ServiceInfo) IsAlive() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return atomic.LoadInt32(&info.isAliveAtomic) != 0
}
func logTimeout(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if e, ok := err.(net.Error); ok {
		if e.Timeout() {
			klog.V(3).Infof("connection to endpoint closed due to inactivity")
			return true
		}
	}
	return false
}

type ProxySocketFunc func(protocol v1.Protocol, ip net.IP, port int) (ProxySocket, error)
type Proxier struct {
	loadBalancer    LoadBalancer
	mu              sync.Mutex
	serviceMap      map[proxy.ServicePortName]*ServiceInfo
	syncPeriod      time.Duration
	minSyncPeriod   time.Duration
	udpIdleTimeout  time.Duration
	portMapMutex    sync.Mutex
	portMap         map[portMapKey]*portMapValue
	numProxyLoops   int32
	listenIP        net.IP
	iptables        iptables.Interface
	hostIP          net.IP
	proxyPorts      PortAllocator
	makeProxySocket ProxySocketFunc
	exec            utilexec.Interface
}

var _ proxy.ProxyProvider = &Proxier{}

type portMapKey struct {
	ip       string
	port     int
	protocol v1.Protocol
}

func (k *portMapKey) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s", net.JoinHostPort(k.ip, strconv.Itoa(k.port)), k.protocol)
}

type portMapValue struct {
	owner  proxy.ServicePortName
	socket interface{ Close() error }
}

var (
	ErrProxyOnLocalhost = fmt.Errorf("cannot proxy on localhost")
)

func IsProxyLocked(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.Contains(err.Error(), "holding the xtables lock")
}
func NewProxier(loadBalancer LoadBalancer, listenIP net.IP, iptables iptables.Interface, exec utilexec.Interface, pr utilnet.PortRange, syncPeriod, minSyncPeriod, udpIdleTimeout time.Duration, nodePortAddresses []string) (*Proxier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewCustomProxier(loadBalancer, listenIP, iptables, exec, pr, syncPeriod, minSyncPeriod, udpIdleTimeout, nodePortAddresses, newProxySocket)
}
func NewCustomProxier(loadBalancer LoadBalancer, listenIP net.IP, iptables iptables.Interface, exec utilexec.Interface, pr utilnet.PortRange, syncPeriod, minSyncPeriod, udpIdleTimeout time.Duration, nodePortAddresses []string, makeProxySocket ProxySocketFunc) (*Proxier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if listenIP.Equal(localhostIPv4) || listenIP.Equal(localhostIPv6) {
		return nil, ErrProxyOnLocalhost
	}
	var err error
	hostIP := listenIP
	if hostIP.Equal(net.IPv4zero) || hostIP.Equal(net.IPv6zero) {
		hostIP, err = utilnet.ChooseHostInterface()
		if err != nil {
			return nil, fmt.Errorf("failed to select a host interface: %v", err)
		}
	}
	err = setRLimit(64 * 1000)
	if err != nil {
		return nil, fmt.Errorf("failed to set open file handler limit: %v", err)
	}
	proxyPorts := newPortAllocator(pr)
	klog.V(2).Infof("Setting proxy IP to %v and initializing iptables", hostIP)
	return createProxier(loadBalancer, listenIP, iptables, exec, hostIP, proxyPorts, syncPeriod, minSyncPeriod, udpIdleTimeout, makeProxySocket)
}
func createProxier(loadBalancer LoadBalancer, listenIP net.IP, iptables iptables.Interface, exec utilexec.Interface, hostIP net.IP, proxyPorts PortAllocator, syncPeriod, minSyncPeriod, udpIdleTimeout time.Duration, makeProxySocket ProxySocketFunc) (*Proxier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if proxyPorts == nil {
		proxyPorts = newPortAllocator(utilnet.PortRange{})
	}
	if err := iptablesInit(iptables); err != nil {
		return nil, fmt.Errorf("failed to initialize iptables: %v", err)
	}
	if err := iptablesFlush(iptables); err != nil {
		return nil, fmt.Errorf("failed to flush iptables: %v", err)
	}
	return &Proxier{loadBalancer: loadBalancer, serviceMap: make(map[proxy.ServicePortName]*ServiceInfo), portMap: make(map[portMapKey]*portMapValue), syncPeriod: syncPeriod, minSyncPeriod: minSyncPeriod, udpIdleTimeout: udpIdleTimeout, listenIP: listenIP, iptables: iptables, hostIP: hostIP, proxyPorts: proxyPorts, makeProxySocket: makeProxySocket, exec: exec}, nil
}
func CleanupLeftovers(ipt iptables.Interface) (encounteredError bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []string{"-m", "comment", "--comment", "handle ClusterIPs; NOTE: this must be before the NodePort rules"}
	if err := ipt.DeleteRule(iptables.TableNAT, iptables.ChainOutput, append(args, "-j", string(iptablesHostPortalChain))...); err != nil {
		if !iptables.IsNotFoundError(err) {
			klog.Errorf("Error removing userspace rule: %v", err)
			encounteredError = true
		}
	}
	if err := ipt.DeleteRule(iptables.TableNAT, iptables.ChainPrerouting, append(args, "-j", string(iptablesContainerPortalChain))...); err != nil {
		if !iptables.IsNotFoundError(err) {
			klog.Errorf("Error removing userspace rule: %v", err)
			encounteredError = true
		}
	}
	args = []string{"-m", "addrtype", "--dst-type", "LOCAL"}
	args = append(args, "-m", "comment", "--comment", "handle service NodePorts; NOTE: this must be the last rule in the chain")
	if err := ipt.DeleteRule(iptables.TableNAT, iptables.ChainOutput, append(args, "-j", string(iptablesHostNodePortChain))...); err != nil {
		if !iptables.IsNotFoundError(err) {
			klog.Errorf("Error removing userspace rule: %v", err)
			encounteredError = true
		}
	}
	if err := ipt.DeleteRule(iptables.TableNAT, iptables.ChainPrerouting, append(args, "-j", string(iptablesContainerNodePortChain))...); err != nil {
		if !iptables.IsNotFoundError(err) {
			klog.Errorf("Error removing userspace rule: %v", err)
			encounteredError = true
		}
	}
	args = []string{"-m", "comment", "--comment", "Ensure that non-local NodePort traffic can flow"}
	if err := ipt.DeleteRule(iptables.TableFilter, iptables.ChainInput, append(args, "-j", string(iptablesNonLocalNodePortChain))...); err != nil {
		if !iptables.IsNotFoundError(err) {
			klog.Errorf("Error removing userspace rule: %v", err)
			encounteredError = true
		}
	}
	tableChains := map[iptables.Table][]iptables.Chain{iptables.TableNAT: {iptablesContainerPortalChain, iptablesHostPortalChain, iptablesHostNodePortChain, iptablesContainerNodePortChain}, iptables.TableFilter: {iptablesNonLocalNodePortChain}}
	for table, chains := range tableChains {
		for _, c := range chains {
			if err := ipt.FlushChain(table, c); err != nil {
				if !iptables.IsNotFoundError(err) {
					klog.Errorf("Error flushing userspace chain: %v", err)
					encounteredError = true
				}
			} else {
				if err = ipt.DeleteChain(table, c); err != nil {
					if !iptables.IsNotFoundError(err) {
						klog.Errorf("Error deleting userspace chain: %v", err)
						encounteredError = true
					}
				}
			}
		}
	}
	return encounteredError
}
func (proxier *Proxier) Sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := iptablesInit(proxier.iptables); err != nil {
		klog.Errorf("Failed to ensure iptables: %v", err)
	}
	proxier.ensurePortals()
	proxier.cleanupStaleStickySessions()
}
func (proxier *Proxier) SyncLoop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t := time.NewTicker(proxier.syncPeriod)
	defer t.Stop()
	for {
		<-t.C
		klog.V(6).Infof("Periodic sync")
		proxier.Sync()
	}
}
func (proxier *Proxier) ensurePortals() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	for name, info := range proxier.serviceMap {
		err := proxier.openPortal(name, info)
		if err != nil {
			klog.Errorf("Failed to ensure portal for %q: %v", name, err)
		}
	}
}
func (proxier *Proxier) cleanupStaleStickySessions() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	for name := range proxier.serviceMap {
		proxier.loadBalancer.CleanupStaleStickySessions(name)
	}
}
func (proxier *Proxier) stopProxy(service proxy.ServicePortName, info *ServiceInfo) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	return proxier.stopProxyInternal(service, info)
}
func (proxier *Proxier) stopProxyInternal(service proxy.ServicePortName, info *ServiceInfo) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	delete(proxier.serviceMap, service)
	info.setAlive(false)
	err := info.socket.Close()
	port := info.socket.ListenPort()
	proxier.proxyPorts.Release(port)
	return err
}
func (proxier *Proxier) getServiceInfo(service proxy.ServicePortName) (*ServiceInfo, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	info, ok := proxier.serviceMap[service]
	return info, ok
}
func (proxier *Proxier) setServiceInfo(service proxy.ServicePortName, info *ServiceInfo) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	proxier.serviceMap[service] = info
}
func (proxier *Proxier) addServiceOnPort(service proxy.ServicePortName, protocol v1.Protocol, proxyPort int, timeout time.Duration) (*ServiceInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sock, err := proxier.makeProxySocket(protocol, proxier.listenIP, proxyPort)
	if err != nil {
		return nil, err
	}
	_, portStr, err := net.SplitHostPort(sock.Addr().String())
	if err != nil {
		sock.Close()
		return nil, err
	}
	portNum, err := strconv.Atoi(portStr)
	if err != nil {
		sock.Close()
		return nil, err
	}
	si := &ServiceInfo{Timeout: timeout, ActiveClients: newClientCache(), isAliveAtomic: 1, proxyPort: portNum, protocol: protocol, socket: sock, sessionAffinityType: v1.ServiceAffinityNone}
	proxier.setServiceInfo(service, si)
	klog.V(2).Infof("Proxying for service %q on %s port %d", service, protocol, portNum)
	go func(service proxy.ServicePortName, proxier *Proxier) {
		defer runtime.HandleCrash()
		atomic.AddInt32(&proxier.numProxyLoops, 1)
		sock.ProxyLoop(service, si, proxier.loadBalancer)
		atomic.AddInt32(&proxier.numProxyLoops, -1)
	}(service, proxier)
	return si, nil
}
func (proxier *Proxier) mergeService(service *v1.Service) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if service == nil {
		return nil
	}
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if !helper.IsServiceIPSet(service) {
		klog.V(3).Infof("Skipping service %s due to clusterIP = %q", svcName, service.Spec.ClusterIP)
		return nil
	}
	existingPorts := sets.NewString()
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		serviceName := proxy.ServicePortName{NamespacedName: svcName, Port: servicePort.Name}
		existingPorts.Insert(servicePort.Name)
		info, exists := proxier.getServiceInfo(serviceName)
		if exists && sameConfig(info, service, servicePort) {
			continue
		}
		if exists {
			klog.V(4).Infof("Something changed for service %q: stopping it", serviceName)
			if err := proxier.closePortal(serviceName, info); err != nil {
				klog.Errorf("Failed to close portal for %q: %v", serviceName, err)
			}
			if err := proxier.stopProxy(serviceName, info); err != nil {
				klog.Errorf("Failed to stop service %q: %v", serviceName, err)
			}
		}
		proxyPort, err := proxier.proxyPorts.AllocateNext()
		if err != nil {
			klog.Errorf("failed to allocate proxy port for service %q: %v", serviceName, err)
			continue
		}
		serviceIP := net.ParseIP(service.Spec.ClusterIP)
		klog.V(1).Infof("Adding new service %q at %s/%s", serviceName, net.JoinHostPort(serviceIP.String(), strconv.Itoa(int(servicePort.Port))), servicePort.Protocol)
		info, err = proxier.addServiceOnPort(serviceName, servicePort.Protocol, proxyPort, proxier.udpIdleTimeout)
		if err != nil {
			klog.Errorf("Failed to start proxy for %q: %v", serviceName, err)
			continue
		}
		info.portal.ip = serviceIP
		info.portal.port = int(servicePort.Port)
		info.externalIPs = service.Spec.ExternalIPs
		info.loadBalancerStatus = *service.Status.LoadBalancer.DeepCopy()
		info.nodePort = int(servicePort.NodePort)
		info.sessionAffinityType = service.Spec.SessionAffinity
		if service.Spec.SessionAffinity == v1.ServiceAffinityClientIP {
			info.stickyMaxAgeSeconds = int(*service.Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds)
		}
		klog.V(4).Infof("info: %#v", info)
		if err := proxier.openPortal(serviceName, info); err != nil {
			klog.Errorf("Failed to open portal for %q: %v", serviceName, err)
		}
		proxier.loadBalancer.NewService(serviceName, info.sessionAffinityType, info.stickyMaxAgeSeconds)
	}
	return existingPorts
}
func (proxier *Proxier) unmergeService(service *v1.Service, existingPorts sets.String) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if service == nil {
		return
	}
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if !helper.IsServiceIPSet(service) {
		klog.V(3).Infof("Skipping service %s due to clusterIP = %q", svcName, service.Spec.ClusterIP)
		return
	}
	staleUDPServices := sets.NewString()
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		if existingPorts.Has(servicePort.Name) {
			continue
		}
		serviceName := proxy.ServicePortName{NamespacedName: svcName, Port: servicePort.Name}
		klog.V(1).Infof("Stopping service %q", serviceName)
		info, exists := proxier.serviceMap[serviceName]
		if !exists {
			klog.Errorf("Service %q is being removed but doesn't exist", serviceName)
			continue
		}
		if proxier.serviceMap[serviceName].protocol == v1.ProtocolUDP {
			staleUDPServices.Insert(proxier.serviceMap[serviceName].portal.ip.String())
		}
		if err := proxier.closePortal(serviceName, info); err != nil {
			klog.Errorf("Failed to close portal for %q: %v", serviceName, err)
		}
		if err := proxier.stopProxyInternal(serviceName, info); err != nil {
			klog.Errorf("Failed to stop service %q: %v", serviceName, err)
		}
		proxier.loadBalancer.DeleteService(serviceName)
	}
	for _, svcIP := range staleUDPServices.UnsortedList() {
		if err := conntrack.ClearEntriesForIP(proxier.exec, svcIP, v1.ProtocolUDP); err != nil {
			klog.Errorf("Failed to delete stale service IP %s connections, error: %v", svcIP, err)
		}
	}
}
func (proxier *Proxier) OnServiceAdd(service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_ = proxier.mergeService(service)
}
func (proxier *Proxier) OnServiceUpdate(oldService, service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	existingPorts := proxier.mergeService(service)
	proxier.unmergeService(oldService, existingPorts)
}
func (proxier *Proxier) OnServiceDelete(service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.unmergeService(service, sets.NewString())
}
func (proxier *Proxier) OnServiceSynced() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func sameConfig(info *ServiceInfo, service *v1.Service, port *v1.ServicePort) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if info.protocol != port.Protocol || info.portal.port != int(port.Port) || info.nodePort != int(port.NodePort) {
		return false
	}
	if !info.portal.ip.Equal(net.ParseIP(service.Spec.ClusterIP)) {
		return false
	}
	if !ipsEqual(info.externalIPs, service.Spec.ExternalIPs) {
		return false
	}
	if !helper.LoadBalancerStatusEqual(&info.loadBalancerStatus, &service.Status.LoadBalancer) {
		return false
	}
	if info.sessionAffinityType != service.Spec.SessionAffinity {
		return false
	}
	return true
}
func ipsEqual(lhs, rhs []string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}
func (proxier *Proxier) openPortal(service proxy.ServicePortName, info *ServiceInfo) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := proxier.openOnePortal(info.portal, info.protocol, proxier.listenIP, info.proxyPort, service)
	if err != nil {
		return err
	}
	for _, publicIP := range info.externalIPs {
		err = proxier.openOnePortal(portal{net.ParseIP(publicIP), info.portal.port, true}, info.protocol, proxier.listenIP, info.proxyPort, service)
		if err != nil {
			return err
		}
	}
	for _, ingress := range info.loadBalancerStatus.Ingress {
		if ingress.IP != "" {
			err = proxier.openOnePortal(portal{net.ParseIP(ingress.IP), info.portal.port, false}, info.protocol, proxier.listenIP, info.proxyPort, service)
			if err != nil {
				return err
			}
		}
	}
	if info.nodePort != 0 {
		err = proxier.openNodePort(info.nodePort, info.protocol, proxier.listenIP, info.proxyPort, service)
		if err != nil {
			return err
		}
	}
	return nil
}
func (proxier *Proxier) openOnePortal(portal portal, protocol v1.Protocol, proxyIP net.IP, proxyPort int, name proxy.ServicePortName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if local, err := utilproxy.IsLocalIP(portal.ip.String()); err != nil {
		return fmt.Errorf("can't determine if IP %s is local, assuming not: %v", portal.ip, err)
	} else if local {
		err := proxier.claimNodePort(portal.ip, portal.port, protocol, name)
		if err != nil {
			return err
		}
	}
	args := proxier.iptablesContainerPortalArgs(portal.ip, portal.isExternal, false, portal.port, protocol, proxyIP, proxyPort, name)
	portalAddress := net.JoinHostPort(portal.ip.String(), strconv.Itoa(portal.port))
	existed, err := proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesContainerPortalChain, args...)
	if err != nil {
		klog.Errorf("Failed to install iptables %s rule for service %q, args:%v", iptablesContainerPortalChain, name, args)
		return err
	}
	if !existed {
		klog.V(3).Infof("Opened iptables from-containers portal for service %q on %s %s", name, protocol, portalAddress)
	}
	if portal.isExternal {
		args := proxier.iptablesContainerPortalArgs(portal.ip, false, true, portal.port, protocol, proxyIP, proxyPort, name)
		existed, err := proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesContainerPortalChain, args...)
		if err != nil {
			klog.Errorf("Failed to install iptables %s rule that opens service %q for local traffic, args:%v", iptablesContainerPortalChain, name, args)
			return err
		}
		if !existed {
			klog.V(3).Infof("Opened iptables from-containers portal for service %q on %s %s for local traffic", name, protocol, portalAddress)
		}
		args = proxier.iptablesHostPortalArgs(portal.ip, true, portal.port, protocol, proxyIP, proxyPort, name)
		existed, err = proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesHostPortalChain, args...)
		if err != nil {
			klog.Errorf("Failed to install iptables %s rule for service %q for dst-local traffic", iptablesHostPortalChain, name)
			return err
		}
		if !existed {
			klog.V(3).Infof("Opened iptables from-host portal for service %q on %s %s for dst-local traffic", name, protocol, portalAddress)
		}
		return nil
	}
	args = proxier.iptablesHostPortalArgs(portal.ip, false, portal.port, protocol, proxyIP, proxyPort, name)
	existed, err = proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesHostPortalChain, args...)
	if err != nil {
		klog.Errorf("Failed to install iptables %s rule for service %q", iptablesHostPortalChain, name)
		return err
	}
	if !existed {
		klog.V(3).Infof("Opened iptables from-host portal for service %q on %s %s", name, protocol, portalAddress)
	}
	return nil
}
func (proxier *Proxier) claimNodePort(ip net.IP, port int, protocol v1.Protocol, owner proxy.ServicePortName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.portMapMutex.Lock()
	defer proxier.portMapMutex.Unlock()
	key := portMapKey{ip: ip.String(), port: port, protocol: protocol}
	existing, found := proxier.portMap[key]
	if !found {
		socket, err := proxier.makeProxySocket(protocol, ip, port)
		if err != nil {
			return fmt.Errorf("can't open node port for %s: %v", key.String(), err)
		}
		proxier.portMap[key] = &portMapValue{owner: owner, socket: socket}
		klog.V(2).Infof("Claimed local port %s", key.String())
		return nil
	}
	if existing.owner == owner {
		return nil
	}
	return fmt.Errorf("Port conflict detected on port %s.  %v vs %v", key.String(), owner, existing)
}
func (proxier *Proxier) releaseNodePort(ip net.IP, port int, protocol v1.Protocol, owner proxy.ServicePortName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.portMapMutex.Lock()
	defer proxier.portMapMutex.Unlock()
	key := portMapKey{ip: ip.String(), port: port, protocol: protocol}
	existing, found := proxier.portMap[key]
	if !found {
		klog.Infof("Ignoring release on unowned port: %v", key)
		return nil
	}
	if existing.owner != owner {
		return fmt.Errorf("Port conflict detected on port %v (unowned unlock).  %v vs %v", key, owner, existing)
	}
	delete(proxier.portMap, key)
	existing.socket.Close()
	return nil
}
func (proxier *Proxier) openNodePort(nodePort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, name proxy.ServicePortName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := proxier.claimNodePort(nil, nodePort, protocol, name)
	if err != nil {
		return err
	}
	args := proxier.iptablesContainerPortalArgs(nil, false, false, nodePort, protocol, proxyIP, proxyPort, name)
	existed, err := proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesContainerNodePortChain, args...)
	if err != nil {
		klog.Errorf("Failed to install iptables %s rule for service %q", iptablesContainerNodePortChain, name)
		return err
	}
	if !existed {
		klog.Infof("Opened iptables from-containers public port for service %q on %s port %d", name, protocol, nodePort)
	}
	args = proxier.iptablesHostNodePortArgs(nodePort, protocol, proxyIP, proxyPort, name)
	existed, err = proxier.iptables.EnsureRule(iptables.Append, iptables.TableNAT, iptablesHostNodePortChain, args...)
	if err != nil {
		klog.Errorf("Failed to install iptables %s rule for service %q", iptablesHostNodePortChain, name)
		return err
	}
	if !existed {
		klog.Infof("Opened iptables from-host public port for service %q on %s port %d", name, protocol, nodePort)
	}
	args = proxier.iptablesNonLocalNodePortArgs(nodePort, protocol, proxyIP, proxyPort, name)
	existed, err = proxier.iptables.EnsureRule(iptables.Append, iptables.TableFilter, iptablesNonLocalNodePortChain, args...)
	if err != nil {
		klog.Errorf("Failed to install iptables %s rule for service %q", iptablesNonLocalNodePortChain, name)
		return err
	}
	if !existed {
		klog.Infof("Opened iptables from-non-local public port for service %q on %s port %d", name, protocol, nodePort)
	}
	return nil
}
func (proxier *Proxier) closePortal(service proxy.ServicePortName, info *ServiceInfo) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	el := proxier.closeOnePortal(info.portal, info.protocol, proxier.listenIP, info.proxyPort, service)
	for _, publicIP := range info.externalIPs {
		el = append(el, proxier.closeOnePortal(portal{net.ParseIP(publicIP), info.portal.port, true}, info.protocol, proxier.listenIP, info.proxyPort, service)...)
	}
	for _, ingress := range info.loadBalancerStatus.Ingress {
		if ingress.IP != "" {
			el = append(el, proxier.closeOnePortal(portal{net.ParseIP(ingress.IP), info.portal.port, false}, info.protocol, proxier.listenIP, info.proxyPort, service)...)
		}
	}
	if info.nodePort != 0 {
		el = append(el, proxier.closeNodePort(info.nodePort, info.protocol, proxier.listenIP, info.proxyPort, service)...)
	}
	if len(el) == 0 {
		klog.V(3).Infof("Closed iptables portals for service %q", service)
	} else {
		klog.Errorf("Some errors closing iptables portals for service %q", service)
	}
	return utilerrors.NewAggregate(el)
}
func (proxier *Proxier) closeOnePortal(portal portal, protocol v1.Protocol, proxyIP net.IP, proxyPort int, name proxy.ServicePortName) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	el := []error{}
	if local, err := utilproxy.IsLocalIP(portal.ip.String()); err != nil {
		el = append(el, fmt.Errorf("can't determine if IP %s is local, assuming not: %v", portal.ip, err))
	} else if local {
		if err := proxier.releaseNodePort(portal.ip, portal.port, protocol, name); err != nil {
			el = append(el, err)
		}
	}
	args := proxier.iptablesContainerPortalArgs(portal.ip, portal.isExternal, false, portal.port, protocol, proxyIP, proxyPort, name)
	if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesContainerPortalChain, args...); err != nil {
		klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesContainerPortalChain, name)
		el = append(el, err)
	}
	if portal.isExternal {
		args := proxier.iptablesContainerPortalArgs(portal.ip, false, true, portal.port, protocol, proxyIP, proxyPort, name)
		if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesContainerPortalChain, args...); err != nil {
			klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesContainerPortalChain, name)
			el = append(el, err)
		}
		args = proxier.iptablesHostPortalArgs(portal.ip, true, portal.port, protocol, proxyIP, proxyPort, name)
		if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesHostPortalChain, args...); err != nil {
			klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesHostPortalChain, name)
			el = append(el, err)
		}
		return el
	}
	args = proxier.iptablesHostPortalArgs(portal.ip, false, portal.port, protocol, proxyIP, proxyPort, name)
	if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesHostPortalChain, args...); err != nil {
		klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesHostPortalChain, name)
		el = append(el, err)
	}
	return el
}
func (proxier *Proxier) closeNodePort(nodePort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, name proxy.ServicePortName) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	el := []error{}
	args := proxier.iptablesContainerPortalArgs(nil, false, false, nodePort, protocol, proxyIP, proxyPort, name)
	if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesContainerNodePortChain, args...); err != nil {
		klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesContainerNodePortChain, name)
		el = append(el, err)
	}
	args = proxier.iptablesHostNodePortArgs(nodePort, protocol, proxyIP, proxyPort, name)
	if err := proxier.iptables.DeleteRule(iptables.TableNAT, iptablesHostNodePortChain, args...); err != nil {
		klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesHostNodePortChain, name)
		el = append(el, err)
	}
	args = proxier.iptablesNonLocalNodePortArgs(nodePort, protocol, proxyIP, proxyPort, name)
	if err := proxier.iptables.DeleteRule(iptables.TableFilter, iptablesNonLocalNodePortChain, args...); err != nil {
		klog.Errorf("Failed to delete iptables %s rule for service %q", iptablesNonLocalNodePortChain, name)
		el = append(el, err)
	}
	if err := proxier.releaseNodePort(nil, nodePort, protocol, name); err != nil {
		el = append(el, err)
	}
	return el
}

var iptablesContainerPortalChain iptables.Chain = "KUBE-PORTALS-CONTAINER"
var iptablesHostPortalChain iptables.Chain = "KUBE-PORTALS-HOST"
var iptablesContainerNodePortChain iptables.Chain = "KUBE-NODEPORT-CONTAINER"
var iptablesHostNodePortChain iptables.Chain = "KUBE-NODEPORT-HOST"
var iptablesNonLocalNodePortChain iptables.Chain = "KUBE-NODEPORT-NON-LOCAL"

func iptablesInit(ipt iptables.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []string{}
	args = append(args, "-m", "comment", "--comment", "handle ClusterIPs; NOTE: this must be before the NodePort rules")
	if _, err := ipt.EnsureChain(iptables.TableNAT, iptablesContainerPortalChain); err != nil {
		return err
	}
	if _, err := ipt.EnsureRule(iptables.Prepend, iptables.TableNAT, iptables.ChainPrerouting, append(args, "-j", string(iptablesContainerPortalChain))...); err != nil {
		return err
	}
	if _, err := ipt.EnsureChain(iptables.TableNAT, iptablesHostPortalChain); err != nil {
		return err
	}
	if _, err := ipt.EnsureRule(iptables.Prepend, iptables.TableNAT, iptables.ChainOutput, append(args, "-j", string(iptablesHostPortalChain))...); err != nil {
		return err
	}
	args = []string{"-m", "addrtype", "--dst-type", "LOCAL"}
	args = append(args, "-m", "comment", "--comment", "handle service NodePorts; NOTE: this must be the last rule in the chain")
	if _, err := ipt.EnsureChain(iptables.TableNAT, iptablesContainerNodePortChain); err != nil {
		return err
	}
	if _, err := ipt.EnsureRule(iptables.Append, iptables.TableNAT, iptables.ChainPrerouting, append(args, "-j", string(iptablesContainerNodePortChain))...); err != nil {
		return err
	}
	if _, err := ipt.EnsureChain(iptables.TableNAT, iptablesHostNodePortChain); err != nil {
		return err
	}
	if _, err := ipt.EnsureRule(iptables.Append, iptables.TableNAT, iptables.ChainOutput, append(args, "-j", string(iptablesHostNodePortChain))...); err != nil {
		return err
	}
	args = []string{"-m", "comment", "--comment", "Ensure that non-local NodePort traffic can flow"}
	if _, err := ipt.EnsureChain(iptables.TableFilter, iptablesNonLocalNodePortChain); err != nil {
		return err
	}
	if _, err := ipt.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainInput, append(args, "-j", string(iptablesNonLocalNodePortChain))...); err != nil {
		return err
	}
	return nil
}
func iptablesFlush(ipt iptables.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	el := []error{}
	if err := ipt.FlushChain(iptables.TableNAT, iptablesContainerPortalChain); err != nil {
		el = append(el, err)
	}
	if err := ipt.FlushChain(iptables.TableNAT, iptablesHostPortalChain); err != nil {
		el = append(el, err)
	}
	if err := ipt.FlushChain(iptables.TableNAT, iptablesContainerNodePortChain); err != nil {
		el = append(el, err)
	}
	if err := ipt.FlushChain(iptables.TableNAT, iptablesHostNodePortChain); err != nil {
		el = append(el, err)
	}
	if err := ipt.FlushChain(iptables.TableFilter, iptablesNonLocalNodePortChain); err != nil {
		el = append(el, err)
	}
	if len(el) != 0 {
		klog.Errorf("Some errors flushing old iptables portals: %v", el)
	}
	return utilerrors.NewAggregate(el)
}

var zeroIPv4 = net.ParseIP("0.0.0.0")
var localhostIPv4 = net.ParseIP("127.0.0.1")
var zeroIPv6 = net.ParseIP("::")
var localhostIPv6 = net.ParseIP("::1")

func iptablesCommonPortalArgs(destIP net.IP, addPhysicalInterfaceMatch bool, addDstLocalMatch bool, destPort int, protocol v1.Protocol, service proxy.ServicePortName) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []string{"-m", "comment", "--comment", service.String(), "-p", strings.ToLower(string(protocol)), "-m", strings.ToLower(string(protocol)), "--dport", fmt.Sprintf("%d", destPort)}
	if destIP != nil {
		args = append(args, "-d", utilproxy.ToCIDR(destIP))
	}
	if addPhysicalInterfaceMatch {
		args = append(args, "-m", "physdev", "!", "--physdev-is-in")
	}
	if addDstLocalMatch {
		args = append(args, "-m", "addrtype", "--dst-type", "LOCAL")
	}
	return args
}
func (proxier *Proxier) iptablesContainerPortalArgs(destIP net.IP, addPhysicalInterfaceMatch bool, addDstLocalMatch bool, destPort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, service proxy.ServicePortName) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := iptablesCommonPortalArgs(destIP, addPhysicalInterfaceMatch, addDstLocalMatch, destPort, protocol, service)
	if proxyIP.Equal(zeroIPv4) || proxyIP.Equal(zeroIPv6) {
		args = append(args, "-j", "REDIRECT", "--to-ports", fmt.Sprintf("%d", proxyPort))
	} else {
		args = append(args, "-j", "DNAT", "--to-destination", net.JoinHostPort(proxyIP.String(), strconv.Itoa(proxyPort)))
	}
	return args
}
func (proxier *Proxier) iptablesHostPortalArgs(destIP net.IP, addDstLocalMatch bool, destPort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, service proxy.ServicePortName) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := iptablesCommonPortalArgs(destIP, false, addDstLocalMatch, destPort, protocol, service)
	if proxyIP.Equal(zeroIPv4) || proxyIP.Equal(zeroIPv6) {
		proxyIP = proxier.hostIP
	}
	args = append(args, "-j", "DNAT", "--to-destination", net.JoinHostPort(proxyIP.String(), strconv.Itoa(proxyPort)))
	return args
}
func (proxier *Proxier) iptablesHostNodePortArgs(nodePort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, service proxy.ServicePortName) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := iptablesCommonPortalArgs(nil, false, false, nodePort, protocol, service)
	if proxyIP.Equal(zeroIPv4) || proxyIP.Equal(zeroIPv6) {
		proxyIP = proxier.hostIP
	}
	args = append(args, "-j", "DNAT", "--to-destination", net.JoinHostPort(proxyIP.String(), strconv.Itoa(proxyPort)))
	return args
}
func (proxier *Proxier) iptablesNonLocalNodePortArgs(nodePort int, protocol v1.Protocol, proxyIP net.IP, proxyPort int, service proxy.ServicePortName) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := iptablesCommonPortalArgs(nil, false, false, proxyPort, protocol, service)
	args = append(args, "-m", "state", "--state", "NEW", "-j", "ACCEPT")
	return args
}
func isTooManyFDsError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.Contains(err.Error(), "too many open files")
}
func isClosedError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasSuffix(err.Error(), "use of closed network connection")
}
