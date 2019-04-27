package analysis

import (
	"testing"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	osgraphtest "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph/test"
	routeedges "github.com/openshift/origin/pkg/oc/lib/graph/routegraph"
)

func TestPortMappingIssues(t *testing.T) {
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
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/missing-route-port.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers := FindPortMappingIssues(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := MissingRoutePortWarning, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
	g, _, err = osgraphtest.BuildGraph("../../../graph/genericgraph/test/lonely-route.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers = FindPortMappingIssues(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := MissingServiceWarning, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
	g, _, err = osgraphtest.BuildGraph("../../../graph/genericgraph/test/wrong-numeric-port.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers = FindPortMappingIssues(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := WrongRoutePortWarning, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
	g, _, err = osgraphtest.BuildGraph("../../../graph/genericgraph/test/wrong-named-port.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers = FindPortMappingIssues(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := WrongRoutePortWarning, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
}
func TestPathBasedPassthroughRoutes(t *testing.T) {
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
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/invalid-route.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers := FindPathBasedPassthroughRoutes(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := PathBasedPassthroughErr, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
}
func TestMissingRouter(t *testing.T) {
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
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/lonely-route.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	routeedges.AddAllRouteEdges(g)
	markers := FindMissingRouter(g, osgraph.DefaultNamer)
	if expected, got := 1, len(markers); expected != got {
		t.Fatalf("expected %d markers, got %d", expected, got)
	}
	if expected, got := MissingRequiredRouterErr, markers[0].Key; expected != got {
		t.Fatalf("expected %s marker key, got %s", expected, got)
	}
}
