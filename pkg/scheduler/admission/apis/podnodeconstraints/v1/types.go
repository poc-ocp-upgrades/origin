package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodNodeConstraintsConfig struct {
	metav1.TypeMeta            `json:",inline"`
	NodeSelectorLabelBlacklist []string `json:"nodeSelectorLabelBlacklist" description:"list of labels which cannot be set by entities without the 'pods/binding' permission"`
}
