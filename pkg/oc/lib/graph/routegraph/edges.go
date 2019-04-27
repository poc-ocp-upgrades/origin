package routegraph

import (
	"github.com/gonum/graph"
	corev1 "k8s.io/api/core/v1"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
	routegraph "github.com/openshift/origin/pkg/oc/lib/graph/routegraph/nodes"
)

const (
	ExposedThroughRouteEdgeKind = "ExposedThroughRoute"
)

func AddRouteEdges(g osgraph.MutableUniqueGraph, node *routegraph.RouteNode) {
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
	syntheticService := &corev1.Service{}
	syntheticService.Namespace = node.Namespace
	syntheticService.Name = node.Spec.To.Name
	serviceNode := kubegraph.FindOrCreateSyntheticServiceNode(g, syntheticService)
	g.AddEdge(node, serviceNode, ExposedThroughRouteEdgeKind)
	for _, svc := range node.Spec.AlternateBackends {
		syntheticService := &corev1.Service{}
		syntheticService.Namespace = node.Namespace
		syntheticService.Name = svc.Name
		serviceNode := kubegraph.FindOrCreateSyntheticServiceNode(g, syntheticService)
		g.AddEdge(node, serviceNode, ExposedThroughRouteEdgeKind)
	}
}
func AddAllRouteEdges(g osgraph.MutableUniqueGraph) {
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
		if routeNode, ok := node.(*routegraph.RouteNode); ok {
			AddRouteEdges(g, routeNode)
		}
	}
}
