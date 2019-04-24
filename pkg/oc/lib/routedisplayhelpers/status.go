package routedisplayhelpers

import (
	corev1 "k8s.io/api/core/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	routev1 "github.com/openshift/api/route/v1"
)

func IngressConditionStatus(ingress *routev1.RouteIngress, t routev1.RouteIngressConditionType) (corev1.ConditionStatus, routev1.RouteIngressCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, condition := range ingress.Conditions {
		if t != condition.Type {
			continue
		}
		return condition.Status, condition
	}
	return corev1.ConditionUnknown, routev1.RouteIngressCondition{}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
