package apps

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/intstr"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type StatefulSet struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   StatefulSetSpec
 Status StatefulSetStatus
}
type PodManagementPolicyType string

const (
 OrderedReadyPodManagement PodManagementPolicyType = "OrderedReady"
 ParallelPodManagement                             = "Parallel"
)

type StatefulSetUpdateStrategy struct {
 Type          StatefulSetUpdateStrategyType
 RollingUpdate *RollingUpdateStatefulSetStrategy
}
type StatefulSetUpdateStrategyType string

const (
 RollingUpdateStatefulSetStrategyType = "RollingUpdate"
 OnDeleteStatefulSetStrategyType      = "OnDelete"
)

type RollingUpdateStatefulSetStrategy struct{ Partition int32 }
type StatefulSetSpec struct {
 Replicas             int32
 Selector             *metav1.LabelSelector
 Template             api.PodTemplateSpec
 VolumeClaimTemplates []api.PersistentVolumeClaim
 ServiceName          string
 PodManagementPolicy  PodManagementPolicyType
 UpdateStrategy       StatefulSetUpdateStrategy
 RevisionHistoryLimit *int32
}
type StatefulSetStatus struct {
 ObservedGeneration *int64
 Replicas           int32
 ReadyReplicas      int32
 CurrentReplicas    int32
 UpdatedReplicas    int32
 CurrentRevision    string
 UpdateRevision     string
 CollisionCount     *int32
 Conditions         []StatefulSetCondition
}
type StatefulSetConditionType string
type StatefulSetCondition struct {
 Type               StatefulSetConditionType
 Status             api.ConditionStatus
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
type StatefulSetList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []StatefulSet
}
type ControllerRevision struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Data     runtime.Object
 Revision int64
}
type ControllerRevisionList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []ControllerRevision
}
type Deployment struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   DeploymentSpec
 Status DeploymentStatus
}
type DeploymentSpec struct {
 Replicas                int32
 Selector                *metav1.LabelSelector
 Template                api.PodTemplateSpec
 Strategy                DeploymentStrategy
 MinReadySeconds         int32
 RevisionHistoryLimit    *int32
 Paused                  bool
 RollbackTo              *RollbackConfig
 ProgressDeadlineSeconds *int32
}
type DeploymentRollback struct {
 metav1.TypeMeta
 Name               string
 UpdatedAnnotations map[string]string
 RollbackTo         RollbackConfig
}
type RollbackConfig struct{ Revision int64 }

const (
 DefaultDeploymentUniqueLabelKey string = "pod-template-hash"
)

type DeploymentStrategy struct {
 Type          DeploymentStrategyType
 RollingUpdate *RollingUpdateDeployment
}
type DeploymentStrategyType string

const (
 RecreateDeploymentStrategyType      DeploymentStrategyType = "Recreate"
 RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
 MaxUnavailable intstr.IntOrString
 MaxSurge       intstr.IntOrString
}
type DeploymentStatus struct {
 ObservedGeneration  int64
 Replicas            int32
 UpdatedReplicas     int32
 ReadyReplicas       int32
 AvailableReplicas   int32
 UnavailableReplicas int32
 Conditions          []DeploymentCondition
 CollisionCount      *int32
}
type DeploymentConditionType string

const (
 DeploymentAvailable      DeploymentConditionType = "Available"
 DeploymentProgressing    DeploymentConditionType = "Progressing"
 DeploymentReplicaFailure DeploymentConditionType = "ReplicaFailure"
)

type DeploymentCondition struct {
 Type               DeploymentConditionType
 Status             api.ConditionStatus
 LastUpdateTime     metav1.Time
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
type DeploymentList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []Deployment
}
type DaemonSetUpdateStrategy struct {
 Type          DaemonSetUpdateStrategyType
 RollingUpdate *RollingUpdateDaemonSet
}
type DaemonSetUpdateStrategyType string

const (
 RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"
 OnDeleteDaemonSetStrategyType      DaemonSetUpdateStrategyType = "OnDelete"
)

type RollingUpdateDaemonSet struct{ MaxUnavailable intstr.IntOrString }
type DaemonSetSpec struct {
 Selector             *metav1.LabelSelector
 Template             api.PodTemplateSpec
 UpdateStrategy       DaemonSetUpdateStrategy
 MinReadySeconds      int32
 TemplateGeneration   int64
 RevisionHistoryLimit *int32
}
type DaemonSetStatus struct {
 CurrentNumberScheduled int32
 NumberMisscheduled     int32
 DesiredNumberScheduled int32
 NumberReady            int32
 ObservedGeneration     int64
 UpdatedNumberScheduled int32
 NumberAvailable        int32
 NumberUnavailable      int32
 CollisionCount         *int32
 Conditions             []DaemonSetCondition
}
type DaemonSetConditionType string
type DaemonSetCondition struct {
 Type               DaemonSetConditionType
 Status             api.ConditionStatus
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
type DaemonSet struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   DaemonSetSpec
 Status DaemonSetStatus
}

const (
 DaemonSetTemplateGenerationKey string = "pod-template-generation"
)

type DaemonSetList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []DaemonSet
}
type ReplicaSet struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   ReplicaSetSpec
 Status ReplicaSetStatus
}
type ReplicaSetList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []ReplicaSet
}
type ReplicaSetSpec struct {
 Replicas        int32
 MinReadySeconds int32
 Selector        *metav1.LabelSelector
 Template        api.PodTemplateSpec
}
type ReplicaSetStatus struct {
 Replicas             int32
 FullyLabeledReplicas int32
 ReadyReplicas        int32
 AvailableReplicas    int32
 ObservedGeneration   int64
 Conditions           []ReplicaSetCondition
}
type ReplicaSetConditionType string

const (
 ReplicaSetReplicaFailure ReplicaSetConditionType = "ReplicaFailure"
)

type ReplicaSetCondition struct {
 Type               ReplicaSetConditionType
 Status             api.ConditionStatus
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
