package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Configuration struct {
	metav1.TypeMeta  `json:",inline"`
	LimitedResources []LimitedResource `json:"limitedResources"`
}
type LimitedResource struct {
	APIGroup      string                                 `json:"apiGroup,omitempty"`
	Resource      string                                 `json:"resource"`
	MatchContains []string                               `json:"matchContains,omitempty"`
	MatchScopes   []v1.ScopedResourceSelectorRequirement `json:"matchScopes,omitempty"`
}
