package upgrade

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/upgrade"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	dryrunutil "k8s.io/kubernetes/cmd/kubeadm/app/util/dryrun"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	defaultImagePullTimeout = 15 * time.Minute
)

type applyFlags struct {
	*applyPlanFlags
	nonInteractiveMode bool
	force              bool
	dryRun             bool
	etcdUpgrade        bool
	criSocket          string
	newK8sVersionStr   string
	newK8sVersion      *version.Version
	imagePullTimeout   time.Duration
}

func (f *applyFlags) SessionIsInteractive() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !f.nonInteractiveMode
}
func NewCmdApply(apf *applyPlanFlags) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := &applyFlags{applyPlanFlags: apf, imagePullTimeout: defaultImagePullTimeout, etcdUpgrade: true, criSocket: kubeadmapiv1beta1.DefaultCRISocket}
	cmd := &cobra.Command{Use: "apply [version]", DisableFlagsInUseLine: true, Short: "Upgrade your Kubernetes cluster to the specified version.", Run: func(cmd *cobra.Command, args []string) {
		var err error
		flags.ignorePreflightErrorsSet, err = validation.ValidateIgnorePreflightErrors(flags.ignorePreflightErrors)
		kubeadmutil.CheckErr(err)
		klog.V(1).Infof("running preflight checks")
		err = runPreflightChecks(flags.ignorePreflightErrorsSet)
		kubeadmutil.CheckErr(err)
		if flags.cfgPath != "" {
			klog.V(1).Infof("fetching configuration from file %s", flags.cfgPath)
			cfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(flags.cfgPath, &kubeadmapiv1beta1.InitConfiguration{})
			kubeadmutil.CheckErr(err)
			if cfg.KubernetesVersion != "" {
				flags.newK8sVersionStr = cfg.KubernetesVersion
			}
		}
		if flags.newK8sVersionStr == "" {
			err = cmdutil.ValidateExactArgNumber(args, []string{"version"})
			kubeadmutil.CheckErr(err)
		}
		if len(args) == 1 {
			flags.newK8sVersionStr = args[0]
		}
		err = SetImplicitFlags(flags)
		kubeadmutil.CheckErr(err)
		err = RunApply(flags)
		kubeadmutil.CheckErr(err)
	}}
	addApplyPlanFlags(cmd.Flags(), flags.applyPlanFlags)
	cmd.Flags().BoolVarP(&flags.nonInteractiveMode, "yes", "y", flags.nonInteractiveMode, "Perform the upgrade and do not prompt for confirmation (non-interactive mode).")
	cmd.Flags().BoolVarP(&flags.force, "force", "f", flags.force, "Force upgrading although some requirements might not be met. This also implies non-interactive mode.")
	cmd.Flags().BoolVar(&flags.dryRun, "dry-run", flags.dryRun, "Do not change any state, just output what actions would be performed.")
	cmd.Flags().BoolVar(&flags.etcdUpgrade, "etcd-upgrade", flags.etcdUpgrade, "Perform the upgrade of etcd.")
	cmd.Flags().DurationVar(&flags.imagePullTimeout, "image-pull-timeout", flags.imagePullTimeout, "The maximum amount of time to wait for the control plane pods to be downloaded.")
	cmd.Flags().StringVar(&flags.criSocket, "cri-socket", flags.criSocket, "Specify the CRI socket to connect to.")
	return cmd
}
func RunApply(flags *applyFlags) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("[upgrade/apply] verifying health of cluster")
	klog.V(1).Infof("[upgrade/apply] retrieving configuration from cluster")
	upgradeVars, err := enforceRequirements(flags.applyPlanFlags, flags.dryRun, flags.newK8sVersionStr)
	if err != nil {
		return err
	}
	if len(flags.criSocket) != 0 {
		fmt.Println("[upgrade/apply] Respecting the --cri-socket flag that is set with higher priority than the config file.")
		upgradeVars.cfg.NodeRegistration.CRISocket = flags.criSocket
	}
	klog.V(1).Infof("[upgrade/apply] validating requested and actual version")
	if err := configutil.NormalizeKubernetesVersion(&upgradeVars.cfg.ClusterConfiguration); err != nil {
		return err
	}
	flags.newK8sVersionStr = upgradeVars.cfg.KubernetesVersion
	k8sVer, err := version.ParseSemantic(flags.newK8sVersionStr)
	if err != nil {
		return errors.Errorf("unable to parse normalized version %q as a semantic version", flags.newK8sVersionStr)
	}
	flags.newK8sVersion = k8sVer
	if err := features.ValidateVersion(features.InitFeatureGates, upgradeVars.cfg.FeatureGates, upgradeVars.cfg.KubernetesVersion); err != nil {
		return err
	}
	klog.V(1).Infof("[upgrade/version] enforcing version skew policies")
	if err := EnforceVersionPolicies(flags, upgradeVars.versionGetter); err != nil {
		return errors.Wrap(err, "[upgrade/version] FATAL")
	}
	if flags.SessionIsInteractive() {
		if err := InteractivelyConfirmUpgrade("Are you sure you want to proceed with the upgrade?"); err != nil {
			return err
		}
	}
	klog.V(1).Infof("[upgrade/apply] creating prepuller")
	prepuller := upgrade.NewDaemonSetPrepuller(upgradeVars.client, upgradeVars.waiter, &upgradeVars.cfg.ClusterConfiguration)
	componentsToPrepull := constants.MasterComponents
	if upgradeVars.cfg.Etcd.External == nil && flags.etcdUpgrade {
		componentsToPrepull = append(componentsToPrepull, constants.Etcd)
	}
	if err := upgrade.PrepullImagesInParallel(prepuller, flags.imagePullTimeout, componentsToPrepull); err != nil {
		return errors.Wrap(err, "[upgrade/prepull] Failed prepulled the images for the control plane components error")
	}
	klog.V(1).Infof("[upgrade/apply] performing upgrade")
	if err := PerformControlPlaneUpgrade(flags, upgradeVars.client, upgradeVars.waiter, upgradeVars.cfg); err != nil {
		return errors.Wrap(err, "[upgrade/apply] FATAL")
	}
	klog.V(1).Infof("[upgrade/postupgrade] upgrading RBAC rules and addons")
	if err := upgrade.PerformPostUpgradeTasks(upgradeVars.client, upgradeVars.cfg, flags.newK8sVersion, flags.dryRun); err != nil {
		return errors.Wrap(err, "[upgrade/postupgrade] FATAL post-upgrade error")
	}
	if flags.dryRun {
		fmt.Println("[dryrun] Finished dryrunning successfully!")
		return nil
	}
	fmt.Println("")
	fmt.Printf("[upgrade/successful] SUCCESS! Your cluster was upgraded to %q. Enjoy!\n", flags.newK8sVersionStr)
	fmt.Println("")
	fmt.Println("[upgrade/kubelet] Now that your control plane is upgraded, please proceed with upgrading your kubelets if you haven't already done so.")
	return nil
}
func SetImplicitFlags(flags *applyFlags) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if flags.dryRun || flags.force {
		flags.nonInteractiveMode = true
	}
	if len(flags.newK8sVersionStr) == 0 {
		return errors.New("version string can't be empty")
	}
	return nil
}
func EnforceVersionPolicies(flags *applyFlags, versionGetter upgrade.VersionGetter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[upgrade/version] You have chosen to change the cluster version to %q\n", flags.newK8sVersionStr)
	versionSkewErrs := upgrade.EnforceVersionPolicies(versionGetter, flags.newK8sVersionStr, flags.newK8sVersion, flags.allowExperimentalUpgrades, flags.allowRCUpgrades)
	if versionSkewErrs != nil {
		if len(versionSkewErrs.Mandatory) > 0 {
			return errors.Errorf("the --version argument is invalid due to these fatal errors:\n\n%v\nPlease fix the misalignments highlighted above and try upgrading again", kubeadmutil.FormatErrMsg(versionSkewErrs.Mandatory))
		}
		if len(versionSkewErrs.Skippable) > 0 {
			if !flags.force {
				return errors.Errorf("the --version argument is invalid due to these errors:\n\n%v\nCan be bypassed if you pass the --force flag", kubeadmutil.FormatErrMsg(versionSkewErrs.Skippable))
			}
			fmt.Printf("[upgrade/version] Found %d potential version compatibility errors but skipping since the --force flag is set: \n\n%v", len(versionSkewErrs.Skippable), kubeadmutil.FormatErrMsg(versionSkewErrs.Skippable))
		}
	}
	return nil
}
func PerformControlPlaneUpgrade(flags *applyFlags, client clientset.Interface, waiter apiclient.Waiter, internalcfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[upgrade/apply] Upgrading your Static Pod-hosted control plane to version %q...\n", flags.newK8sVersionStr)
	if flags.dryRun {
		return DryRunStaticPodUpgrade(internalcfg)
	}
	return PerformStaticPodUpgrade(client, waiter, internalcfg, flags.etcdUpgrade)
}
func GetPathManagerForUpgrade(internalcfg *kubeadmapi.InitConfiguration, etcdUpgrade bool) (upgrade.StaticPodPathManager, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	isHAEtcd := etcdutil.CheckConfigurationIsHA(&internalcfg.Etcd)
	return upgrade.NewKubeStaticPodPathManagerUsingTempDirs(constants.GetStaticPodDirectory(), true, etcdUpgrade && !isHAEtcd)
}
func PerformStaticPodUpgrade(client clientset.Interface, waiter apiclient.Waiter, internalcfg *kubeadmapi.InitConfiguration, etcdUpgrade bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pathManager, err := GetPathManagerForUpgrade(internalcfg, etcdUpgrade)
	if err != nil {
		return err
	}
	return upgrade.StaticPodControlPlane(client, waiter, pathManager, internalcfg, etcdUpgrade, nil, nil)
}
func DryRunStaticPodUpgrade(internalcfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dryRunManifestDir, err := constants.CreateTempDirForKubeadm("kubeadm-upgrade-dryrun")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dryRunManifestDir)
	if err := controlplane.CreateInitStaticPodManifestFiles(dryRunManifestDir, internalcfg); err != nil {
		return err
	}
	files := []dryrunutil.FileToPrint{}
	for _, component := range constants.MasterComponents {
		realPath := constants.GetStaticPodFilepath(component, dryRunManifestDir)
		outputPath := constants.GetStaticPodFilepath(component, constants.GetStaticPodDirectory())
		files = append(files, dryrunutil.NewFileToPrint(realPath, outputPath))
	}
	return dryrunutil.PrintDryRunFiles(files, os.Stdout)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
