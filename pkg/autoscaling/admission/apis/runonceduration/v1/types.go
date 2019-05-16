package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RunOnceDurationConfig struct {
	metav1.TypeMeta               `json:",inline"`
	ActiveDeadlineSecondsOverride *int64 `json:"activeDeadlineSecondsOverride,omitempty" description:"maximum value for activeDeadlineSeconds in run-once pods"`
}
