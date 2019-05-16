package proxy

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuberuntime "k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	rbachelper "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

const (
	KubeProxyClusterRoleName    = "system:node-proxier"
	KubeProxyServiceAccountName = "kube-proxy"
)

func EnsureProxyAddon(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := CreateServiceAccount(client); err != nil {
		return errors.Wrap(err, "error when creating kube-proxy service account")
	}
	masterEndpoint, err := kubeadmutil.GetMasterEndpoint(cfg)
	if err != nil {
		return err
	}
	proxyBytes, err := componentconfigs.Known[componentconfigs.KubeProxyConfigurationKind].Marshal(cfg.ComponentConfigs.KubeProxy)
	if err != nil {
		return errors.Wrap(err, "error when marshaling")
	}
	var prefixBytes bytes.Buffer
	apiclient.PrintBytesWithLinePrefix(&prefixBytes, proxyBytes, "    ")
	var proxyConfigMapBytes, proxyDaemonSetBytes []byte
	proxyConfigMapBytes, err = kubeadmutil.ParseTemplate(KubeProxyConfigMap19, struct {
		MasterEndpoint    string
		ProxyConfig       string
		ProxyConfigMap    string
		ProxyConfigMapKey string
	}{MasterEndpoint: masterEndpoint, ProxyConfig: prefixBytes.String(), ProxyConfigMap: constants.KubeProxyConfigMap, ProxyConfigMapKey: constants.KubeProxyConfigMapKey})
	if err != nil {
		return errors.Wrap(err, "error when parsing kube-proxy configmap template")
	}
	proxyDaemonSetBytes, err = kubeadmutil.ParseTemplate(KubeProxyDaemonSet19, struct{ Image, ProxyConfigMap, ProxyConfigMapKey string }{Image: images.GetKubernetesImage(constants.KubeProxy, &cfg.ClusterConfiguration), ProxyConfigMap: constants.KubeProxyConfigMap, ProxyConfigMapKey: constants.KubeProxyConfigMapKey})
	if err != nil {
		return errors.Wrap(err, "error when parsing kube-proxy daemonset template")
	}
	if err := createKubeProxyAddon(proxyConfigMapBytes, proxyDaemonSetBytes, client); err != nil {
		return err
	}
	if err := CreateRBACRules(client); err != nil {
		return errors.Wrap(err, "error when creating kube-proxy RBAC rules")
	}
	fmt.Println("[addons] Applied essential addon: kube-proxy")
	return nil
}
func CreateServiceAccount(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apiclient.CreateOrUpdateServiceAccount(client, &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: KubeProxyServiceAccountName, Namespace: metav1.NamespaceSystem}})
}
func CreateRBACRules(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return createClusterRoleBindings(client)
}
func createKubeProxyAddon(configMapBytes, daemonSetbytes []byte, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeproxyConfigMap := &v1.ConfigMap{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), configMapBytes, kubeproxyConfigMap); err != nil {
		return errors.Wrap(err, "unable to decode kube-proxy configmap")
	}
	if err := apiclient.CreateOrUpdateConfigMap(client, kubeproxyConfigMap); err != nil {
		return err
	}
	kubeproxyDaemonSet := &apps.DaemonSet{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), daemonSetbytes, kubeproxyDaemonSet); err != nil {
		return errors.Wrap(err, "unable to decode kube-proxy daemonset")
	}
	return apiclient.CreateOrUpdateDaemonSet(client, kubeproxyDaemonSet)
}
func createClusterRoleBindings(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := apiclient.CreateOrUpdateClusterRoleBinding(client, &rbac.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: "kubeadm:node-proxier"}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "ClusterRole", Name: KubeProxyClusterRoleName}, Subjects: []rbac.Subject{{Kind: rbac.ServiceAccountKind, Name: KubeProxyServiceAccountName, Namespace: metav1.NamespaceSystem}}}); err != nil {
		return err
	}
	if err := apiclient.CreateOrUpdateRole(client, &rbac.Role{ObjectMeta: metav1.ObjectMeta{Name: constants.KubeProxyConfigMap, Namespace: metav1.NamespaceSystem}, Rules: []rbac.PolicyRule{rbachelper.NewRule("get").Groups("").Resources("configmaps").Names(constants.KubeProxyConfigMap).RuleOrDie()}}); err != nil {
		return err
	}
	return apiclient.CreateOrUpdateRoleBinding(client, &rbac.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: constants.KubeProxyConfigMap, Namespace: metav1.NamespaceSystem}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "Role", Name: constants.KubeProxyConfigMap}, Subjects: []rbac.Subject{{Kind: rbac.GroupKind, Name: constants.NodeBootstrapTokenAuthGroup}}})
}
