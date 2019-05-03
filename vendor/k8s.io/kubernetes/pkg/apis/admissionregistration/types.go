package admissionregistration

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InitializerConfiguration struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Initializers []Initializer
}
type InitializerConfigurationList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []InitializerConfiguration
}
type Initializer struct {
 Name  string
 Rules []Rule
}
type Rule struct {
 APIGroups   []string
 APIVersions []string
 Resources   []string
}
type FailurePolicyType string

const (
 Ignore FailurePolicyType = "Ignore"
 Fail   FailurePolicyType = "Fail"
)

type SideEffectClass string

const (
 SideEffectClassUnknown      SideEffectClass = "Unknown"
 SideEffectClassNone         SideEffectClass = "None"
 SideEffectClassSome         SideEffectClass = "Some"
 SideEffectClassNoneOnDryRun SideEffectClass = "NoneOnDryRun"
)

type ValidatingWebhookConfiguration struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Webhooks []Webhook
}
type ValidatingWebhookConfigurationList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []ValidatingWebhookConfiguration
}
type MutatingWebhookConfiguration struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Webhooks []Webhook
}
type MutatingWebhookConfigurationList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []MutatingWebhookConfiguration
}
type Webhook struct {
 Name              string
 ClientConfig      WebhookClientConfig
 Rules             []RuleWithOperations
 FailurePolicy     *FailurePolicyType
 NamespaceSelector *metav1.LabelSelector
 SideEffects       *SideEffectClass
}
type RuleWithOperations struct {
 Operations []OperationType
 Rule
}
type OperationType string

const (
 OperationAll OperationType = "*"
 Create       OperationType = "CREATE"
 Update       OperationType = "UPDATE"
 Delete       OperationType = "DELETE"
 Connect      OperationType = "CONNECT"
)

type WebhookClientConfig struct {
 URL      *string
 Service  *ServiceReference
 CABundle []byte
}
type ServiceReference struct {
 Namespace string
 Name      string
 Path      *string
}
