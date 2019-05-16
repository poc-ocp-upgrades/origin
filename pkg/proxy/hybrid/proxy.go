package hybrid

import (
	"fmt"
	goformat "fmt"
	unidlingapi "github.com/openshift/origin/pkg/unidling/api"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/proxy"
	proxyconfig "k8s.io/kubernetes/pkg/proxy/config"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

type HybridProxier struct {
	unidlingServiceHandler   proxyconfig.ServiceHandler
	unidlingEndpointsHandler proxyconfig.EndpointsHandler
	mainEndpointsHandler     proxyconfig.EndpointsHandler
	mainServicesHandler      proxyconfig.ServiceHandler
	mainProxy                proxy.ProxyProvider
	unidlingProxy            proxy.ProxyProvider
	syncPeriod               time.Duration
	serviceLister            corev1listers.ServiceLister
	usingUserspace           map[types.NamespacedName]bool
	usingUserspaceLock       sync.Mutex
	switchedToUserspace      map[types.NamespacedName]bool
	switchedToUserspaceLock  sync.Mutex
}

func NewHybridProxier(unidlingEndpointsHandler proxyconfig.EndpointsHandler, unidlingServiceHandler proxyconfig.ServiceHandler, mainEndpointsHandler proxyconfig.EndpointsHandler, mainServicesHandler proxyconfig.ServiceHandler, mainProxy proxy.ProxyProvider, unidlingProxy proxy.ProxyProvider, syncPeriod time.Duration, serviceLister corev1listers.ServiceLister) (*HybridProxier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &HybridProxier{unidlingEndpointsHandler: unidlingEndpointsHandler, unidlingServiceHandler: unidlingServiceHandler, mainEndpointsHandler: mainEndpointsHandler, mainServicesHandler: mainServicesHandler, mainProxy: mainProxy, unidlingProxy: unidlingProxy, syncPeriod: syncPeriod, serviceLister: serviceLister, usingUserspace: make(map[types.NamespacedName]bool), switchedToUserspace: make(map[types.NamespacedName]bool)}, nil
}
func (p *HybridProxier) OnServiceAdd(service *corev1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	if isUsingUserspace, ok := p.usingUserspace[svcName]; ok && isUsingUserspace {
		klog.V(6).Infof("hybrid proxy: add svc %s/%s in unidling proxy", service.Namespace, service.Name)
		p.unidlingServiceHandler.OnServiceAdd(service)
	} else {
		klog.V(6).Infof("hybrid proxy: add svc %s/%s in main proxy", service.Namespace, service.Name)
		p.mainServicesHandler.OnServiceAdd(service)
	}
}
func (p *HybridProxier) OnServiceUpdate(oldService, service *corev1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	if isUsingUserspace, ok := p.usingUserspace[svcName]; ok && isUsingUserspace {
		klog.V(6).Infof("hybrid proxy: update svc %s/%s in unidling proxy", service.Namespace, service.Name)
		p.unidlingServiceHandler.OnServiceUpdate(oldService, service)
	} else {
		klog.V(6).Infof("hybrid proxy: update svc %s/%s in main proxy", service.Namespace, service.Name)
		p.mainServicesHandler.OnServiceUpdate(oldService, service)
	}
}
func (p *HybridProxier) OnServiceDelete(service *corev1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	p.switchedToUserspaceLock.Lock()
	defer p.switchedToUserspaceLock.Unlock()
	if isUsingUserspace, ok := p.usingUserspace[svcName]; ok && isUsingUserspace {
		klog.V(6).Infof("hybrid proxy: del svc %s/%s in unidling proxy", service.Namespace, service.Name)
		p.unidlingServiceHandler.OnServiceDelete(service)
	} else {
		klog.V(6).Infof("hybrid proxy: del svc %s/%s in main proxy", service.Namespace, service.Name)
		p.mainServicesHandler.OnServiceDelete(service)
	}
	delete(p.switchedToUserspace, svcName)
}
func (p *HybridProxier) OnServiceSynced() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.unidlingServiceHandler.OnServiceSynced()
	p.mainServicesHandler.OnServiceSynced()
	klog.V(6).Infof("hybrid proxy: services synced")
}
func (p *HybridProxier) shouldEndpointsUseUserspace(endpoints *corev1.Endpoints) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hasEndpoints := false
	for _, subset := range endpoints.Subsets {
		if len(subset.Addresses) > 0 {
			hasEndpoints = true
			break
		}
	}
	if !hasEndpoints {
		if _, ok := endpoints.Annotations[unidlingapi.IdledAtAnnotation]; ok {
			return true
		}
	}
	return false
}
func (p *HybridProxier) switchService(name types.NamespacedName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.switchedToUserspaceLock.Lock()
	defer p.switchedToUserspaceLock.Unlock()
	switched, ok := p.switchedToUserspace[name]
	if ok && p.usingUserspace[name] == switched {
		klog.V(6).Infof("hybrid proxy: ignoring duplicate switchService(%s/%s)", name.Namespace, name.Name)
		return
	}
	svc, err := p.serviceLister.Services(name.Namespace).Get(name.Name)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Error while getting service %s/%s from cache: %v", name.Namespace, name.String(), err))
		return
	}
	if p.usingUserspace[name] {
		klog.V(6).Infof("hybrid proxy: switching svc %s/%s to unidling proxy", svc.Namespace, svc.Name)
		p.unidlingServiceHandler.OnServiceAdd(svc)
		p.mainServicesHandler.OnServiceDelete(svc)
	} else {
		klog.V(6).Infof("hybrid proxy: switching svc %s/%s to main proxy", svc.Namespace, svc.Name)
		p.mainServicesHandler.OnServiceAdd(svc)
		p.unidlingServiceHandler.OnServiceDelete(svc)
	}
	p.switchedToUserspace[name] = p.usingUserspace[name]
}
func (p *HybridProxier) OnEndpointsAdd(endpoints *corev1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(6).Infof("hybrid proxy: (always) add ep %s/%s in unidling proxy", endpoints.Namespace, endpoints.Name)
	p.unidlingEndpointsHandler.OnEndpointsAdd(endpoints)
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	svcName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	wasUsingUserspace, knownEndpoints := p.usingUserspace[svcName]
	p.usingUserspace[svcName] = p.shouldEndpointsUseUserspace(endpoints)
	if !p.usingUserspace[svcName] {
		klog.V(6).Infof("hybrid proxy: add ep %s/%s in main proxy", endpoints.Namespace, endpoints.Name)
		p.mainEndpointsHandler.OnEndpointsAdd(endpoints)
	}
	if knownEndpoints && wasUsingUserspace != p.usingUserspace[svcName] {
		p.switchService(svcName)
	}
}
func (p *HybridProxier) OnEndpointsUpdate(oldEndpoints, endpoints *corev1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(6).Infof("hybrid proxy: (always) update ep %s/%s in unidling proxy", endpoints.Namespace, endpoints.Name)
	p.unidlingEndpointsHandler.OnEndpointsUpdate(oldEndpoints, endpoints)
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	svcName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	wasUsingUserspace, knownEndpoints := p.usingUserspace[svcName]
	p.usingUserspace[svcName] = p.shouldEndpointsUseUserspace(endpoints)
	if !knownEndpoints {
		utilruntime.HandleError(fmt.Errorf("received update for unknown endpoints %s", svcName.String()))
		return
	}
	isSwitch := wasUsingUserspace != p.usingUserspace[svcName]
	if !isSwitch && !p.usingUserspace[svcName] {
		klog.V(6).Infof("hybrid proxy: update ep %s/%s in main proxy", endpoints.Namespace, endpoints.Name)
		p.mainEndpointsHandler.OnEndpointsUpdate(oldEndpoints, endpoints)
		return
	}
	if p.usingUserspace[svcName] {
		klog.V(6).Infof("hybrid proxy: del ep %s/%s in main proxy", endpoints.Namespace, endpoints.Name)
		p.mainEndpointsHandler.OnEndpointsDelete(oldEndpoints)
	} else {
		klog.V(6).Infof("hybrid proxy: add ep %s/%s in main proxy", endpoints.Namespace, endpoints.Name)
		p.mainEndpointsHandler.OnEndpointsAdd(endpoints)
	}
	p.switchService(svcName)
}
func (p *HybridProxier) OnEndpointsDelete(endpoints *corev1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(6).Infof("hybrid proxy: (always) del ep %s/%s in unidling proxy", endpoints.Namespace, endpoints.Name)
	p.unidlingEndpointsHandler.OnEndpointsDelete(endpoints)
	p.usingUserspaceLock.Lock()
	defer p.usingUserspaceLock.Unlock()
	svcName := types.NamespacedName{Namespace: endpoints.Namespace, Name: endpoints.Name}
	usingUserspace, knownEndpoints := p.usingUserspace[svcName]
	if !knownEndpoints {
		utilruntime.HandleError(fmt.Errorf("received delete for unknown endpoints %s", svcName.String()))
		return
	}
	if !usingUserspace {
		klog.V(6).Infof("hybrid proxy: del ep %s/%s in main proxy", endpoints.Namespace, endpoints.Name)
		p.mainEndpointsHandler.OnEndpointsDelete(endpoints)
	}
	delete(p.usingUserspace, svcName)
}
func (p *HybridProxier) OnEndpointsSynced() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.unidlingEndpointsHandler.OnEndpointsSynced()
	p.mainEndpointsHandler.OnEndpointsSynced()
	klog.V(6).Infof("hybrid proxy: endpoints synced")
}
func (p *HybridProxier) Sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.mainProxy.Sync()
	p.unidlingProxy.Sync()
	klog.V(6).Infof("hybrid proxy: proxies synced")
}
func (p *HybridProxier) SyncLoop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	go p.mainProxy.SyncLoop()
	go p.unidlingProxy.SyncLoop()
	select {}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
