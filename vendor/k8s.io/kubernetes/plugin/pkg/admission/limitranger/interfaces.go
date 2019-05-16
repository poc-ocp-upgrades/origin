package limitranger

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
)

type LimitRangerActions interface {
	MutateLimit(limitRange *corev1.LimitRange, kind string, obj runtime.Object) error
	ValidateLimit(limitRange *corev1.LimitRange, kind string, obj runtime.Object) error
	SupportsAttributes(attr admission.Attributes) bool
	SupportsLimit(limitRange *corev1.LimitRange) bool
}
