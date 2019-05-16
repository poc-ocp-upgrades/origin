package legacy

import (
	goformat "fmt"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/origin/pkg/apps/apis/apps"
	appsv1helpers "github.com/openshift/origin/pkg/apps/apis/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/apis/extensions"
	extensionsv1beta1conversions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func InstallInternalLegacyApps(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	InstallExternalLegacyApps(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalAppsTypes, core.AddToScheme, extensions.AddToScheme, corev1conversions.AddToScheme, extensionsv1beta1conversions.AddToScheme, appsv1helpers.AddConversionFuncs, appsv1helpers.RegisterDefaults, appsv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyApps(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedAppsTypes, corev1.AddToScheme, rbacv1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedAppsTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	types := []runtime.Object{&appsv1.DeploymentConfig{}, &appsv1.DeploymentConfigList{}, &appsv1.DeploymentConfigRollback{}, &appsv1.DeploymentRequest{}, &appsv1.DeploymentLog{}, &appsv1.DeploymentLogOptions{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalAppsTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(InternalGroupVersion, &apps.DeploymentConfig{}, &apps.DeploymentConfigList{}, &apps.DeploymentConfigRollback{}, &apps.DeploymentRequest{}, &apps.DeploymentLog{}, &apps.DeploymentLogOptions{})
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
