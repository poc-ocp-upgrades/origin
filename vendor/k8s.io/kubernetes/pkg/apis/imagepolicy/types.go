package imagepolicy

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ImageReview struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   ImageReviewSpec
 Status ImageReviewStatus
}
type ImageReviewSpec struct {
 Containers  []ImageReviewContainerSpec
 Annotations map[string]string
 Namespace   string
}
type ImageReviewContainerSpec struct{ Image string }
type ImageReviewStatus struct {
 Allowed          bool
 Reason           string
 AuditAnnotations map[string]string
}
