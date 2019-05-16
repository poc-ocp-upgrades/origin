package proxy

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	apiservice "k8s.io/kubernetes/pkg/api/v1/service"
	utilproxy "k8s.io/kubernetes/pkg/proxy/util"
	utilnet "k8s.io/kubernetes/pkg/util/net"
	"net"
	"reflect"
	"strings"
	"sync"
)

type BaseServiceInfo struct {
	ClusterIP                net.IP
	Port                     int
	Protocol                 v1.Protocol
	NodePort                 int
	LoadBalancerStatus       v1.LoadBalancerStatus
	SessionAffinityType      v1.ServiceAffinity
	StickyMaxAgeSeconds      int
	ExternalIPs              []string
	LoadBalancerSourceRanges []string
	HealthCheckNodePort      int
	OnlyNodeLocalEndpoints   bool
}

var _ ServicePort = &BaseServiceInfo{}

func (info *BaseServiceInfo) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s:%d/%s", info.ClusterIP, info.Port, info.Protocol)
}
func (info *BaseServiceInfo) ClusterIPString() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.ClusterIP.String()
}
func (info *BaseServiceInfo) GetProtocol() v1.Protocol {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.Protocol
}
func (info *BaseServiceInfo) GetHealthCheckNodePort() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.HealthCheckNodePort
}
func (info *BaseServiceInfo) GetNodePort() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.NodePort
}
func (info *BaseServiceInfo) ExternalIPStrings() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return info.ExternalIPs
}
func (sct *ServiceChangeTracker) newBaseServiceInfo(port *v1.ServicePort, service *v1.Service) *BaseServiceInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	onlyNodeLocalEndpoints := false
	if apiservice.RequestsOnlyLocalTraffic(service) {
		onlyNodeLocalEndpoints = true
	}
	var stickyMaxAgeSeconds int
	if service.Spec.SessionAffinity == v1.ServiceAffinityClientIP {
		stickyMaxAgeSeconds = int(*service.Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds)
	}
	info := &BaseServiceInfo{ClusterIP: net.ParseIP(service.Spec.ClusterIP), Port: int(port.Port), Protocol: port.Protocol, NodePort: int(port.NodePort), LoadBalancerStatus: *service.Status.LoadBalancer.DeepCopy(), SessionAffinityType: service.Spec.SessionAffinity, StickyMaxAgeSeconds: stickyMaxAgeSeconds, OnlyNodeLocalEndpoints: onlyNodeLocalEndpoints}
	if sct.isIPv6Mode == nil {
		info.ExternalIPs = make([]string, len(service.Spec.ExternalIPs))
		info.LoadBalancerSourceRanges = make([]string, len(service.Spec.LoadBalancerSourceRanges))
		copy(info.LoadBalancerSourceRanges, service.Spec.LoadBalancerSourceRanges)
		copy(info.ExternalIPs, service.Spec.ExternalIPs)
	} else {
		var incorrectIPs []string
		info.ExternalIPs, incorrectIPs = utilnet.FilterIncorrectIPVersion(service.Spec.ExternalIPs, *sct.isIPv6Mode)
		if len(incorrectIPs) > 0 {
			utilproxy.LogAndEmitIncorrectIPVersionEvent(sct.recorder, "externalIPs", strings.Join(incorrectIPs, ","), service.Namespace, service.Name, service.UID)
		}
		info.LoadBalancerSourceRanges, incorrectIPs = utilnet.FilterIncorrectCIDRVersion(service.Spec.LoadBalancerSourceRanges, *sct.isIPv6Mode)
		if len(incorrectIPs) > 0 {
			utilproxy.LogAndEmitIncorrectIPVersionEvent(sct.recorder, "loadBalancerSourceRanges", strings.Join(incorrectIPs, ","), service.Namespace, service.Name, service.UID)
		}
	}
	if apiservice.NeedsHealthCheck(service) {
		p := service.Spec.HealthCheckNodePort
		if p == 0 {
			klog.Errorf("Service %s/%s has no healthcheck nodeport", service.Namespace, service.Name)
		} else {
			info.HealthCheckNodePort = int(p)
		}
	}
	return info
}

type makeServicePortFunc func(*v1.ServicePort, *v1.Service, *BaseServiceInfo) ServicePort
type serviceChange struct {
	previous ServiceMap
	current  ServiceMap
}
type ServiceChangeTracker struct {
	lock            sync.Mutex
	items           map[types.NamespacedName]*serviceChange
	makeServiceInfo makeServicePortFunc
	isIPv6Mode      *bool
	recorder        record.EventRecorder
}

