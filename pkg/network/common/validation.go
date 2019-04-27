package common

import (
	networkapi "github.com/openshift/api/network/v1"
	internalnetworkapi "github.com/openshift/origin/pkg/network/apis/network"
	internalnetworkv1 "github.com/openshift/origin/pkg/network/apis/network/v1"
	networkvalidation "github.com/openshift/origin/pkg/network/apis/network/validation"
)

func ValidateClusterNetwork(cn *networkapi.ClusterNetwork) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	icn := &internalnetworkapi.ClusterNetwork{}
	if err := internalnetworkv1.Convert_v1_ClusterNetwork_To_network_ClusterNetwork(cn, icn, nil); err != nil {
		return err
	}
	if errs := networkvalidation.ValidateClusterNetwork(icn); len(errs) > 0 {
		return errs.ToAggregate()
	} else {
		return nil
	}
}
func ValidateHostSubnet(hs *networkapi.HostSubnet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ihs := &internalnetworkapi.HostSubnet{}
	if err := internalnetworkv1.Convert_v1_HostSubnet_To_network_HostSubnet(hs, ihs, nil); err != nil {
		return err
	}
	if errs := networkvalidation.ValidateHostSubnet(ihs); len(errs) > 0 {
		return errs.ToAggregate()
	} else {
		return nil
	}
}
