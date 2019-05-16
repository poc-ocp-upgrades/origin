package config

import (
	"crypto/x509"
	"fmt"
	goformat "fmt"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"time"
	gotime "time"
)

func ParseNamespaceAndName(in string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(in) == 0 {
		return "", "", nil
	}
	tokens := strings.Split(in, "/")
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("expected input in the form <namespace>/<resource-name>, not: %v", in)
	}
	return tokens[0], tokens[1], nil
}
func RelativizeMasterConfigPaths(config *MasterConfig, base string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cmdutil.RelativizePathWithNoBacksteps(GetMasterFileReferences(config), base)
}
func ResolveMasterConfigPaths(config *MasterConfig, base string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cmdutil.ResolvePaths(GetMasterFileReferences(config), base)
}
func GetMasterFileReferences(config *MasterConfig) []*string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	refs := []*string{}
	refs = append(refs, &config.ServingInfo.ServerCert.CertFile)
	refs = append(refs, &config.ServingInfo.ServerCert.KeyFile)
	refs = append(refs, &config.ServingInfo.ClientCA)
	for i := range config.ServingInfo.NamedCertificates {
		refs = append(refs, &config.ServingInfo.NamedCertificates[i].CertFile)
		refs = append(refs, &config.ServingInfo.NamedCertificates[i].KeyFile)
	}
	refs = append(refs, &config.EtcdClientInfo.ClientCert.CertFile)
	refs = append(refs, &config.EtcdClientInfo.ClientCert.KeyFile)
	refs = append(refs, &config.EtcdClientInfo.CA)
	refs = append(refs, &config.KubeletClientInfo.ClientCert.CertFile)
	refs = append(refs, &config.KubeletClientInfo.ClientCert.KeyFile)
	refs = append(refs, &config.KubeletClientInfo.CA)
	if config.EtcdConfig != nil {
		refs = append(refs, &config.EtcdConfig.ServingInfo.ServerCert.CertFile)
		refs = append(refs, &config.EtcdConfig.ServingInfo.ServerCert.KeyFile)
		refs = append(refs, &config.EtcdConfig.ServingInfo.ClientCA)
		for i := range config.EtcdConfig.ServingInfo.NamedCertificates {
			refs = append(refs, &config.EtcdConfig.ServingInfo.NamedCertificates[i].CertFile)
			refs = append(refs, &config.EtcdConfig.ServingInfo.NamedCertificates[i].KeyFile)
		}
		refs = append(refs, &config.EtcdConfig.PeerServingInfo.ServerCert.CertFile)
		refs = append(refs, &config.EtcdConfig.PeerServingInfo.ServerCert.KeyFile)
		refs = append(refs, &config.EtcdConfig.PeerServingInfo.ClientCA)
		for i := range config.EtcdConfig.PeerServingInfo.NamedCertificates {
			refs = append(refs, &config.EtcdConfig.PeerServingInfo.NamedCertificates[i].CertFile)
			refs = append(refs, &config.EtcdConfig.PeerServingInfo.NamedCertificates[i].KeyFile)
		}
		refs = append(refs, &config.EtcdConfig.StorageDir)
	}
	if config.OAuthConfig != nil {
		if config.OAuthConfig.MasterCA != nil {
			refs = append(refs, config.OAuthConfig.MasterCA)
		}
		if config.OAuthConfig.SessionConfig != nil {
			refs = append(refs, &config.OAuthConfig.SessionConfig.SessionSecretsFile)
		}
		for _, identityProvider := range config.OAuthConfig.IdentityProviders {
			switch provider := identityProvider.Provider.(type) {
			case *RequestHeaderIdentityProvider:
				refs = append(refs, &provider.ClientCA)
			case *HTPasswdPasswordIdentityProvider:
				refs = append(refs, &provider.File)
			case *LDAPPasswordIdentityProvider:
				refs = append(refs, &provider.CA)
				refs = append(refs, GetStringSourceFileReferences(&provider.BindPassword)...)
			case *BasicAuthPasswordIdentityProvider:
				refs = append(refs, &provider.RemoteConnectionInfo.CA)
				refs = append(refs, &provider.RemoteConnectionInfo.ClientCert.CertFile)
				refs = append(refs, &provider.RemoteConnectionInfo.ClientCert.KeyFile)
			case *KeystonePasswordIdentityProvider:
				refs = append(refs, &provider.RemoteConnectionInfo.CA)
				refs = append(refs, &provider.RemoteConnectionInfo.ClientCert.CertFile)
				refs = append(refs, &provider.RemoteConnectionInfo.ClientCert.KeyFile)
			case *GitLabIdentityProvider:
				refs = append(refs, &provider.CA)
				refs = append(refs, GetStringSourceFileReferences(&provider.ClientSecret)...)
			case *OpenIDIdentityProvider:
				refs = append(refs, &provider.CA)
				refs = append(refs, GetStringSourceFileReferences(&provider.ClientSecret)...)
			case *GoogleIdentityProvider:
				refs = append(refs, GetStringSourceFileReferences(&provider.ClientSecret)...)
			case *GitHubIdentityProvider:
				refs = append(refs, GetStringSourceFileReferences(&provider.ClientSecret)...)
				refs = append(refs, &provider.CA)
			}
		}
		if config.OAuthConfig.Templates != nil {
			refs = append(refs, &config.OAuthConfig.Templates.Login)
			refs = append(refs, &config.OAuthConfig.Templates.ProviderSelection)
			refs = append(refs, &config.OAuthConfig.Templates.Error)
		}
	}
	for k := range config.AdmissionConfig.PluginConfig {
		refs = append(refs, &config.AdmissionConfig.PluginConfig[k].Location)
	}
	refs = append(refs, &config.KubernetesMasterConfig.SchedulerConfigFile)
	refs = append(refs, &config.KubernetesMasterConfig.ProxyClientInfo.CertFile)
	refs = append(refs, &config.KubernetesMasterConfig.ProxyClientInfo.KeyFile)
	refs = appendFlagsWithFileExtensions(refs, config.KubernetesMasterConfig.APIServerArguments)
	refs = appendFlagsWithFileExtensions(refs, config.KubernetesMasterConfig.SchedulerArguments)
	refs = appendFlagsWithFileExtensions(refs, config.KubernetesMasterConfig.ControllerArguments)
	if config.AuthConfig.RequestHeader != nil {
		refs = append(refs, &config.AuthConfig.RequestHeader.ClientCA)
	}
	for k := range config.AuthConfig.WebhookTokenAuthenticators {
		refs = append(refs, &config.AuthConfig.WebhookTokenAuthenticators[k].ConfigFile)
	}
	if len(config.AuthConfig.OAuthMetadataFile) > 0 {
		refs = append(refs, &config.AuthConfig.OAuthMetadataFile)
	}
	refs = append(refs, &config.AggregatorConfig.ProxyClientInfo.CertFile)
	refs = append(refs, &config.AggregatorConfig.ProxyClientInfo.KeyFile)
	refs = append(refs, &config.ServiceAccountConfig.MasterCA)
	refs = append(refs, &config.ServiceAccountConfig.PrivateKeyFile)
	for i := range config.ServiceAccountConfig.PublicKeyFiles {
		refs = append(refs, &config.ServiceAccountConfig.PublicKeyFiles[i])
	}
	refs = append(refs, &config.MasterClients.OpenShiftLoopbackKubeConfig)
	if config.ControllerConfig.ServiceServingCert.Signer != nil {
		refs = append(refs, &config.ControllerConfig.ServiceServingCert.Signer.CertFile)
		refs = append(refs, &config.ControllerConfig.ServiceServingCert.Signer.KeyFile)
	}
	refs = append(refs, &config.AuditConfig.AuditFilePath)
	refs = append(refs, &config.AuditConfig.PolicyFile)
	refs = append(refs, &config.ImagePolicyConfig.AdditionalTrustedCA)
	return refs
}
func appendFlagsWithFileExtensions(refs []*string, args ExtendedArguments) []*string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for key, s := range args {
		if len(s) == 0 {
			continue
		}
		if !strings.HasSuffix(key, "-file") && !strings.HasSuffix(key, "-dir") {
			continue
		}
		for i := range s {
			refs = append(refs, &s[i])
		}
	}
	return refs
}
func RelativizeNodeConfigPaths(config *NodeConfig, base string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cmdutil.RelativizePathWithNoBacksteps(GetNodeFileReferences(config), base)
}
func ResolveNodeConfigPaths(config *NodeConfig, base string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cmdutil.ResolvePaths(GetNodeFileReferences(config), base)
}
func GetNodeFileReferences(config *NodeConfig) []*string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	refs := []*string{}
	refs = append(refs, &config.ServingInfo.ServerCert.CertFile)
	refs = append(refs, &config.ServingInfo.ServerCert.KeyFile)
	refs = append(refs, &config.ServingInfo.ClientCA)
	for i := range config.ServingInfo.NamedCertificates {
		refs = append(refs, &config.ServingInfo.NamedCertificates[i].CertFile)
		refs = append(refs, &config.ServingInfo.NamedCertificates[i].KeyFile)
	}
	refs = append(refs, &config.DNSRecursiveResolvConf)
	refs = append(refs, &config.MasterKubeConfig)
	refs = append(refs, &config.VolumeDirectory)
	if config.PodManifestConfig != nil {
		refs = append(refs, &config.PodManifestConfig.Path)
	}
	refs = appendFlagsWithFileExtensions(refs, config.KubeletArguments)
	return refs
}
func SetProtobufClientDefaults(overrides *ClientConnectionOverrides) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	overrides.AcceptContentTypes = "application/vnd.kubernetes.protobuf,application/json"
	overrides.ContentType = "application/vnd.kubernetes.protobuf"
	overrides.QPS *= 2
	overrides.Burst *= 2
}
func GetKubeConfigOrInClusterConfig(kubeConfigFile string, overrides *ClientConnectionOverrides) (*restclient.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(kubeConfigFile) > 0 {
		return GetClientConfig(kubeConfigFile, overrides)
	}
	clientConfig, err := restclient.InClusterConfig()
	if err != nil {
		return nil, err
	}
	applyClientConnectionOverrides(overrides, clientConfig)
	clientConfig.WrapTransport = DefaultClientTransport
	return clientConfig, nil
}
func GetClientConfig(kubeConfigFile string, overrides *ClientConnectionOverrides) (*restclient.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigBytes, err := ioutil.ReadFile(kubeConfigFile)
	if err != nil {
		return nil, err
	}
	kubeConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		return nil, err
	}
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	applyClientConnectionOverrides(overrides, clientConfig)
	clientConfig.WrapTransport = DefaultClientTransport
	return clientConfig, nil
}
func applyClientConnectionOverrides(overrides *ClientConnectionOverrides, kubeConfig *restclient.Config) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if overrides == nil {
		return
	}
	kubeConfig.QPS = overrides.QPS
	kubeConfig.Burst = int(overrides.Burst)
	kubeConfig.ContentConfig.AcceptContentTypes = overrides.AcceptContentTypes
	kubeConfig.ContentConfig.ContentType = overrides.ContentType
}
func DefaultClientTransport(rt http.RoundTripper) http.RoundTripper {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	transport, ok := rt.(*http.Transport)
	if !ok {
		return rt
	}
	dialer := &net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}
	transport.Dial = dialer.Dial
	transport.MaxIdleConnsPerHost = 100
	return transport
}
func GetOAuthClientCertCAs(options MasterConfig) ([]*x509.Certificate, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allCerts := []*x509.Certificate{}
	if options.OAuthConfig != nil {
		for _, identityProvider := range options.OAuthConfig.IdentityProviders {
			switch provider := identityProvider.Provider.(type) {
			case *RequestHeaderIdentityProvider:
				caFile := provider.ClientCA
				if len(caFile) == 0 {
					continue
				}
				certs, err := cmdutil.CertificatesFromFile(caFile)
				if err != nil {
					return nil, fmt.Errorf("Error reading %s: %s", caFile, err)
				}
				allCerts = append(allCerts, certs...)
			}
		}
	}
	return allCerts, nil
}
func IsPasswordAuthenticator(provider IdentityProvider) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch provider.Provider.(type) {
	case *BasicAuthPasswordIdentityProvider, *AllowAllPasswordIdentityProvider, *DenyAllPasswordIdentityProvider, *HTPasswdPasswordIdentityProvider, *LDAPPasswordIdentityProvider, *KeystonePasswordIdentityProvider:
		return true
	}
	return false
}
func IsIdentityProviderType(provider runtime.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch provider.(type) {
	case *RequestHeaderIdentityProvider, *BasicAuthPasswordIdentityProvider, *AllowAllPasswordIdentityProvider, *DenyAllPasswordIdentityProvider, *HTPasswdPasswordIdentityProvider, *LDAPPasswordIdentityProvider, *KeystonePasswordIdentityProvider, *OpenIDIdentityProvider, *GitHubIdentityProvider, *GitLabIdentityProvider, *GoogleIdentityProvider:
		return true
	}
	return false
}

