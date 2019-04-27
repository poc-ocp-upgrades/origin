package ldaputil

import (
	"fmt"
	"strings"
	"gopkg.in/ldap.v2"
	"k8s.io/apimachinery/pkg/util/sets"
	osinv1 "github.com/openshift/api/osin/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
)

type LDAPUserIdentityFactory interface {
	IdentityFor(user *ldap.Entry) (identity authapi.UserIdentityInfo, err error)
}
type DefaultLDAPUserIdentityFactory struct {
	ProviderName	string
	Definer		LDAPUserAttributeDefiner
}

func (f *DefaultLDAPUserIdentityFactory) IdentityFor(user *ldap.Entry) (identity authapi.UserIdentityInfo, err error) {
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
	uid := f.Definer.ID(user)
	if uid == "" {
		err = fmt.Errorf("Could not retrieve a non-empty value for ID attributes for dn=%q", user.DN)
		return
	}
	id := authapi.NewDefaultUserIdentityInfo(f.ProviderName, uid)
	if name := f.Definer.Name(user); len(name) != 0 {
		id.Extra[authapi.IdentityDisplayNameKey] = name
	}
	if email := f.Definer.Email(user); len(email) != 0 {
		id.Extra[authapi.IdentityEmailKey] = email
	}
	if prefUser := f.Definer.PreferredUsername(user); len(prefUser) != 0 {
		id.Extra[authapi.IdentityPreferredUsernameKey] = prefUser
	}
	identity = id
	return
}
func NewLDAPUserAttributeDefiner(attributeMapping osinv1.LDAPAttributeMapping) LDAPUserAttributeDefiner {
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
	return LDAPUserAttributeDefiner{attributeMapping: attributeMapping}
}

type LDAPUserAttributeDefiner struct{ attributeMapping osinv1.LDAPAttributeMapping }

func (d *LDAPUserAttributeDefiner) AllAttributes() sets.String {
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
	attrs := sets.NewString(d.attributeMapping.Email...)
	attrs.Insert(d.attributeMapping.Name...)
	attrs.Insert(d.attributeMapping.PreferredUsername...)
	attrs.Insert(d.attributeMapping.ID...)
	return attrs
}
func (d *LDAPUserAttributeDefiner) Email(user *ldap.Entry) string {
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
	return GetAttributeValue(user, d.attributeMapping.Email)
}
func (d *LDAPUserAttributeDefiner) Name(user *ldap.Entry) string {
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
	return GetAttributeValue(user, d.attributeMapping.Name)
}
func (d *LDAPUserAttributeDefiner) PreferredUsername(user *ldap.Entry) string {
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
	return GetAttributeValue(user, d.attributeMapping.PreferredUsername)
}
func (d *LDAPUserAttributeDefiner) ID(user *ldap.Entry) string {
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
	return GetAttributeValue(user, d.attributeMapping.ID)
}
func GetAttributeValue(entry *ldap.Entry, attributes []string) string {
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
	for _, k := range attributes {
		if len(k) == 0 {
			continue
		}
		if strings.ToLower(k) == "dn" {
			return entry.DN
		}
		if v := entry.GetAttributeValue(k); len(v) > 0 {
			return v
		}
	}
	return ""
}
