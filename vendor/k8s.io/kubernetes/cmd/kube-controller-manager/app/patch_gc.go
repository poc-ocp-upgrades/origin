package app

import (
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/config"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

func applyOpenShiftGCConfig(controllerManager *config.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerManager.ComponentConfig.GarbageCollectorController.GCIgnoredResources = append(controllerManager.ComponentConfig.GarbageCollectorController.GCIgnoredResources, kubectrlmgrconfig.GroupResource{Group: "authorization.openshift.io", Resource: "rolebindingrestrictions"}, kubectrlmgrconfig.GroupResource{Group: "network.openshift.io", Resource: "clusternetworks"}, kubectrlmgrconfig.GroupResource{Group: "network.openshift.io", Resource: "egressnetworkpolicies"}, kubectrlmgrconfig.GroupResource{Group: "network.openshift.io", Resource: "hostsubnets"}, kubectrlmgrconfig.GroupResource{Group: "network.openshift.io", Resource: "netnamespaces"}, kubectrlmgrconfig.GroupResource{Group: "oauth.openshift.io", Resource: "oauthclientauthorizations"}, kubectrlmgrconfig.GroupResource{Group: "oauth.openshift.io", Resource: "oauthclients"}, kubectrlmgrconfig.GroupResource{Group: "quota.openshift.io", Resource: "clusterresourcequotas"}, kubectrlmgrconfig.GroupResource{Group: "user.openshift.io", Resource: "groups"}, kubectrlmgrconfig.GroupResource{Group: "user.openshift.io", Resource: "identities"}, kubectrlmgrconfig.GroupResource{Group: "user.openshift.io", Resource: "users"}, kubectrlmgrconfig.GroupResource{Group: "image.openshift.io", Resource: "images"}, kubectrlmgrconfig.GroupResource{Group: "project.openshift.io", Resource: "projects"}, kubectrlmgrconfig.GroupResource{Group: "authorization.openshift.io", Resource: "clusterroles"}, kubectrlmgrconfig.GroupResource{Group: "authorization.openshift.io", Resource: "clusterrolebindings"}, kubectrlmgrconfig.GroupResource{Group: "authorization.openshift.io", Resource: "roles"}, kubectrlmgrconfig.GroupResource{Group: "authorization.openshift.io", Resource: "rolebindings"}, kubectrlmgrconfig.GroupResource{Group: "oauth.openshift.io", Resource: "oauthaccesstokens"}, kubectrlmgrconfig.GroupResource{Group: "oauth.openshift.io", Resource: "oauthauthorizetokens"}, kubectrlmgrconfig.GroupResource{Group: "apps", Resource: "deployments"}, kubectrlmgrconfig.GroupResource{Group: "extensions", Resource: "horizontalpodautoscalers"}, kubectrlmgrconfig.GroupResource{Group: "", Resource: "securitycontextconstraints"})
	return nil
}
