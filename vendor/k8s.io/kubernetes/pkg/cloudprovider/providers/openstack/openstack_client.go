package openstack

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func (os *OpenStack) NewNetworkV2() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	network, err := openstack.NewNetworkV2(os.provider, gophercloud.EndpointOpts{Region: os.region})
	if err != nil {
		return nil, fmt.Errorf("failed to find network v2 endpoint for region %s: %v", os.region, err)
	}
	return network, nil
}
func (os *OpenStack) NewComputeV2() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	compute, err := openstack.NewComputeV2(os.provider, gophercloud.EndpointOpts{Region: os.region})
	if err != nil {
		return nil, fmt.Errorf("failed to find compute v2 endpoint for region %s: %v", os.region, err)
	}
	return compute, nil
}
func (os *OpenStack) NewBlockStorageV1() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage, err := openstack.NewBlockStorageV1(os.provider, gophercloud.EndpointOpts{Region: os.region})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize cinder v1 client for region %s: %v", os.region, err)
	}
	return storage, nil
}
func (os *OpenStack) NewBlockStorageV2() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage, err := openstack.NewBlockStorageV2(os.provider, gophercloud.EndpointOpts{Region: os.region})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize cinder v2 client for region %s: %v", os.region, err)
	}
	return storage, nil
}
func (os *OpenStack) NewBlockStorageV3() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage, err := openstack.NewBlockStorageV3(os.provider, gophercloud.EndpointOpts{Region: os.region})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize cinder v3 client for region %s: %v", os.region, err)
	}
	return storage, nil
}
func (os *OpenStack) NewLoadBalancerV2() (*gophercloud.ServiceClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var lb *gophercloud.ServiceClient
	var err error
	if os.lbOpts.UseOctavia {
		lb, err = openstack.NewLoadBalancerV2(os.provider, gophercloud.EndpointOpts{Region: os.region})
	} else {
		lb, err = openstack.NewNetworkV2(os.provider, gophercloud.EndpointOpts{Region: os.region})
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find load-balancer v2 endpoint for region %s: %v", os.region, err)
	}
	return lb, nil
}
