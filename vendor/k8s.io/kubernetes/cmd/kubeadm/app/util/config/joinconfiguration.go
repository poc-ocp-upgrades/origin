package config

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/config/strict"
)

func SetJoinDynamicDefaults(cfg *kubeadmapi.JoinConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addMasterTaint := false
	if cfg.ControlPlane != nil {
		addMasterTaint = true
	}
	if err := SetNodeRegistrationDynamicDefaults(&cfg.NodeRegistration, addMasterTaint); err != nil {
		return err
	}
	if err := SetJoinControlPlaneDefaults(cfg.ControlPlane); err != nil {
		return err
	}
	return nil
}
func SetJoinControlPlaneDefaults(cfg *kubeadmapi.JoinControlPlane) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg != nil {
		if err := SetAPIEndpointDynamicDefaults(&cfg.LocalAPIEndpoint); err != nil {
			return err
		}
	}
	return nil
}
func JoinConfigFileAndDefaultsToInternalConfig(cfgPath string, defaultversionedcfg *kubeadmapiv1beta1.JoinConfiguration) (*kubeadmapi.JoinConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	internalcfg := &kubeadmapi.JoinConfiguration{}
	if cfgPath != "" {
		klog.V(1).Infoln("loading configuration from the given file")
		b, err := ioutil.ReadFile(cfgPath)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read config from %q ", cfgPath)
		}
		if err := DetectUnsupportedVersion(b); err != nil {
			return nil, err
		}
		gvkmap, err := kubeadmutil.SplitYAMLDocuments(b)
		if err != nil {
			return nil, err
		}
		joinBytes := []byte{}
		for gvk, bytes := range gvkmap {
			if gvk.Kind == constants.JoinConfigurationKind {
				joinBytes = bytes
				strict.VerifyUnmarshalStrict(bytes, gvk)
			}
		}
		if len(joinBytes) == 0 {
			return nil, errors.Errorf("no %s found in config file %q", constants.JoinConfigurationKind, cfgPath)
		}
		if err := runtime.DecodeInto(kubeadmscheme.Codecs.UniversalDecoder(), joinBytes, internalcfg); err != nil {
			return nil, err
		}
	} else {
		kubeadmscheme.Scheme.Default(defaultversionedcfg)
		kubeadmscheme.Scheme.Convert(defaultversionedcfg, internalcfg, nil)
	}
	if err := SetJoinDynamicDefaults(internalcfg); err != nil {
		return nil, err
	}
	if err := validation.ValidateJoinConfiguration(internalcfg).ToAggregate(); err != nil {
		return nil, err
	}
	return internalcfg, nil
}
