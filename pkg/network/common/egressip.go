package common

import (
	"fmt"
	"net"
	"os"
	"sync"
	"syscall"
	"time"
	"k8s.io/klog"
	ktypes "k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	networkapi "github.com/openshift/api/network/v1"
	networkinformers "github.com/openshift/client-go/network/informers/externalversions/network/v1"
	"github.com/openshift/origin/pkg/util/netutils"
)

type nodeEgress struct {
	nodeName	string
	nodeIP		string
	sdnIP		string
	requestedIPs	sets.String
	requestedCIDRs	sets.String
	parsedCIDRs	map[string]*net.IPNet
	offline		bool
}
type namespaceEgress struct {
	vnid		uint32
	requestedIPs	[]string
	activeEgressIP	string
}
type egressIPInfo struct {
	ip		string
	parsed		net.IP
	nodes		[]*nodeEgress
	namespaces	[]*namespaceEgress
	assignedNodeIP	string
	assignedVNID	uint32
}
type EgressIPWatcher interface {
	ClaimEgressIP(vnid uint32, egressIP, nodeIP string)
	ReleaseEgressIP(egressIP, nodeIP string)
	SetNamespaceEgressNormal(vnid uint32)
	SetNamespaceEgressDropped(vnid uint32)
	SetNamespaceEgressViaEgressIP(vnid uint32, egressIP, nodeIP string)
	UpdateEgressCIDRs()
}
type EgressIPTracker struct {
	sync.Mutex
	watcher			EgressIPWatcher
	nodes			map[ktypes.UID]*nodeEgress
	nodesByNodeIP		map[string]*nodeEgress
	namespacesByVNID	map[uint32]*namespaceEgress
	egressIPs		map[string]*egressIPInfo
	nodesWithCIDRs		int
	changedEgressIPs	map[*egressIPInfo]bool
	changedNamespaces	map[*namespaceEgress]bool
	updateEgressCIDRs	bool
}

