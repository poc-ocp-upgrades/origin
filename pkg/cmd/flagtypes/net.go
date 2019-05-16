package flagtypes

import (
	"fmt"
	"net"
	"strings"
)

type IP net.IP

func (ip IP) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return net.IP(ip).String()
}
func (ip *IP) Set(value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	*ip = IP(net.ParseIP(strings.TrimSpace(value)))
	if *ip == nil {
		return fmt.Errorf("invalid IP address: '%s'", value)
	}
	return nil
}
func (ip *IP) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "cmd.flagtypes.IP"
}

type IPNet net.IPNet

func DefaultIPNet(value string) IPNet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := IPNet{}
	if err := ret.Set(value); err != nil {
		panic(err)
	}
	return ret
}
func (ipnet IPNet) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n := net.IPNet(ipnet)
	return n.String()
}
func (ipnet *IPNet) Set(value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, n, err := net.ParseCIDR(strings.TrimSpace(value))
	if err != nil {
		return err
	}
	*ipnet = IPNet(*n)
	return nil
}
func (ipnet *IPNet) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "cmd.flagtypes.IPNet"
}
