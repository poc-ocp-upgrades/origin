package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type DeploymentControllerOptions struct {
	ConcurrentDeploymentSyncs      int32
	DeploymentControllerSyncPeriod metav1.Duration
}

func (o *DeploymentControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.Int32Var(&o.ConcurrentDeploymentSyncs, "concurrent-deployment-syncs", o.ConcurrentDeploymentSyncs, "The number of deployment objects that are allowed to sync concurrently. Larger number = more responsive deployments, but more CPU (and network) load")
	fs.DurationVar(&o.DeploymentControllerSyncPeriod.Duration, "deployment-controller-sync-period", o.DeploymentControllerSyncPeriod.Duration, "Period for syncing the deployments.")
}
func (o *DeploymentControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.DeploymentControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ConcurrentDeploymentSyncs = o.ConcurrentDeploymentSyncs
	cfg.DeploymentControllerSyncPeriod = o.DeploymentControllerSyncPeriod
	return nil
}
func (o *DeploymentControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
