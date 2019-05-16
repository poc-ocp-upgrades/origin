package aws

import (
	"context"
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/sts"
	gcfg "gopkg.in/gcfg.v1"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/v1/service"
	"k8s.io/kubernetes/pkg/controller"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"net"
	goos "os"
	"path"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	gotime "time"
)

const NLBHealthCheckRuleDescription = "kubernetes.io/rule/nlb/health"
const NLBClientRuleDescription = "kubernetes.io/rule/nlb/client"
const NLBMtuDiscoveryRuleDescription = "kubernetes.io/rule/nlb/mtu"
const ProviderName = "aws"
const TagNameKubernetesService = "kubernetes.io/service-name"
const TagNameSubnetInternalELB = "kubernetes.io/role/internal-elb"
const TagNameSubnetPublicELB = "kubernetes.io/role/elb"
const ServiceAnnotationLoadBalancerType = "service.beta.kubernetes.io/aws-load-balancer-type"
const ServiceAnnotationLoadBalancerInternal = "service.beta.kubernetes.io/aws-load-balancer-internal"
const ServiceAnnotationLoadBalancerProxyProtocol = "service.beta.kubernetes.io/aws-load-balancer-proxy-protocol"
const ServiceAnnotationLoadBalancerAccessLogEmitInterval = "service.beta.kubernetes.io/aws-load-balancer-access-log-emit-interval"
const ServiceAnnotationLoadBalancerAccessLogEnabled = "service.beta.kubernetes.io/aws-load-balancer-access-log-enabled"
const ServiceAnnotationLoadBalancerAccessLogS3BucketName = "service.beta.kubernetes.io/aws-load-balancer-access-log-s3-bucket-name"
const ServiceAnnotationLoadBalancerAccessLogS3BucketPrefix = "service.beta.kubernetes.io/aws-load-balancer-access-log-s3-bucket-prefix"
const ServiceAnnotationLoadBalancerConnectionDrainingEnabled = "service.beta.kubernetes.io/aws-load-balancer-connection-draining-enabled"
const ServiceAnnotationLoadBalancerConnectionDrainingTimeout = "service.beta.kubernetes.io/aws-load-balancer-connection-draining-timeout"
const ServiceAnnotationLoadBalancerConnectionIdleTimeout = "service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout"
const ServiceAnnotationLoadBalancerCrossZoneLoadBalancingEnabled = "service.beta.kubernetes.io/aws-load-balancer-cross-zone-load-balancing-enabled"
const ServiceAnnotationLoadBalancerExtraSecurityGroups = "service.beta.kubernetes.io/aws-load-balancer-extra-security-groups"
const ServiceAnnotationLoadBalancerSecurityGroups = "service.beta.kubernetes.io/aws-load-balancer-security-groups"
const ServiceAnnotationLoadBalancerCertificate = "service.beta.kubernetes.io/aws-load-balancer-ssl-cert"
const ServiceAnnotationLoadBalancerSSLPorts = "service.beta.kubernetes.io/aws-load-balancer-ssl-ports"
const ServiceAnnotationLoadBalancerSSLNegotiationPolicy = "service.beta.kubernetes.io/aws-load-balancer-ssl-negotiation-policy"
const ServiceAnnotationLoadBalancerBEProtocol = "service.beta.kubernetes.io/aws-load-balancer-backend-protocol"
const ServiceAnnotationLoadBalancerAdditionalTags = "service.beta.kubernetes.io/aws-load-balancer-additional-resource-tags"
const ServiceAnnotationLoadBalancerHCHealthyThreshold = "service.beta.kubernetes.io/aws-load-balancer-healthcheck-healthy-threshold"
const ServiceAnnotationLoadBalancerHCUnhealthyThreshold = "service.beta.kubernetes.io/aws-load-balancer-healthcheck-unhealthy-threshold"
const ServiceAnnotationLoadBalancerHCTimeout = "service.beta.kubernetes.io/aws-load-balancer-healthcheck-timeout"
const ServiceAnnotationLoadBalancerHCInterval = "service.beta.kubernetes.io/aws-load-balancer-healthcheck-interval"
const volumeAttachmentStuck = "VolumeAttachmentStuck"
const nodeWithImpairedVolumes = "NodeWithImpairedVolumes"
const (
	volumeAttachmentStatusConsecutiveErrorLimit = 10
	volumeAttachmentStatusInitialDelay          = 1 * time.Second
	volumeAttachmentStatusFactor                = 1.8
	volumeAttachmentStatusSteps                 = 13
	createTagInitialDelay                       = 1 * time.Second
	createTagFactor                             = 2.0
	createTagSteps                              = 9
	encryptedCheckInterval                      = 1 * time.Second
	encryptedCheckTimeout                       = 30 * time.Second
	filterNodeLimit                             = 150
)

var awsTagNameMasterRoles = sets.NewString("kubernetes.io/role/master", "k8s.io/role/master")
var backendProtocolMapping = map[string]string{"https": "https", "http": "https", "ssl": "ssl", "tcp": "ssl"}

const MaxReadThenCreateRetries = 30
const DefaultVolumeType = "gp2"

var once sync.Once
var _ cloudprovider.PVLabeler = (*Cloud)(nil)