func NewServiceChangeTracker(makeServiceInfo makeServicePortFunc, isIPv6Mode *bool, recorder record.EventRecorder) *ServiceChangeTracker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ServiceChangeTracker{items: make(map[types.NamespacedName]*serviceChange), makeServiceInfo: makeServiceInfo, isIPv6Mode: isIPv6Mode, recorder: recorder}
}
func (sct *ServiceChangeTracker) Update(previous, current *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svc := current
	if svc == nil {
		svc = previous
	}
	if svc == nil {
		return false
	}
	namespacedName := types.NamespacedName{Namespace: svc.Namespace, Name: svc.Name}
	sct.lock.Lock()
	defer sct.lock.Unlock()
	change, exists := sct.items[namespacedName]
	if !exists {
		change = &serviceChange{}
		change.previous = sct.serviceToServiceMap(previous)
		sct.items[namespacedName] = change
	}
	change.current = sct.serviceToServiceMap(current)
	if reflect.DeepEqual(change.previous, change.current) {
		delete(sct.items, namespacedName)
	}
	return len(sct.items) > 0
}

type UpdateServiceMapResult struct {
	HCServiceNodePorts map[types.NamespacedName]uint16
	UDPStaleClusterIP  sets.String
}

func UpdateServiceMap(serviceMap ServiceMap, changes *ServiceChangeTracker) (result UpdateServiceMapResult) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result.UDPStaleClusterIP = sets.NewString()
	serviceMap.apply(changes, result.UDPStaleClusterIP)
	result.HCServiceNodePorts = make(map[types.NamespacedName]uint16)
	for svcPortName, info := range serviceMap {
		if info.GetHealthCheckNodePort() != 0 {
			result.HCServiceNodePorts[svcPortName.NamespacedName] = uint16(info.GetHealthCheckNodePort())
		}
	}
	return result
}

type ServiceMap map[ServicePortName]ServicePort

func (sct *ServiceChangeTracker) serviceToServiceMap(service *v1.Service) ServiceMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if service == nil {
		return nil
	}
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	if utilproxy.ShouldSkipService(svcName, service) {
		return nil
	}
	if len(service.Spec.ClusterIP) != 0 {
		if sct.isIPv6Mode != nil && utilnet.IsIPv6String(service.Spec.ClusterIP) != *sct.isIPv6Mode {
			utilproxy.LogAndEmitIncorrectIPVersionEvent(sct.recorder, "clusterIP", service.Spec.ClusterIP, service.Namespace, service.Name, service.UID)
			return nil
		}
	}
	serviceMap := make(ServiceMap)
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		svcPortName := ServicePortName{NamespacedName: svcName, Port: servicePort.Name}
		baseSvcInfo := sct.newBaseServiceInfo(servicePort, service)
		if sct.makeServiceInfo != nil {
			serviceMap[svcPortName] = sct.makeServiceInfo(servicePort, service, baseSvcInfo)
		} else {
			serviceMap[svcPortName] = baseSvcInfo
		}
	}
	return serviceMap
}
func (serviceMap *ServiceMap) apply(changes *ServiceChangeTracker, UDPStaleClusterIP sets.String) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	changes.lock.Lock()
	defer changes.lock.Unlock()
	for _, change := range changes.items {
		serviceMap.merge(change.current)
		change.previous.filter(change.current)
		serviceMap.unmerge(change.previous, UDPStaleClusterIP)
	}
	changes.items = make(map[types.NamespacedName]*serviceChange)
	return
}
func (sm *ServiceMap) merge(other ServiceMap) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	existingPorts := sets.NewString()
	for svcPortName, info := range other {
		existingPorts.Insert(svcPortName.String())
		_, exists := (*sm)[svcPortName]
		if !exists {
			klog.V(1).Infof("Adding new service port %q at %s", svcPortName, info.String())
		} else {
			klog.V(1).Infof("Updating existing service port %q at %s", svcPortName, info.String())
		}
		(*sm)[svcPortName] = info
	}
	return existingPorts
}
func (sm *ServiceMap) filter(other ServiceMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range *sm {
		if _, ok := other[svcPortName]; ok {
			delete(*sm, svcPortName)
		}
	}
}
func (sm *ServiceMap) unmerge(other ServiceMap, UDPStaleClusterIP sets.String) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for svcPortName := range other {
		info, exists := (*sm)[svcPortName]
		if exists {
			klog.V(1).Infof("Removing service port %q", svcPortName)
			if info.GetProtocol() == v1.ProtocolUDP {
				UDPStaleClusterIP.Insert(info.ClusterIPString())
			}
			delete(*sm, svcPortName)
		} else {
			klog.Errorf("Service port %q doesn't exists", svcPortName)
		}
	}
}
