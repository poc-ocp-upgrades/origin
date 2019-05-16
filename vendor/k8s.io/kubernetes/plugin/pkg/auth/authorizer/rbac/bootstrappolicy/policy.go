package bootstrappolicy

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	"k8s.io/kubernetes/pkg/features"
)

var (
	Write      = []string{"create", "update", "patch", "delete", "deletecollection"}
	ReadWrite  = []string{"get", "list", "watch", "create", "update", "patch", "delete", "deletecollection"}
	Read       = []string{"get", "list", "watch"}
	ReadUpdate = []string{"get", "list", "watch", "update", "patch"}
	Label      = map[string]string{"kubernetes.io/bootstrapping": "rbac-defaults"}
	Annotation = map[string]string{rbacv1.AutoUpdateAnnotationKey: "true"}
)

const (
	legacyGroup         = ""
	appsGroup           = "apps"
	authenticationGroup = "authentication.k8s.io"
	authorizationGroup  = "authorization.k8s.io"
	autoscalingGroup    = "autoscaling"
	batchGroup          = "batch"
	certificatesGroup   = "certificates.k8s.io"
	extensionsGroup     = "extensions"
	policyGroup         = "policy"
	rbacGroup           = "rbac.authorization.k8s.io"
	storageGroup        = "storage.k8s.io"
	resMetricsGroup     = "metrics.k8s.io"
	customMetricsGroup  = "custom.metrics.k8s.io"
	networkingGroup     = "networking.k8s.io"
)

