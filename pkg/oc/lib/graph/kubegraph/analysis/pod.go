package analysis

import (
	"fmt"
	"time"
	"github.com/MakeNowJust/heredoc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	CrashLoopingPodError	= "CrashLoopingPod"
	RestartingPodWarning	= "RestartingPod"
	RestartThreshold	= 5
	RestartRecentDuration	= 10 * time.Minute
)

var nowFn = metav1.Now

func FindRestartingPods(g osgraph.Graph, f osgraph.Namer, logsCommandName, securityPolicyCommandPattern string) []osgraph.Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	markers := []osgraph.Marker{}
	for _, uncastPodNode := range g.NodesByKind(kubegraph.PodNodeKind) {
		podNode := uncastPodNode.(*kubegraph.PodNode)
		pod, ok := podNode.Object().(*corev1.Pod)
		if !ok {
			continue
		}
		for _, containerStatus := range pod.Status.ContainerStatuses {
			containerString := ""
			if len(pod.Spec.Containers) > 1 {
				containerString = fmt.Sprintf("container %q in ", containerStatus.Name)
			}
			switch {
			case containerCrashLoopBackOff(containerStatus):
				var suggestion string
				switch {
				case containerIsNonRoot(pod, containerStatus.Name):
					suggestion = heredoc.Docf(`
						The container is starting and exiting repeatedly. This usually means the container is unable
						to start, misconfigured, or limited by security restrictions. Check the container logs with

						  %s %s -c %s

						Current security policy prevents your containers from being run as the root user. Some images
						may fail expecting to be able to change ownership or permissions on directories. Your admin
						can grant you access to run containers that need to run as the root user with this command:

						  %s
						`, logsCommandName, pod.Name, containerStatus.Name, fmt.Sprintf(securityPolicyCommandPattern, pod.Namespace, pod.Spec.ServiceAccountName))
				default:
					suggestion = heredoc.Docf(`
						The container is starting and exiting repeatedly. This usually means the container is unable
						to start, misconfigured, or limited by security restrictions. Check the container logs with

						  %s %s -c %s
						`, logsCommandName, pod.Name, containerStatus.Name)
				}
				markers = append(markers, osgraph.Marker{Node: podNode, Severity: osgraph.ErrorSeverity, Key: CrashLoopingPodError, Message: fmt.Sprintf("%s%s is crash-looping", containerString, f.ResourceName(podNode)), Suggestion: osgraph.Suggestion(suggestion)})
			case ContainerRestartedRecently(containerStatus, nowFn()):
				markers = append(markers, osgraph.Marker{Node: podNode, Severity: osgraph.WarningSeverity, Key: RestartingPodWarning, Message: fmt.Sprintf("%s%s has restarted within the last 10 minutes", containerString, f.ResourceName(podNode))})
			case containerRestartedFrequently(containerStatus):
				markers = append(markers, osgraph.Marker{Node: podNode, Severity: osgraph.WarningSeverity, Key: RestartingPodWarning, Message: fmt.Sprintf("%s%s has restarted %d times", containerString, f.ResourceName(podNode), containerStatus.RestartCount)})
			}
		}
	}
	return markers
}
func containerIsNonRoot(pod *corev1.Pod, container string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, c := range pod.Spec.Containers {
		if c.Name != container || c.SecurityContext == nil {
			continue
		}
		switch {
		case c.SecurityContext.RunAsUser != nil && *c.SecurityContext.RunAsUser != 0:
			return true
		}
	}
	return false
}
func containerCrashLoopBackOff(status corev1.ContainerStatus) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return status.State.Waiting != nil && status.State.Waiting.Reason == "CrashLoopBackOff"
}
func ContainerRestartedRecently(status corev1.ContainerStatus, now metav1.Time) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status.RestartCount == 0 {
		return false
	}
	if status.LastTerminationState.Terminated != nil && now.Sub(status.LastTerminationState.Terminated.FinishedAt.Time) < RestartRecentDuration {
		return true
	}
	return false
}
func containerRestartedFrequently(status corev1.ContainerStatus) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return status.RestartCount > RestartThreshold
}
