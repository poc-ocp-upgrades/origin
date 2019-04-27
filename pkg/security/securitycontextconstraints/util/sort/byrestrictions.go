package sort

import (
	"strings"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	securityv1 "github.com/openshift/api/security/v1"
)

type ByRestrictions []*securityv1.SecurityContextConstraints

func (s ByRestrictions) Len() int {
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
	return len(s)
}
func (s ByRestrictions) Swap(i, j int) {
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
	s[i], s[j] = s[j], s[i]
}
func (s ByRestrictions) Less(i, j int) bool {
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
	return pointValue(s[i]) < pointValue(s[j])
}

type points int

const (
	privilegedPoints	points	= 1000000
	hostNetworkPoints	points	= 200000
	hostPortsPoints		points	= 400000
	hostVolumePoints	points	= 100000
	nonTrivialVolumePoints	points	= 50000
	runAsAnyUserPoints	points	= 40000
	runAsNonRootPoints	points	= 30000
	runAsRangePoints	points	= 20000
	runAsUserPoints		points	= 10000
	capDefaultPoints	points	= 5000
	capAddOnePoints		points	= 300
	capAllowAllPoints	points	= 4000
	capAllowOnePoints	points	= 10
	capDropAllPoints	points	= -3000
	capDropOnePoints	points	= -50
	capMaxPoints		points	= 9999
	capMinPoints		points	= 0
	noPoints		points	= 0
)

func pointValue(constraint *securityv1.SecurityContextConstraints) points {
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
	totalPoints := noPoints
	if constraint.AllowPrivilegedContainer {
		totalPoints += privilegedPoints
	}
	totalPoints += volumePointValue(constraint)
	if constraint.AllowHostNetwork {
		totalPoints += hostNetworkPoints
	}
	if constraint.AllowHostPorts {
		totalPoints += hostPortsPoints
	}
	totalPoints += capabilitiesPointValue(constraint)
	strategiesPoints := map[string]points{string(securityv1.RunAsUserStrategyRunAsAny): runAsAnyUserPoints, string(securityv1.RunAsUserStrategyMustRunAsNonRoot): runAsNonRootPoints, string(securityv1.RunAsUserStrategyMustRunAsRange): runAsRangePoints, string(securityv1.RunAsUserStrategyMustRunAs): runAsUserPoints}
	strategyType := string(constraint.SELinuxContext.Type)
	points, found := strategiesPoints[strategyType]
	if found {
		totalPoints += points
	} else {
		klog.Warningf("SELinuxContext type %q has no point value, this may cause issues in sorting SCCs by restriction", strategyType)
	}
	strategyType = string(constraint.RunAsUser.Type)
	points, found = strategiesPoints[strategyType]
	if found {
		totalPoints += points
	} else {
		klog.Warningf("RunAsUser type %q has no point value, this may cause issues in sorting SCCs by restriction", strategyType)
	}
	return totalPoints
}
func volumePointValue(scc *securityv1.SecurityContextConstraints) points {
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
	hasHostVolume := false
	hasNonTrivialVolume := false
	for _, v := range scc.Volumes {
		switch v {
		case securityv1.FSTypeHostPath, securityv1.FSTypeAll:
			hasHostVolume = true
			break
		case securityv1.FSTypeSecret, securityv1.FSTypeConfigMap, securityv1.FSTypeEmptyDir, securityv1.FSTypeDownwardAPI, securityv1.FSProjected, securityv1.FSTypeNone:
		default:
			hasNonTrivialVolume = true
		}
	}
	if hasHostVolume {
		return hostVolumePoints
	}
	if hasNonTrivialVolume {
		return nonTrivialVolumePoints
	}
	return noPoints
}
func hasCap(needle string, haystack []corev1.Capability) bool {
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
	for _, c := range haystack {
		if needle == strings.ToUpper(string(c)) {
			return true
		}
	}
	return false
}
func capabilitiesPointValue(scc *securityv1.SecurityContextConstraints) points {
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
	capsPoints := capDefaultPoints
	capsPoints += capAddOnePoints * points(len(scc.DefaultAddCapabilities))
	if hasCap(string(securityv1.AllowAllCapabilities), scc.AllowedCapabilities) {
		capsPoints += capAllowAllPoints
	} else if hasCap("ALL", scc.AllowedCapabilities) {
		capsPoints += capAllowAllPoints
	} else {
		capsPoints += capAllowOnePoints * points(len(scc.AllowedCapabilities))
	}
	if hasCap("ALL", scc.RequiredDropCapabilities) {
		capsPoints += capDropAllPoints
	} else {
		capsPoints += capDropOnePoints * points(len(scc.RequiredDropCapabilities))
	}
	if capsPoints > capMaxPoints {
		return capMaxPoints
	} else if capsPoints < capMinPoints {
		return capMinPoints
	}
	return capsPoints
}
