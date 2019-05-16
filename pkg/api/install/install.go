package install

import (
	goformat "fmt"
	oapps "github.com/openshift/origin/pkg/apps/apis/apps/install"
	authz "github.com/openshift/origin/pkg/authorization/apis/authorization/install"
	build "github.com/openshift/origin/pkg/build/apis/build/install"
	_ "github.com/openshift/origin/pkg/cmd/server/apis/config/install"
	image "github.com/openshift/origin/pkg/image/apis/image/install"
	network "github.com/openshift/origin/pkg/network/apis/network/install"
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth/install"
	project "github.com/openshift/origin/pkg/project/apis/project/install"
	quota "github.com/openshift/origin/pkg/quota/apis/quota/install"
	route "github.com/openshift/origin/pkg/route/apis/route/install"
	security "github.com/openshift/origin/pkg/security/apis/security/install"
	template "github.com/openshift/origin/pkg/template/apis/template/install"
	user "github.com/openshift/origin/pkg/user/apis/user/install"
	crdinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	"k8s.io/apimachinery/pkg/runtime"
	apiregistrationinstall "k8s.io/kube-aggregator/pkg/apis/apiregistration/install"
	kcomponentconfiginstall "k8s.io/kubernetes/cmd/cloud-controller-manager/app/apis/config/scheme"
	kadmissioninstall "k8s.io/kubernetes/pkg/apis/admission/install"
	kadmissionregistrationinstall "k8s.io/kubernetes/pkg/apis/admissionregistration/install"
	kappsinstall "k8s.io/kubernetes/pkg/apis/apps/install"
	kauthenticationinstall "k8s.io/kubernetes/pkg/apis/authentication/install"
	kauthorizationinstall "k8s.io/kubernetes/pkg/apis/authorization/install"
	kautoscalinginstall "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	kbatchinstall "k8s.io/kubernetes/pkg/apis/batch/install"
	kcertificatesinstall "k8s.io/kubernetes/pkg/apis/certificates/install"
	kcoreinstall "k8s.io/kubernetes/pkg/apis/core/install"
	keventsinstall "k8s.io/kubernetes/pkg/apis/events/install"
	kextensionsinstall "k8s.io/kubernetes/pkg/apis/extensions/install"
	kimagepolicyinstall "k8s.io/kubernetes/pkg/apis/imagepolicy/install"
	knetworkinginstall "k8s.io/kubernetes/pkg/apis/networking/install"
	kpolicyinstall "k8s.io/kubernetes/pkg/apis/policy/install"
	krbacinstall "k8s.io/kubernetes/pkg/apis/rbac/install"
	kschedulinginstall "k8s.io/kubernetes/pkg/apis/scheduling/install"
	ksettingsinstall "k8s.io/kubernetes/pkg/apis/settings/install"
	kstorageinstall "k8s.io/kubernetes/pkg/apis/storage/install"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func InstallInternalOpenShift(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oapps.Install(scheme)
	authz.Install(scheme)
	build.Install(scheme)
	image.Install(scheme)
	network.Install(scheme)
	oauth.Install(scheme)
	project.Install(scheme)
	quota.Install(scheme)
	route.Install(scheme)
	security.Install(scheme)
	template.Install(scheme)
	user.Install(scheme)
}
func InstallInternalKube(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	crdinstall.Install(scheme)
	apiregistrationinstall.Install(scheme)
	kadmissioninstall.Install(scheme)
	kadmissionregistrationinstall.Install(scheme)
	kappsinstall.Install(scheme)
	kauthenticationinstall.Install(scheme)
	kauthorizationinstall.Install(scheme)
	kautoscalinginstall.Install(scheme)
	kbatchinstall.Install(scheme)
	kcertificatesinstall.Install(scheme)
	kcomponentconfiginstall.AddToScheme(scheme)
	kcoreinstall.Install(scheme)
	keventsinstall.Install(scheme)
	kextensionsinstall.Install(scheme)
	kimagepolicyinstall.Install(scheme)
	knetworkinginstall.Install(scheme)
	kpolicyinstall.Install(scheme)
	krbacinstall.Install(scheme)
	kschedulinginstall.Install(scheme)
	ksettingsinstall.Install(scheme)
	kstorageinstall.Install(scheme)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
