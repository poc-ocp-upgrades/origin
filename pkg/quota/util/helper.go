package util

import (
	goformat "fmt"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const errQuotaMessageString = `exceeded quota:`
const errQuotaUnknownMessageString = `status unknown for quota:`
const errLimitsMessageString = `exceeds the maximum limit`

func IsErrorQuotaExceeded(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if isForbidden := apierrs.IsForbidden(err); isForbidden || apierrs.IsInvalid(err) {
		lowered := strings.ToLower(err.Error())
		if strings.Contains(lowered, errLimitsMessageString) {
			return true
		}
		if isForbidden && (strings.Contains(lowered, errQuotaMessageString) || strings.Contains(lowered, errQuotaUnknownMessageString)) {
			return true
		}
	}
	return false
}
func IsErrorLimitExceeded(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if isForbidden := apierrs.IsForbidden(err); isForbidden || apierrs.IsInvalid(err) {
		lowered := strings.ToLower(err.Error())
		if strings.Contains(lowered, errLimitsMessageString) {
			return true
		}
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
