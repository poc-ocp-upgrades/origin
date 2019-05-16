package config

import (
	"crypto/x509"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	certutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func FetchConfigFromFileOrCluster(client clientset.Interface, w io.Writer, logPrefix, cfgPath string, newControlPlane bool) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	initcfg, err := loadConfiguration(client, w, logPrefix, cfgPath, newControlPlane)
	if err != nil {
		return nil, err
	}
	if err := SetInitDynamicDefaults(initcfg); err != nil {
		return nil, err
	}
	return initcfg, err
}
func loadConfiguration(client clientset.Interface, w io.Writer, logPrefix, cfgPath string, newControlPlane bool) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfgPath != "" {
		fmt.Fprintf(w, "[%s] Reading configuration options from a file: %s\n", logPrefix, cfgPath)
		return loadInitConfigurationFromFile(cfgPath)
	}
	fmt.Fprintf(w, "[%s] Reading configuration from the cluster...\n", logPrefix)
	fmt.Fprintf(w, "[%s] FYI: You can look at this config file with 'kubectl -n %s get cm %s -oyaml'\n", logPrefix, metav1.NamespaceSystem, constants.KubeadmConfigConfigMap)
	return getInitConfigurationFromCluster(constants.KubernetesDir, client, newControlPlane)
}
func loadInitConfigurationFromFile(cfgPath string) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	initcfg, err := BytesToInternalConfig(configBytes)
	if err != nil {
		return nil, err
	}
	return initcfg, nil
}
func getInitConfigurationFromCluster(kubeconfigDir string, client clientset.Interface, newControlPlane bool) (*kubeadmapi.InitConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMap, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(constants.KubeadmConfigConfigMap, metav1.GetOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get config map")
	}
	initcfg := &kubeadmapi.InitConfiguration{}
	clusterConfigurationData, ok := configMap.Data[constants.ClusterConfigurationConfigMapKey]
	if !ok {
		return nil, errors.Errorf("unexpected error when reading kubeadm-config ConfigMap: %s key value pair missing", constants.ClusterConfigurationConfigMapKey)
	}
	if err := runtime.DecodeInto(kubeadmscheme.Codecs.UniversalDecoder(), []byte(clusterConfigurationData), &initcfg.ClusterConfiguration); err != nil {
		return nil, errors.Wrap(err, "failed to decode cluster configuration data")
	}
	if err := getComponentConfigs(client, &initcfg.ClusterConfiguration); err != nil {
		return nil, errors.Wrap(err, "failed to get component configs")
	}
	if !newControlPlane {
		if err := getNodeRegistration(kubeconfigDir, client, &initcfg.NodeRegistration); err != nil {
			return nil, errors.Wrap(err, "failed to get node registration")
		}
		if err := getAPIEndpoint(configMap.Data, initcfg.NodeRegistration.Name, &initcfg.LocalAPIEndpoint); err != nil {
			return nil, errors.Wrap(err, "failed to getAPIEndpoint")
		}
	}
	return initcfg, nil
}
func getNodeRegistration(kubeconfigDir string, client clientset.Interface, nodeRegistration *kubeadmapi.NodeRegistrationOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeName, err := getNodeNameFromKubeletConfig(kubeconfigDir)
	if err != nil {
		return errors.Wrap(err, "failed to get node name from kubelet config")
	}
	node, err := client.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "faild to get corresponding node")
	}
	criSocket, ok := node.ObjectMeta.Annotations[constants.AnnotationKubeadmCRISocket]
	if !ok {
		return errors.Errorf("node %s doesn't have %s annotation", nodeName, constants.AnnotationKubeadmCRISocket)
	}
	nodeRegistration.Name = nodeName
	nodeRegistration.CRISocket = criSocket
	nodeRegistration.Taints = node.Spec.Taints
	return nil
}
func getNodeNameFromKubeletConfig(kubeconfigDir string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fileName := filepath.Join(kubeconfigDir, constants.KubeletKubeConfigFileName)
	config, err := clientcmd.LoadFromFile(fileName)
	if err != nil {
		return "", err
	}
	authInfo := config.AuthInfos[config.Contexts[config.CurrentContext].AuthInfo]
	var certs []*x509.Certificate
	if len(authInfo.ClientCertificateData) > 0 {
		if certs, err = certutil.ParseCertsPEM(authInfo.ClientCertificateData); err != nil {
			return "", err
		}
	} else if len(authInfo.ClientCertificate) > 0 {
		if certs, err = certutil.CertsFromFile(authInfo.ClientCertificate); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("invalid kubelet.conf. X509 certificate expected")
	}
	cert := certs[0]
	return strings.TrimPrefix(cert.Subject.CommonName, constants.NodesUserPrefix), nil
}
func getAPIEndpoint(data map[string]string, nodeName string, apiEndpoint *kubeadmapi.APIEndpoint) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterStatus, err := unmarshalClusterStatus(data)
	if err != nil {
		return err
	}
	e, ok := clusterStatus.APIEndpoints[nodeName]
	if !ok {
		return errors.New("failed to get APIEndpoint information for this node")
	}
	apiEndpoint.AdvertiseAddress = e.AdvertiseAddress
	apiEndpoint.BindPort = e.BindPort
	return nil
}
func getComponentConfigs(client clientset.Interface, clusterConfiguration *kubeadmapi.ClusterConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	k8sVersion := version.MustParseGeneric(clusterConfiguration.KubernetesVersion)
	for kind, registration := range componentconfigs.Known {
		obj, err := registration.GetFromConfigMap(client, k8sVersion)
		if err != nil {
			return err
		}
		if ok := registration.SetToInternalConfig(obj, clusterConfiguration); !ok {
			return errors.Errorf("couldn't save componentconfig value for kind %q", string(kind))
		}
	}
	return nil
}
func GetClusterStatus(client clientset.Interface) (*kubeadmapi.ClusterStatus, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMap, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(constants.KubeadmConfigConfigMap, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		return &kubeadmapi.ClusterStatus{}, nil
	}
	if err != nil {
		return nil, err
	}
	clusterStatus, err := unmarshalClusterStatus(configMap.Data)
	if err != nil {
		return nil, err
	}
	return clusterStatus, nil
}
func unmarshalClusterStatus(data map[string]string) (*kubeadmapi.ClusterStatus, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterStatusData, ok := data[constants.ClusterStatusConfigMapKey]
	if !ok {
		return nil, errors.Errorf("unexpected error when reading kubeadm-config ConfigMap: %s key value pair missing", constants.ClusterStatusConfigMapKey)
	}
	clusterStatus := &kubeadmapi.ClusterStatus{}
	if err := runtime.DecodeInto(kubeadmscheme.Codecs.UniversalDecoder(), []byte(clusterStatusData), clusterStatus); err != nil {
		return nil, err
	}
	return clusterStatus, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
