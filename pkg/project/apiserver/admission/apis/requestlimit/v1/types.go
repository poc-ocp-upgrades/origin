package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectRequestLimitConfig struct {
	metav1.TypeMeta			`json:",inline"`
	Limits				[]ProjectLimitBySelector	`json:"limits" description:"project request limits"`
	MaxProjectsForSystemUsers	*int				`json:"maxProjectsForSystemUsers"`
	MaxProjectsForServiceAccounts	*int				`json:"maxProjectsForServiceAccounts"`
}
type ProjectLimitBySelector struct {
	Selector	map[string]string	`json:"selector" description:"user label selector"`
	MaxProjects	*int			`json:"maxProjects,omitempty" description:"maximum number of projects, unlimited if nil"`
}
