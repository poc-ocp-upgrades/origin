package graphview

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	"github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/analysis"
	kubenodes "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

type ReplicationController struct {
	RC			*kubenodes.ReplicationControllerNode
	OwnedPods		[]*kubenodes.PodNode
	CreatedPods		[]*kubenodes.PodNode
	ConflictingRCs		[]*kubenodes.ReplicationControllerNode
	ConflictingRCIDToPods	map[int][]*kubenodes.PodNode
}

func AllReplicationControllers(g osgraph.Graph, excludeNodeIDs IntSet) ([]ReplicationController, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	rcViews := []ReplicationController{}
	for _, uncastNode := range g.NodesByKind(kubenodes.ReplicationControllerNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		rcView, covers := NewReplicationController(g, uncastNode.(*kubenodes.ReplicationControllerNode))
		covered.Insert(covers.List()...)
		rcViews = append(rcViews, rcView)
	}
	return rcViews, covered
}
func (rc *ReplicationController) MaxRecentContainerRestarts() int32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var maxRestarts int32
	for _, pod := range rc.OwnedPods {
		for _, status := range pod.Status.ContainerStatuses {
			if status.RestartCount > maxRestarts && analysis.ContainerRestartedRecently(status, metav1.Now()) {
				maxRestarts = status.RestartCount
			}
		}
	}
	return maxRestarts
}
func NewReplicationController(g osgraph.Graph, rcNode *kubenodes.ReplicationControllerNode) (ReplicationController, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(rcNode.ID())
	rcView := ReplicationController{}
	rcView.RC = rcNode
	rcView.ConflictingRCIDToPods = map[int][]*kubenodes.PodNode{}
	for _, uncastPodNode := range g.PredecessorNodesByEdgeKind(rcNode, appsgraph.ManagedByControllerEdgeKind) {
		podNode := uncastPodNode.(*kubenodes.PodNode)
		covered.Insert(podNode.ID())
		rcView.OwnedPods = append(rcView.OwnedPods, podNode)
		uncastOwningRCs := g.SuccessorNodesByEdgeKind(podNode, appsgraph.ManagedByControllerEdgeKind)
		if len(uncastOwningRCs) > 1 {
			for _, uncastOwningRC := range uncastOwningRCs {
				if uncastOwningRC.ID() == rcNode.ID() {
					continue
				}
				conflictingRC := uncastOwningRC.(*kubenodes.ReplicationControllerNode)
				rcView.ConflictingRCs = append(rcView.ConflictingRCs, conflictingRC)
				conflictingPods, ok := rcView.ConflictingRCIDToPods[conflictingRC.ID()]
				if !ok {
					conflictingPods = []*kubenodes.PodNode{}
				}
				conflictingPods = append(conflictingPods, podNode)
				rcView.ConflictingRCIDToPods[conflictingRC.ID()] = conflictingPods
			}
		}
	}
	return rcView, covered
}
func MaxRecentContainerRestartsForRC(g osgraph.Graph, rcNode *kubenodes.ReplicationControllerNode) int32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rcNode == nil {
		return 0
	}
	rc, _ := NewReplicationController(g, rcNode)
	return rc.MaxRecentContainerRestarts()
}
