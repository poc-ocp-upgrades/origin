package winkernel

import (
	"encoding/json"
	"fmt"
	"github.com/Microsoft/hcsshim"
	"github.com/davecgh/go-spew/spew"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	apiservice "k8s.io/kubernetes/pkg/api/v1/service"
	"k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/kubernetes/pkg/proxy"
	"k8s.io/kubernetes/pkg/proxy/healthcheck"
	"k8s.io/kubernetes/pkg/util/async"
	"net"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type KernelCompatTester interface{ IsCompatible() error }

func CanUseWinKernelProxier(kcompat KernelCompatTester) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := kcompat.IsCompatible(); err != nil {
		return false, err
	}
	return true, nil
}

type WindowsKernelCompatTester struct{}

func (lkct WindowsKernelCompatTester) IsCompatible() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := hcsshim.HNSListPolicyListRequest()
	if err != nil {
		return fmt.Errorf("Windows kernel is not compatible for Kernel mode")
	}
	return nil
}

type externalIPInfo struct {
	ip    string
	hnsID string
}
type loadBalancerIngressInfo struct {
	ip    string
	hnsID string
}
type serviceInfo struct {
	clusterIP                net.IP
	port                     int
	protocol                 v1.Protocol
	nodePort                 int
	targetPort               int
	loadBalancerStatus       v1.LoadBalancerStatus
	sessionAffinityType      v1.ServiceAffinity
	stickyMaxAgeSeconds      int
	externalIPs              []*externalIPInfo
	loadBalancerIngressIPs   []*loadBalancerIngressInfo
	loadBalancerSourceRanges []string
	onlyNodeLocalEndpoints   bool
	healthCheckNodePort      int
	hnsID                    string
	nodePorthnsID            string
	policyApplied            bool
}
type hnsNetworkInfo struct {
	name string
	id   string
}

func Log(v interface{}, message string, level klog.Level) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(level).Infof("%s, %s", message, spew.Sdump(v))
}
func LogJson(v interface{}, message string, level klog.Level) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jsonString, err := json.Marshal(v)
	if err == nil {
		klog.V(level).Infof("%s, %s", message, string(jsonString))
	}
}

type endpointsInfo struct {
	ip         string
	port       uint16
	isLocal    bool
	macAddress string
	hnsID      string
	refCount   uint16
}

func conjureMac(macPrefix string, ip net.IP) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ip4 := ip.To4(); ip4 != nil {
		a, b, c, d := ip4[0], ip4[1], ip4[2], ip4[3]
		return fmt.Sprintf("%v-%02x-%02x-%02x-%02x", macPrefix, a, b, c, d)
	}
	return "02-11-22-33-44-55"
}
func newEndpointInfo(ip string, port uint16, isLocal bool) *endpointsInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info := &endpointsInfo{ip: ip, port: port, isLocal: isLocal, macAddress: conjureMac("02-11", net.ParseIP(ip)), refCount: 0, hnsID: ""}
	return info
}
func (ep *endpointsInfo) Cleanup() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	Log(ep, "Endpoint Cleanup", 3)
	ep.refCount--
	if ep.refCount <= 0 && !ep.isLocal {
		klog.V(4).Infof("Removing endpoints for %v, since no one is referencing it", ep)
		deleteHnsEndpoint(ep.hnsID)
		ep.hnsID = ""
	}
}
func newServiceInfo(svcPortName proxy.ServicePortName, port *v1.ServicePort, service *v1.Service) *serviceInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	onlyNodeLocalEndpoints := false
	if apiservice.RequestsOnlyLocalTraffic(service) {
		onlyNodeLocalEndpoints = true
	}
	stickyMaxAgeSeconds := 10800
	if service.Spec.SessionAffinity == v1.ServiceAffinityClientIP {
		stickyMaxAgeSeconds = int(*service.Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds)
	}
	info := &serviceInfo{clusterIP: net.ParseIP(service.Spec.ClusterIP), port: int(port.Port), protocol: port.Protocol, nodePort: int(port.NodePort), targetPort: port.TargetPort.IntValue(), loadBalancerStatus: *service.Status.LoadBalancer.DeepCopy(), sessionAffinityType: service.Spec.SessionAffinity, stickyMaxAgeSeconds: stickyMaxAgeSeconds, loadBalancerSourceRanges: make([]string, len(service.Spec.LoadBalancerSourceRanges)), onlyNodeLocalEndpoints: onlyNodeLocalEndpoints}
	copy(info.loadBalancerSourceRanges, service.Spec.LoadBalancerSourceRanges)
	for _, eip := range service.Spec.ExternalIPs {
		info.externalIPs = append(info.externalIPs, &externalIPInfo{ip: eip})
	}
	for _, ingress := range service.Status.LoadBalancer.Ingress {
		info.loadBalancerIngressIPs = append(info.loadBalancerIngressIPs, &loadBalancerIngressInfo{ip: ingress.IP})
	}
	if apiservice.NeedsHealthCheck(service) {
		p := service.Spec.HealthCheckNodePort
		if p == 0 {
			klog.Errorf("Service %q has no healthcheck nodeport", svcPortName.NamespacedName.String())
		} else {
			info.healthCheckNodePort = int(p)
		}
	}
	return info
}

