package v1

import (
	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	internal "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RegisterDefaults(scheme)
	return nil
}
func SetDefaults_MasterConfig(obj *legacyconfigv1.MasterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.APILevels) == 0 {
		obj.APILevels = internal.DefaultOpenShiftAPILevels
	}
	if len(obj.ControllerConfig.Controllers) == 0 {
		switch {
		case len(obj.Controllers) == 0 || obj.Controllers == legacyconfigv1.ControllersAll:
			obj.ControllerConfig.Controllers = []string{"*"}
		case obj.Controllers == legacyconfigv1.ControllersDisabled:
			obj.ControllerConfig.Controllers = []string{}
		}
	}
	if election := obj.ControllerConfig.Election; election != nil {
		if len(election.LockNamespace) == 0 {
			election.LockNamespace = "kube-system"
		}
		if len(election.LockResource.Group) == 0 && len(election.LockResource.Resource) == 0 {
			election.LockResource.Resource = "configmaps"
		}
	}
	if obj.ServingInfo.RequestTimeoutSeconds == 0 {
		obj.ServingInfo.RequestTimeoutSeconds = 60 * 60
	}
	if obj.ServingInfo.MaxRequestsInFlight == 0 {
		obj.ServingInfo.MaxRequestsInFlight = 1200
	}
	if len(obj.RoutingConfig.Subdomain) == 0 {
		obj.RoutingConfig.Subdomain = "router.default.svc.cluster.local"
	}
	if len(obj.JenkinsPipelineConfig.TemplateNamespace) == 0 {
		obj.JenkinsPipelineConfig.TemplateNamespace = "openshift"
	}
	if len(obj.JenkinsPipelineConfig.TemplateName) == 0 {
		obj.JenkinsPipelineConfig.TemplateName = "jenkins-ephemeral"
	}
	if len(obj.JenkinsPipelineConfig.ServiceName) == 0 {
		obj.JenkinsPipelineConfig.ServiceName = "jenkins"
	}
	if obj.JenkinsPipelineConfig.AutoProvisionEnabled == nil {
		v := true
		obj.JenkinsPipelineConfig.AutoProvisionEnabled = &v
	}
	if obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides == nil {
		obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides = &legacyconfigv1.ClientConnectionOverrides{}
	}
	if obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides.QPS <= 0 {
		obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides.QPS = 150.0
	}
	if obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides.Burst <= 0 {
		obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides.Burst = 300
	}
	SetDefaults_ClientConnectionOverrides(obj.MasterClients.OpenShiftLoopbackClientConnectionOverrides)
	if len(obj.NetworkConfig.ServiceNetworkCIDR) == 0 {
		if len(obj.KubernetesMasterConfig.ServicesSubnet) > 0 {
			obj.NetworkConfig.ServiceNetworkCIDR = obj.KubernetesMasterConfig.ServicesSubnet
		} else {
			obj.NetworkConfig.ServiceNetworkCIDR = "10.0.0.0/24"
		}
	}
	noCloudProvider := (len(obj.KubernetesMasterConfig.ControllerArguments["cloud-provider"]) == 0 || obj.KubernetesMasterConfig.ControllerArguments["cloud-provider"][0] == "")
	if noCloudProvider && len(obj.NetworkConfig.IngressIPNetworkCIDR) == 0 {
		cidr := internal.DefaultIngressIPNetworkCIDR
		cidrOverlap := false
		if internal.CIDRsOverlap(cidr, obj.NetworkConfig.ServiceNetworkCIDR) {
			cidrOverlap = true
		} else {
			for _, entry := range obj.NetworkConfig.ClusterNetworks {
				if internal.CIDRsOverlap(cidr, entry.CIDR) {
					cidrOverlap = true
					break
				}
			}
		}
		if !cidrOverlap {
			obj.NetworkConfig.IngressIPNetworkCIDR = cidr
		}
	}
	if obj.OAuthConfig != nil && obj.OAuthConfig.MasterCA == nil {
		s := obj.ServingInfo.ClientCA
		obj.OAuthConfig.MasterCA = &s
	}
	for pluginName := range obj.AdmissionConfig.PluginConfig {
		if obj.AdmissionConfig.PluginConfig[pluginName] == nil {
			obj.AdmissionConfig.PluginConfig[pluginName] = &legacyconfigv1.AdmissionPluginConfig{}
		}
	}
	for i := range obj.AuthConfig.WebhookTokenAuthenticators {
		if len(obj.AuthConfig.WebhookTokenAuthenticators[i].CacheTTL) == 0 {
			obj.AuthConfig.WebhookTokenAuthenticators[i].CacheTTL = "2m"
		}
	}
}
func SetDefaults_KubernetesMasterConfig(obj *legacyconfigv1.KubernetesMasterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.MasterEndpointReconcileTTL == 0 {
		obj.MasterEndpointReconcileTTL = 15
	}
	if len(obj.APILevels) == 0 {
		obj.APILevels = internal.DefaultKubernetesAPILevels
	}
	if len(obj.ServicesNodePortRange) == 0 {
		obj.ServicesNodePortRange = "30000-32767"
	}
	if len(obj.PodEvictionTimeout) == 0 {
		obj.PodEvictionTimeout = "5m"
	}
}
func SetDefaults_NodeConfig(obj *legacyconfigv1.NodeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.MasterClientConnectionOverrides == nil {
		obj.MasterClientConnectionOverrides = &legacyconfigv1.ClientConnectionOverrides{QPS: 10.0, Burst: 20}
	}
	SetDefaults_ClientConnectionOverrides(obj.MasterClientConnectionOverrides)
	if len(obj.NetworkConfig.NetworkPluginName) == 0 {
		obj.NetworkConfig.NetworkPluginName = obj.DeprecatedNetworkPluginName
	}
	if obj.NetworkConfig.MTU == 0 {
		obj.NetworkConfig.MTU = 1450
	}
	if len(obj.IPTablesSyncPeriod) == 0 {
		obj.IPTablesSyncPeriod = "30s"
	}
	if len(obj.AuthConfig.AuthenticationCacheTTL) == 0 {
		obj.AuthConfig.AuthenticationCacheTTL = "5m"
	}
	if obj.AuthConfig.AuthenticationCacheSize == 0 {
		obj.AuthConfig.AuthenticationCacheSize = 1000
	}
	if len(obj.AuthConfig.AuthorizationCacheTTL) == 0 {
		obj.AuthConfig.AuthorizationCacheTTL = "5m"
	}
	if obj.AuthConfig.AuthorizationCacheSize == 0 {
		obj.AuthConfig.AuthorizationCacheSize = 1000
	}
	if obj.EnableUnidling == nil {
		v := true
		obj.EnableUnidling = &v
	}
}
func SetDefaults_EtcdStorageConfig(obj *legacyconfigv1.EtcdStorageConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.KubernetesStorageVersion) == 0 {
		obj.KubernetesStorageVersion = "v1"
	}
	if len(obj.KubernetesStoragePrefix) == 0 {
		obj.KubernetesStoragePrefix = "kubernetes.io"
	}
	if len(obj.OpenShiftStorageVersion) == 0 {
		obj.OpenShiftStorageVersion = internal.DefaultOpenShiftStorageVersionLevel
	}
	if len(obj.OpenShiftStoragePrefix) == 0 {
		obj.OpenShiftStoragePrefix = "openshift.io"
	}
}
func SetDefaults_DockerConfig(obj *legacyconfigv1.DockerConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.ExecHandlerName) == 0 {
		obj.ExecHandlerName = legacyconfigv1.DockerExecHandlerNative
	}
	if len(obj.DockerShimSocket) == 0 {
		obj.DockerShimSocket = "unix:///var/run/dockershim.sock"
	}
	if len(obj.DockershimRootDirectory) == 0 {
		obj.DockershimRootDirectory = "/var/lib/dockershim"
	}
}
func SetDefaults_ServingInfo(obj *legacyconfigv1.ServingInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.BindNetwork) == 0 {
		obj.BindNetwork = "tcp4"
	}
}
func SetDefaults_ImagePolicyConfig(obj *legacyconfigv1.ImagePolicyConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.MaxImagesBulkImportedPerRepository == 0 {
		obj.MaxImagesBulkImportedPerRepository = 50
	}
	if obj.MaxScheduledImageImportsPerMinute == 0 {
		obj.MaxScheduledImageImportsPerMinute = 60
	}
	if obj.ScheduledImageImportMinimumIntervalSeconds == 0 {
		obj.ScheduledImageImportMinimumIntervalSeconds = 15 * 60
	}
}
func SetDefaults_DNSConfig(obj *legacyconfigv1.DNSConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.BindNetwork) == 0 {
		obj.BindNetwork = "tcp4"
	}
}
func SetDefaults_SecurityAllocator(obj *legacyconfigv1.SecurityAllocator) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.UIDAllocatorRange) == 0 {
		obj.UIDAllocatorRange = "1000000000-1999999999/10000"
	}
	if len(obj.MCSAllocatorRange) == 0 {
		obj.MCSAllocatorRange = "s0:/2"
	}
	if obj.MCSLabelsPerProject == 0 {
		obj.MCSLabelsPerProject = 5
	}
}
func SetDefaults_IdentityProvider(obj *legacyconfigv1.IdentityProvider) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.MappingMethod) == 0 {
		obj.MappingMethod = "claim"
	}
}
func SetDefaults_GrantConfig(obj *legacyconfigv1.GrantConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(obj.ServiceAccountMethod) == 0 {
		obj.ServiceAccountMethod = legacyconfigv1.GrantHandlerPrompt
	}
}
func SetDefaults_ClientConnectionOverrides(overrides *legacyconfigv1.ClientConnectionOverrides) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(overrides.AcceptContentTypes) == 0 {
		overrides.AcceptContentTypes = "application/json"
	}
	if len(overrides.ContentType) == 0 {
		overrides.ContentType = "application/json"
	}
}
