package testclient

import (
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"gopkg.in/ldap.v2"
)

type fakeConfig struct{ client ldap.Client }

func NewConfig(client ldap.Client) ldapclient.Config {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &fakeConfig{client: client}
}
func (c *fakeConfig) Connect() (ldap.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.client, nil
}
func (c *fakeConfig) GetBindCredentials() (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "", ""
}
func (c *fakeConfig) Host() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
