package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
)

func (az *Cloud) ListRoutes(ctx context.Context, clusterName string) ([]*cloudprovider.Route, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(10).Infof("ListRoutes: START clusterName=%q", clusterName)
	routeTable, existsRouteTable, err := az.getRouteTable()
	routes, err := processRoutes(routeTable, existsRouteTable, err)
	if err != nil {
		return nil, err
	}
	unmanagedNodes, err := az.GetUnmanagedNodes()
	if err != nil {
		return nil, err
	}
	az.routeCIDRsLock.Lock()
	defer az.routeCIDRsLock.Unlock()
	for _, nodeName := range unmanagedNodes.List() {
		if cidr, ok := az.routeCIDRs[nodeName]; ok {
			routes = append(routes, &cloudprovider.Route{Name: nodeName, TargetNode: mapRouteNameToNodeName(nodeName), DestinationCIDR: cidr})
		}
	}
	return routes, nil
}
func processRoutes(routeTable network.RouteTable, exists bool, err error) ([]*cloudprovider.Route, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err != nil {
		return nil, err
	}
	if !exists {
		return []*cloudprovider.Route{}, nil
	}
	var kubeRoutes []*cloudprovider.Route
	if routeTable.RouteTablePropertiesFormat != nil && routeTable.Routes != nil {
		kubeRoutes = make([]*cloudprovider.Route, len(*routeTable.Routes))
		for i, route := range *routeTable.Routes {
			instance := mapRouteNameToNodeName(*route.Name)
			cidr := *route.AddressPrefix
			klog.V(10).Infof("ListRoutes: * instance=%q, cidr=%q", instance, cidr)
			kubeRoutes[i] = &cloudprovider.Route{Name: *route.Name, TargetNode: instance, DestinationCIDR: cidr}
		}
	}
	klog.V(10).Info("ListRoutes: FINISH")
	return kubeRoutes, nil
}
func (az *Cloud) createRouteTableIfNotExists(clusterName string, kubeRoute *cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, existsRouteTable, err := az.getRouteTable(); err != nil {
		klog.V(2).Infof("createRouteTableIfNotExists error: couldn't get routetable. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
		return err
	} else if existsRouteTable {
		return nil
	}
	return az.createRouteTable()
}
func (az *Cloud) createRouteTable() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	routeTable := network.RouteTable{Name: to.StringPtr(az.RouteTableName), Location: to.StringPtr(az.Location), RouteTablePropertiesFormat: &network.RouteTablePropertiesFormat{}}
	klog.V(3).Infof("createRouteTableIfNotExists: creating routetable. routeTableName=%q", az.RouteTableName)
	ctx, cancel := getContextWithCancel()
	defer cancel()
	resp, err := az.RouteTablesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, routeTable)
	klog.V(10).Infof("RouteTablesClient.CreateOrUpdate(%q): end", az.RouteTableName)
	if az.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
		klog.V(2).Infof("createRouteTableIfNotExists backing off: creating routetable. routeTableName=%q", az.RouteTableName)
		retryErr := az.CreateOrUpdateRouteTableWithRetry(routeTable)
		if retryErr != nil {
			err = retryErr
			klog.V(2).Infof("createRouteTableIfNotExists abort backoff: creating routetable. routeTableName=%q", az.RouteTableName)
		}
	}
	if err != nil {
		return err
	}
	az.rtCache.Delete(az.RouteTableName)
	return nil
}
func (az *Cloud) CreateRoute(ctx context.Context, clusterName string, nameHint string, kubeRoute *cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeName := string(kubeRoute.TargetNode)
	unmanaged, err := az.IsNodeUnmanaged(nodeName)
	if err != nil {
		return err
	}
	if unmanaged {
		klog.V(2).Infof("CreateRoute: omitting unmanaged node %q", kubeRoute.TargetNode)
		az.routeCIDRsLock.Lock()
		defer az.routeCIDRsLock.Unlock()
		az.routeCIDRs[nodeName] = kubeRoute.DestinationCIDR
		return nil
	}
	klog.V(2).Infof("CreateRoute: creating route. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
	if err := az.createRouteTableIfNotExists(clusterName, kubeRoute); err != nil {
		return err
	}
	targetIP, _, err := az.getIPForMachine(kubeRoute.TargetNode)
	if err != nil {
		return err
	}
	routeName := mapNodeNameToRouteName(kubeRoute.TargetNode)
	route := network.Route{Name: to.StringPtr(routeName), RoutePropertiesFormat: &network.RoutePropertiesFormat{AddressPrefix: to.StringPtr(kubeRoute.DestinationCIDR), NextHopType: network.RouteNextHopTypeVirtualAppliance, NextHopIPAddress: to.StringPtr(targetIP)}}
	klog.V(3).Infof("CreateRoute: creating route: instance=%q cidr=%q", kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
	ctx, cancel := getContextWithCancel()
	defer cancel()
	resp, err := az.RoutesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, *route.Name, route)
	klog.V(10).Infof("RoutesClient.CreateOrUpdate(%q): end", az.RouteTableName)
	if az.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
		klog.V(2).Infof("CreateRoute backing off: creating route: instance=%q cidr=%q", kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
		retryErr := az.CreateOrUpdateRouteWithRetry(route)
		if retryErr != nil {
			err = retryErr
			klog.V(2).Infof("CreateRoute abort backoff: creating route: instance=%q cidr=%q", kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
		}
	}
	if err != nil {
		return err
	}
	klog.V(2).Infof("CreateRoute: route created. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
	return nil
}
func (az *Cloud) DeleteRoute(ctx context.Context, clusterName string, kubeRoute *cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeName := string(kubeRoute.TargetNode)
	unmanaged, err := az.IsNodeUnmanaged(nodeName)
	if err != nil {
		return err
	}
	if unmanaged {
		klog.V(2).Infof("DeleteRoute: omitting unmanaged node %q", kubeRoute.TargetNode)
		az.routeCIDRsLock.Lock()
		defer az.routeCIDRsLock.Unlock()
		delete(az.routeCIDRs, nodeName)
		return nil
	}
	klog.V(2).Infof("DeleteRoute: deleting route. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
	ctx, cancel := getContextWithCancel()
	defer cancel()
	routeName := mapNodeNameToRouteName(kubeRoute.TargetNode)
	resp, err := az.RoutesClient.Delete(ctx, az.ResourceGroup, az.RouteTableName, routeName)
	klog.V(10).Infof("RoutesClient.Delete(%q): end", az.RouteTableName)
	if az.CloudProviderBackoff && shouldRetryHTTPRequest(resp, err) {
		klog.V(2).Infof("DeleteRoute backing off: deleting route. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
		retryErr := az.DeleteRouteWithRetry(routeName)
		if retryErr != nil {
			err = retryErr
			klog.V(2).Infof("DeleteRoute abort backoff: deleting route. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
		}
	}
	if err != nil {
		return err
	}
	klog.V(2).Infof("DeleteRoute: route deleted. clusterName=%q instance=%q cidr=%q", clusterName, kubeRoute.TargetNode, kubeRoute.DestinationCIDR)
	return nil
}
func mapNodeNameToRouteName(nodeName types.NodeName) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s", nodeName)
}
func mapRouteNameToNodeName(routeName string) types.NodeName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return types.NodeName(fmt.Sprintf("%s", routeName))
}
