package identitymapper

import (
	"context"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	corev1 "k8s.io/api/core/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	kuser "k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog"
)

type UserForNewIdentityGetter interface {
	UserForNewIdentity(ctx context.Context, preferredUserName string, identity *userapi.Identity) (*userapi.User, error)
}

var _ = authapi.UserIdentityMapper(&provisioningIdentityMapper{})

type provisioningIdentityMapper struct {
	identity             userclient.IdentityInterface
	user                 userclient.UserInterface
	provisioningStrategy UserForNewIdentityGetter
}

func (p *provisioningIdentityMapper) UserFor(info authapi.UserIdentityInfo) (kuser.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.userForWithRetries(info, 3)
}
func (p *provisioningIdentityMapper) userForWithRetries(info authapi.UserIdentityInfo, allowedRetries int) (kuser.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := apirequest.NewContext()
	identity, err := p.identity.Get(info.GetIdentityName(), metav1.GetOptions{})
	if kerrs.IsNotFound(err) {
		user, err := p.createIdentityAndMapping(ctx, info)
		if (kerrs.IsAlreadyExists(err) || kerrs.IsConflict(err)) && allowedRetries > 0 {
			return p.userForWithRetries(info, allowedRetries-1)
		}
		return user, err
	}
	if err != nil {
		return nil, err
	}
	return p.getMapping(ctx, identity)
}
func (p *provisioningIdentityMapper) createIdentityAndMapping(ctx context.Context, info authapi.UserIdentityInfo) (kuser.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	identity := &userapi.Identity{ObjectMeta: metav1.ObjectMeta{Name: info.GetIdentityName()}, ProviderName: info.GetProviderName(), ProviderUserName: info.GetProviderUserName(), Extra: info.GetExtra()}
	persistedUser, err := p.provisioningStrategy.UserForNewIdentity(ctx, getPreferredUserName(identity), identity)
	if err != nil {
		return nil, err
	}
	identity.User = corev1.ObjectReference{Name: persistedUser.Name, UID: persistedUser.UID}
	if _, err := p.identity.Create(identity); err != nil {
		return nil, err
	}
	return userToInfo(persistedUser), nil
}
func (p *provisioningIdentityMapper) getMapping(ctx context.Context, identity *userapi.Identity) (kuser.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(identity.User.Name) == 0 {
		return nil, kerrs.NewNotFound(userapi.Resource("useridentitymapping"), identity.Name)
	}
	u, err := p.user.Get(identity.User.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if u.UID != identity.User.UID {
		klog.Errorf("identity.user.uid (%s) and user.uid (%s) do not match for identity %s", identity.User.UID, u.UID, identity.Name)
		return nil, kerrs.NewNotFound(userapi.Resource("useridentitymapping"), identity.Name)
	}
	if !sets.NewString(u.Identities...).Has(identity.Name) {
		klog.Errorf("user.identities (%#v) does not include identity (%s)", u, identity.Name)
		return nil, kerrs.NewNotFound(userapi.Resource("useridentitymapping"), identity.Name)
	}
	return userToInfo(u), nil
}
func getPreferredUserName(identity *userapi.Identity) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if login, ok := identity.Extra[authapi.IdentityPreferredUsernameKey]; ok && len(login) > 0 {
		return login
	}
	return identity.ProviderUserName
}
