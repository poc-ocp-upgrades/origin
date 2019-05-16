package ldaputil

import (
	"fmt"
	"github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"gopkg.in/ldap.v2"
	"k8s.io/klog"
	"strings"
)

func NewLDAPQuery(config config.LDAPQuery) (LDAPQuery, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scope, err := DetermineLDAPScope(config.Scope)
	if err != nil {
		return LDAPQuery{}, err
	}
	derefAliases, err := DetermineDerefAliasesBehavior(config.DerefAliases)
	if err != nil {
		return LDAPQuery{}, err
	}
	return LDAPQuery{BaseDN: config.BaseDN, Scope: scope, DerefAliases: derefAliases, TimeLimit: config.TimeLimit, Filter: config.Filter, PageSize: config.PageSize}, nil
}

type LDAPQuery struct {
	BaseDN       string
	Scope        Scope
	DerefAliases DerefAliases
	TimeLimit    int
	Filter       string
	PageSize     int
}

func (q *LDAPQuery) NewSearchRequest(additionalAttributes []string) *ldap.SearchRequest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var controls []ldap.Control
	if q.PageSize > 0 {
		controls = append(controls, ldap.NewControlPaging(uint32(q.PageSize)))
	}
	return ldap.NewSearchRequest(q.BaseDN, int(q.Scope), int(q.DerefAliases), 0, q.TimeLimit, false, q.Filter, additionalAttributes, controls)
}
func NewLDAPQueryOnAttribute(config config.LDAPQuery, attribute string) (LDAPQueryOnAttribute, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ldapQuery, err := NewLDAPQuery(config)
	if err != nil {
		return LDAPQueryOnAttribute{}, err
	}
	return LDAPQueryOnAttribute{LDAPQuery: ldapQuery, QueryAttribute: attribute}, nil
}

type LDAPQueryOnAttribute struct {
	LDAPQuery
	QueryAttribute string
}

func (o *LDAPQueryOnAttribute) NewSearchRequest(attributeValue string, attributes []string) (*ldap.SearchRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if strings.EqualFold(o.QueryAttribute, "dn") {
		dn, err := ldap.ParseDN(attributeValue)
		if err != nil {
			return nil, fmt.Errorf("could not search by dn, invalid dn value: %v", err)
		}
		baseDN, err := ldap.ParseDN(o.BaseDN)
		if err != nil {
			return nil, fmt.Errorf("could not search by dn, invalid dn value: %v", err)
		}
		if !baseDN.AncestorOf(dn) && !baseDN.Equal(dn) {
			return nil, NewQueryOutOfBoundsError(attributeValue, o.BaseDN)
		}
		return o.buildDNQuery(attributeValue, attributes), nil
	} else {
		return o.buildAttributeQuery(attributeValue, attributes), nil
	}
}
func (o *LDAPQueryOnAttribute) buildDNQuery(dn string, attributes []string) *ldap.SearchRequest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var controls []ldap.Control
	if o.PageSize > 0 {
		controls = append(controls, ldap.NewControlPaging(uint32(o.PageSize)))
	}
	return ldap.NewSearchRequest(dn, ldap.ScopeBaseObject, int(o.DerefAliases), 0, o.TimeLimit, false, "(objectClass=*)", attributes, controls)
}
func (o *LDAPQueryOnAttribute) buildAttributeQuery(attributeValue string, attributes []string) *ldap.SearchRequest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	specificFilter := fmt.Sprintf("%s=%s", ldap.EscapeFilter(o.QueryAttribute), ldap.EscapeFilter(attributeValue))
	filter := fmt.Sprintf("(&(%s)(%s))", o.Filter, specificFilter)
	var controls []ldap.Control
	if o.PageSize > 0 {
		controls = append(controls, ldap.NewControlPaging(uint32(o.PageSize)))
	}
	return ldap.NewSearchRequest(o.BaseDN, int(o.Scope), int(o.DerefAliases), 0, o.TimeLimit, false, filter, attributes, controls)
}
func QueryForUniqueEntry(clientConfig ldapclient.Config, query *ldap.SearchRequest) (*ldap.Entry, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := QueryForEntries(clientConfig, query)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, NewEntryNotFoundError(query.BaseDN, query.Filter)
	}
	if len(result) > 1 {
		if query.Scope == ldap.ScopeBaseObject {
			return nil, fmt.Errorf("multiple entries found matching dn=%q:\n%s", query.BaseDN, formatResult(result))
		} else {
			return nil, fmt.Errorf("multiple entries found matching filter %s:\n%s", query.Filter, formatResult(result))
		}
	}
	entry := result[0]
	klog.V(4).Infof("found dn=%q for %s", entry.DN, query.Filter)
	return entry, nil
}
func formatResult(results []*ldap.Entry) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var names []string
	for _, entry := range results {
		names = append(names, entry.DN)
	}
	return "\t" + strings.Join(names[0:10], "\n\t")
}
func QueryForEntries(clientConfig ldapclient.Config, query *ldap.SearchRequest) ([]*ldap.Entry, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	connection, err := clientConfig.Connect()
	if err != nil {
		return nil, fmt.Errorf("could not connect to the LDAP server: %v", err)
	}
	defer connection.Close()
	if bindDN, bindPassword := clientConfig.GetBindCredentials(); len(bindDN) > 0 {
		if err := connection.Bind(bindDN, bindPassword); err != nil {
			return nil, fmt.Errorf("could not bind to the LDAP server: %v", err)
		}
	}
	var searchResult *ldap.SearchResult
	control := ldap.FindControl(query.Controls, ldap.ControlTypePaging)
	if control == nil {
		klog.V(4).Infof("searching LDAP server with config %v with dn=%q and scope %v for %s requesting %v", clientConfig, query.BaseDN, query.Scope, query.Filter, query.Attributes)
		searchResult, err = connection.Search(query)
	} else if pagingControl, ok := control.(*ldap.ControlPaging); ok {
		klog.V(4).Infof("searching LDAP server with config %v with dn=%q and scope %v for %s requesting %v with pageSize=%d", clientConfig, query.BaseDN, query.Scope, query.Filter, query.Attributes, pagingControl.PagingSize)
		searchResult, err = connection.SearchWithPaging(query, pagingControl.PagingSize)
	} else {
		err = fmt.Errorf("invalid paging control type: %v", control)
	}
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, NewNoSuchObjectError(query.BaseDN)
		}
		return nil, err
	}
	for _, entry := range searchResult.Entries {
		klog.V(4).Infof("found dn=%q ", entry.DN)
	}
	return searchResult.Entries, nil
}
