package syncgroups

import (
	"fmt"
	"io"
	"k8s.io/klog"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

type GroupPruner interface{ Prune() (errors []error) }
type LDAPGroupPruner struct {
	GroupLister	interfaces.LDAPGroupLister
	GroupDetector	interfaces.LDAPGroupDetector
	GroupNameMapper	interfaces.LDAPGroupNameMapper
	GroupClient	userv1client.GroupInterface
	Host		string
	DryRun		bool
	Out		io.Writer
	Err		io.Writer
}

var _ GroupPruner = &LDAPGroupPruner{}

func (s *LDAPGroupPruner) Prune() []error {
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
	var errors []error
	klog.V(1).Infof("LDAPGroupPruner listing groups to prune with %v", s.GroupLister)
	ldapGroupUIDs, err := s.GroupLister.ListGroups()
	if err != nil {
		errors = append(errors, err)
		return errors
	}
	klog.V(1).Infof("LDAPGroupPruner will attempt to prune ldapGroupUIDs %v", ldapGroupUIDs)
	for _, ldapGroupUID := range ldapGroupUIDs {
		klog.V(1).Infof("Checking LDAP group %v", ldapGroupUID)
		exists, err := s.GroupDetector.Exists(ldapGroupUID)
		if err != nil {
			fmt.Fprintf(s.Err, "Error determining LDAP group existence for group %q: %v.\n", ldapGroupUID, err)
			errors = append(errors, err)
			continue
		}
		if exists {
			continue
		}
		groupName, err := s.GroupNameMapper.GroupNameFor(ldapGroupUID)
		if err != nil {
			fmt.Fprintf(s.Err, "Error determining OpenShift group name for LDAP group %q: %v.\n", ldapGroupUID, err)
			errors = append(errors, err)
			continue
		}
		if !s.DryRun {
			if err := s.GroupClient.Delete(groupName, nil); err != nil {
				fmt.Fprintf(s.Err, "Error pruning OpenShift group %q: %v.\n", groupName, err)
				errors = append(errors, err)
				continue
			}
		}
		fmt.Fprintf(s.Out, "group/%s\n", groupName)
	}
	return errors
}
