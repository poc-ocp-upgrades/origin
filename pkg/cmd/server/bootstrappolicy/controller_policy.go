package bootstrappolicy

import (
	_ "github.com/openshift/origin/pkg/authorization/apis/authorization/install"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	"strings"
)

const saRolePrefix = "system:openshift:controller:"
const (
	InfraOriginNamespaceServiceAccountName                       = "origin-namespace-controller"
	InfraServiceAccountControllerServiceAccountName              = "serviceaccount-controller"
	InfraServiceAccountPullSecretsControllerServiceAccountName   = "serviceaccount-pull-secrets-controller"
	InfraServiceAccountTokensControllerServiceAccountName        = "serviceaccount-tokens-controller"
	InfraServiceServingCertServiceAccountName                    = "service-serving-cert-controller"
	InfraBuildControllerServiceAccountName                       = "build-controller"
	InfraBuildConfigChangeControllerServiceAccountName           = "build-config-change-controller"
	InfraDeploymentConfigControllerServiceAccountName            = "deploymentconfig-controller"
	InfraDeployerControllerServiceAccountName                    = "deployer-controller"
	InfraImageTriggerControllerServiceAccountName                = "image-trigger-controller"
	InfraImageImportControllerServiceAccountName                 = "image-import-controller"
	InfraSDNControllerServiceAccountName                         = "sdn-controller"
	InfraClusterQuotaReconciliationControllerServiceAccountName  = "cluster-quota-reconciliation-controller"
	InfraUnidlingControllerServiceAccountName                    = "unidling-controller"
	InfraServiceIngressIPControllerServiceAccountName            = "service-ingress-ip-controller"
	InfraPersistentVolumeRecyclerControllerServiceAccountName    = "pv-recycler-controller"
	InfraResourceQuotaControllerServiceAccountName               = "resourcequota-controller"
	InfraDefaultRoleBindingsControllerServiceAccountName         = "default-rolebindings-controller"
	InfraIngressToRouteControllerServiceAccountName              = "ingress-to-route-controller"
	InfraNamespaceSecurityAllocationControllerServiceAccountName = "namespace-security-allocation-controller"
	InfraTemplateInstanceControllerServiceAccountName            = "template-instance-controller"
	InfraTemplateInstanceFinalizerControllerServiceAccountName   = "template-instance-finalizer-controller"
	InfraTemplateServiceBrokerServiceAccountName                 = "template-service-broker"
	InfraHorizontalPodAutoscalerControllerServiceAccountName     = "horizontal-pod-autoscaler"
)

var (
	controllerRoles        = []rbacv1.ClusterRole{}
	controllerRoleBindings = []rbacv1.ClusterRoleBinding{}
)

