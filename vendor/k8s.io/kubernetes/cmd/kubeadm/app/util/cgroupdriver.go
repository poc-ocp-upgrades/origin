package util

import (
	"github.com/pkg/errors"
	utilsexec "k8s.io/utils/exec"
	"strings"
)

func GetCgroupDriverDocker(execer utilsexec.Interface) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info, err := callDockerInfo(execer)
	if err != nil {
		return "", err
	}
	return getCgroupDriverFromDockerInfo(info)
}
func validateCgroupDriver(driver string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if driver != "cgroupfs" && driver != "systemd" {
		return errors.Errorf("unknown cgroup driver %q", driver)
	}
	return nil
}
func callDockerInfo(execer utilsexec.Interface) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out, err := execer.Command("docker", "info").Output()
	if err != nil {
		return "", errors.Wrap(err, "cannot execute 'docker info'")
	}
	return string(out), nil
}
func getCgroupDriverFromDockerInfo(info string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lineSeparator := ": "
	prefix := "Cgroup Driver"
	for _, line := range strings.Split(info, "\n") {
		if !strings.Contains(line, prefix+lineSeparator) {
			continue
		}
		lineSplit := strings.Split(line, lineSeparator)
		driver := lineSplit[1]
		if err := validateCgroupDriver(driver); err != nil {
			return "", err
		}
		return driver, nil
	}
	return "", errors.New("cgroup driver is not defined in 'docker info'")
}
