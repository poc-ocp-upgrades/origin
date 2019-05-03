package strategy

import (
	corev1 "k8s.io/api/core/v1"
	"strconv"
	"strings"
)

type DeploymentStrategy interface {
	Deploy(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error
}
type UpdateAcceptor interface {
	Accept(*corev1.ReplicationController) error
}
type errConditionReached struct{ msg string }

func NewConditionReachedErr(msg string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &errConditionReached{msg: msg}
}
func (e *errConditionReached) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.msg
}
func IsConditionReached(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value, ok := err.(*errConditionReached)
	return ok && value != nil
}
func PercentageBetween(until string, min, max int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
