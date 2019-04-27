package bootstrappolicy

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	securityapiv1 "github.com/openshift/api/security/v1"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityapiinstall "github.com/openshift/origin/pkg/security/apis/security/install"
)

const (
	SecurityContextConstraintPrivileged		= "privileged"
	SecurityContextConstraintPrivilegedDesc		= "privileged allows access to all privileged and host features and the ability to run as any user, any group, any fsGroup, and with any SELinux context.  WARNING: this is the most relaxed SCC and should be used only for cluster administration. Grant with caution."
	SecurityContextConstraintRestricted		= "restricted"
	SecurityContextConstraintRestrictedDesc		= "restricted denies access to all host features and requires pods to be run with a UID, and SELinux context that are allocated to the namespace.  This is the most restrictive SCC and it is used by default for authenticated users."
	SecurityContextConstraintNonRoot		= "nonroot"
	SecurityContextConstraintNonRootDesc		= "nonroot provides all features of the restricted SCC but allows users to run with any non-root UID.  The user must specify the UID or it must be specified on the by the manifest of the container runtime."
	SecurityContextConstraintHostMountAndAnyUID	= "hostmount-anyuid"
	SecurityContextConstraintHostMountAndAnyUIDDesc	= "hostmount-anyuid provides all the features of the restricted SCC but allows host mounts and any UID by a pod.  This is primarily used by the persistent volume recycler. WARNING: this SCC allows host file system access as any UID, including UID 0.  Grant with caution."
	SecurityContextConstraintHostNS			= "hostaccess"
	SecurityContextConstraintHostNSDesc		= "hostaccess allows access to all host namespaces but still requires pods to be run with a UID and SELinux context that are allocated to the namespace. WARNING: this SCC allows host access to namespaces, file systems, and PIDS.  It should only be used by trusted pods.  Grant with caution."
	SecurityContextConstraintsAnyUID		= "anyuid"
	SecurityContextConstraintsAnyUIDDesc		= "anyuid provides all features of the restricted SCC but allows users to run with any UID and any GID."
	SecurityContextConstraintsHostNetwork		= "hostnetwork"
	SecurityContextConstraintsHostNetworkDesc	= "hostnetwork allows using host networking and host ports but still requires pods to be run with a UID and SELinux context that are allocated to the namespace."
	DescriptionAnnotation				= "kubernetes.io/description"
)

