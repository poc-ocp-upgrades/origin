package sync

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	kerrs "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/cmd/server/apis/config"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/ldap"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"github.com/openshift/origin/pkg/oc/lib/groupsync"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/syncerror"
)

const SyncRecommendedName = "sync"

var (
	syncLong	= templates.LongDesc(`
    Sync OpenShift Groups with records from an external provider.

    In order to sync OpenShift Group records with those from an external provider, determine which Groups you wish
    to sync and where their records live. For instance, all or some groups may be selected from the current Groups
    stored in OpenShift that have been synced previously, or similarly all or some groups may be selected from those
    stored on an LDAP server. The path to a sync configuration file is required in order to describe how data is
    requested from the external record store and migrated to OpenShift records. Default behavior is to do a dry-run
    without changing OpenShift records. Passing '--confirm' will sync all groups from the LDAP server returned by the
    LDAP query templates.`)
	syncExamples	= templates.Examples(`
    # Sync all groups from an LDAP server
    %[1]s --sync-config=/path/to/ldap-sync-config.yaml --confirm

    # Sync all groups except the ones from the blacklist file from an LDAP server
    %[1]s --blacklist=/path/to/blacklist.txt --sync-config=/path/to/ldap-sync-config.yaml --confirm

    # Sync specific groups specified in a whitelist file with an LDAP server
    %[1]s --whitelist=/path/to/whitelist.txt --sync-config=/path/to/sync-config.yaml --confirm

    # Sync all OpenShift Groups that have been synced previously with an LDAP server
    %[1]s --type=openshift --sync-config=/path/to/ldap-sync-config.yaml --confirm

    # Sync specific OpenShift Groups if they have been synced previously with an LDAP server
    %[1]s groups/group1 groups/group2 groups/group3 --sync-config=/path/to/sync-config.yaml --confirm`)
)

type GroupSyncSource string

const (
	GroupSyncSourceLDAP		GroupSyncSource	= "ldap"
	GroupSyncSourceOpenShift	GroupSyncSource	= "openshift"
)

var AllowedSourceTypes = []string{string(GroupSyncSourceLDAP), string(GroupSyncSourceOpenShift)}

type SyncOptions struct {
	PrintFlags	*genericclioptions.PrintFlags
	Printer		printers.ResourcePrinter
	Source		GroupSyncSource
	Config		*config.LDAPSyncConfig
	ConfigFile	string
	Whitelist	[]string
	WhitelistFile	string
	Blacklist	[]string
	BlacklistFile	string
	Type		string
	Confirm		bool
	GroupClient	userv1typedclient.GroupsGetter
	genericclioptions.IOStreams
}

