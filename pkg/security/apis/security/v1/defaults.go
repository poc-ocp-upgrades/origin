package v1

import (
	"github.com/openshift/api/security/v1"
	sccutil "github.com/openshift/origin/pkg/security/securitycontextconstraints/util"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
)

func AddDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RegisterDefaults(scheme)
	scheme.AddTypeDefaultingFunc(&v1.SecurityContextConstraints{}, func(obj interface{}) {
		SetDefaults_SCC(obj.(*v1.SecurityContextConstraints))
	})
	return nil
}
func SetDefaults_SCC(scc *v1.SecurityContextConstraints) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(scc.FSGroup.Type) == 0 {
		scc.FSGroup.Type = v1.FSGroupStrategyRunAsAny
	}
	if len(scc.SupplementalGroups.Type) == 0 {
		scc.SupplementalGroups.Type = v1.SupplementalGroupsStrategyRunAsAny
	}
	if scc.Users == nil {
		scc.Users = []string{}
	}
	if scc.Groups == nil {
		scc.Groups = []string{}
	}
	var defaultAllowedVolumes sets.String
	switch {
	case scc.Volumes == nil:
		defaultAllowedVolumes = sets.NewString(string(v1.FSTypeAll))
	case len(scc.Volumes) == 0 && scc.AllowHostDirVolumePlugin:
		defaultAllowedVolumes = sets.NewString(string(v1.FSTypeHostPath))
	case len(scc.Volumes) == 0 && !scc.AllowHostDirVolumePlugin:
		defaultAllowedVolumes = sets.NewString(string(v1.FSTypeNone))
	default:
		defaultAllowedVolumes = fsTypeToStringSet(scc.Volumes)
	}
	if scc.AllowHostDirVolumePlugin {
		if !defaultAllowedVolumes.Has(string(v1.FSTypeAll)) {
			defaultAllowedVolumes.Insert(string(v1.FSTypeHostPath))
		}
	} else {
		shouldDefaultAllVolumes := defaultAllowedVolumes.Has(string(v1.FSTypeAll))
		defaultAllowedVolumes.Delete(string(v1.FSTypeAll))
		defaultAllowedVolumes.Delete(string(v1.FSTypeHostPath))
		if shouldDefaultAllVolumes {
			allVolumes := sccutil.GetAllFSTypesExcept(string(v1.FSTypeHostPath))
			defaultAllowedVolumes.Insert(allVolumes.List()...)
		}
	}
	scc.Volumes = StringSetToFSType(defaultAllowedVolumes)
	if scc.AllowPrivilegeEscalation == nil {
		t := true
		scc.AllowPrivilegeEscalation = &t
	}
}
func StringSetToFSType(set sets.String) []v1.FSType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if set == nil {
		return nil
	}
	volumes := []v1.FSType{}
	for _, v := range set.List() {
		volumes = append(volumes, v1.FSType(v))
	}
	return volumes
}
func fsTypeToStringSet(volumes []v1.FSType) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if volumes == nil {
		return nil
	}
	set := sets.NewString()
	for _, v := range volumes {
		set.Insert(string(v))
	}
	return set
}
