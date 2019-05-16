package node

import (
	"fmt"
	goformat "fmt"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	NodeBootstrapperClusterRoleName                      = "system:node-bootstrapper"
	NodeKubeletBootstrap                                 = "kubeadm:kubelet-bootstrap"
	CSRAutoApprovalClusterRoleName                       = "system:certificates.k8s.io:certificatesigningrequests:nodeclient"
	NodeSelfCSRAutoApprovalClusterRoleName               = "system:certificates.k8s.io:certificatesigningrequests:selfnodeclient"
	NodeAutoApproveBootstrapClusterRoleBinding           = "kubeadm:node-autoapprove-bootstrap"
	NodeAutoApproveCertificateRotationClusterRoleBinding = "kubeadm:node-autoapprove-certificate-rotation"
)

func AllowBootstrapTokensToPostCSRs(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[bootstraptoken] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials")
	return apiclient.CreateOrUpdateClusterRoleBinding(client, &rbac.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: NodeKubeletBootstrap}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "ClusterRole", Name: NodeBootstrapperClusterRoleName}, Subjects: []rbac.Subject{{Kind: rbac.GroupKind, Name: constants.NodeBootstrapTokenAuthGroup}}})
}
func AutoApproveNodeBootstrapTokens(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[bootstraptoken] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token")
	return apiclient.CreateOrUpdateClusterRoleBinding(client, &rbac.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: NodeAutoApproveBootstrapClusterRoleBinding}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "ClusterRole", Name: CSRAutoApprovalClusterRoleName}, Subjects: []rbac.Subject{{Kind: "Group", Name: constants.NodeBootstrapTokenAuthGroup}}})
}
func AutoApproveNodeCertificateRotation(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[bootstraptoken] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster")
	return apiclient.CreateOrUpdateClusterRoleBinding(client, &rbac.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: NodeAutoApproveCertificateRotationClusterRoleBinding}, RoleRef: rbac.RoleRef{APIGroup: rbac.GroupName, Kind: "ClusterRole", Name: NodeSelfCSRAutoApprovalClusterRoleName}, Subjects: []rbac.Subject{{Kind: "Group", Name: constants.NodesGroup}}})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
