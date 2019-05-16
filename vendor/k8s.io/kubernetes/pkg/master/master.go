package master

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	auditregistrationv1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	authenticationv1 "k8s.io/api/authentication/v1"
	authenticationv1beta1 "k8s.io/api/authentication/v1beta1"
	authorizationapiv1 "k8s.io/api/authorization/v1"
	authorizationapiv1beta1 "k8s.io/api/authorization/v1beta1"
	autoscalingapiv1 "k8s.io/api/autoscaling/v1"
	autoscalingapiv2beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingapiv2beta2 "k8s.io/api/autoscaling/v2beta2"
	batchapiv1 "k8s.io/api/batch/v1"
	batchapiv1beta1 "k8s.io/api/batch/v1beta1"
	batchapiv2alpha1 "k8s.io/api/batch/v2alpha1"
	certificatesapiv1beta1 "k8s.io/api/certificates/v1beta1"
	coordinationapiv1beta1 "k8s.io/api/coordination/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	eventsv1beta1 "k8s.io/api/events/v1beta1"
	extensionsapiv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingapiv1 "k8s.io/api/networking/v1"
	policyapiv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"
	schedulingapiv1beta1 "k8s.io/api/scheduling/v1beta1"
	settingsv1alpha1 "k8s.io/api/settings/v1alpha1"
	storageapiv1 "k8s.io/api/storage/v1"
	storageapiv1alpha1 "k8s.io/api/storage/v1alpha1"
	storageapiv1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apiserver/pkg/endpoints/discovery"
	"k8s.io/apiserver/pkg/registry/generic"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/healthz"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	storagefactory "k8s.io/apiserver/pkg/storage/storagebackend/factory"
	"k8s.io/client-go/informers"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	kubeoptions "k8s.io/kubernetes/pkg/kubeapiserver/options"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master/reconcilers"
	"k8s.io/kubernetes/pkg/master/tunneler"
	admissionregistrationrest "k8s.io/kubernetes/pkg/registry/admissionregistration/rest"
	appsrest "k8s.io/kubernetes/pkg/registry/apps/rest"
	auditregistrationrest "k8s.io/kubernetes/pkg/registry/auditregistration/rest"
	authenticationrest "k8s.io/kubernetes/pkg/registry/authentication/rest"
	authorizationrest "k8s.io/kubernetes/pkg/registry/authorization/rest"
	autoscalingrest "k8s.io/kubernetes/pkg/registry/autoscaling/rest"
	batchrest "k8s.io/kubernetes/pkg/registry/batch/rest"
	certificatesrest "k8s.io/kubernetes/pkg/registry/certificates/rest"
	coordinationrest "k8s.io/kubernetes/pkg/registry/coordination/rest"
	corerest "k8s.io/kubernetes/pkg/registry/core/rest"
	eventsrest "k8s.io/kubernetes/pkg/registry/events/rest"
	extensionsrest "k8s.io/kubernetes/pkg/registry/extensions/rest"
	networkingrest "k8s.io/kubernetes/pkg/registry/networking/rest"
	policyrest "k8s.io/kubernetes/pkg/registry/policy/rest"
	rbacrest "k8s.io/kubernetes/pkg/registry/rbac/rest"
	schedulingrest "k8s.io/kubernetes/pkg/registry/scheduling/rest"
	settingsrest "k8s.io/kubernetes/pkg/registry/settings/rest"
	storagerest "k8s.io/kubernetes/pkg/registry/storage/rest"
	"k8s.io/kubernetes/pkg/routes"
	"k8s.io/kubernetes/pkg/serviceaccount"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	DefaultEndpointReconcilerInterval = 10 * time.Second
	DefaultEndpointReconcilerTTL      = 15 * time.Second
)

type ExtraConfig struct {
	ClientCARegistrationHook    ClientCARegistrationHook
	APIResourceConfigSource     serverstorage.APIResourceConfigSource
	StorageFactory              serverstorage.StorageFactory
	EndpointReconcilerConfig    EndpointReconcilerConfig
	EventTTL                    time.Duration
	KubeletClientConfig         kubeletclient.KubeletClientConfig
	Tunneler                    tunneler.Tunneler
	EnableLogsSupport           bool
	ProxyTransport              http.RoundTripper
	ServiceIPRange              net.IPNet
	APIServerServiceIP          net.IP
	APIServerServicePort        int
	ServiceNodePortRange        utilnet.PortRange
	ExtraServicePorts           []apiv1.ServicePort
	ExtraEndpointPorts          []apiv1.EndpointPort
	KubernetesServiceNodePort   int
	MasterCount                 int
	MasterEndpointReconcileTTL  time.Duration
	EndpointReconcilerType      reconcilers.Type
	ServiceAccountIssuer        serviceaccount.TokenGenerator
	ServiceAccountMaxExpiration time.Duration
	VersionedInformers          informers.SharedInformerFactory
}
type Config struct {
	GenericConfig *genericapiserver.Config
	ExtraConfig   ExtraConfig
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }
type EndpointReconcilerConfig struct {
	Reconciler reconcilers.EndpointReconciler
	Interval   time.Duration
}
type Master struct {
	GenericAPIServer         *genericapiserver.GenericAPIServer
	ClientCARegistrationHook ClientCARegistrationHook
}