type endpointsChange struct {
	previous proxyEndpointsMap
	current  proxyEndpointsMap
}
type endpointsChangeMap struct {
	lock     sync.Mutex
	hostname string
	items    map[types.NamespacedName]*endpointsChange
}
type serviceChange struct {
	previous proxyServiceMap
	current  proxyServiceMap
}
type serviceChangeMap struct {
	lock  sync.Mutex
	items map[types.NamespacedName]*serviceChange
}
type updateEndpointMapResult struct {
	hcEndpoints       map[types.NamespacedName]int
	staleEndpoints    map[endpointServicePair]bool
	staleServiceNames map[proxy.ServicePortName]bool
}
type updateServiceMapResult struct {
	hcServices    map[types.NamespacedName]uint16
	staleServices sets.String
}
type proxyServiceMap map[proxy.ServicePortName]*serviceInfo
type proxyEndpointsMap map[proxy.ServicePortName][]*endpointsInfo

func newEndpointsChangeMap(hostname string) endpointsChangeMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return endpointsChangeMap{hostname: hostname, items: make(map[types.NamespacedName]*endpointsChange)}
}
func (ecm *endpointsChangeMap) update(namespacedName *types.NamespacedName, previous, current *v1.Endpoints) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ecm.lock.Lock()
	defer ecm.lock.Unlock()
	change, exists := ecm.items[*namespacedName]
	if !exists {
		change = &endpointsChange{}
		change.previous = endpointsToEndpointsMap(previous, ecm.hostname)
		ecm.items[*namespacedName] = change
	}
	change.current = endpointsToEndpointsMap(current, ecm.hostname)
	if reflect.DeepEqual(change.previous, change.current) {
		delete(ecm.items, *namespacedName)
	}
	return len(ecm.items) > 0
}
func newServiceChangeMap() serviceChangeMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return serviceChangeMap{items: make(map[types.NamespacedName]*serviceChange)}
}
func (scm *serviceChangeMap) update(namespacedName *types.NamespacedName, previous, current *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scm.lock.Lock()
	defer scm.lock.Unlock()
	change, exists := scm.items[*namespacedName]
	if !exists {
		change = &serviceChange{}
		change.previous = serviceToServiceMap(previous)
		scm.items[*namespacedName] = change
	}
	change.current = serviceToServiceMap(current)
	if reflect.DeepEqual(change.previous, change.current) {
		delete(scm.items, *namespacedName)
	}
	return len(scm.items) > 0
}
func (sm *proxyServiceMap) merge(other proxyServiceMap, curEndpoints proxyEndpointsMap) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	existingPorts := sets.NewString()
	for svcPortName, info := range other {
		existingPorts.Insert(svcPortName.Port)
		svcInfo, exists := (*sm)[svcPortName]
		if !exists {
			klog.V(1).Infof("Adding new service port %q at %s:%d/%s", svcPortName, info.clusterIP, info.port, info.protocol)
		} else {
			klog.V(1).Infof("Updating existing service port %q at %s:%d/%s", svcPortName, info.clusterIP, info.port, info.protocol)
			svcInfo.cleanupAllPolicies(curEndpoints[svcPortName])
			delete(*sm, svcPortName)
		}
		(*sm)[svcPortName] = info
	}
	return existingPorts
}
func (sm *proxyServiceMap) unmerge(other proxyServiceMap, existingPorts, staleServices sets.String, curEndpoints proxyEndpointsMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		if existingPorts.Has(svcPortName.Port) {
			continue
		}
		info, exists := (*sm)[svcPortName]
		if exists {
			klog.V(1).Infof("Removing service port %q", svcPortName)
			if info.protocol == v1.ProtocolUDP {
				staleServices.Insert(info.clusterIP.String())
			}
			info.cleanupAllPolicies(curEndpoints[svcPortName])
			delete(*sm, svcPortName)
		} else {
			klog.Errorf("Service port %q removed, but doesn't exists", svcPortName)
		}
	}
}
func (em proxyEndpointsMap) merge(other proxyEndpointsMap, curServices proxyServiceMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		epInfos, exists := em[svcPortName]
		if exists {
			info, exists := curServices[svcPortName]
			klog.V(1).Infof("Updating existing service port %q at %s:%d/%s", svcPortName, info.clusterIP, info.port, info.protocol)
			if exists {
				klog.V(2).Infof("Endpoints are modified. Service [%v] is stale", svcPortName)
				info.cleanupAllPolicies(epInfos)
			} else {
				klog.V(2).Infof("Endpoints are orphaned. Cleaning up")
				for _, ep := range epInfos {
					ep.Cleanup()
				}
			}
			delete(em, svcPortName)
		}
		em[svcPortName] = other[svcPortName]
	}
}
func (em proxyEndpointsMap) unmerge(other proxyEndpointsMap, curServices proxyServiceMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		info, exists := curServices[svcPortName]
		if exists {
			klog.V(2).Infof("Service [%v] is stale", info)
			info.cleanupAllPolicies(em[svcPortName])
		} else {
			klog.V(2).Infof("Endpoints are orphaned. Cleaning up")
			epInfos, exists := em[svcPortName]
			if exists {
				for _, ep := range epInfos {
					ep.Cleanup()
				}
			}
		}
		delete(em, svcPortName)
	}
}