func addDefaultMetadata(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	metadata, err := meta.Accessor(obj)
	if err != nil {
		panic(err)
	}
	labels := metadata.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	for k, v := range Label {
		labels[k] = v
	}
	metadata.SetLabels(labels)
	annotations := metadata.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	for k, v := range Annotation {
		annotations[k] = v
	}
	metadata.SetAnnotations(annotations)
}
func addClusterRoleLabel(roles []rbacv1.ClusterRole) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range roles {
		addDefaultMetadata(&roles[i])
	}
	return
}
func addClusterRoleBindingLabel(rolebindings []rbacv1.ClusterRoleBinding) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range rolebindings {
		addDefaultMetadata(&rolebindings[i])
	}
	return
}
func NodeRules() []rbacv1.PolicyRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodePolicyRules := []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(authenticationGroup).Resources("tokenreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("subjectaccessreviews", "localsubjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("services").RuleOrDie(), rbacv1helpers.NewRule("create", "get", "list", "watch").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("update", "patch").Groups(legacyGroup).Resources("nodes/status").RuleOrDie(), rbacv1helpers.NewRule("update", "patch").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("create", "update", "patch").Groups(legacyGroup).Resources("events").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("create", "delete").Groups(legacyGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("update", "patch").Groups(legacyGroup).Resources("pods/status").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(legacyGroup).Resources("pods/eviction").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("secrets", "configmaps").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("persistentvolumeclaims", "persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("endpoints").RuleOrDie(), rbacv1helpers.NewRule("create", "get", "list", "watch").Groups(certificatesGroup).Resources("certificatesigningrequests").RuleOrDie()}
	if utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) {
		pvcStatusPolicyRule := rbacv1helpers.NewRule("get", "update", "patch").Groups(legacyGroup).Resources("persistentvolumeclaims/status").RuleOrDie()
		nodePolicyRules = append(nodePolicyRules, pvcStatusPolicyRule)
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.TokenRequest) {
		tokenRequestRule := rbacv1helpers.NewRule("create").Groups(legacyGroup).Resources("serviceaccounts/token").RuleOrDie()
		nodePolicyRules = append(nodePolicyRules, tokenRequestRule)
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.CSIPersistentVolume) {
		volAttachRule := rbacv1helpers.NewRule("get").Groups(storageGroup).Resources("volumeattachments").RuleOrDie()
		nodePolicyRules = append(nodePolicyRules, volAttachRule)
		if utilfeature.DefaultFeatureGate.Enabled(features.CSIDriverRegistry) {
			csiDriverRule := rbacv1helpers.NewRule("get", "watch", "list").Groups("csi.storage.k8s.io").Resources("csidrivers").RuleOrDie()
			nodePolicyRules = append(nodePolicyRules, csiDriverRule)
		}
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.KubeletPluginsWatcher) && utilfeature.DefaultFeatureGate.Enabled(features.CSINodeInfo) {
		csiNodeInfoRule := rbacv1helpers.NewRule("get", "create", "update", "patch", "delete").Groups("csi.storage.k8s.io").Resources("csinodeinfos").RuleOrDie()
		nodePolicyRules = append(nodePolicyRules, csiNodeInfoRule)
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.NodeLease) {
		nodePolicyRules = append(nodePolicyRules, rbacv1helpers.NewRule("get", "create", "update", "patch", "delete").Groups("coordination.k8s.io").Resources("leases").RuleOrDie())
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.RuntimeClass) {
		nodePolicyRules = append(nodePolicyRules, rbacv1helpers.NewRule("get", "list", "watch").Groups("node.k8s.io").Resources("runtimeclasses").RuleOrDie())
	}
	return nodePolicyRules
}
func clusterRoles() []rbacv1.ClusterRole {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	roles := []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "cluster-admin"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("*").Groups("*").Resources("*").RuleOrDie(), rbacv1helpers.NewRule("*").URLs("*").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:discovery"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get").URLs("/healthz", "/version", "/version/", "/swaggerapi", "/swaggerapi/*", "/swagger.json", "/swagger-2.0.0.pb-v1", "/openapi", "/openapi/*", "/api", "/api/*", "/apis", "/apis/*").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:openshift:public-info-viewer"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get").URLs("/readyz").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:basic-user"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("selfsubjectaccessreviews", "selfsubjectrulesreviews").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "admin"}, AggregationRule: &rbacv1.AggregationRule{ClusterRoleSelectors: []metav1.LabelSelector{{MatchLabels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-admin": "true"}}}}}, {ObjectMeta: metav1.ObjectMeta{Name: "edit", Labels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-admin": "true"}}, AggregationRule: &rbacv1.AggregationRule{ClusterRoleSelectors: []metav1.LabelSelector{{MatchLabels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-edit": "true"}}}}}, {ObjectMeta: metav1.ObjectMeta{Name: "view", Labels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-edit": "true"}}, AggregationRule: &rbacv1.AggregationRule{ClusterRoleSelectors: []metav1.LabelSelector{{MatchLabels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-view": "true"}}}}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:aggregate-to-admin", Labels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-admin": "true"}}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("localsubjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule(ReadWrite...).Groups(rbacGroup).Resources("roles", "rolebindings").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:aggregate-to-edit", Labels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-edit": "true"}}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("pods/attach", "pods/proxy", "pods/exec", "pods/portforward", "secrets", "services/proxy").RuleOrDie(), rbacv1helpers.NewRule("impersonate").Groups(legacyGroup).Resources("serviceaccounts").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(legacyGroup).Resources("pods", "pods/attach", "pods/proxy", "pods/exec", "pods/portforward").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(legacyGroup).Resources("replicationcontrollers", "replicationcontrollers/scale", "serviceaccounts", "services", "services/proxy", "endpoints", "persistentvolumeclaims", "configmaps", "secrets").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(appsGroup).Resources("statefulsets", "statefulsets/scale", "daemonsets", "deployments", "deployments/scale", "deployments/rollback", "replicasets", "replicasets/scale").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(autoscalingGroup).Resources("horizontalpodautoscalers").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(batchGroup).Resources("jobs", "cronjobs").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(extensionsGroup).Resources("daemonsets", "deployments", "deployments/scale", "deployments/rollback", "ingresses", "replicasets", "replicasets/scale", "replicationcontrollers/scale", "networkpolicies").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(policyGroup).Resources("poddisruptionbudgets").RuleOrDie(), rbacv1helpers.NewRule(Write...).Groups(networkingGroup).Resources("networkpolicies").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:aggregate-to-view", Labels: map[string]string{"rbac.authorization.k8s.io/aggregate-to-view": "true"}}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("pods", "replicationcontrollers", "replicationcontrollers/scale", "serviceaccounts", "services", "endpoints", "persistentvolumeclaims", "configmaps").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("limitranges", "resourcequotas", "bindings", "events", "pods/status", "resourcequotas/status", "namespaces/status", "replicationcontrollers/status", "pods/log").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("namespaces").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(appsGroup).Resources("controllerrevisions", "statefulsets", "statefulsets/scale", "daemonsets", "deployments", "deployments/scale", "replicasets", "replicasets/scale").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(autoscalingGroup).Resources("horizontalpodautoscalers").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(batchGroup).Resources("jobs", "cronjobs").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(extensionsGroup).Resources("daemonsets", "deployments", "deployments/scale", "ingresses", "replicasets", "replicasets/scale", "replicationcontrollers/scale", "networkpolicies").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(policyGroup).Resources("poddisruptionbudgets").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(networkingGroup).Resources("networkpolicies").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:heapster"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("events", "pods", "nodes", "namespaces").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(extensionsGroup).Resources("deployments").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:node"}, Rules: NodeRules()}, {ObjectMeta: metav1.ObjectMeta{Name: "system:node-problem-detector"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("patch").Groups(legacyGroup).Resources("nodes/status").RuleOrDie(), eventsRule()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:node-proxier"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch").Groups(legacyGroup).Resources("services", "endpoints").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("nodes").RuleOrDie(), eventsRule()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:kubelet-api-admin"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("proxy").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("*").Groups(legacyGroup).Resources("nodes/proxy", "nodes/metrics", "nodes/spec", "nodes/stats", "nodes/log").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:node-bootstrapper"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create", "get", "list", "watch").Groups(certificatesGroup).Resources("certificatesigningrequests").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:auth-delegator"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(authenticationGroup).Resources("tokenreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("subjectaccessreviews").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:kube-aggregator"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("services", "endpoints").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:kube-controller-manager"}, Rules: []rbacv1.PolicyRule{eventsRule(), rbacv1helpers.NewRule("create").Groups(legacyGroup).Resources("endpoints", "secrets", "serviceaccounts").RuleOrDie(), rbacv1helpers.NewRule("delete").Groups(legacyGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(legacyGroup).Resources("endpoints", "namespaces", "secrets", "serviceaccounts", "configmaps").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(legacyGroup).Resources("endpoints", "secrets", "serviceaccounts").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authenticationGroup).Resources("tokenreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("subjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule("list", "watch").Groups("*").Resources("*").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:kube-scheduler"}, Rules: []rbacv1.PolicyRule{eventsRule(), rbacv1helpers.NewRule("create").Groups(legacyGroup).Resources("endpoints").RuleOrDie(), rbacv1helpers.NewRule("get", "update", "patch", "delete").Groups(legacyGroup).Resources("endpoints").Names("kube-scheduler").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "delete").Groups(legacyGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(legacyGroup).Resources("pods/binding", "bindings").RuleOrDie(), rbacv1helpers.NewRule("patch", "update").Groups(legacyGroup).Resources("pods/status").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("services", "replicationcontrollers").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(appsGroup, extensionsGroup).Resources("replicasets").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(appsGroup).Resources("statefulsets").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(policyGroup).Resources("poddisruptionbudgets").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(legacyGroup).Resources("persistentvolumeclaims", "persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authenticationGroup).Resources("tokenreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authorizationGroup).Resources("subjectaccessreviews").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:kube-dns"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch").Groups(legacyGroup).Resources("endpoints", "services").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:persistent-volume-provisioner"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "create", "delete").Groups(legacyGroup).Resources("persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "update").Groups(legacyGroup).Resources("persistentvolumeclaims").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(storageGroup).Resources("storageclasses").RuleOrDie(), rbacv1helpers.NewRule("watch").Groups(legacyGroup).Resources("events").RuleOrDie(), eventsRule()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:csi-external-attacher"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "update", "patch").Groups(legacyGroup).Resources("persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "update", "patch").Groups(storageGroup).Resources("volumeattachments").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch").Groups(legacyGroup).Resources("events").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:aws-cloud-provider"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "patch").Groups(legacyGroup).Resources("nodes").RuleOrDie(), eventsRule()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:certificates.k8s.io:certificatesigningrequests:nodeclient"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(certificatesGroup).Resources("certificatesigningrequests/nodeclient").RuleOrDie()}}, {ObjectMeta: metav1.ObjectMeta{Name: "system:certificates.k8s.io:certificatesigningrequests:selfnodeclient"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(certificatesGroup).Resources("certificatesigningrequests/selfnodeclient").RuleOrDie()}}}
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		roles = append(roles, rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: "system:volume-scheduler"}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule(ReadUpdate...).Groups(legacyGroup).Resources("persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule(Read...).Groups(storageGroup).Resources("storageclasses").RuleOrDie(), rbacv1helpers.NewRule(ReadUpdate...).Groups(legacyGroup).Resources("persistentvolumeclaims").RuleOrDie()}})
	}
	externalProvisionerRules := []rbacv1.PolicyRule{rbacv1helpers.NewRule("create", "delete", "get", "list", "watch").Groups(legacyGroup).Resources("persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "update", "patch").Groups(legacyGroup).Resources("persistentvolumeclaims").RuleOrDie(), rbacv1helpers.NewRule("list", "watch").Groups(storageGroup).Resources("storageclasses").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch").Groups(legacyGroup).Resources("events").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(legacyGroup).Resources("nodes").RuleOrDie()}
	if utilfeature.DefaultFeatureGate.Enabled(features.CSINodeInfo) {
		externalProvisionerRules = append(externalProvisionerRules, rbacv1helpers.NewRule("get", "watch", "list").Groups("csi.storage.k8s.io").Resources("csinodeinfos").RuleOrDie())
	}
	roles = append(roles, rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: "system:csi-external-provisioner"}, Rules: externalProvisionerRules})
	addClusterRoleLabel(roles)
	return roles
}

