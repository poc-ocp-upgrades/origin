package network

import (
	"strings"
	"time"
)

const (
	SingleTenantPluginName		= "redhat/openshift-ovs-subnet"
	MultiTenantPluginName		= "redhat/openshift-ovs-multitenant"
	NetworkPolicyPluginName		= "redhat/openshift-ovs-networkpolicy"
	DefaultInformerResyncPeriod	= 30 * time.Minute
)

func IsOpenShiftNetworkPlugin(pluginName string) bool {
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
	switch strings.ToLower(pluginName) {
	case SingleTenantPluginName, MultiTenantPluginName, NetworkPolicyPluginName:
		return true
	}
	return false
}
func IsOpenShiftMultitenantNetworkPlugin(pluginName string) bool {
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
	if strings.ToLower(pluginName) == MultiTenantPluginName {
		return true
	}
	return false
}
