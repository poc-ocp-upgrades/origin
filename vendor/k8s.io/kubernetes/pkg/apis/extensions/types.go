package extensions

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/intstr"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type ReplicationControllerDummy struct{ metav1.TypeMeta }
type Ingress struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   IngressSpec
 Status IngressStatus
}
type IngressList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []Ingress
}
type IngressSpec struct {
 Backend *IngressBackend
 TLS     []IngressTLS
 Rules   []IngressRule
}
type IngressTLS struct {
 Hosts      []string
 SecretName string
}
type IngressStatus struct{ LoadBalancer api.LoadBalancerStatus }
type IngressRule struct {
 Host string
 IngressRuleValue
}
type IngressRuleValue struct{ HTTP *HTTPIngressRuleValue }
type HTTPIngressRuleValue struct{ Paths []HTTPIngressPath }
type HTTPIngressPath struct {
 Path    string
 Backend IngressBackend
}
type IngressBackend struct {
 ServiceName string
 ServicePort intstr.IntOrString
}
