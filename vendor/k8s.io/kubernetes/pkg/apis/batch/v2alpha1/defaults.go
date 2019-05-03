package v2alpha1

import (
 batchv2alpha1 "k8s.io/api/batch/v2alpha1"
 "k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_CronJob(obj *batchv2alpha1.CronJob) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.Spec.ConcurrencyPolicy == "" {
  obj.Spec.ConcurrencyPolicy = batchv2alpha1.AllowConcurrent
 }
 if obj.Spec.Suspend == nil {
  obj.Spec.Suspend = new(bool)
 }
}
