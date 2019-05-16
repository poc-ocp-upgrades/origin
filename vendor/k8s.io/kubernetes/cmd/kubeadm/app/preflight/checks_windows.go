package preflight

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

func (ipuc IsPrivilegedUserCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errorList = []error{}
	args := []string{"[bool](([System.Security.Principal.WindowsIdentity]::GetCurrent()).groups -match \"S-1-5-32-544\")"}
	isAdmin, err := exec.Command("powershell", args...).Output()
	if err != nil {
		errorList = append(errorList, errors.Wrap(err, "unable to determine if user is running as administrator"))
	} else if strings.EqualFold(strings.TrimSpace(string(isAdmin)), "false") {
		errorList = append(errorList, errors.New("user is not running as administrator"))
	}
	return nil, errorList
}
