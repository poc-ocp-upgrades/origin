package cache

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/klog"
	utilnode "k8s.io/kubernetes/pkg/util/node"
	"sync"
)

type NodeTree struct {
	tree      map[string]*nodeArray
	zones     []string
	zoneIndex int
	NumNodes  int
	mu        sync.RWMutex
}
type nodeArray struct {
	nodes     []string
	lastIndex int
}

func (na *nodeArray) next() (nodeName string, exhausted bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(na.nodes) == 0 {
		klog.Error("The nodeArray is empty. It should have been deleted from NodeTree.")
		return "", false
	}
	if na.lastIndex >= len(na.nodes) {
		return "", true
	}
	nodeName = na.nodes[na.lastIndex]
	na.lastIndex++
	return nodeName, false
}
func newNodeTree(nodes []*v1.Node) *NodeTree {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nt := &NodeTree{tree: make(map[string]*nodeArray)}
	for _, n := range nodes {
		nt.AddNode(n)
	}
	return nt
}
func (nt *NodeTree) AddNode(n *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nt.mu.Lock()
	defer nt.mu.Unlock()
	nt.addNode(n)
}
func (nt *NodeTree) addNode(n *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	zone := utilnode.GetZoneKey(n)
	if na, ok := nt.tree[zone]; ok {
		for _, nodeName := range na.nodes {
			if nodeName == n.Name {
				klog.Warningf("node %v already exist in the NodeTree", n.Name)
				return
			}
		}
		na.nodes = append(na.nodes, n.Name)
	} else {
		nt.zones = append(nt.zones, zone)
		nt.tree[zone] = &nodeArray{nodes: []string{n.Name}, lastIndex: 0}
	}
	klog.V(5).Infof("Added node %v in group %v to NodeTree", n.Name, zone)
	nt.NumNodes++
}
func (nt *NodeTree) RemoveNode(n *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nt.mu.Lock()
	defer nt.mu.Unlock()
	return nt.removeNode(n)
}
func (nt *NodeTree) removeNode(n *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	zone := utilnode.GetZoneKey(n)
	if na, ok := nt.tree[zone]; ok {
		for i, nodeName := range na.nodes {
			if nodeName == n.Name {
				na.nodes = append(na.nodes[:i], na.nodes[i+1:]...)
				if len(na.nodes) == 0 {
					nt.removeZone(zone)
				}
				klog.V(5).Infof("Removed node %v in group %v from NodeTree", n.Name, zone)
				nt.NumNodes--
				return nil
			}
		}
	}
	klog.Errorf("Node %v in group %v was not found", n.Name, zone)
	return fmt.Errorf("node %v in group %v was not found", n.Name, zone)
}
func (nt *NodeTree) removeZone(zone string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	delete(nt.tree, zone)
	for i, z := range nt.zones {
		if z == zone {
			nt.zones = append(nt.zones[:i], nt.zones[i+1:]...)
		}
	}
}
func (nt *NodeTree) UpdateNode(old, new *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var oldZone string
	if old != nil {
		oldZone = utilnode.GetZoneKey(old)
	}
	newZone := utilnode.GetZoneKey(new)
	if oldZone == newZone {
		return
	}
	nt.mu.Lock()
	defer nt.mu.Unlock()
	nt.removeNode(old)
	nt.addNode(new)
}
func (nt *NodeTree) resetExhausted() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, na := range nt.tree {
		na.lastIndex = 0
	}
	nt.zoneIndex = 0
}
func (nt *NodeTree) Next() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nt.mu.Lock()
	defer nt.mu.Unlock()
	if len(nt.zones) == 0 {
		return ""
	}
	numExhaustedZones := 0
	for {
		if nt.zoneIndex >= len(nt.zones) {
			nt.zoneIndex = 0
		}
		zone := nt.zones[nt.zoneIndex]
		nt.zoneIndex++
		nodeName, exhausted := nt.tree[zone].next()
		if exhausted {
			numExhaustedZones++
			if numExhaustedZones >= len(nt.zones) {
				nt.resetExhausted()
			}
		} else {
			return nodeName
		}
	}
}
