package util

import (
	"errors"
	"net"
)

var ErrorNoDefaultIP = errors.New("no suitable IP address")

func DefaultLocalIP4() (net.IP, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	devices, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, dev := range devices {
		if (dev.Flags&net.FlagUp != 0) && (dev.Flags&net.FlagLoopback == 0) {
			addrs, err := dev.Addrs()
			if err != nil {
				continue
			}
			for i := range addrs {
				if ip, ok := addrs[i].(*net.IPNet); ok {
					if ip.IP.To4() != nil {
						return ip.IP, nil
					}
				}
			}
		}
	}
	return nil, ErrorNoDefaultIP
}
func AllLocalIP4() ([]net.IP, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	devices, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ret := []net.IP{}
	for _, dev := range devices {
		if dev.Flags&net.FlagUp != 0 {
			addrs, err := dev.Addrs()
			if err != nil {
				continue
			}
			for i := range addrs {
				if ip, ok := addrs[i].(*net.IPNet); ok {
					if ip.IP.To4() != nil {
						ret = append(ret, ip.IP)
					}
				}
			}
		}
	}
	return ret, nil
}
