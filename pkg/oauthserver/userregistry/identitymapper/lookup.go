package identitymapper

import (
	"fmt"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuser "k8s.io/apiserver/pkg/authentication/user"
)

var _ = authapi.UserIdentityMapper(&lookupIdentityMapper{})

type lookupIdentityMapper struct {
	mappings userclient.UserIdentityMappingInterface
	users    userclient.UserInterface
}

func (p *lookupIdentityMapper) UserFor(info authapi.UserIdentityInfo) (kuser.Info, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mapping, err := p.mappings.Get(info.GetIdentityName(), metav1.GetOptions{})
	if err != nil {
		return nil, NewLookupError(info, err)
	}
	u, err := p.users.Get(mapping.User.Name, metav1.GetOptions{})
	if err != nil {
		return nil, NewLookupError(info, err)
	}
	return userToInfo(u), nil
}

type lookupError struct {
	Identity authapi.UserIdentityInfo
	CausedBy error
}

func IsLookupError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, ok := err.(lookupError)
	return ok
}
func NewLookupError(info authapi.UserIdentityInfo, err error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return lookupError{Identity: info, CausedBy: err}
}
func (c lookupError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("lookup of user for %q failed: %v", c.Identity.GetIdentityName(), c.CausedBy)
}
