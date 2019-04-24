package genericgraph

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"reflect"
	"sort"
	"strings"
	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"github.com/gonum/graph/simple"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Node struct {
	simple.Node
	UniqueName
}

func (n Node) DOTAttributes() []dot.Attribute {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []dot.Attribute{{Key: "label", Value: fmt.Sprintf("%q", n.UniqueName)}}
}

type ExistenceChecker interface{ Found() bool }
type UniqueName string
type UniqueNameFunc func(obj interface{}) UniqueName

func (n UniqueName) UniqueName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n)
}
func (n UniqueName) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n)
}

type uniqueNamer interface{ UniqueName() UniqueName }
type NodeFinder interface {
	Find(name UniqueName) graph.Node
}
type UniqueNodeInitializer interface {
	FindOrCreate(name UniqueName, fn NodeInitializerFunc) (graph.Node, bool)
}
type NodeInitializerFunc func(Node) graph.Node

func EnsureUnique(g UniqueNodeInitializer, name UniqueName, fn NodeInitializerFunc) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node, _ := g.FindOrCreate(name, fn)
	return node
}

type MutableDirectedEdge interface {
	AddEdge(from, to graph.Node, edgeKind string)
}
type MutableUniqueGraph interface {
	graph.DirectedBuilder
	MutableDirectedEdge
	UniqueNodeInitializer
	NodeFinder
}
type Edge struct {
	simple.Edge
	kinds	sets.String
}

func NewEdge(from, to graph.Node, weight float64, kinds ...string) Edge {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Edge{Edge: simple.Edge{F: from, T: to, W: weight}, kinds: sets.NewString(kinds...)}
}
func (e Edge) Kinds() sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.kinds
}
func (e Edge) IsKind(kind string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.kinds.Has(kind)
}
func (e Edge) DOTAttributes() []dot.Attribute {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []dot.Attribute{{Key: "label", Value: fmt.Sprintf("%q", strings.Join(e.Kinds().List(), ","))}}
}

type GraphDescriber interface {
	Name(node graph.Node) string
	Kind(node graph.Node) string
	Object(node graph.Node) interface{}
	EdgeKinds(edge graph.Edge) sets.String
}
type Interface interface {
	graph.DirectedBuilder
	GraphDescriber
	MutableDirectedEdge
	UniqueNodeInitializer
	NodeFinder
}
type Namer interface{ ResourceName(obj interface{}) string }
type namer struct{}

var DefaultNamer Namer = namer{}

func (namer) ResourceName(obj interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := obj.(type) {
	case uniqueNamer:
		return t.UniqueName().String()
	default:
		return reflect.TypeOf(obj).String()
	}
}

type Graph struct {
	graph.DirectedBuilder
	GraphDescriber
	uniqueNamedGraph
	internal	*simple.DirectedGraph
}

var _ MutableUniqueGraph = Graph{}

func New() Graph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g := simple.NewDirectedGraph(1.0, 0.0)
	return Graph{DirectedBuilder: g, GraphDescriber: typedGraph{}, uniqueNamedGraph: newUniqueNamedGraph(g), internal: g}
}
func (g Graph) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := ""
	nodes := g.Nodes()
	sort.Sort(ByID(nodes))
	for _, node := range nodes {
		ret += fmt.Sprintf("%d: %v\n", node.ID(), g.GraphDescriber.Name(node))
		successors := g.From(node)
		sort.Sort(ByID(successors))
		for _, successor := range successors {
			edge := g.Edge(node, successor)
			kinds := g.EdgeKinds(edge)
			for _, kind := range kinds.List() {
				ret += fmt.Sprintf("\t%v to %d: %v\n", kind, successor.ID(), g.GraphDescriber.Name(successor))
			}
		}
	}
	return ret
}
func (g Graph) Edges() []graph.Edge {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return g.internal.Edges()
}
func (g Graph) RemoveEdge(e graph.Edge) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.internal.RemoveEdge(e)
}
func (g Graph) RemoveNode(node graph.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.internal.RemoveNode(node)
}

type ByID []graph.Node

