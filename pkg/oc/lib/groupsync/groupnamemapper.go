package syncgroups

import (
	"errors"
	"fmt"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewUserDefinedGroupNameMapper(mapping map[string]string) interfaces.LDAPGroupNameMapper {
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
	return &UserDefinedLDAPGroupNameMapper{nameMapping: mapping}
}

type UserDefinedLDAPGroupNameMapper struct{ nameMapping map[string]string }

func (m *UserDefinedLDAPGroupNameMapper) GroupNameFor(ldapGroupUID string) (string, error) {
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
	openShiftGroupName, exists := m.nameMapping[ldapGroupUID]
	if !exists {
		return "", fmt.Errorf("no OpenShift Group name defined for LDAP group UID: %s", ldapGroupUID)
	}
	return openShiftGroupName, nil
}
func NewEntryAttributeGroupNameMapper(nameAttribute []string, groupGetter interfaces.LDAPGroupGetter) interfaces.LDAPGroupNameMapper {
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
	return &EntryAttributeLDAPGroupNameMapper{nameAttribute: nameAttribute, groupGetter: groupGetter}
}

type EntryAttributeLDAPGroupNameMapper struct {
	nameAttribute	[]string
	groupGetter	interfaces.LDAPGroupGetter
}

func (m *EntryAttributeLDAPGroupNameMapper) GroupNameFor(ldapGroupUID string) (string, error) {
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
	group, err := m.groupGetter.GroupEntryFor(ldapGroupUID)
	if err != nil {
		return "", err
	}
	openShiftGroupName := ldaputil.GetAttributeValue(group, m.nameAttribute)
	if len(openShiftGroupName) == 0 {
		return "", fmt.Errorf("the group entry (%v: %v) does not map to an OpenShift Group name with the given name attribute (%v)", group, group.Attributes, m.nameAttribute)
	}
	return openShiftGroupName, nil
}

type DNLDAPGroupNameMapper struct{}

func (m *DNLDAPGroupNameMapper) GroupNameFor(ldapGroupUID string) (string, error) {
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
	return ldapGroupUID, nil
}

type UnionGroupNameMapper struct {
	GroupNameMappers []interfaces.LDAPGroupNameMapper
}

func (m *UnionGroupNameMapper) GroupNameFor(ldapGroupUID string) (string, error) {
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
	if len(m.GroupNameMappers) == 0 {
		return "", errors.New("no group name mappers defined")
	}
	errs := []error{}
	for _, currMapper := range m.GroupNameMappers {
		ret, err := currMapper.GroupNameFor(ldapGroupUID)
		if err == nil {
			return ret, nil
		}
		errs = append(errs, err)
	}
	return "", kutilerrors.NewAggregate(errs)
}
