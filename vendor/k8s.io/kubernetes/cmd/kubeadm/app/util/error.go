package util

import (
	"fmt"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"os"
	"strings"
)

const (
	DefaultErrorExitCode = 1
	PreFlightExitCode    = 2
	ValidationExitCode   = 3
)

func fatal(msg string, code int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(msg) > 0 {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}
func CheckErr(err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	checkErr(err, fatal)
}

type preflightError interface{ Preflight() bool }

func checkErr(err error, handleErr func(string, int)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch err.(type) {
	case nil:
		return
	case preflightError:
		handleErr(err.Error(), PreFlightExitCode)
	case utilerrors.Aggregate:
		handleErr(err.Error(), ValidationExitCode)
	default:
		handleErr(err.Error(), DefaultErrorExitCode)
	}
}
func FormatErrMsg(errs []error) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errMsg string
	for _, err := range errs {
		errMsg = fmt.Sprintf("%s\t- %s\n", errMsg, err.Error())
	}
	return errMsg
}
