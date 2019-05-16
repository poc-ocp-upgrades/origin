package configprocessing

import (
	routeplugin "github.com/openshift/origin/pkg/route/allocation/simple"
	routeallocationcontroller "github.com/openshift/origin/pkg/route/controller/allocation"
)

func RouteAllocator(routingSubdomain string) (*routeallocationcontroller.RouteAllocationController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	factory := routeallocationcontroller.RouteAllocationControllerFactory{}
	plugin, err := routeplugin.NewSimpleAllocationPlugin(routingSubdomain)
	if err != nil {
		return nil, err
	}
	return factory.Create(plugin), nil
}