const kubeAPIEnablementFlag = "runtime-config"

func GetKubeAPIServerFlagAPIEnablement(flagValue []string) map[schema.GroupVersion]bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versions := map[schema.GroupVersion]bool{}
	for _, val := range flagValue {
		if strings.HasPrefix(val, "api/") {
			continue
		}
		val = strings.TrimPrefix(val, "apis/")
		tokens := strings.Split(val, "=")
		if len(tokens) != 2 {
			continue
		}
		gv, err := schema.ParseGroupVersion(tokens[0])
		if err != nil {
			continue
		}
		enabled, _ := strconv.ParseBool(tokens[1])
		versions[gv] = enabled
	}
	return versions
}
func GetEnabledAPIVersionsForGroup(config KubernetesMasterConfig, apiGroup string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allowedVersions := KubeAPIGroupsToAllowedVersions[apiGroup]
	blacklist := sets.NewString(config.DisabledAPIGroupVersions[apiGroup]...)
	if blacklist.Has(AllVersions) {
		return []string{}
	}
	flagVersions := GetKubeAPIServerFlagAPIEnablement(config.APIServerArguments[kubeAPIEnablementFlag])
	enabledVersions := sets.String{}
	for _, currVersion := range allowedVersions {
		if blacklist.Has(currVersion) {
			continue
		}
		gv := schema.GroupVersion{Group: apiGroup, Version: currVersion}
		if enabled, ok := flagVersions[gv]; ok && !enabled {
			continue
		}
		enabledVersions.Insert(currVersion)
	}
	for currVersion, enabled := range flagVersions {
		if !enabled {
			continue
		}
		if blacklist.Has(currVersion.Version) {
			continue
		}
		if currVersion.Group != apiGroup {
			continue
		}
		enabledVersions.Insert(currVersion.Version)
	}
	return enabledVersions.List()
}
func GetDisabledAPIVersionsForGroup(config KubernetesMasterConfig, apiGroup string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allowedVersions := sets.NewString(KubeAPIGroupsToAllowedVersions[apiGroup]...)
	enabledVersions := sets.NewString(GetEnabledAPIVersionsForGroup(config, apiGroup)...)
	disabledVersions := allowedVersions.Difference(enabledVersions)
	disabledVersions.Insert(config.DisabledAPIGroupVersions[apiGroup]...)
	flagVersions := GetKubeAPIServerFlagAPIEnablement(config.APIServerArguments[kubeAPIEnablementFlag])
	for currVersion, enabled := range flagVersions {
		if enabled {
			continue
		}
		if disabledVersions.Has(currVersion.Version) {
			continue
		}
		if currVersion.Group != apiGroup {
			continue
		}
		disabledVersions.Insert(currVersion.Version)
	}
	return disabledVersions.List()
}
func CIDRsOverlap(cidr1, cidr2 string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, ipNet1, err := net.ParseCIDR(cidr1)
	if err != nil {
		return false
	}
	_, ipNet2, err := net.ParseCIDR(cidr2)
	if err != nil {
		return false
	}
	return ipNet1.Contains(ipNet2.IP) || ipNet2.Contains(ipNet1.IP)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