func (c *Config) createMasterCountReconciler() reconcilers.EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	endpointClient := corev1client.NewForConfigOrDie(c.GenericConfig.LoopbackClientConfig)
	return reconcilers.NewMasterCountEndpointReconciler(c.ExtraConfig.MasterCount, endpointClient)
}
func (c *Config) createNoneReconciler() reconcilers.EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return reconcilers.NewNoneEndpointReconciler()
}
func (c *Config) createLeaseReconciler() reconcilers.EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	endpointClient := corev1client.NewForConfigOrDie(c.GenericConfig.LoopbackClientConfig)
	ttl := c.ExtraConfig.MasterEndpointReconcileTTL
	config, err := c.ExtraConfig.StorageFactory.NewConfig(api.Resource("apiServerIPInfo"))
	if err != nil {
		klog.Fatalf("Error determining service IP ranges: %v", err)
	}
	leaseStorage, _, err := storagefactory.Create(*config)
	if err != nil {
		klog.Fatalf("Error creating storage factory: %v", err)
	}
	masterLeases := reconcilers.NewLeases(leaseStorage, "/masterleases/", ttl)
	return reconcilers.NewLeaseEndpointReconciler(endpointClient, masterLeases)
}
func (c *Config) createEndpointReconciler() reconcilers.EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Using reconciler: %v", c.ExtraConfig.EndpointReconcilerType)
	switch c.ExtraConfig.EndpointReconcilerType {
	case "", reconcilers.MasterCountReconcilerType:
		return c.createMasterCountReconciler()
	case reconcilers.LeaseEndpointReconcilerType:
		return c.createLeaseReconciler()
	case reconcilers.NoneEndpointReconcilerType:
		return c.createNoneReconciler()
	default:
		klog.Fatalf("Reconciler not implemented: %v", c.ExtraConfig.EndpointReconcilerType)
	}
	return nil
}
func (cfg *Config) Complete() CompletedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c := completedConfig{cfg.GenericConfig.Complete(cfg.ExtraConfig.VersionedInformers), &cfg.ExtraConfig}
	serviceIPRange, apiServerServiceIP, err := DefaultServiceIPRange(c.ExtraConfig.ServiceIPRange)
	if err != nil {
		klog.Fatalf("Error determining service IP ranges: %v", err)
	}
	if c.ExtraConfig.ServiceIPRange.IP == nil {
		c.ExtraConfig.ServiceIPRange = serviceIPRange
	}
	if c.ExtraConfig.APIServerServiceIP == nil {
		c.ExtraConfig.APIServerServiceIP = apiServerServiceIP
	}
	discoveryAddresses := discovery.DefaultAddresses{DefaultAddress: c.GenericConfig.ExternalAddress}
	discoveryAddresses.CIDRRules = append(discoveryAddresses.CIDRRules, discovery.CIDRRule{IPRange: c.ExtraConfig.ServiceIPRange, Address: net.JoinHostPort(c.ExtraConfig.APIServerServiceIP.String(), strconv.Itoa(c.ExtraConfig.APIServerServicePort))})
	c.GenericConfig.DiscoveryAddresses = discoveryAddresses
	if c.ExtraConfig.ServiceNodePortRange.Size == 0 {
		c.ExtraConfig.ServiceNodePortRange = kubeoptions.DefaultServiceNodePortRange
		klog.Infof("Node port range unspecified. Defaulting to %v.", c.ExtraConfig.ServiceNodePortRange)
	}
	if c.ExtraConfig.EndpointReconcilerConfig.Interval == 0 {
		c.ExtraConfig.EndpointReconcilerConfig.Interval = DefaultEndpointReconcilerInterval
	}
	if c.ExtraConfig.MasterEndpointReconcileTTL == 0 {
		c.ExtraConfig.MasterEndpointReconcileTTL = DefaultEndpointReconcilerTTL
	}
	if c.ExtraConfig.EndpointReconcilerConfig.Reconciler == nil {
		c.ExtraConfig.EndpointReconcilerConfig.Reconciler = cfg.createEndpointReconciler()
	}
	return CompletedConfig{&c}
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*Master, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if reflect.DeepEqual(c.ExtraConfig.KubeletClientConfig, kubeletclient.KubeletClientConfig{}) {
		return nil, fmt.Errorf("Master.New() called with empty config.KubeletClientConfig")
	}
	s, err := c.GenericConfig.New("kube-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	if c.ExtraConfig.EnableLogsSupport {
		routes.Logs{}.Install(s.Handler.GoRestfulContainer)
	}
	m := &Master{GenericAPIServer: s}
	if c.ExtraConfig.APIResourceConfigSource.VersionEnabled(apiv1.SchemeGroupVersion) {
		legacyRESTStorageProvider := corerest.LegacyRESTStorageProvider{StorageFactory: c.ExtraConfig.StorageFactory, ProxyTransport: c.ExtraConfig.ProxyTransport, KubeletClientConfig: c.ExtraConfig.KubeletClientConfig, EventTTL: c.ExtraConfig.EventTTL, ServiceIPRange: c.ExtraConfig.ServiceIPRange, ServiceNodePortRange: c.ExtraConfig.ServiceNodePortRange, LoopbackClientConfig: c.GenericConfig.LoopbackClientConfig, ServiceAccountIssuer: c.ExtraConfig.ServiceAccountIssuer, ServiceAccountMaxExpiration: c.ExtraConfig.ServiceAccountMaxExpiration, APIAudiences: c.GenericConfig.Authentication.APIAudiences}
		m.InstallLegacyAPI(&c, c.GenericConfig.RESTOptionsGetter, legacyRESTStorageProvider)
	}
	restStorageProviders := []RESTStorageProvider{auditregistrationrest.RESTStorageProvider{}, authenticationrest.RESTStorageProvider{Authenticator: c.GenericConfig.Authentication.Authenticator, APIAudiences: c.GenericConfig.Authentication.APIAudiences}, authorizationrest.RESTStorageProvider{Authorizer: c.GenericConfig.Authorization.Authorizer, RuleResolver: c.GenericConfig.RuleResolver}, autoscalingrest.RESTStorageProvider{}, batchrest.RESTStorageProvider{}, certificatesrest.RESTStorageProvider{}, coordinationrest.RESTStorageProvider{}, extensionsrest.RESTStorageProvider{}, networkingrest.RESTStorageProvider{}, policyrest.RESTStorageProvider{}, rbacrest.RESTStorageProvider{Authorizer: c.GenericConfig.Authorization.Authorizer}, schedulingrest.RESTStorageProvider{}, settingsrest.RESTStorageProvider{}, storagerest.RESTStorageProvider{}, appsrest.RESTStorageProvider{}, admissionregistrationrest.RESTStorageProvider{}, eventsrest.RESTStorageProvider{TTL: c.ExtraConfig.EventTTL}}
	m.InstallAPIs(c.ExtraConfig.APIResourceConfigSource, c.GenericConfig.RESTOptionsGetter, restStorageProviders...)
	if c.ExtraConfig.Tunneler != nil {
		m.installTunneler(c.ExtraConfig.Tunneler, corev1client.NewForConfigOrDie(c.GenericConfig.LoopbackClientConfig).Nodes())
	}
	m.GenericAPIServer.AddPostStartHookOrDie("ca-registration", c.ExtraConfig.ClientCARegistrationHook.PostStartHook)
	return m, nil
}
func (m *Master) InstallLegacyAPI(c *completedConfig, restOptionsGetter generic.RESTOptionsGetter, legacyRESTStorageProvider corerest.LegacyRESTStorageProvider) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	legacyRESTStorage, apiGroupInfo, err := legacyRESTStorageProvider.NewLegacyRESTStorage(restOptionsGetter)
	if err != nil {
		klog.Fatalf("Error building core storage: %v", err)
	}
	controllerName := "bootstrap-controller"
	coreClient := corev1client.NewForConfigOrDie(c.GenericConfig.LoopbackClientConfig)
	bootstrapController := c.NewBootstrapController(legacyRESTStorage, coreClient, coreClient, coreClient, coreClient.RESTClient())
	m.GenericAPIServer.AddPostStartHookOrDie(controllerName, bootstrapController.PostStartHook)
	m.GenericAPIServer.AddPreShutdownHookOrDie(controllerName, bootstrapController.PreShutdownHook)
	if err := m.GenericAPIServer.InstallLegacyAPIGroup(genericapiserver.DefaultLegacyAPIPrefix, &apiGroupInfo); err != nil {
		klog.Fatalf("Error in registering group versions: %v", err)
	}
}
func (m *Master) installTunneler(nodeTunneler tunneler.Tunneler, nodeClient corev1client.NodeInterface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeTunneler.Run(nodeAddressProvider{nodeClient}.externalAddresses)
	m.GenericAPIServer.AddHealthzChecks(healthz.NamedCheck("SSH Tunnel Check", tunneler.TunnelSyncHealthChecker(nodeTunneler)))
	prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "apiserver_proxy_tunnel_sync_latency_secs", Help: "The time since the last successful synchronization of the SSH tunnels for proxy requests."}, func() float64 {
		return float64(nodeTunneler.SecondsSinceSync())
	})
}

