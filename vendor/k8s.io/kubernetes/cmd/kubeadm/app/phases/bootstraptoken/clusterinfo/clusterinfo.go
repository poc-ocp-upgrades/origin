package clusterinfo

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/user"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	rbachelper "k8s.io/kubernetes/pkg/apis/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	BootstrapSignerClusterRoleName = "kubeadm:bootstrap-signer-clusterinfo"
)

func CreateBootstrapConfigMapIfNotExists(client clientset.Interface, file string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[bootstraptoken] creating the %q ConfigMap in the %q namespace\n", bootstrapapi.ConfigMapClusterInfo, metav1.NamespacePublic)
	klog.V(1).Infoln("[bootstraptoken] loading admin kubeconfig")
	adminConfig, err := clientcmd.LoadFromFile(file)
	if err != nil {
		return errors.Wrap(err, "failed to load admin kubeconfig")
	}
	adminCluster := adminConfig.Contexts[adminConfig.CurrentContext].Cluster
	klog.V(1).Infoln("[bootstraptoken] copying the cluster from admin.conf to the bootstrap kubeconfig")
	bootstrapConfig := &clientcmdapi.Config{Clusters: map[string]*clientcmdapi.Cluster{"": adminConfig.Clusters[adminCluster]}}
	bootstrapBytes, err := clientcmd.Write(*bootstrapConfig)
	if err != nil {
		return err
	}
	klog.V(1).Infoln("[bootstraptoken] creating/updating ConfigMap in kube-public namespace")
	return apiclient.CreateOrUpdateConfigMap(client, &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: bootstrapapi.ConfigMapClusterInfo, Namespace: metav1.NamespacePublic}, Data: map[string]string{bootstrapapi.KubeConfigKey: string(bootstrapBytes)}})
}
func CreateClusterInfoRBACRules(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("creating the RBAC rules for exposing the cluster-info ConfigMap in the kube-public namespace")
	err := apiclient.CreateOrUpdateRole(client, &rbac.Role{ObjectMeta: metav1.ObjectMeta{Name: BootstrapSignerClusterRoleName, Namespace: metav1.NamespacePublic}, Rules: []rbac.PolicyRule{rbachelper.NewRule("get").Groups("").Resources("configmaps").Names(bootstrapapi.ConfigMapClusterInfo).RuleOrDie()}})
	if err != nil {
		return err
	}
	return apiclient.CreateOrUpdateRoleBinding(client, &rbac.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: BootstrapSignerClusterRoleName, Namespace: metav1.NamespacePublic}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "Role", Name: BootstrapSignerClusterRoleName}, Subjects: []rbac.Subject{{Kind: rbac.UserKind, Name: user.Anonymous}}})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
