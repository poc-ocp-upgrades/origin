package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TemplateServiceBrokerConfig struct {
	metav1.TypeMeta		`json:",inline"`
	TemplateNamespaces	[]string	`json:"templateNamespaces"`
}
