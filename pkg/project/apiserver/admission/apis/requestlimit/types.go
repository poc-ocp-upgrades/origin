package requestlimit

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectRequestLimitConfig struct {
	metav1.TypeMeta
	Limits				[]ProjectLimitBySelector
	MaxProjectsForSystemUsers	*int
	MaxProjectsForServiceAccounts	*int
}
type ProjectLimitBySelector struct {
	Selector	map[string]string
	MaxProjects	*int
}
