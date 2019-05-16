package upgrade

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	"strings"
)

type applyPlanFlags struct {
	kubeConfigPath            string
	cfgPath                   string
	featureGatesString        string
	allowExperimentalUpgrades bool
	allowRCUpgrades           bool
	printConfig               bool
	ignorePreflightErrors     []string
	ignorePreflightErrorsSet  sets.String
	out                       io.Writer
}

func NewCmdUpgrade(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := &applyPlanFlags{kubeConfigPath: kubeadmconstants.GetAdminKubeConfigPath(), cfgPath: "", featureGatesString: "", allowExperimentalUpgrades: false, allowRCUpgrades: false, printConfig: false, ignorePreflightErrorsSet: sets.NewString(), out: out}
	cmd := &cobra.Command{Use: "upgrade", Short: "Upgrade your cluster smoothly to a newer version with this command.", RunE: cmdutil.SubCmdRunE("upgrade")}
	flags.kubeConfigPath = cmdutil.FindExistingKubeConfig(flags.kubeConfigPath)
	cmd.AddCommand(NewCmdApply(flags))
	cmd.AddCommand(NewCmdPlan(flags))
	cmd.AddCommand(NewCmdDiff(out))
	cmd.AddCommand(NewCmdNode())
	return cmd
}
func addApplyPlanFlags(fs *pflag.FlagSet, flags *applyPlanFlags) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	options.AddKubeConfigFlag(fs, &flags.kubeConfigPath)
	options.AddConfigFlag(fs, &flags.cfgPath)
	fs.BoolVar(&flags.allowExperimentalUpgrades, "allow-experimental-upgrades", flags.allowExperimentalUpgrades, "Show unstable versions of Kubernetes as an upgrade alternative and allow upgrading to an alpha/beta/release candidate versions of Kubernetes.")
	fs.BoolVar(&flags.allowRCUpgrades, "allow-release-candidate-upgrades", flags.allowRCUpgrades, "Show release candidate versions of Kubernetes as an upgrade alternative and allow upgrading to a release candidate versions of Kubernetes.")
	fs.BoolVar(&flags.printConfig, "print-config", flags.printConfig, "Specifies whether the configuration file that will be used in the upgrade should be printed or not.")
	fs.StringVar(&flags.featureGatesString, "feature-gates", flags.featureGatesString, "A set of key=value pairs that describe feature gates for various features. "+"Options are:\n"+strings.Join(features.KnownFeatures(&features.InitFeatureGates), "\n"))
	options.AddIgnorePreflightErrorsFlag(fs, &flags.ignorePreflightErrors)
}
