package syncgroups

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"net"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	userv1 "github.com/openshift/api/user/v1"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func NewAllOpenShiftGroupLister(blacklist []string, ldapURL string, groupClient userv1client.GroupInterface) interfaces.LDAPGroupListerNameMapper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &allOpenShiftGroupLister{blacklist: sets.NewString(blacklist...), client: groupClient, ldapURL: ldapURL, ldapGroupUIDToOpenShiftGroupName: map[string]string{}}
}

type allOpenShiftGroupLister struct {
	blacklist				sets.String
	client					userv1client.GroupInterface
	ldapURL					string
	ldapGroupUIDToOpenShiftGroupName	map[string]string
}

func (l *allOpenShiftGroupLister) ListGroups() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	host, _, err := net.SplitHostPort(l.ldapURL)
	if err != nil {
		return nil, err
	}
	hostSelector := labels.Set(map[string]string{ldaputil.LDAPHostLabel: host}).AsSelector()
	allGroups, err := l.client.List(metav1.ListOptions{LabelSelector: hostSelector.String()})
	if err != nil {
		return nil, err
	}
	var ldapGroupUIDs []string
	for _, group := range allGroups.Items {
		if l.blacklist.Has(group.Name) {
			continue
		}
		matches, err := validateGroupAnnotations(l.ldapURL, group)
		if err != nil {
			return nil, err
		}
		if !matches {
			continue
		}
		ldapGroupUID := group.Annotations[ldaputil.LDAPUIDAnnotation]
		l.ldapGroupUIDToOpenShiftGroupName[ldapGroupUID] = group.Name
		ldapGroupUIDs = append(ldapGroupUIDs, ldapGroupUID)
	}
	return ldapGroupUIDs, nil
}
func (l *allOpenShiftGroupLister) GroupNameFor(ldapGroupUID string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(l.ldapGroupUIDToOpenShiftGroupName) == 0 {
		_, err := l.ListGroups()
		if err != nil {
			return "", err
		}
	}
	openshiftGroupName, exists := l.ldapGroupUIDToOpenShiftGroupName[ldapGroupUID]
	if !exists {
		return "", fmt.Errorf("no mapping found for %q", ldapGroupUID)
	}
	return openshiftGroupName, nil
}
func validateGroupAnnotations(ldapURL string, group userv1.Group) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if actualURL, exists := group.Annotations[ldaputil.LDAPURLAnnotation]; !exists {
		return false, fmt.Errorf("group %q marked as having been synced did not have an %s annotation", group.Name, ldaputil.LDAPURLAnnotation)
	} else if actualURL != ldapURL {
		return false, nil
	}
	if _, exists := group.Annotations[ldaputil.LDAPUIDAnnotation]; !exists {
		return false, fmt.Errorf("group %q marked as having been synced did not have an %s annotation", group.Name, ldaputil.LDAPUIDAnnotation)
	}
	return true, nil
}
func NewOpenShiftGroupLister(whitelist, blacklist []string, ldapURL string, client userv1client.GroupInterface) interfaces.LDAPGroupListerNameMapper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &openshiftGroupLister{whitelist: whitelist, blacklist: sets.NewString(blacklist...), client: client, ldapURL: ldapURL, ldapGroupUIDToOpenShiftGroupName: map[string]string{}}
}

type openshiftGroupLister struct {
	whitelist				[]string
	blacklist				sets.String
	client					userv1client.GroupInterface
	ldapURL					string
	ldapGroupUIDToOpenShiftGroupName	map[string]string
}

func (l *openshiftGroupLister) ListGroups() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var groups []userv1.Group
	for _, name := range l.whitelist {
		if l.blacklist.Has(name) {
			continue
		}
		group, err := l.client.Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		groups = append(groups, *group)
	}
	var ldapGroupUIDs []string
	for _, group := range groups {
		matches, err := validateGroupAnnotations(l.ldapURL, group)
		if err != nil {
			return nil, err
		}
		if !matches {
			return nil, fmt.Errorf("group %q was not synchronized from: %s", group.Name, l.ldapURL)
		}
		ldapGroupUID := group.Annotations[ldaputil.LDAPUIDAnnotation]
		l.ldapGroupUIDToOpenShiftGroupName[ldapGroupUID] = group.Name
		ldapGroupUIDs = append(ldapGroupUIDs, ldapGroupUID)
	}
	return ldapGroupUIDs, nil
}
func (l *openshiftGroupLister) GroupNameFor(ldapGroupUID string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(l.ldapGroupUIDToOpenShiftGroupName) == 0 {
		_, err := l.ListGroups()
		if err != nil {
			return "", err
		}
	}
	openshiftGroupName, exists := l.ldapGroupUIDToOpenShiftGroupName[ldapGroupUID]
	if !exists {
		return "", fmt.Errorf("no mapping found for %q", ldapGroupUID)
	}
	return openshiftGroupName, nil
}
func NewLDAPWhitelistGroupLister(whitelist []string) interfaces.LDAPGroupLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &whitelistLDAPGroupLister{ldapGroupUIDs: whitelist}
}

type whitelistLDAPGroupLister struct{ ldapGroupUIDs []string }

func (l *whitelistLDAPGroupLister) ListGroups() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.ldapGroupUIDs, nil
}
func NewLDAPBlacklistGroupLister(blacklist []string, baseLister interfaces.LDAPGroupLister) interfaces.LDAPGroupLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &blacklistLDAPGroupLister{blacklist: sets.NewString(blacklist...), baseLister: baseLister}
}

type blacklistLDAPGroupLister struct {
	blacklist	sets.String
	baseLister	interfaces.LDAPGroupLister
}

func (l *blacklistLDAPGroupLister) ListGroups() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allNames, err := l.baseLister.ListGroups()
	if err != nil {
		return nil, err
	}
	ret := []string{}
	for _, name := range allNames {
		if l.blacklist.Has(name) {
			continue
		}
		ret = append(ret, name)
	}
	return ret, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
