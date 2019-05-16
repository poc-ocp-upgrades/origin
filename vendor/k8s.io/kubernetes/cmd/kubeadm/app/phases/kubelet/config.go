package kubelet

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	rbachelper "k8s.io/kubernetes/pkg/apis/rbac/v1"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func WriteConfigToDisk(kubeletConfig *kubeletconfig.KubeletConfiguration, kubeletDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeletBytes, err := getConfigBytes(kubeletConfig)
	if err != nil {
		return err
	}
	return writeConfigBytesToDisk(kubeletBytes, kubeletDir)
}
func CreateConfigMap(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	k8sVersion, err := version.ParseSemantic(cfg.KubernetesVersion)
	if err != nil {
		return err
	}
	configMapName := kubeadmconstants.GetKubeletConfigMapName(k8sVersion)
	fmt.Printf("[kubelet] Creating a ConfigMap %q in namespace %s with the configuration for the kubelets in the cluster\n", configMapName, metav1.NamespaceSystem)
	kubeletBytes, err := getConfigBytes(cfg.ComponentConfigs.Kubelet)
	if err != nil {
		return err
	}
	if err := apiclient.CreateOrUpdateConfigMap(client, &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: configMapName, Namespace: metav1.NamespaceSystem}, Data: map[string]string{kubeadmconstants.KubeletBaseConfigurationConfigMapKey: string(kubeletBytes)}}); err != nil {
		return err
	}
	if err := createConfigMapRBACRules(client, k8sVersion); err != nil {
		return errors.Wrap(err, "error creating kubelet configuration configmap RBAC rules")
	}
	return nil
}
func createConfigMapRBACRules(client clientset.Interface, k8sVersion *version.Version) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := apiclient.CreateOrUpdateRole(client, &rbac.Role{ObjectMeta: metav1.ObjectMeta{Name: configMapRBACName(k8sVersion), Namespace: metav1.NamespaceSystem}, Rules: []rbac.PolicyRule{rbachelper.NewRule("get").Groups("").Resources("configmaps").Names(kubeadmconstants.GetKubeletConfigMapName(k8sVersion)).RuleOrDie()}}); err != nil {
		return err
	}
	return apiclient.CreateOrUpdateRoleBinding(client, &rbac.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: configMapRBACName(k8sVersion), Namespace: metav1.NamespaceSystem}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "Role", Name: configMapRBACName(k8sVersion)}, Subjects: []rbac.Subject{{Kind: rbac.GroupKind, Name: kubeadmconstants.NodesGroup}, {Kind: rbac.GroupKind, Name: kubeadmconstants.NodeBootstrapTokenAuthGroup}}})
}
func DownloadConfig(client clientset.Interface, kubeletVersion *version.Version, kubeletDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMapName := kubeadmconstants.GetKubeletConfigMapName(kubeletVersion)
	fmt.Printf("[kubelet] Downloading configuration for the kubelet from the %q ConfigMap in the %s namespace\n", configMapName, metav1.NamespaceSystem)
	kubeletCfg, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(configMapName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) && kubeletVersion.Minor() == 10 {
		return nil
	}
	if err != nil {
		return err
	}
	return writeConfigBytesToDisk([]byte(kubeletCfg.Data[kubeadmconstants.KubeletBaseConfigurationConfigMapKey]), kubeletDir)
}
func configMapRBACName(k8sVersion *version.Version) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s%d.%d", kubeadmconstants.KubeletBaseConfigMapRolePrefix, k8sVersion.Major(), k8sVersion.Minor())
}
func getConfigBytes(kubeletConfig *kubeletconfig.KubeletConfiguration) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return componentconfigs.Known[componentconfigs.KubeletConfigurationKind].Marshal(kubeletConfig)
}
func writeConfigBytesToDisk(b []byte, kubeletDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configFile := filepath.Join(kubeletDir, kubeadmconstants.KubeletConfigurationFileName)
	fmt.Printf("[kubelet-start] Writing kubelet configuration to file %q\n", configFile)
	if err := os.MkdirAll(kubeletDir, 0700); err != nil {
		return errors.Wrapf(err, "failed to create directory %q", kubeletDir)
	}
	if err := ioutil.WriteFile(configFile, b, 0644); err != nil {
		return errors.Wrapf(err, "failed to write kubelet configuration to the file %q", configFile)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
