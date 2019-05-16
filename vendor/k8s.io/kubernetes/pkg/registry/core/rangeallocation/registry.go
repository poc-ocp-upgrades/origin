package rangeallocation

import (
	goformat "fmt"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RangeRegistry interface {
	Get() (*api.RangeAllocation, error)
	CreateOrUpdate(*api.RangeAllocation) error
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
