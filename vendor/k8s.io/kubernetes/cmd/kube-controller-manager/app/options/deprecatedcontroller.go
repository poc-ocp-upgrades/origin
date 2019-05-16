package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type DeprecatedControllerOptions struct {
	DeletingPodsQPS    float32
	DeletingPodsBurst  int32
	RegisterRetryCount int32
}

func (o *DeprecatedControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.Float32Var(&o.DeletingPodsQPS, "deleting-pods-qps", 0.1, "Number of nodes per second on which pods are deleted in case of node failure.")
	fs.MarkDeprecated("deleting-pods-qps", "This flag is currently no-op and will be deleted.")
	fs.Int32Var(&o.DeletingPodsBurst, "deleting-pods-burst", 0, "Number of nodes on which pods are bursty deleted in case of node failure. For more details look into RateLimiter.")
	fs.MarkDeprecated("deleting-pods-burst", "This flag is currently no-op and will be deleted.")
	fs.Int32Var(&o.RegisterRetryCount, "register-retry-count", o.RegisterRetryCount, ""+"The number of retries for initial node registration.  Retry interval equals node-sync-period.")
	fs.MarkDeprecated("register-retry-count", "This flag is currently no-op and will be deleted.")
}
func (o *DeprecatedControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.DeprecatedControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.DeletingPodsQPS = o.DeletingPodsQPS
	cfg.DeletingPodsBurst = o.DeletingPodsBurst
	cfg.RegisterRetryCount = o.RegisterRetryCount
	return nil
}
func (o *DeprecatedControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
