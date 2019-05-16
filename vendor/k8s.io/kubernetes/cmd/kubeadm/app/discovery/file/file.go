package file

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func RetrieveValidatedConfigInfo(filepath, clustername string) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := clientcmd.LoadFromFile(filepath)
	if err != nil {
		return nil, err
	}
	return ValidateConfigInfo(config, clustername)
}
func ValidateConfigInfo(config *clientcmdapi.Config, clustername string) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := validateKubeConfig(config)
	if err != nil {
		return nil, err
	}
	defaultCluster := kubeconfigutil.GetClusterFromKubeConfig(config)
	kubeconfig := kubeconfigutil.CreateBasic(defaultCluster.Server, clustername, "", defaultCluster.CertificateAuthorityData)
	if config.Contexts[config.CurrentContext] != nil && len(config.AuthInfos) > 0 {
		user := config.Contexts[config.CurrentContext].AuthInfo
		authInfo, ok := config.AuthInfos[user]
		if !ok || authInfo == nil {
			return nil, errors.Errorf("empty settings for user %q", user)
		}
		if len(authInfo.ClientCertificateData) == 0 && len(authInfo.ClientCertificate) != 0 {
			clientCert, err := ioutil.ReadFile(authInfo.ClientCertificate)
			if err != nil {
				return nil, err
			}
			authInfo.ClientCertificateData = clientCert
		}
		if len(authInfo.ClientKeyData) == 0 && len(authInfo.ClientKey) != 0 {
			clientKey, err := ioutil.ReadFile(authInfo.ClientKey)
			if err != nil {
				return nil, err
			}
			authInfo.ClientKeyData = clientKey
		}
		if len(authInfo.ClientCertificateData) == 0 || len(authInfo.ClientKeyData) == 0 {
			return nil, errors.New("couldn't read authentication info from the given kubeconfig file")
		}
		kubeconfig = kubeconfigutil.CreateWithCerts(defaultCluster.Server, clustername, "", defaultCluster.CertificateAuthorityData, authInfo.ClientKeyData, authInfo.ClientCertificateData)
	}
	client, err := kubeconfigutil.ToClientSet(kubeconfig)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[discovery] Created cluster-info discovery client, requesting info from %q\n", defaultCluster.Server)
	var clusterinfoCM *v1.ConfigMap
	wait.PollInfinite(constants.DiscoveryRetryInterval, func() (bool, error) {
		var err error
		clusterinfoCM, err = client.CoreV1().ConfigMaps(metav1.NamespacePublic).Get(bootstrapapi.ConfigMapClusterInfo, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsForbidden(err) {
				fmt.Printf("[discovery] Could not access the %s ConfigMap for refreshing the cluster-info information, but the TLS cert is valid so proceeding...\n", bootstrapapi.ConfigMapClusterInfo)
				return true, nil
			}
			fmt.Printf("[discovery] Failed to validate the API Server's identity, will try again: [%v]\n", err)
			return false, nil
		}
		return true, nil
	})
	if clusterinfoCM == nil {
		return kubeconfig, nil
	}
	refreshedBaseKubeConfig, err := tryParseClusterInfoFromConfigMap(clusterinfoCM)
	if err != nil {
		fmt.Printf("[discovery] The %s ConfigMap isn't set up properly (%v), but the TLS cert is valid so proceeding...\n", bootstrapapi.ConfigMapClusterInfo, err)
		return kubeconfig, nil
	}
	fmt.Println("[discovery] Synced cluster-info information from the API Server so we have got the latest information")
	return refreshedBaseKubeConfig, nil
}
func tryParseClusterInfoFromConfigMap(cm *v1.ConfigMap) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigString, ok := cm.Data[bootstrapapi.KubeConfigKey]
	if !ok || len(kubeConfigString) == 0 {
		return nil, errors.Errorf("no %s key in ConfigMap", bootstrapapi.KubeConfigKey)
	}
	parsedKubeConfig, err := clientcmd.Load([]byte(kubeConfigString))
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't parse the kubeconfig file in the %s ConfigMap", bootstrapapi.ConfigMapClusterInfo)
	}
	return parsedKubeConfig, nil
}
func validateKubeConfig(config *clientcmdapi.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(config.Clusters) < 1 {
		return errors.New("the provided cluster-info kubeconfig file must have at least one Cluster defined")
	}
	defaultCluster := kubeconfigutil.GetClusterFromKubeConfig(config)
	if defaultCluster == nil {
		return errors.New("the provided cluster-info kubeconfig file must have an unnamed Cluster or a CurrentContext that specifies a non-nil Cluster")
	}
	return clientcmd.Validate(*config)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
