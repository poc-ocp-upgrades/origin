package v1alpha1

import (
	goformat "fmt"
	auditregistrationv1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilpointer "k8s.io/utils/pointer"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	DefaultQPS   = int64(10)
	DefaultBurst = int64(15)
)

func DefaultThrottle() *auditregistrationv1alpha1.WebhookThrottleConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &auditregistrationv1alpha1.WebhookThrottleConfig{QPS: utilpointer.Int64Ptr(DefaultQPS), Burst: utilpointer.Int64Ptr(DefaultBurst)}
}
func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_AuditSink(obj *auditregistrationv1alpha1.AuditSink) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Spec.Webhook.Throttle != nil {
		if obj.Spec.Webhook.Throttle.QPS == nil {
			obj.Spec.Webhook.Throttle.QPS = utilpointer.Int64Ptr(DefaultQPS)
		}
		if obj.Spec.Webhook.Throttle.Burst == nil {
			obj.Spec.Webhook.Throttle.Burst = utilpointer.Int64Ptr(DefaultBurst)
		}
	} else {
		obj.Spec.Webhook.Throttle = DefaultThrottle()
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
