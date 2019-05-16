package system

import (
	"os/exec"
	"strings"
)

const dockerEndpoint = "unix:///var/run/docker.sock"

var DefaultSysSpec = SysSpec{OS: "Linux", KernelSpec: KernelSpec{Versions: []string{`3\.[1-9][0-9].*`, `4\..*`}, Required: []KernelConfig{{Name: "NAMESPACES"}, {Name: "NET_NS"}, {Name: "PID_NS"}, {Name: "IPC_NS"}, {Name: "UTS_NS"}, {Name: "CGROUPS"}, {Name: "CGROUP_CPUACCT"}, {Name: "CGROUP_DEVICE"}, {Name: "CGROUP_FREEZER"}, {Name: "CGROUP_SCHED"}, {Name: "CPUSETS"}, {Name: "MEMCG"}, {Name: "INET"}, {Name: "EXT4_FS"}, {Name: "PROC_FS"}, {Name: "NETFILTER_XT_TARGET_REDIRECT", Aliases: []string{"IP_NF_TARGET_REDIRECT"}}, {Name: "NETFILTER_XT_MATCH_COMMENT"}}, Optional: []KernelConfig{{Name: "OVERLAY_FS", Aliases: []string{"OVERLAYFS_FS"}, Description: "Required for overlayfs."}, {Name: "AUFS_FS", Description: "Required for aufs."}, {Name: "BLK_DEV_DM", Description: "Required for devicemapper."}}, Forbidden: []KernelConfig{}}, Cgroups: []string{"cpu", "cpuacct", "cpuset", "devices", "freezer", "memory"}, RuntimeSpec: RuntimeSpec{DockerSpec: &DockerSpec{Version: []string{`1\.1[1-3]\..*`, `17\.0[3,6,9]\..*`, `18\.06\..*`}, GraphDriver: []string{"aufs", "overlay", "overlay2", "devicemapper", "zfs"}}}}

type KernelValidatorHelperImpl struct{}

var _ KernelValidatorHelper = &KernelValidatorHelperImpl{}

func (o *KernelValidatorHelperImpl) GetKernelReleaseVersion() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	releaseVersion, err := exec.Command("uname", "-r").CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(releaseVersion)), nil
}
