package configdefault

import (
	goformat "fmt"
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
	"github.com/openshift/library-go/pkg/config/configdefaults"
	"io/ioutil"
	"k8s.io/klog"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func ResolveDirectoriesForSATokenVerification(config *kubecontrolplanev1.KubeAPIServerConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resolvedSATokenValidationCerts := []string{}
	for _, filename := range config.ServiceAccountPublicKeyFiles {
		file, err := os.Open(filename)
		if err != nil {
			resolvedSATokenValidationCerts = append(resolvedSATokenValidationCerts, filename)
			klog.Warningf(err.Error())
			continue
		}
		fileInfo, err := file.Stat()
		if err != nil {
			resolvedSATokenValidationCerts = append(resolvedSATokenValidationCerts, filename)
			klog.Warningf(err.Error())
			continue
		}
		if !fileInfo.IsDir() {
			resolvedSATokenValidationCerts = append(resolvedSATokenValidationCerts, filename)
			continue
		}
		contents, err := ioutil.ReadDir(filename)
		switch {
		case os.IsNotExist(err) || os.IsPermission(err):
			klog.Warningf(err.Error())
		case err != nil:
			panic(err)
		default:
			for _, content := range contents {
				if !content.Mode().IsRegular() {
					continue
				}
				resolvedSATokenValidationCerts = append(resolvedSATokenValidationCerts, filepath.Join(filename, content.Name()))
			}
		}
	}
	config.ServiceAccountPublicKeyFiles = resolvedSATokenValidationCerts
}
func SetRecommendedKubeAPIServerConfigDefaults(config *kubecontrolplanev1.KubeAPIServerConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configdefaults.DefaultString(&config.GenericAPIServerConfig.StorageConfig.StoragePrefix, "kubernetes.io")
	configdefaults.DefaultString(&config.GenericAPIServerConfig.ServingInfo.BindAddress, "0.0.0.0:6443")
	configdefaults.SetRecommendedGenericAPIServerConfigDefaults(&config.GenericAPIServerConfig)
	SetRecommendedMasterAuthConfigDefaults(&config.AuthConfig)
	SetRecommendedAggregatorConfigDefaults(&config.AggregatorConfig)
	SetRecommendedKubeletConnectionInfoDefaults(&config.KubeletClientInfo)
	configdefaults.DefaultString(&config.ServicesSubnet, "10.0.0.0/24")
	configdefaults.DefaultString(&config.ServicesNodePortRange, "30000-32767")
	if len(config.ServiceAccountPublicKeyFiles) == 0 {
		config.ServiceAccountPublicKeyFiles = append([]string{}, "/etc/kubernetes/static-pod-resources/configmaps/sa-token-signing-certs")
	}
	if config.AuthConfig.RequestHeader == nil {
		config.AuthConfig.RequestHeader = &kubecontrolplanev1.RequestHeaderAuthenticationOptions{}
		configdefaults.DefaultStringSlice(&config.AuthConfig.RequestHeader.ClientCommonNames, []string{"system:openshift-aggregator"})
		configdefaults.DefaultString(&config.AuthConfig.RequestHeader.ClientCA, "/var/run/configmaps/aggregator-client-ca/ca-bundle.crt")
		configdefaults.DefaultStringSlice(&config.AuthConfig.RequestHeader.UsernameHeaders, []string{"X-Remote-User"})
		configdefaults.DefaultStringSlice(&config.AuthConfig.RequestHeader.GroupHeaders, []string{"X-Remote-Group"})
		configdefaults.DefaultStringSlice(&config.AuthConfig.RequestHeader.ExtraHeaderPrefixes, []string{"X-Remote-Extra-"})
	}
	for i := range config.AuthConfig.WebhookTokenAuthenticators {
		if len(config.AuthConfig.WebhookTokenAuthenticators[i].CacheTTL) == 0 {
			config.AuthConfig.WebhookTokenAuthenticators[i].CacheTTL = "2m"
		}
	}
	if config.OAuthConfig != nil {
		for i := range config.OAuthConfig.IdentityProviders {
			configdefaults.DefaultString(&config.OAuthConfig.IdentityProviders[i].MappingMethod, "claim")
		}
	}
}
func SetRecommendedMasterAuthConfigDefaults(config *kubecontrolplanev1.MasterAuthConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func SetRecommendedAggregatorConfigDefaults(config *kubecontrolplanev1.AggregatorConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configdefaults.DefaultString(&config.ProxyClientInfo.KeyFile, "/var/run/secrets/aggregator-client/tls.key")
	configdefaults.DefaultString(&config.ProxyClientInfo.CertFile, "/var/run/secrets/aggregator-client/tls.crt")
}
func SetRecommendedKubeletConnectionInfoDefaults(config *kubecontrolplanev1.KubeletConnectionInfo) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if config.Port == 0 {
		config.Port = 10250
	}
	configdefaults.DefaultString(&config.CertInfo.KeyFile, "/var/run/secrets/kubelet-client/tls.key")
	configdefaults.DefaultString(&config.CertInfo.CertFile, "/var/run/secrets/kubelet-client/tls.crt")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
