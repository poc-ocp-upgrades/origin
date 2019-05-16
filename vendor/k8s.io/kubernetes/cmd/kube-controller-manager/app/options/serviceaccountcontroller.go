package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type SAControllerOptions struct {
	ServiceAccountKeyFile  string
	ConcurrentSATokenSyncs int32
	RootCAFile             string
}

func (o *SAControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.StringVar(&o.ServiceAccountKeyFile, "service-account-private-key-file", o.ServiceAccountKeyFile, "Filename containing a PEM-encoded private RSA or ECDSA key used to sign service account tokens.")
	fs.Int32Var(&o.ConcurrentSATokenSyncs, "concurrent-serviceaccount-token-syncs", o.ConcurrentSATokenSyncs, "The number of service account token objects that are allowed to sync concurrently. Larger number = more responsive token generation, but more CPU (and network) load")
	fs.StringVar(&o.RootCAFile, "root-ca-file", o.RootCAFile, "If set, this root certificate authority will be included in service account's token secret. This must be a valid PEM-encoded CA bundle.")
}
func (o *SAControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.SAControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ServiceAccountKeyFile = o.ServiceAccountKeyFile
	cfg.ConcurrentSATokenSyncs = o.ConcurrentSATokenSyncs
	cfg.RootCAFile = o.RootCAFile
	return nil
}
func (o *SAControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
