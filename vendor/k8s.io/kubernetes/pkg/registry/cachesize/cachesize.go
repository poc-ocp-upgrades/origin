package cachesize

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewHeuristicWatchCacheSizes(expectedRAMCapacityMB int) map[schema.GroupResource]int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a > b {
		return a
	}
	return b
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
