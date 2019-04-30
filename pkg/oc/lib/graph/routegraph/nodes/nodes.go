package nodes

import (
	"github.com/gonum/graph"
	routev1 "github.com/openshift/api/route/v1"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
)

func EnsureRouteNode(g osgraph.MutableUniqueGraph, route *routev1.Route) *RouteNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, RouteNodeName(route), func(node osgraph.Node) graph.Node {
		return &RouteNode{Node: node, Route: route}
	}).(*RouteNode)
}