func NewEgressIPTracker(watcher EgressIPWatcher) *EgressIPTracker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &EgressIPTracker{watcher: watcher, nodes: make(map[ktypes.UID]*nodeEgress), nodesByNodeIP: make(map[string]*nodeEgress), namespacesByVNID: make(map[uint32]*namespaceEgress), egressIPs: make(map[string]*egressIPInfo), changedEgressIPs: make(map[*egressIPInfo]bool), changedNamespaces: make(map[*namespaceEgress]bool)}
}
func (eit *EgressIPTracker) Start(hostSubnetInformer networkinformers.HostSubnetInformer, netNamespaceInformer networkinformers.NetNamespaceInformer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.watchHostSubnets(hostSubnetInformer)
	eit.watchNetNamespaces(netNamespaceInformer)
}
func (eit *EgressIPTracker) ensureEgressIPInfo(egressIP string) *egressIPInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eg := eit.egressIPs[egressIP]
	if eg == nil {
		eg = &egressIPInfo{ip: egressIP, parsed: net.ParseIP(egressIP)}
		eit.egressIPs[egressIP] = eg
	}
	return eg
}
func (eit *EgressIPTracker) egressIPChanged(eg *egressIPInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.changedEgressIPs[eg] = true
	for _, ns := range eg.namespaces {
		eit.changedNamespaces[ns] = true
	}
}
func (eit *EgressIPTracker) addNodeEgressIP(node *nodeEgress, egressIP string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eg := eit.ensureEgressIPInfo(egressIP)
	eg.nodes = append(eg.nodes, node)
	eit.egressIPChanged(eg)
}
func (eit *EgressIPTracker) deleteNodeEgressIP(node *nodeEgress, egressIP string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eg := eit.egressIPs[egressIP]
	if eg == nil {
		return
	}
	for i := range eg.nodes {
		if eg.nodes[i] == node {
			eit.egressIPChanged(eg)
			eg.nodes = append(eg.nodes[:i], eg.nodes[i+1:]...)
			return
		}
	}
}
func (eit *EgressIPTracker) addNamespaceEgressIP(ns *namespaceEgress, egressIP string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eg := eit.ensureEgressIPInfo(egressIP)
	eg.namespaces = append(eg.namespaces, ns)
	eit.egressIPChanged(eg)
}
func (eit *EgressIPTracker) deleteNamespaceEgressIP(ns *namespaceEgress, egressIP string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eg := eit.egressIPs[egressIP]
	if eg == nil {
		return
	}
	for i := range eg.namespaces {
		if eg.namespaces[i] == ns {
			eit.egressIPChanged(eg)
			eg.namespaces = append(eg.namespaces[:i], eg.namespaces[i+1:]...)
			return
		}
	}
}
func (eit *EgressIPTracker) watchHostSubnets(hostSubnetInformer networkinformers.HostSubnetInformer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := InformerFuncs(&networkapi.HostSubnet{}, eit.handleAddOrUpdateHostSubnet, eit.handleDeleteHostSubnet)
	hostSubnetInformer.Informer().AddEventHandler(funcs)
}
func (eit *EgressIPTracker) handleAddOrUpdateHostSubnet(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", eventType, hs.Name)
	eit.UpdateHostSubnetEgress(hs)
}
func (eit *EgressIPTracker) handleDeleteHostSubnet(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", watch.Deleted, hs.Name)
	hs = hs.DeepCopy()
	hs.EgressCIDRs = nil
	hs.EgressIPs = nil
	eit.UpdateHostSubnetEgress(hs)
}
func (eit *EgressIPTracker) UpdateHostSubnetEgress(hs *networkapi.HostSubnet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.Lock()
	defer eit.Unlock()
	sdnIP := ""
	if hs.Subnet != "" {
		_, cidr, err := net.ParseCIDR(hs.Subnet)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("could not parse HostSubnet %q CIDR: %v", hs.Name, err))
		}
		sdnIP = netutils.GenerateDefaultGateway(cidr).String()
	}
	node := eit.nodes[hs.UID]
	if node == nil {
		if len(hs.EgressIPs) == 0 && len(hs.EgressCIDRs) == 0 {
			return
		}
		node = &nodeEgress{nodeName: hs.Host, nodeIP: hs.HostIP, sdnIP: sdnIP, requestedIPs: sets.NewString()}
		eit.nodes[hs.UID] = node
		eit.nodesByNodeIP[hs.HostIP] = node
	} else if len(hs.EgressIPs) == 0 && len(hs.EgressCIDRs) == 0 {
		delete(eit.nodes, hs.UID)
		delete(eit.nodesByNodeIP, node.nodeIP)
	}
	newRequestedCIDRs := sets.NewString(hs.EgressCIDRs...)
	if !node.requestedCIDRs.Equal(newRequestedCIDRs) {
		if len(hs.EgressCIDRs) == 0 {
			eit.nodesWithCIDRs--
		} else if node.requestedCIDRs.Len() == 0 {
			eit.nodesWithCIDRs++
		}
		node.requestedCIDRs = newRequestedCIDRs
		node.parsedCIDRs = make(map[string]*net.IPNet)
		for _, cidr := range hs.EgressCIDRs {
			_, parsed, _ := net.ParseCIDR(cidr)
			node.parsedCIDRs[cidr] = parsed
		}
		eit.updateEgressCIDRs = true
	}
	if node.nodeIP != hs.HostIP {
		movedEgressIPs := make([]string, 0, node.requestedIPs.Len())
		for _, ip := range node.requestedIPs.UnsortedList() {
			eg := eit.egressIPs[ip]
			if eg != nil && eg.assignedNodeIP == node.nodeIP {
				movedEgressIPs = append(movedEgressIPs, ip)
				eit.deleteNodeEgressIP(node, ip)
			}
		}
		eit.syncEgressIPs()
		delete(eit.nodesByNodeIP, node.nodeIP)
		node.nodeIP = hs.HostIP
		eit.nodesByNodeIP[node.nodeIP] = node
		for _, ip := range movedEgressIPs {
			eit.addNodeEgressIP(node, ip)
		}
	}
	oldRequestedIPs := node.requestedIPs
	node.requestedIPs = sets.NewString(hs.EgressIPs...)
	for _, ip := range node.requestedIPs.Difference(oldRequestedIPs).UnsortedList() {
		eit.addNodeEgressIP(node, ip)
	}
	for _, ip := range oldRequestedIPs.Difference(node.requestedIPs).UnsortedList() {
		eit.deleteNodeEgressIP(node, ip)
	}
	eit.syncEgressIPs()
}
func (eit *EgressIPTracker) watchNetNamespaces(netNamespaceInformer networkinformers.NetNamespaceInformer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := InformerFuncs(&networkapi.NetNamespace{}, eit.handleAddOrUpdateNetNamespace, eit.handleDeleteNetNamespace)
	netNamespaceInformer.Informer().AddEventHandler(funcs)
}
func (eit *EgressIPTracker) handleAddOrUpdateNetNamespace(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	netns := obj.(*networkapi.NetNamespace)
	klog.V(5).Infof("Watch %s event for NetNamespace %q", eventType, netns.Name)
	eit.UpdateNetNamespaceEgress(netns)
}
func (eit *EgressIPTracker) handleDeleteNetNamespace(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	netns := obj.(*networkapi.NetNamespace)
	klog.V(5).Infof("Watch %s event for NetNamespace %q", watch.Deleted, netns.Name)
	eit.DeleteNetNamespaceEgress(netns.NetID)
}
func (eit *EgressIPTracker) UpdateNetNamespaceEgress(netns *networkapi.NetNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.Lock()
	defer eit.Unlock()
	ns := eit.namespacesByVNID[netns.NetID]
	if ns == nil {
		if len(netns.EgressIPs) == 0 {
			return
		}
		ns = &namespaceEgress{vnid: netns.NetID}
		eit.namespacesByVNID[netns.NetID] = ns
	} else if len(netns.EgressIPs) == 0 {
		delete(eit.namespacesByVNID, netns.NetID)
	}
	oldRequestedIPs := sets.NewString(ns.requestedIPs...)
	newRequestedIPs := sets.NewString(netns.EgressIPs...)
	ns.requestedIPs = netns.EgressIPs
	for _, ip := range newRequestedIPs.Difference(oldRequestedIPs).UnsortedList() {
		eit.addNamespaceEgressIP(ns, ip)
	}
	for _, ip := range oldRequestedIPs.Difference(newRequestedIPs).UnsortedList() {
		eit.deleteNamespaceEgressIP(ns, ip)
	}
	for _, ip := range newRequestedIPs.Intersection(oldRequestedIPs).UnsortedList() {
		if eg := eit.egressIPs[ip]; eg != nil {
			eit.egressIPChanged(eg)
		}
	}
	eit.syncEgressIPs()
}
func (eit *EgressIPTracker) DeleteNetNamespaceEgress(vnid uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.UpdateNetNamespaceEgress(&networkapi.NetNamespace{NetID: vnid})
}
func (eit *EgressIPTracker) egressIPActive(eg *egressIPInfo) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(eg.nodes) == 0 || len(eg.namespaces) == 0 {
		return false, nil
	}
	if len(eg.nodes) > 1 {
		return false, fmt.Errorf("Multiple nodes (%s, %s) claiming EgressIP %s", eg.nodes[0].nodeIP, eg.nodes[1].nodeIP, eg.ip)
	}
	if len(eg.namespaces) > 1 {
		return false, fmt.Errorf("Multiple namespaces (%d, %d) claiming EgressIP %s", eg.namespaces[0].vnid, eg.namespaces[1].vnid, eg.ip)
	}
	for _, ip := range eg.namespaces[0].requestedIPs {
		eg2 := eit.egressIPs[ip]
		if eg2 != nil && eg2 != eg && len(eg2.nodes) == 1 && eg2.nodes[0] == eg.nodes[0] {
			return false, fmt.Errorf("Multiple EgressIPs (%s, %s) for VNID %d on node %s", eg.ip, eg2.ip, eg.namespaces[0].vnid, eg.nodes[0].nodeIP)
		}
	}
	return true, nil
}
func (eit *EgressIPTracker) syncEgressIPs() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	changedEgressIPs := eit.changedEgressIPs
	eit.changedEgressIPs = make(map[*egressIPInfo]bool)
	changedNamespaces := eit.changedNamespaces
	eit.changedNamespaces = make(map[*namespaceEgress]bool)
	for eg := range changedEgressIPs {
		active, err := eit.egressIPActive(eg)
		if err != nil {
			utilruntime.HandleError(err)
		}
		eit.syncEgressNodeState(eg, active)
	}
	for ns := range changedNamespaces {
		eit.syncEgressNamespaceState(ns)
	}
	for eg := range changedEgressIPs {
		if len(eg.namespaces) == 0 && len(eg.nodes) == 0 {
			delete(eit.egressIPs, eg.ip)
		}
	}
	if eit.updateEgressCIDRs {
		eit.updateEgressCIDRs = false
		if eit.nodesWithCIDRs > 0 {
			eit.watcher.UpdateEgressCIDRs()
		}
	}
}
func (eit *EgressIPTracker) syncEgressNodeState(eg *egressIPInfo, active bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if active && eg.assignedNodeIP != eg.nodes[0].nodeIP {
		klog.V(4).Infof("Assigning egress IP %s to node %s", eg.ip, eg.nodes[0].nodeIP)
		eg.assignedNodeIP = eg.nodes[0].nodeIP
		eit.watcher.ClaimEgressIP(eg.namespaces[0].vnid, eg.ip, eg.assignedNodeIP)
	} else if !active && eg.assignedNodeIP != "" {
		klog.V(4).Infof("Removing egress IP %s from node %s", eg.ip, eg.assignedNodeIP)
		eit.watcher.ReleaseEgressIP(eg.ip, eg.assignedNodeIP)
		eg.assignedNodeIP = ""
	}
	if eg.assignedNodeIP == "" {
		eit.updateEgressCIDRs = true
	}
}
func (eit *EgressIPTracker) syncEgressNamespaceState(ns *namespaceEgress) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ns.requestedIPs) == 0 {
		if ns.activeEgressIP != "" {
			ns.activeEgressIP = ""
			eit.watcher.SetNamespaceEgressNormal(ns.vnid)
		}
		return
	}
	var active *egressIPInfo
	for _, ip := range ns.requestedIPs {
		eg := eit.egressIPs[ip]
		if eg == nil {
			continue
		}
		if len(eg.namespaces) > 1 {
			active = nil
			klog.V(4).Infof("VNID %d gets no egress due to multiply-assigned egress IP %s", ns.vnid, eg.ip)
			break
		}
		if active == nil {
			if eg.assignedNodeIP == "" {
				klog.V(4).Infof("VNID %d cannot use unassigned egress IP %s", ns.vnid, eg.ip)
			} else if len(ns.requestedIPs) > 1 && eg.nodes[0].offline {
				klog.V(4).Infof("VNID %d cannot use egress IP %s on offline node %s", ns.vnid, eg.ip, eg.assignedNodeIP)
			} else {
				active = eg
			}
		}
	}
	if active != nil {
		if ns.activeEgressIP != active.ip {
			ns.activeEgressIP = active.ip
			eit.watcher.SetNamespaceEgressViaEgressIP(ns.vnid, active.ip, active.assignedNodeIP)
		}
	} else {
		if ns.activeEgressIP != "dropped" {
			ns.activeEgressIP = "dropped"
			eit.watcher.SetNamespaceEgressDropped(ns.vnid)
		}
	}
}
func (eit *EgressIPTracker) SetNodeOffline(nodeIP string, offline bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.Lock()
	defer eit.Unlock()
	node := eit.nodesByNodeIP[nodeIP]
	if node == nil {
		return
	}
	node.offline = offline
	for _, ip := range node.requestedIPs.UnsortedList() {
		eg := eit.egressIPs[ip]
		if eg != nil {
			eit.egressIPChanged(eg)
		}
	}
	if node.requestedCIDRs.Len() != 0 {
		eit.updateEgressCIDRs = true
	}
	eit.syncEgressIPs()
}
func (eit *EgressIPTracker) lookupNodeIP(ip string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.Lock()
	defer eit.Unlock()
	if node := eit.nodesByNodeIP[ip]; node != nil {
		return node.sdnIP
	}
	return ip
}
func (eit *EgressIPTracker) Ping(ip string, timeout time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip = eit.lookupNodeIP(ip)
	conn, err := net.DialTimeout("tcp", ip+":9", timeout)
	if conn != nil {
		conn.Close()
	}
	if opErr, ok := err.(*net.OpError); ok {
		if opErr.Timeout() {
			return false
		}
		if sysErr, ok := opErr.Err.(*os.SyscallError); ok && sysErr.Err == syscall.EHOSTUNREACH {
			return false
		}
	}
	return true
}
func (eit *EgressIPTracker) findEgressIPAllocation(ip net.IP, allocation map[string][]string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bestNode := ""
	otherNodes := false
	for _, node := range eit.nodes {
		if node.offline {
			continue
		}
		egressIPs := allocation[node.nodeName]
		for _, parsed := range node.parsedCIDRs {
			if parsed.Contains(ip) {
				if bestNode != "" {
					otherNodes = true
					if len(allocation[bestNode]) < len(egressIPs) {
						break
					}
				}
				bestNode = node.nodeName
				break
			}
		}
	}
	return bestNode, otherNodes
}
func (eit *EgressIPTracker) makeEmptyAllocation() (map[string][]string, map[string]bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allocation := make(map[string][]string)
	alreadyAllocated := make(map[string]bool)
	for egressIP, eip := range eit.egressIPs {
		if len(eip.namespaces) == 0 {
			alreadyAllocated[egressIP] = true
		} else if len(eip.nodes) > 1 || len(eip.namespaces) > 1 {
			alreadyAllocated[egressIP] = true
		} else if len(eip.namespaces) == 1 && len(eip.namespaces[0].requestedIPs) > 1 {
			alreadyAllocated[egressIP] = true
		}
	}
	return allocation, alreadyAllocated
}
func (eit *EgressIPTracker) allocateExistingEgressIPs(allocation map[string][]string, alreadyAllocated map[string]bool) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	removedEgressIPs := false
	for _, node := range eit.nodes {
		if len(node.parsedCIDRs) > 0 {
			allocation[node.nodeName] = make([]string, 0, node.requestedIPs.Len())
		}
	}
	for egressIP, eip := range eit.egressIPs {
		if eip.assignedNodeIP == "" || alreadyAllocated[egressIP] {
			continue
		}
		node := eip.nodes[0]
		found := false
		for _, parsed := range node.parsedCIDRs {
			if parsed.Contains(eip.parsed) {
				found = true
				break
			}
		}
		if found && !node.offline {
			allocation[node.nodeName] = append(allocation[node.nodeName], egressIP)
		} else {
			removedEgressIPs = true
		}
		alreadyAllocated[egressIP] = true
	}
	return removedEgressIPs
}
func (eit *EgressIPTracker) allocateNewEgressIPs(allocation map[string][]string, alreadyAllocated map[string]bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for egressIP, eip := range eit.egressIPs {
		if alreadyAllocated[egressIP] {
			continue
		}
		nodeName, otherNodes := eit.findEgressIPAllocation(eip.parsed, allocation)
		if nodeName != "" && !otherNodes {
			allocation[nodeName] = append(allocation[nodeName], egressIP)
			alreadyAllocated[egressIP] = true
		}
	}
	for egressIP, eip := range eit.egressIPs {
		if alreadyAllocated[egressIP] {
			continue
		}
		nodeName, _ := eit.findEgressIPAllocation(eip.parsed, allocation)
		if nodeName != "" {
			allocation[nodeName] = append(allocation[nodeName], egressIP)
		}
	}
}
func (eit *EgressIPTracker) ReallocateEgressIPs() map[string][]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eit.Lock()
	defer eit.Unlock()
	allocation, alreadyAllocated := eit.makeEmptyAllocation()
	removedEgressIPs := eit.allocateExistingEgressIPs(allocation, alreadyAllocated)
	eit.allocateNewEgressIPs(allocation, alreadyAllocated)
	if removedEgressIPs {
		return allocation
	}
	fullReallocation, alreadyAllocated := eit.makeEmptyAllocation()
	eit.allocateNewEgressIPs(fullReallocation, alreadyAllocated)
	emptyNodes := []string{}
	for nodeName, fullEgressIPs := range fullReallocation {
		incrementalEgressIPs := allocation[nodeName]
		if len(incrementalEgressIPs) < len(fullEgressIPs)/2 {
			emptyNodes = append(emptyNodes, nodeName)
		}
	}
	if len(emptyNodes) > 0 {
		allocation, alreadyAllocated = eit.makeEmptyAllocation()
		for _, nodeName := range emptyNodes {
			for _, egressIP := range fullReallocation[nodeName] {
				alreadyAllocated[egressIP] = true
			}
		}
		eit.allocateExistingEgressIPs(allocation, alreadyAllocated)
		eit.allocateNewEgressIPs(allocation, alreadyAllocated)
		eit.updateEgressCIDRs = true
	}
	return allocation
}
