package node

import (
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	pvutil "k8s.io/kubernetes/pkg/api/v1/persistentvolume"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"k8s.io/kubernetes/third_party/forked/gonum/graph"
	"k8s.io/kubernetes/third_party/forked/gonum/graph/simple"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type namedVertex struct {
	name       string
	namespace  string
	id         int
	vertexType vertexType
}

func newNamedVertex(vertexType vertexType, namespace, name string, id int) *namedVertex {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &namedVertex{vertexType: vertexType, name: name, namespace: namespace, id: id}
}
func (n *namedVertex) ID() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return n.id
}
func (n *namedVertex) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(n.namespace) == 0 {
		return vertexTypes[n.vertexType] + ":" + n.name
	}
	return vertexTypes[n.vertexType] + ":" + n.namespace + "/" + n.name
}

type destinationEdge struct {
	F           graph.Node
	T           graph.Node
	Destination graph.Node
}

func newDestinationEdge(from, to, destination graph.Node) graph.Edge {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &destinationEdge{F: from, T: to, Destination: destination}
}
func (e *destinationEdge) From() graph.Node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.F
}
func (e *destinationEdge) To() graph.Node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.T
}
func (e *destinationEdge) Weight() float64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return 0
}
func (e *destinationEdge) DestinationID() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.Destination.ID()
}

type Graph struct {
	lock                     sync.RWMutex
	graph                    *simple.DirectedAcyclicGraph
	vertices                 map[vertexType]namespaceVertexMapping
	destinationEdgeIndex     map[int]*intSet
	destinationEdgeThreshold int
}
type namespaceVertexMapping map[string]nameVertexMapping
type nameVertexMapping map[string]*namedVertex

func NewGraph() *Graph {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Graph{vertices: map[vertexType]namespaceVertexMapping{}, graph: simple.NewDirectedAcyclicGraph(0, 0), destinationEdgeIndex: map[int]*intSet{}, destinationEdgeThreshold: 200}
}

type vertexType byte

const (
	configMapVertexType vertexType = iota
	nodeVertexType
	podVertexType
	pvcVertexType
	pvVertexType
	secretVertexType
	vaVertexType
	serviceAccountVertexType
)

var vertexTypes = map[vertexType]string{configMapVertexType: "configmap", nodeVertexType: "node", podVertexType: "pod", pvcVertexType: "pvc", pvVertexType: "pv", secretVertexType: "secret", vaVertexType: "volumeattachment", serviceAccountVertexType: "serviceAccount"}

