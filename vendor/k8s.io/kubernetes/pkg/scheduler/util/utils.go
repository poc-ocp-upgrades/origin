package util

import (
	"k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/apis/scheduling"
	"k8s.io/kubernetes/pkg/features"
	"sort"
)

func GetContainerPorts(pods ...*v1.Pod) []*v1.ContainerPort {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ports []*v1.ContainerPort
	for _, pod := range pods {
		for j := range pod.Spec.Containers {
			container := &pod.Spec.Containers[j]
			for k := range container.Ports {
				ports = append(ports, &container.Ports[k])
			}
		}
	}
	return ports
}
func PodPriorityEnabled() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return feature.DefaultFeatureGate.Enabled(features.PodPriority)
}
func GetPodFullName(pod *v1.Pod) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pod.Name + "_" + pod.Namespace
}
func GetPodPriority(pod *v1.Pod) int32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod.Spec.Priority != nil {
		return *pod.Spec.Priority
	}
	return scheduling.DefaultPriorityWhenNoDefaultClassExists
}

type SortableList struct {
	Items    []interface{}
	CompFunc LessFunc
}
type LessFunc func(item1, item2 interface{}) bool

var _ = sort.Interface(&SortableList{})

func (l *SortableList) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(l.Items)
}
func (l *SortableList) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.CompFunc(l.Items[i], l.Items[j])
}
func (l *SortableList) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.Items[i], l.Items[j] = l.Items[j], l.Items[i]
}
func (l *SortableList) Sort() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Sort(l)
}
func HigherPriorityPod(pod1, pod2 interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return GetPodPriority(pod1.(*v1.Pod)) > GetPodPriority(pod2.(*v1.Pod))
}
