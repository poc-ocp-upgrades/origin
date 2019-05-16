package options

import (
	"github.com/spf13/pflag"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type NodeIPAMControllerOptions struct {
	ServiceCIDR      string
	NodeCIDRMaskSize int32
}

func (o *NodeIPAMControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.StringVar(&o.ServiceCIDR, "service-cluster-ip-range", o.ServiceCIDR, "CIDR Range for Services in cluster. Requires --allocate-node-cidrs to be true")
	fs.Int32Var(&o.NodeCIDRMaskSize, "node-cidr-mask-size", o.NodeCIDRMaskSize, "Mask size for node cidr in cluster.")
}
func (o *NodeIPAMControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.NodeIPAMControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.ServiceCIDR = o.ServiceCIDR
	cfg.NodeCIDRMaskSize = o.NodeCIDRMaskSize
	return nil
}
func (o *NodeIPAMControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
