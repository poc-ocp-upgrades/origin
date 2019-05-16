package ldappassword

import (
	"context"
	"fmt"
	goformat "fmt"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"gopkg.in/ldap.v2"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"runtime/debug"
	gotime "time"
)

type Options struct {
	URL                  ldaputil.LDAPURL
	ClientConfig         ldapclient.Config
	UserAttributeDefiner ldaputil.LDAPUserAttributeDefiner
}
type Authenticator struct {
	providerName    string
	options         Options
	mapper          authapi.UserIdentityMapper
	identityFactory ldaputil.LDAPUserIdentityFactory
}

func New(providerName string, options Options, mapper authapi.UserIdentityMapper) (authenticator.Password, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	auth := &Authenticator{providerName: providerName, options: options, mapper: mapper, identityFactory: &ldaputil.DefaultLDAPUserIdentityFactory{ProviderName: providerName, Definer: options.UserAttributeDefiner}}
	return auth, nil
}
func (a *Authenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	identity, ok, err := a.getIdentity(username, password)
	if err != nil {
		return nil, false, err
	}
	if !ok {
		return nil, false, nil
	}
	return identitymapper.ResponseFor(a.mapper, identity)
}
func (a *Authenticator) getIdentity(username, password string) (authapi.UserIdentityInfo, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer func() {
		if e := recover(); e != nil {
			utilruntime.HandleError(fmt.Errorf("Recovered panic: %v, %s", e, debug.Stack()))
		}
	}()
	if len(username) == 0 || len(password) == 0 {
		return nil, false, nil
	}
	l, err := a.options.ClientConfig.Connect()
	if err != nil {
		return nil, false, err
	}
	defer l.Close()
	if bindDN, bindPassword := a.options.ClientConfig.GetBindCredentials(); len(bindDN) > 0 {
		if err := l.Bind(bindDN, bindPassword); err != nil {
			utilruntime.HandleError(fmt.Errorf("error binding to %s for search phase: %v", bindDN, err))
			return nil, false, err
		}
	}
	filter := fmt.Sprintf("(&%s(%s=%s))", a.options.URL.Filter, ldap.EscapeFilter(a.options.URL.QueryAttribute), ldap.EscapeFilter(username))
	attrs := sets.NewString(a.options.URL.QueryAttribute)
	attrs.Insert(a.options.UserAttributeDefiner.AllAttributes().List()...)
	searchRequest := ldap.NewSearchRequest(a.options.URL.BaseDN, int(a.options.URL.Scope), ldap.NeverDerefAliases, 2, 0, false, filter, attrs.List(), nil)
	klog.V(4).Infof("searching for %s", filter)
	results, err := l.Search(searchRequest)
	if err != nil {
		return nil, false, err
	}
	if len(results.Entries) == 0 {
		klog.V(4).Infof("no entries matching %s", filter)
		return nil, false, nil
	}
	if len(results.Entries) > 1 {
		return nil, false, fmt.Errorf("multiple entries found matching %q", username)
	}
	entry := results.Entries[0]
	klog.V(4).Infof("found dn=%q for %s", entry.DN, filter)
	if err := l.Bind(entry.DN, password); err != nil {
		klog.V(4).Infof("error binding password for %q: %v", entry.DN, err)
		if err, ok := err.(*ldap.Error); ok {
			switch err.ResultCode {
			case ldap.LDAPResultInappropriateAuthentication:
				fallthrough
			case ldap.LDAPResultInvalidCredentials:
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	identity, err := a.identityFactory.IdentityFor(entry)
	if err != nil {
		return nil, false, err
	}
	return identity, true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