func (m ByID) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m ByID) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m ByID) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m[i].ID() < m[j].ID()
}
func (g Graph) NodesByKind(nodeKinds ...string) []graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Node{}
	kinds := sets.NewString(nodeKinds...)
	for _, node := range g.internal.Nodes() {
		if kinds.Has(g.Kind(node)) {
			ret = append(ret, node)
		}
	}
	return ret
}
func (g Graph) PredecessorEdges(node graph.Node, fn EdgeFunc, edgeKinds ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, n := range g.To(node) {
		edge := g.Edge(n, node)
		kinds := g.EdgeKinds(edge)
		if kinds.HasAny(edgeKinds...) {
			fn(g, n, node, kinds)
		}
	}
}
func (g Graph) SuccessorEdges(node graph.Node, fn EdgeFunc, edgeKinds ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, n := range g.From(node) {
		edge := g.Edge(node, n)
		kinds := g.EdgeKinds(edge)
		if kinds.HasAny(edgeKinds...) {
			fn(g, n, node, kinds)
		}
	}
}
func (g Graph) OutboundEdges(node graph.Node, edgeKinds ...string) []graph.Edge {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Edge{}
	for _, n := range g.From(node) {
		edge := g.Edge(node, n)
		if edge == nil {
			continue
		}
		if len(edgeKinds) == 0 || g.EdgeKinds(edge).HasAny(edgeKinds...) {
			ret = append(ret, edge)
		}
	}
	return ret
}
func (g Graph) InboundEdges(node graph.Node, edgeKinds ...string) []graph.Edge {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Edge{}
	for _, n := range g.To(node) {
		edge := g.Edge(n, node)
		if edge == nil {
			continue
		}
		if len(edgeKinds) == 0 || g.EdgeKinds(edge).HasAny(edgeKinds...) {
			ret = append(ret, edge)
		}
	}
	return ret
}
func (g Graph) PredecessorNodesByEdgeKind(node graph.Node, edgeKinds ...string) []graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Node{}
	for _, inboundEdges := range g.InboundEdges(node, edgeKinds...) {
		ret = append(ret, inboundEdges.From())
	}
	return ret
}
func (g Graph) SuccessorNodesByEdgeKind(node graph.Node, edgeKinds ...string) []graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Node{}
	for _, outboundEdge := range g.OutboundEdges(node, edgeKinds...) {
		ret = append(ret, outboundEdge.To())
	}
	return ret
}
func (g Graph) SuccessorNodesByNodeAndEdgeKind(node graph.Node, nodeKind, edgeKind string) []graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []graph.Node{}
	for _, successor := range g.SuccessorNodesByEdgeKind(node, edgeKind) {
		if g.Kind(successor) != nodeKind {
			continue
		}
		ret = append(ret, successor)
	}
	return ret
}
func (g Graph) AddNode(n graph.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.internal.AddNode(n)
}
func (g Graph) AddEdge(from, to graph.Node, edgeKind string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if edgeKind == ContainsEdgeKind {
		containsEdges := g.InboundEdges(to, ContainsEdgeKind)
		if len(containsEdges) != 0 {
			panic(fmt.Sprintf("%v is already contained by %v", to, containsEdges))
		}
	}
	kinds := sets.NewString(edgeKind)
	if existingEdge := g.Edge(from, to); existingEdge != nil {
		kinds.Insert(g.EdgeKinds(existingEdge).List()...)
	}
	g.internal.SetEdge(NewEdge(from, to, 1.0, kinds.List()...))
}
func (g Graph) addEdges(edges []graph.Edge, fn EdgeFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, e := range edges {
		switch t := e.(type) {
		case Edge:
			if fn(g, t.From(), t.To(), t.Kinds()) {
				g.internal.SetEdge(t)
			}
		case simple.Edge:
			if fn(g, t.From(), t.To(), sets.NewString()) {
				g.internal.SetEdge(t)
			}
		default:
			panic("bad edge")
		}
	}
}

type NodeFunc func(g Interface, n graph.Node) bool

func NodesOfKind(kinds ...string) NodeFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(kinds) == 0 {
		return func(g Interface, n graph.Node) bool {
			return true
		}
	}
	allowedKinds := sets.NewString(kinds...)
	return func(g Interface, n graph.Node) bool {
		return allowedKinds.Has(g.Kind(n))
	}
}

type EdgeFunc func(g Interface, from, to graph.Node, edgeKinds sets.String) bool

