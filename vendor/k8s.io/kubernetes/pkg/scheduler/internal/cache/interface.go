package cache

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

type Cache interface {
	AssumePod(pod *v1.Pod) error
	FinishBinding(pod *v1.Pod) error
	ForgetPod(pod *v1.Pod) error
	AddPod(pod *v1.Pod) error
	UpdatePod(oldPod, newPod *v1.Pod) error
	RemovePod(pod *v1.Pod) error
	GetPod(pod *v1.Pod) (*v1.Pod, error)
	IsAssumedPod(pod *v1.Pod) (bool, error)
	AddNode(node *v1.Node) error
	UpdateNode(oldNode, newNode *v1.Node) error
	RemoveNode(node *v1.Node) error
	UpdateNodeNameToInfoMap(infoMap map[string]*schedulercache.NodeInfo) error
	List(labels.Selector) ([]*v1.Pod, error)
	FilteredList(filter algorithm.PodFilter, selector labels.Selector) ([]*v1.Pod, error)
	Snapshot() *Snapshot
	NodeTree() *NodeTree
}
type Snapshot struct {
	AssumedPods map[string]bool
	Nodes       map[string]*schedulercache.NodeInfo
}
