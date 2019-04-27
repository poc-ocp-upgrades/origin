package graphview

import (
	"fmt"
	"sort"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubeedges "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
	routeedges "github.com/openshift/origin/pkg/oc/lib/graph/routegraph"
	routegraph "github.com/openshift/origin/pkg/oc/lib/graph/routegraph/nodes"
)

type ServiceGroup struct {
	Service				*kubegraph.ServiceNode
	DeploymentConfigPipelines	[]DeploymentConfigPipeline
	ReplicationControllers		[]ReplicationController
	ReplicaSets			[]ReplicaSet
	DaemonSets			[]DaemonSet
	Deployments			[]Deployment
	StatefulSets			[]StatefulSet
	FulfillingStatefulSets		[]*kubegraph.StatefulSetNode
	FulfillingDeployments		[]*kubegraph.DeploymentNode
	FulfillingDCs			[]*appsgraph.DeploymentConfigNode
	FulfillingRCs			[]*kubegraph.ReplicationControllerNode
	FulfillingRSs			[]*kubegraph.ReplicaSetNode
	FulfillingPods			[]*kubegraph.PodNode
	FulfillingDSs			[]*kubegraph.DaemonSetNode
	ExposingRoutes			[]*routegraph.RouteNode
}

func AllServiceGroups(g osgraph.Graph, excludeNodeIDs IntSet) ([]ServiceGroup, IntSet) {
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
	covered := IntSet{}
	services := []ServiceGroup{}
	for _, uncastNode := range g.NodesByKind(kubegraph.ServiceNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		service, covers := NewServiceGroup(g, uncastNode.(*kubegraph.ServiceNode))
		covered.Insert(covers.List()...)
		services = append(services, service)
	}
	sort.Sort(ServiceGroupByObjectMeta(services))
	return services, covered
}
func NewServiceGroup(g osgraph.Graph, serviceNode *kubegraph.ServiceNode) (ServiceGroup, IntSet) {
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
	covered := IntSet{}
	covered.Insert(serviceNode.ID())
	service := ServiceGroup{}
	service.Service = serviceNode
	for _, uncastServiceFulfiller := range g.PredecessorNodesByEdgeKind(serviceNode, kubeedges.ExposedThroughServiceEdgeKind) {
		container := osgraph.GetTopLevelContainerNode(g, uncastServiceFulfiller)
		switch castContainer := container.(type) {
		case *appsgraph.DeploymentConfigNode:
			service.FulfillingDCs = append(service.FulfillingDCs, castContainer)
		case *kubegraph.ReplicationControllerNode:
			service.FulfillingRCs = append(service.FulfillingRCs, castContainer)
		case *kubegraph.ReplicaSetNode:
			service.FulfillingRSs = append(service.FulfillingRSs, castContainer)
		case *kubegraph.PodNode:
			service.FulfillingPods = append(service.FulfillingPods, castContainer)
		case *kubegraph.StatefulSetNode:
			service.FulfillingStatefulSets = append(service.FulfillingStatefulSets, castContainer)
		case *kubegraph.DeploymentNode:
			service.FulfillingDeployments = append(service.FulfillingDeployments, castContainer)
		case *kubegraph.DaemonSetNode:
			service.FulfillingDSs = append(service.FulfillingDSs, castContainer)
		default:
			utilruntime.HandleError(fmt.Errorf("unrecognized container: %v (%T)", castContainer, castContainer))
		}
	}
	for _, uncastServiceFulfiller := range g.PredecessorNodesByEdgeKind(serviceNode, routeedges.ExposedThroughRouteEdgeKind) {
		container := osgraph.GetTopLevelContainerNode(g, uncastServiceFulfiller)
		switch castContainer := container.(type) {
		case *routegraph.RouteNode:
			service.ExposingRoutes = append(service.ExposingRoutes, castContainer)
		default:
			utilruntime.HandleError(fmt.Errorf("unrecognized container: %v", castContainer))
		}
	}
	for _, fulfillingDC := range service.FulfillingDCs {
		dcPipeline, dcCovers := NewDeploymentConfigPipeline(g, fulfillingDC)
		covered.Insert(dcCovers.List()...)
		service.DeploymentConfigPipelines = append(service.DeploymentConfigPipelines, dcPipeline)
	}
	for _, fulfillingRC := range service.FulfillingRCs {
		rcView, rcCovers := NewReplicationController(g, fulfillingRC)
		covered.Insert(rcCovers.List()...)
		service.ReplicationControllers = append(service.ReplicationControllers, rcView)
	}
	for _, fulfillingRS := range service.FulfillingRSs {
		rsView, rsCovers := NewReplicaSet(g, fulfillingRS)
		covered.Insert(rsCovers.List()...)
		service.ReplicaSets = append(service.ReplicaSets, rsView)
	}
	for _, fulfillingDS := range service.FulfillingDSs {
		dsView, dsCovers := NewDaemonSet(g, fulfillingDS)
		covered.Insert(dsCovers.List()...)
		service.DaemonSets = append(service.DaemonSets, dsView)
	}
	for _, fulfillingStatefulSet := range service.FulfillingStatefulSets {
		view, covers := NewStatefulSet(g, fulfillingStatefulSet)
		covered.Insert(covers.List()...)
		service.StatefulSets = append(service.StatefulSets, view)
	}
	for _, fulfillingDeployment := range service.FulfillingDeployments {
		view, covers := NewDeployment(g, fulfillingDeployment)
		covered.Insert(covers.List()...)
		service.Deployments = append(service.Deployments, view)
	}
	for _, fulfillingPod := range service.FulfillingPods {
		_, podCovers := NewPod(g, fulfillingPod)
		covered.Insert(podCovers.List()...)
	}
	return service, covered
}

type ServiceGroupByObjectMeta []ServiceGroup

func (m ServiceGroupByObjectMeta) Len() int {
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
	return len(m)
}
func (m ServiceGroupByObjectMeta) Swap(i, j int) {
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
	m[i], m[j] = m[j], m[i]
}
func (m ServiceGroupByObjectMeta) Less(i, j int) bool {
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
	a, b := m[i], m[j]
	return CompareObjectMeta(&a.Service.Service.ObjectMeta, &b.Service.Service.ObjectMeta)
}
func CompareObjectMeta(a, b *metav1.ObjectMeta) bool {
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
	if a.Namespace == b.Namespace {
		return a.Name < b.Name
	}
	return a.Namespace < b.Namespace
}
