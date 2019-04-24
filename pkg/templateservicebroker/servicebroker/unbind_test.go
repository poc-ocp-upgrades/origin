package servicebroker

import (
	"net/http"
	"reflect"
	"testing"
	templatev1 "github.com/openshift/api/template/v1"
	faketemplatev1 "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1/fake"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/kubernetes/fake"
	clienttesting "k8s.io/client-go/testing"
)

func TestUnbindConflict(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakekc := &fake.Clientset{}
	fakekc.AddReactor("create", "subjectaccessreviews", func(action clienttesting.Action) (bool, runtime.Object, error) {
		return true, &authorizationv1.SubjectAccessReview{Status: authorizationv1.SubjectAccessReviewStatus{Allowed: true}}, nil
	})
	faketemplateclient := &faketemplatev1.FakeTemplateV1{Fake: &clienttesting.Fake{}}
	faketemplateclient.AddReactor("get", "brokertemplateinstances", func(action clienttesting.Action) (bool, runtime.Object, error) {
		return true, &templatev1.BrokerTemplateInstance{Spec: templatev1.BrokerTemplateInstanceSpec{BindingIDs: []string{"bindingid"}}}, nil
	})
	var conflict int
	faketemplateclient.AddReactor("update", "brokertemplateinstances", func(action clienttesting.Action) (bool, runtime.Object, error) {
		if conflict > 0 {
			conflict--
			return true, nil, errors.NewConflict(templatev1.Resource("brokertemplateinstance"), "", nil)
		}
		return true, &templatev1.BrokerTemplateInstance{}, nil
	})
	b := &Broker{templateclient: faketemplateclient, kc: fakekc}
	conflict = 5
	resp := b.Unbind(&user.DefaultInfo{}, "", "bindingid")
	if !reflect.DeepEqual(resp, api.NewResponse(http.StatusUnprocessableEntity, &api.ConcurrencyError, nil)) {
		t.Errorf("got response %#v, expected 422/ConcurrencyError", *resp)
	}
	conflict = 4
	resp = b.Unbind(&user.DefaultInfo{}, "", "bindingid")
	if !reflect.DeepEqual(resp, api.NewResponse(http.StatusOK, &api.UnbindResponse{}, nil)) {
		t.Errorf("got response %#v, expected 200", *resp)
	}
	resp = b.Unbind(&user.DefaultInfo{}, "", "doesnotexist")
	if !reflect.DeepEqual(resp, api.NewResponse(http.StatusGone, &api.UnbindResponse{}, nil)) {
		t.Errorf("got response %#v, expected 410", *resp)
	}
}
