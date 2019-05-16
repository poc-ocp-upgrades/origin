package phases

import (
	"github.com/spf13/cobra"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/version"
)

func runCmdPhase(cmdFunc func(outDir string, cfg *kubeadmapi.InitConfiguration) error, outDir, cfgPath *string, cfg *kubeadmapiv1beta1.InitConfiguration, defaultKubernetesVersion string) func(cmd *cobra.Command, args []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(cmd *cobra.Command, args []string) {
		if err := validation.ValidateMixedArguments(cmd.Flags()); err != nil {
			kubeadmutil.CheckErr(err)
		}
		if defaultKubernetesVersion != "" {
			cfg.KubernetesVersion = defaultKubernetesVersion
		} else {
			SetKubernetesVersion(cfg)
		}
		internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(*cfgPath, cfg)
		kubeadmutil.CheckErr(err)
		err = cmdFunc(*outDir, internalcfg)
		kubeadmutil.CheckErr(err)
	}
}
func SetKubernetesVersion(cfg *kubeadmapiv1beta1.InitConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.KubernetesVersion != kubeadmapiv1beta1.DefaultKubernetesVersion && cfg.KubernetesVersion != "" {
		return
	}
	cfg.KubernetesVersion = version.Get().String()
}
