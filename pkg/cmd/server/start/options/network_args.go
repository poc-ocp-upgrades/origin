package options

import (
	"github.com/spf13/pflag"
)

type NetworkArgs struct {
	NetworkPluginName	string
	ClusterNetworkCIDR	string
	HostSubnetLength	uint32
	ServiceNetworkCIDR	string
}

func BindNetworkArgs(args *NetworkArgs, flags *pflag.FlagSet, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.StringVar(&args.NetworkPluginName, prefix+"network-plugin", args.NetworkPluginName, "The name of the networking plugin to be used for networking.")
	flags.StringVar(&args.ClusterNetworkCIDR, prefix+"network-cidr", args.ClusterNetworkCIDR, "The CIDR string representing the network that all containers should belong to.")
	flags.Uint32Var(&args.HostSubnetLength, prefix+"host-subnet-length", args.HostSubnetLength, "The length of subnet each host is given from the network-cidr.")
	flags.StringVar(&args.ServiceNetworkCIDR, prefix+"portal-net", args.ServiceNetworkCIDR, "The CIDR string representing the network that portal/service IPs will be assigned from. This must not overlap with any IP ranges assigned to nodes for pods.")
}
func NewDefaultMasterNetworkArgs() *NetworkArgs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &NetworkArgs{NetworkPluginName: "", ClusterNetworkCIDR: "10.128.0.0/14", HostSubnetLength: 9, ServiceNetworkCIDR: "172.30.0.0/16"}
	return config
}
