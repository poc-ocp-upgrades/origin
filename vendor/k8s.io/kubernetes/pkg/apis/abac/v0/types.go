package v0

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Policy struct {
 metav1.TypeMeta `json:",inline"`
 User            string `json:"user,omitempty"`
 Group           string `json:"group,omitempty"`
 Readonly        bool   `json:"readonly,omitempty"`
 Resource        string `json:"resource,omitempty"`
 Namespace       string `json:"namespace,omitempty"`
}
