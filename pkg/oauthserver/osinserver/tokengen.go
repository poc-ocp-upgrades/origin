package osinserver

import (
	"github.com/RangelReale/osin"
	"github.com/openshift/origin/pkg/oauthserver/server/crypto"
	"strings"
)

var (
	_ osin.AuthorizeTokenGen = TokenGen{}
	_ osin.AccessTokenGen    = TokenGen{}
)

func randomToken() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for {
		token := crypto.Random256BitsString()
		if strings.HasPrefix(token, "-") {
			continue
		}
		return token
	}
}

type TokenGen struct{}

func (TokenGen) GenerateAuthorizeToken(data *osin.AuthorizeData) (ret string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return randomToken(), nil
}
func (TokenGen) GenerateAccessToken(data *osin.AccessData, generaterefresh bool) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accesstoken := randomToken()
	refreshtoken := ""
	if generaterefresh {
		refreshtoken = randomToken()
	}
	return accesstoken, refreshtoken, nil
}
