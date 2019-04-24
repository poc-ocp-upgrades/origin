package handlers

import (
	"testing"
	"github.com/openshift/origin/pkg/oauthserver/osinserver"
)

func TestGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = osinserver.AuthorizeHandler(&GrantCheck{})
}
func TestEmptyGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = NewEmptyGrant()
}
func TestAutoGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = NewAutoGrant()
}
func TestRedirectGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = NewRedirectGrant("/")
}
