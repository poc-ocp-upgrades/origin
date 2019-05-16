package v1alpha1

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/kube-controller-manager/config/v1alpha1"
	"k8s.io/kubernetes/pkg/controller/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *v1alpha1.GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in, out, s)
}
func Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *v1alpha1.GenericControllerManagerConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in, out, s)
}
func Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in *v1alpha1.KubeCloudSharedConfiguration, out *config.KubeCloudSharedConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in, out, s)
}
func Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in *config.KubeCloudSharedConfiguration, out *v1alpha1.KubeCloudSharedConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in, out, s)
}
func Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in *v1alpha1.ServiceControllerConfiguration, out *config.ServiceControllerConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in, out, s)
}
func Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in *config.ServiceControllerConfiguration, out *v1alpha1.ServiceControllerConfiguration, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in, out, s)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
