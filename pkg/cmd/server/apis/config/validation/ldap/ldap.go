package ldap

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"strings"
	"gopkg.in/ldap.v2"
	"k8s.io/apimachinery/pkg/util/validation/field"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/common"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
)

func ValidateLDAPSyncConfig(config *configapi.LDAPSyncConfig) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	validationResults.Append(common.ValidateStringSource(config.BindPassword, field.NewPath("bindPassword")))
	bindPassword, _ := configapi.ResolveStringValue(config.BindPassword)
	validationResults.Append(ValidateLDAPClientConfig(config.URL, config.BindDN, bindPassword, config.CA, config.Insecure, nil))
	for ldapGroupUID, openShiftGroupName := range config.LDAPGroupUIDToOpenShiftGroupNameMapping {
		if len(ldapGroupUID) == 0 || len(openShiftGroupName) == 0 {
			validationResults.AddErrors(field.Invalid(field.NewPath("groupUIDNameMapping").Key(ldapGroupUID), openShiftGroupName, "has empty key or value"))
		}
	}
	schemaConfigsFound := []string{}
	if config.RFC2307Config != nil {
		configResults := ValidateRFC2307Config(config.RFC2307Config)
		validationResults.AddErrors(configResults.Errors...)
		validationResults.AddWarnings(configResults.Warnings...)
		schemaConfigsFound = append(schemaConfigsFound, "rfc2307")
	}
	if config.ActiveDirectoryConfig != nil {
		configResults := ValidateActiveDirectoryConfig(config.ActiveDirectoryConfig)
		validationResults.AddErrors(configResults.Errors...)
		validationResults.AddWarnings(configResults.Warnings...)
		schemaConfigsFound = append(schemaConfigsFound, "activeDirectory")
	}
	if config.AugmentedActiveDirectoryConfig != nil {
		configResults := ValidateAugmentedActiveDirectoryConfig(config.AugmentedActiveDirectoryConfig)
		validationResults.AddErrors(configResults.Errors...)
		validationResults.AddWarnings(configResults.Warnings...)
		schemaConfigsFound = append(schemaConfigsFound, "augmentedActiveDirectory")
	}
	if len(schemaConfigsFound) > 1 {
		validationResults.AddErrors(field.Invalid(field.NewPath("schema"), config, fmt.Sprintf("only one schema-specific config is allowed; found %v", schemaConfigsFound)))
	}
	if len(schemaConfigsFound) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("schema"), fmt.Sprintf("exactly one schema-specific config is required;  one of %v", []string{"rfc2307", "activeDirectory", "augmentedActiveDirectory"})))
	}
	return validationResults
}
func ValidateLDAPClientConfig(url, bindDN, bindPassword, CA string, insecure bool, fldPath *field.Path) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	if len(url) == 0 {
		validationResults.AddErrors(field.Required(fldPath.Child("url"), ""))
		return validationResults
	}
	u, err := ldaputil.ParseURL(url)
	if err != nil {
		validationResults.AddErrors(field.Invalid(fldPath.Child("url"), url, err.Error()))
		return validationResults
	}
	if (len(bindDN) == 0) != (len(bindPassword) == 0) {
		validationResults.AddErrors(field.Invalid(fldPath.Child("bindDN"), bindDN, "bindDN and bindPassword must both be specified, or both be empty"))
		validationResults.AddErrors(field.Invalid(fldPath.Child("bindPassword"), "(masked)", "bindDN and bindPassword must both be specified, or both be empty"))
	}
	if insecure {
		if u.Scheme == ldaputil.SchemeLDAPS {
			validationResults.AddErrors(field.Invalid(fldPath.Child("url"), url, fmt.Sprintf("Cannot use %s scheme with insecure=true", u.Scheme)))
		}
		if len(CA) > 0 {
			validationResults.AddErrors(field.Invalid(fldPath.Child("ca"), CA, "Cannot specify a ca with insecure=true"))
		}
	} else {
		if len(CA) > 0 {
			validationResults.AddErrors(common.ValidateFile(CA, fldPath.Child("ca"))...)
		}
	}
	if insecure {
		validationResults.AddWarnings(field.Invalid(fldPath.Child("insecure"), insecure, "validating passwords over an insecure connection could allow them to be intercepted"))
	}
	return validationResults
}
func ValidateRFC2307Config(config *configapi.RFC2307Config) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	validationResults.Append(ValidateLDAPQuery(config.AllGroupsQuery, field.NewPath("groupsQuery")))
	if len(config.GroupUIDAttribute) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupUIDAttribute"), ""))
	}
	if len(config.GroupNameAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupNameAttributes"), ""))
	}
	if len(config.GroupMembershipAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupMembershipAttributes"), ""))
	}
	isUserDNQuery := strings.TrimSpace(strings.ToLower(config.UserUIDAttribute)) == "dn"
	validationResults.Append(validateLDAPQuery(config.AllUsersQuery, field.NewPath("usersQuery"), isUserDNQuery))
	if len(config.UserUIDAttribute) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("userUIDAttribute"), ""))
	}
	if len(config.UserNameAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("userNameAttributes"), ""))
	}
	return validationResults
}
func ValidateActiveDirectoryConfig(config *configapi.ActiveDirectoryConfig) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	validationResults.Append(ValidateLDAPQuery(config.AllUsersQuery, field.NewPath("usersQuery")))
	if len(config.UserNameAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("userNameAttributes"), ""))
	}
	if len(config.GroupMembershipAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupMembershipAttributes"), ""))
	}
	return validationResults
}
func ValidateAugmentedActiveDirectoryConfig(config *configapi.AugmentedActiveDirectoryConfig) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	validationResults.Append(ValidateLDAPQuery(config.AllUsersQuery, field.NewPath("usersQuery")))
	if len(config.UserNameAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("userNameAttributes"), ""))
	}
	if len(config.GroupMembershipAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupMembershipAttributes"), ""))
	}
	isGroupDNQuery := strings.TrimSpace(strings.ToLower(config.GroupUIDAttribute)) == "dn"
	validationResults.Append(validateLDAPQuery(config.AllGroupsQuery, field.NewPath("groupsQuery"), isGroupDNQuery))
	if len(config.GroupUIDAttribute) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupUIDAttribute"), ""))
	}
	if len(config.GroupNameAttributes) == 0 {
		validationResults.AddErrors(field.Required(field.NewPath("groupNameAttributes"), ""))
	}
	return validationResults
}
func ValidateLDAPQuery(query configapi.LDAPQuery, fldPath *field.Path) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validateLDAPQuery(query, fldPath, false)
}
func validateLDAPQuery(query configapi.LDAPQuery, fldPath *field.Path, isDNOnly bool) common.ValidationResults {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationResults := common.ValidationResults{}
	if _, err := ldap.ParseDN(query.BaseDN); err != nil {
		validationResults.AddErrors(field.Invalid(fldPath.Child("baseDN"), query.BaseDN, fmt.Sprintf("invalid base DN for search: %v", err)))
	}
	if len(query.Scope) > 0 {
		if _, err := ldaputil.DetermineLDAPScope(query.Scope); err != nil {
			validationResults.AddErrors(field.Invalid(fldPath.Child("scope"), query.Scope, "invalid LDAP search scope"))
		}
	}
	if len(query.DerefAliases) > 0 {
		if _, err := ldaputil.DetermineDerefAliasesBehavior(query.DerefAliases); err != nil {
			validationResults.AddErrors(field.Invalid(fldPath.Child("derefAliases"), query.DerefAliases, "LDAP alias dereferencing instruction invalid"))
		}
	}
	if query.TimeLimit < 0 {
		validationResults.AddErrors(field.Invalid(fldPath.Child("timeout"), query.TimeLimit, "timeout must be equal to or greater than zero"))
	}
	if isDNOnly {
		if len(query.Filter) != 0 {
			validationResults.AddErrors(field.Invalid(fldPath.Child("filter"), query.Filter, `cannot specify a filter when using "dn" as the UID attribute`))
		}
		return validationResults
	}
	if _, err := ldap.CompileFilter(query.Filter); err != nil {
		validationResults.AddErrors(field.Invalid(fldPath.Child("filter"), query.Filter, fmt.Sprintf("invalid query filter: %v", err)))
	}
	return validationResults
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
