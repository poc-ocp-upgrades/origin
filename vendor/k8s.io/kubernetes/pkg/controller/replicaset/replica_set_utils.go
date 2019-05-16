package replicaset

import (
	"fmt"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	appsclient "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/klog"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"reflect"
)

func updateReplicaSetStatus(c appsclient.ReplicaSetInterface, rs *apps.ReplicaSet, newStatus apps.ReplicaSetStatus) (*apps.ReplicaSet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rs.Status.Replicas == newStatus.Replicas && rs.Status.FullyLabeledReplicas == newStatus.FullyLabeledReplicas && rs.Status.ReadyReplicas == newStatus.ReadyReplicas && rs.Status.AvailableReplicas == newStatus.AvailableReplicas && rs.Generation == rs.Status.ObservedGeneration && reflect.DeepEqual(rs.Status.Conditions, newStatus.Conditions) {
		return rs, nil
	}
	newStatus.ObservedGeneration = rs.Generation
	var getErr, updateErr error
	var updatedRS *apps.ReplicaSet
	for i, rs := 0, rs; ; i++ {
		klog.V(4).Infof(fmt.Sprintf("Updating status for %v: %s/%s, ", rs.Kind, rs.Namespace, rs.Name) + fmt.Sprintf("replicas %d->%d (need %d), ", rs.Status.Replicas, newStatus.Replicas, *(rs.Spec.Replicas)) + fmt.Sprintf("fullyLabeledReplicas %d->%d, ", rs.Status.FullyLabeledReplicas, newStatus.FullyLabeledReplicas) + fmt.Sprintf("readyReplicas %d->%d, ", rs.Status.ReadyReplicas, newStatus.ReadyReplicas) + fmt.Sprintf("availableReplicas %d->%d, ", rs.Status.AvailableReplicas, newStatus.AvailableReplicas) + fmt.Sprintf("sequence No: %v->%v", rs.Status.ObservedGeneration, newStatus.ObservedGeneration))
		rs.Status = newStatus
		updatedRS, updateErr = c.UpdateStatus(rs)
		if updateErr == nil {
			return updatedRS, nil
		}
		if i >= statusUpdateRetries {
			break
		}
		if rs, getErr = c.Get(rs.Name, metav1.GetOptions{}); getErr != nil {
			return nil, getErr
		}
	}
	return nil, updateErr
}
func calculateStatus(rs *apps.ReplicaSet, filteredPods []*v1.Pod, manageReplicasErr error) apps.ReplicaSetStatus {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newStatus := rs.Status
	fullyLabeledReplicasCount := 0
	readyReplicasCount := 0
	availableReplicasCount := 0
	templateLabel := labels.Set(rs.Spec.Template.Labels).AsSelectorPreValidated()
	for _, pod := range filteredPods {
		if templateLabel.Matches(labels.Set(pod.Labels)) {
			fullyLabeledReplicasCount++
		}
		if podutil.IsPodReady(pod) {
			readyReplicasCount++
			if podutil.IsPodAvailable(pod, rs.Spec.MinReadySeconds, metav1.Now()) {
				availableReplicasCount++
			}
		}
	}
	failureCond := GetCondition(rs.Status, apps.ReplicaSetReplicaFailure)
	if manageReplicasErr != nil && failureCond == nil {
		var reason string
		if diff := len(filteredPods) - int(*(rs.Spec.Replicas)); diff < 0 {
			reason = "FailedCreate"
		} else if diff > 0 {
			reason = "FailedDelete"
		}
		cond := NewReplicaSetCondition(apps.ReplicaSetReplicaFailure, v1.ConditionTrue, reason, manageReplicasErr.Error())
		SetCondition(&newStatus, cond)
	} else if manageReplicasErr == nil && failureCond != nil {
		RemoveCondition(&newStatus, apps.ReplicaSetReplicaFailure)
	}
	newStatus.Replicas = int32(len(filteredPods))
	newStatus.FullyLabeledReplicas = int32(fullyLabeledReplicasCount)
	newStatus.ReadyReplicas = int32(readyReplicasCount)
	newStatus.AvailableReplicas = int32(availableReplicasCount)
	return newStatus
}
func NewReplicaSetCondition(condType apps.ReplicaSetConditionType, status v1.ConditionStatus, reason, msg string) apps.ReplicaSetCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apps.ReplicaSetCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Reason: reason, Message: msg}
}
func GetCondition(status apps.ReplicaSetStatus, condType apps.ReplicaSetConditionType) *apps.ReplicaSetCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, c := range status.Conditions {
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func SetCondition(status *apps.ReplicaSetStatus, condition apps.ReplicaSetCondition) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	currentCond := GetCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}
func RemoveCondition(status *apps.ReplicaSetStatus, condType apps.ReplicaSetConditionType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	status.Conditions = filterOutCondition(status.Conditions, condType)
}
func filterOutCondition(conditions []apps.ReplicaSetCondition, condType apps.ReplicaSetConditionType) []apps.ReplicaSetCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var newConditions []apps.ReplicaSetCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
