package gce

import (
	"cloud.google.com/go/compute/metadata"
	"context"
	"fmt"
	goformat "fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	computealpha "google.golang.org/api/compute/v0.alpha"
	computebeta "google.golang.org/api/compute/v0.beta"
	compute "google.golang.org/api/compute/v1"
	container "google.golang.org/api/container/v1"
	gcfg "gopkg.in/gcfg.v1"
	"io"
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
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/controller"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/version"
	"net/http"
	goos "os"
	"runtime"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	gotime "time"
)

const (
	ProviderName                 = "gce"
	k8sNodeRouteTag              = "k8s-node-route"
	gceAffinityTypeNone          = "NONE"
	gceAffinityTypeClientIP      = "CLIENT_IP"
	gceAffinityTypeClientIPProto = "CLIENT_IP_PROTO"
	operationPollInterval        = time.Second
	operationPollTimeoutDuration = time.Hour
	maxPages                     = 25
	maxTargetPoolCreateInstances = 200
	gceHcCheckIntervalSeconds    = int64(8)
	gceHcTimeoutSeconds          = int64(1)
	gceHcHealthyThreshold        = int64(1)
	gceHcUnhealthyThreshold      = int64(3)
	gceComputeAPIEndpoint        = "https://www.googleapis.com/compute/v1/"
	gceComputeAPIEndpointBeta    = "https://www.googleapis.com/compute/beta/"
)

type gceObject interface{ MarshalJSON() ([]byte, error) }
type Cloud struct {
	ClusterID                ClusterID
	service                  *compute.Service
	serviceBeta              *computebeta.Service
	serviceAlpha             *computealpha.Service
	containerService         *container.Service
	tpuService               *tpuService
	client                   clientset.Interface
	clientBuilder            controller.ControllerClientBuilder
	eventBroadcaster         record.EventBroadcaster
	eventRecorder            record.EventRecorder
	projectID                string
	region                   string
	regional                 bool
	localZone                string
	managedZones             []string
	networkURL               string
	isLegacyNetwork          bool
	subnetworkURL            string
	secondaryRangeName       string
	networkProjectID         string
	onXPN                    bool
	nodeTags                 []string
	lastComputedNodeTags     []string
	lastKnownNodeNames       sets.String
	computeNodeTagLock       sync.Mutex
	nodeInstancePrefix       string
	useMetadataServer        bool
	operationPollRateLimiter flowcontrol.RateLimiter
	manager                  diskServiceManager
	nodeZonesLock            sync.Mutex
	nodeZones                map[string]sets.String
	nodeInformerSynced       cache.InformerSynced
	sharedResourceLock       sync.Mutex
	AlphaFeatureGate         *AlphaFeatureGate
	c                        cloud.Cloud
	s                        *cloud.Service
}
type ConfigGlobal struct {
	TokenURL             string   `gcfg:"token-url"`
	TokenBody            string   `gcfg:"token-body"`
	ProjectID            string   `gcfg:"project-id"`
	NetworkProjectID     string   `gcfg:"network-project-id"`
	NetworkName          string   `gcfg:"network-name"`
	SubnetworkName       string   `gcfg:"subnetwork-name"`
	SecondaryRangeName   string   `gcfg:"secondary-range-name"`
	NodeTags             []string `gcfg:"node-tags"`
	NodeInstancePrefix   string   `gcfg:"node-instance-prefix"`
	Regional             bool     `gcfg:"regional"`
	Multizone            bool     `gcfg:"multizone"`
	APIEndpoint          string   `gcfg:"api-endpoint"`
	ContainerAPIEndpoint string   `gcfg:"container-api-endpoint"`
	LocalZone            string   `gcfg:"local-zone"`
	AlphaFeatures        []string `gcfg:"alpha-features"`
}
type ConfigFile struct {
	Global ConfigGlobal `gcfg:"global"`
}
type CloudConfig struct {
	APIEndpoint          string
	ContainerAPIEndpoint string
	ProjectID            string
	NetworkProjectID     string
	Region               string
	Regional             bool
	Zone                 string
	ManagedZones         []string
	NetworkName          string
	NetworkURL           string
	SubnetworkName       string
	SubnetworkURL        string
	SecondaryRangeName   string
	NodeTags             []string
	NodeInstancePrefix   string
	TokenSource          oauth2.TokenSource
	UseMetadataServer    bool
	AlphaFeatureGate     *AlphaFeatureGate
}

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newGCECloud(config)
	})
}

type Services struct {
	GA    *compute.Service
	Alpha *computealpha.Service
	Beta  *computebeta.Service
}

