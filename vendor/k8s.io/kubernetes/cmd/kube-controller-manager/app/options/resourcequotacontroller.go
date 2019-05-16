package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type ResourceQuotaControllerOptions struct {
	ResourceQuotaSyncPeriod      metav1.Duration
	ConcurrentResourceQuotaSyncs int32
}

func (o *ResourceQuotaControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.DurationVar(&o.ResourceQuotaSyncPeriod.Duration, "resource-quota-sync-period", o.ResourceQuotaSyncPeriod.Duration, "The period for syncing quota usage status in the system")
	fs.Int32Var(&o.ConcurrentResourceQuotaSyncs, "concurrent-resource-quota-syncs", o.ConcurrentResourceQuotaSyncs, "The number of resource quotas that are allowed to sync concurrently. Larger number = more responsive quota management, but more CPU (and network) load")
}
func (o *ResourceQuotaControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.ResourceQuotaControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ResourceQuotaSyncPeriod = o.ResourceQuotaSyncPeriod
	cfg.ConcurrentResourceQuotaSyncs = o.ConcurrentResourceQuotaSyncs
	return nil
}
func (o *ResourceQuotaControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
