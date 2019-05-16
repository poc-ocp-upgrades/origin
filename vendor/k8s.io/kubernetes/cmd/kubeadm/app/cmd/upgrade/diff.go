package upgrade

import (
	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
)

type diffFlags struct {
	apiServerManifestPath         string
	controllerManagerManifestPath string
	schedulerManifestPath         string
	newK8sVersionStr              string
	contextLines                  int
	cfgPath                       string
	out                           io.Writer
}

var (
	defaultAPIServerManifestPath         = constants.GetStaticPodFilepath(constants.KubeAPIServer, constants.GetStaticPodDirectory())
	defaultControllerManagerManifestPath = constants.GetStaticPodFilepath(constants.KubeControllerManager, constants.GetStaticPodDirectory())
	defaultSchedulerManifestPath         = constants.GetStaticPodFilepath(constants.KubeScheduler, constants.GetStaticPodDirectory())
)

func NewCmdDiff(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := &diffFlags{out: out}
	cmd := &cobra.Command{Use: "diff [version]", Short: "Show what differences would be applied to existing static pod manifests. See also: kubeadm upgrade apply --dry-run", Run: func(cmd *cobra.Command, args []string) {
		kubeadmutil.CheckErr(runDiff(flags, args))
	}}
	options.AddConfigFlag(cmd.Flags(), &flags.cfgPath)
	cmd.Flags().StringVar(&flags.apiServerManifestPath, "api-server-manifest", defaultAPIServerManifestPath, "path to API server manifest")
	cmd.Flags().StringVar(&flags.controllerManagerManifestPath, "controller-manager-manifest", defaultControllerManagerManifestPath, "path to controller manifest")
	cmd.Flags().StringVar(&flags.schedulerManifestPath, "scheduler-manifest", defaultSchedulerManifestPath, "path to scheduler manifest")
	cmd.Flags().IntVarP(&flags.contextLines, "context-lines", "c", 3, "How many lines of context in the diff")
	return cmd
}
func runDiff(flags *diffFlags, args []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("fetching configuration from file %s", flags.cfgPath)
	cfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(flags.cfgPath, &kubeadmapiv1beta1.InitConfiguration{})
	if err != nil {
		return err
	}
	if cfg.KubernetesVersion != "" {
		flags.newK8sVersionStr = cfg.KubernetesVersion
	}
	if flags.newK8sVersionStr == "" {
		if err := cmdutil.ValidateExactArgNumber(args, []string{"version"}); err != nil {
			return err
		}
	}
	if len(args) == 1 {
		flags.newK8sVersionStr = args[0]
	}
	k8sVer, err := version.ParseSemantic(flags.newK8sVersionStr)
	if err != nil {
		return err
	}
	specs := controlplane.GetStaticPodSpecs(cfg, k8sVer)
	for spec, pod := range specs {
		var path string
		switch spec {
		case constants.KubeAPIServer:
			path = flags.apiServerManifestPath
		case constants.KubeControllerManager:
			path = flags.controllerManagerManifestPath
		case constants.KubeScheduler:
			path = flags.schedulerManifestPath
		default:
			klog.Errorf("[diff] unknown spec %v", spec)
			continue
		}
		newManifest, err := kubeadmutil.MarshalToYaml(&pod, corev1.SchemeGroupVersion)
		if err != nil {
			return err
		}
		if path == "" {
			return errors.New("empty manifest path")
		}
		existingManifest, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		diff := difflib.UnifiedDiff{A: difflib.SplitLines(string(existingManifest)), B: difflib.SplitLines(string(newManifest)), FromFile: path, ToFile: "new manifest", Context: flags.contextLines}
		difflib.WriteUnifiedDiff(flags.out, diff)
	}
	return nil
}