type RESTStorageProvider interface {
	GroupName() string
	NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool)
}

func (m *Master) InstallAPIs(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter, restStorageProviders ...RESTStorageProvider) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiGroupsInfo := []genericapiserver.APIGroupInfo{}
	for _, restStorageBuilder := range restStorageProviders {
		groupName := restStorageBuilder.GroupName()
		if !apiResourceConfigSource.AnyVersionForGroupEnabled(groupName) {
			klog.V(1).Infof("Skipping disabled API group %q.", groupName)
			continue
		}
		apiGroupInfo, enabled := restStorageBuilder.NewRESTStorage(apiResourceConfigSource, restOptionsGetter)
		if !enabled {
			klog.Warningf("Problem initializing API group %q, skipping.", groupName)
			continue
		}
		klog.V(1).Infof("Enabling API group %q.", groupName)
		if postHookProvider, ok := restStorageBuilder.(genericapiserver.PostStartHookProvider); ok {
			name, hook, err := postHookProvider.PostStartHook()
			if err != nil {
				klog.Fatalf("Error building PostStartHook: %v", err)
			}
			m.GenericAPIServer.AddPostStartHookOrDie(name, hook)
		}
		apiGroupsInfo = append(apiGroupsInfo, apiGroupInfo)
	}
	for i := range apiGroupsInfo {
		if err := m.GenericAPIServer.InstallAPIGroup(&apiGroupsInfo[i]); err != nil {
			klog.Fatalf("Error in registering group versions: %v", err)
		}
	}
}

