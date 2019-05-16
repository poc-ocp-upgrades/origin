package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	initDoneTempl = template.Must(template.New("init").Parse(dedent.Dedent(`
		Your Kubernetes master has initialized successfully!

		To start using your cluster, you need to run the following as a regular user:

		  mkdir -p $HOME/.kube
		  sudo cp -i {{.KubeConfigPath}} $HOME/.kube/config
		  sudo chown $(id -u):$(id -g) $HOME/.kube/config

		You should now deploy a pod network to the cluster.
		Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
		  https://kubernetes.io/docs/concepts/cluster-administration/addons/

		You can now join any number of machines by running the following on each node
		as root:

		  {{.joinCommand}}

		`)))
	kubeletFailTempl = template.Must(template.New("init").Parse(dedent.Dedent(`
		Unfortunately, an error has occurred:
			{{ .Error }}

		This error is likely caused by:
			- The kubelet is not running
			- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)

		If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
			- 'systemctl status kubelet'
			- 'journalctl -xeu kubelet'

		Additionally, a control plane component may have crashed or exited when started by the container runtime.
		To troubleshoot, list all containers using your preferred container runtimes CLI, e.g. docker.
		Here is one example how you may list all Kubernetes containers running in docker:
			- 'docker ps -a | grep kube | grep -v pause'
			Once you have found the failing container, you can inspect its logs with:
			- 'docker logs CONTAINERID'
		`)))
)

type initOptions struct {
	cfgPath               string
	skipTokenPrint        bool
	dryRun                bool
	kubeconfigDir         string
	kubeconfigPath        string
	featureGatesString    string
	ignorePreflightErrors []string
	bto                   *options.BootstrapTokenOptions
	externalcfg           *kubeadmapiv1beta1.InitConfiguration
}
type initData struct {
	cfg                   *kubeadmapi.InitConfiguration
	skipTokenPrint        bool
	dryRun                bool
	kubeconfigDir         string
	kubeconfigPath        string
	ignorePreflightErrors sets.String
	certificatesDir       string
	dryRunDir             string
	externalCA            bool
	client                clientset.Interface
	waiter                apiclient.Waiter
	outputWriter          io.Writer
}

