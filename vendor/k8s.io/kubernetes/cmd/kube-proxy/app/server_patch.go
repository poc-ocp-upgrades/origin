package app

import (
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
)

func (o *Options) GetConfig() *kubeproxyconfig.KubeProxyConfiguration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.config
}
