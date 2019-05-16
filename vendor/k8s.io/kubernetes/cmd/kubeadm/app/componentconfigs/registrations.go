package componentconfigs

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeproxyconfigv1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	kubeletconfigv1beta1scheme "k8s.io/kubernetes/pkg/kubelet/apis/config/v1beta1"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	kubeproxyconfigv1alpha1scheme "k8s.io/kubernetes/pkg/proxy/apis/config/v1alpha1"
)

type AddToSchemeFunc func(*runtime.Scheme) error
type Registration struct {
	MarshalGroupVersion   schema.GroupVersion
	AddToSchemeFuncs      []AddToSchemeFunc
	DefaulterFunc         func(*kubeadmapi.ClusterConfiguration)
	ValidateFunc          func(*kubeadmapi.ClusterConfiguration, *field.Path) field.ErrorList
	EmptyValue            runtime.Object
	GetFromInternalConfig func(*kubeadmapi.ClusterConfiguration) (runtime.Object, bool)
	SetToInternalConfig   func(runtime.Object, *kubeadmapi.ClusterConfiguration) bool
	GetFromConfigMap      func(clientset.Interface, *version.Version) (runtime.Object, error)
}

func (r Registration) Marshal(obj runtime.Object) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return kubeadmutil.MarshalToYamlForCodecs(obj, r.MarshalGroupVersion, Codecs)
}
func (r Registration) Unmarshal(fileContent []byte) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj := r.EmptyValue.DeepCopyObject()
	if err := unmarshalObject(obj, fileContent); err != nil {
		return nil, err
	}
	return obj, nil
}
func unmarshalObject(obj runtime.Object, fileContent []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := runtime.DecodeInto(Codecs.UniversalDecoder(), fileContent, obj); err != nil {
		return err
	}
	return nil
}

const (
	KubeletConfigurationKind   RegistrationKind = "KubeletConfiguration"
	KubeProxyConfigurationKind RegistrationKind = "KubeProxyConfiguration"
)

type RegistrationKind string
type Registrations map[RegistrationKind]Registration

var Known Registrations = map[RegistrationKind]Registration{KubeProxyConfigurationKind: {MarshalGroupVersion: kubeproxyconfigv1alpha1.SchemeGroupVersion, AddToSchemeFuncs: []AddToSchemeFunc{kubeproxyconfig.AddToScheme, kubeproxyconfigv1alpha1scheme.AddToScheme}, DefaulterFunc: DefaultKubeProxyConfiguration, ValidateFunc: ValidateKubeProxyConfiguration, EmptyValue: &kubeproxyconfig.KubeProxyConfiguration{}, GetFromInternalConfig: func(cfg *kubeadmapi.ClusterConfiguration) (runtime.Object, bool) {
	return cfg.ComponentConfigs.KubeProxy, cfg.ComponentConfigs.KubeProxy != nil
}, SetToInternalConfig: func(obj runtime.Object, cfg *kubeadmapi.ClusterConfiguration) bool {
	kubeproxyConfig, ok := obj.(*kubeproxyconfig.KubeProxyConfiguration)
	if ok {
		cfg.ComponentConfigs.KubeProxy = kubeproxyConfig
	}
	return ok
}, GetFromConfigMap: GetFromKubeProxyConfigMap}, KubeletConfigurationKind: {MarshalGroupVersion: kubeletconfigv1beta1.SchemeGroupVersion, AddToSchemeFuncs: []AddToSchemeFunc{kubeletconfig.AddToScheme, kubeletconfigv1beta1scheme.AddToScheme}, DefaulterFunc: DefaultKubeletConfiguration, ValidateFunc: ValidateKubeletConfiguration, EmptyValue: &kubeletconfig.KubeletConfiguration{}, GetFromInternalConfig: func(cfg *kubeadmapi.ClusterConfiguration) (runtime.Object, bool) {
	return cfg.ComponentConfigs.Kubelet, cfg.ComponentConfigs.Kubelet != nil
}, SetToInternalConfig: func(obj runtime.Object, cfg *kubeadmapi.ClusterConfiguration) bool {
	kubeletConfig, ok := obj.(*kubeletconfig.KubeletConfiguration)
	if ok {
		cfg.ComponentConfigs.Kubelet = kubeletConfig
	}
	return ok
}, GetFromConfigMap: GetFromKubeletConfigMap}}

func (rs *Registrations) AddToScheme(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, registration := range *rs {
		for _, addToSchemeFunc := range registration.AddToSchemeFuncs {
			if err := addToSchemeFunc(scheme); err != nil {
				return err
			}
		}
	}
	return nil
}
func (rs *Registrations) Default(internalcfg *kubeadmapi.ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, registration := range *rs {
		registration.DefaulterFunc(internalcfg)
	}
}
func (rs *Registrations) Validate(internalcfg *kubeadmapi.ClusterConfiguration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for kind, registration := range *rs {
		allErrs = append(allErrs, registration.ValidateFunc(internalcfg, field.NewPath(string(kind)))...)
	}
	return allErrs
}
