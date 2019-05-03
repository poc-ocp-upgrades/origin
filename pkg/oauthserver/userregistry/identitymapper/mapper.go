package identitymapper

import (
	"fmt"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
)

type MappingMethodType string

const (
	MappingMethodLookup   MappingMethodType = "lookup"
	MappingMethodClaim    MappingMethodType = "claim"
	MappingMethodAdd      MappingMethodType = "add"
	MappingMethodGenerate MappingMethodType = "generate"
)

func NewIdentityUserMapper(identities userclient.IdentityInterface, users userclient.UserInterface, userIdentityMapping userclient.UserIdentityMappingInterface, method MappingMethodType) (authapi.UserIdentityMapper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	initUser := NewDefaultUserInitStrategy()
	switch method {
	case MappingMethodLookup:
		return &lookupIdentityMapper{userIdentityMapping, users}, nil
	case MappingMethodClaim:
		return &provisioningIdentityMapper{identities, users, NewStrategyClaim(users, initUser)}, nil
	case MappingMethodAdd:
		return &provisioningIdentityMapper{identities, users, NewStrategyAdd(users, initUser)}, nil
	case MappingMethodGenerate:
		return &provisioningIdentityMapper{identities, users, NewStrategyGenerate(users, initUser)}, nil
	default:
		return nil, fmt.Errorf("unsupported mapping method %q", method)
	}
}
