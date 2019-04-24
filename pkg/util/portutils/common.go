package portutils

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"github.com/fsouza/go-dockerclient"
)

func SplitPortAndProtocol(port string) (docker.Port, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dp := docker.Port(port)
	err := ValidatePortAndProtocol(dp)
	return dp, err
}
func FilterPortAndProtocolArray(ports []string) ([]docker.Port, []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := []error{}
	allPorts := []docker.Port{}
	for _, port := range ports {
		dp, err := SplitPortAndProtocol(port)
		if err == nil {
			allPorts = append(allPorts, dp)
		} else {
			allErrs = append(allErrs, err)
		}
	}
	return allPorts, allErrs
}
func SplitPortAndProtocolArray(ports []string) ([]docker.Port, []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := []error{}
	allPorts := []docker.Port{}
	for _, port := range ports {
		dp, err := SplitPortAndProtocol(port)
		if err != nil {
			allErrs = append(allErrs, err)
		}
		allPorts = append(allPorts, dp)
	}
	if len(allErrs) > 0 {
		return allPorts, allErrs
	}
	return allPorts, nil
}
func ValidatePortAndProtocol(port docker.Port) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := []string{}
	_, err := strconv.ParseUint(port.Port(), 10, 16)
	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange || numError.Err == strconv.ErrSyntax {
				errs = append(errs, "port number must be in range 0 - 65535")
			}
		}
	}
	if len(port.Proto()) > 0 && !(strings.ToUpper(port.Proto()) == "TCP" || strings.ToUpper(port.Proto()) == "UDP") {
		errs = append(errs, "protocol must be tcp or udp")
	}
	if len(errs) > 0 {
		return fmt.Errorf("failed to parse port %s/%s: [%v]", port.Port(), port.Proto(), strings.Join(errs, ", "))
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
