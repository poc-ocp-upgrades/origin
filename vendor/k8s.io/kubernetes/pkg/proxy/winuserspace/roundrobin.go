package winuserspace

import (
 "errors"
 "fmt"
 "net"
 "reflect"
 "strconv"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/proxy"
 "k8s.io/kubernetes/pkg/util/slice"
)

var (
 ErrMissingServiceEntry = errors.New("missing service entry")
 ErrMissingEndpoints    = errors.New("missing endpoints")
)

type affinityState struct {
 clientIP string
 endpoint string
 lastUsed time.Time
}
type affinityPolicy struct {
 affinityType v1.ServiceAffinity
 affinityMap  map[string]*affinityState
 ttlSeconds   int
}
type LoadBalancerRR struct {
 lock     sync.RWMutex
 services map[proxy.ServicePortName]*balancerState
}

var _ LoadBalancer = &LoadBalancerRR{}

type balancerState struct {
 endpoints []string
 index     int
 affinity  affinityPolicy
}

func newAffinityPolicy(affinityType v1.ServiceAffinity, ttlSeconds int) *affinityPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &affinityPolicy{affinityType: affinityType, affinityMap: make(map[string]*affinityState), ttlSeconds: ttlSeconds}
}
func NewLoadBalancerRR() *LoadBalancerRR {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &LoadBalancerRR{services: map[proxy.ServicePortName]*balancerState{}}
}
func (lb *LoadBalancerRR) NewService(svcPort proxy.ServicePortName, affinityType v1.ServiceAffinity, ttlSeconds int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("LoadBalancerRR NewService %q", svcPort)
 lb.lock.Lock()
 defer lb.lock.Unlock()
 lb.newServiceInternal(svcPort, affinityType, ttlSeconds)
 return nil
}
func (lb *LoadBalancerRR) newServiceInternal(svcPort proxy.ServicePortName, affinityType v1.ServiceAffinity, ttlSeconds int) *balancerState {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ttlSeconds == 0 {
  ttlSeconds = int(v1.DefaultClientIPServiceAffinitySeconds)
 }
 if _, exists := lb.services[svcPort]; !exists {
  lb.services[svcPort] = &balancerState{affinity: *newAffinityPolicy(affinityType, ttlSeconds)}
  klog.V(4).Infof("LoadBalancerRR service %q did not exist, created", svcPort)
 } else if affinityType != "" {
  lb.services[svcPort].affinity.affinityType = affinityType
 }
 return lb.services[svcPort]
}
func (lb *LoadBalancerRR) DeleteService(svcPort proxy.ServicePortName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("LoadBalancerRR DeleteService %q", svcPort)
 lb.lock.Lock()
 defer lb.lock.Unlock()
 delete(lb.services, svcPort)
}
func isSessionAffinity(affinity *affinityPolicy) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if affinity.affinityType == "" || affinity.affinityType == v1.ServiceAffinityNone {
  return false
 }
 return true
}
func (lb *LoadBalancerRR) NextEndpoint(svcPort proxy.ServicePortName, srcAddr net.Addr, sessionAffinityReset bool) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 lb.lock.Lock()
 defer lb.lock.Unlock()
 state, exists := lb.services[svcPort]
 if !exists || state == nil {
  return "", ErrMissingServiceEntry
 }
 if len(state.endpoints) == 0 {
  return "", ErrMissingEndpoints
 }
 klog.V(4).Infof("NextEndpoint for service %q, srcAddr=%v: endpoints: %+v", svcPort, srcAddr, state.endpoints)
 sessionAffinityEnabled := isSessionAffinity(&state.affinity)
 var ipaddr string
 if sessionAffinityEnabled {
  var err error
  ipaddr, _, err = net.SplitHostPort(srcAddr.String())
  if err != nil {
   return "", fmt.Errorf("malformed source address %q: %v", srcAddr.String(), err)
  }
  if !sessionAffinityReset {
   sessionAffinity, exists := state.affinity.affinityMap[ipaddr]
   if exists && int(time.Since(sessionAffinity.lastUsed).Seconds()) < state.affinity.ttlSeconds {
    endpoint := sessionAffinity.endpoint
    sessionAffinity.lastUsed = time.Now()
    klog.V(4).Infof("NextEndpoint for service %q from IP %s with sessionAffinity %#v: %s", svcPort, ipaddr, sessionAffinity, endpoint)
    return endpoint, nil
   }
  }
 }
 endpoint := state.endpoints[state.index]
 state.index = (state.index + 1) % len(state.endpoints)
 if sessionAffinityEnabled {
  var affinity *affinityState
  affinity = state.affinity.affinityMap[ipaddr]
  if affinity == nil {
   affinity = new(affinityState)
   state.affinity.affinityMap[ipaddr] = affinity
  }
  affinity.lastUsed = time.Now()
  affinity.endpoint = endpoint
  affinity.clientIP = ipaddr
  klog.V(4).Infof("Updated affinity key %s: %#v", ipaddr, state.affinity.affinityMap[ipaddr])
 }
 return endpoint, nil
}

type hostPortPair struct {
 host string
 port int
}

