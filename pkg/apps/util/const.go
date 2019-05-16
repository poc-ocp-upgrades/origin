package util

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	FailedRcCreateReason                            = "ReplicationControllerCreateError"
	NewReplicationControllerReason                  = "NewReplicationControllerCreated"
	NewRcAvailableReason                            = "NewReplicationControllerAvailable"
	TimedOutReason                                  = "ProgressDeadlineExceeded"
	PausedConfigReason                              = "DeploymentConfigPaused"
	CancelledRolloutReason                          = "RolloutCancelled"
	DeploymentConfigLabel                           = "deploymentconfig"
	DeploymentLabel                                 = "deployment"
	MaxDeploymentDurationSeconds              int64 = 21600
	DefaultRecreateTimeoutSeconds             int64 = 10 * 60
	DefaultRollingTimeoutSeconds              int64 = 10 * 60
	PreHookPodSuffix                                = "hook-pre"
	MidHookPodSuffix                                = "hook-mid"
	PostHookPodSuffix                               = "hook-post"
	DeploymentIgnorePodAnnotation                   = "deploy.openshift.io/deployer-pod.ignore"
	DeploymentReplicasAnnotation                    = "openshift.io/deployment.replicas"
	DeploymentFailedUnrelatedDeploymentExists       = "unrelated pod with the same name as this deployment is already running"
	DeploymentFailedUnableToCreateDeployerPod       = "unable to create deployer pod"
	DeploymentFailedDeployerPodNoLongerExists       = "deployer pod no longer exists"
	deploymentCancelledByUser                       = "cancelled by the user"
	deploymentCancelledNewerDeploymentExists        = "newer deployment was found running"
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
