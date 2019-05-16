package preflight

import (
	"github.com/pkg/errors"
	"os"
)

func (ipuc IsPrivilegedUserCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errorList = []error{}
	if os.Getuid() != 0 {
		errorList = append(errorList, errors.New("user is not running as root"))
	}
	return nil, errorList
}
