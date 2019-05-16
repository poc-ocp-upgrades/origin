package openshift_sdn

import (
	networkclient "github.com/openshift/client-go/network/clientset/versioned"
	networkinformers "github.com/openshift/client-go/network/informers/externalversions"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/network"
	kinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

type informers struct {
	KubeClient       kubernetes.Interface
	NetworkClient    networkclient.Interface
	KubeInformers    kinformers.SharedInformerFactory
	NetworkInformers networkinformers.SharedInformerFactory
}

func (sdn *OpenShiftSDN) buildInformers() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfig, err := configapi.GetKubeConfigOrInClusterConfig(sdn.NodeConfig.MasterKubeConfig, sdn.NodeConfig.MasterClientConnectionOverrides)
	if err != nil {
		return err
	}
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return err
	}
	networkClient, err := networkclient.NewForConfig(kubeConfig)
	if err != nil {
		return err
	}
	kubeInformers := kinformers.NewSharedInformerFactory(kubeClient, sdn.ProxyConfig.IPTables.SyncPeriod.Duration)
	networkInformers := networkinformers.NewSharedInformerFactory(networkClient, network.DefaultInformerResyncPeriod)
	sdn.informers = &informers{KubeClient: kubeClient, NetworkClient: networkClient, KubeInformers: kubeInformers, NetworkInformers: networkInformers}
	return nil
}
func (i *informers) start(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	i.KubeInformers.Start(stopCh)
	i.NetworkInformers.Start(stopCh)
}
