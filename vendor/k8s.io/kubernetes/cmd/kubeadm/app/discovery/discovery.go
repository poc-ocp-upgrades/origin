package discovery

import (
	goformat "fmt"
	"github.com/pkg/errors"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/discovery/file"
	"k8s.io/kubernetes/cmd/kubeadm/app/discovery/https"
	"k8s.io/kubernetes/cmd/kubeadm/app/discovery/token"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const TokenUser = "tls-bootstrap-token-user"

func For(cfg *kubeadmapi.JoinConfiguration) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := DiscoverValidatedKubeConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't validate the identity of the API Server")
	}
	if len(cfg.Discovery.TLSBootstrapToken) == 0 {
		return config, nil
	}
	clusterinfo := kubeconfigutil.GetClusterFromKubeConfig(config)
	return kubeconfigutil.CreateWithToken(clusterinfo.Server, kubeadmapiv1beta1.DefaultClusterName, TokenUser, clusterinfo.CertificateAuthorityData, cfg.Discovery.TLSBootstrapToken), nil
}
func DiscoverValidatedKubeConfig(cfg *kubeadmapi.JoinConfiguration) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case cfg.Discovery.File != nil:
		kubeConfigPath := cfg.Discovery.File.KubeConfigPath
		if isHTTPSURL(kubeConfigPath) {
			return https.RetrieveValidatedConfigInfo(kubeConfigPath, kubeadmapiv1beta1.DefaultClusterName)
		}
		return file.RetrieveValidatedConfigInfo(kubeConfigPath, kubeadmapiv1beta1.DefaultClusterName)
	case cfg.Discovery.BootstrapToken != nil:
		return token.RetrieveValidatedConfigInfo(cfg)
	default:
		return nil, errors.New("couldn't find a valid discovery configuration")
	}
}
func isHTTPSURL(s string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	u, err := url.Parse(s)
	return err == nil && u.Scheme == "https"
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