type Services interface {
	Compute(region string) (EC2, error)
	LoadBalancing(region string) (ELB, error)
	LoadBalancingV2(region string) (ELBV2, error)
	Autoscaling(region string) (ASG, error)
	Metadata() (EC2Metadata, error)
	KeyManagement(region string) (KMS, error)
}
type EC2 interface {
	DescribeInstances(request *ec2.DescribeInstancesInput) ([]*ec2.Instance, error)
	AttachVolume(*ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error)
	DetachVolume(request *ec2.DetachVolumeInput) (resp *ec2.VolumeAttachment, err error)
	DescribeVolumes(request *ec2.DescribeVolumesInput) ([]*ec2.Volume, error)
	CreateVolume(request *ec2.CreateVolumeInput) (resp *ec2.Volume, err error)
	DeleteVolume(*ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)
	ModifyVolume(*ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error)
	DescribeVolumeModifications(*ec2.DescribeVolumesModificationsInput) ([]*ec2.VolumeModification, error)
	DescribeSecurityGroups(request *ec2.DescribeSecurityGroupsInput) ([]*ec2.SecurityGroup, error)
	CreateSecurityGroup(*ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error)
	DeleteSecurityGroup(request *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error)
	AuthorizeSecurityGroupIngress(*ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupIngress(*ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error)
	DescribeSubnets(*ec2.DescribeSubnetsInput) ([]*ec2.Subnet, error)
	CreateTags(*ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error)
	DescribeRouteTables(request *ec2.DescribeRouteTablesInput) ([]*ec2.RouteTable, error)
	CreateRoute(request *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error)
	DeleteRoute(request *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error)
	ModifyInstanceAttribute(request *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error)
	DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
}
type ELB interface {
	CreateLoadBalancer(*elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)
	DeleteLoadBalancer(*elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)
	DescribeLoadBalancers(*elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)
	AddTags(*elb.AddTagsInput) (*elb.AddTagsOutput, error)
	RegisterInstancesWithLoadBalancer(*elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	DeregisterInstancesFromLoadBalancer(*elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	CreateLoadBalancerPolicy(*elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)
	SetLoadBalancerPoliciesForBackendServer(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesOfListener(input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
	DescribeLoadBalancerPolicies(input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DetachLoadBalancerFromSubnets(*elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	AttachLoadBalancerToSubnets(*elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	CreateLoadBalancerListeners(*elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)
	DeleteLoadBalancerListeners(*elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)
	ApplySecurityGroupsToLoadBalancer(*elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	ConfigureHealthCheck(*elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)
	DescribeLoadBalancerAttributes(*elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributes(*elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)
}
type ELBV2 interface {
	AddTags(input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error)
	CreateLoadBalancer(*elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
	DescribeLoadBalancers(*elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error)
	ModifyLoadBalancerAttributes(*elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributes(*elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	CreateTargetGroup(*elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error)
	DescribeTargetGroups(*elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	ModifyTargetGroup(*elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error)
	DeleteTargetGroup(*elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error)
	DescribeTargetHealth(input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetGroupAttributes(*elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributes(*elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	RegisterTargets(*elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
	DeregisterTargets(*elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error)
	CreateListener(*elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error)
	DescribeListeners(*elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error)
	DeleteListener(*elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error)
	ModifyListener(*elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error)
	WaitUntilLoadBalancersDeleted(*elbv2.DescribeLoadBalancersInput) error
}
type ASG interface {
	UpdateAutoScalingGroup(*autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error)
	DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
}
type KMS interface {
	DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)
}
type EC2Metadata interface {
	GetMetadata(path string) (string, error)
}

const (
	VolumeTypeIO1 = "io1"
	VolumeTypeGP2 = "gp2"
	VolumeTypeSC1 = "sc1"
	VolumeTypeST1 = "st1"
)
const (
	MinTotalIOPS = 100
	MaxTotalIOPS = 20000
)

type VolumeOptions struct {
	CapacityGB       int
	Tags             map[string]string
	VolumeType       string
	AvailabilityZone string
	IOPSPerGB        int
	Encrypted        bool
	KmsKeyID         string
}
type Volumes interface {
	AttachDisk(diskName KubernetesVolumeID, nodeName types.NodeName) (string, error)
	DetachDisk(diskName KubernetesVolumeID, nodeName types.NodeName) (string, error)
	CreateDisk(volumeOptions *VolumeOptions) (volumeName KubernetesVolumeID, err error)
	DeleteDisk(volumeName KubernetesVolumeID) (bool, error)
	GetVolumeLabels(volumeName KubernetesVolumeID) (map[string]string, error)
	GetDiskPath(volumeName KubernetesVolumeID) (string, error)
	DiskIsAttached(diskName KubernetesVolumeID, nodeName types.NodeName) (bool, error)
	DisksAreAttached(map[types.NodeName][]KubernetesVolumeID) (map[types.NodeName]map[KubernetesVolumeID]bool, error)
	ResizeDisk(diskName KubernetesVolumeID, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error)
}
type InstanceGroups interface {
	ResizeInstanceGroup(instanceGroupName string, size int) error
	DescribeInstanceGroup(instanceGroupName string) (InstanceGroupInfo, error)
}
type InstanceGroupInfo interface{ CurrentSize() (int, error) }
type Cloud struct {
	ec2              EC2
	elb              ELB
	elbv2            ELBV2
	asg              ASG
	kms              KMS
	metadata         EC2Metadata
	cfg              *CloudConfig
	region           string
	vpcID            string
	tagging          awsTagging
	selfAWSInstance  *awsInstance
	instanceCache    instanceCache
	clientBuilder    controller.ControllerClientBuilder
	kubeClient       clientset.Interface
	eventBroadcaster record.EventBroadcaster
	eventRecorder    record.EventRecorder
	attachingMutex   sync.Mutex
	attaching        map[types.NodeName]map[mountDevice]EBSVolumeID
	deviceAllocators map[types.NodeName]DeviceAllocator
}

var _ Volumes = &Cloud{}

type CloudConfig struct {
	Global struct {
		Zone                        string
		VPC                         string
		SubnetID                    string
		RouteTableID                string
		RoleARN                     string
		KubernetesClusterTag        string
		KubernetesClusterID         string
		DisableSecurityGroupIngress bool
		ElbSecurityGroup            string
		DisableStrictZoneCheck      bool
	}
	ServiceOverride map[string]*struct {
		Service       string
		Region        string
		URL           string
		SigningRegion string
		SigningMethod string
		SigningName   string
	}
}

func (cfg *CloudConfig) validateOverrides() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(cfg.ServiceOverride) == 0 {
		return nil
	}
	set := make(map[string]bool)
	for onum, ovrd := range cfg.ServiceOverride {
		name := strings.TrimSpace(ovrd.Service)
		if name == "" {
			return fmt.Errorf("service name is missing [Service is \"\"] in override %s", onum)
		}
		ovrd.Service = name
		region := strings.TrimSpace(ovrd.Region)
		if region == "" {
			return fmt.Errorf("service region is missing [Region is \"\"] in override %s", onum)
		}
		ovrd.Region = region
		url := strings.TrimSpace(ovrd.URL)
		if url == "" {
			return fmt.Errorf("url is missing [URL is \"\"] in override %s", onum)
		}
		signingRegion := strings.TrimSpace(ovrd.SigningRegion)
		if signingRegion == "" {
			return fmt.Errorf("signingRegion is missing [SigningRegion is \"\"] in override %s", onum)
		}
		signature := name + "_" + region
		if set[signature] {
			return fmt.Errorf("duplicate entry found for service override [%s] (%s in %s)", onum, name, region)
		}
		set[signature] = true
	}
	return nil
}
func (cfg *CloudConfig) getResolver() endpoints.ResolverFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultResolver := endpoints.DefaultResolver()
	defaultResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		return defaultResolver.EndpointFor(service, region, optFns...)
	}
	if len(cfg.ServiceOverride) == 0 {
		return defaultResolverFn
	}
	return func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		for _, override := range cfg.ServiceOverride {
			if override.Service == service && override.Region == region {
				return endpoints.ResolvedEndpoint{URL: override.URL, SigningRegion: override.SigningRegion, SigningMethod: override.SigningMethod, SigningName: override.SigningName}, nil
			}
		}
		return defaultResolver.EndpointFor(service, region, optFns...)
	}
}

type awsSdkEC2 struct{ ec2 *ec2.EC2 }
type awsCloudConfigProvider interface{ getResolver() endpoints.ResolverFunc }
type awsSDKProvider struct {
	creds          *credentials.Credentials
	cfg            awsCloudConfigProvider
	mutex          sync.Mutex
	regionDelayers map[string]*CrossRequestRetryDelay
}

func newAWSSDKProvider(creds *credentials.Credentials, cfg *CloudConfig) *awsSDKProvider {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &awsSDKProvider{creds: creds, cfg: cfg, regionDelayers: make(map[string]*CrossRequestRetryDelay)}
}
func (p *awsSDKProvider) addHandlers(regionName string, h *request.Handlers) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h.Sign.PushFrontNamed(request.NamedHandler{Name: "k8s/logger", Fn: awsHandlerLogger})
	delayer := p.getCrossRequestRetryDelay(regionName)
	if delayer != nil {
		h.Sign.PushFrontNamed(request.NamedHandler{Name: "k8s/delay-presign", Fn: delayer.BeforeSign})
		h.AfterRetry.PushFrontNamed(request.NamedHandler{Name: "k8s/delay-afterretry", Fn: delayer.AfterRetry})
	}
	p.addAPILoggingHandlers(h)
}
func (p *awsSDKProvider) addAPILoggingHandlers(h *request.Handlers) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h.Send.PushBackNamed(request.NamedHandler{Name: "k8s/api-request", Fn: awsSendHandlerLogger})
	h.ValidateResponse.PushFrontNamed(request.NamedHandler{Name: "k8s/api-validate-response", Fn: awsValidateResponseHandlerLogger})
}
func (p *awsSDKProvider) getCrossRequestRetryDelay(regionName string) *CrossRequestRetryDelay {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delayer, found := p.regionDelayers[regionName]
	if !found {
		delayer = NewCrossRequestRetryDelay()
		p.regionDelayers[regionName] = delayer
	}
	return delayer
}
func (p *awsSDKProvider) Compute(regionName string) (EC2, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsConfig := &aws.Config{Region: &regionName, Credentials: p.creds}
	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true).WithEndpointResolver(p.cfg.getResolver())
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	service := ec2.New(sess)
	p.addHandlers(regionName, &service.Handlers)
	ec2 := &awsSdkEC2{ec2: service}
	return ec2, nil
}
func (p *awsSDKProvider) LoadBalancing(regionName string) (ELB, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsConfig := &aws.Config{Region: &regionName, Credentials: p.creds}
	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true).WithEndpointResolver(p.cfg.getResolver())
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	elbClient := elb.New(sess)
	p.addHandlers(regionName, &elbClient.Handlers)
	return elbClient, nil
}
func (p *awsSDKProvider) LoadBalancingV2(regionName string) (ELBV2, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsConfig := &aws.Config{Region: &regionName, Credentials: p.creds}
	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true).WithEndpointResolver(p.cfg.getResolver())
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	elbClient := elbv2.New(sess)
	p.addHandlers(regionName, &elbClient.Handlers)
	return elbClient, nil
}
func (p *awsSDKProvider) Autoscaling(regionName string) (ASG, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsConfig := &aws.Config{Region: &regionName, Credentials: p.creds}
	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true).WithEndpointResolver(p.cfg.getResolver())
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	client := autoscaling.New(sess)
	p.addHandlers(regionName, &client.Handlers)
	return client, nil
}
func (p *awsSDKProvider) Metadata() (EC2Metadata, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sess, err := session.NewSession(&aws.Config{EndpointResolver: p.cfg.getResolver()})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	client := ec2metadata.New(sess)
	p.addAPILoggingHandlers(&client.Handlers)
	return client, nil
}
func (p *awsSDKProvider) KeyManagement(regionName string) (KMS, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsConfig := &aws.Config{Region: &regionName, Credentials: p.creds}
	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true).WithEndpointResolver(p.cfg.getResolver())
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
	}
	kmsClient := kms.New(sess)
	p.addHandlers(regionName, &kmsClient.Handlers)
	return kmsClient, nil
}
func newEc2Filter(name string, values ...string) *ec2.Filter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filter := &ec2.Filter{Name: aws.String(name)}
	for _, value := range values {
		filter.Values = append(filter.Values, aws.String(value))
	}
	return filter
}
func (c *Cloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cloudprovider.NotImplemented
}
func (c *Cloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.selfAWSInstance.nodeName, nil
}
func (s *awsSdkEC2) DescribeInstances(request *ec2.DescribeInstancesInput) ([]*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	results := []*ec2.Instance{}
	var nextToken *string
	requestTime := time.Now()
	for {
		response, err := s.ec2.DescribeInstances(request)
		if err != nil {
			recordAWSMetric("describe_instance", 0, err)
			return nil, fmt.Errorf("error listing AWS instances: %q", err)
		}
		for _, reservation := range response.Reservations {
			results = append(results, reservation.Instances...)
		}
		nextToken = response.NextToken
		if aws.StringValue(nextToken) == "" {
			break
		}
		request.NextToken = nextToken
	}
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("describe_instance", timeTaken, nil)
	return results, nil
}
func (s *awsSdkEC2) DescribeSecurityGroups(request *ec2.DescribeSecurityGroupsInput) ([]*ec2.SecurityGroup, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	response, err := s.ec2.DescribeSecurityGroups(request)
	if err != nil {
		return nil, fmt.Errorf("error listing AWS security groups: %q", err)
	}
	return response.SecurityGroups, nil
}
func (s *awsSdkEC2) AttachVolume(request *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.AttachVolume(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("attach_volume", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) DetachVolume(request *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.DetachVolume(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("detach_volume", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) DescribeVolumes(request *ec2.DescribeVolumesInput) ([]*ec2.Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	results := []*ec2.Volume{}
	var nextToken *string
	requestTime := time.Now()
	for {
		response, err := s.ec2.DescribeVolumes(request)
		if err != nil {
			recordAWSMetric("describe_volume", 0, err)
			return nil, err
		}
		results = append(results, response.Volumes...)
		nextToken = response.NextToken
		if aws.StringValue(nextToken) == "" {
			break
		}
		request.NextToken = nextToken
	}
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("describe_volume", timeTaken, nil)
	return results, nil
}
func (s *awsSdkEC2) CreateVolume(request *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.CreateVolume(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("create_volume", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) DeleteVolume(request *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.DeleteVolume(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("delete_volume", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) ModifyVolume(request *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.ModifyVolume(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("modify_volume", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) DescribeVolumeModifications(request *ec2.DescribeVolumesModificationsInput) ([]*ec2.VolumeModification, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	results := []*ec2.VolumeModification{}
	var nextToken *string
	for {
		resp, err := s.ec2.DescribeVolumesModifications(request)
		if err != nil {
			recordAWSMetric("describe_volume_modification", 0, err)
			return nil, fmt.Errorf("error listing volume modifictions : %v", err)
		}
		results = append(results, resp.VolumesModifications...)
		nextToken = resp.NextToken
		if aws.StringValue(nextToken) == "" {
			break
		}
		request.NextToken = nextToken
	}
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("describe_volume_modification", timeTaken, nil)
	return results, nil
}
func (s *awsSdkEC2) DescribeSubnets(request *ec2.DescribeSubnetsInput) ([]*ec2.Subnet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	response, err := s.ec2.DescribeSubnets(request)
	if err != nil {
		return nil, fmt.Errorf("error listing AWS subnets: %q", err)
	}
	return response.Subnets, nil
}
func (s *awsSdkEC2) CreateSecurityGroup(request *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.CreateSecurityGroup(request)
}
func (s *awsSdkEC2) DeleteSecurityGroup(request *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.DeleteSecurityGroup(request)
}
func (s *awsSdkEC2) AuthorizeSecurityGroupIngress(request *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.AuthorizeSecurityGroupIngress(request)
}
func (s *awsSdkEC2) RevokeSecurityGroupIngress(request *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.RevokeSecurityGroupIngress(request)
}
func (s *awsSdkEC2) CreateTags(request *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestTime := time.Now()
	resp, err := s.ec2.CreateTags(request)
	timeTaken := time.Since(requestTime).Seconds()
	recordAWSMetric("create_tags", timeTaken, err)
	return resp, err
}
func (s *awsSdkEC2) DescribeRouteTables(request *ec2.DescribeRouteTablesInput) ([]*ec2.RouteTable, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	response, err := s.ec2.DescribeRouteTables(request)
	if err != nil {
		return nil, fmt.Errorf("error listing AWS route tables: %q", err)
	}
	return response.RouteTables, nil
}
func (s *awsSdkEC2) CreateRoute(request *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.CreateRoute(request)
}
func (s *awsSdkEC2) DeleteRoute(request *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.DeleteRoute(request)
}
func (s *awsSdkEC2) ModifyInstanceAttribute(request *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.ModifyInstanceAttribute(request)
}
func (s *awsSdkEC2) DescribeVpcs(request *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ec2.DescribeVpcs(request)
}
func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerMetrics()
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := readAWSCloudConfig(config)
		if err != nil {
			return nil, fmt.Errorf("unable to read AWS cloud provider config file: %v", err)
		}
		if err = cfg.validateOverrides(); err != nil {
			return nil, fmt.Errorf("unable to validate custom endpoint overrides: %v", err)
		}
		sess, err := session.NewSession(&aws.Config{})
		if err != nil {
			return nil, fmt.Errorf("unable to initialize AWS session: %v", err)
		}
		var provider credentials.Provider
		if cfg.Global.RoleARN == "" {
			provider = &ec2rolecreds.EC2RoleProvider{Client: ec2metadata.New(sess)}
		} else {
			klog.Infof("Using AWS assumed role %v", cfg.Global.RoleARN)
			provider = &stscreds.AssumeRoleProvider{Client: sts.New(sess), RoleARN: cfg.Global.RoleARN}
		}
		creds := credentials.NewChainCredentials([]credentials.Provider{&credentials.EnvProvider{}, provider, &credentials.SharedCredentialsProvider{}})
		aws := newAWSSDKProvider(creds, cfg)
		return newAWSCloud(*cfg, aws)
	})
}
func readAWSCloudConfig(config io.Reader) (*CloudConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cfg CloudConfig
	var err error
	if config != nil {
		err = gcfg.ReadInto(&cfg, config)
		if err != nil {
			return nil, err
		}
	}
	return &cfg, nil
}
func updateConfigZone(cfg *CloudConfig, metadata EC2Metadata) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.Global.Zone == "" {
		if metadata != nil {
			klog.Info("Zone not specified in configuration file; querying AWS metadata service")
			var err error
			cfg.Global.Zone, err = getAvailabilityZone(metadata)
			if err != nil {
				return err
			}
		}
		if cfg.Global.Zone == "" {
			return fmt.Errorf("no zone specified in configuration file")
		}
	}
	return nil
}
func getAvailabilityZone(metadata EC2Metadata) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return metadata.GetMetadata("placement/availability-zone")
}
func azToRegion(az string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(az) < 1 {
		return "", fmt.Errorf("invalid (empty) AZ")
	}
	region := az[:len(az)-1]
	return region, nil
}
func newAWSCloud(cfg CloudConfig, awsServices Services) (*Cloud, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Building AWS cloudprovider")
	metadata, err := awsServices.Metadata()
	if err != nil {
		return nil, fmt.Errorf("error creating AWS metadata client: %q", err)
	}
	err = updateConfigZone(&cfg, metadata)
	if err != nil {
		return nil, fmt.Errorf("unable to determine AWS zone from cloud provider config or EC2 instance metadata: %v", err)
	}
	zone := cfg.Global.Zone
	if len(zone) <= 1 {
		return nil, fmt.Errorf("invalid AWS zone in config file: %s", zone)
	}
	regionName, err := azToRegion(zone)
	if err != nil {
		return nil, err
	}
	recognizeRegion(regionName)
	if !cfg.Global.DisableStrictZoneCheck {
		valid := isRegionValid(regionName)
		if !valid {
			return nil, fmt.Errorf("not a valid AWS zone (unknown region): %s", zone)
		}
	} else {
		klog.Warningf("Strict AWS zone checking is disabled.  Proceeding with zone: %s", zone)
	}
	ec2, err := awsServices.Compute(regionName)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS EC2 client: %v", err)
	}
	elb, err := awsServices.LoadBalancing(regionName)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS ELB client: %v", err)
	}
	elbv2, err := awsServices.LoadBalancingV2(regionName)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS ELBV2 client: %v", err)
	}
	asg, err := awsServices.Autoscaling(regionName)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS autoscaling client: %v", err)
	}
	kms, err := awsServices.KeyManagement(regionName)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS key management client: %v", err)
	}
	awsCloud := &Cloud{ec2: ec2, elb: elb, elbv2: elbv2, asg: asg, metadata: metadata, kms: kms, cfg: &cfg, region: regionName, attaching: make(map[types.NodeName]map[mountDevice]EBSVolumeID), deviceAllocators: make(map[types.NodeName]DeviceAllocator)}
	awsCloud.instanceCache.cloud = awsCloud
	tagged := cfg.Global.KubernetesClusterTag != "" || cfg.Global.KubernetesClusterID != ""
	if cfg.Global.VPC != "" && (cfg.Global.SubnetID != "" || cfg.Global.RoleARN != "") && tagged {
		klog.Info("Master is configured to run on a different AWS account, different cloud provider or on-premises")
		awsCloud.selfAWSInstance = &awsInstance{nodeName: "master-dummy", vpcID: cfg.Global.VPC, subnetID: cfg.Global.SubnetID}
		awsCloud.vpcID = cfg.Global.VPC
	} else {
		selfAWSInstance, err := awsCloud.buildSelfAWSInstance()
		if err != nil {
			return nil, err
		}
		awsCloud.selfAWSInstance = selfAWSInstance
		awsCloud.vpcID = selfAWSInstance.vpcID
	}
	if cfg.Global.KubernetesClusterTag != "" || cfg.Global.KubernetesClusterID != "" {
		if err := awsCloud.tagging.init(cfg.Global.KubernetesClusterTag, cfg.Global.KubernetesClusterID); err != nil {
			return nil, err
		}
	} else {
		info, err := awsCloud.selfAWSInstance.describeInstance()
		if err != nil {
			return nil, err
		}
		if err := awsCloud.tagging.initFromTags(info.Tags); err != nil {
			return nil, err
		}
	}
	once.Do(func() {
		recognizeWellKnownRegions()
	})
	return awsCloud, nil
}
func (c *Cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.clientBuilder = clientBuilder
	c.kubeClient = clientBuilder.ClientOrDie("aws-cloud-provider")
	c.eventBroadcaster = record.NewBroadcaster()
	c.eventBroadcaster.StartLogging(klog.Infof)
	c.eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: c.kubeClient.CoreV1().Events("")})
	c.eventRecorder = c.eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "aws-cloud-provider"})
}
func (c *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (c *Cloud) ProviderName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ProviderName
}
func (c *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c, true
}
func (c *Cloud) Instances() (cloudprovider.Instances, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c, true
}
func (c *Cloud) Zones() (cloudprovider.Zones, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c, true
}
func (c *Cloud) Routes() (cloudprovider.Routes, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c, true
}
func (c *Cloud) HasClusterID() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(c.tagging.clusterID()) > 0
}
func (c *Cloud) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.selfAWSInstance.nodeName == name || len(name) == 0 {
		addresses := []v1.NodeAddress{}
		macs, err := c.metadata.GetMetadata("network/interfaces/macs/")
		if err != nil {
			return nil, fmt.Errorf("error querying AWS metadata for %q: %q", "network/interfaces/macs", err)
		}
		for _, macID := range strings.Split(macs, "\n") {
			if macID == "" {
				continue
			}
			macPath := path.Join("network/interfaces/macs/", macID, "local-ipv4s")
			internalIPs, err := c.metadata.GetMetadata(macPath)
			if err != nil {
				return nil, fmt.Errorf("error querying AWS metadata for %q: %q", macPath, err)
			}
			for _, internalIP := range strings.Split(internalIPs, "\n") {
				if internalIP == "" {
					continue
				}
				addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalIP, Address: internalIP})
			}
		}
		externalIP, err := c.metadata.GetMetadata("public-ipv4")
		if err != nil {
			klog.V(4).Info("Could not determine public IP from AWS metadata.")
		} else {
			addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: externalIP})
		}
		internalDNS, err := c.metadata.GetMetadata("local-hostname")
		if err != nil || len(internalDNS) == 0 {
			klog.V(4).Info("Could not determine private DNS from AWS metadata.")
		} else {
			addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalDNS, Address: internalDNS})
			addresses = append(addresses, v1.NodeAddress{Type: v1.NodeHostName, Address: internalDNS})
		}
		externalDNS, err := c.metadata.GetMetadata("public-hostname")
		if err != nil || len(externalDNS) == 0 {
			klog.V(4).Info("Could not determine public DNS from AWS metadata.")
		} else {
			addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalDNS, Address: externalDNS})
		}
		return addresses, nil
	}
	instance, err := c.getInstanceByNodeName(name)
	if err != nil {
		return nil, fmt.Errorf("getInstanceByNodeName failed for %q with %q", name, err)
	}
	return extractNodeAddresses(instance)
}
func extractNodeAddresses(instance *ec2.Instance) ([]v1.NodeAddress, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if instance == nil {
		return nil, fmt.Errorf("nil instance passed to extractNodeAddresses")
	}
	addresses := []v1.NodeAddress{}
	for _, networkInterface := range instance.NetworkInterfaces {
		if aws.StringValue(networkInterface.Status) != ec2.NetworkInterfaceStatusInUse {
			continue
		}
		for _, internalIP := range networkInterface.PrivateIpAddresses {
			if ipAddress := aws.StringValue(internalIP.PrivateIpAddress); ipAddress != "" {
				ip := net.ParseIP(ipAddress)
				if ip == nil {
					return nil, fmt.Errorf("EC2 instance had invalid private address: %s (%q)", aws.StringValue(instance.InstanceId), ipAddress)
				}
				addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalIP, Address: ip.String()})
			}
		}
	}
	publicIPAddress := aws.StringValue(instance.PublicIpAddress)
	if publicIPAddress != "" {
		ip := net.ParseIP(publicIPAddress)
		if ip == nil {
			return nil, fmt.Errorf("EC2 instance had invalid public address: %s (%s)", aws.StringValue(instance.InstanceId), publicIPAddress)
		}
		addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: ip.String()})
	}
	privateDNSName := aws.StringValue(instance.PrivateDnsName)
	if privateDNSName != "" {
		addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalDNS, Address: privateDNSName})
		addresses = append(addresses, v1.NodeAddress{Type: v1.NodeHostName, Address: privateDNSName})
	}
	publicDNSName := aws.StringValue(instance.PublicDnsName)
	if publicDNSName != "" {
		addresses = append(addresses, v1.NodeAddress{Type: v1.NodeExternalDNS, Address: publicDNSName})
	}
	return addresses, nil
}
func (c *Cloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID, err := kubernetesInstanceID(providerID).mapToAWSInstanceID()
	if err != nil {
		return nil, err
	}
	instance, err := describeInstance(c.ec2, instanceID)
	if err != nil {
		return nil, err
	}
	return extractNodeAddresses(instance)
}
func (c *Cloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID, err := kubernetesInstanceID(providerID).mapToAWSInstanceID()
	if err != nil {
		return false, err
	}
	request := &ec2.DescribeInstancesInput{InstanceIds: []*string{instanceID.awsString()}}
	instances, err := c.ec2.DescribeInstances(request)
	if err != nil {
		return false, err
	}
	if len(instances) == 0 {
		return false, nil
	}
	if len(instances) > 1 {
		return false, fmt.Errorf("multiple instances found for instance: %s", instanceID)
	}
	state := instances[0].State.Name
	if *state == ec2.InstanceStateNameTerminated {
		klog.Warningf("the instance %s is terminated", instanceID)
		return false, nil
	}
	return true, nil
}
func (c *Cloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID, err := kubernetesInstanceID(providerID).mapToAWSInstanceID()
	if err != nil {
		return false, err
	}
	request := &ec2.DescribeInstancesInput{InstanceIds: []*string{instanceID.awsString()}}
	instances, err := c.ec2.DescribeInstances(request)
	if err != nil {
		return false, err
	}
	if len(instances) == 0 {
		klog.Warningf("the instance %s does not exist anymore", providerID)
		return false, nil
	}
	if len(instances) > 1 {
		return false, fmt.Errorf("multiple instances found for instance: %s", instanceID)
	}
	instance := instances[0]
	if instance.State != nil {
		state := aws.StringValue(instance.State.Name)
		if state == ec2.InstanceStateNameStopped {
			return true, nil
		}
	}
	return false, nil
}
func (c *Cloud) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.selfAWSInstance.nodeName == nodeName {
		return "/" + c.selfAWSInstance.availabilityZone + "/" + c.selfAWSInstance.awsID, nil
	}
	inst, err := c.getInstanceByNodeName(nodeName)
	if err != nil {
		if err == cloudprovider.InstanceNotFound {
			return "", err
		}
		return "", fmt.Errorf("getInstanceByNodeName failed for %q with %q", nodeName, err)
	}
	return "/" + aws.StringValue(inst.Placement.AvailabilityZone) + "/" + aws.StringValue(inst.InstanceId), nil
}
func (c *Cloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID, err := kubernetesInstanceID(providerID).mapToAWSInstanceID()
	if err != nil {
		return "", err
	}
	instance, err := describeInstance(c.ec2, instanceID)
	if err != nil {
		return "", err
	}
	return aws.StringValue(instance.InstanceType), nil
}
func (c *Cloud) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.selfAWSInstance.nodeName == nodeName {
		return c.selfAWSInstance.instanceType, nil
	}
	inst, err := c.getInstanceByNodeName(nodeName)
	if err != nil {
		return "", fmt.Errorf("getInstanceByNodeName failed for %q with %q", nodeName, err)
	}
	return aws.StringValue(inst.InstanceType), nil
}
func (c *Cloud) GetCandidateZonesForDynamicVolume() (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filters := []*ec2.Filter{newEc2Filter("instance-state-name", "running")}
	instances, err := c.describeInstances(filters)
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instances returned")
	}
	zones := sets.NewString()
	for _, instance := range instances {
		master := false
		for _, tag := range instance.Tags {
			tagKey := aws.StringValue(tag.Key)
			if awsTagNameMasterRoles.Has(tagKey) {
				master = true
			}
		}
		if master {
			klog.V(4).Infof("Ignoring master instance %q in zone discovery", aws.StringValue(instance.InstanceId))
			continue
		}
		if instance.Placement != nil {
			zone := aws.StringValue(instance.Placement.AvailabilityZone)
			zones.Insert(zone)
		}
	}
	klog.V(2).Infof("Found instances in zones %s", zones)
	return zones, nil
}
func (c *Cloud) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cloudprovider.Zone{FailureDomain: c.selfAWSInstance.availabilityZone, Region: c.region}, nil
}
func (c *Cloud) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID, err := kubernetesInstanceID(providerID).mapToAWSInstanceID()
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	instance, err := c.getInstanceByID(string(instanceID))
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	zone := cloudprovider.Zone{FailureDomain: *(instance.Placement.AvailabilityZone), Region: c.region}
	return zone, nil
}
func (c *Cloud) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instance, err := c.getInstanceByNodeName(nodeName)
	if err != nil {
		return cloudprovider.Zone{}, err
	}
	zone := cloudprovider.Zone{FailureDomain: *(instance.Placement.AvailabilityZone), Region: c.region}
	return zone, nil
}

