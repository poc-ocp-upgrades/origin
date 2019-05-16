package system

import (
	"os/exec"
	"strings"
)

const dockerEndpoint = "npipe:////./pipe/docker_engine"

var DefaultSysSpec = SysSpec{OS: "Microsoft Windows Server 2016", KernelSpec: KernelSpec{Versions: []string{`10\.0\.1439[3-9]`, `10\.0\.14[4-9][0-9]{2}`, `10\.0\.1[5-9][0-9]{3}`, `10\.0\.[2-9][0-9]{4}`, `10\.[1-9]+\.[0-9]+`}, Required: []KernelConfig{}, Optional: []KernelConfig{}, Forbidden: []KernelConfig{}}, Cgroups: []string{}, RuntimeSpec: RuntimeSpec{DockerSpec: &DockerSpec{Version: []string{`18\.06\..*`}, GraphDriver: []string{"windowsfilter"}}}}

type KernelValidatorHelperImpl struct{}

var _ KernelValidatorHelper = &KernelValidatorHelperImpl{}

func (o *KernelValidatorHelperImpl) GetKernelReleaseVersion() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []string{"(Get-CimInstance Win32_OperatingSystem).Version"}
	releaseVersion, err := exec.Command("powershell", args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(releaseVersion)), nil
}
