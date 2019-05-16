package proxy

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	utilproxy "k8s.io/kubernetes/pkg/proxy/util"
	utilnet "k8s.io/kubernetes/pkg/util/net"
	"net"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strconv"
	"sync"
	gotime "time"
)

type BaseEndpointInfo struct {
	Endpoint string
	IsLocal  bool
}

var _ Endpoint = &BaseEndpointInfo{}

func (info *BaseEndpointInfo) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.Endpoint
}
func (info *BaseEndpointInfo) GetIsLocal() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.IsLocal
}
func (info *BaseEndpointInfo) IP() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return utilproxy.IPPart(info.Endpoint)
}
func (info *BaseEndpointInfo) Port() (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return utilproxy.PortPart(info.Endpoint)
}
func (info *BaseEndpointInfo) Equal(other Endpoint) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.String() == other.String() && info.GetIsLocal() == other.GetIsLocal()
}
func newBaseEndpointInfo(IP string, port int, isLocal bool) *BaseEndpointInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &BaseEndpointInfo{Endpoint: net.JoinHostPort(IP, strconv.Itoa(port)), IsLocal: isLocal}
}

type makeEndpointFunc func(info *BaseEndpointInfo) Endpoint
type EndpointChangeTracker struct {
	lock             sync.Mutex
	hostname         string
	items            map[types.NamespacedName]*endpointsChange
	makeEndpointInfo makeEndpointFunc
	isIPv6Mode       *bool
	recorder         record.EventRecorder
}

func NewEndpointChangeTracker(hostname string, makeEndpointInfo makeEndpointFunc, isIPv6Mode *bool, recorder record.EventRecorder) *EndpointChangeTracker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &EndpointChangeTracker{hostname: hostname, items: make(map[types.NamespacedName]*endpointsChange), makeEndpointInfo: makeEndpointInfo, isIPv6Mode: isIPv6Mode, recorder: recorder}
}
func (ect *EndpointChangeTracker) Update(previous, current *v1.Endpoints) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	endpoints := current
	if endpoints == nil {
		endpoints = previous
	}
	if endpoints == nil {
		return false
	}
	namespacedName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	ect.lock.Lock()
	defer ect.lock.Unlock()
	change, exists := ect.items[namespacedName]
	if !exists {
		change = &endpointsChange{}
		change.previous = ect.endpointsToEndpointsMap(previous)
		ect.items[namespacedName] = change
	}
	change.current = ect.endpointsToEndpointsMap(current)
	if reflect.DeepEqual(change.previous, change.current) {
		delete(ect.items, namespacedName)
	}
	return len(ect.items) > 0
}

type endpointsChange struct {
	previous EndpointsMap
	current  EndpointsMap
}
type UpdateEndpointMapResult struct {
	HCEndpointsLocalIPSize map[types.NamespacedName]int
	StaleEndpoints         []ServiceEndpoint
	StaleServiceNames      []ServicePortName
}

func UpdateEndpointsMap(endpointsMap EndpointsMap, changes *EndpointChangeTracker) (result UpdateEndpointMapResult) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result.StaleEndpoints = make([]ServiceEndpoint, 0)
	result.StaleServiceNames = make([]ServicePortName, 0)
	endpointsMap.apply(changes, &result.StaleEndpoints, &result.StaleServiceNames)
	result.HCEndpointsLocalIPSize = make(map[types.NamespacedName]int)
	localIPs := GetLocalEndpointIPs(endpointsMap)
	for nsn, ips := range localIPs {
		result.HCEndpointsLocalIPSize[nsn] = len(ips)
	}
	return result
}

type EndpointsMap map[ServicePortName][]Endpoint