type mountDevice string
type awsInstance struct {
	ec2              EC2
	awsID            string
	nodeName         types.NodeName
	availabilityZone string
	vpcID            string
	subnetID         string
	instanceType     string
}

func newAWSInstance(ec2Service EC2, instance *ec2.Instance) *awsInstance {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	az := ""
	if instance.Placement != nil {
		az = aws.StringValue(instance.Placement.AvailabilityZone)
	}
	self := &awsInstance{ec2: ec2Service, awsID: aws.StringValue(instance.InstanceId), nodeName: mapInstanceToNodeName(instance), availabilityZone: az, instanceType: aws.StringValue(instance.InstanceType), vpcID: aws.StringValue(instance.VpcId), subnetID: aws.StringValue(instance.SubnetId)}
	return self
}
func (i *awsInstance) describeInstance() (*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return describeInstance(i.ec2, awsInstanceID(i.awsID))
}
func (c *Cloud) getMountDevice(i *awsInstance, info *ec2.Instance, volumeID EBSVolumeID, assign bool) (assigned mountDevice, alreadyAttached bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	deviceMappings := map[mountDevice]EBSVolumeID{}
	for _, blockDevice := range info.BlockDeviceMappings {
		name := aws.StringValue(blockDevice.DeviceName)
		if strings.HasPrefix(name, "/dev/sd") {
			name = name[7:]
		}
		if strings.HasPrefix(name, "/dev/xvd") {
			name = name[8:]
		}
		if len(name) < 1 || len(name) > 2 {
			klog.Warningf("Unexpected EBS DeviceName: %q", aws.StringValue(blockDevice.DeviceName))
		}
		deviceMappings[mountDevice(name)] = EBSVolumeID(aws.StringValue(blockDevice.Ebs.VolumeId))
	}
	klog.V(2).Infof("volume ID: %s: device mappings from EC2 Instance: %+v", volumeID, deviceMappings)
	c.attachingMutex.Lock()
	defer c.attachingMutex.Unlock()
	for mountDevice, volume := range c.attaching[i.nodeName] {
		deviceMappings[mountDevice] = volume
	}
	klog.V(2).Infof("volume ID: %s: Full device mappings: %+v", volumeID, deviceMappings)
	for mountDevice, mappingVolumeID := range deviceMappings {
		if volumeID == mappingVolumeID {
			if assign {
				klog.Warningf("Got assignment call for already-assigned volume: %s@%s", mountDevice, mappingVolumeID)
			}
			return mountDevice, true, nil
		}
	}
	if !assign {
		return mountDevice(""), false, nil
	}
	deviceAllocator := c.deviceAllocators[i.nodeName]
	if deviceAllocator == nil {
		deviceAllocator = NewDeviceAllocator()
		c.deviceAllocators[i.nodeName] = deviceAllocator
	}
	deviceAllocator.Lock()
	defer deviceAllocator.Unlock()
	chosen, err := deviceAllocator.GetNext(deviceMappings)
	if err != nil {
		klog.Warningf("Could not assign a mount device.  mappings=%v, error: %v", deviceMappings, err)
		return "", false, fmt.Errorf("too many EBS volumes attached to node %s", i.nodeName)
	}
	attaching := c.attaching[i.nodeName]
	if attaching == nil {
		attaching = make(map[mountDevice]EBSVolumeID)
		c.attaching[i.nodeName] = attaching
	}
	attaching[chosen] = volumeID
	klog.V(2).Infof("Assigned mount device %s -> volume %s", chosen, volumeID)
	return chosen, false, nil
}
func (c *Cloud) endAttaching(i *awsInstance, volumeID EBSVolumeID, mountDevice mountDevice) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.attachingMutex.Lock()
	defer c.attachingMutex.Unlock()
	existingVolumeID, found := c.attaching[i.nodeName][mountDevice]
	if !found {
		return false
	}
	if volumeID != existingVolumeID {
		klog.Infof("endAttaching on device %q assigned to different volume: %q vs %q", mountDevice, volumeID, existingVolumeID)
		return false
	}
	klog.V(2).Infof("Releasing in-process attachment entry: %s -> volume %s", mountDevice, volumeID)
	delete(c.attaching[i.nodeName], mountDevice)
	return true
}

type awsDisk struct {
	ec2   EC2
	name  KubernetesVolumeID
	awsID EBSVolumeID
}

