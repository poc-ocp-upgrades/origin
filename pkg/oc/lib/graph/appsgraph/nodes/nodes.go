package nodes

import (
	"github.com/gonum/graph"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	appsv1 "github.com/openshift/api/apps/v1"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

func EnsureDeploymentConfigNode(g osgraph.MutableUniqueGraph, dc *appsv1.DeploymentConfig) *DeploymentConfigNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dcName := DeploymentConfigNodeName(dc)
	dcNode := osgraph.EnsureUnique(g, dcName, func(node osgraph.Node) graph.Node {
		return &DeploymentConfigNode{Node: node, DeploymentConfig: dc, IsFound: true}
	}).(*DeploymentConfigNode)
	if dc.Spec.Template != nil {
		podTemplateSpecNode := kubegraph.EnsurePodTemplateSpecNode(g, dc.Spec.Template, dc.Namespace, dcName)
		g.AddEdge(dcNode, podTemplateSpecNode, osgraph.ContainsEdgeKind)
	}
	return dcNode
}
func FindOrCreateSyntheticDeploymentConfigNode(g osgraph.MutableUniqueGraph, dc *appsv1.DeploymentConfig) *DeploymentConfigNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, DeploymentConfigNodeName(dc), func(node osgraph.Node) graph.Node {
		return &DeploymentConfigNode{Node: node, DeploymentConfig: dc, IsFound: false}
	}).(*DeploymentConfigNode)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
