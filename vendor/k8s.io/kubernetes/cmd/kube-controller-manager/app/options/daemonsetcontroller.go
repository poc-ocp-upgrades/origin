package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type DaemonSetControllerOptions struct{ ConcurrentDaemonSetSyncs int32 }

func (o *DaemonSetControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
}
func (o *DaemonSetControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.DaemonSetControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ConcurrentDaemonSetSyncs = o.ConcurrentDaemonSetSyncs
	return nil
}
func (o *DaemonSetControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
