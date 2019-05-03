package v1beta1

import (
 v1beta1 "k8s.io/api/certificates/v1beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1beta1.CertificateSigningRequest{}, func(obj interface{}) {
  SetObjectDefaults_CertificateSigningRequest(obj.(*v1beta1.CertificateSigningRequest))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.CertificateSigningRequestList{}, func(obj interface{}) {
  SetObjectDefaults_CertificateSigningRequestList(obj.(*v1beta1.CertificateSigningRequestList))
 })
 return nil
}
func SetObjectDefaults_CertificateSigningRequest(in *v1beta1.CertificateSigningRequest) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_CertificateSigningRequestSpec(&in.Spec)
}
func SetObjectDefaults_CertificateSigningRequestList(in *v1beta1.CertificateSigningRequestList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_CertificateSigningRequest(a)
 }
}
