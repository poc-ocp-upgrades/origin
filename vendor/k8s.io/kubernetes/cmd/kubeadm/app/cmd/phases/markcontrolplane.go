package phases

import (
	"github.com/pkg/errors"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	markcontrolplanephase "k8s.io/kubernetes/cmd/kubeadm/app/phases/markcontrolplane"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	markControlPlaneExample = normalizer.Examples(`
		# Applies control-plane label and taint to the current node, functionally equivalent to what executed by kubeadm init.
		kubeadm init phase mark-control-plane --config config.yml

		# Applies control-plane label and taint to a specific node
		kubeadm init phase mark-control-plane --node-name myNode
		`)
)

type markControlPlaneData interface {
	Cfg() *kubeadmapi.InitConfiguration
	Client() (clientset.Interface, error)
	DryRun() bool
}

func NewMarkControlPlanePhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "mark-control-plane", Short: "Mark a node as a control-plane", Example: markControlPlaneExample, InheritFlags: []string{options.NodeName, options.CfgPath}, Run: runMarkControlPlane}
}
func runMarkControlPlane(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(markControlPlaneData)
	if !ok {
		return errors.New("mark-control-plane phase invoked with an invalid data struct")
	}
	client, err := data.Client()
	if err != nil {
		return err
	}
	nodeRegistration := data.Cfg().NodeRegistration
	if err := markcontrolplanephase.MarkControlPlane(client, nodeRegistration.Name, nodeRegistration.Taints); err != nil {
		return err
	}
	return nil
}
