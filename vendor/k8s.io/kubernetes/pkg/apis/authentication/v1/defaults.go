package v1

import (
 authenticationv1 "k8s.io/api/authentication/v1"
 "k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_TokenRequestSpec(obj *authenticationv1.TokenRequestSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.ExpirationSeconds == nil {
  hour := int64(60 * 60)
  obj.ExpirationSeconds = &hour
 }
}
