package options

import (
	"fmt"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"net"
)

func NewSecureServingOptions() *genericoptions.SecureServingOptionsWithLoopback {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o := genericoptions.SecureServingOptions{BindAddress: net.ParseIP("0.0.0.0"), BindPort: 6443, Required: true, ServerCert: genericoptions.GeneratableKeyCert{PairName: "apiserver", CertDirectory: "/var/run/kubernetes"}}
	return o.WithLoopback()
}
func NewInsecureServingOptions() *genericoptions.DeprecatedInsecureServingOptionsWithLoopback {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o := genericoptions.DeprecatedInsecureServingOptions{BindAddress: net.ParseIP("127.0.0.1"), BindPort: 8080}
	return o.WithLoopback()
}
func DefaultAdvertiseAddress(s *genericoptions.ServerRunOptions, insecure *genericoptions.DeprecatedInsecureServingOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if insecure == nil {
		return nil
	}
	if s.AdvertiseAddress == nil || s.AdvertiseAddress.IsUnspecified() {
		hostIP, err := utilnet.ChooseBindAddress(insecure.BindAddress)
		if err != nil {
			return fmt.Errorf("unable to find suitable network address.error='%v'. "+"Try to set the AdvertiseAddress directly or provide a valid BindAddress to fix this", err)
		}
		s.AdvertiseAddress = hostIP
	}
	return nil
}
