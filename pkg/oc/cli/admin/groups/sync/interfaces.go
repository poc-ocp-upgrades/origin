package sync

import (
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

type SyncBuilder interface {
	GetGroupLister() (interfaces.LDAPGroupLister, error)
	GetGroupNameMapper() (interfaces.LDAPGroupNameMapper, error)
	GetUserNameMapper() (interfaces.LDAPUserNameMapper, error)
	GetGroupMemberExtractor() (interfaces.LDAPMemberExtractor, error)
}
type PruneBuilder interface {
	GetGroupLister() (interfaces.LDAPGroupLister, error)
	GetGroupNameMapper() (interfaces.LDAPGroupNameMapper, error)
	GetGroupDetector() (interfaces.LDAPGroupDetector, error)
}
type GroupNameRestrictions interface {
	GetWhitelist() []string
	GetBlacklist() []string
}
type OpenShiftGroupNameRestrictions interface {
	GroupNameRestrictions
	GetClient() userv1client.GroupInterface
}
type MappedNameRestrictions interface{ GetGroupNameMappings() map[string]string }
