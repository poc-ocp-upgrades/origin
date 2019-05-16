package kubeadmission

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/customresourcevalidationregistration"
	authorizationrestrictusers "github.com/openshift/origin/pkg/authorization/apiserver/admission/restrictusers"
	quotaclusterresourceoverride "github.com/openshift/origin/pkg/autoscaling/admission/clusterresourceoverride"
	quotarunonceduration "github.com/openshift/origin/pkg/autoscaling/admission/runonceduration"
	imagepolicyapiv1 "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
	"github.com/openshift/origin/pkg/image/apiserver/admission/imagepolicy"
	"github.com/openshift/origin/pkg/network/admission/externalipranger"
	"github.com/openshift/origin/pkg/network/admission/restrictedendpoints"
	quotaclusterresourcequota "github.com/openshift/origin/pkg/quota/apiserver/admission/clusterresourcequota"
	ingressadmission "github.com/openshift/origin/pkg/route/apiserver/admission"
	projectnodeenv "github.com/openshift/origin/pkg/scheduler/admission/nodeenv"
	schedulerpodnodeconstraints "github.com/openshift/origin/pkg/scheduler/admission/podnodeconstraints"
	securityadmission "github.com/openshift/origin/pkg/security/apiserver/admission/sccadmission"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	mutatingwebhook "k8s.io/apiserver/pkg/admission/plugin/webhook/mutating"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func RegisterOpenshiftKubeAdmissionPlugins(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authorizationrestrictusers.Register(plugins)
	imagepolicy.Register(plugins)
	ingressadmission.Register(plugins)
	projectnodeenv.Register(plugins)
	quotaclusterresourceoverride.Register(plugins)
	quotaclusterresourcequota.Register(plugins)
	quotarunonceduration.Register(plugins)
	schedulerpodnodeconstraints.Register(plugins)
	securityadmission.Register(plugins)
	securityadmission.RegisterSCCExecRestrictions(plugins)
	externalipranger.RegisterExternalIP(plugins)
	restrictedendpoints.RegisterRestrictedEndpoints(plugins)
}

var (
	SkipRunLevelZeroPlugins          = sets.NewString()
	SkipRunLevelOnePlugins           = sets.NewString(imagepolicyapiv1.PluginName, "quota.openshift.io/ClusterResourceQuota", "security.openshift.io/SecurityContextConstraint", "security.openshift.io/SCCExecRestrictions")
	openshiftAdmissionPluginsForKube = []string{"autoscaling.openshift.io/ClusterResourceOverride", "authorization.openshift.io/RestrictSubjectBindings", "autoscaling.openshift.io/RunOnceDuration", "scheduling.openshift.io/PodNodeConstraints", "scheduling.openshift.io/OriginPodNodeEnvironment", "network.openshift.io/ExternalIPRanger", "network.openshift.io/RestrictedEndpointsAdmission", imagepolicyapiv1.PluginName, "security.openshift.io/SecurityContextConstraint", "security.openshift.io/SCCExecRestrictions", "route.openshift.io/IngressAdmission", "quota.openshift.io/ClusterResourceQuota"}
	additionalDefaultOnPlugins       = sets.NewString("NodeRestriction", "OwnerReferencesPermissionEnforcement", "PersistentVolumeLabel", "PodNodeSelector", "PodTolerationRestriction", "Priority", imagepolicyapiv1.PluginName, "StorageObjectInUseProtection")
)

func NewOrderedKubeAdmissionPlugins(kubeAdmissionOrder []string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := []string{}
	for _, curr := range kubeAdmissionOrder {
		if curr == mutatingwebhook.PluginName {
			ret = append(ret, openshiftAdmissionPluginsForKube...)
			ret = append(ret, customresourcevalidationregistration.AllCustomResourceValidators...)
		}
		ret = append(ret, curr)
	}
	return ret
}
func NewDefaultOffPluginsFunc(kubeDefaultOffAdmission sets.String) func() sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() sets.String {
		kubeOff := sets.NewString(kubeDefaultOffAdmission.UnsortedList()...)
		kubeOff.Delete(additionalDefaultOnPlugins.List()...)
		kubeOff.Delete(openshiftAdmissionPluginsForKube...)
		kubeOff.Delete(customresourcevalidationregistration.AllCustomResourceValidators...)
		return kubeOff
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
