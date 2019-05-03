package v1beta1

import (
 v1beta1 "k8s.io/api/storage/v1beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1beta1.StorageClass{}, func(obj interface{}) {
  SetObjectDefaults_StorageClass(obj.(*v1beta1.StorageClass))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.StorageClassList{}, func(obj interface{}) {
  SetObjectDefaults_StorageClassList(obj.(*v1beta1.StorageClassList))
 })
 return nil
}
func SetObjectDefaults_StorageClass(in *v1beta1.StorageClass) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_StorageClass(in)
}
func SetObjectDefaults_StorageClassList(in *v1beta1.StorageClassList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_StorageClass(a)
 }
}
