package runonceduration

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RunOnceDurationConfig struct {
	metav1.TypeMeta
	ActiveDeadlineSecondsLimit	*int64
}

const ActiveDeadlineSecondsLimitAnnotation = "openshift.io/active-deadline-seconds-override"
