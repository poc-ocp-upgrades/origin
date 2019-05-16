package options

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type CloudProviderOptions struct {
	CloudConfigFile string
	Name            string
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
	fs.StringVar(&s.Name, "cloud-provider", s.Name, "The provider for cloud services. Empty string for no provider.")
	fs.StringVar(&s.CloudConfigFile, "cloud-config", s.CloudConfigFile, "The path to the cloud provider configuration file. Empty string for no configuration file.")
}
func (s *CloudProviderOptions) ApplyTo(cfg *kubectrlmgrconfig.CloudProviderConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s == nil {
		return nil
	}
	cfg.Name = s.Name
	cfg.CloudConfigFile = s.CloudConfigFile
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
