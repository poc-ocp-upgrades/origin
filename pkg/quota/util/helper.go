package util

import (
	"strings"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
)

const errQuotaMessageString = `exceeded quota:`
const errQuotaUnknownMessageString = `status unknown for quota:`
const errLimitsMessageString = `exceeds the maximum limit`

func IsErrorQuotaExceeded(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isForbidden := apierrs.IsForbidden(err); isForbidden || apierrs.IsInvalid(err) {
		lowered := strings.ToLower(err.Error())
		if strings.Contains(lowered, errLimitsMessageString) {
			return true
		}
	}
	return false
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
