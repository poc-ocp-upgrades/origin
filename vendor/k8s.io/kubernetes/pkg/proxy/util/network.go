package util

import (
	"net"
)

type NetworkInterfacer interface {
	Addrs(intf *net.Interface) ([]net.Addr, error)
	Interfaces() ([]net.Interface, error)
}
type RealNetwork struct{}

func (_ RealNetwork) Addrs(intf *net.Interface) ([]net.Addr, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return intf.Addrs()
}
func (_ RealNetwork) Interfaces() ([]net.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return net.Interfaces()
}

var _ NetworkInterfacer = &RealNetwork{}
