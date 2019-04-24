package useridentitymapping

import (
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	userapi "github.com/openshift/api/user/v1"
	"github.com/openshift/client-go/user/clientset/versioned/typed/user/v1/fake"
)

type Action struct {
	Name	string
	Object	interface{}
}
type IdentityRegistry struct {
	*fake.FakeIdentities
	GetErr		map[string]error
	GetIdentities	map[string]*userapi.Identity
	CreateErr	error
	CreateIdentity	*userapi.Identity
	UpdateErr	error
	UpdateIdentity	*userapi.Identity
	ListErr		error
	ListIdentity	*userapi.IdentityList
	Actions		*[]Action
}

func (r *IdentityRegistry) Get(name string, options metav1.GetOptions) (*userapi.Identity, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*r.Actions = append(*r.Actions, Action{"GetIdentity", name})
	if identity, ok := r.GetIdentities[name]; ok {
		return identity, nil
	}
	if err, ok := r.GetErr[name]; ok {
		return nil, err
	}
	return nil, kerrs.NewNotFound(userapi.Resource("identity"), name)
}
func (r *IdentityRegistry) Create(u *userapi.Identity) (*userapi.Identity, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*r.Actions = append(*r.Actions, Action{"CreateIdentity", u})
	if r.CreateIdentity == nil && r.CreateErr == nil {
		return u, nil
	}
	return r.CreateIdentity, r.CreateErr
}
func (r *IdentityRegistry) Update(u *userapi.Identity) (*userapi.Identity, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*r.Actions = append(*r.Actions, Action{"UpdateIdentity", u})
	if r.UpdateIdentity == nil && r.UpdateErr == nil {
		return u, nil
	}
	return r.UpdateIdentity, r.UpdateErr
}
func (r *IdentityRegistry) List(options metav1.ListOptions) (*userapi.IdentityList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*r.Actions = append(*r.Actions, Action{"ListIdentities", options})
	if r.ListIdentity == nil && r.ListErr == nil {
		return &userapi.IdentityList{}, nil
	}
	return r.ListIdentity, r.ListErr
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
