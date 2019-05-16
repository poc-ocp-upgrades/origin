package leaderelectionconfig

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	apiserverconfig "k8s.io/apiserver/pkg/apis/config"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	DefaultLeaseDuration = 15 * time.Second
)

func BindFlags(l *apiserverconfig.LeaderElectionConfiguration, fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.BoolVar(&l.LeaderElect, "leader-elect", l.LeaderElect, ""+"Start a leader election client and gain leadership before "+"executing the main loop. Enable this when running replicated "+"components for high availability.")
	fs.DurationVar(&l.LeaseDuration.Duration, "leader-elect-lease-duration", l.LeaseDuration.Duration, ""+"The duration that non-leader candidates will wait after observing a leadership "+"renewal until attempting to acquire leadership of a led but unrenewed leader "+"slot. This is effectively the maximum duration that a leader can be stopped "+"before it is replaced by another candidate. This is only applicable if leader "+"election is enabled.")
	fs.DurationVar(&l.RenewDeadline.Duration, "leader-elect-renew-deadline", l.RenewDeadline.Duration, ""+"The interval between attempts by the acting master to renew a leadership slot "+"before it stops leading. This must be less than or equal to the lease duration. "+"This is only applicable if leader election is enabled.")
	fs.DurationVar(&l.RetryPeriod.Duration, "leader-elect-retry-period", l.RetryPeriod.Duration, ""+"The duration the clients should wait between attempting acquisition and renewal "+"of a leadership. This is only applicable if leader election is enabled.")
	fs.StringVar(&l.ResourceLock, "leader-elect-resource-lock", l.ResourceLock, ""+"The type of resource object that is used for locking during "+"leader election. Supported options are `endpoints` (default) and `configmaps`.")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
