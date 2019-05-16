package strategy

import (
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

type DeploymentStrategy interface {
	Deploy(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error
}
type UpdateAcceptor interface {
	Accept(*corev1.ReplicationController) error
}
type errConditionReached struct{ msg string }

func NewConditionReachedErr(msg string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &errConditionReached{msg: msg}
}
func (e *errConditionReached) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.msg
}
func IsConditionReached(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value, ok := err.(*errConditionReached)
	return ok && value != nil
}
func PercentageBetween(until string, min, max int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasSuffix(until, "%") {
		return false
	}
	until = until[:len(until)-1]
	i, err := strconv.Atoi(until)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}
func Percentage(until string) (int, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasSuffix(until, "%") {
		return 0, false
	}
	until = until[:len(until)-1]
	i, err := strconv.Atoi(until)
	if err != nil {
		return 0, false
	}
	return i, true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