func (g *Graph) getOrCreateVertex_locked(vertexType vertexType, namespace, name string) *namedVertex {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if vertex, exists := g.getVertex_rlocked(vertexType, namespace, name); exists {
		return vertex
	}
	return g.createVertex_locked(vertexType, namespace, name)
}
func (g *Graph) getVertex_rlocked(vertexType vertexType, namespace, name string) (*namedVertex, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vertex, exists := g.vertices[vertexType][namespace][name]
	return vertex, exists
}
func (g *Graph) createVertex_locked(vertexType vertexType, namespace, name string) *namedVertex {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	typedVertices, exists := g.vertices[vertexType]
	if !exists {
		typedVertices = namespaceVertexMapping{}
		g.vertices[vertexType] = typedVertices
	}
	namespacedVertices, exists := typedVertices[namespace]
	if !exists {
		namespacedVertices = map[string]*namedVertex{}
		typedVertices[namespace] = namespacedVertices
	}
	vertex := newNamedVertex(vertexType, namespace, name, g.graph.NewNodeID())
	namespacedVertices[name] = vertex
	g.graph.AddNode(vertex)
	return vertex
}
func (g *Graph) deleteVertex_locked(vertexType vertexType, namespace, name string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vertex, exists := g.getVertex_rlocked(vertexType, namespace, name)
	if !exists {
		return
	}
	neighborsToRemove := []graph.Node{}
	neighborsToRecompute := []graph.Node{}
	g.graph.VisitFrom(vertex, func(neighbor graph.Node) bool {
		if g.graph.Degree(neighbor) == 1 {
			neighborsToRemove = append(neighborsToRemove, neighbor)
		}
		return true
	})
	g.graph.VisitTo(vertex, func(neighbor graph.Node) bool {
		if g.graph.Degree(neighbor) == 1 {
			neighborsToRemove = append(neighborsToRemove, neighbor)
		} else {
			neighborsToRecompute = append(neighborsToRemove, neighbor)
		}
		return true
	})
	g.removeVertex_locked(vertex)
	for _, neighbor := range neighborsToRemove {
		g.removeVertex_locked(neighbor.(*namedVertex))
	}
	for _, neighbor := range neighborsToRecompute {
		g.recomputeDestinationIndex_locked(neighbor)
	}
}
func (g *Graph) deleteEdges_locked(fromType, toType vertexType, toNamespace, toName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	toVert, exists := g.getVertex_rlocked(toType, toNamespace, toName)
	if !exists {
		return
	}
	neighborsToRemove := []*namedVertex{}
	neighborsToRecompute := []*namedVertex{}
	g.graph.VisitTo(toVert, func(from graph.Node) bool {
		fromVert := from.(*namedVertex)
		if fromVert.vertexType != fromType {
			return true
		}
		g.graph.RemoveEdge(simple.Edge{F: fromVert, T: toVert})
		if g.graph.Degree(fromVert) == 0 {
			neighborsToRemove = append(neighborsToRemove, fromVert)
		} else {
			neighborsToRecompute = append(neighborsToRecompute, fromVert)
		}
		return true
	})
	for _, v := range neighborsToRemove {
		g.removeVertex_locked(v)
	}
	for _, v := range neighborsToRecompute {
		g.recomputeDestinationIndex_locked(v)
	}
}
func (g *Graph) removeVertex_locked(v *namedVertex) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.graph.RemoveNode(v)
	delete(g.destinationEdgeIndex, v.ID())
	delete(g.vertices[v.vertexType][v.namespace], v.name)
	if len(g.vertices[v.vertexType][v.namespace]) == 0 {
		delete(g.vertices[v.vertexType], v.namespace)
	}
}
func (g *Graph) recomputeDestinationIndex_locked(n graph.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	edgeCount := g.graph.Degree(n)
	if edgeCount < g.destinationEdgeThreshold {
		delete(g.destinationEdgeIndex, n.ID())
		return
	}
	index := g.destinationEdgeIndex[n.ID()]
	if index == nil {
		index = newIntSet()
	} else {
		index.startNewGeneration()
	}
	g.graph.VisitFrom(n, func(dest graph.Node) bool {
		if destinationEdge, ok := g.graph.EdgeBetween(n, dest).(*destinationEdge); ok {
			index.mark(destinationEdge.DestinationID())
		}
		return true
	})
	index.sweep()
	if len(index.members) < g.destinationEdgeThreshold {
		delete(g.destinationEdgeIndex, n.ID())
	} else {
		g.destinationEdgeIndex[n.ID()] = index
	}
}
func (g *Graph) AddPod(pod *corev1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(podVertexType, pod.Namespace, pod.Name)
	podVertex := g.getOrCreateVertex_locked(podVertexType, pod.Namespace, pod.Name)
	nodeVertex := g.getOrCreateVertex_locked(nodeVertexType, "", pod.Spec.NodeName)
	g.graph.SetEdge(newDestinationEdge(podVertex, nodeVertex, nodeVertex))
	if _, isMirrorPod := pod.Annotations[corev1.MirrorPodAnnotationKey]; isMirrorPod {
		return
	}
	if len(pod.Spec.ServiceAccountName) > 0 {
		serviceAccountVertex := g.getOrCreateVertex_locked(serviceAccountVertexType, pod.Namespace, pod.Spec.ServiceAccountName)
		g.graph.SetEdge(newDestinationEdge(serviceAccountVertex, podVertex, nodeVertex))
		g.recomputeDestinationIndex_locked(serviceAccountVertex)
	}
	podutil.VisitPodSecretNames(pod, func(secret string) bool {
		secretVertex := g.getOrCreateVertex_locked(secretVertexType, pod.Namespace, secret)
		g.graph.SetEdge(newDestinationEdge(secretVertex, podVertex, nodeVertex))
		g.recomputeDestinationIndex_locked(secretVertex)
		return true
	})
	podutil.VisitPodConfigmapNames(pod, func(configmap string) bool {
		configmapVertex := g.getOrCreateVertex_locked(configMapVertexType, pod.Namespace, configmap)
		g.graph.SetEdge(newDestinationEdge(configmapVertex, podVertex, nodeVertex))
		g.recomputeDestinationIndex_locked(configmapVertex)
		return true
	})
	for _, v := range pod.Spec.Volumes {
		if v.PersistentVolumeClaim != nil {
			pvcVertex := g.getOrCreateVertex_locked(pvcVertexType, pod.Namespace, v.PersistentVolumeClaim.ClaimName)
			g.graph.SetEdge(newDestinationEdge(pvcVertex, podVertex, nodeVertex))
			g.recomputeDestinationIndex_locked(pvcVertex)
		}
	}
}
func (g *Graph) DeletePod(name, namespace string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(podVertexType, namespace, name)
}
func (g *Graph) AddPV(pv *corev1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(pvVertexType, "", pv.Name)
	if pv.Spec.ClaimRef != nil {
		pvVertex := g.getOrCreateVertex_locked(pvVertexType, "", pv.Name)
		g.graph.SetEdge(simple.Edge{F: pvVertex, T: g.getOrCreateVertex_locked(pvcVertexType, pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name)})
		pvutil.VisitPVSecretNames(pv, func(namespace, secret string, kubeletVisible bool) bool {
			if kubeletVisible {
				g.graph.SetEdge(simple.Edge{F: g.getOrCreateVertex_locked(secretVertexType, namespace, secret), T: pvVertex})
			}
			return true
		})
	}
}
func (g *Graph) DeletePV(name string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(pvVertexType, "", name)
}
func (g *Graph) AddVolumeAttachment(attachmentName, nodeName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(vaVertexType, "", attachmentName)
	if len(nodeName) > 0 {
		vaVertex := g.getOrCreateVertex_locked(vaVertexType, "", attachmentName)
		nodeVertex := g.getOrCreateVertex_locked(nodeVertexType, "", nodeName)
		g.graph.SetEdge(newDestinationEdge(vaVertex, nodeVertex, nodeVertex))
	}
}
func (g *Graph) DeleteVolumeAttachment(name string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteVertex_locked(vaVertexType, "", name)
}
func (g *Graph) SetNodeConfigMap(nodeName, configMapName, configMapNamespace string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.lock.Lock()
	defer g.lock.Unlock()
	g.deleteEdges_locked(configMapVertexType, nodeVertexType, "", nodeName)
	if len(configMapName) > 0 && len(configMapNamespace) > 0 {
		configmapVertex := g.getOrCreateVertex_locked(configMapVertexType, configMapNamespace, configMapName)
		nodeVertex := g.getOrCreateVertex_locked(nodeVertexType, "", nodeName)
		g.graph.SetEdge(newDestinationEdge(configmapVertex, nodeVertex, nodeVertex))
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
