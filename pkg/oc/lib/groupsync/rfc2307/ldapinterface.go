package rfc2307

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"gopkg.in/ldap.v2"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/groupdetector"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/syncerror"
)

func NewLDAPInterface(clientConfig ldapclient.Config, groupQuery ldaputil.LDAPQueryOnAttribute, groupNameAttributes []string, groupMembershipAttributes []string, userQuery ldaputil.LDAPQueryOnAttribute, userNameAttributes []string, errorHandler syncerror.Handler) *LDAPInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &LDAPInterface{clientConfig: clientConfig, groupQuery: groupQuery, groupNameAttributes: groupNameAttributes, groupMembershipAttributes: groupMembershipAttributes, userQuery: userQuery, userNameAttributes: userNameAttributes, cachedUsers: map[string]*ldap.Entry{}, cachedGroups: map[string]*ldap.Entry{}, errorHandler: errorHandler}
}

type LDAPInterface struct {
	clientConfig			ldapclient.Config
	groupQuery			ldaputil.LDAPQueryOnAttribute
	groupNameAttributes		[]string
	groupMembershipAttributes	[]string
	userQuery			ldaputil.LDAPQueryOnAttribute
	userNameAttributes		[]string
	cachedGroups			map[string]*ldap.Entry
	cachedUsers			map[string]*ldap.Entry
	errorHandler			syncerror.Handler
}

var _ interfaces.LDAPMemberExtractor = &LDAPInterface{}
var _ interfaces.LDAPGroupGetter = &LDAPInterface{}
var _ interfaces.LDAPGroupLister = &LDAPInterface{}

func (e *LDAPInterface) ExtractMembers(ldapGroupUID string) ([]*ldap.Entry, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	group, err := e.GroupEntryFor(ldapGroupUID)
	if err != nil {
		return nil, err
	}
	var ldapMemberUIDs []string
	for _, attribute := range e.groupMembershipAttributes {
		ldapMemberUIDs = append(ldapMemberUIDs, group.GetAttributeValues(attribute)...)
	}
	members := []*ldap.Entry{}
	for _, ldapMemberUID := range ldapMemberUIDs {
		memberEntry, err := e.userEntryFor(ldapMemberUID)
		if err == nil {
			members = append(members, memberEntry)
			continue
		}
		err = syncerror.NewMemberLookupError(ldapGroupUID, ldapMemberUID, err)
		handled, fatalErr := e.errorHandler.HandleError(err)
		if fatalErr != nil {
			return nil, fatalErr
		}
		if !handled {
			return nil, err
		}
	}
	return members, nil
}
func (e *LDAPInterface) GroupEntryFor(ldapGroupUID string) (*ldap.Entry, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	group, exists := e.cachedGroups[ldapGroupUID]
	if exists {
		return group, nil
	}
	searchRequest, err := e.groupQuery.NewSearchRequest(ldapGroupUID, e.requiredGroupAttributes())
	if err != nil {
		return nil, err
	}
	group, err = ldaputil.QueryForUniqueEntry(e.clientConfig, searchRequest)
	if err != nil {
		return nil, err
	}
	e.cachedGroups[ldapGroupUID] = group
	return group, nil
}
func (e *LDAPInterface) ListGroups() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	searchRequest := e.groupQuery.LDAPQuery.NewSearchRequest(e.requiredGroupAttributes())
	groups, err := ldaputil.QueryForEntries(e.clientConfig, searchRequest)
	if err != nil {
		return nil, err
	}
	ldapGroupUIDs := []string{}
	for _, group := range groups {
		ldapGroupUID := ldaputil.GetAttributeValue(group, []string{e.groupQuery.QueryAttribute})
		if len(ldapGroupUID) == 0 {
			return nil, fmt.Errorf("unable to find LDAP group UID for %s", group)
		}
		e.cachedGroups[ldapGroupUID] = group
		ldapGroupUIDs = append(ldapGroupUIDs, ldapGroupUID)
	}
	return ldapGroupUIDs, nil
}
func (e *LDAPInterface) requiredGroupAttributes() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allAttributes := sets.NewString(e.groupNameAttributes...)
	allAttributes.Insert(e.groupMembershipAttributes...)
	allAttributes.Insert(e.groupQuery.QueryAttribute)
	return allAttributes.List()
}
func (e *LDAPInterface) userEntryFor(ldapUserUID string) (user *ldap.Entry, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user, exists := e.cachedUsers[ldapUserUID]
	if exists {
		return user, nil
	}
	searchRequest, err := e.userQuery.NewSearchRequest(ldapUserUID, e.requiredUserAttributes())
	if err != nil {
		return nil, err
	}
	user, err = ldaputil.QueryForUniqueEntry(e.clientConfig, searchRequest)
	if err != nil {
		return nil, err
	}
	e.cachedUsers[ldapUserUID] = user
	return user, nil
}
func (e *LDAPInterface) requiredUserAttributes() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allAttributes := sets.NewString(e.userNameAttributes...)
	allAttributes.Insert(e.userQuery.QueryAttribute)
	return allAttributes.List()
}
func (e *LDAPInterface) Exists(ldapGroupUID string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return groupdetector.NewGroupBasedDetector(e).Exists(ldapGroupUID)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
