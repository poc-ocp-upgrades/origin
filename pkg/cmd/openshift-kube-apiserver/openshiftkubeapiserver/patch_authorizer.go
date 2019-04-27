package openshiftkubeapiserver

import (
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/authorization/authorizerfactory"
	authorizerunion "k8s.io/apiserver/pkg/authorization/union"
	"k8s.io/client-go/informers"
	"k8s.io/kubernetes/pkg/auth/nodeidentifier"
	"k8s.io/kubernetes/plugin/pkg/auth/authorizer/node"
	rbacauthorizer "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	kbootstrappolicy "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy"
	"github.com/openshift/origin/pkg/authorization/authorizer/browsersafe"
	"github.com/openshift/origin/pkg/authorization/authorizer/scope"
)

func NewAuthorizer(versionedInformers informers.SharedInformerFactory) authorizer.Authorizer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rbacInformers := versionedInformers.Rbac().V1()
	scopeLimitedAuthorizer := scope.NewAuthorizer(rbacInformers.ClusterRoles().Lister())
	kubeAuthorizer := rbacauthorizer.New(&rbacauthorizer.RoleGetter{Lister: rbacInformers.Roles().Lister()}, &rbacauthorizer.RoleBindingLister{Lister: rbacInformers.RoleBindings().Lister()}, &rbacauthorizer.ClusterRoleGetter{Lister: rbacInformers.ClusterRoles().Lister()}, &rbacauthorizer.ClusterRoleBindingLister{Lister: rbacInformers.ClusterRoleBindings().Lister()})
	graph := node.NewGraph()
	node.AddGraphEventHandlers(graph, versionedInformers.Core().V1().Nodes(), versionedInformers.Core().V1().Pods(), versionedInformers.Core().V1().PersistentVolumes(), versionedInformers.Storage().V1beta1().VolumeAttachments())
	nodeAuthorizer := node.NewAuthorizer(graph, nodeidentifier.NewDefaultNodeIdentifier(), kbootstrappolicy.NodeRules())
	openshiftAuthorizer := authorizerunion.New(browsersafe.NewBrowserSafeAuthorizer(scopeLimitedAuthorizer, user.AllAuthenticated), authorizerfactory.NewPrivilegedGroups(user.SystemPrivilegedGroup), nodeAuthorizer, browsersafe.NewBrowserSafeAuthorizer(kubeAuthorizer, user.AllAuthenticated))
	return openshiftAuthorizer
}
