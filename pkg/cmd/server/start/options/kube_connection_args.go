package options

import (
	"errors"
	"github.com/openshift/origin/pkg/cmd/flagtypes"
	"github.com/spf13/pflag"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/url"
)

type KubeConnectionArgs struct {
	KubernetesAddr           flagtypes.Addr
	ClientConfig             clientcmd.ClientConfig
	ClientConfigLoadingRules clientcmd.ClientConfigLoadingRules
}

func BindKubeConnectionArgs(args *KubeConnectionArgs, flags *pflag.FlagSet, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.Var(&args.KubernetesAddr, prefix+"kubernetes", "removed in favor of --"+prefix+"kubeconfig")
	flags.StringVar(&args.ClientConfigLoadingRules.ExplicitPath, prefix+"kubeconfig", "", "Path to the kubeconfig file to use for requests to the Kubernetes API.")
}
func NewDefaultKubeConnectionArgs() *KubeConnectionArgs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &KubeConnectionArgs{}
	config.KubernetesAddr = flagtypes.Addr{Value: "localhost:8443", DefaultScheme: "https", DefaultPort: 8443, AllowPrefix: true}.Default()
	config.ClientConfig = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&config.ClientConfigLoadingRules, &clientcmd.ConfigOverrides{})
	return config
}
func (args KubeConnectionArgs) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if args.KubernetesAddr.Provided {
		return errors.New("--kubernetes is no longer allowed, try using --kubeconfig")
	}
	return nil
}
func (args KubeConnectionArgs) GetExternalKubernetesClientConfig() (*restclient.Config, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args.ClientConfigLoadingRules.ExplicitPath) == 0 || args.ClientConfig == nil {
		return nil, false, nil
	}
	clientConfig, err := args.ClientConfig.ClientConfig()
	if err != nil {
		return nil, false, err
	}
	return clientConfig, true, nil
}
func (args KubeConnectionArgs) GetKubernetesAddress(defaultAddress *url.URL) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config, ok, err := args.GetExternalKubernetesClientConfig()
	if err != nil {
		return nil, err
	}
	if ok && len(config.Host) > 0 {
		return url.Parse(config.Host)
	}
	return defaultAddress, nil
}