func newAWSDisk(aws *Cloud, name KubernetesVolumeID) (*awsDisk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsID, err := name.MapToAWSVolumeID()
	if err != nil {
		return nil, err
	}
	disk := &awsDisk{ec2: aws.ec2, name: name, awsID: awsID}
	return disk, nil
}
func isAWSErrorVolumeNotFound(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			if awsError.Code() == "InvalidVolume.NotFound" {
				return true
			}
		}
	}
	return false
}
func (d *awsDisk) describeVolume() (*ec2.Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeID := d.awsID
	request := &ec2.DescribeVolumesInput{VolumeIds: []*string{volumeID.awsString()}}
	volumes, err := d.ec2.DescribeVolumes(request)
	if err != nil {
		return nil, err
	}
	if len(volumes) == 0 {
		return nil, fmt.Errorf("no volumes found")
	}
	if len(volumes) > 1 {
		return nil, fmt.Errorf("multiple volumes found")
	}
	return volumes[0], nil
}
func (d *awsDisk) describeVolumeModification() (*ec2.VolumeModification, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeID := d.awsID
	request := &ec2.DescribeVolumesModificationsInput{VolumeIds: []*string{volumeID.awsString()}}
	volumeMods, err := d.ec2.DescribeVolumeModifications(request)
	if err != nil {
		return nil, fmt.Errorf("error describing volume modification %s with %v", volumeID, err)
	}
	if len(volumeMods) == 0 {
		return nil, fmt.Errorf("no volume modifications found for %s", volumeID)
	}
	lastIndex := len(volumeMods) - 1
	return volumeMods[lastIndex], nil
}
func (d *awsDisk) modifyVolume(requestGiB int64) (int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeID := d.awsID
	request := &ec2.ModifyVolumeInput{VolumeId: volumeID.awsString(), Size: aws.Int64(requestGiB)}
	output, err := d.ec2.ModifyVolume(request)
	if err != nil {
		modifyError := fmt.Errorf("AWS modifyVolume failed for %s with %v", volumeID, err)
		return requestGiB, modifyError
	}
	volumeModification := output.VolumeModification
	if aws.StringValue(volumeModification.ModificationState) == ec2.VolumeModificationStateCompleted {
		return aws.Int64Value(volumeModification.TargetSize), nil
	}
	backoff := wait.Backoff{Duration: 1 * time.Second, Factor: 2, Steps: 10}
	checkForResize := func() (bool, error) {
		volumeModification, err := d.describeVolumeModification()
		if err != nil {
			return false, err
		}
		if aws.StringValue(volumeModification.ModificationState) == ec2.VolumeModificationStateOptimizing {
			return true, nil
		}
		return false, nil
	}
	waitWithErr := wait.ExponentialBackoff(backoff, checkForResize)
	return requestGiB, waitWithErr
}
func (c *Cloud) applyUnSchedulableTaint(nodeName types.NodeName, reason string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node, fetchErr := c.kubeClient.CoreV1().Nodes().Get(string(nodeName), metav1.GetOptions{})
	if fetchErr != nil {
		klog.Errorf("Error fetching node %s with %v", nodeName, fetchErr)
		return
	}
	taint := &v1.Taint{Key: nodeWithImpairedVolumes, Value: "true", Effect: v1.TaintEffectNoSchedule}
	err := controller.AddOrUpdateTaintOnNode(c.kubeClient, string(nodeName), taint)
	if err != nil {
		klog.Errorf("Error applying taint to node %s with error %v", nodeName, err)
		return
	}
	c.eventRecorder.Eventf(node, v1.EventTypeWarning, volumeAttachmentStuck, reason)
}
func (d *awsDisk) waitForAttachmentStatus(status string) (*ec2.VolumeAttachment, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backoff := wait.Backoff{Duration: volumeAttachmentStatusInitialDelay, Factor: volumeAttachmentStatusFactor, Steps: volumeAttachmentStatusSteps}
	describeErrorCount := 0
	var attachment *ec2.VolumeAttachment
	err := wait.ExponentialBackoff(backoff, func() (bool, error) {
		info, err := d.describeVolume()
		if err != nil {
			if isAWSErrorVolumeNotFound(err) {
				if status == "detached" {
					klog.Warningf("Waiting for volume %q to be detached but the volume does not exist", d.awsID)
					stateStr := "detached"
					attachment = &ec2.VolumeAttachment{State: &stateStr}
					return true, nil
				}
				if status == "attached" {
					klog.Warningf("Waiting for volume %q to be attached but the volume does not exist", d.awsID)
					return false, err
				}
			}
			describeErrorCount++
			if describeErrorCount > volumeAttachmentStatusConsecutiveErrorLimit {
				return false, err
			}
			klog.Warningf("Ignoring error from describe volume for volume %q; will retry: %q", d.awsID, err)
			return false, nil
		}
		describeErrorCount = 0
		if len(info.Attachments) > 1 {
			klog.Warningf("Found multiple attachments for volume %q: %v", d.awsID, info)
		}
		attachmentStatus := ""
		for _, a := range info.Attachments {
			if attachmentStatus != "" {
				klog.Warningf("Found multiple attachments for volume %q: %v", d.awsID, info)
			}
			if a.State != nil {
				attachment = a
				attachmentStatus = *a.State
			} else {
				klog.Warningf("Ignoring nil attachment state for volume %q: %v", d.awsID, a)
			}
		}
		if attachmentStatus == "" {
			attachmentStatus = "detached"
		}
		if attachmentStatus == status {
			return true, nil
		}
		klog.V(2).Infof("Waiting for volume %q state: actual=%s, desired=%s", d.awsID, attachmentStatus, status)
		return false, nil
	})
	return attachment, err
}
func (d *awsDisk) deleteVolume() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &ec2.DeleteVolumeInput{VolumeId: d.awsID.awsString()}
	_, err := d.ec2.DeleteVolume(request)
	if err != nil {
		if isAWSErrorVolumeNotFound(err) {
			return false, nil
		}
		if awsError, ok := err.(awserr.Error); ok {
			if awsError.Code() == "VolumeInUse" {
				return false, volume.NewDeletedVolumeInUseError(err.Error())
			}
		}
		return false, fmt.Errorf("error deleting EBS volume %q: %q", d.awsID, err)
	}
	return true, nil
}
func (c *Cloud) buildSelfAWSInstance() (*awsInstance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.selfAWSInstance != nil {
		panic("do not call buildSelfAWSInstance directly")
	}
	instanceID, err := c.metadata.GetMetadata("instance-id")
	if err != nil {
		return nil, fmt.Errorf("error fetching instance-id from ec2 metadata service: %q", err)
	}
	instance, err := c.getInstanceByID(instanceID)
	if err != nil {
		return nil, fmt.Errorf("error finding instance %s: %q", instanceID, err)
	}
	return newAWSInstance(c.ec2, instance), nil
}
func wrapAttachError(err error, disk *awsDisk, instance string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if awsError, ok := err.(awserr.Error); ok {
		if awsError.Code() == "VolumeInUse" {
			info, err := disk.describeVolume()
			if err != nil {
				klog.Errorf("Error describing volume %q: %q", disk.awsID, err)
			} else {
				for _, a := range info.Attachments {
					if disk.awsID != EBSVolumeID(aws.StringValue(a.VolumeId)) {
						klog.Warningf("Expected to get attachment info of volume %q but instead got info of %q", disk.awsID, aws.StringValue(a.VolumeId))
					} else if aws.StringValue(a.State) == "attached" {
						return fmt.Errorf("Error attaching EBS volume %q to instance %q: %q. The volume is currently attached to instance %q", disk.awsID, instance, awsError, aws.StringValue(a.InstanceId))
					}
				}
			}
		}
	}
	return fmt.Errorf("Error attaching EBS volume %q to instance %q: %q", disk.awsID, instance, err)
}
func (c *Cloud) AttachDisk(diskName KubernetesVolumeID, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := newAWSDisk(c, diskName)
	if err != nil {
		return "", err
	}
	awsInstance, info, err := c.getFullInstance(nodeName)
	if err != nil {
		return "", fmt.Errorf("error finding instance %s: %q", nodeName, err)
	}
	klog.V(2).Infof("volumeID: %s, AttachDisk got AWS instance %+v", disk.awsID, info.BlockDeviceMappings)
	var mountDevice mountDevice
	var alreadyAttached bool
	attachEnded := false
	defer func() {
		if attachEnded {
			if !c.endAttaching(awsInstance, disk.awsID, mountDevice) {
				klog.Errorf("endAttaching called for disk %q when attach not in progress", disk.awsID)
			}
		}
	}()
	mountDevice, alreadyAttached, err = c.getMountDevice(awsInstance, info, disk.awsID, true)
	if err != nil {
		return "", err
	}
	hostDevice := "/dev/xvd" + string(mountDevice)
	ec2Device := "/dev/xvd" + string(mountDevice)
	if !alreadyAttached {
		available, err := c.checkIfAvailable(disk, "attaching", awsInstance.awsID)
		if err != nil {
			klog.Error(err)
		}
		if !available {
			attachEnded = true
			return "", err
		}
		request := &ec2.AttachVolumeInput{Device: aws.String(ec2Device), InstanceId: aws.String(awsInstance.awsID), VolumeId: disk.awsID.awsString()}
		attachResponse, err := c.ec2.AttachVolume(request)
		if err != nil {
			attachEnded = true
			return "", wrapAttachError(err, disk, awsInstance.awsID)
		}
		if da, ok := c.deviceAllocators[awsInstance.nodeName]; ok {
			da.Deprioritize(mountDevice)
		}
		klog.V(2).Infof("AttachVolume volume=%q instance=%q request returned %v", disk.awsID, awsInstance.awsID, attachResponse)
	}
	attachment, err := disk.waitForAttachmentStatus("attached")
	if err != nil {
		if err == wait.ErrWaitTimeout {
			c.applyUnSchedulableTaint(nodeName, "Volume stuck in attaching state - node needs reboot to fix impaired state.")
		}
		return "", err
	}
	attachEnded = true
	if attachment == nil {
		return "", fmt.Errorf("unexpected state: attachment nil after attached %q to %q", diskName, nodeName)
	}
	if ec2Device != aws.StringValue(attachment.Device) {
		return "", fmt.Errorf("disk attachment of %q to %q failed: requested device %q but found %q", diskName, nodeName, ec2Device, aws.StringValue(attachment.Device))
	}
	if awsInstance.awsID != aws.StringValue(attachment.InstanceId) {
		return "", fmt.Errorf("disk attachment of %q to %q failed: requested instance %q but found %q", diskName, nodeName, awsInstance.awsID, aws.StringValue(attachment.InstanceId))
	}
	return hostDevice, nil
}
func (c *Cloud) DetachDisk(diskName KubernetesVolumeID, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	diskInfo, attached, err := c.checkIfAttachedToNode(diskName, nodeName)
	if err != nil {
		if isAWSErrorVolumeNotFound(err) {
			klog.Warningf("DetachDisk %s called for node %s but volume does not exist; assuming the volume is detached", diskName, nodeName)
			return "", nil
		}
		return "", err
	}
	if !attached && diskInfo.ec2Instance != nil {
		klog.Warningf("DetachDisk %s called for node %s but volume is attached to node %s", diskName, nodeName, diskInfo.nodeName)
		return "", nil
	}
	if !attached {
		return "", nil
	}
	awsInstance := newAWSInstance(c.ec2, diskInfo.ec2Instance)
	mountDevice, alreadyAttached, err := c.getMountDevice(awsInstance, diskInfo.ec2Instance, diskInfo.disk.awsID, false)
	if err != nil {
		return "", err
	}
	if !alreadyAttached {
		klog.Warningf("DetachDisk called on non-attached disk: %s", diskName)
	}
	request := ec2.DetachVolumeInput{InstanceId: &awsInstance.awsID, VolumeId: diskInfo.disk.awsID.awsString()}
	response, err := c.ec2.DetachVolume(&request)
	if err != nil {
		return "", fmt.Errorf("error detaching EBS volume %q from %q: %q", diskInfo.disk.awsID, awsInstance.awsID, err)
	}
	if response == nil {
		return "", errors.New("no response from DetachVolume")
	}
	attachment, err := diskInfo.disk.waitForAttachmentStatus("detached")
	if err != nil {
		return "", err
	}
	if da, ok := c.deviceAllocators[awsInstance.nodeName]; ok {
		da.Deprioritize(mountDevice)
	}
	if attachment != nil {
		klog.V(2).Infof("waitForAttachmentStatus returned non-nil attachment with state=detached: %v", attachment)
	}
	if mountDevice != "" {
		c.endAttaching(awsInstance, diskInfo.disk.awsID, mountDevice)
	}
	_, info, err2 := c.getFullInstance(nodeName)
	if err2 != nil {
		klog.V(2).Infof("volume ID: %s: error finding instance %s: %q", diskInfo.disk.awsID, nodeName, err2)
	}
	klog.V(2).Infof("volume ID: %s: DetachDisk: instance after detach: %+v", diskInfo.disk.awsID, info.BlockDeviceMappings)
	hostDevicePath := "/dev/xvd" + string(mountDevice)
	return hostDevicePath, err
}
func (c *Cloud) CreateDisk(volumeOptions *VolumeOptions) (KubernetesVolumeID, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var createType string
	var iops int64
	switch volumeOptions.VolumeType {
	case VolumeTypeGP2, VolumeTypeSC1, VolumeTypeST1:
		createType = volumeOptions.VolumeType
	case VolumeTypeIO1:
		createType = volumeOptions.VolumeType
		iops = int64(volumeOptions.CapacityGB * volumeOptions.IOPSPerGB)
		if iops < MinTotalIOPS {
			iops = MinTotalIOPS
		}
		if iops > MaxTotalIOPS {
			iops = MaxTotalIOPS
		}
	case "":
		createType = DefaultVolumeType
	default:
		return "", fmt.Errorf("invalid AWS VolumeType %q", volumeOptions.VolumeType)
	}
	request := &ec2.CreateVolumeInput{}
	request.AvailabilityZone = aws.String(volumeOptions.AvailabilityZone)
	request.Size = aws.Int64(int64(volumeOptions.CapacityGB))
	request.VolumeType = aws.String(createType)
	request.Encrypted = aws.Bool(volumeOptions.Encrypted)
	if len(volumeOptions.KmsKeyID) > 0 {
		request.KmsKeyId = aws.String(volumeOptions.KmsKeyID)
		request.Encrypted = aws.Bool(true)
	}
	if iops > 0 {
		request.Iops = aws.Int64(iops)
	}
	tags := volumeOptions.Tags
	tags = c.tagging.buildTags(ResourceLifecycleOwned, tags)
	var tagList []*ec2.Tag
	for k, v := range tags {
		tagList = append(tagList, &ec2.Tag{Key: aws.String(k), Value: aws.String(v)})
	}
	request.TagSpecifications = append(request.TagSpecifications, &ec2.TagSpecification{Tags: tagList, ResourceType: aws.String(ec2.ResourceTypeVolume)})
	response, err := c.ec2.CreateVolume(request)
	if err != nil {
		return "", err
	}
	awsID := EBSVolumeID(aws.StringValue(response.VolumeId))
	if awsID == "" {
		return "", fmt.Errorf("VolumeID was not returned by CreateVolume")
	}
	volumeName := KubernetesVolumeID("aws://" + aws.StringValue(response.AvailabilityZone) + "/" + string(awsID))
	if len(volumeOptions.KmsKeyID) > 0 {
		err := c.waitUntilVolumeAvailable(volumeName)
		if err != nil {
			if isAWSErrorVolumeNotFound(err) {
				err = fmt.Errorf("failed to create encrypted volume: the volume disappeared after creation, most likely due to inaccessible KMS encryption key")
			}
			return "", err
		}
	}
	return volumeName, nil
}
func (c *Cloud) waitUntilVolumeAvailable(volumeName KubernetesVolumeID) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := newAWSDisk(c, volumeName)
	if err != nil {
		return err
	}
	err = wait.Poll(encryptedCheckInterval, encryptedCheckTimeout, func() (done bool, err error) {
		vol, err := disk.describeVolume()
		if err != nil {
			return true, err
		}
		if vol.State != nil {
			switch *vol.State {
			case "available":
				return true, nil
			case "creating":
				return false, nil
			default:
				return true, fmt.Errorf("unexpected State of newly created AWS EBS volume %s: %q", volumeName, *vol.State)
			}
		}
		return false, nil
	})
	return err
}
func (c *Cloud) DeleteDisk(volumeName KubernetesVolumeID) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsDisk, err := newAWSDisk(c, volumeName)
	if err != nil {
		return false, err
	}
	available, err := c.checkIfAvailable(awsDisk, "deleting", "")
	if err != nil {
		if isAWSErrorVolumeNotFound(err) {
			klog.V(2).Infof("Volume %s not found when deleting it, assuming it's deleted", awsDisk.awsID)
			return false, nil
		}
		klog.Error(err)
	}
	if !available {
		return false, err
	}
	return awsDisk.deleteVolume()
}
func (c *Cloud) checkIfAvailable(disk *awsDisk, opName string, instance string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info, err := disk.describeVolume()
	if err != nil {
		klog.Errorf("Error describing volume %q: %q", disk.awsID, err)
		return false, err
	}
	volumeState := aws.StringValue(info.State)
	opError := fmt.Sprintf("Error %s EBS volume %q", opName, disk.awsID)
	if len(instance) != 0 {
		opError = fmt.Sprintf("%q to instance %q", opError, instance)
	}
	if volumeState != "available" {
		if len(info.Attachments) > 0 {
			attachment := info.Attachments[0]
			instanceID := aws.StringValue(attachment.InstanceId)
			attachedInstance, ierr := c.getInstanceByID(instanceID)
			attachErr := fmt.Sprintf("%s since volume is currently attached to %q", opError, instanceID)
			if ierr != nil {
				klog.Error(attachErr)
				return false, errors.New(attachErr)
			}
			devicePath := aws.StringValue(attachment.Device)
			nodeName := mapInstanceToNodeName(attachedInstance)
			danglingErr := volumeutil.NewDanglingError(attachErr, nodeName, devicePath)
			return false, danglingErr
		}
		attachErr := fmt.Errorf("%s since volume is in %q state", opError, volumeState)
		return false, attachErr
	}
	return true, nil
}
func (c *Cloud) GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.AWSElasticBlockStore.VolumeID == volume.ProvisionedVolumeName {
		return nil, nil
	}
	spec := KubernetesVolumeID(pv.Spec.AWSElasticBlockStore.VolumeID)
	labels, err := c.GetVolumeLabels(spec)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
