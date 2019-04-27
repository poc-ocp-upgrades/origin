package sync

import (
	"github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"github.com/openshift/origin/pkg/oc/lib/groupsync"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/ad"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

var _ SyncBuilder = &AugmentedADBuilder{}
var _ PruneBuilder = &AugmentedADBuilder{}

type AugmentedADBuilder struct {
	ClientConfig			ldapclient.Config
	Config				*config.AugmentedActiveDirectoryConfig
	augmentedADLDAPInterface	*ad.AugmentedADLDAPInterface
}

func (b *AugmentedADBuilder) GetGroupLister() (interfaces.LDAPGroupLister, error) {
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
	return b.getAugmentedADLDAPInterface()
}
func (b *AugmentedADBuilder) GetGroupNameMapper() (interfaces.LDAPGroupNameMapper, error) {
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
	ldapInterface, err := b.getAugmentedADLDAPInterface()
	if err != nil {
		return nil, err
	}
	if b.Config.GroupNameAttributes != nil {
		return syncgroups.NewEntryAttributeGroupNameMapper(b.Config.GroupNameAttributes, ldapInterface), nil
	}
	return nil, nil
}
func (b *AugmentedADBuilder) GetUserNameMapper() (interfaces.LDAPUserNameMapper, error) {
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
	return syncgroups.NewUserNameMapper(b.Config.UserNameAttributes), nil
}
func (b *AugmentedADBuilder) GetGroupMemberExtractor() (interfaces.LDAPMemberExtractor, error) {
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
	return b.getAugmentedADLDAPInterface()
}
func (b *AugmentedADBuilder) getAugmentedADLDAPInterface() (*ad.AugmentedADLDAPInterface, error) {
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
	if b.augmentedADLDAPInterface != nil {
		return b.augmentedADLDAPInterface, nil
	}
	userQuery, err := ldaputil.NewLDAPQuery(b.Config.AllUsersQuery)
	if err != nil {
		return nil, err
	}
	groupQuery, err := ldaputil.NewLDAPQueryOnAttribute(b.Config.AllGroupsQuery, b.Config.GroupUIDAttribute)
	if err != nil {
		return nil, err
	}
	b.augmentedADLDAPInterface = ad.NewAugmentedADLDAPInterface(b.ClientConfig, userQuery, b.Config.GroupMembershipAttributes, b.Config.UserNameAttributes, groupQuery, b.Config.GroupNameAttributes)
	return b.augmentedADLDAPInterface, nil
}
func (b *AugmentedADBuilder) GetGroupDetector() (interfaces.LDAPGroupDetector, error) {
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
	return b.getAugmentedADLDAPInterface()
}
