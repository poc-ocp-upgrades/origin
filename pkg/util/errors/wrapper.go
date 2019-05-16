package errors

import (
	"context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"strings"
)

func SyncStatusError(ctx context.Context, err error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		return nil
	}
	statusErr, isStatusErr := err.(apierrors.APIStatus)
	if !isStatusErr {
		return err
	}
	info, hasInfo := apirequest.RequestInfoFrom(ctx)
	if !hasInfo {
		return err
	}
	status := statusErr.Status()
	if status.Details == nil {
		return err
	}
	oldGR := (&schema.GroupResource{Group: status.Details.Group, Resource: status.Details.Kind}).String()
	newGR := (&schema.GroupResource{Group: info.APIGroup, Resource: info.Resource}).String()
	status.Message = strings.Replace(status.Message, oldGR, newGR, 1)
	status.Details.Group = info.APIGroup
	status.Details.Kind = info.Resource
	return &apierrors.StatusError{ErrStatus: status}
}
