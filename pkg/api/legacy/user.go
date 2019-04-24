package legacy

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	userv1 "github.com/openshift/api/user/v1"
	"github.com/openshift/origin/pkg/user/apis/user"
	userv1helpers "github.com/openshift/origin/pkg/user/apis/user/v1"
	"k8s.io/kubernetes/pkg/apis/core"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
)

func InstallInternalLegacyUser(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyUser(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalUserTypes, core.AddToScheme, corev1conversions.AddToScheme, addLegacyUserFieldSelectorKeyConversions, userv1helpers.RegisterDefaults, userv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyUser(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedUserTypes, corev1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedUserTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&userv1.User{}, &userv1.UserList{}, &userv1.Identity{}, &userv1.IdentityList{}, &userv1.UserIdentityMapping{}, &userv1.Group{}, &userv1.GroupList{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalUserTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &user.User{}, &user.UserList{}, &user.Identity{}, &user.IdentityList{}, &user.UserIdentityMapping{}, &user.Group{}, &user.GroupList{})
	return nil
}
func addLegacyUserFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("Identity"), legacyIdentityFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func legacyIdentityFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "providerName", "providerUserName", "user.name", "user.uid":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
