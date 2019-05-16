package system

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

var _ Validator = &OSValidator{}

type OSValidator struct{ Reporter Reporter }

func (o *OSValidator) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "os"
}
func (o *OSValidator) Validate(spec SysSpec) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	os, err := exec.Command("uname").CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get os name")
	}
	return nil, o.validateOS(strings.TrimSpace(string(os)), spec.OS)
}
func (o *OSValidator) validateOS(os, specOS string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if os != specOS {
		o.Reporter.Report("OS", os, bad)
		return errors.Errorf("unsupported operating system: %s", os)
	}
	o.Reporter.Report("OS", os, good)
	return nil
}
