package google

import (
	"testing"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
)

func TestGoogle(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p, err := NewProvider("google", "clientid", "clientsecret", "", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	_ = external.Provider(p)
}