type nodeAddressProvider struct{ nodeClient corev1client.NodeInterface }

func (n nodeAddressProvider) externalAddresses() ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	preferredAddressTypes := []apiv1.NodeAddressType{apiv1.NodeExternalIP}
	nodes, err := n.nodeClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	addrs := []string{}
	for ix := range nodes.Items {
		node := &nodes.Items[ix]
		addr, err := nodeutil.GetPreferredNodeAddress(node, preferredAddressTypes)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}
func DefaultAPIResourceConfigSource() *serverstorage.ResourceConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := serverstorage.NewResourceConfig()
	ret.EnableVersions(admissionregistrationv1beta1.SchemeGroupVersion, apiv1.SchemeGroupVersion, appsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion, appsv1.SchemeGroupVersion, authenticationv1.SchemeGroupVersion, authenticationv1beta1.SchemeGroupVersion, authorizationapiv1.SchemeGroupVersion, authorizationapiv1beta1.SchemeGroupVersion, autoscalingapiv1.SchemeGroupVersion, autoscalingapiv2beta1.SchemeGroupVersion, autoscalingapiv2beta2.SchemeGroupVersion, batchapiv1.SchemeGroupVersion, batchapiv1beta1.SchemeGroupVersion, certificatesapiv1beta1.SchemeGroupVersion, coordinationapiv1beta1.SchemeGroupVersion, eventsv1beta1.SchemeGroupVersion, extensionsapiv1beta1.SchemeGroupVersion, networkingapiv1.SchemeGroupVersion, policyapiv1beta1.SchemeGroupVersion, rbacv1.SchemeGroupVersion, rbacv1beta1.SchemeGroupVersion, storageapiv1.SchemeGroupVersion, storageapiv1beta1.SchemeGroupVersion, schedulingapiv1beta1.SchemeGroupVersion)
	ret.DisableVersions(auditregistrationv1alpha1.SchemeGroupVersion, admissionregistrationv1alpha1.SchemeGroupVersion, batchapiv2alpha1.SchemeGroupVersion, rbacv1alpha1.SchemeGroupVersion, schedulingv1alpha1.SchemeGroupVersion, settingsv1alpha1.SchemeGroupVersion, storageapiv1alpha1.SchemeGroupVersion)
	return ret
}
