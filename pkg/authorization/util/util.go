package util

import (
	"errors"
	authorizationv1 "k8s.io/api/authorization/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authentication/user"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

func AddUserToSAR(user user.Info, sar *authorizationv1.SubjectAccessReview) *authorizationv1.SubjectAccessReview {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sar.Spec.User = user.GetName()
	sar.Spec.Groups = make([]string, len(user.GetGroups()))
	copy(sar.Spec.Groups, user.GetGroups())
	sar.Spec.Extra = map[string]authorizationv1.ExtraValue{}
	for k, v := range user.GetExtra() {
		sar.Spec.Extra[k] = authorizationv1.ExtraValue(v)
	}
	return sar
}
func Authorize(sarClient authorizationclient.SubjectAccessReviewInterface, user user.Info, resourceAttributes *authorizationv1.ResourceAttributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sar := AddUserToSAR(user, &authorizationv1.SubjectAccessReview{Spec: authorizationv1.SubjectAccessReviewSpec{ResourceAttributes: resourceAttributes}})
	resp, err := sarClient.Create(sar)
	if err == nil && resp != nil && resp.Status.Allowed {
		return nil
	}
	if err == nil {
		err = errors.New(resp.Status.Reason)
	}
	return kerrors.NewForbidden(schema.GroupResource{Group: resourceAttributes.Group, Resource: resourceAttributes.Resource}, resourceAttributes.Name, err)
}
