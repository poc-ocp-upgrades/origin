package identitymapper

import (
	userapi "github.com/openshift/api/user/v1"
)

type Initializer interface {
	InitializeUser(identity *userapi.Identity, user *userapi.User) error
}
