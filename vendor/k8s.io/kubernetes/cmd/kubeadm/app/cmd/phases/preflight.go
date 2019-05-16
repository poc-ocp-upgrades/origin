package phases

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	"k8s.io/kubernetes/cmd/kubeadm/app/preflight"
	"k8s.io/kubernetes/pkg/util/normalizer"
	utilsexec "k8s.io/utils/exec"
)

var (
	masterPreflightExample = normalizer.Examples(`
		# Run master pre-flight checks using a config file.
		kubeadm init phase preflight --config kubeadm-config.yml
		`)
)

type preflightMasterData interface {
	Cfg() *kubeadmapi.InitConfiguration
	DryRun() bool
	IgnorePreflightErrors() sets.String
}

func NewPreflightMasterPhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "preflight", Short: "Run master pre-flight checks", Long: "Run master pre-flight checks, functionally equivalent to what implemented by kubeadm init.", Example: masterPreflightExample, Run: runPreflightMaster, InheritFlags: []string{options.CfgPath, options.IgnorePreflightErrors}}
}
func runPreflightMaster(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(preflightMasterData)
	if !ok {
		return errors.New("preflight phase invoked with an invalid data struct")
	}
	fmt.Println("[preflight] Running pre-flight checks")
	if err := preflight.RunInitMasterChecks(utilsexec.New(), data.Cfg(), data.IgnorePreflightErrors()); err != nil {
		return err
	}
	if !data.DryRun() {
		fmt.Println("[preflight] Pulling images required for setting up a Kubernetes cluster")
		fmt.Println("[preflight] This might take a minute or two, depending on the speed of your internet connection")
		fmt.Println("[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'")
		if err := preflight.RunPullImagesCheck(utilsexec.New(), data.Cfg(), data.IgnorePreflightErrors()); err != nil {
			return err
		}
	} else {
		fmt.Println("[preflight] Would pull the required images (like 'kubeadm config images pull')")
	}
	return nil
}
