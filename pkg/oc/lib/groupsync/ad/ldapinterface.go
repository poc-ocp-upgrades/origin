package ad

import (
	"gopkg.in/ldap.v2"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/groupdetector"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewADLDAPInterface(clientConfig ldapclient.Config, userQuery ldaputil.LDAPQuery, groupMembershipAttributes []string, userNameAttributes []string) *ADLDAPInterface {
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
	return &ADLDAPInterface{clientConfig: clientConfig, userQuery: userQuery, userNameAttributes: userNameAttributes, groupMembershipAttributes: groupMembershipAttributes, ldapGroupToLDAPMembers: map[string][]*ldap.Entry{}}
}

type ADLDAPInterface struct {
	clientConfig			ldapclient.Config
	userQuery			ldaputil.LDAPQuery
	groupMembershipAttributes	[]string
	userNameAttributes		[]string
	cacheFullyPopulated		bool
	ldapGroupToLDAPMembers		map[string][]*ldap.Entry
}

var _ interfaces.LDAPMemberExtractor = &ADLDAPInterface{}
var _ interfaces.LDAPGroupLister = &ADLDAPInterface{}

func (e *ADLDAPInterface) ExtractMembers(ldapGroupUID string) ([]*ldap.Entry, error) {
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
	if members, present := e.ldapGroupToLDAPMembers[ldapGroupUID]; present {
		return members, nil
	}
	usersInGroup := []*ldap.Entry{}
	for _, currAttribute := range e.groupMembershipAttributes {
		currQuery := ldaputil.LDAPQueryOnAttribute{LDAPQuery: e.userQuery, QueryAttribute: currAttribute}
		searchRequest, err := currQuery.NewSearchRequest(ldapGroupUID, e.requiredUserAttributes())
		if err != nil {
			return nil, err
		}
		currEntries, err := ldaputil.QueryForEntries(e.clientConfig, searchRequest)
		if err != nil {
			return nil, err
		}
		for _, currEntry := range currEntries {
			if !isEntryPresent(usersInGroup, currEntry) {
				usersInGroup = append(usersInGroup, currEntry)
			}
		}
	}
	e.ldapGroupToLDAPMembers[ldapGroupUID] = usersInGroup
	return usersInGroup, nil
}
func (e *ADLDAPInterface) ListGroups() ([]string, error) {
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
	if err := e.populateCache(); err != nil {
		return nil, err
	}
	return sets.StringKeySet(e.ldapGroupToLDAPMembers).List(), nil
}
func (e *ADLDAPInterface) populateCache() error {
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
	if e.cacheFullyPopulated {
		return nil
	}
	searchRequest := e.userQuery.NewSearchRequest(e.requiredUserAttributes())
	userEntries, err := ldaputil.QueryForEntries(e.clientConfig, searchRequest)
	if err != nil {
		return err
	}
	for _, userEntry := range userEntries {
		if userEntry == nil {
			continue
		}
		for _, groupAttribute := range e.groupMembershipAttributes {
			for _, groupUID := range userEntry.GetAttributeValues(groupAttribute) {
				if _, exists := e.ldapGroupToLDAPMembers[groupUID]; !exists {
					e.ldapGroupToLDAPMembers[groupUID] = []*ldap.Entry{}
				}
				if !isEntryPresent(e.ldapGroupToLDAPMembers[groupUID], userEntry) {
					e.ldapGroupToLDAPMembers[groupUID] = append(e.ldapGroupToLDAPMembers[groupUID], userEntry)
				}
			}
		}
	}
	e.cacheFullyPopulated = true
	return nil
}
func isEntryPresent(haystack []*ldap.Entry, needle *ldap.Entry) bool {
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
	for _, curr := range haystack {
		if curr.DN == needle.DN {
			return true
		}
	}
	return false
}
func (e *ADLDAPInterface) requiredUserAttributes() []string {
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
	allAttributes := sets.NewString(e.userNameAttributes...)
	allAttributes.Insert(e.groupMembershipAttributes...)
	return allAttributes.List()
}
func (e *ADLDAPInterface) Exists(ldapGrouUID string) (bool, error) {
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
	return groupdetector.NewMemberBasedDetector(e).Exists(ldapGrouUID)
}
