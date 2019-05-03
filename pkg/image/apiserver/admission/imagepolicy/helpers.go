package imagepolicy

import (
	imagepolicyapiv1 "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
)

func RequestsResolution(imageResolutionType imagepolicyapiv1.ImageResolutionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch imageResolutionType {
	case imagepolicyapiv1.RequiredRewrite, imagepolicyapiv1.Required, imagepolicyapiv1.AttemptRewrite, imagepolicyapiv1.Attempt:
		return true
	}
	return false
}
func FailOnResolutionFailure(imageResolutionType imagepolicyapiv1.ImageResolutionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch imageResolutionType {
	case imagepolicyapiv1.RequiredRewrite, imagepolicyapiv1.Required:
		return true
	}
	return false
}
func RewriteImagePullSpec(imageResolutionType imagepolicyapiv1.ImageResolutionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch imageResolutionType {
	case imagepolicyapiv1.RequiredRewrite, imagepolicyapiv1.AttemptRewrite:
		return true
	}
	return false
}
