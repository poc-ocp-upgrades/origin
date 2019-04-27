package genericgraph

import (
	"fmt"
	"github.com/gonum/graph"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	UnknownNodeKind = "UnknownNode"
)
const (
	UnknownEdgeKind		= "UnknownEdge"
	ReferencedByEdgeKind	= "ReferencedBy"
	ContainsEdgeKind	= "Contains"
)

func GetUniqueRuntimeObjectNodeName(nodeKind string, obj runtime.Object) UniqueName {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	meta, err := meta.Accessor(obj)
	if err != nil {
		panic(err)
	}
	return UniqueName(fmt.Sprintf("%s|%s/%s", nodeKind, meta.GetNamespace(), meta.GetName()))
}
func GetTopLevelContainerNode(g Graph, containedNode graph.Node) graph.Node {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	visited := map[int]bool{}
	prevContainingNode := containedNode
	for {
		visited[prevContainingNode.ID()] = true
		currContainingNode := GetContainingNode(g, prevContainingNode)
		if currContainingNode == nil {
			return prevContainingNode
		}
		if _, alreadyVisited := visited[currContainingNode.ID()]; alreadyVisited {
			panic(fmt.Sprintf("contains cycle in %v", visited))
		}
		prevContainingNode = currContainingNode
	}
}
func GetContainingNode(g Graph, containedNode graph.Node) graph.Node {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, node := range g.To(containedNode) {
		edge := g.Edge(node, containedNode)
		if g.EdgeKinds(edge).Has(ContainsEdgeKind) {
			return node
		}
	}
	return nil
}
