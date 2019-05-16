package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

const (
	DefaultClusterSigningCertFile = "/etc/kubernetes/ca/ca.pem"
	DefaultClusterSigningKeyFile  = "/etc/kubernetes/ca/ca.key"
)

type CSRSigningControllerOptions struct {
	ClusterSigningDuration metav1.Duration
	ClusterSigningKeyFile  string
	ClusterSigningCertFile string
}

func (o *CSRSigningControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.StringVar(&o.ClusterSigningCertFile, "cluster-signing-cert-file", o.ClusterSigningCertFile, "Filename containing a PEM-encoded X509 CA certificate used to issue cluster-scoped certificates")
	fs.StringVar(&o.ClusterSigningKeyFile, "cluster-signing-key-file", o.ClusterSigningKeyFile, "Filename containing a PEM-encoded RSA or ECDSA private key used to sign cluster-scoped certificates")
	fs.DurationVar(&o.ClusterSigningDuration.Duration, "experimental-cluster-signing-duration", o.ClusterSigningDuration.Duration, "The length of duration signed certificates will be given.")
}
func (o *CSRSigningControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.CSRSigningControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ClusterSigningCertFile = o.ClusterSigningCertFile
	cfg.ClusterSigningKeyFile = o.ClusterSigningKeyFile
	cfg.ClusterSigningDuration = o.ClusterSigningDuration
	return nil
}
func (o *CSRSigningControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
