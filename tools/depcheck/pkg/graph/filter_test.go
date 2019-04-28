package graph

import (
	"strings"
	"testing"
	"github.com/gonum/graph/simple"
)

type testFilterNode struct {
	name		string
	outboundEdges	[]string
}

func getVendorNodes() []*testFilterNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []*testFilterNode{{name: "github.com/test/repo/vendor/github.com/testvendor/prefix", outboundEdges: []string{"github.com/test/repo/vendor/github.com/testvendor/prefix/one"}}, {name: "github.com/test/repo/vendor/github.com/testvendor/prefix/one", outboundEdges: []string{"github.com/test/repo/vendor/github.com/testvendor/prefix2/one"}}, {name: "github.com/test/repo/vendor/github.com/testvendor/prefix2", outboundEdges: []string{"github.com/test/repo/vendor/github.com/testvendor/prefix2/one"}}, {name: "github.com/test/repo/vendor/github.com/testvendor/prefix2/one", outboundEdges: []string{}}, {name: "github.com/test/repo/vendor/github.com/docker/docker-test-util", outboundEdges: []string{"github.com/test/repo/vendor/github.com/docker/docker-test-util/api"}}, {name: "github.com/test/repo/vendor/github.com/docker/docker-test-util/api", outboundEdges: []string{"github.com/test/repo/vendor/github.com/google/glog"}}, {name: "github.com/test/repo/vendor/github.com/google/glog", outboundEdges: []string{}}}
}
func getNonVendorNodes() []*testFilterNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []*testFilterNode{{name: "github.com/test/repo/pkg/prefix", outboundEdges: []string{"github.com/test/repo/pkg/prefix/one"}}, {name: "github.com/test/repo/pkg/prefix/one", outboundEdges: []string{"github.com/test/repo/vendor/github.com/testvendor/prefix"}}, {name: "github.com/test/repo/pkg/prefix2", outboundEdges: []string{}}}
}
func buildTestGraph(nodes []*testFilterNode) (*MutableDirectedGraph, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g := NewMutableDirectedGraph(nil)
	for _, n := range nodes {
		err := g.AddNode(&Node{UniqueName: n.name, Id: g.NewNodeID()})
		if err != nil {
			return nil, err
		}
	}
	for _, n := range nodes {
		from, exists := g.NodeByName(n.name)
		if !exists {
			continue
		}
		for _, dep := range n.outboundEdges {
			to, exists := g.NodeByName(dep)
			if !exists {
				continue
			}
			g.SetEdge(simple.Edge{F: from, T: to})
		}
	}
	return g, nil
}
func TestVendorPackagesCollapsedIntoRepo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g, err := buildTestGraph(getVendorNodes())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedRepoNodeCount := 0
	for _, n := range g.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatalf("expected node to be of type *Node")
		}
		if strings.Contains(node.UniqueName, "/vendor/") {
			continue
		}
		expectedRepoNodeCount++
	}
	vendorRoots := []string{"github.com/test/repo/vendor/github.com/testvendor/prefix", "github.com/test/repo/vendor/github.com/testvendor/prefix2", "github.com/test/repo/vendor/github.com/google/glog", "github.com/test/repo/vendor/github.com/docker/docker-test-util"}
	filteredGraph, err := FilterPackages(g, vendorRoots)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actualRepoNodeCount := 0
	actualVendorNodeCount := 0
	for _, n := range filteredGraph.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatalf("expected node to be of type *Node")
		}
		if strings.Contains(node.UniqueName, "/vendor/") {
			actualVendorNodeCount++
			continue
		}
		actualRepoNodeCount++
	}
	if actualVendorNodeCount != len(vendorRoots) {
		t.Fatalf("expected filtered graph to have been reduced to %v vendor nodes, but saw %v", len(vendorRoots), actualVendorNodeCount)
	}
	if expectedRepoNodeCount != actualRepoNodeCount {
		t.Fatalf("expected reduced graph to have original amount of non-vendor nodes (%v), but saw %v", expectedRepoNodeCount, actualRepoNodeCount)
	}
	for _, n := range filteredGraph.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatal("expected node to be of type *Node")
		}
		seen := false
		for _, root := range vendorRoots {
			if node.UniqueName == root {
				seen = true
				break
			}
		}
		if !seen {
			t.Fatalf("expected node with name %q to exist in the known vendor roots set %v", node.UniqueName, vendorRoots)
		}
	}
	expectedOutgoingEdges := map[string][]string{"github.com/test/repo/vendor/github.com/docker/docker-test-util": {"github.com/test/repo/vendor/github.com/google/glog"}, "github.com/test/repo/vendor/github.com/testvendor/prefix": {"github.com/test/repo/vendor/github.com/testvendor/prefix2"}}
	for _, n := range filteredGraph.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			continue
		}
		expectedNodes, exists := expectedOutgoingEdges[node.UniqueName]
		if !exists {
			continue
		}
		actualNodes := filteredGraph.From(n)
		if len(expectedNodes) != len(actualNodes) {
			t.Fatalf("expected node with name %q to have %v outward edges, but saw %v\n", node.UniqueName, len(expectedNodes), len(actualNodes))
		}
		for idx := range expectedNodes {
			actual, ok := actualNodes[idx].(*Node)
			if !ok {
				t.Fatal("expected node to be of type *Node")
			}
			if expectedNodes[idx] != actual.UniqueName {
				t.Fatalf("expected outgoing edge for node with name %q to point towards node with name %q, saw instead a node with name %q", node.UniqueName, expectedNodes[idx], actual.UniqueName)
			}
		}
	}
}
func TestCollapsedGraphPreservesNonVendorNodes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g, err := buildTestGraph(append(getVendorNodes(), getNonVendorNodes()...))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedRepoNodeCount := 0
	for _, n := range g.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatalf("expected node to be of type *Node")
		}
		if strings.Contains(node.UniqueName, "/vendor/") {
			continue
		}
		expectedRepoNodeCount++
	}
	vendorRoots := []string{"github.com/test/repo/vendor/github.com/testvendor/prefix", "github.com/test/repo/vendor/github.com/testvendor/prefix2", "github.com/test/repo/vendor/github.com/google/glog", "github.com/test/repo/vendor/github.com/docker/docker-test-util"}
	filteredGraph, err := FilterPackages(g, vendorRoots)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actualRepoNodeCount := 0
	for _, n := range g.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatalf("expected node to be of type *Node")
		}
		if strings.Contains(node.UniqueName, "/vendor/") {
			continue
		}
		actualRepoNodeCount++
	}
	if expectedRepoNodeCount != actualRepoNodeCount {
		t.Fatalf("expected reduced graph to contain %v nodes, but saw %v", expectedRepoNodeCount, actualRepoNodeCount)
	}
	expectedOutgoingEdges := map[string][]string{"github.com/test/repo/pkg/prefix": {"github.com/test/repo/pkg/prefix/one"}, "github.com/test/repo/pkg/prefix/one": {"github.com/test/repo/vendor/github.com/testvendor/prefix"}}
	for _, n := range filteredGraph.Nodes() {
		node, ok := n.(*Node)
		if !ok {
			t.Fatalf("expected node to be of type *Node")
		}
		expectedEdges, exists := expectedOutgoingEdges[node.UniqueName]
		if !exists {
			continue
		}
		actualEdges := filteredGraph.From(n)
		if len(expectedEdges) != len(actualEdges) {
			t.Fatalf("expeced node with name %q to contain %v outgoing edges, but saw %v", node.UniqueName, len(expectedEdges), len(actualEdges))
		}
		for _, expected := range expectedEdges {
			seen := false
			for _, n := range actualEdges {
				actual, ok := n.(*Node)
				if !ok {
					t.Fatalf("expected node to be of type *Node")
				}
				if expected == actual.UniqueName {
					seen = true
				}
			}
			if !seen {
				t.Fatalf("expected edge from %q to %q to exist in reduced graph", node.UniqueName, expected)
			}
		}
	}
}
