package v1

import (
 v1 "k8s.io/api/authentication/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1.TokenRequest{}, func(obj interface{}) {
  SetObjectDefaults_TokenRequest(obj.(*v1.TokenRequest))
 })
 return nil
}
func SetObjectDefaults_TokenRequest(in *v1.TokenRequest) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_TokenRequestSpec(&in.Spec)
}
