package predicates

import (
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/labels"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

func FindLabelsInSet(labelsToKeep []string, selector labels.Set) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 aL := make(map[string]string)
 for _, l := range labelsToKeep {
  if selector.Has(l) {
   aL[l] = selector.Get(l)
  }
 }
 return aL
}
func AddUnsetLabelsToMap(aL map[string]string, labelsToAdd []string, labelSet labels.Set) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, l := range labelsToAdd {
  if _, exists := aL[l]; exists {
   continue
  }
  if labelSet.Has(l) {
   aL[l] = labelSet.Get(l)
  }
 }
}
func FilterPodsByNamespace(pods []*v1.Pod, ns string) []*v1.Pod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 filtered := []*v1.Pod{}
 for _, nsPod := range pods {
  if nsPod.Namespace == ns {
   filtered = append(filtered, nsPod)
  }
 }
 return filtered
}
func CreateSelectorFromLabels(aL map[string]string) labels.Selector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if aL == nil || len(aL) == 0 {
  return labels.Everything()
 }
 return labels.Set(aL).AsSelector()
}
func portsConflict(existingPorts schedulercache.HostPortInfo, wantPorts []*v1.ContainerPort) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, cp := range wantPorts {
  if existingPorts.CheckConflict(cp.HostIP, string(cp.Protocol), cp.HostPort) {
   return true
  }
 }
 return false
}
