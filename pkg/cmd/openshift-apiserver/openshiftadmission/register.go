package openshiftadmission

import (
	buildsecretinjector "github.com/openshift/origin/pkg/build/apiserver/admission/secretinjector"
	buildstrategyrestrictions "github.com/openshift/origin/pkg/build/apiserver/admission/strategyrestrictions"
	"github.com/openshift/origin/pkg/image/apiserver/admission/imagepolicy"
	imageadmission "github.com/openshift/origin/pkg/image/apiserver/admission/limitrange"
	projectrequestlimit "github.com/openshift/origin/pkg/project/apiserver/admission/requestlimit"
	quotaclusterresourcequota "github.com/openshift/origin/pkg/quota/apiserver/admission/clusterresourcequota"
	schedulerpodnodeconstraints "github.com/openshift/origin/pkg/scheduler/admission/podnodeconstraints"
	"k8s.io/apiserver/pkg/admission"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/kubernetes/plugin/pkg/admission/gc"
	"k8s.io/kubernetes/plugin/pkg/admission/resourcequota"
)

var OriginAdmissionPlugins = admission.NewPlugins()

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	RegisterAllAdmissionPlugins(OriginAdmissionPlugins)
}
func RegisterAllAdmissionPlugins(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gc.Register(plugins)
	resourcequota.Register(plugins)
	genericapiserver.RegisterAllAdmissionPlugins(plugins)
	RegisterOpenshiftAdmissionPlugins(plugins)
}
func RegisterOpenshiftAdmissionPlugins(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projectrequestlimit.Register(plugins)
	buildsecretinjector.Register(plugins)
	buildstrategyrestrictions.Register(plugins)
	imageadmission.Register(plugins)
	imagepolicy.Register(plugins)
	schedulerpodnodeconstraints.Register(plugins)
	quotaclusterresourcequota.Register(plugins)
}

var (
	OpenShiftAdmissionPlugins = []string{"NamespaceLifecycle", "OwnerReferencesPermissionEnforcement", "project.openshift.io/ProjectRequestLimit", "build.openshift.io/BuildConfigSecretInjector", "build.openshift.io/BuildByStrategy", "image.openshift.io/ImageLimitRange", "image.openshift.io/ImagePolicy", "quota.openshift.io/ClusterResourceQuota", "MutatingAdmissionWebhook", "ValidatingAdmissionWebhook", "ResourceQuota"}
)
