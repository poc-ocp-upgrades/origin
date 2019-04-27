package v1

import (
	v1 "github.com/openshift/api/legacyconfig/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.BuildDefaultsConfig{}, func(obj interface{}) {
		SetObjectDefaults_BuildDefaultsConfig(obj.(*v1.BuildDefaultsConfig))
	})
	scheme.AddTypeDefaultingFunc(&v1.MasterConfig{}, func(obj interface{}) {
		SetObjectDefaults_MasterConfig(obj.(*v1.MasterConfig))
	})
	scheme.AddTypeDefaultingFunc(&v1.NodeConfig{}, func(obj interface{}) {
		SetObjectDefaults_NodeConfig(obj.(*v1.NodeConfig))
	})
	return nil
}
func SetObjectDefaults_BuildDefaultsConfig(in *v1.BuildDefaultsConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Env {
		a := &in.Env[i]
		if a.ValueFrom != nil {
			if a.ValueFrom.FieldRef != nil {
				corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
			}
		}
	}
	corev1.SetDefaults_ResourceList(&in.Resources.Limits)
	corev1.SetDefaults_ResourceList(&in.Resources.Requests)
}
func SetObjectDefaults_MasterConfig(in *v1.MasterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_MasterConfig(in)
	SetDefaults_ServingInfo(&in.ServingInfo.ServingInfo)
	SetDefaults_EtcdStorageConfig(&in.EtcdStorageConfig)
	SetDefaults_KubernetesMasterConfig(&in.KubernetesMasterConfig)
	if in.EtcdConfig != nil {
		SetDefaults_ServingInfo(&in.EtcdConfig.ServingInfo)
		SetDefaults_ServingInfo(&in.EtcdConfig.PeerServingInfo)
	}
	if in.OAuthConfig != nil {
		for i := range in.OAuthConfig.IdentityProviders {
			a := &in.OAuthConfig.IdentityProviders[i]
			SetDefaults_IdentityProvider(a)
		}
		SetDefaults_GrantConfig(&in.OAuthConfig.GrantConfig)
	}
	if in.DNSConfig != nil {
		SetDefaults_DNSConfig(in.DNSConfig)
	}
	if in.MasterClients.OpenShiftLoopbackClientConnectionOverrides != nil {
		SetDefaults_ClientConnectionOverrides(in.MasterClients.OpenShiftLoopbackClientConnectionOverrides)
	}
	SetDefaults_ImagePolicyConfig(&in.ImagePolicyConfig)
	if in.ProjectConfig.SecurityAllocator != nil {
		SetDefaults_SecurityAllocator(in.ProjectConfig.SecurityAllocator)
	}
}
func SetObjectDefaults_NodeConfig(in *v1.NodeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_NodeConfig(in)
	SetDefaults_ServingInfo(&in.ServingInfo)
	if in.MasterClientConnectionOverrides != nil {
		SetDefaults_ClientConnectionOverrides(in.MasterClientConnectionOverrides)
	}
	SetDefaults_DockerConfig(&in.DockerConfig)
}