func isValidEndpoint(hpp *hostPortPair) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return hpp.host != "" && hpp.port > 0
}
func flattenValidEndpoints(endpoints []hostPortPair) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var result []string
 for i := range endpoints {
  hpp := &endpoints[i]
  if isValidEndpoint(hpp) {
   result = append(result, net.JoinHostPort(hpp.host, strconv.Itoa(hpp.port)))
  }
 }
 return result
}
func removeSessionAffinityByEndpoint(state *balancerState, svcPort proxy.ServicePortName, endpoint string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, affinity := range state.affinity.affinityMap {
  if affinity.endpoint == endpoint {
   klog.V(4).Infof("Removing client: %s from affinityMap for service %q", affinity.endpoint, svcPort)
   delete(state.affinity.affinityMap, affinity.clientIP)
  }
 }
}
func (lb *LoadBalancerRR) updateAffinityMap(svcPort proxy.ServicePortName, newEndpoints []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allEndpoints := map[string]int{}
 for _, newEndpoint := range newEndpoints {
  allEndpoints[newEndpoint] = 1
 }
 state, exists := lb.services[svcPort]
 if !exists {
  return
 }
 for _, existingEndpoint := range state.endpoints {
  allEndpoints[existingEndpoint] = allEndpoints[existingEndpoint] + 1
 }
 for mKey, mVal := range allEndpoints {
  if mVal == 1 {
   klog.V(2).Infof("Delete endpoint %s for service %q", mKey, svcPort)
   removeSessionAffinityByEndpoint(state, svcPort, mKey)
  }
 }
}
func buildPortsToEndpointsMap(endpoints *v1.Endpoints) map[string][]hostPortPair {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portsToEndpoints := map[string][]hostPortPair{}
 for i := range endpoints.Subsets {
  ss := &endpoints.Subsets[i]
  for i := range ss.Ports {
   port := &ss.Ports[i]
   for i := range ss.Addresses {
    addr := &ss.Addresses[i]
    portsToEndpoints[port.Name] = append(portsToEndpoints[port.Name], hostPortPair{addr.IP, int(port.Port)})
   }
  }
 }
 return portsToEndpoints
}
func (lb *LoadBalancerRR) OnEndpointsAdd(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portsToEndpoints := buildPortsToEndpointsMap(endpoints)
 lb.lock.Lock()
 defer lb.lock.Unlock()
 for portname := range portsToEndpoints {
  svcPort := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: portname}
  newEndpoints := flattenValidEndpoints(portsToEndpoints[portname])
  state, exists := lb.services[svcPort]
  if !exists || state == nil || len(newEndpoints) > 0 {
   klog.V(1).Infof("LoadBalancerRR: Setting endpoints for %s to %+v", svcPort, newEndpoints)
   lb.updateAffinityMap(svcPort, newEndpoints)
   state = lb.newServiceInternal(svcPort, v1.ServiceAffinity(""), 0)
   state.endpoints = slice.ShuffleStrings(newEndpoints)
   state.index = 0
  }
 }
}
func (lb *LoadBalancerRR) OnEndpointsUpdate(oldEndpoints, endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portsToEndpoints := buildPortsToEndpointsMap(endpoints)
 oldPortsToEndpoints := buildPortsToEndpointsMap(oldEndpoints)
 registeredEndpoints := make(map[proxy.ServicePortName]bool)
 lb.lock.Lock()
 defer lb.lock.Unlock()
 for portname := range portsToEndpoints {
  svcPort := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: portname}
  newEndpoints := flattenValidEndpoints(portsToEndpoints[portname])
  state, exists := lb.services[svcPort]
  curEndpoints := []string{}
  if state != nil {
   curEndpoints = state.endpoints
  }
  if !exists || state == nil || len(curEndpoints) != len(newEndpoints) || !slicesEquiv(slice.CopyStrings(curEndpoints), newEndpoints) {
   klog.V(1).Infof("LoadBalancerRR: Setting endpoints for %s to %+v", svcPort, newEndpoints)
   lb.updateAffinityMap(svcPort, newEndpoints)
   state = lb.newServiceInternal(svcPort, v1.ServiceAffinity(""), 0)
   state.endpoints = slice.ShuffleStrings(newEndpoints)
   state.index = 0
  }
  registeredEndpoints[svcPort] = true
 }
 for portname := range oldPortsToEndpoints {
  svcPort := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: portname}
  if _, exists := registeredEndpoints[svcPort]; !exists {
   klog.V(2).Infof("LoadBalancerRR: Removing endpoints for %s", svcPort)
   state := lb.services[svcPort]
   state.endpoints = []string{}
   state.index = 0
   state.affinity.affinityMap = map[string]*affinityState{}
  }
 }
}
func (lb *LoadBalancerRR) OnEndpointsDelete(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 portsToEndpoints := buildPortsToEndpointsMap(endpoints)
 lb.lock.Lock()
 defer lb.lock.Unlock()
 for portname := range portsToEndpoints {
  svcPort := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: portname}
  klog.V(2).Infof("LoadBalancerRR: Removing endpoints for %s", svcPort)
  if state, ok := lb.services[svcPort]; ok {
   state.endpoints = []string{}
   state.index = 0
   state.affinity.affinityMap = map[string]*affinityState{}
  }
 }
}
func (lb *LoadBalancerRR) OnEndpointsSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func slicesEquiv(lhs, rhs []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(lhs) != len(rhs) {
  return false
 }
 if reflect.DeepEqual(slice.SortStrings(lhs), slice.SortStrings(rhs)) {
  return true
 }
 return false
}
func (lb *LoadBalancerRR) CleanupStaleStickySessions(svcPort proxy.ServicePortName) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 lb.lock.Lock()
 defer lb.lock.Unlock()
 state, exists := lb.services[svcPort]
 if !exists {
  return
 }
 for ip, affinity := range state.affinity.affinityMap {
  if int(time.Since(affinity.lastUsed).Seconds()) >= state.affinity.ttlSeconds {
   klog.V(4).Infof("Removing client %s from affinityMap for service %q", affinity.clientIP, svcPort)
   delete(state.affinity.affinityMap, ip)
  }
 }
}
