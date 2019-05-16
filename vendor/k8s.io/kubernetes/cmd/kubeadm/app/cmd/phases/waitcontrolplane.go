package phases

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/renstrom/dedent"
	"io"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	dryrunutil "k8s.io/kubernetes/cmd/kubeadm/app/util/dryrun"
	"path/filepath"
	"text/template"
	"time"
)

var (
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

type waitControlPlaneData interface {
	Cfg() *kubeadmapi.InitConfiguration
	ManifestDir() string
	DryRun() bool
	Client() (clientset.Interface, error)
	OutputWriter() io.Writer
}

func NewWaitControlPlanePhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	phase := workflow.Phase{Name: "wait-control-plane", Run: runWaitControlPlanePhase, Hidden: true}
	return phase
}
func runWaitControlPlanePhase(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(waitControlPlaneData)
	if !ok {
		return errors.New("wait-control-plane phase invoked with an invalid data struct")
	}
	if err := printFilesIfDryRunning(data); err != nil {
		return errors.Wrap(err, "error printing files on dryrun")
	}
	klog.V(1).Infof("[wait-control-plane] Waiting for the API server to be healthy")
	client, err := data.Client()
	if err != nil {
		return errors.Wrap(err, "cannot obtain client")
	}
	timeout := data.Cfg().ClusterConfiguration.APIServer.TimeoutForControlPlane.Duration
	waiter, err := newControlPlaneWaiter(data.DryRun(), timeout, client, data.OutputWriter())
	if err != nil {
		return errors.Wrap(err, "error creating waiter")
	}
	fmt.Printf("[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory %q. This can take up to %v\n", data.ManifestDir(), timeout)
	if err := waiter.WaitForKubeletAndFunc(waiter.WaitForAPI); err != nil {
		ctx := map[string]string{"Error": fmt.Sprintf("%v", err)}
		kubeletFailTempl.Execute(data.OutputWriter(), ctx)
		return errors.New("couldn't initialize a Kubernetes cluster")
	}
	return nil
}
func printFilesIfDryRunning(data waitControlPlaneData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !data.DryRun() {
		return nil
	}
	manifestDir := data.ManifestDir()
	fmt.Printf("[dryrun] Wrote certificates, kubeconfig files and control plane manifests to the %q directory\n", manifestDir)
	fmt.Println("[dryrun] The certificates or kubeconfig files would not be printed due to their sensitive nature")
	fmt.Printf("[dryrun] Please examine the %q directory for details about what would be written\n", manifestDir)
	files := []dryrunutil.FileToPrint{}
	for _, component := range kubeadmconstants.MasterComponents {
		realPath := kubeadmconstants.GetStaticPodFilepath(component, manifestDir)
		outputPath := kubeadmconstants.GetStaticPodFilepath(component, kubeadmconstants.GetStaticPodDirectory())
		files = append(files, dryrunutil.NewFileToPrint(realPath, outputPath))
	}
	kubeletConfigFiles := []string{kubeadmconstants.KubeletConfigurationFileName, kubeadmconstants.KubeletEnvFileName}
	for _, filename := range kubeletConfigFiles {
		realPath := filepath.Join(manifestDir, filename)
		outputPath := filepath.Join(kubeadmconstants.KubeletRunDirectory, filename)
		files = append(files, dryrunutil.NewFileToPrint(realPath, outputPath))
	}
	return dryrunutil.PrintDryRunFiles(files, data.OutputWriter())
}
func newControlPlaneWaiter(dryRun bool, timeout time.Duration, client clientset.Interface, out io.Writer) (apiclient.Waiter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		return dryrunutil.NewWaiter(), nil
	}
	return apiclient.NewKubeWaiter(client, timeout, out), nil
}
