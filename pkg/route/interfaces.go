package route

import (
	goformat "fmt"
	api "github.com/openshift/origin/pkg/route/apis/route"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type AllocationPlugin interface {
	Allocate(*api.Route) (*api.RouterShard, error)
	GenerateHostname(*api.Route, *api.RouterShard) string
}
type RouteAllocator interface {
	AllocateRouterShard(*api.Route) (*api.RouterShard, error)
	GenerateHostname(*api.Route, *api.RouterShard) string
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
