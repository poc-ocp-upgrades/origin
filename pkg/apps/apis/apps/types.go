package apps

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	kapi "k8s.io/kubernetes/pkg/apis/core"
)

const (
	DefaultRollingTimeoutSeconds      int64 = 10 * 60
	DefaultRecreateTimeoutSeconds     int64 = 10 * 60
	DefaultRollingIntervalSeconds     int64 = 1
	DefaultRollingUpdatePeriodSeconds int64 = 1
	MaxDeploymentDurationSeconds      int64 = 21600
	DefaultRevisionHistoryLimit       int32 = 10
)

type DeploymentConfig struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   DeploymentConfigSpec
	Status DeploymentConfigStatus
}
type DeploymentConfigSpec struct {
	Strategy             DeploymentStrategy
	MinReadySeconds      int32
	Triggers             []DeploymentTriggerPolicy
	Replicas             int32
	RevisionHistoryLimit *int32
	Test                 bool
	Paused               bool
	Selector             map[string]string
	Template             *kapi.PodTemplateSpec
}
type DeploymentStrategy struct {
	Type                  DeploymentStrategyType
	CustomParams          *CustomDeploymentStrategyParams
	RecreateParams        *RecreateDeploymentStrategyParams
	RollingParams         *RollingDeploymentStrategyParams
	Resources             kapi.ResourceRequirements
	Labels                map[string]string
	Annotations           map[string]string
	ActiveDeadlineSeconds *int64
}
type DeploymentStrategyType string

const (
	DeploymentStrategyTypeRecreate DeploymentStrategyType = "Recreate"
	DeploymentStrategyTypeCustom   DeploymentStrategyType = "Custom"
	DeploymentStrategyTypeRolling  DeploymentStrategyType = "Rolling"
)

type CustomDeploymentStrategyParams struct {
	Image       string
	Environment []kapi.EnvVar
	Command     []string
}
type RecreateDeploymentStrategyParams struct {
	TimeoutSeconds *int64
	Pre            *LifecycleHook
	Mid            *LifecycleHook
	Post           *LifecycleHook
}
type RollingDeploymentStrategyParams struct {
	UpdatePeriodSeconds *int64
	IntervalSeconds     *int64
	TimeoutSeconds      *int64
	MaxUnavailable      intstr.IntOrString
	MaxSurge            intstr.IntOrString
	Pre                 *LifecycleHook
	Post                *LifecycleHook
}
type LifecycleHook struct {
	FailurePolicy LifecycleHookFailurePolicy
	ExecNewPod    *ExecNewPodHook
	TagImages     []TagImageHook
}
type LifecycleHookFailurePolicy string

const (
	LifecycleHookFailurePolicyRetry  LifecycleHookFailurePolicy = "Retry"
	LifecycleHookFailurePolicyAbort  LifecycleHookFailurePolicy = "Abort"
	LifecycleHookFailurePolicyIgnore LifecycleHookFailurePolicy = "Ignore"
)

type ExecNewPodHook struct {
	Command       []string
	Env           []kapi.EnvVar
	ContainerName string
	Volumes       []string
}
type TagImageHook struct {
	ContainerName string
	To            kapi.ObjectReference
}
type DeploymentTriggerPolicy struct {
	Type              DeploymentTriggerType
	ImageChangeParams *DeploymentTriggerImageChangeParams
}
type DeploymentTriggerType string

const (
	DeploymentTriggerManual         DeploymentTriggerType = "Manual"
	DeploymentTriggerOnImageChange  DeploymentTriggerType = "ImageChange"
	DeploymentTriggerOnConfigChange DeploymentTriggerType = "ConfigChange"
)

type DeploymentTriggerImageChangeParams struct {
	Automatic          bool
	ContainerNames     []string
	From               kapi.ObjectReference
	LastTriggeredImage string
}
type DeploymentConfigStatus struct {
	LatestVersion       int64
	ObservedGeneration  int64
	Replicas            int32
	UpdatedReplicas     int32
	AvailableReplicas   int32
	UnavailableReplicas int32
	Details             *DeploymentDetails
	Conditions          []DeploymentCondition
	ReadyReplicas       int32
}
type DeploymentDetails struct {
	Message string
	Causes  []DeploymentCause
}
type DeploymentCause struct {
	Type         DeploymentTriggerType
	ImageTrigger *DeploymentCauseImageTrigger
}
type DeploymentCauseImageTrigger struct{ From kapi.ObjectReference }
type DeploymentConditionType string
type DeploymentConditionReason string
type DeploymentCondition struct {
	Type               DeploymentConditionType
	Status             kapi.ConditionStatus
	LastUpdateTime     metav1.Time
	LastTransitionTime metav1.Time
	Reason             DeploymentConditionReason
	Message            string
}
type DeploymentConfigList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []DeploymentConfig
}
type DeploymentConfigRollback struct {
	metav1.TypeMeta
	Name               string
	UpdatedAnnotations map[string]string
	Spec               DeploymentConfigRollbackSpec
}
type DeploymentConfigRollbackSpec struct {
	From                   kapi.ObjectReference
	Revision               int64
	IncludeTriggers        bool
	IncludeTemplate        bool
	IncludeReplicationMeta bool
	IncludeStrategy        bool
}
type DeploymentRequest struct {
	metav1.TypeMeta
	Name            string
	Latest          bool
	Force           bool
	ExcludeTriggers []DeploymentTriggerType
}
type DeploymentLog struct{ metav1.TypeMeta }
type DeploymentLogOptions struct {
	metav1.TypeMeta
	Container    string
	Follow       bool
	Previous     bool
	SinceSeconds *int64
	SinceTime    *metav1.Time
	Timestamps   bool
	TailLines    *int64
	LimitBytes   *int64
	NoWait       bool
	Version      *int64
}