func (c *Cloud) GetVolumeLabels(volumeName KubernetesVolumeID) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsDisk, err := newAWSDisk(c, volumeName)
	if err != nil {
		return nil, err
	}
	info, err := awsDisk.describeVolume()
	if err != nil {
		return nil, err
	}
	labels := make(map[string]string)
	az := aws.StringValue(info.AvailabilityZone)
	if az == "" {
		return nil, fmt.Errorf("volume did not have AZ information: %q", aws.StringValue(info.VolumeId))
	}
	labels[kubeletapis.LabelZoneFailureDomain] = az
	region, err := azToRegion(az)
	if err != nil {
		return nil, err
	}
	labels[kubeletapis.LabelZoneRegion] = region
	return labels, nil
}
func (c *Cloud) GetDiskPath(volumeName KubernetesVolumeID) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsDisk, err := newAWSDisk(c, volumeName)
	if err != nil {
		return "", err
	}
	info, err := awsDisk.describeVolume()
	if err != nil {
		return "", err
	}
	if len(info.Attachments) == 0 {
		return "", fmt.Errorf("No attachment to volume %s", volumeName)
	}
	return aws.StringValue(info.Attachments[0].Device), nil
}
func (c *Cloud) DiskIsAttached(diskName KubernetesVolumeID, nodeName types.NodeName) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, attached, err := c.checkIfAttachedToNode(diskName, nodeName)
	if err != nil {
		if isAWSErrorVolumeNotFound(err) {
			klog.Warningf("DiskIsAttached called for volume %s on node %s but the volume does not exist", diskName, nodeName)
			return false, nil
		}
		return true, err
	}
	return attached, nil
}
func (c *Cloud) DisksAreAttached(nodeDisks map[types.NodeName][]KubernetesVolumeID) (map[types.NodeName]map[KubernetesVolumeID]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attached := make(map[types.NodeName]map[KubernetesVolumeID]bool)
	if len(nodeDisks) == 0 {
		return attached, nil
	}
	nodeNames := []string{}
	for nodeName, diskNames := range nodeDisks {
		for _, diskName := range diskNames {
			setNodeDisk(attached, diskName, nodeName, false)
		}
		nodeNames = append(nodeNames, mapNodeNameToPrivateDNSName(nodeName))
	}
	awsInstances, err := c.getInstancesByNodeNames(nodeNames)
	if err != nil {
		return nil, err
	}
	if len(awsInstances) == 0 {
		klog.V(2).Infof("DisksAreAttached found no instances matching node names; will assume disks not attached")
		return attached, nil
	}
	for _, awsInstance := range awsInstances {
		nodeName := mapInstanceToNodeName(awsInstance)
		diskNames := nodeDisks[nodeName]
		if len(diskNames) == 0 {
			continue
		}
		awsInstanceState := "<nil>"
		if awsInstance != nil && awsInstance.State != nil {
			awsInstanceState = aws.StringValue(awsInstance.State.Name)
		}
		if awsInstanceState == "terminated" {
			continue
		}
		idToDiskName := make(map[EBSVolumeID]KubernetesVolumeID)
		for _, diskName := range diskNames {
			volumeID, err := diskName.MapToAWSVolumeID()
			if err != nil {
				return nil, fmt.Errorf("error mapping volume spec %q to aws id: %v", diskName, err)
			}
			idToDiskName[volumeID] = diskName
		}
		for _, blockDevice := range awsInstance.BlockDeviceMappings {
			volumeID := EBSVolumeID(aws.StringValue(blockDevice.Ebs.VolumeId))
			diskName, found := idToDiskName[volumeID]
			if found {
				setNodeDisk(attached, diskName, nodeName, true)
			}
		}
	}
	return attached, nil
}
func (c *Cloud) ResizeDisk(diskName KubernetesVolumeID, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	awsDisk, err := newAWSDisk(c, diskName)
	if err != nil {
		return oldSize, err
	}
	volumeInfo, err := awsDisk.describeVolume()
	if err != nil {
		descErr := fmt.Errorf("AWS.ResizeDisk Error describing volume %s with %v", diskName, err)
		return oldSize, descErr
	}
	requestGiB := volumeutil.RoundUpToGiB(newSize)
	newSizeQuant := resource.MustParse(fmt.Sprintf("%dGi", requestGiB))
	if aws.Int64Value(volumeInfo.Size) >= requestGiB {
		return newSizeQuant, nil
	}
	_, err = awsDisk.modifyVolume(requestGiB)
	if err != nil {
		return oldSize, err
	}
	return newSizeQuant, nil
}
func (c *Cloud) describeLoadBalancer(name string) (*elb.LoadBalancerDescription, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &elb.DescribeLoadBalancersInput{}
	request.LoadBalancerNames = []*string{&name}
	response, err := c.elb.DescribeLoadBalancers(request)
	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			if awsError.Code() == "LoadBalancerNotFound" {
				return nil, nil
			}
		}
		return nil, err
	}
	var ret *elb.LoadBalancerDescription
	for _, loadBalancer := range response.LoadBalancerDescriptions {
		if ret != nil {
			klog.Errorf("Found multiple load balancers with name: %s", name)
		}
		ret = loadBalancer
	}
	return ret, nil
}
func (c *Cloud) addLoadBalancerTags(loadBalancerName string, requested map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var tags []*elb.Tag
	for k, v := range requested {
		tag := &elb.Tag{Key: aws.String(k), Value: aws.String(v)}
		tags = append(tags, tag)
	}
	request := &elb.AddTagsInput{}
	request.LoadBalancerNames = []*string{&loadBalancerName}
	request.Tags = tags
	_, err := c.elb.AddTags(request)
	if err != nil {
		return fmt.Errorf("error adding tags to load balancer: %v", err)
	}
	return nil
}
func (c *Cloud) describeLoadBalancerv2(name string) (*elbv2.LoadBalancer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &elbv2.DescribeLoadBalancersInput{Names: []*string{aws.String(name)}}
	response, err := c.elbv2.DescribeLoadBalancers(request)
	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			if awsError.Code() == elbv2.ErrCodeLoadBalancerNotFoundException {
				return nil, nil
			}
		}
		return nil, fmt.Errorf("Error describing load balancer: %q", err)
	}
	for i := range response.LoadBalancers {
		if aws.StringValue(response.LoadBalancers[i].Type) == elbv2.LoadBalancerTypeEnumNetwork {
			return response.LoadBalancers[i], nil
		}
	}
	return nil, fmt.Errorf("NLB '%s' could not be found", name)
}
func (c *Cloud) findVPCID() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	macs, err := c.metadata.GetMetadata("network/interfaces/macs/")
	if err != nil {
		return "", fmt.Errorf("Could not list interfaces of the instance: %q", err)
	}
	for _, macPath := range strings.Split(macs, "\n") {
		if len(macPath) == 0 {
			continue
		}
		url := fmt.Sprintf("network/interfaces/macs/%svpc-id", macPath)
		vpcID, err := c.metadata.GetMetadata(url)
		if err != nil {
			continue
		}
		return vpcID, nil
	}
	return "", fmt.Errorf("Could not find VPC ID in instance metadata")
}
func (c *Cloud) findSecurityGroup(securityGroupID string) (*ec2.SecurityGroup, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	describeSecurityGroupsRequest := &ec2.DescribeSecurityGroupsInput{GroupIds: []*string{&securityGroupID}}
	groups, err := c.ec2.DescribeSecurityGroups(describeSecurityGroupsRequest)
	if err != nil {
		klog.Warningf("Error retrieving security group: %q", err)
		return nil, err
	}
	if len(groups) == 0 {
		return nil, nil
	}
	if len(groups) != 1 {
		return nil, fmt.Errorf("multiple security groups found with same id %q", securityGroupID)
	}
	group := groups[0]
	return group, nil
}
func isEqualIntPointer(l, r *int64) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if l == nil {
		return r == nil
	}
	if r == nil {
		return l == nil
	}
	return *l == *r
}
func isEqualStringPointer(l, r *string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if l == nil {
		return r == nil
	}
	if r == nil {
		return l == nil
	}
	return *l == *r
}
func ipPermissionExists(newPermission, existing *ec2.IpPermission, compareGroupUserIDs bool) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !isEqualIntPointer(newPermission.FromPort, existing.FromPort) {
		return false
	}
	if !isEqualIntPointer(newPermission.ToPort, existing.ToPort) {
		return false
	}
	if !isEqualStringPointer(newPermission.IpProtocol, existing.IpProtocol) {
		return false
	}
	klog.V(4).Infof("Comparing %v to %v", newPermission, existing)
	if len(newPermission.IpRanges) > len(existing.IpRanges) {
		return false
	}
	for j := range newPermission.IpRanges {
		found := false
		for k := range existing.IpRanges {
			if isEqualStringPointer(newPermission.IpRanges[j].CidrIp, existing.IpRanges[k].CidrIp) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	for _, leftPair := range newPermission.UserIdGroupPairs {
		found := false
		for _, rightPair := range existing.UserIdGroupPairs {
			if isEqualUserGroupPair(leftPair, rightPair, compareGroupUserIDs) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func isEqualUserGroupPair(l, r *ec2.UserIdGroupPair, compareGroupUserIDs bool) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("Comparing %v to %v", *l.GroupId, *r.GroupId)
	if isEqualStringPointer(l.GroupId, r.GroupId) {
		if compareGroupUserIDs {
			if isEqualStringPointer(l.UserId, r.UserId) {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
func (c *Cloud) setSecurityGroupIngress(securityGroupID string, permissions IPPermissionSet) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if securityGroupID == c.cfg.Global.ElbSecurityGroup {
		return false, nil
	}
	group, err := c.findSecurityGroup(securityGroupID)
	if err != nil {
		klog.Warningf("Error retrieving security group %q", err)
		return false, err
	}
	if group == nil {
		return false, fmt.Errorf("security group not found: %s", securityGroupID)
	}
	klog.V(2).Infof("Existing security group ingress: %s %v", securityGroupID, group.IpPermissions)
	actual := NewIPPermissionSet(group.IpPermissions...)
	permissions = permissions.Ungroup()
	actual = actual.Ungroup()
	remove := actual.Difference(permissions)
	add := permissions.Difference(actual)
	if add.Len() == 0 && remove.Len() == 0 {
		return false, nil
	}
	if add.Len() != 0 {
		klog.V(2).Infof("Adding security group ingress: %s %v", securityGroupID, add.List())
		request := &ec2.AuthorizeSecurityGroupIngressInput{}
		request.GroupId = &securityGroupID
		request.IpPermissions = add.List()
		_, err = c.ec2.AuthorizeSecurityGroupIngress(request)
		if err != nil {
			return false, fmt.Errorf("error authorizing security group ingress: %q", err)
		}
	}
	if remove.Len() != 0 {
		klog.V(2).Infof("Remove security group ingress: %s %v", securityGroupID, remove.List())
		request := &ec2.RevokeSecurityGroupIngressInput{}
		request.GroupId = &securityGroupID
		request.IpPermissions = remove.List()
		_, err = c.ec2.RevokeSecurityGroupIngress(request)
		if err != nil {
			return false, fmt.Errorf("error revoking security group ingress: %q", err)
		}
	}
	return true, nil
}
func (c *Cloud) addSecurityGroupIngress(securityGroupID string, addPermissions []*ec2.IpPermission) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if securityGroupID == c.cfg.Global.ElbSecurityGroup {
		return false, nil
	}
	group, err := c.findSecurityGroup(securityGroupID)
	if err != nil {
		klog.Warningf("Error retrieving security group: %q", err)
		return false, err
	}
	if group == nil {
		return false, fmt.Errorf("security group not found: %s", securityGroupID)
	}
	klog.V(2).Infof("Existing security group ingress: %s %v", securityGroupID, group.IpPermissions)
	changes := []*ec2.IpPermission{}
	for _, addPermission := range addPermissions {
		hasUserID := false
		for i := range addPermission.UserIdGroupPairs {
			if addPermission.UserIdGroupPairs[i].UserId != nil {
				hasUserID = true
			}
		}
		found := false
		for _, groupPermission := range group.IpPermissions {
			if ipPermissionExists(addPermission, groupPermission, hasUserID) {
				found = true
				break
			}
		}
		if !found {
			changes = append(changes, addPermission)
		}
	}
	if len(changes) == 0 {
		return false, nil
	}
	klog.V(2).Infof("Adding security group ingress: %s %v", securityGroupID, changes)
	request := &ec2.AuthorizeSecurityGroupIngressInput{}
	request.GroupId = &securityGroupID
	request.IpPermissions = changes
	_, err = c.ec2.AuthorizeSecurityGroupIngress(request)
	if err != nil {
		klog.Warningf("Error authorizing security group ingress %q", err)
		return false, fmt.Errorf("error authorizing security group ingress: %q", err)
	}
	return true, nil
}
func (c *Cloud) removeSecurityGroupIngress(securityGroupID string, removePermissions []*ec2.IpPermission) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if securityGroupID == c.cfg.Global.ElbSecurityGroup {
		return false, nil
	}
	group, err := c.findSecurityGroup(securityGroupID)
	if err != nil {
		klog.Warningf("Error retrieving security group: %q", err)
		return false, err
	}
	if group == nil {
		klog.Warning("Security group not found: ", securityGroupID)
		return false, nil
	}
	changes := []*ec2.IpPermission{}
	for _, removePermission := range removePermissions {
		hasUserID := false
		for i := range removePermission.UserIdGroupPairs {
			if removePermission.UserIdGroupPairs[i].UserId != nil {
				hasUserID = true
			}
		}
		var found *ec2.IpPermission
		for _, groupPermission := range group.IpPermissions {
			if ipPermissionExists(removePermission, groupPermission, hasUserID) {
				found = removePermission
				break
			}
		}
		if found != nil {
			changes = append(changes, found)
		}
	}
	if len(changes) == 0 {
		return false, nil
	}
	klog.V(2).Infof("Removing security group ingress: %s %v", securityGroupID, changes)
	request := &ec2.RevokeSecurityGroupIngressInput{}
	request.GroupId = &securityGroupID
	request.IpPermissions = changes
	_, err = c.ec2.RevokeSecurityGroupIngress(request)
	if err != nil {
		klog.Warningf("Error revoking security group ingress: %q", err)
		return false, err
	}
	return true, nil
}
func (c *Cloud) ensureSecurityGroup(name string, description string, additionalTags map[string]string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	groupID := ""
	attempt := 0
	for {
		attempt++
		request := &ec2.DescribeSecurityGroupsInput{}
		filters := []*ec2.Filter{newEc2Filter("group-name", name), newEc2Filter("vpc-id", c.vpcID)}
		request.Filters = filters
		securityGroups, err := c.ec2.DescribeSecurityGroups(request)
		if err != nil {
			return "", err
		}
		if len(securityGroups) >= 1 {
			if len(securityGroups) > 1 {
				klog.Warningf("Found multiple security groups with name: %q", name)
			}
			err := c.tagging.readRepairClusterTags(c.ec2, aws.StringValue(securityGroups[0].GroupId), ResourceLifecycleOwned, nil, securityGroups[0].Tags)
			if err != nil {
				return "", err
			}
			return aws.StringValue(securityGroups[0].GroupId), nil
		}
		createRequest := &ec2.CreateSecurityGroupInput{}
		createRequest.VpcId = &c.vpcID
		createRequest.GroupName = &name
		createRequest.Description = &description
		createResponse, err := c.ec2.CreateSecurityGroup(createRequest)
		if err != nil {
			ignore := false
			switch err := err.(type) {
			case awserr.Error:
				if err.Code() == "InvalidGroup.Duplicate" && attempt < MaxReadThenCreateRetries {
					klog.V(2).Infof("Got InvalidGroup.Duplicate while creating security group (race?); will retry")
					ignore = true
				}
			}
			if !ignore {
				klog.Errorf("Error creating security group: %q", err)
				return "", err
			}
			time.Sleep(1 * time.Second)
		} else {
			groupID = aws.StringValue(createResponse.GroupId)
			break
		}
	}
	if groupID == "" {
		return "", fmt.Errorf("created security group, but id was not returned: %s", name)
	}
	err := c.tagging.createTags(c.ec2, groupID, ResourceLifecycleOwned, additionalTags)
	if err != nil {
		return "", fmt.Errorf("error tagging security group: %q", err)
	}
	return groupID, nil
}
func findTag(tags []*ec2.Tag, key string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, tag := range tags {
		if aws.StringValue(tag.Key) == key {
			return aws.StringValue(tag.Value), true
		}
	}
	return "", false
}
func (c *Cloud) findSubnets() ([]*ec2.Subnet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &ec2.DescribeSubnetsInput{}
	filters := []*ec2.Filter{newEc2Filter("vpc-id", c.vpcID)}
	request.Filters = c.tagging.addFilters(filters)
	subnets, err := c.ec2.DescribeSubnets(request)
	if err != nil {
		return nil, fmt.Errorf("error describing subnets: %q", err)
	}
	var matches []*ec2.Subnet
	for _, subnet := range subnets {
		if c.tagging.hasClusterTag(subnet.Tags) {
			matches = append(matches, subnet)
		}
	}
	if len(matches) != 0 {
		return matches, nil
	}
	klog.Warningf("No tagged subnets found; will fall-back to the current subnet only.  This is likely to be an error in a future version of k8s.")
	request = &ec2.DescribeSubnetsInput{}
	filters = []*ec2.Filter{newEc2Filter("subnet-id", c.selfAWSInstance.subnetID)}
	request.Filters = filters
	subnets, err = c.ec2.DescribeSubnets(request)
	if err != nil {
		return nil, fmt.Errorf("error describing subnets: %q", err)
	}
	return subnets, nil
}
func (c *Cloud) findELBSubnets(internalELB bool) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vpcIDFilter := newEc2Filter("vpc-id", c.vpcID)
	subnets, err := c.findSubnets()
	if err != nil {
		return nil, err
	}
	rRequest := &ec2.DescribeRouteTablesInput{}
	rRequest.Filters = []*ec2.Filter{vpcIDFilter}
	rt, err := c.ec2.DescribeRouteTables(rRequest)
	if err != nil {
		return nil, fmt.Errorf("error describe route table: %q", err)
	}
	subnetsByAZ := make(map[string]*ec2.Subnet)
	for _, subnet := range subnets {
		az := aws.StringValue(subnet.AvailabilityZone)
		id := aws.StringValue(subnet.SubnetId)
		if az == "" || id == "" {
			klog.Warningf("Ignoring subnet with empty az/id: %v", subnet)
			continue
		}
		isPublic, err := isSubnetPublic(rt, id)
		if err != nil {
			return nil, err
		}
		if !internalELB && !isPublic {
			klog.V(2).Infof("Ignoring private subnet for public ELB %q", id)
			continue
		}
		existing := subnetsByAZ[az]
		if existing == nil {
			subnetsByAZ[az] = subnet
			continue
		}
		var tagName string
		if internalELB {
			tagName = TagNameSubnetInternalELB
		} else {
			tagName = TagNameSubnetPublicELB
		}
		_, existingHasTag := findTag(existing.Tags, tagName)
		_, subnetHasTag := findTag(subnet.Tags, tagName)
		if existingHasTag != subnetHasTag {
			if subnetHasTag {
				subnetsByAZ[az] = subnet
			}
			continue
		}
		if strings.Compare(*existing.SubnetId, *subnet.SubnetId) > 0 {
			klog.Warningf("Found multiple subnets in AZ %q; choosing %q between subnets %q and %q", az, *subnet.SubnetId, *existing.SubnetId, *subnet.SubnetId)
			subnetsByAZ[az] = subnet
			continue
		}
		klog.Warningf("Found multiple subnets in AZ %q; choosing %q between subnets %q and %q", az, *existing.SubnetId, *existing.SubnetId, *subnet.SubnetId)
		continue
	}
	var subnetIDs []string
	for _, subnet := range subnetsByAZ {
		subnetIDs = append(subnetIDs, aws.StringValue(subnet.SubnetId))
	}
	return subnetIDs, nil
}
func isSubnetPublic(rt []*ec2.RouteTable, subnetID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var subnetTable *ec2.RouteTable
	for _, table := range rt {
		for _, assoc := range table.Associations {
			if aws.StringValue(assoc.SubnetId) == subnetID {
				subnetTable = table
				break
			}
		}
	}
	if subnetTable == nil {
		for _, table := range rt {
			for _, assoc := range table.Associations {
				if aws.BoolValue(assoc.Main) == true {
					klog.V(4).Infof("Assuming implicit use of main routing table %s for %s", aws.StringValue(table.RouteTableId), subnetID)
					subnetTable = table
					break
				}
			}
		}
	}
	if subnetTable == nil {
		return false, fmt.Errorf("Could not locate routing table for subnet %s", subnetID)
	}
	for _, route := range subnetTable.Routes {
		if strings.HasPrefix(aws.StringValue(route.GatewayId), "igw") {
			return true, nil
		}
	}
	return false, nil
}

type portSets struct {
	names   sets.String
	numbers sets.Int64
}

func getPortSets(annotation string) (ports *portSets) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if annotation != "" && annotation != "*" {
		ports = &portSets{sets.NewString(), sets.NewInt64()}
		portStringSlice := strings.Split(annotation, ",")
		for _, item := range portStringSlice {
			port, err := strconv.Atoi(item)
			if err != nil {
				ports.names.Insert(item)
			} else {
				ports.numbers.Insert(int64(port))
			}
		}
	}
	return
}
func (c *Cloud) buildELBSecurityGroupList(serviceName types.NamespacedName, loadBalancerName string, annotations map[string]string) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	var securityGroupID string
	if c.cfg.Global.ElbSecurityGroup != "" {
		securityGroupID = c.cfg.Global.ElbSecurityGroup
	} else {
		sgName := "k8s-elb-" + loadBalancerName
		sgDescription := fmt.Sprintf("Security group for Kubernetes ELB %s (%v)", loadBalancerName, serviceName)
		securityGroupID, err = c.ensureSecurityGroup(sgName, sgDescription, getLoadBalancerAdditionalTags(annotations))
		if err != nil {
			klog.Errorf("Error creating load balancer security group: %q", err)
			return nil, err
		}
	}
	sgList := []string{}
	for _, extraSG := range strings.Split(annotations[ServiceAnnotationLoadBalancerSecurityGroups], ",") {
		extraSG = strings.TrimSpace(extraSG)
		if len(extraSG) > 0 {
			sgList = append(sgList, extraSG)
		}
	}
	if len(sgList) == 0 {
		sgList = append(sgList, securityGroupID)
	}
	for _, extraSG := range strings.Split(annotations[ServiceAnnotationLoadBalancerExtraSecurityGroups], ",") {
		extraSG = strings.TrimSpace(extraSG)
		if len(extraSG) > 0 {
			sgList = append(sgList, extraSG)
		}
	}
	return sgList, nil
}
func buildListener(port v1.ServicePort, annotations map[string]string, sslPorts *portSets) (*elb.Listener, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	loadBalancerPort := int64(port.Port)
	portName := strings.ToLower(port.Name)
	instancePort := int64(port.NodePort)
	protocol := strings.ToLower(string(port.Protocol))
	instanceProtocol := protocol
	listener := &elb.Listener{}
	listener.InstancePort = &instancePort
	listener.LoadBalancerPort = &loadBalancerPort
	certID := annotations[ServiceAnnotationLoadBalancerCertificate]
	if certID != "" && (sslPorts == nil || sslPorts.numbers.Has(loadBalancerPort) || sslPorts.names.Has(portName)) {
		instanceProtocol = annotations[ServiceAnnotationLoadBalancerBEProtocol]
		if instanceProtocol == "" {
			protocol = "ssl"
			instanceProtocol = "tcp"
		} else {
			protocol = backendProtocolMapping[instanceProtocol]
			if protocol == "" {
				return nil, fmt.Errorf("Invalid backend protocol %s for %s in %s", instanceProtocol, certID, ServiceAnnotationLoadBalancerBEProtocol)
			}
		}
		listener.SSLCertificateId = &certID
	} else if annotationProtocol := annotations[ServiceAnnotationLoadBalancerBEProtocol]; annotationProtocol == "http" {
		instanceProtocol = annotationProtocol
		protocol = "http"
	}
	listener.Protocol = &protocol
	listener.InstanceProtocol = &instanceProtocol
	return listener, nil
}
func (c *Cloud) EnsureLoadBalancer(ctx context.Context, clusterName string, apiService *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	annotations := apiService.Annotations
	klog.V(2).Infof("EnsureLoadBalancer(%v, %v, %v, %v, %v, %v, %v)", clusterName, apiService.Namespace, apiService.Name, c.region, apiService.Spec.LoadBalancerIP, apiService.Spec.Ports, annotations)
	if apiService.Spec.SessionAffinity != v1.ServiceAffinityNone {
		return nil, fmt.Errorf("unsupported load balancer affinity: %v", apiService.Spec.SessionAffinity)
	}
	if len(apiService.Spec.Ports) == 0 {
		return nil, fmt.Errorf("requested load balancer with no ports")
	}
	listeners := []*elb.Listener{}
	v2Mappings := []nlbPortMapping{}
	portList := getPortSets(annotations[ServiceAnnotationLoadBalancerSSLPorts])
	for _, port := range apiService.Spec.Ports {
		if port.Protocol != v1.ProtocolTCP {
			return nil, fmt.Errorf("Only TCP LoadBalancer is supported for AWS ELB")
		}
		if port.NodePort == 0 {
			klog.Errorf("Ignoring port without NodePort defined: %v", port)
			continue
		}
		if isNLB(annotations) {
			v2Mappings = append(v2Mappings, nlbPortMapping{FrontendPort: int64(port.Port), TrafficPort: int64(port.NodePort), HealthCheckPort: int64(port.NodePort), HealthCheckProtocol: elbv2.ProtocolEnumTcp})
		}
		listener, err := buildListener(port, annotations, portList)
		if err != nil {
			return nil, err
		}
		listeners = append(listeners, listener)
	}
	if apiService.Spec.LoadBalancerIP != "" {
		return nil, fmt.Errorf("LoadBalancerIP cannot be specified for AWS ELB")
	}
	instances, err := c.findInstancesForELB(nodes)
	if err != nil {
		return nil, err
	}
	sourceRanges, err := service.GetLoadBalancerSourceRanges(apiService)
	if err != nil {
		return nil, err
	}
	internalELB := false
	internalAnnotation := apiService.Annotations[ServiceAnnotationLoadBalancerInternal]
	if internalAnnotation == "false" {
		internalELB = false
	} else if internalAnnotation != "" {
		internalELB = true
	}
	if isNLB(annotations) {
		if path, healthCheckNodePort := service.GetServiceHealthCheckPathPort(apiService); path != "" {
			for i := range v2Mappings {
				v2Mappings[i].HealthCheckPort = int64(healthCheckNodePort)
				v2Mappings[i].HealthCheckPath = path
				v2Mappings[i].HealthCheckProtocol = elbv2.ProtocolEnumHttp
			}
		}
		subnetIDs, err := c.findELBSubnets(internalELB)
		if err != nil {
			klog.Errorf("Error listing subnets in VPC: %q", err)
			return nil, err
		}
		if len(subnetIDs) == 0 {
			return nil, fmt.Errorf("could not find any suitable subnets for creating the ELB")
		}
		loadBalancerName := c.GetLoadBalancerName(ctx, clusterName, apiService)
		serviceName := types.NamespacedName{Namespace: apiService.Namespace, Name: apiService.Name}
		instanceIDs := []string{}
		for id := range instances {
			instanceIDs = append(instanceIDs, string(id))
		}
		v2LoadBalancer, err := c.ensureLoadBalancerv2(serviceName, loadBalancerName, v2Mappings, instanceIDs, subnetIDs, internalELB, annotations)
		if err != nil {
			return nil, err
		}
		sourceRangeCidrs := []string{}
		for cidr := range sourceRanges {
			sourceRangeCidrs = append(sourceRangeCidrs, cidr)
		}
		if len(sourceRangeCidrs) == 0 {
			sourceRangeCidrs = append(sourceRangeCidrs, "0.0.0.0/0")
		}
		err = c.updateInstanceSecurityGroupsForNLB(v2Mappings, instances, loadBalancerName, sourceRangeCidrs)
		if err != nil {
			klog.Warningf("Error opening ingress rules for the load balancer to the instances: %q", err)
			return nil, err
		}
		return v2toStatus(v2LoadBalancer), nil
	}
	proxyProtocol := false
	proxyProtocolAnnotation := apiService.Annotations[ServiceAnnotationLoadBalancerProxyProtocol]
	if proxyProtocolAnnotation != "" {
		if proxyProtocolAnnotation != "*" {
			return nil, fmt.Errorf("annotation %q=%q detected, but the only value supported currently is '*'", ServiceAnnotationLoadBalancerProxyProtocol, proxyProtocolAnnotation)
		}
		proxyProtocol = true
	}
	loadBalancerAttributes := &elb.LoadBalancerAttributes{AccessLog: &elb.AccessLog{Enabled: aws.Bool(false)}, ConnectionDraining: &elb.ConnectionDraining{Enabled: aws.Bool(false)}, ConnectionSettings: &elb.ConnectionSettings{IdleTimeout: aws.Int64(60)}, CrossZoneLoadBalancing: &elb.CrossZoneLoadBalancing{Enabled: aws.Bool(false)}}
	accessLogEmitIntervalAnnotation := annotations[ServiceAnnotationLoadBalancerAccessLogEmitInterval]
	if accessLogEmitIntervalAnnotation != "" {
		accessLogEmitInterval, err := strconv.ParseInt(accessLogEmitIntervalAnnotation, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerAccessLogEmitInterval, accessLogEmitIntervalAnnotation)
		}
		loadBalancerAttributes.AccessLog.EmitInterval = &accessLogEmitInterval
	}
	accessLogEnabledAnnotation := annotations[ServiceAnnotationLoadBalancerAccessLogEnabled]
	if accessLogEnabledAnnotation != "" {
		accessLogEnabled, err := strconv.ParseBool(accessLogEnabledAnnotation)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerAccessLogEnabled, accessLogEnabledAnnotation)
		}
		loadBalancerAttributes.AccessLog.Enabled = &accessLogEnabled
	}
	accessLogS3BucketNameAnnotation := annotations[ServiceAnnotationLoadBalancerAccessLogS3BucketName]
	if accessLogS3BucketNameAnnotation != "" {
		loadBalancerAttributes.AccessLog.S3BucketName = &accessLogS3BucketNameAnnotation
	}
	accessLogS3BucketPrefixAnnotation := annotations[ServiceAnnotationLoadBalancerAccessLogS3BucketPrefix]
	if accessLogS3BucketPrefixAnnotation != "" {
		loadBalancerAttributes.AccessLog.S3BucketPrefix = &accessLogS3BucketPrefixAnnotation
	}
	connectionDrainingEnabledAnnotation := annotations[ServiceAnnotationLoadBalancerConnectionDrainingEnabled]
	if connectionDrainingEnabledAnnotation != "" {
		connectionDrainingEnabled, err := strconv.ParseBool(connectionDrainingEnabledAnnotation)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerConnectionDrainingEnabled, connectionDrainingEnabledAnnotation)
		}
		loadBalancerAttributes.ConnectionDraining.Enabled = &connectionDrainingEnabled
	}
	connectionDrainingTimeoutAnnotation := annotations[ServiceAnnotationLoadBalancerConnectionDrainingTimeout]
	if connectionDrainingTimeoutAnnotation != "" {
		connectionDrainingTimeout, err := strconv.ParseInt(connectionDrainingTimeoutAnnotation, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerConnectionDrainingTimeout, connectionDrainingTimeoutAnnotation)
		}
		loadBalancerAttributes.ConnectionDraining.Timeout = &connectionDrainingTimeout
	}
	connectionIdleTimeoutAnnotation := annotations[ServiceAnnotationLoadBalancerConnectionIdleTimeout]
	if connectionIdleTimeoutAnnotation != "" {
		connectionIdleTimeout, err := strconv.ParseInt(connectionIdleTimeoutAnnotation, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerConnectionIdleTimeout, connectionIdleTimeoutAnnotation)
		}
		loadBalancerAttributes.ConnectionSettings.IdleTimeout = &connectionIdleTimeout
	}
	crossZoneLoadBalancingEnabledAnnotation := annotations[ServiceAnnotationLoadBalancerCrossZoneLoadBalancingEnabled]
	if crossZoneLoadBalancingEnabledAnnotation != "" {
		crossZoneLoadBalancingEnabled, err := strconv.ParseBool(crossZoneLoadBalancingEnabledAnnotation)
		if err != nil {
			return nil, fmt.Errorf("error parsing service annotation: %s=%s", ServiceAnnotationLoadBalancerCrossZoneLoadBalancingEnabled, crossZoneLoadBalancingEnabledAnnotation)
		}
		loadBalancerAttributes.CrossZoneLoadBalancing.Enabled = &crossZoneLoadBalancingEnabled
	}
	subnetIDs, err := c.findELBSubnets(internalELB)
	if err != nil {
		klog.Errorf("Error listing subnets in VPC: %q", err)
		return nil, err
	}
	if len(subnetIDs) == 0 {
		return nil, fmt.Errorf("could not find any suitable subnets for creating the ELB")
	}
	loadBalancerName := c.GetLoadBalancerName(ctx, clusterName, apiService)
	serviceName := types.NamespacedName{Namespace: apiService.Namespace, Name: apiService.Name}
	securityGroupIDs, err := c.buildELBSecurityGroupList(serviceName, loadBalancerName, annotations)
	if err != nil {
		return nil, err
	}
	if len(securityGroupIDs) == 0 {
		return nil, fmt.Errorf("[BUG] ELB can't have empty list of Security Groups to be assigned, this is a Kubernetes bug, please report")
	}
	{
		ec2SourceRanges := []*ec2.IpRange{}
		for _, sourceRange := range sourceRanges.StringSlice() {
			ec2SourceRanges = append(ec2SourceRanges, &ec2.IpRange{CidrIp: aws.String(sourceRange)})
		}
		permissions := NewIPPermissionSet()
		for _, port := range apiService.Spec.Ports {
			portInt64 := int64(port.Port)
			protocol := strings.ToLower(string(port.Protocol))
			permission := &ec2.IpPermission{}
			permission.FromPort = &portInt64
			permission.ToPort = &portInt64
			permission.IpRanges = ec2SourceRanges
			permission.IpProtocol = &protocol
			permissions.Insert(permission)
		}
		{
			permission := &ec2.IpPermission{IpProtocol: aws.String("icmp"), FromPort: aws.Int64(3), ToPort: aws.Int64(4), IpRanges: ec2SourceRanges}
			permissions.Insert(permission)
		}
		_, err = c.setSecurityGroupIngress(securityGroupIDs[0], permissions)
		if err != nil {
			return nil, err
		}
	}
	loadBalancer, err := c.ensureLoadBalancer(serviceName, loadBalancerName, listeners, subnetIDs, securityGroupIDs, internalELB, proxyProtocol, loadBalancerAttributes, annotations)
	if err != nil {
		return nil, err
	}
	if sslPolicyName, ok := annotations[ServiceAnnotationLoadBalancerSSLNegotiationPolicy]; ok {
		err := c.ensureSSLNegotiationPolicy(loadBalancer, sslPolicyName)
		if err != nil {
			return nil, err
		}
		for _, port := range c.getLoadBalancerTLSPorts(loadBalancer) {
			err := c.setSSLNegotiationPolicy(loadBalancerName, sslPolicyName, port)
			if err != nil {
				return nil, err
			}
		}
	}
	if path, healthCheckNodePort := service.GetServiceHealthCheckPathPort(apiService); path != "" {
		klog.V(4).Infof("service %v (%v) needs health checks on :%d%s)", apiService.Name, loadBalancerName, healthCheckNodePort, path)
		err = c.ensureLoadBalancerHealthCheck(loadBalancer, "HTTP", healthCheckNodePort, path, annotations)
		if err != nil {
			return nil, fmt.Errorf("Failed to ensure health check for localized service %v on node port %v: %q", loadBalancerName, healthCheckNodePort, err)
		}
	} else {
		klog.V(4).Infof("service %v does not need custom health checks", apiService.Name)
		var tcpHealthCheckPort int32
		for _, listener := range listeners {
			if listener.InstancePort == nil {
				continue
			}
			tcpHealthCheckPort = int32(*listener.InstancePort)
			break
		}
		err = c.ensureLoadBalancerHealthCheck(loadBalancer, "TCP", tcpHealthCheckPort, "", annotations)
		if err != nil {
			return nil, err
		}
	}
	err = c.updateInstanceSecurityGroupsForLoadBalancer(loadBalancer, instances)
	if err != nil {
		klog.Warningf("Error opening ingress rules for the load balancer to the instances: %q", err)
		return nil, err
	}
	err = c.ensureLoadBalancerInstances(aws.StringValue(loadBalancer.LoadBalancerName), loadBalancer.Instances, instances)
	if err != nil {
		klog.Warningf("Error registering instances with the load balancer: %q", err)
		return nil, err
	}
	klog.V(1).Infof("Loadbalancer %s (%v) has DNS name %s", loadBalancerName, serviceName, aws.StringValue(loadBalancer.DNSName))
	status := toStatus(loadBalancer)
	return status, nil
}
func (c *Cloud) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	loadBalancerName := c.GetLoadBalancerName(ctx, clusterName, service)
	if isNLB(service.Annotations) {
		lb, err := c.describeLoadBalancerv2(loadBalancerName)
		if err != nil {
			return nil, false, err
		}
		if lb == nil {
			return nil, false, nil
		}
		return v2toStatus(lb), true, nil
	}
	lb, err := c.describeLoadBalancer(loadBalancerName)
	if err != nil {
		return nil, false, err
	}
	if lb == nil {
		return nil, false, nil
	}
	status := toStatus(lb)
	return status, true, nil
}
func (c *Cloud) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cloudprovider.DefaultLoadBalancerName(service)
}
func toStatus(lb *elb.LoadBalancerDescription) *v1.LoadBalancerStatus {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	status := &v1.LoadBalancerStatus{}
	if aws.StringValue(lb.DNSName) != "" {
		var ingress v1.LoadBalancerIngress
		ingress.Hostname = aws.StringValue(lb.DNSName)
		status.Ingress = []v1.LoadBalancerIngress{ingress}
	}
	return status
}
func v2toStatus(lb *elbv2.LoadBalancer) *v1.LoadBalancerStatus {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	status := &v1.LoadBalancerStatus{}
	if lb == nil {
		klog.Error("[BUG] v2toStatus got nil input, this is a Kubernetes bug, please report")
		return status
	}
	if aws.StringValue(lb.DNSName) != "" && (aws.StringValue(lb.State.Code) == elbv2.LoadBalancerStateEnumActive || aws.StringValue(lb.State.Code) == elbv2.LoadBalancerStateEnumProvisioning) {
		var ingress v1.LoadBalancerIngress
		ingress.Hostname = aws.StringValue(lb.DNSName)
		status.Ingress = []v1.LoadBalancerIngress{ingress}
	}
	return status
}
func findSecurityGroupForInstance(instance *ec2.Instance, taggedSecurityGroups map[string]*ec2.SecurityGroup) (*ec2.GroupIdentifier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceID := aws.StringValue(instance.InstanceId)
	var tagged []*ec2.GroupIdentifier
	var untagged []*ec2.GroupIdentifier
	for _, group := range instance.SecurityGroups {
		groupID := aws.StringValue(group.GroupId)
		if groupID == "" {
			klog.Warningf("Ignoring security group without id for instance %q: %v", instanceID, group)
			continue
		}
		_, isTagged := taggedSecurityGroups[groupID]
		if isTagged {
			tagged = append(tagged, group)
		} else {
			untagged = append(untagged, group)
		}
	}
	if len(tagged) > 0 {
		if len(tagged) != 1 {
			taggedGroups := ""
			for _, v := range tagged {
				taggedGroups += fmt.Sprintf("%s(%s) ", *v.GroupId, *v.GroupName)
			}
			return nil, fmt.Errorf("Multiple tagged security groups found for instance %s; ensure only the k8s security group is tagged; the tagged groups were %v", instanceID, taggedGroups)
		}
		return tagged[0], nil
	}
	if len(untagged) > 0 {
		if len(untagged) != 1 {
			return nil, fmt.Errorf("Multiple untagged security groups found for instance %s; ensure the k8s security group is tagged", instanceID)
		}
		return untagged[0], nil
	}
	klog.Warningf("No security group found for instance %q", instanceID)
	return nil, nil
}
func (c *Cloud) getTaggedSecurityGroups() (map[string]*ec2.SecurityGroup, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &ec2.DescribeSecurityGroupsInput{}
	request.Filters = c.tagging.addFilters(nil)
	groups, err := c.ec2.DescribeSecurityGroups(request)
	if err != nil {
		return nil, fmt.Errorf("error querying security groups: %q", err)
	}
	m := make(map[string]*ec2.SecurityGroup)
	for _, group := range groups {
		if !c.tagging.hasClusterTag(group.Tags) {
			continue
		}
		id := aws.StringValue(group.GroupId)
		if id == "" {
			klog.Warningf("Ignoring group without id: %v", group)
			continue
		}
		m[id] = group
	}
	return m, nil
}
func (c *Cloud) updateInstanceSecurityGroupsForLoadBalancer(lb *elb.LoadBalancerDescription, instances map[awsInstanceID]*ec2.Instance) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.cfg.Global.DisableSecurityGroupIngress {
		return nil
	}
	loadBalancerSecurityGroupID := ""
	for _, securityGroup := range lb.SecurityGroups {
		if aws.StringValue(securityGroup) == "" {
			continue
		}
		if loadBalancerSecurityGroupID != "" {
			klog.Warningf("Multiple security groups for load balancer: %q", aws.StringValue(lb.LoadBalancerName))
		}
		loadBalancerSecurityGroupID = *securityGroup
	}
	if loadBalancerSecurityGroupID == "" {
		return fmt.Errorf("Could not determine security group for load balancer: %s", aws.StringValue(lb.LoadBalancerName))
	}
	var actualGroups []*ec2.SecurityGroup
	{
		describeRequest := &ec2.DescribeSecurityGroupsInput{}
		filters := []*ec2.Filter{newEc2Filter("ip-permission.group-id", loadBalancerSecurityGroupID)}
		describeRequest.Filters = c.tagging.addFilters(filters)
		response, err := c.ec2.DescribeSecurityGroups(describeRequest)
		if err != nil {
			return fmt.Errorf("error querying security groups for ELB: %q", err)
		}
		for _, sg := range response {
			if !c.tagging.hasClusterTag(sg.Tags) {
				continue
			}
			actualGroups = append(actualGroups, sg)
		}
	}
	taggedSecurityGroups, err := c.getTaggedSecurityGroups()
	if err != nil {
		return fmt.Errorf("error querying for tagged security groups: %q", err)
	}
	instanceSecurityGroupIds := map[string]bool{}
	for _, instance := range instances {
		securityGroup, err := findSecurityGroupForInstance(instance, taggedSecurityGroups)
		if err != nil {
			return err
		}
		if securityGroup == nil {
			klog.Warning("Ignoring instance without security group: ", aws.StringValue(instance.InstanceId))
			continue
		}
		id := aws.StringValue(securityGroup.GroupId)
		if id == "" {
			klog.Warningf("found security group without id: %v", securityGroup)
			continue
		}
		instanceSecurityGroupIds[id] = true
	}
	for _, actualGroup := range actualGroups {
		actualGroupID := aws.StringValue(actualGroup.GroupId)
		if actualGroupID == "" {
			klog.Warning("Ignoring group without ID: ", actualGroup)
			continue
		}
		adding, found := instanceSecurityGroupIds[actualGroupID]
		if found && adding {
			delete(instanceSecurityGroupIds, actualGroupID)
		} else {
			instanceSecurityGroupIds[actualGroupID] = false
		}
	}
	for instanceSecurityGroupID, add := range instanceSecurityGroupIds {
		if add {
			klog.V(2).Infof("Adding rule for traffic from the load balancer (%s) to instances (%s)", loadBalancerSecurityGroupID, instanceSecurityGroupID)
		} else {
			klog.V(2).Infof("Removing rule for traffic from the load balancer (%s) to instance (%s)", loadBalancerSecurityGroupID, instanceSecurityGroupID)
		}
		sourceGroupID := &ec2.UserIdGroupPair{}
		sourceGroupID.GroupId = &loadBalancerSecurityGroupID
		allProtocols := "-1"
		permission := &ec2.IpPermission{}
		permission.IpProtocol = &allProtocols
		permission.UserIdGroupPairs = []*ec2.UserIdGroupPair{sourceGroupID}
		permissions := []*ec2.IpPermission{permission}
		if add {
			changed, err := c.addSecurityGroupIngress(instanceSecurityGroupID, permissions)
			if err != nil {
				return err
			}
			if !changed {
				klog.Warning("Allowing ingress was not needed; concurrent change? groupId=", instanceSecurityGroupID)
			}
		} else {
			changed, err := c.removeSecurityGroupIngress(instanceSecurityGroupID, permissions)
			if err != nil {
				return err
			}
			if !changed {
				klog.Warning("Revoking ingress was not needed; concurrent change? groupId=", instanceSecurityGroupID)
			}
		}
	}
	return nil
}
func (c *Cloud) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	loadBalancerName := c.GetLoadBalancerName(ctx, clusterName, service)
	if isNLB(service.Annotations) {
		lb, err := c.describeLoadBalancerv2(loadBalancerName)
		if err != nil {
			return err
		}
		if lb == nil {
			klog.Info("Load balancer already deleted: ", loadBalancerName)
			return nil
		}
		{
			targetGroups, err := c.elbv2.DescribeTargetGroups(&elbv2.DescribeTargetGroupsInput{LoadBalancerArn: lb.LoadBalancerArn})
			if err != nil {
				return fmt.Errorf("Error listing target groups before deleting load balancer: %q", err)
			}
			_, err = c.elbv2.DeleteLoadBalancer(&elbv2.DeleteLoadBalancerInput{LoadBalancerArn: lb.LoadBalancerArn})
			if err != nil {
				return fmt.Errorf("Error deleting load balancer %q: %v", loadBalancerName, err)
			}
			for _, group := range targetGroups.TargetGroups {
				_, err := c.elbv2.DeleteTargetGroup(&elbv2.DeleteTargetGroupInput{TargetGroupArn: group.TargetGroupArn})
				if err != nil {
					return fmt.Errorf("Error deleting target groups after deleting load balancer: %q", err)
				}
			}
		}
		{
			var matchingGroups []*ec2.SecurityGroup
			{
				describeRequest := &ec2.DescribeSecurityGroupsInput{}
				filters := []*ec2.Filter{newEc2Filter("ip-permission.protocol", "tcp")}
				describeRequest.Filters = c.tagging.addFilters(filters)
				response, err := c.ec2.DescribeSecurityGroups(describeRequest)
				if err != nil {
					return fmt.Errorf("Error querying security groups for NLB: %q", err)
				}
				for _, sg := range response {
					if !c.tagging.hasClusterTag(sg.Tags) {
						continue
					}
					matchingGroups = append(matchingGroups, sg)
				}
				matchingGroups = filterForIPRangeDescription(matchingGroups, loadBalancerName)
			}
			{
				clientRule := fmt.Sprintf("%s=%s", NLBClientRuleDescription, loadBalancerName)
				mtuRule := fmt.Sprintf("%s=%s", NLBMtuDiscoveryRuleDescription, loadBalancerName)
				healthRule := fmt.Sprintf("%s=%s", NLBHealthCheckRuleDescription, loadBalancerName)
				for i := range matchingGroups {
					removes := []*ec2.IpPermission{}
					for j := range matchingGroups[i].IpPermissions {
						v4rangesToRemove := []*ec2.IpRange{}
						v6rangesToRemove := []*ec2.Ipv6Range{}
						for k := range matchingGroups[i].IpPermissions[j].IpRanges {
							description := aws.StringValue(matchingGroups[i].IpPermissions[j].IpRanges[k].Description)
							if description == clientRule || description == mtuRule || description == healthRule {
								v4rangesToRemove = append(v4rangesToRemove, matchingGroups[i].IpPermissions[j].IpRanges[k])
							}
						}
						for k := range matchingGroups[i].IpPermissions[j].Ipv6Ranges {
							description := aws.StringValue(matchingGroups[i].IpPermissions[j].Ipv6Ranges[k].Description)
							if description == clientRule || description == mtuRule || description == healthRule {
								v6rangesToRemove = append(v6rangesToRemove, matchingGroups[i].IpPermissions[j].Ipv6Ranges[k])
							}
						}
						if len(v4rangesToRemove) > 0 {
							removedPermission := &ec2.IpPermission{FromPort: matchingGroups[i].IpPermissions[j].FromPort, IpProtocol: matchingGroups[i].IpPermissions[j].IpProtocol, IpRanges: v4rangesToRemove, ToPort: matchingGroups[i].IpPermissions[j].ToPort}
							removes = append(removes, removedPermission)
						}
						if len(v6rangesToRemove) > 0 {
							removedPermission := &ec2.IpPermission{FromPort: matchingGroups[i].IpPermissions[j].FromPort, IpProtocol: matchingGroups[i].IpPermissions[j].IpProtocol, Ipv6Ranges: v6rangesToRemove, ToPort: matchingGroups[i].IpPermissions[j].ToPort}
							removes = append(removes, removedPermission)
						}
					}
					if len(removes) > 0 {
						changed, err := c.removeSecurityGroupIngress(aws.StringValue(matchingGroups[i].GroupId), removes)
						if err != nil {
							return err
						}
						if !changed {
							klog.Warning("Revoking ingress was not needed; concurrent change? groupId=", *matchingGroups[i].GroupId)
						}
					}
				}
			}
		}
		return nil
	}
	lb, err := c.describeLoadBalancer(loadBalancerName)
	if err != nil {
		return err
	}
	if lb == nil {
		klog.Info("Load balancer already deleted: ", loadBalancerName)
		return nil
	}
	{
		err = c.updateInstanceSecurityGroupsForLoadBalancer(lb, nil)
		if err != nil {
			klog.Errorf("Error deregistering load balancer from instance security groups: %q", err)
			return err
		}
	}
	{
		request := &elb.DeleteLoadBalancerInput{}
		request.LoadBalancerName = lb.LoadBalancerName
		_, err = c.elb.DeleteLoadBalancer(request)
		if err != nil {
			klog.Errorf("Error deleting load balancer: %q", err)
			return err
		}
	}
	{
		securityGroupIDs := map[string]struct{}{}
		for _, securityGroupID := range lb.SecurityGroups {
			if *securityGroupID == c.cfg.Global.ElbSecurityGroup {
				continue
			}
			if aws.StringValue(securityGroupID) == "" {
				klog.Warning("Ignoring empty security group in ", service.Name)
				continue
			}
			securityGroupIDs[*securityGroupID] = struct{}{}
		}
		timeoutAt := time.Now().Add(time.Second * 600)
		for {
			for securityGroupID := range securityGroupIDs {
				request := &ec2.DeleteSecurityGroupInput{}
				request.GroupId = &securityGroupID
				_, err := c.ec2.DeleteSecurityGroup(request)
				if err == nil {
					delete(securityGroupIDs, securityGroupID)
				} else {
					ignore := false
					if awsError, ok := err.(awserr.Error); ok {
						if awsError.Code() == "DependencyViolation" {
							klog.V(2).Infof("Ignoring DependencyViolation while deleting load-balancer security group (%s), assuming because LB is in process of deleting", securityGroupID)
							ignore = true
						}
					}
					if !ignore {
						return fmt.Errorf("error while deleting load balancer security group (%s): %q", securityGroupID, err)
					}
				}
			}
			if len(securityGroupIDs) == 0 {
				klog.V(2).Info("Deleted all security groups for load balancer: ", service.Name)
				break
			}
			if time.Now().After(timeoutAt) {
				ids := []string{}
				for id := range securityGroupIDs {
					ids = append(ids, id)
				}
				return fmt.Errorf("timed out deleting ELB: %s. Could not delete security groups %v", service.Name, strings.Join(ids, ","))
			}
			klog.V(2).Info("Waiting for load-balancer to delete so we can delete security groups: ", service.Name)
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}
func (c *Cloud) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instances, err := c.findInstancesForELB(nodes)
	if err != nil {
		return err
	}
	loadBalancerName := c.GetLoadBalancerName(ctx, clusterName, service)
	if isNLB(service.Annotations) {
		lb, err := c.describeLoadBalancerv2(loadBalancerName)
		if err != nil {
			return err
		}
		if lb == nil {
			return fmt.Errorf("Load balancer not found")
		}
		_, err = c.EnsureLoadBalancer(ctx, clusterName, service, nodes)
		return err
	}
	lb, err := c.describeLoadBalancer(loadBalancerName)
	if err != nil {
		return err
	}
	if lb == nil {
		return fmt.Errorf("Load balancer not found")
	}
	if sslPolicyName, ok := service.Annotations[ServiceAnnotationLoadBalancerSSLNegotiationPolicy]; ok {
		err := c.ensureSSLNegotiationPolicy(lb, sslPolicyName)
		if err != nil {
			return err
		}
		for _, port := range c.getLoadBalancerTLSPorts(lb) {
			err := c.setSSLNegotiationPolicy(loadBalancerName, sslPolicyName, port)
			if err != nil {
				return err
			}
		}
	}
	err = c.ensureLoadBalancerInstances(aws.StringValue(lb.LoadBalancerName), lb.Instances, instances)
	if err != nil {
		return nil
	}
	err = c.updateInstanceSecurityGroupsForLoadBalancer(lb, instances)
	if err != nil {
		return err
	}
	return nil
}
func (c *Cloud) getInstanceByID(instanceID string) (*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instances, err := c.getInstancesByIDs([]*string{&instanceID})
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, cloudprovider.InstanceNotFound
	}
	if len(instances) > 1 {
		return nil, fmt.Errorf("multiple instances found for instance: %s", instanceID)
	}
	return instances[instanceID], nil
}
func (c *Cloud) getInstancesByIDs(instanceIDs []*string) (map[string]*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instancesByID := make(map[string]*ec2.Instance)
	if len(instanceIDs) == 0 {
		return instancesByID, nil
	}
	request := &ec2.DescribeInstancesInput{InstanceIds: instanceIDs}
	instances, err := c.ec2.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	for _, instance := range instances {
		instanceID := aws.StringValue(instance.InstanceId)
		if instanceID == "" {
			continue
		}
		instancesByID[instanceID] = instance
	}
	return instancesByID, nil
}
func (c *Cloud) getInstancesByNodeNames(nodeNames []string, states ...string) ([]*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	names := aws.StringSlice(nodeNames)
	ec2Instances := []*ec2.Instance{}
	for i := 0; i < len(names); i += filterNodeLimit {
		end := i + filterNodeLimit
		if end > len(names) {
			end = len(names)
		}
		nameSlice := names[i:end]
		nodeNameFilter := &ec2.Filter{Name: aws.String("private-dns-name"), Values: nameSlice}
		filters := []*ec2.Filter{nodeNameFilter}
		if len(states) > 0 {
			filters = append(filters, newEc2Filter("instance-state-name", states...))
		}
		instances, err := c.describeInstances(filters)
		if err != nil {
			klog.V(2).Infof("Failed to describe instances %v", nodeNames)
			return nil, err
		}
		ec2Instances = append(ec2Instances, instances...)
	}
	if len(ec2Instances) == 0 {
		klog.V(3).Infof("Failed to find any instances %v", nodeNames)
		return nil, nil
	}
	return ec2Instances, nil
}
func (c *Cloud) describeInstances(filters []*ec2.Filter) ([]*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filters = c.tagging.addFilters(filters)
	request := &ec2.DescribeInstancesInput{Filters: filters}
	response, err := c.ec2.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	var matches []*ec2.Instance
	for _, instance := range response {
		if c.tagging.hasClusterTag(instance.Tags) {
			matches = append(matches, instance)
		}
	}
	return matches, nil
}
func mapNodeNameToPrivateDNSName(nodeName types.NodeName) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(nodeName)
}
func mapInstanceToNodeName(i *ec2.Instance) types.NodeName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return types.NodeName(aws.StringValue(i.PrivateDnsName))
}

