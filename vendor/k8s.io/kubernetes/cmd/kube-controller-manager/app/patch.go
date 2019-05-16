package app

import (
	"github.com/golang/glog"
	"k8s.io/client-go/informers"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/config"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	utilflag "k8s.io/kubernetes/pkg/util/flag"
	"path"
)

var InformerFactoryOverride informers.SharedInformerFactory

func ShimForOpenShift(controllerManagerOptions *options.KubeControllerManagerOptions, controllerManager *config.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(controllerManager.OpenShiftContext.OpenShiftConfig) == 0 {
		return nil
	}
	openshiftConfig, err := getOpenShiftConfig(controllerManager.OpenShiftContext.OpenShiftConfig)
	if err != nil {
		return err
	}
	if err := applyOpenShiftServiceServingCertCAFunc(path.Dir(controllerManager.OpenShiftContext.OpenShiftConfig), openshiftConfig); err != nil {
		return err
	}
	if err := applyOpenShiftGCConfig(controllerManager); err != nil {
		return err
	}
	if informers, err := newInformerFactory(controllerManager.Kubeconfig); err != nil {
		return err
	} else {
		InformerFactoryOverride = informers
	}
	return nil
}
func ShimFlagsForOpenShift(controllerManagerOptions *options.KubeControllerManagerOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(controllerManagerOptions.OpenShiftContext.OpenShiftConfig) == 0 {
		return nil
	}
	openshiftConfig, err := getOpenShiftConfig(controllerManagerOptions.OpenShiftContext.OpenShiftConfig)
	if err != nil {
		return err
	}
	if err := applyOpenShiftConfigFlags(controllerManagerOptions, openshiftConfig); err != nil {
		return err
	}
	for name, fs := range controllerManagerOptions.Flags(KnownControllers(), ControllersDisabledByDefault.List()).FlagSets {
		glog.V(1).Infof("FLAGSET: %s", name)
		utilflag.PrintFlags(fs)
	}
	return nil
}