type Proxier struct {
	endpointsChanges         endpointsChangeMap
	serviceChanges           serviceChangeMap
	mu                       sync.Mutex
	serviceMap               proxyServiceMap
	endpointsMap             proxyEndpointsMap
	portsMap                 map[localPort]closeable
	endpointsSynced          bool
	servicesSynced           bool
	initialized              int32
	syncRunner               *async.BoundedFrequencyRunner
	masqueradeAll            bool
	masqueradeMark           string
	clusterCIDR              string
	hostname                 string
	nodeIP                   net.IP
	recorder                 record.EventRecorder
	healthChecker            healthcheck.Server
	healthzServer            healthcheck.HealthzUpdater
	precomputedProbabilities []string
	network                  hnsNetworkInfo
}
type localPort struct {
	desc     string
	ip       string
	port     int
	protocol string
}

func (lp *localPort) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%q (%s:%d/%s)", lp.desc, lp.ip, lp.port, lp.protocol)
}
func Enum(p v1.Protocol) uint16 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p == v1.ProtocolTCP {
		return 6
	}
	if p == v1.ProtocolUDP {
		return 17
	}
	if p == v1.ProtocolSCTP {
		return 132
	}
	return 0
}

type closeable interface{ Close() error }

var _ proxy.ProxyProvider = &Proxier{}

