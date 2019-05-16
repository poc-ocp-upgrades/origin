package phases

import (
	"github.com/pkg/errors"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	kubeletphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubelet"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	kubeletStartPhaseExample = normalizer.Examples(`
		# Writes a dynamic environment file with kubelet flags from a InitConfiguration file.
		kubeadm init phase kubelet-start --config masterconfig.yaml
		`)
)

type kubeletStartData interface {
	Cfg() *kubeadmapi.InitConfiguration
	DryRun() bool
	KubeletDir() string
}

func NewKubeletStartPhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "kubelet-start", Short: "Writes kubelet settings and (re)starts the kubelet", Long: "Writes a file with KubeletConfiguration and an environment file with node specific kubelet settings, and then (re)starts kubelet.", Example: kubeletStartPhaseExample, Run: runKubeletStart, InheritFlags: []string{options.CfgPath, options.NodeCRISocket, options.NodeName}}
}
func runKubeletStart(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(kubeletStartData)
	if !ok {
		return errors.New("kubelet-start phase invoked with an invalid data struct")
	}
	if !data.DryRun() {
		klog.V(1).Infof("Stopping the kubelet")
		kubeletphase.TryStopKubelet()
	}
	if err := kubeletphase.WriteKubeletDynamicEnvFile(data.Cfg(), false, data.KubeletDir()); err != nil {
		return errors.Wrap(err, "error writing a dynamic environment file for the kubelet")
	}
	if err := kubeletphase.WriteConfigToDisk(data.Cfg().ComponentConfigs.Kubelet, data.KubeletDir()); err != nil {
		return errors.Wrap(err, "error writing kubelet configuration to disk")
	}
	if !data.DryRun() {
		klog.V(1).Infof("Starting the kubelet")
		kubeletphase.TryStartKubelet()
	}
	return nil
}
