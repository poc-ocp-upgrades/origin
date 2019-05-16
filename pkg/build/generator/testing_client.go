package generator

import (
	"context"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TestingClient struct {
	GetBuildConfigFunc      func(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.BuildConfig, error)
	UpdateBuildConfigFunc   func(ctx context.Context, buildConfig *buildv1.BuildConfig) error
	GetBuildFunc            func(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.Build, error)
	CreateBuildFunc         func(ctx context.Context, build *buildv1.Build) error
	UpdateBuildFunc         func(ctx context.Context, build *buildv1.Build) error
	GetImageStreamFunc      func(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStream, error)
	GetImageStreamImageFunc func(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamImage, error)
	GetImageStreamTagFunc   func(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamTag, error)
}

func (c TestingClient) GetBuildConfig(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.BuildConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetBuildConfigFunc(ctx, name, options)
}
func (c TestingClient) UpdateBuildConfig(ctx context.Context, buildConfig *buildv1.BuildConfig) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.UpdateBuildConfigFunc(ctx, buildConfig)
}
func (c TestingClient) GetBuild(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.Build, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetBuildFunc(ctx, name, options)
}
func (c TestingClient) CreateBuild(ctx context.Context, build *buildv1.Build) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.CreateBuildFunc(ctx, build)
}
func (c TestingClient) UpdateBuild(ctx context.Context, build *buildv1.Build) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.UpdateBuildFunc(ctx, build)
}
func (c TestingClient) GetImageStream(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStream, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetImageStreamFunc(ctx, name, options)
}
func (c TestingClient) GetImageStreamImage(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamImage, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetImageStreamImageFunc(ctx, name, options)
}
func (c TestingClient) GetImageStreamTag(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamTag, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetImageStreamTagFunc(ctx, name, options)
}
