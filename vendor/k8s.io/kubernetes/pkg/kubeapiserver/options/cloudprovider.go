package options

import (
	"github.com/spf13/pflag"
)

type CloudProviderOptions struct {
	CloudConfigFile string
	CloudProvider   string
}

func NewCloudProviderOptions() *CloudProviderOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &CloudProviderOptions{}
}
func (s *CloudProviderOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrors := []error{}
	return allErrors
}
func (s *CloudProviderOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(&s.CloudProvider, "cloud-provider", s.CloudProvider, "The provider for cloud services. Empty string for no provider.")
	fs.StringVar(&s.CloudConfigFile, "cloud-config", s.CloudConfigFile, "The path to the cloud provider configuration file. Empty string for no configuration file.")
}
