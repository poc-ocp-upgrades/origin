package tokenreview

import (
	"context"
	"errors"
	"fmt"
	goformat "fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/apis/authentication"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var badAuthenticatorAuds = apierrors.NewInternalError(errors.New("error validating audiences"))

type REST struct {
	tokenAuthenticator authenticator.Request
	apiAudiences       []string
}

func NewREST(tokenAuthenticator authenticator.Request, apiAudiences []string) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{tokenAuthenticator: tokenAuthenticator, apiAudiences: apiAudiences}
}
func (r *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &authentication.TokenReview{}
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokenReview, ok := obj.(*authentication.TokenReview)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("not a TokenReview: %#v", obj))
	}
	namespace := genericapirequest.NamespaceValue(ctx)
	if len(namespace) != 0 {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("namespace is not allowed on this type: %v", namespace))
	}
	if len(tokenReview.Spec.Token) == 0 {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("token is required for TokenReview in authentication"))
	}
	if r.tokenAuthenticator == nil {
		return tokenReview, nil
	}
	fakeReq := &http.Request{Header: http.Header{}}
	fakeReq.Header.Add("Authorization", "Bearer "+tokenReview.Spec.Token)
	auds := tokenReview.Spec.Audiences
	if len(auds) == 0 {
		auds = r.apiAudiences
	}
	if len(auds) > 0 {
		fakeReq = fakeReq.WithContext(authenticator.WithAudiences(fakeReq.Context(), auds))
	}
	resp, ok, err := r.tokenAuthenticator.AuthenticateRequest(fakeReq)
	tokenReview.Status.Authenticated = ok
	if err != nil {
		tokenReview.Status.Error = err.Error()
	}
	if len(auds) > 0 && resp != nil && len(authenticator.Audiences(auds).Intersect(resp.Audiences)) == 0 {
		klog.Errorf("error validating audience. want=%q got=%q", auds, resp.Audiences)
		return nil, badAuthenticatorAuds
	}
	if resp != nil && resp.User != nil {
		tokenReview.Status.User = authentication.UserInfo{Username: resp.User.GetName(), UID: resp.User.GetUID(), Groups: resp.User.GetGroups(), Extra: map[string]authentication.ExtraValue{}}
		for k, v := range resp.User.GetExtra() {
			tokenReview.Status.User.Extra[k] = authentication.ExtraValue(v)
		}
		tokenReview.Status.Audiences = resp.Audiences
	}
	return tokenReview, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