func (g *Cloud) ComputeServices() *Services {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Services{g.service, g.serviceAlpha, g.serviceBeta}
}
func (g *Cloud) Compute() cloud.Cloud {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.c
}
func (g *Cloud) ContainerService() *container.Service {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.containerService
}
func newGCECloud(config io.Reader) (gceCloud *Cloud, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cloudConfig *CloudConfig
	var configFile *ConfigFile
	if config != nil {
		configFile, err = readConfig(config)
		if err != nil {
			return nil, err
		}
		klog.Infof("Using GCE provider config %+v", configFile)
	}
	cloudConfig, err = generateCloudConfig(configFile)
	if err != nil {
		return nil, err
	}
	return CreateGCECloud(cloudConfig)
}
func readConfig(reader io.Reader) (*ConfigFile, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := &ConfigFile{}
	if err := gcfg.FatalOnly(gcfg.ReadInto(cfg, reader)); err != nil {
		klog.Errorf("Couldn't read config: %v", err)
		return nil, err
	}
	return cfg, nil
}
func generateCloudConfig(configFile *ConfigFile) (cloudConfig *CloudConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cloudConfig = &CloudConfig{}
	cloudConfig.TokenSource = google.ComputeTokenSource("")
	cloudConfig.UseMetadataServer = true
	cloudConfig.AlphaFeatureGate = NewAlphaFeatureGate([]string{})
	if configFile != nil {
		if configFile.Global.APIEndpoint != "" {
			cloudConfig.APIEndpoint = configFile.Global.APIEndpoint
		}
		if configFile.Global.ContainerAPIEndpoint != "" {
			cloudConfig.ContainerAPIEndpoint = configFile.Global.ContainerAPIEndpoint
		}
		if configFile.Global.TokenURL != "" {
			if configFile.Global.TokenURL == "nil" {
				cloudConfig.TokenSource = nil
			} else {
				cloudConfig.TokenSource = NewAltTokenSource(configFile.Global.TokenURL, configFile.Global.TokenBody)
			}
		}
		cloudConfig.NodeTags = configFile.Global.NodeTags
		cloudConfig.NodeInstancePrefix = configFile.Global.NodeInstancePrefix
		cloudConfig.AlphaFeatureGate = NewAlphaFeatureGate(configFile.Global.AlphaFeatures)
	}
	if configFile == nil || configFile.Global.ProjectID == "" || configFile.Global.LocalZone == "" {
		cloudConfig.ProjectID, cloudConfig.Zone, err = getProjectAndZone()
		if err != nil {
			return nil, err
		}
	}
	if configFile != nil {
		if configFile.Global.ProjectID != "" {
			cloudConfig.ProjectID = configFile.Global.ProjectID
		}
		if configFile.Global.LocalZone != "" {
			cloudConfig.Zone = configFile.Global.LocalZone
		}
		if configFile.Global.NetworkProjectID != "" {
			cloudConfig.NetworkProjectID = configFile.Global.NetworkProjectID
		}
	}
	cloudConfig.Region, err = GetGCERegion(cloudConfig.Zone)
	if err != nil {
		return nil, err
	}
	if configFile != nil && configFile.Global.Regional {
		cloudConfig.Regional = true
	}
	cloudConfig.ManagedZones = []string{cloudConfig.Zone}
	if configFile != nil && (configFile.Global.Multizone || configFile.Global.Regional) {
		cloudConfig.ManagedZones = nil
	}
	if configFile != nil && configFile.Global.NetworkName != "" {
		if strings.Contains(configFile.Global.NetworkName, "/") {
			cloudConfig.NetworkURL = configFile.Global.NetworkName
		} else {
			cloudConfig.NetworkName = configFile.Global.NetworkName
		}
	} else {
		cloudConfig.NetworkName, err = getNetworkNameViaMetadata()
		if err != nil {
			return nil, err
		}
	}
	if configFile != nil && configFile.Global.SubnetworkName != "" {
		if strings.Contains(configFile.Global.SubnetworkName, "/") {
			cloudConfig.SubnetworkURL = configFile.Global.SubnetworkName
		} else {
			cloudConfig.SubnetworkName = configFile.Global.SubnetworkName
		}
	}
	if configFile != nil {
		cloudConfig.SecondaryRangeName = configFile.Global.SecondaryRangeName
	}
	return cloudConfig, err
}
func CreateGCECloud(config *CloudConfig) (*Cloud, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	version := strings.TrimLeft(strings.Split(strings.Split(version.Get().GitVersion, "-")[0], "+")[0], "v")
	userAgent := fmt.Sprintf("Kubernetes/%s (%s %s)", version, runtime.GOOS, runtime.GOARCH)
	if config.NetworkProjectID == "" {
		config.NetworkProjectID = config.ProjectID
	}
	client, err := newOauthClient(config.TokenSource)
	if err != nil {
		return nil, err
	}
	service, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	service.UserAgent = userAgent
	client, err = newOauthClient(config.TokenSource)
	if err != nil {
		return nil, err
	}
	serviceBeta, err := computebeta.New(client)
	if err != nil {
		return nil, err
	}
	serviceBeta.UserAgent = userAgent
	client, err = newOauthClient(config.TokenSource)
	if err != nil {
		return nil, err
	}
	serviceAlpha, err := computealpha.New(client)
	if err != nil {
		return nil, err
	}
	serviceAlpha.UserAgent = userAgent
	if config.APIEndpoint != "" {
		service.BasePath = fmt.Sprintf("%sprojects/", config.APIEndpoint)
		serviceBeta.BasePath = fmt.Sprintf("%sprojects/", strings.Replace(config.APIEndpoint, "v1", "beta", -1))
		serviceAlpha.BasePath = fmt.Sprintf("%sprojects/", strings.Replace(config.APIEndpoint, "v1", "alpha", -1))
	}
	containerService, err := container.New(client)
	if err != nil {
		return nil, err
	}
	containerService.UserAgent = userAgent
	if config.ContainerAPIEndpoint != "" {
		containerService.BasePath = config.ContainerAPIEndpoint
	}
	tpuService, err := newTPUService(client)
	if err != nil {
		return nil, err
	}
	projID, netProjID := tryConvertToProjectNames(config.ProjectID, config.NetworkProjectID, service)
	onXPN := projID != netProjID
	var networkURL string
	var subnetURL string
	var isLegacyNetwork bool
	if config.NetworkURL != "" {
		networkURL = config.NetworkURL
	} else if config.NetworkName != "" {
		networkURL = gceNetworkURL(config.APIEndpoint, netProjID, config.NetworkName)
	} else {
		klog.Warningf("No network name or URL specified.")
	}
	if config.SubnetworkURL != "" {
		subnetURL = config.SubnetworkURL
	} else if config.SubnetworkName != "" {
		subnetURL = gceSubnetworkURL(config.APIEndpoint, netProjID, config.Region, config.SubnetworkName)
	} else {
		if networkName := lastComponent(networkURL); networkName != "" {
			var n *compute.Network
			if n, err = getNetwork(service, netProjID, networkName); err != nil {
				klog.Warningf("Could not retrieve network %q; err: %v", networkName, err)
			} else {
				switch typeOfNetwork(n) {
				case netTypeLegacy:
					klog.Infof("Network %q is type legacy - no subnetwork", networkName)
					isLegacyNetwork = true
				case netTypeCustom:
					klog.Warningf("Network %q is type custom - cannot auto select a subnetwork", networkName)
				case netTypeAuto:
					subnetURL, err = determineSubnetURL(service, netProjID, networkName, config.Region)
					if err != nil {
						klog.Warningf("Could not determine subnetwork for network %q and region %v; err: %v", networkName, config.Region, err)
					} else {
						klog.Infof("Auto selecting subnetwork %q", subnetURL)
					}
				}
			}
		}
	}
	if len(config.ManagedZones) == 0 {
		config.ManagedZones, err = getZonesForRegion(service, config.ProjectID, config.Region)
		if err != nil {
			return nil, err
		}
	}
	if len(config.ManagedZones) > 1 {
		klog.Infof("managing multiple zones: %v", config.ManagedZones)
	}
	operationPollRateLimiter := flowcontrol.NewTokenBucketRateLimiter(5, 5)
	gce := &Cloud{service: service, serviceAlpha: serviceAlpha, serviceBeta: serviceBeta, containerService: containerService, tpuService: tpuService, projectID: projID, networkProjectID: netProjID, onXPN: onXPN, region: config.Region, regional: config.Regional, localZone: config.Zone, managedZones: config.ManagedZones, networkURL: networkURL, isLegacyNetwork: isLegacyNetwork, subnetworkURL: subnetURL, secondaryRangeName: config.SecondaryRangeName, nodeTags: config.NodeTags, nodeInstancePrefix: config.NodeInstancePrefix, useMetadataServer: config.UseMetadataServer, operationPollRateLimiter: operationPollRateLimiter, AlphaFeatureGate: config.AlphaFeatureGate, nodeZones: map[string]sets.String{}}
	gce.manager = &gceServiceManager{gce}
	gce.s = &cloud.Service{GA: service, Alpha: serviceAlpha, Beta: serviceBeta, ProjectRouter: &gceProjectRouter{gce}, RateLimiter: &gceRateLimiter{gce}}
	gce.c = cloud.NewGCE(gce.s)
	return gce, nil
}
func (g *Cloud) SetRateLimiter(rl cloud.RateLimiter) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rl != nil {
		g.s.RateLimiter = rl
	}
}
func determineSubnetURL(service *compute.Service, networkProjectID, networkName, region string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subnets, err := listSubnetworksOfNetwork(service, networkProjectID, networkName, region)
	if err != nil {
		return "", err
	}
	autoSubnets, err := subnetsInCIDR(subnets, autoSubnetIPRange)
	if err != nil {
		return "", err
	}
	if len(autoSubnets) == 0 {
		return "", fmt.Errorf("no subnet exists in auto CIDR")
	}
	if len(autoSubnets) > 1 {
		return "", fmt.Errorf("multiple subnetworks in the same region exist in auto CIDR")
	}
	return autoSubnets[0].SelfLink, nil
}
func tryConvertToProjectNames(configProject, configNetworkProject string, service *compute.Service) (projID, netProjID string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projID = configProject
	if isProjectNumber(projID) {
		projName, err := getProjectID(service, projID)
		if err != nil {
			klog.Warningf("Failed to retrieve project %v while trying to retrieve its name. err %v", projID, err)
		} else {
			projID = projName
		}
	}
	netProjID = projID
	if configNetworkProject != configProject {
		netProjID = configNetworkProject
	}
	if isProjectNumber(netProjID) {
		netProjName, err := getProjectID(service, netProjID)
		if err != nil {
			klog.Warningf("Failed to retrieve network project %v while trying to retrieve its name. err %v", netProjID, err)
		} else {
			netProjID = netProjName
		}
	}
	return projID, netProjID
}
func (g *Cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.clientBuilder = clientBuilder
	g.client = clientBuilder.ClientOrDie("cloud-provider")
	if g.OnXPN() {
		g.eventBroadcaster = record.NewBroadcaster()
		g.eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: g.client.CoreV1().Events("")})
		g.eventRecorder = g.eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "g-cloudprovider"})
	}
	go g.watchClusterID(stop)
}
func (g *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g, true
}
func (g *Cloud) Instances() (cloudprovider.Instances, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g, true
}
func (g *Cloud) Zones() (cloudprovider.Zones, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g, true
}
func (g *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g, true
}
func (g *Cloud) Routes() (cloudprovider.Routes, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g, true
}
func (g *Cloud) ProviderName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ProviderName
}
func (g *Cloud) ProjectID() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.projectID
}
func (g *Cloud) NetworkProjectID() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.networkProjectID
}
func (g *Cloud) Region() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.region
}
func (g *Cloud) OnXPN() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.onXPN
}
func (g *Cloud) NetworkURL() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.networkURL
}
func (g *Cloud) SubnetworkURL() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.subnetworkURL
}
func (g *Cloud) IsLegacyNetwork() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.isLegacyNetwork
}
func (g *Cloud) SetInformers(informerFactory informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Setting up informers for Cloud")
	nodeInformer := informerFactory.Core().V1().Nodes().Informer()
	nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		node := obj.(*v1.Node)
		g.updateNodeZones(nil, node)
	}, UpdateFunc: func(prev, obj interface{}) {
		prevNode := prev.(*v1.Node)
		newNode := obj.(*v1.Node)
		if newNode.Labels[kubeletapis.LabelZoneFailureDomain] == prevNode.Labels[kubeletapis.LabelZoneFailureDomain] {
			return
		}
		g.updateNodeZones(prevNode, newNode)
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
		g.updateNodeZones(node, nil)
	}})
	g.nodeInformerSynced = nodeInformer.HasSynced
}
func (g *Cloud) updateNodeZones(prevNode, newNode *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.nodeZonesLock.Lock()
	defer g.nodeZonesLock.Unlock()
	if prevNode != nil {
		prevZone, ok := prevNode.ObjectMeta.Labels[kubeletapis.LabelZoneFailureDomain]
		if ok {
			g.nodeZones[prevZone].Delete(prevNode.ObjectMeta.Name)
			if g.nodeZones[prevZone].Len() == 0 {
				g.nodeZones[prevZone] = nil
			}
		}
	}
	if newNode != nil {
		newZone, ok := newNode.ObjectMeta.Labels[kubeletapis.LabelZoneFailureDomain]
		if ok {
			if g.nodeZones[newZone] == nil {
				g.nodeZones[newZone] = sets.NewString()
			}
			g.nodeZones[newZone].Insert(newNode.ObjectMeta.Name)
		}
	}
}
func (g *Cloud) HasClusterID() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func isProjectNumber(idOrNumber string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := strconv.ParseUint(idOrNumber, 10, 64)
	return err == nil
}

