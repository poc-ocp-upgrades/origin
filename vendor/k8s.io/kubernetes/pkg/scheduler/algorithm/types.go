package algorithm

import (
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

var NodeFieldSelectorKeys = map[string]func(*v1.Node) string{schedulerapi.NodeFieldSelectorKeyNodeName: func(n *v1.Node) string {
	return n.Name
}}

type FitPredicate func(pod *v1.Pod, meta PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []PredicateFailureReason, error)
type PriorityMapFunction func(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error)
type PriorityReduceFunction func(pod *v1.Pod, meta interface{}, nodeNameToInfo map[string]*schedulercache.NodeInfo, result schedulerapi.HostPriorityList) error
type PredicateMetadataProducer func(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo) PredicateMetadata
type PriorityMetadataProducer func(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo) interface{}
type PriorityFunction func(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo, nodes []*v1.Node) (schedulerapi.HostPriorityList, error)
type PriorityConfig struct {
	Name     string
	Map      PriorityMapFunction
	Reduce   PriorityReduceFunction
	Function PriorityFunction
	Weight   int
}

func EmptyPredicateMetadataProducer(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo) PredicateMetadata {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func EmptyPriorityMetadataProducer(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo) interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}

type PredicateFailureReason interface{ GetReason() string }
type NodeLister interface{ List() ([]*v1.Node, error) }
type PodFilter func(*v1.Pod) bool
type PodLister interface {
	List(labels.Selector) ([]*v1.Pod, error)
	FilteredList(podFilter PodFilter, selector labels.Selector) ([]*v1.Pod, error)
}
type ServiceLister interface {
	List(labels.Selector) ([]*v1.Service, error)
	GetPodServices(*v1.Pod) ([]*v1.Service, error)
}
type ControllerLister interface {
	List(labels.Selector) ([]*v1.ReplicationController, error)
	GetPodControllers(*v1.Pod) ([]*v1.ReplicationController, error)
}
type ReplicaSetLister interface {
	GetPodReplicaSets(*v1.Pod) ([]*apps.ReplicaSet, error)
}
type PDBLister interface {
	List(labels.Selector) ([]*policyv1beta1.PodDisruptionBudget, error)
}

var _ ControllerLister = &EmptyControllerLister{}

type EmptyControllerLister struct{}

func (f EmptyControllerLister) List(labels.Selector) ([]*v1.ReplicationController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (f EmptyControllerLister) GetPodControllers(pod *v1.Pod) (controllers []*v1.ReplicationController, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}

var _ ReplicaSetLister = &EmptyReplicaSetLister{}

type EmptyReplicaSetLister struct{}

func (f EmptyReplicaSetLister) GetPodReplicaSets(pod *v1.Pod) (rss []*apps.ReplicaSet, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}

type StatefulSetLister interface {
	GetPodStatefulSets(*v1.Pod) ([]*apps.StatefulSet, error)
}

var _ StatefulSetLister = &EmptyStatefulSetLister{}

type EmptyStatefulSetLister struct{}

func (f EmptyStatefulSetLister) GetPodStatefulSets(pod *v1.Pod) (sss []*apps.StatefulSet, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}

type PredicateMetadata interface {
	ShallowCopy() PredicateMetadata
	AddPod(addedPod *v1.Pod, nodeInfo *schedulercache.NodeInfo) error
	RemovePod(deletedPod *v1.Pod) error
}