func NewSyncOptions(streams genericclioptions.IOStreams) *SyncOptions {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SyncOptions{Whitelist: []string{}, Type: string(GroupSyncSourceLDAP), PrintFlags: genericclioptions.NewPrintFlags("").WithTypeSetter(scheme.Scheme).WithDefaultOutput("yaml"), IOStreams: streams}
}
func NewCmdSync(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewSyncOptions(streams)
	cmd := &cobra.Command{Use: fmt.Sprintf("%s [--type=TYPE] [WHITELIST] [--whitelist=WHITELIST-FILE] --sync-config=CONFIG-FILE [--confirm]", name), Short: "Sync OpenShift groups with records from an external provider.", Long: syncLong, Example: fmt.Sprintf(syncExamples, fullName), Run: func(c *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(f, args))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.Run())
	}}
	cmd.Flags().StringVar(&o.WhitelistFile, "whitelist", o.WhitelistFile, "path to the group whitelist file")
	cmd.MarkFlagFilename("whitelist", "txt")
	cmd.Flags().StringVar(&o.BlacklistFile, "blacklist", o.BlacklistFile, "path to the group blacklist file")
	cmd.MarkFlagFilename("blacklist", "txt")
	cmd.Flags().StringVar(&o.ConfigFile, "sync-config", o.ConfigFile, "path to the sync config")
	cmd.MarkFlagFilename("sync-config", "yaml", "yml")
	cmd.Flags().StringVar(&o.Type, "type", o.Type, "which groups white- and blacklist entries refer to: "+strings.Join(AllowedSourceTypes, ","))
	cmd.Flags().BoolVar(&o.Confirm, "confirm", o.Confirm, "if true, modify OpenShift groups; if false, display results of a dry-run")
	return cmd
}
func (o *SyncOptions) Complete(f kcmdutil.Factory, args []string) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch o.Type {
	case string(GroupSyncSourceLDAP):
		o.Source = GroupSyncSourceLDAP
	case string(GroupSyncSourceOpenShift):
		o.Source = GroupSyncSourceOpenShift
	default:
		return fmt.Errorf("unrecognized --type %q; allowed types %v", o.Type, strings.Join(AllowedSourceTypes, ","))
	}
	var err error
	o.Config, err = decodeSyncConfigFromFile(o.ConfigFile)
	if err != nil {
		return err
	}
	if o.Source == GroupSyncSourceOpenShift {
		o.Whitelist, err = buildOpenShiftGroupNameList(args, o.WhitelistFile, o.Config.LDAPGroupUIDToOpenShiftGroupNameMapping)
		if err != nil {
			return err
		}
		o.Blacklist, err = buildOpenShiftGroupNameList([]string{}, o.BlacklistFile, o.Config.LDAPGroupUIDToOpenShiftGroupNameMapping)
		if err != nil {
			return err
		}
	} else {
		o.Whitelist, err = buildNameList(args, o.WhitelistFile)
		if err != nil {
			return err
		}
		o.Blacklist, err = buildNameList([]string{}, o.BlacklistFile)
		if err != nil {
			return err
		}
	}
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	o.GroupClient, err = userv1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	if !o.Confirm {
		o.PrintFlags.Complete("%s (dry run)")
	}
	o.Printer, err = o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	return nil
}
func buildOpenShiftGroupNameList(args []string, file string, nameMapping map[string]string) ([]string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	rawList, err := buildNameList(args, file)
	if err != nil {
		return nil, err
	}
	namesList, err := openshiftGroupNamesOnlyList(rawList)
	if err != nil {
		return nil, err
	}
	if len(nameMapping) > 0 {
		for i, name := range namesList {
			if nameOverride, ok := nameMapping[name]; ok {
				namesList[i] = nameOverride
			}
		}
	}
	return namesList, nil
}
func buildNameList(args []string, file string) ([]string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var list []string
	if len(args) > 0 {
		list = append(list, args...)
	}
	if len(file) != 0 {
		listData, err := readLines(file)
		if err != nil {
			return nil, err
		}
		list = append(list, listData...)
	}
	return list, nil
}
func decodeSyncConfigFromFile(configFile string) (*config.LDAPSyncConfig, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var config config.LDAPSyncConfig
	yamlConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", configFile, err)
	}
	jsonConfig, err := kyaml.ToJSON(yamlConfig)
	if err != nil {
		return nil, fmt.Errorf("could not parse file %s: %v", configFile, err)
	}
	if err := runtime.DecodeInto(configapilatest.Codec, jsonConfig, &config); err != nil {
		return nil, fmt.Errorf("could not decode file into config: %v", err)
	}
	return &config, nil
}
func openshiftGroupNamesOnlyList(list []string) ([]string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []string{}
	errs := []error{}
	for _, curr := range list {
		tokens := strings.SplitN(curr, "/", 2)
		if len(tokens) == 1 {
			ret = append(ret, tokens[0])
			continue
		}
		if tokens[0] != "group" && tokens[0] != "groups" {
			errs = append(errs, fmt.Errorf("%q is not a valid OpenShift group", curr))
			continue
		}
		ret = append(ret, tokens[1])
	}
	if len(errs) > 0 {
		return nil, kerrs.NewAggregate(errs)
	}
	return ret, nil
}
func readLines(path string) ([]string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", path, err)
	}
	rawLines := strings.Split(string(bytes), "\n")
	var trimmedLines []string
	for _, line := range rawLines {
		if len(strings.TrimSpace(line)) > 0 {
			trimmedLines = append(trimmedLines, strings.TrimSpace(line))
		}
	}
	return trimmedLines, nil
}
func ValidateSource(source GroupSyncSource) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sets.NewString(AllowedSourceTypes...).Has(string(source))
}
func (o *SyncOptions) Validate() error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !ValidateSource(o.Source) {
		return fmt.Errorf("sync source must be one of the following: %v", strings.Join(AllowedSourceTypes, ","))
	}
	results := ldap.ValidateLDAPSyncConfig(o.Config)
	if o.GroupClient == nil {
		results.Errors = append(results.Errors, field.Required(field.NewPath("groupInterface"), ""))
	}
	if len(results.Errors) > 0 {
		return fmt.Errorf("validation of LDAP sync config failed: %v", results.Errors.ToAggregate())
	}
	return nil
}
func (o *SyncOptions) CreateErrorHandler() syncerror.Handler {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	components := []syncerror.Handler{}
	if o.Config.RFC2307Config != nil {
		if o.Config.RFC2307Config.TolerateMemberOutOfScopeErrors {
			components = append(components, syncerror.NewMemberLookupOutOfBoundsSuppressor(o.ErrOut))
		}
		if o.Config.RFC2307Config.TolerateMemberNotFoundErrors {
			components = append(components, syncerror.NewMemberLookupMemberNotFoundSuppressor(o.ErrOut))
		}
	}
	return syncerror.NewCompoundHandler(components...)
}
func (o *SyncOptions) Run() error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	bindPassword, err := config.ResolveStringValue(o.Config.BindPassword)
	if err != nil {
		return err
	}
	clientConfig, err := ldaputil.NewLDAPClientConfig(o.Config.URL, o.Config.BindDN, bindPassword, o.Config.CA, o.Config.Insecure)
	if err != nil {
		return fmt.Errorf("could not determine LDAP client configuration: %v", err)
	}
	errorHandler := o.CreateErrorHandler()
	syncBuilder, err := buildSyncBuilder(clientConfig, o.Config, errorHandler)
	if err != nil {
		return err
	}
	syncer := &syncgroups.LDAPGroupSyncer{Host: clientConfig.Host(), GroupClient: o.GroupClient.Groups(), DryRun: !o.Confirm, Out: o.Out, Err: o.ErrOut}
	switch o.Source {
	case GroupSyncSourceOpenShift:
		listerMapper, err := getOpenShiftGroupListerMapper(clientConfig.Host(), o)
		if err != nil {
			return err
		}
		syncer.GroupLister = listerMapper
		syncer.GroupNameMapper = listerMapper
	case GroupSyncSourceLDAP:
		syncer.GroupLister, err = getLDAPGroupLister(syncBuilder, o)
		if err != nil {
			return err
		}
		syncer.GroupNameMapper, err = getGroupNameMapper(syncBuilder, o)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid group source: %v", o.Source)
	}
	syncer.GroupMemberExtractor, err = syncBuilder.GetGroupMemberExtractor()
	if err != nil {
		return err
	}
	syncer.UserNameMapper, err = syncBuilder.GetUserNameMapper()
	if err != nil {
		return err
	}
	openshiftGroups, syncErrors := syncer.Sync()
	if !o.Confirm {
		list := &unstructured.UnstructuredList{Object: map[string]interface{}{"kind": "List", "apiVersion": "v1", "metadata": map[string]interface{}{}}}
		for _, item := range openshiftGroups {
			unstructuredItem, err := runtime.DefaultUnstructuredConverter.ToUnstructured(item)
			if err != nil {
				return err
			}
			list.Items = append(list.Items, unstructured.Unstructured{Object: unstructuredItem})
		}
		if err := o.Printer.PrintObj(list, o.Out); err != nil {
			return err
		}
	}
	for _, err := range syncErrors {
		fmt.Fprintf(o.ErrOut, "%s\n", err)
	}
	return kerrs.NewAggregate(syncErrors)
}
func buildSyncBuilder(clientConfig ldapclient.Config, syncConfig *config.LDAPSyncConfig, errorHandler syncerror.Handler) (SyncBuilder, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case syncConfig.RFC2307Config != nil:
		return &RFC2307Builder{ClientConfig: clientConfig, Config: syncConfig.RFC2307Config, ErrorHandler: errorHandler}, nil
	case syncConfig.ActiveDirectoryConfig != nil:
		return &ADBuilder{ClientConfig: clientConfig, Config: syncConfig.ActiveDirectoryConfig}, nil
	case syncConfig.AugmentedActiveDirectoryConfig != nil:
		return &AugmentedADBuilder{ClientConfig: clientConfig, Config: syncConfig.AugmentedActiveDirectoryConfig}, nil
	default:
		return nil, errors.New("invalid sync config type")
	}
}
func getOpenShiftGroupListerMapper(host string, info OpenShiftGroupNameRestrictions) (interfaces.LDAPGroupListerNameMapper, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(info.GetWhitelist()) != 0 {
		return syncgroups.NewOpenShiftGroupLister(info.GetWhitelist(), info.GetBlacklist(), host, info.GetClient()), nil
	} else {
		return syncgroups.NewAllOpenShiftGroupLister(info.GetBlacklist(), host, info.GetClient()), nil
	}
}
func getLDAPGroupLister(syncBuilder SyncBuilder, info GroupNameRestrictions) (interfaces.LDAPGroupLister, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(info.GetWhitelist()) != 0 {
		ldapWhitelist := syncgroups.NewLDAPWhitelistGroupLister(info.GetWhitelist())
		if len(info.GetBlacklist()) == 0 {
			return ldapWhitelist, nil
		}
		return syncgroups.NewLDAPBlacklistGroupLister(info.GetBlacklist(), ldapWhitelist), nil
	}
	syncLister, err := syncBuilder.GetGroupLister()
	if err != nil {
		return nil, err
	}
	if len(info.GetBlacklist()) == 0 {
		return syncLister, nil
	}
	return syncgroups.NewLDAPBlacklistGroupLister(info.GetBlacklist(), syncLister), nil
}
func getGroupNameMapper(syncBuilder SyncBuilder, info MappedNameRestrictions) (interfaces.LDAPGroupNameMapper, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	syncNameMapper, err := syncBuilder.GetGroupNameMapper()
	if err != nil {
		return nil, err
	}
	if len(info.GetGroupNameMappings()) > 0 {
		userDefinedMapper := syncgroups.NewUserDefinedGroupNameMapper(info.GetGroupNameMappings())
		if syncNameMapper == nil {
			return userDefinedMapper, nil
		}
		return &syncgroups.UnionGroupNameMapper{GroupNameMappers: []interfaces.LDAPGroupNameMapper{userDefinedMapper, syncNameMapper}}, nil
	}
	return syncNameMapper, nil
}
func (o *SyncOptions) GetWhitelist() []string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Whitelist
}
func (o *SyncOptions) GetBlacklist() []string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Blacklist
}
func (o *SyncOptions) GetClient() userv1typedclient.GroupInterface {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.GroupClient.Groups()
}
func (o *SyncOptions) GetGroupNameMappings() map[string]string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Config.LDAPGroupUIDToOpenShiftGroupNameMapping
}
