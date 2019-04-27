package v1

import (
	unsafe "unsafe"
	buildv1 "github.com/openshift/api/build/v1"
	v1 "github.com/openshift/api/legacyconfig/v1"
	build "github.com/openshift/origin/pkg/build/apis/build"
	config "github.com/openshift/origin/pkg/cmd/server/apis/config"
	apicorev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.ActiveDirectoryConfig)(nil), (*config.ActiveDirectoryConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ActiveDirectoryConfig_To_config_ActiveDirectoryConfig(a.(*v1.ActiveDirectoryConfig), b.(*config.ActiveDirectoryConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ActiveDirectoryConfig)(nil), (*v1.ActiveDirectoryConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ActiveDirectoryConfig_To_v1_ActiveDirectoryConfig(a.(*config.ActiveDirectoryConfig), b.(*v1.ActiveDirectoryConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AdmissionConfig)(nil), (*config.AdmissionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AdmissionConfig_To_config_AdmissionConfig(a.(*v1.AdmissionConfig), b.(*config.AdmissionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AdmissionConfig)(nil), (*v1.AdmissionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AdmissionConfig_To_v1_AdmissionConfig(a.(*config.AdmissionConfig), b.(*v1.AdmissionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AdmissionPluginConfig)(nil), (*config.AdmissionPluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AdmissionPluginConfig_To_config_AdmissionPluginConfig(a.(*v1.AdmissionPluginConfig), b.(*config.AdmissionPluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AdmissionPluginConfig)(nil), (*v1.AdmissionPluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AdmissionPluginConfig_To_v1_AdmissionPluginConfig(a.(*config.AdmissionPluginConfig), b.(*v1.AdmissionPluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AggregatorConfig)(nil), (*config.AggregatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AggregatorConfig_To_config_AggregatorConfig(a.(*v1.AggregatorConfig), b.(*config.AggregatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AggregatorConfig)(nil), (*v1.AggregatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AggregatorConfig_To_v1_AggregatorConfig(a.(*config.AggregatorConfig), b.(*v1.AggregatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AllowAllPasswordIdentityProvider)(nil), (*config.AllowAllPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AllowAllPasswordIdentityProvider_To_config_AllowAllPasswordIdentityProvider(a.(*v1.AllowAllPasswordIdentityProvider), b.(*config.AllowAllPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AllowAllPasswordIdentityProvider)(nil), (*v1.AllowAllPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AllowAllPasswordIdentityProvider_To_v1_AllowAllPasswordIdentityProvider(a.(*config.AllowAllPasswordIdentityProvider), b.(*v1.AllowAllPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AuditConfig)(nil), (*config.AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AuditConfig_To_config_AuditConfig(a.(*v1.AuditConfig), b.(*config.AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AuditConfig)(nil), (*v1.AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AuditConfig_To_v1_AuditConfig(a.(*config.AuditConfig), b.(*v1.AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AugmentedActiveDirectoryConfig)(nil), (*config.AugmentedActiveDirectoryConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AugmentedActiveDirectoryConfig_To_config_AugmentedActiveDirectoryConfig(a.(*v1.AugmentedActiveDirectoryConfig), b.(*config.AugmentedActiveDirectoryConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AugmentedActiveDirectoryConfig)(nil), (*v1.AugmentedActiveDirectoryConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AugmentedActiveDirectoryConfig_To_v1_AugmentedActiveDirectoryConfig(a.(*config.AugmentedActiveDirectoryConfig), b.(*v1.AugmentedActiveDirectoryConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BasicAuthPasswordIdentityProvider)(nil), (*config.BasicAuthPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BasicAuthPasswordIdentityProvider_To_config_BasicAuthPasswordIdentityProvider(a.(*v1.BasicAuthPasswordIdentityProvider), b.(*config.BasicAuthPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BasicAuthPasswordIdentityProvider)(nil), (*v1.BasicAuthPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BasicAuthPasswordIdentityProvider_To_v1_BasicAuthPasswordIdentityProvider(a.(*config.BasicAuthPasswordIdentityProvider), b.(*v1.BasicAuthPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildDefaultsConfig)(nil), (*config.BuildDefaultsConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildDefaultsConfig_To_config_BuildDefaultsConfig(a.(*v1.BuildDefaultsConfig), b.(*config.BuildDefaultsConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BuildDefaultsConfig)(nil), (*v1.BuildDefaultsConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BuildDefaultsConfig_To_v1_BuildDefaultsConfig(a.(*config.BuildDefaultsConfig), b.(*v1.BuildDefaultsConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildOverridesConfig)(nil), (*config.BuildOverridesConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildOverridesConfig_To_config_BuildOverridesConfig(a.(*v1.BuildOverridesConfig), b.(*config.BuildOverridesConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BuildOverridesConfig)(nil), (*v1.BuildOverridesConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BuildOverridesConfig_To_v1_BuildOverridesConfig(a.(*config.BuildOverridesConfig), b.(*v1.BuildOverridesConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CertInfo)(nil), (*config.CertInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CertInfo_To_config_CertInfo(a.(*v1.CertInfo), b.(*config.CertInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CertInfo)(nil), (*v1.CertInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CertInfo_To_v1_CertInfo(a.(*config.CertInfo), b.(*v1.CertInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClientConnectionOverrides)(nil), (*config.ClientConnectionOverrides)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClientConnectionOverrides_To_config_ClientConnectionOverrides(a.(*v1.ClientConnectionOverrides), b.(*config.ClientConnectionOverrides), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClientConnectionOverrides)(nil), (*v1.ClientConnectionOverrides)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClientConnectionOverrides_To_v1_ClientConnectionOverrides(a.(*config.ClientConnectionOverrides), b.(*v1.ClientConnectionOverrides), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterNetworkEntry)(nil), (*config.ClusterNetworkEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterNetworkEntry_To_config_ClusterNetworkEntry(a.(*v1.ClusterNetworkEntry), b.(*config.ClusterNetworkEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClusterNetworkEntry)(nil), (*v1.ClusterNetworkEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(a.(*config.ClusterNetworkEntry), b.(*v1.ClusterNetworkEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ControllerConfig)(nil), (*config.ControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ControllerConfig_To_config_ControllerConfig(a.(*v1.ControllerConfig), b.(*config.ControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfig)(nil), (*v1.ControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfig_To_v1_ControllerConfig(a.(*config.ControllerConfig), b.(*v1.ControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ControllerElectionConfig)(nil), (*config.ControllerElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ControllerElectionConfig_To_config_ControllerElectionConfig(a.(*v1.ControllerElectionConfig), b.(*config.ControllerElectionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerElectionConfig)(nil), (*v1.ControllerElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerElectionConfig_To_v1_ControllerElectionConfig(a.(*config.ControllerElectionConfig), b.(*v1.ControllerElectionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DNSConfig)(nil), (*config.DNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DNSConfig_To_config_DNSConfig(a.(*v1.DNSConfig), b.(*config.DNSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DNSConfig)(nil), (*v1.DNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DNSConfig_To_v1_DNSConfig(a.(*config.DNSConfig), b.(*v1.DNSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DefaultAdmissionConfig)(nil), (*config.DefaultAdmissionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DefaultAdmissionConfig_To_config_DefaultAdmissionConfig(a.(*v1.DefaultAdmissionConfig), b.(*config.DefaultAdmissionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DefaultAdmissionConfig)(nil), (*v1.DefaultAdmissionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DefaultAdmissionConfig_To_v1_DefaultAdmissionConfig(a.(*config.DefaultAdmissionConfig), b.(*v1.DefaultAdmissionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DenyAllPasswordIdentityProvider)(nil), (*config.DenyAllPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DenyAllPasswordIdentityProvider_To_config_DenyAllPasswordIdentityProvider(a.(*v1.DenyAllPasswordIdentityProvider), b.(*config.DenyAllPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DenyAllPasswordIdentityProvider)(nil), (*v1.DenyAllPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DenyAllPasswordIdentityProvider_To_v1_DenyAllPasswordIdentityProvider(a.(*config.DenyAllPasswordIdentityProvider), b.(*v1.DenyAllPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DockerConfig)(nil), (*config.DockerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DockerConfig_To_config_DockerConfig(a.(*v1.DockerConfig), b.(*config.DockerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DockerConfig)(nil), (*v1.DockerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DockerConfig_To_v1_DockerConfig(a.(*config.DockerConfig), b.(*v1.DockerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EtcdConfig)(nil), (*config.EtcdConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EtcdConfig_To_config_EtcdConfig(a.(*v1.EtcdConfig), b.(*config.EtcdConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.EtcdConfig)(nil), (*v1.EtcdConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_EtcdConfig_To_v1_EtcdConfig(a.(*config.EtcdConfig), b.(*v1.EtcdConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EtcdConnectionInfo)(nil), (*config.EtcdConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EtcdConnectionInfo_To_config_EtcdConnectionInfo(a.(*v1.EtcdConnectionInfo), b.(*config.EtcdConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.EtcdConnectionInfo)(nil), (*v1.EtcdConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_EtcdConnectionInfo_To_v1_EtcdConnectionInfo(a.(*config.EtcdConnectionInfo), b.(*v1.EtcdConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EtcdStorageConfig)(nil), (*config.EtcdStorageConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EtcdStorageConfig_To_config_EtcdStorageConfig(a.(*v1.EtcdStorageConfig), b.(*config.EtcdStorageConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.EtcdStorageConfig)(nil), (*v1.EtcdStorageConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_EtcdStorageConfig_To_v1_EtcdStorageConfig(a.(*config.EtcdStorageConfig), b.(*v1.EtcdStorageConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitHubIdentityProvider)(nil), (*config.GitHubIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitHubIdentityProvider_To_config_GitHubIdentityProvider(a.(*v1.GitHubIdentityProvider), b.(*config.GitHubIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GitHubIdentityProvider)(nil), (*v1.GitHubIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GitHubIdentityProvider_To_v1_GitHubIdentityProvider(a.(*config.GitHubIdentityProvider), b.(*v1.GitHubIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitLabIdentityProvider)(nil), (*config.GitLabIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitLabIdentityProvider_To_config_GitLabIdentityProvider(a.(*v1.GitLabIdentityProvider), b.(*config.GitLabIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GitLabIdentityProvider)(nil), (*v1.GitLabIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GitLabIdentityProvider_To_v1_GitLabIdentityProvider(a.(*config.GitLabIdentityProvider), b.(*v1.GitLabIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GoogleIdentityProvider)(nil), (*config.GoogleIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GoogleIdentityProvider_To_config_GoogleIdentityProvider(a.(*v1.GoogleIdentityProvider), b.(*config.GoogleIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GoogleIdentityProvider)(nil), (*v1.GoogleIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GoogleIdentityProvider_To_v1_GoogleIdentityProvider(a.(*config.GoogleIdentityProvider), b.(*v1.GoogleIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GrantConfig)(nil), (*config.GrantConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GrantConfig_To_config_GrantConfig(a.(*v1.GrantConfig), b.(*config.GrantConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GrantConfig)(nil), (*v1.GrantConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GrantConfig_To_v1_GrantConfig(a.(*config.GrantConfig), b.(*v1.GrantConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GroupResource)(nil), (*config.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupResource_To_config_GroupResource(a.(*v1.GroupResource), b.(*config.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GroupResource)(nil), (*v1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GroupResource_To_v1_GroupResource(a.(*config.GroupResource), b.(*v1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.HTPasswdPasswordIdentityProvider)(nil), (*config.HTPasswdPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HTPasswdPasswordIdentityProvider_To_config_HTPasswdPasswordIdentityProvider(a.(*v1.HTPasswdPasswordIdentityProvider), b.(*config.HTPasswdPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.HTPasswdPasswordIdentityProvider)(nil), (*v1.HTPasswdPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_HTPasswdPasswordIdentityProvider_To_v1_HTPasswdPasswordIdentityProvider(a.(*config.HTPasswdPasswordIdentityProvider), b.(*v1.HTPasswdPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.HTTPServingInfo)(nil), (*config.HTTPServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HTTPServingInfo_To_config_HTTPServingInfo(a.(*v1.HTTPServingInfo), b.(*config.HTTPServingInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.HTTPServingInfo)(nil), (*v1.HTTPServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_HTTPServingInfo_To_v1_HTTPServingInfo(a.(*config.HTTPServingInfo), b.(*v1.HTTPServingInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.IdentityProvider)(nil), (*config.IdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_IdentityProvider_To_config_IdentityProvider(a.(*v1.IdentityProvider), b.(*config.IdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.IdentityProvider)(nil), (*v1.IdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_IdentityProvider_To_v1_IdentityProvider(a.(*config.IdentityProvider), b.(*v1.IdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageConfig)(nil), (*config.ImageConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageConfig_To_config_ImageConfig(a.(*v1.ImageConfig), b.(*config.ImageConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ImageConfig)(nil), (*v1.ImageConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ImageConfig_To_v1_ImageConfig(a.(*config.ImageConfig), b.(*v1.ImageConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImagePolicyConfig)(nil), (*config.ImagePolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImagePolicyConfig_To_config_ImagePolicyConfig(a.(*v1.ImagePolicyConfig), b.(*config.ImagePolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ImagePolicyConfig)(nil), (*v1.ImagePolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ImagePolicyConfig_To_v1_ImagePolicyConfig(a.(*config.ImagePolicyConfig), b.(*v1.ImagePolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JenkinsPipelineConfig)(nil), (*config.JenkinsPipelineConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JenkinsPipelineConfig_To_config_JenkinsPipelineConfig(a.(*v1.JenkinsPipelineConfig), b.(*config.JenkinsPipelineConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.JenkinsPipelineConfig)(nil), (*v1.JenkinsPipelineConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_JenkinsPipelineConfig_To_v1_JenkinsPipelineConfig(a.(*config.JenkinsPipelineConfig), b.(*v1.JenkinsPipelineConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.KeystonePasswordIdentityProvider)(nil), (*config.KeystonePasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_KeystonePasswordIdentityProvider_To_config_KeystonePasswordIdentityProvider(a.(*v1.KeystonePasswordIdentityProvider), b.(*config.KeystonePasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KeystonePasswordIdentityProvider)(nil), (*v1.KeystonePasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KeystonePasswordIdentityProvider_To_v1_KeystonePasswordIdentityProvider(a.(*config.KeystonePasswordIdentityProvider), b.(*v1.KeystonePasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.KubeletConnectionInfo)(nil), (*config.KubeletConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_KubeletConnectionInfo_To_config_KubeletConnectionInfo(a.(*v1.KubeletConnectionInfo), b.(*config.KubeletConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletConnectionInfo)(nil), (*v1.KubeletConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletConnectionInfo_To_v1_KubeletConnectionInfo(a.(*config.KubeletConnectionInfo), b.(*v1.KubeletConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.KubernetesMasterConfig)(nil), (*config.KubernetesMasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_KubernetesMasterConfig_To_config_KubernetesMasterConfig(a.(*v1.KubernetesMasterConfig), b.(*config.KubernetesMasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubernetesMasterConfig)(nil), (*v1.KubernetesMasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubernetesMasterConfig_To_v1_KubernetesMasterConfig(a.(*config.KubernetesMasterConfig), b.(*v1.KubernetesMasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LDAPAttributeMapping)(nil), (*config.LDAPAttributeMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LDAPAttributeMapping_To_config_LDAPAttributeMapping(a.(*v1.LDAPAttributeMapping), b.(*config.LDAPAttributeMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LDAPAttributeMapping)(nil), (*v1.LDAPAttributeMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LDAPAttributeMapping_To_v1_LDAPAttributeMapping(a.(*config.LDAPAttributeMapping), b.(*v1.LDAPAttributeMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LDAPPasswordIdentityProvider)(nil), (*config.LDAPPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LDAPPasswordIdentityProvider_To_config_LDAPPasswordIdentityProvider(a.(*v1.LDAPPasswordIdentityProvider), b.(*config.LDAPPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LDAPPasswordIdentityProvider)(nil), (*v1.LDAPPasswordIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LDAPPasswordIdentityProvider_To_v1_LDAPPasswordIdentityProvider(a.(*config.LDAPPasswordIdentityProvider), b.(*v1.LDAPPasswordIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LDAPQuery)(nil), (*config.LDAPQuery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LDAPQuery_To_config_LDAPQuery(a.(*v1.LDAPQuery), b.(*config.LDAPQuery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LDAPQuery)(nil), (*v1.LDAPQuery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LDAPQuery_To_v1_LDAPQuery(a.(*config.LDAPQuery), b.(*v1.LDAPQuery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LDAPSyncConfig)(nil), (*config.LDAPSyncConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LDAPSyncConfig_To_config_LDAPSyncConfig(a.(*v1.LDAPSyncConfig), b.(*config.LDAPSyncConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LDAPSyncConfig)(nil), (*v1.LDAPSyncConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LDAPSyncConfig_To_v1_LDAPSyncConfig(a.(*config.LDAPSyncConfig), b.(*v1.LDAPSyncConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LocalQuota)(nil), (*config.LocalQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalQuota_To_config_LocalQuota(a.(*v1.LocalQuota), b.(*config.LocalQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LocalQuota)(nil), (*v1.LocalQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LocalQuota_To_v1_LocalQuota(a.(*config.LocalQuota), b.(*v1.LocalQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MasterAuthConfig)(nil), (*config.MasterAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterAuthConfig_To_config_MasterAuthConfig(a.(*v1.MasterAuthConfig), b.(*config.MasterAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MasterAuthConfig)(nil), (*v1.MasterAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterAuthConfig_To_v1_MasterAuthConfig(a.(*config.MasterAuthConfig), b.(*v1.MasterAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MasterClients)(nil), (*config.MasterClients)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterClients_To_config_MasterClients(a.(*v1.MasterClients), b.(*config.MasterClients), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MasterClients)(nil), (*v1.MasterClients)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterClients_To_v1_MasterClients(a.(*config.MasterClients), b.(*v1.MasterClients), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MasterConfig)(nil), (*config.MasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterConfig_To_config_MasterConfig(a.(*v1.MasterConfig), b.(*config.MasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MasterConfig)(nil), (*v1.MasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterConfig_To_v1_MasterConfig(a.(*config.MasterConfig), b.(*v1.MasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MasterNetworkConfig)(nil), (*config.MasterNetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterNetworkConfig_To_config_MasterNetworkConfig(a.(*v1.MasterNetworkConfig), b.(*config.MasterNetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MasterNetworkConfig)(nil), (*v1.MasterNetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterNetworkConfig_To_v1_MasterNetworkConfig(a.(*config.MasterNetworkConfig), b.(*v1.MasterNetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MasterVolumeConfig)(nil), (*config.MasterVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterVolumeConfig_To_config_MasterVolumeConfig(a.(*v1.MasterVolumeConfig), b.(*config.MasterVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MasterVolumeConfig)(nil), (*v1.MasterVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterVolumeConfig_To_v1_MasterVolumeConfig(a.(*config.MasterVolumeConfig), b.(*v1.MasterVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NamedCertificate)(nil), (*config.NamedCertificate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NamedCertificate_To_config_NamedCertificate(a.(*v1.NamedCertificate), b.(*config.NamedCertificate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NamedCertificate)(nil), (*v1.NamedCertificate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NamedCertificate_To_v1_NamedCertificate(a.(*config.NamedCertificate), b.(*v1.NamedCertificate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NodeAuthConfig)(nil), (*config.NodeAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeAuthConfig_To_config_NodeAuthConfig(a.(*v1.NodeAuthConfig), b.(*config.NodeAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeAuthConfig)(nil), (*v1.NodeAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeAuthConfig_To_v1_NodeAuthConfig(a.(*config.NodeAuthConfig), b.(*v1.NodeAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NodeConfig)(nil), (*config.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeConfig_To_config_NodeConfig(a.(*v1.NodeConfig), b.(*config.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeConfig)(nil), (*v1.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeConfig_To_v1_NodeConfig(a.(*config.NodeConfig), b.(*v1.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NodeNetworkConfig)(nil), (*config.NodeNetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeNetworkConfig_To_config_NodeNetworkConfig(a.(*v1.NodeNetworkConfig), b.(*config.NodeNetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeNetworkConfig)(nil), (*v1.NodeNetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeNetworkConfig_To_v1_NodeNetworkConfig(a.(*config.NodeNetworkConfig), b.(*v1.NodeNetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NodeVolumeConfig)(nil), (*config.NodeVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeVolumeConfig_To_config_NodeVolumeConfig(a.(*v1.NodeVolumeConfig), b.(*config.NodeVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeVolumeConfig)(nil), (*v1.NodeVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeVolumeConfig_To_v1_NodeVolumeConfig(a.(*config.NodeVolumeConfig), b.(*v1.NodeVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthConfig)(nil), (*config.OAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthConfig_To_config_OAuthConfig(a.(*v1.OAuthConfig), b.(*config.OAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OAuthConfig)(nil), (*v1.OAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OAuthConfig_To_v1_OAuthConfig(a.(*config.OAuthConfig), b.(*v1.OAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthTemplates)(nil), (*config.OAuthTemplates)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthTemplates_To_config_OAuthTemplates(a.(*v1.OAuthTemplates), b.(*config.OAuthTemplates), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OAuthTemplates)(nil), (*v1.OAuthTemplates)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OAuthTemplates_To_v1_OAuthTemplates(a.(*config.OAuthTemplates), b.(*v1.OAuthTemplates), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OpenIDClaims)(nil), (*config.OpenIDClaims)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OpenIDClaims_To_config_OpenIDClaims(a.(*v1.OpenIDClaims), b.(*config.OpenIDClaims), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OpenIDClaims)(nil), (*v1.OpenIDClaims)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OpenIDClaims_To_v1_OpenIDClaims(a.(*config.OpenIDClaims), b.(*v1.OpenIDClaims), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OpenIDIdentityProvider)(nil), (*config.OpenIDIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OpenIDIdentityProvider_To_config_OpenIDIdentityProvider(a.(*v1.OpenIDIdentityProvider), b.(*config.OpenIDIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OpenIDIdentityProvider)(nil), (*v1.OpenIDIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OpenIDIdentityProvider_To_v1_OpenIDIdentityProvider(a.(*config.OpenIDIdentityProvider), b.(*v1.OpenIDIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OpenIDURLs)(nil), (*config.OpenIDURLs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OpenIDURLs_To_config_OpenIDURLs(a.(*v1.OpenIDURLs), b.(*config.OpenIDURLs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OpenIDURLs)(nil), (*v1.OpenIDURLs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OpenIDURLs_To_v1_OpenIDURLs(a.(*config.OpenIDURLs), b.(*v1.OpenIDURLs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodManifestConfig)(nil), (*config.PodManifestConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodManifestConfig_To_config_PodManifestConfig(a.(*v1.PodManifestConfig), b.(*config.PodManifestConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PodManifestConfig)(nil), (*v1.PodManifestConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PodManifestConfig_To_v1_PodManifestConfig(a.(*config.PodManifestConfig), b.(*v1.PodManifestConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PolicyConfig)(nil), (*config.PolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyConfig_To_config_PolicyConfig(a.(*v1.PolicyConfig), b.(*config.PolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PolicyConfig)(nil), (*v1.PolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PolicyConfig_To_v1_PolicyConfig(a.(*config.PolicyConfig), b.(*v1.PolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProjectConfig)(nil), (*config.ProjectConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProjectConfig_To_config_ProjectConfig(a.(*v1.ProjectConfig), b.(*config.ProjectConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ProjectConfig)(nil), (*v1.ProjectConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ProjectConfig_To_v1_ProjectConfig(a.(*config.ProjectConfig), b.(*v1.ProjectConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RFC2307Config)(nil), (*config.RFC2307Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RFC2307Config_To_config_RFC2307Config(a.(*v1.RFC2307Config), b.(*config.RFC2307Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RFC2307Config)(nil), (*v1.RFC2307Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RFC2307Config_To_v1_RFC2307Config(a.(*config.RFC2307Config), b.(*v1.RFC2307Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RegistryLocation)(nil), (*config.RegistryLocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RegistryLocation_To_config_RegistryLocation(a.(*v1.RegistryLocation), b.(*config.RegistryLocation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RegistryLocation)(nil), (*v1.RegistryLocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RegistryLocation_To_v1_RegistryLocation(a.(*config.RegistryLocation), b.(*v1.RegistryLocation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RemoteConnectionInfo)(nil), (*config.RemoteConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RemoteConnectionInfo_To_config_RemoteConnectionInfo(a.(*v1.RemoteConnectionInfo), b.(*config.RemoteConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RemoteConnectionInfo)(nil), (*v1.RemoteConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RemoteConnectionInfo_To_v1_RemoteConnectionInfo(a.(*config.RemoteConnectionInfo), b.(*v1.RemoteConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RequestHeaderAuthenticationOptions)(nil), (*config.RequestHeaderAuthenticationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RequestHeaderAuthenticationOptions_To_config_RequestHeaderAuthenticationOptions(a.(*v1.RequestHeaderAuthenticationOptions), b.(*config.RequestHeaderAuthenticationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RequestHeaderAuthenticationOptions)(nil), (*v1.RequestHeaderAuthenticationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RequestHeaderAuthenticationOptions_To_v1_RequestHeaderAuthenticationOptions(a.(*config.RequestHeaderAuthenticationOptions), b.(*v1.RequestHeaderAuthenticationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RequestHeaderIdentityProvider)(nil), (*config.RequestHeaderIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RequestHeaderIdentityProvider_To_config_RequestHeaderIdentityProvider(a.(*v1.RequestHeaderIdentityProvider), b.(*config.RequestHeaderIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RequestHeaderIdentityProvider)(nil), (*v1.RequestHeaderIdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RequestHeaderIdentityProvider_To_v1_RequestHeaderIdentityProvider(a.(*config.RequestHeaderIdentityProvider), b.(*v1.RequestHeaderIdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoutingConfig)(nil), (*config.RoutingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoutingConfig_To_config_RoutingConfig(a.(*v1.RoutingConfig), b.(*config.RoutingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RoutingConfig)(nil), (*v1.RoutingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RoutingConfig_To_v1_RoutingConfig(a.(*config.RoutingConfig), b.(*v1.RoutingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecurityAllocator)(nil), (*config.SecurityAllocator)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecurityAllocator_To_config_SecurityAllocator(a.(*v1.SecurityAllocator), b.(*config.SecurityAllocator), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SecurityAllocator)(nil), (*v1.SecurityAllocator)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SecurityAllocator_To_v1_SecurityAllocator(a.(*config.SecurityAllocator), b.(*v1.SecurityAllocator), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountConfig)(nil), (*config.ServiceAccountConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAccountConfig_To_config_ServiceAccountConfig(a.(*v1.ServiceAccountConfig), b.(*config.ServiceAccountConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServiceAccountConfig)(nil), (*v1.ServiceAccountConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServiceAccountConfig_To_v1_ServiceAccountConfig(a.(*config.ServiceAccountConfig), b.(*v1.ServiceAccountConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceServingCert)(nil), (*config.ServiceServingCert)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceServingCert_To_config_ServiceServingCert(a.(*v1.ServiceServingCert), b.(*config.ServiceServingCert), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServiceServingCert)(nil), (*v1.ServiceServingCert)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServiceServingCert_To_v1_ServiceServingCert(a.(*config.ServiceServingCert), b.(*v1.ServiceServingCert), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServingInfo)(nil), (*config.ServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServingInfo_To_config_ServingInfo(a.(*v1.ServingInfo), b.(*config.ServingInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServingInfo)(nil), (*v1.ServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServingInfo_To_v1_ServingInfo(a.(*config.ServingInfo), b.(*v1.ServingInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SessionConfig)(nil), (*config.SessionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SessionConfig_To_config_SessionConfig(a.(*v1.SessionConfig), b.(*config.SessionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SessionConfig)(nil), (*v1.SessionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SessionConfig_To_v1_SessionConfig(a.(*config.SessionConfig), b.(*v1.SessionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SessionSecret)(nil), (*config.SessionSecret)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SessionSecret_To_config_SessionSecret(a.(*v1.SessionSecret), b.(*config.SessionSecret), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SessionSecret)(nil), (*v1.SessionSecret)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SessionSecret_To_v1_SessionSecret(a.(*config.SessionSecret), b.(*v1.SessionSecret), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SessionSecrets)(nil), (*config.SessionSecrets)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SessionSecrets_To_config_SessionSecrets(a.(*v1.SessionSecrets), b.(*config.SessionSecrets), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SessionSecrets)(nil), (*v1.SessionSecrets)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SessionSecrets_To_v1_SessionSecrets(a.(*config.SessionSecrets), b.(*v1.SessionSecrets), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SourceStrategyDefaultsConfig)(nil), (*config.SourceStrategyDefaultsConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceStrategyDefaultsConfig_To_config_SourceStrategyDefaultsConfig(a.(*v1.SourceStrategyDefaultsConfig), b.(*config.SourceStrategyDefaultsConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SourceStrategyDefaultsConfig)(nil), (*v1.SourceStrategyDefaultsConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SourceStrategyDefaultsConfig_To_v1_SourceStrategyDefaultsConfig(a.(*config.SourceStrategyDefaultsConfig), b.(*v1.SourceStrategyDefaultsConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.StringSource)(nil), (*config.StringSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_StringSource_To_config_StringSource(a.(*v1.StringSource), b.(*config.StringSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.StringSource)(nil), (*v1.StringSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_StringSource_To_v1_StringSource(a.(*config.StringSource), b.(*v1.StringSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.StringSourceSpec)(nil), (*config.StringSourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_StringSourceSpec_To_config_StringSourceSpec(a.(*v1.StringSourceSpec), b.(*config.StringSourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.StringSourceSpec)(nil), (*v1.StringSourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_StringSourceSpec_To_v1_StringSourceSpec(a.(*config.StringSourceSpec), b.(*v1.StringSourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TokenConfig)(nil), (*config.TokenConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TokenConfig_To_config_TokenConfig(a.(*v1.TokenConfig), b.(*config.TokenConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TokenConfig)(nil), (*v1.TokenConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TokenConfig_To_v1_TokenConfig(a.(*config.TokenConfig), b.(*v1.TokenConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserAgentDenyRule)(nil), (*config.UserAgentDenyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserAgentDenyRule_To_config_UserAgentDenyRule(a.(*v1.UserAgentDenyRule), b.(*config.UserAgentDenyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.UserAgentDenyRule)(nil), (*v1.UserAgentDenyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_UserAgentDenyRule_To_v1_UserAgentDenyRule(a.(*config.UserAgentDenyRule), b.(*v1.UserAgentDenyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserAgentMatchRule)(nil), (*config.UserAgentMatchRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserAgentMatchRule_To_config_UserAgentMatchRule(a.(*v1.UserAgentMatchRule), b.(*config.UserAgentMatchRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.UserAgentMatchRule)(nil), (*v1.UserAgentMatchRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_UserAgentMatchRule_To_v1_UserAgentMatchRule(a.(*config.UserAgentMatchRule), b.(*v1.UserAgentMatchRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserAgentMatchingConfig)(nil), (*config.UserAgentMatchingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserAgentMatchingConfig_To_config_UserAgentMatchingConfig(a.(*v1.UserAgentMatchingConfig), b.(*config.UserAgentMatchingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.UserAgentMatchingConfig)(nil), (*v1.UserAgentMatchingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_UserAgentMatchingConfig_To_v1_UserAgentMatchingConfig(a.(*config.UserAgentMatchingConfig), b.(*v1.UserAgentMatchingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.WebhookTokenAuthenticator)(nil), (*config.WebhookTokenAuthenticator)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_WebhookTokenAuthenticator_To_config_WebhookTokenAuthenticator(a.(*v1.WebhookTokenAuthenticator), b.(*config.WebhookTokenAuthenticator), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.WebhookTokenAuthenticator)(nil), (*v1.WebhookTokenAuthenticator)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_WebhookTokenAuthenticator_To_v1_WebhookTokenAuthenticator(a.(*config.WebhookTokenAuthenticator), b.(*v1.WebhookTokenAuthenticator), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.AdmissionPluginConfig)(nil), (*v1.AdmissionPluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AdmissionPluginConfig_To_v1_AdmissionPluginConfig(a.(*config.AdmissionPluginConfig), b.(*v1.AdmissionPluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.AuditConfig)(nil), (*v1.AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AuditConfig_To_v1_AuditConfig(a.(*config.AuditConfig), b.(*v1.AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.EtcdConnectionInfo)(nil), (*v1.EtcdConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_EtcdConnectionInfo_To_v1_EtcdConnectionInfo(a.(*config.EtcdConnectionInfo), b.(*v1.EtcdConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.IdentityProvider)(nil), (*v1.IdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_IdentityProvider_To_v1_IdentityProvider(a.(*config.IdentityProvider), b.(*v1.IdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.ImagePolicyConfig)(nil), (*v1.ImagePolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ImagePolicyConfig_To_v1_ImagePolicyConfig(a.(*config.ImagePolicyConfig), b.(*v1.ImagePolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.KubeletConnectionInfo)(nil), (*v1.KubeletConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletConnectionInfo_To_v1_KubeletConnectionInfo(a.(*config.KubeletConnectionInfo), b.(*v1.KubeletConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.KubernetesMasterConfig)(nil), (*v1.KubernetesMasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubernetesMasterConfig_To_v1_KubernetesMasterConfig(a.(*config.KubernetesMasterConfig), b.(*v1.KubernetesMasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.MasterVolumeConfig)(nil), (*v1.MasterVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MasterVolumeConfig_To_v1_MasterVolumeConfig(a.(*config.MasterVolumeConfig), b.(*v1.MasterVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.NodeConfig)(nil), (*v1.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeConfig_To_v1_NodeConfig(a.(*config.NodeConfig), b.(*v1.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.RemoteConnectionInfo)(nil), (*v1.RemoteConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RemoteConnectionInfo_To_v1_RemoteConnectionInfo(a.(*config.RemoteConnectionInfo), b.(*v1.RemoteConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.ServingInfo)(nil), (*v1.ServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServingInfo_To_v1_ServingInfo(a.(*config.ServingInfo), b.(*v1.ServingInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.AdmissionPluginConfig)(nil), (*config.AdmissionPluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AdmissionPluginConfig_To_config_AdmissionPluginConfig(a.(*v1.AdmissionPluginConfig), b.(*config.AdmissionPluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.AuditConfig)(nil), (*config.AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AuditConfig_To_config_AuditConfig(a.(*v1.AuditConfig), b.(*config.AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.EtcdConnectionInfo)(nil), (*config.EtcdConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EtcdConnectionInfo_To_config_EtcdConnectionInfo(a.(*v1.EtcdConnectionInfo), b.(*config.EtcdConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.IdentityProvider)(nil), (*config.IdentityProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_IdentityProvider_To_config_IdentityProvider(a.(*v1.IdentityProvider), b.(*config.IdentityProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ImagePolicyConfig)(nil), (*config.ImagePolicyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImagePolicyConfig_To_config_ImagePolicyConfig(a.(*v1.ImagePolicyConfig), b.(*config.ImagePolicyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.KubeletConnectionInfo)(nil), (*config.KubeletConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_KubeletConnectionInfo_To_config_KubeletConnectionInfo(a.(*v1.KubeletConnectionInfo), b.(*config.KubeletConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.KubernetesMasterConfig)(nil), (*config.KubernetesMasterConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_KubernetesMasterConfig_To_config_KubernetesMasterConfig(a.(*v1.KubernetesMasterConfig), b.(*config.KubernetesMasterConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.MasterNetworkConfig)(nil), (*config.MasterNetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterNetworkConfig_To_config_MasterNetworkConfig(a.(*v1.MasterNetworkConfig), b.(*config.MasterNetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.MasterVolumeConfig)(nil), (*config.MasterVolumeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MasterVolumeConfig_To_config_MasterVolumeConfig(a.(*v1.MasterVolumeConfig), b.(*config.MasterVolumeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.NodeConfig)(nil), (*config.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeConfig_To_config_NodeConfig(a.(*v1.NodeConfig), b.(*config.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.RemoteConnectionInfo)(nil), (*config.RemoteConnectionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RemoteConnectionInfo_To_config_RemoteConnectionInfo(a.(*v1.RemoteConnectionInfo), b.(*config.RemoteConnectionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ServingInfo)(nil), (*config.ServingInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServingInfo_To_config_ServingInfo(a.(*v1.ServingInfo), b.(*config.ServingInfo), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_ActiveDirectoryConfig_To_config_ActiveDirectoryConfig(in *v1.ActiveDirectoryConfig, out *config.ActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_LDAPQuery_To_config_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	return nil
}
func Convert_v1_ActiveDirectoryConfig_To_config_ActiveDirectoryConfig(in *v1.ActiveDirectoryConfig, out *config.ActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ActiveDirectoryConfig_To_config_ActiveDirectoryConfig(in, out, s)
}
func autoConvert_config_ActiveDirectoryConfig_To_v1_ActiveDirectoryConfig(in *config.ActiveDirectoryConfig, out *v1.ActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_LDAPQuery_To_v1_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	return nil
}
func Convert_config_ActiveDirectoryConfig_To_v1_ActiveDirectoryConfig(in *config.ActiveDirectoryConfig, out *v1.ActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ActiveDirectoryConfig_To_v1_ActiveDirectoryConfig(in, out, s)
}
func autoConvert_v1_AdmissionConfig_To_config_AdmissionConfig(in *v1.AdmissionConfig, out *config.AdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make(map[string]*config.AdmissionPluginConfig, len(*in))
		for key, val := range *in {
			newVal := new(*config.AdmissionPluginConfig)
			if err := s.Convert(&val, newVal, 0); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.PluginConfig = nil
	}
	out.PluginOrderOverride = *(*[]string)(unsafe.Pointer(&in.PluginOrderOverride))
	return nil
}
func Convert_v1_AdmissionConfig_To_config_AdmissionConfig(in *v1.AdmissionConfig, out *config.AdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AdmissionConfig_To_config_AdmissionConfig(in, out, s)
}
func autoConvert_config_AdmissionConfig_To_v1_AdmissionConfig(in *config.AdmissionConfig, out *v1.AdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make(map[string]*v1.AdmissionPluginConfig, len(*in))
		for key, val := range *in {
			newVal := new(*v1.AdmissionPluginConfig)
			if err := s.Convert(&val, newVal, 0); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.PluginConfig = nil
	}
	out.PluginOrderOverride = *(*[]string)(unsafe.Pointer(&in.PluginOrderOverride))
	return nil
}
func Convert_config_AdmissionConfig_To_v1_AdmissionConfig(in *config.AdmissionConfig, out *v1.AdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_AdmissionConfig_To_v1_AdmissionConfig(in, out, s)
}
func autoConvert_v1_AdmissionPluginConfig_To_config_AdmissionPluginConfig(in *v1.AdmissionPluginConfig, out *config.AdmissionPluginConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Location = in.Location
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Configuration, &out.Configuration, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_config_AdmissionPluginConfig_To_v1_AdmissionPluginConfig(in *config.AdmissionPluginConfig, out *v1.AdmissionPluginConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Location = in.Location
	return nil
}
func autoConvert_v1_AggregatorConfig_To_config_AggregatorConfig(in *v1.AggregatorConfig, out *config.AggregatorConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_CertInfo_To_config_CertInfo(&in.ProxyClientInfo, &out.ProxyClientInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_AggregatorConfig_To_config_AggregatorConfig(in *v1.AggregatorConfig, out *config.AggregatorConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AggregatorConfig_To_config_AggregatorConfig(in, out, s)
}
func autoConvert_config_AggregatorConfig_To_v1_AggregatorConfig(in *config.AggregatorConfig, out *v1.AggregatorConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_CertInfo_To_v1_CertInfo(&in.ProxyClientInfo, &out.ProxyClientInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_AggregatorConfig_To_v1_AggregatorConfig(in *config.AggregatorConfig, out *v1.AggregatorConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_AggregatorConfig_To_v1_AggregatorConfig(in, out, s)
}
func autoConvert_v1_AllowAllPasswordIdentityProvider_To_config_AllowAllPasswordIdentityProvider(in *v1.AllowAllPasswordIdentityProvider, out *config.AllowAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_v1_AllowAllPasswordIdentityProvider_To_config_AllowAllPasswordIdentityProvider(in *v1.AllowAllPasswordIdentityProvider, out *config.AllowAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AllowAllPasswordIdentityProvider_To_config_AllowAllPasswordIdentityProvider(in, out, s)
}
func autoConvert_config_AllowAllPasswordIdentityProvider_To_v1_AllowAllPasswordIdentityProvider(in *config.AllowAllPasswordIdentityProvider, out *v1.AllowAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_config_AllowAllPasswordIdentityProvider_To_v1_AllowAllPasswordIdentityProvider(in *config.AllowAllPasswordIdentityProvider, out *v1.AllowAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_AllowAllPasswordIdentityProvider_To_v1_AllowAllPasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_AuditConfig_To_config_AuditConfig(in *v1.AuditConfig, out *config.AuditConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Enabled = in.Enabled
	out.AuditFilePath = in.AuditFilePath
	out.MaximumFileRetentionDays = in.MaximumFileRetentionDays
	out.MaximumRetainedFiles = in.MaximumRetainedFiles
	out.MaximumFileSizeMegabytes = in.MaximumFileSizeMegabytes
	out.PolicyFile = in.PolicyFile
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.PolicyConfiguration, &out.PolicyConfiguration, s); err != nil {
		return err
	}
	out.LogFormat = config.LogFormatType(in.LogFormat)
	out.WebHookKubeConfig = in.WebHookKubeConfig
	out.WebHookMode = config.WebHookModeType(in.WebHookMode)
	return nil
}
func autoConvert_config_AuditConfig_To_v1_AuditConfig(in *config.AuditConfig, out *v1.AuditConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Enabled = in.Enabled
	out.AuditFilePath = in.AuditFilePath
	out.MaximumFileRetentionDays = in.MaximumFileRetentionDays
	out.MaximumRetainedFiles = in.MaximumRetainedFiles
	out.MaximumFileSizeMegabytes = in.MaximumFileSizeMegabytes
	out.PolicyFile = in.PolicyFile
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.PolicyConfiguration, &out.PolicyConfiguration, s); err != nil {
		return err
	}
	out.LogFormat = v1.LogFormatType(in.LogFormat)
	out.WebHookKubeConfig = in.WebHookKubeConfig
	out.WebHookMode = v1.WebHookModeType(in.WebHookMode)
	return nil
}
func autoConvert_v1_AugmentedActiveDirectoryConfig_To_config_AugmentedActiveDirectoryConfig(in *v1.AugmentedActiveDirectoryConfig, out *config.AugmentedActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_LDAPQuery_To_config_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	if err := Convert_v1_LDAPQuery_To_config_LDAPQuery(&in.AllGroupsQuery, &out.AllGroupsQuery, s); err != nil {
		return err
	}
	out.GroupUIDAttribute = in.GroupUIDAttribute
	out.GroupNameAttributes = *(*[]string)(unsafe.Pointer(&in.GroupNameAttributes))
	return nil
}
func Convert_v1_AugmentedActiveDirectoryConfig_To_config_AugmentedActiveDirectoryConfig(in *v1.AugmentedActiveDirectoryConfig, out *config.AugmentedActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AugmentedActiveDirectoryConfig_To_config_AugmentedActiveDirectoryConfig(in, out, s)
}
func autoConvert_config_AugmentedActiveDirectoryConfig_To_v1_AugmentedActiveDirectoryConfig(in *config.AugmentedActiveDirectoryConfig, out *v1.AugmentedActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_LDAPQuery_To_v1_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	if err := Convert_config_LDAPQuery_To_v1_LDAPQuery(&in.AllGroupsQuery, &out.AllGroupsQuery, s); err != nil {
		return err
	}
	out.GroupUIDAttribute = in.GroupUIDAttribute
	out.GroupNameAttributes = *(*[]string)(unsafe.Pointer(&in.GroupNameAttributes))
	return nil
}
func Convert_config_AugmentedActiveDirectoryConfig_To_v1_AugmentedActiveDirectoryConfig(in *config.AugmentedActiveDirectoryConfig, out *v1.AugmentedActiveDirectoryConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_AugmentedActiveDirectoryConfig_To_v1_AugmentedActiveDirectoryConfig(in, out, s)
}
func autoConvert_v1_BasicAuthPasswordIdentityProvider_To_config_BasicAuthPasswordIdentityProvider(in *v1.BasicAuthPasswordIdentityProvider, out *config.BasicAuthPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_RemoteConnectionInfo_To_config_RemoteConnectionInfo(&in.RemoteConnectionInfo, &out.RemoteConnectionInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_BasicAuthPasswordIdentityProvider_To_config_BasicAuthPasswordIdentityProvider(in *v1.BasicAuthPasswordIdentityProvider, out *config.BasicAuthPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BasicAuthPasswordIdentityProvider_To_config_BasicAuthPasswordIdentityProvider(in, out, s)
}
func autoConvert_config_BasicAuthPasswordIdentityProvider_To_v1_BasicAuthPasswordIdentityProvider(in *config.BasicAuthPasswordIdentityProvider, out *v1.BasicAuthPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_RemoteConnectionInfo_To_v1_RemoteConnectionInfo(&in.RemoteConnectionInfo, &out.RemoteConnectionInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_BasicAuthPasswordIdentityProvider_To_v1_BasicAuthPasswordIdentityProvider(in *config.BasicAuthPasswordIdentityProvider, out *v1.BasicAuthPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_BasicAuthPasswordIdentityProvider_To_v1_BasicAuthPasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_BuildDefaultsConfig_To_config_BuildDefaultsConfig(in *v1.BuildDefaultsConfig, out *config.BuildDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.GitHTTPProxy = in.GitHTTPProxy
	out.GitHTTPSProxy = in.GitHTTPSProxy
	out.GitNoProxy = in.GitNoProxy
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.SourceStrategyDefaults = (*config.SourceStrategyDefaultsConfig)(unsafe.Pointer(in.SourceStrategyDefaults))
	out.ImageLabels = *(*[]build.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if err := corev1.Convert_v1_ResourceRequirements_To_core_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_BuildDefaultsConfig_To_config_BuildDefaultsConfig(in *v1.BuildDefaultsConfig, out *config.BuildDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildDefaultsConfig_To_config_BuildDefaultsConfig(in, out, s)
}
func autoConvert_config_BuildDefaultsConfig_To_v1_BuildDefaultsConfig(in *config.BuildDefaultsConfig, out *v1.BuildDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.GitHTTPProxy = in.GitHTTPProxy
	out.GitHTTPSProxy = in.GitHTTPSProxy
	out.GitNoProxy = in.GitNoProxy
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.SourceStrategyDefaults = (*v1.SourceStrategyDefaultsConfig)(unsafe.Pointer(in.SourceStrategyDefaults))
	out.ImageLabels = *(*[]buildv1.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if err := corev1.Convert_core_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_BuildDefaultsConfig_To_v1_BuildDefaultsConfig(in *config.BuildDefaultsConfig, out *v1.BuildDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_BuildDefaultsConfig_To_v1_BuildDefaultsConfig(in, out, s)
}
func autoConvert_v1_BuildOverridesConfig_To_config_BuildOverridesConfig(in *v1.BuildOverridesConfig, out *config.BuildOverridesConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ForcePull = in.ForcePull
	out.ImageLabels = *(*[]build.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core.Toleration, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_Toleration_To_core_Toleration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Tolerations = nil
	}
	return nil
}
func Convert_v1_BuildOverridesConfig_To_config_BuildOverridesConfig(in *v1.BuildOverridesConfig, out *config.BuildOverridesConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildOverridesConfig_To_config_BuildOverridesConfig(in, out, s)
}
func autoConvert_config_BuildOverridesConfig_To_v1_BuildOverridesConfig(in *config.BuildOverridesConfig, out *v1.BuildOverridesConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ForcePull = in.ForcePull
	out.ImageLabels = *(*[]buildv1.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]apicorev1.Toleration, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_Toleration_To_v1_Toleration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Tolerations = nil
	}
	return nil
}
func Convert_config_BuildOverridesConfig_To_v1_BuildOverridesConfig(in *config.BuildOverridesConfig, out *v1.BuildOverridesConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_BuildOverridesConfig_To_v1_BuildOverridesConfig(in, out, s)
}
func autoConvert_v1_CertInfo_To_config_CertInfo(in *v1.CertInfo, out *config.CertInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}
func Convert_v1_CertInfo_To_config_CertInfo(in *v1.CertInfo, out *config.CertInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_CertInfo_To_config_CertInfo(in, out, s)
}
func autoConvert_config_CertInfo_To_v1_CertInfo(in *config.CertInfo, out *v1.CertInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}
func Convert_config_CertInfo_To_v1_CertInfo(in *config.CertInfo, out *v1.CertInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_CertInfo_To_v1_CertInfo(in, out, s)
}
func autoConvert_v1_ClientConnectionOverrides_To_config_ClientConnectionOverrides(in *v1.ClientConnectionOverrides, out *config.ClientConnectionOverrides, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}
func Convert_v1_ClientConnectionOverrides_To_config_ClientConnectionOverrides(in *v1.ClientConnectionOverrides, out *config.ClientConnectionOverrides, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClientConnectionOverrides_To_config_ClientConnectionOverrides(in, out, s)
}
func autoConvert_config_ClientConnectionOverrides_To_v1_ClientConnectionOverrides(in *config.ClientConnectionOverrides, out *v1.ClientConnectionOverrides, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}
func Convert_config_ClientConnectionOverrides_To_v1_ClientConnectionOverrides(in *config.ClientConnectionOverrides, out *v1.ClientConnectionOverrides, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ClientConnectionOverrides_To_v1_ClientConnectionOverrides(in, out, s)
}
func autoConvert_v1_ClusterNetworkEntry_To_config_ClusterNetworkEntry(in *v1.ClusterNetworkEntry, out *config.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDR = in.CIDR
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}
func Convert_v1_ClusterNetworkEntry_To_config_ClusterNetworkEntry(in *v1.ClusterNetworkEntry, out *config.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterNetworkEntry_To_config_ClusterNetworkEntry(in, out, s)
}
func autoConvert_config_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *config.ClusterNetworkEntry, out *v1.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDR = in.CIDR
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}
func Convert_config_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *config.ClusterNetworkEntry, out *v1.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in, out, s)
}
func autoConvert_v1_ControllerConfig_To_config_ControllerConfig(in *v1.ControllerConfig, out *config.ControllerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	out.Election = (*config.ControllerElectionConfig)(unsafe.Pointer(in.Election))
	if err := Convert_v1_ServiceServingCert_To_config_ServiceServingCert(&in.ServiceServingCert, &out.ServiceServingCert, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ControllerConfig_To_config_ControllerConfig(in *v1.ControllerConfig, out *config.ControllerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ControllerConfig_To_config_ControllerConfig(in, out, s)
}
func autoConvert_config_ControllerConfig_To_v1_ControllerConfig(in *config.ControllerConfig, out *v1.ControllerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	out.Election = (*v1.ControllerElectionConfig)(unsafe.Pointer(in.Election))
	if err := Convert_config_ServiceServingCert_To_v1_ServiceServingCert(&in.ServiceServingCert, &out.ServiceServingCert, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_ControllerConfig_To_v1_ControllerConfig(in *config.ControllerConfig, out *v1.ControllerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ControllerConfig_To_v1_ControllerConfig(in, out, s)
}
func autoConvert_v1_ControllerElectionConfig_To_config_ControllerElectionConfig(in *v1.ControllerElectionConfig, out *config.ControllerElectionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LockName = in.LockName
	out.LockNamespace = in.LockNamespace
	if err := Convert_v1_GroupResource_To_config_GroupResource(&in.LockResource, &out.LockResource, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ControllerElectionConfig_To_config_ControllerElectionConfig(in *v1.ControllerElectionConfig, out *config.ControllerElectionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ControllerElectionConfig_To_config_ControllerElectionConfig(in, out, s)
}
func autoConvert_config_ControllerElectionConfig_To_v1_ControllerElectionConfig(in *config.ControllerElectionConfig, out *v1.ControllerElectionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LockName = in.LockName
	out.LockNamespace = in.LockNamespace
	if err := Convert_config_GroupResource_To_v1_GroupResource(&in.LockResource, &out.LockResource, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_ControllerElectionConfig_To_v1_ControllerElectionConfig(in *config.ControllerElectionConfig, out *v1.ControllerElectionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ControllerElectionConfig_To_v1_ControllerElectionConfig(in, out, s)
}
func autoConvert_v1_DNSConfig_To_config_DNSConfig(in *v1.DNSConfig, out *config.DNSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BindAddress = in.BindAddress
	out.BindNetwork = in.BindNetwork
	out.AllowRecursiveQueries = in.AllowRecursiveQueries
	return nil
}
func Convert_v1_DNSConfig_To_config_DNSConfig(in *v1.DNSConfig, out *config.DNSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DNSConfig_To_config_DNSConfig(in, out, s)
}
func autoConvert_config_DNSConfig_To_v1_DNSConfig(in *config.DNSConfig, out *v1.DNSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BindAddress = in.BindAddress
	out.BindNetwork = in.BindNetwork
	out.AllowRecursiveQueries = in.AllowRecursiveQueries
	return nil
}
func Convert_config_DNSConfig_To_v1_DNSConfig(in *config.DNSConfig, out *v1.DNSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_DNSConfig_To_v1_DNSConfig(in, out, s)
}
func autoConvert_v1_DefaultAdmissionConfig_To_config_DefaultAdmissionConfig(in *v1.DefaultAdmissionConfig, out *config.DefaultAdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Disable = in.Disable
	return nil
}
func Convert_v1_DefaultAdmissionConfig_To_config_DefaultAdmissionConfig(in *v1.DefaultAdmissionConfig, out *config.DefaultAdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DefaultAdmissionConfig_To_config_DefaultAdmissionConfig(in, out, s)
}
func autoConvert_config_DefaultAdmissionConfig_To_v1_DefaultAdmissionConfig(in *config.DefaultAdmissionConfig, out *v1.DefaultAdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Disable = in.Disable
	return nil
}
func Convert_config_DefaultAdmissionConfig_To_v1_DefaultAdmissionConfig(in *config.DefaultAdmissionConfig, out *v1.DefaultAdmissionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_DefaultAdmissionConfig_To_v1_DefaultAdmissionConfig(in, out, s)
}
func autoConvert_v1_DenyAllPasswordIdentityProvider_To_config_DenyAllPasswordIdentityProvider(in *v1.DenyAllPasswordIdentityProvider, out *config.DenyAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_v1_DenyAllPasswordIdentityProvider_To_config_DenyAllPasswordIdentityProvider(in *v1.DenyAllPasswordIdentityProvider, out *config.DenyAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DenyAllPasswordIdentityProvider_To_config_DenyAllPasswordIdentityProvider(in, out, s)
}
func autoConvert_config_DenyAllPasswordIdentityProvider_To_v1_DenyAllPasswordIdentityProvider(in *config.DenyAllPasswordIdentityProvider, out *v1.DenyAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_config_DenyAllPasswordIdentityProvider_To_v1_DenyAllPasswordIdentityProvider(in *config.DenyAllPasswordIdentityProvider, out *v1.DenyAllPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_DenyAllPasswordIdentityProvider_To_v1_DenyAllPasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_DockerConfig_To_config_DockerConfig(in *v1.DockerConfig, out *config.DockerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExecHandlerName = config.DockerExecHandlerType(in.ExecHandlerName)
	out.DockerShimSocket = in.DockerShimSocket
	out.DockershimRootDirectory = in.DockershimRootDirectory
	return nil
}
func Convert_v1_DockerConfig_To_config_DockerConfig(in *v1.DockerConfig, out *config.DockerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DockerConfig_To_config_DockerConfig(in, out, s)
}
func autoConvert_config_DockerConfig_To_v1_DockerConfig(in *config.DockerConfig, out *v1.DockerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExecHandlerName = v1.DockerExecHandlerType(in.ExecHandlerName)
	out.DockerShimSocket = in.DockerShimSocket
	out.DockershimRootDirectory = in.DockershimRootDirectory
	return nil
}
func Convert_config_DockerConfig_To_v1_DockerConfig(in *config.DockerConfig, out *v1.DockerConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_DockerConfig_To_v1_DockerConfig(in, out, s)
}
func autoConvert_v1_EtcdConfig_To_config_EtcdConfig(in *v1.EtcdConfig, out *config.EtcdConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_ServingInfo_To_config_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.Address = in.Address
	if err := Convert_v1_ServingInfo_To_config_ServingInfo(&in.PeerServingInfo, &out.PeerServingInfo, s); err != nil {
		return err
	}
	out.PeerAddress = in.PeerAddress
	out.StorageDir = in.StorageDir
	return nil
}
func Convert_v1_EtcdConfig_To_config_EtcdConfig(in *v1.EtcdConfig, out *config.EtcdConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EtcdConfig_To_config_EtcdConfig(in, out, s)
}
func autoConvert_config_EtcdConfig_To_v1_EtcdConfig(in *config.EtcdConfig, out *v1.EtcdConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_ServingInfo_To_v1_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.Address = in.Address
	if err := Convert_config_ServingInfo_To_v1_ServingInfo(&in.PeerServingInfo, &out.PeerServingInfo, s); err != nil {
		return err
	}
	out.PeerAddress = in.PeerAddress
	out.StorageDir = in.StorageDir
	return nil
}
func Convert_config_EtcdConfig_To_v1_EtcdConfig(in *config.EtcdConfig, out *v1.EtcdConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_EtcdConfig_To_v1_EtcdConfig(in, out, s)
}
func autoConvert_v1_EtcdConnectionInfo_To_config_EtcdConnectionInfo(in *v1.EtcdConnectionInfo, out *config.EtcdConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URLs = *(*[]string)(unsafe.Pointer(&in.URLs))
	out.CA = in.CA
	return nil
}
func autoConvert_config_EtcdConnectionInfo_To_v1_EtcdConnectionInfo(in *config.EtcdConnectionInfo, out *v1.EtcdConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URLs = *(*[]string)(unsafe.Pointer(&in.URLs))
	out.CA = in.CA
	return nil
}
func autoConvert_v1_EtcdStorageConfig_To_config_EtcdStorageConfig(in *v1.EtcdStorageConfig, out *config.EtcdStorageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.KubernetesStorageVersion = in.KubernetesStorageVersion
	out.KubernetesStoragePrefix = in.KubernetesStoragePrefix
	out.OpenShiftStorageVersion = in.OpenShiftStorageVersion
	out.OpenShiftStoragePrefix = in.OpenShiftStoragePrefix
	return nil
}
func Convert_v1_EtcdStorageConfig_To_config_EtcdStorageConfig(in *v1.EtcdStorageConfig, out *config.EtcdStorageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EtcdStorageConfig_To_config_EtcdStorageConfig(in, out, s)
}
func autoConvert_config_EtcdStorageConfig_To_v1_EtcdStorageConfig(in *config.EtcdStorageConfig, out *v1.EtcdStorageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.KubernetesStorageVersion = in.KubernetesStorageVersion
	out.KubernetesStoragePrefix = in.KubernetesStoragePrefix
	out.OpenShiftStorageVersion = in.OpenShiftStorageVersion
	out.OpenShiftStoragePrefix = in.OpenShiftStoragePrefix
	return nil
}
func Convert_config_EtcdStorageConfig_To_v1_EtcdStorageConfig(in *config.EtcdStorageConfig, out *v1.EtcdStorageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_EtcdStorageConfig_To_v1_EtcdStorageConfig(in, out, s)
}
func autoConvert_v1_GitHubIdentityProvider_To_config_GitHubIdentityProvider(in *v1.GitHubIdentityProvider, out *config.GitHubIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientID = in.ClientID
	if err := Convert_v1_StringSource_To_config_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.Organizations = *(*[]string)(unsafe.Pointer(&in.Organizations))
	out.Teams = *(*[]string)(unsafe.Pointer(&in.Teams))
	out.Hostname = in.Hostname
	out.CA = in.CA
	return nil
}
func Convert_v1_GitHubIdentityProvider_To_config_GitHubIdentityProvider(in *v1.GitHubIdentityProvider, out *config.GitHubIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitHubIdentityProvider_To_config_GitHubIdentityProvider(in, out, s)
}
func autoConvert_config_GitHubIdentityProvider_To_v1_GitHubIdentityProvider(in *config.GitHubIdentityProvider, out *v1.GitHubIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientID = in.ClientID
	if err := Convert_config_StringSource_To_v1_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.Organizations = *(*[]string)(unsafe.Pointer(&in.Organizations))
	out.Teams = *(*[]string)(unsafe.Pointer(&in.Teams))
	out.Hostname = in.Hostname
	out.CA = in.CA
	return nil
}
func Convert_config_GitHubIdentityProvider_To_v1_GitHubIdentityProvider(in *config.GitHubIdentityProvider, out *v1.GitHubIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_GitHubIdentityProvider_To_v1_GitHubIdentityProvider(in, out, s)
}
func autoConvert_v1_GitLabIdentityProvider_To_config_GitLabIdentityProvider(in *v1.GitLabIdentityProvider, out *config.GitLabIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CA = in.CA
	out.URL = in.URL
	out.ClientID = in.ClientID
	if err := Convert_v1_StringSource_To_config_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.Legacy = (*bool)(unsafe.Pointer(in.Legacy))
	return nil
}
func Convert_v1_GitLabIdentityProvider_To_config_GitLabIdentityProvider(in *v1.GitLabIdentityProvider, out *config.GitLabIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitLabIdentityProvider_To_config_GitLabIdentityProvider(in, out, s)
}
func autoConvert_config_GitLabIdentityProvider_To_v1_GitLabIdentityProvider(in *config.GitLabIdentityProvider, out *v1.GitLabIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CA = in.CA
	out.URL = in.URL
	out.ClientID = in.ClientID
	if err := Convert_config_StringSource_To_v1_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.Legacy = (*bool)(unsafe.Pointer(in.Legacy))
	return nil
}
func Convert_config_GitLabIdentityProvider_To_v1_GitLabIdentityProvider(in *config.GitLabIdentityProvider, out *v1.GitLabIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_GitLabIdentityProvider_To_v1_GitLabIdentityProvider(in, out, s)
}
func autoConvert_v1_GoogleIdentityProvider_To_config_GoogleIdentityProvider(in *v1.GoogleIdentityProvider, out *config.GoogleIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientID = in.ClientID
	if err := Convert_v1_StringSource_To_config_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.HostedDomain = in.HostedDomain
	return nil
}
func Convert_v1_GoogleIdentityProvider_To_config_GoogleIdentityProvider(in *v1.GoogleIdentityProvider, out *config.GoogleIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GoogleIdentityProvider_To_config_GoogleIdentityProvider(in, out, s)
}
func autoConvert_config_GoogleIdentityProvider_To_v1_GoogleIdentityProvider(in *config.GoogleIdentityProvider, out *v1.GoogleIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientID = in.ClientID
	if err := Convert_config_StringSource_To_v1_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.HostedDomain = in.HostedDomain
	return nil
}
func Convert_config_GoogleIdentityProvider_To_v1_GoogleIdentityProvider(in *config.GoogleIdentityProvider, out *v1.GoogleIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_GoogleIdentityProvider_To_v1_GoogleIdentityProvider(in, out, s)
}
func autoConvert_v1_GrantConfig_To_config_GrantConfig(in *v1.GrantConfig, out *config.GrantConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Method = config.GrantHandlerType(in.Method)
	out.ServiceAccountMethod = config.GrantHandlerType(in.ServiceAccountMethod)
	return nil
}
func Convert_v1_GrantConfig_To_config_GrantConfig(in *v1.GrantConfig, out *config.GrantConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GrantConfig_To_config_GrantConfig(in, out, s)
}
func autoConvert_config_GrantConfig_To_v1_GrantConfig(in *config.GrantConfig, out *v1.GrantConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Method = v1.GrantHandlerType(in.Method)
	out.ServiceAccountMethod = v1.GrantHandlerType(in.ServiceAccountMethod)
	return nil
}
func Convert_config_GrantConfig_To_v1_GrantConfig(in *config.GrantConfig, out *v1.GrantConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_GrantConfig_To_v1_GrantConfig(in, out, s)
}
func autoConvert_v1_GroupResource_To_config_GroupResource(in *v1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}
func Convert_v1_GroupResource_To_config_GroupResource(in *v1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GroupResource_To_config_GroupResource(in, out, s)
}
func autoConvert_config_GroupResource_To_v1_GroupResource(in *config.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}
func Convert_config_GroupResource_To_v1_GroupResource(in *config.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_GroupResource_To_v1_GroupResource(in, out, s)
}
func autoConvert_v1_HTPasswdPasswordIdentityProvider_To_config_HTPasswdPasswordIdentityProvider(in *v1.HTPasswdPasswordIdentityProvider, out *config.HTPasswdPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.File = in.File
	return nil
}
func Convert_v1_HTPasswdPasswordIdentityProvider_To_config_HTPasswdPasswordIdentityProvider(in *v1.HTPasswdPasswordIdentityProvider, out *config.HTPasswdPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_HTPasswdPasswordIdentityProvider_To_config_HTPasswdPasswordIdentityProvider(in, out, s)
}
func autoConvert_config_HTPasswdPasswordIdentityProvider_To_v1_HTPasswdPasswordIdentityProvider(in *config.HTPasswdPasswordIdentityProvider, out *v1.HTPasswdPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.File = in.File
	return nil
}
func Convert_config_HTPasswdPasswordIdentityProvider_To_v1_HTPasswdPasswordIdentityProvider(in *config.HTPasswdPasswordIdentityProvider, out *v1.HTPasswdPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_HTPasswdPasswordIdentityProvider_To_v1_HTPasswdPasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_HTTPServingInfo_To_config_HTTPServingInfo(in *v1.HTTPServingInfo, out *config.HTTPServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_ServingInfo_To_config_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.MaxRequestsInFlight = in.MaxRequestsInFlight
	out.RequestTimeoutSeconds = in.RequestTimeoutSeconds
	return nil
}
func Convert_v1_HTTPServingInfo_To_config_HTTPServingInfo(in *v1.HTTPServingInfo, out *config.HTTPServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_HTTPServingInfo_To_config_HTTPServingInfo(in, out, s)
}
func autoConvert_config_HTTPServingInfo_To_v1_HTTPServingInfo(in *config.HTTPServingInfo, out *v1.HTTPServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_ServingInfo_To_v1_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.MaxRequestsInFlight = in.MaxRequestsInFlight
	out.RequestTimeoutSeconds = in.RequestTimeoutSeconds
	return nil
}
func Convert_config_HTTPServingInfo_To_v1_HTTPServingInfo(in *config.HTTPServingInfo, out *v1.HTTPServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_HTTPServingInfo_To_v1_HTTPServingInfo(in, out, s)
}
func autoConvert_v1_IdentityProvider_To_config_IdentityProvider(in *v1.IdentityProvider, out *config.IdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.UseAsChallenger = in.UseAsChallenger
	out.UseAsLogin = in.UseAsLogin
	out.MappingMethod = in.MappingMethod
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Provider, &out.Provider, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_config_IdentityProvider_To_v1_IdentityProvider(in *config.IdentityProvider, out *v1.IdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.UseAsChallenger = in.UseAsChallenger
	out.UseAsLogin = in.UseAsLogin
	out.MappingMethod = in.MappingMethod
	return nil
}
func autoConvert_v1_ImageConfig_To_config_ImageConfig(in *v1.ImageConfig, out *config.ImageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Format = in.Format
	out.Latest = in.Latest
	return nil
}
func Convert_v1_ImageConfig_To_config_ImageConfig(in *v1.ImageConfig, out *config.ImageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageConfig_To_config_ImageConfig(in, out, s)
}
func autoConvert_config_ImageConfig_To_v1_ImageConfig(in *config.ImageConfig, out *v1.ImageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Format = in.Format
	out.Latest = in.Latest
	return nil
}
func Convert_config_ImageConfig_To_v1_ImageConfig(in *config.ImageConfig, out *v1.ImageConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ImageConfig_To_v1_ImageConfig(in, out, s)
}
func autoConvert_v1_ImagePolicyConfig_To_config_ImagePolicyConfig(in *v1.ImagePolicyConfig, out *config.ImagePolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.MaxImagesBulkImportedPerRepository = in.MaxImagesBulkImportedPerRepository
	out.DisableScheduledImport = in.DisableScheduledImport
	out.ScheduledImageImportMinimumIntervalSeconds = in.ScheduledImageImportMinimumIntervalSeconds
	out.MaxScheduledImageImportsPerMinute = in.MaxScheduledImageImportsPerMinute
	out.AllowedRegistriesForImport = (*config.AllowedRegistries)(unsafe.Pointer(in.AllowedRegistriesForImport))
	out.InternalRegistryHostname = in.InternalRegistryHostname
	out.AdditionalTrustedCA = in.AdditionalTrustedCA
	return nil
}
func autoConvert_config_ImagePolicyConfig_To_v1_ImagePolicyConfig(in *config.ImagePolicyConfig, out *v1.ImagePolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.MaxImagesBulkImportedPerRepository = in.MaxImagesBulkImportedPerRepository
	out.DisableScheduledImport = in.DisableScheduledImport
	out.ScheduledImageImportMinimumIntervalSeconds = in.ScheduledImageImportMinimumIntervalSeconds
	out.MaxScheduledImageImportsPerMinute = in.MaxScheduledImageImportsPerMinute
	out.AllowedRegistriesForImport = (*v1.AllowedRegistries)(unsafe.Pointer(in.AllowedRegistriesForImport))
	out.InternalRegistryHostname = in.InternalRegistryHostname
	out.AdditionalTrustedCA = in.AdditionalTrustedCA
	return nil
}
func autoConvert_v1_JenkinsPipelineConfig_To_config_JenkinsPipelineConfig(in *v1.JenkinsPipelineConfig, out *config.JenkinsPipelineConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AutoProvisionEnabled = (*bool)(unsafe.Pointer(in.AutoProvisionEnabled))
	out.TemplateNamespace = in.TemplateNamespace
	out.TemplateName = in.TemplateName
	out.ServiceName = in.ServiceName
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	return nil
}
func Convert_v1_JenkinsPipelineConfig_To_config_JenkinsPipelineConfig(in *v1.JenkinsPipelineConfig, out *config.JenkinsPipelineConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_JenkinsPipelineConfig_To_config_JenkinsPipelineConfig(in, out, s)
}
func autoConvert_config_JenkinsPipelineConfig_To_v1_JenkinsPipelineConfig(in *config.JenkinsPipelineConfig, out *v1.JenkinsPipelineConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AutoProvisionEnabled = (*bool)(unsafe.Pointer(in.AutoProvisionEnabled))
	out.TemplateNamespace = in.TemplateNamespace
	out.TemplateName = in.TemplateName
	out.ServiceName = in.ServiceName
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	return nil
}
func Convert_config_JenkinsPipelineConfig_To_v1_JenkinsPipelineConfig(in *config.JenkinsPipelineConfig, out *v1.JenkinsPipelineConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_JenkinsPipelineConfig_To_v1_JenkinsPipelineConfig(in, out, s)
}
func autoConvert_v1_KeystonePasswordIdentityProvider_To_config_KeystonePasswordIdentityProvider(in *v1.KeystonePasswordIdentityProvider, out *config.KeystonePasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_RemoteConnectionInfo_To_config_RemoteConnectionInfo(&in.RemoteConnectionInfo, &out.RemoteConnectionInfo, s); err != nil {
		return err
	}
	out.DomainName = in.DomainName
	out.UseKeystoneIdentity = in.UseKeystoneIdentity
	return nil
}
func Convert_v1_KeystonePasswordIdentityProvider_To_config_KeystonePasswordIdentityProvider(in *v1.KeystonePasswordIdentityProvider, out *config.KeystonePasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_KeystonePasswordIdentityProvider_To_config_KeystonePasswordIdentityProvider(in, out, s)
}
func autoConvert_config_KeystonePasswordIdentityProvider_To_v1_KeystonePasswordIdentityProvider(in *config.KeystonePasswordIdentityProvider, out *v1.KeystonePasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_RemoteConnectionInfo_To_v1_RemoteConnectionInfo(&in.RemoteConnectionInfo, &out.RemoteConnectionInfo, s); err != nil {
		return err
	}
	out.DomainName = in.DomainName
	out.UseKeystoneIdentity = in.UseKeystoneIdentity
	return nil
}
func Convert_config_KeystonePasswordIdentityProvider_To_v1_KeystonePasswordIdentityProvider(in *config.KeystonePasswordIdentityProvider, out *v1.KeystonePasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_KeystonePasswordIdentityProvider_To_v1_KeystonePasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_KubeletConnectionInfo_To_config_KubeletConnectionInfo(in *v1.KubeletConnectionInfo, out *config.KubeletConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Port = in.Port
	out.CA = in.CA
	return nil
}
func autoConvert_config_KubeletConnectionInfo_To_v1_KubeletConnectionInfo(in *config.KubeletConnectionInfo, out *v1.KubeletConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Port = in.Port
	out.CA = in.CA
	return nil
}
func autoConvert_v1_KubernetesMasterConfig_To_config_KubernetesMasterConfig(in *v1.KubernetesMasterConfig, out *config.KubernetesMasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DisabledAPIGroupVersions = *(*map[string][]string)(unsafe.Pointer(&in.DisabledAPIGroupVersions))
	out.MasterIP = in.MasterIP
	out.MasterEndpointReconcileTTL = in.MasterEndpointReconcileTTL
	out.ServicesSubnet = in.ServicesSubnet
	out.ServicesNodePortRange = in.ServicesNodePortRange
	out.SchedulerConfigFile = in.SchedulerConfigFile
	out.PodEvictionTimeout = in.PodEvictionTimeout
	if err := Convert_v1_CertInfo_To_config_CertInfo(&in.ProxyClientInfo, &out.ProxyClientInfo, s); err != nil {
		return err
	}
	out.APIServerArguments = *(*config.ExtendedArguments)(unsafe.Pointer(&in.APIServerArguments))
	out.ControllerArguments = *(*config.ExtendedArguments)(unsafe.Pointer(&in.ControllerArguments))
	out.SchedulerArguments = *(*config.ExtendedArguments)(unsafe.Pointer(&in.SchedulerArguments))
	return nil
}
func autoConvert_config_KubernetesMasterConfig_To_v1_KubernetesMasterConfig(in *config.KubernetesMasterConfig, out *v1.KubernetesMasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DisabledAPIGroupVersions = *(*map[string][]string)(unsafe.Pointer(&in.DisabledAPIGroupVersions))
	out.MasterIP = in.MasterIP
	out.MasterEndpointReconcileTTL = in.MasterEndpointReconcileTTL
	out.ServicesSubnet = in.ServicesSubnet
	out.ServicesNodePortRange = in.ServicesNodePortRange
	out.SchedulerConfigFile = in.SchedulerConfigFile
	out.PodEvictionTimeout = in.PodEvictionTimeout
	if err := Convert_config_CertInfo_To_v1_CertInfo(&in.ProxyClientInfo, &out.ProxyClientInfo, s); err != nil {
		return err
	}
	out.APIServerArguments = *(*v1.ExtendedArguments)(unsafe.Pointer(&in.APIServerArguments))
	out.ControllerArguments = *(*v1.ExtendedArguments)(unsafe.Pointer(&in.ControllerArguments))
	out.SchedulerArguments = *(*v1.ExtendedArguments)(unsafe.Pointer(&in.SchedulerArguments))
	return nil
}
func autoConvert_v1_LDAPAttributeMapping_To_config_LDAPAttributeMapping(in *v1.LDAPAttributeMapping, out *config.LDAPAttributeMapping, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ID = *(*[]string)(unsafe.Pointer(&in.ID))
	out.PreferredUsername = *(*[]string)(unsafe.Pointer(&in.PreferredUsername))
	out.Name = *(*[]string)(unsafe.Pointer(&in.Name))
	out.Email = *(*[]string)(unsafe.Pointer(&in.Email))
	return nil
}
func Convert_v1_LDAPAttributeMapping_To_config_LDAPAttributeMapping(in *v1.LDAPAttributeMapping, out *config.LDAPAttributeMapping, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LDAPAttributeMapping_To_config_LDAPAttributeMapping(in, out, s)
}
func autoConvert_config_LDAPAttributeMapping_To_v1_LDAPAttributeMapping(in *config.LDAPAttributeMapping, out *v1.LDAPAttributeMapping, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ID = *(*[]string)(unsafe.Pointer(&in.ID))
	out.PreferredUsername = *(*[]string)(unsafe.Pointer(&in.PreferredUsername))
	out.Name = *(*[]string)(unsafe.Pointer(&in.Name))
	out.Email = *(*[]string)(unsafe.Pointer(&in.Email))
	return nil
}
func Convert_config_LDAPAttributeMapping_To_v1_LDAPAttributeMapping(in *config.LDAPAttributeMapping, out *v1.LDAPAttributeMapping, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_LDAPAttributeMapping_To_v1_LDAPAttributeMapping(in, out, s)
}
func autoConvert_v1_LDAPPasswordIdentityProvider_To_config_LDAPPasswordIdentityProvider(in *v1.LDAPPasswordIdentityProvider, out *config.LDAPPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.BindDN = in.BindDN
	if err := Convert_v1_StringSource_To_config_StringSource(&in.BindPassword, &out.BindPassword, s); err != nil {
		return err
	}
	out.Insecure = in.Insecure
	out.CA = in.CA
	if err := Convert_v1_LDAPAttributeMapping_To_config_LDAPAttributeMapping(&in.Attributes, &out.Attributes, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_LDAPPasswordIdentityProvider_To_config_LDAPPasswordIdentityProvider(in *v1.LDAPPasswordIdentityProvider, out *config.LDAPPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LDAPPasswordIdentityProvider_To_config_LDAPPasswordIdentityProvider(in, out, s)
}
func autoConvert_config_LDAPPasswordIdentityProvider_To_v1_LDAPPasswordIdentityProvider(in *config.LDAPPasswordIdentityProvider, out *v1.LDAPPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.BindDN = in.BindDN
	if err := Convert_config_StringSource_To_v1_StringSource(&in.BindPassword, &out.BindPassword, s); err != nil {
		return err
	}
	out.Insecure = in.Insecure
	out.CA = in.CA
	if err := Convert_config_LDAPAttributeMapping_To_v1_LDAPAttributeMapping(&in.Attributes, &out.Attributes, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_LDAPPasswordIdentityProvider_To_v1_LDAPPasswordIdentityProvider(in *config.LDAPPasswordIdentityProvider, out *v1.LDAPPasswordIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_LDAPPasswordIdentityProvider_To_v1_LDAPPasswordIdentityProvider(in, out, s)
}
func autoConvert_v1_LDAPQuery_To_config_LDAPQuery(in *v1.LDAPQuery, out *config.LDAPQuery, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BaseDN = in.BaseDN
	out.Scope = in.Scope
	out.DerefAliases = in.DerefAliases
	out.TimeLimit = in.TimeLimit
	out.Filter = in.Filter
	out.PageSize = in.PageSize
	return nil
}
func Convert_v1_LDAPQuery_To_config_LDAPQuery(in *v1.LDAPQuery, out *config.LDAPQuery, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LDAPQuery_To_config_LDAPQuery(in, out, s)
}
func autoConvert_config_LDAPQuery_To_v1_LDAPQuery(in *config.LDAPQuery, out *v1.LDAPQuery, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BaseDN = in.BaseDN
	out.Scope = in.Scope
	out.DerefAliases = in.DerefAliases
	out.TimeLimit = in.TimeLimit
	out.Filter = in.Filter
	out.PageSize = in.PageSize
	return nil
}
func Convert_config_LDAPQuery_To_v1_LDAPQuery(in *config.LDAPQuery, out *v1.LDAPQuery, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_LDAPQuery_To_v1_LDAPQuery(in, out, s)
}
func autoConvert_v1_LDAPSyncConfig_To_config_LDAPSyncConfig(in *v1.LDAPSyncConfig, out *config.LDAPSyncConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.BindDN = in.BindDN
	if err := Convert_v1_StringSource_To_config_StringSource(&in.BindPassword, &out.BindPassword, s); err != nil {
		return err
	}
	out.Insecure = in.Insecure
	out.CA = in.CA
	out.LDAPGroupUIDToOpenShiftGroupNameMapping = *(*map[string]string)(unsafe.Pointer(&in.LDAPGroupUIDToOpenShiftGroupNameMapping))
	out.RFC2307Config = (*config.RFC2307Config)(unsafe.Pointer(in.RFC2307Config))
	out.ActiveDirectoryConfig = (*config.ActiveDirectoryConfig)(unsafe.Pointer(in.ActiveDirectoryConfig))
	out.AugmentedActiveDirectoryConfig = (*config.AugmentedActiveDirectoryConfig)(unsafe.Pointer(in.AugmentedActiveDirectoryConfig))
	return nil
}
func Convert_v1_LDAPSyncConfig_To_config_LDAPSyncConfig(in *v1.LDAPSyncConfig, out *config.LDAPSyncConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LDAPSyncConfig_To_config_LDAPSyncConfig(in, out, s)
}
func autoConvert_config_LDAPSyncConfig_To_v1_LDAPSyncConfig(in *config.LDAPSyncConfig, out *v1.LDAPSyncConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.BindDN = in.BindDN
	if err := Convert_config_StringSource_To_v1_StringSource(&in.BindPassword, &out.BindPassword, s); err != nil {
		return err
	}
	out.Insecure = in.Insecure
	out.CA = in.CA
	out.LDAPGroupUIDToOpenShiftGroupNameMapping = *(*map[string]string)(unsafe.Pointer(&in.LDAPGroupUIDToOpenShiftGroupNameMapping))
	out.RFC2307Config = (*v1.RFC2307Config)(unsafe.Pointer(in.RFC2307Config))
	out.ActiveDirectoryConfig = (*v1.ActiveDirectoryConfig)(unsafe.Pointer(in.ActiveDirectoryConfig))
	out.AugmentedActiveDirectoryConfig = (*v1.AugmentedActiveDirectoryConfig)(unsafe.Pointer(in.AugmentedActiveDirectoryConfig))
	return nil
}
func Convert_config_LDAPSyncConfig_To_v1_LDAPSyncConfig(in *config.LDAPSyncConfig, out *v1.LDAPSyncConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_LDAPSyncConfig_To_v1_LDAPSyncConfig(in, out, s)
}
func autoConvert_v1_LocalQuota_To_config_LocalQuota(in *v1.LocalQuota, out *config.LocalQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.PerFSGroup = (*resource.Quantity)(unsafe.Pointer(in.PerFSGroup))
	return nil
}
func Convert_v1_LocalQuota_To_config_LocalQuota(in *v1.LocalQuota, out *config.LocalQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LocalQuota_To_config_LocalQuota(in, out, s)
}
func autoConvert_config_LocalQuota_To_v1_LocalQuota(in *config.LocalQuota, out *v1.LocalQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.PerFSGroup = (*resource.Quantity)(unsafe.Pointer(in.PerFSGroup))
	return nil
}
func Convert_config_LocalQuota_To_v1_LocalQuota(in *config.LocalQuota, out *v1.LocalQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_LocalQuota_To_v1_LocalQuota(in, out, s)
}
func autoConvert_v1_MasterAuthConfig_To_config_MasterAuthConfig(in *v1.MasterAuthConfig, out *config.MasterAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RequestHeader = (*config.RequestHeaderAuthenticationOptions)(unsafe.Pointer(in.RequestHeader))
	out.WebhookTokenAuthenticators = *(*[]config.WebhookTokenAuthenticator)(unsafe.Pointer(&in.WebhookTokenAuthenticators))
	out.OAuthMetadataFile = in.OAuthMetadataFile
	return nil
}
func Convert_v1_MasterAuthConfig_To_config_MasterAuthConfig(in *v1.MasterAuthConfig, out *config.MasterAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_MasterAuthConfig_To_config_MasterAuthConfig(in, out, s)
}
func autoConvert_config_MasterAuthConfig_To_v1_MasterAuthConfig(in *config.MasterAuthConfig, out *v1.MasterAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RequestHeader = (*v1.RequestHeaderAuthenticationOptions)(unsafe.Pointer(in.RequestHeader))
	out.WebhookTokenAuthenticators = *(*[]v1.WebhookTokenAuthenticator)(unsafe.Pointer(&in.WebhookTokenAuthenticators))
	out.OAuthMetadataFile = in.OAuthMetadataFile
	return nil
}
func Convert_config_MasterAuthConfig_To_v1_MasterAuthConfig(in *config.MasterAuthConfig, out *v1.MasterAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_MasterAuthConfig_To_v1_MasterAuthConfig(in, out, s)
}
func autoConvert_v1_MasterClients_To_config_MasterClients(in *v1.MasterClients, out *config.MasterClients, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.OpenShiftLoopbackKubeConfig = in.OpenShiftLoopbackKubeConfig
	out.OpenShiftLoopbackClientConnectionOverrides = (*config.ClientConnectionOverrides)(unsafe.Pointer(in.OpenShiftLoopbackClientConnectionOverrides))
	return nil
}
func Convert_v1_MasterClients_To_config_MasterClients(in *v1.MasterClients, out *config.MasterClients, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_MasterClients_To_config_MasterClients(in, out, s)
}
func autoConvert_config_MasterClients_To_v1_MasterClients(in *config.MasterClients, out *v1.MasterClients, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.OpenShiftLoopbackKubeConfig = in.OpenShiftLoopbackKubeConfig
	out.OpenShiftLoopbackClientConnectionOverrides = (*v1.ClientConnectionOverrides)(unsafe.Pointer(in.OpenShiftLoopbackClientConnectionOverrides))
	return nil
}
func Convert_config_MasterClients_To_v1_MasterClients(in *config.MasterClients, out *v1.MasterClients, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_MasterClients_To_v1_MasterClients(in, out, s)
}
func autoConvert_v1_MasterConfig_To_config_MasterConfig(in *v1.MasterConfig, out *config.MasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_HTTPServingInfo_To_config_HTTPServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	if err := Convert_v1_MasterAuthConfig_To_config_MasterAuthConfig(&in.AuthConfig, &out.AuthConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_AggregatorConfig_To_config_AggregatorConfig(&in.AggregatorConfig, &out.AggregatorConfig, s); err != nil {
		return err
	}
	out.CORSAllowedOrigins = *(*[]string)(unsafe.Pointer(&in.CORSAllowedOrigins))
	out.APILevels = *(*[]string)(unsafe.Pointer(&in.APILevels))
	out.MasterPublicURL = in.MasterPublicURL
	out.Controllers = in.Controllers
	if err := Convert_v1_AdmissionConfig_To_config_AdmissionConfig(&in.AdmissionConfig, &out.AdmissionConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_ControllerConfig_To_config_ControllerConfig(&in.ControllerConfig, &out.ControllerConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_EtcdStorageConfig_To_config_EtcdStorageConfig(&in.EtcdStorageConfig, &out.EtcdStorageConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_EtcdConnectionInfo_To_config_EtcdConnectionInfo(&in.EtcdClientInfo, &out.EtcdClientInfo, s); err != nil {
		return err
	}
	if err := Convert_v1_KubeletConnectionInfo_To_config_KubeletConnectionInfo(&in.KubeletClientInfo, &out.KubeletClientInfo, s); err != nil {
		return err
	}
	if err := Convert_v1_KubernetesMasterConfig_To_config_KubernetesMasterConfig(&in.KubernetesMasterConfig, &out.KubernetesMasterConfig, s); err != nil {
		return err
	}
	if in.EtcdConfig != nil {
		in, out := &in.EtcdConfig, &out.EtcdConfig
		*out = new(config.EtcdConfig)
		if err := Convert_v1_EtcdConfig_To_config_EtcdConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.EtcdConfig = nil
	}
	if in.OAuthConfig != nil {
		in, out := &in.OAuthConfig, &out.OAuthConfig
		*out = new(config.OAuthConfig)
		if err := Convert_v1_OAuthConfig_To_config_OAuthConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OAuthConfig = nil
	}
	out.DNSConfig = (*config.DNSConfig)(unsafe.Pointer(in.DNSConfig))
	if err := Convert_v1_ServiceAccountConfig_To_config_ServiceAccountConfig(&in.ServiceAccountConfig, &out.ServiceAccountConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_MasterClients_To_config_MasterClients(&in.MasterClients, &out.MasterClients, s); err != nil {
		return err
	}
	if err := Convert_v1_ImageConfig_To_config_ImageConfig(&in.ImageConfig, &out.ImageConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_ImagePolicyConfig_To_config_ImagePolicyConfig(&in.ImagePolicyConfig, &out.ImagePolicyConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_PolicyConfig_To_config_PolicyConfig(&in.PolicyConfig, &out.PolicyConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectConfig_To_config_ProjectConfig(&in.ProjectConfig, &out.ProjectConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_RoutingConfig_To_config_RoutingConfig(&in.RoutingConfig, &out.RoutingConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_MasterNetworkConfig_To_config_MasterNetworkConfig(&in.NetworkConfig, &out.NetworkConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_MasterVolumeConfig_To_config_MasterVolumeConfig(&in.VolumeConfig, &out.VolumeConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_JenkinsPipelineConfig_To_config_JenkinsPipelineConfig(&in.JenkinsPipelineConfig, &out.JenkinsPipelineConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_AuditConfig_To_config_AuditConfig(&in.AuditConfig, &out.AuditConfig, s); err != nil {
		return err
	}
	out.DisableOpenAPI = in.DisableOpenAPI
	return nil
}
func Convert_v1_MasterConfig_To_config_MasterConfig(in *v1.MasterConfig, out *config.MasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_MasterConfig_To_config_MasterConfig(in, out, s)
}
func autoConvert_config_MasterConfig_To_v1_MasterConfig(in *config.MasterConfig, out *v1.MasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_HTTPServingInfo_To_v1_HTTPServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	if err := Convert_config_MasterAuthConfig_To_v1_MasterAuthConfig(&in.AuthConfig, &out.AuthConfig, s); err != nil {
		return err
	}
	if err := Convert_config_AggregatorConfig_To_v1_AggregatorConfig(&in.AggregatorConfig, &out.AggregatorConfig, s); err != nil {
		return err
	}
	out.CORSAllowedOrigins = *(*[]string)(unsafe.Pointer(&in.CORSAllowedOrigins))
	out.APILevels = *(*[]string)(unsafe.Pointer(&in.APILevels))
	out.MasterPublicURL = in.MasterPublicURL
	if err := Convert_config_AdmissionConfig_To_v1_AdmissionConfig(&in.AdmissionConfig, &out.AdmissionConfig, s); err != nil {
		return err
	}
	out.Controllers = in.Controllers
	if err := Convert_config_ControllerConfig_To_v1_ControllerConfig(&in.ControllerConfig, &out.ControllerConfig, s); err != nil {
		return err
	}
	if err := Convert_config_EtcdStorageConfig_To_v1_EtcdStorageConfig(&in.EtcdStorageConfig, &out.EtcdStorageConfig, s); err != nil {
		return err
	}
	if err := Convert_config_EtcdConnectionInfo_To_v1_EtcdConnectionInfo(&in.EtcdClientInfo, &out.EtcdClientInfo, s); err != nil {
		return err
	}
	if err := Convert_config_KubeletConnectionInfo_To_v1_KubeletConnectionInfo(&in.KubeletClientInfo, &out.KubeletClientInfo, s); err != nil {
		return err
	}
	if err := Convert_config_KubernetesMasterConfig_To_v1_KubernetesMasterConfig(&in.KubernetesMasterConfig, &out.KubernetesMasterConfig, s); err != nil {
		return err
	}
	if in.EtcdConfig != nil {
		in, out := &in.EtcdConfig, &out.EtcdConfig
		*out = new(v1.EtcdConfig)
		if err := Convert_config_EtcdConfig_To_v1_EtcdConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.EtcdConfig = nil
	}
	if in.OAuthConfig != nil {
		in, out := &in.OAuthConfig, &out.OAuthConfig
		*out = new(v1.OAuthConfig)
		if err := Convert_config_OAuthConfig_To_v1_OAuthConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OAuthConfig = nil
	}
	out.DNSConfig = (*v1.DNSConfig)(unsafe.Pointer(in.DNSConfig))
	if err := Convert_config_ServiceAccountConfig_To_v1_ServiceAccountConfig(&in.ServiceAccountConfig, &out.ServiceAccountConfig, s); err != nil {
		return err
	}
	if err := Convert_config_MasterClients_To_v1_MasterClients(&in.MasterClients, &out.MasterClients, s); err != nil {
		return err
	}
	if err := Convert_config_ImageConfig_To_v1_ImageConfig(&in.ImageConfig, &out.ImageConfig, s); err != nil {
		return err
	}
	if err := Convert_config_ImagePolicyConfig_To_v1_ImagePolicyConfig(&in.ImagePolicyConfig, &out.ImagePolicyConfig, s); err != nil {
		return err
	}
	if err := Convert_config_PolicyConfig_To_v1_PolicyConfig(&in.PolicyConfig, &out.PolicyConfig, s); err != nil {
		return err
	}
	if err := Convert_config_ProjectConfig_To_v1_ProjectConfig(&in.ProjectConfig, &out.ProjectConfig, s); err != nil {
		return err
	}
	if err := Convert_config_RoutingConfig_To_v1_RoutingConfig(&in.RoutingConfig, &out.RoutingConfig, s); err != nil {
		return err
	}
	if err := Convert_config_MasterNetworkConfig_To_v1_MasterNetworkConfig(&in.NetworkConfig, &out.NetworkConfig, s); err != nil {
		return err
	}
	if err := Convert_config_MasterVolumeConfig_To_v1_MasterVolumeConfig(&in.VolumeConfig, &out.VolumeConfig, s); err != nil {
		return err
	}
	if err := Convert_config_JenkinsPipelineConfig_To_v1_JenkinsPipelineConfig(&in.JenkinsPipelineConfig, &out.JenkinsPipelineConfig, s); err != nil {
		return err
	}
	if err := Convert_config_AuditConfig_To_v1_AuditConfig(&in.AuditConfig, &out.AuditConfig, s); err != nil {
		return err
	}
	out.DisableOpenAPI = in.DisableOpenAPI
	return nil
}
func Convert_config_MasterConfig_To_v1_MasterConfig(in *config.MasterConfig, out *v1.MasterConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_MasterConfig_To_v1_MasterConfig(in, out, s)
}
func autoConvert_v1_MasterNetworkConfig_To_config_MasterNetworkConfig(in *v1.MasterNetworkConfig, out *config.MasterNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NetworkPluginName = in.NetworkPluginName
	out.DeprecatedClusterNetworkCIDR = in.DeprecatedClusterNetworkCIDR
	out.ClusterNetworks = *(*[]config.ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterNetworks))
	out.DeprecatedHostSubnetLength = in.DeprecatedHostSubnetLength
	out.ServiceNetworkCIDR = in.ServiceNetworkCIDR
	out.ExternalIPNetworkCIDRs = *(*[]string)(unsafe.Pointer(&in.ExternalIPNetworkCIDRs))
	out.IngressIPNetworkCIDR = in.IngressIPNetworkCIDR
	out.VXLANPort = in.VXLANPort
	return nil
}
func autoConvert_config_MasterNetworkConfig_To_v1_MasterNetworkConfig(in *config.MasterNetworkConfig, out *v1.MasterNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NetworkPluginName = in.NetworkPluginName
	out.DeprecatedClusterNetworkCIDR = in.DeprecatedClusterNetworkCIDR
	out.ClusterNetworks = *(*[]v1.ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterNetworks))
	out.DeprecatedHostSubnetLength = in.DeprecatedHostSubnetLength
	out.ServiceNetworkCIDR = in.ServiceNetworkCIDR
	out.ExternalIPNetworkCIDRs = *(*[]string)(unsafe.Pointer(&in.ExternalIPNetworkCIDRs))
	out.IngressIPNetworkCIDR = in.IngressIPNetworkCIDR
	out.VXLANPort = in.VXLANPort
	return nil
}
func Convert_config_MasterNetworkConfig_To_v1_MasterNetworkConfig(in *config.MasterNetworkConfig, out *v1.MasterNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_MasterNetworkConfig_To_v1_MasterNetworkConfig(in, out, s)
}
func autoConvert_v1_MasterVolumeConfig_To_config_MasterVolumeConfig(in *v1.MasterVolumeConfig, out *config.MasterVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := metav1.Convert_Pointer_bool_To_bool(&in.DynamicProvisioningEnabled, &out.DynamicProvisioningEnabled, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_config_MasterVolumeConfig_To_v1_MasterVolumeConfig(in *config.MasterVolumeConfig, out *v1.MasterVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := metav1.Convert_bool_To_Pointer_bool(&in.DynamicProvisioningEnabled, &out.DynamicProvisioningEnabled, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_NamedCertificate_To_config_NamedCertificate(in *v1.NamedCertificate, out *config.NamedCertificate, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Names = *(*[]string)(unsafe.Pointer(&in.Names))
	if err := Convert_v1_CertInfo_To_config_CertInfo(&in.CertInfo, &out.CertInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_NamedCertificate_To_config_NamedCertificate(in *v1.NamedCertificate, out *config.NamedCertificate, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NamedCertificate_To_config_NamedCertificate(in, out, s)
}
func autoConvert_config_NamedCertificate_To_v1_NamedCertificate(in *config.NamedCertificate, out *v1.NamedCertificate, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Names = *(*[]string)(unsafe.Pointer(&in.Names))
	if err := Convert_config_CertInfo_To_v1_CertInfo(&in.CertInfo, &out.CertInfo, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_NamedCertificate_To_v1_NamedCertificate(in *config.NamedCertificate, out *v1.NamedCertificate, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_NamedCertificate_To_v1_NamedCertificate(in, out, s)
}
func autoConvert_v1_NodeAuthConfig_To_config_NodeAuthConfig(in *v1.NodeAuthConfig, out *config.NodeAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AuthenticationCacheTTL = in.AuthenticationCacheTTL
	out.AuthenticationCacheSize = in.AuthenticationCacheSize
	out.AuthorizationCacheTTL = in.AuthorizationCacheTTL
	out.AuthorizationCacheSize = in.AuthorizationCacheSize
	return nil
}
func Convert_v1_NodeAuthConfig_To_config_NodeAuthConfig(in *v1.NodeAuthConfig, out *config.NodeAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NodeAuthConfig_To_config_NodeAuthConfig(in, out, s)
}
func autoConvert_config_NodeAuthConfig_To_v1_NodeAuthConfig(in *config.NodeAuthConfig, out *v1.NodeAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AuthenticationCacheTTL = in.AuthenticationCacheTTL
	out.AuthenticationCacheSize = in.AuthenticationCacheSize
	out.AuthorizationCacheTTL = in.AuthorizationCacheTTL
	out.AuthorizationCacheSize = in.AuthorizationCacheSize
	return nil
}
func Convert_config_NodeAuthConfig_To_v1_NodeAuthConfig(in *config.NodeAuthConfig, out *v1.NodeAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_NodeAuthConfig_To_v1_NodeAuthConfig(in, out, s)
}
func autoConvert_v1_NodeConfig_To_config_NodeConfig(in *v1.NodeConfig, out *config.NodeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NodeName = in.NodeName
	out.NodeIP = in.NodeIP
	if err := Convert_v1_ServingInfo_To_config_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.MasterKubeConfig = in.MasterKubeConfig
	out.MasterClientConnectionOverrides = (*config.ClientConnectionOverrides)(unsafe.Pointer(in.MasterClientConnectionOverrides))
	out.DNSDomain = in.DNSDomain
	out.DNSIP = in.DNSIP
	out.DNSBindAddress = in.DNSBindAddress
	out.DNSNameservers = *(*[]string)(unsafe.Pointer(&in.DNSNameservers))
	out.DNSRecursiveResolvConf = in.DNSRecursiveResolvConf
	if err := Convert_v1_NodeNetworkConfig_To_config_NodeNetworkConfig(&in.NetworkConfig, &out.NetworkConfig, s); err != nil {
		return err
	}
	out.VolumeDirectory = in.VolumeDirectory
	if err := Convert_v1_ImageConfig_To_config_ImageConfig(&in.ImageConfig, &out.ImageConfig, s); err != nil {
		return err
	}
	out.AllowDisabledDocker = in.AllowDisabledDocker
	out.PodManifestConfig = (*config.PodManifestConfig)(unsafe.Pointer(in.PodManifestConfig))
	if err := Convert_v1_NodeAuthConfig_To_config_NodeAuthConfig(&in.AuthConfig, &out.AuthConfig, s); err != nil {
		return err
	}
	if err := Convert_v1_DockerConfig_To_config_DockerConfig(&in.DockerConfig, &out.DockerConfig, s); err != nil {
		return err
	}
	out.KubeletArguments = *(*config.ExtendedArguments)(unsafe.Pointer(&in.KubeletArguments))
	out.ProxyArguments = *(*config.ExtendedArguments)(unsafe.Pointer(&in.ProxyArguments))
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	if err := metav1.Convert_Pointer_bool_To_bool(&in.EnableUnidling, &out.EnableUnidling, s); err != nil {
		return err
	}
	if err := Convert_v1_NodeVolumeConfig_To_config_NodeVolumeConfig(&in.VolumeConfig, &out.VolumeConfig, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_config_NodeConfig_To_v1_NodeConfig(in *config.NodeConfig, out *v1.NodeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NodeName = in.NodeName
	out.NodeIP = in.NodeIP
	if err := Convert_config_ServingInfo_To_v1_ServingInfo(&in.ServingInfo, &out.ServingInfo, s); err != nil {
		return err
	}
	out.MasterKubeConfig = in.MasterKubeConfig
	out.MasterClientConnectionOverrides = (*v1.ClientConnectionOverrides)(unsafe.Pointer(in.MasterClientConnectionOverrides))
	out.DNSDomain = in.DNSDomain
	out.DNSIP = in.DNSIP
	out.DNSBindAddress = in.DNSBindAddress
	out.DNSNameservers = *(*[]string)(unsafe.Pointer(&in.DNSNameservers))
	out.DNSRecursiveResolvConf = in.DNSRecursiveResolvConf
	if err := Convert_config_NodeNetworkConfig_To_v1_NodeNetworkConfig(&in.NetworkConfig, &out.NetworkConfig, s); err != nil {
		return err
	}
	out.VolumeDirectory = in.VolumeDirectory
	if err := Convert_config_ImageConfig_To_v1_ImageConfig(&in.ImageConfig, &out.ImageConfig, s); err != nil {
		return err
	}
	out.AllowDisabledDocker = in.AllowDisabledDocker
	out.PodManifestConfig = (*v1.PodManifestConfig)(unsafe.Pointer(in.PodManifestConfig))
	if err := Convert_config_NodeAuthConfig_To_v1_NodeAuthConfig(&in.AuthConfig, &out.AuthConfig, s); err != nil {
		return err
	}
	if err := Convert_config_DockerConfig_To_v1_DockerConfig(&in.DockerConfig, &out.DockerConfig, s); err != nil {
		return err
	}
	out.KubeletArguments = *(*v1.ExtendedArguments)(unsafe.Pointer(&in.KubeletArguments))
	out.ProxyArguments = *(*v1.ExtendedArguments)(unsafe.Pointer(&in.ProxyArguments))
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	if err := metav1.Convert_bool_To_Pointer_bool(&in.EnableUnidling, &out.EnableUnidling, s); err != nil {
		return err
	}
	if err := Convert_config_NodeVolumeConfig_To_v1_NodeVolumeConfig(&in.VolumeConfig, &out.VolumeConfig, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_NodeNetworkConfig_To_config_NodeNetworkConfig(in *v1.NodeNetworkConfig, out *config.NodeNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NetworkPluginName = in.NetworkPluginName
	out.MTU = in.MTU
	return nil
}
func Convert_v1_NodeNetworkConfig_To_config_NodeNetworkConfig(in *v1.NodeNetworkConfig, out *config.NodeNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NodeNetworkConfig_To_config_NodeNetworkConfig(in, out, s)
}
func autoConvert_config_NodeNetworkConfig_To_v1_NodeNetworkConfig(in *config.NodeNetworkConfig, out *v1.NodeNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.NetworkPluginName = in.NetworkPluginName
	out.MTU = in.MTU
	return nil
}
func Convert_config_NodeNetworkConfig_To_v1_NodeNetworkConfig(in *config.NodeNetworkConfig, out *v1.NodeNetworkConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_NodeNetworkConfig_To_v1_NodeNetworkConfig(in, out, s)
}
func autoConvert_v1_NodeVolumeConfig_To_config_NodeVolumeConfig(in *v1.NodeVolumeConfig, out *config.NodeVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_LocalQuota_To_config_LocalQuota(&in.LocalQuota, &out.LocalQuota, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_NodeVolumeConfig_To_config_NodeVolumeConfig(in *v1.NodeVolumeConfig, out *config.NodeVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NodeVolumeConfig_To_config_NodeVolumeConfig(in, out, s)
}
func autoConvert_config_NodeVolumeConfig_To_v1_NodeVolumeConfig(in *config.NodeVolumeConfig, out *v1.NodeVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_LocalQuota_To_v1_LocalQuota(&in.LocalQuota, &out.LocalQuota, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_NodeVolumeConfig_To_v1_NodeVolumeConfig(in *config.NodeVolumeConfig, out *v1.NodeVolumeConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_NodeVolumeConfig_To_v1_NodeVolumeConfig(in, out, s)
}
func autoConvert_v1_OAuthConfig_To_config_OAuthConfig(in *v1.OAuthConfig, out *config.OAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.MasterCA = (*string)(unsafe.Pointer(in.MasterCA))
	out.MasterURL = in.MasterURL
	out.MasterPublicURL = in.MasterPublicURL
	out.AssetPublicURL = in.AssetPublicURL
	out.AlwaysShowProviderSelection = in.AlwaysShowProviderSelection
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]config.IdentityProvider, len(*in))
		for i := range *in {
			if err := Convert_v1_IdentityProvider_To_config_IdentityProvider(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.IdentityProviders = nil
	}
	if err := Convert_v1_GrantConfig_To_config_GrantConfig(&in.GrantConfig, &out.GrantConfig, s); err != nil {
		return err
	}
	out.SessionConfig = (*config.SessionConfig)(unsafe.Pointer(in.SessionConfig))
	if err := Convert_v1_TokenConfig_To_config_TokenConfig(&in.TokenConfig, &out.TokenConfig, s); err != nil {
		return err
	}
	out.Templates = (*config.OAuthTemplates)(unsafe.Pointer(in.Templates))
	return nil
}
func Convert_v1_OAuthConfig_To_config_OAuthConfig(in *v1.OAuthConfig, out *config.OAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthConfig_To_config_OAuthConfig(in, out, s)
}
func autoConvert_config_OAuthConfig_To_v1_OAuthConfig(in *config.OAuthConfig, out *v1.OAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.MasterCA = (*string)(unsafe.Pointer(in.MasterCA))
	out.MasterURL = in.MasterURL
	out.MasterPublicURL = in.MasterPublicURL
	out.AssetPublicURL = in.AssetPublicURL
	out.AlwaysShowProviderSelection = in.AlwaysShowProviderSelection
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]v1.IdentityProvider, len(*in))
		for i := range *in {
			if err := Convert_config_IdentityProvider_To_v1_IdentityProvider(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.IdentityProviders = nil
	}
	if err := Convert_config_GrantConfig_To_v1_GrantConfig(&in.GrantConfig, &out.GrantConfig, s); err != nil {
		return err
	}
	out.SessionConfig = (*v1.SessionConfig)(unsafe.Pointer(in.SessionConfig))
	if err := Convert_config_TokenConfig_To_v1_TokenConfig(&in.TokenConfig, &out.TokenConfig, s); err != nil {
		return err
	}
	out.Templates = (*v1.OAuthTemplates)(unsafe.Pointer(in.Templates))
	return nil
}
func Convert_config_OAuthConfig_To_v1_OAuthConfig(in *config.OAuthConfig, out *v1.OAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_OAuthConfig_To_v1_OAuthConfig(in, out, s)
}
func autoConvert_v1_OAuthTemplates_To_config_OAuthTemplates(in *v1.OAuthTemplates, out *config.OAuthTemplates, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Login = in.Login
	out.ProviderSelection = in.ProviderSelection
	out.Error = in.Error
	return nil
}
func Convert_v1_OAuthTemplates_To_config_OAuthTemplates(in *v1.OAuthTemplates, out *config.OAuthTemplates, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthTemplates_To_config_OAuthTemplates(in, out, s)
}
func autoConvert_config_OAuthTemplates_To_v1_OAuthTemplates(in *config.OAuthTemplates, out *v1.OAuthTemplates, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Login = in.Login
	out.ProviderSelection = in.ProviderSelection
	out.Error = in.Error
	return nil
}
func Convert_config_OAuthTemplates_To_v1_OAuthTemplates(in *config.OAuthTemplates, out *v1.OAuthTemplates, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_OAuthTemplates_To_v1_OAuthTemplates(in, out, s)
}
func autoConvert_v1_OpenIDClaims_To_config_OpenIDClaims(in *v1.OpenIDClaims, out *config.OpenIDClaims, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ID = *(*[]string)(unsafe.Pointer(&in.ID))
	out.PreferredUsername = *(*[]string)(unsafe.Pointer(&in.PreferredUsername))
	out.Name = *(*[]string)(unsafe.Pointer(&in.Name))
	out.Email = *(*[]string)(unsafe.Pointer(&in.Email))
	return nil
}
func Convert_v1_OpenIDClaims_To_config_OpenIDClaims(in *v1.OpenIDClaims, out *config.OpenIDClaims, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OpenIDClaims_To_config_OpenIDClaims(in, out, s)
}
func autoConvert_config_OpenIDClaims_To_v1_OpenIDClaims(in *config.OpenIDClaims, out *v1.OpenIDClaims, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ID = *(*[]string)(unsafe.Pointer(&in.ID))
	out.PreferredUsername = *(*[]string)(unsafe.Pointer(&in.PreferredUsername))
	out.Name = *(*[]string)(unsafe.Pointer(&in.Name))
	out.Email = *(*[]string)(unsafe.Pointer(&in.Email))
	return nil
}
func Convert_config_OpenIDClaims_To_v1_OpenIDClaims(in *config.OpenIDClaims, out *v1.OpenIDClaims, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_OpenIDClaims_To_v1_OpenIDClaims(in, out, s)
}
func autoConvert_v1_OpenIDIdentityProvider_To_config_OpenIDIdentityProvider(in *v1.OpenIDIdentityProvider, out *config.OpenIDIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CA = in.CA
	out.ClientID = in.ClientID
	if err := Convert_v1_StringSource_To_config_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.ExtraScopes = *(*[]string)(unsafe.Pointer(&in.ExtraScopes))
	out.ExtraAuthorizeParameters = *(*map[string]string)(unsafe.Pointer(&in.ExtraAuthorizeParameters))
	if err := Convert_v1_OpenIDURLs_To_config_OpenIDURLs(&in.URLs, &out.URLs, s); err != nil {
		return err
	}
	if err := Convert_v1_OpenIDClaims_To_config_OpenIDClaims(&in.Claims, &out.Claims, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_OpenIDIdentityProvider_To_config_OpenIDIdentityProvider(in *v1.OpenIDIdentityProvider, out *config.OpenIDIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OpenIDIdentityProvider_To_config_OpenIDIdentityProvider(in, out, s)
}
func autoConvert_config_OpenIDIdentityProvider_To_v1_OpenIDIdentityProvider(in *config.OpenIDIdentityProvider, out *v1.OpenIDIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CA = in.CA
	out.ClientID = in.ClientID
	if err := Convert_config_StringSource_To_v1_StringSource(&in.ClientSecret, &out.ClientSecret, s); err != nil {
		return err
	}
	out.ExtraScopes = *(*[]string)(unsafe.Pointer(&in.ExtraScopes))
	out.ExtraAuthorizeParameters = *(*map[string]string)(unsafe.Pointer(&in.ExtraAuthorizeParameters))
	if err := Convert_config_OpenIDURLs_To_v1_OpenIDURLs(&in.URLs, &out.URLs, s); err != nil {
		return err
	}
	if err := Convert_config_OpenIDClaims_To_v1_OpenIDClaims(&in.Claims, &out.Claims, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_OpenIDIdentityProvider_To_v1_OpenIDIdentityProvider(in *config.OpenIDIdentityProvider, out *v1.OpenIDIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_OpenIDIdentityProvider_To_v1_OpenIDIdentityProvider(in, out, s)
}
func autoConvert_v1_OpenIDURLs_To_config_OpenIDURLs(in *v1.OpenIDURLs, out *config.OpenIDURLs, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Authorize = in.Authorize
	out.Token = in.Token
	out.UserInfo = in.UserInfo
	return nil
}
func Convert_v1_OpenIDURLs_To_config_OpenIDURLs(in *v1.OpenIDURLs, out *config.OpenIDURLs, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OpenIDURLs_To_config_OpenIDURLs(in, out, s)
}
func autoConvert_config_OpenIDURLs_To_v1_OpenIDURLs(in *config.OpenIDURLs, out *v1.OpenIDURLs, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Authorize = in.Authorize
	out.Token = in.Token
	out.UserInfo = in.UserInfo
	return nil
}
func Convert_config_OpenIDURLs_To_v1_OpenIDURLs(in *config.OpenIDURLs, out *v1.OpenIDURLs, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_OpenIDURLs_To_v1_OpenIDURLs(in, out, s)
}
func autoConvert_v1_PodManifestConfig_To_config_PodManifestConfig(in *v1.PodManifestConfig, out *config.PodManifestConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Path = in.Path
	out.FileCheckIntervalSeconds = in.FileCheckIntervalSeconds
	return nil
}
func Convert_v1_PodManifestConfig_To_config_PodManifestConfig(in *v1.PodManifestConfig, out *config.PodManifestConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodManifestConfig_To_config_PodManifestConfig(in, out, s)
}
func autoConvert_config_PodManifestConfig_To_v1_PodManifestConfig(in *config.PodManifestConfig, out *v1.PodManifestConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Path = in.Path
	out.FileCheckIntervalSeconds = in.FileCheckIntervalSeconds
	return nil
}
func Convert_config_PodManifestConfig_To_v1_PodManifestConfig(in *config.PodManifestConfig, out *v1.PodManifestConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_PodManifestConfig_To_v1_PodManifestConfig(in, out, s)
}
func autoConvert_v1_PolicyConfig_To_config_PolicyConfig(in *v1.PolicyConfig, out *config.PolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_UserAgentMatchingConfig_To_config_UserAgentMatchingConfig(&in.UserAgentMatchingConfig, &out.UserAgentMatchingConfig, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PolicyConfig_To_config_PolicyConfig(in *v1.PolicyConfig, out *config.PolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PolicyConfig_To_config_PolicyConfig(in, out, s)
}
func autoConvert_config_PolicyConfig_To_v1_PolicyConfig(in *config.PolicyConfig, out *v1.PolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_UserAgentMatchingConfig_To_v1_UserAgentMatchingConfig(&in.UserAgentMatchingConfig, &out.UserAgentMatchingConfig, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_PolicyConfig_To_v1_PolicyConfig(in *config.PolicyConfig, out *v1.PolicyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_PolicyConfig_To_v1_PolicyConfig(in, out, s)
}
func autoConvert_v1_ProjectConfig_To_config_ProjectConfig(in *v1.ProjectConfig, out *config.ProjectConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DefaultNodeSelector = in.DefaultNodeSelector
	out.ProjectRequestMessage = in.ProjectRequestMessage
	out.ProjectRequestTemplate = in.ProjectRequestTemplate
	out.SecurityAllocator = (*config.SecurityAllocator)(unsafe.Pointer(in.SecurityAllocator))
	return nil
}
func Convert_v1_ProjectConfig_To_config_ProjectConfig(in *v1.ProjectConfig, out *config.ProjectConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ProjectConfig_To_config_ProjectConfig(in, out, s)
}
func autoConvert_config_ProjectConfig_To_v1_ProjectConfig(in *config.ProjectConfig, out *v1.ProjectConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DefaultNodeSelector = in.DefaultNodeSelector
	out.ProjectRequestMessage = in.ProjectRequestMessage
	out.ProjectRequestTemplate = in.ProjectRequestTemplate
	out.SecurityAllocator = (*v1.SecurityAllocator)(unsafe.Pointer(in.SecurityAllocator))
	return nil
}
func Convert_config_ProjectConfig_To_v1_ProjectConfig(in *config.ProjectConfig, out *v1.ProjectConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ProjectConfig_To_v1_ProjectConfig(in, out, s)
}
func autoConvert_v1_RFC2307Config_To_config_RFC2307Config(in *v1.RFC2307Config, out *config.RFC2307Config, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_LDAPQuery_To_config_LDAPQuery(&in.AllGroupsQuery, &out.AllGroupsQuery, s); err != nil {
		return err
	}
	out.GroupUIDAttribute = in.GroupUIDAttribute
	out.GroupNameAttributes = *(*[]string)(unsafe.Pointer(&in.GroupNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	if err := Convert_v1_LDAPQuery_To_config_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserUIDAttribute = in.UserUIDAttribute
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.TolerateMemberNotFoundErrors = in.TolerateMemberNotFoundErrors
	out.TolerateMemberOutOfScopeErrors = in.TolerateMemberOutOfScopeErrors
	return nil
}
func Convert_v1_RFC2307Config_To_config_RFC2307Config(in *v1.RFC2307Config, out *config.RFC2307Config, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RFC2307Config_To_config_RFC2307Config(in, out, s)
}
func autoConvert_config_RFC2307Config_To_v1_RFC2307Config(in *config.RFC2307Config, out *v1.RFC2307Config, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_LDAPQuery_To_v1_LDAPQuery(&in.AllGroupsQuery, &out.AllGroupsQuery, s); err != nil {
		return err
	}
	out.GroupUIDAttribute = in.GroupUIDAttribute
	out.GroupNameAttributes = *(*[]string)(unsafe.Pointer(&in.GroupNameAttributes))
	out.GroupMembershipAttributes = *(*[]string)(unsafe.Pointer(&in.GroupMembershipAttributes))
	if err := Convert_config_LDAPQuery_To_v1_LDAPQuery(&in.AllUsersQuery, &out.AllUsersQuery, s); err != nil {
		return err
	}
	out.UserUIDAttribute = in.UserUIDAttribute
	out.UserNameAttributes = *(*[]string)(unsafe.Pointer(&in.UserNameAttributes))
	out.TolerateMemberNotFoundErrors = in.TolerateMemberNotFoundErrors
	out.TolerateMemberOutOfScopeErrors = in.TolerateMemberOutOfScopeErrors
	return nil
}
func Convert_config_RFC2307Config_To_v1_RFC2307Config(in *config.RFC2307Config, out *v1.RFC2307Config, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_RFC2307Config_To_v1_RFC2307Config(in, out, s)
}
func autoConvert_v1_RegistryLocation_To_config_RegistryLocation(in *v1.RegistryLocation, out *config.RegistryLocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DomainName = in.DomainName
	out.Insecure = in.Insecure
	return nil
}
func Convert_v1_RegistryLocation_To_config_RegistryLocation(in *v1.RegistryLocation, out *config.RegistryLocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RegistryLocation_To_config_RegistryLocation(in, out, s)
}
func autoConvert_config_RegistryLocation_To_v1_RegistryLocation(in *config.RegistryLocation, out *v1.RegistryLocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.DomainName = in.DomainName
	out.Insecure = in.Insecure
	return nil
}
func Convert_config_RegistryLocation_To_v1_RegistryLocation(in *config.RegistryLocation, out *v1.RegistryLocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_RegistryLocation_To_v1_RegistryLocation(in, out, s)
}
func autoConvert_v1_RemoteConnectionInfo_To_config_RemoteConnectionInfo(in *v1.RemoteConnectionInfo, out *config.RemoteConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.CA = in.CA
	return nil
}
func autoConvert_config_RemoteConnectionInfo_To_v1_RemoteConnectionInfo(in *config.RemoteConnectionInfo, out *v1.RemoteConnectionInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.CA = in.CA
	return nil
}
func autoConvert_v1_RequestHeaderAuthenticationOptions_To_config_RequestHeaderAuthenticationOptions(in *v1.RequestHeaderAuthenticationOptions, out *config.RequestHeaderAuthenticationOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientCA = in.ClientCA
	out.ClientCommonNames = *(*[]string)(unsafe.Pointer(&in.ClientCommonNames))
	out.UsernameHeaders = *(*[]string)(unsafe.Pointer(&in.UsernameHeaders))
	out.GroupHeaders = *(*[]string)(unsafe.Pointer(&in.GroupHeaders))
	out.ExtraHeaderPrefixes = *(*[]string)(unsafe.Pointer(&in.ExtraHeaderPrefixes))
	return nil
}
func Convert_v1_RequestHeaderAuthenticationOptions_To_config_RequestHeaderAuthenticationOptions(in *v1.RequestHeaderAuthenticationOptions, out *config.RequestHeaderAuthenticationOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RequestHeaderAuthenticationOptions_To_config_RequestHeaderAuthenticationOptions(in, out, s)
}
func autoConvert_config_RequestHeaderAuthenticationOptions_To_v1_RequestHeaderAuthenticationOptions(in *config.RequestHeaderAuthenticationOptions, out *v1.RequestHeaderAuthenticationOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClientCA = in.ClientCA
	out.ClientCommonNames = *(*[]string)(unsafe.Pointer(&in.ClientCommonNames))
	out.UsernameHeaders = *(*[]string)(unsafe.Pointer(&in.UsernameHeaders))
	out.GroupHeaders = *(*[]string)(unsafe.Pointer(&in.GroupHeaders))
	out.ExtraHeaderPrefixes = *(*[]string)(unsafe.Pointer(&in.ExtraHeaderPrefixes))
	return nil
}
func Convert_config_RequestHeaderAuthenticationOptions_To_v1_RequestHeaderAuthenticationOptions(in *config.RequestHeaderAuthenticationOptions, out *v1.RequestHeaderAuthenticationOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_RequestHeaderAuthenticationOptions_To_v1_RequestHeaderAuthenticationOptions(in, out, s)
}
func autoConvert_v1_RequestHeaderIdentityProvider_To_config_RequestHeaderIdentityProvider(in *v1.RequestHeaderIdentityProvider, out *config.RequestHeaderIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LoginURL = in.LoginURL
	out.ChallengeURL = in.ChallengeURL
	out.ClientCA = in.ClientCA
	out.ClientCommonNames = *(*[]string)(unsafe.Pointer(&in.ClientCommonNames))
	out.Headers = *(*[]string)(unsafe.Pointer(&in.Headers))
	out.PreferredUsernameHeaders = *(*[]string)(unsafe.Pointer(&in.PreferredUsernameHeaders))
	out.NameHeaders = *(*[]string)(unsafe.Pointer(&in.NameHeaders))
	out.EmailHeaders = *(*[]string)(unsafe.Pointer(&in.EmailHeaders))
	return nil
}
func Convert_v1_RequestHeaderIdentityProvider_To_config_RequestHeaderIdentityProvider(in *v1.RequestHeaderIdentityProvider, out *config.RequestHeaderIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RequestHeaderIdentityProvider_To_config_RequestHeaderIdentityProvider(in, out, s)
}
func autoConvert_config_RequestHeaderIdentityProvider_To_v1_RequestHeaderIdentityProvider(in *config.RequestHeaderIdentityProvider, out *v1.RequestHeaderIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LoginURL = in.LoginURL
	out.ChallengeURL = in.ChallengeURL
	out.ClientCA = in.ClientCA
	out.ClientCommonNames = *(*[]string)(unsafe.Pointer(&in.ClientCommonNames))
	out.Headers = *(*[]string)(unsafe.Pointer(&in.Headers))
	out.PreferredUsernameHeaders = *(*[]string)(unsafe.Pointer(&in.PreferredUsernameHeaders))
	out.NameHeaders = *(*[]string)(unsafe.Pointer(&in.NameHeaders))
	out.EmailHeaders = *(*[]string)(unsafe.Pointer(&in.EmailHeaders))
	return nil
}
func Convert_config_RequestHeaderIdentityProvider_To_v1_RequestHeaderIdentityProvider(in *config.RequestHeaderIdentityProvider, out *v1.RequestHeaderIdentityProvider, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_RequestHeaderIdentityProvider_To_v1_RequestHeaderIdentityProvider(in, out, s)
}
func autoConvert_v1_RoutingConfig_To_config_RoutingConfig(in *v1.RoutingConfig, out *config.RoutingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Subdomain = in.Subdomain
	return nil
}
func Convert_v1_RoutingConfig_To_config_RoutingConfig(in *v1.RoutingConfig, out *config.RoutingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoutingConfig_To_config_RoutingConfig(in, out, s)
}
func autoConvert_config_RoutingConfig_To_v1_RoutingConfig(in *config.RoutingConfig, out *v1.RoutingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Subdomain = in.Subdomain
	return nil
}
func Convert_config_RoutingConfig_To_v1_RoutingConfig(in *config.RoutingConfig, out *v1.RoutingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_RoutingConfig_To_v1_RoutingConfig(in, out, s)
}
func autoConvert_v1_SecurityAllocator_To_config_SecurityAllocator(in *v1.SecurityAllocator, out *config.SecurityAllocator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UIDAllocatorRange = in.UIDAllocatorRange
	out.MCSAllocatorRange = in.MCSAllocatorRange
	out.MCSLabelsPerProject = in.MCSLabelsPerProject
	return nil
}
func Convert_v1_SecurityAllocator_To_config_SecurityAllocator(in *v1.SecurityAllocator, out *config.SecurityAllocator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SecurityAllocator_To_config_SecurityAllocator(in, out, s)
}
func autoConvert_config_SecurityAllocator_To_v1_SecurityAllocator(in *config.SecurityAllocator, out *v1.SecurityAllocator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UIDAllocatorRange = in.UIDAllocatorRange
	out.MCSAllocatorRange = in.MCSAllocatorRange
	out.MCSLabelsPerProject = in.MCSLabelsPerProject
	return nil
}
func Convert_config_SecurityAllocator_To_v1_SecurityAllocator(in *config.SecurityAllocator, out *v1.SecurityAllocator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_SecurityAllocator_To_v1_SecurityAllocator(in, out, s)
}
func autoConvert_v1_ServiceAccountConfig_To_config_ServiceAccountConfig(in *v1.ServiceAccountConfig, out *config.ServiceAccountConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ManagedNames = *(*[]string)(unsafe.Pointer(&in.ManagedNames))
	out.LimitSecretReferences = in.LimitSecretReferences
	out.PrivateKeyFile = in.PrivateKeyFile
	out.PublicKeyFiles = *(*[]string)(unsafe.Pointer(&in.PublicKeyFiles))
	out.MasterCA = in.MasterCA
	return nil
}
func Convert_v1_ServiceAccountConfig_To_config_ServiceAccountConfig(in *v1.ServiceAccountConfig, out *config.ServiceAccountConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ServiceAccountConfig_To_config_ServiceAccountConfig(in, out, s)
}
func autoConvert_config_ServiceAccountConfig_To_v1_ServiceAccountConfig(in *config.ServiceAccountConfig, out *v1.ServiceAccountConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ManagedNames = *(*[]string)(unsafe.Pointer(&in.ManagedNames))
	out.LimitSecretReferences = in.LimitSecretReferences
	out.PrivateKeyFile = in.PrivateKeyFile
	out.PublicKeyFiles = *(*[]string)(unsafe.Pointer(&in.PublicKeyFiles))
	out.MasterCA = in.MasterCA
	return nil
}
func Convert_config_ServiceAccountConfig_To_v1_ServiceAccountConfig(in *config.ServiceAccountConfig, out *v1.ServiceAccountConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ServiceAccountConfig_To_v1_ServiceAccountConfig(in, out, s)
}
func autoConvert_v1_ServiceServingCert_To_config_ServiceServingCert(in *v1.ServiceServingCert, out *config.ServiceServingCert, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Signer = (*config.CertInfo)(unsafe.Pointer(in.Signer))
	return nil
}
func Convert_v1_ServiceServingCert_To_config_ServiceServingCert(in *v1.ServiceServingCert, out *config.ServiceServingCert, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ServiceServingCert_To_config_ServiceServingCert(in, out, s)
}
func autoConvert_config_ServiceServingCert_To_v1_ServiceServingCert(in *config.ServiceServingCert, out *v1.ServiceServingCert, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Signer = (*v1.CertInfo)(unsafe.Pointer(in.Signer))
	return nil
}
func Convert_config_ServiceServingCert_To_v1_ServiceServingCert(in *config.ServiceServingCert, out *v1.ServiceServingCert, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_ServiceServingCert_To_v1_ServiceServingCert(in, out, s)
}
func autoConvert_v1_ServingInfo_To_config_ServingInfo(in *v1.ServingInfo, out *config.ServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BindAddress = in.BindAddress
	out.BindNetwork = in.BindNetwork
	out.ClientCA = in.ClientCA
	out.NamedCertificates = *(*[]config.NamedCertificate)(unsafe.Pointer(&in.NamedCertificates))
	out.MinTLSVersion = in.MinTLSVersion
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	return nil
}
func autoConvert_config_ServingInfo_To_v1_ServingInfo(in *config.ServingInfo, out *v1.ServingInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.BindAddress = in.BindAddress
	out.BindNetwork = in.BindNetwork
	out.ClientCA = in.ClientCA
	out.NamedCertificates = *(*[]v1.NamedCertificate)(unsafe.Pointer(&in.NamedCertificates))
	out.MinTLSVersion = in.MinTLSVersion
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	return nil
}
func autoConvert_v1_SessionConfig_To_config_SessionConfig(in *v1.SessionConfig, out *config.SessionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SessionSecretsFile = in.SessionSecretsFile
	out.SessionMaxAgeSeconds = in.SessionMaxAgeSeconds
	out.SessionName = in.SessionName
	return nil
}
func Convert_v1_SessionConfig_To_config_SessionConfig(in *v1.SessionConfig, out *config.SessionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SessionConfig_To_config_SessionConfig(in, out, s)
}
func autoConvert_config_SessionConfig_To_v1_SessionConfig(in *config.SessionConfig, out *v1.SessionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SessionSecretsFile = in.SessionSecretsFile
	out.SessionMaxAgeSeconds = in.SessionMaxAgeSeconds
	out.SessionName = in.SessionName
	return nil
}
func Convert_config_SessionConfig_To_v1_SessionConfig(in *config.SessionConfig, out *v1.SessionConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_SessionConfig_To_v1_SessionConfig(in, out, s)
}
func autoConvert_v1_SessionSecret_To_config_SessionSecret(in *v1.SessionSecret, out *config.SessionSecret, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Authentication = in.Authentication
	out.Encryption = in.Encryption
	return nil
}
func Convert_v1_SessionSecret_To_config_SessionSecret(in *v1.SessionSecret, out *config.SessionSecret, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SessionSecret_To_config_SessionSecret(in, out, s)
}
func autoConvert_config_SessionSecret_To_v1_SessionSecret(in *config.SessionSecret, out *v1.SessionSecret, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Authentication = in.Authentication
	out.Encryption = in.Encryption
	return nil
}
func Convert_config_SessionSecret_To_v1_SessionSecret(in *config.SessionSecret, out *v1.SessionSecret, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_SessionSecret_To_v1_SessionSecret(in, out, s)
}
func autoConvert_v1_SessionSecrets_To_config_SessionSecrets(in *v1.SessionSecrets, out *config.SessionSecrets, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Secrets = *(*[]config.SessionSecret)(unsafe.Pointer(&in.Secrets))
	return nil
}
func Convert_v1_SessionSecrets_To_config_SessionSecrets(in *v1.SessionSecrets, out *config.SessionSecrets, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SessionSecrets_To_config_SessionSecrets(in, out, s)
}
func autoConvert_config_SessionSecrets_To_v1_SessionSecrets(in *config.SessionSecrets, out *v1.SessionSecrets, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Secrets = *(*[]v1.SessionSecret)(unsafe.Pointer(&in.Secrets))
	return nil
}
func Convert_config_SessionSecrets_To_v1_SessionSecrets(in *config.SessionSecrets, out *v1.SessionSecrets, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_SessionSecrets_To_v1_SessionSecrets(in, out, s)
}
func autoConvert_v1_SourceStrategyDefaultsConfig_To_config_SourceStrategyDefaultsConfig(in *v1.SourceStrategyDefaultsConfig, out *config.SourceStrategyDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	return nil
}
func Convert_v1_SourceStrategyDefaultsConfig_To_config_SourceStrategyDefaultsConfig(in *v1.SourceStrategyDefaultsConfig, out *config.SourceStrategyDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SourceStrategyDefaultsConfig_To_config_SourceStrategyDefaultsConfig(in, out, s)
}
func autoConvert_config_SourceStrategyDefaultsConfig_To_v1_SourceStrategyDefaultsConfig(in *config.SourceStrategyDefaultsConfig, out *v1.SourceStrategyDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	return nil
}
func Convert_config_SourceStrategyDefaultsConfig_To_v1_SourceStrategyDefaultsConfig(in *config.SourceStrategyDefaultsConfig, out *v1.SourceStrategyDefaultsConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_SourceStrategyDefaultsConfig_To_v1_SourceStrategyDefaultsConfig(in, out, s)
}
func autoConvert_v1_StringSource_To_config_StringSource(in *v1.StringSource, out *config.StringSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_StringSourceSpec_To_config_StringSourceSpec(&in.StringSourceSpec, &out.StringSourceSpec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_StringSource_To_config_StringSource(in *v1.StringSource, out *config.StringSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_StringSource_To_config_StringSource(in, out, s)
}
func autoConvert_config_StringSource_To_v1_StringSource(in *config.StringSource, out *v1.StringSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_StringSourceSpec_To_v1_StringSourceSpec(&in.StringSourceSpec, &out.StringSourceSpec, s); err != nil {
		return err
	}
	return nil
}
func Convert_config_StringSource_To_v1_StringSource(in *config.StringSource, out *v1.StringSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_StringSource_To_v1_StringSource(in, out, s)
}
func autoConvert_v1_StringSourceSpec_To_config_StringSourceSpec(in *v1.StringSourceSpec, out *config.StringSourceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Value = in.Value
	out.Env = in.Env
	out.File = in.File
	out.KeyFile = in.KeyFile
	return nil
}
func Convert_v1_StringSourceSpec_To_config_StringSourceSpec(in *v1.StringSourceSpec, out *config.StringSourceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_StringSourceSpec_To_config_StringSourceSpec(in, out, s)
}
func autoConvert_config_StringSourceSpec_To_v1_StringSourceSpec(in *config.StringSourceSpec, out *v1.StringSourceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Value = in.Value
	out.Env = in.Env
	out.File = in.File
	out.KeyFile = in.KeyFile
	return nil
}
func Convert_config_StringSourceSpec_To_v1_StringSourceSpec(in *config.StringSourceSpec, out *v1.StringSourceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_StringSourceSpec_To_v1_StringSourceSpec(in, out, s)
}
func autoConvert_v1_TokenConfig_To_config_TokenConfig(in *v1.TokenConfig, out *config.TokenConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AuthorizeTokenMaxAgeSeconds = in.AuthorizeTokenMaxAgeSeconds
	out.AccessTokenMaxAgeSeconds = in.AccessTokenMaxAgeSeconds
	out.AccessTokenInactivityTimeoutSeconds = (*int32)(unsafe.Pointer(in.AccessTokenInactivityTimeoutSeconds))
	return nil
}
func Convert_v1_TokenConfig_To_config_TokenConfig(in *v1.TokenConfig, out *config.TokenConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TokenConfig_To_config_TokenConfig(in, out, s)
}
func autoConvert_config_TokenConfig_To_v1_TokenConfig(in *config.TokenConfig, out *v1.TokenConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AuthorizeTokenMaxAgeSeconds = in.AuthorizeTokenMaxAgeSeconds
	out.AccessTokenMaxAgeSeconds = in.AccessTokenMaxAgeSeconds
	out.AccessTokenInactivityTimeoutSeconds = (*int32)(unsafe.Pointer(in.AccessTokenInactivityTimeoutSeconds))
	return nil
}
func Convert_config_TokenConfig_To_v1_TokenConfig(in *config.TokenConfig, out *v1.TokenConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_TokenConfig_To_v1_TokenConfig(in, out, s)
}
func autoConvert_v1_UserAgentDenyRule_To_config_UserAgentDenyRule(in *v1.UserAgentDenyRule, out *config.UserAgentDenyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_UserAgentMatchRule_To_config_UserAgentMatchRule(&in.UserAgentMatchRule, &out.UserAgentMatchRule, s); err != nil {
		return err
	}
	out.RejectionMessage = in.RejectionMessage
	return nil
}
func Convert_v1_UserAgentDenyRule_To_config_UserAgentDenyRule(in *v1.UserAgentDenyRule, out *config.UserAgentDenyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_UserAgentDenyRule_To_config_UserAgentDenyRule(in, out, s)
}
func autoConvert_config_UserAgentDenyRule_To_v1_UserAgentDenyRule(in *config.UserAgentDenyRule, out *v1.UserAgentDenyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_config_UserAgentMatchRule_To_v1_UserAgentMatchRule(&in.UserAgentMatchRule, &out.UserAgentMatchRule, s); err != nil {
		return err
	}
	out.RejectionMessage = in.RejectionMessage
	return nil
}
func Convert_config_UserAgentDenyRule_To_v1_UserAgentDenyRule(in *config.UserAgentDenyRule, out *v1.UserAgentDenyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_UserAgentDenyRule_To_v1_UserAgentDenyRule(in, out, s)
}
func autoConvert_v1_UserAgentMatchRule_To_config_UserAgentMatchRule(in *v1.UserAgentMatchRule, out *config.UserAgentMatchRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Regex = in.Regex
	out.HTTPVerbs = *(*[]string)(unsafe.Pointer(&in.HTTPVerbs))
	return nil
}
func Convert_v1_UserAgentMatchRule_To_config_UserAgentMatchRule(in *v1.UserAgentMatchRule, out *config.UserAgentMatchRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_UserAgentMatchRule_To_config_UserAgentMatchRule(in, out, s)
}
func autoConvert_config_UserAgentMatchRule_To_v1_UserAgentMatchRule(in *config.UserAgentMatchRule, out *v1.UserAgentMatchRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Regex = in.Regex
	out.HTTPVerbs = *(*[]string)(unsafe.Pointer(&in.HTTPVerbs))
	return nil
}
func Convert_config_UserAgentMatchRule_To_v1_UserAgentMatchRule(in *config.UserAgentMatchRule, out *v1.UserAgentMatchRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_UserAgentMatchRule_To_v1_UserAgentMatchRule(in, out, s)
}
func autoConvert_v1_UserAgentMatchingConfig_To_config_UserAgentMatchingConfig(in *v1.UserAgentMatchingConfig, out *config.UserAgentMatchingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RequiredClients = *(*[]config.UserAgentMatchRule)(unsafe.Pointer(&in.RequiredClients))
	out.DeniedClients = *(*[]config.UserAgentDenyRule)(unsafe.Pointer(&in.DeniedClients))
	out.DefaultRejectionMessage = in.DefaultRejectionMessage
	return nil
}
func Convert_v1_UserAgentMatchingConfig_To_config_UserAgentMatchingConfig(in *v1.UserAgentMatchingConfig, out *config.UserAgentMatchingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_UserAgentMatchingConfig_To_config_UserAgentMatchingConfig(in, out, s)
}
func autoConvert_config_UserAgentMatchingConfig_To_v1_UserAgentMatchingConfig(in *config.UserAgentMatchingConfig, out *v1.UserAgentMatchingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RequiredClients = *(*[]v1.UserAgentMatchRule)(unsafe.Pointer(&in.RequiredClients))
	out.DeniedClients = *(*[]v1.UserAgentDenyRule)(unsafe.Pointer(&in.DeniedClients))
	out.DefaultRejectionMessage = in.DefaultRejectionMessage
	return nil
}
func Convert_config_UserAgentMatchingConfig_To_v1_UserAgentMatchingConfig(in *config.UserAgentMatchingConfig, out *v1.UserAgentMatchingConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_UserAgentMatchingConfig_To_v1_UserAgentMatchingConfig(in, out, s)
}
func autoConvert_v1_WebhookTokenAuthenticator_To_config_WebhookTokenAuthenticator(in *v1.WebhookTokenAuthenticator, out *config.WebhookTokenAuthenticator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ConfigFile = in.ConfigFile
	out.CacheTTL = in.CacheTTL
	return nil
}
func Convert_v1_WebhookTokenAuthenticator_To_config_WebhookTokenAuthenticator(in *v1.WebhookTokenAuthenticator, out *config.WebhookTokenAuthenticator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_WebhookTokenAuthenticator_To_config_WebhookTokenAuthenticator(in, out, s)
}
func autoConvert_config_WebhookTokenAuthenticator_To_v1_WebhookTokenAuthenticator(in *config.WebhookTokenAuthenticator, out *v1.WebhookTokenAuthenticator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ConfigFile = in.ConfigFile
	out.CacheTTL = in.CacheTTL
	return nil
}
func Convert_config_WebhookTokenAuthenticator_To_v1_WebhookTokenAuthenticator(in *config.WebhookTokenAuthenticator, out *v1.WebhookTokenAuthenticator, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_config_WebhookTokenAuthenticator_To_v1_WebhookTokenAuthenticator(in, out, s)
}
