package auditregistration

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Level string

const (
	LevelNone            Level = "None"
	LevelMetadata        Level = "Metadata"
	LevelRequest         Level = "Request"
	LevelRequestResponse Level = "RequestResponse"
)

type Stage string

const (
	StageRequestReceived  = "RequestReceived"
	StageResponseStarted  = "ResponseStarted"
	StageResponseComplete = "ResponseComplete"
	StagePanic            = "Panic"
)

type AuditSink struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec AuditSinkSpec
}
type AuditSinkSpec struct {
	Policy  Policy
	Webhook Webhook
}
type AuditSinkList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []AuditSink
}
type Policy struct {
	Level  Level
	Stages []Stage
}
type Webhook struct {
	Throttle     *WebhookThrottleConfig
	ClientConfig WebhookClientConfig
}
type WebhookThrottleConfig struct {
	QPS   *int64
	Burst *int64
}
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
