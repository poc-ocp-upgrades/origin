package v1

import "github.com/openshift/api/oauth/v1"

func SetDefaults_OAuthAuthorizeToken(obj *v1.OAuthAuthorizeToken) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(obj.CodeChallenge) > 0 && len(obj.CodeChallengeMethod) == 0 {
		obj.CodeChallengeMethod = "plain"
	}
}
