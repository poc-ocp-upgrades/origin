package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type LimitType string

const (
	ServerLimitType          LimitType = "Server"
	NamespaceLimitType       LimitType = "Namespace"
	UserLimitType            LimitType = "User"
	SourceAndObjectLimitType LimitType = "SourceAndObject"
)

type Configuration struct {
	metav1.TypeMeta `json:",inline"`
	Limits          []Limit `json:"limits"`
}
type Limit struct {
	Type      LimitType `json:"type"`
	QPS       int32     `json:"qps"`
	Burst     int32     `json:"burst"`
	CacheSize int32     `json:"cacheSize,omitempty"`
}
