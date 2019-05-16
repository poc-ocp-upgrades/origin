package util

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/validation"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"net"
	"net/url"
	"strconv"
)

func GetMasterEndpoint(cfg *kubeadmapi.InitConfiguration) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bindPortString := strconv.Itoa(int(cfg.LocalAPIEndpoint.BindPort))
	if _, err := ParsePort(bindPortString); err != nil {
		return "", errors.Wrapf(err, "invalid value %q given for api.bindPort", cfg.LocalAPIEndpoint.BindPort)
	}
	var ip = net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress)
	if ip == nil {
		return "", errors.Errorf("invalid value `%s` given for api.advertiseAddress", cfg.LocalAPIEndpoint.AdvertiseAddress)
	}
	masterURL := &url.URL{Scheme: "https", Host: net.JoinHostPort(ip.String(), bindPortString)}
	if len(cfg.ControlPlaneEndpoint) > 0 {
		var host, port string
		var err error
		if host, port, err = ParseHostPort(cfg.ControlPlaneEndpoint); err != nil {
			return "", errors.Wrapf(err, "invalid value %q given for controlPlaneEndpoint", cfg.ControlPlaneEndpoint)
		}
		if port != "" {
			if port != bindPortString {
				fmt.Println("[endpoint] WARNING: port specified in controlPlaneEndpoint overrides bindPort in the controlplane address")
			}
		} else {
			port = bindPortString
		}
		masterURL = &url.URL{Scheme: "https", Host: net.JoinHostPort(host, port)}
	}
	return masterURL.String(), nil
}
func ParseHostPort(hostport string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var host, port string
	var err error
	if host, port, err = net.SplitHostPort(hostport); err != nil {
		host = hostport
	}
	if port != "" {
		if _, err := ParsePort(port); err != nil {
			return "", "", errors.New("port must be a valid number between 1 and 65535, inclusive")
		}
	}
	if ip := net.ParseIP(host); ip != nil {
		return host, port, nil
	}
	if errs := validation.IsDNS1123Subdomain(host); len(errs) == 0 {
		return host, port, nil
	}
	return "", "", errors.New("host must be a valid IP address or a valid RFC-1123 DNS subdomain")
}
func ParsePort(port string) (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	portInt, err := strconv.Atoi(port)
	if err == nil && (1 <= portInt && portInt <= 65535) {
		return portInt, nil
	}
	return 0, errors.New("port must be a valid number between 1 and 65535, inclusive")
}
