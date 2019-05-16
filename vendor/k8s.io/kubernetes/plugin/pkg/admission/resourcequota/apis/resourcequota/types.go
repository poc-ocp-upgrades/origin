package resourcequota

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Configuration struct {
	metav1.TypeMeta
	LimitedResources []LimitedResource
}
type LimitedResource struct {
	APIGroup      string `json:"apiGroup,omitempty"`
	Resource      string `json:"resource"`
	MatchContains []string
	MatchScopes   []corev1.ScopedResourceSelectorRequirement `json:"matchScopes,omitempty"`
}
