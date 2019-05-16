package selfhosting

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	selfHostedKubeConfigDir = "/etc/kubernetes/kubeconfig"
)

type PodSpecMutatorFunc func(*v1.PodSpec)

func GetDefaultMutators() map[string][]PodSpecMutatorFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return map[string][]PodSpecMutatorFunc{kubeadmconstants.KubeAPIServer: {addNodeSelectorToPodSpec, setMasterTolerationOnPodSpec, setRightDNSPolicyOnPodSpec, setHostIPOnPodSpec}, kubeadmconstants.KubeControllerManager: {addNodeSelectorToPodSpec, setMasterTolerationOnPodSpec, setRightDNSPolicyOnPodSpec}, kubeadmconstants.KubeScheduler: {addNodeSelectorToPodSpec, setMasterTolerationOnPodSpec, setRightDNSPolicyOnPodSpec}}
}
func GetMutatorsFromFeatureGates(certsInSecrets bool) map[string][]PodSpecMutatorFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mutators := GetDefaultMutators()
	if certsInSecrets {
		mutators[kubeadmconstants.KubeAPIServer] = append(mutators[kubeadmconstants.KubeAPIServer], setSelfHostedVolumesForAPIServer)
		mutators[kubeadmconstants.KubeControllerManager] = append(mutators[kubeadmconstants.KubeControllerManager], setSelfHostedVolumesForControllerManager)
		mutators[kubeadmconstants.KubeScheduler] = append(mutators[kubeadmconstants.KubeScheduler], setSelfHostedVolumesForScheduler)
	}
	return mutators
}
func mutatePodSpec(mutators map[string][]PodSpecMutatorFunc, name string, podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mutatorsForComponent := mutators[name]
	for _, mutateFunc := range mutatorsForComponent {
		mutateFunc(podSpec)
	}
}
func addNodeSelectorToPodSpec(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if podSpec.NodeSelector == nil {
		podSpec.NodeSelector = map[string]string{kubeadmconstants.LabelNodeRoleMaster: ""}
		return
	}
	podSpec.NodeSelector[kubeadmconstants.LabelNodeRoleMaster] = ""
}
func setMasterTolerationOnPodSpec(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if podSpec.Tolerations == nil {
		podSpec.Tolerations = []v1.Toleration{kubeadmconstants.MasterToleration}
		return
	}
	podSpec.Tolerations = append(podSpec.Tolerations, kubeadmconstants.MasterToleration)
}
func setHostIPOnPodSpec(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	envVar := v1.EnvVar{Name: "HOST_IP", ValueFrom: &v1.EnvVarSource{FieldRef: &v1.ObjectFieldSelector{FieldPath: "status.hostIP"}}}
	podSpec.Containers[0].Env = append(podSpec.Containers[0].Env, envVar)
	for i := range podSpec.Containers[0].Command {
		if strings.Contains(podSpec.Containers[0].Command[i], "advertise-address") {
			podSpec.Containers[0].Command[i] = "--advertise-address=$(HOST_IP)"
		}
	}
}
func setRightDNSPolicyOnPodSpec(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podSpec.DNSPolicy = v1.DNSClusterFirstWithHostNet
}
func setSelfHostedVolumesForAPIServer(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, v := range podSpec.Volumes {
		if v.Name == kubeadmconstants.KubeCertificatesVolumeName {
			podSpec.Volumes[i].VolumeSource = apiServerCertificatesVolumeSource()
		}
	}
}
func setSelfHostedVolumesForControllerManager(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, v := range podSpec.Volumes {
		if v.Name == kubeadmconstants.KubeCertificatesVolumeName {
			podSpec.Volumes[i].VolumeSource = controllerManagerCertificatesVolumeSource()
		} else if v.Name == kubeadmconstants.KubeConfigVolumeName {
			podSpec.Volumes[i].VolumeSource = kubeConfigVolumeSource(kubeadmconstants.ControllerManagerKubeConfigFileName)
		}
	}
	for i, vm := range podSpec.Containers[0].VolumeMounts {
		if vm.Name == kubeadmconstants.KubeConfigVolumeName {
			podSpec.Containers[0].VolumeMounts[i].MountPath = selfHostedKubeConfigDir
		}
	}
	podSpec.Containers[0].Command = kubeadmutil.ReplaceArgument(podSpec.Containers[0].Command, func(argMap map[string]string) map[string]string {
		argMap["kubeconfig"] = filepath.Join(selfHostedKubeConfigDir, kubeadmconstants.ControllerManagerKubeConfigFileName)
		return argMap
	})
}
func setSelfHostedVolumesForScheduler(podSpec *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, v := range podSpec.Volumes {
		if v.Name == kubeadmconstants.KubeConfigVolumeName {
			podSpec.Volumes[i].VolumeSource = kubeConfigVolumeSource(kubeadmconstants.SchedulerKubeConfigFileName)
		}
	}
	for i, vm := range podSpec.Containers[0].VolumeMounts {
		if vm.Name == kubeadmconstants.KubeConfigVolumeName {
			podSpec.Containers[0].VolumeMounts[i].MountPath = selfHostedKubeConfigDir
		}
	}
	podSpec.Containers[0].Command = kubeadmutil.ReplaceArgument(podSpec.Containers[0].Command, func(argMap map[string]string) map[string]string {
		argMap["kubeconfig"] = filepath.Join(selfHostedKubeConfigDir, kubeadmconstants.SchedulerKubeConfigFileName)
		return argMap
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
