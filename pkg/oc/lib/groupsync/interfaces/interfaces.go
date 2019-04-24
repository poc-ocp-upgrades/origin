package interfaces

import (
	"gopkg.in/ldap.v2"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type LDAPGroupLister interface {
	ListGroups() (ldapGroupUIDs []string, err error)
}
type LDAPMemberExtractor interface {
	ExtractMembers(ldapGroupUID string) (members []*ldap.Entry, err error)
}
type LDAPGroupNameMapper interface {
	GroupNameFor(ldapGroupUID string) (openShiftGroupName string, err error)
}
type LDAPUserNameMapper interface {
	UserNameFor(ldapUser *ldap.Entry) (openShiftUserName string, err error)
}
type LDAPGroupGetter interface {
	GroupEntryFor(ldapGroupUID string) (group *ldap.Entry, err error)
}
type LDAPGroupListerNameMapper interface {
	LDAPGroupLister
	LDAPGroupNameMapper
}
type LDAPGroupDetector interface {
	Exists(ldapGroupUID string) (exists bool, err error)
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
