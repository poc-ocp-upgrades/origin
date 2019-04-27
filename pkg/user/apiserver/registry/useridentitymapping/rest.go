package useridentitymapping

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"github.com/openshift/api/user"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	userinternal "github.com/openshift/origin/pkg/user/apis/user"
)

type REST struct {
	userClient	userclient.UserInterface
	identityClient	userclient.IdentityInterface
}

var _ rest.Getter = &REST{}
var _ rest.CreaterUpdater = &REST{}
var _ rest.GracefulDeleter = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(userClient userclient.UserInterface, identityClient userclient.IdentityInterface) *REST {
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
	return &REST{userClient: userClient, identityClient: identityClient}
}
func (r *REST) New() runtime.Object {
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
	return &userinternal.UserIdentityMapping{}
}
func (s *REST) NamespaceScoped() bool {
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
	return false
}
func (s *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
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
	_, _, _, _, mapping, err := s.getRelatedObjects(ctx, name, options)
	return mapping, err
}
func (s *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
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
	mapping, ok := obj.(*userinternal.UserIdentityMapping)
	if !ok {
		return nil, kerrs.NewBadRequest("invalid type")
	}
	Strategy.PrepareForCreate(ctx, mapping)
	createdMapping, _, err := s.createOrUpdate(ctx, obj, true)
	return createdMapping, err
}
func (s *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, _ rest.ValidateObjectFunc, _ rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
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
	obj, err := objInfo.UpdatedObject(ctx, nil)
	if err != nil {
		return nil, false, err
	}
	mapping, ok := obj.(*userinternal.UserIdentityMapping)
	if !ok {
		return nil, false, kerrs.NewBadRequest("invalid type")
	}
	Strategy.PrepareForUpdate(ctx, mapping, nil)
	return s.createOrUpdate(ctx, mapping, false)
}
func (s *REST) createOrUpdate(ctx context.Context, obj runtime.Object, forceCreate bool) (runtime.Object, bool, error) {
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
	mapping := obj.(*userinternal.UserIdentityMapping)
	identity, identityErr, oldUser, oldUserErr, oldMapping, oldMappingErr := s.getRelatedObjects(ctx, mapping.Name, &metav1.GetOptions{})
	if !(oldMappingErr == nil || kerrs.IsNotFound(oldMappingErr)) {
		return nil, false, oldMappingErr
	}
	if !(identityErr == nil || kerrs.IsNotFound(identityErr)) {
		return nil, false, identityErr
	}
	if !(oldUserErr == nil || kerrs.IsNotFound(oldUserErr)) {
		return nil, false, oldUserErr
	}
	if forceCreate && oldMappingErr == nil {
		return nil, false, kerrs.NewAlreadyExists(userapi.Resource("useridentitymapping"), oldMapping.Name)
	}
	creating := forceCreate || kerrs.IsNotFound(oldMappingErr)
	if creating {
		if err := rest.BeforeCreate(Strategy, ctx, mapping); err != nil {
			return nil, false, err
		}
		if len(mapping.ResourceVersion) > 0 {
			return nil, false, kerrs.NewNotFound(userapi.Resource("useridentitymapping"), mapping.Name)
		}
	} else {
		if err := rest.BeforeUpdate(Strategy, ctx, mapping, oldMapping); err != nil {
			return nil, false, err
		}
		if len(mapping.ResourceVersion) > 0 && mapping.ResourceVersion != oldMapping.ResourceVersion {
			return nil, false, kerrs.NewConflict(userapi.Resource("useridentitymapping"), mapping.Name, fmt.Errorf("the resource was updated to %s", oldMapping.ResourceVersion))
		}
		if mapping.User.Name == oldMapping.User.Name {
			return oldMapping, false, nil
		}
	}
	if kerrs.IsNotFound(identityErr) {
		errs := field.ErrorList{field.Invalid(field.NewPath("identity", "name"), mapping.Identity.Name, "referenced identity does not exist")}
		return nil, false, kerrs.NewInvalid(user.Kind("UserIdentityMapping"), mapping.Name, errs)
	}
	newUser, err := s.userClient.Get(mapping.User.Name, metav1.GetOptions{})
	if kerrs.IsNotFound(err) {
		errs := field.ErrorList{field.Invalid(field.NewPath("user", "name"), mapping.User.Name, "referenced user does not exist")}
		return nil, false, kerrs.NewInvalid(user.Kind("UserIdentityMapping"), mapping.Name, errs)
	}
	if err != nil {
		return nil, false, err
	}
	if addIdentityToUser(identity, newUser) {
		if _, err := s.userClient.Update(newUser); err != nil {
			return nil, false, err
		}
	}
	if setIdentityUser(identity, newUser) {
		if updatedIdentity, err := s.identityClient.Update(identity); err != nil {
			return nil, false, err
		} else {
			identity = updatedIdentity
		}
	}
	if oldUser != nil && removeIdentityFromUser(identity, oldUser) {
		if _, err := s.userClient.Update(oldUser); err != nil {
			utilruntime.HandleError(fmt.Errorf("error removing identity reference %s from user %s: %v", identity.Name, oldUser.Name, err))
		}
	}
	updatedMapping, err := mappingFor(newUser, identity)
	return updatedMapping, creating, err
}
func (s *REST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
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
	identity, _, user, _, _, mappingErr := s.getRelatedObjects(ctx, name, &metav1.GetOptions{})
	if mappingErr != nil {
		return nil, false, mappingErr
	}
	if removeIdentityFromUser(identity, user) {
		if _, err := s.userClient.Update(user); err != nil {
			return nil, false, err
		}
	}
	if unsetIdentityUser(identity) {
		if _, err := s.identityClient.Update(identity); err != nil {
			utilruntime.HandleError(fmt.Errorf("error removing user reference %s from identity %s: %v", user.Name, identity.Name, err))
		}
	}
	return &metav1.Status{Status: metav1.StatusSuccess}, true, nil
}
func (s *REST) getRelatedObjects(ctx context.Context, name string, options *metav1.GetOptions) (identity *userapi.Identity, identityErr error, user *userapi.User, userErr error, mapping *userinternal.UserIdentityMapping, mappingErr error) {
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
	identityErr = kerrs.NewNotFound(userapi.Resource("identity"), name)
	userErr = kerrs.NewNotFound(userapi.Resource("user"), "")
	mappingErr = kerrs.NewNotFound(userapi.Resource("useridentitymapping"), name)
	identity, identityErr = s.identityClient.Get(name, *options)
	if identityErr != nil {
		return
	}
	if !hasUserMapping(identity) {
		return
	}
	user, userErr = s.userClient.Get(identity.User.Name, *options)
	if userErr != nil {
		return
	}
	if !identityReferencesUser(identity, user) {
		return
	}
	if !userReferencesIdentity(user, identity) {
		return
	}
	mapping, mappingErr = mappingFor(user, identity)
	return
}
func hasUserMapping(identity *userapi.Identity) bool {
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
	return len(identity.User.Name) > 0
}
func identityReferencesUser(identity *userapi.Identity, user *userapi.User) bool {
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
	return identity.User.Name == user.Name && identity.User.UID == user.UID
}
func userReferencesIdentity(user *userapi.User, identity *userapi.Identity) bool {
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
	return sets.NewString(user.Identities...).Has(identity.Name)
}
func addIdentityToUser(identity *userapi.Identity, user *userapi.User) bool {
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
	identities := sets.NewString(user.Identities...)
	if identities.Has(identity.Name) {
		return false
	}
	identities.Insert(identity.Name)
	user.Identities = identities.List()
	return true
}
func removeIdentityFromUser(identity *userapi.Identity, user *userapi.User) bool {
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
	identities := sets.NewString(user.Identities...)
	if !identities.Has(identity.Name) {
		return false
	}
	identities.Delete(identity.Name)
	user.Identities = identities.List()
	return true
}
func setIdentityUser(identity *userapi.Identity, user *userapi.User) bool {
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
	if identityReferencesUser(identity, user) {
		return false
	}
	identity.User = corev1.ObjectReference{Name: user.Name, UID: user.UID}
	return true
}
func unsetIdentityUser(identity *userapi.Identity) bool {
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
	if !hasUserMapping(identity) {
		return false
	}
	identity.User = corev1.ObjectReference{}
	return true
}
func mappingFor(user *userapi.User, identity *userapi.Identity) (*userinternal.UserIdentityMapping, error) {
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
	return &userinternal.UserIdentityMapping{ObjectMeta: metav1.ObjectMeta{Name: identity.Name, ResourceVersion: identity.ResourceVersion, UID: identity.UID}, Identity: kapi.ObjectReference{Name: identity.Name, UID: identity.UID}, User: kapi.ObjectReference{Name: user.Name, UID: user.UID}}, nil
}
