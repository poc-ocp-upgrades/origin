package algorithm

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type SchedulerExtender interface {
	Name() string
	Filter(pod *v1.Pod, nodes []*v1.Node, nodeNameToInfo map[string]*schedulercache.NodeInfo) (filteredNodes []*v1.Node, failedNodesMap schedulerapi.FailedNodesMap, err error)
	Prioritize(pod *v1.Pod, nodes []*v1.Node) (hostPriorities *schedulerapi.HostPriorityList, weight int, err error)
	Bind(binding *v1.Binding) error
	IsBinder() bool
	IsInterested(pod *v1.Pod) bool
	ProcessPreemption(pod *v1.Pod, nodeToVictims map[*v1.Node]*schedulerapi.Victims, nodeNameToInfo map[string]*schedulercache.NodeInfo) (map[*v1.Node]*schedulerapi.Victims, error)
	SupportsPreemption() bool
	IsIgnorable() bool
}
type ScheduleAlgorithm interface {
	Schedule(*v1.Pod, NodeLister) (selectedMachine string, err error)
	Preempt(*v1.Pod, NodeLister, error) (selectedNode *v1.Node, preemptedPods []*v1.Pod, cleanupNominatedPods []*v1.Pod, err error)
	Predicates() map[string]FitPredicate
	Prioritizers() []PriorityConfig
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
