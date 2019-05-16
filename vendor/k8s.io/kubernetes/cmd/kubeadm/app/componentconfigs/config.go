package componentconfigs

import (
	goformat "fmt"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func GetFromKubeletConfigMap(client clientset.Interface, version *version.Version) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMapName := kubeadmconstants.GetKubeletConfigMapName(version)
	kubeletCfg, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kubeletConfigData, ok := kubeletCfg.Data[kubeadmconstants.KubeletBaseConfigurationConfigMapKey]
	if !ok {
		return nil, errors.Errorf("unexpected error when reading %s ConfigMap: %s key value pair missing", configMapName, kubeadmconstants.KubeletBaseConfigurationConfigMapKey)
	}
	obj := &kubeletconfig.KubeletConfiguration{}
	err = unmarshalObject(obj, []byte(kubeletConfigData))
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func GetFromKubeProxyConfigMap(client clientset.Interface, version *version.Version) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeproxyCfg, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(kubeadmconstants.KubeProxyConfigMap, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kubeproxyConfigData, ok := kubeproxyCfg.Data[kubeadmconstants.KubeProxyConfigMapKey]
	if !ok {
		return nil, errors.Errorf("unexpected error when reading %s ConfigMap: %s key value pair missing", kubeadmconstants.KubeProxyConfigMap, kubeadmconstants.KubeProxyConfigMapKey)
	}
	obj := &kubeproxyconfig.KubeProxyConfiguration{}
	err = unmarshalObject(obj, []byte(kubeproxyConfigData))
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