func (ect *EndpointChangeTracker) endpointsToEndpointsMap(endpoints *v1.Endpoints) EndpointsMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if endpoints == nil {
		return nil
	}
	endpointsMap := make(EndpointsMap)
	for i := range endpoints.Subsets {
		ss := &endpoints.Subsets[i]
		for i := range ss.Ports {
			port := &ss.Ports[i]
			if port.Port == 0 {
				klog.Warningf("ignoring invalid endpoint port %s", port.Name)
				continue
			}
			svcPortName := ServicePortName{NamespacedName: types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}, Port: port.Name}
			for i := range ss.Addresses {
				addr := &ss.Addresses[i]
				if addr.IP == "" {
					klog.Warningf("ignoring invalid endpoint port %s with empty host", port.Name)
					continue
				}
				if ect.isIPv6Mode != nil && utilnet.IsIPv6String(addr.IP) != *ect.isIPv6Mode {
					utilproxy.LogAndEmitIncorrectIPVersionEvent(ect.recorder, "endpoints", addr.IP, endpoints.Name, endpoints.Namespace, "")
					continue
				}
				isLocal := addr.NodeName != nil && *addr.NodeName == ect.hostname
				baseEndpointInfo := newBaseEndpointInfo(addr.IP, int(port.Port), isLocal)
				if ect.makeEndpointInfo != nil {
					endpointsMap[svcPortName] = append(endpointsMap[svcPortName], ect.makeEndpointInfo(baseEndpointInfo))
				} else {
					endpointsMap[svcPortName] = append(endpointsMap[svcPortName], baseEndpointInfo)
				}
			}
			if klog.V(3) {
				newEPList := []string{}
				for _, ep := range endpointsMap[svcPortName] {
					newEPList = append(newEPList, ep.String())
				}
				klog.Infof("Setting endpoints for %q to %+v", svcPortName, newEPList)
			}
		}
	}
	return endpointsMap
}
func (endpointsMap EndpointsMap) apply(changes *EndpointChangeTracker, staleEndpoints *[]ServiceEndpoint, staleServiceNames *[]ServicePortName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if changes == nil {
		return
	}
	changes.lock.Lock()
	defer changes.lock.Unlock()
	for _, change := range changes.items {
		endpointsMap.Unmerge(change.previous)
		endpointsMap.Merge(change.current)
		detectStaleConnections(change.previous, change.current, staleEndpoints, staleServiceNames)
	}
	changes.items = make(map[types.NamespacedName]*endpointsChange)
}
func (em EndpointsMap) Merge(other EndpointsMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		em[svcPortName] = other[svcPortName]
	}
}
func (em EndpointsMap) Unmerge(other EndpointsMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		delete(em, svcPortName)
	}
}
func GetLocalEndpointIPs(endpointsMap EndpointsMap) map[types.NamespacedName]sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	localIPs := make(map[types.NamespacedName]sets.String)
	for svcPortName, epList := range endpointsMap {
		for _, ep := range epList {
			if ep.GetIsLocal() {
				nsn := svcPortName.NamespacedName
				if localIPs[nsn] == nil {
					localIPs[nsn] = sets.NewString()
				}
				localIPs[nsn].Insert(ep.IP())
			}
		}
	}
	return localIPs
}
func detectStaleConnections(oldEndpointsMap, newEndpointsMap EndpointsMap, staleEndpoints *[]ServiceEndpoint, staleServiceNames *[]ServicePortName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName, epList := range oldEndpointsMap {
		for _, ep := range epList {
			stale := true
			for i := range newEndpointsMap[svcPortName] {
				if newEndpointsMap[svcPortName][i].Equal(ep) {
					stale = false
					break
				}
			}
			if stale {
				klog.V(4).Infof("Stale endpoint %v -> %v", svcPortName, ep.String())
				*staleEndpoints = append(*staleEndpoints, ServiceEndpoint{Endpoint: ep.String(), ServicePortName: svcPortName})
			}
		}
	}
	for svcPortName, epList := range newEndpointsMap {
		if len(epList) > 0 && len(oldEndpointsMap[svcPortName]) == 0 {
			*staleServiceNames = append(*staleServiceNames, svcPortName)
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
