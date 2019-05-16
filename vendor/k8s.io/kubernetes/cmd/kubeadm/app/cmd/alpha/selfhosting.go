package alpha

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/selfhosting"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/pkg/util/normalizer"
	"os"
	"strings"
	"time"
)

var (
	selfhostingLongDesc = normalizer.LongDesc(`
		Converts static Pod files for control plane components into self-hosted DaemonSets configured via the Kubernetes API.

		See the documentation for self-hosting limitations.

		` + cmdutil.AlphaDisclaimer)
	selfhostingExample = normalizer.Examples(`
		# Converts a static Pod-hosted control plane into a self-hosted one. 

		kubeadm alpha phase self-hosting convert-from-staticpods
		`)
)

func NewCmdSelfhosting(in io.Reader) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "selfhosting", Aliases: []string{"selfhosted", "self-hosting"}, Short: "Makes a kubeadm cluster self-hosted", Long: cmdutil.MacroCommandLongDescription}
	cmd.AddCommand(getSelfhostingSubCommand(in))
	return cmd
}
func getSelfhostingSubCommand(in io.Reader) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := &kubeadmapiv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(cfg)
	var cfgPath, featureGatesString string
	forcePivot, certsInSecrets := false, false
	kubeConfigFile := constants.GetAdminKubeConfigPath()
	cmd := &cobra.Command{Use: "pivot", Aliases: []string{"from-staticpods"}, Short: "Converts a static Pod-hosted control plane into a self-hosted one", Long: selfhostingLongDesc, Example: selfhostingExample, Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !forcePivot {
			fmt.Println("WARNING: self-hosted clusters are not supported by kubeadm upgrade and by other kubeadm commands!")
			fmt.Print("[pivot] are you sure you want to proceed? [y/n]: ")
			s := bufio.NewScanner(in)
			s.Scan()
			err = s.Err()
			kubeadmutil.CheckErr(err)
			if strings.ToLower(s.Text()) != "y" {
				kubeadmutil.CheckErr(errors.New("aborted pivot operation"))
			}
		}
		fmt.Println("[pivot] pivoting cluster to self-hosted")
		if cfg.FeatureGates, err = features.NewFeatureGate(&features.InitFeatureGates, featureGatesString); err != nil {
			kubeadmutil.CheckErr(err)
		}
		if err := validation.ValidateMixedArguments(cmd.Flags()); err != nil {
			kubeadmutil.CheckErr(err)
		}
		kubeConfigFile = cmdutil.FindExistingKubeConfig(kubeConfigFile)
		client, err := kubeconfigutil.ClientSetFromFile(kubeConfigFile)
		kubeadmutil.CheckErr(err)
		phases.SetKubernetesVersion(cfg)
		internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(cfgPath, cfg)
		kubeadmutil.CheckErr(err)
		waiter := apiclient.NewKubeWaiter(client, 2*time.Minute, os.Stdout)
		err = selfhosting.CreateSelfHostedControlPlane(constants.GetStaticPodDirectory(), constants.KubernetesDir, internalcfg, client, waiter, false, certsInSecrets)
		kubeadmutil.CheckErr(err)
	}}
	cmd.Flags().StringVar(&cfg.CertificatesDir, "cert-dir", cfg.CertificatesDir, `The path where certificates are stored`)
	cmd.Flags().StringVar(&cfgPath, "config", cfgPath, "Path to a kubeadm config file. WARNING: Usage of a configuration file is experimental")
	cmd.Flags().BoolVarP(&certsInSecrets, "store-certs-in-secrets", "s", false, "Enable storing certs in secrets")
	cmd.Flags().BoolVarP(&forcePivot, "force", "f", false, "Pivot the cluster without prompting for confirmation")
	options.AddKubeConfigFlag(cmd.Flags(), &kubeConfigFile)
	return cmd
}
