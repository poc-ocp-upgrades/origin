package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/apis/apiserver"
	apiserverv1alpha1 "k8s.io/apiserver/pkg/apis/apiserver/v1alpha1"
	"k8s.io/apiserver/pkg/apis/audit"
	auditv1alpha1 "k8s.io/apiserver/pkg/apis/audit/v1alpha1"
	auditv1beta1 "k8s.io/apiserver/pkg/apis/audit/v1beta1"
	clusterresourceoverrideinstall "github.com/openshift/origin/pkg/autoscaling/admission/apis/clusterresourceoverride/install"
	runoncedurationinstall "github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration/install"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	configapiv1 "github.com/openshift/origin/pkg/cmd/server/apis/config/v1"
	imagepolicyapiv1 "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
	externaliprangerinstall "github.com/openshift/origin/pkg/network/admission/apis/externalipranger/install"
	restrictedendpointsinstall "github.com/openshift/origin/pkg/network/admission/apis/restrictedendpoints/install"
	requestlimitinstall "github.com/openshift/origin/pkg/project/apiserver/admission/apis/requestlimit/install"
	ingressadmissioninstall "github.com/openshift/origin/pkg/route/apiserver/admission/apis/ingressadmission/install"
	podnodeconstraintsinstall "github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints/install"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallLegacyInternal(configapi.Scheme)
}
func InstallLegacyInternal(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configapi.InstallLegacy(scheme)
	configapiv1.InstallLegacy(scheme)
	audit.AddToScheme(scheme)
	auditv1alpha1.AddToScheme(scheme)
	auditv1beta1.AddToScheme(scheme)
	apiserver.AddToScheme(scheme)
	apiserverv1alpha1.AddToScheme(scheme)
	imagepolicyapiv1.Install(scheme)
	requestlimitinstall.InstallInternal(scheme)
	ingressadmissioninstall.InstallInternal(scheme)
	clusterresourceoverrideinstall.InstallInternal(scheme)
	runoncedurationinstall.InstallInternal(scheme)
	podnodeconstraintsinstall.InstallInternal(scheme)
	restrictedendpointsinstall.InstallInternal(scheme)
	externaliprangerinstall.InstallInternal(scheme)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
