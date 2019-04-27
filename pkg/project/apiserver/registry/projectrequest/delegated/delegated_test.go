package delegated

import (
	"sync"
	"testing"
	"time"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	rbacv1listers "k8s.io/client-go/listers/rbac/v1"
	"github.com/go-openapi/errors"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
)

func TestDelegatedWait(t *testing.T) {
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
	cache := &testRoleBindingLister{}
	storage := &REST{roleBindings: cache}
	cache.namespacelister = &testRoleBindingNamespaceLister{}
	cache.namespacelister.bindings = map[string]*rbacv1.RoleBinding{}
	cache.namespacelister.bindings["anything"] = nil
	waitReturnedCh := waitForResultChannel(storage)
	select {
	case <-waitReturnedCh:
		t.Error("waitForRoleBinding() failed to block pending rolebinding creation")
	case <-time.After(1 * time.Second):
	}
	cache.addAdminRolebinding()
	select {
	case <-waitReturnedCh:
	case <-time.After(1 * time.Second):
		t.Error("waitForRoleBinding() failed to unblock after rolebinding creation")
	}
}
func waitForResultChannel(storage *REST) chan struct{} {
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
	ret := make(chan struct{})
	go func() {
		storage.waitForRoleBinding("foo", bootstrappolicy.AdminRoleName)
		close(ret)
	}()
	return ret
}

type testRoleBindingNamespaceLister struct {
	bindings	map[string]*rbacv1.RoleBinding
	lock		sync.Mutex
}

func (t *testRoleBindingNamespaceLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
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
	return ret, nil
}
func (t *testRoleBindingNamespaceLister) Get(name string) (*rbacv1.RoleBinding, error) {
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
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.bindings[bootstrappolicy.AdminRoleName] != nil {
		return t.bindings[bootstrappolicy.AdminRoleName], nil
	}
	return nil, errors.NotFound("could not find role " + bootstrappolicy.AdminRoleName)
}

type testRoleBindingLister struct {
	namespacelister *testRoleBindingNamespaceLister
}

func (t *testRoleBindingLister) RoleBindings(namespace string) rbacv1listers.RoleBindingNamespaceLister {
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
	return t.namespacelister
}
func (t *testRoleBindingLister) List(selector labels.Selector) ([]*rbacv1.RoleBinding, error) {
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
	return nil, nil
}
func (t *testRoleBindingLister) addAdminRolebinding() {
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
	t.namespacelister.lock.Lock()
	defer t.namespacelister.lock.Unlock()
	t.namespacelister.bindings[bootstrappolicy.AdminRoleName] = &rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: bootstrappolicy.AdminRoleName}}
}
