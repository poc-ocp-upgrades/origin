package options

import (
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/util/globalflag"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers"
)

func AddCustomGlobalFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	globalflag.Register(fs, "cloud-provider-gce-lb-src-cidrs")
}