var aliveFilter = []string{ec2.InstanceStateNamePending, ec2.InstanceStateNameRunning, ec2.InstanceStateNameShuttingDown, ec2.InstanceStateNameStopping, ec2.InstanceStateNameStopped}

func (c *Cloud) findInstanceByNodeName(nodeName types.NodeName) (*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	privateDNSName := mapNodeNameToPrivateDNSName(nodeName)
	filters := []*ec2.Filter{newEc2Filter("private-dns-name", privateDNSName), newEc2Filter("instance-state-name", aliveFilter...)}
	instances, err := c.describeInstances(filters)
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, nil
	}
	if len(instances) > 1 {
		return nil, fmt.Errorf("multiple instances found for name: %s", nodeName)
	}
	return instances[0], nil
}
func (c *Cloud) getInstanceByNodeName(nodeName types.NodeName) (*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instance, err := c.findInstanceByNodeName(nodeName)
	if err == nil && instance == nil {
		return nil, cloudprovider.InstanceNotFound
	}
	return instance, err
}
func (c *Cloud) getFullInstance(nodeName types.NodeName) (*awsInstance, *ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if nodeName == "" {
		instance, err := c.getInstanceByID(c.selfAWSInstance.awsID)
		return c.selfAWSInstance, instance, err
	}
	instance, err := c.getInstanceByNodeName(nodeName)
	if err != nil {
		return nil, nil, err
	}
	awsInstance := newAWSInstance(c.ec2, instance)
	return awsInstance, instance, err
}
func setNodeDisk(nodeDiskMap map[types.NodeName]map[KubernetesVolumeID]bool, volumeID KubernetesVolumeID, nodeName types.NodeName, check bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeMap := nodeDiskMap[nodeName]
	if volumeMap == nil {
		volumeMap = make(map[KubernetesVolumeID]bool)
		nodeDiskMap[nodeName] = volumeMap
	}
	volumeMap[volumeID] = check
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
