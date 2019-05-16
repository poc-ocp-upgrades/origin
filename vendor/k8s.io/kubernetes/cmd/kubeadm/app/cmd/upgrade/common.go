package upgrade

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	fakediscovery "k8s.io/client-go/discovery/fake"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/upgrade"
	"k8s.io/kubernetes/cmd/kubeadm/app/preflight"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	dryrunutil "k8s.io/kubernetes/cmd/kubeadm/app/util/dryrun"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"os"
	"strings"
)

type upgradeVariables struct {
	client        clientset.Interface
	cfg           *kubeadmapi.InitConfiguration
	versionGetter upgrade.VersionGetter
	waiter        apiclient.Waiter
}

func enforceRequirements(flags *applyPlanFlags, dryRun bool, newK8sVersion string) (*upgradeVariables, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client, err := getClient(flags.kubeConfigPath, dryRun)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't create a Kubernetes client from file %q", flags.kubeConfigPath)
	}
	if upgrade.IsControlPlaneSelfHosted(client) {
		return nil, errors.Errorf("cannot upgrade a self-hosted control plane")
	}
	if err := upgrade.CheckClusterHealth(client, flags.ignorePreflightErrorsSet); err != nil {
		return nil, errors.Wrap(err, "[upgrade/health] FATAL")
	}
	fmt.Println("[upgrade/config] Making sure the configuration is correct:")
	cfg, err := configutil.FetchConfigFromFileOrCluster(client, os.Stdout, "upgrade/config", flags.cfgPath, false)
	if err != nil {
		if apierrors.IsNotFound(err) {
			fmt.Printf("[upgrade/config] In order to upgrade, a ConfigMap called %q in the %s namespace must exist.\n", constants.KubeadmConfigConfigMap, metav1.NamespaceSystem)
			fmt.Println("[upgrade/config] Without this information, 'kubeadm upgrade' won't know how to configure your upgraded cluster.")
			fmt.Println("")
			fmt.Println("[upgrade/config] Next steps:")
			fmt.Printf("\t- OPTION 1: Run 'kubeadm config upload from-flags' and specify the same CLI arguments you passed to 'kubeadm init' when you created your master.\n")
			fmt.Printf("\t- OPTION 2: Run 'kubeadm config upload from-file' and specify the same config file you passed to 'kubeadm init' when you created your master.\n")
			fmt.Printf("\t- OPTION 3: Pass a config file to 'kubeadm upgrade' using the --config flag.\n")
			fmt.Println("")
			err = errors.Errorf("the ConfigMap %q in the %s namespace used for getting configuration information was not found", constants.KubeadmConfigConfigMap, metav1.NamespaceSystem)
		}
		return nil, errors.Wrap(err, "[upgrade/config] FATAL")
	}
	if len(newK8sVersion) != 0 {
		cfg.KubernetesVersion = newK8sVersion
	}
	if flags.featureGatesString != "" {
		cfg.FeatureGates, err = features.NewFeatureGate(&features.InitFeatureGates, flags.featureGatesString)
		if err != nil {
			return nil, errors.Wrap(err, "[upgrade/config] FATAL")
		}
	}
	if msg := features.CheckDeprecatedFlags(&features.InitFeatureGates, cfg.FeatureGates); len(msg) > 0 {
		for _, m := range msg {
			fmt.Printf("[upgrade/config] %s\n", m)
		}
		return nil, errors.New("[upgrade/config] FATAL. Unable to upgrade a cluster using deprecated feature-gate flags. Please see the release notes")
	}
	if flags.printConfig {
		printConfiguration(&cfg.ClusterConfiguration, os.Stdout)
	}
	return &upgradeVariables{client: client, cfg: cfg, versionGetter: upgrade.NewOfflineVersionGetter(upgrade.NewKubeVersionGetter(client, os.Stdout), newK8sVersion), waiter: getWaiter(dryRun, client)}, nil
}
func printConfiguration(clustercfg *kubeadmapi.ClusterConfiguration, w io.Writer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if clustercfg == nil {
		return
	}
	cfgYaml, err := configutil.MarshalKubeadmConfigObject(clustercfg)
	if err == nil {
		fmt.Fprintln(w, "[upgrade/config] Configuration used:")
		scanner := bufio.NewScanner(bytes.NewReader(cfgYaml))
		for scanner.Scan() {
			fmt.Fprintf(w, "\t%s\n", scanner.Text())
		}
	}
}
func runPreflightChecks(ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[preflight] Running pre-flight checks.")
	return preflight.RunRootCheckOnly(ignorePreflightErrors)
}
func getClient(file string, dryRun bool) (clientset.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		dryRunGetter, err := apiclient.NewClientBackedDryRunGetterFromKubeconfig(file)
		if err != nil {
			return nil, err
		}
		realServerVersion, err := dryRunGetter.Client().Discovery().ServerVersion()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get server version")
		}
		dryRunOpts := apiclient.GetDefaultDryRunClientOptions(dryRunGetter, os.Stdout)
		dryRunOpts.PrintGETAndLIST = true
		fakeclient := apiclient.NewDryRunClientWithOpts(dryRunOpts)
		fakeclientDiscovery, ok := fakeclient.Discovery().(*fakediscovery.FakeDiscovery)
		if !ok {
			return nil, errors.New("couldn't set fake discovery's server version")
		}
		fakeclientDiscovery.FakedServerVersion = realServerVersion
		return fakeclient, nil
	}
	return kubeconfigutil.ClientSetFromFile(file)
}
func getWaiter(dryRun bool, client clientset.Interface) apiclient.Waiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		return dryrunutil.NewWaiter()
	}
	return apiclient.NewKubeWaiter(client, upgrade.UpgradeManifestTimeout, os.Stdout)
}
func InteractivelyConfirmUpgrade(question string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[upgrade/confirm] %s [y/N]: ", question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "couldn't read from standard input")
	}
	answer := scanner.Text()
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		return nil
	}
	return errors.New("won't proceed; the user didn't answer (Y|y) in order to continue")
}
