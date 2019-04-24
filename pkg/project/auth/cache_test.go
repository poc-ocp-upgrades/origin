package auth

import (
	"fmt"
	"strconv"
	"testing"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/controller"
)

type mockReview struct {
	users	[]string
	groups	[]string
	err	string
}

func (r *mockReview) Users() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.users
}
func (r *mockReview) Groups() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.groups
}
func (r *mockReview) EvaluationError() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.err
}

var (
	alice	= &user.DefaultInfo{Name: "Alice", UID: "alice-uid", Groups: []string{}}
	bob	= &user.DefaultInfo{Name: "Bob", UID: "bob-uid", Groups: []string{"employee"}}
	eve	= &user.DefaultInfo{Name: "Eve", UID: "eve-uid", Groups: []string{"employee"}}
	frank	= &user.DefaultInfo{Name: "Frank", UID: "frank-uid", Groups: []string{}}
)

type mockReviewer struct{ expectedResults map[string]*mockReview }

func (mr *mockReviewer) Review(name string) (Review, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	review := mr.expectedResults[name]
	if review == nil {
		return nil, fmt.Errorf("Item %s does not exist", name)
	}
	return review, nil
}
func validateList(t *testing.T, lister Lister, user user.Info, expectedSet sets.String) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaceList, err := lister.List(user, labels.Everything())
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	results := sets.String{}
	for _, namespace := range namespaceList.Items {
		results.Insert(namespace.Name)
	}
	if results.Len() != expectedSet.Len() || !results.HasAll(expectedSet.List()...) {
		t.Errorf("User %v, Expected: %v, Actual: %v", user.GetName(), expectedSet, results)
	}
}
func TestSyncNamespace(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaceList := corev1.NamespaceList{Items: []corev1.Namespace{{ObjectMeta: metav1.ObjectMeta{Name: "foo", ResourceVersion: "1"}}, {ObjectMeta: metav1.ObjectMeta{Name: "bar", ResourceVersion: "2"}}, {ObjectMeta: metav1.ObjectMeta{Name: "car", ResourceVersion: "3"}}}}
	mockKubeClient := fake.NewSimpleClientset(&namespaceList)
	reviewer := &mockReviewer{expectedResults: map[string]*mockReview{"foo": {users: []string{alice.GetName(), bob.GetName()}, groups: eve.GetGroups()}, "bar": {users: []string{frank.GetName(), eve.GetName()}, groups: []string{"random"}}, "car": {users: []string{}, groups: []string{}}}}
	informers := informers.NewSharedInformerFactory(mockKubeClient, controller.NoResyncPeriodFunc())
	nsIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	nsLister := corev1listers.NewNamespaceLister(nsIndexer)
	authorizationCache := NewAuthorizationCache(nsLister, informers.Core().V1().Namespaces().Informer(), reviewer, informers.Rbac().V1())
	for i := range namespaceList.Items {
		nsIndexer.Add(&namespaceList.Items[i])
	}
	authorizationCache.synchronize()
	validateList(t, authorizationCache, alice, sets.NewString("foo"))
	validateList(t, authorizationCache, bob, sets.NewString("foo"))
	validateList(t, authorizationCache, eve, sets.NewString("foo", "bar"))
	validateList(t, authorizationCache, frank, sets.NewString("bar"))
	reviewer.expectedResults["foo"].users = []string{bob.GetName()}
	reviewer.expectedResults["foo"].groups = []string{"random"}
	reviewer.expectedResults["bar"].users = []string{alice.GetName(), eve.GetName()}
	reviewer.expectedResults["bar"].groups = []string{"employee"}
	reviewer.expectedResults["car"].users = []string{bob.GetName(), eve.GetName()}
	reviewer.expectedResults["car"].groups = []string{"employee"}
	for i := range namespaceList.Items {
		namespace := namespaceList.Items[i]
		oldVersion, err := strconv.Atoi(namespace.ResourceVersion)
		if err != nil {
			t.Errorf("Bad test setup, resource versions should be numbered, %v", err)
		}
		newVersion := strconv.Itoa(oldVersion + 1)
		namespace.ResourceVersion = newVersion
		nsIndexer.Add(&namespace)
	}
	authorizationCache.synchronize()
	validateList(t, authorizationCache, alice, sets.NewString("bar"))
	validateList(t, authorizationCache, bob, sets.NewString("foo", "bar", "car"))
	validateList(t, authorizationCache, eve, sets.NewString("bar", "car"))
	validateList(t, authorizationCache, frank, sets.NewString())
}
