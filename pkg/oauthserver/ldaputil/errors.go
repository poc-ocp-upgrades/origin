package ldaputil

import (
	"fmt"
	"gopkg.in/ldap.v2"
)

func NewNoSuchObjectError(baseDN string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &errNoSuchObject{baseDN: baseDN}
}

type errNoSuchObject struct{ baseDN string }

func (e *errNoSuchObject) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("search for entry with base dn=%q refers to a non-existent entry", e.baseDN)
}
func IsNoSuchObjectError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return false
	}
	_, ok := err.(*errNoSuchObject)
	return ok || ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject)
}
func NewEntryNotFoundError(baseDN, filter string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &errEntryNotFound{baseDN: baseDN, filter: filter}
}

type errEntryNotFound struct {
	baseDN string
	filter string
}

func (e *errEntryNotFound) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("search for entry with base dn=%q and filter %q did not return any results", e.baseDN, e.filter)
}
func IsEntryNotFoundError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return false
	}
	_, ok := err.(*errEntryNotFound)
	return ok
}
func NewQueryOutOfBoundsError(queryDN, baseDN string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &errQueryOutOfBounds{baseDN: baseDN, queryDN: queryDN}
}

type errQueryOutOfBounds struct {
	baseDN  string
	queryDN string
}

func (q *errQueryOutOfBounds) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("search for entry with dn=%q would search outside of the base dn specified (dn=%q)", q.queryDN, q.baseDN)
}
func IsQueryOutOfBoundsError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return false
	}
	_, ok := err.(*errQueryOutOfBounds)
	return ok
}
