package uploadconfig

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	rbachelper "k8s.io/kubernetes/pkg/apis/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	NodesKubeadmConfigClusterRoleName = "kubeadm:nodes-kubeadm-config"
)

func UploadConfiguration(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[uploadconfig] storing the configuration used in ConfigMap %q in the %q Namespace\n", kubeadmconstants.KubeadmConfigConfigMap, metav1.NamespaceSystem)
	clusterConfigurationToUpload := cfg.ClusterConfiguration.DeepCopy()
	clusterConfigurationToUpload.ComponentConfigs = kubeadmapi.ComponentConfigs{}
	clusterConfigurationYaml, err := configutil.MarshalKubeadmConfigObject(clusterConfigurationToUpload)
	if err != nil {
		return err
	}
	clusterStatus, err := configutil.GetClusterStatus(client)
	if err != nil {
		return err
	}
	if clusterStatus.APIEndpoints == nil {
		clusterStatus.APIEndpoints = map[string]kubeadmapi.APIEndpoint{}
	}
	clusterStatus.APIEndpoints[cfg.NodeRegistration.Name] = cfg.LocalAPIEndpoint
	clusterStatusYaml, err := configutil.MarshalKubeadmConfigObject(clusterStatus)
	if err != nil {
		return err
	}
	err = apiclient.CreateOrUpdateConfigMap(client, &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: kubeadmconstants.KubeadmConfigConfigMap, Namespace: metav1.NamespaceSystem}, Data: map[string]string{kubeadmconstants.ClusterConfigurationConfigMapKey: string(clusterConfigurationYaml), kubeadmconstants.ClusterStatusConfigMapKey: string(clusterStatusYaml)}})
	if err != nil {
		return err
	}
	err = apiclient.CreateOrUpdateRole(client, &rbac.Role{ObjectMeta: metav1.ObjectMeta{Name: NodesKubeadmConfigClusterRoleName, Namespace: metav1.NamespaceSystem}, Rules: []rbac.PolicyRule{rbachelper.NewRule("get").Groups("").Resources("configmaps").Names(kubeadmconstants.KubeadmConfigConfigMap).RuleOrDie()}})
	if err != nil {
		return err
	}
	return apiclient.CreateOrUpdateRoleBinding(client, &rbac.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: NodesKubeadmConfigClusterRoleName, Namespace: metav1.NamespaceSystem}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "Role", Name: NodesKubeadmConfigClusterRoleName}, Subjects: []rbac.Subject{{Kind: rbac.GroupKind, Name: kubeadmconstants.NodeBootstrapTokenAuthGroup}, {Kind: rbac.GroupKind, Name: kubeadmconstants.NodesGroup}}})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
