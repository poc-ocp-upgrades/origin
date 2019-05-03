package sysctl

import (
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type SysctlsStrategy interface {
 Validate(pod *api.Pod) field.ErrorList
}
