package syncgroups

import (
	"fmt"
	"gopkg.in/ldap.v2"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewUserNameMapper(nameAttributes []string) interfaces.LDAPUserNameMapper {
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
	return &DefaultLDAPUserNameMapper{nameAttributes: nameAttributes}
}

type DefaultLDAPUserNameMapper struct{ nameAttributes []string }

func (m *DefaultLDAPUserNameMapper) UserNameFor(ldapUser *ldap.Entry) (string, error) {
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
	openShiftUserName := ldaputil.GetAttributeValue(ldapUser, m.nameAttributes)
	if len(openShiftUserName) == 0 {
		return "", fmt.Errorf("the user entry (%v) does not map to a OpenShift User name with the given mapping", ldapUser)
	}
	return openShiftUserName, nil
}