func NewCmdInit(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	initOptions := newInitOptions()
	initRunner := workflow.NewRunner()
	cmd := &cobra.Command{Use: "init", Short: "Run this command in order to set up the Kubernetes master.", Run: func(cmd *cobra.Command, args []string) {
		c, err := initRunner.InitData()
		kubeadmutil.CheckErr(err)
		data := c.(initData)
		fmt.Printf("[init] Using Kubernetes version: %s\n", data.cfg.KubernetesVersion)
		err = initRunner.Run()
		kubeadmutil.CheckErr(err)
		err = showJoinCommand(&data, out)
		kubeadmutil.CheckErr(err)
	}}
	AddInitConfigFlags(cmd.Flags(), initOptions.externalcfg, &initOptions.featureGatesString)
	AddInitOtherFlags(cmd.Flags(), &initOptions.cfgPath, &initOptions.skipTokenPrint, &initOptions.dryRun, &initOptions.ignorePreflightErrors)
	initOptions.bto.AddTokenFlag(cmd.Flags())
	initOptions.bto.AddTTLFlag(cmd.Flags())
	options.AddImageMetaFlags(cmd.Flags(), &initOptions.externalcfg.ImageRepository)
	initRunner.SetAdditionalFlags(func(flags *flag.FlagSet) {
		options.AddKubeConfigFlag(flags, &initOptions.kubeconfigPath)
		options.AddKubeConfigDirFlag(flags, &initOptions.kubeconfigDir)
		options.AddControlPlanExtraArgsFlags(flags, &initOptions.externalcfg.APIServer.ExtraArgs, &initOptions.externalcfg.ControllerManager.ExtraArgs, &initOptions.externalcfg.Scheduler.ExtraArgs)
	})
	initRunner.AppendPhase(phases.NewPreflightMasterPhase())
	initRunner.AppendPhase(phases.NewKubeletStartPhase())
	initRunner.AppendPhase(phases.NewCertsPhase())
	initRunner.AppendPhase(phases.NewKubeConfigPhase())
	initRunner.AppendPhase(phases.NewControlPlanePhase())
	initRunner.AppendPhase(phases.NewEtcdPhase())
	initRunner.AppendPhase(phases.NewWaitControlPlanePhase())
	initRunner.AppendPhase(phases.NewUploadConfigPhase())
	initRunner.AppendPhase(phases.NewMarkControlPlanePhase())
	initRunner.AppendPhase(phases.NewBootstrapTokenPhase())
	initRunner.AppendPhase(phases.NewAddonPhase())
	initRunner.SetDataInitializer(func(cmd *cobra.Command) (workflow.RunData, error) {
		return newInitData(cmd, initOptions, out)
	})
	initRunner.BindToCommand(cmd)
	return cmd
}
func AddInitConfigFlags(flagSet *flag.FlagSet, cfg *kubeadmapiv1beta1.InitConfiguration, featureGatesString *string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flagSet.StringVar(&cfg.LocalAPIEndpoint.AdvertiseAddress, options.APIServerAdvertiseAddress, cfg.LocalAPIEndpoint.AdvertiseAddress, "The IP address the API Server will advertise it's listening on. Specify '0.0.0.0' to use the address of the default network interface.")
	flagSet.Int32Var(&cfg.LocalAPIEndpoint.BindPort, options.APIServerBindPort, cfg.LocalAPIEndpoint.BindPort, "Port for the API Server to bind to.")
	flagSet.StringVar(&cfg.Networking.ServiceSubnet, options.NetworkingServiceSubnet, cfg.Networking.ServiceSubnet, "Use alternative range of IP address for service VIPs.")
	flagSet.StringVar(&cfg.Networking.PodSubnet, options.NetworkingPodSubnet, cfg.Networking.PodSubnet, "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	flagSet.StringVar(&cfg.Networking.DNSDomain, options.NetworkingDNSDomain, cfg.Networking.DNSDomain, `Use alternative domain for services, e.g. "myorg.internal".`)
	flagSet.StringVar(&cfg.KubernetesVersion, options.KubernetesVersion, cfg.KubernetesVersion, `Choose a specific Kubernetes version for the control plane.`)
	flagSet.StringVar(&cfg.CertificatesDir, options.CertificatesDir, cfg.CertificatesDir, `The path where to save and store the certificates.`)
	flagSet.StringSliceVar(&cfg.APIServer.CertSANs, options.APIServerCertSANs, cfg.APIServer.CertSANs, `Optional extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names.`)
	flagSet.StringVar(&cfg.NodeRegistration.Name, options.NodeName, cfg.NodeRegistration.Name, `Specify the node name.`)
	flagSet.StringVar(&cfg.NodeRegistration.CRISocket, options.NodeCRISocket, cfg.NodeRegistration.CRISocket, `Specify the CRI socket to connect to.`)
	flagSet.StringVar(featureGatesString, options.FeatureGatesString, *featureGatesString, "A set of key=value pairs that describe feature gates for various features. "+"Options are:\n"+strings.Join(features.KnownFeatures(&features.InitFeatureGates), "\n"))
}
func AddInitOtherFlags(flagSet *flag.FlagSet, cfgPath *string, skipTokenPrint, dryRun *bool, ignorePreflightErrors *[]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flagSet.StringVar(cfgPath, options.CfgPath, *cfgPath, "Path to kubeadm config file. WARNING: Usage of a configuration file is experimental.")
	flagSet.StringSliceVar(ignorePreflightErrors, options.IgnorePreflightErrors, *ignorePreflightErrors, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	flagSet.BoolVar(skipTokenPrint, options.SkipTokenPrint, *skipTokenPrint, "Skip printing of the default bootstrap token generated by 'kubeadm init'.")
	flagSet.BoolVar(dryRun, options.DryRun, *dryRun, "Don't apply any changes; just output what would be done.")
}
func newInitOptions() *initOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	externalcfg := &kubeadmapiv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(externalcfg)
	bto := options.NewBootstrapTokenOptions()
	bto.Description = "The default bootstrap token generated by 'kubeadm init'."
	return &initOptions{externalcfg: externalcfg, bto: bto, kubeconfigDir: kubeadmconstants.KubernetesDir, kubeconfigPath: kubeadmconstants.GetAdminKubeConfigPath()}
}
func newInitData(cmd *cobra.Command, options *initOptions, out io.Writer) (initData, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeadmscheme.Scheme.Default(options.externalcfg)
	var err error
	if options.externalcfg.FeatureGates, err = features.NewFeatureGate(&features.InitFeatureGates, options.featureGatesString); err != nil {
		return initData{}, err
	}
	ignorePreflightErrorsSet, err := validation.ValidateIgnorePreflightErrors(options.ignorePreflightErrors)
	kubeadmutil.CheckErr(err)
	if err = validation.ValidateMixedArguments(cmd.Flags()); err != nil {
		return initData{}, err
	}
	if err = options.bto.ApplyTo(options.externalcfg); err != nil {
		return initData{}, err
	}
	cfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(options.cfgPath, options.externalcfg)
	if err != nil {
		return initData{}, err
	}
	if options.externalcfg.NodeRegistration.Name != "" {
		cfg.NodeRegistration.Name = options.externalcfg.NodeRegistration.Name
	}
	if options.externalcfg.NodeRegistration.CRISocket != kubeadmapiv1beta1.DefaultCRISocket {
		cfg.NodeRegistration.CRISocket = options.externalcfg.NodeRegistration.CRISocket
	}
	if err := configutil.VerifyAPIServerBindAddress(cfg.LocalAPIEndpoint.AdvertiseAddress); err != nil {
		return initData{}, err
	}
	if err := features.ValidateVersion(features.InitFeatureGates, cfg.FeatureGates, cfg.KubernetesVersion); err != nil {
		return initData{}, err
	}
	dryRunDir := ""
	if options.dryRun {
		if dryRunDir, err = ioutil.TempDir("", "kubeadm-init-dryrun"); err != nil {
			return initData{}, errors.Wrap(err, "couldn't create a temporary directory")
		}
	}
	externalCA, _ := certsphase.UsingExternalCA(cfg)
	if externalCA {
		kubeconfigDir := kubeadmconstants.KubernetesDir
		if options.dryRun {
			kubeconfigDir = dryRunDir
		}
		if err := kubeconfigphase.ValidateKubeconfigsForExternalCA(kubeconfigDir, cfg); err != nil {
			return initData{}, err
		}
	}
	return initData{cfg: cfg, certificatesDir: cfg.CertificatesDir, skipTokenPrint: options.skipTokenPrint, dryRun: options.dryRun, dryRunDir: dryRunDir, kubeconfigDir: options.kubeconfigDir, kubeconfigPath: options.kubeconfigPath, ignorePreflightErrors: ignorePreflightErrorsSet, externalCA: externalCA, outputWriter: out}, nil
}
func (d initData) Cfg() *kubeadmapi.InitConfiguration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.cfg
}
func (d initData) DryRun() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.dryRun
}
func (d initData) SkipTokenPrint() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.skipTokenPrint
}
func (d initData) IgnorePreflightErrors() sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.ignorePreflightErrors
}
func (d initData) CertificateWriteDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.dryRun {
		return d.dryRunDir
	}
	return d.certificatesDir
}
func (d initData) CertificateDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.certificatesDir
}
func (d initData) KubeConfigDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.dryRun {
		return d.dryRunDir
	}
	return d.kubeconfigDir
}
func (d initData) KubeConfigPath() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.dryRun {
		d.kubeconfigPath = filepath.Join(d.dryRunDir, kubeadmconstants.AdminKubeConfigFileName)
	}
	return d.kubeconfigPath
}
func (d initData) ManifestDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.dryRun {
		return d.dryRunDir
	}
	return kubeadmconstants.GetStaticPodDirectory()
}
func (d initData) KubeletDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.dryRun {
		return d.dryRunDir
	}
	return kubeadmconstants.KubeletRunDirectory
}
func (d initData) ExternalCA() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.externalCA
}
func (d initData) OutputWriter() io.Writer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.outputWriter
}
func (d initData) Client() (clientset.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.client == nil {
		if d.dryRun {
			dryRunGetter := apiclient.NewInitDryRunGetter(d.cfg.NodeRegistration.Name, d.cfg.Networking.ServiceSubnet)
			d.client = apiclient.NewDryRunClient(dryRunGetter, os.Stdout)
		} else {
			var err error
			d.client, err = kubeconfigutil.ClientSetFromFile(d.KubeConfigPath())
			if err != nil {
				return nil, err
			}
		}
	}
	return d.client, nil
}
func (d initData) Tokens() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokens := []string{}
	for _, bt := range d.cfg.BootstrapTokens {
		tokens = append(tokens, bt.Token.String())
	}
	return tokens
}
func printJoinCommand(out io.Writer, adminKubeConfigPath, token string, skipTokenPrint bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	joinCommand, err := cmdutil.GetJoinCommand(adminKubeConfigPath, token, skipTokenPrint)
	if err != nil {
		return err
	}
	ctx := map[string]string{"KubeConfigPath": adminKubeConfigPath, "joinCommand": joinCommand}
	return initDoneTempl.Execute(out, ctx)
}
func getDirectoriesToUse(dryRun bool, dryRunDir string, defaultPkiDir string) (string, string, string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		return dryRunDir, dryRunDir, dryRunDir, dryRunDir, nil
	}
	return defaultPkiDir, kubeadmconstants.KubernetesDir, kubeadmconstants.GetStaticPodDirectory(), kubeadmconstants.KubeletRunDirectory, nil
}
func showJoinCommand(i *initData, out io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	adminKubeConfigPath := i.KubeConfigPath()
	for _, token := range i.Tokens() {
		if err := printJoinCommand(out, adminKubeConfigPath, token, i.skipTokenPrint); err != nil {
			return errors.Wrap(err, "failed to print join command")
		}
	}
	return nil
}
