package predicates

import (
	"fmt"
	"k8s.io/api/core/v1"
)

var (
	ErrDiskConflict                          = newPredicateFailureError("NoDiskConflict", "node(s) had no available disk")
	ErrVolumeZoneConflict                    = newPredicateFailureError("NoVolumeZoneConflict", "node(s) had no available volume zone")
	ErrNodeSelectorNotMatch                  = newPredicateFailureError("MatchNodeSelector", "node(s) didn't match node selector")
	ErrPodAffinityNotMatch                   = newPredicateFailureError("MatchInterPodAffinity", "node(s) didn't match pod affinity/anti-affinity")
	ErrPodAffinityRulesNotMatch              = newPredicateFailureError("PodAffinityRulesNotMatch", "node(s) didn't match pod affinity rules")
	ErrPodAntiAffinityRulesNotMatch          = newPredicateFailureError("PodAntiAffinityRulesNotMatch", "node(s) didn't match pod anti-affinity rules")
	ErrExistingPodsAntiAffinityRulesNotMatch = newPredicateFailureError("ExistingPodsAntiAffinityRulesNotMatch", "node(s) didn't satisfy existing pods anti-affinity rules")
	ErrTaintsTolerationsNotMatch             = newPredicateFailureError("PodToleratesNodeTaints", "node(s) had taints that the pod didn't tolerate")
	ErrPodNotMatchHostName                   = newPredicateFailureError("HostName", "node(s) didn't match the requested hostname")
	ErrPodNotFitsHostPorts                   = newPredicateFailureError("PodFitsHostPorts", "node(s) didn't have free ports for the requested pod ports")
	ErrNodeLabelPresenceViolated             = newPredicateFailureError("CheckNodeLabelPresence", "node(s) didn't have the requested labels")
	ErrServiceAffinityViolated               = newPredicateFailureError("CheckServiceAffinity", "node(s) didn't match service affinity")
	ErrMaxVolumeCountExceeded                = newPredicateFailureError("MaxVolumeCount", "node(s) exceed max volume count")
	ErrNodeUnderMemoryPressure               = newPredicateFailureError("NodeUnderMemoryPressure", "node(s) had memory pressure")
	ErrNodeUnderDiskPressure                 = newPredicateFailureError("NodeUnderDiskPressure", "node(s) had disk pressure")
	ErrNodeUnderPIDPressure                  = newPredicateFailureError("NodeUnderPIDPressure", "node(s) had pid pressure")
	ErrNodeOutOfDisk                         = newPredicateFailureError("NodeOutOfDisk", "node(s) were out of disk space")
	ErrNodeNotReady                          = newPredicateFailureError("NodeNotReady", "node(s) were not ready")
	ErrNodeNetworkUnavailable                = newPredicateFailureError("NodeNetworkUnavailable", "node(s) had unavailable network")
	ErrNodeUnschedulable                     = newPredicateFailureError("NodeUnschedulable", "node(s) were unschedulable")
	ErrNodeUnknownCondition                  = newPredicateFailureError("NodeUnknownCondition", "node(s) had unknown conditions")
	ErrVolumeNodeConflict                    = newPredicateFailureError("VolumeNodeAffinityConflict", "node(s) had volume node affinity conflict")
	ErrVolumeBindConflict                    = newPredicateFailureError("VolumeBindingNoMatch", "node(s) didn't find available persistent volumes to bind")
	ErrFakePredicate                         = newPredicateFailureError("FakePredicateError", "Nodes failed the fake predicate")
)

type InsufficientResourceError struct {
	ResourceName v1.ResourceName
	requested    int64
	used         int64
	capacity     int64
}

func NewInsufficientResourceError(resourceName v1.ResourceName, requested, used, capacity int64) *InsufficientResourceError {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &InsufficientResourceError{ResourceName: resourceName, requested: requested, used: used, capacity: capacity}
}
func (e *InsufficientResourceError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("Node didn't have enough resource: %s, requested: %d, used: %d, capacity: %d", e.ResourceName, e.requested, e.used, e.capacity)
}
func (e *InsufficientResourceError) GetReason() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("Insufficient %v", e.ResourceName)
}
func (e *InsufficientResourceError) GetInsufficientAmount() int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.requested - (e.capacity - e.used)
}

type PredicateFailureError struct {
	PredicateName string
	PredicateDesc string
}

func newPredicateFailureError(predicateName, predicateDesc string) *PredicateFailureError {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PredicateFailureError{PredicateName: predicateName, PredicateDesc: predicateDesc}
}
func (e *PredicateFailureError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("Predicate %s failed", e.PredicateName)
}
func (e *PredicateFailureError) GetReason() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.PredicateDesc
}

type FailureReason struct{ reason string }

func NewFailureReason(msg string) *FailureReason {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FailureReason{reason: msg}
}
func (e *FailureReason) GetReason() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.reason
}
