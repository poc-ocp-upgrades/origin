package legacy

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupName            = ""
	GroupVersion         = schema.GroupVersion{Group: GroupName, Version: "v1"}
	InternalGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
)

func Kind(kind string) schema.GroupKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.GroupKind{Group: GroupName, Kind: kind}
}
func GroupVersionKind(kind string) schema.GroupVersionKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.GroupVersionKind{Group: GroupName, Version: GroupVersion.Version, Kind: kind}
}
func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.GroupResource{Group: GroupName, Resource: resource}
}
func InstallInternalLegacyAll(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallInternalLegacyApps(scheme)
	InstallInternalLegacyAuthorization(scheme)
	InstallInternalLegacyBuild(scheme)
	InstallInternalLegacyImage(scheme)
	InstallInternalLegacyNetwork(scheme)
	InstallInternalLegacyOAuth(scheme)
	InstallInternalLegacyProject(scheme)
	InstallInternalLegacyQuota(scheme)
	InstallInternalLegacyRoute(scheme)
	InstallInternalLegacySecurity(scheme)
	InstallInternalLegacyTemplate(scheme)
	InstallInternalLegacyUser(scheme)
}
func InstallExternalLegacyAll(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyApps(scheme)
	InstallExternalLegacyAuthorization(scheme)
	InstallExternalLegacyBuild(scheme)
	InstallExternalLegacyImage(scheme)
	InstallExternalLegacyNetwork(scheme)
	InstallExternalLegacyOAuth(scheme)
	InstallExternalLegacyProject(scheme)
	InstallExternalLegacyQuota(scheme)
	InstallExternalLegacyRoute(scheme)
	InstallExternalLegacySecurity(scheme)
	InstallExternalLegacyTemplate(scheme)
	InstallExternalLegacyUser(scheme)
}
