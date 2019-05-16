package alpha

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/preflight"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/util/normalizer"
	utilsexec "k8s.io/utils/exec"
)

var (
	nodePreflightLongDesc = normalizer.LongDesc(`
		Run node pre-flight checks, functionally equivalent to what implemented by kubeadm join.
		` + cmdutil.AlphaDisclaimer)
	nodePreflightExample = normalizer.Examples(`
		# Run node pre-flight checks.
		kubeadm alpha preflight node
	`)
	errorMissingConfigFlag = errors.New("the --config flag is mandatory")
)

func newCmdPreFlightUtility() *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "preflight", Short: "Commands related to pre-flight checks", Long: cmdutil.MacroCommandLongDescription}
	cmd.AddCommand(newCmdPreFlightNode())
	return cmd
}
func newCmdPreFlightNode() *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cfgPath string
	var ignorePreflightErrors []string
	cmd := &cobra.Command{Use: "node", Short: "Run node pre-flight checks", Long: nodePreflightLongDesc, Example: nodePreflightExample, Run: func(cmd *cobra.Command, args []string) {
		if len(cfgPath) == 0 {
			kubeadmutil.CheckErr(errorMissingConfigFlag)
		}
		ignorePreflightErrorsSet, err := validation.ValidateIgnorePreflightErrors(ignorePreflightErrors)
		kubeadmutil.CheckErr(err)
		cfg := &kubeadmapiv1beta1.JoinConfiguration{}
		kubeadmscheme.Scheme.Default(cfg)
		internalcfg, err := configutil.JoinConfigFileAndDefaultsToInternalConfig(cfgPath, cfg)
		kubeadmutil.CheckErr(err)
		if internalcfg.ControlPlane != nil {
			err = configutil.VerifyAPIServerBindAddress(internalcfg.ControlPlane.LocalAPIEndpoint.AdvertiseAddress)
			kubeadmutil.CheckErr(err)
		}
		fmt.Println("[preflight] running pre-flight checks")
		err = preflight.RunJoinNodeChecks(utilsexec.New(), internalcfg, ignorePreflightErrorsSet)
		kubeadmutil.CheckErr(err)
		fmt.Println("[preflight] pre-flight checks passed")
	}}
	options.AddConfigFlag(cmd.PersistentFlags(), &cfgPath)
	options.AddIgnorePreflightErrorsFlag(cmd.PersistentFlags(), &ignorePreflightErrors)
	return cmd
}
