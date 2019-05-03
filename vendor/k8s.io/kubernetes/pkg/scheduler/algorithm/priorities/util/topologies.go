package util

import (
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/util/sets"
)

func GetNamespacesFromPodAffinityTerm(pod *v1.Pod, podAffinityTerm *v1.PodAffinityTerm) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 names := sets.String{}
 if len(podAffinityTerm.Namespaces) == 0 {
  names.Insert(pod.Namespace)
 } else {
  names.Insert(podAffinityTerm.Namespaces...)
 }
 return names
}
func PodMatchesTermsNamespaceAndSelector(pod *v1.Pod, namespaces sets.String, selector labels.Selector) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !namespaces.Has(pod.Namespace) {
  return false
 }
 if !selector.Matches(labels.Set(pod.Labels)) {
  return false
 }
 return true
}
func NodesHaveSameTopologyKey(nodeA, nodeB *v1.Node, topologyKey string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(topologyKey) == 0 {
  return false
 }
 if nodeA.Labels == nil || nodeB.Labels == nil {
  return false
 }
 nodeALabel, okA := nodeA.Labels[topologyKey]
 nodeBLabel, okB := nodeB.Labels[topologyKey]
 if okB && okA {
  return nodeALabel == nodeBLabel
 }
 return false
}

type Topologies struct{ DefaultKeys []string }

func (tps *Topologies) NodesHaveSameTopologyKey(nodeA, nodeB *v1.Node, topologyKey string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return NodesHaveSameTopologyKey(nodeA, nodeB, topologyKey)
}
