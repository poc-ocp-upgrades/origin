package graphview

import (
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

type Pod struct{ Pod *kubegraph.PodNode }

func AllPods(g osgraph.Graph, excludeNodeIDs IntSet) ([]Pod, IntSet) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	pods := []Pod{}
	for _, uncastNode := range g.NodesByKind(kubegraph.PodNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		pod, covers := NewPod(g, uncastNode.(*kubegraph.PodNode))
		covered.Insert(covers.List()...)
		pods = append(pods, pod)
	}
	return pods, covered
}
func NewPod(g osgraph.Graph, podNode *kubegraph.PodNode) (Pod, IntSet) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(podNode.ID())
	podView := Pod{}
	podView.Pod = podNode
	return podView, covered
}
