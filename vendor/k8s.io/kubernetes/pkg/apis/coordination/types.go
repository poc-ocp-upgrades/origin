package coordination

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Lease struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec LeaseSpec
}
type LeaseSpec struct {
 HolderIdentity       *string
 LeaseDurationSeconds *int32
 AcquireTime          *metav1.MicroTime
 RenewTime            *metav1.MicroTime
 LeaseTransitions     *int32
}
type LeaseList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []Lease
}
