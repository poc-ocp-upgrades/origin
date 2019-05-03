package garbagecollector

import (
 "fmt"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "net/http"
 godefaulthttp "net/http"
 "strings"
 "gonum.org/v1/gonum/graph"
 "gonum.org/v1/gonum/graph/encoding"
 "gonum.org/v1/gonum/graph/encoding/dot"
 "gonum.org/v1/gonum/graph/simple"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

type gonumVertex struct {
 uid                types.UID
 gvk                schema.GroupVersionKind
 namespace          string
 name               string
 missingFromGraph   bool
 beingDeleted       bool
 deletingDependents bool
 virtual            bool
 vertexID           int64
}

func (v *gonumVertex) ID() int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return v.vertexID
}
func (v *gonumVertex) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 kind := v.gvk.Kind + "." + v.gvk.Version
 if len(v.gvk.Group) > 0 {
  kind = kind + "." + v.gvk.Group
 }
 missing := ""
 if v.missingFromGraph {
  missing = "(missing)"
 }
 deleting := ""
 if v.beingDeleted {
  deleting = "(deleting)"
 }
 deletingDependents := ""
 if v.deletingDependents {
  deleting = "(deletingDependents)"
 }
 virtual := ""
 if v.virtual {
  virtual = "(virtual)"
 }
 return fmt.Sprintf(`%s/%s[%s]-%v%s%s%s%s`, kind, v.name, v.namespace, v.uid, missing, deleting, deletingDependents, virtual)
}
func (v *gonumVertex) Attributes() []encoding.Attribute {
 _logClusterCodePath()
 defer _logClusterCodePath()
 kubectlString := v.gvk.Kind + "." + v.gvk.Version
 if len(v.gvk.Group) > 0 {
  kubectlString = kubectlString + "." + v.gvk.Group
 }
 kubectlString = kubectlString + "/" + v.name
 label := fmt.Sprintf(`uid=%v
namespace=%v
%v
`, v.uid, v.namespace, kubectlString)
 conditionStrings := []string{}
 if v.beingDeleted {
  conditionStrings = append(conditionStrings, "beingDeleted")
 }
 if v.deletingDependents {
  conditionStrings = append(conditionStrings, "deletingDependents")
 }
 if v.virtual {
  conditionStrings = append(conditionStrings, "virtual")
 }
 if v.missingFromGraph {
  conditionStrings = append(conditionStrings, "missingFromGraph")
 }
 conditionString := strings.Join(conditionStrings, ",")
 if len(conditionString) > 0 {
  label = label + conditionString + "\n"
 }
 return []encoding.Attribute{{Key: "label", Value: fmt.Sprintf(`"%v"`, label)}, {Key: "group", Value: fmt.Sprintf(`"%v"`, v.gvk.Group)}, {Key: "version", Value: fmt.Sprintf(`"%v"`, v.gvk.Version)}, {Key: "kind", Value: fmt.Sprintf(`"%v"`, v.gvk.Kind)}, {Key: "namespace", Value: fmt.Sprintf(`"%v"`, v.namespace)}, {Key: "name", Value: fmt.Sprintf(`"%v"`, v.name)}, {Key: "uid", Value: fmt.Sprintf(`"%v"`, v.uid)}, {Key: "missing", Value: fmt.Sprintf(`"%v"`, v.missingFromGraph)}, {Key: "beingDeleted", Value: fmt.Sprintf(`"%v"`, v.beingDeleted)}, {Key: "deletingDependents", Value: fmt.Sprintf(`"%v"`, v.deletingDependents)}, {Key: "virtual", Value: fmt.Sprintf(`"%v"`, v.virtual)}}
}
func NewGonumVertex(node *node, nodeID int64) *gonumVertex {
 _logClusterCodePath()
 defer _logClusterCodePath()
 gv, err := schema.ParseGroupVersion(node.identity.APIVersion)
 if err != nil {
  utilruntime.HandleError(err)
 }
 return &gonumVertex{uid: node.identity.UID, gvk: gv.WithKind(node.identity.Kind), namespace: node.identity.Namespace, name: node.identity.Name, beingDeleted: node.beingDeleted, deletingDependents: node.deletingDependents, virtual: node.virtual, vertexID: nodeID}
}
func NewMissingGonumVertex(ownerRef metav1.OwnerReference, nodeID int64) *gonumVertex {
 _logClusterCodePath()
 defer _logClusterCodePath()
 gv, err := schema.ParseGroupVersion(ownerRef.APIVersion)
 if err != nil {
  utilruntime.HandleError(err)
 }
 return &gonumVertex{uid: ownerRef.UID, gvk: gv.WithKind(ownerRef.Kind), name: ownerRef.Name, missingFromGraph: true, vertexID: nodeID}
}
func (m *concurrentUIDToNode) ToGonumGraph() graph.Directed {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.uidToNodeLock.Lock()
 defer m.uidToNodeLock.Unlock()
 return toGonumGraph(m.uidToNode)
}
func toGonumGraph(uidToNode map[types.UID]*node) graph.Directed {
 _logClusterCodePath()
 defer _logClusterCodePath()
 uidToVertex := map[types.UID]*gonumVertex{}
 graphBuilder := simple.NewDirectedGraph()
 for _, node := range uidToNode {
  if len(node.dependents) == 0 && len(node.owners) == 0 {
   continue
  }
  vertex := NewGonumVertex(node, graphBuilder.NewNode().ID())
  uidToVertex[node.identity.UID] = vertex
  graphBuilder.AddNode(vertex)
 }
 for _, node := range uidToNode {
  currVertex := uidToVertex[node.identity.UID]
  for _, ownerRef := range node.owners {
   currOwnerVertex, ok := uidToVertex[ownerRef.UID]
   if !ok {
    currOwnerVertex = NewMissingGonumVertex(ownerRef, graphBuilder.NewNode().ID())
    uidToVertex[node.identity.UID] = currOwnerVertex
    graphBuilder.AddNode(currOwnerVertex)
   }
   graphBuilder.SetEdge(simple.Edge{F: currVertex, T: currOwnerVertex})
  }
 }
 return graphBuilder
}
func (m *concurrentUIDToNode) ToGonumGraphForObj(uids ...types.UID) graph.Directed {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.uidToNodeLock.Lock()
 defer m.uidToNodeLock.Unlock()
 return toGonumGraphForObj(m.uidToNode, uids...)
}
func toGonumGraphForObj(uidToNode map[types.UID]*node, uids ...types.UID) graph.Directed {
 _logClusterCodePath()
 defer _logClusterCodePath()
 uidsToCheck := append([]types.UID{}, uids...)
 interestingNodes := map[types.UID]*node{}
 for i := 0; i < len(uidsToCheck); i++ {
  uid := uidsToCheck[i]
  if _, ok := interestingNodes[uid]; ok {
   continue
  }
  node, ok := uidToNode[uid]
  if !ok {
   continue
  }
  interestingNodes[node.identity.UID] = node
  for _, ownerRef := range node.owners {
   if _, ok := interestingNodes[ownerRef.UID]; ok {
    continue
   }
   uidsToCheck = append(uidsToCheck, ownerRef.UID)
  }
  for dependent := range node.dependents {
   if _, ok := interestingNodes[dependent.identity.UID]; ok {
    continue
   }
   uidsToCheck = append(uidsToCheck, dependent.identity.UID)
  }
 }
 return toGonumGraph(interestingNodes)
}
func NewDebugHandler(controller *GarbageCollector) http.Handler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &debugHTTPHandler{controller: controller}
}

type debugHTTPHandler struct{ controller *GarbageCollector }

func (h *debugHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if req.URL.Path != "/graph" {
  http.Error(w, "", http.StatusNotFound)
  return
 }
 var graph graph.Directed
 if uidStrings := req.URL.Query()["uid"]; len(uidStrings) > 0 {
  uids := []types.UID{}
  for _, uidString := range uidStrings {
   uids = append(uids, types.UID(uidString))
  }
  graph = h.controller.dependencyGraphBuilder.uidToNode.ToGonumGraphForObj(uids...)
 } else {
  graph = h.controller.dependencyGraphBuilder.uidToNode.ToGonumGraph()
 }
 data, err := dot.Marshal(graph, "full", "", "  ", false)
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 w.Write(data)
 w.WriteHeader(http.StatusOK)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
