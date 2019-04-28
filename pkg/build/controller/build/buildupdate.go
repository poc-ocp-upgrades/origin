package build

import (
	"fmt"
	"strings"
	"time"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/build/controller/common"
)

type buildUpdate struct {
	podNameAnnotation	*string
	phase			*buildv1.BuildPhase
	reason			*buildv1.StatusReason
	message			*string
	startTime		*metav1.Time
	completionTime		*metav1.Time
	duration		*time.Duration
	outputRef		*string
	logSnippet		*string
	pushSecret		*corev1.LocalObjectReference
}

func (u *buildUpdate) setPhase(phase buildv1.BuildPhase) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.phase = &phase
}
func (u *buildUpdate) setReason(reason buildv1.StatusReason) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.reason = &reason
}
func (u *buildUpdate) setMessage(message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.message = &message
}
func (u *buildUpdate) setStartTime(startTime metav1.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.startTime = &startTime
}
func (u *buildUpdate) setCompletionTime(completionTime metav1.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.completionTime = &completionTime
}
func (u *buildUpdate) setDuration(duration time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.duration = &duration
}
func (u *buildUpdate) setOutputRef(ref string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.outputRef = &ref
}
func (u *buildUpdate) setPodNameAnnotation(podName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.podNameAnnotation = &podName
}
func (u *buildUpdate) setLogSnippet(message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.logSnippet = &message
}
func (u *buildUpdate) setPushSecret(pushSecret corev1.LocalObjectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.pushSecret = &pushSecret
}
func (u *buildUpdate) reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.podNameAnnotation = nil
	u.phase = nil
	u.reason = nil
	u.message = nil
	u.startTime = nil
	u.completionTime = nil
	u.duration = nil
	u.outputRef = nil
	u.logSnippet = nil
	u.pushSecret = nil
}
func (u *buildUpdate) isEmpty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return u.podNameAnnotation == nil && u.phase == nil && u.reason == nil && u.message == nil && u.startTime == nil && u.completionTime == nil && u.duration == nil && u.outputRef == nil && u.logSnippet == nil && u.pushSecret == nil
}
func (u *buildUpdate) apply(build *buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if u.phase != nil {
		build.Status.Phase = *u.phase
	}
	if u.reason != nil {
		build.Status.Reason = *u.reason
	}
	if u.message != nil {
		build.Status.Message = *u.message
	}
	if u.startTime != nil {
		build.Status.StartTimestamp = u.startTime
	}
	if u.completionTime != nil {
		build.Status.CompletionTimestamp = u.completionTime
	}
	if u.duration != nil {
		build.Status.Duration = *u.duration
	}
	if u.podNameAnnotation != nil {
		common.SetBuildPodNameAnnotation(build, *u.podNameAnnotation)
	}
	if u.outputRef != nil {
		build.Status.OutputDockerImageReference = *u.outputRef
	}
	if u.logSnippet != nil {
		build.Status.LogSnippet = *u.logSnippet
	}
	if u.pushSecret != nil {
		build.Spec.Output.PushSecret = u.pushSecret
	}
}
func (u *buildUpdate) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	updates := []string{}
	if u.phase != nil {
		updates = append(updates, fmt.Sprintf("phase: %q", *u.phase))
	}
	if u.reason != nil {
		updates = append(updates, fmt.Sprintf("reason: %q", *u.reason))
	}
	if u.message != nil {
		updates = append(updates, fmt.Sprintf("message: %q", *u.message))
	}
	if u.startTime != nil {
		updates = append(updates, fmt.Sprintf("startTime: %q", u.startTime.String()))
	}
	if u.completionTime != nil {
		updates = append(updates, fmt.Sprintf("completionTime: %q", u.completionTime.String()))
	}
	if u.duration != nil {
		updates = append(updates, fmt.Sprintf("duration: %q", u.duration.String()))
	}
	if u.outputRef != nil {
		updates = append(updates, fmt.Sprintf("outputRef: %q", *u.outputRef))
	}
	if u.podNameAnnotation != nil {
		updates = append(updates, fmt.Sprintf("podName: %q", *u.podNameAnnotation))
	}
	if u.logSnippet != nil {
		updates = append(updates, fmt.Sprintf("logSnippet: %q", *u.logSnippet))
	}
	if u.pushSecret != nil {
		updates = append(updates, fmt.Sprintf("pushSecret: %v", *u.pushSecret))
	}
	return fmt.Sprintf("buildUpdate(%s)", strings.Join(updates, ", "))
}
