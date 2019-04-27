package etcd

import (
	"testing"
	authorizationapi "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	genericregistrytest "k8s.io/apiserver/pkg/registry/generic/testing"
	"k8s.io/apiserver/pkg/registry/rest"
	etcdtesting "k8s.io/apiserver/pkg/storage/etcd/testing"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	routetypes "github.com/openshift/origin/pkg/route"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	_ "github.com/openshift/origin/pkg/route/apis/route/install"
	"github.com/openshift/origin/pkg/route/apiserver/registry/route"
	"k8s.io/apiserver/pkg/registry/generic"
)

type testAllocator struct {
	Hostname	string
	Err		error
	Allocate	bool
	Generate	bool
}

func (a *testAllocator) AllocateRouterShard(*routeapi.Route) (*routeapi.RouterShard, error) {
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
	a.Allocate = true
	return nil, a.Err
}
func (a *testAllocator) GenerateHostname(*routeapi.Route, *routeapi.RouterShard) string {
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
	a.Generate = true
	return a.Hostname
}

type testSAR struct {
	allow	bool
	err	error
	sar	*authorizationapi.SubjectAccessReview
}

func (t *testSAR) Create(subjectAccessReview *authorizationapi.SubjectAccessReview) (*authorizationapi.SubjectAccessReview, error) {
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
	t.sar = subjectAccessReview
	return &authorizationapi.SubjectAccessReview{Status: authorizationapi.SubjectAccessReviewStatus{Allowed: t.allow}}, t.err
}
func newStorage(t *testing.T, allocator routetypes.RouteAllocator) (*REST, *etcdtesting.EtcdTestServer) {
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
	server, etcdStorage := etcdtesting.NewUnsecuredEtcd3TestClientServer(t)
	etcdStorage.Codec = legacyscheme.Codecs.LegacyCodec(schema.GroupVersion{Group: "route.openshift.io", Version: "v1"})
	restOptions := generic.RESTOptions{StorageConfig: etcdStorage, Decorator: generic.UndecoratedStorage, DeleteCollectionWorkers: 1, ResourcePrefix: "routes"}
	storage, _, err := NewREST(restOptions, allocator, &testSAR{allow: true})
	if err != nil {
		t.Fatal(err)
	}
	return storage, server
}
func validRoute() *routeapi.Route {
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
	return &routeapi.Route{ObjectMeta: metav1.ObjectMeta{Name: "foo"}, Spec: routeapi.RouteSpec{To: routeapi.RouteTargetReference{Name: "test", Kind: "Service"}}}
}
func TestCreate(t *testing.T) {
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
	storage, server := newStorage(t, nil)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestCreate(validRoute(), &routeapi.Route{ObjectMeta: metav1.ObjectMeta{Name: "_-a123-a_"}})
}
func TestCreateWithAllocation(t *testing.T) {
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
	allocator := &testAllocator{Hostname: "bar"}
	storage, server := newStorage(t, allocator)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	obj, err := storage.Create(apirequest.NewDefaultContext(), validRoute(), rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("unable to create object: %v", err)
	}
	result := obj.(*routeapi.Route)
	if result.Spec.Host != "bar" {
		t.Fatalf("unexpected route: %#v", result)
	}
	if v, ok := result.Annotations[route.HostGeneratedAnnotationKey]; !ok || v != "true" {
		t.Fatalf("unexpected route: %#v", result)
	}
	if !allocator.Allocate || !allocator.Generate {
		t.Fatalf("unexpected allocator: %#v", allocator)
	}
}
func TestUpdate(t *testing.T) {
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
	storage, server := newStorage(t, nil)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestUpdate(validRoute(), func(obj runtime.Object) runtime.Object {
		object := obj.(*routeapi.Route)
		if object.Annotations == nil {
			object.Annotations = map[string]string{}
		}
		object.Annotations["updated"] = "true"
		return object
	}, func(obj runtime.Object) runtime.Object {
		object := obj.(*routeapi.Route)
		object.Spec.Path = "invalid/path"
		return object
	})
}
func TestList(t *testing.T) {
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
	storage, server := newStorage(t, nil)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestList(validRoute())
}
func TestGet(t *testing.T) {
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
	storage, server := newStorage(t, &testAllocator{})
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestGet(validRoute())
}
func TestDelete(t *testing.T) {
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
	storage, server := newStorage(t, nil)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestDelete(validRoute())
}
func TestWatch(t *testing.T) {
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
	storage, server := newStorage(t, nil)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	valid := validRoute()
	valid.Name = "foo"
	valid.Labels = map[string]string{"foo": "bar"}
	test.TestWatch(valid, []labels.Set{{"foo": "bar"}}, []labels.Set{{"foo": "baz"}}, []fields.Set{{"metadata.name": "foo"}}, []fields.Set{{"metadata.name": "bar"}})
}
