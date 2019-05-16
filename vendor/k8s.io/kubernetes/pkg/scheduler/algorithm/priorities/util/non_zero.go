package util

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const DefaultMilliCPURequest int64 = 100
const DefaultMemoryRequest int64 = 200 * 1024 * 1024

func GetNonzeroRequests(requests *v1.ResourceList) (int64, int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var outMilliCPU, outMemory int64
	if _, found := (*requests)[v1.ResourceCPU]; !found {
		outMilliCPU = DefaultMilliCPURequest
	} else {
		outMilliCPU = requests.Cpu().MilliValue()
	}
	if _, found := (*requests)[v1.ResourceMemory]; !found {
		outMemory = DefaultMemoryRequest
	} else {
		outMemory = requests.Memory().Value()
	}
	return outMilliCPU, outMemory
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
