package identitymapper

import (
	"context"
	"errors"
	"fmt"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type UserNameGenerator func(base string, sequence int) string

var (
	MaxGenerateAttempts	= 100
	DefaultGenerator	= UserNameGenerator(func(base string, sequence int) string {
		if sequence == 0 {
			return base
		}
		return fmt.Sprintf("%s%d", base, sequence+1)
	})
)
var _ = UserForNewIdentityGetter(&StrategyGenerate{})

type StrategyGenerate struct {
	user		userclient.UserInterface
	generator	UserNameGenerator
	initializer	Initializer
}

func NewStrategyGenerate(user userclient.UserInterface, initializer Initializer) UserForNewIdentityGetter {
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
	return &StrategyGenerate{user, DefaultGenerator, initializer}
}
func (s *StrategyGenerate) UserForNewIdentity(ctx context.Context, preferredUserName string, identity *userapi.Identity) (*userapi.User, error) {
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
UserSearch:
	for sequence := 0; sequence < MaxGenerateAttempts; sequence++ {
		potentialUserName := s.generator(preferredUserName, sequence)
		persistedUser, err := s.user.Get(potentialUserName, metav1.GetOptions{})
		switch {
		case kerrs.IsNotFound(err):
			desiredUser := &userapi.User{}
			desiredUser.Name = potentialUserName
			desiredUser.Identities = []string{identity.Name}
			s.initializer.InitializeUser(identity, desiredUser)
			return s.user.Create(desiredUser)
		case err == nil:
			if sets.NewString(persistedUser.Identities...).Has(identity.Name) {
				return persistedUser, nil
			}
			continue UserSearch
		default:
			return nil, err
		}
	}
	return nil, errors.New("Could not create user, max attempts exceeded")
}
