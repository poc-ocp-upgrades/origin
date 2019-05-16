package util

import (
	goformat "fmt"
	pkgerrors "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/errors"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	utilsexec "k8s.io/utils/exec"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	goruntime "runtime"
	"strings"
	gotime "time"
)

type ContainerRuntime interface {
	IsDocker() bool
	IsRunning() error
	ListKubeContainers() ([]string, error)
	RemoveContainers(containers []string) error
	PullImage(image string) error
	ImageExists(image string) (bool, error)
}
type CRIRuntime struct {
	exec      utilsexec.Interface
	criSocket string
}
type DockerRuntime struct{ exec utilsexec.Interface }

func NewContainerRuntime(execer utilsexec.Interface, criSocket string) (ContainerRuntime, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var toolName string
	var runtime ContainerRuntime
	if criSocket != kubeadmapiv1beta1.DefaultCRISocket {
		toolName = "crictl"
		if filepath.IsAbs(criSocket) && goruntime.GOOS != "windows" {
			criSocket = "unix://" + criSocket
		}
		runtime = &CRIRuntime{execer, criSocket}
	} else {
		toolName = "docker"
		runtime = &DockerRuntime{execer}
	}
	if _, err := execer.LookPath(toolName); err != nil {
		return nil, pkgerrors.Wrapf(err, "%s is required for container runtime", toolName)
	}
	return runtime, nil
}
func (runtime *CRIRuntime) IsDocker() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (runtime *DockerRuntime) IsDocker() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (runtime *CRIRuntime) IsRunning() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if out, err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "info").CombinedOutput(); err != nil {
		return pkgerrors.Wrapf(err, "container runtime is not running: output: %s, error", string(out))
	}
	return nil
}
func (runtime *DockerRuntime) IsRunning() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if out, err := runtime.exec.Command("docker", "info").CombinedOutput(); err != nil {
		return pkgerrors.Wrapf(err, "container runtime is not running: output: %s, error", string(out))
	}
	return nil
}
func (runtime *CRIRuntime) ListKubeContainers() ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out, err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "pods", "-q").CombinedOutput()
	if err != nil {
		return nil, pkgerrors.Wrapf(err, "output: %s, error", string(out))
	}
	pods := []string{}
	for _, pod := range strings.Fields(string(out)) {
		pods = append(pods, pod)
	}
	return pods, nil
}
func (runtime *DockerRuntime) ListKubeContainers() ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	output, err := runtime.exec.Command("docker", "ps", "-a", "--filter", "name=k8s_", "-q").CombinedOutput()
	return strings.Fields(string(output)), err
}
func (runtime *CRIRuntime) RemoveContainers(containers []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	for _, container := range containers {
		out, err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "stopp", container).CombinedOutput()
		if err != nil {
			errs = append(errs, pkgerrors.Wrapf(err, "failed to stop running pod %s: output: %s, error", container, string(out)))
		} else {
			out, err = runtime.exec.Command("crictl", "-r", runtime.criSocket, "rmp", container).CombinedOutput()
			if err != nil {
				errs = append(errs, pkgerrors.Wrapf(err, "failed to remove running container %s: output: %s, error", container, string(out)))
			}
		}
	}
	return errors.NewAggregate(errs)
}
func (runtime *DockerRuntime) RemoveContainers(containers []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	for _, container := range containers {
		out, err := runtime.exec.Command("docker", "rm", "--force", "--volumes", container).CombinedOutput()
		if err != nil {
			errs = append(errs, pkgerrors.Wrapf(err, "failed to remove running container %s: output: %s, error", container, string(out)))
		}
	}
	return errors.NewAggregate(errs)
}
func (runtime *CRIRuntime) PullImage(image string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out, err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "pull", image).CombinedOutput()
	if err != nil {
		return pkgerrors.Wrapf(err, "output: %s, error", string(out))
	}
	return nil
}
func (runtime *DockerRuntime) PullImage(image string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out, err := runtime.exec.Command("docker", "pull", image).CombinedOutput()
	if err != nil {
		return pkgerrors.Wrapf(err, "output: %s, error", string(out))
	}
	return nil
}
func (runtime *CRIRuntime) ImageExists(image string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "inspecti", image).Run()
	return err == nil, nil
}
func (runtime *DockerRuntime) ImageExists(image string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := runtime.exec.Command("docker", "inspect", image).Run()
	return err == nil, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