func bindControllerRole(saName string, roleName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	roleBinding := rbacv1helpers.NewClusterBinding(roleName).SAs(DefaultOpenShiftInfraNamespace, saName).BindingOrDie()
	addDefaultMetadata(&roleBinding)
	controllerRoleBindings = append(controllerRoleBindings, roleBinding)
}
func addControllerRole(role rbacv1.ClusterRole) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(role.Name, saRolePrefix) {
		klog.Fatalf(`role %q must start with %q`, role.Name, saRolePrefix)
	}
	addControllerRoleToSA(DefaultOpenShiftInfraNamespace, role.Name[len(saRolePrefix):], role)
}
func addControllerRoleToSA(saNamespace, saName string, role rbacv1.ClusterRole) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(role.Name, saRolePrefix) {
		klog.Fatalf(`role %q must start with %q`, role.Name, saRolePrefix)
	}
	for _, existingRole := range controllerRoles {
		if role.Name == existingRole.Name {
			klog.Fatalf("role %q was already registered", role.Name)
		}
	}
	addDefaultMetadata(&role)
	controllerRoles = append(controllerRoles, role)
	roleBinding := rbacv1helpers.NewClusterBinding(role.Name).SAs(saNamespace, saName).BindingOrDie()
	addDefaultMetadata(&roleBinding)
	controllerRoleBindings = append(controllerRoleBindings, roleBinding)
}
func eventsRule() rbacv1.PolicyRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rbacv1helpers.NewRule("create", "update", "patch").Groups(kapiGroup).Resources("events").RuleOrDie()
}
func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraBuildControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "patch", "update", "delete").Groups(buildGroup, legacyBuildGroup).Resources("builds").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(buildGroup, legacyBuildGroup).Resources("builds/finalizers").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(buildGroup, legacyBuildGroup).Resources("buildconfigs").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(buildGroup, legacyBuildGroup).Resources("builds/optimizeddocker", "builds/docker", "builds/source", "builds/custom", "builds/jenkinspipeline").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(imageGroup, legacyImageGroup).Resources("imagestreams").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(kapiGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "create").Groups(kapiGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "create", "delete").Groups(kapiGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(kapiGroup).Resources("namespaces").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(kapiGroup).Resources("serviceaccounts").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(securityGroup, legacySecurityGroup).Resources("podsecuritypolicysubjectreviews").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(configGroup).Resources("builds").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraBuildConfigChangeControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(buildGroup, legacyBuildGroup).Resources("buildconfigs").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(buildGroup, legacyBuildGroup).Resources("buildconfigs/instantiate").RuleOrDie(), rbacv1helpers.NewRule("delete").Groups(buildGroup, legacyBuildGroup).Resources("builds").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraDeployerControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create", "get", "list", "watch", "patch", "delete").Groups(kapiGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("delete").Groups(kapiGroup).Resources("replicationcontrollers").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "update").Groups(kapiGroup).Resources("replicationcontrollers").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(kapiGroup).Resources("replicationcontrollers/scale").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraDeploymentConfigControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create", "get", "list", "watch", "update", "patch", "delete").Groups(kapiGroup).Resources("replicationcontrollers").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(kapiGroup).Resources("replicationcontrollers/scale").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs/status").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs/finalizers").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraTemplateInstanceControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(kAuthzGroup).Resources("subjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(templateGroup).Resources("templateinstances/status").RuleOrDie()}})
	templateInstanceController := rbacv1helpers.NewClusterBinding(AdminRoleName).SAs(DefaultOpenShiftInfraNamespace, InfraTemplateInstanceControllerServiceAccountName).BindingOrDie()
	templateInstanceController.Name = "system:openshift:controller:" + InfraTemplateInstanceControllerServiceAccountName + ":admin"
	addDefaultMetadata(&templateInstanceController)
	controllerRoleBindings = append(controllerRoleBindings, templateInstanceController)
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraTemplateInstanceFinalizerControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("update").Groups(templateGroup).Resources("templateinstances/status").RuleOrDie()}})
	templateInstanceFinalizerController := rbacv1helpers.NewClusterBinding(AdminRoleName).SAs(DefaultOpenShiftInfraNamespace, InfraTemplateInstanceFinalizerControllerServiceAccountName).BindingOrDie()
	templateInstanceFinalizerController.Name = "system:openshift:controller:" + InfraTemplateInstanceFinalizerControllerServiceAccountName + ":admin"
	addDefaultMetadata(&templateInstanceFinalizerController)
	controllerRoleBindings = append(controllerRoleBindings, templateInstanceFinalizerController)
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraOriginNamespaceServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("namespaces").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("namespaces/finalize", "namespaces/status").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraServiceAccountControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch", "delete").Groups(kapiGroup).Resources("serviceaccounts").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraServiceAccountPullSecretsControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "create", "update").Groups(kapiGroup).Resources("serviceaccounts").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch", "delete").Groups(kapiGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("services").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraImageTriggerControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch").Groups(imageGroup, legacyImageGroup).Resources("imagestreams").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(extensionsGroup).Resources("daemonsets").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(extensionsGroup, appsGroup).Resources("deployments").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(appsGroup).Resources("statefulsets").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(batchGroup).Resources("cronjobs").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(buildGroup, legacyBuildGroup).Resources("buildconfigs/instantiate").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(buildGroup, legacyBuildGroup).Resources(SourceBuildResource, DockerBuildResource, CustomBuildResource, OptimizedDockerBuildResource, JenkinsPipelineBuildResource).RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraServiceServingCertServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch", "update").Groups(kapiGroup).Resources("services").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "delete").Groups(kapiGroup).Resources("secrets").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraImageImportControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch", "create", "update").Groups(imageGroup, legacyImageGroup).Resources("imagestreams").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch", "delete").Groups(imageGroup, legacyImageGroup).Resources("images").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(imageGroup, legacyImageGroup).Resources("imagestreamimports").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraSDNControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "create", "update").Groups(networkGroup, legacyNetworkGroup).Resources("clusternetworks").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "delete").Groups(networkGroup, legacyNetworkGroup).Resources("hostsubnets").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "delete").Groups(networkGroup, legacyNetworkGroup).Resources("netnamespaces").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(kapiGroup).Resources("pods").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("services").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("namespaces").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("nodes").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("nodes/status").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraClusterQuotaReconciliationControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list").Groups(kapiGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("get", "list").Groups(kapiGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(quotaGroup, legacyQuotaGroup).Resources("clusterresourcequotas/status").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraUnidlingControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "update").Groups(kapiGroup).Resources("replicationcontrollers/scale", "endpoints").RuleOrDie(), rbacv1helpers.NewRule("get", "update", "patch").Groups(kapiGroup).Resources("replicationcontrollers").RuleOrDie(), rbacv1helpers.NewRule("get", "update", "patch").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(extensionsGroup, appsGroup).Resources("replicasets/scale", "deployments/scale").RuleOrDie(), rbacv1helpers.NewRule("get", "update").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs/scale").RuleOrDie(), rbacv1helpers.NewRule("watch", "list").Groups(kapiGroup).Resources("events").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraServiceIngressIPControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch", "update").Groups(kapiGroup).Resources("services").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("services/status").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraIngressToRouteControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("secrets", "services").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(extensionsGroup).Resources("ingress").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "create", "update", "patch", "delete").Groups(routeGroup).Resources("routes").RuleOrDie(), rbacv1helpers.NewRule("create", "update").Groups(routeGroup).Resources("routes/custom-host").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraPersistentVolumeRecyclerControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "update", "create", "delete", "list", "watch").Groups(kapiGroup).Resources("persistentvolumes").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("persistentvolumes/status").RuleOrDie(), rbacv1helpers.NewRule("get", "update", "list", "watch").Groups(kapiGroup).Resources("persistentvolumeclaims").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("persistentvolumeclaims/status").RuleOrDie(), rbacv1helpers.NewRule("get", "create", "delete", "list", "watch").Groups(kapiGroup).Resources("pods").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraResourceQuotaControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("update").Groups(kapiGroup).Resources("resourcequotas/status").RuleOrDie(), rbacv1helpers.NewRule("list").Groups(kapiGroup).Resources("resourcequotas").RuleOrDie(), rbacv1helpers.NewRule("list").Groups(kapiGroup).Resources("services").RuleOrDie(), rbacv1helpers.NewRule("list").Groups(kapiGroup).Resources("configmaps").RuleOrDie(), rbacv1helpers.NewRule("list").Groups(kapiGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("list").Groups(kapiGroup).Resources("replicationcontrollers").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraHorizontalPodAutoscalerControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "update").Groups(deployGroup, legacyDeployGroup).Resources("deploymentconfigs/scale").RuleOrDie()}})
	bindControllerRole(InfraHorizontalPodAutoscalerControllerServiceAccountName, "system:controller:horizontal-pod-autoscaler")
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraTemplateServiceBrokerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(kAuthzGroup).Resources("subjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authzGroup).Resources("subjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule("get", "create", "update", "delete").Groups(templateGroup).Resources("brokertemplateinstances").RuleOrDie(), rbacv1helpers.NewRule("update").Groups(templateGroup).Resources("brokertemplateinstances/finalizers").RuleOrDie(), rbacv1helpers.NewRule("get", "create", "delete", "assign").Groups(templateGroup).Resources("templateinstances").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(templateGroup).Resources("templates").RuleOrDie(), rbacv1helpers.NewRule("get", "create", "delete").Groups(kapiGroup).Resources("secrets").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(kapiGroup).Resources("services", "configmaps").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(legacyRouteGroup).Resources("routes").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(routeGroup).Resources("routes").RuleOrDie(), eventsRule()}})
	bindControllerRole(InfraDefaultRoleBindingsControllerServiceAccountName, ImagePullerRoleName)
	bindControllerRole(InfraDefaultRoleBindingsControllerServiceAccountName, ImageBuilderRoleName)
	bindControllerRole(InfraDefaultRoleBindingsControllerServiceAccountName, DeployerRoleName)
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraDefaultRoleBindingsControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(rbacGroup).Resources("rolebindings").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(kapiGroup).Resources("namespaces").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch").Groups(rbacGroup).Resources("rolebindings").RuleOrDie(), eventsRule()}})
	addControllerRole(rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: saRolePrefix + InfraNamespaceSecurityAllocationControllerServiceAccountName}, Rules: []rbacv1.PolicyRule{rbacv1helpers.NewRule("get", "create", "update").Groups(securityGroup).Resources("rangeallocations").RuleOrDie(), rbacv1helpers.NewRule("get", "list", "watch", "update").Groups(kapiGroup).Resources("namespaces").RuleOrDie(), eventsRule()}})
}
func ControllerRoles() []rbacv1.ClusterRole {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return controllerRoles
}
func ControllerRoleBindings() []rbacv1.ClusterRoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return controllerRoleBindings
}
