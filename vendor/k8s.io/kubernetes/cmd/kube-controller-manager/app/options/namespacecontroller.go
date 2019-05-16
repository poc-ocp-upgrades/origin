package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type NamespaceControllerOptions struct {
	NamespaceSyncPeriod      metav1.Duration
	ConcurrentNamespaceSyncs int32
}

func (o *NamespaceControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.DurationVar(&o.NamespaceSyncPeriod.Duration, "namespace-sync-period", o.NamespaceSyncPeriod.Duration, "The period for syncing namespace life-cycle updates")
	fs.Int32Var(&o.ConcurrentNamespaceSyncs, "concurrent-namespace-syncs", o.ConcurrentNamespaceSyncs, "The number of namespace objects that are allowed to sync concurrently. Larger number = more responsive namespace termination, but more CPU (and network) load")
}
func (o *NamespaceControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.NamespaceControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.NamespaceSyncPeriod = o.NamespaceSyncPeriod
	cfg.ConcurrentNamespaceSyncs = o.ConcurrentNamespaceSyncs
	return nil
}
func (o *NamespaceControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
