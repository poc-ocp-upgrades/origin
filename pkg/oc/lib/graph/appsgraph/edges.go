package appsgraph

import (
	"github.com/gonum/graph"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	imagegraph "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph/nodes"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	TriggersDeploymentEdgeKind	= "TriggersDeployment"
	UsedInDeploymentEdgeKind	= "UsedInDeployment"
	DeploymentEdgeKind		= "Deployment"
	VolumeClaimEdgeKind		= "VolumeClaim"
	ManagedByControllerEdgeKind	= "ManagedByController"
)

func AddTriggerDeploymentConfigsEdges(g osgraph.MutableUniqueGraph, node *appsgraph.DeploymentConfigNode) *appsgraph.DeploymentConfigNode {
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
	podTemplate := node.DeploymentConfig.Spec.Template
	if podTemplate == nil {
		return node
	}
	EachTemplateImage(&podTemplate.Spec, DeploymentConfigHasTrigger(node.DeploymentConfig), func(image TemplateImage, err error) {
		if err != nil {
			return
		}
		if image.From != nil {
			if len(image.From.Name) == 0 {
				return
			}
			name, tag, _ := imageapi.SplitImageStreamTag(image.From.Name)
			in := imagegraph.FindOrCreateSyntheticImageStreamTagNode(g, imagegraph.MakeImageStreamTagObjectMeta(image.From.Namespace, name, tag))
			g.AddEdge(in, node, TriggersDeploymentEdgeKind)
			return
		}
		tag := image.Ref.Tag
		image.Ref.Tag = ""
		in := imagegraph.EnsureDockerRepositoryNode(g, image.Ref.String(), tag)
		g.AddEdge(in, node, UsedInDeploymentEdgeKind)
	})
	return node
}
func AddAllTriggerDeploymentConfigsEdges(g osgraph.MutableUniqueGraph) {
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
	for _, node := range g.(graph.Graph).Nodes() {
		if dcNode, ok := node.(*appsgraph.DeploymentConfigNode); ok {
			AddTriggerDeploymentConfigsEdges(g, dcNode)
		}
	}
}
func AddDeploymentConfigsDeploymentEdges(g osgraph.MutableUniqueGraph, node *appsgraph.DeploymentConfigNode) *appsgraph.DeploymentConfigNode {
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
	for _, n := range g.(graph.Graph).Nodes() {
		if rcNode, ok := n.(*kubegraph.ReplicationControllerNode); ok {
			if rcNode.ReplicationController.Namespace != node.DeploymentConfig.Namespace {
				continue
			}
			if BelongsToDeploymentConfig(node.DeploymentConfig, rcNode.ReplicationController) {
				g.AddEdge(node, rcNode, DeploymentEdgeKind)
				g.AddEdge(rcNode, node, ManagedByControllerEdgeKind)
			}
		}
	}
	return node
}
func AddAllDeploymentConfigsDeploymentEdges(g osgraph.MutableUniqueGraph) {
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
	for _, node := range g.(graph.Graph).Nodes() {
		if dcNode, ok := node.(*appsgraph.DeploymentConfigNode); ok {
			AddDeploymentConfigsDeploymentEdges(g, dcNode)
		}
	}
}
func AddVolumeClaimEdges(g osgraph.Graph, dcNode *appsgraph.DeploymentConfigNode) {
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
	podTemplate := dcNode.DeploymentConfig.Spec.Template
	if podTemplate == nil {
		klog.Warningf("DeploymentConfig %s/%s template should not be empty", dcNode.DeploymentConfig.Namespace, dcNode.DeploymentConfig.Name)
		return
	}
	for _, volume := range podTemplate.Spec.Volumes {
		source := volume.VolumeSource
		if source.PersistentVolumeClaim == nil {
			continue
		}
		syntheticClaim := &corev1.PersistentVolumeClaim{ObjectMeta: metav1.ObjectMeta{Name: source.PersistentVolumeClaim.ClaimName, Namespace: dcNode.DeploymentConfig.Namespace}}
		pvcNode := kubegraph.FindOrCreateSyntheticPVCNode(g, syntheticClaim)
		g.AddEdge(dcNode, pvcNode, VolumeClaimEdgeKind)
	}
}
func AddAllVolumeClaimEdges(g osgraph.Graph) {
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
	for _, node := range g.Nodes() {
		if dcNode, ok := node.(*appsgraph.DeploymentConfigNode); ok && dcNode.IsFound {
			AddVolumeClaimEdges(g, dcNode)
		}
	}
}
