package replication

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewReplicationControllerCondition(condType v1.ReplicationControllerConditionType, status v1.ConditionStatus, reason, msg string) v1.ReplicationControllerCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.ReplicationControllerCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Reason: reason, Message: msg}
}
func GetCondition(status v1.ReplicationControllerStatus, condType v1.ReplicationControllerConditionType) *v1.ReplicationControllerCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func SetCondition(status *v1.ReplicationControllerStatus, condition v1.ReplicationControllerCondition) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	currentCond := GetCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}
func RemoveCondition(status *v1.ReplicationControllerStatus, condType v1.ReplicationControllerConditionType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	status.Conditions = filterOutCondition(status.Conditions, condType)
}
func filterOutCondition(conditions []v1.ReplicationControllerCondition, condType v1.ReplicationControllerConditionType) []v1.ReplicationControllerCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var newConditions []v1.ReplicationControllerCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
