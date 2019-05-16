package options

import (
	utilnet "k8s.io/apimachinery/pkg/util/net"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	api "k8s.io/kubernetes/pkg/apis/core"
	_ "k8s.io/kubernetes/pkg/features"
	kubeoptions "k8s.io/kubernetes/pkg/kubeapiserver/options"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master/ports"
	"k8s.io/kubernetes/pkg/master/reconcilers"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"net"
	"strings"
	"time"
)

type ServerRunOptions struct {
	GenericServerRunOptions          *genericoptions.ServerRunOptions
	Etcd                             *genericoptions.EtcdOptions
	SecureServing                    *genericoptions.SecureServingOptionsWithLoopback
	InsecureServing                  *genericoptions.DeprecatedInsecureServingOptionsWithLoopback
	Audit                            *genericoptions.AuditOptions
	Features                         *genericoptions.FeatureOptions
	Admission                        *kubeoptions.AdmissionOptions
	Authentication                   *kubeoptions.BuiltInAuthenticationOptions
	Authorization                    *kubeoptions.BuiltInAuthorizationOptions
	CloudProvider                    *kubeoptions.CloudProviderOptions
	StorageSerialization             *kubeoptions.StorageSerializationOptions
	APIEnablement                    *genericoptions.APIEnablementOptions
	AllowPrivileged                  bool
	EnableLogsHandler                bool
	EventTTL                         time.Duration
	KubeletConfig                    kubeletclient.KubeletClientConfig
	KubernetesServiceNodePort        int
	MaxConnectionBytesPerSec         int64
	ServiceClusterIPRange            net.IPNet
	ServiceNodePortRange             utilnet.PortRange
	SSHKeyfile                       string
	SSHUser                          string
	ProxyClientCertFile              string
	ProxyClientKeyFile               string
	EnableAggregatorRouting          bool
	MasterCount                      int
	EndpointReconcilerType           string
	ServiceAccountSigningKeyFile     string
	ServiceAccountIssuer             serviceaccount.TokenGenerator
	ServiceAccountTokenMaxExpiration time.Duration
}

