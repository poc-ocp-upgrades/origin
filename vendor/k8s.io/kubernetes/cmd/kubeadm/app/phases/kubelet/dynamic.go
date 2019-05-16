package kubelet

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
)

func EnableDynamicConfigForNode(client clientset.Interface, nodeName string, kubeletVersion *version.Version) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMapName := kubeadmconstants.GetKubeletConfigMapName(kubeletVersion)
	fmt.Printf("[kubelet] Enabling Dynamic Kubelet Config for Node %q; config sourced from ConfigMap %q in namespace %s\n", nodeName, configMapName, metav1.NamespaceSystem)
	fmt.Println("[kubelet] WARNING: The Dynamic Kubelet Config feature is beta, but off by default. It hasn't been well-tested yet at this stage, use with caution.")
	_, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "couldn't get the kubelet configuration ConfigMap")
	}
	return apiclient.PatchNode(client, nodeName, func(n *v1.Node) {
		patchNodeForDynamicConfig(n, configMapName)
	})
}
func patchNodeForDynamicConfig(n *v1.Node, configMapName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.Spec.ConfigSource = &v1.NodeConfigSource{ConfigMap: &v1.ConfigMapNodeConfigSource{Name: configMapName, Namespace: metav1.NamespaceSystem, KubeletConfigKey: kubeadmconstants.KubeletBaseConfigurationConfigMapKey}}
}
