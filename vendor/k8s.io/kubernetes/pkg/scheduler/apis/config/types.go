package config

import (
 apimachineryconfig "k8s.io/apimachinery/pkg/apis/config"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 apiserverconfig "k8s.io/apiserver/pkg/apis/config"
)

const (
 SchedulerDefaultLockObjectNamespace string = metav1.NamespaceSystem
 SchedulerDefaultLockObjectName             = "kube-scheduler"
 SchedulerPolicyConfigMapKey                = "policy.cfg"
 SchedulerDefaultProviderName               = "DefaultProvider"
)

type KubeSchedulerConfiguration struct {
 metav1.TypeMeta
 SchedulerName                  string
 AlgorithmSource                SchedulerAlgorithmSource
 HardPodAffinitySymmetricWeight int32
 LeaderElection                 KubeSchedulerLeaderElectionConfiguration
 ClientConnection               apimachineryconfig.ClientConnectionConfiguration
 HealthzBindAddress             string
 MetricsBindAddress             string
 apiserverconfig.DebuggingConfiguration
 DisablePreemption        bool
 PercentageOfNodesToScore int32
 FailureDomains           string
 BindTimeoutSeconds       *int64
}
type SchedulerAlgorithmSource struct {
 Policy   *SchedulerPolicySource
 Provider *string
}
type SchedulerPolicySource struct {
 File      *SchedulerPolicyFileSource
 ConfigMap *SchedulerPolicyConfigMapSource
}
type SchedulerPolicyFileSource struct{ Path string }
type SchedulerPolicyConfigMapSource struct {
 Namespace string
 Name      string
}
type KubeSchedulerLeaderElectionConfiguration struct {
 apiserverconfig.LeaderElectionConfiguration
 LockObjectNamespace string
 LockObjectName      string
}
