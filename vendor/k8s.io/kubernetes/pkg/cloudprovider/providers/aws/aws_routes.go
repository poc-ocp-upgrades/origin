package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
)

func (c *Cloud) findRouteTable(clusterName string) (*ec2.RouteTable, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var tables []*ec2.RouteTable
	if c.cfg.Global.RouteTableID != "" {
		request := &ec2.DescribeRouteTablesInput{Filters: []*ec2.Filter{newEc2Filter("route-table-id", c.cfg.Global.RouteTableID)}}
		response, err := c.ec2.DescribeRouteTables(request)
		if err != nil {
			return nil, err
		}
		tables = response
	} else {
		request := &ec2.DescribeRouteTablesInput{Filters: c.tagging.addFilters(nil)}
		response, err := c.ec2.DescribeRouteTables(request)
		if err != nil {
			return nil, err
		}
		for _, table := range response {
			if c.tagging.hasClusterTag(table.Tags) {
				tables = append(tables, table)
			}
		}
	}
	if len(tables) == 0 {
		return nil, fmt.Errorf("unable to find route table for AWS cluster: %s", clusterName)
	}
	if len(tables) != 1 {
		return nil, fmt.Errorf("found multiple matching AWS route tables for AWS cluster: %s", clusterName)
	}
	return tables[0], nil
}
func (c *Cloud) ListRoutes(ctx context.Context, clusterName string) ([]*cloudprovider.Route, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	table, err := c.findRouteTable(clusterName)
	if err != nil {
		return nil, err
	}
	var routes []*cloudprovider.Route
	var instanceIDs []*string
	for _, r := range table.Routes {
		instanceID := aws.StringValue(r.InstanceId)
		if instanceID == "" {
			continue
		}
		instanceIDs = append(instanceIDs, &instanceID)
	}
	instances, err := c.getInstancesByIDs(instanceIDs)
	if err != nil {
		return nil, err
	}
	for _, r := range table.Routes {
		destinationCIDR := aws.StringValue(r.DestinationCidrBlock)
		if destinationCIDR == "" {
			continue
		}
		route := &cloudprovider.Route{Name: clusterName + "-" + destinationCIDR, DestinationCIDR: destinationCIDR}
		if aws.StringValue(r.State) == ec2.RouteStateBlackhole {
			route.Blackhole = true
			routes = append(routes, route)
			continue
		}
		instanceID := aws.StringValue(r.InstanceId)
		if instanceID != "" {
			instance, found := instances[instanceID]
			if found {
				route.TargetNode = mapInstanceToNodeName(instance)
				routes = append(routes, route)
			} else {
				klog.Warningf("unable to find instance ID %s in the list of instances being routed to", instanceID)
			}
		}
	}
	return routes, nil
}
func (c *Cloud) configureInstanceSourceDestCheck(instanceID string, sourceDestCheck bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &ec2.ModifyInstanceAttributeInput{}
	request.InstanceId = aws.String(instanceID)
	request.SourceDestCheck = &ec2.AttributeBooleanValue{Value: aws.Bool(sourceDestCheck)}
	_, err := c.ec2.ModifyInstanceAttribute(request)
	if err != nil {
		return fmt.Errorf("error configuring source-dest-check on instance %s: %q", instanceID, err)
	}
	return nil
}
func (c *Cloud) CreateRoute(ctx context.Context, clusterName string, nameHint string, route *cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instance, err := c.getInstanceByNodeName(route.TargetNode)
	if err != nil {
		return err
	}
	err = c.configureInstanceSourceDestCheck(aws.StringValue(instance.InstanceId), false)
	if err != nil {
		return err
	}
	table, err := c.findRouteTable(clusterName)
	if err != nil {
		return err
	}
	var deleteRoute *ec2.Route
	for _, r := range table.Routes {
		destinationCIDR := aws.StringValue(r.DestinationCidrBlock)
		if destinationCIDR != route.DestinationCIDR {
			continue
		}
		if aws.StringValue(r.State) == ec2.RouteStateBlackhole {
			deleteRoute = r
		}
	}
	if deleteRoute != nil {
		klog.Infof("deleting blackholed route: %s", aws.StringValue(deleteRoute.DestinationCidrBlock))
		request := &ec2.DeleteRouteInput{}
		request.DestinationCidrBlock = deleteRoute.DestinationCidrBlock
		request.RouteTableId = table.RouteTableId
		_, err = c.ec2.DeleteRoute(request)
		if err != nil {
			return fmt.Errorf("error deleting blackholed AWS route (%s): %q", aws.StringValue(deleteRoute.DestinationCidrBlock), err)
		}
	}
	request := &ec2.CreateRouteInput{}
	request.DestinationCidrBlock = aws.String(route.DestinationCIDR)
	request.InstanceId = instance.InstanceId
	request.RouteTableId = table.RouteTableId
	_, err = c.ec2.CreateRoute(request)
	if err != nil {
		return fmt.Errorf("error creating AWS route (%s): %q", route.DestinationCIDR, err)
	}
	return nil
}
func (c *Cloud) DeleteRoute(ctx context.Context, clusterName string, route *cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	table, err := c.findRouteTable(clusterName)
	if err != nil {
		return err
	}
	request := &ec2.DeleteRouteInput{}
	request.DestinationCidrBlock = aws.String(route.DestinationCIDR)
	request.RouteTableId = table.RouteTableId
	_, err = c.ec2.DeleteRoute(request)
	if err != nil {
		return fmt.Errorf("error deleting AWS route (%s): %q", route.DestinationCIDR, err)
	}
	return nil
}