const systemNodeRoleName = "system:node"

func clusterRoleBindings() []rbacv1.ClusterRoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rolebindings := []rbacv1.ClusterRoleBinding{rbacv1helpers.NewClusterBinding("cluster-admin").Groups(user.SystemPrivilegedGroup).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:discovery").Groups(user.AllAuthenticated, user.AllUnauthenticated).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:openshift:public-info-viewer").Groups(user.AllAuthenticated, user.AllUnauthenticated).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:basic-user").Groups(user.AllAuthenticated, user.AllUnauthenticated).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:node-proxier").Users(user.KubeProxy).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:kube-controller-manager").Users(user.KubeControllerManager).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:kube-dns").SAs("kube-system", "kube-dns").BindingOrDie(), rbacv1helpers.NewClusterBinding("system:kube-scheduler").Users(user.KubeScheduler).BindingOrDie(), rbacv1helpers.NewClusterBinding("system:aws-cloud-provider").SAs("kube-system", "aws-cloud-provider").BindingOrDie(), {ObjectMeta: metav1.ObjectMeta{Name: systemNodeRoleName}, RoleRef: rbacv1.RoleRef{APIGroup: rbacv1.GroupName, Kind: "ClusterRole", Name: systemNodeRoleName}}}
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		rolebindings = append(rolebindings, rbacv1helpers.NewClusterBinding("system:volume-scheduler").Users(user.KubeScheduler).BindingOrDie())
	}
	addClusterRoleBindingLabel(rolebindings)
	return rolebindings
}
func ClusterRolesToAggregate() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return map[string]string{"admin": "system:aggregate-to-admin", "edit": "system:aggregate-to-edit", "view": "system:aggregate-to-view"}
}
