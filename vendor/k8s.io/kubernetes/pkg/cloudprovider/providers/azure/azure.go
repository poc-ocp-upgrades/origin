package azure

import (
	"fmt"
	goformat "fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"io"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/flowcontrol"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/azure/auth"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/version"
	goos "os"
	godefaultruntime "runtime"
	"sigs.k8s.io/yaml"
	"strings"
	"sync"
	"time"
	gotime "time"
)

const (
	CloudProviderName            = "azure"
	rateLimitQPSDefault          = 1.0
	rateLimitBucketDefault       = 5
	backoffRetriesDefault        = 6
	backoffExponentDefault       = 1.5
	backoffDurationDefault       = 5
	backoffJitterDefault         = 1.0
	maximumLoadBalancerRuleCount = 148
	vmTypeVMSS                   = "vmss"
	vmTypeStandard               = "standard"
	loadBalancerSkuBasic         = "basic"
	loadBalancerSkuStandard      = "standard"
	externalResourceGroupLabel   = "kubernetes.azure.com/resource-group"
	managedByAzureLabel          = "kubernetes.azure.com/managed"
)

var (
	defaultExcludeMasterFromStandardLB = true
)
var _ cloudprovider.PVLabeler = (*Cloud)(nil)

type Config struct {
	auth.AzureAuthConfig
	ResourceGroup                     string  `json:"resourceGroup" yaml:"resourceGroup"`
	Location                          string  `json:"location" yaml:"location"`
	VnetName                          string  `json:"vnetName" yaml:"vnetName"`
	VnetResourceGroup                 string  `json:"vnetResourceGroup" yaml:"vnetResourceGroup"`
	SubnetName                        string  `json:"subnetName" yaml:"subnetName"`
	SecurityGroupName                 string  `json:"securityGroupName" yaml:"securityGroupName"`
	RouteTableName                    string  `json:"routeTableName" yaml:"routeTableName"`
	PrimaryAvailabilitySetName        string  `json:"primaryAvailabilitySetName" yaml:"primaryAvailabilitySetName"`
	VMType                            string  `json:"vmType" yaml:"vmType"`
	PrimaryScaleSetName               string  `json:"primaryScaleSetName" yaml:"primaryScaleSetName"`
	CloudProviderBackoff              bool    `json:"cloudProviderBackoff" yaml:"cloudProviderBackoff"`
	CloudProviderBackoffRetries       int     `json:"cloudProviderBackoffRetries" yaml:"cloudProviderBackoffRetries"`
	CloudProviderBackoffExponent      float64 `json:"cloudProviderBackoffExponent" yaml:"cloudProviderBackoffExponent"`
	CloudProviderBackoffDuration      int     `json:"cloudProviderBackoffDuration" yaml:"cloudProviderBackoffDuration"`
	CloudProviderBackoffJitter        float64 `json:"cloudProviderBackoffJitter" yaml:"cloudProviderBackoffJitter"`
	CloudProviderRateLimit            bool    `json:"cloudProviderRateLimit" yaml:"cloudProviderRateLimit"`
	CloudProviderRateLimitQPS         float32 `json:"cloudProviderRateLimitQPS" yaml:"cloudProviderRateLimitQPS"`
	CloudProviderRateLimitBucket      int     `json:"cloudProviderRateLimitBucket" yaml:"cloudProviderRateLimitBucket"`
	CloudProviderRateLimitQPSWrite    float32 `json:"cloudProviderRateLimitQPSWrite" yaml:"cloudProviderRateLimitQPSWrite"`
	CloudProviderRateLimitBucketWrite int     `json:"cloudProviderRateLimitBucketWrite" yaml:"cloudProviderRateLimitBucketWrite"`
	UseInstanceMetadata               bool    `json:"useInstanceMetadata" yaml:"useInstanceMetadata"`
	LoadBalancerSku                   string  `json:"loadBalancerSku" yaml:"loadBalancerSku"`
	ExcludeMasterFromStandardLB       *bool   `json:"excludeMasterFromStandardLB" yaml:"excludeMasterFromStandardLB"`
	MaximumLoadBalancerRuleCount      int     `json:"maximumLoadBalancerRuleCount" yaml:"maximumLoadBalancerRuleCount"`
}
type Cloud struct {
	Config
	Environment                     azure.Environment
	RoutesClient                    RoutesClient
	SubnetsClient                   SubnetsClient
	InterfacesClient                InterfacesClient
	RouteTablesClient               RouteTablesClient
	LoadBalancerClient              LoadBalancersClient
	PublicIPAddressesClient         PublicIPAddressesClient
	SecurityGroupsClient            SecurityGroupsClient
	VirtualMachinesClient           VirtualMachinesClient
	StorageAccountClient            StorageAccountClient
	DisksClient                     DisksClient
	FileClient                      FileClient
	resourceRequestBackoff          wait.Backoff
	metadata                        *InstanceMetadataService
	vmSet                           VMSet
	nodeCachesLock                  sync.Mutex
	nodeZones                       map[string]sets.String
	nodeResourceGroups              map[string]string
	unmanagedNodes                  sets.String
	nodeInformerSynced              cache.InformerSynced
	routeCIDRsLock                  sync.Mutex
	routeCIDRs                      map[string]string
	VirtualMachineScaleSetsClient   VirtualMachineScaleSetsClient
	VirtualMachineScaleSetVMsClient VirtualMachineScaleSetVMsClient
	VirtualMachineSizesClient       VirtualMachineSizesClient
	kubeClient                      clientset.Interface
	eventBroadcaster                record.EventBroadcaster
	eventRecorder                   record.EventRecorder
	vmCache                         *timedCache
	lbCache                         *timedCache
	nsgCache                        *timedCache
	rtCache                         *timedCache
	*BlobDiskController
	*ManagedDiskController
	*controllerCommon
}

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cloudprovider.RegisterCloudProvider(CloudProviderName, NewCloud)
}
func NewCloud(configReader io.Reader) (cloudprovider.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := parseConfig(configReader)
	if err != nil {
		return nil, err
	}
	if config.VMType == "" {
		config.VMType = vmTypeStandard
	}
	env, err := auth.ParseAzureEnvironment(config.Cloud)
	if err != nil {
		return nil, err
	}
	servicePrincipalToken, err := auth.GetServicePrincipalToken(&config.AzureAuthConfig, env)
	if err != nil {
		return nil, err
	}
	operationPollRateLimiter := flowcontrol.NewFakeAlwaysRateLimiter()
	operationPollRateLimiterWrite := flowcontrol.NewFakeAlwaysRateLimiter()
	if config.CloudProviderRateLimit {
		if config.CloudProviderRateLimitQPS == 0 {
			config.CloudProviderRateLimitQPS = rateLimitQPSDefault
		}
		if config.CloudProviderRateLimitBucket == 0 {
			config.CloudProviderRateLimitBucket = rateLimitBucketDefault
		}
		if config.CloudProviderRateLimitQPSWrite == 0 {
			config.CloudProviderRateLimitQPSWrite = rateLimitQPSDefault
		}
		if config.CloudProviderRateLimitBucketWrite == 0 {
			config.CloudProviderRateLimitBucketWrite = rateLimitBucketDefault
		}
		operationPollRateLimiter = flowcontrol.NewTokenBucketRateLimiter(config.CloudProviderRateLimitQPS, config.CloudProviderRateLimitBucket)
		operationPollRateLimiterWrite = flowcontrol.NewTokenBucketRateLimiter(config.CloudProviderRateLimitQPSWrite, config.CloudProviderRateLimitBucketWrite)
		klog.V(2).Infof("Azure cloudprovider (read ops) using rate limit config: QPS=%g, bucket=%d", config.CloudProviderRateLimitQPS, config.CloudProviderRateLimitBucket)
		klog.V(2).Infof("Azure cloudprovider (write ops) using rate limit config: QPS=%g, bucket=%d", config.CloudProviderRateLimitQPSWrite, config.CloudProviderRateLimitBucketWrite)
	}
	if config.ExcludeMasterFromStandardLB == nil {
		config.ExcludeMasterFromStandardLB = &defaultExcludeMasterFromStandardLB
	}
	azClientConfig := &azClientConfig{subscriptionID: config.SubscriptionID, resourceManagerEndpoint: env.ResourceManagerEndpoint, servicePrincipalToken: servicePrincipalToken, rateLimiterReader: operationPollRateLimiter, rateLimiterWriter: operationPollRateLimiterWrite}
	az := Cloud{Config: *config, Environment: *env, nodeZones: map[string]sets.String{}, nodeResourceGroups: map[string]string{}, unmanagedNodes: sets.NewString(), routeCIDRs: map[string]string{}, DisksClient: newAzDisksClient(azClientConfig), RoutesClient: newAzRoutesClient(azClientConfig), SubnetsClient: newAzSubnetsClient(azClientConfig), InterfacesClient: newAzInterfacesClient(azClientConfig), RouteTablesClient: newAzRouteTablesClient(azClientConfig), LoadBalancerClient: newAzLoadBalancersClient(azClientConfig), SecurityGroupsClient: newAzSecurityGroupsClient(azClientConfig), StorageAccountClient: newAzStorageAccountClient(azClientConfig), VirtualMachinesClient: newAzVirtualMachinesClient(azClientConfig), PublicIPAddressesClient: newAzPublicIPAddressesClient(azClientConfig), VirtualMachineSizesClient: newAzVirtualMachineSizesClient(azClientConfig), VirtualMachineScaleSetsClient: newAzVirtualMachineScaleSetsClient(azClientConfig), VirtualMachineScaleSetVMsClient: newAzVirtualMachineScaleSetVMsClient(azClientConfig), FileClient: &azureFileClient{env: *env}}
	if az.CloudProviderBackoff {
		if az.CloudProviderBackoffRetries == 0 {
			az.CloudProviderBackoffRetries = backoffRetriesDefault
		}
		if az.CloudProviderBackoffExponent == 0 {
			az.CloudProviderBackoffExponent = backoffExponentDefault
		}
		if az.CloudProviderBackoffDuration == 0 {
			az.CloudProviderBackoffDuration = backoffDurationDefault
		}
		if az.CloudProviderBackoffJitter == 0 {
			az.CloudProviderBackoffJitter = backoffJitterDefault
		}
		az.resourceRequestBackoff = wait.Backoff{Steps: az.CloudProviderBackoffRetries, Factor: az.CloudProviderBackoffExponent, Duration: time.Duration(az.CloudProviderBackoffDuration) * time.Second, Jitter: az.CloudProviderBackoffJitter}
		klog.V(2).Infof("Azure cloudprovider using try backoff: retries=%d, exponent=%f, duration=%d, jitter=%f", az.CloudProviderBackoffRetries, az.CloudProviderBackoffExponent, az.CloudProviderBackoffDuration, az.CloudProviderBackoffJitter)
	}
	az.metadata, err = NewInstanceMetadataService(metadataURL)
	if err != nil {
		return nil, err
	}
	if az.MaximumLoadBalancerRuleCount == 0 {
		az.MaximumLoadBalancerRuleCount = maximumLoadBalancerRuleCount
	}
	if strings.EqualFold(vmTypeVMSS, az.Config.VMType) {
		az.vmSet, err = newScaleSet(&az)
		if err != nil {
			return nil, err
		}
	} else {
		az.vmSet = newAvailabilitySet(&az)
	}
	az.vmCache, err = az.newVMCache()
	if err != nil {
		return nil, err
	}
	az.lbCache, err = az.newLBCache()
	if err != nil {
		return nil, err
	}
	az.nsgCache, err = az.newNSGCache()
	if err != nil {
		return nil, err
	}
	az.rtCache, err = az.newRouteTableCache()
	if err != nil {
		return nil, err
	}
	if err := initDiskControllers(&az); err != nil {
		return nil, err
	}
	return &az, nil
}
func parseConfig(configReader io.Reader) (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var config Config
	if configReader == nil {
		return &config, nil
	}
	configContents, err := ioutil.ReadAll(configReader)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configContents, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
func (az *Cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	az.kubeClient = clientBuilder.ClientOrDie("azure-cloud-provider")
	az.eventBroadcaster = record.NewBroadcaster()
	az.eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: az.kubeClient.CoreV1().Events("")})
	az.eventRecorder = az.eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "azure-cloud-provider"})
}
func (az *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az, true
}
func (az *Cloud) Instances() (cloudprovider.Instances, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az, true
}
func (az *Cloud) Zones() (cloudprovider.Zones, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az, true
}
func (az *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (az *Cloud) Routes() (cloudprovider.Routes, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az, true
}
func (az *Cloud) HasClusterID() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (az *Cloud) ProviderName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return CloudProviderName
}
func configureUserAgent(client *autorest.Client) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	k8sVersion := version.Get().GitVersion
	client.UserAgent = fmt.Sprintf("%s; kubernetes-cloudprovider/%s", client.UserAgent, k8sVersion)
}
func initDiskControllers(az *Cloud) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	common := &controllerCommon{location: az.Location, storageEndpointSuffix: az.Environment.StorageEndpointSuffix, resourceGroup: az.ResourceGroup, subscriptionID: az.SubscriptionID, cloud: az}
	az.BlobDiskController = &BlobDiskController{common: common}
	az.ManagedDiskController = &ManagedDiskController{common: common}
	az.controllerCommon = common
	return nil
}
func (az *Cloud) SetInformers(informerFactory informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Setting up informers for Azure cloud provider")
	nodeInformer := informerFactory.Core().V1().Nodes().Informer()
	nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		node := obj.(*v1.Node)
		az.updateNodeCaches(nil, node)
	}, UpdateFunc: func(prev, obj interface{}) {
		prevNode := prev.(*v1.Node)
		newNode := obj.(*v1.Node)
		if newNode.Labels[kubeletapis.LabelZoneFailureDomain] == prevNode.Labels[kubeletapis.LabelZoneFailureDomain] {
			return
		}
		az.updateNodeCaches(prevNode, newNode)
	}, DeleteFunc: func(obj interface{}) {
		node, isNode := obj.(*v1.Node)
		if !isNode {
			deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				klog.Errorf("Received unexpected object: %v", obj)
				return
			}
			node, ok = deletedState.Obj.(*v1.Node)
			if !ok {
				klog.Errorf("DeletedFinalStateUnknown contained non-Node object: %v", deletedState.Obj)
				return
			}
		}
		az.updateNodeCaches(node, nil)
	}})
	az.nodeInformerSynced = nodeInformer.HasSynced
}
func (az *Cloud) updateNodeCaches(prevNode, newNode *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	az.nodeCachesLock.Lock()
	defer az.nodeCachesLock.Unlock()
	if prevNode != nil {
		prevZone, ok := prevNode.ObjectMeta.Labels[kubeletapis.LabelZoneFailureDomain]
		if ok && az.isAvailabilityZone(prevZone) {
			az.nodeZones[prevZone].Delete(prevNode.ObjectMeta.Name)
			if az.nodeZones[prevZone].Len() == 0 {
				az.nodeZones[prevZone] = nil
			}
		}
		_, ok = prevNode.ObjectMeta.Labels[externalResourceGroupLabel]
		if ok {
			delete(az.nodeResourceGroups, prevNode.ObjectMeta.Name)
		}
		managed, ok := prevNode.ObjectMeta.Labels[managedByAzureLabel]
		if ok && managed == "false" {
			az.unmanagedNodes.Delete(prevNode.ObjectMeta.Name)
		}
	}
	if newNode != nil {
		newZone, ok := newNode.ObjectMeta.Labels[kubeletapis.LabelZoneFailureDomain]
		if ok && az.isAvailabilityZone(newZone) {
			if az.nodeZones[newZone] == nil {
				az.nodeZones[newZone] = sets.NewString()
			}
			az.nodeZones[newZone].Insert(newNode.ObjectMeta.Name)
		}
		newRG, ok := newNode.ObjectMeta.Labels[externalResourceGroupLabel]
		if ok && len(newRG) > 0 {
			az.nodeResourceGroups[newNode.ObjectMeta.Name] = newRG
		}
		managed, ok := newNode.ObjectMeta.Labels[managedByAzureLabel]
		if ok && managed == "false" {
			az.unmanagedNodes.Insert(newNode.ObjectMeta.Name)
		}
	}
}
func (az *Cloud) GetActiveZones() (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.nodeInformerSynced == nil {
		return nil, fmt.Errorf("Azure cloud provider doesn't have informers set")
	}
	az.nodeCachesLock.Lock()
	defer az.nodeCachesLock.Unlock()
	if !az.nodeInformerSynced() {
		return nil, fmt.Errorf("node informer is not synced when trying to GetActiveZones")
	}
	zones := sets.NewString()
	for zone, nodes := range az.nodeZones {
		if len(nodes) > 0 {
			zones.Insert(zone)
		}
	}
	return zones, nil
}
func (az *Cloud) GetLocation() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az.Location
}
func (az *Cloud) GetNodeResourceGroup(nodeName string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.nodeInformerSynced == nil {
		return az.ResourceGroup, nil
	}
	az.nodeCachesLock.Lock()
	defer az.nodeCachesLock.Unlock()
	if !az.nodeInformerSynced() {
		return "", fmt.Errorf("node informer is not synced when trying to GetNodeResourceGroup")
	}
	if cachedRG, ok := az.nodeResourceGroups[nodeName]; ok {
		return cachedRG, nil
	}
	return az.ResourceGroup, nil
}
func (az *Cloud) GetResourceGroups() (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.nodeInformerSynced == nil {
		return sets.NewString(az.ResourceGroup), nil
	}
	az.nodeCachesLock.Lock()
	defer az.nodeCachesLock.Unlock()
	if !az.nodeInformerSynced() {
		return nil, fmt.Errorf("node informer is not synced when trying to GetResourceGroups")
	}
	resourceGroups := sets.NewString(az.ResourceGroup)
	for _, rg := range az.nodeResourceGroups {
		resourceGroups.Insert(rg)
	}
	return resourceGroups, nil
}
func (az *Cloud) GetUnmanagedNodes() (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if az.nodeInformerSynced == nil {
		return nil, nil
	}
	az.nodeCachesLock.Lock()
	defer az.nodeCachesLock.Unlock()
	if !az.nodeInformerSynced() {
		return nil, fmt.Errorf("node informer is not synced when trying to GetUnmanagedNodes")
	}
	return sets.NewString(az.unmanagedNodes.List()...), nil
}
func (az *Cloud) ShouldNodeExcludedFromLoadBalancer(node *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	labels := node.ObjectMeta.Labels
	if rg, ok := labels[externalResourceGroupLabel]; ok && rg != az.ResourceGroup {
		return true
	}
	if managed, ok := labels[managedByAzureLabel]; ok && managed == "false" {
		return true
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