var bootstrapSCCScheme = runtime.NewScheme()

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	securityapiinstall.Install(bootstrapSCCScheme)
}
func GetBootstrapSecurityContextConstraints(sccNameToAdditionalGroups map[string][]string, sccNameToAdditionalUsers map[string][]string) []*securityapi.SecurityContextConstraints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		securityContextConstraintsAnyUIDPriority = int32(10)
	)
	constraints := []*securityapi.SecurityContextConstraints{{ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintPrivileged, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintPrivilegedDesc}}, AllowPrivilegedContainer: true, AllowedCapabilities: []kapi.Capability{securityapi.AllowAllCapabilities}, Volumes: []securityapi.FSType{securityapi.FSTypeAll}, AllowHostNetwork: true, AllowHostPorts: true, AllowHostPID: true, AllowHostIPC: true, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyRunAsAny}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyRunAsAny}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyRunAsAny}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, SeccompProfiles: []string{"*"}, AllowedUnsafeSysctls: []string{"*"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintNonRoot, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintNonRootDesc}}, Volumes: []securityapi.FSType{securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSProjected}, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyMustRunAsNonRoot}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyRunAsAny}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, RequiredDropCapabilities: []kapi.Capability{"KILL", "MKNOD", "SETUID", "SETGID"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintHostMountAndAnyUID, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintHostMountAndAnyUIDDesc}}, Volumes: []securityapi.FSType{securityapi.FSTypeHostPath, securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSTypeNFS, securityapi.FSProjected}, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyRunAsAny}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyRunAsAny}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, RequiredDropCapabilities: []kapi.Capability{"MKNOD"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintHostNS, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintHostNSDesc}}, Volumes: []securityapi.FSType{securityapi.FSTypeHostPath, securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSProjected}, AllowHostNetwork: true, AllowHostPorts: true, AllowHostPID: true, AllowHostIPC: true, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyMustRunAsRange}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyMustRunAs}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, RequiredDropCapabilities: []kapi.Capability{"KILL", "MKNOD", "SETUID", "SETGID"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintRestricted, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintRestrictedDesc}}, Volumes: []securityapi.FSType{securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSProjected}, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyMustRunAsRange}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyMustRunAs}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, RequiredDropCapabilities: []kapi.Capability{"KILL", "MKNOD", "SETUID", "SETGID"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintsAnyUID, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintsAnyUIDDesc}}, Volumes: []securityapi.FSType{securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSProjected}, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyRunAsAny}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyRunAsAny}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyRunAsAny}, Priority: &securityContextConstraintsAnyUIDPriority, RequiredDropCapabilities: []kapi.Capability{"MKNOD"}}, {ObjectMeta: metav1.ObjectMeta{Name: SecurityContextConstraintsHostNetwork, Annotations: map[string]string{DescriptionAnnotation: SecurityContextConstraintsHostNetworkDesc}}, AllowHostNetwork: true, AllowHostPorts: true, Volumes: []securityapi.FSType{securityapi.FSTypeEmptyDir, securityapi.FSTypeSecret, securityapi.FSTypeDownwardAPI, securityapi.FSTypeConfigMap, securityapi.FSTypePersistentVolumeClaim, securityapi.FSProjected}, SELinuxContext: securityapi.SELinuxContextStrategyOptions{Type: securityapi.SELinuxStrategyMustRunAs}, RunAsUser: securityapi.RunAsUserStrategyOptions{Type: securityapi.RunAsUserStrategyMustRunAsRange}, FSGroup: securityapi.FSGroupStrategyOptions{Type: securityapi.FSGroupStrategyMustRunAs}, SupplementalGroups: securityapi.SupplementalGroupsStrategyOptions{Type: securityapi.SupplementalGroupsStrategyMustRunAs}, RequiredDropCapabilities: []kapi.Capability{"KILL", "MKNOD", "SETUID", "SETGID"}}}
	for i := range constraints {
		v1constraint := &securityapiv1.SecurityContextConstraints{}
		constraint := &securityapi.SecurityContextConstraints{}
		if err := bootstrapSCCScheme.Convert(constraints[i], v1constraint, nil); err != nil {
			panic(err)
		}
		bootstrapSCCScheme.Default(v1constraint)
		if err := bootstrapSCCScheme.Convert(v1constraint, constraint, nil); err != nil {
			panic(err)
		}
		constraints[i] = constraint
		if usersToAdd, ok := sccNameToAdditionalUsers[constraint.Name]; ok {
			constraints[i].Users = append(constraints[i].Users, usersToAdd...)
		}
		if groupsToAdd, ok := sccNameToAdditionalGroups[constraint.Name]; ok {
			constraints[i].Groups = append(constraints[i].Groups, groupsToAdd...)
		}
	}
	return constraints
}
func GetBoostrapSCCAccess(infraNamespace string) (map[string][]string, map[string][]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	groups := map[string][]string{SecurityContextConstraintPrivileged: {ClusterAdminGroup, NodesGroup, MastersGroup}, SecurityContextConstraintsAnyUID: {ClusterAdminGroup}, SecurityContextConstraintRestricted: {AuthenticatedGroup}}
	buildControllerUsername := serviceaccount.MakeUsername(infraNamespace, InfraBuildControllerServiceAccountName)
	pvRecyclerControllerUsername := serviceaccount.MakeUsername(infraNamespace, InfraPersistentVolumeRecyclerControllerServiceAccountName)
	users := map[string][]string{SecurityContextConstraintPrivileged: {SystemAdminUsername, buildControllerUsername}, SecurityContextConstraintHostMountAndAnyUID: {pvRecyclerControllerUsername}}
	return groups, users
}
