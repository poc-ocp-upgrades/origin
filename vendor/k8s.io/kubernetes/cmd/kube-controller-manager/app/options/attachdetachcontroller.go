package options

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type AttachDetachControllerOptions struct {
	ReconcilerSyncLoopPeriod          metav1.Duration
	DisableAttachDetachReconcilerSync bool
}

func (o *AttachDetachControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.BoolVar(&o.DisableAttachDetachReconcilerSync, "disable-attach-detach-reconcile-sync", false, "Disable volume attach detach reconciler sync. Disabling this may cause volumes to be mismatched with pods. Use wisely.")
	fs.DurationVar(&o.ReconcilerSyncLoopPeriod.Duration, "attach-detach-reconcile-sync-period", o.ReconcilerSyncLoopPeriod.Duration, "The reconciler sync wait time between volume attach detach. This duration must be larger than one second, and increasing this value from the default may allow for volumes to be mismatched with pods.")
}
func (o *AttachDetachControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.AttachDetachControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.DisableAttachDetachReconcilerSync = o.DisableAttachDetachReconcilerSync
	cfg.ReconcilerSyncLoopPeriod = o.ReconcilerSyncLoopPeriod
	return nil
}
func (o *AttachDetachControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
