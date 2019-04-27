package osinserver

import (
	"strings"
	"github.com/RangelReale/osin"
	"github.com/openshift/origin/pkg/oauthserver/server/crypto"
)

var (
	_	osin.AuthorizeTokenGen	= TokenGen{}
	_	osin.AccessTokenGen	= TokenGen{}
)

func randomToken() string {
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
	return randomToken(), nil
}
func (TokenGen) GenerateAccessToken(data *osin.AccessData, generaterefresh bool) (string, string, error) {
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
	accesstoken := randomToken()
	refreshtoken := ""
	if generaterefresh {
		refreshtoken = randomToken()
	}
	return accesstoken, refreshtoken, nil
}
