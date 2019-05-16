package kubelet

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"k8s.io/kubernetes/pkg/util/procfs"
	utilsexec "k8s.io/utils/exec"
	"os"
	"path/filepath"
	"strings"
)

type kubeletFlagsOpts struct {
	nodeRegOpts              *kubeadmapi.NodeRegistrationOptions
	featureGates             map[string]bool
	pauseImage               string
	registerTaintsUsingFlags bool
	execer                   utilsexec.Interface
	pidOfFunc                func(string) ([]int, error)
	defaultHostname          string
}

func WriteKubeletDynamicEnvFile(cfg *kubeadmapi.InitConfiguration, registerTaintsUsingFlags bool, kubeletDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hostName, err := nodeutil.GetHostname("")
	if err != nil {
		return err
	}
	flagOpts := kubeletFlagsOpts{nodeRegOpts: &cfg.NodeRegistration, featureGates: cfg.FeatureGates, pauseImage: images.GetPauseImage(&cfg.ClusterConfiguration), registerTaintsUsingFlags: registerTaintsUsingFlags, execer: utilsexec.New(), pidOfFunc: procfs.PidOf, defaultHostname: hostName}
	stringMap := buildKubeletArgMap(flagOpts)
	argList := kubeadmutil.BuildArgumentListFromMap(stringMap, cfg.NodeRegistration.KubeletExtraArgs)
	envFileContent := fmt.Sprintf("%s=%s\n", constants.KubeletEnvFileVariableName, strings.Join(argList, " "))
	return writeKubeletFlagBytesToDisk([]byte(envFileContent), kubeletDir)
}
func buildKubeletArgMap(opts kubeletFlagsOpts) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeletFlags := map[string]string{}
	if opts.nodeRegOpts.CRISocket == kubeadmapiv1beta1.DefaultCRISocket {
		kubeletFlags["network-plugin"] = "cni"
		driver, err := kubeadmutil.GetCgroupDriverDocker(opts.execer)
		if err != nil {
			klog.Warningf("cannot automatically assign a '--cgroup-driver' value when starting the Kubelet: %v\n", err)
		} else {
			kubeletFlags["cgroup-driver"] = driver
		}
		if opts.pauseImage != "" {
			kubeletFlags["pod-infra-container-image"] = opts.pauseImage
		}
	} else {
		kubeletFlags["container-runtime"] = "remote"
		kubeletFlags["container-runtime-endpoint"] = opts.nodeRegOpts.CRISocket
	}
	if opts.registerTaintsUsingFlags && opts.nodeRegOpts.Taints != nil && len(opts.nodeRegOpts.Taints) > 0 {
		taintStrs := []string{}
		for _, taint := range opts.nodeRegOpts.Taints {
			taintStrs = append(taintStrs, taint.ToString())
		}
		kubeletFlags["register-with-taints"] = strings.Join(taintStrs, ",")
	}
	if pids, _ := opts.pidOfFunc("systemd-resolved"); len(pids) > 0 {
		kubeletFlags["resolv-conf"] = "/run/systemd/resolve/resolv.conf"
	}
	if opts.nodeRegOpts.Name != "" && opts.nodeRegOpts.Name != opts.defaultHostname {
		klog.V(1).Infof("setting kubelet hostname-override to %q", opts.nodeRegOpts.Name)
		kubeletFlags["hostname-override"] = opts.nodeRegOpts.Name
	}
	return kubeletFlags
}
func writeKubeletFlagBytesToDisk(b []byte, kubeletDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeletEnvFilePath := filepath.Join(kubeletDir, constants.KubeletEnvFileName)
	fmt.Printf("[kubelet-start] Writing kubelet environment file with flags to file %q\n", kubeletEnvFilePath)
	if err := os.MkdirAll(kubeletDir, 0700); err != nil {
		return errors.Wrapf(err, "failed to create directory %q", kubeletDir)
	}
	if err := ioutil.WriteFile(kubeletEnvFilePath, b, 0644); err != nil {
		return errors.Wrapf(err, "failed to write kubelet configuration to the file %q", kubeletEnvFilePath)
	}
	return nil
}
