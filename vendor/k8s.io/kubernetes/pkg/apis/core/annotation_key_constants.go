package core

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	ImagePolicyFailedOpenKey              string = "alpha.image-policy.k8s.io/failed-open"
	PodPresetOptOutAnnotationKey          string = "podpreset.admission.kubernetes.io/exclude"
	MirrorPodAnnotationKey                string = "kubernetes.io/config.mirror"
	TolerationsAnnotationKey              string = "scheduler.alpha.kubernetes.io/tolerations"
	TaintsAnnotationKey                   string = "scheduler.alpha.kubernetes.io/taints"
	SeccompPodAnnotationKey               string = "seccomp.security.alpha.kubernetes.io/pod"
	SeccompContainerAnnotationKeyPrefix   string = "container.seccomp.security.alpha.kubernetes.io/"
	SeccompProfileRuntimeDefault          string = "runtime/default"
	DeprecatedSeccompProfileDockerDefault string = "docker/default"
	PreferAvoidPodsAnnotationKey          string = "scheduler.alpha.kubernetes.io/preferAvoidPods"
	ObjectTTLAnnotationKey                string = "node.alpha.kubernetes.io/ttl"
	BootstrapCheckpointAnnotationKey      string = "node.kubernetes.io/bootstrap-checkpoint"
	NonConvertibleAnnotationPrefix               = "non-convertible.kubernetes.io"
	kubectlPrefix                                = "kubectl.kubernetes.io/"
	LastAppliedConfigAnnotation                  = kubectlPrefix + "last-applied-configuration"
	AnnotationLoadBalancerSourceRangesKey        = "service.beta.kubernetes.io/load-balancer-source-ranges"
	EndpointsLastChangeTriggerTime               = "endpoints.kubernetes.io/last-change-trigger-time"
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
