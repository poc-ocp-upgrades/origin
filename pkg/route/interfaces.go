package route

import (
	api "github.com/openshift/origin/pkg/route/apis/route"
)

type AllocationPlugin interface {
	Allocate(*api.Route) (*api.RouterShard, error)
	GenerateHostname(*api.Route, *api.RouterShard) string
}
type RouteAllocator interface {
	AllocateRouterShard(*api.Route) (*api.RouterShard, error)
	GenerateHostname(*api.Route, *api.RouterShard) string
}
