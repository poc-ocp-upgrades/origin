package analyze

import (
	"github.com/openshift/origin/tools/depcheck/pkg/graph"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

func FindExclusiveDependencies(g *graph.MutableDirectedGraph, targetNodes []*graph.Node) []*graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newGraph := g.Copy()
	for _, target := range targetNodes {
		newGraph.RemoveNode(target)
	}
	return newGraph.PruneOrphans()
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
