package identitymapper

import (
	"context"
	"fmt"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

var _ = UserForNewIdentityGetter(&StrategyClaim{})

type StrategyClaim struct {
	user		userclient.UserInterface
	initializer	Initializer
}
type claimError struct {
	User		*userapi.User
	Identity	*userapi.Identity
}

func IsClaimError(err error) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := err.(claimError)
	return ok
}
func NewClaimError(user *userapi.User, identity *userapi.Identity) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return claimError{User: user, Identity: identity}
}
func (c claimError) Error() string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("user %q cannot be claimed by identity %q because it is already mapped to %v", c.User.Name, c.Identity.Name, c.User.Identities)
}
func NewStrategyClaim(user userclient.UserInterface, initializer Initializer) UserForNewIdentityGetter {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &StrategyClaim{user, initializer}
}
func (s *StrategyClaim) UserForNewIdentity(ctx context.Context, preferredUserName string, identity *userapi.Identity) (*userapi.User, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	persistedUser, err := s.user.Get(preferredUserName, metav1.GetOptions{})
	switch {
	case kerrs.IsNotFound(err):
		desiredUser := &userapi.User{}
		desiredUser.Name = preferredUserName
		desiredUser.Identities = []string{identity.Name}
		s.initializer.InitializeUser(identity, desiredUser)
		return s.user.Create(desiredUser)
	case err == nil:
		if sets.NewString(persistedUser.Identities...).Has(identity.Name) {
			return persistedUser, nil
		}
		if len(persistedUser.Identities) == 0 {
			persistedUser.Identities = []string{identity.Name}
			s.initializer.InitializeUser(identity, persistedUser)
			return s.user.Update(persistedUser)
		}
		return nil, NewClaimError(persistedUser, identity)
	default:
		return nil, err
	}
}
