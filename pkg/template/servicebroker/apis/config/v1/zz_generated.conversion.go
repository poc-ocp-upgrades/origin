package v1

import (
	config "github.com/openshift/origin/pkg/template/servicebroker/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*TemplateServiceBrokerConfig)(nil), (*config.TemplateServiceBrokerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateServiceBrokerConfig_To_config_TemplateServiceBrokerConfig(a.(*TemplateServiceBrokerConfig), b.(*config.TemplateServiceBrokerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TemplateServiceBrokerConfig)(nil), (*TemplateServiceBrokerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TemplateServiceBrokerConfig_To_v1_TemplateServiceBrokerConfig(a.(*config.TemplateServiceBrokerConfig), b.(*TemplateServiceBrokerConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_TemplateServiceBrokerConfig_To_config_TemplateServiceBrokerConfig(in *TemplateServiceBrokerConfig, out *config.TemplateServiceBrokerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TemplateNamespaces = *(*[]string)(unsafe.Pointer(&in.TemplateNamespaces))
	return nil
}
func Convert_v1_TemplateServiceBrokerConfig_To_config_TemplateServiceBrokerConfig(in *TemplateServiceBrokerConfig, out *config.TemplateServiceBrokerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateServiceBrokerConfig_To_config_TemplateServiceBrokerConfig(in, out, s)
}
func autoConvert_config_TemplateServiceBrokerConfig_To_v1_TemplateServiceBrokerConfig(in *config.TemplateServiceBrokerConfig, out *TemplateServiceBrokerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TemplateNamespaces = *(*[]string)(unsafe.Pointer(&in.TemplateNamespaces))
	return nil
}
func Convert_config_TemplateServiceBrokerConfig_To_v1_TemplateServiceBrokerConfig(in *config.TemplateServiceBrokerConfig, out *TemplateServiceBrokerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_TemplateServiceBrokerConfig_To_v1_TemplateServiceBrokerConfig(in, out, s)
}
