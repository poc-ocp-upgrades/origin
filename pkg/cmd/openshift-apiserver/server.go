package openshift_apiserver

import (
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftapiserver"
	"github.com/openshift/origin/pkg/cmd/util"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/pkg/version"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/capabilities"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus"
	kubelettypes "k8s.io/kubernetes/pkg/kubelet/types"
)

var featureKeepRemovedNetworkingAPI = true

func RunOpenShiftAPIServer(serverConfig *openshiftcontrolplanev1.OpenShiftAPIServerConfig, stopCh <-chan struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	util.InitLogrus()
	capabilities.Initialize(capabilities.Capabilities{AllowPrivileged: true, PrivilegedSources: capabilities.PrivilegedSources{HostNetworkSources: []string{kubelettypes.ApiserverSource, kubelettypes.FileSource}, HostPIDSources: []string{kubelettypes.ApiserverSource, kubelettypes.FileSource}, HostIPCSources: []string{kubelettypes.ApiserverSource, kubelettypes.FileSource}}})
	openshiftAPIServerRuntimeConfig, err := openshiftapiserver.NewOpenshiftAPIConfig(serverConfig)
	if err != nil {
		return err
	}
	openshiftAPIServer, err := openshiftAPIServerRuntimeConfig.Complete().New(genericapiserver.NewEmptyDelegate(), featureKeepRemovedNetworkingAPI)
	if err != nil {
		return err
	}
	preparedOpenshiftAPIServer := openshiftAPIServer.GenericAPIServer.PrepareRun()
	klog.Infof("Starting master on %s (%s)", serverConfig.ServingInfo.BindAddress, version.Get().String())
	return preparedOpenshiftAPIServer.Run(stopCh)
}
