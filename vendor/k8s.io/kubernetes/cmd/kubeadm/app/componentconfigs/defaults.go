package componentconfigs

import (
	kubeproxyconfigv1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	utilpointer "k8s.io/utils/pointer"
)

const (
	KubeproxyKubeConfigFileName = "/var/lib/kube-proxy/kubeconfig.conf"
)

func DefaultKubeProxyConfiguration(internalcfg *kubeadmapi.ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	externalproxycfg := &kubeproxyconfigv1alpha1.KubeProxyConfiguration{}
	if internalcfg.ComponentConfigs.KubeProxy != nil {
		Scheme.Convert(internalcfg.ComponentConfigs.KubeProxy, externalproxycfg, nil)
	}
	if externalproxycfg.ClusterCIDR == "" && internalcfg.Networking.PodSubnet != "" {
		externalproxycfg.ClusterCIDR = internalcfg.Networking.PodSubnet
	}
	if externalproxycfg.ClientConnection.Kubeconfig == "" {
		externalproxycfg.ClientConnection.Kubeconfig = KubeproxyKubeConfigFileName
	}
	Scheme.Default(externalproxycfg)
	if internalcfg.ComponentConfigs.KubeProxy == nil {
		internalcfg.ComponentConfigs.KubeProxy = &kubeproxyconfig.KubeProxyConfiguration{}
	}
	Scheme.Convert(externalproxycfg, internalcfg.ComponentConfigs.KubeProxy, nil)
}
func DefaultKubeletConfiguration(internalcfg *kubeadmapi.ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	externalkubeletcfg := &kubeletconfigv1beta1.KubeletConfiguration{}
	if internalcfg.ComponentConfigs.Kubelet != nil {
		Scheme.Convert(internalcfg.ComponentConfigs.Kubelet, externalkubeletcfg, nil)
	}
	if externalkubeletcfg.StaticPodPath == "" {
		externalkubeletcfg.StaticPodPath = kubeadmapiv1beta1.DefaultManifestsDir
	}
	if externalkubeletcfg.ClusterDNS == nil {
		dnsIP, err := constants.GetDNSIP(internalcfg.Networking.ServiceSubnet)
		if err != nil {
			externalkubeletcfg.ClusterDNS = []string{kubeadmapiv1beta1.DefaultClusterDNSIP}
		} else {
			externalkubeletcfg.ClusterDNS = []string{dnsIP.String()}
		}
	}
	if externalkubeletcfg.ClusterDomain == "" {
		externalkubeletcfg.ClusterDomain = internalcfg.Networking.DNSDomain
	}
	externalkubeletcfg.Authentication.X509.ClientCAFile = kubeadmapiv1beta1.DefaultCACertPath
	externalkubeletcfg.Authentication.Anonymous.Enabled = utilpointer.BoolPtr(false)
	externalkubeletcfg.Authorization.Mode = kubeletconfigv1beta1.KubeletAuthorizationModeWebhook
	externalkubeletcfg.Authentication.Webhook.Enabled = utilpointer.BoolPtr(true)
	externalkubeletcfg.ReadOnlyPort = 0
	externalkubeletcfg.RotateCertificates = true
	externalkubeletcfg.HealthzBindAddress = "127.0.0.1"
	externalkubeletcfg.HealthzPort = utilpointer.Int32Ptr(constants.KubeletHealthzPort)
	Scheme.Default(externalkubeletcfg)
	if internalcfg.ComponentConfigs.Kubelet == nil {
		internalcfg.ComponentConfigs.Kubelet = &kubeletconfig.KubeletConfiguration{}
	}
	Scheme.Convert(externalkubeletcfg, internalcfg.ComponentConfigs.Kubelet, nil)
}