var _ cloudprovider.Interface = (*Cloud)(nil)

func gceNetworkURL(apiEndpoint, project, network string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if apiEndpoint == "" {
		apiEndpoint = gceComputeAPIEndpoint
	}
	return apiEndpoint + strings.Join([]string{"projects", project, "global", "networks", network}, "/")
}
func gceSubnetworkURL(apiEndpoint, project, region, subnetwork string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if apiEndpoint == "" {
		apiEndpoint = gceComputeAPIEndpoint
	}
	return apiEndpoint + strings.Join([]string{"projects", project, "regions", region, "subnetworks", subnetwork}, "/")
}
func getRegionInURL(urlStr string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fields := strings.Split(urlStr, "/")
	for i, v := range fields {
		if v == "regions" && i < len(fields)-1 {
			return fields[i+1]
		}
	}
	return ""
}
func getNetworkNameViaMetadata() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := metadata.Get("instance/network-interfaces/0/network")
	if err != nil {
		return "", err
	}
	parts := strings.Split(result, "/")
	if len(parts) != 4 {
		return "", fmt.Errorf("unexpected response: %s", result)
	}
	return parts[3], nil
}
func getNetwork(svc *compute.Service, networkProjectID, networkID string) (*compute.Network, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return svc.Networks.Get(networkProjectID, networkID).Do()
}
func listSubnetworksOfNetwork(svc *compute.Service, networkProjectID, networkID, region string) ([]*compute.Subnetwork, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var subnets []*compute.Subnetwork
	err := svc.Subnetworks.List(networkProjectID, region).Filter(fmt.Sprintf("network eq .*/%v$", networkID)).Pages(context.Background(), func(res *compute.SubnetworkList) error {
		subnets = append(subnets, res.Items...)
		return nil
	})
	return subnets, err
}
func getProjectID(svc *compute.Service, projectNumberOrID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proj, err := svc.Projects.Get(projectNumberOrID).Do()
	if err != nil {
		return "", err
	}
	return proj.Name, nil
}
func getZonesForRegion(svc *compute.Service, projectID, region string) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	listCall := svc.Zones.List(projectID)
	res, err := listCall.Do()
	if err != nil {
		return nil, fmt.Errorf("unexpected response listing zones: %v", err)
	}
	zones := []string{}
	for _, zone := range res.Items {
		regionName := lastComponent(zone.Region)
		if regionName == region {
			zones = append(zones, zone.Name)
		}
	}
	return zones, nil
}
func findSubnetForRegion(subnetURLs []string, region string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, url := range subnetURLs {
		if thisRegion := getRegionInURL(url); thisRegion == region {
			return url
		}
	}
	return ""
}
func newOauthClient(tokenSource oauth2.TokenSource) (*http.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if tokenSource == nil {
		var err error
		tokenSource, err = google.DefaultTokenSource(oauth2.NoContext, compute.CloudPlatformScope, compute.ComputeScope)
		klog.Infof("Using DefaultTokenSource %#v", tokenSource)
		if err != nil {
			return nil, err
		}
	} else {
		klog.Infof("Using existing Token Source %#v", tokenSource)
	}
	backoff := wait.Backoff{Duration: time.Second, Factor: 1.4, Steps: 10}
	if err := wait.ExponentialBackoff(backoff, func() (bool, error) {
		if _, err := tokenSource.Token(); err != nil {
			klog.Errorf("error fetching initial token: %v", err)
			return false, nil
		}
		return true, nil
	}); err != nil {
		return nil, err
	}
	return oauth2.NewClient(oauth2.NoContext, tokenSource), nil
}
func (manager *gceServiceManager) getProjectsAPIEndpoint() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projectsAPIEndpoint := gceComputeAPIEndpoint + "projects/"
	if manager.gce.service != nil {
		projectsAPIEndpoint = manager.gce.service.BasePath
	}
	return projectsAPIEndpoint
}
func (manager *gceServiceManager) getProjectsAPIEndpointBeta() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projectsAPIEndpoint := gceComputeAPIEndpointBeta + "projects/"
	if manager.gce.service != nil {
		projectsAPIEndpoint = manager.gce.serviceBeta.BasePath
	}
	return projectsAPIEndpoint
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
