package groupdetector

import (
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewGroupBasedDetector(groupGetter interfaces.LDAPGroupGetter) interfaces.LDAPGroupDetector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GroupBasedDetector{groupGetter: groupGetter}
}

type GroupBasedDetector struct{ groupGetter interfaces.LDAPGroupGetter }

func (l *GroupBasedDetector) Exists(ldapGroupUID string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	group, err := l.groupGetter.GroupEntryFor(ldapGroupUID)
	if ldaputil.IsQueryOutOfBoundsError(err) || ldaputil.IsEntryNotFoundError(err) || ldaputil.IsNoSuchObjectError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if group == nil {
		return false, nil
	}
	return true, nil
}
func NewMemberBasedDetector(memberExtractor interfaces.LDAPMemberExtractor) interfaces.LDAPGroupDetector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MemberBasedDetector{memberExtractor: memberExtractor}
}

type MemberBasedDetector struct {
	memberExtractor interfaces.LDAPMemberExtractor
}

func (l *MemberBasedDetector) Exists(ldapGrouUID string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	members, err := l.memberExtractor.ExtractMembers(ldapGrouUID)
	if ldaputil.IsQueryOutOfBoundsError(err) || ldaputil.IsEntryNotFoundError(err) || ldaputil.IsNoSuchObjectError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if len(members) == 0 {
		return false, nil
	}
	return true, nil
}
func NewCompoundDetector(locators ...interfaces.LDAPGroupDetector) interfaces.LDAPGroupDetector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &CompoundDetector{locators: locators}
}

type CompoundDetector struct {
	locators []interfaces.LDAPGroupDetector
}

func (l *CompoundDetector) Exists(ldapGrouUID string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(l.locators) == 0 {
		return false, nil
	}
	conclusion := true
	for _, locator := range l.locators {
		opinion, err := locator.Exists(ldapGrouUID)
		if err != nil {
			return false, err
		}
		conclusion = conclusion && opinion
	}
	return conclusion, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
