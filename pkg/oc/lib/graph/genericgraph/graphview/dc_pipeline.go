package graphview

import (
	"sort"
	appsedges "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

type DeploymentConfigPipeline struct {
	DeploymentConfig	*appsgraph.DeploymentConfigNode
	ActiveDeployment	*kubegraph.ReplicationControllerNode
	InactiveDeployments	[]*kubegraph.ReplicationControllerNode
	Images			[]ImagePipeline
}

func AllDeploymentConfigPipelines(g osgraph.Graph, excludeNodeIDs IntSet) ([]DeploymentConfigPipeline, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	dcPipelines := []DeploymentConfigPipeline{}
	for _, uncastNode := range g.NodesByKind(appsgraph.DeploymentConfigNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		pipeline, covers := NewDeploymentConfigPipeline(g, uncastNode.(*appsgraph.DeploymentConfigNode))
		covered.Insert(covers.List()...)
		dcPipelines = append(dcPipelines, pipeline)
	}
	sort.Sort(SortedDeploymentConfigPipeline(dcPipelines))
	return dcPipelines, covered
}
func NewDeploymentConfigPipeline(g osgraph.Graph, dcNode *appsgraph.DeploymentConfigNode) (DeploymentConfigPipeline, IntSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(dcNode.ID())
	dcPipeline := DeploymentConfigPipeline{}
	dcPipeline.DeploymentConfig = dcNode
	for _, istNode := range g.PredecessorNodesByEdgeKind(dcNode, appsedges.TriggersDeploymentEdgeKind) {
		imagePipeline, covers := NewImagePipelineFromImageTagLocation(g, istNode, istNode.(ImageTagLocation))
		covered.Insert(covers.List()...)
		dcPipeline.Images = append(dcPipeline.Images, imagePipeline)
	}
	for _, tagNode := range g.PredecessorNodesByEdgeKind(dcNode, appsedges.UsedInDeploymentEdgeKind) {
		imagePipeline, covers := NewImagePipelineFromImageTagLocation(g, tagNode, tagNode.(ImageTagLocation))
		covered.Insert(covers.List()...)
		dcPipeline.Images = append(dcPipeline.Images, imagePipeline)
	}
	dcPipeline.ActiveDeployment, dcPipeline.InactiveDeployments = appsedges.RelevantDeployments(g, dcNode)
	for _, rc := range dcPipeline.InactiveDeployments {
		_, covers := NewReplicationController(g, rc)
		covered.Insert(covers.List()...)
	}
	if dcPipeline.ActiveDeployment != nil {
		_, covers := NewReplicationController(g, dcPipeline.ActiveDeployment)
		covered.Insert(covers.List()...)
	}
	return dcPipeline, covered
}

type SortedDeploymentConfigPipeline []DeploymentConfigPipeline

func (m SortedDeploymentConfigPipeline) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m SortedDeploymentConfigPipeline) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m SortedDeploymentConfigPipeline) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return CompareObjectMeta(&m[i].DeploymentConfig.DeploymentConfig.ObjectMeta, &m[j].DeploymentConfig.DeploymentConfig.ObjectMeta)
}
