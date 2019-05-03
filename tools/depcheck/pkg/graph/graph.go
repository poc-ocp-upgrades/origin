package graph

import (
	"fmt"
	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"github.com/gonum/graph/simple"
	"strings"
)

type Node struct {
	Id         int
	UniqueName string
	Color      string
}

func (n Node) ID() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Id
}
func (n Node) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return labelNameForNode(n.UniqueName)
}
func (n Node) DOTAttributes() []dot.Attribute {
	_logClusterCodePath()
	defer _logClusterCodePath()
	color := n.Color
	if len(color) == 0 {
		color = "black"
	}
	return []dot.Attribute{{Key: "label", Value: fmt.Sprintf("%q", n)}, {Key: "color", Value: color}}
}
func labelNameForNode(importPath string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	segs := strings.Split(importPath, "/vendor/")
	if len(segs) > 1 {
		return segs[1]
	}
	return importPath
}
func NewMutableDirectedGraph(roots []string) *MutableDirectedGraph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MutableDirectedGraph{DirectedGraph: simple.NewDirectedGraph(1.0, 0.0), nodesByName: make(map[string]graph.Node), rootNodeNames: roots}
}

type MutableDirectedGraph struct {
	*simple.DirectedGraph
	nodesByName   map[string]graph.Node
	rootNodeNames []string
}

func (g *MutableDirectedGraph) AddNode(n *Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, exists := g.nodesByName[n.UniqueName]; exists {
		return fmt.Errorf("node .UniqueName collision: %s", n.UniqueName)
	}
	g.nodesByName[n.UniqueName] = n
	g.DirectedGraph.AddNode(n)
	return nil
}
func (g *MutableDirectedGraph) RemoveNode(n *Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(g.nodesByName, n.UniqueName)
	g.DirectedGraph.RemoveNode(n)
}
func (g *MutableDirectedGraph) NodeByName(name string) (graph.Node, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, exists := g.nodesByName[name]
	return n, exists && g.DirectedGraph.Has(n)
}
func (g *MutableDirectedGraph) PruneOrphans() []*Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	removed := []*Node{}
	for _, n := range g.nodesByName {
		node, ok := n.(*Node)
		if !ok {
			continue
		}
		if len(g.To(n)) > 0 {
			continue
		}
		if contains(node.UniqueName, g.rootNodeNames) {
			continue
		}
		g.RemoveNode(node)
		removed = append(removed, node)
	}
	if len(removed) == 0 {
		return []*Node{}
	}
	return append(removed, g.PruneOrphans()...)
}
func contains(needle string, haystack []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}
	return false
}
func (g *MutableDirectedGraph) Copy() *MutableDirectedGraph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newGraph := NewMutableDirectedGraph(g.rootNodeNames)
	for _, n := range g.Nodes() {
		newNode, ok := n.(*Node)
		if !ok {
			continue
		}
		if err := newGraph.AddNode(newNode); err != nil {
			panic(fmt.Errorf("unexpected error while copying graph: %v", err))
		}
	}
	for _, n := range g.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			continue
		}
		if _, exists := newGraph.NodeByName(node.UniqueName); !exists {
			continue
		}
		from := g.From(n)
		for _, to := range from {
			if newGraph.HasEdgeFromTo(n, to) {
				continue
			}
			newGraph.SetEdge(simple.Edge{F: n, T: to})
		}
	}
	return newGraph
}
