package identitymapper

import (
	"context"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

var _ = UserForNewIdentityGetter(&StrategyAdd{})

type StrategyAdd struct {
	user        userclient.UserInterface
	initializer Initializer
}

func NewStrategyAdd(user userclient.UserInterface, initializer Initializer) UserForNewIdentityGetter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &StrategyAdd{user, initializer}
}
func (s *StrategyAdd) UserForNewIdentity(ctx context.Context, preferredUserName string, identity *userapi.Identity) (*userapi.User, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
		persistedUser.Identities = append(persistedUser.Identities, identity.Name)
		if len(persistedUser.Identities) == 1 {
			s.initializer.InitializeUser(identity, persistedUser)
		}
		return s.user.Update(persistedUser)
	default:
		return nil, err
	}
}
