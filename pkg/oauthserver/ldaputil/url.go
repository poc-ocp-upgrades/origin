package ldaputil

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"gopkg.in/ldap.v2"
)

type Scheme string

const (
	SchemeLDAP	Scheme	= "ldap"
	SchemeLDAPS	Scheme	= "ldaps"
)

type Scope int

const (
	ScopeWholeSubtree	Scope	= ldap.ScopeWholeSubtree
	ScopeSingleLevel	Scope	= ldap.ScopeSingleLevel
	ScopeBaseObject		Scope	= ldap.ScopeBaseObject
)

type DerefAliases int

const (
	DerefAliasesNever	= ldap.NeverDerefAliases
	DerefAliasesSearching	= ldap.DerefInSearching
	DerefAliasesFinding	= ldap.DerefFindingBaseObj
	DerefAliasesAlways	= ldap.DerefAlways
)
const (
	defaultLDAPPort		= "389"
	defaultLDAPSPort	= "636"
	defaultHost		= "localhost"
	defaultQueryAttribute	= "uid"
	defaultFilter		= "(objectClass=*)"
	scopeWholeSubtreeString	= "sub"
	scopeSingleLevelString	= "one"
	scopeBaseObjectString	= "base"
	criticalExtensionPrefix	= "!"
)

type LDAPURL struct {
	Scheme		Scheme
	Host		string
	BaseDN		string
	QueryAttribute	string
	Scope		Scope
	Filter		string
}

func ParseURL(ldapURL string) (LDAPURL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parsedURL, err := url.Parse(ldapURL)
	if err != nil {
		return LDAPURL{}, err
	}
	opts := LDAPURL{}
	determinedScheme, err := DetermineLDAPScheme(parsedURL.Scheme)
	if err != nil {
		return LDAPURL{}, err
	}
	opts.Scheme = determinedScheme
	determinedHost, err := DetermineLDAPHost(parsedURL.Host, opts.Scheme)
	if err != nil {
		return LDAPURL{}, err
	}
	opts.Host = determinedHost
	opts.BaseDN = strings.TrimLeft(parsedURL.Path, "/")
	attributes, scope, filter, extensions, err := SplitLDAPQuery(parsedURL.RawQuery)
	if err != nil {
		return LDAPURL{}, err
	}
	opts.QueryAttribute = strings.Split(attributes, ",")[0]
	if len(opts.QueryAttribute) == 0 {
		opts.QueryAttribute = defaultQueryAttribute
	}
	determinedScope, err := DetermineLDAPScope(scope)
	if err != nil {
		return LDAPURL{}, err
	}
	opts.Scope = determinedScope
	determinedFilter, err := DetermineLDAPFilter(filter)
	if err != nil {
		return LDAPURL{}, err
	}
	opts.Filter = determinedFilter
	if len(extensions) > 0 {
		for _, extension := range strings.Split(extensions, ",") {
			exttype := strings.SplitN(extension, "=", 2)[0]
			if strings.HasPrefix(exttype, criticalExtensionPrefix) {
				return LDAPURL{}, fmt.Errorf("unsupported critical extension %s", extension)
			}
		}
	}
	return opts, nil
}
func DetermineLDAPScheme(scheme string) (Scheme, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch Scheme(scheme) {
	case SchemeLDAP, SchemeLDAPS:
		return Scheme(scheme), nil
	default:
		return "", fmt.Errorf("invalid scheme %q", scheme)
	}
}
func DetermineLDAPHost(hostport string, scheme Scheme) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(hostport) == 0 {
		hostport = defaultHost
	}
	if _, _, err := net.SplitHostPort(hostport); err != nil {
		switch scheme {
		case SchemeLDAPS:
			return net.JoinHostPort(hostport, defaultLDAPSPort), nil
		case SchemeLDAP:
			return net.JoinHostPort(hostport, defaultLDAPPort), nil
		default:
			return "", fmt.Errorf("no default port for scheme %q", scheme)
		}
	}
	return hostport, nil
}
func SplitLDAPQuery(query string) (attributes, scope, filter, extensions string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.Split(query, "?")
	switch len(parts) {
	case 4:
		extensions = parts[3]
		fallthrough
	case 3:
		if v, err := url.QueryUnescape(parts[2]); err != nil {
			return "", "", "", "", err
		} else {
			filter = v
		}
		fallthrough
	case 2:
		if v, err := url.QueryUnescape(parts[1]); err != nil {
			return "", "", "", "", err
		} else {
			scope = v
		}
		fallthrough
	case 1:
		if v, err := url.QueryUnescape(parts[0]); err != nil {
			return "", "", "", "", err
		} else {
			attributes = v
		}
		return attributes, scope, filter, extensions, nil
	case 0:
		return
	default:
		err = fmt.Errorf("too many query options %q", query)
		return "", "", "", "", err
	}
}
func DetermineLDAPScope(scope string) (Scope, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch scope {
	case "", scopeWholeSubtreeString:
		return ScopeWholeSubtree, nil
	case scopeSingleLevelString:
		return ScopeSingleLevel, nil
	case scopeBaseObjectString:
		return ScopeBaseObject, nil
	default:
		return -1, fmt.Errorf("invalid scope %q", scope)
	}
}
func DetermineLDAPFilter(filter string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(filter) == 0 {
		return defaultFilter, nil
	}
	if _, err := ldap.CompileFilter(filter); err != nil {
		return "", fmt.Errorf("invalid filter: %v", err)
	}
	return filter, nil
}
func DetermineDerefAliasesBehavior(derefAliasesString string) (DerefAliases, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mapping := map[string]DerefAliases{"never": DerefAliasesNever, "search": DerefAliasesSearching, "base": DerefAliasesFinding, "always": DerefAliasesAlways}
	derefAliases, exists := mapping[derefAliasesString]
	if !exists {
		return -1, fmt.Errorf("not a valid LDAP alias dereferncing behavior: %s", derefAliasesString)
	}
	return derefAliases, nil
}
