package kubemark

import (
 "fmt"
 "time"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 clientset "k8s.io/client-go/kubernetes"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/record"
 proxyapp "k8s.io/kubernetes/cmd/kube-proxy/app"
 "k8s.io/kubernetes/pkg/proxy"
 proxyconfig "k8s.io/kubernetes/pkg/proxy/config"
 "k8s.io/kubernetes/pkg/proxy/iptables"
 utiliptables "k8s.io/kubernetes/pkg/util/iptables"
 utilnode "k8s.io/kubernetes/pkg/util/node"
 utilsysctl "k8s.io/kubernetes/pkg/util/sysctl"
 utilexec "k8s.io/utils/exec"
 utilpointer "k8s.io/utils/pointer"
 "k8s.io/klog"
)

type HollowProxy struct{ ProxyServer *proxyapp.ProxyServer }
type FakeProxier struct{}

func (*FakeProxier) Sync() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) SyncLoop() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 select {}
}
func (*FakeProxier) OnServiceAdd(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnServiceUpdate(oldService, service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnServiceDelete(service *v1.Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnServiceSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnEndpointsAdd(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnEndpointsUpdate(oldEndpoints, endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnEndpointsDelete(endpoints *v1.Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (*FakeProxier) OnEndpointsSynced() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func NewHollowProxyOrDie(nodeName string, client clientset.Interface, eventClient v1core.EventsGetter, iptInterface utiliptables.Interface, sysctl utilsysctl.Interface, execer utilexec.Interface, broadcaster record.EventBroadcaster, recorder record.EventRecorder, useRealProxier bool, proxierSyncPeriod time.Duration, proxierMinSyncPeriod time.Duration) (*HollowProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var proxier proxy.ProxyProvider
 var serviceHandler proxyconfig.ServiceHandler
 var endpointsHandler proxyconfig.EndpointsHandler
 if useRealProxier {
  proxierIPTables, err := iptables.NewProxier(iptInterface, sysctl, execer, proxierSyncPeriod, proxierMinSyncPeriod, false, 0, "10.0.0.0/8", nodeName, utilnode.GetNodeIP(client, nodeName), recorder, nil, []string{})
  if err != nil {
   return nil, fmt.Errorf("unable to create proxier: %v", err)
  }
  proxier = proxierIPTables
  serviceHandler = proxierIPTables
  endpointsHandler = proxierIPTables
 } else {
  proxier = &FakeProxier{}
  serviceHandler = &FakeProxier{}
  endpointsHandler = &FakeProxier{}
 }
 nodeRef := &v1.ObjectReference{Kind: "Node", Name: nodeName, UID: types.UID(nodeName), Namespace: ""}
 return &HollowProxy{ProxyServer: &proxyapp.ProxyServer{Client: client, EventClient: eventClient, IptInterface: iptInterface, Proxier: proxier, Broadcaster: broadcaster, Recorder: recorder, ProxyMode: "fake", NodeRef: nodeRef, OOMScoreAdj: utilpointer.Int32Ptr(0), ResourceContainer: "", ConfigSyncPeriod: 30 * time.Second, ServiceEventHandler: serviceHandler, EndpointsEventHandler: endpointsHandler}}, nil
}
func (hp *HollowProxy) Run() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := hp.ProxyServer.Run(); err != nil {
  klog.Fatalf("Error while running proxy: %v\n", err)
 }
}
