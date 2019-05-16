package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type ReplicationControllerOptions struct{ ConcurrentRCSyncs int32 }

func (o *ReplicationControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.Int32Var(&o.ConcurrentRCSyncs, "concurrent_rc_syncs", o.ConcurrentRCSyncs, "The number of replication controllers that are allowed to sync concurrently. Larger number = more responsive replica management, but more CPU (and network) load")
}
func (o *ReplicationControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.ReplicationControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ConcurrentRCSyncs = o.ConcurrentRCSyncs
	return nil
}
func (o *ReplicationControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
