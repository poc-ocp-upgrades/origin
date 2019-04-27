package graphview

import (
	"github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	"github.com/openshift/origin/pkg/oc/lib/graph/kubegraph"
	kubenodes "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

type Job struct {
	Job		*kubenodes.JobNode
	OwnedPods	[]*kubenodes.PodNode
	CreatedPods	[]*kubenodes.PodNode
	Images		[]ImagePipeline
}

func AllJobs(g osgraph.Graph, excludeNodeIDs IntSet) ([]Job, IntSet) {
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
	covered := IntSet{}
	views := []Job{}
	for _, uncastNode := range g.NodesByKind(kubenodes.JobNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		view, covers := NewJob(g, uncastNode.(*kubenodes.JobNode))
		covered.Insert(covers.List()...)
		views = append(views, view)
	}
	return views, covered
}
func NewJob(g osgraph.Graph, node *kubenodes.JobNode) (Job, IntSet) {
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
	covered := IntSet{}
	covered.Insert(node.ID())
	view := Job{}
	view.Job = node
	for _, uncastPodNode := range g.PredecessorNodesByEdgeKind(node, appsgraph.ManagedByControllerEdgeKind) {
		podNode := uncastPodNode.(*kubenodes.PodNode)
		covered.Insert(podNode.ID())
		view.OwnedPods = append(view.OwnedPods, podNode)
	}
	for _, istNode := range g.PredecessorNodesByEdgeKind(node, kubegraph.TriggersDeploymentEdgeKind) {
		imagePipeline, covers := NewImagePipelineFromImageTagLocation(g, istNode, istNode.(ImageTagLocation))
		covered.Insert(covers.List()...)
		view.Images = append(view.Images, imagePipeline)
	}
	for _, tagNode := range g.PredecessorNodesByEdgeKind(node, appsgraph.UsedInDeploymentEdgeKind) {
		imagePipeline, covers := NewImagePipelineFromImageTagLocation(g, tagNode, tagNode.(ImageTagLocation))
		covered.Insert(covers.List()...)
		view.Images = append(view.Images, imagePipeline)
	}
	return view, covered
}
