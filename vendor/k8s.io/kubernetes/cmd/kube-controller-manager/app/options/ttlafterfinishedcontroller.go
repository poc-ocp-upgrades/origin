package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type TTLAfterFinishedControllerOptions struct{ ConcurrentTTLSyncs int32 }

func (o *TTLAfterFinishedControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.Int32Var(&o.ConcurrentTTLSyncs, "concurrent-ttl-after-finished-syncs", o.ConcurrentTTLSyncs, "The number of TTL-after-finished controller workers that are allowed to sync concurrently.")
}
func (o *TTLAfterFinishedControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.TTLAfterFinishedControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ConcurrentTTLSyncs = o.ConcurrentTTLSyncs
	return nil
}
func (o *TTLAfterFinishedControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
