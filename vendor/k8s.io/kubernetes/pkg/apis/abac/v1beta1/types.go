package v1beta1

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Policy struct {
 metav1.TypeMeta `json:",inline"`
 Spec            PolicySpec `json:"spec"`
}
type PolicySpec struct {
 User            string `json:"user,omitempty"`
 Group           string `json:"group,omitempty"`
 Readonly        bool   `json:"readonly,omitempty"`
 APIGroup        string `json:"apiGroup,omitempty"`
 Resource        string `json:"resource,omitempty"`
 Namespace       string `json:"namespace,omitempty"`
 NonResourcePath string `json:"nonResourcePath,omitempty"`
}