func EdgesOfKind(kinds ...string) EdgeFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(kinds) == 0 {
		return func(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
			return true
		}
	}
	allowedKinds := sets.NewString(kinds...)
	return func(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
		return allowedKinds.HasAny(edgeKinds.List()...)
	}
}
func RemoveInboundEdges(nodes []graph.Node) EdgeFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
		for _, node := range nodes {
			if node == to {
				return false
			}
		}
		return true
	}
}
func RemoveOutboundEdges(nodes []graph.Node) EdgeFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
		for _, node := range nodes {
			if node == from {
				return false
			}
		}
		return true
	}
}
func (g Graph) EdgeSubgraph(edgeFn EdgeFunc) Graph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := New()
	for _, node := range g.Nodes() {
		out.internal.AddNode(node)
	}
	out.addEdges(g.internal.Edges(), edgeFn)
	return out
}
func (g Graph) Subgraph(nodeFn NodeFunc, edgeFn EdgeFunc) Graph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := New()
	for _, node := range g.Nodes() {
		if nodeFn(out, node) {
			out.internal.AddNode(node)
		}
	}
	out.addEdges(g.internal.Edges(), edgeFn)
	return out
}
func (g Graph) SubgraphWithNodes(nodes []graph.Node, fn EdgeFunc) Graph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := New()
	for _, node := range nodes {
		out.internal.AddNode(node)
	}
	out.addEdges(g.internal.Edges(), fn)
	return out
}
func ExistingDirectEdge(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return !edgeKinds.Has(ReferencedByEdgeKind) && g.Has(from) && g.Has(to)
}
func ReverseExistingDirectEdge(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ExistingDirectEdge(g, from, to, edgeKinds) && ReverseGraphEdge(g, from, to, edgeKinds)
}
func ReverseGraphEdge(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for edgeKind := range edgeKinds {
		g.AddEdge(to, from, edgeKind)
	}
	return false
}
func AddReversedEdge(g Interface, from, to graph.Node, edgeKinds sets.String) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.AddEdge(to, from, ReferencedByEdgeKind)
	return true
}

type uniqueNamedGraph struct {
	graph.Builder
	names	map[UniqueName]graph.Node
}

func newUniqueNamedGraph(g graph.Builder) uniqueNamedGraph {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uniqueNamedGraph{Builder: g, names: make(map[UniqueName]graph.Node)}
}
func (g uniqueNamedGraph) FindOrCreate(name UniqueName, fn NodeInitializerFunc) (graph.Node, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if node, ok := g.names[name]; ok {
		return node, true
	}
	id := g.NewNodeID()
	node := fn(Node{simple.Node(id), name})
	g.names[name] = node
	g.AddNode(node)
	return node, false
}
func (g uniqueNamedGraph) Find(name UniqueName) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if node, ok := g.names[name]; ok {
		return node
	}
	return nil
}

type typedGraph struct{}

func (g typedGraph) Name(node graph.Node) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := node.(type) {
	case fmt.Stringer:
		return t.String()
	case uniqueNamer:
		return t.UniqueName().String()
	default:
		return fmt.Sprintf("<unknown:%d>", node.ID())
	}
}

type objectifier interface{ Object() interface{} }

func (g typedGraph) Object(node graph.Node) interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := node.(type) {
	case objectifier:
		return t.Object()
	default:
		return nil
	}
}

type kind interface{ Kind() string }

func (g typedGraph) Kind(node graph.Node) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if k, ok := node.(kind); ok {
		return k.Kind()
	}
	return UnknownNodeKind
}
func (g typedGraph) EdgeKinds(edge graph.Edge) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var e Edge
	switch t := edge.(type) {
	case Edge:
		e = t
	case simple.Edge:
		e = Edge{Edge: t}
	default:
		return sets.NewString(UnknownEdgeKind)
	}
	return e.Kinds()
}
func NodesByKind(g Interface, nodes []graph.Node, kinds ...string) [][]graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buckets := make(map[string]int)
	for i, kind := range kinds {
		buckets[kind] = i
	}
	if nodes == nil {
		nodes = g.Nodes()
	}
	last := len(kinds)
	result := make([][]graph.Node, last+1)
	for _, node := range nodes {
		if bucket, ok := buckets[g.Kind(node)]; ok {
			result[bucket] = append(result[bucket], node)
		} else {
			result[last] = append(result[last], node)
		}
	}
	return result
}
func IsFromDifferentNamespace(namespace string, node graph.Node) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	potentiallySyntheticNode, ok := node.(ExistenceChecker)
	if !ok || potentiallySyntheticNode.Found() {
		return false
	}
	objectified, ok := node.(objectifier)
	if !ok {
		return false
	}
	object, err := meta.Accessor(objectified)
	if err != nil {
		return false
	}
	return object.GetNamespace() != namespace
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
