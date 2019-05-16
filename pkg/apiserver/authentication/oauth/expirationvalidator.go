package oauth

import (
	"errors"
	oauthv1 "github.com/openshift/api/oauth/v1"
	userv1 "github.com/openshift/api/user/v1"
	"time"
)

var errExpired = errors.New("token is expired")

func NewExpirationValidator() OAuthTokenValidator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return OAuthTokenValidatorFunc(func(token *oauthv1.OAuthAccessToken, _ *userv1.User) error {
		if token.ExpiresIn > 0 {
			if expire(token).Before(time.Now()) {
				return errExpired
			}
		}
		if token.DeletionTimestamp != nil {
			return errExpired
		}
		return nil
	})
}
func expire(token *oauthv1.OAuthAccessToken) time.Time {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return token.CreationTimestamp.Add(time.Duration(token.ExpiresIn) * time.Second)
}
