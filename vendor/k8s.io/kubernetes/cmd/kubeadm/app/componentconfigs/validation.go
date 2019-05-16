package componentconfigs

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeletvalidation "k8s.io/kubernetes/pkg/kubelet/apis/config/validation"
	proxyvalidation "k8s.io/kubernetes/pkg/proxy/apis/config/validation"
)

func ValidateKubeProxyConfiguration(internalcfg *kubeadmapi.ClusterConfiguration, _ *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if internalcfg.ComponentConfigs.KubeProxy == nil {
		return allErrs
	}
	return proxyvalidation.Validate(internalcfg.ComponentConfigs.KubeProxy)
}
func ValidateKubeletConfiguration(internalcfg *kubeadmapi.ClusterConfiguration, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if internalcfg.ComponentConfigs.Kubelet == nil {
		return allErrs
	}
	if err := kubeletvalidation.ValidateKubeletConfiguration(internalcfg.ComponentConfigs.Kubelet); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "", err.Error()))
	}
	return allErrs
}