func NewProxier(syncPeriod time.Duration, minSyncPeriod time.Duration, masqueradeAll bool, masqueradeBit int, clusterCIDR string, hostname string, nodeIP net.IP, recorder record.EventRecorder, healthzServer healthcheck.HealthzUpdater) (*Proxier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	masqueradeValue := 1 << uint(masqueradeBit)
	masqueradeMark := fmt.Sprintf("%#08x/%#08x", masqueradeValue, masqueradeValue)
	if nodeIP == nil {
		klog.Warningf("invalid nodeIP, initializing kube-proxy with 127.0.0.1 as nodeIP")
		nodeIP = net.ParseIP("127.0.0.1")
	}
	if len(clusterCIDR) == 0 {
		klog.Warningf("clusterCIDR not specified, unable to distinguish between internal and external traffic")
	}
	healthChecker := healthcheck.NewServer(hostname, recorder, nil, nil)
	hnsNetworkName := os.Getenv("KUBE_NETWORK")
	if len(hnsNetworkName) == 0 {
		return nil, fmt.Errorf("Environment variable KUBE_NETWORK not initialized")
	}
	hnsNetwork, err := getHnsNetworkInfo(hnsNetworkName)
	if err != nil {
		klog.Fatalf("Unable to find Hns Network specified by %s. Please check environment variable KUBE_NETWORK", hnsNetworkName)
		return nil, err
	}
	klog.V(1).Infof("Hns Network loaded with info = %v", hnsNetwork)
	proxier := &Proxier{portsMap: make(map[localPort]closeable), serviceMap: make(proxyServiceMap), serviceChanges: newServiceChangeMap(), endpointsMap: make(proxyEndpointsMap), endpointsChanges: newEndpointsChangeMap(hostname), masqueradeAll: masqueradeAll, masqueradeMark: masqueradeMark, clusterCIDR: clusterCIDR, hostname: hostname, nodeIP: nodeIP, recorder: recorder, healthChecker: healthChecker, healthzServer: healthzServer, network: *hnsNetwork}
	burstSyncs := 2
	klog.V(3).Infof("minSyncPeriod: %v, syncPeriod: %v, burstSyncs: %d", minSyncPeriod, syncPeriod, burstSyncs)
	proxier.syncRunner = async.NewBoundedFrequencyRunner("sync-runner", proxier.syncProxyRules, minSyncPeriod, syncPeriod, burstSyncs)
	return proxier, nil
}
func CleanupLeftovers() (encounteredError bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	deleteAllHnsLoadBalancerPolicy()
	return encounteredError
}
func (svcInfo *serviceInfo) cleanupAllPolicies(endpoints []*endpointsInfo) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	Log(svcInfo, "Service Cleanup", 3)
	svcInfo.deleteAllHnsLoadBalancerPolicy()
	for _, ep := range endpoints {
		ep.Cleanup()
	}
	svcInfo.policyApplied = false
}
func (svcInfo *serviceInfo) deleteAllHnsLoadBalancerPolicy() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	deleteHnsLoadBalancerPolicy(svcInfo.hnsID)
	svcInfo.hnsID = ""
	deleteHnsLoadBalancerPolicy(svcInfo.nodePorthnsID)
	svcInfo.nodePorthnsID = ""
	for _, externalIp := range svcInfo.externalIPs {
		deleteHnsLoadBalancerPolicy(externalIp.hnsID)
		externalIp.hnsID = ""
	}
	for _, lbIngressIp := range svcInfo.loadBalancerIngressIPs {
		deleteHnsLoadBalancerPolicy(lbIngressIp.hnsID)
		lbIngressIp.hnsID = ""
	}
}
func deleteAllHnsLoadBalancerPolicy() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plists, err := hcsshim.HNSListPolicyListRequest()
	if err != nil {
		return
	}
	for _, plist := range plists {
		LogJson(plist, "Remove Policy", 3)
		_, err = plist.Delete()
		if err != nil {
			klog.Errorf("%v", err)
		}
	}
}
func getHnsLoadBalancer(endpoints []hcsshim.HNSEndpoint, isILB bool, vip string, protocol uint16, internalPort uint16, externalPort uint16) (*hcsshim.PolicyList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plists, err := hcsshim.HNSListPolicyListRequest()
	if err != nil {
		return nil, err
	}
	for _, plist := range plists {
		if len(plist.EndpointReferences) != len(endpoints) {
			continue
		}
		elbPolicy := hcsshim.ELBPolicy{}
		if err = json.Unmarshal(plist.Policies[0], &elbPolicy); err != nil {
			continue
		}
		if elbPolicy.Protocol == protocol && elbPolicy.InternalPort == internalPort && elbPolicy.ExternalPort == externalPort && elbPolicy.ILB == isILB {
			if len(vip) > 0 {
				if len(elbPolicy.VIPs) == 0 || elbPolicy.VIPs[0] != vip {
					continue
				}
			}
			LogJson(plist, "Found existing Hns loadbalancer policy resource", 1)
			return &plist, nil
		}
	}
	var sourceVip string
	lb, err := hcsshim.AddLoadBalancer(endpoints, isILB, sourceVip, vip, protocol, internalPort, externalPort)
	if err == nil {
		LogJson(lb, "Hns loadbalancer policy resource", 1)
	}
	return lb, err
}
func deleteHnsLoadBalancerPolicy(hnsID string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(hnsID) == 0 {
		return
	}
	hnsloadBalancer, err := hcsshim.GetPolicyListByID(hnsID)
	if err != nil {
		klog.Errorf("%v", err)
		return
	}
	LogJson(hnsloadBalancer, "Removing Policy", 2)
	_, err = hnsloadBalancer.Delete()
	if err != nil {
		klog.Errorf("%v", err)
	}
}
func deleteHnsEndpoint(hnsID string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hnsendpoint, err := hcsshim.GetHNSEndpointByID(hnsID)
	if err != nil {
		klog.Errorf("%v", err)
		return
	}
	_, err = hnsendpoint.Delete()
	if err != nil {
		klog.Errorf("%v", err)
	}
	klog.V(3).Infof("Remote endpoint resource deleted id %s", hnsID)
}
func getHnsNetworkInfo(hnsNetworkName string) (*hnsNetworkInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hnsnetwork, err := hcsshim.GetHNSNetworkByName(hnsNetworkName)
	if err != nil {
		klog.Errorf("%v", err)
		return nil, err
	}
	return &hnsNetworkInfo{id: hnsnetwork.Id, name: hnsnetwork.Name}, nil
}
func getHnsEndpointByIpAddress(ip net.IP, networkName string) (*hcsshim.HNSEndpoint, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hnsnetwork, err := hcsshim.GetHNSNetworkByName(networkName)
	if err != nil {
		klog.Errorf("%v", err)
		return nil, err
	}
	endpoints, err := hcsshim.HNSListEndpointRequest()
	for _, endpoint := range endpoints {
		equal := reflect.DeepEqual(endpoint.IPAddress, ip)
		if equal && endpoint.VirtualNetwork == hnsnetwork.Id {
			return &endpoint, nil
		}
	}
	return nil, fmt.Errorf("Endpoint %v not found on network %s", ip, networkName)
}
func (proxier *Proxier) Sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.syncRunner.Run()
}
func (proxier *Proxier) SyncLoop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if proxier.healthzServer != nil {
		proxier.healthzServer.UpdateTimestamp()
	}
	proxier.syncRunner.Loop(wait.NeverStop)
}
func (proxier *Proxier) setInitialized(value bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var initialized int32
	if value {
		initialized = 1
	}
	atomic.StoreInt32(&proxier.initialized, initialized)
}
func (proxier *Proxier) isInitialized() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return atomic.LoadInt32(&proxier.initialized) > 0
}
func (proxier *Proxier) OnServiceAdd(service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if proxier.serviceChanges.update(&namespacedName, nil, service) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnServiceUpdate(oldService, service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if proxier.serviceChanges.update(&namespacedName, oldService, service) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnServiceDelete(service *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if proxier.serviceChanges.update(&namespacedName, service, nil) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnServiceSynced() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	proxier.servicesSynced = true
	proxier.setInitialized(proxier.servicesSynced && proxier.endpointsSynced)
	proxier.mu.Unlock()
	proxier.syncProxyRules()
}
func shouldSkipService(svcName types.NamespacedName, service *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func (proxier *Proxier) updateServiceMap() (result updateServiceMapResult) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result.staleServices = sets.NewString()
	var serviceMap proxyServiceMap = proxier.serviceMap
	var changes *serviceChangeMap = &proxier.serviceChanges
	func() {
		changes.lock.Lock()
		defer changes.lock.Unlock()
		for _, change := range changes.items {
			existingPorts := serviceMap.merge(change.current, proxier.endpointsMap)
			serviceMap.unmerge(change.previous, existingPorts, result.staleServices, proxier.endpointsMap)
		}
		changes.items = make(map[types.NamespacedName]*serviceChange)
	}()
	result.hcServices = make(map[types.NamespacedName]uint16)
	for svcPortName, info := range serviceMap {
		if info.healthCheckNodePort != 0 {
			result.hcServices[svcPortName.NamespacedName] = uint16(info.healthCheckNodePort)
		}
	}
	return result
}
func (proxier *Proxier) OnEndpointsAdd(endpoints *v1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	if proxier.endpointsChanges.update(&namespacedName, nil, endpoints) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnEndpointsUpdate(oldEndpoints, endpoints *v1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	if proxier.endpointsChanges.update(&namespacedName, oldEndpoints, endpoints) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnEndpointsDelete(endpoints *v1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespacedName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	if proxier.endpointsChanges.update(&namespacedName, endpoints, nil) && proxier.isInitialized() {
		proxier.syncRunner.Run()
	}
}
func (proxier *Proxier) OnEndpointsSynced() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	proxier.endpointsSynced = true
	proxier.setInitialized(proxier.servicesSynced && proxier.endpointsSynced)
	proxier.mu.Unlock()
	proxier.syncProxyRules()
}
func (proxier *Proxier) updateEndpointsMap() (result updateEndpointMapResult) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result.staleEndpoints = make(map[endpointServicePair]bool)
	result.staleServiceNames = make(map[proxy.ServicePortName]bool)
	var endpointsMap proxyEndpointsMap = proxier.endpointsMap
	var changes *endpointsChangeMap = &proxier.endpointsChanges
	func() {
		changes.lock.Lock()
		defer changes.lock.Unlock()
		for _, change := range changes.items {
			endpointsMap.unmerge(change.previous, proxier.serviceMap)
			endpointsMap.merge(change.current, proxier.serviceMap)
		}
		changes.items = make(map[types.NamespacedName]*endpointsChange)
	}()
	result.hcEndpoints = make(map[types.NamespacedName]int)
	localIPs := getLocalIPs(endpointsMap)
	for nsn, ips := range localIPs {
		result.hcEndpoints[nsn] = len(ips)
	}
	return result
}
func getLocalIPs(endpointsMap proxyEndpointsMap) map[types.NamespacedName]sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	localIPs := make(map[types.NamespacedName]sets.String)
	for svcPortName := range endpointsMap {
		for _, ep := range endpointsMap[svcPortName] {
			if ep.isLocal {
				nsn := svcPortName.NamespacedName
				if localIPs[nsn] == nil {
					localIPs[nsn] = sets.NewString()
				}
				localIPs[nsn].Insert(ep.ip)
			}
		}
	}
	return localIPs
}
func endpointsToEndpointsMap(endpoints *v1.Endpoints, hostname string) proxyEndpointsMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if endpoints == nil {
		return nil
	}
	endpointsMap := make(proxyEndpointsMap)
	for i := range endpoints.Subsets {
		ss := &endpoints.Subsets[i]
		for i := range ss.Ports {
			port := &ss.Ports[i]
			if port.Port == 0 {
				klog.Warningf("ignoring invalid endpoint port %s", port.Name)
				continue
			}
			svcPortName := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: port.Name}
			for i := range ss.Addresses {
				addr := &ss.Addresses[i]
				if addr.IP == "" {
					klog.Warningf("ignoring invalid endpoint port %s with empty host", port.Name)
					continue
				}
				isLocal := addr.NodeName != nil && *addr.NodeName == hostname
				epInfo := newEndpointInfo(addr.IP, uint16(port.Port), isLocal)
				endpointsMap[svcPortName] = append(endpointsMap[svcPortName], epInfo)
			}
			if klog.V(3) {
				newEPList := []*endpointsInfo{}
				for _, ep := range endpointsMap[svcPortName] {
					newEPList = append(newEPList, ep)
				}
				klog.Infof("Setting endpoints for %q to %+v", svcPortName, newEPList)
			}
		}
	}
	return endpointsMap
}
func serviceToServiceMap(service *v1.Service) proxyServiceMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if service == nil {
		return nil
	}
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if shouldSkipService(svcName, service) {
		return nil
	}
	serviceMap := make(proxyServiceMap)
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		svcPortName := proxy.ServicePortName{NamespacedName: svcName, Port: servicePort.Name}
		serviceMap[svcPortName] = newServiceInfo(svcPortName, servicePort, service)
	}
	return serviceMap
}
func (proxier *Proxier) syncProxyRules() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxier.mu.Lock()
	defer proxier.mu.Unlock()
	start := time.Now()
	defer func() {
		SyncProxyRulesLatency.Observe(sinceInMicroseconds(start))
		klog.V(4).Infof("syncProxyRules took %v", time.Since(start))
	}()
	if !proxier.endpointsSynced || !proxier.servicesSynced {
		klog.V(2).Info("Not syncing hns until Services and Endpoints have been received from master")
		return
	}
	serviceUpdateResult := proxier.updateServiceMap()
	endpointUpdateResult := proxier.updateEndpointsMap()
	staleServices := serviceUpdateResult.staleServices
	for svcPortName := range endpointUpdateResult.staleServiceNames {
		if svcInfo, ok := proxier.serviceMap[svcPortName]; ok && svcInfo != nil && svcInfo.protocol == v1.ProtocolUDP {
			klog.V(2).Infof("Stale udp service %v -> %s", svcPortName, svcInfo.clusterIP.String())
			staleServices.Insert(svcInfo.clusterIP.String())
		}
	}
	klog.V(3).Infof("Syncing Policies")
	for svcName, svcInfo := range proxier.serviceMap {
		if svcInfo.policyApplied {
			klog.V(4).Infof("Policy already applied for %s", spew.Sdump(svcInfo))
			continue
		}
		var hnsEndpoints []hcsshim.HNSEndpoint
		klog.V(4).Infof("====Applying Policy for %s====", svcName)
		for _, ep := range proxier.endpointsMap[svcName] {
			var newHnsEndpoint *hcsshim.HNSEndpoint
			hnsNetworkName := proxier.network.name
			var err error
			if svcInfo.targetPort == 0 {
				svcInfo.targetPort = int(ep.port)
			}
			if len(ep.hnsID) > 0 {
				newHnsEndpoint, err = hcsshim.GetHNSEndpointByID(ep.hnsID)
			}
			if newHnsEndpoint == nil {
				newHnsEndpoint, err = getHnsEndpointByIpAddress(net.ParseIP(ep.ip), hnsNetworkName)
			}
			if newHnsEndpoint == nil {
				if ep.isLocal {
					klog.Errorf("Local endpoint not found for %v: err: %v on network %s", ep.ip, err, hnsNetworkName)
					continue
				}
				hnsnetwork, err := hcsshim.GetHNSNetworkByName(hnsNetworkName)
				if err != nil {
					klog.Errorf("%v", err)
					continue
				}
				hnsEndpoint := &hcsshim.HNSEndpoint{MacAddress: ep.macAddress, IPAddress: net.ParseIP(ep.ip)}
				newHnsEndpoint, err = hnsnetwork.CreateRemoteEndpoint(hnsEndpoint)
				if err != nil {
					klog.Errorf("Remote endpoint creation failed: %v", err)
					continue
				}
			}
			LogJson(newHnsEndpoint, "Hns Endpoint resource", 1)
			hnsEndpoints = append(hnsEndpoints, *newHnsEndpoint)
			ep.hnsID = newHnsEndpoint.Id
			ep.refCount++
			Log(ep, "Endpoint resource found", 3)
		}
		klog.V(3).Infof("Associated endpoints [%s] for service [%s]", spew.Sdump(hnsEndpoints), svcName)
		if len(svcInfo.hnsID) > 0 {
			klog.Warningf("Load Balancer already exists %s -- Debug ", svcInfo.hnsID)
		}
		if len(hnsEndpoints) == 0 {
			klog.Errorf("Endpoint information not available for service %s. Not applying any policy", svcName)
			continue
		}
		klog.V(4).Infof("Trying to Apply Policies for service %s", spew.Sdump(svcInfo))
		var hnsLoadBalancer *hcsshim.PolicyList
		hnsLoadBalancer, err := getHnsLoadBalancer(hnsEndpoints, false, svcInfo.clusterIP.String(), Enum(svcInfo.protocol), uint16(svcInfo.targetPort), uint16(svcInfo.port))
		if err != nil {
			klog.Errorf("Policy creation failed: %v", err)
			continue
		}
		svcInfo.hnsID = hnsLoadBalancer.ID
		klog.V(3).Infof("Hns LoadBalancer resource created for cluster ip resources %v, Id [%s]", svcInfo.clusterIP, hnsLoadBalancer.ID)
		if svcInfo.nodePort > 0 {
			hnsLoadBalancer, err := getHnsLoadBalancer(hnsEndpoints, false, "", Enum(svcInfo.protocol), uint16(svcInfo.targetPort), uint16(svcInfo.nodePort))
			if err != nil {
				klog.Errorf("Policy creation failed: %v", err)
				continue
			}
			svcInfo.nodePorthnsID = hnsLoadBalancer.ID
			klog.V(3).Infof("Hns LoadBalancer resource created for nodePort resources %v, Id [%s]", svcInfo.clusterIP, hnsLoadBalancer.ID)
		}
		for _, externalIp := range svcInfo.externalIPs {
			hnsLoadBalancer, err := getHnsLoadBalancer(hnsEndpoints, false, externalIp.ip, Enum(svcInfo.protocol), uint16(svcInfo.targetPort), uint16(svcInfo.port))
			if err != nil {
				klog.Errorf("Policy creation failed: %v", err)
				continue
			}
			externalIp.hnsID = hnsLoadBalancer.ID
			klog.V(3).Infof("Hns LoadBalancer resource created for externalIp resources %v, Id[%s]", externalIp, hnsLoadBalancer.ID)
		}
		for _, lbIngressIp := range svcInfo.loadBalancerIngressIPs {
			hnsLoadBalancer, err := getHnsLoadBalancer(hnsEndpoints, false, lbIngressIp.ip, Enum(svcInfo.protocol), uint16(svcInfo.targetPort), uint16(svcInfo.port))
			if err != nil {
				klog.Errorf("Policy creation failed: %v", err)
				continue
			}
			lbIngressIp.hnsID = hnsLoadBalancer.ID
			klog.V(3).Infof("Hns LoadBalancer resource created for loadBalancer Ingress resources %v", lbIngressIp)
		}
		svcInfo.policyApplied = true
		Log(svcInfo, "+++Policy Successfully applied for service +++", 2)
	}
	if proxier.healthzServer != nil {
		proxier.healthzServer.UpdateTimestamp()
	}
	if err := proxier.healthChecker.SyncServices(serviceUpdateResult.hcServices); err != nil {
		klog.Errorf("Error syncing healthcheck services: %v", err)
	}
	if err := proxier.healthChecker.SyncEndpoints(endpointUpdateResult.hcEndpoints); err != nil {
		klog.Errorf("Error syncing healthcheck endpoints: %v", err)
	}
	for _, svcIP := range staleServices.UnsortedList() {
		klog.V(5).Infof("Pending delete stale service IP %s connections", svcIP)
	}
}

type endpointServicePair struct {
	endpoint        string
	servicePortName proxy.ServicePortName
}
