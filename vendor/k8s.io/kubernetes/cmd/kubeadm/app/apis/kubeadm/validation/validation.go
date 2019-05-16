package validation

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"net"
	"net/url"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

func ValidateInitConfiguration(c *kubeadm.InitConfiguration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateNodeRegistrationOptions(&c.NodeRegistration, field.NewPath("nodeRegistration"))...)
	allErrs = append(allErrs, ValidateBootstrapTokens(c.BootstrapTokens, field.NewPath("bootstrapTokens"))...)
	allErrs = append(allErrs, ValidateClusterConfiguration(&c.ClusterConfiguration)...)
	allErrs = append(allErrs, ValidateAPIEndpoint(&c.LocalAPIEndpoint, field.NewPath("localAPIEndpoint"))...)
	return allErrs
}
func ValidateClusterConfiguration(c *kubeadm.ClusterConfiguration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateNetworking(&c.Networking, field.NewPath("networking"))...)
	allErrs = append(allErrs, ValidateAPIServer(&c.APIServer, field.NewPath("apiServer"))...)
	allErrs = append(allErrs, ValidateAbsolutePath(c.CertificatesDir, field.NewPath("certificatesDir"))...)
	allErrs = append(allErrs, ValidateFeatureGates(c.FeatureGates, field.NewPath("featureGates"))...)
	allErrs = append(allErrs, ValidateHostPort(c.ControlPlaneEndpoint, field.NewPath("controlPlaneEndpoint"))...)
	allErrs = append(allErrs, ValidateEtcd(&c.Etcd, field.NewPath("etcd"))...)
	allErrs = append(allErrs, componentconfigs.Known.Validate(c)...)
	return allErrs
}
func ValidateAPIServer(a *kubeadm.APIServer, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateCertSANs(a.CertSANs, fldPath.Child("certSANs"))...)
	return allErrs
}
func ValidateJoinConfiguration(c *kubeadm.JoinConfiguration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateDiscovery(&c.Discovery, field.NewPath("discovery"))...)
	allErrs = append(allErrs, ValidateNodeRegistrationOptions(&c.NodeRegistration, field.NewPath("nodeRegistration"))...)
	allErrs = append(allErrs, ValidateJoinControlPlane(c.ControlPlane, field.NewPath("controlPlane"))...)
	if !filepath.IsAbs(c.CACertPath) || !strings.HasSuffix(c.CACertPath, ".crt") {
		allErrs = append(allErrs, field.Invalid(field.NewPath("caCertPath"), c.CACertPath, "the ca certificate path must be an absolute path"))
	}
	return allErrs
}
func ValidateJoinControlPlane(c *kubeadm.JoinControlPlane, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if c != nil {
		allErrs = append(allErrs, ValidateAPIEndpoint(&c.LocalAPIEndpoint, fldPath.Child("localAPIEndpoint"))...)
	}
	return allErrs
}
func ValidateNodeRegistrationOptions(nro *kubeadm.NodeRegistrationOptions, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(nro.Name) == 0 {
		allErrs = append(allErrs, field.Required(fldPath, "--node-name or .nodeRegistration.name in the config file is a required value. It seems like this value couldn't be automatically detected in your environment, please specify the desired value using the CLI or config file."))
	} else {
		allErrs = append(allErrs, apivalidation.ValidateDNS1123Subdomain(nro.Name, field.NewPath("name"))...)
	}
	allErrs = append(allErrs, ValidateSocketPath(nro.CRISocket, fldPath.Child("criSocket"))...)
	return allErrs
}
func ValidateDiscovery(d *kubeadm.Discovery, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if d.BootstrapToken == nil && d.File == nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "", "bootstrapToken or file must be set"))
	}
	if d.BootstrapToken != nil && d.File != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "", "bootstrapToken and file cannot both be set"))
	}
	if d.BootstrapToken != nil {
		allErrs = append(allErrs, ValidateDiscoveryBootstrapToken(d.BootstrapToken, fldPath.Child("bootstrapToken"))...)
		allErrs = append(allErrs, ValidateToken(d.TLSBootstrapToken, fldPath.Child("tlsBootstrapToken"))...)
	}
	if d.File != nil {
		allErrs = append(allErrs, ValidateDiscoveryFile(d.File, fldPath.Child("file"))...)
		if len(d.TLSBootstrapToken) != 0 {
			allErrs = append(allErrs, ValidateToken(d.TLSBootstrapToken, fldPath.Child("tlsBootstrapToken"))...)
		}
	}
	return allErrs
}
func ValidateDiscoveryBootstrapToken(b *kubeadm.BootstrapTokenDiscovery, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(b.APIServerEndpoint) == 0 {
		allErrs = append(allErrs, field.Required(fldPath, "APIServerEndpoint is not set"))
	}
	if len(b.CACertHashes) == 0 && !b.UnsafeSkipCAVerification {
		allErrs = append(allErrs, field.Invalid(fldPath, "", "using token-based discovery without caCertHashes can be unsafe. Set unsafeSkipCAVerification to continue"))
	}
	allErrs = append(allErrs, ValidateToken(b.Token, fldPath.Child("token"))...)
	allErrs = append(allErrs, ValidateDiscoveryTokenAPIServer(b.APIServerEndpoint, fldPath.Child("apiServerEndpoints"))...)
	return allErrs
}
func ValidateDiscoveryFile(f *kubeadm.FileDiscovery, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateDiscoveryKubeConfigPath(f.KubeConfigPath, fldPath.Child("kubeConfigPath"))...)
	return allErrs
}
func ValidateDiscoveryTokenAPIServer(apiServer string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	_, _, err := net.SplitHostPort(apiServer)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, apiServer, err.Error()))
	}
	return allErrs
}
func ValidateDiscoveryKubeConfigPath(discoveryFile string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	u, err := url.Parse(discoveryFile)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, discoveryFile, "not a valid HTTPS URL or a file on disk"))
		return allErrs
	}
	if u.Scheme == "" {
		if _, err := os.Stat(discoveryFile); os.IsNotExist(err) {
			allErrs = append(allErrs, field.Invalid(fldPath, discoveryFile, "not a valid HTTPS URL or a file on disk"))
		}
		return allErrs
	}
	if u.Scheme != "https" {
		allErrs = append(allErrs, field.Invalid(fldPath, discoveryFile, "if a URL is used, the scheme must be https"))
	}
	return allErrs
}
func ValidateBootstrapTokens(bts []kubeadm.BootstrapToken, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for i, bt := range bts {
		btPath := fldPath.Child(fmt.Sprintf("%d", i))
		allErrs = append(allErrs, ValidateToken(bt.Token.String(), btPath.Child("token"))...)
		allErrs = append(allErrs, ValidateTokenUsages(bt.Usages, btPath.Child("usages"))...)
		allErrs = append(allErrs, ValidateTokenGroups(bt.Usages, bt.Groups, btPath.Child("groups"))...)
		if bt.Expires != nil && bt.TTL != nil {
			allErrs = append(allErrs, field.Invalid(btPath, "", "the BootstrapToken .TTL and .Expires fields are mutually exclusive"))
		}
	}
	return allErrs
}
func ValidateToken(token string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if !bootstraputil.IsValidBootstrapToken(token) {
		allErrs = append(allErrs, field.Invalid(fldPath, token, "the bootstrap token is invalid"))
	}
	return allErrs
}
func ValidateTokenGroups(usages []string, groups []string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	usagesSet := sets.NewString(usages...)
	usageAuthentication := strings.TrimPrefix(bootstrapapi.BootstrapTokenUsageAuthentication, bootstrapapi.BootstrapTokenUsagePrefix)
	if len(groups) > 0 && !usagesSet.Has(usageAuthentication) {
		allErrs = append(allErrs, field.Invalid(fldPath, groups, fmt.Sprintf("token groups cannot be specified unless --usages includes %q", usageAuthentication)))
	}
	for _, group := range groups {
		if err := bootstraputil.ValidateBootstrapGroupName(group); err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath, groups, err.Error()))
		}
	}
	return allErrs
}
func ValidateTokenUsages(usages []string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if err := bootstraputil.ValidateUsages(usages); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, usages, err.Error()))
	}
	return allErrs
}
func ValidateEtcd(e *kubeadm.Etcd, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	localPath := fldPath.Child("local")
	externalPath := fldPath.Child("external")
	if e.Local == nil && e.External == nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "", "either .Etcd.Local or .Etcd.External is required"))
		return allErrs
	}
	if e.Local != nil && e.External != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "", ".Etcd.Local and .Etcd.External are mutually exclusive"))
		return allErrs
	}
	if e.Local != nil {
		allErrs = append(allErrs, ValidateAbsolutePath(e.Local.DataDir, localPath.Child("dataDir"))...)
		allErrs = append(allErrs, ValidateCertSANs(e.Local.ServerCertSANs, localPath.Child("serverCertSANs"))...)
		allErrs = append(allErrs, ValidateCertSANs(e.Local.PeerCertSANs, localPath.Child("peerCertSANs"))...)
	}
	if e.External != nil {
		requireHTTPS := true
		if e.External.CAFile == "" && e.External.CertFile == "" && e.External.KeyFile == "" {
			requireHTTPS = false
		}
		if (e.External.CertFile == "" && e.External.KeyFile != "") || (e.External.CertFile != "" && e.External.KeyFile == "") {
			allErrs = append(allErrs, field.Invalid(externalPath, "", "either both or none of .Etcd.External.CertFile and .Etcd.External.KeyFile must be set"))
		}
		if e.External.CertFile != "" && e.External.KeyFile != "" && e.External.CAFile == "" {
			allErrs = append(allErrs, field.Invalid(externalPath, "", "setting .Etcd.External.CertFile and .Etcd.External.KeyFile requires .Etcd.External.CAFile"))
		}
		allErrs = append(allErrs, ValidateURLs(e.External.Endpoints, requireHTTPS, externalPath.Child("endpoints"))...)
		if e.External.CAFile != "" {
			allErrs = append(allErrs, ValidateAbsolutePath(e.External.CAFile, externalPath.Child("caFile"))...)
		}
		if e.External.CertFile != "" {
			allErrs = append(allErrs, ValidateAbsolutePath(e.External.CertFile, externalPath.Child("certFile"))...)
		}
		if e.External.KeyFile != "" {
			allErrs = append(allErrs, ValidateAbsolutePath(e.External.KeyFile, externalPath.Child("keyFile"))...)
		}
	}
	return allErrs
}
func ValidateCertSANs(altnames []string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for _, altname := range altnames {
		if len(validation.IsDNS1123Subdomain(altname)) != 0 && net.ParseIP(altname) == nil {
			allErrs = append(allErrs, field.Invalid(fldPath, altname, "altname is not a valid dns label or ip address"))
		}
	}
	return allErrs
}
func ValidateURLs(urls []string, requireHTTPS bool, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for _, urlstr := range urls {
		u, err := url.Parse(urlstr)
		if err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath, urlstr, fmt.Sprintf("URL parse error: %v", err)))
			continue
		}
		if requireHTTPS && u.Scheme != "https" {
			allErrs = append(allErrs, field.Invalid(fldPath, urlstr, "the URL must be using the HTTPS scheme"))
		}
		if u.Scheme == "" {
			allErrs = append(allErrs, field.Invalid(fldPath, urlstr, "the URL without scheme is not allowed"))
		}
	}
	return allErrs
}
func ValidateIPFromString(ipaddr string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if net.ParseIP(ipaddr) == nil {
		allErrs = append(allErrs, field.Invalid(fldPath, ipaddr, "ip address is not valid"))
	}
	return allErrs
}
func ValidatePort(port int32, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if _, err := kubeadmutil.ParsePort(strconv.Itoa(int(port))); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, port, "port number is not valid"))
	}
	return allErrs
}
func ValidateHostPort(endpoint string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if _, _, err := kubeadmutil.ParseHostPort(endpoint); endpoint != "" && err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, endpoint, "endpoint is not valid"))
	}
	return allErrs
}
func ValidateIPNetFromString(subnet string, minAddrs int64, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	_, svcSubnet, err := net.ParseCIDR(subnet)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, subnet, "couldn't parse subnet"))
		return allErrs
	}
	numAddresses := ipallocator.RangeSize(svcSubnet)
	if numAddresses < minAddrs {
		allErrs = append(allErrs, field.Invalid(fldPath, subnet, "subnet is too small"))
	}
	return allErrs
}
func ValidateNetworking(c *kubeadm.Networking, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateDNS1123Subdomain(c.DNSDomain, field.NewPath("dnsDomain"))...)
	allErrs = append(allErrs, ValidateIPNetFromString(c.ServiceSubnet, constants.MinimumAddressesInServiceSubnet, field.NewPath("serviceSubnet"))...)
	if len(c.PodSubnet) != 0 {
		allErrs = append(allErrs, ValidateIPNetFromString(c.PodSubnet, constants.MinimumAddressesInServiceSubnet, field.NewPath("podSubnet"))...)
	}
	return allErrs
}
func ValidateAbsolutePath(path string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if !filepath.IsAbs(path) {
		allErrs = append(allErrs, field.Invalid(fldPath, path, "path is not absolute"))
	}
	return allErrs
}
func ValidateMixedArguments(flag *pflag.FlagSet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !flag.Changed("config") {
		return nil
	}
	mixedInvalidFlags := []string{}
	flag.Visit(func(f *pflag.Flag) {
		if f.Name == "config" || f.Name == "ignore-preflight-errors" || strings.HasPrefix(f.Name, "skip-") || f.Name == "dry-run" || f.Name == "kubeconfig" || f.Name == "v" || f.Name == "rootfs" || f.Name == "print-join-command" || f.Name == "node-name" || f.Name == "cri-socket" {
			return
		}
		mixedInvalidFlags = append(mixedInvalidFlags, f.Name)
	})
	if len(mixedInvalidFlags) != 0 {
		return errors.Errorf("can not mix '--config' with arguments %v", mixedInvalidFlags)
	}
	return nil
}
func ValidateFeatureGates(featureGates map[string]bool, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for k := range featureGates {
		if !features.Supports(features.InitFeatureGates, k) {
			allErrs = append(allErrs, field.Invalid(fldPath, featureGates, fmt.Sprintf("%s is not a valid feature name.", k)))
		}
	}
	return allErrs
}
func ValidateAPIEndpoint(c *kubeadm.APIEndpoint, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateIPFromString(c.AdvertiseAddress, fldPath.Child("advertiseAddress"))...)
	allErrs = append(allErrs, ValidatePort(c.BindPort, fldPath.Child("bindPort"))...)
	return allErrs
}
func ValidateIgnorePreflightErrors(ignorePreflightErrors []string) (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ignoreErrors := sets.NewString()
	allErrs := field.ErrorList{}
	for _, item := range ignorePreflightErrors {
		ignoreErrors.Insert(strings.ToLower(item))
	}
	if ignoreErrors.Has("all") && ignoreErrors.Len() > 1 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("ignore-preflight-errors"), strings.Join(ignoreErrors.List(), ","), "don't specify individual checks if 'all' is used"))
	}
	return ignoreErrors, allErrs.ToAggregate()
}
func ValidateSocketPath(socket string, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	u, err := url.Parse(socket)
	if err != nil {
		return append(allErrs, field.Invalid(fldPath, socket, fmt.Sprintf("URL parsing error: %v", err)))
	}
	if u.Scheme == "" {
		if !filepath.IsAbs(u.Path) {
			return append(allErrs, field.Invalid(fldPath, socket, fmt.Sprintf("path is not absolute: %s", socket)))
		}
	} else if u.Scheme != kubeadmapiv1beta1.DefaultUrlScheme {
		return append(allErrs, field.Invalid(fldPath, socket, fmt.Sprintf("URL scheme %s is not supported", u.Scheme)))
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