func NewServerRunOptions() *ServerRunOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := ServerRunOptions{GenericServerRunOptions: genericoptions.NewServerRunOptions(), Etcd: genericoptions.NewEtcdOptions(storagebackend.NewDefaultConfig(kubeoptions.DefaultEtcdPathPrefix, nil)), SecureServing: kubeoptions.NewSecureServingOptions(), InsecureServing: kubeoptions.NewInsecureServingOptions(), Audit: genericoptions.NewAuditOptions(), Features: genericoptions.NewFeatureOptions(), Admission: kubeoptions.NewAdmissionOptions(), Authentication: kubeoptions.NewBuiltInAuthenticationOptions().WithAll(), Authorization: kubeoptions.NewBuiltInAuthorizationOptions(), CloudProvider: kubeoptions.NewCloudProviderOptions(), StorageSerialization: kubeoptions.NewStorageSerializationOptions(), APIEnablement: genericoptions.NewAPIEnablementOptions(), EnableLogsHandler: true, EventTTL: 1 * time.Hour, MasterCount: 1, EndpointReconcilerType: string(reconcilers.LeaseEndpointReconcilerType), KubeletConfig: kubeletclient.KubeletClientConfig{Port: ports.KubeletPort, ReadOnlyPort: ports.KubeletReadOnlyPort, PreferredAddressTypes: []string{string(api.NodeHostName), string(api.NodeInternalDNS), string(api.NodeInternalIP), string(api.NodeExternalDNS), string(api.NodeExternalIP)}, EnableHttps: true, HTTPTimeout: time.Duration(5) * time.Second}, ServiceNodePortRange: kubeoptions.DefaultServiceNodePortRange}
	s.ServiceClusterIPRange = kubeoptions.DefaultServiceIPCIDR
	s.Etcd.DefaultStorageMediaType = "application/vnd.kubernetes.protobuf"
	return &s
}
func (s *ServerRunOptions) Flags() (fss apiserverflag.NamedFlagSets) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.GenericServerRunOptions.AddUniversalFlags(fss.FlagSet("generic"))
	s.Etcd.AddFlags(fss.FlagSet("etcd"))
	s.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	s.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	s.InsecureServing.AddUnqualifiedFlags(fss.FlagSet("insecure serving"))
	s.Audit.AddFlags(fss.FlagSet("auditing"))
	s.Features.AddFlags(fss.FlagSet("features"))
	s.Authentication.AddFlags(fss.FlagSet("authentication"))
	s.Authorization.AddFlags(fss.FlagSet("authorization"))
	s.CloudProvider.AddFlags(fss.FlagSet("cloud provider"))
	s.StorageSerialization.AddFlags(fss.FlagSet("storage"))
	s.APIEnablement.AddFlags(fss.FlagSet("api enablement"))
	s.Admission.AddFlags(fss.FlagSet("admission"))
	fs := fss.FlagSet("misc")
	fs.DurationVar(&s.EventTTL, "event-ttl", s.EventTTL, "Amount of time to retain events.")
	fs.BoolVar(&s.AllowPrivileged, "allow-privileged", s.AllowPrivileged, "If true, allow privileged containers. [default=false]")
	fs.BoolVar(&s.EnableLogsHandler, "enable-logs-handler", s.EnableLogsHandler, "If true, install a /logs handler for the apiserver logs.")
	fs.StringVar(&s.SSHUser, "ssh-user", s.SSHUser, "If non-empty, use secure SSH proxy to the nodes, using this user name")
	fs.MarkDeprecated("ssh-user", "This flag will be removed in a future version.")
	fs.StringVar(&s.SSHKeyfile, "ssh-keyfile", s.SSHKeyfile, "If non-empty, use secure SSH proxy to the nodes, using this user keyfile")
	fs.MarkDeprecated("ssh-keyfile", "This flag will be removed in a future version.")
	fs.Int64Var(&s.MaxConnectionBytesPerSec, "max-connection-bytes-per-sec", s.MaxConnectionBytesPerSec, ""+"If non-zero, throttle each user connection to this number of bytes/sec. "+"Currently only applies to long-running requests.")
	fs.IntVar(&s.MasterCount, "apiserver-count", s.MasterCount, "The number of apiservers running in the cluster, must be a positive number. (In use when --endpoint-reconciler-type=master-count is enabled.)")
	fs.StringVar(&s.EndpointReconcilerType, "endpoint-reconciler-type", string(s.EndpointReconcilerType), "Use an endpoint reconciler ("+strings.Join(reconcilers.AllTypes.Names(), ", ")+")")
	fs.IntVar(&s.KubernetesServiceNodePort, "kubernetes-service-node-port", s.KubernetesServiceNodePort, ""+"If non-zero, the Kubernetes master service (which apiserver creates/maintains) will be "+"of type NodePort, using this as the value of the port. If zero, the Kubernetes master "+"service will be of type ClusterIP.")
	fs.IPNetVar(&s.ServiceClusterIPRange, "service-cluster-ip-range", s.ServiceClusterIPRange, ""+"A CIDR notation IP range from which to assign service cluster IPs. This must not "+"overlap with any IP ranges assigned to nodes for pods.")
	fs.Var(&s.ServiceNodePortRange, "service-node-port-range", ""+"A port range to reserve for services with NodePort visibility. "+"Example: '30000-32767'. Inclusive at both ends of the range.")
	fs.BoolVar(&s.KubeletConfig.EnableHttps, "kubelet-https", s.KubeletConfig.EnableHttps, "Use https for kubelet connections.")
	fs.StringSliceVar(&s.KubeletConfig.PreferredAddressTypes, "kubelet-preferred-address-types", s.KubeletConfig.PreferredAddressTypes, "List of the preferred NodeAddressTypes to use for kubelet connections.")
	fs.UintVar(&s.KubeletConfig.Port, "kubelet-port", s.KubeletConfig.Port, "DEPRECATED: kubelet port.")
	fs.MarkDeprecated("kubelet-port", "kubelet-port is deprecated and will be removed.")
	fs.UintVar(&s.KubeletConfig.ReadOnlyPort, "kubelet-read-only-port", s.KubeletConfig.ReadOnlyPort, "DEPRECATED: kubelet port.")
	fs.DurationVar(&s.KubeletConfig.HTTPTimeout, "kubelet-timeout", s.KubeletConfig.HTTPTimeout, "Timeout for kubelet operations.")
	fs.StringVar(&s.KubeletConfig.CertFile, "kubelet-client-certificate", s.KubeletConfig.CertFile, "Path to a client cert file for TLS.")
	fs.StringVar(&s.KubeletConfig.KeyFile, "kubelet-client-key", s.KubeletConfig.KeyFile, "Path to a client key file for TLS.")
	fs.StringVar(&s.KubeletConfig.CAFile, "kubelet-certificate-authority", s.KubeletConfig.CAFile, "Path to a cert file for the certificate authority.")
	repair := false
	fs.BoolVar(&repair, "repair-malformed-updates", false, "deprecated")
	fs.MarkDeprecated("repair-malformed-updates", "This flag will be removed in a future version")
	fs.StringVar(&s.ProxyClientCertFile, "proxy-client-cert-file", s.ProxyClientCertFile, ""+"Client certificate used to prove the identity of the aggregator or kube-apiserver "+"when it must call out during a request. This includes proxying requests to a user "+"api-server and calling out to webhook admission plugins. It is expected that this "+"cert includes a signature from the CA in the --requestheader-client-ca-file flag. "+"That CA is published in the 'extension-apiserver-authentication' configmap in "+"the kube-system namespace. Components receiving calls from kube-aggregator should "+"use that CA to perform their half of the mutual TLS verification.")
	fs.StringVar(&s.ProxyClientKeyFile, "proxy-client-key-file", s.ProxyClientKeyFile, ""+"Private key for the client certificate used to prove the identity of the aggregator or kube-apiserver "+"when it must call out during a request. This includes proxying requests to a user "+"api-server and calling out to webhook admission plugins.")
	fs.BoolVar(&s.EnableAggregatorRouting, "enable-aggregator-routing", s.EnableAggregatorRouting, "Turns on aggregator routing requests to endpoints IP rather than cluster IP.")
	fs.StringVar(&s.ServiceAccountSigningKeyFile, "service-account-signing-key-file", s.ServiceAccountSigningKeyFile, ""+"Path to the file that contains the current private key of the service account token issuer. The issuer will sign issued ID tokens with this private key. (Requires the 'TokenRequest' feature gate.)")
	return fss
}
