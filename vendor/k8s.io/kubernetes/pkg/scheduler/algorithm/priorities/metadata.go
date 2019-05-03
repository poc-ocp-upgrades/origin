package priorities

import (
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

type PriorityMetadataFactory struct {
 serviceLister     algorithm.ServiceLister
 controllerLister  algorithm.ControllerLister
 replicaSetLister  algorithm.ReplicaSetLister
 statefulSetLister algorithm.StatefulSetLister
}

func NewPriorityMetadataFactory(serviceLister algorithm.ServiceLister, controllerLister algorithm.ControllerLister, replicaSetLister algorithm.ReplicaSetLister, statefulSetLister algorithm.StatefulSetLister) algorithm.PriorityMetadataProducer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 factory := &PriorityMetadataFactory{serviceLister: serviceLister, controllerLister: controllerLister, replicaSetLister: replicaSetLister, statefulSetLister: statefulSetLister}
 return factory.PriorityMetadata
}

type priorityMetadata struct {
 nonZeroRequest          *schedulercache.Resource
 podTolerations          []v1.Toleration
 affinity                *v1.Affinity
 podSelectors            []labels.Selector
 controllerRef           *metav1.OwnerReference
 podFirstServiceSelector labels.Selector
 totalNumNodes           int
}

func (pmf *PriorityMetadataFactory) PriorityMetadata(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo) interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pod == nil {
  return nil
 }
 return &priorityMetadata{nonZeroRequest: getNonZeroRequests(pod), podTolerations: getAllTolerationPreferNoSchedule(pod.Spec.Tolerations), affinity: pod.Spec.Affinity, podSelectors: getSelectors(pod, pmf.serviceLister, pmf.controllerLister, pmf.replicaSetLister, pmf.statefulSetLister), controllerRef: metav1.GetControllerOf(pod), podFirstServiceSelector: getFirstServiceSelector(pod, pmf.serviceLister), totalNumNodes: len(nodeNameToInfo)}
}
func getFirstServiceSelector(pod *v1.Pod, sl algorithm.ServiceLister) (firstServiceSelector labels.Selector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if services, err := sl.GetPodServices(pod); err == nil && len(services) > 0 {
  return labels.SelectorFromSet(services[0].Spec.Selector)
 }
 return nil
}
func getSelectors(pod *v1.Pod, sl algorithm.ServiceLister, cl algorithm.ControllerLister, rsl algorithm.ReplicaSetLister, ssl algorithm.StatefulSetLister) []labels.Selector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var selectors []labels.Selector
 if services, err := sl.GetPodServices(pod); err == nil {
  for _, service := range services {
   selectors = append(selectors, labels.SelectorFromSet(service.Spec.Selector))
  }
 }
 if rcs, err := cl.GetPodControllers(pod); err == nil {
  for _, rc := range rcs {
   selectors = append(selectors, labels.SelectorFromSet(rc.Spec.Selector))
  }
 }
 if rss, err := rsl.GetPodReplicaSets(pod); err == nil {
  for _, rs := range rss {
   if selector, err := metav1.LabelSelectorAsSelector(rs.Spec.Selector); err == nil {
    selectors = append(selectors, selector)
   }
  }
 }
 if sss, err := ssl.GetPodStatefulSets(pod); err == nil {
  for _, ss := range sss {
   if selector, err := metav1.LabelSelectorAsSelector(ss.Spec.Selector); err == nil {
    selectors = append(selectors, selector)
   }
  }
 }
 return selectors
}
