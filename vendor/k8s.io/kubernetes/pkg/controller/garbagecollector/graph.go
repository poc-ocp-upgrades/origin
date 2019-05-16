package garbagecollector

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sync"
)

type objectReference struct {
	metav1.OwnerReference
	Namespace string
}

func (s objectReference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("[%s/%s, namespace: %s, name: %s, uid: %s]", s.APIVersion, s.Kind, s.Namespace, s.Name, s.UID)
}

type node struct {
	identity               objectReference
	dependentsLock         sync.RWMutex
	dependents             map[*node]struct{}
	deletingDependents     bool
	deletingDependentsLock sync.RWMutex
	beingDeleted           bool
	beingDeletedLock       sync.RWMutex
	virtual                bool
	virtualLock            sync.RWMutex
	owners                 []metav1.OwnerReference
}

func (n *node) markBeingDeleted() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.beingDeletedLock.Lock()
	defer n.beingDeletedLock.Unlock()
	n.beingDeleted = true
}
func (n *node) isBeingDeleted() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.beingDeletedLock.RLock()
	defer n.beingDeletedLock.RUnlock()
	return n.beingDeleted
}
func (n *node) markObserved() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.virtualLock.Lock()
	defer n.virtualLock.Unlock()
	n.virtual = false
}
func (n *node) isObserved() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.virtualLock.RLock()
	defer n.virtualLock.RUnlock()
	return n.virtual == false
}
func (n *node) markDeletingDependents() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.deletingDependentsLock.Lock()
	defer n.deletingDependentsLock.Unlock()
	n.deletingDependents = true
}
func (n *node) isDeletingDependents() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.deletingDependentsLock.RLock()
	defer n.deletingDependentsLock.RUnlock()
	return n.deletingDependents
}
func (ownerNode *node) addDependent(dependent *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ownerNode.dependentsLock.Lock()
	defer ownerNode.dependentsLock.Unlock()
	ownerNode.dependents[dependent] = struct{}{}
}
func (ownerNode *node) deleteDependent(dependent *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ownerNode.dependentsLock.Lock()
	defer ownerNode.dependentsLock.Unlock()
	delete(ownerNode.dependents, dependent)
}
func (ownerNode *node) dependentsLength() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ownerNode.dependentsLock.RLock()
	defer ownerNode.dependentsLock.RUnlock()
	return len(ownerNode.dependents)
}
func (ownerNode *node) getDependents() []*node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ownerNode.dependentsLock.RLock()
	defer ownerNode.dependentsLock.RUnlock()
	var ret []*node
	for dep := range ownerNode.dependents {
		ret = append(ret, dep)
	}
	return ret
}
func (n *node) blockingDependents() []*node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dependents := n.getDependents()
	var ret []*node
	for _, dep := range dependents {
		for _, owner := range dep.owners {
			if owner.UID == n.identity.UID && owner.BlockOwnerDeletion != nil && *owner.BlockOwnerDeletion {
				ret = append(ret, dep)
			}
		}
	}
	return ret
}
func (n *node) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.dependentsLock.RLock()
	defer n.dependentsLock.RUnlock()
	return fmt.Sprintf("%#v", n)
}

type concurrentUIDToNode struct {
	uidToNodeLock sync.RWMutex
	uidToNode     map[types.UID]*node
}

func (m *concurrentUIDToNode) Write(node *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.uidToNodeLock.Lock()
	defer m.uidToNodeLock.Unlock()
	m.uidToNode[node.identity.UID] = node
}
func (m *concurrentUIDToNode) Read(uid types.UID) (*node, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.uidToNodeLock.RLock()
	defer m.uidToNodeLock.RUnlock()
	n, ok := m.uidToNode[uid]
	return n, ok
}
func (m *concurrentUIDToNode) Delete(uid types.UID) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.uidToNodeLock.Lock()
	defer m.uidToNodeLock.Unlock()
	delete(m.uidToNode, uid)
}
