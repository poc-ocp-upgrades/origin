package podnodeconstraints

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodNodeConstraintsConfig struct {
    metav1.TypeMeta
    NodeSelectorLabelBlacklist []string
}
