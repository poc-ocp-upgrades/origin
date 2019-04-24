package etcd

import (
	"testing"
	authorizationapi "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistrytest "k8s.io/apiserver/pkg/registry/generic/testing"
	"k8s.io/apiserver/pkg/registry/rest"
	etcdtesting "k8s.io/apiserver/pkg/storage/etcd/testing"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation/fake"
	admfake "github.com/openshift/origin/pkg/image/apiserver/admission/fake"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	_ "github.com/openshift/origin/pkg/api/install"
)

const (
	name = "foo"
)

var (
	testDefaultRegistry	= func() (string, bool) {
		return "test", true
	}
	noDefaultRegistry	= func() (string, bool) {
		return "", false
	}
)

type fakeSubjectAccessReviewRegistry struct {
	err			error
	allow			bool
	request			*authorizationapi.SubjectAccessReview
	requestNamespace	string
}

func (f *fakeSubjectAccessReviewRegistry) Create(subjectAccessReview *authorizationapi.SubjectAccessReview) (*authorizationapi.SubjectAccessReview, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.request = subjectAccessReview
	f.requestNamespace = subjectAccessReview.Spec.ResourceAttributes.Namespace
	return &authorizationapi.SubjectAccessReview{Status: authorizationapi.SubjectAccessReviewStatus{Allowed: f.allow}}, f.err
}
func newStorage(t *testing.T) (*REST, *StatusREST, *InternalREST, *etcdtesting.EtcdTestServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	server, etcdStorage := etcdtesting.NewUnsecuredEtcd3TestClientServer(t)
	etcdStorage.Codec = legacyscheme.Codecs.LegacyCodec(schema.GroupVersion{Group: "image.openshift.io", Version: "v1"})
	imagestreamRESTOptions := generic.RESTOptions{StorageConfig: etcdStorage, Decorator: generic.UndecoratedStorage, DeleteCollectionWorkers: 1, ResourcePrefix: "imagestreams"}
	registry := registryhostname.TestingRegistryHostnameRetriever(noDefaultRegistry, "", "")
	imageStorage, _, statusStorage, internalStorage, err := NewRESTWithLimitVerifier(imagestreamRESTOptions, registry, &fakeSubjectAccessReviewRegistry{}, &admfake.ImageStreamLimitVerifier{}, &fake.RegistryWhitelister{}, NewEmptyLayerIndex())
	if err != nil {
		t.Fatal(err)
	}
	return imageStorage, statusStorage, internalStorage, server
}
func validImageStream() *imageapi.ImageStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: name}}
}
func create(t *testing.T, storage *REST, obj *imageapi.ImageStream) *imageapi.ImageStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := apirequest.WithUser(apirequest.NewDefaultContext(), &fakeUser{})
	newObj, err := storage.Create(ctx, obj, rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	return newObj.(*imageapi.ImageStream)
}
func TestCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage, _, _, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	create(t, storage, validImageStream())
}
func TestList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage, _, _, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := genericregistrytest.New(t, storage.Store)
	test.TestList(validImageStream())
}
func TestGetImageStreamError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage, _, _, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	image, err := storage.Get(apirequest.NewDefaultContext(), "image1", &metav1.GetOptions{})
	if !errors.IsNotFound(err) {
		t.Errorf("Expected not-found error, got %v", err)
	}
	if image != nil {
		t.Errorf("Unexpected non-nil image stream: %#v", image)
	}
}
func TestGetImageStreamOK(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage, _, _, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	image := create(t, storage, validImageStream())
	obj, err := storage.Get(apirequest.NewDefaultContext(), name, &metav1.GetOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
	if obj == nil {
		t.Fatalf("Unexpected nil stream")
	}
	got := obj.(*imageapi.ImageStream)
	got.ResourceVersion = image.ResourceVersion
	if !kapihelper.Semantic.DeepEqual(image, got) {
		t.Errorf("Expected %#v, got %#v", image, got)
	}
}

type fakeUser struct{}

var _ user.Info = &fakeUser{}

func (u *fakeUser) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "user"
}
func (u *fakeUser) GetUID() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "uid"
}
func (u *fakeUser) GetGroups() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"group1"}
}
func (u *fakeUser) GetExtra() map[string][]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return map[string][]string{}
}
