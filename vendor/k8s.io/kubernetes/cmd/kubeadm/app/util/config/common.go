package config

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"net"
	"reflect"
	"strings"
)

func MarshalKubeadmConfigObject(obj runtime.Object) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch internalcfg := obj.(type) {
	case *kubeadmapi.InitConfiguration:
		return MarshalInitConfigurationToBytes(internalcfg, kubeadmapiv1beta1.SchemeGroupVersion)
	case *kubeadmapi.ClusterConfiguration:
		return MarshalClusterConfigurationToBytes(internalcfg, kubeadmapiv1beta1.SchemeGroupVersion)
	default:
		return kubeadmutil.MarshalToYamlForCodecs(obj, kubeadmapiv1beta1.SchemeGroupVersion, kubeadmscheme.Codecs)
	}
}
func DetectUnsupportedVersion(b []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gvks, err := kubeadmutil.GroupVersionKindsFromBytes(b)
	if err != nil {
		return err
	}
	oldKnownAPIVersions := map[string]string{"kubeadm.k8s.io/v1alpha1": "v1.11", "kubeadm.k8s.io/v1alpha2": "v1.12"}
	knownKinds := map[string]bool{}
	for _, gvk := range gvks {
		if useKubeadmVersion := oldKnownAPIVersions[gvk.GroupVersion().String()]; len(useKubeadmVersion) != 0 {
			return errors.Errorf("your configuration file uses an old API spec: %q. Please use kubeadm %s instead and run 'kubeadm config migrate --old-config old.yaml --new-config new.yaml', which will write the new, similar spec using a newer API version.", gvk.GroupVersion().String(), useKubeadmVersion)
		}
		knownKinds[gvk.Kind] = true
	}
	mutuallyExclusive := []string{constants.InitConfigurationKind, constants.JoinConfigurationKind}
	mutuallyExclusiveCount := 0
	for _, kind := range mutuallyExclusive {
		if knownKinds[kind] {
			mutuallyExclusiveCount++
		}
	}
	if mutuallyExclusiveCount > 1 {
		klog.Warningf("WARNING: Detected resource kinds that may not apply: %v", mutuallyExclusive)
	}
	return nil
}
func NormalizeKubernetesVersion(cfg *kubeadmapi.ClusterConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeadmutil.KubernetesIsCIVersion(cfg.KubernetesVersion) {
		cfg.CIImageRepository = constants.DefaultCIImageRepository
	}
	ver, err := kubeadmutil.KubernetesReleaseVersion(cfg.KubernetesVersion)
	if err != nil {
		return err
	}
	cfg.KubernetesVersion = ver
	k8sVersion, err := version.ParseSemantic(cfg.KubernetesVersion)
	if err != nil {
		return errors.Wrapf(err, "couldn't parse Kubernetes version %q", cfg.KubernetesVersion)
	}
	if k8sVersion.LessThan(constants.MinimumControlPlaneVersion) {
		return errors.Errorf("this version of kubeadm only supports deploying clusters with the control plane version >= %s. Current version: %s", constants.MinimumControlPlaneVersion.String(), cfg.KubernetesVersion)
	}
	return nil
}
func LowercaseSANs(sans []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, san := range sans {
		lowercase := strings.ToLower(san)
		if lowercase != san {
			klog.V(1).Infof("lowercasing SAN %q to %q", san, lowercase)
			sans[i] = lowercase
		}
	}
}
func VerifyAPIServerBindAddress(address string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ip := net.ParseIP(address)
	if ip == nil {
		return errors.Errorf("cannot parse IP address: %s", address)
	}
	if !ip.IsGlobalUnicast() {
		return errors.Errorf("cannot use %q as the bind address for the API Server", address)
	}
	return nil
}
func ChooseAPIServerBindAddress(bindAddress net.IP) (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ip, err := netutil.ChooseBindAddress(bindAddress)
	if err != nil {
		if netutil.IsNoRoutesError(err) {
			klog.Warningf("WARNING: could not obtain a bind address for the API Server: %v; using: %s", err, constants.DefaultAPIServerBindAddress)
			defaultIP := net.ParseIP(constants.DefaultAPIServerBindAddress)
			if defaultIP == nil {
				return nil, errors.Errorf("cannot parse default IP address: %s", constants.DefaultAPIServerBindAddress)
			}
			return defaultIP, nil
		}
		return nil, err
	}
	if bindAddress != nil && !bindAddress.IsUnspecified() && !reflect.DeepEqual(ip, bindAddress) {
		klog.Warningf("WARNING: overriding requested API server bind address: requested %q, actual %q", bindAddress, ip)
	}
	return ip, nil
}
func MigrateOldConfigFromFile(cfgPath string) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newConfig := [][]byte{}
	cfgBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return []byte{}, err
	}
	gvks, err := kubeadmutil.GroupVersionKindsFromBytes(cfgBytes)
	if err != nil {
		return []byte{}, err
	}
	if kubeadmutil.GroupVersionKindsHasInitConfiguration(gvks...) || kubeadmutil.GroupVersionKindsHasClusterConfiguration(gvks...) {
		o, err := ConfigFileAndDefaultsToInternalConfig(cfgPath, &kubeadmapiv1beta1.InitConfiguration{})
		if err != nil {
			return []byte{}, err
		}
		b, err := MarshalKubeadmConfigObject(o)
		if err != nil {
			return []byte{}, err
		}
		newConfig = append(newConfig, b)
	}
	if kubeadmutil.GroupVersionKindsHasJoinConfiguration(gvks...) {
		o, err := JoinConfigFileAndDefaultsToInternalConfig(cfgPath, &kubeadmapiv1beta1.JoinConfiguration{})
		if err != nil {
			return []byte{}, err
		}
		b, err := MarshalKubeadmConfigObject(o)
		if err != nil {
			return []byte{}, err
		}
		newConfig = append(newConfig, b)
	}
	return bytes.Join(newConfig, []byte(constants.YAMLDocumentSeparator)), nil
}
