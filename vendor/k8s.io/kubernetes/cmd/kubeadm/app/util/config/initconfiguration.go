package config

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/config/strict"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"net"
	"reflect"
	"sort"
	"strconv"
)

func SetInitDynamicDefaults(cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := SetBootstrapTokensDynamicDefaults(&cfg.BootstrapTokens); err != nil {
		return err
	}
	if err := SetNodeRegistrationDynamicDefaults(&cfg.NodeRegistration, true); err != nil {
		return err
	}
	if err := SetAPIEndpointDynamicDefaults(&cfg.LocalAPIEndpoint); err != nil {
		return err
	}
	if err := SetClusterDynamicDefaults(&cfg.ClusterConfiguration, cfg.LocalAPIEndpoint.AdvertiseAddress, cfg.LocalAPIEndpoint.BindPort); err != nil {
		return err
	}
	return nil
}
func SetBootstrapTokensDynamicDefaults(cfg *[]kubeadmapi.BootstrapToken) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, bt := range *cfg {
		if bt.Token != nil && len(bt.Token.String()) > 0 {
			continue
		}
		tokenStr, err := bootstraputil.GenerateBootstrapToken()
		if err != nil {
			return errors.Wrap(err, "couldn't generate random token")
		}
		token, err := kubeadmapi.NewBootstrapTokenString(tokenStr)
		if err != nil {
			return err
		}
		(*cfg)[i].Token = token
	}
	return nil
}
func SetNodeRegistrationDynamicDefaults(cfg *kubeadmapi.NodeRegistrationOptions, masterTaint bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	cfg.Name, err = nodeutil.GetHostname(cfg.Name)
	if err != nil {
		return err
	}
	if masterTaint && cfg.Taints == nil {
		cfg.Taints = []v1.Taint{kubeadmconstants.MasterTaint}
	}
	return nil
}
func SetAPIEndpointDynamicDefaults(cfg *kubeadmapi.APIEndpoint) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addressIP := net.ParseIP(cfg.AdvertiseAddress)
	if addressIP == nil && cfg.AdvertiseAddress != "" {
		return errors.Errorf("couldn't use \"%s\" as \"apiserver-advertise-address\", must be ipv4 or ipv6 address", cfg.AdvertiseAddress)
	}
	ip, err := ChooseAPIServerBindAddress(addressIP)
	if err != nil {
		return err
	}
	cfg.AdvertiseAddress = ip.String()
	return nil
}
func SetClusterDynamicDefaults(cfg *kubeadmapi.ClusterConfiguration, advertiseAddress string, bindPort int32) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	componentconfigs.Known.Default(cfg)
	ip := net.ParseIP(advertiseAddress)
	if ip.To4() != nil {
		cfg.ComponentConfigs.KubeProxy.BindAddress = kubeadmapiv1beta1.DefaultProxyBindAddressv4
	} else {
		cfg.ComponentConfigs.KubeProxy.BindAddress = kubeadmapiv1beta1.DefaultProxyBindAddressv6
	}
	if err := NormalizeKubernetesVersion(cfg); err != nil {
		return err
	}
	if cfg.ControlPlaneEndpoint != "" {
		host, port, err := kubeadmutil.ParseHostPort(cfg.ControlPlaneEndpoint)
		if err != nil {
			return err
		}
		if port == "" {
			cfg.ControlPlaneEndpoint = net.JoinHostPort(host, strconv.FormatInt(int64(bindPort), 10))
		}
	}
	LowercaseSANs(cfg.APIServer.CertSANs)
	return nil
}
func ConfigFileAndDefaultsToInternalConfig(cfgPath string, defaultversionedcfg *kubeadmapiv1beta1.InitConfiguration) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	internalcfg := &kubeadmapi.InitConfiguration{}
	if cfgPath != "" {
		klog.V(1).Infoln("loading configuration from the given file")
		b, err := ioutil.ReadFile(cfgPath)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read config from %q ", cfgPath)
		}
		internalcfg, err = BytesToInternalConfig(b)
		if err != nil {
			return nil, err
		}
	} else {
		kubeadmscheme.Scheme.Default(defaultversionedcfg)
		kubeadmscheme.Scheme.Convert(defaultversionedcfg, internalcfg, nil)
	}
	if err := SetInitDynamicDefaults(internalcfg); err != nil {
		return nil, err
	}
	if err := validation.ValidateInitConfiguration(internalcfg).ToAggregate(); err != nil {
		return nil, err
	}
	return internalcfg, nil
}
func BytesToInternalConfig(b []byte) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var initcfg *kubeadmapi.InitConfiguration
	var clustercfg *kubeadmapi.ClusterConfiguration
	decodedComponentConfigObjects := map[componentconfigs.RegistrationKind]runtime.Object{}
	if err := DetectUnsupportedVersion(b); err != nil {
		return nil, err
	}
	gvkmap, err := kubeadmutil.SplitYAMLDocuments(b)
	if err != nil {
		return nil, err
	}
	for gvk, fileContent := range gvkmap {
		strict.VerifyUnmarshalStrict(fileContent, gvk)
		regKind := componentconfigs.RegistrationKind(gvk.Kind)
		registration, found := componentconfigs.Known[regKind]
		if found {
			obj, err := registration.Unmarshal(fileContent)
			if err != nil {
				return nil, err
			}
			decodedComponentConfigObjects[regKind] = obj
			continue
		}
		if kubeadmutil.GroupVersionKindsHasInitConfiguration(gvk) {
			initcfg = &kubeadmapi.InitConfiguration{}
			if err := runtime.DecodeInto(kubeadmscheme.Codecs.UniversalDecoder(), fileContent, initcfg); err != nil {
				return nil, err
			}
			continue
		}
		if kubeadmutil.GroupVersionKindsHasClusterConfiguration(gvk) {
			clustercfg = &kubeadmapi.ClusterConfiguration{}
			if err := runtime.DecodeInto(kubeadmscheme.Codecs.UniversalDecoder(), fileContent, clustercfg); err != nil {
				return nil, err
			}
			continue
		}
		fmt.Printf("[config] WARNING: Ignored YAML document with GroupVersionKind %v\n", gvk)
	}
	if initcfg == nil && clustercfg == nil {
		return nil, errors.New("no InitConfiguration or ClusterConfiguration kind was found in the YAML file")
	}
	if initcfg == nil {
		extinitcfg := &kubeadmapiv1beta1.InitConfiguration{}
		kubeadmscheme.Scheme.Default(extinitcfg)
		initcfg = &kubeadmapi.InitConfiguration{}
		kubeadmscheme.Scheme.Convert(extinitcfg, initcfg, nil)
	}
	if clustercfg != nil {
		initcfg.ClusterConfiguration = *clustercfg
	}
	for kind, obj := range decodedComponentConfigObjects {
		registration, found := componentconfigs.Known[kind]
		if found {
			if ok := registration.SetToInternalConfig(obj, &initcfg.ClusterConfiguration); !ok {
				return nil, errors.Errorf("couldn't save componentconfig value for kind %q", string(kind))
			}
		} else {
			fmt.Printf("[config] WARNING: Decoded a kind that couldn't be saved to the internal configuration: %q\n", string(kind))
		}
	}
	return initcfg, nil
}
func defaultedInternalConfig() *kubeadmapi.ClusterConfiguration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	externalcfg := &kubeadmapiv1beta1.ClusterConfiguration{}
	internalcfg := &kubeadmapi.ClusterConfiguration{}
	kubeadmscheme.Scheme.Default(externalcfg)
	kubeadmscheme.Scheme.Convert(externalcfg, internalcfg, nil)
	componentconfigs.Known.Default(internalcfg)
	return internalcfg
}
func MarshalInitConfigurationToBytes(cfg *kubeadmapi.InitConfiguration, gv schema.GroupVersion) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	initbytes, err := kubeadmutil.MarshalToYamlForCodecs(cfg, gv, kubeadmscheme.Codecs)
	if err != nil {
		return []byte{}, err
	}
	allFiles := [][]byte{initbytes}
	if gv.Version != runtime.APIVersionInternal {
		clusterbytes, err := MarshalClusterConfigurationToBytes(&cfg.ClusterConfiguration, gv)
		if err != nil {
			return []byte{}, err
		}
		allFiles = append(allFiles, clusterbytes)
	}
	return bytes.Join(allFiles, []byte(kubeadmconstants.YAMLDocumentSeparator)), nil
}
func MarshalClusterConfigurationToBytes(clustercfg *kubeadmapi.ClusterConfiguration, gv schema.GroupVersion) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterbytes, err := kubeadmutil.MarshalToYamlForCodecs(clustercfg, gv, kubeadmscheme.Codecs)
	if err != nil {
		return []byte{}, err
	}
	allFiles := [][]byte{clusterbytes}
	componentConfigContent := map[string][]byte{}
	defaultedcfg := defaultedInternalConfig()
	for kind, registration := range componentconfigs.Known {
		realobj, ok := registration.GetFromInternalConfig(clustercfg)
		if !ok {
			continue
		}
		defaultedobj, ok := registration.GetFromInternalConfig(defaultedcfg)
		if !ok {
			return []byte{}, errors.New("couldn't create a default componentconfig object")
		}
		if !reflect.DeepEqual(realobj, defaultedobj) {
			contentBytes, err := registration.Marshal(realobj)
			if err != nil {
				return []byte{}, err
			}
			componentConfigContent[string(kind)] = contentBytes
		}
	}
	sortedComponentConfigFiles := consistentOrderByteSlice(componentConfigContent)
	allFiles = append(allFiles, sortedComponentConfigFiles...)
	return bytes.Join(allFiles, []byte(kubeadmconstants.YAMLDocumentSeparator)), nil
}
func consistentOrderByteSlice(content map[string][]byte) [][]byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	keys := []string{}
	sortedContent := [][]byte{}
	for key := range content {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		sortedContent = append(sortedContent, content[key])
	}
	return sortedContent
}
