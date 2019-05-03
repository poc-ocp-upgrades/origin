package options

import (
 "github.com/spf13/pflag"
)

type CloudProviderOptions struct {
 CloudConfigFile string
 CloudProvider   string
}

func NewCloudProviderOptions() *CloudProviderOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &CloudProviderOptions{}
}
func (s *CloudProviderOptions) Validate() []error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := []error{}
 return allErrors
}
func (s *CloudProviderOptions) AddFlags(fs *pflag.FlagSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fs.StringVar(&s.CloudProvider, "cloud-provider", s.CloudProvider, "The provider for cloud services. Empty string for no provider.")
 fs.StringVar(&s.CloudConfigFile, "cloud-config", s.CloudConfigFile, "The path to the cloud provider configuration file. Empty string for no configuration file.")
}
