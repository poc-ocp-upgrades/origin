package phases

import (
	"fmt"
	"github.com/pkg/errors"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	kubeconfigFilePhaseProperties = map[string]struct {
		name  string
		short string
		long  string
	}{kubeadmconstants.AdminKubeConfigFileName: {name: "admin", short: "Generates a kubeconfig file for the admin to use and for kubeadm itself", long: "Generates the kubeconfig file for the admin and for kubeadm itself, and saves it to %s file."}, kubeadmconstants.KubeletKubeConfigFileName: {name: "kubelet", short: "Generates a kubeconfig file for the kubelet to use *only* for cluster bootstrapping purposes", long: normalizer.LongDesc(`
					Generates the kubeconfig file for the kubelet to use and saves it to %s file.
			
					Please note that this should *only* be used for cluster bootstrapping purposes. After your control plane is up,
					you should request all kubelet credentials from the CSR API.`)}, kubeadmconstants.ControllerManagerKubeConfigFileName: {name: "controller-manager", short: "Generates a kubeconfig file for the controller manager to use", long: "Generates the kubeconfig file for the controller manager to use and saves it to %s file"}, kubeadmconstants.SchedulerKubeConfigFileName: {name: "scheduler", short: "Generates a kubeconfig file for the scheduler to use", long: "Generates the kubeconfig file for the scheduler to use and saves it to %s file."}}
)

type kubeConfigData interface {
	Cfg() *kubeadmapi.InitConfiguration
	ExternalCA() bool
	CertificateDir() string
	CertificateWriteDir() string
	KubeConfigDir() string
}

func NewKubeConfigPhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "kubeconfig", Short: "Generates all kubeconfig files necessary to establish the control plane and the admin kubeconfig file", Long: cmdutil.MacroCommandLongDescription, Phases: []workflow.Phase{{Name: "all", Short: "Generates all kubeconfig files", InheritFlags: getKubeConfigPhaseFlags("all"), RunAllSiblings: true}, NewKubeConfigFilePhase(kubeadmconstants.AdminKubeConfigFileName), NewKubeConfigFilePhase(kubeadmconstants.KubeletKubeConfigFileName), NewKubeConfigFilePhase(kubeadmconstants.ControllerManagerKubeConfigFileName), NewKubeConfigFilePhase(kubeadmconstants.SchedulerKubeConfigFileName)}, Run: runKubeConfig}
}
func NewKubeConfigFilePhase(kubeConfigFileName string) workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: kubeconfigFilePhaseProperties[kubeConfigFileName].name, Short: kubeconfigFilePhaseProperties[kubeConfigFileName].short, Long: fmt.Sprintf(kubeconfigFilePhaseProperties[kubeConfigFileName].long, kubeConfigFileName), Run: runKubeConfigFile(kubeConfigFileName), InheritFlags: getKubeConfigPhaseFlags(kubeConfigFileName)}
}
func getKubeConfigPhaseFlags(name string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := []string{options.APIServerAdvertiseAddress, options.APIServerBindPort, options.CertificatesDir, options.CfgPath, options.KubeconfigDir}
	if name == "all" || name == kubeadmconstants.KubeletKubeConfigFileName {
		flags = append(flags, options.NodeName)
	}
	return flags
}
func runKubeConfig(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(kubeConfigData)
	if !ok {
		return errors.New("kubeconfig phase invoked with an invalid data struct")
	}
	fmt.Printf("[kubeconfig] Using kubeconfig folder %q\n", data.KubeConfigDir())
	return nil
}
func runKubeConfigFile(kubeConfigFileName string) func(workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(c workflow.RunData) error {
		data, ok := c.(kubeConfigData)
		if !ok {
			return errors.New("kubeconfig phase invoked with an invalid data struct")
		}
		if data.ExternalCA() {
			fmt.Printf("[kubeconfig] External CA mode: Using user provided %s\n", kubeConfigFileName)
			return nil
		}
		cfg := data.Cfg()
		cfg.CertificatesDir = data.CertificateWriteDir()
		defer func() {
			cfg.CertificatesDir = data.CertificateDir()
		}()
		return kubeconfigphase.CreateKubeConfigFile(kubeConfigFileName, data.KubeConfigDir(), data.Cfg())
	}
}
